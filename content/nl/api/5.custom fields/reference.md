---
title: Referentie Aangepast Veld
description: Maak referentievelden die linken naar records in andere projecten voor cross-projectrelaties
---

Referentie aangepaste velden stellen je in staat om links te maken tussen records in verschillende projecten, waardoor cross-projectrelaties en gegevensdeling mogelijk worden. Ze bieden een krachtige manier om gerelateerde werkzaamheden binnen de projectstructuur van je organisatie te verbinden.

## Basisvoorbeeld

Maak een eenvoudig referentieveld:

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

## Geavanceerd Voorbeeld

Maak een referentieveld met filtering en meervoudige selectie:

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

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het referentieveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `REFERENCE` |
| `referenceProjectId` | String | Nee | ID van het project om naar te verwijzen |
| `referenceMultiple` | Boolean | Nee | Meerdere recordselectie toestaan (standaard: false) |
| `referenceFilter` | TodoFilterInput | Nee | Filtercriteria voor verwezen records |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

**Opmerking**: Aangepaste velden worden automatisch gekoppeld aan het project op basis van de huidige projectcontext van de gebruiker.

## Referentieconfiguratie

### Enkele vs Meerdere Referenties

**Enkele Referentie (standaard):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Gebruikers kunnen één record selecteren uit het verwezen project
- Retourneert een enkel Todo-object

**Meerdere Referenties:**
```graphql
{
  referenceMultiple: true
}
```
- Gebruikers kunnen meerdere records selecteren uit het verwezen project
- Retourneert een array van Todo-objecten

### Referentiefiltering

Gebruik `referenceFilter` om te beperken welke records kunnen worden geselecteerd:

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

## Referentiewaarden Instellen

### Enkele Referentie

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Meerdere Referenties

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

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het referentie aangepaste veld |
| `customFieldReferenceTodoIds` | [String!] | ✅ Ja | Array van verwezen record-ID's |

## Records Maken met Referenties

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

## Responsvelden

### TodoCustomField Respons

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | ID! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het referentieveld |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

**Opmerking**: Verweze todos zijn toegankelijk via `customField.selectedTodos`, niet direct op TodoCustomField.

### Verweze Todo Velden

Elke verwezen Todo omvat:

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | ID! | Unieke identificatie van het verwezen record |
| `title` | String! | Titel van het verwezen record |
| `status` | TodoStatus! | Huidige status (ACTIEF, VOLTOOID, enz.) |
| `description` | String | Beschrijving van het verwezen record |
| `dueDate` | DateTime | Vervaldatum indien ingesteld |
| `assignees` | [User!] | Toegewezen gebruikers |
| `tags` | [Tag!] | Geassocieerde tags |
| `project` | Project! | Project dat het verwezen record bevat |

## Referentiegegevens Opvragen

### Basisquery

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

### Geavanceerde Query met Geneste Gegevens

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

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Belangrijk**: Gebruikers moeten kijkmachtigingen hebben op het verwezen project om de gelinkte records te zien.

## Cross-Project Toegang

### Projectzichtbaarheid

- Gebruikers kunnen alleen records verwijzen uit projecten waar ze toegang toe hebben
- Verweze records respecteren de machtigingen van het oorspronkelijke project
- Wijzigingen aan verweze records verschijnen in realtime
- Het verwijderen van verweze records verwijdert ze uit referentievelden

### Machtigingsovererving

- Referentievelden erven machtigingen van beide projecten
- Gebruikers hebben kijktoegang nodig tot het verwezen project
- Bewerkingmachtigingen zijn gebaseerd op de regels van het huidige project
- Verweze gegevens zijn alleen-lezen in de context van het referentieveld

## Foutresponsen

### Ongeldig Referentieproject

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

### Verwezend Record Niet Gevonden

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

### Toegang Geweigerd

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

### Veldontwerp

1. **Duidelijke naamgeving** - Gebruik beschrijvende namen die de relatie aangeven
2. **Geschikte filtering** - Stel filters in om alleen relevante records te tonen
3. **Overweeg machtigingen** - Zorg ervoor dat gebruikers toegang hebben tot verwezen projecten
4. **Documenteer relaties** - Geef duidelijke beschrijvingen van de verbinding

### Prestatieoverwegingen

1. **Beperk referentiebereik** - Gebruik filters om het aantal selecteerbare records te verminderen
2. **Vermijd diepe nesting** - Maak geen complexe ketens van referenties
3. **Overweeg caching** - Verweze gegevens worden gecached voor prestatie
4. **Monitor gebruik** - Volg hoe referenties worden gebruikt in projecten

### Gegevensintegriteit

1. **Omgaan met verwijderingen** - Plan voor wanneer verweze records worden verwijderd
2. **Valideer machtigingen** - Zorg ervoor dat gebruikers toegang hebben tot verwezen projecten
3. **Update afhankelijkheden** - Overweeg de impact bij het wijzigen van verweze records
4. **Audit trails** - Volg referentierelaties voor naleving

## Veelvoorkomende Gebruikscases

### Projectafhankelijkheden

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

### Klantvereisten

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

### Hulpbronnenallocatie

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

### Kwaliteitsborging

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

## Integratie met Lookup

Referentievelden werken met [Lookup-velden](/api/custom-fields/lookup) om gegevens uit verwezen records te halen. Lookup-velden kunnen waarden extraheren uit records die zijn geselecteerd in referentievelden, maar ze zijn alleen gegevensextractoren (geen aggregatiefuncties zoals SOM worden ondersteund).

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

## Beperkingen

- Verweze projecten moeten toegankelijk zijn voor de gebruiker
- Wijzigingen in de machtigingen van verwezen projecten beïnvloeden de toegang tot referentievelden
- Diepe nesting van referenties kan de prestaties beïnvloeden
- Geen ingebouwde validatie voor circulaire referenties
- Geen automatische beperking die dezelfde-projectreferenties voorkomt
- Filtervalidatie wordt niet afgedwongen bij het instellen van referentiewaarden

## Gerelateerde Bronnen

- [Lookup Velden](/api/custom-fields/lookup) - Gegevens extraheren uit verwezen records
- [Projecten API](/api/projects) - Beheren van projecten die verwijzingen bevatten
- [Records API](/api/records) - Werken met records die verwijzingen hebben
- [Overzicht van Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten