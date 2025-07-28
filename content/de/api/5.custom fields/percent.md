---
title: Prozentbenutzerdefiniertes Feld
description: Erstellen Sie Prozentfelder, um numerische Werte mit automatischer Handhabung des %-Symbols und Anzeigeformatierung zu speichern
---

Prozentbenutzerdefinierte Felder ermöglichen es Ihnen, Prozentwerte für Datensätze zu speichern. Sie behandeln automatisch das %-Symbol für Eingaben und Anzeigen, während der rohe numerische Wert intern gespeichert wird. Perfekt für Abschlussquoten, Erfolgsquoten oder andere prozentbasierte Kennzahlen.

## Einfaches Beispiel

Erstellen Sie ein einfaches Prozentfeld:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Prozentfeld mit Beschreibung:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
  }) {
    id
    name
    type
    description
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Prozentfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `PERCENT` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Der Projektkontext wird automatisch aus Ihren Authentifizierungsheadern bestimmt. Kein `projectId`-Parameter ist erforderlich.

**Hinweis**: PERCENT-Felder unterstützen keine min/max-Beschränkungen oder Präfixformatierungen wie NUMBER-Felder.

## Prozentwerte festlegen

Prozentfelder speichern numerische Werte mit automatischer Handhabung des %-Symbols:

### Mit Prozentzeichen

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Direkter numerischer Wert

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### SetTodoCustomFieldInput-Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Prozentbenutzerdefinierten Feldes |
| `number` | Float | Nein | Numerischer Prozentwert (z. B. 75,5 für 75,5%) |

## Wertspeicherung und Anzeige

### Speicherformat
- **Interne Speicherung**: Roher numerischer Wert (z. B. 75,5)
- **Datenbank**: Als `Decimal` in der `number`-Spalte gespeichert
- **GraphQL**: Als `Float`-Typ zurückgegeben

### Anzeigeformat
- **Benutzeroberfläche**: Clientanwendungen müssen das %-Symbol anhängen (z. B. "75,5%")
- **Diagramme**: Wird mit %-Symbol angezeigt, wenn der Ausgabetyp PERCENTAGE ist
- **API-Antworten**: Roher numerischer Wert ohne %-Symbol (z. B. 75,5)

## Datensätze mit Prozentwerten erstellen

Beim Erstellen eines neuen Datensatzes mit Prozentwerten:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Unterstützte Eingabeformate

| Format | Beispiel | Ergebnis |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Hinweis**: Das %-Symbol wird automatisch von der Eingabe entfernt und während der Anzeige wieder hinzugefügt.

## Abfragen von Prozentwerten

Beim Abfragen von Datensätzen mit Prozentbenutzerfeldern greifen Sie über den `customField.value.number`-Pfad auf den Wert zu:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

Die Antwort enthält den Prozentsatz als rohe Zahl:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Antwortfelder

### TodoCustomField-Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | ID! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes (enthält den Prozentwert) |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

**Wichtig**: Prozentwerte werden über das `customField.value.number`-Feld abgerufen. Das %-Symbol ist nicht in den gespeicherten Werten enthalten und muss von Clientanwendungen für die Anzeige hinzugefügt werden.

## Filtern und Abfragen

Prozentfelder unterstützen dasselbe Filtern wie NUMBER-Felder:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Unterstützte Operatoren

| Operator | Beschreibung | Beispiel |
|----------|-------------|---------|
| `EQ` | Gleich | `percentage = 75` |
| `NE` | Ungleich | `percentage ≠ 75` |
| `GT` | Größer als | `percentage > 75` |
| `GTE` | Größer oder gleich | `percentage ≥ 75` |
| `LT` | Kleiner als | `percentage < 75` |
| `LTE` | Kleiner oder gleich | `percentage ≤ 75` |
| `IN` | Wert in der Liste | `percentage in [50, 75, 100]` |
| `NIN` | Wert nicht in der Liste | `percentage not in [0, 25]` |
| `IS` | Überprüfung auf null mit `values: null` | `percentage is null` |
| `NOT` | Überprüfung auf nicht null mit `values: null` | `percentage is not null` |

### Bereichsfilterung

Für die Bereichsfilterung verwenden Sie mehrere Operatoren:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Prozentwertbereiche

### Häufige Bereiche

| Bereich | Beschreibung | Anwendungsfall |
|-------|-------------|----------|
| `0-100` | Standardprozent | Completion rates, success rates |
| `0-∞` | Unbegrenztes Prozent | Growth rates, performance metrics |
| `-∞-∞` | Jeder Wert | Change rates, variance |

### Beispielwerte

| Eingabe | Gespeichert | Anzeige |
|-------|--------|---------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Diagrammaggregation

Prozentfelder unterstützen die Aggregation in Dashboard-Diagrammen und Berichten. Verfügbare Funktionen umfassen:

- `AVERAGE` - Durchschnittlicher Prozentwert
- `COUNT` - Anzahl der Datensätze mit Werten
- `MIN` - Niedrigster Prozentwert
- `MAX` - Höchster Prozentwert 
- `SUM` - Summe aller Prozentwerte

Diese Aggregationen sind beim Erstellen von Diagrammen und Dashboards verfügbar, nicht in direkten GraphQL-Abfragen.

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Fehlermeldungen

### Ungültiges Prozentformat
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Keine Zahl
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Best Practices

### Werteingabe
- Erlauben Sie Benutzern die Eingabe mit oder ohne %-Symbol
- Validieren Sie angemessene Bereiche für Ihren Anwendungsfall
- Geben Sie klaren Kontext darüber, was 100% darstellt

### Anzeige
- Zeigen Sie immer das %-Symbol in Benutzeroberflächen an
- Verwenden Sie angemessene Dezimalgenauigkeit
- Berücksichtigen Sie Farbcode für Bereiche (rot/gelb/grün)

### Dateninterpretation
- Dokumentieren Sie, was 100% in Ihrem Kontext bedeutet
- Gehen Sie angemessen mit Werten über 100% um
- Überlegen Sie, ob negative Werte gültig sind

## Häufige Anwendungsfälle

1. **Projektmanagement**
   - Abschlussquoten von Aufgaben
   - Projektfortschritt
   - Ressourcenauslastung
   - Sprintgeschwindigkeit

2. **Leistungsüberwachung**
   - Erfolgsquoten
   - Fehlerquoten
   - Effizienzkennzahlen
   - Qualitätsbewertungen

3. **Finanzkennzahlen**
   - Wachstumsraten
   - Gewinnmargen
   - Rabattbeträge
   - Änderungsprozentsätze

4. **Analytik**
   - Konversionsraten
   - Klickrate
   - Engagement-Kennzahlen
   - Leistungsindikatoren

## Integrationsmerkmale

### Mit Formeln
- Verweisen Sie in Berechnungen auf PERCENT-Felder
- Automatische %-Symbolformatierung in Formel-Ausgaben
- Kombinieren Sie mit anderen numerischen Feldern

### Mit Automatisierungen
- Aktionen basierend auf Prozentgrenzen auslösen
- Benachrichtigungen für Meilensteinprozentsätze senden
- Status basierend auf Abschlussquoten aktualisieren

### Mit Nachschlägen
- Aggregieren Sie Prozentsätze aus verwandten Datensätzen
- Berechnen Sie durchschnittliche Erfolgsquoten
- Finden Sie die am besten/schlechtesten abschneidenden Elemente

### Mit Diagrammen
- Erstellen Sie prozentbasierte Visualisierungen
- Verfolgen Sie den Fortschritt über die Zeit
- Vergleichen Sie Leistungskennzahlen

## Unterschiede zu NUMBER-Feldern

### Was ist anders
- **Eingabeverarbeitung**: Entfernt automatisch das %-Symbol
- **Anzeige**: Fügt automatisch das %-Symbol hinzu
- **Einschränkungen**: Keine min/max-Validierung
- **Formatierung**: Keine Präfixunterstützung

### Was ist gleich
- **Speicherung**: Gleiche Datenbankspalte und -typ
- **Filtern**: Gleiche Abfrageoperatoren
- **Aggregation**: Gleiche Aggregationsfunktionen
- **Berechtigungen**: Gleiches Berechtigungsmodell

## Einschränkungen

- Keine min/max-Wertbeschränkungen
- Keine Präfixformatierungsoptionen
- Keine automatische Validierung des 0-100%-Bereichs
- Keine Umwandlung zwischen Prozentformaten (z. B. 0,75 ↔ 75%)
- Werte über 100% sind erlaubt

## Verwandte Ressourcen

- [Überblick über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte zu benutzerdefinierten Feldern
- [Zahlenbenutzerdefiniertes Feld](/api/custom-fields/number) - Für rohe numerische Werte
- [Automatisierungs-API](/api/automations/index) - Erstellen Sie prozentbasierte Automatisierungen