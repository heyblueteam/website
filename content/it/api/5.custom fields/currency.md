---
title: Campo Personalizzato per Valuta
description: Crea campi di valuta per tracciare valori monetari con formattazione e convalida appropriate
---

I campi personalizzati per la valuta ti consentono di memorizzare e gestire valori monetari con codici di valuta associati. Il campo supporta 72 valute diverse, comprese le principali valute fiat e criptovalute, con formattazione automatica e vincoli min/max opzionali.

## Esempio Base

Crea un semplice campo di valuta:

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Esempio Avanzato

Crea un campo di valuta con vincoli di convalida:

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di valuta |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `CURRENCY` |
| `currency` | String | No | Codice di valuta predefinito (codice ISO a 3 lettere) |
| `min` | Float | No | Valore minimo consentito (memorizzato ma non applicato agli aggiornamenti) |
| `max` | Float | No | Valore massimo consentito (memorizzato ma non applicato agli aggiornamenti) |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: Il contesto del progetto viene determinato automaticamente dalla tua autenticazione. Devi avere accesso al progetto in cui stai creando il campo.

## Impostazione dei Valori di Valuta

Per impostare o aggiornare un valore di valuta su un record:

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### Parametri di SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato per la valuta |
| `number` | Float! | ✅ Sì | L'importo monetario |
| `currency` | String! | ✅ Sì | Codice di valuta a 3 lettere |

## Creazione di Record con Valori di Valuta

Quando crei un nuovo record con valori di valuta:

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### Formato di Input per la Creazione

Quando crei record, i valori di valuta vengono passati in modo diverso:

| Parametro | Tipo | Descrizione |
|-----------|------|-------------|
| `customFieldId` | String! | ID del campo di valuta |
| `value` | String! | Importo come stringa (es. "1500.50") |
| `currency` | String! | Codice di valuta a 3 lettere |

## Valute Supportate

Blue supporta 72 valute, comprese 70 valute fiat e 2 criptovalute:

### Valute Fiat

#### Americhe
| Valuta | Codice | Nome |
|--------|--------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Bolívar Soberano Venezuelano |
| Boliviano Boliviano | `BOB` | Boliviano Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europa
| Valuta | Codice | Nome |
|--------|--------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Corona Norvegese | `NOK` | Corona Norvegese |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Asia-Pacifico
| Valuta | Codice | Nome |
|--------|--------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### Medio Oriente e Africa
| Valuta | Codice | Nome |
|--------|--------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Criptovalute
| Valuta | Codice |
|--------|--------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Campi di Risposta

### TodoCustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `number` | Float | L'importo monetario |
| `currency` | String | Il codice di valuta a 3 lettere |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

## Formattazione della Valuta

Il sistema formatta automaticamente i valori di valuta in base alla localizzazione:

- **Posizionamento del simbolo**: Posiziona correttamente i simboli di valuta (prima/dopo)
- **Separatori decimali**: Utilizza separatori specifici per la localizzazione (. o ,)
- **Separatori delle migliaia**: Applica il raggruppamento appropriato
- **Decimali**: Mostra 0-2 decimali in base all'importo
- **Gestione speciale**: USD/CAD mostrano il prefisso del codice di valuta per chiarezza

### Esempi di Formattazione

| Valore | Valuta | Visualizzazione |
|--------|--------|----------------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Convalida

### Convalida dell'Importo
- Deve essere un numero valido
- I vincoli min/max sono memorizzati con la definizione del campo ma non applicati durante gli aggiornamenti dei valori
- Supporta fino a 2 decimali per la visualizzazione (la precisione completa è memorizzata internamente)

### Convalida del Codice di Valuta
- Deve essere uno dei 72 codici di valuta supportati
- Sensibile al maiuscolo (usa lettere maiuscole)
- Codici non validi restituiscono un errore

## Funzionalità di Integrazione

### Formule
I campi di valuta possono essere utilizzati nei campi personalizzati FORMULA per calcoli:
- Somma di più campi di valuta
- Calcolo delle percentuali
- Esecuzione di operazioni aritmetiche

### Conversione di Valuta
Utilizza i campi CURRENCY_CONVERSION per convertire automaticamente tra valute (vedi [Campi di Conversione di Valuta](/api/custom-fields/currency-conversion))

### Automazioni
I valori di valuta possono attivare automazioni basate su:
- Soglie di importo
- Tipo di valuta
- Cambiamenti di valore

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|---------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Nota**: Anche se qualsiasi membro del progetto può creare campi personalizzati, la possibilità di impostare valori dipende dai permessi basati sui ruoli configurati per ciascun campo.

## Risposte di Errore

### Valore di Valuta Non Valido
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

Questo errore si verifica quando:
- Il codice di valuta non è uno dei 72 codici supportati
- Il formato del numero non è valido
- Il valore non può essere analizzato correttamente

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

### Selezione della Valuta
- Imposta una valuta predefinita che corrisponda al tuo mercato principale
- Utilizza i codici di valuta ISO 4217 in modo coerente
- Considera la posizione dell'utente quando scegli i predefiniti

### Vincoli di Valore
- Imposta valori min/max ragionevoli per prevenire errori di inserimento dati
- Usa 0 come minimo per i campi che non dovrebbero essere negativi
- Considera il tuo caso d'uso quando imposti i massimi

### Progetti Multi-Valuta
- Utilizza una valuta base coerente per la reportistica
- Implementa i campi CURRENCY_CONVERSION per la conversione automatica
- Documenta quale valuta dovrebbe essere utilizzata per ciascun campo

## Casi d'Uso Comuni

1. **Budgeting del Progetto**
   - Monitoraggio del budget del progetto
   - Stime dei costi
   - Monitoraggio delle spese

2. **Vendite e Affari**
   - Valori degli affari
   - Importi dei contratti
   - Monitoraggio dei ricavi

3. **Pianificazione Finanziaria**
   - Importi degli investimenti
   - Round di finanziamento
   - Obiettivi finanziari

4. **Affari Internazionali**
   - Prezzi in multi-valuta
   - Monitoraggio del cambio estero
   - Transazioni transfrontaliere

## Limitazioni

- Massimo 2 decimali per la visualizzazione (anche se viene memorizzata maggiore precisione)
- Nessuna conversione di valuta integrata nei campi CURRENCY standard
- Non è possibile mescolare valute in un singolo valore di campo
- Nessun aggiornamento automatico dei tassi di cambio (usa CURRENCY_CONVERSION per questo)
- I simboli di valuta non sono personalizzabili

## Risorse Correlate

- [Campi di Conversione di Valuta](/api/custom-fields/currency-conversion) - Conversione automatica di valuta
- [Campi Numerici](/api/custom-fields/number) - Per valori numerici non monetari
- [Campi Formula](/api/custom-fields/formula) - Calcola con valori di valuta
- [Campi Personalizzati di Elenco](/api/custom-fields/list-custom-fields) - Interroga tutti i campi personalizzati in un progetto