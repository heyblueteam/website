---
title: Skapa återanvändbara checklistor med hjälp av automatiseringar
description: Lär dig hur du skapar automatiseringar för projektledning för återanvändbara checklistor.
category: "Best Practices"
date: 2024-07-08
---



I många projekt och processer kan det vara nödvändigt att använda samma checklista över flera poster eller uppgifter.

Det är dock inte särskilt effektivt att manuellt skriva om checklistan varje gång du vill lägga till den i en post. Här kan du utnyttja [kraftfulla automatiseringar för projektledning](/platform/features/automations) för att automatiskt göra detta åt dig!

Som en påminnelse kräver automatiseringar i Blue två viktiga saker:

1. En Trigger — Vad som ska hända för att starta automatiseringen. Detta kan vara när en post får en specifik tagg, när den flyttas till en specifik 
2. En eller flera Åtgärder — I det här fallet skulle det vara den automatiska skapelsen av en eller flera checklistor.

Låt oss börja med åtgärden först, och sedan diskutera de möjliga triggers som du kan använda.

## Åtgärd för checklistautomatisering

Du kan skapa en ny automatisering, och du kan ställa in en eller flera checklistor som ska skapas, enligt exemplet nedan:

![](/insights/checklist-automation.png)

Detta skulle vara checklistan/ checklistorna som du vill ska skapas varje gång du vidtar åtgärden.

## Triggers för checklistautomatisering

Det finns flera sätt att utlösa skapandet av dina återanvändbara checklistor. Här är några populära alternativ:

- **Lägga till en specifik tagg:** Du kan ställa in automatiseringen för att utlösas när en viss tagg läggs till en post. Till exempel, när taggen "Nytt Projekt" läggs till, kan det automatiskt skapa din checklista för projektinitiering.
- **Posttilldelning:** Checklistskapandet kan utlösas när en post tilldelas en specifik person eller till någon. Detta är användbart för onboarding-checklistor eller uppgiftsspecifika procedurer.
- **Flytta till en specifik lista:** När en post flyttas till en viss lista i din projektstyrningsbräda kan det utlösa skapandet av en relevant checklista. Till exempel, att flytta en punkt till en "Kvalitetssäkring"-lista kan utlösa en QA-checklista.
- **Anpassat kryssruta-fält:** Skapa ett anpassat kryssruta-fält och ställ in automatiseringen för att utlösas när denna ruta är ikryssad. Detta ger dig manuell kontroll över när checklistan ska läggas till.
- **Enkel- eller flervalsanpassade fält:** Du kan skapa ett enkelt eller flervalsanpassat fält med olika alternativ. Varje alternativ kan kopplas till en specifik checklistmall genom separata automatiseringar. Detta möjliggör mer detaljerad kontroll och möjligheten att ha flera checklistmallar redo för olika scenarier.

För att öka kontrollen över vem som kan utlösa dessa automatiseringar kan du dölja dessa anpassade fält från vissa användare med hjälp av anpassade användarroller. Detta säkerställer att endast projektadministratörer eller annan auktoriserad personal kan utlösa dessa alternativ.

Kom ihåg, nyckeln till effektiv användning av återanvändbara checklistor med automatiseringar är att utforma dina triggers genomtänkt. Tänk på ditt teams arbetsflöde, de typer av projekt du hanterar och vem som bör ha möjlighet att initiera olika processer. Med välplanerade automatiseringar kan du avsevärt strömlinjeforma din projektledning och säkerställa konsekvens i dina operationer.

## Användbara resurser

- [Dokumentation för automatisering av projektledning](https://documentation.blue.cc/automations)
- [Dokumentation för anpassade användarroller](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Dokumentation för anpassade fält](https://documentation.blue.cc/custom-fields)