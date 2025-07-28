---
title: Campo Personalizzato Data
description: Crea campi data per tracciare singole date o intervalli di date con supporto per i fusi orari
---

I campi personalizzati data ti consentono di memorizzare singole date o intervalli di date per i record. Supportano la gestione dei fusi orari, la formattazione intelligente e possono essere utilizzati per tenere traccia di scadenze, date di eventi o qualsiasi informazione basata sul tempo.

## Esempio Base

Crea un semplice campo data:

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo data di scadenza con descrizione:

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo data |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `DATE` |
| `isDueDate` | Boolean | No | Se questo campo rappresenta una data di scadenza |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: I campi personalizzati sono automaticamente associati al progetto in base al contesto del progetto attuale dell'utente. Non è richiesto alcun parametro `projectId`.

## Impostazione dei Valori di Data

I campi data possono memorizzare sia una singola data che un intervallo di date:

### Data Singola

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Intervallo di Date

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Evento Tutto il Giorno

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Parametri SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato data |
| `startDate` | DateTime | No | Data/ora di inizio in formato ISO 8601 |
| `endDate` | DateTime | No | Data/ora di fine in formato ISO 8601 |
| `timezone` | String | No | Identificatore del fuso orario (es. "America/New_York") |

**Nota**: Se viene fornito solo `startDate`, `endDate` predefinisce automaticamente lo stesso valore.

## Formati di Data

### Formato ISO 8601
Tutte le date devono essere fornite in formato ISO 8601:
- `2025-01-15T14:30:00Z` - ora UTC
- `2025-01-15T14:30:00+05:00` - Con offset del fuso orario
- `2025-01-15T14:30:00.123Z` - Con millisecondi

### Identificatori di Fuso Orario
Utilizza identificatori di fuso orario standard:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Se non viene fornito alcun fuso orario, il sistema predefinisce il fuso orario rilevato dell'utente.

## Creazione di Record con Valori di Data

Quando crei un nuovo record con valori di data:

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### Formati di Input Supportati

Quando crei record, le date possono essere fornite in vari formati:

| Formato | Esempio | Risultato |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | Identificatore univoco per il valore del campo |
| `uid` | String! | Stringa identificativa univoca |
| `customField` | CustomField! | La definizione del campo personalizzato (contiene i valori di data) |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

**Importante**: I valori di data (`startDate`, `endDate`, `timezone`) sono accessibili tramite il campo `customField.value`, non direttamente su TodoCustomField.

### Struttura dell'Oggetto Valore

I valori di data vengono restituiti tramite il campo `customField.value` come oggetto JSON:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**Nota**: Il campo `value` è di tipo `CustomField`, non di `TodoCustomField`.

## Interrogazione dei Valori di Data

Quando interroghi record con campi personalizzati data, accedi ai valori di data tramite il campo `customField.value`:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

La risposta includerà i valori di data nel campo `value`:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## Intelligenza di Visualizzazione della Data

Il sistema formatta automaticamente le date in base all'intervallo:

| Scenario | Formato di Visualizzazione |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (nessun orario mostrato) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Rilevamento di tutto il giorno**: Gli eventi dalle 00:00 alle 23:59 vengono automaticamente rilevati come eventi di tutto il giorno.

## Gestione dei Fusi Orari

### Archiviazione
- Tutte le date sono memorizzate in UTC nel database
- Le informazioni sul fuso orario sono conservate separatamente
- La conversione avviene in fase di visualizzazione

### Migliori Pratiche
- Fornisci sempre il fuso orario per precisione
- Usa fusi orari coerenti all'interno di un progetto
- Considera le posizioni degli utenti per team globali

### Fusi Orari Comuni

| Regione | ID Fuso Orario | Offset UTC |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filtraggio e Interrogazione

I campi data supportano il filtraggio complesso:

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### Controllo dei Campi Data Vuoti

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### Operatori Supportati

| Operatore | Utilizzo | Descrizione |
|----------|-------|-------------|
| `EQ` | Con dateRange | La data sovrappone l'intervallo specificato (qualsiasi intersezione) |
| `NE` | Con dateRange | La data non sovrappone l'intervallo |
| `IS` | Con `values: null` | Il campo data è vuoto (startDate o endDate è nullo) |
| `NOT` | Con `values: null` | Il campo data ha un valore (entrambe le date non sono nulle) |

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Risposte di Errore

### Formato Data Non Valido
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
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
      "code": "NOT_FOUND"
    }
  }]
}
```


## Limitazioni

- Nessun supporto per date ricorrenti (usa automazioni per eventi ricorrenti)
- Impossibile impostare l'ora senza data
- Nessun calcolo integrato dei giorni lavorativi
- Gli intervalli di date non convalidano automaticamente fine > inizio
- La massima precisione è al secondo (nessuna memorizzazione dei millisecondi)

## Risorse Correlate

- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali sui campi personalizzati
- [API Automazioni](/api/automations/index) - Crea automazioni basate su date