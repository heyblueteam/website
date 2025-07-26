---
title:  Schalen van CSV-imports en -exports tot 250.000+ records
description: Ontdek hoe Blue CSV-imports en -exports 10x heeft opgeschaald met behulp van Rust en schaalbare architectuur en strategische technologische keuzes in B2B SaaS.
category: "Engineering"
date: 2024-07-18
---


Bij Blue zijn we [constant de grenzen aan het verleggen](/platform/roadmap) van wat mogelijk is in projectmanagementsoftware. In de loop der jaren hebben we [honderden functies uitgebracht](/platform/changelog).

Onze laatste technische prestatie?

Een complete herziening van ons [CSV-import](https://documentation.blue.cc/integrations/csv-import) en [export](https://documentation.blue.cc/integrations/csv-export) systeem, wat de prestaties en schaalbaarheid dramatisch heeft verbeterd.

Deze post neemt je mee achter de schermen van hoe we deze uitdaging hebben aangepakt, de technologieën die we hebben gebruikt en de indrukwekkende resultaten die we hebben behaald.

Het meest interessante is dat we buiten onze typische [technologiestack](https://sop.blue.cc/product/technology-stack) moesten stappen om de gewenste resultaten te behalen. Dit is een beslissing die zorgvuldig moet worden genomen, omdat de langetermijngevolgen ernstig kunnen zijn in termen van technologische schulden en langdurige onderhoudskosten.

<video autoplay loop muted playsinline>
  <source src="/videos/import-export-video.mp4" type="video/mp4">
</video>

## Schalen voor Enterprise Behoeften

Onze reis begon met een verzoek van een enterprise klant in de evenementenindustrie. Deze klant gebruikt Blue als hun centrale hub voor het beheren van enorme lijsten van evenementen, locaties en sprekers, en integreert het naadloos met hun website.

Voor hen is Blue niet zomaar een tool — het is de enige bron van waarheid voor hun hele operatie.

Hoewel we altijd trots zijn om te horen dat klanten ons gebruiken voor zulke cruciale behoeften, ligt er ook een grote verantwoordelijkheid bij ons om een snel en betrouwbaar systeem te waarborgen.

Toen deze klant hun operaties opschaalde, stuitten ze op een aanzienlijke hindernis: **het importeren en exporteren van grote CSV-bestanden met 100.000 tot 200.000+ records.**

Dit was op dat moment buiten de mogelijkheden van ons systeem. Sterker nog, ons vorige import/export systeem had al moeite met importen en exporten die meer dan 10.000 tot 20.000 records bevatten! Dus 200.000+ records was uitgesloten.

Gebruikers ervaarden frustrerend lange wachttijden, en in sommige gevallen zouden importen of exporten *helemaal niet voltooien.* Dit had een aanzienlijke impact op hun operaties, aangezien ze afhankelijk waren van dagelijkse importen en exporten om bepaalde aspecten van hun operaties te beheren.

> Multi-tenancy is een architectuur waarbij een enkele instantie van software meerdere klanten (huurders) bedient. Hoewel het efficiënt is, vereist het zorgvuldige resourcebeheer om ervoor te zorgen dat de acties van de ene huurder de andere niet negatief beïnvloeden.

En deze beperking had niet alleen invloed op deze specifieke klant.

Vanwege onze multi-tenant architectuur — waarbij meerdere klanten dezelfde infrastructuur delen — kon een enkele resource-intensieve import of export de operaties voor andere gebruikers vertragen, wat in de praktijk vaak gebeurde.

Zoals gebruikelijk hebben we een build vs buy-analyse uitgevoerd om te begrijpen of we de tijd moesten besteden aan het upgraden van ons eigen systeem of een systeem van iemand anders moesten kopen. We hebben verschillende mogelijkheden bekeken.

De leverancier die eruit sprong was een SaaS-provider genaamd [Flatfile](https://flatfile.com/). Hun systeem en mogelijkheden leken precies te zijn wat we nodig hadden.

Maar na het bekijken van hun [prijzen](https://flatfile.com/pricing/) besloten we dat dit uiteindelijk een extreem dure oplossing zou zijn voor een applicatie van onze schaal — *$2/bestand loopt echt snel op!* — en het was beter om onze ingebouwde CSV-import/export-engine uit te breiden.

Om deze uitdaging aan te pakken, hebben we een gedurfde beslissing genomen: introduceer Rust in onze primaire Javascript-technologiestack. Deze systeemprogrammeertaal, bekend om zijn prestaties en veiligheid, was het perfecte hulpmiddel voor onze prestatiekritische CSV-parsing en datamappingbehoeften.

Hier is hoe we de oplossing hebben benaderd.

### Introductie van Achtergrondservices

De basis van onze oplossing was de introductie van achtergrondservices om resource-intensieve taken af te handelen. Deze aanpak stelde ons in staat om zware verwerking van onze hoofdserver af te leiden, wat de algehele systeemprestaties aanzienlijk verbeterde. 
Onze achtergrondservices-architectuur is ontworpen met schaalbaarheid in gedachten. Zoals alle componenten van onze infrastructuur, schalen deze services automatisch op basis van de vraag.

Dit betekent dat tijdens piektijden, wanneer meerdere grote importen of exporten gelijktijdig worden verwerkt, het systeem automatisch meer middelen toewijst om de verhoogde belasting aan te kunnen. Omgekeerd, tijdens rustigere periodes, schaalt het systeem af om het middelengebruik te optimaliseren.

Deze schaalbare achtergrondservice-architectuur heeft Blue niet alleen ten goede gekomen voor CSV-imports en -exports. In de loop der tijd hebben we een aanzienlijk aantal functies naar achtergrondservices verplaatst om de belasting van onze hoofdservers te verlichten:

- **[Formuleberekeningen](https://documentation.blue.cc/custom-fields/formula)**: Verplaatst complexe wiskundige bewerkingen om snelle updates van afgeleide velden te waarborgen zonder de prestaties van de hoofdserver te beïnvloeden.
- **[Dashboard/Grafieken](/platform/features/dashboards)**: Verwerkt grote datasets op de achtergrond om actuele visualisaties te genereren zonder de gebruikersinterface te vertragen.
- **[Zoekindex](https://documentation.blue.cc/projects/search)**: Werkt de zoekindex continu op de achtergrond bij, waardoor snelle en nauwkeurige zoekresultaten worden gegarandeerd zonder de systeemprestaties te beïnvloeden.
- **[Projecten Kopiëren](https://documentation.blue.cc/projects/copying-projects)**: Behandelt de replicatie van grote, complexe projecten op de achtergrond, zodat gebruikers kunnen doorgaan met werken terwijl de kopie wordt gemaakt.
- **[Projectmanagementautomatiseringen](/platform/features/automations)**: Voert door de gebruiker gedefinieerde geautomatiseerde workflows op de achtergrond uit, waardoor tijdige acties worden gegarandeerd zonder andere operaties te blokkeren.
- **[Herhalende Records](https://documentation.blue.cc/records/repeat)**: Genereert terugkerende taken of evenementen op de achtergrond, waardoor de nauwkeurigheid van de planning behouden blijft zonder de hoofdapplicatie te belasten.
- **[Tijdsduur Aangepaste Velden](https://documentation.blue.cc/custom-fields/duration)**: Berekent en werkt continu het tijdsverschil tussen twee evenementen in Blue bij, waardoor realtime duurgegevens worden verstrekt zonder de systeemprestaties te beïnvloeden.

## Nieuwe Rust-module voor Gegevensparsing

Het hart van onze CSV-verwerkingsoplossing is een aangepaste Rust-module. Hoewel dit onze eerste stap buiten onze kerntechnologiestack van Javascript markeerde, was de beslissing om Rust te gebruiken gedreven door de uitzonderlijke prestaties in gelijktijdige bewerkingen en bestandsverwerkingstaken.

De sterke punten van Rust sluiten perfect aan bij de eisen van CSV-parsing en datamapping. De zero-cost abstracties maken hoog-niveau programmeren mogelijk zonder in te boeten op prestaties, terwijl het eigendommodel zorgt voor geheugenveiligheid zonder dat garbage collection nodig is. Deze functies maken Rust bijzonder geschikt voor het efficiënt en veilig verwerken van grote datasets.

Voor CSV-parsing hebben we gebruikgemaakt van de csv crate van Rust, die hoge prestaties biedt bij het lezen en schrijven van CSV-gegevens. We hebben dit gecombineerd met aangepaste datamappinglogica om een naadloze integratie met de datastructuren van Blue te waarborgen.

De leercurve voor Rust was steil maar beheersbaar. Ons team heeft ongeveer twee weken besteed aan intensieve training hiervoor.

De verbeteringen waren indrukwekkend:

![](/insights/import-export.png)

Ons nieuwe systeem kan dezelfde hoeveelheid records verwerken die ons oude systeem in 15 minuten kon verwerken in ongeveer 30 seconden.

## Webserver en Database-interactie

Voor de webservercomponent van onze Rust-implementatie kozen we Rocket als ons framework. Rocket viel op vanwege de combinatie van prestaties en ontwikkelaarsvriendelijke functies. De statische typing en compile-tijd controle sluiten goed aan bij de veiligheidsprincipes van Rust, waardoor we potentiële problemen vroeg in het ontwikkelingsproces kunnen opsporen. 
Aan de databasekant kozen we voor SQLx. Deze async SQL-bibliotheek voor Rust biedt verschillende voordelen die het ideaal maakten voor onze behoeften:

- Type-veilige SQL: SQLx stelt ons in staat om ruwe SQL te schrijven met compile-tijd gecontroleerde queries, waardoor typeveiligheid wordt gegarandeerd zonder in te boeten op prestaties.
- Async-ondersteuning: Dit sluit goed aan bij Rocket en onze behoefte aan efficiënte, niet-blokkerende databasebewerkingen.
- Database-agnostisch: Hoewel we voornamelijk [AWS Aurora](https://aws.amazon.com/rds/aurora/) gebruiken, dat MySQL-compatibel is, biedt de ondersteuning van SQLx voor meerdere databases ons flexibiliteit voor de toekomst, voor het geval we ooit besluiten te veranderen.

## Optimalisatie van Batching

Onze reis naar de optimale batchingconfiguratie was er een van rigoureuze tests en zorgvuldige analyse. We hebben uitgebreide benchmarks uitgevoerd met verschillende combinaties van gelijktijdige transacties en chunk-groottes, waarbij we niet alleen de ruwe snelheid, maar ook het middelengebruik en de systeemstabiliteit hebben gemeten.

Het proces omvatte het creëren van testdatasets van verschillende groottes en complexiteit, waarbij we realistische gebruikspatronen simuleerden. We hebben deze datasets vervolgens door ons systeem gehaald, waarbij we het aantal gelijktijdige transacties en de chunk-grootte voor elke run hebben aangepast.

Na het analyseren van de resultaten ontdekten we dat het verwerken van 5 gelijktijdige transacties met een chunk-grootte van 500 records de beste balans bood tussen snelheid en middelengebruik. Deze configuratie stelt ons in staat om een hoge doorvoer te behouden zonder onze database te overweldigen of overmatig geheugen te verbruiken.

Interessant genoeg ontdekten we dat het verhogen van de gelijktijdigheid boven de 5 transacties geen significante prestatieverbeteringen opleverde en soms leidde tot verhoogde databasecontentie. Evenzo verbeterden grotere chunk-groottes de ruwe snelheid, maar ten koste van een hoger geheugenverbruik en langere responstijden voor kleine tot middelgrote importen/exporten.

## CSV-exports via E-mail Links

Het laatste onderdeel van onze oplossing pakt de uitdaging aan van het leveren van grote geëxporteerde bestanden aan gebruikers. In plaats van een directe download vanuit onze webapp te bieden, wat zou kunnen leiden tot time-outproblemen en een verhoogde serverbelasting, hebben we een systeem van gemailde downloadlinks geïmplementeerd.

Wanneer een gebruiker een grote export initieert, verwerkt ons systeem het verzoek op de achtergrond. Zodra het is voltooid, in plaats van de verbinding open te houden of het bestand op onze webservers op te slaan, uploaden we het bestand naar een veilige, tijdelijke opslaglocatie. We genereren vervolgens een unieke, veilige downloadlink en e-mailen deze naar de gebruiker.

Deze downloadlinks zijn 2 uur geldig, wat een balans biedt tussen gebruiksgemak en informatiebeveiliging. Deze tijdspanne geeft gebruikers voldoende gelegenheid om hun gegevens op te halen, terwijl ervoor wordt gezorgd dat gevoelige informatie niet onbeperkt toegankelijk blijft.

De beveiliging van deze downloadlinks was een topprioriteit in ons ontwerp. Elke link is:

- Uniek en willekeurig gegenereerd, waardoor het praktisch onmogelijk is om te raden
- Slechts 2 uur geldig
- Versleuteld tijdens verzending, waardoor de veiligheid van gegevens tijdens het downloaden wordt gewaarborgd

Deze aanpak biedt verschillende voordelen:

- Het vermindert de belasting van onze webservers, aangezien ze geen grote bestandsdownloads rechtstreeks hoeven af te handelen
- Het verbetert de gebruikerservaring, vooral voor gebruikers met langzamere internetverbindingen die mogelijk tegen browser time-outproblemen aanlopen bij directe downloads
- Het biedt een betrouwbaardere oplossing voor zeer grote exports die de typische web time-outlimieten kunnen overschrijden

De feedback van gebruikers over deze functie is overweldigend positief, met velen die de flexibiliteit waarderen die het biedt bij het beheren van grote data-exporten.

## Exporteren van Gefilterde Gegevens

De andere voor de hand liggende verbetering was om gebruikers alleen gegevens te laten exporteren die al in hun projectweergave waren gefilterd. Dit betekent dat als er een actieve tag "prioriteit" is, alleen records met deze tag in de CSV-export zouden eindigen. Dit betekent minder tijd besteden aan het manipuleren van gegevens in Excel om dingen eruit te filteren die niet belangrijk zijn, en helpt ons ook het aantal rijen dat moet worden verwerkt te verminderen.

## Vooruitkijken

Hoewel we geen onmiddellijke plannen hebben om ons gebruik van Rust uit te breiden, heeft dit project ons het potentieel van deze technologie voor prestatiekritische operaties laten zien. Het is een spannende optie die we nu in onze toolkit hebben voor toekomstige optimalisatiebehoeften. Deze herziening van CSV-import en -export sluit perfect aan bij Blue's toewijding aan schaalbaarheid.

We zijn toegewijd aan het bieden van een platform dat meegroeit met onze klanten, dat hun groeiende data-behoeften aankan zonder in te boeten op prestaties.

De beslissing om Rust in onze technologiestack te introduceren werd niet lichtvaardig genomen. Het stelde een belangrijke vraag die veel engineeringteams zich stellen: Wanneer is het gepast om buiten je kerntechnologiestack te treden, en wanneer moet je bij vertrouwde tools blijven?

Er is geen pasklare oplossing, maar bij Blue hebben we een kader ontwikkeld voor het nemen van deze cruciale beslissingen:

- **Probleem-First Aanpak:** We beginnen altijd met het duidelijk definiëren van het probleem dat we proberen op te lossen. In dit geval moesten we de prestaties van CSV-imports en -exports voor grote datasets drastisch verbeteren.
- **Uitputting van Bestaande Oplossingen:** Voordat we buiten onze kernstack kijken, verkennen we grondig wat er kan worden bereikt met onze bestaande technologieën. Dit omvat vaak profilering, optimalisatie en het heroverwegen van onze aanpak binnen vertrouwde beperkingen.
- **Kwantiseren van de Potentiële Winst:** Als we een nieuwe technologie overwegen, moeten we in staat zijn om de voordelen duidelijk te articuleren en idealiter te kwantificeren. Voor ons CSV-project projekteerden we verbeteringen van een orde van grootte in verwerkingssnelheid.
- **Beoordelen van de Kosten:** Het introduceren van een nieuwe technologie gaat niet alleen om het onmiddellijke project. We overwegen de langetermijnkosten:
  - Leercurve voor het team
  - Doorlopende onderhouds- en ondersteuningskosten
  - Potentiële complicaties bij implementatie en operaties
  - Impact op werving en team samenstelling
- **Beperking en Integratie:** Als we een nieuwe technologie introduceren, streven we ernaar deze te beperken tot een specifiek, goed gedefinieerd deel van ons systeem. We zorgen ook voor een duidelijk plan voor hoe het zal integreren met onze bestaande stack.
- **Toekomstbestendigheid:** We overwegen of deze technologiekeuze toekomstige kansen opent of ons mogelijk in een hoek kan drijven.

Een van de belangrijkste risico's van het frequent adopteren van nieuwe technologieën is dat je eindigt met wat we een *"technologie-dierentuin"* noemen - een gefragmenteerd ecosysteem waarbij verschillende delen van je applicatie in verschillende talen of frameworks zijn geschreven, wat een breed scala aan gespecialiseerde vaardigheden vereist om te onderhouden.

## Conclusie

Dit project is een voorbeeld van Blue's benadering van engineering: *we zijn niet bang om buiten onze comfortzone te treden en nieuwe technologieën te adopteren wanneer dit betekent dat we een aanzienlijk betere ervaring voor onze gebruikers kunnen bieden.*

Door ons CSV-import- en exportproces opnieuw te bedenken, hebben we niet alleen een dringende behoefte voor één enterprise klant opgelost, maar ook de ervaring voor al onze gebruikers die met grote datasets werken verbeterd.

Terwijl we blijven streven naar het verleggen van de grenzen van wat mogelijk is in [projectmanagementsoftware](/solutions/use-case/project-management), zijn we enthousiast om meer uitdagingen zoals deze aan te pakken.

Blijf op de hoogte voor meer [diepgaande artikelen over de engineering die Blue aandrijft!](/insights/engineering-blog)