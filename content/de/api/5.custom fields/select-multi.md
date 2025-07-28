---
title: Mehrfachauswahl benutzerdefiniertes Feld
description: Erstellen Sie Mehrfachauswahlfelder, um Benutzern die Auswahl mehrerer Optionen aus einer vordefinierten Liste zu ermöglichen
---

Mehrfachauswahl benutzerdefinierte Felder ermöglichen es Benutzern, mehrere Optionen aus einer vordefinierten Liste auszuwählen. Sie sind ideal für Kategorien, Tags, Fähigkeiten, Funktionen oder jedes Szenario, in dem mehrere Auswahlen aus einer kontrollierten Menge von Optionen erforderlich sind.

## Einfaches Beispiel

Erstellen Sie ein einfaches Mehrfachauswahlfeld:

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Mehrfachauswahlfeld und fügen Sie dann Optionen separat hinzu:

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Mehrfachauswahlfelds |
| `type` | CustomFieldType! | ✅ Ja | Muss sein `SELECT_MULTI` |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |
| `projectId` | String! | ✅ Ja | ID des Projekts für dieses Feld |

### CreateCustomFieldOptionInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `customFieldId` | String! | ✅ Ja | ID des benutzerdefinierten Felds |
| `title` | String! | ✅ Ja | Anzeigetext für die Option |
| `color` | String | Nein | Farbe für die Option (beliebiger String) |
| `position` | Float | Nein | Sortierreihenfolge für die Option |

## Hinzufügen von Optionen zu bestehenden Feldern

Fügen Sie neue Optionen zu einem bestehenden Mehrfachauswahlfeld hinzu:

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## Festlegen von Mehrfachauswahlwerten

Um mehrere ausgewählte Optionen in einem Datensatz festzulegen:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des benutzerdefinierten Mehrfachauswahlfelds |
| `customFieldOptionIds` | [String!] | ✅ Ja | Array von Option IDs, die ausgewählt werden sollen |

## Erstellen von Datensätzen mit Mehrfachauswahlwerten

Beim Erstellen eines neuen Datensatzes mit Mehrfachauswahlwerten:

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
    }
  }
}
```

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Felds |
| `selectedOptions` | [CustomFieldOption!] | Array von ausgewählten Optionen |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

### CustomFieldOption Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für die Option |
| `title` | String! | Anzeigetext für die Option |
| `color` | String | Hex-Farbcode zur visuellen Darstellung |
| `position` | Float | Sortierreihenfolge für die Option |
| `customField` | CustomField! | Das benutzerdefinierte Feld, zu dem diese Option gehört |

### CustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für das Feld |
| `name` | String! | Anzeigename des Mehrfachauswahlfelds |
| `type` | CustomFieldType! | Immer `SELECT_MULTI` |
| `description` | String | Hilfetext für das Feld |
| `customFieldOptions` | [CustomFieldOption!] | Alle verfügbaren Optionen |

## Wertformat

### Eingabeformat
- **API-Parameter**: Array von Option IDs (`["option1", "option2", "option3"]`)
- **String-Format**: Komma-getrennte Option IDs (`"option1,option2,option3"`)

### Ausgabeformat
- **GraphQL-Antwort**: Array von CustomFieldOption-Objekten
- **Aktivitätsprotokoll**: Komma-getrennte Optionstitel
- **Automatisierungsdaten**: Array von Optionstiteln

## Verwaltung von Optionen

### Aktualisieren von Optionen
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
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

### Optionen neu anordnen
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## Validierungsregeln

### Optionen Validierung
- Alle angegebenen Option IDs müssen existieren
- Optionen müssen dem angegebenen benutzerdefinierten Feld zugeordnet sein
- Nur SELECT_MULTI-Felder können mehrere Optionen ausgewählt haben
- Leeres Array ist gültig (keine Auswahlen)

### Feldvalidierung
- Es muss mindestens eine Option definiert sein, um verwendbar zu sein
- Optionstitel müssen innerhalb des Feldes eindeutig sein
- Das Farbfeld akzeptiert jeden Stringwert (keine Hex-Validierung)

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|---------------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Fehlermeldungen

### Ungültige Option ID
```json
{
  "errors": [{
    "message": "Custom field option not found",
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Mehrere Optionen in einem Nicht-Mehrfachfeld
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Beste Praktiken

### Optionen Design
- Verwenden Sie beschreibende, prägnante Optionstitel
- Wenden Sie konsistente Farbcodierungsschemata an
- Halten Sie die Optionslisten überschaubar (typischerweise 3-20 Optionen)
- Ordnen Sie Optionen logisch (alphabetisch, nach Häufigkeit usw.)

### Datenmanagement
- Überprüfen und bereinigen Sie ungenutzte Optionen regelmäßig
- Verwenden Sie konsistente Namenskonventionen über Projekte hinweg
- Berücksichtigen Sie die Wiederverwendbarkeit von Optionen bei der Erstellung von Feldern
- Planen Sie für Optionenaktualisierungen und Migrationen

### Benutzererfahrung
- Geben Sie klare Feldbeschreibungen an
- Verwenden Sie Farben, um visuelle Unterscheidungen zu verbessern
- Gruppieren Sie verwandte Optionen zusammen
- Berücksichtigen Sie Standardauswahlen für häufige Fälle

## Häufige Anwendungsfälle

1. **Projektmanagement**
   - Aufgaben Kategorien und Tags
   - Prioritätsstufen und -typen
   - Zuweisungen von Teammitgliedern
   - Statusindikatoren

2. **Inhaltsmanagement**
   - Artikelkategorien und -themen
   - Inhaltstypen und -formate
   - Veröffentlichungswege
   - Genehmigungsworkflows

3. **Kundensupport**
   - Problemlösungs-Kategorien und -typen
   - Betroffene Produkte oder Dienstleistungen
   - Lösungsansätze
   - Kundensegmente

4. **Produktentwicklung**
   - Funktionskategorien
   - Technische Anforderungen
   - Testumgebungen
   - Veröffentlichungswege

## Integrationsfunktionen

### Mit Automatisierungen
- Aktionen auslösen, wenn bestimmte Optionen ausgewählt werden
- Arbeiten basierend auf ausgewählten Kategorien leiten
- Benachrichtigungen für hochpriorisierte Auswahlen senden
- Nachverfolgungsaufgaben basierend auf Optionskombinationen erstellen

### Mit Nachschlagewerten
- Datensätze nach ausgewählten Optionen filtern
- Daten über Optionsauswahlen aggregieren
- Optionsdaten aus anderen Datensätzen referenzieren
- Berichte basierend auf Optionskombinationen erstellen

### Mit Formularen
- Mehrfachauswahl-Eingabesteuerungen
- Optionenvalidierung und -filterung
- Dynamisches Laden von Optionen
- Bedingte Feldanzeige

## Aktivitätsverfolgung

Änderungen an Mehrfachauswahlfeldern werden automatisch verfolgt:
- Zeigt hinzugefügte und entfernte Optionen an
- Zeigt Optionstitel im Aktivitätsprotokoll an
- Zeitstempel für alle Auswahländerungen
- Benutzerzuordnung für Änderungen

## Einschränkungen

- Maximale praktische Anzahl von Optionen hängt von der UI-Leistung ab
- Keine hierarchische oder geschachtelte Optionsstruktur
- Optionen werden über alle Datensätze, die das Feld verwenden, geteilt
- Keine integrierte Optionenanalyse oder Nutzungsverfolgung
- Farbfeld akzeptiert jeden String (keine Hex-Validierung)
- Es können keine unterschiedlichen Berechtigungen pro Option festgelegt werden
- Optionen müssen separat erstellt werden, nicht inline mit der Feld Erstellung
- Keine dedizierte Neuanordnungsmutation (verwenden Sie editCustomFieldOption mit Position)

## Verwandte Ressourcen

- [Einzelauswahlfelder](/api/custom-fields/select-single) - Für Einzelwahl-Auswahlen
- [Checkbox-Felder](/api/custom-fields/checkbox) - Für einfache boolesche Entscheidungen
- [Textfelder](/api/custom-fields/text-single) - Für Freitext-Eingaben
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/2.list-custom-fields) - Allgemeine Konzepte