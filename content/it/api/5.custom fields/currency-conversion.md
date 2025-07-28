---
title: Campo Personalizzato di Conversione Valutaria
description: Crea campi che convertono automaticamente i valori delle valute utilizzando i tassi di cambio in tempo reale
---

I campi personalizzati di conversione valutaria convertono automaticamente i valori da un campo CURRENCY sorgente a diverse valute target utilizzando i tassi di cambio in tempo reale. Questi campi si aggiornano automaticamente ogni volta che il valore della valuta sorgente cambia.

I tassi di conversione sono forniti dall'[API Frankfurter](https://github.com/hakanensari/frankfurter), un servizio open-source che traccia i tassi di cambio di riferimento pubblicati dalla [Banca Centrale Europea](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Questo garantisce conversioni valutarie accurate, affidabili e aggiornate per le tue esigenze aziendali internazionali.

## Esempio Base

Crea un semplice campo di conversione valutaria:

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## Esempio Avanzato

Crea un campo di conversione con una data specifica per i tassi storici:

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## Processo Completo di Configurazione

Impostare un campo di conversione valutaria richiede tre passaggi:

### Passaggio 1: Crea un Campo CURRENCY Sorgente

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### Passaggio 2: Crea il Campo CURRENCY_CONVERSION

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### Passaggio 3: Crea Opzioni di Conversione

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|-----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di conversione |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | No | ID del campo CURRENCY sorgente da cui convertire |
| `conversionDateType` | String | No | Strategia di data per i tassi di cambio (vedi sotto) |
| `conversionDate` | String | No | Stringa di data per la conversione (basata su conversionDateType) |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: I campi personalizzati sono automaticamente associati al progetto in base al contesto del progetto attuale dell'utente. Non è richiesto alcun parametro `projectId`.

### Tipi di Data di Conversione

| Tipo | Descrizione | Parametro conversionDate |
|------|-------------|-------------------------|
| `currentDate` | Utilizza tassi di cambio in tempo reale | Non richiesto |
| `specificDate` | Utilizza tassi da una data fissa | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Utilizza la data da un altro campo | "todoDueDate" or DATE field ID |

## Creazione di Opzioni di Conversione

Le opzioni di conversione definiscono quali coppie di valute possono essere convertite:

### CreateCustomFieldOptionInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|-----------|-------------|
| `customFieldId` | String! | ✅ Sì | ID del campo CURRENCY_CONVERSION |
| `title` | String! | ✅ Sì | Nome visualizzato per questa opzione di conversione |
| `currencyConversionFrom` | String! | ✅ Sì | Codice della valuta sorgente o "Qualsiasi" |
| `currencyConversionTo` | String! | ✅ Sì | Codice della valuta target |

### Utilizzare "Qualsiasi" come Sorgente

Il valore speciale "Qualsiasi" come `currencyConversionFrom` crea un'opzione di fallback:

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

Questa opzione sarà utilizzata quando non viene trovata alcuna corrispondenza specifica per la coppia di valute.

## Come Funziona la Conversione Automatica

1. **Aggiornamento del Valore**: Quando un valore viene impostato nel campo CURRENCY sorgente
2. **Corrispondenza dell'Opzione**: Il sistema trova l'opzione di conversione corrispondente in base alla valuta sorgente
3. **Recupero del Tasso**: Recupera il tasso di cambio dall'API Frankfurter
4. **Calcolo**: Moltiplica l'importo sorgente per il tasso di cambio
5. **Memorizzazione**: Salva il valore convertito con il codice della valuta target

### Flusso di Esempio

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## Conversioni Basate sulla Data

### Utilizzare la Data Corrente

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

Le conversioni si aggiornano con i tassi di cambio correnti ogni volta che il valore sorgente cambia.

### Utilizzare una Data Specifica

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

Utilizza sempre i tassi di cambio dalla data specificata.

### Utilizzare la Data da un Campo

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

Utilizza la data da un altro campo (sia la data di scadenza di un'attività o un campo personalizzato di DATA).

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo di conversione |
| `number` | Float | L'importo convertito |
| `currency` | String | Il codice della valuta target |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato aggiornato l'ultimo valore |

## Fonte del Tasso di Cambio

Blue utilizza l'**API Frankfurter** per i tassi di cambio:
- API open-source ospitata dalla Banca Centrale Europea
- Aggiornamenti quotidiani con tassi di cambio ufficiali
- Supporta tassi storici fino al 1999
- Gratuito e affidabile per uso aziendale

## Gestione degli Errori

### Fallimenti di Conversione

Quando la conversione fallisce (errore API, valuta non valida, ecc.):
- Il valore convertito è impostato su `0`
- La valuta target è comunque memorizzata
- Non viene generato alcun errore per l'utente

### Scenari Comuni

| Scenario | Risultato |
|----------|-----------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Nessuna opzione corrispondente | Uses "Any" option if available |
| Missing source value | Nessuna conversione eseguita |

## Permessi Richiesti

La gestione dei campi personalizzati richiede accesso a livello di progetto:

| Ruolo | Può Creare/Aggiornare Campi |
|-------|------------------------------|
| `OWNER` | ✅ Sì |
| `ADMIN` | ✅ Sì |
| `MEMBER` | ❌ No |
| `CLIENT` | ❌ No |

Le autorizzazioni di visualizzazione per i valori convertiti seguono le regole standard di accesso ai record.

## Migliori Pratiche

### Configurazione delle Opzioni
- Crea coppie di valute specifiche per conversioni comuni
- Aggiungi un'opzione di fallback "Qualsiasi" per flessibilità
- Usa titoli descrittivi per le opzioni

### Selezione della Strategia di Data
- Usa `currentDate` per il monitoraggio finanziario in tempo reale
- Usa `specificDate` per reportistica storica
- Usa `fromDateField` per tassi specifici delle transazioni

### Considerazioni sulle Prestazioni
- Più campi di conversione si aggiornano in parallelo
- Le chiamate API vengono effettuate solo quando il valore sorgente cambia
- Le conversioni nella stessa valuta saltano le chiamate API

## Casi d'Uso Comuni

1. **Progetti Multi-Valuta**
   - Traccia i costi del progetto in valute locali
   - Riporta il budget totale nella valuta aziendale
   - Confronta i valori tra le regioni

2. **Vendite Internazionali**
   - Converte i valori degli affari nella valuta di reporting
   - Traccia le entrate in più valute
   - Conversione storica per affari chiusi

3. **Reportistica Finanziaria**
   - Conversioni valutarie a fine periodo
   - Bilanci finanziari consolidati
   - Budget vs. reale in valuta locale

4. **Gestione dei Contratti**
   - Converte i valori dei contratti alla data di firma
   - Traccia i piani di pagamento in più valute
   - Valutazione del rischio valutario

## Limitazioni

- Nessun supporto per conversioni di criptovalute
- Non è possibile impostare valori convertiti manualmente (sempre calcolati)
- Precisione fissa di 2 decimali per tutti gli importi convertiti
- Nessun supporto per tassi di cambio personalizzati
- Nessuna memorizzazione dei tassi di cambio (chiamata API fresca per ogni conversione)
- Dipende dalla disponibilità dell'API Frankfurter

## Risorse Correlate

- [Campi Valuta](/api/custom-fields/currency) - Campi sorgente per le conversioni
- [Campi Data](/api/custom-fields/date) - Per conversioni basate sulla data
- [Campi Formula](/api/custom-fields/formula) - Calcoli alternativi
- [Panoramica dei Campi Personalizzati](/custom-fields/list-custom-fields) - Concetti generali