---
title: Campo Personalizzato di Valutazione
description: Crea campi di valutazione per memorizzare valutazioni numeriche con scale e convalide configurabili
---

I campi personalizzati di valutazione consentono di memorizzare valutazioni numeriche nei record con valori minimi e massimi configurabili. Sono ideali per valutazioni delle prestazioni, punteggi di soddisfazione, livelli di priorità o qualsiasi dato basato su scale numeriche nei tuoi progetti.

## Esempio di Base

Crea un semplice campo di valutazione con scala predefinita 0-5:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Esempio Avanzato

Crea un campo di valutazione con scala e descrizione personalizzate:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di valutazione |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `RATING` |
| `projectId` | String! | ✅ Sì | L'ID del progetto in cui verrà creato questo campo |
| `description` | String | No | Testo di aiuto mostrato agli utenti |
| `min` | Float | No | Valore minimo di valutazione (nessun valore predefinito) |
| `max` | Float | No | Valore massimo di valutazione |

## Impostazione dei Valori di Valutazione

Per impostare o aggiornare un valore di valutazione su un record:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput Parametri

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato di valutazione |
| `value` | String! | ✅ Sì | Valore di valutazione come stringa (all'interno dell'intervallo configurato) |

## Creazione di Record con Valori di Valutazione

Quando crei un nuovo record con valori di valutazione:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
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
| `value` | Float | Il valore di valutazione memorizzato (accessibile tramite customField.value) |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato per l'ultima volta il valore |

**Nota**: Il valore di valutazione è effettivamente accessibile tramite `customField.value.number` nelle query.

### CustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il campo |
| `name` | String! | Nome visualizzato del campo di valutazione |
| `type` | CustomFieldType! | Sempre `RATING` |
| `min` | Float | Valore minimo di valutazione consentito |
| `max` | Float | Valore massimo di valutazione consentito |
| `description` | String | Testo di aiuto per il campo |

## Convalida della Valutazione

### Vincoli di Valore
- I valori di valutazione devono essere numerici (tipo Float)
- I valori devono essere all'interno dell'intervallo min/max configurato
- Se non viene specificato un minimo, non c'è un valore predefinito
- Il valore massimo è facoltativo ma raccomandato

### Regole di Convalida
**Importante**: La convalida avviene solo quando si inviano moduli, non quando si utilizza `setTodoCustomField` direttamente.

- L'input viene analizzato come un numero float (quando si utilizzano moduli)
- Deve essere maggiore o uguale al valore minimo (quando si utilizzano moduli)
- Deve essere minore o uguale al valore massimo (quando si utilizzano moduli)
- `setTodoCustomField` accetta qualsiasi valore stringa senza convalida

### Esempi di Valutazione Valida
Per un campo con min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Esempi di Valutazione Non Valida
Per un campo con min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Opzioni di Configurazione

### Impostazione della Scala di Valutazione
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Scale di Valutazione Comuni
- **1-5 Stelle**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 Prestazione**: `min: 1, max: 10`
- **0-100 Percentuale**: `min: 0, max: 100`
- **Scala Personalizzata**: Qualsiasi intervallo numerico

## Permessi Richiesti

Le operazioni sui campi personalizzati seguono i permessi standard basati sui ruoli:

| Azione | Ruolo Richiesto |
|--------|------------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Nota**: I ruoli specifici richiesti dipendono dalla configurazione dei ruoli personalizzati del tuo progetto e dai permessi a livello di campo.

## Risposte di Errore

### Errore di Convalida (Solo Moduli)
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

**Importante**: La convalida del valore di valutazione (vincoli min/max) avviene solo quando si inviano moduli, non quando si utilizza `setTodoCustomField` direttamente.

### Campo Personalizzato Non Trovato
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

## Migliori Pratiche

### Progettazione della Scala
- Utilizza scale di valutazione coerenti tra campi simili
- Considera la familiarità dell'utente (1-5 stelle, 0-10 NPS)
- Imposta valori minimi appropriati (0 vs 1)
- Definisci un significato chiaro per ciascun livello di valutazione

### Qualità dei Dati
- Convalida i valori di valutazione prima di memorizzarli
- Utilizza la precisione decimale in modo appropriato
- Considera l'arrotondamento per scopi di visualizzazione
- Fornisci indicazioni chiare sui significati delle valutazioni

### Esperienza Utente
- Mostra le scale di valutazione visivamente (stelle, barre di progresso)
- Mostra il valore attuale e i limiti della scala
- Fornisci contesto per i significati delle valutazioni
- Considera valori predefiniti per nuovi record

## Casi d'Uso Comuni

1. **Gestione delle Prestazioni**
   - Valutazioni delle prestazioni dei dipendenti
   - Punteggi di qualità del progetto
   - Valutazioni di completamento delle attività
   - Valutazioni del livello di competenza

2. **Feedback dei Clienti**
   - Valutazioni di soddisfazione
   - Punteggi di qualità del prodotto
   - Valutazioni dell'esperienza di servizio
   - Net Promoter Score (NPS)

3. **Priorità e Importanza**
   - Livelli di priorità delle attività
   - Valutazioni di urgenza
   - Punteggi di valutazione del rischio
   - Valutazioni di impatto

4. **Assicurazione della Qualità**
   - Valutazioni delle revisioni del codice
   - Punteggi di qualità dei test
   - Qualità della documentazione
   - Valutazioni di aderenza ai processi

## Funzionalità di Integrazione

### Con Automazioni
- Attiva azioni basate su soglie di valutazione
- Invia notifiche per valutazioni basse
- Crea attività di follow-up per valutazioni alte
- Smista il lavoro in base ai valori di valutazione

### Con Ricerche
- Calcola le valutazioni medie tra i record
- Trova record per intervalli di valutazione
- Riferisci dati di valutazione da altri record
- Aggrega statistiche di valutazione

### Con il Frontend di Blue
- Convalida automatica degli intervalli nei contesti dei moduli
- Controlli di input di valutazione visivi
- Feedback di convalida in tempo reale
- Opzioni di input a stelle o cursore

## Tracciamento delle Attività

Le modifiche ai campi di valutazione vengono tracciate automaticamente:
- I vecchi e i nuovi valori di valutazione vengono registrati
- L'attività mostra cambiamenti numerici
- Timestamp per tutti gli aggiornamenti di valutazione
- Attribuzione utente per le modifiche

## Limitazioni

- Sono supportati solo valori numerici
- Nessuna visualizzazione di valutazione integrata (stelle, ecc.)
- La precisione decimale dipende dalla configurazione del database
- Nessuna memorizzazione di metadati di valutazione (commenti, contesto)
- Nessuna aggregazione automatica delle valutazioni o statistiche
- Nessuna conversione di valutazione integrata tra scale
- **Critico**: La convalida min/max funziona solo nei moduli, non tramite `setTodoCustomField`

## Risorse Correlate

- [Campi Numerici](/api/5.custom%20fields/number) - Per dati numerici generali
- [Campi Percentuali](/api/5.custom%20fields/percent) - Per valori percentuali
- [Campi di Selezione](/api/5.custom%20fields/select-single) - Per valutazioni a scelta discreta
- [Panoramica dei Campi Personalizzati](/api/5.custom%20fields/2.list-custom-fields) - Concetti generali