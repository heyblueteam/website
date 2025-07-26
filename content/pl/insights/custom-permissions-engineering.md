---
title: Tworzenie niestandardowego silnika uprawnień Blue
description: Zajrzyj za kulisy z zespołem inżynieryjnym Blue, gdy wyjaśniają, jak zbudować funkcję automatycznej kategoryzacji i tagowania opartą na AI.
category: "Engineering"
date: 2024-07-25
---


Skuteczne zarządzanie projektami i procesami jest kluczowe dla organizacji każdej wielkości.

W Blue, [postawiliśmy sobie za cel](/about) zorganizowanie pracy na świecie, budując najlepszą platformę do zarządzania projektami na planecie — prostą, potężną, elastyczną i przystępną dla wszystkich.

Oznacza to, że nasza platforma musi dostosować się do unikalnych potrzeb każdego zespołu. Dziś z radością odsłaniamy jedną z naszych najpotężniejszych funkcji: Niestandardowe Uprawnienia.

Narzędzia do zarządzania projektami są kręgosłupem nowoczesnych przepływów pracy, przechowując wrażliwe dane, kluczowe komunikacje i strategiczne plany. W związku z tym możliwość precyzyjnego kontrolowania dostępu do tych informacji nie jest tylko luksusem — to konieczność.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>

Niestandardowe uprawnienia odgrywają kluczową rolę w platformach B2B SaaS, szczególnie w narzędziach do zarządzania projektami, gdzie równowaga między współpracą a bezpieczeństwem może przesądzić o sukcesie projektu.

Ale tutaj Blue przyjmuje inne podejście: **wierzymy, że funkcje klasy enterprise nie powinny być zarezerwowane tylko dla budżetów wielkich przedsiębiorstw.**

W erze, w której AI umożliwia małym zespołom działanie na niespotykaną dotąd skalę, dlaczego solidne bezpieczeństwo i personalizacja miałyby być poza zasięgiem?

W tym spojrzeniu za kulisy zbadamy, jak opracowaliśmy naszą funkcję Niestandardowych Uprawnień, kwestionując status quo poziomów cenowych SaaS i wprowadzając potężne, elastyczne opcje bezpieczeństwa dla firm każdej wielkości.

Niezależnie od tego, czy jesteś startupem z wielkimi marzeniami, czy ugruntowanym graczem, który chce zoptymalizować swoje procesy, niestandardowe uprawnienia mogą umożliwić nowe przypadki użycia, o których nawet nie wiedziałeś, że są możliwe.

## Zrozumienie niestandardowych uprawnień użytkowników

Zanim zanurzymy się w naszą podróż tworzenia niestandardowych uprawnień dla Blue, poświęćmy chwilę na zrozumienie, czym są niestandardowe uprawnienia użytkowników i dlaczego są tak kluczowe w oprogramowaniu do zarządzania projektami.

Niestandardowe uprawnienia użytkowników odnoszą się do możliwości dostosowania praw dostępu dla poszczególnych użytkowników lub grup w systemie oprogramowania. Zamiast polegać na zdefiniowanych rolach z ustalonymi zestawami uprawnień, niestandardowe uprawnienia pozwalają administratorom tworzyć wysoce specyficzne profile dostępu, które idealnie odpowiadają strukturze organizacji i potrzebom przepływu pracy.

W kontekście oprogramowania do zarządzania projektami, takiego jak Blue, niestandardowe uprawnienia obejmują:

* **Granularna kontrola dostępu**: Określenie, kto może przeglądać, edytować lub usuwać konkretne typy danych projektowych.
* **Ograniczenia oparte na funkcjach**: Włączanie lub wyłączanie określonych funkcji dla konkretnych użytkowników lub zespołów.
* **Poziomy wrażliwości danych**: Ustalanie różnych poziomów dostępu do wrażliwych informacji w ramach projektów.
* **Uprawnienia specyficzne dla przepływu pracy**: Dostosowanie możliwości użytkowników do konkretnych etapów lub aspektów przepływu pracy projektu.

Znaczenie niestandardowych uprawnień w zarządzaniu projektami jest nie do przecenienia:

* **Zwiększone bezpieczeństwo**: Oferując użytkownikom tylko niezbędny dostęp, zmniejszasz ryzyko naruszenia danych lub nieautoryzowanych zmian.
* **Poprawa zgodności**: Niestandardowe uprawnienia pomagają organizacjom spełniać specyficzne wymagania regulacyjne branży poprzez kontrolowanie dostępu do danych.
* **Usprawniona współpraca**: Zespoły mogą pracować efektywniej, gdy każdy członek ma odpowiedni poziom dostępu do wykonywania swojej roli bez zbędnych ograniczeń lub przytłaczających przywilejów.
* **Elastyczność dla złożonych organizacji**: W miarę jak firmy rosną i ewoluują, niestandardowe uprawnienia pozwalają oprogramowaniu dostosować się do zmieniających się struktur organizacyjnych i procesów.

## Droga do TAK

[Już wcześniej pisaliśmy](/insights/value-proposition-blue), że każda funkcja w Blue musi być **twardym** TAK, zanim zdecydujemy się ją zbudować. Nie mamy luksusu setek inżynierów, aby marnować czas i pieniądze na budowanie rzeczy, których klienci nie potrzebują.

I tak, droga do wdrożenia niestandardowych uprawnień w Blue nie była prostą linią. Jak wiele potężnych funkcji, zaczęła się od wyraźnej potrzeby naszych użytkowników i ewoluowała przez staranne rozważania i planowanie.

Przez lata nasi klienci domagali się bardziej granularnej kontroli nad uprawnieniami użytkowników. W miarę jak organizacje każdej wielkości zaczęły obsługiwać coraz bardziej złożone i wrażliwe projekty, ograniczenia naszego standardowego systemu kontroli dostępu opartego na rolach stały się oczywiste.

Małe startupy pracujące z zewnętrznymi klientami, średnie firmy z złożonymi procesami zatwierdzania oraz duże przedsiębiorstwa z rygorystycznymi wymaganiami zgodności wszystkie wyrażały tę samą potrzebę:

Więcej elastyczności w zarządzaniu dostępem użytkowników.

Pomimo wyraźnego zapotrzebowania, początkowo wahałyśmy się przed przystąpieniem do opracowywania niestandardowych uprawnień.

Dlaczego?

Rozumiałyśmy złożoność, która się z tym wiązała!

Niestandardowe uprawnienia dotykają każdej części systemu zarządzania projektami, od interfejsu użytkownika po strukturę bazy danych. Wiedziałyśmy, że wdrożenie tej funkcji wymagałoby znaczących zmian w naszej architekturze podstawowej oraz starannego rozważenia implikacji dotyczących wydajności.

Gdy analizowałyśmy rynek, zauważyłyśmy, że bardzo niewielu naszych konkurentów próbowało wdrożyć potężny silnik niestandardowych uprawnień, jakiego domagali się nasi klienci. Ci, którzy to zrobili, często zarezerwowali go dla swoich najwyższych planów enterprise.

Stało się jasne, dlaczego: wysiłek rozwojowy jest znaczny, a stawka jest wysoka.

Nieprawidłowe wdrożenie niestandardowych uprawnień mogłoby wprowadzić krytyczne błędy lub luki w zabezpieczeniach, potencjalnie zagrażając całemu systemowi. Ta realizacja podkreśliła wielkość wyzwania, które rozważałyśmy.

### Kwestionowanie status quo

Jednak w miarę jak kontynuowałyśmy rozwój i ewolucję, doszłyśmy do kluczowej realizacji:

**Tradycyjny model SaaS, który rezerwuje potężne funkcje dla klientów enterprise, przestał mieć sens w dzisiejszym krajobrazie biznesowym.**

W 2024 roku, dzięki mocy AI i zaawansowanym narzędziom, małe zespoły mogą działać na skali i złożoności, które rywalizują z dużo większymi organizacjami. Startup może obsługiwać wrażliwe dane klientów w wielu krajach. Mała agencja marketingowa może zarządzać dziesiątkami projektów klientów o różnych wymaganiach dotyczących poufności. Te firmy potrzebują tego samego poziomu bezpieczeństwa i personalizacji, co *jakiekolwiek* duże przedsiębiorstwo.

Zadaliśmy sobie pytanie: Dlaczego wielkość siły roboczej lub budżetu firmy ma decydować o jej zdolności do zabezpieczenia danych i efektywności procesów?

### Funkcje klasy enterprise dla wszystkich

Ta realizacja doprowadziła nas do podstawowej filozofii, która teraz napędza większość naszego rozwoju w Blue: Funkcje klasy enterprise powinny być dostępne dla firm każdej wielkości.

Wierzymy, że:

- **Bezpieczeństwo nie powinno być luksusem.** Każda firma, niezależnie od wielkości, zasługuje na narzędzia do ochrony swoich danych i procesów.
- **Elastyczność napędza innowacje.** Dając wszystkim naszym użytkownikom potężne narzędzia, umożliwiamy im tworzenie przepływów pracy i systemów, które posuwają ich branże naprzód.
- **Wzrost nie powinien wymagać zmian w platformie.** W miarę jak nasi klienci rosną, ich narzędzia powinny płynnie rosnąć razem z nimi.

Z tym nastawieniem postanowiłyśmy zmierzyć się z wyzwaniem niestandardowych uprawnień, zobowiązując się do udostępnienia go wszystkim naszym użytkownikom, a nie tylko tym na wyższych planach.

Ta decyzja wyznaczyła nas na ścieżkę starannego projektowania, iteracyjnego rozwoju i ciągłej informacji zwrotnej od użytkowników, co ostatecznie doprowadziło do funkcji niestandardowych uprawnień, z której jesteśmy dumne dzisiaj.

W następnej sekcji przyjrzymy się, jak podeszłyśmy do procesu projektowania i rozwoju, aby ożywić tę złożoną funkcję.

### Projektowanie i rozwój

Gdy zdecydowałyśmy się zająć niestandardowymi uprawnieniami, szybko zdałyśmy sobie sprawę, że stajemy przed ogromnym zadaniem.

Na pierwszy rzut oka "niestandardowe uprawnienia" mogą brzmieć prosto, ale to zwodniczo złożona funkcja, która dotyka każdego aspektu naszego systemu.

Wyzwanie było ogromne: musiałyśmy wdrożyć kaskadowe uprawnienia, umożliwić edycje w locie, dokonać znaczących zmian w schemacie bazy danych i zapewnić płynne funkcjonowanie w całym naszym ekosystemie – aplikacjach webowych, Mac, Windows, iOS i Android, a także w naszym API i webhookach.

Złożoność była wystarczająca, aby nawet najbardziej doświadczonych programistów zatrzymać na chwilę.

Nasze podejście koncentrowało się na dwóch kluczowych zasadach:

1. Rozbicie funkcji na zarządzalne wersje
2. Przyjęcie stopniowego wdrażania.

Stając przed złożonością pełnoskalowych niestandardowych uprawnień, zadałyśmy sobie kluczowe pytanie:

> Jaka byłaby najprostsza możliwa pierwsza wersja tej funkcji?

To podejście jest zgodne z zasadą agile dostarczania Minimalnego Produktu Wartościowego (MVP) i iterowania na podstawie opinii.

Nasza odpowiedź była zaskakująco prosta:

1. Wprowadzenie przełącznika do ukrywania zakładki aktywności projektu
2. Dodanie kolejnego przełącznika do ukrywania zakładki formularzy

**To wszystko.**

Bez dzwonków i świecidełek, bez skomplikowanych macierzy uprawnień — tylko dwa proste przełączniki włącz/wyłącz.

Choć może to wydawać się nieco rozczarowujące na pierwszy rzut oka, to podejście oferowało kilka istotnych zalet:

* **Szybka implementacja**: Te proste przełączniki mogły być szybko opracowane i przetestowane, co pozwoliło nam szybko wprowadzić podstawową wersję niestandardowych uprawnień do rąk użytkowników.
* **Wyraźna wartość dla użytkowników**: Nawet z tymi dwoma opcjami dostarczaliśmy namacalną wartość. Niektóre zespoły mogą chcieć ukryć strumień aktywności przed klientami, podczas gdy inne mogą potrzebować ograniczyć dostęp do formularzy dla określonych grup użytkowników.
* **Podstawa do rozwoju**: Ten prosty początek położył fundamenty dla bardziej złożonych uprawnień. Pozwoliło nam to ustawić podstawową infrastrukturę dla niestandardowych uprawnień, nie wpadając w złożoność od samego początku.
* **Informacje zwrotne od użytkowników**: Wydając tę prostą wersję, mogłyśmy zbierać rzeczywiste opinie na temat tego, jak użytkownicy wchodzą w interakcje z niestandardowymi uprawnieniami, co informowało nasz przyszły rozwój.
* **Nauka techniczna**: Ta początkowa implementacja dała naszemu zespołowi deweloperskiemu praktyczne doświadczenie w modyfikowaniu uprawnień w całej naszej platformie, przygotowując nas do bardziej złożonych iteracji.

I wiesz, to naprawdę dość pokorne mieć ogromną wizję czegoś, a potem dostarczyć coś, co jest tak małym procentem tej wizji.

Po wdrożeniu tych pierwszych dwóch przełączników postanowiłyśmy zająć się czymś bardziej zaawansowanym. Skupiłyśmy się na dwóch nowych niestandardowych rolach użytkowników.

Pierwsza to możliwość ograniczenia użytkowników do przeglądania tylko tych rekordów, które zostały im przypisane. To bardzo przydatne, jeśli masz klienta w projekcie i chcesz, aby widział tylko rekordy, które są mu przypisane, a nie wszystko, nad czym pracujesz dla niego.

Druga to opcja dla administratorów projektu, aby zablokować grupy użytkowników przed możliwością zapraszania innych użytkowników. To dobre, jeśli masz wrażliwy projekt, który chcesz, aby pozostał na zasadzie "potrzebujesz wiedzieć".

Gdy to wdrożyłyśmy, zyskałyśmy więcej pewności, a w naszej trzeciej wersji zająłyśmy się uprawnieniami na poziomie kolumn, co oznacza możliwość decydowania, które niestandardowe pola konkretna grupa użytkowników może przeglądać lub edytować.

To jest niezwykle potężne. Wyobraź sobie, że masz projekt CRM, a w nim dane, które są nie tylko związane z kwotami, które klient zapłaci, ale także z twoimi kosztami i marżami zysku. Możesz nie chcieć, aby twoje pola kosztów i pole formuły marży projektowej były widoczne dla młodszych pracowników, a niestandardowe uprawnienia pozwalają ci zablokować te pola, aby nie były wyświetlane.

Następnie przeszłyśmy do tworzenia uprawnień opartych na listach, gdzie administratorzy projektu mogą decydować, czy grupa użytkowników może przeglądać, edytować i usuwać określoną listę. Jeśli ukryją listę, wszystkie rekordy w tej liście również stają się ukryte, co jest świetne, ponieważ oznacza, że możesz ukryć niektóre części swojego procesu przed członkami zespołu lub klientami.

Oto końcowy rezultat:

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Rozważania techniczne

W sercu technicznej architektury Blue leży GraphQL, kluczowy wybór, który znacząco wpłynął na naszą zdolność do wdrażania złożonych funkcji, takich jak niestandardowe uprawnienia. Ale zanim przejdziemy do szczegółów, zróbmy krok wstecz i zrozummy, czym jest GraphQL i jak różni się od bardziej tradycyjnego podejścia REST API.
GraphQL vs REST API: Przystępne wyjaśnienie

Wyobraź sobie, że jesteś w restauracji. Z REST API to tak, jakby zamawiać z ustalonego menu. Proszisz o konkretne danie (punkt końcowy), a dostajesz wszystko, co się z nim wiąże, niezależnie od tego, czy tego chcesz, czy nie. Jeśli chcesz dostosować swoje danie, może być konieczne złożenie wielu zamówień (wywołań API) lub poproszenie o specjalnie przygotowane danie (niestandardowy punkt końcowy).

GraphQL, z drugiej strony, jest jak rozmowa z szefem kuchni, który może przygotować cokolwiek. Mówisz szefowi kuchni dokładnie, jakie składniki chcesz (pola danych) i w jakich ilościach. Szef kuchni przygotowuje danie, które jest dokładnie tym, o co prosiłeś — ani więcej, ani mniej. To zasadniczo to, co robi GraphQL — pozwala klientowi zapytać o dokładnie te dane, których potrzebuje, a serwer dostarcza tylko to.

### Ważny lunch

Około sześciu tygodni po początkowym rozwoju Blue, nasz główny inżynier i CEO poszli na lunch.

Tematem dyskusji?

Czy przejść z REST API na GraphQL. To nie była decyzja, którą można podjąć lekko — przyjęcie GraphQL oznaczałoby porzucenie sześciu tygodni początkowej pracy.

W drodze powrotnej do biura CEO zadał kluczowe pytanie głównemu inżynierowi: "Czy będziemy żałować, że tego nie zrobiliśmy za pięć lat?"

Odpowiedź stała się jasna: GraphQL był drogą naprzód.

Wczesne dostrzegłyśmy potencjał tej technologii, widząc, jak może wspierać naszą wizję elastycznej, potężnej platformy do zarządzania projektami.

Nasza przewidywalność w przyjęciu GraphQL przyniosła korzyści, gdy przyszło do wdrażania niestandardowych uprawnień. Z REST API potrzebowałybyśmy innego punktu końcowego dla każdej możliwej konfiguracji niestandardowych uprawnień — podejścia, które szybko stałoby się nieporęczne i trudne do utrzymania.

GraphQL jednak pozwala nam dynamicznie obsługiwać niestandardowe uprawnienia. Oto jak to działa:

- **Sprawdzanie uprawnień w locie**: Gdy klient składa żądanie, nasz serwer GraphQL może sprawdzić uprawnienia użytkownika bezpośrednio z naszej bazy danych.
- **Precyzyjne pobieranie danych**: Na podstawie tych uprawnień GraphQL zwraca tylko żądane dane, które pasują do praw dostępu użytkownika.
- **Elastyczne zapytania**: W miarę jak uprawnienia się zmieniają, nie musimy tworzyć nowych punktów końcowych ani zmieniać istniejących. To samo zapytanie GraphQL może dostosować się do różnych konfiguracji uprawnień.
- **Efektywne pobieranie danych**: GraphQL pozwala klientom żądać dokładnie tego, czego potrzebują. Oznacza to, że nie pobieramy nadmiarowych danych, co mogłoby potencjalnie ujawnić informacje, do których użytkownik nie powinien mieć dostępu.

Ta elastyczność jest kluczowa dla funkcji tak złożonej jak niestandardowe uprawnienia. Pozwala nam oferować granularną kontrolę *bez* poświęcania wydajności lub łatwości utrzymania.

## Wyzwania

Wdrożenie niestandardowych uprawnień w Blue przyniosło swoje wyzwania, z których każde zmuszało nas do innowacji i udoskonalania naszego podejścia. Optymalizacja wydajności szybko stała się kluczową kwestią. W miarę dodawania bardziej granularnych sprawdzeń uprawnień, ryzykowałyśmy spowolnienie naszego systemu, szczególnie w przypadku dużych projektów z wieloma użytkownikami i złożonymi konfiguracjami uprawnień. Aby temu zaradzić, wdrożyłyśmy strategię wielowarstwowego buforowania, zoptymalizowałyśmy nasze zapytania do bazy danych i wykorzystałyśmy zdolność GraphQL do żądania tylko niezbędnych danych. To podejście pozwoliło nam utrzymać szybkie czasy odpowiedzi, nawet gdy projekty rosły, a złożoność uprawnień wzrastała.

Interfejs użytkownika dla niestandardowych uprawnień stanowił kolejną znaczącą przeszkodę. Musiałyśmy uczynić interfejs intuicyjnym i zarządzalnym dla administratorów, nawet gdy dodawałyśmy więcej opcji i zwiększałyśmy złożoność systemu.

Naszym rozwiązaniem były wielokrotne rundy testów użytkowników i iteracyjne projektowanie.

Wprowadziłyśmy wizualną macierz uprawnień, która pozwalała administratorom szybko przeglądać i modyfikować uprawnienia w różnych rolach i obszarach projektów.

Zapewnienie spójności między platformami stanowiło własny zestaw wyzwań. Musiałyśmy wdrożyć niestandardowe uprawnienia jednolicie w naszych aplikacjach webowych, desktopowych i mobilnych, z każdą z unikalnym interfejsem i uwagami dotyczącymi doświadczeń użytkowników. To było szczególnie trudne dla naszych aplikacji mobilnych, które musiały dynamicznie ukrywać i pokazywać funkcje w zależności od uprawnień użytkownika. Rozwiązałyśmy to, centralizując naszą logikę uprawnień w warstwie API, zapewniając, że wszystkie platformy otrzymują spójne dane dotyczące uprawnień.

Następnie opracowałyśmy elastyczny framework UI, który mógł dostosować się do tych zmian uprawnień w czasie rzeczywistym, zapewniając płynne doświadczenie niezależnie od używanej platformy.

Edukacja użytkowników i adopcja stanowiły ostatnią przeszkodę w naszej podróży z niestandardowymi uprawnieniami. Wprowadzenie tak potężnej funkcji oznaczało, że musiałyśmy pomóc naszym użytkownikom zrozumieć i skutecznie wykorzystywać niestandardowe uprawnienia.

Początkowo uruchomiłyśmy niestandardowe uprawnienia dla części naszej bazy użytkowników, starannie monitorując ich doświadczenia i zbierając informacje. To podejście pozwoliło nam udoskonalić funkcję i nasze materiały edukacyjne na podstawie rzeczywistego użytkowania, zanim uruchomiliśmy ją dla całej bazy użytkowników.

Fazowe wdrożenie okazało się nieocenione, pomagając nam zidentyfikować i rozwiązać drobne problemy oraz punkty dezorientacji użytkowników, których nie przewidziałyśmy, co ostatecznie prowadziło do bardziej dopracowanej i przyjaznej dla użytkownika funkcji dla wszystkich naszych użytkowników.

To podejście uruchamiania dla podzbioru użytkowników, a także nasz typowy 2-3 tygodniowy okres "Beta" w naszym publicznym Beta, pomaga nam spać spokojnie. :)

## Patrząc w przyszłość

Jak w przypadku wszystkich funkcji, nic nigdy nie jest *"zrobione"*.

Nasza długoterminowa wizja dla funkcji niestandardowych uprawnień rozciąga się na tagi, filtry pól niestandardowych, konfigurowalną nawigację projektów i kontrolę komentarzy.

Zanurzmy się w każdy z tych aspektów.

### Uprawnienia tagów

Uważamy, że byłoby niesamowicie móc tworzyć uprawnienia na podstawie tego, czy rekord ma jeden lub więcej tagów. Najbardziej oczywistym przypadkiem użycia byłoby stworzenie niestandardowej roli użytkownika o nazwie "Klienci" i pozwolenie tylko użytkownikom w tej roli na widzenie rekordów, które mają tag "Klienci".

To daje ci widok na pierwszy rzut oka, czy rekord może być widziany przez twoich klientów, czy nie.

To mogłoby stać się jeszcze potężniejsze dzięki kombinatorom AND/OR, gdzie można określić bardziej złożone zasady. Na przykład, można ustawić regułę, która pozwala na dostęp do rekordów oznaczonych zarówno "Klienci", jak i "Publiczne", lub rekordów oznaczonych "Wewnętrzne" lub "Poufne". Ten poziom elastyczności pozwoliłby na niezwykle zniuansowane ustawienia uprawnień, dostosowując się do nawet najbardziej złożonych struktur organizacyjnych i przepływów pracy.

Potencjalne zastosowania są ogromne. Menedżerowie projektów mogliby łatwo segregować wrażliwe informacje, zespoły sprzedażowe mogłyby mieć automatyczny dostęp do odpowiednich danych klientów, a zewnętrzni współpracownicy mogliby być płynnie integrowani w określone części projektu bez ryzyka ujawnienia wrażliwych informacji wewnętrznych.

### Filtry pól niestandardowych

Nasza wizja dla filtrów pól niestandardowych stanowi znaczący krok naprzód w granularnej kontroli dostępu. Ta funkcja umożliwi administratorom projektów definiowanie, które rekordy konkretne grupy użytkowników mogą widzieć na podstawie wartości pól niestandardowych. Chodzi o tworzenie dynamicznych, opartych na danych granic dostępu do informacji.

Wyobraź sobie możliwość ustawienia uprawnień w ten sposób:

- Pokaż tylko rekordy, w których rozwijana lista "Status projektu" jest ustawiona na "Publiczny"
- Ogranicz widoczność do elementów, w których pole wielokrotnego wyboru "Departament" zawiera "Marketing"
- Pozwól na dostęp do zadań, w których pole wyboru "Priorytet" jest zaznaczone
- Wyświetl projekty, w których pole liczby "Budżet" jest powyżej określonego progu

### Konfigurowalna nawigacja projektów

To po prostu rozszerzenie przełączników, które już mamy. Zamiast mieć tylko przełączniki dla "aktywności" i "formularzy", chcemy rozszerzyć to na każdą część nawigacji projektu. W ten sposób administratorzy projektów mogą tworzyć skoncentrowane interfejsy i usuwać narzędzia, których nie potrzebują.

### Kontrola komentarzy

W przyszłości chcemy być kreatywne w tym, jak pozwalamy naszym klientom decydować, kto może, a kto nie może widzieć komentarzy. Możemy pozwolić na wiele zakładkowych obszarów komentarzy pod jednym rekordem, a każdy z nich może być widoczny lub niewidoczny dla różnych grup użytkowników.

Dodatkowo, możemy również pozwolić na funkcję, w której widoczne są tylko komentarze, w których użytkownik jest *specjalnie* wspomniany, a nic więcej. To pozwoliłoby zespołom, które mają klientów w projektach, upewnić się, że tylko komentarze, które chcą, aby klienci widzieli, są widoczne.

## Podsumowanie

I oto mamy to, w ten sposób podeszłyśmy do budowania jednej z najbardziej interesujących i potężnych funkcji! [Jak widać w naszym narzędziu porównawczym do zarządzania projektami](/compare), bardzo niewiele systemów do zarządzania projektami ma tak potężne ustawienia macierzy uprawnień, a te, które je mają, rezerwują je dla swoich najdroższych planów enterprise, co czyni je niedostępnymi dla typowej małej lub średniej firmy.

Z Blue masz *wszystkie* funkcje dostępne w naszym planie — nie wierzymy, że funkcje na poziomie enterprise powinny być zarezerwowane dla klientów enterprise!