---
title: Cerca Campo Personalizzato
description: Crea campi di ricerca che estraggono automaticamente dati da record di riferimento
---

I campi personalizzati di ricerca estraggono automaticamente dati dai record a cui si fa riferimento tramite [Campi di riferimento](/api/custom-fields/reference), visualizzando informazioni dai record collegati senza copia manuale. Si aggiornano automaticamente quando i dati di riferimento cambiano.

## Esempio di Base

Crea un campo di ricerca per visualizzare tag dai record di riferimento:

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

## Esempio Avanzato

Crea un campo di ricerca per estrarre valori di campi personalizzati dai record di riferimento:

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

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di ricerca |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Sì | Configurazione della ricerca |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

## Configurazione della Ricerca

### CustomFieldLookupOptionInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `referenceId` | String! | ✅ Sì | ID del campo di riferimento da cui estrarre i dati |
| `lookupId` | String | No | ID del campo personalizzato specifico da cercare (richiesto per il tipo TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Sì | Tipo di dati da estrarre dai record di riferimento |

## Tipi di Ricerca

### Valori CustomFieldLookupType

| Tipo | Descrizione | Restituisce |
|------|-------------|-------------|
| `TODO_DUE_DATE` | Date di scadenza dai todo di riferimento | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Date di creazione dai todo di riferimento | Array of creation timestamps |
| `TODO_UPDATED_AT` | Date dell'ultimo aggiornamento dai todo di riferimento | Array of update timestamps |
| `TODO_TAG` | Tag dai todo di riferimento | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Assegnatari dai todo di riferimento | Array of user objects |
| `TODO_DESCRIPTION` | Descrizioni dai todo di riferimento | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Nomi delle liste todo dai todo di riferimento | Array of list titles |
| `TODO_CUSTOM_FIELD` | Valori dei campi personalizzati dai todo di riferimento | Array of values based on the field type |

## Campi di Risposta

### Risposta CustomField (per campi di ricerca)

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il campo |
| `name` | String! | Nome visualizzato del campo di ricerca |
| `type` | CustomFieldType! | Sarà `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Configurazione e risultati della ricerca |
| `createdAt` | DateTime! | Quando è stato creato il campo |
| `updatedAt` | DateTime! | Quando è stato aggiornato l'ultimo campo |

### Struttura CustomFieldLookupOption

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | Tipo di ricerca in corso |
| `lookupResult` | JSON | I dati estratti dai record di riferimento |
| `reference` | CustomField | Il campo di riferimento utilizzato come sorgente |
| `lookup` | CustomField | Il campo specifico da cercare (per TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Il campo di ricerca principale |
| `parentLookup` | CustomField | Ricerca principale nella catena (per ricerche annidate) |

## Come Funzionano le Ricerche

1. **Estrazione Dati**: Le ricerche estraggono dati specifici da tutti i record collegati tramite un campo di riferimento
2. **Aggiornamenti Automatici**: Quando i record di riferimento cambiano, i valori di ricerca si aggiornano automaticamente
3. **Solo Lettura**: I campi di ricerca non possono essere modificati direttamente - riflettono sempre i dati di riferimento attuali
4. **Nessun Calcolo**: Le ricerche estraggono e visualizzano i dati così come sono, senza aggregazioni o calcoli

## Ricerche TODO_CUSTOM_FIELD

Quando si utilizza il tipo `TODO_CUSTOM_FIELD`, è necessario specificare quale campo personalizzato estrarre utilizzando il parametro `lookupId`:

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

Questo estrae i valori del campo personalizzato specificato da tutti i record di riferimento.

## Interrogazione Dati di Ricerca

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

## Esempio di Risultati di Ricerca

### Risultato di Ricerca Tag
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

### Risultato di Ricerca Assegnatario
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

### Risultato di Ricerca Campo Personalizzato
I risultati variano in base al tipo di campo personalizzato che viene cercato. Ad esempio, una ricerca di campo valuta potrebbe restituire:
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

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Importante**: Gli utenti devono avere permessi di visualizzazione sia sul progetto corrente che sul progetto di riferimento per vedere i risultati della ricerca.

## Risposte di Errore

### Campo di Riferimento Non Valido
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

### Ricerca Circolare Rilevata
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

### ID di Ricerca Mancante per TODO_CUSTOM_FIELD
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

## Migliori Pratiche

1. **Nomi Chiari**: Usa nomi descrittivi che indicano quali dati vengono cercati
2. **Tipi Appropriati**: Scegli il tipo di ricerca che corrisponde alle tue esigenze di dati
3. **Prestazioni**: Le ricerche elaborano tutti i record di riferimento, quindi fai attenzione ai campi di riferimento con molti collegamenti
4. **Permessi**: Assicurati che gli utenti abbiano accesso ai progetti di riferimento affinché le ricerche funzionino

## Casi d'Uso Comuni

### Visibilità Trasversale ai Progetti
Visualizza tag, assegnatari o stati da progetti correlati senza sincronizzazione manuale.

### Monitoraggio delle Dipendenze
Mostra le date di scadenza o lo stato di completamento delle attività su cui dipende il lavoro attuale.

### Panoramica delle Risorse
Visualizza tutti i membri del team assegnati alle attività di riferimento per la pianificazione delle risorse.

### Aggregazione degli Stati
Raccogli tutti gli stati unici dalle attività correlate per vedere la salute del progetto a colpo d'occhio.

## Limitazioni

- I campi di ricerca sono di sola lettura e non possono essere modificati direttamente
- Nessuna funzione di aggregazione (SOMMA, CONTEGGIO, MEDIA) - le ricerche estraggono solo dati
- Nessuna opzione di filtraggio - tutti i record di riferimento sono inclusi
- Le catene di ricerca circolari sono prevenute per evitare loop infiniti
- I risultati riflettono i dati attuali e si aggiornano automaticamente

## Risorse Correlate

- [Campi di Riferimento](/api/custom-fields/reference) - Crea collegamenti a record per fonti di ricerca
- [Valori dei Campi Personalizzati](/api/custom-fields/custom-field-values) - Imposta valori su campi personalizzati modificabili
- [Elenca Campi Personalizzati](/api/custom-fields/list-custom-fields) - Interroga tutti i campi personalizzati in un progetto