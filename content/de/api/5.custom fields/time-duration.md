---
title: Zeitdauer benutzerdefiniertes Feld
description: Erstellen Sie berechnete Zeitdauerfelder, die die Zeit zwischen Ereignissen in Ihrem Workflow verfolgen
---

Zeitdauer benutzerdefinierte Felder berechnen und zeigen automatisch die Dauer zwischen zwei Ereignissen in Ihrem Workflow an. Sie sind ideal, um Bearbeitungszeiten, Reaktionszeiten, Zykluszeiten oder andere zeitbasierte Kennzahlen in Ihren Projekten zu verfolgen.

## Einfaches Beispiel

Erstellen Sie ein einfaches Zeitdauerfeld, das verfolgt, wie lange Aufgaben zur Fertigstellung benötigen:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein komplexes Zeitdauerfeld, das die Zeit zwischen Änderungen des benutzerdefinierten Feldes mit einem SLA-Ziel verfolgt:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput (TIME_DURATION)

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Dauerfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `TIME_DURATION` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Ja | Wie die Dauer angezeigt werden soll |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Ja | Konfiguration des Startereignisses |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Ja | Konfiguration des Endereignisses |
| `timeDurationTargetTime` | Float | Nein | Zielzeit in Sekunden für die SLA-Überwachung |

### CustomFieldTimeDurationInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ Ja | Art des zu verfolgenden Ereignisses |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Ja | `FIRST` oder `LAST` Vorkommen |
| `customFieldId` | String | Conditional | Erforderlich für `TODO_CUSTOM_FIELD` Typ |
| `customFieldOptionIds` | [String!] | Conditional | Erforderlich für Änderungen an Auswahlfeldern |
| `todoListId` | String | Conditional | Erforderlich für `TODO_MOVED` Typ |
| `tagId` | String | Conditional | Erforderlich für `TODO_TAG_ADDED` Typ |
| `assigneeId` | String | Conditional | Erforderlich für `TODO_ASSIGNEE_ADDED` Typ |

### CustomFieldTimeDurationType Werte

| Wert | Beschreibung |
|-------|-------------|
| `TODO_CREATED_AT` | Wann der Datensatz erstellt wurde |
| `TODO_CUSTOM_FIELD` | Wann sich ein benutzerdefinierter Feldwert geändert hat |
| `TODO_DUE_DATE` | Wann das Fälligkeitsdatum festgelegt wurde |
| `TODO_MARKED_AS_COMPLETE` | Wann der Datensatz als abgeschlossen markiert wurde |
| `TODO_MOVED` | Wann der Datensatz in eine andere Liste verschoben wurde |
| `TODO_TAG_ADDED` | Wann ein Tag zum Datensatz hinzugefügt wurde |
| `TODO_ASSIGNEE_ADDED` | Wann ein Zuweisender zum Datensatz hinzugefügt wurde |

### CustomFieldTimeDurationCondition Werte

| Wert | Beschreibung |
|-------|-------------|
| `FIRST` | Verwenden Sie das erste Vorkommen des Ereignisses |
| `LAST` | Verwenden Sie das letzte Vorkommen des Ereignisses |

### CustomFieldTimeDurationDisplayType Werte

| Wert | Beschreibung | Beispiel |
|-------|-------------|---------|
| `FULL_DATE` | Tage:Stunden:Minuten:Sekunden Format | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Vollständig ausgeschrieben | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numerisch mit Einheiten | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Dauer nur in Tagen | `"2.5"` (2.5 days) |
| `HOURS` | Dauer nur in Stunden | `"60"` (60 hours) |
| `MINUTES` | Dauer nur in Minuten | `"3600"` (3600 minutes) |
| `SECONDS` | Dauer nur in Sekunden | `"216000"` (216000 seconds) |

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `number` | Float | Die Dauer in Sekunden |
| `value` | Float | Alias für Zahl (Dauer in Sekunden) |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt aktualisiert wurde |

### CustomField Antwort (TIME_DURATION)

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Anzeigeformat für die Dauer |
| `timeDurationStart` | CustomFieldTimeDuration | Konfiguration des Startereignisses |
| `timeDurationEnd` | CustomFieldTimeDuration | Konfiguration des Endereignisses |
| `timeDurationTargetTime` | Float | Zielzeit in Sekunden (für SLA-Überwachung) |

## Dauerberechnung

### So funktioniert es
1. **Startereignis**: Das System überwacht das angegebene Startereignis
2. **Endereignis**: Das System überwacht das angegebene Endereignis
3. **Berechnung**: Dauer = Endzeit - Startzeit
4. **Speicherung**: Dauer wird in Sekunden als Zahl gespeichert
5. **Anzeige**: Formatiert gemäß der Einstellung `timeDurationDisplay`

### Aktualisierungsereignisse
Dauerwerte werden automatisch neu berechnet, wenn:
- Datensätze erstellt oder aktualisiert werden
- Werte benutzerdefinierter Felder sich ändern
- Tags hinzugefügt oder entfernt werden
- Zuweisende hinzugefügt oder entfernt werden
- Datensätze zwischen Listen verschoben werden
- Datensätze als abgeschlossen/nicht abgeschlossen markiert werden

## Dauerwerte lesen

### Abfrage von Dauerfeldern
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Formatierte Anzeigewerte
Dauerwerte werden automatisch basierend auf der Einstellung `timeDurationDisplay` formatiert:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Häufige Konfigurationsbeispiele

### Abschlusszeit von Aufgaben
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Dauer von Statusänderungen
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Zeit in spezifischer Liste
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Reaktionszeit bei Zuweisungen
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Fehlermeldungen

### Ungültige Konfiguration
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Referenziertes Feld nicht gefunden
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

### Fehlende erforderliche Optionen
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Wichtige Hinweise

### Automatische Berechnung
- Dauerfelder sind **schreibgeschützt** - Werte werden automatisch berechnet
- Sie können Dauerwerte nicht manuell über die API festlegen
- Berechnungen erfolgen asynchron über Hintergrundjobs
- Werte werden automatisch aktualisiert, wenn Triggerereignisse auftreten

### Leistungsüberlegungen
- Dauerberechnungen werden in Warteschlangen gestellt und asynchron verarbeitet
- Eine große Anzahl von Dauerfeldern kann die Leistung beeinträchtigen
- Berücksichtigen Sie die Häufigkeit von Triggerereignissen bei der Gestaltung von Dauerfeldern
- Verwenden Sie spezifische Bedingungen, um unnötige Neuberechnungen zu vermeiden

### Nullwerte
Dauerfelder zeigen `null` an, wenn:
- Das Startereignis noch nicht aufgetreten ist
- Das Endereignis noch nicht aufgetreten ist
- Die Konfiguration auf nicht vorhandene Entitäten verweist
- Bei der Berechnung ein Fehler auftritt

## Best Practices

### Konfigurationsdesign
- Verwenden Sie spezifische Ereignistypen anstelle von generischen, wenn möglich
- Wählen Sie geeignete `FIRST` vs `LAST` Bedingungen basierend auf Ihrem Workflow
- Testen Sie Dauerberechnungen mit Beispieldaten vor der Bereitstellung
- Dokumentieren Sie Ihre Logik für Dauerfelder für Teammitglieder

### Anzeigeformatierung
- Verwenden Sie `FULL_DATE_SUBSTRING` für das lesbarste Format
- Verwenden Sie `FULL_DATE` für eine kompakte, konsistente Breite
- Verwenden Sie `FULL_DATE_STRING` für formale Berichte und Dokumente
- Verwenden Sie `DAYS`, `HOURS`, `MINUTES` oder `SECONDS` für einfache numerische Anzeigen
- Berücksichtigen Sie Ihre UI-Platzbeschränkungen bei der Auswahl des Formats

### SLA-Überwachung mit Zielzeit
Bei der Verwendung von `timeDurationTargetTime`:
- Stellen Sie die Zielzeit in Sekunden ein
- Vergleichen Sie die tatsächliche Dauer mit dem Ziel für die SLA-Konformität
- Verwenden Sie in Dashboards, um überfällige Elemente hervorzuheben
- Beispiel: 24-Stunden-Reaktions-SLA = 86400 Sekunden

### Workflow-Integration
- Gestalten Sie Dauerfelder so, dass sie mit Ihren tatsächlichen Geschäftsprozessen übereinstimmen
- Verwenden Sie Dauerinformationen zur Prozessverbesserung und -optimierung
- Überwachen Sie Dauertrends, um Engpässe im Workflow zu identifizieren
- Richten Sie bei Bedarf Warnungen für Dauergrenzen ein

## Häufige Anwendungsfälle

1. **Prozessleistung**
   - Abschlusszeiten von Aufgaben
   - Überprüfungszykluszeiten
   - Genehmigungsbearbeitungszeiten
   - Reaktionszeiten

2. **SLA-Überwachung**
   - Zeit bis zur ersten Antwort
   - Lösungszeiten
   - Eskalationszeiträume
   - Einhaltung des Servicelevels

3. **Workflow-Analytik**
   - Engpassidentifikation
   - Prozessoptimierung
   - Teamleistungskennzahlen
   - Qualitätssicherungszeit

4. **Projektmanagement**
   - Phasendauern
   - Meilensteinzeit
   - Zeit für Ressourcenallokation
   - Lieferzeiträume

## Einschränkungen

- Dauerfelder sind **schreibgeschützt** und können nicht manuell festgelegt werden
- Werte werden asynchron berechnet und sind möglicherweise nicht sofort verfügbar
- Erfordert, dass die richtigen Ereignistrigger in Ihrem Workflow eingerichtet sind
- Kann keine Dauer für Ereignisse berechnen, die noch nicht aufgetreten sind
- Beschränkt auf die Verfolgung der Zeit zwischen diskreten Ereignissen (nicht kontinuierliche Zeitverfolgung)
- Keine integrierten SLA-Warnungen oder -Benachrichtigungen
- Kann mehrere Dauerberechnungen nicht in einem einzigen Feld aggregieren

## Verwandte Ressourcen

- [Zahlenfelder](/api/custom-fields/number) - Für manuelle numerische Werte
- [Datumsfelder](/api/custom-fields/date) - Für die Verfolgung spezifischer Daten
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte
- [Automatisierungen](/api/automations) - Für das Auslösen von Aktionen basierend auf Dauergrenzen