---
title: Campo Personalizzato a Selezione Singola
description: Crea campi a selezione singola per consentire agli utenti di scegliere un'opzione da un elenco predefinito
---

I campi personalizzati a selezione singola consentono agli utenti di scegliere esattamente un'opzione da un elenco predefinito. Sono ideali per campi di stato, categorie, priorità o qualsiasi scenario in cui dovrebbe essere effettuata solo una scelta da un insieme controllato di opzioni.

## Esempio di Base

Crea un semplice campo a selezione singola:

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

## Esempio Avanzato

Crea un campo a selezione singola con opzioni predefinite:

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

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo a selezione singola |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `SELECT_SINGLE` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | No | Opzioni iniziali per il campo |

### CreateCustomFieldOptionInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `title` | String! | ✅ Sì | Testo visualizzato per l'opzione |
| `color` | String | No | Codice colore esadecimale per l'opzione |

## Aggiunta di Opzioni a Campi Esistenti

Aggiungi nuove opzioni a un campo a selezione singola esistente:

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

## Impostazione dei Valori a Selezione Singola

Per impostare l'opzione selezionata su un record:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### Parametri di SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato a selezione singola |
| `customFieldOptionId` | String | No | ID dell'opzione da selezionare (preferito per la selezione singola) |
| `customFieldOptionIds` | [String!] | No | Array di ID delle opzioni (usa il primo elemento per la selezione singola) |

## Interrogazione dei Valori a Selezione Singola

Interroga il valore a selezione singola di un record:

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

Il campo `value` restituisce un oggetto JSON con i dettagli dell'opzione selezionata.

## Creazione di Record con Valori a Selezione Singola

Quando si crea un nuovo record con valori a selezione singola:

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

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `value` | JSON | Contiene l'oggetto dell'opzione selezionata con id, titolo, colore, posizione |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato per l'ultima volta il valore |

### Risposta CustomFieldOption

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per l'opzione |
| `title` | String! | Testo visualizzato per l'opzione |
| `color` | String | Codice colore esadecimale per la rappresentazione visiva |
| `position` | Float | Ordine di ordinamento per l'opzione |
| `customField` | CustomField! | Il campo personalizzato a cui appartiene questa opzione |

### Risposta CustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il campo |
| `name` | String! | Nome visualizzato del campo a selezione singola |
| `type` | CustomFieldType! | Sempre `SELECT_SINGLE` |
| `description` | String | Testo di aiuto per il campo |
| `customFieldOptions` | [CustomFieldOption!] | Tutte le opzioni disponibili |

## Formato del Valore

### Formato di Input
- **Parametro API**: Usa `customFieldOptionId` per l'ID dell'opzione singola
- **Alternativa**: Usa `customFieldOptionIds` array (prende il primo elemento)
- **Cancellazione della Selezione**: Ometti entrambi i campi o passa valori vuoti

### Formato di Output
- **Risposta GraphQL**: Oggetto JSON nel campo `value` contenente {id, titolo, colore, posizione}
- **Registro Attività**: Titolo dell'opzione come stringa
- **Dati di Automazione**: Titolo dell'opzione come stringa

## Comportamento di Selezione

### Selezione Esclusiva
- Impostare una nuova opzione rimuove automaticamente la selezione precedente
- Solo un'opzione può essere selezionata alla volta
- Impostare `null` o un valore vuoto cancella la selezione

### Logica di Fallback
- Se viene fornito un array `customFieldOptionIds`, viene utilizzata solo la prima opzione
- Questo garantisce la compatibilità con i formati di input a selezione multipla
- Array vuoti o valori nulli cancellano la selezione

## Gestione delle Opzioni

### Aggiorna Proprietà dell'Opzione
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

### Elimina Opzione
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**Nota**: Eliminare un'opzione la cancellerà da tutti i record in cui è stata selezionata.

### Riordina Opzioni
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

## Regole di Validazione

### Validazione dell'Opzione
- L'ID dell'opzione fornito deve esistere
- L'opzione deve appartenere al campo personalizzato specificato
- Solo un'opzione può essere selezionata (imposto automaticamente)
- Valori nulli/vuoti sono validi (nessuna selezione)

### Validazione del Campo
- Deve avere almeno un'opzione definita per essere utilizzabile
- I titoli delle opzioni devono essere unici all'interno del campo
- I codici colore devono essere in formato esadecimale valido (se forniti)

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Risposte di Errore

### ID Opzione Non Valido
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

### L'Opzione Non Appartiene al Campo
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

### Campo Non Trovato
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

### Impossibile Analizzare il Valore
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

## Migliori Pratiche

### Design dell'Opzione
- Usa titoli di opzione chiari e descrittivi
- Applica codifica colore significativa
- Mantieni le liste di opzioni focalizzate e pertinenti
- Ordina le opzioni in modo logico (per priorità, frequenza, ecc.)

### Modelli di Campo di Stato
- Usa flussi di lavoro di stato coerenti tra i progetti
- Considera la progressione naturale delle opzioni
- Includi stati finali chiari (Fatto, Annullato, ecc.)
- Usa colori che riflettono il significato dell'opzione

### Gestione dei Dati
- Rivedi e pulisci periodicamente le opzioni non utilizzate
- Usa convenzioni di denominazione coerenti
- Considera l'impatto dell'eliminazione delle opzioni sui record esistenti
- Pianifica aggiornamenti e migrazioni delle opzioni

## Casi d'Uso Comuni

1. **Stato e Flusso di Lavoro**
   - Stato del compito (Da Fare, In Corso, Fatto)
   - Stato di approvazione (In Attesa, Approvato, Rifiutato)
   - Fase del progetto (Pianificazione, Sviluppo, Test, Rilasciato)
   - Stato di risoluzione dei problemi

2. **Classificazione e Categorizzazione**
   - Livelli di priorità (Basso, Medio, Alto, Critico)
   - Tipi di compito (Bug, Funzionalità, Miglioramento, Documentazione)
   - Categorie di progetto (Interno, Cliente, Ricerca)
   - Assegnazioni di dipartimento

3. **Qualità e Valutazione**
   - Stato di revisione (Non Iniziato, In Revisione, Approvato)
   - Valutazioni di qualità (Scarso, Discreto, Buono, Eccellente)
   - Livelli di rischio (Basso, Medio, Alto)
   - Livelli di fiducia

4. **Assegnazione e Proprietà**
   - Assegnazioni di team
   - Proprietà del dipartimento
   - Assegnazioni basate sul ruolo
   - Assegnazioni regionali

## Funzionalità di Integrazione

### Con Automazioni
- Attiva azioni quando vengono selezionate opzioni specifiche
- Instrada il lavoro in base alle categorie selezionate
- Invia notifiche per cambiamenti di stato
- Crea flussi di lavoro condizionali basati sulle selezioni

### Con Ricerche
- Filtra i record in base alle opzioni selezionate
- Riferisci i dati delle opzioni da altri record
- Crea report basati sulle selezioni delle opzioni
- Raggruppa i record per valori selezionati

### Con Moduli
- Controlli di input a discesa
- Interfacce a pulsante radio
- Validazione e filtraggio delle opzioni
- Visualizzazione condizionale dei campi in base alle selezioni

## Monitoraggio delle Attività

Le modifiche ai campi a selezione singola vengono tracciate automaticamente:
- Mostra le selezioni di opzioni vecchie e nuove
- Visualizza i titoli delle opzioni nel registro delle attività
- Timestamp per tutte le modifiche di selezione
- Attribuzione utente per le modifiche

## Differenze rispetto alla Selezione Multipla

| Caratteristica | Selezione Singola | Selezione Multipla |
|----------------|-------------------|--------------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Limitazioni

- Solo un'opzione può essere selezionata alla volta
- Nessuna struttura di opzione gerarchica o nidificata
- Le opzioni sono condivise tra tutti i record che utilizzano il campo
- Nessuna analisi delle opzioni integrata o monitoraggio dell'uso
- I codici colore sono solo per visualizzazione, senza impatto funzionale
- Non è possibile impostare permessi diversi per ogni opzione

## Risorse Correlate

- [Campi a Selezione Multipla](/api/custom-fields/select-multi) - Per selezioni a scelta multipla
- [Campi Checkbox](/api/custom-fields/checkbox) - Per scelte booleane semplici
- [Campi di Testo](/api/custom-fields/text-single) - Per input di testo libero
- [Panoramica dei Campi Personalizzati](/api/custom-fields/1.index) - Concetti generali