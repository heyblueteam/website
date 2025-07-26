---
title:  Het creëren van Blue's Aangepaste Toegangsrechten Engine
description: Ga achter de schermen met het Blue engineeringteam terwijl ze uitleggen hoe ze een AI-gestuurde auto-categorisatie en tagging functie hebben gebouwd.
category: "Engineering"
date: 2024-07-25
---



Effectief project- en procesmanagement is cruciaal voor organisaties van elke omvang.

Bij Blue hebben [we het tot onze missie gemaakt](/about) om het werk van de wereld te organiseren door het beste projectmanagementplatform op de planeet te bouwen—simpel, krachtig, flexibel en betaalbaar voor iedereen.

Dit betekent dat ons platform zich moet aanpassen aan de unieke behoeften van elk team. Vandaag zijn we enthousiast om een kijkje achter de schermen te geven van een van onze krachtigste functies: Aangepaste Toegangsrechten.

Projectmanagementtools zijn de ruggengraat van moderne workflows, waarin gevoelige gegevens, cruciale communicatie en strategische plannen zijn ondergebracht. Daarom is de mogelijkheid om de toegang tot deze informatie nauwkeurig te controleren niet alleen een luxe—het is een noodzaak.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>


Aangepaste toegangsrechten spelen een cruciale rol in B2B SaaS-platforms, vooral in projectmanagementtools, waar de balans tussen samenwerking en beveiliging het succes van een project kan maken of breken.

Maar hier neemt Blue een andere benadering: **we geloven dat functies van ondernemingskwaliteit niet voorbehouden moeten zijn aan ondernemingsbudgetten.**

In een tijdperk waarin AI kleine teams in staat stelt om op ongekende schaal te opereren, waarom zouden robuuste beveiliging en aanpassing dan buiten bereik moeten zijn?

In deze blik achter de schermen zullen we verkennen hoe we onze functie voor Aangepaste Toegangsrechten hebben ontwikkeld, de status quo van SaaS-prijsniveaus uitdaagden en krachtige, flexibele beveiligingsopties naar bedrijven van elke omvang brachten.

Of je nu een startup met grote dromen bent of een gevestigde speler die zijn processen wil optimaliseren, aangepaste toegangsrechten kunnen nieuwe gebruiksgevallen mogelijk maken waarvan je niet eens wist dat ze mogelijk waren.

## Begrijpen van Aangepaste Gebruikersrechten

Voordat we ons verdiepen in onze reis naar het ontwikkelen van aangepaste toegangsrechten voor Blue, laten we even stilstaan bij wat aangepaste gebruikersrechten zijn en waarom ze zo cruciaal zijn in projectmanagementsoftware.

Aangepaste gebruikersrechten verwijzen naar de mogelijkheid om toegangsrechten voor individuele gebruikers of groepen binnen een softwaresysteem op maat te maken. In plaats van te vertrouwen op vooraf gedefinieerde rollen met vaste sets van rechten, stellen aangepaste rechten beheerders in staat om zeer specifieke toegangsprofielen te creëren die perfect aansluiten bij de structuur en workflowbehoeften van hun organisatie.

In de context van projectmanagementsoftware zoals Blue omvatten aangepaste rechten:

* **Fijne toegangscontrole**: Bepalen wie specifieke soorten projectgegevens kan bekijken, bewerken of verwijderen.
* **Functie-gebaseerde beperkingen**: Bepalen of bepaalde functies voor specifieke gebruikers of teams zijn ingeschakeld of uitgeschakeld.
* **Gevoeligheidsniveaus van gegevens**: Instellen van verschillende niveaus van toegang tot gevoelige informatie binnen projecten.
* **Workflow-specifieke rechten**: De mogelijkheden van gebruikers afstemmen op specifieke fasen of aspecten van je projectworkflow.

Het belang van aangepaste rechten in projectmanagement kan niet genoeg worden benadrukt:

* **Verbeterde beveiliging**: Door gebruikers alleen de toegang te geven die ze nodig hebben, verklein je het risico op datalekken of ongeautoriseerde wijzigingen.
* **Verbeterde naleving**: Aangepaste rechten helpen organisaties om te voldoen aan specifieke regelgeving door de toegang tot gegevens te controleren.
* **Gestroomlijnde samenwerking**: Teams kunnen efficiënter werken wanneer elk lid het juiste niveau van toegang heeft om zijn rol uit te voeren zonder onnodige beperkingen of overweldigende privileges.
* **Flexibiliteit voor complexe organisaties**: Naarmate bedrijven groeien en evolueren, stellen aangepaste rechten de software in staat om zich aan te passen aan veranderende organisatiestructuren en processen.

## Naar JA

[We hebben eerder geschreven](/insights/value-proposition-blue) dat elke functie in Blue een **hard** JA moet zijn voordat we besluiten deze te bouwen. We hebben niet de luxe van honderden ingenieurs en kunnen geen tijd en geld verspillen aan het bouwen van dingen die klanten niet nodig hebben.

En zo was de weg naar de implementatie van aangepaste rechten in Blue geen rechte lijn. Zoals veel krachtige functies begon het met een duidelijke behoefte van onze gebruikers en evolueerde het door zorgvuldige overweging en planning.

Jarenlang hadden onze klanten gevraagd om meer gedetailleerde controle over gebruikersrechten. Terwijl organisaties van elke omvang steeds complexere en gevoelige projecten begonnen te beheren, werden de beperkingen van onze standaard rolgebaseerde toegangscontrole duidelijk.

Kleine startups die met externe klanten werken, middelgrote bedrijven met ingewikkelde goedkeuringsprocessen en grote ondernemingen met strikte nalevingsvereisten gaven allemaal dezelfde behoefte aan:

Meer flexibiliteit in het beheren van gebruikersaccess.

Ondanks de duidelijke vraag aarzelden we aanvankelijk om in te gaan op de ontwikkeling van aangepaste rechten.

Waarom?

We begrepen de complexiteit die erbij kwam kijken!

Aangepaste rechten raken elk onderdeel van een projectmanagementsysteem, van de gebruikersinterface tot de database-structuur. We wisten dat het implementeren van deze functie aanzienlijke wijzigingen in onze kernarchitectuur zou vereisen en zorgvuldige overweging van de prestatie-implicaties.

Toen we het landschap in kaart brachten, merkten we dat zeer weinig van onze concurrenten hadden geprobeerd een krachtige engine voor aangepaste rechten te implementeren zoals onze klanten vroegen. Degenen die dat deden, reserveerden het vaak voor hun hoogste ondernemingsplannen.

Het werd duidelijk waarom: de ontwikkelingsinspanning is aanzienlijk en de inzet is hoog.

Aangepaste rechten onjuist implementeren kan kritieke bugs of beveiligingskwetsbaarheden introduceren, wat mogelijk het hele systeem in gevaar brengt. Deze realisatie benadrukte de omvang van de uitdaging die we overwogen.

### De Status Quo Uitdagen

Echter, terwijl we bleven groeien en evolueren, bereikten we een cruciale realisatie:

**Het traditionele SaaS-model van het reserveren van krachtige functies voor ondernemingsklanten maakt in het huidige zakelijke landschap niet langer zin.**

In 2024, met de kracht van AI en geavanceerde tools, kunnen kleine teams opereren op een schaal en complexiteit die rivalen van veel grotere organisaties. Een startup kan gevoelige klantgegevens beheren in meerdere landen. Een klein marketingbureau kan tientallen klantprojecten met verschillende vertrouwelijkheidsvereisten jongleren. Deze bedrijven hebben hetzelfde niveau van beveiliging en aanpassing nodig als *elke* grote onderneming.

We stelden onszelf de vraag: Waarom zou de grootte van een bedrijf of budget bepalen of ze hun gegevens veilig kunnen houden en hun processen efficiënt kunnen laten verlopen?

### Ondernemingskwaliteit voor Iedereen

Deze realisatie leidde ons tot een kernfilosofie die nu veel van onze ontwikkeling bij Blue aanstuurt: Ondernemingsfuncties moeten toegankelijk zijn voor bedrijven van elke omvang.

We geloven dat:

- **Beveiliging geen luxe zou moeten zijn.** Elk bedrijf, ongeacht de grootte, verdient de tools om hun gegevens en processen te beschermen.
- **Flexibiliteit innovatie stimuleert.** Door al onze gebruikers krachtige tools te geven, stellen we hen in staat om workflows en systemen te creëren die hun industrieën vooruit helpen.
- **Groei geen platformwijzigingen zou moeten vereisen.** Naarmate onze klanten groeien, zouden hun tools naadloos met hen mee moeten groeien.

Met deze mindset besloten we de uitdaging van aangepaste rechten direct aan te pakken, vastbesloten om het beschikbaar te maken voor al onze gebruikers, niet alleen voor degenen met hogere plannen.

Deze beslissing zette ons op een pad van zorgvuldige ontwerpeisen, iteratieve ontwikkeling en continue gebruikersfeedback die uiteindelijk leidde tot de functie voor aangepaste rechten die we vandaag met trots aanbieden.

In de volgende sectie zullen we ingaan op hoe we de ontwerp- en ontwikkelingsprocessen hebben benaderd om deze complexe functie tot leven te brengen.

### Ontwerp en Ontwikkeling

Toen we besloten om aangepaste rechten aan te pakken, realiseerden we ons al snel dat we voor een enorme taak stonden.

Op het eerste gezicht klinkt "aangepaste rechten" misschien eenvoudig, maar het is een bedrieglijk complexe functie die elk aspect van ons systeem raakt.

De uitdaging was ontmoedigend: we moesten cascaderende rechten implementeren, on-the-fly bewerkingen mogelijk maken, aanzienlijke wijzigingen in de databaseschema's aanbrengen en naadloze functionaliteit over ons hele ecosysteem waarborgen – web, Mac, Windows, iOS en Android-apps, evenals onze API en webhooks.

De complexiteit was genoeg om zelfs de meest ervaren ontwikkelaars te laten aarzelen.

Onze aanpak was gebaseerd op twee kernprincipes:

1. De functie opdelen in beheersbare versies
2. Incrementele verzending omarmen.

Geconfronteerd met de complexiteit van volledige aangepaste rechten, stelden we onszelf een cruciale vraag:

> Wat zou de eenvoudigste mogelijke eerste versie van deze functie zijn?

Deze aanpak sluit aan bij het agile principe van het leveren van een Minimum Viable Product (MVP) en itereren op basis van feedback.

Ons antwoord was verfrissend eenvoudig:

1. Introduceer een schakelaar om het projectactiviteitstabblad te verbergen
2. Voeg een andere schakelaar toe om het formuliertabblad te verbergen

**Dat was het.**

Geen toeters en bellen, geen complexe rechtenmatrices—slechts twee eenvoudige aan/uit-schakelaars.

Hoewel het op het eerste gezicht misschien teleurstellend lijkt, bood deze aanpak verschillende aanzienlijke voordelen:

* **Snelle Implementatie**: Deze eenvoudige schakelaars konden snel worden ontwikkeld en getest, waardoor we een basisversie van aangepaste rechten snel in handen van gebruikers konden krijgen.
* **Duidelijke Gebruikerswaarde**: Zelfs met slechts deze twee opties boden we tastbare waarde. Sommige teams willen misschien de activiteitsoverzicht voor klanten verbergen, terwijl anderen de toegang tot formulieren voor bepaalde gebruikersgroepen moeten beperken.
* **Basis voor Groei**: Deze eenvoudige start legde de basis voor complexere rechten. Het stelde ons in staat om de basisinfrastructuur voor aangepaste rechten op te zetten zonder ons vanaf het begin te laten verlammen door complexiteit.
* **Gebruikersfeedback**: Door deze eenvoudige versie uit te brengen, konden we feedback uit de echte wereld verzamelen over hoe gebruikers met aangepaste rechten omgingen, wat onze toekomstige ontwikkeling informeerde.
* **Technische Lering**: Deze initiële implementatie gaf ons ontwikkelingsteam praktische ervaring in het wijzigen van rechten over ons platform, waardoor we ons voorbereidden op complexere iteraties.

En je weet, het is eigenlijk best nederig om een groot visioen voor iets te hebben, en dan iets te verzenden dat zo'n klein percentage van dat visioen is.

Na het verzenden van deze eerste twee schakelaars besloten we om iets geavanceerders aan te pakken. We kwamen op twee nieuwe aangepaste gebruikersrolrechten.

De eerste was de mogelijkheid om gebruikers te beperken tot alleen het bekijken van records die specifiek aan hen zijn toegewezen. Dit is zeer nuttig als je een klant in een project hebt en je wilt dat ze alleen de records zien die specifiek aan hen zijn toegewezen in plaats van alles waar je voor hen aan werkt.

De tweede was een optie voor projectbeheerders om gebruikersgroepen te blokkeren van het uitnodigen van andere gebruikers. Dit is goed als je een gevoelig project hebt dat je wilt waarborgen dat het op een "need to see" basis blijft.

Zodra we dit hadden verzonden, kregen we meer vertrouwen en voor onze derde versie pakten we kolomniveau-rechten aan, wat betekent dat je kunt beslissen welke aangepaste velden een specifieke gebruikersgroep kan bekijken of bewerken.

Dit is extreem krachtig. Stel je voor dat je een CRM-project hebt, en je hebt gegevens daarin die niet alleen gerelateerd zijn aan de bedragen die de klant zal betalen, maar ook aan je kosten en winstmarges. Je wilt misschien niet dat je kostenvelden en projectmargeformuleveld zichtbaar zijn voor junior personeel, en aangepaste rechten stellen je in staat om die velden te vergrendelen zodat ze niet worden weergegeven.

Vervolgens gingen we verder met het creëren van lijstgebaseerde rechten, waarbij projectbeheerders kunnen beslissen of een gebruikersgroep een specifieke lijst kan bekijken, bewerken en verwijderen. Als ze een lijst verbergen, worden ook alle records binnen die lijst verborgen, wat geweldig is omdat het betekent dat je bepaalde delen van je proces kunt verbergen voor je teamleden of klanten.

Dit is het eindresultaat:

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Technische Overwegingen

In het hart van Blue's technische architectuur ligt GraphQL, een cruciale keuze die onze mogelijkheid om complexe functies zoals aangepaste rechten te implementeren aanzienlijk heeft beïnvloed. Maar voordat we in de details duiken, laten we een stap terugnemen en begrijpen wat GraphQL is en hoe het verschilt van de meer traditionele REST API-aanpak.
GraphQL vs REST API: Een Toegankelijke Uitleg

Stel je voor dat je in een restaurant bent. Met een REST API is het alsof je van een vast menu bestelt. Je vraagt om een specifiek gerecht (endpoint), en je krijgt alles wat erbij hoort, of je het nu wilt of niet. Als je je maaltijd wilt aanpassen, moet je misschien meerdere bestellingen plaatsen (API-aanroepen) of vragen om een speciaal bereid gerecht (aangepast endpoint).

GraphQL daarentegen is als een gesprek met een chef-kok die alles kan bereiden. Je vertelt de chef precies welke ingrediënten je wilt (gegevensvelden), en in welke hoeveelheden. De chef bereidt vervolgens een gerecht dat precies is wat je vroeg - niet meer, niet minder. Dit is in wezen wat GraphQL doet - het stelt de cliënt in staat om precies de gegevens te vragen die hij nodig heeft, en de server levert precies dat.

### Een Belangrijke Lunch

Ongeveer zes weken na de initiële ontwikkeling van Blue ging onze hoofdingenieur en CEO uit lunchen.

Het onderwerp van discussie?

Of we moesten overstappen van REST API's naar GraphQL. Dit was geen beslissing om lichtvaardig te nemen - het aannemen van GraphQL zou betekenen dat we zes weken aan initiële werkzaamheden moesten afschrijven.

Tijdens de terugweg naar kantoor stelde de CEO een cruciale vraag aan de hoofdingenieur: "Zullen we er over vijf jaar spijt van hebben dat we dit niet hebben gedaan?"

Het antwoord werd duidelijk: GraphQL was de weg vooruit.

We herkenden het potentieel van deze technologie vroeg, omdat we zagen hoe het onze visie voor een flexibel, krachtig projectmanagementplatform kon ondersteunen.

Onze vooruitziende blik bij het aannemen van GraphQL betaalde zich uit bij het implementeren van aangepaste rechten. Met een REST API zouden we voor elke mogelijke configuratie van aangepaste rechten een ander endpoint nodig hebben - een aanpak die snel onhandelbaar en moeilijk te onderhouden zou worden.

GraphQL stelt ons echter in staat om aangepaste rechten dynamisch te beheren. Hier is hoe het werkt:

- **On-the-fly Toegangscontroles**: Wanneer een cliënt een verzoek indient, kan onze GraphQL-server de rechten van de gebruiker direct vanuit onze database controleren.
- **Nauwkeurige Gegevensopvraging**: Op basis van deze rechten retourneert GraphQL alleen de gevraagde gegevens die binnen de toegangsrechten van de gebruiker passen.
- **Flexibele Vragen**: Naarmate de rechten veranderen, hoeven we geen nieuwe endpoints te creëren of bestaande te wijzigen. Dezelfde GraphQL-query kan zich aanpassen aan verschillende rechteninstellingen.
- **Efficiënte Gegevensopvraging**: GraphQL stelt cliënten in staat om precies te vragen wat ze nodig hebben. Dit betekent dat we geen gegevens overvragen, wat mogelijk informatie zou kunnen blootstellen die de gebruiker niet zou moeten kunnen zien.

Deze flexibiliteit is cruciaal voor een functie die zo complex is als aangepaste rechten. Het stelt ons in staat om gedetailleerde controle te bieden *zonder* in te boeten op prestaties of onderhoudbaarheid.

## Uitdagingen

Het implementeren van aangepaste rechten in Blue bracht zijn eigen uitdagingen met zich mee, die ons allemaal dwongen om te innoveren en onze aanpak te verfijnen. Prestatieoptimalisatie kwam snel naar voren als een kritieke zorg. Terwijl we meer gedetailleerde toegangscontroles toevoegden, riskeerden we ons systeem te vertragen, vooral voor grote projecten met veel gebruikers en complexe rechteninstellingen. Om dit aan te pakken, implementeerden we een gelaagde cachingstrategie, optimaliseerden we onze databasequery's en maakten we gebruik van de mogelijkheid van GraphQL om alleen noodzakelijke gegevens op te vragen. Deze aanpak stelde ons in staat om snelle responstijden te behouden, zelfs terwijl projecten groeiden en de complexiteit van de rechten toenam.

De gebruikersinterface voor aangepaste rechten vormde een andere aanzienlijke hindernis. We moesten de interface intuïtief en beheersbaar maken voor beheerders, zelfs terwijl we meer opties toevoegden en de complexiteit van het systeem verhoogden.

Onze oplossing omvatte meerdere rondes van gebruikerstests en iteratief ontwerp.

We introduceerden een visuele rechtenmatrix die beheerders in staat stelde om snel de rechten over verschillende rollen en projectgebieden te bekijken en te wijzigen.

Zorgdragen voor consistentie over verschillende platforms presenteerde zijn eigen uitdagingen. We moesten aangepaste rechten uniform implementeren over onze web-, desktop- en mobiele applicaties, elk met zijn unieke interface en gebruikerservaring. Dit was bijzonder lastig voor onze mobiele apps, die dynamisch functies moesten verbergen en tonen op basis van de rechten van de gebruiker. We losten dit op door onze logica voor rechten te centraliseren in de API-laag, zodat alle platforms consistente gegevens over rechten ontvingen.

Daarna ontwikkelden we een flexibel UI-framework dat zich in realtime kon aanpassen aan deze wijzigingen in rechten, waardoor een naadloze ervaring werd geboden, ongeacht het gebruikte platform.

Gebruikerseducatie en adoptie vormden de laatste hindernis in onze reis naar aangepaste rechten. Het introduceren van een dergelijke krachtige functie betekende dat we onze gebruikers moesten helpen om aangepaste rechten te begrijpen en effectief te benutten.

We lanceerden aanvankelijk aangepaste rechten voor een subset van onze gebruikersbasis, waarbij we hun ervaringen zorgvuldig monitoren en inzichten verzamelden. Deze aanpak stelde ons in staat om de functie en onze educatieve materialen te verfijnen op basis van het gebruik in de echte wereld voordat we deze aan onze gehele gebruikersbasis lanceerden.

De gefaseerde uitrol bleek van onschatbare waarde, omdat het ons hielp om kleine problemen en verwarring bij gebruikers te identificeren en aan te pakken die we niet hadden voorzien, wat uiteindelijk leidde tot een meer verfijnde en gebruiksvriendelijke functie voor al onze gebruikers.

Deze aanpak van het lanceren voor een subset van gebruikers, evenals onze typische 2-3 weken "Beta"-periode op onze openbare Beta, helpt ons om 's nachts rustig te slapen. :)

## Vooruitkijken

Zoals met alle functies, is niets ooit *"af"*.

Onze langetermijnvisie voor de functie voor aangepaste rechten strekt zich uit over tags, filters voor aangepaste velden, aanpasbare projectnavigatie en commentaarcontroles.

Laten we elk aspect bekijken.

### Tagrechten

We denken dat het geweldig zou zijn om rechten te kunnen creëren op basis van of een record een of meer tags heeft. De meest voor de hand liggende use case zou zijn dat je een aangepaste gebruikersrol genaamd "Klanten" creëert en alleen gebruikers in die rol toestaat om records te zien die de tag "Klanten" hebben.

Dit geeft je een snel overzicht van of een record wel of niet door je klanten kan worden gezien.

Dit zou nog krachtiger kunnen worden met EN/OF-combinatoren, waarbij je meer complexe regels kunt specificeren. Bijvoorbeeld, je zou een regel kunnen instellen die toegang toestaat tot records die zowel "Klanten" ALS "Openbaar" zijn getagd, of records die ofwel "Intern" OF "Vertrouwelijk" zijn getagd. Dit niveau van flexibiliteit zou voor ongelooflijk genuanceerde instellingen voor rechten zorgen, die zelfs de meest complexe organisatiestructuren en workflows bedienen.

De potentiële toepassingen zijn enorm. Projectmanagers zouden gevoelige informatie gemakkelijk kunnen scheiden, verkoopteams zouden automatisch toegang kunnen hebben tot relevante klantgegevens, en externe samenwerkingspartners zouden naadloos in specifieke delen van een project kunnen worden geïntegreerd zonder het risico van blootstelling aan gevoelige interne informatie.

### Filters voor Aangepaste Velden

Onze visie voor Filters voor Aangepaste Velden vertegenwoordigt een aanzienlijke sprong voorwaarts in gedetailleerde toegangscontrole. Deze functie zal projectbeheerders in staat stellen om te definiëren welke records specifieke gebruikersgroepen kunnen zien op basis van de waarden van aangepaste velden. Het gaat om het creëren van dynamische, datagestuurde grenzen voor informatie-toegang.

Stel je voor dat je rechten kunt instellen zoals:

- Alleen records tonen waar de dropdown "Projectstatus" is ingesteld op "Openbaar"
- Zichtbaarheid beperken tot items waar het multi-select veld "Afdeling" "Marketing" bevat
- Toegang toestaan tot taken waar het selectievakje "Prioriteit" is aangevinkt
- Projecten tonen waar het nummerveld "Budget" boven een bepaalde drempel ligt

### Aanpasbare Projectnavigatie

Dit is simpelweg een uitbreiding van de schakelaars die we al hebben. In plaats van alleen schakelaars voor "activiteit" en "formulieren", willen we dat uitbreiden naar elk onderdeel van de projectnavigatie. Op deze manier kunnen projectbeheerders gefocuste interfaces creëren en tools verwijderen die ze niet nodig hebben.

### Commentaarcontroles

In de toekomst willen we creatief zijn in hoe we onze klanten toestaan om te beslissen wie wel en niet commentaar kan zien. We kunnen meerdere tabbladen voor commentaar onder één record toestaan, en elk kan zichtbaar of niet zichtbaar zijn voor verschillende gebruikersgroepen.

Bovendien kunnen we ook een functie toestaan waarbij alleen commentaar waar een gebruiker *specifiek* wordt genoemd zichtbaar is, en niets anders. Dit zou teams die klanten op projecten hebben in staat stellen om ervoor te zorgen dat alleen commentaar dat ze willen dat klanten zien zichtbaar is.

## Conclusie

Dus daar hebben we het, zo hebben we het bouwen van een van de meest interessante en krachtige functies benaderd! [Zoals je kunt zien op onze projectmanagementvergelijkingstool](/compare), hebben zeer weinig projectmanagementsystemen zo'n krachtige opstelling van rechtenmatrices, en degenen die dat wel doen, reserveren het voor hun duurste ondernemingsplannen, waardoor het ontoegankelijk is voor een typisch klein of middelgroot bedrijf.

Met Blue heb je *alle* functies beschikbaar met ons plan — we geloven niet dat functies van ondernemingsniveau voorbehouden moeten zijn aan ondernemingsklanten!