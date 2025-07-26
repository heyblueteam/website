---
title: Projectmanagementautomatisering — e-mails naar belanghebbenden.
description: Vaak wil je de controle hebben over je projectmanagementautomatiseringen
category: "Product Updates"
date: 2024-07-08
---


We hebben eerder besproken hoe je [e-mailautomatiseringen kunt maken.](/insights/email-automations)

Echter, vaak zijn er belanghebbenden in projecten die alleen moeten worden gewaarschuwd wanneer er iets *echt* belangrijks is. 

Zou het niet fijn zijn als er een projectmanagementautomatisering was waarbij jij, als projectmanager, precies kon bepalen *wanneer* je een belangrijke belanghebbende met de druk op een knop kon waarschuwen?

Nou, het blijkt dat je dit met Blue precies kunt doen! 

Vandaag gaan we leren hoe we een echt nuttige projectmanagementautomatisering kunnen maken: 

Een selectievakje dat automatisch één of meer belangrijke belanghebbenden waarschuwt, en hen alle belangrijke context geeft van waarover je hen waarschuwt. Als bonus leren we ook hoe we deze mogelijkheid kunnen vergrendelen, zodat alleen bepaalde leden van jouw project deze e-mailwaarschuwing kunnen activeren.

Dit zal er ongeveer zo uitzien als je klaar bent:

![](/insights/checkbox-email-automation.png)

En door dit selectievakje aan te vinken, kun je een projectmanagementautomatisering activeren om een aangepaste notificatie-e-mail naar belanghebbenden te sturen. 

Laten we stap voor stap gaan.

## 1. Maak je aangepaste veld voor het selectievakje

Dit is heel eenvoudig, je kunt onze [gedetailleerde documentatie](https://documentation.blue.cc/custom-fields/introduction#creating-custom-fields) bekijken over het maken van aangepaste velden.

Zorg ervoor dat je dit veld een voor de hand liggende naam geeft die je zult onthouden, zoals “beheer waarschuwen” of “belanghebbenden waarschuwen”. 

## 2. Maak je trigger voor projectmanagementautomatisering.

Klik in de recordweergave van je project op de kleine robot rechtsboven om de automatiseringsinstellingen te openen:

<video autoplay loop muted playsinline>
  <source src="/videos/notify-stakeholders-automation-setup.mp4" type="video/mp4">
</video>

## 3. Maak je actie voor projectmanagementautomatisering.

In dit geval zal onze actie zijn om een aangepaste e-mailnotificatie naar één of meer e-mailadressen te sturen. Het is goed om hier op te merken dat deze mensen **niet** in Blue hoeven te zijn om deze e-mails te ontvangen, je kunt e-mails naar *elke* e-mailadres sturen.  

Je kunt meer leren in onze [gedetailleerde documentatiehandleiding over het instellen van e-mailautomatiseringen](https://documentation.blue.cc/automations/actions/email-automations)

Je eindresultaat zou er ongeveer zo uit moeten zien:

![](/insights/email-automation-example.png)

## 4. Bonus: Beperk de toegang tot het selectievakje.

Je kunt [aangepaste gebruikersrollen in Blue](/platform/features/user-permissions) gebruiken om de toegang tot de aangepaste velden voor het selectievakje te beperken, zodat alleen geautoriseerde teamleden e-mailnotificaties kunnen activeren.

Blue stelt projectbeheerders in staat om rollen te definiëren en machtigingen aan gebruikersgroepen toe te wijzen. Dit systeem is cruciaal voor het behouden van controle over wie kan interageren met specifieke elementen van jouw project, inclusief aangepaste velden zoals het notificatie-selectievakje.

1. Navigeer naar de sectie Gebruikersbeheer in Blue en selecteer "Aangepaste gebruikersrollen."
2. Maak een nieuwe rol aan door een beschrijvende naam en een optionele beschrijving te geven.
3. Zoek binnen de rolmachtigingen de sectie voor Toegang tot aangepaste velden.
4. Geef aan of de rol het selectievakje kan bekijken of bewerken. Beperk bijvoorbeeld de bewerkingsrechten tot rollen zoals "Projectbeheerder" terwijl je een nieuw aangemaakte aangepaste rol toestaat om dit veld te beheren.
5. Wijs de nieuw aangemaakte rol toe aan de juiste gebruikers of gebruikersgroepen. Dit zorgt ervoor dat alleen de aangewezen personen de mogelijkheid hebben om met het notificatie-selectievakje te interageren.

[Lees meer op onze officiële documentatiesite.](https://documentation.blue.cc/user-management/roles/custom-user-roles)

Door deze aangepaste rollen te implementeren, verbeter je de beveiliging en integriteit van je projectmanagementprocessen. Alleen geautoriseerde teamleden kunnen kritieke e-mailnotificaties activeren, zodat belanghebbenden belangrijke updates ontvangen zonder onnodige waarschuwingen. 

## Conclusie

Door de hierboven beschreven projectmanagementautomatisering te implementeren, krijg je nauwkeurige controle over wanneer en hoe je belangrijke belanghebbenden waarschuwt. Deze aanpak zorgt ervoor dat belangrijke updates effectief worden gecommuniceerd, zonder je belanghebbenden te overweldigen met onnodige informatie. Door gebruik te maken van de aangepaste velden en automatiseringsfuncties van Blue, kun je je projectmanagementproces stroomlijnen, de communicatie verbeteren en een hoog niveau van efficiëntie behouden.

Met slechts een eenvoudig selectievakje kun je aangepaste e-mailnotificaties activeren die zijn afgestemd op de behoeften van jouw project, zodat de juiste mensen op het juiste moment geïnformeerd zijn. Bovendien voegt de mogelijkheid om deze functionaliteit te beperken tot specifieke teamleden een extra laag van controle en beveiliging toe.

Begin vandaag nog met het benutten van deze krachtige functie in Blue om je belanghebbenden geïnformeerd te houden en je projecten soepel te laten verlopen. Voor meer gedetailleerde stappen en aanvullende aanpassingsopties, verwijs naar de verstrekte documentatielinks. Veel succes met automatiseren!