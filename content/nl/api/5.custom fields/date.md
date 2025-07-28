---
title: Datum Aangepast Veld
description: Maak datumvelden om enkele datums of datumbereiken bij te houden met ondersteuning voor tijdzones
---

Datum aangepaste velden stellen je in staat om enkele datums of datumbereiken voor records op te slaan. Ze ondersteunen tijdzonebeheer, intelligente opmaak en kunnen worden gebruikt om deadlines, evenementdatums of andere tijdgebonden informatie bij te houden.

## Basisvoorbeeld

Maak een eenvoudig datumveld aan:

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een vervaldatumveld met beschrijving:

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## Invoervelden

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het datumveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `DATE` |
| `isDueDate` | Boolean | Nee | Of dit veld een vervaldatum vertegenwoordigt |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |

**Opmerking**: Aangepaste velden worden automatisch gekoppeld aan het project op basis van de huidige projectcontext van de gebruiker. Geen `projectId` parameter is vereist.

## Datumwaarden Instellen

Datumvelden kunnen een enkele datum of een datumbereik opslaan:

### Enkele Datum

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Datumbereik

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Hele Dag Evenement

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het datum aangepaste veld |
| `startDate` | DateTime | Nee | Startdatum/tijd in ISO 8601-indeling |
| `endDate` | DateTime | Nee | Einddatum/tijd in ISO 8601-indeling |
| `timezone` | String | Nee | Tijdzone-identificator (bijv. "America/New_York") |

**Opmerking**: Als alleen `startDate` wordt opgegeven, wordt `endDate` automatisch ingesteld op dezelfde waarde.

## Datumformaten

### ISO 8601 Indeling
Alle datums moeten worden opgegeven in ISO 8601-indeling:
- `2025-01-15T14:30:00Z` - UTC tijd
- `2025-01-15T14:30:00+05:00` - Met tijdzone-offset
- `2025-01-15T14:30:00.123Z` - Met milliseconden

### Tijdzone Identificatoren
Gebruik standaard tijdzone-identificatoren:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Als er geen tijdzone wordt opgegeven, valt het systeem terug op de gedetecteerde tijdzone van de gebruiker.

## Records Maken met Datumwaarden

Bij het maken van een nieuw record met datumwaarden:

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### Ondersteunde Invoervormen

Bij het maken van records kunnen datums in verschillende indelingen worden opgegeven:

| Indeling | Voorbeeld | Resultaat |
|----------|-----------|-----------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | ID! | Unieke identificatie voor de veldwaarde |
| `uid` | String! | Unieke identificatiestring |
| `customField` | CustomField! | De definitie van het aangepaste veld (bevat de datumwaarden) |
| `todo` | Todo! | Het record waar deze waarde bij hoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

**Belangrijk**: Datumwaarden (`startDate`, `endDate`, `timezone`) worden benaderd via het `customField.value` veld, niet direct op TodoCustomField.

### Waarde Objectstructuur

Datumwaarden worden teruggegeven via het `customField.value` veld als een JSON-object:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**Opmerking**: Het `value` veld is van het `CustomField` type, niet van `TodoCustomField`.

## Datumwaarden Opvragen

Bij het opvragen van records met datum aangepaste velden, krijg je toegang tot de datumwaarden via het `customField.value` veld:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

De reactie bevat de datumwaarden in het `value` veld:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## Datumweergave Intelligentie

Het systeem formatteert datums automatisch op basis van het bereik:

| Scenario | Weergaveformaat |
|----------|-----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (geen tijd weergegeven) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Hele dag detectie**: Evenementen van 00:00 tot 23:59 worden automatisch gedetecteerd als hele dagen evenementen.

## Tijdzonebeheer

### Opslag
- Alle datums worden in UTC in de database opgeslagen
- Tijdzone-informatie wordt apart bewaard
- Conversie gebeurt bij weergave

### Beste Praktijken
- Geef altijd een tijdzone op voor nauwkeurigheid
- Gebruik consistente tijdzones binnen een project
- Houd rekening met de locaties van gebruikers voor wereldwijde teams

### Veelvoorkomende Tijdzones

| Regio | Tijdzone ID | UTC Offset |
|-------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filteren en Opvragen

Datumvelden ondersteunen complexe filtering:

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### Controleren op Lege Datumvelden

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### Ondersteunde Operators

| Operator | Gebruik | Beschrijving |
|----------|---------|--------------|
| `EQ` | Met dateRange | Datum overlapt met opgegeven bereik (elke intersectie) |
| `NE` | Met dateRange | Datum overlapt niet met bereik |
| `IS` | Met `values: null` | Datumveld is leeg (startDate of endDate is null) |
| `NOT` | Met `values: null` | Datumveld heeft een waarde (beide datums zijn niet null) |

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Foutantwoorden

### Ongeldig Datumformaat
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Veld Niet Gevonden
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```


## Beperkingen

- Geen ondersteuning voor terugkerende datums (gebruik automatiseringen voor terugkerende evenementen)
- Tijd kan niet worden ingesteld zonder datum
- Geen ingebouwde berekening van werkdagen
- Datumbereiken valideren niet automatisch eind > start
- Maximale precisie is tot de seconde (geen milliseconde opslag)

## Gerelateerde Bronnen

- [Overzicht Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten van aangepaste velden
- [Automatiseringen API](/api/automations/index) - Maak datumgebaseerde automatiseringen