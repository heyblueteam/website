---
title: Campo Personalizzato Checkbox
description: Crea campi checkbox booleani per dati sì/no o vero/falso
---

I campi personalizzati checkbox forniscono un input booleano semplice (vero/falso) per i compiti. Sono perfetti per scelte binarie, indicatori di stato o per tenere traccia del completamento di qualcosa.

## Esempio di Base

Crea un semplice campo checkbox:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo checkbox con descrizione e validazione:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
  }) {
    id
    name
    type
    description
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato della checkbox |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `CHECKBOX` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

## Impostazione dei Valori Checkbox

Per impostare o aggiornare un valore checkbox su un compito:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Per deselezionare una checkbox:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### Parametri di SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del compito da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato checkbox |
| `checked` | Boolean | No | Vero per selezionare, falso per deselezionare |

## Creazione di Compiti con Valori Checkbox

Quando si crea un nuovo compito con valori checkbox:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Valori Stringa Accettati

Quando si creano compiti, i valori checkbox devono essere passati come stringhe:

| Valore Stringa | Risultato |
|----------------|-----------|
| `"true"` | ✅ Selezionato (case-sensitive) |
| `"1"` | ✅ Selezionato |
| `"checked"` | ✅ Selezionato (case-sensitive) |
| Any other value | ❌ Deselezionato |

**Nota**: I confronti delle stringhe durante la creazione dei compiti sono sensibili al maiuscolo. I valori devono corrispondere esattamente a `"true"`, `"1"`, o `"checked"` per risultare in uno stato selezionato.

## Campi di Risposta

### TodoCustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | Identificatore unico per il valore del campo |
| `uid` | String! | Identificatore unico alternativo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `checked` | Boolean | Lo stato della checkbox (vero/falso/null) |
| `todo` | Todo! | Il compito a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

## Integrazione Automazione

I campi checkbox attivano diversi eventi di automazione in base ai cambiamenti di stato:

| Azione | Evento Attivato | Descrizione |
|--------|----------------|-------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Attivato quando la checkbox è selezionata |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Attivato quando la checkbox è deselezionata |

Questo ti consente di creare automazioni che rispondono ai cambiamenti di stato della checkbox, come ad esempio:
- Inviare notifiche quando gli elementi sono approvati
- Spostare compiti quando le checkbox di revisione sono selezionate
- Aggiornare campi correlati in base agli stati delle checkbox

## Importazione/Esportazione Dati

### Importazione di Valori Checkbox

Quando si importano dati tramite CSV o altri formati:
- `"true"`, `"yes"` → Selezionato (case-insensitive)
- Qualsiasi altro valore (inclusi `"false"`, `"no"`, `"0"`, vuoto) → Deselezionato

### Esportazione di Valori Checkbox

Quando si esportano dati:
- Le checkbox selezionate vengono esportate come `"X"`
- Le checkbox deselezionate vengono esportate come stringa vuota `""`

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Risposte di Errore

### Tipo di Valore Non Valido
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Campo Non Trovato
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

## Migliori Pratiche

### Convenzioni di Nominazione
- Usa nomi chiari e orientati all'azione: "Approvato", "Revisionato", "È Completo"
- Evita nomi negativi che confondono gli utenti: preferisci "È Attivo" piuttosto che "È Inattivo"
- Sii specifico su cosa rappresenta la checkbox

### Quando Usare le Checkbox
- **Scelte binarie**: Sì/No, Vero/Falso, Fatto/Non Fatto
- **Indicatori di stato**: Approvato, Revisionato, Pubblicato
- **Flag di funzionalità**: Ha Supporto Prioritario, Richiede Firma
- **Tracciamento semplice**: Email Inviata, Fattura Pagata, Articolo Spedito

### Quando NON Usare le Checkbox
- Quando hai bisogno di più di due opzioni (usa SELECT_SINGLE invece)
- Per dati numerici o testuali (usa campi NUMBER o TEXT)
- Quando hai bisogno di tenere traccia di chi l'ha selezionata o quando (usa log di audit)

## Casi d'Uso Comuni

1. **Flussi di Approvazione**
   - "Manager Approvato"
   - "Firma del Cliente"
   - "Revisione Legale Completa"

2. **Gestione dei Compiti**
   - "È Bloccato"
   - "Pronto per Revisione"
   - "Alta Priorità"

3. **Controllo Qualità**
   - "QA Superato"
   - "Documentazione Completa"
   - "Test Scritti"

4. **Flag Amministrativi**
   - "Fattura Inviata"
   - "Contratto Firmato"
   - "Follow-up Richiesto"

## Limitazioni

- I campi checkbox possono memorizzare solo valori vero/falso (niente tri-stato o null dopo l'impostazione iniziale)
- Nessuna configurazione del valore predefinito (inizia sempre come null fino a quando non viene impostato)
- Non possono memorizzare metadati aggiuntivi come chi l'ha selezionata o quando
- Nessuna visibilità condizionale basata su altri valori di campo

## Risorse Correlate

- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali sui campi personalizzati
- [API Automazioni](/api/automations) - Crea automazioni attivate dai cambiamenti delle checkbox