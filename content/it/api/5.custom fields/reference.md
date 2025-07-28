---
title: Campo Personalizzato di Riferimento
description: Crea campi di riferimento che collegano a record in altri progetti per relazioni tra progetti
---

I campi personalizzati di riferimento consentono di creare collegamenti tra record in progetti diversi, abilitando relazioni tra progetti e condivisione di dati. Forniscono un modo potente per connettere il lavoro correlato all'interno della struttura dei progetti della tua organizzazione.

## Esempio di Base

Crea un semplice campo di riferimento:

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

## Esempio Avanzato

Crea un campo di riferimento con filtraggio e selezione multipla:

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

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di riferimento |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `REFERENCE` |
| `referenceProjectId` | String | No | ID del progetto da riferire |
| `referenceMultiple` | Boolean | No | Consenti selezione multipla dei record (predefinito: falso) |
| `referenceFilter` | TodoFilterInput | No | Criteri di filtro per i record di riferimento |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: I campi personalizzati sono automaticamente associati al progetto in base al contesto del progetto attuale dell'utente.

## Configurazione di Riferimento

### Riferimenti Singoli vs Multipli

**Riferimento Singolo (predefinito):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Gli utenti possono selezionare un record dal progetto di riferimento
- Restituisce un singolo oggetto Todo

**Riferimenti Multipli:**
```graphql
{
  referenceMultiple: true
}
```
- Gli utenti possono selezionare più record dal progetto di riferimento
- Restituisce un array di oggetti Todo

### Filtro di Riferimento

Usa `referenceFilter` per limitare quali record possono essere selezionati:

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

## Impostazione dei Valori di Riferimento

### Riferimento Singolo

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Riferimenti Multipli

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

### Parametri di SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato di riferimento |
| `customFieldReferenceTodoIds` | [String!] | ✅ Sì | Array di ID dei record di riferimento |

## Creazione di Record con Riferimenti

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

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo di riferimento |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

**Nota**: I todo di riferimento sono accessibili tramite `customField.selectedTodos`, non direttamente su TodoCustomField.

### Campi Todo di Riferimento

Ogni Todo di riferimento include:

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | Identificatore unico del record di riferimento |
| `title` | String! | Titolo del record di riferimento |
| `status` | TodoStatus! | Stato attuale (ATTIVO, COMPLETATO, ecc.) |
| `description` | String | Descrizione del record di riferimento |
| `dueDate` | DateTime | Data di scadenza se impostata |
| `assignees` | [User!] | Utenti assegnati |
| `tags` | [Tag!] | Tag associati |
| `project` | Project! | Progetto contenente il record di riferimento |

## Interrogazione dei Dati di Riferimento

### Interrogazione di Base

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

### Interrogazione Avanzata con Dati Annidati

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

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Importante**: Gli utenti devono avere permessi di visualizzazione sul progetto di riferimento per vedere i record collegati.

## Accesso tra Progetti

### Visibilità del Progetto

- Gli utenti possono solo fare riferimento a record di progetti a cui hanno accesso
- I record di riferimento rispettano i permessi del progetto originale
- Le modifiche ai record di riferimento appaiono in tempo reale
- Eliminare record di riferimento li rimuove dai campi di riferimento

### Ereditarietà dei Permessi

- I campi di riferimento ereditano i permessi da entrambi i progetti
- Gli utenti necessitano di accesso in visualizzazione al progetto di riferimento
- I permessi di modifica si basano sulle regole del progetto attuale
- I dati di riferimento sono in sola lettura nel contesto del campo di riferimento

## Risposte di Errore

### Progetto di Riferimento Non Valido

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

### Record di Riferimento Non Trovato

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

### Permesso Negato

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

## Migliori Pratiche

### Design del Campo

1. **Nomi chiari** - Usa nomi descrittivi che indicano la relazione
2. **Filtraggio appropriato** - Imposta filtri per mostrare solo i record pertinenti
3. **Considera i permessi** - Assicurati che gli utenti abbiano accesso ai progetti di riferimento
4. **Documenta le relazioni** - Fornisci descrizioni chiare della connessione

### Considerazioni sulle Prestazioni

1. **Limita l'ambito di riferimento** - Usa filtri per ridurre il numero di record selezionabili
2. **Evita annidamenti profondi** - Non creare catene complesse di riferimenti
3. **Considera la memorizzazione nella cache** - I dati di riferimento sono memorizzati nella cache per prestazioni
4. **Monitora l'uso** - Tieni traccia di come vengono utilizzati i riferimenti tra i progetti

### Integrità dei Dati

1. **Gestisci le eliminazioni** - Pianifica quando i record di riferimento vengono eliminati
2. **Valida i permessi** - Assicurati che gli utenti possano accedere ai progetti di riferimento
3. **Aggiorna le dipendenze** - Considera l'impatto quando cambi i record di riferimento
4. **Audit trail** - Tieni traccia delle relazioni di riferimento per la conformità

## Casi d'Uso Comuni

### Dipendenze di Progetto

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

### Requisiti del Cliente

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

### Assegnazione delle Risorse

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

### Assicurazione della Qualità

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

## Integrazione con i Lookup

I campi di riferimento funzionano con [Campi di Lookup](/api/custom-fields/lookup) per estrarre dati dai record di riferimento. I campi di lookup possono estrarre valori dai record selezionati nei campi di riferimento, ma sono solo estrattori di dati (non sono supportate funzioni di aggregazione come SUM).

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

## Limitazioni

- I progetti di riferimento devono essere accessibili all'utente
- Le modifiche ai permessi del progetto di riferimento influenzano l'accesso ai campi di riferimento
- Annidamenti profondi di riferimenti possono influenzare le prestazioni
- Nessuna validazione integrata per riferimenti circolari
- Nessuna restrizione automatica che impedisca riferimenti allo stesso progetto
- La validazione dei filtri non è applicata quando si impostano i valori di riferimento

## Risorse Correlate

- [Campi di Lookup](/api/custom-fields/lookup) - Estrai dati dai record di riferimento
- [API Progetti](/api/projects) - Gestire progetti che contengono riferimenti
- [API Record](/api/records) - Lavorare con record che hanno riferimenti
- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali