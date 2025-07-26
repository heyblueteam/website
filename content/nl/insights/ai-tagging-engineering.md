---
title: AI Auto-Categorisatie (Engineering Deep Dive)
category: "Engineering"
description: Ga achter de schermen met het Blue engineering team terwijl ze uitleggen hoe ze een AI-gestuurde auto-categorisatie en tagging functie hebben gebouwd.
date: 2024-12-07
---

We hebben onlangs [AI Auto-Categorisatie](/insights/ai-auto-categorization) uitgerold naar alle Blue gebruikers. Dit is een AI-functie die is gebundeld in het kernabonnement van Blue, zonder extra kosten. In dit bericht duiken we in de engineering achter het realiseren van deze functie.

---
Bij Blue is onze aanpak van functie-ontwikkeling geworteld in een diep begrip van gebruikersbehoeften en markttrends, gekoppeld aan een toewijding om de eenvoud en het gebruiksgemak te behouden die ons platform definiëren. Dit is wat onze [roadmap](/platform/roadmap) aandrijft, en wat ons [in staat heeft gesteld om jarenlang consequent elke maand functies te leveren](/platform/changelog).

De introductie van AI-gestuurde auto-tagging in Blue is een perfect voorbeeld van deze filosofie in actie. Voordat we in de technische details duiken van hoe we deze functie hebben gebouwd, is het cruciaal om het probleem te begrijpen dat we oplosten en de zorgvuldige overweging die in de ontwikkeling ervan ging.

Het projectmanagementlandschap evolueert snel, waarbij AI-mogelijkheden steeds centraler komen te staan in gebruikersverwachtingen. Onze klanten, vooral degenen die grootschalige [projecten](/platform) beheren met miljoenen [records](/platform/features/records), waren uitgesproken over hun wens voor slimmere, efficiëntere manieren om hun data te organiseren en te categoriseren.

Bij Blue voegen we echter niet zomaar functies toe omdat ze trendy zijn of gevraagd worden. Onze filosofie is dat elke nieuwe toevoeging zijn waarde moet bewijzen, waarbij het standaardantwoord een resoluut *"nee"* is totdat een functie sterke vraag en duidelijk nut aantoont.

Om de diepte van het probleem en het potentieel van AI auto-tagging echt te begrijpen, hebben we uitgebreide klantinterviews uitgevoerd, gericht op langdurige gebruikers die complexe, datarijke projecten beheren over meerdere domeinen.

Deze gesprekken onthulden een gemeenschappelijke draad: *hoewel tagging onschatbaar was voor organisatie en doorzoekbaarheid, werd de handmatige aard van het proces een knelpunt, vooral voor teams die met grote volumes records werken.*

Maar we zagen verder dan alleen het oplossen van het directe pijnpunt van handmatig taggen.

We zagen een toekomst voor ons waarin AI-gestuurde tagging de basis zou kunnen worden voor intelligentere, geautomatiseerde workflows.

De echte kracht van deze functie, realiseerden we ons, lag in het potentieel voor integratie met ons [projectmanagement automatiseringssysteem](/platform/features/automations). Stel je een projectmanagementtool voor die niet alleen informatie intelligent categoriseert, maar ook die categorieën gebruikt om taken te routeren, acties te triggeren en workflows in real-time aan te passen.

Deze visie sloot perfect aan bij ons doel om Blue eenvoudig maar krachtig te houden.

Bovendien erkenden we het potentieel om deze mogelijkheid uit te breiden buiten de grenzen van ons platform. Door een robuust AI-taggingsysteem te ontwikkelen, legden we de basis voor een "categorisatie-API" die out-of-the-box zou kunnen werken, wat mogelijk nieuwe wegen opent voor hoe onze gebruikers Blue gebruiken en benutten in hun bredere tech-ecosystemen.

Deze functie ging dus niet alleen over het toevoegen van een AI-checkbox aan onze functielijst.

Het ging over het nemen van een belangrijke stap naar een intelligenter, adaptief projectmanagementplatform terwijl we trouw bleven aan onze kernfilosofie van eenvoud en gebruikersgerichtheid.

In de volgende secties duiken we in de technische uitdagingen die we tegenkwamen bij het tot leven brengen van deze visie, de architectuur die we ontwierpen om het te ondersteunen, en de oplossingen die we implementeerden. We verkennen ook de toekomstige mogelijkheden die deze functie opent, en laten zien hoe een zorgvuldig overwogen toevoeging de weg kan vrijmaken voor transformatieve veranderingen in projectmanagement.

---
## Het Probleem

Zoals hierboven besproken, kan handmatig taggen van projectrecords tijdrovend en inconsistent zijn.

We wilden dit oplossen door AI te gebruiken om automatisch tags voor te stellen op basis van recordinhoud.

De belangrijkste uitdagingen waren:

1. Een geschikt AI-model kiezen
2. Efficiënt verwerken van grote volumes records
3. Zorgen voor dataprivacy en beveiliging
4. De functie naadloos integreren in onze bestaande architectuur

## Het AI-Model Selecteren

We evalueerden verschillende AI-platforms, waaronder [OpenAI](https://openai.com), open-source modellen op [HuggingFace](https://huggingface.co/), en [Replicate](https://replicate.com).

Onze criteria waren onder andere:

- Kosteneffectiviteit
- Nauwkeurigheid in het begrijpen van context
- Vermogen om aan specifieke outputformaten te voldoen
- Garanties voor dataprivacy

Na grondig testen kozen we voor OpenAI's [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo). Hoewel [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) marginale verbeteringen in nauwkeurigheid zou kunnen bieden, toonden onze tests aan dat de prestaties van GPT-3.5 meer dan voldoende waren voor onze auto-tagging behoeften. De balans tussen kosteneffectiviteit en sterke categorisatiemogelijkheden maakte GPT-3.5 de ideale keuze voor deze functie.

De hogere kosten van GPT-4 zouden ons hebben gedwongen om de functie als betaalde add-on aan te bieden, wat in strijd zou zijn met ons doel om **AI binnen ons hoofdproduct te bundelen zonder extra kosten voor eindgebruikers.**

Vanaf onze implementatie is de prijsstelling voor GPT-3.5 Turbo:

- $0.0005 per 1K input tokens (of $0.50 per 1M input tokens)
- $0.0015 per 1K output tokens (of $1.50 per 1M output tokens)

Laten we enkele aannames maken over een gemiddeld record in Blue:

- **Titel**: ~10 tokens
- **Beschrijving**: ~50 tokens
- **2 opmerkingen**: ~30 tokens elk
- **5 aangepaste velden**: ~10 tokens elk
- **Lijstnaam, vervaldatum en andere metadata**: ~20 tokens
- **Systeemprompt en beschikbare tags**: ~50 tokens

Totaal input tokens per record: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 tokens

Voor de output, laten we uitgaan van gemiddeld 3 voorgestelde tags per record, wat mogelijk rond de 20 output tokens zou zijn inclusief de JSON-opmaak.

Voor 1 miljoen records:

- Input kosten: (240 * 1,000,000 / 1,000,000) * $0.50 = $120
- Output kosten: (20 * 1,000,000 / 1,000,000) * $1.50 = $30

**Totale kosten voor auto-tagging van 1 miljoen records: $120 + $30 = $150**

## GPT3.5 Turbo Prestaties

Categorisatie is een taak waarin grote taalmodellen (LLM's) zoals GPT-3.5 Turbo uitblinken, waardoor ze bijzonder geschikt zijn voor onze auto-tagging functie. LLM's zijn getraind op enorme hoeveelheden tekstdata, waardoor ze context, semantiek en relaties tussen concepten kunnen begrijpen. Deze brede kennisbasis stelt hen in staat om categorisatietaken met hoge nauwkeurigheid uit te voeren over een breed scala aan domeinen.

Voor ons specifieke gebruik van projectmanagement-tagging toont GPT-3.5 Turbo verschillende belangrijke sterke punten:

- **Contextueel Begrip:** Kan de algehele context van een projectrecord begrijpen, waarbij niet alleen individuele woorden maar de betekenis van de gehele beschrijving, opmerkingen en andere velden wordt overwogen.
- **Flexibiliteit:** Kan zich aanpassen aan verschillende projecttypes en industrieën zonder uitgebreide herprogrammering.
- **Omgaan met Ambiguïteit:** Kan meerdere factoren afwegen om genuanceerde beslissingen te nemen.
- **Leren van Voorbeelden:** Kan snel nieuwe categorisatieschema's begrijpen en toepassen zonder extra training.
- **Multi-label Classificatie:** Kan meerdere relevante tags voor één record voorstellen, wat cruciaal was voor onze vereisten.

GPT-3.5 Turbo viel ook op door zijn betrouwbaarheid in het naleven van ons vereiste JSON-outputformaat, wat *cruciaal* was voor naadloze integratie met onze bestaande systemen. Open-source modellen, hoewel veelbelovend, voegden vaak extra opmerkingen toe of weken af van het verwachte formaat, wat extra nabewerking zou hebben vereist. Deze consistentie in outputformaat was een sleutelfactor in onze beslissing, omdat het onze implementatie aanzienlijk vereenvoudigde en potentiële faalpunten verminderde.

Door te kiezen voor GPT-3.5 Turbo met zijn consistente JSON-output konden we een eenvoudigere, betrouwbaardere en beter onderhoudbare oplossing implementeren.

Als we hadden gekozen voor een model met minder betrouwbare opmaak, zouden we zijn geconfronteerd met een cascade van complicaties: de noodzaak voor robuuste parsing-logica om verschillende outputformaten te verwerken, uitgebreide foutafhandeling voor inconsistente outputs, potentiële prestatie-impact door extra verwerking, verhoogde testcomplexiteit om alle outputvariaties te dekken, en een grotere onderhoudsbelasting op lange termijn.

Parsing-fouten kunnen leiden tot onjuiste tagging, wat de gebruikerservaring negatief beïnvloedt. Door deze valkuilen te vermijden, konden we onze engineering-inspanningen richten op kritieke aspecten zoals prestatie-optimalisatie en gebruikersinterfaceontwerp, in plaats van te worstelen met onvoorspelbare AI-outputs.

## Systeemarchitectuur

Onze AI auto-tagging functie is gebouwd op een robuuste, schaalbare architectuur die is ontworpen om grote volumes aanvragen efficiënt te verwerken terwijl een naadloze gebruikerservaring wordt geboden. Zoals bij al onze systemen, hebben we deze functie gearchitecteerd om één orde van grootte meer verkeer te ondersteunen dan we momenteel ervaren. Deze aanpak, hoewel schijnbaar overengineered voor huidige behoeften, is een best practice die ons in staat stelt om plotselinge pieken in gebruik naadloos op te vangen en ons voldoende ruimte geeft voor groei zonder grote architecturale aanpassingen. Anders zouden we al onze systemen elke 18 maanden opnieuw moeten engineeren — iets wat we in het verleden op de harde manier hebben geleerd!

Laten we de componenten en flow van ons systeem uiteenzetten:

- **Gebruikersinteractie:** Het proces begint wanneer een gebruiker op de "Autotag" knop drukt in de Blue interface. Deze actie triggert de auto-tagging workflow.
- **Blue API Call:** De actie van de gebruiker wordt vertaald naar een API-aanroep naar onze Blue backend. Dit API-eindpunt is ontworpen om auto-tagging verzoeken af te handelen.
- **Queue Management:** In plaats van het verzoek onmiddellijk te verwerken, wat onder hoge belasting tot prestatieproblemen kan leiden, voegen we het tagging-verzoek toe aan een wachtrij. We gebruiken Redis voor dit wachtrijmechanisme, wat ons in staat stelt om de belasting effectief te beheren en systeemschaalbaarheid te garanderen.
- **Achtergrondservice:** We hebben een achtergrondservice geïmplementeerd die continu de wachtrij monitort op nieuwe verzoeken. Deze service is verantwoordelijk voor het verwerken van verzoeken in de wachtrij.
- **OpenAI API Integratie:** De achtergrondservice bereidt de benodigde data voor en maakt API-aanroepen naar OpenAI's GPT-3.5 model. Dit is waar de daadwerkelijke AI-gestuurde tagging plaatsvindt. We sturen relevante projectdata en ontvangen voorgestelde tags terug.
- **Resultaatverwerking:** De achtergrondservice verwerkt de ontvangen resultaten van OpenAI. Dit omvat het parsen van de AI-respons en het voorbereiden van de data voor toepassing op het project.
- **Tag Toepassing:** De verwerkte resultaten worden gebruikt om de nieuwe tags toe te passen op de relevante items in het project. Deze stap werkt onze database bij met de door AI voorgestelde tags.
- **Gebruikersinterface Reflectie:** Ten slotte verschijnen de nieuwe tags in de projectweergave van de gebruiker, waarmee het auto-tagging proces vanuit het perspectief van de gebruiker wordt voltooid.

Deze architectuur biedt verschillende belangrijke voordelen die zowel systeemprestaties als gebruikerservaring verbeteren. Door gebruik te maken van een wachtrij en achtergrondverwerking hebben we indrukwekkende schaalbaarheid bereikt, waardoor we talloze verzoeken tegelijkertijd kunnen afhandelen zonder ons systeem te overbelasten of de rate limits van de OpenAI API te bereiken. Het implementeren van deze architectuur vereiste zorgvuldige overweging van verschillende factoren om optimale prestaties en betrouwbaarheid te garanderen. Voor wachtrijbeheer kozen we Redis, waarbij we gebruikmaken van zijn snelheid en betrouwbaarheid bij het afhandelen van gedistribueerde wachtrijen.

Deze aanpak draagt ook bij aan de algehele responsiviteit van de functie. Gebruikers ontvangen onmiddellijke feedback dat hun verzoek wordt verwerkt, zelfs als de daadwerkelijke tagging enige tijd kost, wat een gevoel van real-time interactie creëert. De fouttolerantie van de architectuur is een ander cruciaal voordeel. Als een deel van het proces problemen ondervindt, zoals tijdelijke OpenAI API-onderbrekingen, kunnen we de fout netjes opnieuw proberen of afhandelen zonder het hele systeem te beïnvloeden.

Deze robuustheid, gecombineerd met het real-time verschijnen van tags, verbetert de gebruikerservaring en geeft de indruk van AI "magie" aan het werk.

## Data & Prompts

Een cruciale stap in ons auto-tagging proces is het voorbereiden van de data die naar het GPT-3.5 model wordt gestuurd. Deze stap vereiste zorgvuldige overweging om een balans te vinden tussen het bieden van voldoende context voor nauwkeurige tagging terwijl efficiëntie wordt behouden en gebruikersprivacy wordt beschermd. Hier is een gedetailleerd overzicht van ons datavoorbereidingsproces.

Voor elk record compileren we de volgende informatie:

- **Lijstnaam**: Biedt context over de bredere categorie of fase van het project.
- **Record Titel**: Bevat vaak belangrijke informatie over het doel of de inhoud van het record.
- **Aangepaste Velden**: We nemen tekst- en nummergebaseerde [aangepaste velden](/platform/features/custom-fields) op, die vaak cruciale projectspecifieke informatie bevatten.
- **Beschrijving**: Bevat meestal de meest gedetailleerde informatie over het record.
- **Opmerkingen**: Kunnen aanvullende context of updates bieden die relevant kunnen zijn voor tagging.
- **Vervaldatum**: Temporele informatie die de tagselectie kan beïnvloeden.

Interessant genoeg sturen we geen bestaande tagdata naar GPT-3.5, en we doen dit om het model niet te beïnvloeden.

De kern van onze auto-tagging functie ligt in hoe we interacteren met het GPT-3.5 model en zijn antwoorden verwerken. Dit deel van onze pijplijn vereiste zorgvuldig ontwerp om nauwkeurige, consistente en efficiënte tagging te garanderen.

We gebruiken een zorgvuldig opgesteld systeemprompt om de AI te instrueren over zijn taak. Hier is een uitsplitsing van ons prompt en de redenering achter elk component:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Taakdefinitie:** We stellen duidelijk de taak van de AI om gefocuste reacties te garanderen.
- **Onzekerheidsafhandeling:** We staan expliciet lege reacties toe, waardoor gedwongen tagging wordt voorkomen wanneer de AI onzeker is.
- **Beschikbare Tags:** We bieden een lijst van geldige tags (${tags}) om de keuzes van de AI te beperken tot bestaande projecttags.
- **Huidige Datum:** Het opnemen van ${currentDate} helpt de AI de temporele context te begrijpen, wat cruciaal kan zijn voor bepaalde soorten projecten.
- **Reactieformaat:** We specificeren een JSON-formaat voor gemakkelijke parsing en foutcontrole.

Dit prompt is het resultaat van uitgebreid testen en iteratie. We ontdekten dat expliciet zijn over de taak, beschikbare opties en gewenst outputformaat de nauwkeurigheid en consistentie van de AI-reacties aanzienlijk verbeterde — eenvoud is de sleutel!

De lijst met beschikbare tags wordt server-side gegenereerd en gevalideerd voor opname in het prompt. We implementeren strikte karakterlimieten op tagnamen om te grote prompts te voorkomen.

Zoals hierboven vermeld, hadden we geen probleem met GPT-3.5 Turbo om 100% van de tijd de pure JSON-respons in het juiste formaat terug te krijgen.

Dus samengevat,

- We combineren het systeemprompt met de voorbereide recorddata.
- Dit gecombineerde prompt wordt naar het GPT-3.5 model gestuurd via OpenAI's API.
- We gebruiken een temperatuurinstelling van 0.3 om creativiteit en consistentie in de AI-reacties in balans te brengen.
- Onze API-aanroep bevat een max_tokens parameter om de responsgrootte te beperken en kosten te beheersen.

Zodra we de AI-respons ontvangen, doorlopen we verschillende stappen om de voorgestelde tags te verwerken en toe te passen:

* **JSON Parsing**: We proberen de respons als JSON te parsen. Als parsing mislukt, loggen we de fout en slaan we tagging voor dat record over.
* **Schema Validatie**: We valideren de geparseerde JSON tegen ons verwachte schema (een object met een "tags" array). Dit vangt structurele afwijkingen in de AI-respons op.
* **Tag Validatie**: We kruisverwijzen de voorgestelde tags met onze lijst van geldige projecttags. Deze stap filtert tags eruit die niet in het project bestaan, wat kan voorkomen als de AI het verkeerd begrepen heeft of als projecttags zijn veranderd tussen promptcreatie en responsverwerking.
* **Deduplicatie**: We verwijderen dubbele tags uit de AI-suggestie om redundante tagging te vermijden.
* **Toepassing**: De gevalideerde en gededupliceerde tags worden vervolgens toegepast op het record in onze database.
* **Logging en Analytics**: We loggen de uiteindelijk toegepaste tags. Deze data is waardevol voor het monitoren van de systeemprestaties en het verbeteren ervan in de loop van de tijd.

## Uitdagingen

Het implementeren van AI-gestuurde auto-tagging in Blue presenteerde verschillende unieke uitdagingen, die elk innovatieve oplossingen vereisten om een robuuste, efficiënte en gebruiksvriendelijke functie te garanderen.

### Bulk Operatie Ongedaan Maken

De AI Tagging functie kan zowel op individuele records als in bulk worden uitgevoerd. Het probleem met de bulkoperatie is dat als de gebruiker het resultaat niet bevalt, ze handmatig door duizenden records zouden moeten gaan en het werk van de AI ongedaan zouden moeten maken. Dat is duidelijk onaanvaardbaar.

Om dit op te lossen, hebben we een innovatief tagging-sessiesysteem geïmplementeerd. Elke bulk tagging-operatie krijgt een unieke sessie-ID toegewezen, die is gekoppeld aan alle tags die tijdens die sessie zijn toegepast. Dit stelt ons in staat om ongedaan maken-operaties efficiënt te beheren door simpelweg alle tags te verwijderen die aan een bepaalde sessie-ID zijn gekoppeld. We verwijderen ook gerelateerde audit trails, zodat ongedaan gemaakte operaties geen spoor achterlaten in het systeem. Deze aanpak geeft gebruikers het vertrouwen om te experimenteren met AI-tagging, wetende dat ze wijzigingen gemakkelijk kunnen terugdraaien indien nodig.

### Dataprivacy

Dataprivacy was een andere kritieke uitdaging die we tegenkwamen.

Onze gebruikers vertrouwen ons hun projectdata toe, en het was van het grootste belang om ervoor te zorgen dat deze informatie niet werd bewaard of gebruikt voor modeltraining door OpenAI. We pakten dit op meerdere fronten aan.

Eerst vormden we een overeenkomst met OpenAI die expliciet het gebruik van onze data voor modeltraining verbiedt. Bovendien verwijdert OpenAI de data na verwerking, wat een extra laag privacybescherming biedt.

Aan onze kant namen we de voorzorgsmaatregel om gevoelige informatie, zoals toegewezen personen, uit te sluiten van de data die naar de AI wordt gestuurd, zodat dit ervoor zorgt dat specifieke persoonsnamen niet samen met andere data naar derden worden gestuurd.

Deze uitgebreide aanpak stelt ons in staat om AI-mogelijkheden te benutten terwijl we de hoogste normen voor dataprivacy en beveiliging handhaven.

### Rate Limits en Fouten Opvangen

Een van onze primaire zorgen was schaalbaarheid en rate limiting. Directe API-aanroepen naar OpenAI voor elk record zouden inefficiënt zijn geweest en zouden snel rate limits kunnen bereiken, vooral voor grote projecten of tijdens piekgebruikstijden. Om dit aan te pakken, ontwikkelden we een achtergrondservice-architectuur die ons in staat stelt om verzoeken te batchen en ons eigen wachtrijsysteem te implementeren. Deze aanpak helpt ons de frequentie van API-aanroepen te beheren en maakt efficiëntere verwerking van grote volumes records mogelijk, wat zorgt voor soepele prestaties zelfs onder zware belasting.

De aard van AI-interacties betekende dat we ons ook moesten voorbereiden op occasionele fouten of onverwachte outputs. Er waren gevallen waarin de AI ongeldige JSON of outputs zou kunnen produceren die niet overeenkwamen met ons verwachte formaat. Om dit te verwerken, implementeerden we robuuste foutafhandeling en parsing-logica door ons hele systeem. Als de AI-respons geen geldige JSON is of niet de verwachte "tags" sleutel bevat, is ons systeem ontworpen om het te behandelen alsof er geen tags werden voorgesteld, in plaats van te proberen mogelijk corrupte data te verwerken. Dit zorgt ervoor dat zelfs bij AI-onvoorspelbaarheid ons systeem stabiel en betrouwbaar blijft.

## Toekomstige Ontwikkelingen

We geloven dat functies, en het Blue product als geheel, nooit "af" zijn — er is altijd ruimte voor verbetering.

Er waren enkele functies die we overwogen in de initiële build die de scopingfase niet haalden, maar die interessant zijn om op te merken omdat we waarschijnlijk een versie ervan in de toekomst zullen implementeren.

De eerste is het toevoegen van tagbeschrijvingen. Dit zou eindgebruikers in staat stellen om tags niet alleen een naam en een kleur te geven, maar ook een optionele beschrijving. Dit zou ook aan de AI worden doorgegeven om verdere context te helpen bieden en mogelijk de nauwkeurigheid te verbeteren.

Hoewel aanvullende context waardevol zou kunnen zijn, zijn we ons bewust van de potentiële complexiteit die het zou kunnen introduceren. Er is een delicate balans om te vinden tussen het bieden van nuttige informatie en het overweldigen van gebruikers met te veel detail. Terwijl we deze functie ontwikkelen, zullen we ons richten op het vinden van die sweet spot waar toegevoegde context de gebruikerservaring verbetert in plaats van compliceert.

Misschien wel de meest opwindende verbetering aan onze horizon is de integratie van AI auto-tagging met ons [projectmanagement automatiseringssysteem](/platform/features/automations).

Dit betekent dat de AI-tagging functie een trigger of een actie van een automatisering zou kunnen zijn. Dit kan enorm zijn omdat het deze vrij eenvoudige AI-categorisatiefunctie zou kunnen veranderen in een AI-gebaseerd routeringssysteem voor werk.

Stel je een automatisering voor die stelt:

Wanneer AI een record tagt als "Kritiek" -> Toewijzen aan "Manager" en Stuur een Aangepaste E-mail

Dit betekent dat wanneer je een record AI-tagt, als de AI beslist dat het een kritiek probleem is, het automatisch de projectmanager kan toewijzen en hen een aangepaste e-mail kan sturen. Dit breidt de [voordelen van ons projectmanagement automatiseringssysteem](/platform/features/automations) uit van puur een regelgebaseerd systeem naar een echt flexibel AI-systeem.

Door continu de grenzen van AI in projectmanagement te verkennen, streven we ernaar onze gebruikers tools te bieden die niet alleen aan hun huidige behoeften voldoen, maar ook de toekomst van werk anticiperen en vormgeven. Zoals altijd zullen we deze functies ontwikkelen in nauwe samenwerking met onze gebruikersgemeenschap, om ervoor te zorgen dat elke verbetering echte, praktische waarde toevoegt aan het projectmanagementproces.

## Conclusie

Dus dat is het!

Dit was een leuke functie om te implementeren, en een van onze eerste stappen in AI, naast de [AI Content Samenvatting](/insights/ai-content-summarization) die we eerder hebben gelanceerd. We weten dat AI een steeds grotere rol zal spelen in projectmanagement in de toekomst, en we kunnen niet wachten om meer innovatieve functies uit te rollen die gebruikmaken van geavanceerde LLMs (Large Language Models).

Er was behoorlijk wat om over na te denken tijdens de implementatie, en we zijn vooral enthousiast over hoe we deze functie in de toekomst kunnen benutten met Blue's bestaande [projectmanagement automatiseringsengine](/insights/benefits-project-management-automation).

We hopen ook dat het een interessante lezing is geweest, en dat het je een glimp geeft van hoe we nadenken over het engineeren van de functies die je elke dag gebruikt.
