---
title: Campo Personalizzato Multi-Select
description: Crea campi multi-select per consentire agli utenti di scegliere più opzioni da un elenco predefinito
---

I campi personalizzati multi-select consentono agli utenti di scegliere più opzioni da un elenco predefinito. Sono ideali per categorie, tag, competenze, funzionalità o qualsiasi scenario in cui siano necessarie più selezioni da un insieme controllato di opzioni.

## Esempio di Base

Crea un semplice campo multi-select:

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

## Esempio Avanzato

Crea un campo multi-select e poi aggiungi opzioni separatamente:

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

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo multi-select |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `SELECT_MULTI` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |
| `projectId` | String! | ✅ Sì | ID del progetto per questo campo |

### CreateCustomFieldOptionInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato |
| `title` | String! | ✅ Sì | Testo visualizzato per l'opzione |
| `color` | String | No | Colore per l'opzione (qualsiasi stringa) |
| `position` | Float | No | Ordine di ordinamento per l'opzione |

## Aggiungere Opzioni a Campi Esistenti

Aggiungi nuove opzioni a un campo multi-select esistente:

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

## Impostare Valori Multi-Select

Per impostare più opzioni selezionate su un record:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### Parametri SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato multi-select |
| `customFieldOptionIds` | [String!] | ✅ Sì | Array di ID delle opzioni da selezionare |

## Creare Record con Valori Multi-Select

Quando si crea un nuovo record con valori multi-select:

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

## Campi di Risposta

### TodoCustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `selectedOptions` | [CustomFieldOption!] | Array di opzioni selezionate |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato per l'ultima volta il valore |

### CustomFieldOption Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per l'opzione |
| `title` | String! | Testo visualizzato per l'opzione |
| `color` | String | Codice colore esadecimale per la rappresentazione visiva |
| `position` | Float | Ordine di ordinamento per l'opzione |
| `customField` | CustomField! | Il campo personalizzato a cui appartiene questa opzione |

### CustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il campo |
| `name` | String! | Nome visualizzato del campo multi-select |
| `type` | CustomFieldType! | Sempre `SELECT_MULTI` |
| `description` | String | Testo di aiuto per il campo |
| `customFieldOptions` | [CustomFieldOption!] | Tutte le opzioni disponibili |

## Formato del Valore

### Formato di Input
- **Parametro API**: Array di ID delle opzioni (`["option1", "option2", "option3"]`)
- **Formato Stringa**: ID delle opzioni separati da virgola (`"option1,option2,option3"`)

### Formato di Output
- **Risposta GraphQL**: Array di oggetti CustomFieldOption
- **Registro Attività**: Titoli delle opzioni separati da virgola
- **Dati di Automazione**: Array di titoli delle opzioni

## Gestione delle Opzioni

### Aggiornare Proprietà dell'Opzione
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

### Eliminare Opzione
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### Riordinare Opzioni
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

## Regole di Validazione

### Validazione dell'Opzione
- Tutti gli ID delle opzioni forniti devono esistere
- Le opzioni devono appartenere al campo personalizzato specificato
- Solo i campi SELECT_MULTI possono avere più opzioni selezionate
- L'array vuoto è valido (nessuna selezione)

### Validazione del Campo
- Deve avere almeno un'opzione definita per essere utilizzabile
- I titoli delle opzioni devono essere unici all'interno del campo
- Il campo colore accetta qualsiasi valore di stringa (nessuna validazione esadecimale)

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Risposte di Errore

### ID Opzione Non Valido
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Opzioni Multiple su Campo Non Multi
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

## Migliori Pratiche

### Progettazione dell'Opzione
- Utilizza titoli delle opzioni descrittivi e concisi
- Applica schemi di codifica colore coerenti
- Mantieni le liste delle opzioni gestibili (tipicamente 3-20 opzioni)
- Ordina le opzioni in modo logico (alfabeticamente, per frequenza, ecc.)

### Gestione dei Dati
- Rivedi e pulisci periodicamente le opzioni non utilizzate
- Utilizza convenzioni di denominazione coerenti tra i progetti
- Considera la riutilizzabilità delle opzioni quando crei campi
- Pianifica aggiornamenti e migrazioni delle opzioni

### Esperienza Utente
- Fornisci descrizioni chiare dei campi
- Usa colori per migliorare la distinzione visiva
- Raggruppa insieme opzioni correlate
- Considera selezioni predefinite per casi comuni

## Casi d'Uso Comuni

1. **Gestione Progetti**
   - Categorie e tag delle attività
   - Livelli e tipi di priorità
   - Assegnazioni dei membri del team
   - Indicatori di stato

2. **Gestione dei Contenuti**
   - Categorie e argomenti degli articoli
   - Tipi e formati di contenuto
   - Canali di pubblicazione
   - Flussi di lavoro di approvazione

3. **Assistenza Clienti**
   - Categorie e tipi di problemi
   - Prodotti o servizi interessati
   - Metodi di risoluzione
   - Segmenti di clienti

4. **Sviluppo Prodotto**
   - Categorie di funzionalità
   - Requisiti tecnici
   - Ambienti di test
   - Canali di rilascio

## Caratteristiche di Integrazione

### Con Automazioni
- Attiva azioni quando vengono selezionate opzioni specifiche
- Smista il lavoro in base alle categorie selezionate
- Invia notifiche per selezioni ad alta priorità
- Crea attività di follow-up basate su combinazioni di opzioni

### Con Ricerche
- Filtra i record in base alle opzioni selezionate
- Aggrega dati attraverso le selezioni delle opzioni
- Riferisci i dati delle opzioni da altri record
- Crea report basati su combinazioni di opzioni

### Con Moduli
- Controlli di input multi-select
- Validazione e filtraggio delle opzioni
- Caricamento dinamico delle opzioni
- Visualizzazione condizionale dei campi

## Monitoraggio delle Attività

Le modifiche ai campi multi-select vengono automaticamente tracciate:
- Mostra opzioni aggiunte e rimosse
- Visualizza i titoli delle opzioni nel registro delle attività
- Timestamp per tutte le modifiche di selezione
- Attribuzione all'utente per le modifiche

## Limitazioni

- Il limite pratico massimo delle opzioni dipende dalle prestazioni dell'interfaccia utente
- Nessuna struttura di opzione gerarchica o nidificata
- Le opzioni sono condivise tra tutti i record che utilizzano il campo
- Nessuna analisi delle opzioni integrata o monitoraggio dell'uso
- Il campo colore accetta qualsiasi stringa (nessuna validazione esadecimale)
- Non è possibile impostare permessi diversi per opzione
- Le opzioni devono essere create separatamente, non in linea con la creazione del campo
- Nessuna mutazione di riordino dedicata (usa editCustomFieldOption con posizione)

## Risorse Correlate

- [Campi Single-Select](/api/custom-fields/select-single) - Per selezioni a scelta singola
- [Campi Checkbox](/api/custom-fields/checkbox) - Per scelte booleane semplici
- [Campi di Testo](/api/custom-fields/text-single) - Per input di testo libero
- [Panoramica dei Campi Personalizzati](/api/custom-fields/2.list-custom-fields) - Concetti generali