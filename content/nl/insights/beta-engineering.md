---
title: Waarom Blue een Open Beta heeft
description: Leer waarom ons projectmanagementsysteem een doorlopende open beta heeft.
category: "Engineering"
date: 2024-08-03
---



Veel B2B SaaS-startups lanceren in Beta, en dat heeft goede redenen. Het is een deel van het traditionele Silicon Valley-motto *“beweeg snel en breek dingen”*.

Een “beta”-label op een product verlaagt de verwachtingen.

Is er iets kapot? Ach, het is maar een beta.

Is het systeem traag? Ach, het is maar een beta.

[De documentatie](https://blue.cc/docs) is niet bestaand? Ach, je snapt het punt.

En dit is *eigenlijk* een goede zaak. Reid Hoffman, de oprichter van LinkedIn, zei ooit:

> Als je je niet schaamt voor de eerste versie van je product, heb je te laat gelanceerd.

En het beta-label is ook goed voor klanten. Het helpt hen om zichzelf te selecteren.

De klanten die beta-producten proberen, bevinden zich in de vroege fasen van de Technology Adoption Lifecycle, ook wel bekend als de Product Adoption Curve.

De Technology Adoption Lifecycle is typisch verdeeld in vijf hoofdsegmenten:

1. Innovators
2. Early Adopters
3. Early Majority
4. Late Majority
5. Laggards

![](/insights/technology-adoption-lifecycle-graph.png)


Uiteindelijk moet het product echter rijpen, en klanten verwachten een stabiel, werkend product. Ze willen geen toegang tot een “beta”-omgeving waar dingen kapot gaan.

Of toch?

*Dit* is de vraag die we onszelf hebben gesteld.

We geloven dat we ons deze vraag hebben gesteld vanwege de manier waarop Blue aanvankelijk is gebouwd. [Blue begon als een tak van een drukke ontwerpstudio](/insights/agency-success-playbook), en dus werkten we *binnen* het kantoor van een bedrijf dat Blue actief gebruikte om al hun projecten te beheren.

Dit betekent dat we jarenlang hebben kunnen observeren hoe *echte* mensen — die naast ons zaten! — Blue in hun dagelijks leven gebruikten.

En omdat ze Blue vanaf de vroege dagen gebruikten, gebruikte dit team altijd Blue Beta!

En dus was het natuurlijk voor ons om al onze andere klanten het ook te laten gebruiken.

**En dit is waarom we geen dedicated testteam hebben.**

Dat klopt.

Niemand bij Blue heeft de *enige* verantwoordelijkheid voor het waarborgen dat ons platform goed en stabiel draait.

Dit is om verschillende redenen.

De eerste is een lagere kostenbasis.

Het niet hebben van een fulltime testteam verlaagt onze kosten aanzienlijk, en we kunnen deze besparingen doorgeven aan onze klanten met de laagste prijzen in de industrie.

Om dit in perspectief te plaatsen, bieden we enterprise-niveau functie sets die onze concurrentie $30-$55/gebruiker/maand kost voor slechts $7/maand.

Dit gebeurt niet per ongeluk, het is *met opzet*.

Het is echter geen goede strategie om een goedkoper product te verkopen als het niet werkt.

Dus de *echte vraag is*, hoe slagen we erin een stabiel platform te creëren dat duizenden klanten kunnen gebruiken zonder een dedicated testteam?

Natuurlijk is onze aanpak van een open Beta cruciaal hiervoor, maar voordat we hierop ingaan, willen we het hebben over de verantwoordelijkheid van ontwikkelaars.

We hebben bij Blue vroeg de beslissing genomen dat we nooit verantwoordelijkheden voor front-end en back-end technologieën zouden splitsen. We zouden alleen full stack ontwikkelaars aannemen of opleiden.

De reden dat we deze beslissing hebben genomen, was om ervoor te zorgen dat een ontwikkelaar volledig eigenaar zou zijn van de functie waaraan ze werkten. Zo zou er geen *“gooi het probleem over het tuinhek”* mentaliteit zijn die je soms krijgt wanneer er gezamenlijke verantwoordelijkheden voor functies zijn.

En dit strekt zich uit tot het testen van de functie, het begrijpen van de klantgebruikscenario's en verzoeken, en het lezen en commentaar geven op de specificaties.

Met andere woorden, elke ontwikkelaar bouwt een diep en intuïtief begrip op van de functie die ze aan het bouwen zijn.

Oké, laten we nu praten over onze open beta.

Wanneer we zeggen dat het “open” is — dan menen we dat. Elke klant kan het proberen door simpelweg “beta” voor onze webapplicatie-URL toe te voegen.

Dus “app.blue.cc” wordt “beta.app.blue.cc”

Wanneer ze dit doen, kunnen ze hun gebruikelijke gegevens zien, aangezien zowel de Beta- als de Productieomgevingen dezelfde database delen, maar ze zullen ook nieuwe functies kunnen zien.

Klanten kunnen gemakkelijk werken, zelfs als sommige teamleden op Productie zitten en andere nieuwsgierige op Beta.

We hebben doorgaans een paar honderd klanten die op elk moment Beta gebruiken, en we plaatsen functievoorbeelden op onze communityforums, zodat ze kunnen zien wat nieuw is en het kunnen uitproberen.

En dit is het punt: we hebben *enkele honderden* testers!

Al deze klanten zullen functies in hun workflows uitproberen en behoorlijk vocal zijn als er iets niet helemaal goed is, omdat ze *al* de functie binnen hun bedrijf implementeren!

De meest voorkomende feedback zijn kleine maar zeer nuttige wijzigingen die randgevallen aanpakken die we niet hadden overwogen.

We laten nieuwe functies 2-4 weken op Beta staan. Wanneer we voelen dat ze stabiel zijn, brengen we ze naar productie.

We hebben ook de mogelijkheid om Beta over te slaan indien nodig, met behulp van een fast-track vlag. Dit wordt meestal gedaan voor bugfixes die we niet willen vasthouden voor 2-4 weken voordat ze naar productie worden verzonden.

Het resultaat?

Naar productie duwen voelt… nou ja, saai! Zoals niets — het is gewoon geen big deal voor ons.

En het betekent dat dit onze releaseplanning versoepelt, wat ons in staat heeft gesteld om [maandelijks functies als een klok te verzenden gedurende de afgelopen zes jaar.](/changelog).

Echter, zoals bij elke keuze, zijn er enkele afwegingen.

Klantenondersteuning is iets complexer, omdat we klanten moeten ondersteunen over twee versies van ons platform. Soms kan dit verwarring veroorzaken bij klanten die teamleden hebben die twee verschillende versies gebruiken.

Een ander pijnpunt is dat deze aanpak soms de algehele releaseplanning naar productie kan vertragen. Dit is vooral waar voor grotere functies die “vast kunnen komen te zitten” in Beta als er een andere gerelateerde functie is die problemen heeft en verdere werkzaamheden nodig heeft.

Maar in balans denken we dat deze afwegingen de voordelen van een lagere kostenbasis en grotere klantbetrokkenheid waard zijn.

We zijn een van de weinige softwarebedrijven die deze aanpak omarmen, maar het is nu een fundamenteel onderdeel van onze productontwikkelingsaanpak.