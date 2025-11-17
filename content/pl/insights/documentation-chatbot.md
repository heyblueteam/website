---
title: Dlaczego zbudowaliśmy naszego własnego chatbota AI do dokumentacji
description: Zbudowaliśmy naszego własnego chatbota AI do dokumentacji, który jest szkolony na dokumentacji platformy Blue.
category: "Product Updates"
date: 2024-07-09
---



W Blue zawsze szukamy sposobów, aby ułatwić życie naszym klientom. Mamy [szczegółową dokumentację każdej funkcji](https://documentation.blue.cc), [filmy na YouTube](https://www.youtube.com/@workwithblue), [Porady i wskazówki](/insights/tips-tricks) oraz [różne kanały wsparcia](/support). 

Uważnie obserwujemy rozwój AI (sztucznej inteligencji), ponieważ bardzo interesujemy się [automatyzacją zarządzania projektami](/platform/features/automations). Wprowadziliśmy również funkcje takie jak [AI Auto Kategoryzacja](/insights/ai-auto-categorization) i [AI Podsumowania](/insights/ai-content-summarization), aby ułatwić pracę naszym tysiącom klientów. 

Jedno jest pewne: AI jest tutaj, aby zostać, i będzie miało niesamowity wpływ na większość branż, a zarządzanie projektami nie jest wyjątkiem. Dlatego zadaliśmy sobie pytanie, jak możemy jeszcze bardziej wykorzystać AI, aby pomóc w pełnym cyklu życia klienta, od odkrycia, przez przedsprzedaż, onboarding, aż po bieżące pytania.

Odpowiedź była dość jasna: **Potrzebowaliśmy chatbota AI, który byłby szkolony na naszej dokumentacji.**

Spójrzmy prawdzie w oczy: *każda* organizacja powinna mieć chatbota. To doskonały sposób dla klientów, aby uzyskać natychmiastowe odpowiedzi na typowe pytania, bez konieczności przeszukiwania stron gęstej dokumentacji lub Twojej strony internetowej. Znaczenie chatbotów w nowoczesnych stronach marketingowych nie może być przeceniane. 

![](/insights/ai-chatbot-regular.png)

Dla firm programistycznych nie należy traktować strony marketingowej jako osobnej "rzeczy" — *jest* częścią Twojego produktu. Dzieje się tak, ponieważ wpisuje się w typowy cykl życia klienta:

- **Świadomość** (Odkrycie): To tutaj potencjalni klienci po raz pierwszy natrafiają na Twój wspaniały produkt. Twój chatbot może być ich przyjaznym przewodnikiem, wskazującym kluczowe funkcje i korzyści od razu.
- **Rozważanie** (Edukacja): Teraz są ciekawi i chcą dowiedzieć się więcej. Twój chatbot staje się ich osobistym nauczycielem, dostarczając informacje dostosowane do ich specyficznych potrzeb i pytań.
- **Zakup/Konwersja**: To moment prawdy - kiedy potencjalny klient decyduje się na zakup i staje się klientem. Twój chatbot może wygładzić wszelkie ostatnie przeszkody, odpowiedzieć na pytania "tuż przed zakupem" i może nawet zaproponować korzystną ofertę, aby sfinalizować transakcję.
- **Onboarding**: Kupili, co dalej? Twój chatbot przekształca się w pomocnego pomocnika, prowadząc nowych użytkowników przez konfigurację, pokazując im zasady i upewniając się, że nie czują się zagubieni w krainie cudów Twojego produktu.
- **Utrzymanie**: Utrzymanie klientów w dobrym nastroju to klucz do sukcesu. Twój chatbot jest dostępny 24/7, gotowy do rozwiązywania problemów, oferowania wskazówek i trików oraz upewniania się, że Twoi klienci czują się doceniani.
- **Ekspansja**: Czas na rozwój! Twój chatbot może subtelnie sugerować nowe funkcje, upselling lub cross-selling, które są zgodne z tym, jak klient już korzysta z Twojego produktu. To jak posiadanie naprawdę inteligentnego, nieprzeszkadzającego sprzedawcy zawsze w gotowości.
- **Adwokatura**: Zadowoleni klienci stają się Twoimi największymi zwolennikami. Twój chatbot może zachęcać zadowolonych użytkowników do rozpowszechniania informacji, zostawiania recenzji lub uczestniczenia w programach poleceń. To jak posiadanie maszyny do hype'u wbudowanej w Twój produkt!

## Decyzja: Zbudować czy Kupić

Gdy zdecydowaliśmy się na wdrożenie chatbota AI, kolejne ważne pytanie brzmiało: zbudować czy kupić? Jako mały zespół skoncentrowany na naszym podstawowym produkcie, generalnie preferujemy rozwiązania "jako usługa" lub popularne platformy open-source. W końcu nie zajmujemy się wynajdywaniem koła na nowo dla każdej części naszego stosu technologicznego.
Zatem, zakasaliśmy rękawy i zanurzyliśmy się w rynek, poszukując zarówno płatnych, jak i open-source'owych rozwiązań chatbotów AI. 

Nasze wymagania były proste, ale niepodlegające negocjacjom:

- **Niebrandowe doświadczenie**: Ten chatbot to nie tylko miły dodatek; będzie na naszej stronie marketingowej i ostatecznie w naszym produkcie. Nie jesteśmy zainteresowani reklamowaniem marki kogoś innego na naszej własnej przestrzeni cyfrowej.
- **Świetne UX**: Dla wielu potencjalnych klientów ten chatbot może być ich pierwszym punktem kontaktu z Blue. Ustala ton ich postrzegania naszej firmy. Spójrzmy prawdzie w oczy: jeśli nie potrafimy stworzyć odpowiedniego chatbota na naszej stronie, jak możemy oczekiwać, że klienci zaufają nam w kwestii ich kluczowych projektów i procesów?
- **Rozsądny koszt**: Przy dużej bazie użytkowników i planach na integrację chatbota z naszym podstawowym produktem, potrzebowaliśmy rozwiązania, które nie zrujnuje nas finansowo w miarę wzrostu użytkowania. Idealnie, chcieliśmy opcję **BYOK (Bring Your Own Key)**. To pozwoliłoby nam korzystać z naszego własnego klucza OpenAI lub innego klucza usługi AI, płacąc tylko za bezpośrednie koszty zmienne, a nie marżę dla zewnętrznego dostawcy, który tak naprawdę nie uruchamia modeli.
- **Kompatybilność z API OpenAI Assistants**: Jeśli mieliśmy iść w stronę oprogramowania open-source, nie chcieliśmy mieć kłopotów z zarządzaniem pipeline'em do wczytywania dokumentów, indeksowania, baz wektorowych i tym podobnych. Chcieliśmy korzystać z [API OpenAI Assistants](https://platform.openai.com/docs/assistants/overview), które ukrywałoby całą złożoność za API. Szczerze mówiąc — to jest naprawdę dobrze zrobione. 
- **Skalowalność**: Chcemy mieć tego chatbota w wielu miejscach, z potencjalnie dziesiątkami tysięcy użytkowników rocznie. Oczekujemy znacznego użytkowania i nie chcemy być związani z rozwiązaniem, które nie może rosnąć wraz z naszymi potrzebami.

## Komercyjne Chatboty AI

Te, które przeglądaliśmy, miały tendencję do lepszego UX niż rozwiązania open-source — co niestety często się zdarza. Prawdopodobnie pewnego dnia będzie osobna dyskusja na temat *dlaczego* wiele rozwiązań open-source ignoruje lub umniejsza znaczenie UX. 

Podamy tutaj listę, na wypadek gdybyś szukał solidnych komercyjnych ofert:

- **[Chatbase](https://chatbase.co):** Chatbase pozwala na budowanie niestandardowego chatbota AI szkolonego na Twojej bazie wiedzy i dodanie go do Twojej strony internetowej lub interakcję z nim za pośrednictwem ich API. Oferuje funkcje takie jak wiarygodne odpowiedzi, generowanie leadów, zaawansowana analityka i możliwość łączenia z wieloma źródłami danych. Dla nas wydawało się to jedną z najbardziej dopracowanych komercyjnych ofert na rynku. 
- **[DocsBot AI](https://docsbot.ai/):** DocsBot AI tworzy niestandardowe boty ChatGPT szkolone na Twojej dokumentacji i treściach do wsparcia, przedsprzedaży, badań i więcej. Oferuje osadzane widgety, aby łatwo dodać chatbota do Twojej strony internetowej, możliwość automatycznego odpowiadania na zgłoszenia wsparcia oraz potężne API do integracji.
- **[CustomGPT.ai](https://customgpt.ai):** CustomGPT.ai tworzy osobiste doświadczenie chatbota, wczytując Twoje dane biznesowe, w tym treści strony internetowej, helpdesk, bazy wiedzy, dokumenty i inne. Umożliwia leadom zadawanie pytań i uzyskiwanie natychmiastowych odpowiedzi na podstawie Twoich treści, bez potrzeby przeszukiwania. Co ciekawe, również [twierdzą, że pokonują OpenAI w RAG (Retrieval Augmented Generation)!](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: To interesująca oferta komercyjna, ponieważ *również* jest oprogramowaniem open-source. Wydaje się, że jest to jeszcze w wczesnej fazie, a ceny nie wydawały się realistyczne (27 USD/miesiąc za nieograniczone wiadomości nigdy nie zadziała komercyjnie dla nich).

Przyjrzeliśmy się również [InterCom Fin](https://www.intercom.com/fin), który jest częścią ich oprogramowania do wsparcia klienta. To oznaczałoby przejście z [HelpScout](https://wwww.helpscout.com), którego używamy od początku działalności Blue. To mogłoby być możliwe, ale InterCom Fin ma szalone ceny, które po prostu wykluczyły go z rozważania.

I to jest właściwie problem z wieloma ofertami komercyjnymi. InterCom Fin pobiera 0,99 USD za każde zgłoszenie wsparcia, a ChatBase 399 USD/miesiąc za 40 000 wiadomości. To prawie 5 tys. USD rocznie za prosty widget czatu. 

Biorąc pod uwagę, że ceny za inferencję AI spadają jak szalone. OpenAI znacznie obniżyło swoje ceny:

- Oryginalny GPT-4 (8k kontekst) był wyceniony na 0,03 USD za 1K tokenów prompt.
- GPT-4 Turbo (128k kontekst) był wyceniony na 0,01 USD za 1K tokenów prompt, co stanowi 50% redukcję w porównaniu do oryginalnego GPT-4.
- Model GPT-4o jest wyceniony na 0,005 USD za 1K tokenów, co stanowi dalszą 50% redukcję w porównaniu do cen GPT-4 Turbo.

To 83% redukcji kosztów, a nie spodziewamy się, że to pozostanie stagnacyjne. 

Biorąc pod uwagę, że szukaliśmy skalowalnego rozwiązania, które będzie używane przez dziesiątki tysięcy użytkowników rocznie z znaczną ilością wiadomości, ma sens iść bezpośrednio do źródła i płacić za koszty API bez korzystania z wersji komercyjnej, która podnosi koszty.

## Chatboty AI Open Source

Jak wspomniano, opcje open source, które przeglądaliśmy, były głównie rozczarowujące w odniesieniu do wymagań "Świetne UX". 

Przyjrzeliśmy się:

- **[Deepchat](https://deepchat.dev/)**: To komponent czatu niezależny od frameworka dla usług AI, który łączy się z różnymi API AI, w tym OpenAI. Ma również możliwość, aby użytkownicy pobrali model AI, który działa bezpośrednio w przeglądarce. Bawiliśmy się tym i udało nam się uruchomić wersję, ale zaimplementowane API OpenAI Assistants wydawało się dość wadliwe z wieloma problemami. Niemniej jednak, to bardzo obiecujący projekt, a ich plac zabaw jest naprawdę dobrze zrobiony. 
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Patrząc na to ponownie z perspektywy open-source, wymagałoby to od nas uruchomienia sporej infrastruktury, czego nie chcieliśmy robić, ponieważ chcieliśmy polegać tak bardzo, jak to możliwe, na API OpenAI Assistants. 


## Budowanie naszego własnego ChatBota

I tak, nie znajdując czegoś, co spełniałoby wszystkie nasze wymagania, zdecydowaliśmy się zbudować własnego chatbota AI, który mógłby komunikować się z API OpenAI Assistants. To, ostatecznie, okazało się stosunkowo bezbolesne! 

Nasza strona internetowa korzysta z [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (który jest tym samym frameworkiem, co platforma Blue) oraz [TailwindUI](https://tailwindui.com/).

Pierwszym krokiem było stworzenie API (Interfejs Programowania Aplikacji) w Nuxt3, które może "rozmawiać" z API OpenAI Assistants. Było to konieczne, ponieważ nie chcieliśmy robić wszystkiego po stronie front-end, ponieważ narażałoby to nasz klucz API OpenAI na świat, z potencjalnym ryzykiem nadużyć. 

Nasze API backendowe działa jako bezpieczny pośrednik między przeglądarką użytkownika a OpenAI. Oto, co robi:

- **Zarządzanie rozmowami:** Tworzy i zarządza "wątkami" dla każdej rozmowy. Pomyśl o wątku jako o unikalnej sesji czatu, która pamięta wszystko, co powiedziałeś.
- **Obsługa wiadomości:** Kiedy wysyłasz wiadomość, nasze API dodaje ją do odpowiedniego wątku i prosi asystenta OpenAI o sformułowanie odpowiedzi.
- **Inteligentne czekanie:** Zamiast zmuszać Cię do wpatrywania się w ekran ładowania, nasze API sprawdza co sekundę, czy Twoja odpowiedź jest gotowa. To jak mieć kelnera, który pilnuje Twojego zamówienia, nie przeszkadzając kucharzowi co dwie sekundy.
- **Bezpieczeństwo przede wszystkim:** Obsługując to wszystko na serwerze, chronimy Twoje dane i nasze klucze API.

Następnie przeszliśmy do front-endu i doświadczenia użytkownika. Jak wcześniej omówiono, to było *krytycznie* ważne, ponieważ nie mamy drugiej szansy na zrobienie pierwszego wrażenia! 

Projektując naszego chatbota, zwróciliśmy szczególną uwagę na doświadczenie użytkownika, zapewniając, że każda interakcja jest płynna, intuicyjna i odzwierciedla zaangażowanie Blue w jakość. Interfejs chatbota zaczyna się od prostego, eleganckiego niebieskiego okręgu, używając [HeroIcons do naszych ikon](https://heroicons.com/) (które używamy w całej stronie Blue) jako naszego widgetu otwierającego chatbota. Ten wybór projektowy zapewnia spójność wizualną i natychmiastowe rozpoznawanie marki.

![](/insights/ai-chatbot-circle.png)

Rozumiemy, że czasami użytkownicy mogą potrzebować dodatkowego wsparcia lub bardziej szczegółowych informacji. Dlatego w interfejsie chatbota uwzględniliśmy wygodne linki. Link e-mailowy do wsparcia jest łatwo dostępny, umożliwiając użytkownikom bezpośredni kontakt z naszym zespołem, jeśli potrzebują bardziej spersonalizowanej pomocy. Dodatkowo wprowadziliśmy link do dokumentacji, zapewniając łatwy dostęp do bardziej kompleksowych zasobów dla tych, którzy chcą zgłębić ofertę Blue.

Doświadczenie użytkownika jest dodatkowo wzbogacone przez gustowne animacje fade-in i fade-up podczas otwierania okna chatbota. Te subtelne animacje dodają odrobinę wyrafinowania do interfejsu, sprawiając, że interakcja wydaje się bardziej dynamiczna i angażująca. Wprowadziliśmy również wskaźnik pisania, małą, ale kluczową funkcję, która informuje użytkowników, że chatbot przetwarza ich zapytanie i tworzy odpowiedź. Ten wizualny sygnał pomaga zarządzać oczekiwaniami użytkowników i utrzymuje poczucie aktywnej komunikacji.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>

Zauważając, że niektóre rozmowy mogą wymagać więcej miejsca na ekranie, dodaliśmy możliwość otwierania rozmowy w większym oknie. Ta funkcja jest szczególnie przydatna w przypadku dłuższych wymian lub przeglądania szczegółowych informacji, dając użytkownikom elastyczność dostosowania chatbota do swoich potrzeb.

Za kulisami wprowadziliśmy inteligentne przetwarzanie, aby zoptymalizować odpowiedzi chatbota. Nasz system automatycznie analizuje odpowiedzi AI, aby usunąć odniesienia do naszych wewnętrznych dokumentów, zapewniając, że prezentowane informacje są czyste, istotne i skoncentrowane wyłącznie na odpowiadaniu na zapytanie użytkownika.
Aby poprawić czytelność i umożliwić bardziej zniuansowaną komunikację, wprowadziliśmy wsparcie dla markdowna za pomocą biblioteki 'marked'. Ta funkcja umożliwia naszemu AI dostarczanie bogato formatowanego tekstu, w tym pogrubień i kursyw, uporządkowanych list oraz nawet fragmentów kodu, gdy jest to konieczne. To jak otrzymywanie dobrze sformatowanego, dostosowanego mini-dokumentu w odpowiedzi na Twoje pytania.

Ostatnie, ale z pewnością nie mniej ważne, priorytetem w naszej implementacji było bezpieczeństwo. Używając biblioteki DOMPurify, sanitizujemy HTML generowany z analizy markdowna. Ten kluczowy krok zapewnia, że wszelkie potencjalnie szkodliwe skrypty lub kody są usuwane przed wyświetleniem treści. To nasz sposób na gwarantowanie, że pomocne informacje, które otrzymujesz, są nie tylko informacyjne, ale także bezpieczne do spożycia.


## Przyszłe Rozwój

To dopiero początek, mamy kilka ekscytujących rzeczy na horyzoncie dla tej funkcji. 

Jedną z naszych nadchodzących funkcji jest możliwość strumieniowego przesyłania odpowiedzi w czasie rzeczywistym. Wkrótce zobaczysz, jak odpowiedzi chatbota pojawiają się znak po znaku, co sprawi, że rozmowy będą wydawać się bardziej naturalne i dynamiczne. To jak obserwowanie, jak AI myśli, tworząc bardziej angażujące i interaktywne doświadczenie, które trzyma Cię w pętli na każdym kroku.

Dla naszych cenionych użytkowników Blue pracujemy nad personalizacją. Chatbot rozpozna, kiedy jesteś zalogowany, dostosowując swoje odpowiedzi na podstawie informacji o Twoim koncie, historii użytkowania i preferencjach. Wyobraź sobie chatbota, który nie tylko odpowiada na Twoje pytania, ale rozumie Twój specyficzny kontekst w ekosystemie Blue, oferując bardziej istotną i spersonalizowaną pomoc.

Rozumiemy, że możesz pracować nad wieloma projektami lub mieć różne zapytania. Dlatego rozwijamy możliwość utrzymywania kilku odrębnych wątków rozmów z naszym chatbotem. Ta funkcja pozwoli Ci płynnie przełączać się między różnymi tematami, nie tracąc kontekstu – tak jakbyś miał otwarte wiele kart w przeglądarce.

Aby uczynić Twoje interakcje jeszcze bardziej produktywnymi, tworzymy funkcję, która będzie oferować sugerowane pytania uzupełniające na podstawie Twojej obecnej rozmowy. To pomoże Ci zgłębić tematy głębiej i odkryć powiązane informacje, o które mogłeś nie pomyśleć, aby zapytać, czyniąc każdą sesję czatu bardziej kompleksową i wartościową.

Cieszymy się również na myśl o stworzeniu zestawu wyspecjalizowanych asystentów AI, z których każdy będzie dostosowany do konkretnych potrzeb. Niezależnie od tego, czy szukasz odpowiedzi na pytania przedsprzedażowe, czy chcesz skonfigurować nowy projekt, czy rozwiązać zaawansowane funkcje, będziesz mógł wybrać asystenta, który najlepiej odpowiada Twoim obecnym potrzebom. To jak posiadanie zespołu ekspertów Blue na wyciągnięcie ręki, z których każdy specjalizuje się w różnych aspektach naszej platformy.

Na koniec pracujemy nad umożliwieniem przesyłania zrzutów ekranu bezpośrednio do czatu. AI przeanalizuje obraz i dostarczy wyjaśnienia lub kroki rozwiązywania problemów na podstawie tego, co widzi. Ta funkcja ułatwi uzyskiwanie pomocy w konkretnych problemach, które napotykasz podczas korzystania z Blue, łącząc wizualne informacje z pomocą tekstową.

## Podsumowanie

Mamy nadzieję, że ta głęboka analiza naszego procesu rozwoju chatbota AI dostarczyła cennych informacji na temat naszego myślenia o rozwoju produktu w Blue. Nasza podróż od zidentyfikowania potrzeby chatbota do zbudowania własnego rozwiązania pokazuje, jak podchodzimy do podejmowania decyzji i innowacji.

![](/insights/ai-chatbot-modal.png)

W Blue starannie rozważamy opcje budowania w porównaniu do kupowania, zawsze z myślą o tym, co najlepiej służy naszym użytkownikom i jest zgodne z naszymi długoterminowymi celami. W tym przypadku zidentyfikowaliśmy znaczną lukę na rynku dla opłacalnego, a jednocześnie wizualnie atrakcyjnego chatbota, który mógłby spełnić nasze specyficzne potrzeby. Chociaż generalnie opowiadamy się za wykorzystywaniem istniejących rozwiązań zamiast wynajdywaniem koła na nowo, czasami najlepszą drogą naprzód jest stworzenie czegoś dostosowanego do Twoich unikalnych wymagań.

Nasza decyzja o zbudowaniu własnego chatbota nie została podjęta lekko. Była wynikiem dokładnych badań rynkowych, jasnego zrozumienia naszych potrzeb i zobowiązania do zapewnienia jak najlepszego doświadczenia dla naszych użytkowników. Dzięki rozwojowi wewnętrznemu mogliśmy stworzyć rozwiązanie, które nie tylko spełnia nasze obecne potrzeby, ale także kładzie podwaliny pod przyszłe ulepszenia i integracje.

Ten projekt ilustruje nasze podejście w Blue: nie boimy się zakasać rękawów i zbudować coś od podstaw, gdy jest to właściwy wybór dla naszego produktu i naszych użytkowników. To gotowość do podjęcia dodatkowego wysiłku pozwala nam dostarczać innowacyjne rozwiązania, które naprawdę spełniają potrzeby naszych klientów.
Cieszymy się na przyszłość naszego chatbota AI i wartość, jaką przyniesie zarówno potencjalnym, jak i obecnym użytkownikom Blue. W miarę jak będziemy kontynuować udoskonalanie i rozszerzanie jego możliwości, pozostajemy zobowiązani do przesuwania granic tego, co możliwe w zarządzaniu projektami i interakcji z klientami.

Dziękujemy za dołączenie do nas w tej podróży przez nasz proces rozwoju. Mamy nadzieję, że dało to Wam wgląd w przemyślane, zorientowane na użytkownika podejście, które stosujemy w każdym aspekcie Blue. Bądźcie na bieżąco z kolejnymi aktualizacjami, gdy będziemy kontynuować rozwój i ulepszanie naszej platformy, aby lepiej służyć Wam.

Jeśli jesteś zainteresowany, możesz znaleźć link do kodu źródłowego tego projektu tutaj:

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)**: To komponent Vue, który napędza sam widget czatu. 
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)**: To middleware, które działa pomiędzy komponentem czatu a API OpenAI Assistants.