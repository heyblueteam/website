---
title: Nachschlagefeld
description: Erstellen Sie Nachschlagefelder, die automatisch Daten aus referenzierten Datensätzen abrufen
---

Nachschlagefelder ziehen automatisch Daten aus Datensätzen, die durch [Referenzfelder](/api/custom-fields/reference) referenziert werden, und zeigen Informationen aus verknüpften Datensätzen an, ohne dass eine manuelle Kopie erforderlich ist. Sie aktualisieren sich automatisch, wenn sich die referenzierten Daten ändern.

## Einfaches Beispiel

Erstellen Sie ein Nachschlagefeld, um Tags aus referenzierten Datensätzen anzuzeigen:

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

## Fortgeschrittenes Beispiel

Erstellen Sie ein Nachschlagefeld, um benutzerdefinierte Feldwerte aus referenzierten Datensätzen zu extrahieren:

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

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Nachschlagefelds |
| `type` | CustomFieldType! | ✅ Ja | Muss `LOOKUP` sein |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Ja | Nachschlagekonfiguration |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

## Nachschlagekonfiguration

### CustomFieldLookupOptionInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `referenceId` | String! | ✅ Ja | ID des Referenzfelds, aus dem Daten abgerufen werden |
| `lookupId` | String | Nein | ID des spezifischen benutzerdefinierten Felds, das nachgeschlagen werden soll (erforderlich für den Typ TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Ja | Art der Daten, die aus referenzierten Datensätzen extrahiert werden sollen |

## Nachschlagetypen

### CustomFieldLookupType Werte

| Typ | Beschreibung | Gibt zurück |
|------|-------------|---------|
| `TODO_DUE_DATE` | Fälligkeitsdaten aus referenzierten Todos | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Erstellungsdaten aus referenzierten Todos | Array of creation timestamps |
| `TODO_UPDATED_AT` | Letzte Aktualisierungsdaten aus referenzierten Todos | Array of update timestamps |
| `TODO_TAG` | Tags aus referenzierten Todos | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Zuweisungen aus referenzierten Todos | Array of user objects |
| `TODO_DESCRIPTION` | Beschreibungen aus referenzierten Todos | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Todo-Listen-Namen aus referenzierten Todos | Array of list titles |
| `TODO_CUSTOM_FIELD` | Werte aus benutzerdefinierten Feldern aus referenzierten Todos | Array of values based on the field type |

## Antwortfelder

### CustomField Antwort (für Nachschlagefelder)

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für das Feld |
| `name` | String! | Anzeigename des Nachschlagefelds |
| `type` | CustomFieldType! | Wird `LOOKUP` sein |
| `customFieldLookupOption` | CustomFieldLookupOption | Nachschlagekonfiguration und Ergebnisse |
| `createdAt` | DateTime! | Wann das Feld erstellt wurde |
| `updatedAt` | DateTime! | Wann das Feld zuletzt aktualisiert wurde |

### CustomFieldLookupOption Struktur

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | Art des durchgeführten Nachschlags |
| `lookupResult` | JSON | Die extrahierten Daten aus referenzierten Datensätzen |
| `reference` | CustomField | Das Referenzfeld, das als Quelle verwendet wird |
| `lookup` | CustomField | Das spezifische Feld, das nachgeschlagen wird (für TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Das übergeordnete Nachschlagefeld |
| `parentLookup` | CustomField | Übergeordnetes Nachschlagefeld in der Kette (für geschachtelte Nachschläge) |

## Wie Nachschläge funktionieren

1. **Datenextraktion**: Nachschläge extrahieren spezifische Daten aus allen Datensätzen, die über ein Referenzfeld verknüpft sind
2. **Automatische Updates**: Wenn sich referenzierte Datensätze ändern, aktualisieren sich die Nachschlagewerte automatisch
3. **Schreibgeschützt**: Nachschlagefelder können nicht direkt bearbeitet werden - sie spiegeln immer die aktuellen referenzierten Daten wider
4. **Keine Berechnungen**: Nachschläge extrahieren und zeigen Daten unverändert an, ohne Aggregationen oder Berechnungen

## TODO_CUSTOM_FIELD Nachschläge

Bei Verwendung des Typs `TODO_CUSTOM_FIELD` müssen Sie angeben, welches benutzerdefinierte Feld mit dem Parameter `lookupId` extrahiert werden soll:

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

Dies extrahiert die Werte des angegebenen benutzerdefinierten Felds aus allen referenzierten Datensätzen.

## Abfragen von Nachschlagdaten

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

## Beispiel für Nachschlagergebnisse

### Tag-Nachschlagergebnis
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

### Zuweisungs-Nachschlagergebnis
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

### Nachschlagergebnis für benutzerdefiniertes Feld
Die Ergebnisse variieren je nach Typ des nachgeschlagenen benutzerdefinierten Felds. Ein Währungsfeld-Nachschlag könnte beispielsweise zurückgeben:
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

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Wichtig**: Benutzer müssen über Anzeige-Berechtigungen sowohl für das aktuelle Projekt als auch für das referenzierte Projekt verfügen, um Nachschlageergebnisse zu sehen.

## Fehlerantworten

### Ungültiges Referenzfeld
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

### Zirkulärer Nachschlag erkannt
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

### Fehlende Nachschlag-ID für TODO_CUSTOM_FIELD
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

## Best Practices

1. **Klare Benennung**: Verwenden Sie beschreibende Namen, die anzeigen, welche Daten nachgeschlagen werden
2. **Geeignete Typen**: Wählen Sie den Nachschlagetyp, der Ihren Datenanforderungen entspricht
3. **Leistung**: Nachschläge verarbeiten alle referenzierten Datensätze, daher sollten Sie auf Referenzfelder mit vielen Links achten
4. **Berechtigungen**: Stellen Sie sicher, dass Benutzer Zugriff auf referenzierte Projekte haben, damit Nachschläge funktionieren

## Häufige Anwendungsfälle

### Projektübergreifende Sichtbarkeit
Zeigen Sie Tags, Zuweisungen oder Status aus verwandten Projekten ohne manuelle Synchronisierung an.

### Abhängigkeitsverfolgung
Zeigen Sie Fälligkeitsdaten oder den Abschlussstatus von Aufgaben an, von denen die aktuelle Arbeit abhängt.

### Ressourcenübersicht
Zeigen Sie alle Teammitglieder an, die den referenzierten Aufgaben zugewiesen sind, für die Ressourcenplanung.

### Statusaggregation
Sammeln Sie alle eindeutigen Status von verwandten Aufgaben, um die Projektgesundheit auf einen Blick zu sehen.

## Einschränkungen

- Nachschlagefelder sind schreibgeschützt und können nicht direkt bearbeitet werden
- Keine Aggregationsfunktionen (SUMME, ANZAHL, DURCHSCHNITT) - Nachschläge extrahieren nur Daten
- Keine Filteroptionen - alle referenzierten Datensätze sind enthalten
- Zirkuläre Nachschlageketten werden verhindert, um unendliche Schleifen zu vermeiden
- Ergebnisse spiegeln aktuelle Daten wider und aktualisieren sich automatisch

## Verwandte Ressourcen

- [Referenzfelder](/api/custom-fields/reference) - Erstellen Sie Links zu Datensätzen für Nachschlagequellen
- [Werte für benutzerdefinierte Felder](/api/custom-fields/custom-field-values) - Werte für bearbeitbare benutzerdefinierte Felder festlegen
- [Benutzerdefinierte Felder auflisten](/api/custom-fields/list-custom-fields) - Abfragen aller benutzerdefinierten Felder in einem Projekt