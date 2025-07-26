---
title: AI-autokategorisering (teknisk djupdykning)
category: "Engineering"
description: Följ med Blue-ingenjörsteamet bakom kulisserna när de förklarar hur de byggde en AI-driven funktion för autokategorisering och taggning.
date: 2024-12-07
---

Vi släppte nyligen [AI-autokategorisering](/insights/ai-auto-categorization) till alla Blue-användare. Detta är en AI-funktion som ingår i grundabonnemanget för Blue, utan extra kostnader. I det här inlägget dyker vi djupare in i tekniken bakom denna funktion.

---
På Blue grundar sig vår utvecklingsfilosofi i en djup förståelse för användarbehov och marknadstrender, kombinerat med ett åtagande att behålla den enkelhet och användarvänlighet som definierar vår plattform. Detta driver vår [färdplan](/platform/roadmap) och har [gjort det möjligt för oss att konsekvent leverera funktioner varje månad i flera år](/platform/changelog).

Införandet av AI-driven autotaggning i Blue är ett perfekt exempel på denna filosofi i praktiken. Innan vi dyker in i de tekniska detaljerna om hur vi byggde denna funktion är det avgörande att förstå problemet vi löste och de noggranna överväganden som gjordes under utvecklingen.

Projektledningslandskapet utvecklas snabbt, där AI-funktioner blir allt mer centrala för användarnas förväntningar. Våra kunder, särskilt de som hanterar storskaliga [projekt](/platform) med miljontals [poster](/platform/features/records), hade varit tydliga med sin önskan om smartare, mer effektiva sätt att organisera och kategorisera sin data.

Men på Blue lägger vi inte bara till funktioner för att de är trendiga eller efterfrågade. Vår filosofi är att varje nytt tillägg måste bevisa sitt värde, med standardsvaret som ett bestämt *"nej"* tills en funktion visar stark efterfrågan och tydlig nytta.

För att verkligen förstå problemets djup och potentialen med AI-autotaggning genomförde vi omfattande kundintervjuer, med fokus på långtidsanvändare som hanterar komplexa, datarika projekt över flera domäner.

Dessa samtal avslöjade en gemensam tråd: *medan taggning var ovärderlig för organisation och sökbarhet, hade den manuella processen blivit en flaskhals, särskilt för team som hanterar stora volymer poster.*

Men vi såg bortom att bara lösa det omedelbara problemet med manuell taggning.

Vi föreställde oss en framtid där AI-driven taggning kunde bli grunden för mer intelligenta, automatiserade arbetsflöden.

Den verkliga kraften i denna funktion, insåg vi, låg i dess potentiella integration med vårt [automationssystem för projektledning](/platform/features/automations). Föreställ dig ett projektledningsverktyg som inte bara kategoriserar information intelligent utan också använder dessa kategorier för att dirigera uppgifter, utlösa åtgärder och anpassa arbetsflöden i realtid.

Denna vision stämde perfekt överens med vårt mål att hålla Blue enkelt men kraftfullt.

Dessutom såg vi potentialen att utvidga denna förmåga bortom gränserna för vår plattform. Genom att utveckla ett robust AI-taggningssystem lade vi grunden för ett "kategoriserings-API" som kunde fungera direkt ur lådan, vilket potentiellt öppnar nya vägar för hur våra användare interagerar med och utnyttjar Blue i sina bredare tekosystem.

Denna funktion handlade därför inte bara om att lägga till en AI-kryssruta i vår funktionslista.

Det handlade om att ta ett betydande steg mot en mer intelligent, anpassningsbar projektledningsplattform samtidigt som vi förblev trogna vår kärnfilosofi om enkelhet och användarcentrering.

I följande avsnitt kommer vi att dyka in i de tekniska utmaningar vi mötte när vi förverkligade denna vision, arkitekturen vi designade för att stödja den och lösningarna vi implementerade. Vi kommer också att utforska de framtida möjligheter som denna funktion öppnar upp, och visa hur ett noggrant övervägt tillägg kan bana väg för transformativa förändringar inom projektledning.

---
## Problemet

Som diskuterat ovan kan manuell taggning av projektposter vara tidskrävande och inkonsekvent.

Vi satte oss för att lösa detta genom att utnyttja AI för att automatiskt föreslå taggar baserat på postinnehåll.

De huvudsakliga utmaningarna var:

1. Att välja en lämplig AI-modell
2. Effektivt bearbeta stora volymer poster
3. Säkerställa datasekretess och säkerhet
4. Integrera funktionen sömlöst i vår befintliga arkitektur

## Val av AI-modell

Vi utvärderade flera AI-plattformar, inklusive [OpenAI](https://openai.com), open source-modeller på [HuggingFace](https://huggingface.co/) och [Replicate](https://replicate.com).

Våra kriterier inkluderade:

- Kostnadseffektivitet
- Noggrannhet i att förstå sammanhang
- Förmåga att följa specifika utdataformat
- Garantier för datasekretess

Efter grundlig testning valde vi OpenAI:s [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo). Medan [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) kan erbjuda marginella förbättringar i noggrannhet visade våra tester att GPT-3.5:s prestanda var mer än tillräcklig för våra autotaggningsbehov. Balansen mellan kostnadseffektivitet och starka kategoriseringsförmågor gjorde GPT-3.5 till det idealiska valet för denna funktion.

Den högre kostnaden för GPT-4 skulle ha tvingat oss att erbjuda funktionen som ett betalt tillägg, vilket strider mot vårt mål att **inkludera AI i vår huvudprodukt utan extra kostnad för slutanvändare.**

Vid tidpunkten för vår implementation är prissättningen för GPT-3.5 Turbo:

- $0.0005 per 1K indata-tokens (eller $0.50 per 1M indata-tokens)
- $0.0015 per 1K utdata-tokens (eller $1.50 per 1M utdata-tokens)

Låt oss göra några antaganden om en genomsnittlig post i Blue:

- **Titel**: ~10 tokens
- **Beskrivning**: ~50 tokens
- **2 kommentarer**: ~30 tokens vardera
- **5 anpassade fält**: ~10 tokens vardera
- **Listnamn, förfallodatum och annan metadata**: ~20 tokens
- **Systemprompt och tillgängliga taggar**: ~50 tokens

Totala indata-tokens per post: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 tokens

För utdata, låt oss anta i genomsnitt 3 taggar föreslagna per post, vilket kan totalt vara runt 20 utdata-tokens inklusive JSON-formatering.

För 1 miljon poster:

- Indatakostnad: (240 * 1,000,000 / 1,000,000) * $0.50 = $120
- Utdatakostnad: (20 * 1,000,000 / 1,000,000) * $1.50 = $30

**Total kostnad för autotaggning av 1 miljon poster: $120 + $30 = $150**

## GPT3.5 Turbo-prestanda

Kategorisering är en uppgift som stora språkmodeller (LLMs) som GPT-3.5 Turbo utmärker sig i, vilket gör dem särskilt lämpade för vår autotaggningsfunktion. LLMs är tränade på enorma mängder textdata, vilket gör att de kan förstå sammanhang, semantik och relationer mellan koncept. Denna breda kunskapsbas gör det möjligt för dem att utföra kategoriseringsuppgifter med hög noggrannhet över ett brett spektrum av domäner.

För vårt specifika användningsfall med projektledningstaggning visar GPT-3.5 Turbo flera viktiga styrkor:

- **Kontextuell förståelse:** Kan förstå det övergripande sammanhanget för en projektpost och ta hänsyn till inte bara enskilda ord utan meningen som förmedlas av hela beskrivningen, kommentarer och andra fält.
- **Flexibilitet:** Kan anpassa sig till olika projekttyper och branscher utan att kräva omfattande omprogrammering.
- **Hantering av tvetydighet:** Kan väga flera faktorer för att fatta nyanserade beslut.
- **Lärande från exempel:** Kan snabbt förstå och tillämpa nya kategoriseringsscheman utan ytterligare träning.
- **Fleretikettsklassificering:** Kan föreslå flera relevanta taggar för en enskild post, vilket var avgörande för våra krav.

GPT-3.5 Turbo utmärkte sig också för sin tillförlitlighet i att följa vårt krävda JSON-utdataformat, vilket var *avgörande* för sömlös integration med våra befintliga system. Open source-modeller, även om de är lovande, lade ofta till extra kommentarer eller avvek från det förväntade formatet, vilket skulle ha krävt ytterligare efterbearbetning. Denna konsistens i utdataformat var en nyckelfaktor i vårt beslut, eftersom det avsevärt förenklade vår implementation och minskade potentiella felpunkter.

Att välja GPT-3.5 Turbo med dess konsekventa JSON-utdata gjorde det möjligt för oss att implementera en mer enkel, pålitlig och underhållbar lösning.

Hade vi valt en modell med mindre tillförlitlig formatering skulle vi ha mött en kaskad av komplikationer: behovet av robust parsningslogik för att hantera olika utdataformat, omfattande felhantering för inkonsekventa utdata, potentiella prestandaeffekter från ytterligare bearbetning, ökad testkomplexitet för att täcka alla utdatavariationer och en större långsiktig underhållsbörda.

Parsningsfel kan leda till felaktig taggning, vilket negativt påverkar användarupplevelsen. Genom att undvika dessa fallgropar kunde vi fokusera våra ingenjörsinsatser på kritiska aspekter som prestandaoptimering och användargränssnittsdesign, snarare än att brottas med oförutsägbara AI-utdata.

## Systemarkitektur

Vår AI-autotaggningsfunktion bygger på en robust, skalbar arkitektur designad för att hantera stora volymer förfrågningar effektivt samtidigt som den ger en sömlös användarupplevelse. Som med alla våra system har vi arkitekterat denna funktion för att stödja en storleksordning mer trafik än vi för närvarande upplever. Detta tillvägagångssätt, även om det verkar överdesignat för nuvarande behov, är en bästa praxis som gör att vi sömlöst kan hantera plötsliga trafikspikar och ger oss gott om utrymme för tillväxt utan större arkitektoniska omarbetningar. Annars skulle vi behöva omdesigna alla våra system var 18:e månad — något vi lärt oss den hårda vägen tidigare!

Låt oss bryta ner komponenterna och flödet i vårt system:

- **Användarinteraktion:** Processen börjar när en användare trycker på "Autotagga"-knappen i Blue-gränssnittet. Denna åtgärd utlöser autotaggningsarbetsflödet.
- **Blue API-anrop:** Användarens åtgärd översätts till ett API-anrop till vår Blue-backend. Denna API-endpoint är designad för att hantera autotaggningsförfrågningar.
- **Köhantering:** Istället för att bearbeta förfrågan omedelbart, vilket kan leda till prestandaproblem under hög belastning, lägger vi till taggningsförfrågan i en kö. Vi använder Redis för denna kömekanism, vilket gör att vi kan hantera belastning effektivt och säkerställa systemskalbarhet.
- **Bakgrundstjänst:** Vi implementerade en bakgrundstjänst som kontinuerligt övervakar kön för nya förfrågningar. Denna tjänst ansvarar för att bearbeta köade förfrågningar.
- **OpenAI API-integration:** Bakgrundstjänsten förbereder nödvändig data och gör API-anrop till OpenAI:s GPT-3.5-modell. Det är här den faktiska AI-drivna taggningen sker. Vi skickar relevant projektdata och får föreslagna taggar i retur.
- **Resultatbearbetning:** Bakgrundstjänsten bearbetar resultaten från OpenAI. Detta innebär att tolka AI:s svar och förbereda data för tillämpning på projektet.
- **Taggtillämpning:** De bearbetade resultaten används för att tillämpa de nya taggarna på relevanta objekt i projektet. Detta steg uppdaterar vår databas med AI-föreslagna taggar.
- **Uppdatering av användargränssnitt:** Slutligen dyker de nya taggarna upp i användarens projektvy, vilket fullbordar autotaggningsprocessen från användarens perspektiv.

Denna arkitektur erbjuder flera viktiga fördelar som förbättrar både systemprestanda och användarupplevelse. Genom att använda en kö och bakgrundsbearbetning har vi uppnått imponerande skalbarhet, vilket gör att vi kan hantera många förfrågningar samtidigt utan att överväldiga vårt system eller nå hastighetsgränserna för OpenAI API. Implementering av denna arkitektur krävde noggrann hänsyn till olika faktorer för att säkerställa optimal prestanda och tillförlitlighet. För köhantering valde vi Redis, vilket utnyttjar dess hastighet och tillförlitlighet vid hantering av distribuerade köer.

Detta tillvägagångssätt bidrar också till funktionens övergripande responsivitet. Användare får omedelbar feedback om att deras förfrågan bearbetas, även om själva taggningen tar lite tid, vilket skapar en känsla av realtidsinteraktion. Arkitekturens feltolerans är en annan avgörande fördel. Om någon del av processen stöter på problem, såsom tillfälliga störningar i OpenAI API, kan vi elegant försöka igen eller hantera felet utan att påverka hela systemet.

Denna robusthet, kombinerad med taggarnas realtidsvisning, förbättrar användarupplevelsen och ger intrycket av AI-"magi" i arbete.

## Data och prompter

Ett avgörande steg i vår autotaggningsprocess är att förbereda data som ska skickas till GPT-3.5-modellen. Detta steg krävde noggrann övervägning för att balansera tillhandahållande av tillräcklig kontext för korrekt taggning samtidigt som effektivitet bibehölls och användarnas integritet skyddades. Här är en detaljerad titt på vår dataförberedningsprocess.

För varje post sammanställer vi följande information:

- **Listnamn**: Ger kontext om den bredare kategorin eller fasen av projektet.
- **Posttitel**: Innehåller ofta nyckelinformation om postens syfte eller innehåll.
- **Anpassade fält**: Vi inkluderar text- och nummerbaserade [anpassade fält](/platform/features/custom-fields), som ofta innehåller avgörande projektspecifik information.
- **Beskrivning**: Innehåller vanligtvis den mest detaljerade informationen om posten.
- **Kommentarer**: Kan ge ytterligare kontext eller uppdateringar som kan vara relevanta för taggning.
- **Förfallodatum**: Tidsmässig information som kan påverka taggval.

Intressant nog skickar vi inte befintliga taggdata till GPT-3.5, och vi gör detta för att undvika att förfördelning av modellen.

Kärnan i vår autotaggningsfunktion ligger i hur vi interagerar med GPT-3.5-modellen och bearbetar dess svar. Denna del av vår pipeline krävde noggrann design för att säkerställa korrekt, konsekvent och effektiv taggning.

Vi använder en noggrant utformad systemprompt för att instruera AI om dess uppgift. Här är en uppdelning av vår prompt och logiken bakom varje komponent:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Uppgiftsdefinition:** Vi anger tydligt AI:s uppgift för att säkerställa fokuserade svar.
- **Osäkerhetshantering:** Vi tillåter uttryckligen tomma svar, vilket förhindrar tvingad taggning när AI är osäker.
- **Tillgängliga taggar:** Vi tillhandahåller en lista över giltiga taggar (${tags}) för att begränsa AI:s val till befintliga projekttaggar.
- **Aktuellt datum:** Att inkludera ${currentDate} hjälper AI att förstå den tidsmässiga kontexten, vilket kan vara avgörande för vissa typer av projekt.
- **Svarsformat:** Vi specificerar ett JSON-format för enkel tolkning och felkontroll.

Denna prompt är resultatet av omfattande testning och iteration. Vi fann att att vara explicit om uppgiften, tillgängliga alternativ och önskat utdataformat avsevärt förbättrade noggrannheten och konsekvensen i AI:s svar — enkelhet är nyckeln!

Listan över tillgängliga taggar genereras på serversidan och valideras innan inkludering i prompten. Vi implementerar strikta teckengränser på taggnamn för att förhindra överdimensionerade prompter.

Som nämnts ovan hade vi inga problem med GPT-3.5 Turbo att få tillbaka det rena JSON-svaret i rätt format 100% av tiden.

Så sammanfattningsvis,

- Vi kombinerar systemprompten med förberedd postdata.
- Denna kombinerade prompt skickas till GPT-3.5-modellen via OpenAI:s API.
- Vi använder en temperaturinställning på 0,3 för att balansera kreativitet och konsekvens i AI:s svar.
- Vårt API-anrop inkluderar en max_tokens-parameter för att begränsa svarsstorleken och kontrollera kostnader.

När vi får AI:s svar går vi igenom flera steg för att bearbeta och tillämpa de föreslagna taggarna:

* **JSON-tolkning**: Vi försöker tolka svaret som JSON. Om tolkning misslyckas loggar vi felet och hoppar över taggning för den posten.
* **Schemavalidering**: Vi validerar den tolkade JSON mot vårt förväntade schema (ett objekt med en "tags"-array). Detta fångar upp strukturella avvikelser i AI:s svar.
* **Taggvalidering**: Vi korsrefererar de föreslagna taggarna mot vår lista över giltiga projekttaggar. Detta steg filtrerar bort taggar som inte finns i projektet, vilket kan inträffa om AI missförstod eller om projekttaggar ändrades mellan promptskapande och svarsbearbetning.
* **Borttagning av dubbletter**: Vi tar bort eventuella dubblettaggar från AI:s förslag för att undvika redundant taggning.
* **Tillämpning**: De validerade och avdubblettaggarna tillämpas sedan på posten i vår databas.
* **Loggning och analys**: Vi loggar de slutliga tillämpade taggarna. Denna data är värdefull för att övervaka systemets prestanda och förbättra det över tid.

## Utmaningar

Att implementera AI-driven autotaggning i Blue presenterade flera unika utmaningar, som var och en krävde innovativa lösningar för att säkerställa en robust, effektiv och användarvänlig funktion.

### Ångra bulkoperation

AI-taggningsfunktionen kan göras både på enskilda poster såväl som i bulk. Problemet med bulkoperationen är att om användaren inte gillar resultatet skulle de behöva manuellt gå igenom tusentals poster och ångra AI:s arbete. Det är uppenbart oacceptabelt.

För att lösa detta implementerade vi ett innovativt taggningssessionssystem. Varje bulktaggningsoperation tilldelas ett unikt sessions-ID, som är kopplat till alla taggar som tillämpas under den sessionen. Detta gör att vi effektivt kan hantera ångringsoperationer genom att helt enkelt ta bort alla taggar som är kopplade till ett visst sessions-ID. Vi tar också bort relaterade revisionsspår, vilket säkerställer att ångrade operationer inte lämnar några spår i systemet. Detta tillvägagångssätt ger användare förtroendet att experimentera med AI-taggning, med vetskapen om att de enkelt kan återställa ändringar om det behövs.

### Datasekretess

Datasekretess var en annan kritisk utmaning vi mötte.

Våra användare litar på oss med sin projektdata, och det var av yttersta vikt att säkerställa att denna information inte bevarades eller användes för modellträning av OpenAI. Vi tacklade detta på flera fronter.

Först bildade vi ett avtal med OpenAI som uttryckligen förbjuder användning av våra data för modellträning. Dessutom raderar OpenAI data efter bearbetning, vilket ger ett extra lager av integritetsskydd.

På vår sida tog vi försiktighetsåtgärden att utesluta känslig information, såsom tilldelningsdetaljer, från data som skickas till AI så detta säkerställer att specifika individers namn inte skickas till tredje part tillsammans med andra data.

Detta omfattande tillvägagångssätt gör det möjligt för oss att utnyttja AI-funktioner samtidigt som vi upprätthåller de högsta standarderna för datasekretess och säkerhet.

### Hastighetsbegränsningar och felhantering

En av våra främsta bekymmer var skalbarhet och hastighetsbegränsning. Direkta API-anrop till OpenAI för varje post skulle ha varit ineffektivt och kunde snabbt nå hastighetsgränser, särskilt för stora projekt eller under topptider. För att ta itu med detta utvecklade vi en bakgrundstjänstarkitektur som gör att vi kan samla förfrågningar och implementera vårt eget kösystem. Detta tillvägagångssätt hjälper oss att hantera API-anropsfrekvens och möjliggör mer effektiv bearbetning av stora volymer poster, vilket säkerställer smidig prestanda även under hög belastning.

Naturen av AI-interaktioner innebar att vi också behövde förbereda oss för tillfälliga fel eller oväntade utdata. Det fanns tillfällen då AI kan producera ogiltig JSON eller utdata som inte matchade vårt förväntade format. För att hantera detta implementerade vi robust felhantering och tolkningslogik genom hela vårt system. Om AI-svaret inte är giltig JSON eller inte innehåller den förväntade "tags"-nyckeln är vårt system designat för att behandla det som om inga taggar föreslogs, snarare än att försöka bearbeta potentiellt korrupt data. Detta säkerställer att även inför AI:s oförutsägbarhet förblir vårt system stabilt och tillförlitligt.

## Framtida utveckling

Vi tror att funktioner, och Blue-produkten som helhet, aldrig är "klar" — det finns alltid utrymme för förbättring.

Det fanns några funktioner som vi övervägde i den initiala byggnationen som inte klarade avgränsningsfasen, men som är intressanta att notera eftersom vi sannolikt kommer att implementera någon version av dem i framtiden.

Den första är att lägga till taggbeskrivning. Detta skulle göra det möjligt för slutanvändare att inte bara ge taggar ett namn och en färg, utan också en valfri beskrivning. Detta skulle också skickas till AI för att hjälpa till att ge ytterligare kontext och potentiellt förbättra noggrannheten.

Även om ytterligare kontext kan vara värdefullt är vi medvetna om den potentiella komplexitet det kan introducera. Det finns en delikat balans att hitta mellan att tillhandahålla användbar information och övervälda användare med för mycket detaljer. När vi utvecklar denna funktion kommer vi att fokusera på att hitta den där sweet spot där tillagd kontext förbättrar snarare än komplicerar användarupplevelsen.

Kanske den mest spännande förbättringen på vår horisont är integrationen av AI-autotaggning med vårt [automationssystem för projektledning](/platform/features/automations).

Detta innebär att AI-taggningsfunktionen kan vara antingen en utlösare eller en åtgärd från en automation. Detta kan vara enormt eftersom det kan förvandla denna ganska enkla AI-kategoriseringsfunktion till ett AI-baserat dirigeringssystem för arbete.

Föreställ dig en automation som säger:

När AI taggar en post som "Kritisk" -> Tilldela till "Chef" och Skicka ett anpassat e-postmeddelande

Detta innebär att när du AI-taggar en post, om AI beslutar att det är ett kritiskt problem, kan den automatiskt tilldela projektledaren och skicka dem ett anpassat e-postmeddelande. Detta utökar [fördelarna med vårt automationssystem för projektledning](/platform/features/automations) från ett rent regelbaserat system till ett verkligt flexibelt AI-system.

Genom att kontinuerligt utforska gränserna för AI inom projektledning strävar vi efter att förse våra användare med verktyg som inte bara möter deras nuvarande behov utan förutser och formar framtidens arbete. Som alltid kommer vi att utveckla dessa funktioner i nära samarbete med vår användargemenskap, vilket säkerställer att varje förbättring tillför verkligt, praktiskt värde till projektledningsprocessen.

## Slutsats

Så där var det!

Detta var en rolig funktion att implementera, och ett av våra första steg inom AI, tillsammans med [AI-innehållssammanfattning](/insights/ai-content-summarization) som vi tidigare lanserat. Vi vet att AI kommer att spela en allt större roll inom projektledning i framtiden, och vi kan inte vänta med att lansera fler innovativa funktioner som utnyttjar avancerade LLMs (Large Language Models).

Det var en hel del att tänka på när vi implementerade detta, och vi är särskilt entusiastiska över hur vi kan utnyttja denna funktion i framtiden med Blues befintliga [automationsmotor för projektledning](/insights/benefits-project-management-automation).

Vi hoppas också att det har varit en intressant läsning och att det ger dig en inblick i hur vi tänker kring utvecklingen av de funktioner du använder varje dag.
