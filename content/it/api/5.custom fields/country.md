---
title: Campo Personalizzato per il Paese
description: Crea campi di selezione del paese con validazione del codice ISO del paese
---

I campi personalizzati per il paese ti consentono di memorizzare e gestire informazioni sui paesi per i record. Il campo supporta sia i nomi dei paesi che i codici ISO Alpha-2 dei paesi.

**Importante**: Il comportamento di validazione e conversione del paese differisce significativamente tra le mutazioni:
- **createTodo**: Valida e converte automaticamente i nomi dei paesi in codici ISO
- **setTodoCustomField**: Accetta qualsiasi valore senza validazione

## Esempio di Base

Crea un semplice campo per il paese:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo per il paese con descrizione:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo paese |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `COUNTRY` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: Il `projectId` non viene passato nell'input ma è determinato dal contesto GraphQL (tipicamente dagli header della richiesta o dall'autenticazione).

## Impostazione dei Valori del Paese

I campi per il paese memorizzano i dati in due campi del database:
- **`countryCodes`**: Memorizza i codici ISO Alpha-2 dei paesi come stringa separata da virgole nel database (restituita come array tramite API)
- **`text`**: Memorizza il testo visualizzato o i nomi dei paesi come stringa

### Comprendere i Parametri

La mutazione `setTodoCustomField` accetta due parametri opzionali per i campi paese:

| Parametro | Tipo | Richiesto | Descrizione | Cosa fa |
|-----------|------|----------|-------------|--------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare | - |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato per il paese | - |
| `countryCodes` | [String!] | No | Array di codici ISO Alpha-2 dei paesi | Stored in the `countryCodes` field |
| `text` | String | No | Testo visualizzato o nomi dei paesi | Stored in the `text` field |

**Importante**: 
- In `setTodoCustomField`: Entrambi i parametri sono opzionali e memorizzati in modo indipendente
- In `createTodo`: Il sistema imposta automaticamente entrambi i campi in base al tuo input (non puoi controllarli in modo indipendente)

### Opzione 1: Utilizzare Solo Codici Paese

Memorizza codici ISO validati senza testo visualizzato:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Risultato: `countryCodes` = `["US"]`, `text` = `null`

### Opzione 2: Utilizzare Solo Testo

Memorizza testo visualizzato senza codici validati:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Risultato: `countryCodes` = `null`, `text` = `"United States"`

**Nota**: Quando si utilizza `setTodoCustomField`, non viene effettuata alcuna validazione indipendentemente dal parametro utilizzato. I valori vengono memorizzati esattamente come forniti.

### Opzione 3: Utilizzare Entrambi (Consigliato)

Memorizza sia codici validati che testo visualizzato:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Risultato: `countryCodes` = `["US"]`, `text` = `"United States"`

### Paesi Multipli

Memorizza più paesi utilizzando array:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Creazione di Record con Valori di Paese

Quando si creano record, la mutazione `createTodo` **valida e converte automaticamente** i valori dei paesi. Questa è l'unica mutazione che esegue la validazione del paese:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      text
      countryCodes
    }
  }
}
```

### Formati di Input Accettati

| Tipo di Input | Esempio | Risultato |
|---------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Non supportato** - trattato come valore singolo non valido |
| Mixed format | `"United States, CA"` | **Non supportato** - trattato come valore singolo non valido |

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `text` | String | Testo visualizzato (nomi dei paesi) |
| `countryCodes` | [String!] | Array di codici ISO Alpha-2 dei paesi |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

## Standard dei Paesi

Blue utilizza lo standard **ISO 3166-1 Alpha-2** per i codici dei paesi:

- Codici dei paesi di due lettere (ad es., US, GB, FR, DE)
- La validazione utilizzando la libreria `i18n-iso-countries` **si verifica solo in createTodo**
- Supporta tutti i paesi ufficialmente riconosciuti

### Esempi di Codici Paese

| Paese | Codice ISO |
|-------|------------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Per l'elenco ufficiale completo dei codici ISO 3166-1 alpha-2 dei paesi, visita la [Piattaforma di Navigazione Online ISO](https://www.iso.org/obp/ui/#search/code/).

## Validazione

**La validazione si verifica solo nella mutazione `createTodo`**:

1. **Codice ISO Valido**: Accetta qualsiasi codice ISO Alpha-2 valido
2. **Nome del Paese**: Converte automaticamente i nomi dei paesi riconosciuti in codici
3. **Input Non Valido**: Genera `CustomFieldValueParseError` per valori non riconosciuti

**Nota**: La mutazione `setTodoCustomField` non esegue ALCUNA validazione e accetta qualsiasi valore di stringa.

### Esempio di Errore

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Funzionalità di Integrazione

### Campi di Ricerca
I campi paese possono essere referenziati da campi personalizzati di RICERCA, consentendo di estrarre dati sui paesi da record correlati.

### Automazioni
Utilizza i valori dei paesi nelle condizioni di automazione:
- Filtra le azioni per paesi specifici
- Invia notifiche in base al paese
- Instrada i compiti in base alle regioni geografiche

### Moduli
I campi paese nei moduli validano automaticamente l'input degli utenti e convertono i nomi dei paesi in codici.

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Risposte di Errore

### Valore del Paese Non Valido
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Incompatibilità del Tipo di Campo
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Migliori Pratiche

### Gestione degli Input
- Utilizza `createTodo` per la validazione e conversione automatica
- Utilizza `setTodoCustomField` con cautela poiché salta la validazione
- Considera di validare gli input nella tua applicazione prima di utilizzare `setTodoCustomField`
- Mostra i nomi completi dei paesi nell'interfaccia utente per chiarezza

### Qualità dei Dati
- Valida gli input dei paesi al punto di ingresso
- Utilizza formati coerenti nel tuo sistema
- Considera i raggruppamenti regionali per la reportistica

### Paesi Multipli
- Utilizza il supporto per array in `setTodoCustomField` per più paesi
- Più paesi in `createTodo` **non sono supportati** tramite il campo valore
- Memorizza i codici dei paesi come array in `setTodoCustomField` per una corretta gestione

## Casi d'Uso Comuni

1. **Gestione Clienti**
   - Posizione della sede del cliente
   - Destinazioni di spedizione
   - Giurisdizioni fiscali

2. **Monitoraggio Progetti**
   - Posizione del progetto
   - Posizioni dei membri del team
   - Obiettivi di mercato

3. **Conformità e Legale**
   - Giurisdizioni normative
   - Requisiti di residenza dei dati
   - Controlli sulle esportazioni

4. **Vendite e Marketing**
   - Assegnazioni territoriali
   - Segmentazione di mercato
   - Targeting delle campagne

## Limitazioni

- Supporta solo codici ISO 3166-1 Alpha-2 (codici di 2 lettere)
- Nessun supporto integrato per le suddivisioni dei paesi (stati/province)
- Nessuna icona della bandiera del paese automatica (solo basata su testo)
- Non è possibile validare codici storici dei paesi
- Nessun raggruppamento integrato per regione o continente
- **La validazione funziona solo in `createTodo`, non in `setTodoCustomField`**
- **I paesi multipli non sono supportati nel campo valore di `createTodo`**
- **I codici dei paesi sono memorizzati come stringa separata da virgole, non come vero array**

## Risorse Correlate

- [Panoramica dei Campi Personalizzati](/custom-fields/list-custom-fields) - Concetti generali sui campi personalizzati
- [Campi di Ricerca](/api/custom-fields/lookup) - Riferisci i dati sui paesi da altri record
- [API Moduli](/api/forms) - Includi campi paese in moduli personalizzati