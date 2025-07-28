---
title: Zoekopdracht Aangepast Veld
description: Maak zoekvelden die automatisch gegevens ophalen uit verwezen records
---

Zoekopdracht aangepaste velden halen automatisch gegevens op uit records die worden verwezen door [Verwijsvelden](/api/custom-fields/reference), en tonen informatie van gekoppelde records zonder handmatig kopiëren. Ze worden automatisch bijgewerkt wanneer de verwezen gegevens veranderen.

## Basisvoorbeeld

Maak een zoekveld om tags van verwezen records weer te geven:

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Geavanceerd Voorbeeld

Maak een zoekveld om aangepaste veldwaarden uit verwezen records te extraheren:

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Invoervelden

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het zoekveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Ja | Zoekconfiguratie |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |

## Zoekconfiguratie

### CustomFieldLookupOptionInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `referenceId` | String! | ✅ Ja | ID van het verwijsveld waaruit gegevens worden opgehaald |
| `lookupId` | String | Nee | ID van het specifieke aangepaste veld om op te zoeken (vereist voor TODO_CUSTOM_FIELD type) |
| `lookupType` | CustomFieldLookupType! | ✅ Ja | Type gegevens dat uit verwezen records moet worden geëxtraheerd |

## Zoektypes

### CustomFieldLookupType Waarden

| Type | Beschrijving | Retourneert |
|------|--------------|-------------|
| `TODO_DUE_DATE` | Vervaldatums van verwezen taken | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Aanmaakdatums van verwezen taken | Array of creation timestamps |
| `TODO_UPDATED_AT` | Laatst bijgewerkte datums van verwezen taken | Array of update timestamps |
| `TODO_TAG` | Tags van verwezen taken | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Toegewezen personen van verwezen taken | Array of user objects |
| `TODO_DESCRIPTION` | Beschrijvingen van verwezen taken | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Namen van takenlijsten van verwezen taken | Array of list titles |
| `TODO_CUSTOM_FIELD` | Aangepaste veldwaarden van verwezen taken | Array of values based on the field type |

## Antwoordvelden

### CustomField Antwoord (voor zoekvelden)

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificatie voor het veld |
| `name` | String! | Weergavenaam van het zoekveld |
| `type` | CustomFieldType! | Zal zijn `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Zoekconfiguratie en resultaten |
| `createdAt` | DateTime! | Wanneer het veld is aangemaakt |
| `updatedAt` | DateTime! | Wanneer het veld voor het laatst is bijgewerkt |

### CustomFieldLookupOption Structuur

| Veld | Type | Beschrijving |
|------|------|--------------|
| `lookupType` | CustomFieldLookupType! | Type zoekopdracht die wordt uitgevoerd |
| `lookupResult` | JSON | De geëxtraheerde gegevens uit verwezen records |
| `reference` | CustomField | Het verwijsveld dat als bron wordt gebruikt |
| `lookup` | CustomField | Het specifieke veld dat wordt opgezocht (voor TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Het bovenliggende zoekveld |
| `parentLookup` | CustomField | Bovenliggend zoekveld in keten (voor geneste zoekopdrachten) |

## Hoe Zoekopdrachten Werken

1. **Gegevensextractie**: Zoekopdrachten extraheren specifieke gegevens uit alle records die zijn gekoppeld via een verwijsveld
2. **Automatische Updates**: Wanneer verwezen records veranderen, worden de zoekwaarden automatisch bijgewerkt
3. **Alleen-lezen**: Zoekvelden kunnen niet direct worden bewerkt - ze weerspiegelen altijd de huidige verwezen gegevens
4. **Geen Berekeningen**: Zoekopdrachten extraheren en tonen gegevens zoals ze zijn zonder aggregaties of berekeningen

## TODO_CUSTOM_FIELD Zoekopdrachten

Bij gebruik van `TODO_CUSTOM_FIELD` type, moet je specificeren welk aangepast veld je wilt extraheren met behulp van de `lookupId` parameter:

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

Dit extraheert de waarden van het opgegeven aangepaste veld uit alle verwezen records.

## Opvragen van Zoekgegevens

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## Voorbeeld Zoekresultaten

### Tag Zoekresultaat
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### Toegewezen Persoon Zoekresultaat
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### Aangepast Veld Zoekresultaat
Resultaten variëren op basis van het aangepaste veldtype dat wordt opgezocht. Bijvoorbeeld, een valuta veld zoekopdracht kan retourneren:
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Belangrijk**: Gebruikers moeten kijkmachtigingen hebben op zowel het huidige project als het verwezen project om zoekresultaten te zien.

## Foutantwoorden

### Ongeldig Verwijsveld
```json
{
  "errors": [{
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Circulaire Zoekopdracht Gedetecteerd
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Ontbrekende Zoek-ID voor TODO_CUSTOM_FIELD
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## Beste Praktijken

1. **Duidelijke Namen**: Gebruik beschrijvende namen die aangeven welke gegevens worden opgezocht
2. **Geschikte Types**: Kies het zoektype dat overeenkomt met je gegevensbehoeften
3. **Prestaties**: Zoekopdrachten verwerken alle verwezen records, dus wees voorzichtig met verwijsvelden met veel koppelingen
4. **Machtigingen**: Zorg ervoor dat gebruikers toegang hebben tot verwezen projecten zodat zoekopdrachten kunnen werken

## Veelvoorkomende Gebruiksscenario's

### Cross-Project Zichtbaarheid
Toon tags, toegewezen personen of statussen van gerelateerde projecten zonder handmatige synchronisatie.

### Afhankelijkheid Tracking
Toon vervaldatums of voltooiingsstatus van taken waarop het huidige werk afhankelijk is.

### Resource Overzicht
Toon alle teamleden die zijn toegewezen aan verwezen taken voor resourceplanning.

### Status Aggregatie
Verzamel alle unieke statussen van gerelateerde taken om de projectgezondheid in één oogopslag te zien.

## Beperkingen

- Zoekvelden zijn alleen-lezen en kunnen niet direct worden bewerkt
- Geen aggregatiefuncties (SOM, AANTAL, GEMIDDELDE) - zoekopdrachten extraheren alleen gegevens
- Geen filteropties - alle verwezen records zijn inbegrepen
- Circulaire zoekketens worden voorkomen om oneindige lussen te vermijden
- Resultaten weerspiegelen huidige gegevens en worden automatisch bijgewerkt

## Gerelateerde Bronnen

- [Verwijsvelden](/api/custom-fields/reference) - Maak koppelingen naar records voor zoekbronnen
- [Aangepaste Veldwaarden](/api/custom-fields/custom-field-values) - Stel waarden in op bewerkbare aangepaste velden
- [Lijst Aangepaste Velden](/api/custom-fields/list-custom-fields) - Vraag alle aangepaste velden in een project op