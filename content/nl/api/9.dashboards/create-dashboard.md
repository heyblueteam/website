---
title: Maak Dashboard
description: Maak een nieuw dashboard voor datavisualisatie en rapportage in Blue
---

## Maak een Dashboard

De `createDashboard` mutatie stelt je in staat om een nieuw dashboard binnen jouw bedrijf of project te creëren. Dashboards zijn krachtige visualisatietools die teams helpen om metrics te volgen, voortgang te monitoren en datagestuurde beslissingen te nemen.

### Basisvoorbeeld

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

### Project-specifiek Dashboard

Maak een dashboard dat is gekoppeld aan een specifiek project:

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

## Invoervelden

### CreateDashboardInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `companyId` | String! | ✅ Ja | De ID van het bedrijf waar het dashboard zal worden aangemaakt |
| `title` | String! | ✅ Ja | De naam van het dashboard. Moet een niet-lege string zijn |
| `projectId` | String | Nee | Optionele ID van een project om aan dit dashboard te koppelen |

## Responsvelden

De mutatie retourneert een compleet `Dashboard` object:

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor het gemaakte dashboard |
| `title` | String! | De dashboardtitel zoals opgegeven |
| `companyId` | String! | Het bedrijf waartoe dit dashboard behoort |
| `projectId` | String | De bijbehorende project-ID (indien opgegeven) |
| `project` | Project | Het bijbehorende projectobject (indien projectId is opgegeven) |
| `createdBy` | User! | De gebruiker die het dashboard heeft gemaakt (jij) |
| `dashboardUsers` | [DashboardUser!]! | Lijst van gebruikers met toegang (aanvankelijk alleen de maker) |
| `createdAt` | DateTime! | Tijdstempel van wanneer het dashboard is aangemaakt |
| `updatedAt` | DateTime! | Tijdstempel van de laatste wijziging (zelfde als createdAt voor nieuwe dashboards) |

### DashboardUser Velden

Wanneer een dashboard wordt aangemaakt, wordt de maker automatisch toegevoegd als dashboardgebruiker:

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de dashboardgebruiker relatie |
| `user` | User! | Het gebruikersobject met toegang tot het dashboard |
| `role` | DashboardRole! | De rol van de gebruiker (maker krijgt volledige toegang) |
| `dashboard` | Dashboard! | Referentie terug naar het dashboard |

## Vereiste Machtigingen

Elke geauthenticeerde gebruiker die tot het opgegeven bedrijf behoort, kan dashboards maken. Er zijn geen speciale rolvereisten.

| Gebruikersstatus | Kan Dashboard Maken |
|------------------|---------------------|
| Company Member | ✅ Ja |
| Niet-Bedrijfslid | ❌ Nee |
| Unauthenticated | ❌ Nee |

## Foutreacties

### Ongeldig Bedrijf
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

### Gebruiker Niet in Bedrijf
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

### Ongeldig Project
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

### Lege Titel
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

## Belangrijke Notities

- **Automatisch eigendom**: De gebruiker die het dashboard aanmaakt, wordt automatisch de eigenaar met volledige machtigingen
- **Projectkoppeling**: Als je een `projectId` opgeeft, moet deze tot hetzelfde bedrijf behoren
- **Initiële machtigingen**: Alleen de maker heeft aanvankelijk toegang. Gebruik `editDashboard` om meer gebruikers toe te voegen
- **Titelvereisten**: Dashboardtitels moeten niet-lege strings zijn. Er is geen uniciteitsvereiste
- **Bedrijfsleden**: Je moet lid zijn van het bedrijf om dashboards daarin te kunnen maken

## Workflow voor Dashboardcreatie

1. **Maak het dashboard** met behulp van deze mutatie
2. **Configureer grafieken en widgets** met behulp van de dashboardbouwer UI
3. **Voeg teamleden toe** met behulp van de `editDashboard` mutatie met `dashboardUsers`
4. **Stel filters en datumbereiken in** via de dashboardinterface
5. **Deel of embed** het dashboard met behulp van de unieke ID

## Gebruikscases

1. **Executive dashboards**: Maak overzichten op hoog niveau van bedrijfsmetrics
2. **Projecttracking**: Bouw project-specifieke dashboards om de voortgang te monitoren
3. **Team prestaties**: Volg de productiviteit en prestatiemetrics van het team
4. **Klantrapportage**: Maak dashboards voor klantgerichte rapporten
5. **Realtime monitoring**: Stel dashboards in voor live operationele gegevens

## Beste Praktijken

1. **Naamgevingsconventies**: Gebruik duidelijke, beschrijvende titels die het doel van het dashboard aangeven
2. **Projectkoppeling**: Koppel dashboards aan projecten wanneer ze project-specifiek zijn
3. **Toegangsbeheer**: Voeg teamleden onmiddellijk na creatie toe voor samenwerking
4. **Organisatie**: Creëer een dashboardhiërarchie met behulp van consistente naamgevingspatronen

## Gerelateerde Operaties

- [Lijst Dashboards](/api/dashboards/) - Haal alle dashboards op voor een bedrijf of project
- [Bewerk Dashboard](/api/dashboards/rename-dashboard) - Hernoem dashboard of beheer gebruikers
- [Kopieer Dashboard](/api/dashboards/copy-dashboard) - Dupliceer een bestaand dashboard
- [Verwijder Dashboard](/api/dashboards/delete-dashboard) - Verwijder een dashboard