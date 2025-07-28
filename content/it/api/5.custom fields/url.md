---
title: Campo Personalizzato URL
description: Crea campi URL per memorizzare indirizzi e link di siti web
---

I campi personalizzati URL ti consentono di memorizzare indirizzi e link di siti web nei tuoi record. Sono ideali per tenere traccia dei siti web dei progetti, dei link di riferimento, degli URL della documentazione o di qualsiasi risorsa web correlata al tuo lavoro.

## Esempio Base

Crea un semplice campo URL:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo URL con descrizione:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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
| `name` | String! | ✅ Sì | Nome visualizzato del campo URL |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `URL` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota:** L'`projectId` viene passato come argomento separato alla mutazione, non come parte dell'oggetto di input.

## Impostazione dei Valori URL

Per impostare o aggiornare un valore URL su un record:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### Parametri SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato URL |
| `text` | String! | ✅ Sì | Indirizzo URL da memorizzare |

## Creazione di Record con Valori URL

Quando crei un nuovo record con valori URL:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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

### TodoCustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `text` | String | L'indirizzo URL memorizzato |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato l'ultimo valore |

## Validazione URL

### Implementazione Corrente
- **API Diretta**: Attualmente non viene applicata alcuna validazione del formato URL
- **Moduli**: La validazione degli URL è pianificata ma non attualmente attiva
- **Memorizzazione**: Qualsiasi valore stringa può essere memorizzato nei campi URL

### Validazione Pianificata
Le versioni future includeranno:
- Validazione del protocollo HTTP/HTTPS
- Verifica del formato URL valido
- Validazione del nome di dominio
- Aggiunta automatica del prefisso del protocollo

### Formati URL Raccomandati
Sebbene non siano attualmente applicati, utilizza questi formati standard:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Note Importanti

### Formato di Memorizzazione
- Gli URL sono memorizzati come testo semplice senza modifiche
- Nessuna aggiunta automatica del protocollo (http://, https://)
- La sensibilità alle maiuscole è preservata come inserita
- Nessuna codifica/decodifica URL eseguita

### API Diretta vs Moduli
- **Moduli**: Validazione URL pianificata (non attualmente attiva)
- **API Diretta**: Nessuna validazione - qualsiasi testo può essere memorizzato
- **Raccomandazione**: Valida gli URL nella tua applicazione prima di memorizzarli

### Campi URL vs Testo
- **URL**: Inteso semanticamente per indirizzi web
- **TEXT_SINGLE**: Testo generale su una sola riga
- **Backend**: Attualmente identica memorizzazione e validazione
- **Frontend**: Componenti UI diversi per l'inserimento dei dati

## Permessi Richiesti

Le operazioni sui campi personalizzati utilizzano permessi basati sui ruoli:

| Azione | Ruolo Richiesto |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Nota:** I permessi vengono verificati in base ai ruoli degli utenti nel progetto, non a costanti di permesso specifiche.

## Risposte di Errore

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

## Migliori Pratiche

### Standard di Formato URL
- Includi sempre il protocollo (http:// o https://)
- Utilizza HTTPS quando possibile per la sicurezza
- Testa gli URL prima di memorizzarli per assicurarti che siano accessibili
- Considera di utilizzare URL accorciati per scopi di visualizzazione

### Qualità dei Dati
- Valida gli URL nella tua applicazione prima di memorizzarli
- Controlla errori di battitura comuni (protocollo mancante, domini errati)
- Standardizza i formati URL nella tua organizzazione
- Considera l'accessibilità e la disponibilità degli URL

### Considerazioni di Sicurezza
- Fai attenzione agli URL forniti dagli utenti
- Valida i domini se limiti a siti specifici
- Considera la scansione degli URL per contenuti dannosi
- Utilizza URL HTTPS quando gestisci dati sensibili

## Filtraggio e Ricerca

### Ricerca Contiene
I campi URL supportano la ricerca di sottostringhe:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Capacità di Ricerca
- Corrispondenza di sottostringhe senza distinzione tra maiuscole e minuscole
- Corrispondenza parziale del dominio
- Ricerca di percorsi e parametri
- Nessun filtraggio specifico per protocollo

## Casi d'Uso Comuni

1. **Gestione Progetti**
   - Siti web dei progetti
   - Link alla documentazione
   - URL dei repository
   - Siti di dimostrazione

2. **Gestione dei Contenuti**
   - Materiali di riferimento
   - Link sorgente
   - Risorse multimediali
   - Articoli esterni

3. **Supporto Clienti**
   - Siti web dei clienti
   - Documentazione di supporto
   - Articoli della base di conoscenza
   - Tutorial video

4. **Vendite e Marketing**
   - Siti web aziendali
   - Pagine prodotto
   - Materiali di marketing
   - Profili sui social media

## Funzionalità di Integrazione

### Con Ricerche
- URL di riferimento da altri record
- Trova record per dominio o modello URL
- Visualizza risorse web correlate
- Aggrega link da più fonti

### Con Moduli
- Componenti di input specifici per URL
- Validazione pianificata per un formato URL corretto
- Capacità di anteprima del link (frontend)
- Visualizzazione URL cliccabili

### Con Reporting
- Traccia l'uso e i modelli degli URL
- Monitora link rotti o inaccessibili
- Categorizza per dominio o protocollo
- Esporta elenchi di URL per analisi

## Limitazioni

### Limitazioni Correnti
- Nessuna validazione attiva del formato URL
- Nessuna aggiunta automatica del protocollo
- Nessuna verifica dei link o controllo di accessibilità
- Nessun accorciamento o espansione degli URL
- Nessuna generazione di favicon o anteprime

### Restrizioni di Automazione
- Non disponibile come campi di attivazione per automazione
- Non possono essere utilizzati negli aggiornamenti dei campi di automazione
- Possono essere riferiti nelle condizioni di automazione
- Disponibili nei modelli di email e webhook

### Vincoli Generali
- Nessuna funzionalità di anteprima dei link integrata
- Nessun accorciamento automatico degli URL
- Nessun tracciamento dei clic o analisi
- Nessun controllo di scadenza degli URL
- Nessuna scansione di URL dannosi

## Miglioramenti Futuri

### Funzionalità Pianificate
- Validazione del protocollo HTTP/HTTPS
- Modelli di validazione regex personalizzati
- Aggiunta automatica del prefisso del protocollo
- Controllo di accessibilità degli URL

### Potenziali Miglioramenti
- Generazione di anteprime dei link
- Visualizzazione di favicon
- Integrazione per l'accorciamento degli URL
- Capacità di tracciamento dei clic
- Rilevamento di link rotti

## Risorse Correlate

- [Campi di Testo](/api/custom-fields/text-single) - Per dati di testo non URL
- [Campi Email](/api/custom-fields/email) - Per indirizzi email
- [Panoramica dei Campi Personalizzati](/api/custom-fields/2.list-custom-fields) - Concetti generali

## Migrazione da Campi di Testo

Se stai migrando da campi di testo a campi URL:

1. **Crea il campo URL** con lo stesso nome e configurazione
2. **Esporta i valori di testo esistenti** per verificare che siano URL validi
3. **Aggiorna i record** per utilizzare il nuovo campo URL
4. **Elimina il vecchio campo di testo** dopo la migrazione riuscita
5. **Aggiorna le applicazioni** per utilizzare componenti UI specifici per URL

### Esempio di Migrazione
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```