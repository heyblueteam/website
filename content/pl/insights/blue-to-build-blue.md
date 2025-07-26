---
title: Jak używamy Blue do budowania Blue.
description: Dowiedz się, jak korzystamy z naszej własnej platformy do zarządzania projektami, aby zbudować naszą platformę do zarządzania projektami!
category: "CEO Blog"
date: 2024-08-07
---


Zaraz dostaniesz wgląd w to, jak Blue buduje Blue.

W Blue korzystamy z własnych produktów.

Oznacza to, że używamy Blue do *budowania* Blue.

Ten dziwnie brzmiący termin, często określany jako "dogfooding", jest często przypisywany Paulowi Maritzowi, menedżerowi w Microsoft w latach 80. XX wieku. Podobno wysłał e-mail z tematem *"Jemy nasze własne jedzenie dla psów"*, aby zachęcić pracowników Microsoftu do korzystania z produktów firmy.

Pomysł korzystania z własnych narzędzi do budowania narzędzi polega na tym, że prowadzi do pozytywnego cyklu informacji zwrotnej.

Pomysł korzystania z własnych narzędzi do budowania narzędzi prowadzi do pozytywnego cyklu informacji zwrotnej, tworząc liczne korzyści:

- **Pomaga nam szybko zidentyfikować problemy z użytecznością w rzeczywistym świecie.** Korzystając z Blue codziennie, napotykamy te same wyzwania, z którymi mogą się zmagać nasi użytkownicy, co pozwala nam proaktywnie je rozwiązywać.
- **Przyspiesza odkrywanie błędów.** Wewnętrzne użytkowanie często ujawnia błędy, zanim dotrą do naszych klientów, poprawiając ogólną jakość produktu.
- **Zwiększa naszą empatię wobec końcowych użytkowników.** Nasz zespół zdobywa bezpośrednie doświadczenie w mocnych i słabych stronach Blue, co pomaga nam podejmować bardziej zorientowane na użytkownika decyzje.
- **Napędza kulturę jakości w naszej organizacji.** Kiedy wszyscy korzystają z produktu, istnieje wspólny interes w jego doskonałości.
- **Sprzyja innowacjom.** Regularne użytkowanie często wywołuje pomysły na nowe funkcje lub ulepszenia, utrzymując Blue na czołowej pozycji.

[Już wcześniej rozmawialiśmy o tym, dlaczego nie mamy dedykowanego zespołu testowego](/insights/open-beta) i to jest kolejny powód.

Jeśli w naszym systemie są błędy, prawie zawsze je znajdujemy w naszym codziennym użytkowaniu platformy. A to także tworzy funkcję wymuszającą ich naprawę, ponieważ z pewnością uznamy je za bardzo irytujące, jako że prawdopodobnie jesteśmy jednym z głównych użytkowników Blue!

To podejście pokazuje nasze zaangażowanie w produkt. Polegając na Blue, pokazujemy naszym klientom, że naprawdę wierzymy w to, co budujemy. To nie tylko produkt, który sprzedajemy – to narzędzie, na którym polegamy każdego dnia.

## Główny proces

Mamy jeden projekt w Blue, odpowiednio nazwany "Produkt".

**Wszystko** związane z naszym rozwojem produktu jest śledzone tutaj. Opinie klientów, błędy, pomysły na funkcje, bieżąca praca itd. Pomysł posiadania jednego projektu, w którym śledzimy wszystko, polega na tym, że [promuje lepszą współpracę zespołową.](/insights/great-teamwork)

Każdy rekord to funkcja lub część funkcji. Tak przechodzimy od "czy nie byłoby fajnie, gdyby..." do "sprawdź tę niesamowitą nową funkcję!"

Projekt ma następujące listy:

- **Pomysły/Opinie**: To lista pomysłów zespołu lub opinii klientów na podstawie rozmów lub wymiany e-maili. Śmiało dodawaj tutaj swoje pomysły! Na tej liście nie zdecydowaliśmy jeszcze, że zbudujemy którąkolwiek z tych funkcji, ale regularnie przeglądamy ją w poszukiwaniu pomysłów, które chcemy zbadać dalej.
- **Backlog (długoterminowy)**: To tutaj trafiają funkcje z listy Pomysły/Opinie, jeśli zdecydujemy, że będą dobrym dodatkiem do Blue.
- **{Bieżący kwartał}**: Zwykle jest to zorganizowane jako "Qx YYYY" i pokazuje nasze priorytety kwartalne.
- **Błędy**: To lista znanych błędów zgłoszonych przez zespół lub klientów. Błędy dodane tutaj automatycznie otrzymają tag "Błąd".
- **Specyfikacje**: Te funkcje są obecnie specyfikowane. Nie każda funkcja wymaga specyfikacji lub projektu; zależy to od oczekiwanej wielkości funkcji i poziomu pewności, jaki mamy w odniesieniu do przypadków brzegowych i złożoności.
- **Backlog projektowy**: To backlog dla projektantów, za każdym razem, gdy skończą coś, co jest w toku, mogą wybrać dowolny element z tej listy.
- **W trakcie projektowania**: To obecne funkcje, które projektanci projektują.
- **Przegląd projektu**: To miejsce, w którym znajdują się funkcje, których projekty są obecnie przeglądane.
- **Backlog (krótkoterminowy)**: To lista funkcji, nad którymi prawdopodobnie zaczniemy pracować w ciągu najbliższych kilku tygodni. To tutaj odbywają się przypisania. CEO i szef inżynierii decydują, które funkcje są przypisane do którego inżyniera na podstawie wcześniejszych doświadczeń i obciążenia pracą. [Członkowie zespołu mogą następnie przenieść je do W trakcie](/insights/push-vs-pull-kanban), gdy zakończą swoją bieżącą pracę.
- **W trakcie**: To funkcje, które są obecnie rozwijane.
- **Przegląd kodu**: Gdy funkcja zakończy rozwój, przechodzi przegląd kodu. Następnie zostanie przeniesiona z powrotem do "W trakcie", jeśli potrzebne są poprawki, lub wdrożona do środowiska deweloperskiego.
- **Dev**: To wszystkie funkcje obecnie w środowisku deweloperskim. Inni członkowie zespołu i niektórzy klienci mogą je przeglądać.
- **Beta**: To wszystkie funkcje obecnie w [środowisku Beta](https://beta.app.blue.cc). Wielu klientów korzysta z tego jako swojej codziennej platformy Blue i również dostarcza opinie.
- **Produkcja**: Gdy funkcja osiąga produkcję, jest wtedy uważana za zakończoną.

Czasami, gdy rozwijamy funkcję, zdajemy sobie sprawę, że niektóre podfunkcje są trudniejsze do wdrożenia niż początkowo oczekiwano, i możemy zdecydować się ich nie realizować w początkowej wersji, którą wdrażamy do klientów. W takim przypadku możemy utworzyć nowy rekord z nazwą w formacie "{NazwaFunkcji} V2" i uwzględnić wszystkie podfunkcje jako elementy listy kontrolnej.

## Tagowanie

- **Mobilny**: Oznacza to, że funkcja jest specyficzna dla naszych aplikacji iOS, Android lub iPad.
- **{NazwaKlientaEnterprise}**: Funkcja jest specjalnie tworzona dla klienta korporacyjnego. Śledzenie jest ważne, ponieważ zazwyczaj istnieją dodatkowe umowy handlowe dla każdej funkcji.
- **Błąd**: Oznacza to, że jest to błąd, który wymaga naprawy.
- **Szybka ścieżka**: Oznacza to, że jest to zmiana na szybkiej ścieżce, która nie musi przechodzić przez pełny cykl wydania, jak opisano powyżej.
- **Główna**: To rozwój głównej funkcji. Zwykle jest zarezerwowane dla dużych prac infrastrukturalnych, dużych aktualizacji zależności i znaczących nowych modułów w Blue.
- **AI**: Ta funkcja zawiera komponent sztucznej inteligencji.
- **Bezpieczeństwo**: Oznacza to, że musi być przeglądany wpływ na bezpieczeństwo lub wymagana jest łatka.

Tag szybkiej ścieżki jest szczególnie interesujący. Jest zarezerwowany dla mniejszych, mniej skomplikowanych aktualizacji, które nie wymagają pełnego cyklu wydania i które chcemy dostarczyć klientom w ciągu 24-48 godzin.

Zmiany na szybkiej ścieżce to zazwyczaj drobne poprawki, które mogą znacznie poprawić doświadczenie użytkownika bez zmiany podstawowej funkcjonalności. Myśl o poprawianiu literówek w interfejsie użytkownika, dostosowywaniu odstępów przycisków lub dodawaniu nowych ikon dla lepszej wizualnej wskazówki. To są zmiany, które, mimo że małe, mogą mieć duże znaczenie dla tego, jak użytkownicy postrzegają i wchodzą w interakcję z naszym produktem. Są również irytujące, jeśli zajmują dużo czasu na wdrożenie!

Nasz proces szybkiej ścieżki jest prosty.

Zaczyna się od utworzenia nowej gałęzi z głównej, wdrożenia zmian, a następnie utworzenia żądań scalania dla każdej docelowej gałęzi - Dev, Beta i Produkcja. Generujemy link do podglądu do przeglądu, zapewniając, że nawet te małe zmiany spełniają nasze standardy jakości. Po zatwierdzeniu zmiany są scalane jednocześnie we wszystkich gałęziach, utrzymując nasze środowiska w synchronizacji.

## Pola niestandardowe

Nie mamy wielu pól niestandardowych w naszym projekcie Produkt.

- **Specyfikacje**: To łączy się z dokumentem Blue, który zawiera specyfikację dla danej funkcji. Nie zawsze jest to robione, ponieważ zależy to od złożoności funkcji.
- **MR**: To link do Żądania Scalania w [Gitlab](https://gitlab.com), gdzie hostujemy nasz kod.
- **Link do podglądu**: Dla funkcji, które głównie zmieniają front-end, możemy utworzyć unikalny URL, który ma te zmiany dla każdego zatwierdzenia, abyśmy mogli łatwo przeglądać zmiany.
- **Lider**: To pole informuje nas, który starszy inżynier zajmuje się przeglądem kodu. Zapewnia, że każda funkcja otrzymuje fachową uwagę, na jaką zasługuje, i zawsze jest jasna osoba kontaktowa w przypadku pytań lub wątpliwości.

## Listy kontrolne

Podczas naszych cotygodniowych demonstracji umieszczamy omówione opinie w liście kontrolnej o nazwie "opinie", a także będzie inna lista kontrolna, która zawiera główny [WBS (Zakres Rozbicia Pracy)](/insights/simple-work-breakdown-structure) funkcji, abyśmy mogli łatwo określić, co jest zrobione, a co jeszcze do zrobienia.

## Podsumowanie

I to wszystko!

Uważamy, że czasami ludzie są zaskoczeni, jak prosty jest nasz proces, ale wierzymy, że proste procesy są często znacznie lepsze niż zbyt skomplikowane procesy, których nie można łatwo zrozumieć.

Ta prostota jest zamierzona. Pozwala nam pozostać zwinnymi, szybko reagować na potrzeby klientów i utrzymywać cały zespół w zgodzie.

Korzystając z Blue do budowania Blue, nie tylko rozwijamy produkt – my go przeżywamy.

Więc następnym razem, gdy korzystasz z Blue, pamiętaj: nie korzystasz tylko z produktu, który zbudowaliśmy. Korzystasz z produktu, na którym osobiście polegamy każdego dnia.

I to robi całą różnicę.