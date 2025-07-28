---
title: Referenzbenutzerdefiniertes Feld
description: Erstellen Sie Referenzfelder, die auf Datensätze in anderen Projekten verlinken, um projektübergreifende Beziehungen herzustellen
---

Referenzbenutzerdefinierte Felder ermöglichen es Ihnen, Links zwischen Datensätzen in verschiedenen Projekten zu erstellen, wodurch projektübergreifende Beziehungen und Datenaustausch ermöglicht werden. Sie bieten eine leistungsstarke Möglichkeit, verwandte Arbeiten innerhalb der Projektstruktur Ihrer Organisation zu verbinden.

## Einfaches Beispiel

Erstellen Sie ein einfaches Referenzfeld:

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Referenzfeld mit Filterung und Mehrfachauswahl:

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Referenzfelds |
| `type` | CustomFieldType! | ✅ Ja | Muss `REFERENCE` sein |
| `referenceProjectId` | String | Nein | ID des Projekts, auf das verwiesen wird |
| `referenceMultiple` | Boolean | Nein | Mehrfache Datensatzauswahl zulassen (Standard: false) |
| `referenceFilter` | TodoFilterInput | Nein | Filterkriterien für referenzierte Datensätze |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Benutzerdefinierte Felder sind automatisch mit dem Projekt verknüpft, basierend auf dem aktuellen Projektkontext des Benutzers.

## Referenzkonfiguration

### Einzel- vs. Mehrfachreferenzen

**Einzelreferenz (Standard):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Benutzer können einen Datensatz aus dem referenzierten Projekt auswählen
- Gibt ein einzelnes Todo-Objekt zurück

**Mehrfachreferenzen:**
```graphql
{
  referenceMultiple: true
}
```
- Benutzer können mehrere Datensätze aus dem referenzierten Projekt auswählen
- Gibt ein Array von Todo-Objekten zurück

### Referenzfilterung

Verwenden Sie `referenceFilter`, um einzuschränken, welche Datensätze ausgewählt werden können:

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## Festlegen von Referenzwerten

### Einzelreferenz

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Mehrfachreferenzen

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des referenzierten benutzerdefinierten Feldes |
| `customFieldReferenceTodoIds` | [String!] | ✅ Ja | Array von referenzierten Datensatz-IDs |

## Erstellen von Datensätzen mit Referenzen

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | ID! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des Referenzfeldes |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

**Hinweis**: Referenzierte Todos werden über `customField.selectedTodos` abgerufen, nicht direkt auf TodoCustomField.

### Referenzierte Todo-Felder

Jedes referenzierte Todo enthält:

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | ID! | Eindeutige Kennung des referenzierten Datensatzes |
| `title` | String! | Titel des referenzierten Datensatzes |
| `status` | TodoStatus! | Aktueller Status (AKTIV, ABGESCHLOSSEN usw.) |
| `description` | String | Beschreibung des referenzierten Datensatzes |
| `dueDate` | DateTime | Fälligkeitsdatum, falls festgelegt |
| `assignees` | [User!] | Zugewiesene Benutzer |
| `tags` | [Tag!] | Zugeordnete Tags |
| `project` | Project! | Projekt, das den referenzierten Datensatz enthält |

## Abfragen von Referenzdaten

### Grundlegende Abfrage

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### Erweiterte Abfrage mit geschachtelten Daten

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Wichtig**: Benutzer müssen über Anzeige-Berechtigungen für das referenzierte Projekt verfügen, um die verlinkten Datensätze zu sehen.

## Projektübergreifender Zugriff

### Projekt Sichtbarkeit

- Benutzer können nur Datensätze aus Projekten referenzieren, auf die sie Zugriff haben
- Referenzierte Datensätze respektieren die Berechtigungen des ursprünglichen Projekts
- Änderungen an referenzierten Datensätzen erscheinen in Echtzeit
- Das Löschen referenzierter Datensätze entfernt sie aus den Referenzfeldern

### Berechtigungsübertragung

- Referenzfelder erben Berechtigungen von beiden Projekten
- Benutzer benötigen Anzeigezugriff auf das referenzierte Projekt
- Bearbeitungsberechtigungen basieren auf den Regeln des aktuellen Projekts
- Referenzierte Daten sind im Kontext des Referenzfelds schreibgeschützt

## Fehlermeldungen

### Ungültiges Referenzprojekt

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### Referenzierter Datensatz nicht gefunden

```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Berechtigung verweigert

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## Best Practices

### Feldgestaltung

1. **Klare Benennung** - Verwenden Sie beschreibende Namen, die die Beziehung anzeigen
2. **Angemessene Filterung** - Setzen Sie Filter, um nur relevante Datensätze anzuzeigen
3. **Berechtigungen berücksichtigen** - Stellen Sie sicher, dass Benutzer Zugriff auf referenzierte Projekte haben
4. **Dokumentieren Sie Beziehungen** - Geben Sie klare Beschreibungen der Verbindung an

### Leistungsüberlegungen

1. **Begrenzen Sie den Referenzbereich** - Verwenden Sie Filter, um die Anzahl der auswählbaren Datensätze zu reduzieren
2. **Vermeiden Sie tiefe Verschachtelungen** - Erstellen Sie keine komplexen Ketten von Referenzen
3. **Caching in Betracht ziehen** - Referenzierte Daten werden zur Verbesserung der Leistung zwischengespeichert
4. **Nutzung überwachen** - Verfolgen Sie, wie Referenzen in Projekten verwendet werden

### Datenintegrität

1. **Löschen behandeln** - Planen Sie, was passiert, wenn referenzierte Datensätze gelöscht werden
2. **Berechtigungen validieren** - Stellen Sie sicher, dass Benutzer auf referenzierte Projekte zugreifen können
3. **Abhängigkeiten aktualisieren** - Berücksichtigen Sie die Auswirkungen beim Ändern referenzierter Datensätze
4. **Audit-Trails** - Verfolgen Sie Referenzbeziehungen zur Einhaltung von Vorschriften

## Häufige Anwendungsfälle

### Projektabhängigkeiten

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### Kundenanforderungen

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### Ressourcenallokation

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### Qualitätssicherung

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## Integration mit Nachschlagefeldern

Referenzfelder arbeiten mit [Nachschlagefeldern](/api/custom-fields/lookup) zusammen, um Daten aus referenzierten Datensätzen abzurufen. Nachschlagefelder können Werte aus in Referenzfeldern ausgewählten Datensätzen extrahieren, sind jedoch nur Datenextraktoren (keine Aggregatfunktionen wie SUM werden unterstützt).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## Einschränkungen

- Referenzierte Projekte müssen für den Benutzer zugänglich sein
- Änderungen an den Berechtigungen des referenzierten Projekts wirken sich auf den Zugriff auf das Referenzfeld aus
- Tiefe Verschachtelungen von Referenzen können die Leistung beeinträchtigen
- Keine integrierte Validierung für zirkuläre Referenzen
- Keine automatische Einschränkung, die Referenzen im selben Projekt verhindert
- Filtervalidierung wird nicht durchgesetzt, wenn Referenzwerte festgelegt werden

## Verwandte Ressourcen

- [Nachschlagefelder](/api/custom-fields/lookup) - Daten aus referenzierten Datensätzen extrahieren
- [Projekte API](/api/projects) - Verwaltung von Projekten, die Referenzen enthalten
- [Datensätze API](/api/records) - Arbeiten mit Datensätzen, die Referenzen haben
- [Überblick über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte