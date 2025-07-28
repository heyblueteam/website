---
title: Campo Personalizzato di Durata del Tempo
description: Crea campi di durata del tempo calcolati che tracciano il tempo tra gli eventi nel tuo flusso di lavoro
---

I campi personalizzati di durata del tempo calcolano e visualizzano automaticamente la durata tra due eventi nel tuo flusso di lavoro. Sono ideali per monitorare i tempi di elaborazione, i tempi di risposta, i tempi di ciclo o qualsiasi metrica basata sul tempo nei tuoi progetti.

## Esempio Base

Crea un semplice campo di durata del tempo che traccia quanto tempo impiegano i compiti per essere completati:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Esempio Avanzato

Crea un campo di durata del tempo complesso che traccia il tempo tra le modifiche ai campi personalizzati con un obiettivo SLA:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Parametri di Input

### CreateCustomFieldInput (TIME_DURATION)

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di durata |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `TIME_DURATION` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Sì | Come visualizzare la durata |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Sì | Configurazione dell'evento di inizio |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Sì | Configurazione dell'evento di fine |
| `timeDurationTargetTime` | Float | No | Durata target in secondi per il monitoraggio SLA |

### CustomFieldTimeDurationInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ Sì | Tipo di evento da tracciare |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Sì | Occorrenza `FIRST` o `LAST` |
| `customFieldId` | String | Conditional | Richiesto per il tipo `TODO_CUSTOM_FIELD` |
| `customFieldOptionIds` | [String!] | Conditional | Richiesto per le modifiche ai campi di selezione |
| `todoListId` | String | Conditional | Richiesto per il tipo `TODO_MOVED` |
| `tagId` | String | Conditional | Richiesto per il tipo `TODO_TAG_ADDED` |
| `assigneeId` | String | Conditional | Richiesto per il tipo `TODO_ASSIGNEE_ADDED` |

### Valori di CustomFieldTimeDurationType

| Valore | Descrizione |
|-------|-------------|
| `TODO_CREATED_AT` | Quando il record è stato creato |
| `TODO_CUSTOM_FIELD` | Quando un valore di campo personalizzato è cambiato |
| `TODO_DUE_DATE` | Quando la data di scadenza è stata impostata |
| `TODO_MARKED_AS_COMPLETE` | Quando il record è stato contrassegnato come completo |
| `TODO_MOVED` | Quando il record è stato spostato in un'altra lista |
| `TODO_TAG_ADDED` | Quando un tag è stato aggiunto al record |
| `TODO_ASSIGNEE_ADDED` | Quando un assegnatario è stato aggiunto al record |

### Valori di CustomFieldTimeDurationCondition

| Valore | Descrizione |
|-------|-------------|
| `FIRST` | Usa la prima occorrenza dell'evento |
| `LAST` | Usa l'ultima occorrenza dell'evento |

### Valori di CustomFieldTimeDurationDisplayType

| Valore | Descrizione | Esempio |
|-------|-------------|---------|
| `FULL_DATE` | Formato Giorni:Ore:Minuti:Secondi | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Scritto per esteso | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numerico con unità | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Durata solo in giorni | `"2.5"` (2.5 days) |
| `HOURS` | Durata solo in ore | `"60"` (60 hours) |
| `MINUTES` | Durata solo in minuti | `"3600"` (3600 minutes) |
| `SECONDS` | Durata solo in secondi | `"216000"` (216000 seconds) |

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `number` | Float | La durata in secondi |
| `value` | Float | Alias per numero (durata in secondi) |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato aggiornato l'ultimo valore |

### Risposta CustomField (TIME_DURATION)

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Formato di visualizzazione per la durata |
| `timeDurationStart` | CustomFieldTimeDuration | Configurazione dell'evento di inizio |
| `timeDurationEnd` | CustomFieldTimeDuration | Configurazione dell'evento di fine |
| `timeDurationTargetTime` | Float | Durata target in secondi (per il monitoraggio SLA) |

## Calcolo della Durata

### Come Funziona
1. **Evento di Inizio**: Il sistema monitora l'evento di inizio specificato
2. **Evento di Fine**: Il sistema monitora l'evento di fine specificato
3. **Calcolo**: Durata = Tempo di Fine - Tempo di Inizio
4. **Memorizzazione**: Durata memorizzata in secondi come numero
5. **Visualizzazione**: Formattata secondo l'impostazione `timeDurationDisplay`

### Attivatori di Aggiornamento
I valori di durata vengono ricalcolati automaticamente quando:
- I record vengono creati o aggiornati
- I valori dei campi personalizzati cambiano
- I tag vengono aggiunti o rimossi
- Gli assegnatari vengono aggiunti o rimossi
- I record vengono spostati tra le liste
- I record vengono contrassegnati come completi/incompleti

## Lettura dei Valori di Durata

### Query dei Campi di Durata
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Valori di Visualizzazione Formattati
I valori di durata vengono automaticamente formattati in base all'impostazione `timeDurationDisplay`:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Esempi di Configurazione Comuni

### Tempo di Completamento del Compito
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Durata della Modifica di Stato
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Tempo in Lista Specifica
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Tempo di Risposta per Assegnazione
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Risposte di Errore

### Configurazione Non Valida
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Campo Referenziato Non Trovato
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Opzioni Richieste Mancanti
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Note Importanti

### Calcolo Automatico
- I campi di durata sono **sola lettura** - i valori vengono calcolati automaticamente
- Non puoi impostare manualmente i valori di durata tramite API
- I calcoli avvengono in modo asincrono tramite lavori in background
- I valori si aggiornano automaticamente quando si verificano eventi di attivazione

### Considerazioni sulle Prestazioni
- I calcoli di durata sono messi in coda e elaborati in modo asincrono
- Un gran numero di campi di durata può influenzare le prestazioni
- Considera la frequenza degli eventi di attivazione quando progetti i campi di durata
- Usa condizioni specifiche per evitare ricalcoli non necessari

### Valori Null
I campi di durata mostreranno `null` quando:
- L'evento di inizio non è ancora avvenuto
- L'evento di fine non è ancora avvenuto
- La configurazione fa riferimento a entità non esistenti
- Il calcolo incontra un errore

## Migliori Pratiche

### Progettazione della Configurazione
- Usa tipi di eventi specifici piuttosto che generici quando possibile
- Scegli `FIRST` rispetto a `LAST` condizioni appropriate in base al tuo flusso di lavoro
- Testa i calcoli di durata con dati di esempio prima del deployment
- Documenta la logica del tuo campo di durata per i membri del team

### Formattazione della Visualizzazione
- Usa `FULL_DATE_SUBSTRING` per il formato più leggibile
- Usa `FULL_DATE` per una visualizzazione compatta e di larghezza costante
- Usa `FULL_DATE_STRING` per report e documenti formali
- Usa `DAYS`, `HOURS`, `MINUTES`, o `SECONDS` per visualizzazioni numeriche semplici
- Considera i vincoli di spazio della tua interfaccia utente quando scegli il formato

### Monitoraggio SLA con Tempo Target
Quando usi `timeDurationTargetTime`:
- Imposta la durata target in secondi
- Confronta la durata effettiva con quella target per la conformità SLA
- Usa nei dashboard per evidenziare gli elementi scaduti
- Esempio: SLA di risposta di 24 ore = 86400 secondi

### Integrazione nel Flusso di Lavoro
- Progetta i campi di durata per corrispondere ai tuoi processi aziendali effettivi
- Usa i dati di durata per il miglioramento e l'ottimizzazione dei processi
- Monitora le tendenze di durata per identificare i colli di bottiglia nel flusso di lavoro
- Imposta avvisi per le soglie di durata se necessario

## Casi d'Uso Comuni

1. **Prestazioni del Processo**
   - Tempi di completamento dei compiti
   - Tempi di ciclo di revisione
   - Tempi di elaborazione delle approvazioni
   - Tempi di risposta

2. **Monitoraggio SLA**
   - Tempo per la prima risposta
   - Tempi di risoluzione
   - Tempi di escalation
   - Conformità al livello di servizio

3. **Analisi del Flusso di Lavoro**
   - Identificazione dei colli di bottiglia
   - Ottimizzazione dei processi
   - Metriche delle prestazioni del team
   - Tempistiche per il controllo qualità

4. **Gestione dei Progetti**
   - Durate delle fasi
   - Tempistiche dei traguardi
   - Tempo di allocazione delle risorse
   - Tempistiche di consegna

## Limitazioni

- I campi di durata sono **sola lettura** e non possono essere impostati manualmente
- I valori vengono calcolati in modo asincrono e potrebbero non essere immediatamente disponibili
- Richiede che i trigger degli eventi siano impostati correttamente nel tuo flusso di lavoro
- Non può calcolare durate per eventi che non si sono verificati
- Limitato al monitoraggio del tempo tra eventi discreti (non monitoraggio del tempo continuo)
- Nessun avviso o notifica SLA integrati
- Non può aggregare più calcoli di durata in un singolo campo

## Risorse Correlate

- [Campi Numerici](/api/custom-fields/number) - Per valori numerici manuali
- [Campi Data](/api/custom-fields/date) - Per il monitoraggio di date specifiche
- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali
- [Automazioni](/api/automations) - Per attivare azioni basate su soglie di durata