---
title: Campo Personalizzato ID Unico
description: Crea campi identificatori unici generati automaticamente con numerazione sequenziale e formattazione personalizzata
---

I campi personalizzati ID unico generano automaticamente identificatori unici e sequenziali per i tuoi record. Sono perfetti per creare numeri di ticket, ID ordine, numeri di fattura o qualsiasi sistema di identificazione sequenziale nel tuo flusso di lavoro.

## Esempio Base

Crea un semplice campo ID unico con sequenziamento automatico:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Esempio Avanzato

Crea un campo ID unico formattato con prefisso e padding zero:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Parametri di Input

### CreateCustomFieldInput (UNIQUE_ID)

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo ID unico |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `UNIQUE_ID` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |
| `useSequenceUniqueId` | Boolean | No | Abilita il sequenziamento automatico (predefinito: falso) |
| `prefix` | String | No | Prefisso di testo per gli ID generati (es. "TASK-") |
| `sequenceDigits` | Int | No | Numero di cifre per il padding zero |
| `sequenceStartingNumber` | Int | No | Numero iniziale per la sequenza |

## Opzioni di Configurazione

### Sequenziamento Automatico (`useSequenceUniqueId`)
- **true**: Genera automaticamente ID sequenziali quando vengono creati i record
- **false** o **undefined**: Richiesta di inserimento manuale (funziona come un campo di testo)

### Prefisso (`prefix`)
- Prefisso di testo opzionale aggiunto a tutti gli ID generati
- Esempi: "TASK-", "ORD-", "BUG-", "REQ-"
- Nessun limite di lunghezza, ma mantenere ragionevole per la visualizzazione

### Cifre della Sequenza (`sequenceDigits`)
- Numero di cifre per il padding zero del numero di sequenza
- Esempio: `sequenceDigits: 3` produce `001`, `002`, `003`
- Se non specificato, non viene applicato alcun padding

### Numero Iniziale (`sequenceStartingNumber`)
- Il primo numero nella sequenza
- Esempio: `sequenceStartingNumber: 1000` inizia a 1000, 1001, 1002...
- Se non specificato, inizia a 1 (comportamento predefinito)

## Formato ID Generato

Il formato finale dell'ID combina tutte le opzioni di configurazione:

```
{prefix}{paddedSequenceNumber}
```

### Esempi di Formato

| Configurazione | ID Generati |
|----------------|-------------|
| Nessuna opzione | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Lettura dei Valori ID Unici

### Query Records con ID Unici
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Formato di Risposta
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Generazione Automatica degli ID

### Quando Vengono Generati gli ID
- **Creazione Record**: Gli ID vengono assegnati automaticamente quando vengono creati nuovi record
- **Aggiunta Campo**: Quando si aggiunge un campo UNIQUE_ID a record esistenti, viene messo in coda un lavoro in background (implementazione del lavoratore in attesa)
- **Elaborazione in Background**: La generazione degli ID per i nuovi record avviene in modo sincrono tramite trigger del database

### Processo di Generazione
1. **Trigger**: Viene creato un nuovo record o viene aggiunto un campo UNIQUE_ID
2. **Ricerca Sequenza**: Il sistema trova il prossimo numero di sequenza disponibile
3. **Assegnazione ID**: Il numero di sequenza viene assegnato al record
4. **Aggiornamento Contatore**: Il contatore di sequenza viene incrementato per i record futuri
5. **Formattazione**: L'ID viene formattato con prefisso e padding quando viene visualizzato

### Garanzie di Unicità
- **Vincoli del Database**: Vincolo unico sugli ID di sequenza all'interno di ciascun campo
- **Operazioni Atomiche**: La generazione della sequenza utilizza blocchi del database per prevenire duplicati
- **Ambito del Progetto**: Le sequenze sono indipendenti per progetto
- **Protezione da Condizioni di Gara**: Le richieste concorrenti vengono gestite in modo sicuro

## Modalità Manuale vs Automatica

### Modalità Automatica (`useSequenceUniqueId: true`)
- Gli ID vengono generati automaticamente tramite trigger del database
- La numerazione sequenziale è garantita
- La generazione atomica della sequenza previene duplicati
- Gli ID formattati combinano prefisso + numero di sequenza con padding

### Modalità Manuale (`useSequenceUniqueId: false` o `undefined`)
- Funziona come un normale campo di testo
- Gli utenti possono inserire valori personalizzati tramite `setTodoCustomField` con parametro `text`
- Nessuna generazione automatica
- Nessuna enforcement di unicità oltre ai vincoli del database

## Impostazione Valori Manuali (Solo Modalità Manuale)

Quando `useSequenceUniqueId` è falso, puoi impostare valori manualmente:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Campi di Risposta

### Risposta TodoCustomField (UNIQUE_ID)

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `sequenceId` | Int | Il numero di sequenza generato (popolato per i campi UNIQUE_ID) |
| `text` | String | Il valore di testo formattato (combina prefisso + sequenza con padding) |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando il valore è stato aggiornato l'ultima volta |

### Risposta CustomField (UNIQUE_ID)

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | Se il sequenziamento automatico è abilitato |
| `prefix` | String | Prefisso di testo per gli ID generati |
| `sequenceDigits` | Int | Numero di cifre per il padding zero |
| `sequenceStartingNumber` | Int | Numero iniziale per la sequenza |

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Risposte di Errore

### Errore di Configurazione del Campo
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Errore di Permesso
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

## Note Importanti

### ID Generati Automaticamente
- **Solo Lettura**: Gli ID generati automaticamente non possono essere modificati manualmente
- **Permanenti**: Una volta assegnati, gli ID di sequenza non cambiano
- **Cronologici**: Gli ID riflettono l'ordine di creazione
- **Scoped**: Le sequenze sono indipendenti per progetto

### Considerazioni sulle Prestazioni
- La generazione degli ID per i nuovi record è sincrona tramite trigger del database
- La generazione della sequenza utilizza `FOR UPDATE` blocchi per operazioni atomiche
- Esiste un sistema di lavoro in background, ma l'implementazione del lavoratore è in attesa
- Considera i numeri iniziali della sequenza per progetti ad alto volume

### Migrazione e Aggiornamenti
- Aggiungere il sequenziamento automatico a record esistenti mette in coda un lavoro in background (lavoratore in attesa)
- Cambiare le impostazioni della sequenza influisce solo sui record futuri
- Gli ID esistenti rimangono invariati quando vengono aggiornate le configurazioni
- I contatori di sequenza continuano dal massimo attuale

## Migliori Pratiche

### Progettazione della Configurazione
- Scegli prefissi descrittivi che non entreranno in conflitto con altri sistemi
- Usa un padding di cifre appropriato per il volume previsto
- Imposta numeri iniziali ragionevoli per evitare conflitti
- Testa la configurazione con dati di esempio prima del deployment

### Linee Guida per il Prefisso
- Mantieni i prefissi brevi e memorabili (2-5 caratteri)
- Usa maiuscole per coerenza
- Includi separatori (trattini, underscore) per leggibilità
- Evita caratteri speciali che potrebbero causare problemi in URL o sistemi

### Pianificazione della Sequenza
- Stima il volume dei tuoi record per scegliere un padding di cifre appropriato
- Considera la crescita futura quando imposti i numeri iniziali
- Pianifica diverse gamme di sequenze per diversi tipi di record
- Documenta i tuoi schemi ID per riferimento del team

## Casi d'Uso Comuni

1. **Sistemi di Supporto**
   - Numeri di ticket: `TICK-001`, `TICK-002`
   - ID caso: `CASE-2024-001`
   - Richieste di supporto: `SUP-001`

2. **Gestione Progetti**
   - ID attività: `TASK-001`, `TASK-002`
   - Elementi di sprint: `SPRINT-001`
   - Numeri di deliverable: `DEL-001`

3. **Operazioni Aziendali**
   - Numeri d'ordine: `ORD-2024-001`
   - ID fattura: `INV-001`
   - Ordini di acquisto: `PO-001`

4. **Gestione della Qualità**
   - Rapporti di bug: `BUG-001`
   - ID casi di test: `TEST-001`
   - Numeri di revisione: `REV-001`

## Funzionalità di Integrazione

### Con Automazioni
- Attiva azioni quando vengono assegnati ID unici
- Usa modelli ID nelle regole di automazione
- Riferisci ID nei modelli di email e notifiche

### Con Ricerche
- Riferisci ID unici da altri record
- Trova record per ID unico
- Visualizza identificatori di record correlati

### Con Reporting
- Raggruppa e filtra per modelli ID
- Monitora le tendenze di assegnazione ID
- Monitora l'uso della sequenza e le lacune

## Limitazioni

- **Solo Sequenziale**: Gli ID vengono assegnati in ordine cronologico
- **Nessun Gap**: I record eliminati lasciano gap nelle sequenze
- **Nessuna Riutilizzazione**: I numeri di sequenza non vengono mai riutilizzati
- **Scoped per Progetto**: Non è possibile condividere sequenze tra progetti
- **Vincoli di Formato**: Opzioni di formattazione limitate
- **Nessun Aggiornamento di Massa**: Non è possibile aggiornare in massa gli ID di sequenza esistenti
- **Nessuna Logica Personalizzata**: Non è possibile implementare regole di generazione ID personalizzate

## Risorse Correlate

- [Campi di Testo](/api/custom-fields/text-single) - Per identificatori di testo manuali
- [Campi Numerici](/api/custom-fields/number) - Per sequenze numeriche
- [Panoramica dei Campi Personalizzati](/api/custom-fields/2.list-custom-fields) - Concetti generali
- [Automazioni](/api/automations) - Per regole di automazione basate su ID