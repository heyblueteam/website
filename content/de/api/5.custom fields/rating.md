---
title: Bewertungsbenutzerdefiniertes Feld
description: Erstellen Sie Bewertungsfelder, um numerische Bewertungen mit konfigurierbaren Skalen und Validierungen zu speichern
---

Bewertungsbenutzerdefinierte Felder ermöglichen es Ihnen, numerische Bewertungen in Datensätzen mit konfigurierbaren Mindest- und Höchstwerten zu speichern. Sie sind ideal für Leistungsbewertungen, Zufriedenheitswerte, Prioritätsstufen oder beliebige datengestützte Werte in Ihren Projekten.

## Einfaches Beispiel

Erstellen Sie ein einfaches Bewertungsfeld mit der Standard-Skala 0-5:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Bewertungsfeld mit benutzerdefinierter Skala und Beschreibung:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Bewertungsfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `RATING` sein |
| `projectId` | String! | ✅ Ja | Die Projekt-ID, in der dieses Feld erstellt wird |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |
| `min` | Float | Nein | Mindestbewertungswert (kein Standard) |
| `max` | Float | Nein | Höchstbewertungswert |

## Bewertungswerte festlegen

Um einen Bewertungswert in einem Datensatz festzulegen oder zu aktualisieren:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Bewertungsbenutzerdefinierten Feldes |
| `value` | String! | ✅ Ja | Bewertungswert als Zeichenfolge (innerhalb des konfigurierten Bereichs) |

## Erstellen von Datensätzen mit Bewertungswerten

Beim Erstellen eines neuen Datensatzes mit Bewertungswerten:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
    }
  }
}
```

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `value` | Float | Der gespeicherte Bewertungswert (über customField.value zugänglich) |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

**Hinweis**: Der Bewertungswert wird tatsächlich über `customField.value.number` in Abfragen abgerufen.

### CustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für das Feld |
| `name` | String! | Anzeigename des Bewertungsfeldes |
| `type` | CustomFieldType! | Immer `RATING` |
| `min` | Float | Mindestbewertungswert |
| `max` | Float | Höchstbewertungswert |
| `description` | String | Hilfetext für das Feld |

## Bewertungsvalidierung

### Wertbeschränkungen
- Bewertungswerte müssen numerisch (Float-Typ) sein
- Werte müssen innerhalb des konfigurierten Min/Max-Bereichs liegen
- Wenn kein Minimum angegeben ist, gibt es keinen Standardwert
- Höchstwert ist optional, aber empfohlen

### Validierungsregeln
**Wichtig**: Die Validierung erfolgt nur beim Einreichen von Formularen, nicht bei der direkten Verwendung von `setTodoCustomField`.

- Eingaben werden als Fließkommazahl interpretiert (bei Verwendung von Formularen)
- Muss größer oder gleich dem Mindestwert sein (bei Verwendung von Formularen)
- Muss kleiner oder gleich dem Höchstwert sein (bei Verwendung von Formularen)
- `setTodoCustomField` akzeptiert jeden Zeichenwert ohne Validierung

### Gültige Bewertungsbeispiele
Für ein Feld mit min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Ungültige Bewertungsbeispiele
Für ein Feld mit min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Konfigurationsoptionen

### Bewertungsskalen-Setup
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Häufige Bewertungsskalen
- **1-5 Sterne**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 Leistung**: `min: 1, max: 10`
- **0-100 Prozent**: `min: 0, max: 100`
- **Benutzerdefinierte Skala**: Jeder numerische Bereich

## Erforderliche Berechtigungen

Benutzerdefinierte Feldoperationen folgen den standardmäßigen rollenbasierten Berechtigungen:

| Aktion | Erforderliche Rolle |
|--------|---------------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Hinweis**: Die spezifischen erforderlichen Rollen hängen von der benutzerdefinierten Rollenkonfiguration und den bereichsbezogenen Berechtigungen Ihres Projekts ab.

## Fehlermeldungen

### Validierungsfehler (nur Formulare)
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

**Wichtig**: Die Validierung des Bewertungswerts (Min/Max-Beschränkungen) erfolgt nur beim Einreichen von Formularen, nicht bei der direkten Verwendung von `setTodoCustomField`.

### Benutzerdefiniertes Feld nicht gefunden
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

## Best Practices

### Skalen-Design
- Verwenden Sie konsistente Bewertungsskalen über ähnliche Felder hinweg
- Berücksichtigen Sie die Vertrautheit der Benutzer (1-5 Sterne, 0-10 NPS)
- Setzen Sie angemessene Mindestwerte (0 vs 1)
- Definieren Sie eine klare Bedeutung für jede Bewertungsstufe

### Datenqualität
- Validieren Sie Bewertungswerte vor der Speicherung
- Verwenden Sie Dezimalgenauigkeit angemessen
- Berücksichtigen Sie Rundungen für Anzeigezwecke
- Geben Sie klare Hinweise zu den Bedeutungen der Bewertungen

### Benutzererfahrung
- Stellen Sie Bewertungsskalen visuell dar (Sterne, Fortschrittsbalken)
- Zeigen Sie den aktuellen Wert und die Skalenlimits an
- Geben Sie Kontext für die Bedeutungen der Bewertungen
- Berücksichtigen Sie Standardwerte für neue Datensätze

## Häufige Anwendungsfälle

1. **Leistungsmanagement**
   - Mitarbeiterleistungsbewertungen
   - Projektqualitätswerte
   - Bewertungen zum Abschluss von Aufgaben
   - Bewertungen des Fähigkeitsniveaus

2. **Kundenfeedback**
   - Zufriedenheitsbewertungen
   - Produktqualitätswerte
   - Bewertungen der Serviceerfahrung
   - Net Promoter Score (NPS)

3. **Priorität und Wichtigkeit**
   - Prioritätsstufen von Aufgaben
   - Dringlichkeitsbewertungen
   - Risikobewertungsergebnisse
   - Auswirkungenbewertungen

4. **Qualitätssicherung**
   - Bewertungen von Code-Überprüfungen
   - Qualitätswerte von Tests
   - Dokumentationsqualität
   - Bewertungen der Prozesskonformität

## Integrationsmerkmale

### Mit Automatisierungen
- Aktionen basierend auf Bewertungsgrenzen auslösen
- Benachrichtigungen für niedrige Bewertungen senden
- Nachverfolgungsaufgaben für hohe Bewertungen erstellen
- Arbeiten basierend auf Bewertungswerten zuweisen

### Mit Nachschlägen
- Durchschnittswerte von Bewertungen über Datensätze berechnen
- Datensätze nach Bewertungsbereichen finden
- Bewertungsdaten aus anderen Datensätzen referenzieren
- Bewertungsstatistiken aggregieren

### Mit Blue-Frontend
- Automatische Bereichsvalidierung in Formular-Kontexten
- Visuelle Bewertungs-Eingabesteuerungen
- Echtzeit-Validierungsfeedback
- Sterne- oder Schieberegler-Eingabeoptionen

## Aktivitätsverfolgung

Änderungen an Bewertungsfeldern werden automatisch verfolgt:
- Alte und neue Bewertungswerte werden protokolliert
- Aktivitäten zeigen numerische Änderungen
- Zeitstempel für alle Bewertungsaktualisierungen
- Benutzerzuordnung für Änderungen

## Einschränkungen

- Nur numerische Werte werden unterstützt
- Keine integrierte visuelle Bewertungsanzeige (Sterne usw.)
- Dezimalgenauigkeit hängt von der Datenbankkonfiguration ab
- Keine Speicherung von Bewertungsmetadaten (Kommentare, Kontext)
- Keine automatische Aggregation oder Statistiken von Bewertungen
- Keine integrierte Bewertungsumwandlung zwischen Skalen
- **Kritisch**: Min/Max-Validierung funktioniert nur in Formularen, nicht über `setTodoCustomField`

## Verwandte Ressourcen

- [Zahlenfelder](/api/5.custom%20fields/number) - Für allgemeine numerische Daten
- [Prozentfelder](/api/5.custom%20fields/percent) - Für Prozentwerte
- [Auswahlfelder](/api/5.custom%20fields/select-single) - Für diskrete Bewertungsoptionen
- [Übersicht über benutzerdefinierte Felder](/api/5.custom%20fields/2.list-custom-fields) - Allgemeine Konzepte