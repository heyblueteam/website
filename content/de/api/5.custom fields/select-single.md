---
title: Einzelne Auswahl benutzerdefiniertes Feld
description: Erstellen Sie einzelne Auswahlfelder, um Benutzern die Auswahl einer Option aus einer vordefinierten Liste zu ermöglichen
---

Einzelne Auswahl benutzerdefinierte Felder ermöglichen es Benutzern, genau eine Option aus einer vordefinierten Liste auszuwählen. Sie sind ideal für Statusfelder, Kategorien, Prioritäten oder jedes Szenario, in dem nur eine Wahl aus einer kontrollierten Menge von Optionen getroffen werden sollte.

## Basisbeispiel

Erstellen Sie ein einfaches einzelnes Auswahlfeld:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein einzelnes Auswahlfeld mit vordefinierten Optionen:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des einzelnen Auswahlfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `SELECT_SINGLE` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | Nein | Anfangsoptionen für das Feld |

### CreateCustomFieldOptionInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `title` | String! | ✅ Ja | Anzeigetext für die Option |
| `color` | String | Nein | Hex-Farbcode für die Option |

## Optionen zu bestehenden Feldern hinzufügen

Fügen Sie neue Optionen zu einem bestehenden einzelnen Auswahlfeld hinzu:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## Einzelne Auswahlwerte festlegen

Um die ausgewählte Option in einem Datensatz festzulegen:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des einzelnen Auswahlbenutzerfeldes |
| `customFieldOptionId` | String | Nein | ID der auszuwählenden Option (bevorzugt für Einzelne Auswahl) |
| `customFieldOptionIds` | [String!] | Nein | Array von Options-IDs (verwendet das erste Element für Einzelne Auswahl) |

## Einzelne Auswahlwerte abfragen

Fragen Sie den Einzelne Auswahlwert eines Datensatzes ab:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

Das `value` Feld gibt ein JSON-Objekt mit den Details der ausgewählten Option zurück.

## Datensätze mit einzelnen Auswahlwerten erstellen

Beim Erstellen eines neuen Datensatzes mit einzelnen Auswahlwerten:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # Contains the selected option object
    }
  }
}
```

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für den Feldwert |
| `customField` | CustomField! | Die benutzerdefinierte Felddefinition |
| `value` | JSON | Enthält das ausgewählte Optionsobjekt mit id, titel, farbe, position |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

### CustomFieldOption Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für die Option |
| `title` | String! | Anzeigetext für die Option |
| `color` | String | Hex-Farbcode für die visuelle Darstellung |
| `position` | Float | Sortierreihenfolge für die Option |
| `customField` | CustomField! | Das benutzerdefinierte Feld, zu dem diese Option gehört |

### CustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für das Feld |
| `name` | String! | Anzeigename des einzelnen Auswahlfeldes |
| `type` | CustomFieldType! | Immer `SELECT_SINGLE` |
| `description` | String | Hilfetext für das Feld |
| `customFieldOptions` | [CustomFieldOption!] | Alle verfügbaren Optionen |

## Wertformat

### Eingabeformat
- **API-Parameter**: Verwenden Sie `customFieldOptionId` für die ID der einzelnen Option
- **Alternative**: Verwenden Sie `customFieldOptionIds` Array (nimmt das erste Element)
- **Auswahl löschen**: Beide Felder weglassen oder leere Werte übergeben

### Ausgabeformat
- **GraphQL-Antwort**: JSON-Objekt im `value` Feld, das {id, titel, farbe, position} enthält
- **Aktivitätsprotokoll**: Optionsbezeichnung als Zeichenfolge
- **Automatisierungsdaten**: Optionsbezeichnung als Zeichenfolge

## Auswahlverhalten

### Exklusive Auswahl
- Das Festlegen einer neuen Option entfernt automatisch die vorherige Auswahl
- Es kann immer nur eine Option gleichzeitig ausgewählt werden
- Das Festlegen von `null` oder einem leeren Wert löscht die Auswahl

### Fallback-Logik
- Wenn das `customFieldOptionIds` Array bereitgestellt wird, wird nur die erste Option verwendet
- Dies gewährleistet die Kompatibilität mit Mehrfachauswahl-Eingabeformaten
- Leere Arrays oder null-Werte löschen die Auswahl

## Optionen verwalten

### Optionseigenschaften aktualisieren
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### Option löschen
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**Hinweis**: Das Löschen einer Option entfernt sie aus allen Datensätzen, in denen sie ausgewählt wurde.

### Optionen neu anordnen
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## Validierungsregeln

### Optionenvalidierung
- Die bereitgestellte Options-ID muss existieren
- Die Option muss dem angegebenen benutzerdefinierten Feld angehören
- Es kann immer nur eine Option ausgewählt werden (automatisch durchgesetzt)
- Null/leere Werte sind gültig (keine Auswahl)

### Feldvalidierung
- Es muss mindestens eine Option definiert sein, um verwendbar zu sein
- Optionstitel müssen innerhalb des Feldes eindeutig sein
- Farbcode muss im gültigen Hex-Format vorliegen (sofern bereitgestellt)

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|---------------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Fehlerantworten

### Ungültige Options-ID
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### Option gehört nicht zum Feld
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

### Wert kann nicht geparst werden
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Best Practices

### Optionen-Design
- Verwenden Sie klare, beschreibende Optionstitel
- Wenden Sie sinnvolle Farbkennzeichnungen an
- Halten Sie die Optionslisten fokussiert und relevant
- Ordnen Sie Optionen logisch (nach Priorität, Häufigkeit usw.)

### Statusfeldmuster
- Verwenden Sie konsistente Status-Workflows über Projekte hinweg
- Berücksichtigen Sie den natürlichen Verlauf der Optionen
- Schließen Sie klare Endzustände ein (Fertig, Abgebrochen usw.)
- Verwenden Sie Farben, die die Bedeutung der Optionen widerspiegeln

### Datenmanagement
- Überprüfen und bereinigen Sie ungenutzte Optionen regelmäßig
- Verwenden Sie konsistente Namenskonventionen
- Berücksichtigen Sie die Auswirkungen des Löschens von Optionen auf bestehende Datensätze
- Planen Sie für Optionenaktualisierungen und -migrationen

## Häufige Anwendungsfälle

1. **Status und Workflow**
   - Aufgabenstatus (Zu Erledigen, In Bearbeitung, Fertig)
   - Genehmigungsstatus (Ausstehend, Genehmigt, Abgelehnt)
   - Projektphase (Planung, Entwicklung, Test, Freigegeben)
   - Status der Problemlösung

2. **Klassifizierung und Kategorisierung**
   - Prioritätsstufen (Niedrig, Mittel, Hoch, Kritisch)
   - Aufgabentypen (Fehler, Funktion, Verbesserung, Dokumentation)
   - Projektkategorien (Intern, Kunde, Forschung)
   - Abteilungszuweisungen

3. **Qualität und Bewertung**
   - Überprüfungsstatus (Nicht gestartet, In Überprüfung, Genehmigt)
   - Qualitätsbewertungen (Schlecht, Ausreichend, Gut, Ausgezeichnet)
   - Risikostufen (Niedrig, Mittel, Hoch)
   - Vertrauensniveaus

4. **Zuweisung und Eigentum**
   - Teamzuweisungen
   - Abteilungseigentum
   - Rollenbasierte Zuweisungen
   - Regionale Zuweisungen

## Integrationsfunktionen

### Mit Automatisierungen
- Aktionen auslösen, wenn bestimmte Optionen ausgewählt werden
- Arbeiten basierend auf ausgewählten Kategorien leiten
- Benachrichtigungen für Statusänderungen senden
- Bedingte Workflows basierend auf Auswahlen erstellen

### Mit Nachschlägen
- Datensätze nach ausgewählten Optionen filtern
- Optionsdaten aus anderen Datensätzen referenzieren
- Berichte basierend auf Optionsauswahlen erstellen
- Datensätze nach ausgewählten Werten gruppieren

### Mit Formularen
- Dropdown-Eingabesteuerungen
- Radio-Button-Schnittstellen
- Optionenvalidierung und -filterung
- Bedingte Feldanzeige basierend auf Auswahlen

## Aktivitätsverfolgung

Änderungen an einzelnen Auswahlfeldern werden automatisch verfolgt:
- Zeigt alte und neue Optionsauswahlen an
- Zeigt Optionsbezeichnungen im Aktivitätsprotokoll an
- Zeitstempel für alle Auswahländerungen
- Benutzerzuordnung für Änderungen

## Unterschiede zur Mehrfachauswahl

| Funktion | Einzelne Auswahl | Mehrfachauswahl |
|---------|-----------------|----------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Einschränkungen

- Es kann immer nur eine Option gleichzeitig ausgewählt werden
- Keine hierarchische oder verschachtelte Optionsstruktur
- Optionen werden über alle Datensätze hinweg, die das Feld verwenden, geteilt
- Keine integrierte Optionenanalytik oder Nutzungstracking
- Farbcode dient nur zur Anzeige, hat keine funktionalen Auswirkungen
- Es können keine unterschiedlichen Berechtigungen pro Option festgelegt werden

## Verwandte Ressourcen

- [Mehrfachauswahlfelder](/api/custom-fields/select-multi) - Für Mehrfachauswahlen
- [Checkboxfelder](/api/custom-fields/checkbox) - Für einfache boolesche Entscheidungen
- [Textfelder](/api/custom-fields/text-single) - Für Freitext-Eingaben
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/1.index) - Allgemeine Konzepte