---
title: Campo Personalizzato di Testo Multi-Linea
description: Crea campi di testo multi-linea per contenuti più lunghi come descrizioni, note e commenti
---

I campi personalizzati di testo multi-linea consentono di memorizzare contenuti testuali più lunghi con interruzioni di riga e formattazione. Sono ideali per descrizioni, note, commenti o qualsiasi dato testuale che necessiti di più righe.

## Esempio Base

Crea un semplice campo di testo multi-linea:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo di testo multi-linea con descrizione:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `name` | String! | ✅ Sì | Nome visualizzato del campo di testo |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `TEXT_MULTI` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota:** Il `projectId` viene passato come argomento separato alla mutazione, non come parte dell'oggetto di input. In alternativa, il contesto del progetto può essere determinato dall'intestazione `X-Bloo-Project-ID` nella tua richiesta GraphQL.

## Impostazione dei Valori di Testo

Per impostare o aggiornare un valore di testo multi-linea su un record:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### Parametri SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo di testo personalizzato |
| `text` | String | No | Contenuto di testo multi-linea da memorizzare |

## Creazione di Record con Valori di Testo

Quando si crea un nuovo record con valori di testo multi-linea:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
    }
  }
}
```

## Campi di Risposta

### Risposta TodoCustomField

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `text` | String | Il contenuto di testo multi-linea memorizzato |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

## Validazione del Testo

### Validazione del Modulo
Quando i campi di testo multi-linea vengono utilizzati nei moduli:
- Gli spazi bianchi all'inizio e alla fine vengono automaticamente rimossi
- La validazione richiesta viene applicata se il campo è contrassegnato come richiesto
- Non viene applicata alcuna validazione di formato specifico

### Regole di Validazione
- Accetta qualsiasi contenuto di stringa, comprese le interruzioni di riga
- Nessun limite di lunghezza dei caratteri (fino ai limiti del database)
- Supporta caratteri Unicode e simboli speciali
- Le interruzioni di riga vengono preservate in memorizzazione

### Esempi di Testo Valido
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Note Importanti

### Capacità di Memorizzazione
- Memorizzato utilizzando il tipo MySQL `MediumText`
- Supporta fino a 16MB di contenuto testuale
- Le interruzioni di riga e la formattazione vengono preservate
- Codifica UTF-8 per caratteri internazionali

### API Diretta vs Moduli
- **Moduli**: Rimozione automatica degli spazi bianchi e validazione richiesta
- **API Diretta**: Il testo viene memorizzato esattamente come fornito
- **Raccomandazione**: Utilizzare i moduli per l'input degli utenti per garantire una formattazione coerente

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI**: Input di textarea multi-linea, ideale per contenuti più lunghi
- **TEXT_SINGLE**: Input di testo su una sola riga, ideale per valori brevi
- **Backend**: Entrambi i tipi sono identici - stesso campo di memorizzazione, validazione e elaborazione
- **Frontend**: Componenti UI diversi per l'immissione dei dati (textarea vs campo di input)
- **Importante**: La distinzione tra TEXT_MULTI e TEXT_SINGLE esiste puramente per scopi UI

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

## Risposte di Errore

### Validazione del Campo Richiesto (Solo Moduli)
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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Migliori Pratiche

### Organizzazione dei Contenuti
- Utilizzare una formattazione coerente per contenuti strutturati
- Considerare l'uso di una sintassi simile a markdown per la leggibilità
- Suddividere contenuti lunghi in sezioni logiche
- Utilizzare interruzioni di riga per migliorare la leggibilità

### Immissione Dati
- Fornire descrizioni chiare dei campi per guidare gli utenti
- Utilizzare moduli per l'input degli utenti per garantire la validazione
- Considerare i limiti di caratteri in base al proprio caso d'uso
- Validare il formato del contenuto nella propria applicazione se necessario

### Considerazioni sulle Prestazioni
- Contenuti testuali molto lunghi possono influenzare le prestazioni delle query
- Considerare la paginazione per visualizzare grandi campi di testo
- Considerazioni sugli indici per la funzionalità di ricerca
- Monitorare l'uso della memoria per campi con contenuti grandi

## Filtraggio e Ricerca

### Ricerca Contiene
I campi di testo multi-linea supportano la ricerca di sottostringhe tramite filtri di campo personalizzati:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Capacità di Ricerca
- Corrispondenza di sottostringhe all'interno dei campi di testo utilizzando l'operatore `CONTAINS`
- Ricerca senza distinzione tra maiuscole e minuscole utilizzando l'operatore `NCONTAINS`
- Corrispondenza esatta utilizzando l'operatore `IS`
- Corrispondenza negativa utilizzando l'operatore `NOT`
- Ricerche su tutte le righe di testo
- Supporta la corrispondenza di parole parziali

## Casi d'Uso Comuni

1. **Gestione Progetti**
   - Descrizioni delle attività
   - Requisiti del progetto
   - Note delle riunioni
   - Aggiornamenti sullo stato

2. **Assistenza Clienti**
   - Descrizioni dei problemi
   - Note di risoluzione
   - Feedback dei clienti
   - Registri di comunicazione

3. **Gestione dei Contenuti**
   - Contenuto degli articoli
   - Descrizioni dei prodotti
   - Commenti degli utenti
   - Dettagli delle recensioni

4. **Documentazione**
   - Descrizioni dei processi
   - Istruzioni
   - Linee guida
   - Materiali di riferimento

## Funzionalità di Integrazione

### Con Automazioni
- Attivare azioni quando il contenuto testuale cambia
- Estrarre parole chiave dal contenuto testuale
- Creare riassunti o notifiche
- Elaborare contenuti testuali con servizi esterni

### Con Ricerche
- Riferire dati testuali da altri record
- Aggregare contenuti testuali da più fonti
- Trovare record per contenuto testuale
- Visualizzare informazioni testuali correlate

### Con Moduli
- Rimozione automatica degli spazi bianchi
- Validazione dei campi richiesti
- UI di textarea multi-linea
- Visualizzazione del conteggio dei caratteri (se configurato)

## Limitazioni

- Nessuna formattazione del testo incorporata o modifica del testo ricco
- Nessuna rilevazione automatica dei link o conversione
- Nessuna correzione ortografica o validazione grammaticale
- Nessuna analisi o elaborazione del testo incorporata
- Nessuna versioning o tracciamento delle modifiche
- Capacità di ricerca limitate (nessuna ricerca full-text)
- Nessuna compressione del contenuto per testi molto grandi

## Risorse Correlate

- [Campi di Testo a Riga Singola](/api/custom-fields/text-single) - Per valori di testo brevi
- [Campi Email](/api/custom-fields/email) - Per indirizzi email
- [Campi URL](/api/custom-fields/url) - Per indirizzi web
- [Panoramica dei Campi Personalizzati](/api/custom-fields/2.list-custom-fields) - Concetti generali