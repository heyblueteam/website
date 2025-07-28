---
title: Datumsbenutzerfeld
description: Erstellen Sie Datumsfelder, um einzelne Daten oder Datumsbereiche mit Zeitzonenunterstützung zu verfolgen
---

Datumsbenutzerfelder ermöglichen es Ihnen, einzelne Daten oder Datumsbereiche für Datensätze zu speichern. Sie unterstützen die Handhabung von Zeitzonen, intelligentes Formatieren und können verwendet werden, um Fristen, Veranstaltungsdaten oder zeitbasierte Informationen zu verfolgen.

## Einfaches Beispiel

Erstellen Sie ein einfaches Datumsfeld:

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

## Fortgeschrittenes Beispiel

Erstellen Sie ein Fälligkeitsdatumfeld mit Beschreibung:

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

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Datumsfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `DATE` sein |
| `isDueDate` | Boolean | Nein | Ob dieses Feld ein Fälligkeitsdatum darstellt |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Benutzerfelder sind automatisch mit dem Projekt verknüpft, basierend auf dem aktuellen Projektkontext des Benutzers. Kein `projectId` Parameter ist erforderlich.

## Datumswerte festlegen

Datumsfelder können entweder ein einzelnes Datum oder einen Datumsbereich speichern:

### Einzelnes Datum

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

### Datumsbereich

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

### Ganztägiges Ereignis

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

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Datumsbenutzerfeldes |
| `startDate` | DateTime | Nein | Startdatum/-uhrzeit im ISO 8601-Format |
| `endDate` | DateTime | Nein | Enddatum/-uhrzeit im ISO 8601-Format |
| `timezone` | String | Nein | Zeitzonenbezeichner (z. B. "America/New_York") |

**Hinweis**: Wenn nur `startDate` angegeben ist, wird `endDate` automatisch auf denselben Wert gesetzt.

## Datumsformate

### ISO 8601 Format
Alle Daten müssen im ISO 8601-Format angegeben werden:
- `2025-01-15T14:30:00Z` - UTC-Zeit
- `2025-01-15T14:30:00+05:00` - Mit Zeitzonenoffset
- `2025-01-15T14:30:00.123Z` - Mit Millisekunden

### Zeitzonenbezeichner
Verwenden Sie standardisierte Zeitzonenbezeichner:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Wenn keine Zeitzone angegeben ist, verwendet das System standardmäßig die erkannte Zeitzone des Benutzers.

## Erstellen von Datensätzen mit Datumswerten

Beim Erstellen eines neuen Datensatzes mit Datumswerten:

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

### Unterstützte Eingabeformate

Beim Erstellen von Datensätzen können Daten in verschiedenen Formaten angegeben werden:

| Format | Beispiel | Ergebnis |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | ID! | Eindeutige Kennung für den Feldwert |
| `uid` | String! | Eindeutige Identifikationszeichenfolge |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes (enthält die Datumswerte) |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

**Wichtig**: Datumswerte (`startDate`, `endDate`, `timezone`) werden über das `customField.value` Feld abgerufen, nicht direkt auf TodoCustomField.

### Wertobjektstruktur

Datumswerte werden über das `customField.value` Feld als JSON-Objekt zurückgegeben:

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

**Hinweis**: Das `value` Feld ist vom `CustomField` Typ, nicht von `TodoCustomField`.

## Abfragen von Datumswerten

Beim Abfragen von Datensätzen mit Datumsbenutzerfeldern greifen Sie auf die Datumswerte über das `customField.value` Feld zu:

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

Die Antwort enthält die Datumswerte im `value` Feld:

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

## Datumsanzeigeintelligenz

Das System formatiert Daten automatisch basierend auf dem Bereich:

| Szenario | Anzeigeformat |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (keine Zeit angezeigt) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Ganztägige Erkennung**: Ereignisse von 00:00 bis 23:59 werden automatisch als ganztägige Ereignisse erkannt.

## Zeitzonenhandhabung

### Speicherung
- Alle Daten werden in UTC in der Datenbank gespeichert
- Zeitzoneninformationen werden separat gespeichert
- Die Umwandlung erfolgt bei der Anzeige

### Beste Praktiken
- Geben Sie immer die Zeitzone für Genauigkeit an
- Verwenden Sie konsistente Zeitzonen innerhalb eines Projekts
- Berücksichtigen Sie die Standorte der Benutzer für globale Teams

### Häufige Zeitzonen

| Region | Zeitzonen-ID | UTC-Offset |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filtern und Abfragen

Datumsfelder unterstützen komplexes Filtern:

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

### Überprüfen auf leere Datumsfelder

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

### Unterstützte Operatoren

| Operator | Verwendung | Beschreibung |
|----------|-------|-------------|
| `EQ` | Mit dateRange | Datum überschneidet sich mit dem angegebenen Bereich (jede Überschneidung) |
| `NE` | Mit dateRange | Datum überschneidet sich nicht mit dem Bereich |
| `IS` | Mit `values: null` | Datumsfeld ist leer (startDate oder endDate ist null) |
| `NOT` | Mit `values: null` | Datumsfeld hat einen Wert (beide Daten sind nicht null) |

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Fehlermeldungen

### Ungültiges Datumsformat
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

### Feld nicht gefunden
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


## Einschränkungen

- Keine Unterstützung für wiederkehrende Daten (verwenden Sie Automatisierungen für wiederkehrende Ereignisse)
- Keine Zeit ohne Datum festlegbar
- Keine eingebaute Berechnung der Arbeitstage
- Datumsbereiche validieren nicht automatisch end > start
- Maximale Genauigkeit ist auf die Sekunde (keine Millisekundenspeicherung)

## Verwandte Ressourcen

- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte zu benutzerdefinierten Feldern
- [Automations-API](/api/automations/index) - Erstellen Sie datumsbasierte Automatisierungen