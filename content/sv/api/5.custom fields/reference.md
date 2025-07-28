---
title: Referens anpassat fält
description: Skapa referensfält som länkar till poster i andra projekt för tvärprojektrelationer
---

Referens anpassade fält gör att du kan skapa länkar mellan poster i olika projekt, vilket möjliggör tvärprojektrelationer och datadelning. De ger ett kraftfullt sätt att koppla samman relaterat arbete över din organisations projektstruktur.

## Grundläggande exempel

Skapa ett enkelt referensfält:

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

## Avancerat exempel

Skapa ett referensfält med filtrering och flera val:

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

## Indata parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för referensfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `REFERENCE` |
| `referenceProjectId` | String | Nej | ID för projektet att referera till |
| `referenceMultiple` | Boolean | Nej | Tillåt flera postval (standard: falskt) |
| `referenceFilter` | TodoFilterInput | Nej | Filterkriterier för refererade poster |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: Anpassade fält är automatiskt kopplade till projektet baserat på användarens aktuella projektkontext.

## Referenskonfiguration

### Enkel vs flera referenser

**Enkel referens (standard):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Användare kan välja en post från det refererade projektet
- Returnerar ett enda Todo-objekt

**Flera referenser:**
```graphql
{
  referenceMultiple: true
}
```
- Användare kan välja flera poster från det refererade projektet
- Returnerar en array av Todo-objekt

### Referensfiltrering

Använd `referenceFilter` för att begränsa vilka poster som kan väljas:

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

## Ställa in referensvärden

### Enkel referens

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Flera referenser

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

### SetTodoCustomFieldInput parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det referens anpassade fältet |
| `customFieldReferenceTodoIds` | [String!] | ✅ Ja | Array av refererade post-ID:n |

## Skapa poster med referenser

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

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | ID! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Definitionen av referensfältet |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

**Notera**: Refererade todos nås via `customField.selectedTodos`, inte direkt på TodoCustomField.

### Refererade Todo-fält

Varje refererad Todo inkluderar:

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | ID! | Unik identifierare för den refererade posten |
| `title` | String! | Titel på den refererade posten |
| `status` | TodoStatus! | Aktuell status (AKTIV, AVSLUTAD, etc.) |
| `description` | String | Beskrivning av den refererade posten |
| `dueDate` | DateTime | Förfallodatum om det är inställt |
| `assignees` | [User!] | Tilldelade användare |
| `tags` | [Tag!] | Associerade taggar |
| `project` | Project! | Projekt som innehåller den refererade posten |

## Fråga referensdata

### Grundläggande fråga

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

### Avancerad fråga med nästlad data

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

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk behörighet |
|--------|------------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Viktigt**: Användare måste ha visningsbehörigheter på det refererade projektet för att se de länkade posterna.

## Tvärprojektåtkomst

### Projektets synlighet

- Användare kan endast referera till poster från projekt som de har tillgång till
- Refererade poster respekterar det ursprungliga projektets behörigheter
- Ändringar av refererade poster visas i realtid
- Att ta bort refererade poster tar bort dem från referensfälten

### Behörighetsarv

- Referensfält ärver behörigheter från båda projekten
- Användare behöver visningsåtkomst till det refererade projektet
- Redigeringsbehörigheter baseras på det aktuella projektets regler
- Refererad data är skrivskyddad i kontexten av referensfältet

## Felrespons

### Ogiltigt referensprojekt

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

### Refererad post hittades inte

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

### Behörighet nekad

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

## Bästa praxis

### Fältdesign

1. **Tydlig namngivning** - Använd beskrivande namn som indikerar relationen
2. **Lämplig filtrering** - Ställ in filter för att visa endast relevanta poster
3. **Överväg behörigheter** - Se till att användare har tillgång till refererade projekt
4. **Dokumentera relationer** - Ge tydliga beskrivningar av kopplingen

### Prestandaöverväganden

1. **Begränsa referensens omfattning** - Använd filter för att minska antalet valbara poster
2. **Undvik djup nästling** - Skapa inte komplexa kedjor av referenser
3. **Överväg caching** - Refererad data cachas för prestanda
4. **Övervaka användning** - Spåra hur referenser används över projekt

### Dataintegritet

1. **Hantera borttagningar** - Planera för när refererade poster tas bort
2. **Validera behörigheter** - Se till att användare kan få åtkomst till refererade projekt
3. **Uppdatera beroenden** - Överväg påverkan vid ändring av refererade poster
4. **Revisionsspår** - Spåra referensrelationer för efterlevnad

## Vanliga användningsfall

### Projektberoenden

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

### Klientkrav

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

### Resursallokering

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

### Kvalitetssäkring

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

## Integration med uppslag

Referensfält fungerar med [Uppslagsfält](/api/custom-fields/lookup) för att hämta data från refererade poster. Uppslagsfält kan extrahera värden från poster som valts i referensfält, men de är endast datainsamlare (inga aggregeringsfunktioner som SUM stöds).

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

## Begränsningar

- Refererade projekt måste vara tillgängliga för användaren
- Ändringar av behörigheter för refererade projekt påverkar åtkomsten till referensfält
- Djup nästling av referenser kan påverka prestandan
- Ingen inbyggd validering för cirkulära referenser
- Ingen automatisk begränsning som förhindrar referenser inom samma projekt
- Filtervalidering tillämpas inte vid inställning av referensvärden

## Relaterade resurser

- [Uppslagsfält](/api/custom-fields/lookup) - Extrahera data från refererade poster
- [Projekt API](/api/projects) - Hantera projekt som innehåller referenser
- [Post API](/api/records) - Arbeta med poster som har referenser
- [Översikt över anpassade fält](/api/custom-fields/list-custom-fields) - Allmänna begrepp