---
title:  Skalning av CSV-importer och -exporter till 250 000+ poster
description: Upptäck hur Blue skalade CSV-importer och -exporter 10x med hjälp av Rust och skalbar arkitektur samt strategiska teknologival i B2B SaaS.
category: "Engineering"
date: 2024-07-18
---


På Blue [pressar vi ständigt gränserna](/platform/roadmap) för vad som är möjligt inom projektledningsprogramvara. Under åren har vi [släppt hundratals funktioner](/platform/changelog).

Vår senaste ingenjörsbedrift? 

En fullständig översyn av vårt [CSV-import](https://documentation.blue.cc/integrations/csv-import) och [export](https://documentation.blue.cc/integrations/csv-export) system, vilket dramatiskt förbättrade prestanda och skalbarhet. 

Detta inlägg tar dig bakom kulisserna av hur vi tacklade denna utmaning, de teknologier vi använde och de imponerande resultat vi uppnådde.

Det mest intressanta här är att vi var tvungna att gå utanför vår typiska [teknologistack](https://sop.blue.cc/product/technology-stack) för att uppnå de resultat vi ville ha. Detta är ett beslut som måste fattas genomtänkt, eftersom de långsiktiga konsekvenserna kan vara allvarliga när det gäller teknologisk skuld och långsiktig underhållsöverhead. 

<video autoplay loop muted playsinline>
  <source src="/videos/import-export-video.mp4" type="video/mp4">
</video>

## Skalning för företagsbehov

Vår resa började med en begäran från en företagskund inom evenemangsbranschen. Denna kund använder Blue som sin centrala hubb för att hantera stora listor av evenemang, platser och talare, och integrerar det sömlöst med sin webbplats. 

För dem är Blue inte bara ett verktyg — det är den enda sanningskällan för hela deras verksamhet.

Även om vi alltid är stolta över att höra att kunder använder oss för sådana kritiska behov, finns det också ett stort ansvar på vår sida att säkerställa ett snabbt och pålitligt system.

När denna kund skalade sina operationer stötte de på ett betydande hinder: **importera och exportera stora CSV-filer som innehåller 100 000 till 200 000+ poster.**

Detta var bortom kapabiliteten hos vårt system vid den tiden. Faktum är att vårt tidigare import/export-system redan hade problem med importer och exporter som innehöll mer än 10 000 till 20 000 poster! Så 200 000+ poster var uteslutet. 

Användare upplevde frustrerande långa väntetider, och i vissa fall skulle importer eller exporter *inte slutföras alls.* Detta påverkade deras verksamhet avsevärt eftersom de var beroende av dagliga importer och exporter för att hantera vissa aspekter av sin verksamhet. 

> Multi-tenancy är en arkitektur där en enda instans av programvara betjänar flera kunder (hyresgäster). Även om det är effektivt, kräver det noggrant resursförvaltning för att säkerställa att en hyresgästs åtgärder inte negativt påverkar andra.

Och denna begränsning påverkade inte bara denna specifika kund. 

På grund av vår multi-tenant-arkitektur—där flera kunder delar samma infrastruktur—kunde en enda resurskrävande import eller export potentiellt sakta ner operationerna för andra användare, vilket i praktiken ofta hände. 

Som vanligt gjorde vi en analys av bygga vs köpa, för att förstå om vi skulle spendera tid på att uppgradera vårt eget system eller köpa ett system från någon annan. Vi tittade på olika möjligheter.

Leverantören som verkligen stack ut var en SaaS-leverantör som heter [Flatfile](https://flatfile.com/). Deras system och kapabiliteter såg ut att vara exakt vad vi behövde. 

Men efter att ha granskat deras [priser](https://flatfile.com/pricing/) beslutade vi att detta skulle bli en extremt dyr lösning för en applikation av vår skala — *$2/fil börjar snabbt bli mycket!* — och det var bättre att utöka vår inbyggda CSV-import/export-motor. 

För att tackla denna utmaning fattade vi ett djärvt beslut: att introducera Rust i vår primära Javascript-teknologistack. Detta systemprogrammeringsspråk, känt för sin prestanda och säkerhet, var det perfekta verktyget för våra prestandakritiska CSV-parsing och datakartläggningsbehov.

Så här närmade vi oss lösningen.

### Introduktion av bakgrundstjänster

Grunden för vår lösning var introduktionen av bakgrundstjänster för att hantera resurskrävande uppgifter. Denna metod gjorde det möjligt för oss att avlasta tung bearbetning från vår huvudserver, vilket betydligt förbättrade den övergripande systemprestandan.
Vår arkitektur för bakgrundstjänster är utformad med skalbarhet i åtanke. Precis som alla komponenter i vår infrastruktur, auto-skalas dessa tjänster baserat på efterfrågan. 

Detta innebär att under högtrafikperioder, när flera stora importer eller exporter bearbetas samtidigt, tilldelar systemet automatiskt fler resurser för att hantera den ökade belastningen. Omvänt, under lugnare perioder, skalas det ner för att optimera resursanvändningen.

Denna skalbara bakgrundstjänstarkitektur har gynnat Blue inte bara för CSV-importer och -exporter. Med tiden har vi flyttat ett betydande antal funktioner till bakgrundstjänster för att avlasta våra huvudservrar:

- **[Formelberäkningar](https://documentation.blue.cc/custom-fields/formula)**: Avlastar komplexa matematiska operationer för att säkerställa snabba uppdateringar av härledda fält utan att påverka huvudserverns prestanda.
- **[Dashboard/Diagram](/platform/features/dashboards)**: Bearbetar stora datamängder i bakgrunden för att generera aktuella visualiseringar utan att sakta ner användargränssnittet.
- **[Sökindex](https://documentation.blue.cc/projects/search)**: Uppdaterar kontinuerligt sökindexet i bakgrunden, vilket säkerställer snabba och exakta sökresultat utan att påverka systemets prestanda.
- **[Kopiera projekt](https://documentation.blue.cc/projects/copying-projects)**: Hanterar replikeringen av stora, komplexa projekt i bakgrunden, vilket gör att användare kan fortsätta arbeta medan kopian skapas.
- **[Automatisering av projektledning](/platform/features/automations)**: Utför användardefinierade automatiserade arbetsflöden i bakgrunden, vilket säkerställer tidsenliga åtgärder utan att blockera andra operationer.
- **[Upprepande poster](https://documentation.blue.cc/records/repeat)**: Genererar återkommande uppgifter eller händelser i bakgrunden, vilket upprätthåller schemats noggrannhet utan att belasta huvudapplikationen.
- **[Anpassade fält för tidsduration](https://documentation.blue.cc/custom-fields/duration)**: Beräknar och uppdaterar kontinuerligt tidsdifferensen mellan två händelser i Blue, vilket ger realtidsdata om duration utan att påverka systemets responsivitet.

## Ny Rust-modul för dataparsering

Kärnan i vår CSV-bearbetningslösning är en anpassad Rust-modul. Även om detta markerade vårt första steg utanför vår kärnteknologistack av Javascript, drevs beslutet att använda Rust av dess exceptionella prestanda i samtidiga operationer och filbearbetningsuppgifter.

Rusts styrkor passar perfekt med kraven på CSV-parsing och datakartläggning. Dess kostnadsfria abstraktioner möjliggör hög nivå programmering utan att offra prestanda, medan dess ägarmodell säkerställer minnessäkerhet utan behov av skräpsamling. Dessa funktioner gör Rust särskilt skicklig på att hantera stora datamängder effektivt och säkert.

För CSV-parsing utnyttjade vi Rusts csv crate, som erbjuder högpresterande läsning och skrivning av CSV-data. Vi kombinerade detta med anpassad datakartläggningslogik för att säkerställa sömlös integration med Blues datastrukturer.

Inlärningskurvan för Rust var brant men hanterbar. Vårt team ägnade cirka två veckor åt intensiv inlärning för detta.

Förbättringarna var imponerande:

![](/insights/import-export.png)

Vårt nya system kan bearbeta samma mängd poster som vårt gamla system kunde bearbeta på 15 minuter på cirka 30 sekunder. 

## Webbserver och databasinteraktion

För webbserverkomponenten av vår Rust-implementering valde vi Rocket som vårt ramverk. Rocket stack ut för sin kombination av prestanda och utvecklarvänliga funktioner. Dess statiska typning och kompileringstidkontroll passar bra med Rusts säkerhetsprinciper, vilket hjälper oss att fånga potentiella problem tidigt i utvecklingsprocessen.
På databasfronten valde vi SQLx. Detta asynkrona SQL-bibliotek för Rust erbjuder flera fördelar som gjorde det idealiskt för våra behov:

- Typ-säker SQL: SQLx gör det möjligt för oss att skriva rå SQL med kompileringstidkontrollerade frågor, vilket säkerställer typ-säkerhet utan att offra prestanda.
- Asynkron support: Detta passar bra med Rocket och vårt behov av effektiva, icke-blockerande databasoperationer.
- Databasagnostisk: Även om vi främst använder [AWS Aurora](https://aws.amazon.com/rds/aurora/), som är MySQL-kompatibel, ger SQLx:s stöd för flera databaser oss flexibilitet för framtiden om vi någonsin beslutar att byta. 

## Optimering av batchning

Vår resa till den optimala batchkonfigurationen var en av rigorös testning och noggrann analys. Vi körde omfattande prestandatester med olika kombinationer av samtidiga transaktioner och chunk-storlekar, och mätte inte bara rå hastighet utan även resursutnyttjande och systemstabilitet.

Processen involverade att skapa testdatamängder av varierande storlek och komplexitet, som simulerade verkliga användningsmönster. Vi körde sedan dessa datamängder genom vårt system, justerade antalet samtidiga transaktioner och chunk-storleken för varje körning.

Efter att ha analyserat resultaten fann vi att bearbetning av 5 samtidiga transaktioner med en chunk-storlek på 500 poster gav den bästa balansen mellan hastighet och resursutnyttjande. Denna konfiguration gör att vi kan upprätthålla hög genomströmning utan att överbelasta vår databas eller konsumera överdriven minne.

Intressant nog fann vi att ökad samtidighet bortom 5 transaktioner inte gav betydande prestandavinster och ibland ledde till ökad databasbelastning. På samma sätt förbättrade större chunk-storlekar rå hastighet men till priset av högre minnesanvändning och längre svarstider för små till medelstora importer/exporter.

## CSV-exporter via e-postlänkar

Den sista delen av vår lösning adresserar utmaningen att leverera stora exporterade filer till användare. Istället för att tillhandahålla en direkt nedladdning från vår webbapp, vilket skulle kunna leda till timeout-problem och ökad serverbelastning, implementerade vi ett system med e-postade nedladdningslänkar.

När en användare initierar en stor export bearbetar vårt system begäran i bakgrunden. När den är klar, istället för att hålla anslutningen öppen eller lagra filen på våra webbservrar, laddar vi upp filen till en säker, temporär lagringsplats. Vi genererar sedan en unik, säker nedladdningslänk och e-postar den till användaren.

Dessa nedladdningslänkar är giltiga i 2 timmar, vilket ger en balans mellan användarvänlighet och informationssäkerhet. Denna tidsram ger användare gott om tid att hämta sina data samtidigt som den säkerställer att känslig information inte lämnas tillgänglig på obestämd tid.

Säkerheten för dessa nedladdningslänkar var en hög prioritet i vår design. Varje länk är:

- Unik och slumpmässigt genererad, vilket gör det praktiskt taget omöjligt att gissa
- Giltig i endast 2 timmar
- Krypterad under överföring, vilket säkerställer säkerheten för data när den laddas ner

Denna metod erbjuder flera fördelar:

- Den minskar belastningen på våra webbservrar, eftersom de inte behöver hantera stora filnedladdningar direkt
- Den förbättrar användarupplevelsen, särskilt för användare med långsammare internetanslutningar som kan möta timeout-problem med direkta nedladdningar
- Den ger en mer pålitlig lösning för mycket stora exporter som kan överskrida typiska webbtimeoutgränser

Användarfeedback på denna funktion har varit överväldigande positiv, med många som uppskattar den flexibilitet den erbjuder i hanteringen av stora dataexporter.

## Exportera filtrerad data

Den andra uppenbara förbättringen var att tillåta användare att endast exportera data som redan var filtrerad i deras projektvy. Detta innebär att om det finns en aktiv tagg "prioritet", så skulle endast poster som har denna tagg hamna i CSV-exporten. Detta innebär mindre tid att manipulera data i Excel för att filtrera bort sådant som inte är viktigt, och hjälper oss också att minska antalet rader att bearbeta.

## Framtidsutsikter

Även om vi inte har några omedelbara planer på att utöka vår användning av Rust, har detta projekt visat oss potentialen för denna teknologi för prestandakritiska operationer. Det är ett spännande alternativ vi nu har i vår verktygslåda för framtida optimeringsbehov. Denna översyn av CSV-import och -export stämmer perfekt överens med Blues åtagande för skalbarhet. 

Vi är dedikerade till att tillhandahålla en plattform som växer med våra kunder, och hanterar deras växande databehov utan att kompromissa med prestanda.

Beslutet att introducera Rust i vår teknologistack togs inte lättvindigt. Det väckte en viktig fråga som många ingenjörsteam står inför: När är det lämpligt att ge sig utanför din kärnteknologistack, och när bör du hålla dig till bekanta verktyg?

Det finns inget universellt svar, men på Blue har vi utvecklat en ram för att fatta dessa avgörande beslut:

- **Problem-först tillvägagångssätt:** Vi börjar alltid med att tydligt definiera problemet vi försöker lösa. I detta fall behövde vi dramatiskt förbättra prestandan för CSV-importer och -exporter för stora datamängder.
- **Utnyttja befintliga lösningar:** Innan vi tittar utanför vår kärnstack, utforskar vi grundligt vad som kan uppnås med våra befintliga teknologier. Detta involverar ofta profilering, optimering och att tänka om vår metod inom bekanta begränsningar.
- **Kvantifiera den potentiella vinsten:** Om vi överväger en ny teknologi, måste vi kunna tydligt formulera och, helst, kvantifiera fördelarna. För vårt CSV-projekt förutspådde vi förbättringar i bearbetningshastighet med en ordning av storleksordning.
- **Bedöma kostnaderna:** Att introducera en ny teknologi handlar inte bara om det omedelbara projektet. Vi överväger de långsiktiga kostnaderna:
  - Inlärningskurva för teamet
  - Löpande underhåll och support
  - Potentiella komplikationer i distribution och drift
  - Påverkan på rekrytering och teamkomposition
- **Inneslutning och integration:** Om vi introducerar en ny teknologi, syftar vi till att begränsa den till en specifik, väldefinierad del av vårt system. Vi säkerställer också att vi har en tydlig plan för hur den kommer att integreras med vår befintliga stack.
- **Framtidssäkring:** Vi överväger om detta teknologival öppnar upp framtida möjligheter eller om det kan måla in oss i ett hörn.

En av de primära riskerna med att ofta anta nya teknologier är att hamna i vad vi kallar en *"teknologizoo"* - ett fragmenterat ekosystem där olika delar av din applikation är skrivna i olika språk eller ramverk, vilket kräver en bred uppsättning specialiserade färdigheter för att underhålla.


## Slutsats

Detta projekt exemplifierar Blues tillvägagångssätt till ingenjörskonst: *vi är inte rädda för att gå utanför vår komfortzon och anta nya teknologier när det innebär att leverera en betydligt bättre upplevelse för våra användare.* 

Genom att omforma vår CSV-import och -exportprocess har vi inte bara löst ett pressande behov för en företagskund utan förbättrat upplevelsen för alla våra användare som hanterar stora datamängder.

När vi fortsätter att pressa gränserna för vad som är möjligt inom [projektledningsprogramvara](/solutions/use-case/project-management), ser vi fram emot att ta oss an fler utmaningar som denna. 

Håll utkik efter fler [djupdykningar i ingenjörskonsten som driver Blue!](/insights/engineering-blog)