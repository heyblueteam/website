---
title: Campo Personalizzato Numero
description: Crea campi numerici per memorizzare valori numerici con vincoli min/max opzionali e formattazione del prefisso
---

I campi personalizzati numerici consentono di memorizzare valori numerici per i record. Supportano vincoli di validazione, precisione decimale e possono essere utilizzati per quantità, punteggi, misurazioni o qualsiasi dato numerico che non richieda una formattazione speciale.

## Esempio Base

Crea un semplice campo numerico:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo numerico con vincoli e prefisso:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo numerico |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `NUMBER` |
| `projectId` | String! | ✅ Sì | ID del progetto in cui creare il campo |
| `min` | Float | No | Vincolo di valore minimo (solo guida UI) |
| `max` | Float | No | Vincolo di valore massimo (solo guida UI) |
| `prefix` | String | No | Prefisso di visualizzazione (es. "#", "~", "$") |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

## Impostazione dei Valori Numerici

I campi numerici memorizzano valori decimali con validazione opzionale:

### Valore Numerico Semplice

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Valore Intero

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### Parametri SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato numerico |
| `number` | Float | No | Valore numerico da memorizzare |

## Vincoli di Valore

### Vincoli Min/Max (Guida UI)

**Importante**: I vincoli min/max sono memorizzati ma NON applicati lato server. Servono come guida UI per le applicazioni frontend.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Validazione Lato Client Richiesta**: Le applicazioni frontend devono implementare la logica di validazione per applicare i vincoli min/max.

### Tipi di Valore Supportati

| Tipo | Esempio | Descrizione |
|------|---------|-------------|
| Integer | `42` | Numeri interi |
| Decimal | `42.5` | Numeri con decimali |
| Negative | `-10` | Valori negativi (se non c'è vincolo min) |
| Zero | `0` | Valore zero |

**Nota**: I vincoli min/max NON sono convalidati lato server. I valori al di fuori dell'intervallo specificato saranno accettati e memorizzati.

## Creazione di Record con Valori Numerici

Quando si crea un nuovo record con valori numerici:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### Formati di Input Supportati

Quando si creano record, utilizzare il parametro `number` (non `value`) nell'array dei campi personalizzati:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `number` | Float | Il valore numerico |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando il valore è stato modificato l'ultima volta |

### Risposta CustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per la definizione del campo |
| `name` | String! | Nome visualizzato del campo |
| `type` | CustomFieldType! | Sempre `NUMBER` |
| `min` | Float | Valore minimo consentito |
| `max` | Float | Valore massimo consentito |
| `prefix` | String | Prefisso di visualizzazione |
| `description` | String | Testo di aiuto |

**Nota**: Se il valore numerico non è impostato, il campo `number` sarà `null`.

## Filtraggio e Querying

I campi numerici supportano un filtraggio numerico completo:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Operatori Supportati

| Operatore | Descrizione | Esempio |
|-----------|-------------|---------|
| `EQ` | Uguale a | `number = 42` |
| `NE` | Diverso da | `number ≠ 42` |
| `GT` | Maggiore di | `number > 42` |
| `GTE` | Maggiore o uguale | `number ≥ 42` |
| `LT` | Minore di | `number < 42` |
| `LTE` | Minore o uguale | `number ≤ 42` |
| `IN` | In array | `number in [1, 2, 3]` |
| `NIN` | Non in array | `number not in [1, 2, 3]` |
| `IS` | È nullo/non nullo | `number is null` |

### Filtraggio per Intervallo

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Formattazione di Visualizzazione

### Con Prefisso

Se è impostato un prefisso, verrà visualizzato:

| Valore | Prefisso | Visualizzazione |
|--------|----------|-----------------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Precisione Decimale

I numeri mantengono la loro precisione decimale:

| Input | Memorizzato | Visualizzato |
|-------|-------------|--------------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|---------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Risposte di Errore

### Formato Numero Non Valido
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Nota**: Gli errori di validazione min/max NON si verificano lato server. La validazione dei vincoli deve essere implementata nella tua applicazione frontend.

### Non è un Numero
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Migliori Pratiche

### Progettazione dei Vincoli
- Imposta valori min/max realistici per la guida UI
- Implementa la validazione lato client per applicare i vincoli
- Usa i vincoli per fornire feedback agli utenti nei moduli
- Considera se i valori negativi sono validi per il tuo caso d'uso

### Precisione del Valore
- Usa la precisione decimale appropriata per le tue esigenze
- Considera l'arrotondamento per scopi di visualizzazione
- Sii coerente con la precisione tra i campi correlati

### Miglioramento della Visualizzazione
- Usa prefissi significativi per il contesto
- Considera le unità nei nomi dei campi (es. "Peso (kg)")
- Fornisci descrizioni chiare per le regole di validazione

## Casi d'Uso Comuni

1. **Sistemi di Punteggio**
   - Valutazioni delle prestazioni
   - Punteggi di qualità
   - Livelli di priorità
   - Valutazioni di soddisfazione del cliente

2. **Misurazioni**
   - Quantità e importi
   - Dimensioni e grandezze
   - Durate (in formato numerico)
   - Capacità e limiti

3. **Metriche Aziendali**
   - Fatturato
   - Tassi di conversione
   - Allocazioni di budget
   - Numeri obiettivo

4. **Dati Tecnici**
   - Numeri di versione
   - Valori di configurazione
   - Metriche di prestazione
   - Impostazioni di soglia

## Caratteristiche di Integrazione

### Con Grafici e Dashboard
- Usa i campi NUMERO nei calcoli dei grafici
- Crea visualizzazioni numeriche
- Monitora le tendenze nel tempo

### Con Automazioni
- Attiva azioni in base a soglie numeriche
- Aggiorna i campi correlati in base ai cambiamenti numerici
- Invia notifiche per valori specifici

### Con Lookups
- Aggrega numeri da record correlati
- Calcola totali e medie
- Trova valori min/max attraverso le relazioni

### Con Grafici
- Crea visualizzazioni numeriche
- Monitora le tendenze nel tempo
- Confronta valori tra i record

## Limitazioni

- **Nessuna validazione lato server** dei vincoli min/max
- **Validazione lato client richiesta** per l'applicazione dei vincoli
- Nessuna formattazione della valuta integrata (usa il tipo CURRENCY invece)
- Nessun simbolo di percentuale automatico (usa il tipo PERCENT invece)
- Nessuna capacità di conversione delle unità
- Precisione decimale limitata dal tipo Decimal del database
- Nessuna valutazione di formule matematiche nel campo stesso

## Risorse Correlate

- [Panoramica dei Campi Personalizzati](/api/custom-fields/1.index) - Concetti generali dei campi personalizzati
- [Campo Personalizzato Valuta](/api/custom-fields/currency) - Per valori monetari
- [Campo Personalizzato Percentuale](/api/custom-fields/percent) - Per valori percentuali
- [API Automazioni](/api/automations/1.index) - Crea automazioni basate su numeri