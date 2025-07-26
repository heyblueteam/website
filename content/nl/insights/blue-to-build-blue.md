---
title:  Hoe we Blue gebruiken om Blue te bouwen. 
description: Leer hoe we ons eigen projectmanagementplatform gebruiken om ons projectmanagementplatform te bouwen!
category: "CEO Blog"
date: 2024-08-07
---



Je staat op het punt een insider-tour te krijgen van hoe Blue Blue bouwt.

Bij Blue gebruiken we onze eigen software.

Dit betekent dat we Blue gebruiken om *Blue* te *bouwen*.

Deze vreemd klinkende term, vaak aangeduid als "dogfooding", wordt vaak toegeschreven aan Paul Maritz, een manager bij Microsoft in de jaren '80. Hij zou een e-mail hebben gestuurd met de onderwerpregel *"Onze eigen hondenvoer eten"* om Microsoft-medewerkers aan te moedigen de producten van het bedrijf te gebruiken.

Het idee om je eigen tools te gebruiken om je tools te bouwen, leidt tot een positieve feedbackcyclus.

Het idee om je eigen tools te gebruiken om je tools te bouwen, leidt tot een positieve feedbackcyclus, wat tal van voordelen met zich meebrengt:

- **Het helpt ons om snel echte bruikbaarheidsproblemen te identificeren.** Aangezien we Blue dagelijks gebruiken, komen we dezelfde uitdagingen tegen als onze gebruikers, waardoor we deze proactief kunnen aanpakken.
- **Het versnelt het ontdekken van bugs.** Intern gebruik onthult vaak bugs voordat ze onze klanten bereiken, wat de algehele productkwaliteit verbetert.
- **Het vergroot onze empathie voor eindgebruikers.** Ons team krijgt uit eerste hand ervaring met de sterke en zwakke punten van Blue, wat ons helpt om meer gebruiksvriendelijke beslissingen te nemen.
- **Het bevordert een cultuur van kwaliteit binnen onze organisatie.** Wanneer iedereen het product gebruikt, is er een gedeeld belang in de excellentie ervan.
- **Het stimuleert innovatie.** Regelmatig gebruik leidt vaak tot ideeën voor nieuwe functies of verbeteringen, waardoor Blue aan de top blijft.

[We hebben eerder gesproken over waarom we geen dedicated testteam hebben](/insights/open-beta) en dit is weer een reden.

Als er bugs in ons systeem zijn, vinden we ze bijna altijd in ons constante dagelijkse gebruik van het platform. En dit creëert ook een drijfveer om ze op te lossen, aangezien we ze uiteraard erg vervelend vinden omdat we waarschijnlijk een van de grootste gebruikers van Blue zijn!

Deze aanpak toont onze toewijding aan het product aan. Door zelf op Blue te vertrouwen, tonen we onze klanten dat we echt geloven in wat we bouwen. Het is niet zomaar een product dat we verkopen – het is een tool waarop we elke dag vertrouwen.

## Hoofdproces

We hebben één project in Blue, toepasselijk genaamd "Product".

**Alles** wat met onze productontwikkeling te maken heeft, wordt hier bijgehouden. Klantfeedback, bugs, functie-ideeën, lopend werk, enzovoort. Het idee om één project te hebben waarin we alles bijhouden, is dat het [betere samenwerking bevordert.](/insights/great-teamwork)

Elke record is een functie of een onderdeel van een functie. Dit is hoe we van "zou het niet leuk zijn als..." naar "kijk eens naar deze geweldige nieuwe functie!" gaan.

Het project heeft de volgende lijsten:

- **Ideeën/Feedback**: Dit is een lijst van teamideeën of klantfeedback op basis van gesprekken of e-mailuitwisselingen. Voel je vrij om hier ideeën toe te voegen! In deze lijst hebben we nog niet besloten dat we een van deze functies zullen bouwen, maar we bekijken dit regelmatig voor ideeën die we verder willen verkennen.
- **Backlog (Langetermijn)**: Dit is waar functies van de Ideeën/Feedback-lijst naartoe gaan als we besluiten dat ze een goede aanvulling op Blue zouden zijn.
- **{Huidig Kwartaal}**: Dit is meestal gestructureerd als "Qx YYYY" en toont onze kwartaalprioriteiten.
- **Bugs**: Dit is een lijst van bekende bugs die door het team of klanten zijn gerapporteerd. Bugs die hier worden toegevoegd, krijgen automatisch het "Bug"-label.
- **Specificaties**: Deze functies worden momenteel gespecificeerd. Niet elke functie vereist een specificatie of ontwerp; het hangt af van de verwachte grootte van de functie en het niveau van vertrouwen dat we hebben met betrekking tot randgevallen en complexiteit.
- **Ontwerp Backlog**: Dit is de backlog voor de ontwerpers; elke keer dat ze iets hebben afgerond dat in uitvoering is, kunnen ze een item uit deze lijst kiezen.
- **In Ontwerp**: Dit zijn de huidige functies die de ontwerpers aan het ontwerpen zijn.
- **Ontwerp Review**: Dit is waar de functies wiens ontwerpen momenteel worden beoordeeld.
- **Backlog (Korte Termijn)**: Dit is een lijst van functies waar we waarschijnlijk in de komende weken aan zullen beginnen. Dit is waar de toewijzingen plaatsvinden. De CEO en Hoofd Engineering beslissen welke functies aan welke ingenieur worden toegewezen op basis van eerdere ervaringen en werklast. [Teamleden kunnen deze vervolgens in de In Progress trekken](/insights/push-vs-pull-kanban) zodra ze hun huidige werk hebben voltooid.
- **In Uitvoering**: Dit zijn functies die momenteel worden ontwikkeld.
- **Code Review**: Zodra een functie is afgerond, ondergaat deze een code review. Dan wordt deze ofwel teruggezet naar "In Uitvoering" als er aanpassingen nodig zijn, of gedeployed naar de ontwikkelomgeving.
- **Dev**: Dit zijn alle functies die momenteel in de ontwikkelomgeving zijn. Andere teamleden en bepaalde klanten kunnen deze bekijken.
- **Beta**: Dit zijn alle functies die momenteel in de [Beta-omgeving](https://beta.app.blue.cc) zijn. Veel klanten gebruiken dit als hun dagelijkse Blue-platform en zullen ook feedback geven.
- **Productie**: Wanneer een functie de productie bereikt, wordt deze als voltooid beschouwd.

Soms, terwijl we een functie ontwikkelen, realiseren we ons dat bepaalde subfuncties moeilijker te implementeren zijn dan aanvankelijk verwacht, en we kunnen ervoor kiezen om ze niet in de initiële versie die we naar klanten uitrollen te doen. In dit geval kunnen we een nieuwe record aanmaken met een naam volgens het formaat "{FeatureName} V2" en alle subfuncties als checklistitems opnemen.

## Labels

- **Mobiel**: Dit betekent dat de functie specifiek is voor onze iOS-, Android- of iPad-apps.
- **{EnterpriseCustomerName}**: Een functie wordt specifiek gebouwd voor een zakelijke klant. Tracking is belangrijk, aangezien er meestal aanvullende commerciële overeenkomsten voor elke functie zijn.
- **Bug**: Dit betekent dat dit een bug is die moet worden opgelost.
- **Snelle Wijziging**: Dit betekent dat dit een Snelle Wijziging is die niet door de volledige releasecyclus hoeft te gaan zoals hierboven beschreven.
- **Hoofd**: Dit is een belangrijke functieontwikkeling. Het is meestal gereserveerd voor belangrijke infrastructuurwerkzaamheden, grote afhankelijkheidsupgrades en significante nieuwe modules binnen Blue.
- **AI**: Deze functie bevat een component van kunstmatige intelligentie.
- **Beveiliging**: Dit betekent dat een beveiligingsimplicatie moet worden beoordeeld of dat er een patch nodig is.

Het snelle wijzigingslabel is bijzonder interessant. Dit is gereserveerd voor kleinere, minder complexe updates die onze volledige releasecyclus niet vereisen en die we binnen 24-48 uur naar klanten willen verzenden.

Snelle wijzigingen zijn meestal kleine aanpassingen die de gebruikerservaring aanzienlijk kunnen verbeteren zonder de kernfunctionaliteit te veranderen. Denk aan het corrigeren van typfouten in de gebruikersinterface, het aanpassen van de padding van knoppen of het toevoegen van nieuwe pictogrammen voor betere visuele begeleiding. Dit zijn de soort wijzigingen die, hoewel klein, een groot verschil kunnen maken in hoe gebruikers ons product waarnemen en ermee omgaan. Ze zijn ook vervelend als ze lang duren om te verzenden!

Ons snelle wijzigingsproces is eenvoudig.

Het begint met het creëren van een nieuwe branch vanuit de hoofdbranch, het implementeren van de wijzigingen en vervolgens het aanmaken van merge-aanvragen voor elke doelbranch - Dev, Beta en Productie. We genereren een previewlink voor beoordeling, zodat zelfs deze kleine wijzigingen aan onze kwaliteitsnormen voldoen. Zodra ze zijn goedgekeurd, worden de wijzigingen gelijktijdig in alle branches samengevoegd, waardoor onze omgevingen synchroon blijven.

## Aangepaste Velden

We hebben niet veel aangepaste velden in ons Productproject.

- **Specificaties**: Dit linkt naar een Blue-document dat de specificatie voor die specifieke functie bevat. Dit wordt niet altijd gedaan, aangezien het afhangt van de complexiteit van de functie.
- **MR**: Dit is de link naar de Merge-aanvraag in [Gitlab](https://gitlab.com) waar we onze code hosten.
- **Preview Link**: Voor functies die voornamelijk de front-end wijzigen, kunnen we een unieke URL maken die die wijzigingen voor elke commit heeft, zodat we de wijzigingen gemakkelijk kunnen beoordelen.
- **Lead**: Dit veld vertelt ons welke senior engineer de leiding heeft over de code review. Het zorgt ervoor dat elke functie de deskundige aandacht krijgt die het verdient, en er is altijd een duidelijke contactpersoon voor vragen of zorgen.

## Checklists

Tijdens onze wekelijkse demo's zullen we de besproken feedback in een checklist genaamd "feedback" plaatsen en er zal ook een andere checklist zijn die de belangrijkste [WBS (Work Breakdown Scope)](/insights/simple-work-breakdown-structure) van de functie bevat, zodat we gemakkelijk kunnen zien wat gedaan is en wat nog gedaan moet worden.

## Conclusie

En dat is het!

We denken dat mensen soms verrast zijn over hoe eenvoudig ons proces is, maar we geloven dat eenvoudige processen vaak veel beter zijn dan te complexe processen die je niet gemakkelijk kunt begrijpen.

Deze eenvoud is opzettelijk. Het stelt ons in staat om wendbaar te blijven, snel te reageren op klantbehoeften en ons hele team op één lijn te houden.

Door Blue te gebruiken om Blue te bouwen, ontwikkelen we niet alleen een product – we leven het.

Dus de volgende keer dat je Blue gebruikt, herinner je dan: je gebruikt niet alleen een product dat we hebben gebouwd. Je gebruikt een product waarop we persoonlijk elke dag vertrouwen.

En dat maakt het verschil.