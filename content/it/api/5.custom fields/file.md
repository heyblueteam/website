---
title: Campo Personalizzato File
description: Crea campi file per allegare documenti, immagini e altri file ai record
---

I campi personalizzati file ti consentono di allegare più file ai record. I file sono archiviati in modo sicuro in AWS S3 con un tracciamento completo dei metadati, validazione del tipo di file e controlli di accesso appropriati.

## Esempio di Base

Crea un semplice campo file:

```graphql
mutation CreateFileField {
  createCustomField(input: {
    name: "Attachments"
    type: FILE
  }) {
    id
    name
    type
  }
}
```

## Esempio Avanzato

Crea un campo file con descrizione:

```graphql
mutation CreateDetailedFileField {
  createCustomField(input: {
    name: "Project Documents"
    type: FILE
    description: "Upload project-related documents, images, and files"
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
| `name` | String! | ✅ Sì | Nome visualizzato del campo file |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `FILE` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

**Nota**: I campi personalizzati sono automaticamente associati al progetto in base al contesto del progetto attuale dell'utente. Non è richiesto alcun parametro `projectId`.

## Processo di Caricamento File

### Passo 1: Carica File

Prima, carica il file per ottenere un UID file:

```graphql
mutation UploadFile {
  uploadFile(input: {
    file: $file  # File upload variable
    companyId: "company_123"
    projectId: "proj_123"
  }) {
    id
    uid
    name
    size
    type
    extension
    status
  }
}
```

### Passo 2: Allegare File al Record

Poi allega il file caricato a un record:

```graphql
mutation AttachFileToRecord {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "file_field_456"
    fileUid: "file_uid_from_upload"
  }) {
    id
    file {
      uid
      name
      size
      type
    }
  }
}
```

## Gestione degli Allegati File

### Aggiunta di File Singoli

```graphql
mutation AddFileToField {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  }) {
    id
    position
    file {
      uid
      name
      size
      type
      extension
    }
  }
}
```

### Rimozione di File

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Operazioni di File in Massa

Aggiorna più file contemporaneamente utilizzando customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Parametri di Input per il Caricamento File

### UploadFileInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `file` | Upload! | ✅ Sì | File da caricare |
| `companyId` | String! | ✅ Sì | ID aziendale per l'archiviazione dei file |
| `projectId` | String | No | ID progetto per file specifici del progetto |

### Parametri di Input per la Gestione dei File

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record |
| `customFieldId` | String! | ✅ Sì | ID del campo file personalizzato |
| `fileUid` | String! | ✅ Sì | Identificatore unico del file caricato |

## Archiviazione e Limiti dei File

### Limiti di Dimensione dei File

| Tipo di Limite | Dimensione |
|----------------|------------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Tipi di File Supportati

#### Immagini
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Video
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Audio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Documenti
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Archivi
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Codice/Testo
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Architettura di Archiviazione

- **Archiviazione**: AWS S3 con struttura di cartelle organizzata
- **Formato del Percorso**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Sicurezza**: URL firmati per accesso sicuro
- **Backup**: Ridondanza automatica S3

## Campi di Risposta

### Risposta File

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | ID del database |
| `uid` | String! | Identificatore unico del file |
| `name` | String! | Nome originale del file |
| `size` | Float! | Dimensione del file in byte |
| `type` | String! | Tipo MIME |
| `extension` | String! | Estensione del file |
| `status` | FileStatus | PENDING o CONFIRMED (nullable) |
| `shared` | Boolean! | Se il file è condiviso |
| `createdAt` | DateTime! | Timestamp di caricamento |

### Risposta TodoCustomFieldFile

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | ID! | ID del record di giunzione |
| `uid` | String! | Identificatore unico |
| `position` | Float! | Ordine di visualizzazione |
| `file` | File! | Oggetto file associato |
| `todoCustomField` | TodoCustomField! | Campo personalizzato padre |
| `createdAt` | DateTime! | Quando è stato allegato il file |

## Creazione di Record con File

Quando crei record, puoi allegare file utilizzando i loro UID:

```graphql
mutation CreateRecordWithFiles {
  createTodo(input: {
    title: "Project deliverables"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "file_field_id"
      customFieldOptionIds: ["file_uid_1", "file_uid_2"]
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
      todoCustomFieldFiles {
        id
        position
        file {
          uid
          name
          size
          type
        }
      }
    }
  }
}
```

## Validazione e Sicurezza dei File

### Validazione del Caricamento

- **Controllo del Tipo MIME**: Valida rispetto ai tipi consentiti
- **Validazione dell'Estensione del File**: Fallback per `application/octet-stream`
- **Limiti di Dimensione**: Applicati al momento del caricamento
- **Sanitizzazione del Nome del File**: Rimuove caratteri speciali

### Controllo degli Accessi

- **Permessi di Caricamento**: Richiesta di appartenenza al progetto/azienda
- **Associazione File**: Ruoli ADMIN, OWNER, MEMBER, CLIENT
- **Accesso ai File**: Ereditato dai permessi del progetto/azienda
- **URL Sicuri**: URL firmati a tempo limitato per l'accesso ai file

## Permessi Richiesti

| Azione | Permesso Richiesto |
|--------|-------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Risposte di Errore

### File Troppo Grande
```json
{
  "errors": [{
    "message": "File \"filename.pdf\": Size exceeds maximum limit of 256MB",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### File Non Trovato
```json
{
  "errors": [{
    "message": "File not found",
    "extensions": {
      "code": "FILE_NOT_FOUND"
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

### Gestione dei File
- Carica i file prima di allegarli ai record
- Usa nomi di file descrittivi
- Organizza i file per progetto/scopo
- Pulisci periodicamente i file non utilizzati

### Prestazioni
- Carica file in batch quando possibile
- Usa formati di file appropriati per il tipo di contenuto
- Comprimi file di grandi dimensioni prima del caricamento
- Considera i requisiti di anteprima dei file

### Sicurezza
- Valida i contenuti del file, non solo le estensioni
- Usa la scansione antivirus per i file caricati
- Implementa controlli di accesso appropriati
- Monitora i modelli di caricamento dei file

## Casi d'Uso Comuni

1. **Gestione Documenti**
   - Specifiche di progetto
   - Contratti e accordi
   - Appunti e presentazioni delle riunioni
   - Documentazione tecnica

2. **Gestione delle Risorse**
   - File di design e mockup
   - Risorse e loghi del marchio
   - Materiale di marketing
   - Immagini dei prodotti

3. **Conformità e Registrazioni**
   - Documenti legali
   - Tracce di audit
   - Certificati e licenze
   - Registrazioni finanziarie

4. **Collaborazione**
   - Risorse condivise
   - Documenti con controllo delle versioni
   - Feedback e annotazioni
   - Materiali di riferimento

## Funzionalità di Integrazione

### Con Automazioni
- Attiva azioni quando i file vengono aggiunti/rimossi
- Elabora file in base al tipo o ai metadati
- Invia notifiche per le modifiche ai file
- Archivia file in base a condizioni

### Con Immagini di Copertura
- Usa campi file come fonti di immagini di copertura
- Elaborazione automatica delle immagini e miniature
- Aggiornamenti dinamici della copertura quando i file cambiano

### Con Ricerche
- Riferisci file da altri record
- Aggrega conteggi e dimensioni dei file
- Trova record per metadati dei file
- Riferimento incrociato degli allegati file

## Limitazioni

- Massimo 256MB per file
- Dipendente dalla disponibilità di S3
- Nessuna versioning file integrato
- Nessuna conversione automatica dei file
- Capacità limitate di anteprima dei file
- Nessuna modifica collaborativa in tempo reale

## Risorse Correlate

- [API Caricamento File](/api/upload-files) - Endpoint di caricamento file
- [Panoramica Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali
- [API Automazioni](/api/automations) - Automazioni basate su file
- [Documentazione AWS S3](https://docs.aws.amazon.com/s3/) - Backend di archiviazione