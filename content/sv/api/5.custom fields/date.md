---
title: Datum Anpassat Fält
description: Skapa datumfält för att spåra enskilda datum eller datumintervall med tidszonsstöd
---

Datum anpassade fält gör att du kan lagra enskilda datum eller datumintervall för poster. De stödjer hantering av tidszoner, intelligent formatering och kan användas för att spåra deadlines, evenemangsdatum eller annan tidsbaserad information.

## Grundläggande Exempel

Skapa ett enkelt datumfält:

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

## Avancerat Exempel

Skapa ett förfallodatumfält med beskrivning:

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

## Indata Parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för datumfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `DATE` |
| `isDueDate` | Boolean | Nej | Om detta fält representerar ett förfallodatum |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: Anpassade fält är automatiskt kopplade till projektet baserat på användarens aktuella projektkontext. Ingen `projectId` parameter krävs.

## Ställa in Datumvärden

Datumfält kan lagra antingen ett enskilt datum eller ett datumintervall:

### Enskilt Datum

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

### Datumintervall

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

### Hela Dagen Evenemang

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

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade datumfältet |
| `startDate` | DateTime | Nej | Startdatum/tid i ISO 8601-format |
| `endDate` | DateTime | Nej | Slutdatum/tid i ISO 8601-format |
| `timezone` | String | Nej | Tidszonsidentifierare (t.ex. "America/New_York") |

**Notera**: Om endast `startDate` anges, kommer `endDate` automatiskt att anta samma värde.

## Datumformat

### ISO 8601 Format
Alla datum måste anges i ISO 8601-format:
- `2025-01-15T14:30:00Z` - UTC tid
- `2025-01-15T14:30:00+05:00` - Med tidszonsförskjutning
- `2025-01-15T14:30:00.123Z` - Med millisekunder

### Tidszonsidentifierare
Använd standard tidszonsidentifierare:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Om ingen tidszon anges, kommer systemet automatiskt att använda användarens upptäckta tidszon.

## Skapa Poster med Datumvärden

När du skapar en ny post med datumvärden:

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

### Stödda Indataformat

När du skapar poster kan datum anges i olika format:

| Format | Exempel | Resultat |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | ID! | Unik identifierare för fältvärdet |
| `uid` | String! | Unik identifierarsträng |
| `customField` | CustomField! | Definition av det anpassade fältet (innehåller datumvärdena) |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

**Viktigt**: Datumvärden (`startDate`, `endDate`, `timezone`) nås genom `customField.value` fältet, inte direkt på TodoCustomField.

### Värde Objekt Struktur

Datumvärden returneras genom `customField.value` fältet som ett JSON-objekt:

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

**Notera**: Fältet `value` är av `CustomField` typ, inte av `TodoCustomField`.

## Fråga Datumvärden

När du frågar poster med datum anpassade fält, nå datumvärdena genom `customField.value` fältet:

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

Svaret kommer att inkludera datumvärdena i `value` fältet:

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

## Datumvisningsintelligens

Systemet formaterar automatiskt datum baserat på intervallet:

| Scenario | Visningsformat |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (ingen tid visas) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Hela dagen detektion**: Evenemang från 00:00 till 23:59 upptäcks automatiskt som hela dagen evenemang.

## Tidszonsbehandling

### Lagring
- Alla datum lagras i UTC i databasen
- Tidszonsinformation bevaras separat
- Konvertering sker vid visning

### Bästa Praxis
- Ange alltid tidszon för noggrannhet
- Använd konsekventa tidszoner inom ett projekt
- Tänk på användarens platser för globala team

### Vanliga Tidszoner

| Region | Tidszons-ID | UTC Offset |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filtrering och Fråga

Datumfält stödjer komplex filtrering:

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

### Kontrollera Tomma Datumfält

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

### Stödda Operatörer

| Operatör | Användning | Beskrivning |
|----------|------------|-------------|
| `EQ` | Med dateRange | Datum överlappar med angivet intervall (vilken som helst korsning) |
| `NE` | Med dateRange | Datum överlappar inte med intervallet |
| `IS` | Med `values: null` | Datumfältet är tomt (startDate eller endDate är null) |
| `NOT` | Med `values: null` | Datumfältet har ett värde (båda datumen är inte null) |

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|------------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Fel Svar

### Ogiltigt Datumformat
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

### Fält Inte Hittat
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


## Begränsningar

- Ingen stöd för återkommande datum (använd automatiseringar för återkommande evenemang)
- Kan inte ställa in tid utan datum
- Ingen inbyggd beräkning av arbetsdagar
- Datumintervall validerar inte slut > start automatiskt
- Maximal precision är till sekunden (ingen millisekund lagring)

## Relaterade Resurser

- [Översikt över Anpassade Fält](/api/custom-fields/list-custom-fields) - Allmänna koncept för anpassade fält
- [Automations API](/api/automations/index) - Skapa datumbaserade automatiseringar