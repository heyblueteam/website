---
title: Zahlenbenutzerfeld
description: Erstellen Sie Zahlenfelder, um numerische Werte mit optionalen Min/Max-Beschränkungen und Präfixformatierung zu speichern
---

Zahlenbenutzerfelder ermöglichen es Ihnen, numerische Werte für Datensätze zu speichern. Sie unterstützen Validierungsbeschränkungen, Dezimalgenauigkeit und können für Mengen, Punktzahlen, Messungen oder beliebige numerische Daten verwendet werden, die keine spezielle Formatierung erfordern.

## Einfaches Beispiel

Erstellen Sie ein einfaches Zahlenfeld:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Zahlenfeld mit Beschränkungen und Präfix:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Zahlenfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `NUMBER` sein |
| `projectId` | String! | ✅ Ja | ID des Projekts, in dem das Feld erstellt werden soll |
| `min` | Float | Nein | Mindestwertbeschränkung (nur UI-Hinweis) |
| `max` | Float | Nein | Höchstwertbeschränkung (nur UI-Hinweis) |
| `prefix` | String | Nein | Anzeigepräfix (z.B. "#", "~", "$") |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

## Zahlenwerte festlegen

Zahlenfelder speichern Dezimalwerte mit optionaler Validierung:

### Einfacher Zahlenwert

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Ganzzahlwert

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput-Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des Datensatzes, der aktualisiert werden soll |
| `customFieldId` | String! | ✅ Ja | ID des Zahlenbenutzerfeldes |
| `number` | Float | Nein | Numerischer Wert, der gespeichert werden soll |

## Wertbeschränkungen

### Min/Max-Beschränkungen (UI-Hinweis)

**Wichtig**: Min/Max-Beschränkungen werden gespeichert, aber NICHT serverseitig durchgesetzt. Sie dienen als UI-Hinweis für Frontend-Anwendungen.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Clientseitige Validierung erforderlich**: Frontend-Anwendungen müssen eine Validierungslogik implementieren, um Min/Max-Beschränkungen durchzusetzen.

### Unterstützte Werttypen

| Typ | Beispiel | Beschreibung |
|------|---------|-------------|
| Integer | `42` | Ganze Zahlen |
| Decimal | `42.5` | Zahlen mit Dezimalstellen |
| Negative | `-10` | Negative Werte (wenn keine Min-Beschränkung) |
| Zero | `0` | Nullwert |

**Hinweis**: Min/Max-Beschränkungen werden NICHT serverseitig validiert. Werte außerhalb des angegebenen Bereichs werden akzeptiert und gespeichert.

## Erstellen von Datensätzen mit Zahlenwerten

Beim Erstellen eines neuen Datensatzes mit Zahlenwerten:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### Unterstützte Eingabeformate

Beim Erstellen von Datensätzen verwenden Sie den `number`-Parameter (nicht `value`) im Array der benutzerdefinierten Felder:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Antwortfelder

### TodoCustomField-Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `number` | Float | Der numerische Wert |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

### CustomField-Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für die Felddefinition |
| `name` | String! | Anzeigename des Feldes |
| `type` | CustomFieldType! | Immer `NUMBER` |
| `min` | Float | Mindestzulässiger Wert |
| `max` | Float | Höchstzulässiger Wert |
| `prefix` | String | Anzeigepräfix |
| `description` | String | Hilfetext |

**Hinweis**: Wenn der Zahlenwert nicht festgelegt ist, wird das `number`-Feld auf `null` gesetzt.

## Filtern und Abfragen

Zahlenfelder unterstützen umfassendes numerisches Filtern:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
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
| `EQ` | Gleich | `number = 42` |
| `NE` | Ungleich | `number ≠ 42` |
| `GT` | Größer als | `number > 42` |
| `GTE` | Größer oder gleich | `number ≥ 42` |
| `LT` | Kleiner als | `number < 42` |
| `LTE` | Kleiner oder gleich | `number ≤ 42` |
| `IN` | In Array | `number in [1, 2, 3]` |
| `NIN` | Nicht im Array | `number not in [1, 2, 3]` |
| `IS` | Ist null/nicht null | `number is null` |

### Bereichsfilterung

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Anzeigeformatierung

### Mit Präfix

Wenn ein Präfix festgelegt ist, wird es angezeigt:

| Wert | Präfix | Anzeige |
|-------|--------|---------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Dezimalgenauigkeit

Zahlen behalten ihre Dezimalgenauigkeit:

| Eingabe | Gespeichert | Angezeigt |
|-------|--------|-----------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|---------------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Fehlermeldungen

### Ungültiges Zahlenformat
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Hinweis**: Min/Max-Validierungsfehler treten NICHT serverseitig auf. Die Validierung der Beschränkungen muss in Ihrer Frontend-Anwendung implementiert werden.

### Ist keine Zahl
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

## Beste Praktiken

### Beschränkungsdesign
- Setzen Sie realistische Min/Max-Werte als UI-Hinweis
- Implementieren Sie clientseitige Validierung, um Beschränkungen durchzusetzen
- Verwenden Sie Beschränkungen, um Benutzerrückmeldungen in Formularen zu geben
- Überlegen Sie, ob negative Werte für Ihren Anwendungsfall gültig sind

### Wertgenauigkeit
- Verwenden Sie die geeignete Dezimalgenauigkeit für Ihre Bedürfnisse
- Berücksichtigen Sie Rundungen zu Anzeigezwecken
- Seien Sie konsistent mit der Genauigkeit über verwandte Felder hinweg

### Anzeigeverbesserung
- Verwenden Sie bedeutungsvolle Präfixe für den Kontext
- Berücksichtigen Sie Einheiten in Feldnamen (z.B. "Gewicht (kg)")
- Geben Sie klare Beschreibungen für Validierungsregeln an

## Häufige Anwendungsfälle

1. **Bewertungssysteme**
   - Leistungsbewertungen
   - Qualitätsbewertungen
   - Prioritätsstufen
   - Kundenzufriedenheitsbewertungen

2. **Messungen**
   - Mengen und Beträge
   - Abmessungen und Größen
   - Zeitdauern (im numerischen Format)
   - Kapazitäten und Grenzen

3. **Geschäftskennzahlen**
   - Umsatzzahlen
   - Konversionsraten
   - Budgetzuweisungen
   - Zielzahlen

4. **Technische Daten**
   - Versionsnummern
   - Konfigurationswerte
   - Leistungskennzahlen
   - Schwellenwerteinstellungen

## Integrationsfunktionen

### Mit Diagrammen und Dashboards
- Verwenden Sie ZAHL-Felder in Diagramm-Berechnungen
- Erstellen Sie numerische Visualisierungen
- Verfolgen Sie Trends im Laufe der Zeit

### Mit Automatisierungen
- Auslösen von Aktionen basierend auf Zahlen-Schwellenwerten
- Aktualisieren verwandter Felder basierend auf Zahlenänderungen
- Benachrichtigungen für spezifische Werte senden

### Mit Nachschlagen
- Aggregieren von Zahlen aus verwandten Datensätzen
- Berechnen von Summen und Durchschnitten
- Finden von Min/Max-Werten über Beziehungen hinweg

### Mit Diagrammen
- Erstellen von numerischen Visualisierungen
- Verfolgen von Trends im Laufe der Zeit
- Vergleichen von Werten über Datensätze hinweg

## Einschränkungen

- **Keine serverseitige Validierung** von Min/Max-Beschränkungen
- **Clientseitige Validierung erforderlich** zur Durchsetzung von Beschränkungen
- Keine integrierte Währungsformatierung (verwenden Sie stattdessen den WÄHRUNG-Typ)
- Kein automatisches Prozentzeichen (verwenden Sie stattdessen den PROZENT-Typ)
- Keine Einheitenumwandlungsfunktionen
- Dezimalgenauigkeit durch den Datenbank-Dezimaltyp begrenzt
- Keine mathematische Formelbewertung im Feld selbst

## Verwandte Ressourcen

- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/1.index) - Allgemeine Konzepte zu benutzerdefinierten Feldern
- [Währungsbenutzerfeld](/api/custom-fields/currency) - Für Geldwerte
- [Prozentbenutzerfeld](/api/custom-fields/percent) - Für Prozentwerte
- [Automatisierungs-API](/api/automations/1.index) - Erstellen von automatisierungen auf Basis von Zahlen