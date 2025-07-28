---
title: Sök anpassat fält
description: Skapa sökfält som automatiskt hämtar data från refererade poster
---

Sök anpassade fält hämtar automatiskt data från poster som refereras av [Referensfält](/api/custom-fields/reference), och visar information från länkade poster utan manuell kopiering. De uppdateras automatiskt när refererad data ändras.

## Grundläggande exempel

Skapa ett sökfält för att visa taggar från refererade poster:

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

## Avancerat exempel

Skapa ett sökfält för att extrahera värden för anpassade fält från refererade poster:

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

## Indata parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för sökfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Ja | Sök konfiguration |
| `description` | String | Nej | Hjälptext som visas för användare |

## Sök konfiguration

### CustomFieldLookupOptionInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `referenceId` | String! | ✅ Ja | ID för referensfältet att hämta data från |
| `lookupId` | String | Nej | ID för det specifika anpassade fältet att söka (obligatoriskt för TODO_CUSTOM_FIELD-typ) |
| `lookupType` | CustomFieldLookupType! | ✅ Ja | Typ av data att extrahera från refererade poster |

## Söktyper

### CustomFieldLookupType Värden

| Typ | Beskrivning | Returnerar |
|------|-------------|---------|
| `TODO_DUE_DATE` | Förfallodatum från refererade att-göra-poster | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Skapelsedatum från refererade att-göra-poster | Array of creation timestamps |
| `TODO_UPDATED_AT` | Senast uppdaterade datum från refererade att-göra-poster | Array of update timestamps |
| `TODO_TAG` | Taggar från refererade att-göra-poster | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Tilldelade från refererade att-göra-poster | Array of user objects |
| `TODO_DESCRIPTION` | Beskrivningar från refererade att-göra-poster | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Att-göra-listnamn från refererade att-göra-poster | Array of list titles |
| `TODO_CUSTOM_FIELD` | Värden för anpassade fält från refererade att-göra-poster | Array of values based on the field type |

## Svarsfält

### CustomField Svar (för sökfält)

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältet |
| `name` | String! | Visningsnamn för sökfältet |
| `type` | CustomFieldType! | Kommer att vara `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Sök konfiguration och resultat |
| `createdAt` | DateTime! | När fältet skapades |
| `updatedAt` | DateTime! | När fältet senast uppdaterades |

### CustomFieldLookupOption Struktur

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | Typ av sökning som utförs |
| `lookupResult` | JSON | Den extraherade datan från refererade poster |
| `reference` | CustomField | Referensfältet som används som källa |
| `lookup` | CustomField | Det specifika fältet som söks (för TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Det överordnade sökfältet |
| `parentLookup` | CustomField | Överordnat sökfält i kedjan (för nästlade sökningar) |

## Hur sökningar fungerar

1. **Dataextraktion**: Sökningar extraherar specifik data från alla poster som är kopplade genom ett referensfält
2. **Automatiska uppdateringar**: När refererade poster ändras, uppdateras sökvärden automatiskt
3. **Skrivskyddat**: Sökfält kan inte redigeras direkt - de återspeglar alltid aktuell refererad data
4. **Inga beräkningar**: Sökningar extraherar och visar data som den är utan aggregeringar eller beräkningar

## TODO_CUSTOM_FIELD Sökningar

När du använder `TODO_CUSTOM_FIELD`-typ, måste du specificera vilket anpassat fält som ska extraheras med hjälp av `lookupId`-parametern:

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

Detta extraherar värdena för det angivna anpassade fältet från alla refererade poster.

## Fråga sökdata

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

## Exempel på sökresultat

### Tagg Sökresultat
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

### Tilldelad Sökresultat
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

### Sökresultat för anpassat fält
Resultaten varierar beroende på vilken typ av anpassat fält som söks. Till exempel kan en sökning av ett valutafält returnera:
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

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk behörighet |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Viktigt**: Användare måste ha visningsbehörighet för både det aktuella projektet och det refererade projektet för att se sökresultaten.

## Felrespons

### Ogiltigt referensfält
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

### Cirkulär sökning upptäcktes
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

### Saknas sök-ID för TODO_CUSTOM_FIELD
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

## Bästa praxis

1. **Tydlig namngivning**: Använd beskrivande namn som indikerar vilken data som söks
2. **Lämpliga typer**: Välj den söktyp som matchar dina databehov
3. **Prestanda**: Sökningar bearbetar alla refererade poster, så var medveten om referensfält med många länkar
4. **Behörigheter**: Se till att användare har tillgång till refererade projekt för att sökningar ska fungera

## Vanliga användningsfall

### Tvärprojektssynlighet
Visa taggar, tilldelningar eller statusar från relaterade projekt utan manuell synkronisering.

### Beroendespårning
Visa förfallodatum eller slutförandestatus för uppgifter som det aktuella arbetet är beroende av.

### Resursöversikt
Visa alla teammedlemmar som är tilldelade refererade uppgifter för resursplanering.

### Statusaggregat
Samla alla unika statusar från relaterade uppgifter för att se projektets hälsa vid en ögonkast.

## Begränsningar

- Sökfält är skrivskyddade och kan inte redigeras direkt
- Inga aggregationsfunktioner (SUM, COUNT, AVG) - sökningar extraherar endast data
- Inga filtreringsalternativ - alla refererade poster ingår
- Cirkulära sökningskedjor förhindras för att undvika oändliga loopar
- Resultaten återspeglar aktuell data och uppdateras automatiskt

## Relaterade resurser

- [Referensfält](/api/custom-fields/reference) - Skapa länkar till poster för sökkällor
- [Värden för anpassade fält](/api/custom-fields/custom-field-values) - Ställ in värden på redigerbara anpassade fält
- [Lista anpassade fält](/api/custom-fields/list-custom-fields) - Fråga alla anpassade fält i ett projekt