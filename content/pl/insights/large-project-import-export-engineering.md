---
title:  Skalowanie importów i eksportów CSV do 250 000+ rekordów
description: Odkryj, jak Blue skalowało importy i eksporty CSV 10x, wykorzystując Rust, skalowalną architekturę i strategiczne wybory technologiczne w B2B SaaS.
category: "Engineering"
date: 2024-07-18
---


W Blue [nieustannie przesuwamy granice](/platform/roadmap) tego, co możliwe w oprogramowaniu do zarządzania projektami. Na przestrzeni lat [wprowadziliśmy setki funkcji](/platform/changelog).

Nasze najnowsze osiągnięcie inżynieryjne? 

Całkowita przebudowa naszego systemu [importu CSV](https://documentation.blue.cc/integrations/csv-import) i [eksportu](https://documentation.blue.cc/integrations/csv-export), która dramatycznie poprawiła wydajność i skalowalność. 

Ten post zabierze cię za kulisy tego, jak podjęliśmy to wyzwanie, jakich technologii użyliśmy i jakie imponujące wyniki osiągnęliśmy.

Najciekawszą rzeczą jest to, że musieliśmy wyjść poza nasz typowy [stos technologiczny](https://sop.blue.cc/product/technology-stack), aby osiągnąć pożądane rezultaty. To decyzja, którą należy podjąć z rozwagą, ponieważ długoterminowe reperkusje mogą być poważne w kontekście długu technologicznego i długoterminowych kosztów utrzymania. 

<video autoplay loop muted playsinline>
  <source src="/videos/import-export-video.mp4" type="video/mp4">
</video>

## Skalowanie dla potrzeb przedsiębiorstw

Nasza podróż rozpoczęła się od prośby od klienta korporacyjnego z branży wydarzeń. Ten klient korzysta z Blue jako centralnego hubu do zarządzania ogromnymi listami wydarzeń, miejsc i prelegentów, integrując go bezproblemowo z ich stroną internetową. 

Dla nich Blue to nie tylko narzędzie — to jedyne źródło prawdy dla całej ich operacji.

Chociaż zawsze cieszy nas, gdy słyszymy, że klienci korzystają z nas w tak krytycznych dla misji potrzebach, to również spoczywa na nas duża odpowiedzialność, aby zapewnić szybki i niezawodny system.

Gdy ten klient skalował swoje operacje, napotkał znaczną przeszkodę: **importowanie i eksportowanie dużych plików CSV zawierających od 100 000 do 200 000+ rekordów.**

To przekraczało możliwości naszego systemu w tamtym czasie. W rzeczywistości nasz poprzedni system importu/eksportu już miał problemy z importami i eksportami zawierającymi więcej niż 10 000 do 20 000 rekordów! Tak więc 200 000+ rekordów było poza zasięgiem. 

Użytkownicy doświadczali frustrująco długich czasów oczekiwania, a w niektórych przypadkach importy lub eksporty *nie kończyły się w ogóle.* To znacząco wpłynęło na ich operacje, ponieważ polegali na codziennych importach i eksportach, aby zarządzać niektórymi aspektami swojej działalności. 

> Multi-tenancy to architektura, w której jedna instancja oprogramowania obsługuje wielu klientów (najemców). Choć jest to wydajne, wymaga starannego zarządzania zasobami, aby działania jednego najemcy nie wpływały negatywnie na innych.

A ta ograniczenie nie dotyczyło tylko tego konkretnego klienta. 

Z powodu naszej architektury wielonajemnej — gdzie wielu klientów dzieli tę samą infrastrukturę — pojedynczy zasobożerny import lub eksport mógł potencjalnie spowolnić operacje dla innych użytkowników, co w praktyce często miało miejsce. 

Jak zwykle, przeprowadziliśmy analizę budowy vs zakupu, aby zrozumieć, czy powinniśmy poświęcić czas na ulepszenie naszego własnego systemu, czy kupić system od kogoś innego. Rozważaliśmy różne możliwości.

Dostawca, który się wyróżniał, to dostawca SaaS o nazwie [Flatfile](https://flatfile.com/). Ich system i możliwości wyglądały na dokładnie to, czego potrzebowaliśmy. 

Jednak po przeglądzie ich [cen](https://flatfile.com/pricing/) zdecydowaliśmy, że to byłoby niezwykle kosztowne rozwiązanie dla aplikacji naszej skali — *2 USD/plik szybko się sumują!* — i lepiej byłoby rozszerzyć nasz wbudowany silnik importu/eksportu CSV. 

Aby stawić czoła temu wyzwaniu, podjęliśmy odważną decyzję: wprowadzić Rust do naszego głównego stosu technologicznego opartego na Javascript. Ten język programowania systemowego, znany ze swojej wydajności i bezpieczeństwa, był idealnym narzędziem do naszych krytycznych pod względem wydajności potrzeb związanych z analizą CSV i mapowaniem danych.

Oto jak podeszliśmy do rozwiązania.

### Wprowadzenie usług w tle

Fundamentem naszego rozwiązania było wprowadzenie usług w tle do obsługi zadań wymagających dużych zasobów. To podejście pozwoliło nam odciążyć ciężkie przetwarzanie z naszego głównego serwera, znacznie poprawiając ogólną wydajność systemu. 
Nasza architektura usług w tle została zaprojektowana z myślą o skalowalności. Jak wszystkie komponenty naszej infrastruktury, te usługi automatycznie skalują się w zależności od popytu. 

Oznacza to, że w okresach szczytowych, gdy jednocześnie przetwarzane są wiele dużych importów lub eksportów, system automatycznie przydziela więcej zasobów, aby poradzić sobie z zwiększonym obciążeniem. Z kolei w spokojniejszych okresach skaluje się w dół, aby zoptymalizować wykorzystanie zasobów.

Ta skalowalna architektura usług w tle przyniosła korzyści Blue nie tylko w przypadku importów i eksportów CSV. Z biegiem czasu przenieśliśmy znaczną liczbę funkcji do usług w tle, aby odciążyć nasze główne serwery:

- **[Obliczenia Formuł](https://documentation.blue.cc/custom-fields/formula)**: Odciąża złożone operacje matematyczne, aby zapewnić szybkie aktualizacje pól pochodnych bez wpływu na wydajność głównego serwera.
- **[Dashboardy/Wykresy](/platform/features/dashboards)**: Przetwarza duże zbiory danych w tle, aby generować aktualne wizualizacje bez spowalniania interfejsu użytkownika.
- **[Indeks Wyszukiwania](https://documentation.blue.cc/projects/search)**: Ciągle aktualizuje indeks wyszukiwania w tle, zapewniając szybkie i dokładne wyniki wyszukiwania bez wpływu na wydajność systemu.
- **[Kopiowanie Projektów](https://documentation.blue.cc/projects/copying-projects)**: Obsługuje replikację dużych, złożonych projektów w tle, pozwalając użytkownikom kontynuować pracę podczas tworzenia kopii.
- **[Automatyzacje Zarządzania Projektami](/platform/features/automations)**: Wykonuje zdefiniowane przez użytkownika zautomatyzowane przepływy pracy w tle, zapewniając terminowe działania bez blokowania innych operacji.
- **[Powtarzające się Rekordy](https://documentation.blue.cc/records/repeat)**: Generuje powtarzające się zadania lub wydarzenia w tle, utrzymując dokładność harmonogramu bez obciążania głównej aplikacji.
- **[Pola Niestandardowe Czasu Trwania](https://documentation.blue.cc/custom-fields/duration)**: Ciągle oblicza i aktualizuje różnicę czasu między dwoma wydarzeniami w Blue, zapewniając dane o czasie trwania w czasie rzeczywistym bez wpływu na responsywność systemu.

## Nowy moduł Rust do analizy danych

Serce naszego rozwiązania do przetwarzania CSV stanowi niestandardowy moduł Rust. Choć oznaczało to nasze pierwsze kroki poza naszym podstawowym stosem technologicznym opartym na Javascript, decyzja o użyciu Rust była podyktowana jego wyjątkową wydajnością w operacjach współbieżnych i zadaniach przetwarzania plików.

Mocne strony Rust idealnie odpowiadają wymaganiom analizy CSV i mapowania danych. Jego zerowe koszty abstrakcji pozwalają na programowanie na wysokim poziomie bez poświęcania wydajności, podczas gdy jego model własności zapewnia bezpieczeństwo pamięci bez potrzeby zbierania śmieci. Te cechy sprawiają, że Rust jest szczególnie skuteczny w efektywnym i bezpiecznym przetwarzaniu dużych zbiorów danych.

Do analizy CSV wykorzystaliśmy crate csv Rust, który oferuje wysokowydajne odczytywanie i zapisywanie danych CSV. Połączyliśmy to z niestandardową logiką mapowania danych, aby zapewnić bezproblemową integrację z strukturami danych Blue.

Krzywa uczenia się Rust była stroma, ale wykonalna. Nasz zespół poświęcił około dwóch tygodni na intensywne uczenie się tego języka.

Poprawa była imponująca:

![](/insights/import-export.png)

Nasz nowy system może przetwarzać tę samą liczbę rekordów, którą nasz stary system mógł przetworzyć w 15 minut, w około 30 sekund. 

## Interakcja z serwerem WWW i bazą danych

Dla komponentu serwera WWW naszej implementacji Rust wybraliśmy Rocket jako nasz framework. Rocket wyróżniał się połączeniem wydajności i przyjaznych dla dewelopera funkcji. Jego statyczne typowanie i sprawdzanie w czasie kompilacji dobrze współgrają z zasadami bezpieczeństwa Rust, pomagając nam wcześnie wychwytywać potencjalne problemy w procesie rozwoju.
W kwestii bazy danych zdecydowaliśmy się na SQLx. Ta asynchroniczna biblioteka SQL dla Rust oferuje kilka zalet, które uczyniły ją idealną dla naszych potrzeb:

- Bezpieczny typowo SQL: SQLx pozwala nam pisać surowe zapytania SQL z kontrolowanymi w czasie kompilacji zapytaniami, zapewniając bezpieczeństwo typów bez poświęcania wydajności.
- Wsparcie asynchroniczne: To dobrze współgra z Rocket i naszą potrzebą efektywnych, nieblokujących operacji na bazie danych.
- Niezależność od bazy danych: Chociaż głównie korzystamy z [AWS Aurora](https://aws.amazon.com/rds/aurora/), która jest zgodna z MySQL, wsparcie SQLx dla wielu baz danych daje nam elastyczność na przyszłość, w przypadku gdybyśmy zdecydowali się na zmianę. 

## Optymalizacja wsadowa

Nasza podróż do optymalnej konfiguracji wsadowej była pełna rygorystycznych testów i starannej analizy. Przeprowadziliśmy obszerne benchmarki z różnymi kombinacjami równoległych transakcji i rozmiarów bloków, mierząc nie tylko surową prędkość, ale także wykorzystanie zasobów i stabilność systemu.

Proces obejmował tworzenie zestawów testowych o różnej wielkości i złożoności, symulując rzeczywiste wzorce użytkowania. Następnie przetwarzaliśmy te zestawy przez nasz system, dostosowując liczbę równoległych transakcji i rozmiar bloku dla każdego uruchomienia.

Po analizie wyników stwierdziliśmy, że przetwarzanie 5 równoległych transakcji z rozmiarem bloku 500 rekordów zapewnia najlepszą równowagę między prędkością a wykorzystaniem zasobów. Ta konfiguracja pozwala nam utrzymać wysoką przepustowość bez przytłaczania naszej bazy danych lub zużywania nadmiernej ilości pamięci.

Interesujące jest to, że zwiększenie równoległości powyżej 5 transakcji nie przyniosło znaczących zysków wydajności i czasami prowadziło do zwiększonej kontencji bazy danych. Podobnie, większe rozmiary bloków poprawiały surową prędkość, ale kosztem wyższego zużycia pamięci i dłuższych czasów odpowiedzi dla małych i średnich importów/eksportów.

## Eksporty CSV za pomocą linków e-mailowych

Ostatni element naszego rozwiązania dotyczy wyzwania dostarczania dużych plików eksportowanych do użytkowników. Zamiast zapewniać bezpośrednie pobieranie z naszej aplikacji internetowej, co mogłoby prowadzić do problemów z czasem oczekiwania i zwiększonego obciążenia serwera, wdrożyliśmy system linków do pobrania wysyłanych e-mailem.

Gdy użytkownik inicjuje duży eksport, nasz system przetwarza żądanie w tle. Po zakończeniu, zamiast trzymać otwarte połączenie lub przechowywać plik na naszych serwerach internetowych, przesyłamy plik do bezpiecznej, tymczasowej lokalizacji przechowywania. Następnie generujemy unikalny, bezpieczny link do pobrania i wysyłamy go do użytkownika e-mailem.

Te linki do pobrania są ważne przez 2 godziny, co stanowi równowagę między wygodą użytkownika a bezpieczeństwem informacji. Ten czas daje użytkownikom wystarczającą możliwość pobrania danych, zapewniając jednocześnie, że wrażliwe informacje nie pozostają dostępne na zawsze.

Bezpieczeństwo tych linków do pobrania było priorytetem w naszym projekcie. Każdy link jest:

- Unikalny i losowo generowany, co sprawia, że praktycznie niemożliwe jest jego odgadnięcie
- Ważny tylko przez 2 godziny
- Szyfrowany w tranzycie, zapewniając bezpieczeństwo danych podczas pobierania

To podejście oferuje kilka korzyści:

- Zmniejsza obciążenie naszych serwerów internetowych, ponieważ nie muszą one bezpośrednio obsługiwać dużych pobrań plików
- Poprawia doświadczenie użytkownika, zwłaszcza dla użytkowników z wolniejszymi połączeniami internetowymi, którzy mogą napotykać problemy z czasem oczekiwania w przeglądarkach przy bezpośrednich pobraniach
- Zapewnia bardziej niezawodne rozwiązanie dla bardzo dużych eksportów, które mogą przekraczać typowe limity czasów oczekiwania w sieci

Opinie użytkowników na temat tej funkcji były niezwykle pozytywne, a wielu doceniło elastyczność, jaką oferuje w zarządzaniu dużymi eksportami danych.

## Eksportowanie przefiltrowanych danych

Inną oczywistą poprawą było umożliwienie użytkownikom eksportowania tylko danych, które były już przefiltrowane w ich widoku projektu. Oznacza to, że jeśli istnieje aktywna etykieta "priorytet", to tylko rekordy, które mają tę etykietę, trafią do eksportu CSV. Oznacza to mniej czasu na manipulowanie danymi w Excelu, aby odfiltrować rzeczy, które nie są ważne, a także pomaga nam zmniejszyć liczbę wierszy do przetworzenia.

## Patrząc w przyszłość

Chociaż nie mamy natychmiastowych planów na rozszerzenie naszego użycia Rust, ten projekt pokazał nam potencjał tej technologii dla operacji krytycznych pod względem wydajności. To ekscytująca opcja, którą teraz mamy w naszym zestawie narzędzi do przyszłych potrzeb optymalizacyjnych. Ta przebudowa importu i eksportu CSV idealnie wpisuje się w zobowiązanie Blue do skalowalności. 

Jesteśmy zdeterminowani, aby zapewnić platformę, która rośnie wraz z naszymi klientami, obsługując ich rosnące potrzeby danych bez kompromisów w wydajności.

Decyzja o wprowadzeniu Rust do naszego stosu technologicznego nie została podjęta lekko. Postawiła ważne pytanie, przed którym stają wiele zespołów inżynieryjnych: Kiedy jest odpowiedni moment, aby wyjść poza swój podstawowy stos technologiczny, a kiedy powinieneś pozostać przy znanych narzędziach?

Nie ma jednego rozwiązania dla wszystkich, ale w Blue opracowaliśmy ramy do podejmowania tych kluczowych decyzji:

- **Podejście skoncentrowane na problemie:** Zawsze zaczynamy od wyraźnego zdefiniowania problemu, który próbujemy rozwiązać. W tym przypadku musieliśmy dramatycznie poprawić wydajność importów i eksportów CSV dla dużych zbiorów danych.
- **Wykończenie istniejących rozwiązań:** Zanim spojrzymy poza nasz podstawowy stos, dokładnie badamy, co można osiągnąć za pomocą naszych istniejących technologii. Często wiąże się to z profilowaniem, optymalizacją i przemyśleniem naszego podejścia w ramach znanych ograniczeń.
- **Kwotowanie potencjalnych zysków:** Jeśli rozważamy nową technologię, musimy być w stanie jasno określić i, idealnie, kwantyfikować korzyści. Dla naszego projektu CSV przewidywaliśmy poprawę wydajności o rząd wielkości.
- **Ocena kosztów:** Wprowadzenie nowej technologii to nie tylko kwestia bieżącego projektu. Rozważamy długoterminowe koszty:
  - Krzywa uczenia się dla zespołu
  - Ciągłe utrzymanie i wsparcie
  - Potencjalne komplikacje w wdrożeniu i operacjach
  - Wpływ na zatrudnianie i skład zespołu
- **Ograniczenie i integracja:** Jeśli wprowadzimy nową technologię, dążymy do ograniczenia jej do konkretnej, dobrze zdefiniowanej części naszego systemu. Upewniamy się również, że mamy jasny plan, jak będzie integrować się z naszym istniejącym stosem.
- **Przyszłościowe planowanie:** Rozważamy, czy ten wybór technologii otwiera przyszłe możliwości, czy może zamknąć nas w kącie.

Jednym z głównych ryzyk częstego przyjmowania nowych technologii jest zakończenie w tym, co nazywamy *"zoo technologii"* - fragmentowanym ekosystemem, w którym różne części aplikacji są napisane w różnych językach lub frameworkach, wymagając szerokiego zakresu specjalistycznych umiejętności do utrzymania.

## Podsumowanie

Ten projekt jest przykładem podejścia Blue do inżynierii: *nie boimy się wyjść poza naszą strefę komfortu i przyjąć nowe technologie, gdy oznacza to dostarczenie znacznie lepszego doświadczenia dla naszych użytkowników.* 

Przekształcając nasz proces importu i eksportu CSV, nie tylko rozwiązaliśmy pilną potrzebę jednego klienta korporacyjnego, ale także poprawiliśmy doświadczenie wszystkich naszych użytkowników zajmujących się dużymi zbiorami danych.

Gdy nadal przesuwamy granice tego, co możliwe w [oprogramowaniu do zarządzania projektami](/solutions/use-case/project-management), cieszymy się na kolejne wyzwania jak to. 

Bądź na bieżąco z kolejnymi [dogłębnymi analizami inżynierii, która napędza Blue!](/insights/engineering-blog)