---
title: Automatyczna kategoryzacja AI (szczegółowa analiza techniczna)
category: "Engineering"
description: Zajrzyj za kulisy z zespołem inżynieryjnym Blue, który wyjaśnia, jak zbudowali funkcję automatycznej kategoryzacji i tagowania wspieraną przez AI.
date: 2024-12-07
---

Niedawno udostępniliśmy [Automatyczną kategoryzację AI](/insights/ai-auto-categorization) wszystkim użytkownikom Blue. Jest to funkcja AI, która jest zawarta w podstawowej subskrypcji Blue, bez dodatkowych kosztów. W tym wpisie zagłębiamy się w inżynierię stojącą za stworzeniem tej funkcji.

---
W Blue nasze podejście do rozwoju funkcji jest zakorzenione w głębokim zrozumieniu potrzeb użytkowników i trendów rynkowych, połączonym z zaangażowaniem w utrzymanie prostoty i łatwości użytkowania, które definiują naszą platformę. To właśnie napędza naszą [mapę drogową](/platform/roadmap) i pozwoliło nam [konsekwentnie dostarczać funkcje każdego miesiąca przez lata](/platform/changelog).

Wprowadzenie automatycznego tagowania wspieranego przez AI do Blue jest doskonałym przykładem tej filozofii w działaniu. Zanim zagłębimy się w szczegóły techniczne tego, jak zbudowaliśmy tę funkcję, kluczowe jest zrozumienie problemu, który rozwiązywaliśmy, oraz starannego namysłu, który włożyliśmy w jej rozwój.

Krajobraz zarządzania projektami szybko ewoluuje, a możliwości AI stają się coraz bardziej centralne dla oczekiwań użytkowników. Nasi klienci, szczególnie ci zarządzający dużymi [projektami](/platform) z milionami [rekordów](/platform/features/records), byli bardzo wyraźni w swoim pragnieniu inteligentniejszych, bardziej wydajnych sposobów organizowania i kategoryzowania swoich danych.

Jednak w Blue nie dodajemy po prostu funkcji, ponieważ są modne lub o nie proszono. Nasza filozofia polega na tym, że każde nowe dodatek musi udowodnić swoją wartość, z domyślną odpowiedzią będącą stanowczym *"nie"*, dopóki funkcja nie wykaże silnego popytu i wyraźnej użyteczności.

Aby naprawdę zrozumieć głębię problemu i potencjał automatycznego tagowania AI, przeprowadziliśmy obszerne wywiady z klientami, koncentrując się na długoletnich użytkownikach, którzy zarządzają złożonymi, bogatymi w dane projektami w wielu domenach.

Te rozmowy ujawniły wspólny wątek: *podczas gdy tagowanie było nieocenione dla organizacji i możliwości wyszukiwania, ręczny charakter procesu stawał się wąskim gardłem, szczególnie dla zespołów radzących sobie z dużymi wolumenami rekordów.*

Ale widzieliśmy więcej niż tylko rozwiązanie bezpośredniego problemu ręcznego tagowania.

Wyobraziliśmy sobie przyszłość, w której tagowanie wspierane przez AI mogłoby stać się fundamentem bardziej inteligentnych, zautomatyzowanych przepływów pracy.

Prawdziwa moc tej funkcji, zdaliśmy sobie sprawę, leżała w jej potencjalnej integracji z naszym [systemem automatyzacji zarządzania projektami](/platform/features/automations). Wyobraź sobie narzędzie do zarządzania projektami, które nie tylko inteligentnie kategoryzuje informacje, ale także wykorzystuje te kategorie do kierowania zadań, wyzwalania akcji i adaptowania przepływów pracy w czasie rzeczywistym.

Ta wizja idealnie zgadzała się z naszym celem utrzymania Blue prostym, ale potężnym.

Co więcej, rozpoznaliśmy potencjał rozszerzenia tej możliwości poza granice naszej platformy. Rozwijając solidny system tagowania AI, kładliśmy fundamenty pod "API kategoryzacji", które mogłoby działać od razu po wyjęciu z pudełka, potencjalnie otwierając nowe drogi dla sposobu, w jaki nasi użytkownicy wchodzą w interakcje i wykorzystują Blue w swoich szerszych ekosystemach technologicznych.

Ta funkcja nie była więc tylko o dodaniu pola wyboru AI do naszej listy funkcji.

Chodziło o wykonanie znaczącego kroku w kierunku bardziej inteligentnej, adaptacyjnej platformy zarządzania projektami, pozostając jednocześnie wiernym naszej podstawowej filozofii prostoty i koncentracji na użytkowniku.

W kolejnych sekcjach zagłębimy się w wyzwania techniczne, z którymi zmierzyliśmy się, realizując tę wizję, architekturę, którą zaprojektowaliśmy, aby ją wspierać, oraz rozwiązania, które wdrożyliśmy. Zbadamy również przyszłe możliwości, które otwiera ta funkcja, pokazując, jak starannie przemyślany dodatek może utorować drogę dla transformacyjnych zmian w zarządzaniu projektami.

---
## Problem

Jak omówiono powyżej, ręczne tagowanie rekordów projektowych może być czasochłonne i niespójne.

Postanowiliśmy rozwiązać ten problem, wykorzystując AI do automatycznego sugerowania tagów na podstawie zawartości rekordów.

Główne wyzwania to:

1. Wybór odpowiedniego modelu AI
2. Efektywne przetwarzanie dużych wolumenów rekordów
3. Zapewnienie prywatności i bezpieczeństwa danych
4. Bezproblemowa integracja funkcji z naszą istniejącą architekturą

## Wybór modelu AI

Oceniliśmy kilka platform AI, w tym [OpenAI](https://openai.com), modele open-source na [HuggingFace](https://huggingface.co/) i [Replicate](https://replicate.com).

Nasze kryteria obejmowały:

- Opłacalność
- Dokładność w rozumieniu kontekstu
- Zdolność do przestrzegania określonych formatów wyjściowych
- Gwarancje prywatności danych

Po dokładnych testach wybraliśmy [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo) od OpenAI. Chociaż [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) może oferować marginalne ulepszenia w dokładności, nasze testy pokazały, że wydajność GPT-3.5 była więcej niż wystarczająca dla naszych potrzeb automatycznego tagowania. Równowaga opłacalności i silnych możliwości kategoryzacji sprawiła, że GPT-3.5 był idealnym wyborem dla tej funkcji.

Wyższy koszt GPT-4 zmusiłby nas do oferowania funkcji jako płatnego dodatku, co byłoby sprzeczne z naszym celem **włączenia AI do naszego głównego produktu bez dodatkowych kosztów dla użytkowników końcowych.**

W momencie naszej implementacji ceny za GPT-3.5 Turbo wynoszą:

- $0.0005 za 1K tokenów wejściowych (lub $0.50 za 1M tokenów wejściowych)
- $0.0015 za 1K tokenów wyjściowych (lub $1.50 za 1M tokenów wyjściowych)

Przyjmijmy pewne założenia dotyczące przeciętnego rekordu w Blue:

- **Tytuł**: ~10 tokenów
- **Opis**: ~50 tokenów
- **2 komentarze**: ~30 tokenów każdy
- **5 pól niestandardowych**: ~10 tokenów każde
- **Nazwa listy, termin wykonania i inne metadane**: ~20 tokenów
- **Prompt systemowy i dostępne tagi**: ~50 tokenów

Łączna liczba tokenów wejściowych na rekord: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 tokenów

Dla wyjścia przyjmijmy średnio 3 tagi sugerowane na rekord, co może wynosić około 20 tokenów wyjściowych, włączając formatowanie JSON.

Dla 1 miliona rekordów:

- Koszt wejścia: (240 * 1,000,000 / 1,000,000) * $0.50 = $120
- Koszt wyjścia: (20 * 1,000,000 / 1,000,000) * $1.50 = $30

**Całkowity koszt automatycznego tagowania 1 miliona rekordów: $120 + $30 = $150**

## Wydajność GPT3.5 Turbo

Kategoryzacja to zadanie, w którym duże modele językowe (LLMs) takie jak GPT-3.5 Turbo doskonale sobie radzą, co czyni je szczególnie dobrze dopasowanymi do naszej funkcji automatycznego tagowania. LLMs są trenowane na ogromnych ilościach danych tekstowych, co pozwala im rozumieć kontekst, semantykę i relacje między pojęciami. Ta szeroka baza wiedzy umożliwia im wykonywanie zadań kategoryzacji z wysoką dokładnością w szerokim zakresie dziedzin.

Dla naszego konkretnego przypadku użycia tagowania zarządzania projektami, GPT-3.5 Turbo wykazuje kilka kluczowych mocnych stron:

- **Zrozumienie kontekstu:** Potrafi uchwycić ogólny kontekst rekordu projektu, biorąc pod uwagę nie tylko pojedyncze słowa, ale znaczenie przekazywane przez cały opis, komentarze i inne pola.
- **Elastyczność:** Może dostosować się do różnych typów projektów i branż bez konieczności rozległego przeprogramowania.
- **Radzenie sobie z niejednoznacznością:** Może rozważyć wiele czynników, aby podjąć niuansowane decyzje.
- **Uczenie się z przykładów:** Może szybko zrozumieć i zastosować nowe schematy kategoryzacji bez dodatkowego treningu.
- **Klasyfikacja wieloetykietowa:** Może sugerować wiele istotnych tagów dla pojedynczego rekordu, co było kluczowe dla naszych wymagań.

GPT-3.5 Turbo wyróżniał się również niezawodnością w przestrzeganiu naszego wymaganego formatu wyjściowego JSON, co było *kluczowe* dla bezproblemowej integracji z naszymi istniejącymi systemami. Modele open-source, choć obiecujące, często dodawały dodatkowe komentarze lub odbiegały od oczekiwanego formatu, co wymagałoby dodatkowego przetwarzania po stronie. Ta spójność w formacie wyjściowym była kluczowym czynnikiem w naszej decyzji, ponieważ znacznie uprościła naszą implementację i zmniejszyła potencjalne punkty awarii.

Wybór GPT-3.5 Turbo ze spójnym wyjściem JSON pozwolił nam wdrożyć bardziej bezpośrednie, niezawodne i łatwe w utrzymaniu rozwiązanie.

Gdybyśmy wybrali model z mniej niezawodnym formatowaniem, stanęlibyśmy przed kaskadą komplikacji: potrzebą solidnej logiki parsowania do obsługi różnych formatów wyjściowych, rozbudowaną obsługą błędów dla niespójnych wyjść, potencjalnym wpływem na wydajność z powodu dodatkowego przetwarzania, zwiększoną złożonością testowania, aby pokryć wszystkie warianty wyjściowe, i większym długoterminowym obciążeniem konserwacyjnym.

Błędy parsowania mogłyby prowadzić do nieprawidłowego tagowania, negatywnie wpływając na doświadczenie użytkownika. Unikając tych pułapek, byliśmy w stanie skupić nasze wysiłki inżynieryjne na krytycznych aspektach, takich jak optymalizacja wydajności i projektowanie interfejsu użytkownika, zamiast zmagać się z nieprzewidywalnymi wyjściami AI.

## Architektura systemu

Nasza funkcja automatycznego tagowania AI jest zbudowana na solidnej, skalowalnej architekturze zaprojektowanej do efektywnej obsługi dużych wolumenów żądań, zapewniając jednocześnie płynne doświadczenie użytkownika. Jak w przypadku wszystkich naszych systemów, zaprojektowaliśmy tę funkcję tak, aby wspierała ruch o jeden rząd wielkości większy niż obecnie doświadczamy. To podejście, choć pozornie przesadzone dla obecnych potrzeb, jest najlepszą praktyką, która pozwala nam bezproblemowo obsługiwać nagłe skoki użytkowania i daje nam wystarczającą przestrzeń do wzrostu bez większych przebudów architektonicznych. W przeciwnym razie musielibyśmy przebudowywać wszystkie nasze systemy co 18 miesięcy — coś, czego nauczyliśmy się na własnej skórze w przeszłości!

Rozłóżmy komponenty i przepływ naszego systemu:

- **Interakcja użytkownika:** Proces rozpoczyna się, gdy użytkownik naciska przycisk "Autotaguj" w interfejsie Blue. Ta akcja uruchamia przepływ pracy automatycznego tagowania.
- **Wywołanie API Blue:** Akcja użytkownika jest tłumaczona na wywołanie API do naszego backendu Blue. Ten punkt końcowy API jest zaprojektowany do obsługi żądań automatycznego tagowania.
- **Zarządzanie kolejką:** Zamiast przetwarzać żądanie natychmiast, co mogłoby prowadzić do problemów z wydajnością przy dużym obciążeniu, dodajemy żądanie tagowania do kolejki. Używamy Redis do tego mechanizmu kolejkowania, co pozwala nam efektywnie zarządzać obciążeniem i zapewniać skalowalność systemu.
- **Usługa w tle:** Wdrożyliśmy usługę działającą w tle, która stale monitoruje kolejkę w poszukiwaniu nowych żądań. Ta usługa jest odpowiedzialna za przetwarzanie żądań w kolejce.
- **Integracja z API OpenAI:** Usługa w tle przygotowuje niezbędne dane i wykonuje wywołania API do modelu GPT-3.5 OpenAI. To tutaj następuje faktyczne tagowanie wspierane przez AI. Wysyłamy istotne dane projektu i otrzymujemy w zamian sugerowane tagi.
- **Przetwarzanie wyników:** Usługa w tle przetwarza wyniki otrzymane z OpenAI. Obejmuje to parsowanie odpowiedzi AI i przygotowanie danych do zastosowania w projekcie.
- **Zastosowanie tagów:** Przetworzone wyniki są używane do zastosowania nowych tagów do odpowiednich elementów w projekcie. Ten krok aktualizuje naszą bazę danych o tagi sugerowane przez AI.
- **Odzwierciedlenie w interfejsie użytkownika:** Na koniec nowe tagi pojawiają się w widoku projektu użytkownika, kończąc proces automatycznego tagowania z perspektywy użytkownika.

Ta architektura oferuje kilka kluczowych korzyści, które poprawiają zarówno wydajność systemu, jak i doświadczenie użytkownika. Wykorzystując kolejkę i przetwarzanie w tle, osiągnęliśmy imponującą skalowalność, pozwalając nam obsługiwać liczne żądania jednocześnie bez przeciążania naszego systemu lub przekraczania limitów szybkości API OpenAI. Wdrożenie tej architektury wymagało starannego rozważenia różnych czynników, aby zapewnić optymalną wydajność i niezawodność. Do zarządzania kolejką wybraliśmy Redis, wykorzystując jego szybkość i niezawodność w obsłudze rozproszonych kolejek.

To podejście przyczynia się również do ogólnej responsywności funkcji. Użytkownicy otrzymują natychmiastową informację zwrotną, że ich żądanie jest przetwarzane, nawet jeśli faktyczne tagowanie zajmuje trochę czasu, tworząc poczucie interakcji w czasie rzeczywistym. Tolerancja na błędy architektury to kolejna kluczowa zaleta. Jeśli jakakolwiek część procesu napotka problemy, takie jak tymczasowe zakłócenia API OpenAI, możemy elegancko ponowić próbę lub obsłużyć awarię bez wpływu na cały system.

Ta solidność, w połączeniu z pojawianiem się tagów w czasie rzeczywistym, poprawia doświadczenie użytkownika, dając wrażenie "magii" AI w działaniu.

## Dane i prompty

Kluczowym krokiem w naszym procesie automatycznego tagowania jest przygotowanie danych do wysłania do modelu GPT-3.5. Ten krok wymagał starannego rozważenia, aby zrównoważyć dostarczanie wystarczającego kontekstu dla dokładnego tagowania przy jednoczesnym zachowaniu wydajności i ochronie prywatności użytkowników. Oto szczegółowy przegląd naszego procesu przygotowania danych.

Dla każdego rekordu kompilujemy następujące informacje:

- **Nazwa listy**: Zapewnia kontekst dotyczący szerszej kategorii lub fazy projektu.
- **Tytuł rekordu**: Często zawiera kluczowe informacje o celu lub zawartości rekordu.
- **Pola niestandardowe**: Włączamy tekstowe i liczbowe [pola niestandardowe](/platform/features/custom-fields), które często zawierają kluczowe informacje specyficzne dla projektu.
- **Opis**: Zazwyczaj zawiera najbardziej szczegółowe informacje o rekordzie.
- **Komentarze**: Mogą dostarczyć dodatkowy kontekst lub aktualizacje, które mogą być istotne dla tagowania.
- **Termin wykonania**: Informacje czasowe, które mogą wpływać na wybór tagu.

Co ciekawe, nie wysyłamy istniejących danych tagów do GPT-3.5, i robimy to, aby uniknąć uprzedzeń modelu.

Rdzeń naszej funkcji automatycznego tagowania leży w sposobie, w jaki wchodzimy w interakcję z modelem GPT-3.5 i przetwarzamy jego odpowiedzi. Ta sekcja naszego potoku wymagała starannego projektowania, aby zapewnić dokładne, spójne i wydajne tagowanie.

Używamy starannie przygotowanego promptu systemowego, aby poinstruować AI o jego zadaniu. Oto rozkład naszego promptu i uzasadnienie każdego komponentu:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Definicja zadania:** Jasno określamy zadanie AI, aby zapewnić skoncentrowane odpowiedzi.
- **Obsługa niepewności:** Wyraźnie pozwalamy na puste odpowiedzi, zapobiegając wymuszaniu tagowania, gdy AI nie jest pewne.
- **Dostępne tagi:** Dostarczamy listę ważnych tagów (${tags}), aby ograniczyć wybory AI do istniejących tagów projektu.
- **Bieżąca data:** Włączenie ${currentDate} pomaga AI zrozumieć kontekst czasowy, który może być kluczowy dla niektórych typów projektów.
- **Format odpowiedzi:** Określamy format JSON dla łatwego parsowania i sprawdzania błędów.

Ten prompt jest wynikiem rozległych testów i iteracji. Odkryliśmy, że bycie jednoznacznym co do zadania, dostępnych opcji i pożądanego formatu wyjściowego znacznie poprawiło dokładność i spójność odpowiedzi AI — prostota jest kluczowa!

Lista dostępnych tagów jest generowana po stronie serwera i walidowana przed włączeniem do promptu. Wdrażamy ścisłe limity znaków w nazwach tagów, aby zapobiec zbyt dużym promptom.

Jak wspomniano powyżej, nie mieliśmy żadnych problemów z GPT-3.5 Turbo w otrzymywaniu czystej odpowiedzi JSON w prawidłowym formacie w 100% przypadków.

Podsumowując:

- Łączymy prompt systemowy z przygotowanymi danymi rekordu.
- Ten połączony prompt jest wysyłany do modelu GPT-3.5 przez API OpenAI.
- Używamy ustawienia temperatury 0.3, aby zrównoważyć kreatywność i spójność w odpowiedziach AI.
- Nasze wywołanie API zawiera parametr max_tokens, aby ograniczyć rozmiar odpowiedzi i kontrolować koszty.

Po otrzymaniu odpowiedzi AI przechodzimy przez kilka kroków, aby przetworzyć i zastosować sugerowane tagi:

* **Parsowanie JSON**: Próbujemy sparsować odpowiedź jako JSON. Jeśli parsowanie się nie powiedzie, logujemy błąd i pomijamy tagowanie dla tego rekordu.
* **Walidacja schematu**: Walidujemy sparsowany JSON względem naszego oczekiwanego schematu (obiekt z tablicą "tags"). To wyłapuje wszelkie odchylenia strukturalne w odpowiedzi AI.
* **Walidacja tagów**: Porównujemy sugerowane tagi z naszą listą ważnych tagów projektu. Ten krok filtruje wszystkie tagi, które nie istnieją w projekcie, co może się zdarzyć, jeśli AI źle zrozumiało lub jeśli tagi projektu zmieniły się między utworzeniem promptu a przetwarzaniem odpowiedzi.
* **Deduplikacja**: Usuwamy wszelkie zduplikowane tagi z sugestii AI, aby uniknąć redundantnego tagowania.
* **Zastosowanie**: Zwalidowane i zdeduplikowane tagi są następnie stosowane do rekordu w naszej bazie danych.
* **Logowanie i analityka**: Logujemy ostatecznie zastosowane tagi. Te dane są cenne dla monitorowania wydajności systemu i jego ulepszania w czasie.

## Wyzwania

Wdrożenie automatycznego tagowania wspieranego przez AI w Blue przedstawiło kilka unikalnych wyzwań, z których każde wymagało innowacyjnych rozwiązań, aby zapewnić solidną, wydajną i przyjazną dla użytkownika funkcję.

### Cofnij operację masową

Funkcja tagowania AI może być wykonywana zarówno na pojedynczych rekordach, jak i masowo. Problem z operacją masową polega na tym, że jeśli użytkownikowi nie podoba się wynik, musiałby ręcznie przejść przez tysiące rekordów i cofnąć pracę AI. Oczywiście to jest niedopuszczalne.

Aby to rozwiązać, wdrożyliśmy innowacyjny system sesji tagowania. Każdej operacji masowego tagowania przypisywany jest unikalny identyfikator sesji, który jest powiązany ze wszystkimi tagami zastosowanymi podczas tej sesji. Pozwala nam to efektywnie zarządzać operacjami cofania poprzez proste usunięcie wszystkich tagów powiązanych z określonym identyfikatorem sesji. Usuwamy również powiązane ślady audytu, zapewniając, że cofnięte operacje nie pozostawiają śladu w systemie. To podejście daje użytkownikom pewność eksperymentowania z tagowaniem AI, wiedząc, że mogą łatwo cofnąć zmiany, jeśli będzie to potrzebne.

### Prywatność danych

Prywatność danych była kolejnym kluczowym wyzwaniem, przed którym stanęliśmy.

Nasi użytkownicy powierzają nam swoje dane projektowe i było niezwykle ważne, aby zapewnić, że te informacje nie były zachowywane ani wykorzystywane do trenowania modeli przez OpenAI. Poradziliśmy sobie z tym na wielu frontach.

Po pierwsze, zawarliśmy umowę z OpenAI, która wyraźnie zabrania wykorzystywania naszych danych do trenowania modeli. Dodatkowo OpenAI usuwa dane po przetworzeniu, zapewniając dodatkową warstwę ochrony prywatności.

Z naszej strony podjęliśmy środek ostrożności polegający na wykluczeniu wrażliwych informacji, takich jak szczegóły osób przypisanych, z danych wysyłanych do AI, więc zapewnia to, że konkretne nazwiska osób nie są wysyłane do stron trzecich wraz z innymi danymi.

To kompleksowe podejście pozwala nam wykorzystać możliwości AI przy jednoczesnym zachowaniu najwyższych standardów prywatności i bezpieczeństwa danych.

### Limity szybkości i wyłapywanie błędów

Jedną z naszych głównych obaw była skalowalność i ograniczanie szybkości. Bezpośrednie wywołania API do OpenAI dla każdego rekordu byłyby nieefektywne i mogłyby szybko osiągnąć limity szybkości, szczególnie w przypadku dużych projektów lub podczas szczytowego użytkowania. Aby temu zaradzić, opracowaliśmy architekturę usługi działającej w tle, która pozwala nam grupować żądania i wdrożyć nasz własny system kolejkowania. To podejście pomaga nam zarządzać częstotliwością wywołań API i umożliwia bardziej wydajne przetwarzanie dużych wolumenów rekordów, zapewniając płynną wydajność nawet przy dużym obciążeniu.

Charakter interakcji AI oznaczał, że musieliśmy również przygotować się na sporadyczne błędy lub nieoczekiwane wyniki. Zdarzały się przypadki, gdy AI mogło generować nieprawidłowy JSON lub wyniki, które nie pasowały do naszego oczekiwanego formatu. Aby sobie z tym poradzić, wdrożyliśmy solidną obsługę błędów i logikę parsowania w całym naszym systemie. Jeśli odpowiedź AI nie jest prawidłowym JSON lub nie zawiera oczekiwanego klucza "tags", nasz system jest zaprojektowany tak, aby traktować to tak, jakby nie sugerowano żadnych tagów, zamiast próbować przetwarzać potencjalnie uszkodzone dane. To zapewnia, że nawet w obliczu nieprzewidywalności AI nasz system pozostaje stabilny i niezawodny.

## Przyszłe rozwinięcia

Wierzymy, że funkcje i produkt Blue jako całość nigdy nie są "skończone" — zawsze jest miejsce na ulepszenia.

Były pewne funkcje, które rozważaliśmy w początkowej budowie, które nie przeszły fazy określania zakresu, ale warto zauważyć, że prawdopodobnie wdrożymy pewną ich wersję w przyszłości.

Pierwszą jest dodanie opisu tagu. Pozwoliłoby to użytkownikom końcowym nie tylko nadać tagom nazwę i kolor, ale także opcjonalny opis. Byłby on również przekazywany do AI, aby pomóc zapewnić dalszy kontekst i potencjalnie poprawić dokładność.

Chociaż dodatkowy kontekst mógłby być cenny, jesteśmy świadomi potencjalnej złożoności, którą może wprowadzić. Istnieje delikatna równowaga do osiągnięcia między dostarczaniem użytecznych informacji a przytłaczaniem użytkowników zbyt dużą ilością szczegółów. Rozwijając tę funkcję, skupimy się na znalezieniu tego słodkiego punktu, w którym dodany kontekst wzbogaca, a nie komplikuje doświadczenie użytkownika.

Być może najbardziej ekscytującym ulepszeniem na naszym horyzoncie jest integracja automatycznego tagowania AI z naszym [systemem automatyzacji zarządzania projektami](/platform/features/automations).

Oznacza to, że funkcja tagowania AI mogłaby być wyzwalaczem lub akcją z automatyzacji. To może być ogromne, ponieważ może przekształcić tę dość prostą funkcję kategoryzacji AI w system routingu pracy oparty na AI.

Wyobraź sobie automatyzację, która stwierdza:

Gdy AI otaguje rekord jako "Krytyczny" -> Przypisz do "Menedżera" i Wyślij niestandardowy e-mail

Oznacza to, że gdy otagowujesz rekord AI, jeśli AI zdecyduje, że jest to krytyczny problem, może automatycznie przypisać menedżera projektu i wysłać mu niestandardowy e-mail. To rozszerza [korzyści naszego systemu automatyzacji zarządzania projektami](/platform/features/automations) z czysto opartego na regułach systemu do prawdziwego elastycznego systemu AI.

Kontynuując eksplorację granic AI w zarządzaniu projektami, dążymy do zapewnienia naszym użytkownikom narzędzi, które nie tylko spełniają ich obecne potrzeby, ale przewidują i kształtują przyszłość pracy. Jak zawsze, będziemy rozwijać te funkcje w ścisłej współpracy z naszą społecznością użytkowników, zapewniając, że każde ulepszenie dodaje rzeczywistą, praktyczną wartość do procesu zarządzania projektami.

## Podsumowanie

Więc to wszystko!

To była zabawna funkcja do wdrożenia i jeden z naszych pierwszych kroków w AI, obok [AI Podsumowania treści](/insights/ai-content-summarization), które wcześniej uruchomiliśmy. Wiemy, że AI będzie odgrywać coraz większą rolę w zarządzaniu projektami w przyszłości i nie możemy się doczekać wprowadzenia kolejnych innowacyjnych funkcji wykorzystujących zaawansowane LLMs (duże modele językowe).

Było sporo do przemyślenia podczas wdrażania tego i jesteśmy szczególnie podekscytowani tym, jak możemy wykorzystać tę funkcję w przyszłości z istniejącym [silnikiem automatyzacji zarządzania projektami](/insights/benefits-project-management-automation) Blue.

Mamy również nadzieję, że była to interesująca lektura i że daje wgląd w to, jak myślimy o inżynierii funkcji, których używasz codziennie.
