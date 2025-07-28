---
title: Skapa Instrumentpanel
description: Skapa en ny instrumentpanel för datavisualisering och rapportering i Blue
---

## Skapa en Instrumentpanel

Mutation `createDashboard` gör att du kan skapa en ny instrumentpanel inom ditt företag eller projekt. Instrumentpaneler är kraftfulla visualiseringsverktyg som hjälper team att spåra mätvärden, övervaka framsteg och fatta datadrivna beslut.

### Grundläggande Exempel

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### Projektspecifik Instrumentpanel

Skapa en instrumentpanel kopplad till ett specifikt projekt:

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## Inmatningsparametrar

### CreateDashboardInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `companyId` | String! | ✅ Ja | ID:t för företaget där instrumentpanelen kommer att skapas |
| `title` | String! | ✅ Ja | Namnet på instrumentpanelen. Måste vara en icke-tom sträng |
| `projectId` | String | Nej | Valfritt ID för ett projekt att koppla till denna instrumentpanel |

## Svarsfält

Mutation returnerar ett komplett `Dashboard` objekt:

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unikt identifierare för den skapade instrumentpanelen |
| `title` | String! | Instrumentpanelens titel som angivits |
| `companyId` | String! | Företaget som denna instrumentpanel tillhör |
| `projectId` | String | Det kopplade projektets ID (om angivet) |
| `project` | Project | Det kopplade projektobjektet (om projectId angavs) |
| `createdBy` | User! | Användaren som skapade instrumentpanelen (du) |
| `dashboardUsers` | [DashboardUser!]! | Lista över användare med åtkomst (initialt bara skaparen) |
| `createdAt` | DateTime! | Tidsstämpel för när instrumentpanelen skapades |
| `updatedAt` | DateTime! | Tidsstämpel för senaste ändring (samma som createdAt för nya instrumentpaneler) |

### DashboardUser Fält

När en instrumentpanel skapas läggs skaparen automatiskt till som en instrumentpanelsanvändare:

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unikt identifierare för relationen mellan instrumentpanelsanvändare |
| `user` | User! | Användarobjektet med åtkomst till instrumentpanelen |
| `role` | DashboardRole! | Användarens roll (skaparen får full åtkomst) |
| `dashboard` | Dashboard! | Referens tillbaka till instrumentpanelen |

## Obligatoriska Behörigheter

Alla autentiserade användare som tillhör det angivna företaget kan skapa instrumentpaneler. Det finns inga speciella rollkrav.

| Användarstatus | Kan Skapa Instrumentpanel |
|----------------|--------------------------|
| Company Member | ✅ Ja |
| Icke-Företagsmedlem | ❌ Nej |
| Unauthenticated | ❌ Nej |

## Felrespons

### Ogiltigt Företag
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Användare Inte i Företaget
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Ogiltigt Projekt
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Tom Titel
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Viktiga Anteckningar

- **Automatisk äganderätt**: Användaren som skapar instrumentpanelen blir automatiskt dess ägare med fullständiga behörigheter
- **Projektkoppling**: Om du anger en `projectId`, måste den tillhöra samma företag
- **Initiala behörigheter**: Endast skaparen har åtkomst initialt. Använd `editDashboard` för att lägga till fler användare
- **Titelkrav**: Instrumentpanelstitlar måste vara icke-tomma strängar. Det finns inget krav på unikhet
- **Företagsmedlemskap**: Du måste vara medlem i företaget för att skapa instrumentpaneler i det

## Arbetsflöde för Instrumentpanelsskapande

1. **Skapa instrumentpanelen** med denna mutation
2. **Konfigurera diagram och widgets** med hjälp av instrumentpanelens byggargränssnitt
3. **Lägg till teammedlemmar** med hjälp av `editDashboard` mutation med `dashboardUsers`
4. **Ställ in filter och datumintervall** genom instrumentpanelens gränssnitt
5. **Dela eller bädda in** instrumentpanelen med dess unika ID

## Användningsfall

1. **Verkställande instrumentpaneler**: Skapa övergripande översikter av företagsmätvärden
2. **Projektspårning**: Bygg projektspecifika instrumentpaneler för att övervaka framsteg
3. **Teamets prestation**: Spåra teamets produktivitet och prestationsmätvärden
4. **Kundrapportering**: Skapa instrumentpaneler för kundinriktade rapporter
5. **Real-tidsövervakning**: Ställ in instrumentpaneler för live operativa data

## Bästa Praxis

1. **Namngivningskonventioner**: Använd tydliga, beskrivande titlar som anger instrumentpanelens syfte
2. **Projektkoppling**: Koppla instrumentpaneler till projekt när de är projektspecifika
3. **Åtkomsthantering**: Lägg till teammedlemmar omedelbart efter skapandet för samarbete
4. **Organisation**: Skapa en hierarki av instrumentpaneler med konsekventa namnmönster

## Relaterade Åtgärder

- [Lista Instrumentpaneler](/api/dashboards/) - Hämta alla instrumentpaneler för ett företag eller projekt
- [Redigera Instrumentpanel](/api/dashboards/rename-dashboard) - Byt namn på instrumentpanelen eller hantera användare
- [Kopiera Instrumentpanel](/api/dashboards/copy-dashboard) - Duplicera en befintlig instrumentpanel
- [Radera Instrumentpanel](/api/dashboards/delete-dashboard) - Ta bort en instrumentpanel