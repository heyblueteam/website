---
title: Campo Personalizzato Percentuale
description: Crea campi percentuali per memorizzare valori numerici con gestione automatica del simbolo % e formattazione di visualizzazione
---

I campi personalizzati percentuali ti consentono di memorizzare valori percentuali per i record. Gestiscono automaticamente il simbolo % per l'input e la visualizzazione, mentre memorizzano il valore numerico grezzo internamente. Perfetto per tassi di completamento, tassi di successo o qualsiasi metrica basata su percentuali.

## Esempio di Base

Crea un semplice campo percentuale:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo percentuale con descrizione:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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
| `name` | String! | ✅ Sì | Nome visualizzato del campo percentuale |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `PERCENT` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: Il contesto del progetto è determinato automaticamente dalle tue intestazioni di autenticazione. Non è necessario alcun parametro `projectId`.

**Nota**: I campi PERCENT non supportano vincoli min/max o formattazione del prefisso come i campi NUMBER.

## Impostazione dei Valori Percentuali

I campi percentuali memorizzano valori numerici con gestione automatica del simbolo %:

### Con Simbolo Percentuale

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Valore Numerico Diretto

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### Parametri SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato percentuale |
| `number` | Float | No | Valore percentuale numerico (ad es., 75.5 per 75.5%) |

## Memorizzazione e Visualizzazione dei Valori

### Formato di Memorizzazione
- **Memorizzazione interna**: Valore numerico grezzo (ad es., 75.5)
- **Database**: Memorizzato come `Decimal` nella colonna `number`
- **GraphQL**: Restituito come `Float` tipo

### Formato di Visualizzazione
- **Interfaccia utente**: Le applicazioni client devono aggiungere il simbolo % (ad es., "75.5%")
- **Grafici**: Visualizza con simbolo % quando il tipo di output è PERCENTAGE
- **Risposte API**: Valore numerico grezzo senza simbolo % (ad es., 75.5)

## Creazione di Record con Valori Percentuali

Quando si crea un nuovo record con valori percentuali:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Formati di Input Supportati

| Formato | Esempio | Risultato |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Nota**: Il simbolo % viene automaticamente rimosso dall'input e aggiunto nuovamente durante la visualizzazione.

## Interrogazione dei Valori Percentuali

Quando si interrogano record con campi personalizzati percentuali, accedi al valore tramite il percorso `customField.value.number`:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

La risposta includerà la percentuale come numero grezzo:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | Identificatore univoco per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato (contiene il valore percentuale) |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

**Importante**: I valori percentuali vengono accessi tramite il campo `customField.value.number`. Il simbolo % non è incluso nei valori memorizzati e deve essere aggiunto dalle applicazioni client per la visualizzazione.

## Filtraggio e Interrogazione

I campi percentuali supportano lo stesso filtraggio dei campi NUMBER:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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
|----------|-------------|---------|
| `EQ` | Uguale a | `percentage = 75` |
| `NE` | Diverso da | `percentage ≠ 75` |
| `GT` | Maggiore di | `percentage > 75` |
| `GTE` | Maggiore o uguale | `percentage ≥ 75` |
| `LT` | Minore di | `percentage < 75` |
| `LTE` | Minore o uguale | `percentage ≤ 75` |
| `IN` | Valore nella lista | `percentage in [50, 75, 100]` |
| `NIN` | Valore non nella lista | `percentage not in [0, 25]` |
| `IS` | Controlla per null con `values: null` | `percentage is null` |
| `NOT` | Controlla per non null con `values: null` | `percentage is not null` |

### Filtraggio per Intervallo

Per il filtraggio per intervallo, utilizza più operatori:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Intervalli di Valore Percentuale

### Intervalli Comuni

| Intervallo | Descrizione | Caso d'uso |
|-------|-------------|----------|
| `0-100` | Percentuale standard | Completion rates, success rates |
| `0-∞` | Percentuale illimitata | Growth rates, performance metrics |
| `-∞-∞` | Qualsiasi valore | Change rates, variance |

### Valori di Esempio

| Input | Memorizzato | Visualizzato |
|-------|--------|---------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Aggregazione dei Grafici

I campi percentuali supportano l'aggregazione nei grafici e nei report del dashboard. Le funzioni disponibili includono:

- `AVERAGE` - Valore percentuale medio
- `COUNT` - Numero di record con valori
- `MIN` - Valore percentuale più basso
- `MAX` - Valore percentuale più alto 
- `SUM` - Totale di tutti i valori percentuali

Queste aggregazioni sono disponibili durante la creazione di grafici e dashboard, non in query GraphQL dirette.

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Risposte di Errore

### Formato Percentuale Non Valido
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

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

### Inserimento Valori
- Consenti agli utenti di inserire con o senza simbolo %
- Valida intervalli ragionevoli per il tuo caso d'uso
- Fornisci un contesto chiaro su cosa rappresenta il 100%

### Visualizzazione
- Mostra sempre il simbolo % nelle interfacce utente
- Usa una precisione decimale appropriata
- Considera la codifica a colori per gli intervalli (rosso/giallo/verde)

### Interpretazione dei Dati
- Documenta cosa significa il 100% nel tuo contesto
- Gestisci i valori superiori al 100% in modo appropriato
- Considera se i valori negativi sono validi

## Casi d'Uso Comuni

1. **Gestione Progetti**
   - Tassi di completamento delle attività
   - Progresso del progetto
   - Utilizzo delle risorse
   - Velocità dello sprint

2. **Monitoraggio delle Prestazioni**
   - Tassi di successo
   - Tassi di errore
   - Metriche di efficienza
   - Punteggi di qualità

3. **Metriche Finanziarie**
   - Tassi di crescita
   - Margini di profitto
   - Importi degli sconti
   - Percentuali di cambiamento

4. **Analisi**
   - Tassi di conversione
   - Tassi di clic
   - Metriche di coinvolgimento
   - Indicatori di prestazione

## Funzionalità di Integrazione

### Con Formule
- Riferisci i campi PERCENT nei calcoli
- Formattazione automatica del simbolo % negli output delle formule
- Combina con altri campi numerici

### Con Automazioni
- Attiva azioni basate su soglie percentuali
- Invia notifiche per percentuali di traguardo
- Aggiorna lo stato in base ai tassi di completamento

### Con Ricerche
- Aggrega percentuali da record correlati
- Calcola tassi di successo medi
- Trova gli elementi con le migliori/peggiori prestazioni

### Con Grafici
- Crea visualizzazioni basate su percentuali
- Monitora i progressi nel tempo
- Confronta metriche di prestazione

## Differenze dai Campi NUMBER

### Cosa è Diverso
- **Gestione dell'input**: Rimuove automaticamente il simbolo %
- **Visualizzazione**: Aggiunge automaticamente il simbolo %
- **Vincoli**: Nessuna validazione min/max
- **Formattazione**: Nessun supporto per il prefisso

### Cosa è Uguale
- **Memorizzazione**: Stessa colonna e tipo di database
- **Filtraggio**: Stessi operatori di query
- **Aggregazione**: Stesse funzioni di aggregazione
- **Permessi**: Stesso modello di permessi

## Limitazioni

- Nessun vincolo di valore min/max
- Nessuna opzione di formattazione del prefisso
- Nessuna validazione automatica dell'intervallo 0-100%
- Nessuna conversione tra formati percentuali (ad es., 0.75 ↔ 75%)
- Valori superiori al 100% sono consentiti

## Risorse Correlate

- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali sui campi personalizzati
- [Campo Personalizzato Numero](/api/custom-fields/number) - Per valori numerici grezzi
- [API Automazioni](/api/automations/index) - Crea automazioni basate su percentuali