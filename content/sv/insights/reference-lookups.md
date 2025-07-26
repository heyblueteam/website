---
title: Referens- och uppslagsanpassade fält
description: Skapa enkelt sammanlänkade projekt i Blue, vilket förvandlar det till en enda sanningskälla för ditt företag med de nya Referens- och Uppslagsfälten.
category: "Product Updates"
date: 2023-11-01
---



Projekt i Blue är redan ett kraftfullt sätt att hantera dina affärsdata och driva arbetet framåt.

Idag tar vi nästa logiska steg och låter dig sammanlänka dina data *mellan* projekt för ultimat flexibilitet och kraft.

Att sammanlänka projekt inom Blue förvandlar det till en enda sanningskälla för ditt företag. Denna funktion möjliggör skapandet av en omfattande och sammanlänkad dataset, vilket möjliggör sömlös dataflöde och ökad synlighet över projekt. Genom att länka projekt kan teamen uppnå en enhetlig bild av verksamheten, vilket förbättrar beslutsfattande och operativ effektivitet.

## Ett exempel

Tänk på ACME Company, som använder Blues Referens- och Uppslagsanpassade fält för att skapa ett sammanlänkat dataekosystem över sina projekt för Kunder, Försäljning och Lager. Kundregister i projektet Kunder är länkade via Referensfält till försäljningstransaktioner i projektet Försäljning. Denna länkning gör att Uppslagsfält kan hämta relaterade kunduppgifter, såsom telefonnummer och kontostatusar, direkt in i varje försäljningspost. Dessutom visas sålda lagerartiklar i försäljningsposten genom ett Uppslagsfält som refererar till data om Sålda Kvantiteter från projektet Lager. Slutligen är lageruttag kopplade till relevanta försäljningar via ett Referensfält i Lager, som pekar tillbaka på försäljningsposterna. Denna uppsättning ger full synlighet över vilken försäljning som utlöste lageruttaget, vilket skapar en integrerad 360-graders vy över projekten.

## Hur Referensfält fungerar

Referensanpassade fält gör att du kan skapa relationer mellan poster över olika projekt i Blue. När du skapar ett Referensfält väljer Projektadministratören det specifika projekt som ska tillhandahålla listan över referensposter. Konfigurationsalternativ inkluderar:

* **Enkelval**: Tillåter val av en referenspost.
* **Flerval**: Tillåter val av flera referensposter.
* **Filtrering**: Ställ in filter för att låta användare välja endast poster som matchar filterkriterierna.

När det är inställt kan användare välja specifika poster från rullgardinsmenyn inom Referensfältet, vilket etablerar en länk mellan projekten.

## Utöka referensfält med hjälp av uppslag

Uppslagsanpassade fält används för att importera data från poster i andra projekt, vilket skapar en envägsvisibilitet. De är alltid skrivskyddade och är kopplade till ett specifikt Referensanpassat fält. När en användare väljer en eller flera poster med hjälp av ett Referensanpassat fält, kommer Uppslagsfältet att visa data från dessa poster. Uppslag kan visa data såsom:

* Skapad den
* Uppdaterad den
* Förfallodatum
* Beskrivning
* Lista
* Tagg
* Tilldelad
* Vilket som helst stödd anpassat fält från den refererade posten — inklusive andra uppslagsfält!


Till exempel, föreställ dig ett scenario där du har tre projekt: **Projekt A** är ett försäljningsprojekt, **Projekt B** är ett lagerhanteringsprojekt, och **Projekt C** är ett kundrelationsprojekt. I Projekt A har du ett Referensanpassat fält som länkar försäljningsposter till motsvarande kundposter i Projekt C. I Projekt B har du ett Uppslagsanpassat fält som importerar information från Projekt A, såsom sålda kvantiteter. På så sätt, när en försäljningspost skapas i Projekt A, dras kundinformationen kopplad till den försäljningen automatiskt in från Projekt C, och den sålda kvantiteten dras automatiskt in från Projekt B. Detta gör att du kan hålla all relevant information på ett ställe och se utan att behöva skapa duplicerade data eller manuellt uppdatera poster över projekten.

Ett verkligt exempel på detta är ett e-handelsföretag som använder Blue för att hantera sin försäljning, lager och kundrelationer. I deras **Försäljning**-projekt har de ett Referensanpassat fält som länkar varje försäljningspost till motsvarande **Kund**-post i deras **Kunder**-projekt. I deras **Lager**-projekt har de ett Uppslagsanpassat fält som importerar information från Försäljningsprojektet, såsom sålda kvantiteter, och visar det i lagerposten. Detta gör att de enkelt kan se vilka försäljningar som driver lageruttag och hålla sina lagernivåer uppdaterade utan att behöva manuellt uppdatera poster över projekten.

## Slutsats

Föreställ dig en värld där dina projektdata inte är isolerade utan flödar fritt mellan projekten, vilket ger omfattande insikter och driver effektivitet. Det är kraften i Blues Referens- och Uppslagsfält. Genom att möjliggöra sömlösa datakopplingar och ge realtidsvisibilitet över projekten, förändrar dessa funktioner hur team samarbetar och fattar beslut. Oavsett om du hanterar kundrelationer, spårar försäljning eller övervakar lager, ger Referens- och Uppslagsfält i Blue ditt team möjlighet att arbeta smartare, snabbare och mer effektivt. Dyk in i den sammanlänkade världen av Blue och se din produktivitet skjuta i höjden.

[Kolla in dokumentationen](https://documentation.blue.cc/custom-fields/reference) eller [registrera dig och prova själv.](https://app.blue.cc)