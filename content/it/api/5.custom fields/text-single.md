---
title: Campo Personalizzato di Testo su Una Linea
description: Crea campi di testo su una linea per valori di testo brevi come nomi, titoli e etichette
---

I campi personalizzati di testo su una linea ti consentono di memorizzare valori di testo brevi destinati all'input su una sola linea. Sono ideali per nomi, titoli, etichette o qualsiasi dato testuale che dovrebbe essere visualizzato su una sola linea.

## Esempio Base

Crea un semplice campo di testo su una linea:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo di testo su una linea con descrizione:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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

| Parametro | Tipo | Obbligatorio | Descrizione |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di testo |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `TEXT_SINGLE` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: Il contesto del progetto è determinato automaticamente dalle intestazioni di autenticazione. Non è necessario alcun parametro `projectId`.

## Impostazione dei Valori di Testo

Per impostare o aggiornare un valore di testo su una linea in un record:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### Parametri di SetTodoCustomFieldInput

| Parametro | Tipo | Obbligatorio | Descrizione |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo di testo personalizzato |
| `text` | String | No | Contenuto di testo su una linea da memorizzare |

## Creazione di Record con Valori di Testo

Quando crei un nuovo record con valori di testo su una linea:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato (contiene il valore di testo) |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando il valore è stato modificato per l'ultima volta |

**Importante**: I valori di testo sono accessibili attraverso il campo `customField.value.text`, non direttamente su TodoCustomField.

## Interrogazione dei Valori di Testo

Quando interroghi record con campi di testo personalizzati, accedi al testo attraverso il percorso `customField.value.text`:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

La risposta includerà il testo nella struttura nidificata:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Validazione del Testo

### Validazione del Modulo
Quando i campi di testo su una linea sono utilizzati nei moduli:
- Gli spazi bianchi iniziali e finali vengono automaticamente rimossi
- La validazione obbligatoria viene applicata se il campo è contrassegnato come obbligatorio
- Non viene applicata alcuna validazione di formato specifico

### Regole di Validazione
- Accetta qualsiasi contenuto di stringa, inclusi i ritorni a capo (anche se non raccomandato)
- Nessun limite di lunghezza dei caratteri (fino ai limiti del database)
- Supporta caratteri Unicode e simboli speciali
- I ritorni a capo sono preservati ma non sono intesi per questo tipo di campo

### Esempi di Testo Tipici
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Note Importanti

### Capacità di Archiviazione
- Memorizzato utilizzando il tipo MySQL `MediumText`
- Supporta fino a 16MB di contenuto testuale
- Archiviazione identica ai campi di testo su più linee
- Codifica UTF-8 per caratteri internazionali

### API Diretta vs Moduli
- **Moduli**: Rimozione automatica degli spazi bianchi e validazione obbligatoria
- **API Diretta**: Il testo è memorizzato esattamente come fornito
- **Raccomandazione**: Utilizza i moduli per l'input degli utenti per garantire un formato coerente

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: Input di testo su una linea, ideale per valori brevi
- **TEXT_MULTI**: Input di textarea su più linee, ideale per contenuti più lunghi
- **Backend**: Entrambi utilizzano archiviazione e validazione identiche
- **Frontend**: Componenti UI diversi per l'inserimento dei dati
- **Intento**: TEXT_SINGLE è semanticamente destinato a valori su una linea

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## Risposte di Errore

### Validazione del Campo Obbligatorio (Solo Moduli)
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

## Migliori Pratiche

### Linee Guida per il Contenuto
- Mantieni il testo conciso e appropriato per una sola linea
- Evita i ritorni a capo per una visualizzazione su una sola linea
- Usa un formato coerente per tipi di dati simili
- Considera i limiti di caratteri in base ai requisiti della tua UI

### Inserimento Dati
- Fornisci descrizioni chiare dei campi per guidare gli utenti
- Usa i moduli per l'input degli utenti per garantire la validazione
- Valida il formato del contenuto nella tua applicazione se necessario
- Considera l'uso di menu a discesa per valori standardizzati

### Considerazioni sulle Prestazioni
- I campi di testo su una linea sono leggeri e performanti
- Considera l'indicizzazione per campi frequentemente cercati
- Usa larghezze di visualizzazione appropriate nella tua UI
- Monitora la lunghezza del contenuto per scopi di visualizzazione

## Filtraggio e Ricerca

### Ricerca per Contenuto
I campi di testo su una linea supportano la ricerca di sottostringhe:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Capacità di Ricerca
- Corrispondenza di sottostringhe senza distinzione tra maiuscole e minuscole
- Supporta la corrispondenza di parole parziali
- Corrispondenza esatta dei valori
- Nessuna ricerca full-text o ranking

## Casi d'Uso Comuni

1. **Identificatori e Codici**
   - SKU dei prodotti
   - Numeri d'ordine
   - Codici di riferimento
   - Numeri di versione

2. **Nomi e Titoli**
   - Nomi dei clienti
   - Titoli dei progetti
   - Nomi dei prodotti
   - Etichette delle categorie

3. **Brevi Descrizioni**
   - Brevi riassunti
   - Etichette di stato
   - Indicatori di priorità
   - Tag di classificazione

4. **Riferimenti Esterni**
   - Numeri di ticket
   - Riferimenti a fatture
   - ID di sistemi esterni
   - Numeri di documenti

## Caratteristiche di Integrazione

### Con Ricerche
- Riferisci dati testuali da altri record
- Trova record per contenuto testuale
- Visualizza informazioni testuali correlate
- Aggrega valori testuali da più fonti

### Con Moduli
- Rimozione automatica degli spazi bianchi
- Validazione dei campi obbligatori
- UI di input di testo su una linea
- Visualizzazione del limite di caratteri (se configurato)

### Con Importazioni/Esportazioni
- Mappatura diretta delle colonne CSV
- Assegnazione automatica dei valori di testo
- Supporto per l'importazione di dati in blocco
- Esportazione in formati di foglio di calcolo

## Limitazioni

### Restrizioni di Automazione
- Non direttamente disponibili come campi di attivazione per l'automazione
- Non possono essere utilizzati negli aggiornamenti dei campi di automazione
- Possono essere referenziati nelle condizioni di automazione
- Disponibili nei modelli di email e webhook

### Limitazioni Generali
- Nessuna formattazione o stile di testo incorporato
- Nessuna validazione automatica oltre ai campi obbligatori
- Nessuna enforcement di unicità incorporata
- Nessuna compressione del contenuto per testi molto grandi
- Nessuna versioning o tracciamento delle modifiche
- Capacità di ricerca limitate (nessuna ricerca full-text)

## Risorse Correlate

- [Campi di Testo su Più Linee](/api/custom-fields/text-multi) - Per contenuti testuali più lunghi
- [Campi Email](/api/custom-fields/email) - Per indirizzi email
- [Campi URL](/api/custom-fields/url) - Per indirizzi web
- [Campi ID Unici](/api/custom-fields/unique-id) - Per identificatori generati automaticamente
- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali