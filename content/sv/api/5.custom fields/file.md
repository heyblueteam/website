---
title: Filanpassat fält
description: Skapa fält för filer för att bifoga dokument, bilder och andra filer till poster
---

Filanpassade fält gör att du kan bifoga flera filer till poster. Filer lagras säkert i AWS S3 med omfattande metadataövervakning, validering av filtyp och korrekta åtkomstkontroller.

## Grundläggande exempel

Skapa ett enkelt filfält:

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

## Avancerat exempel

Skapa ett filfält med beskrivning:

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

## Indata parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för filfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `FILE` |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: Anpassade fält är automatiskt kopplade till projektet baserat på användarens aktuella projektkontext. Ingen `projectId` parameter krävs.

## Filuppladdningsprocess

### Steg 1: Ladda upp fil

Först, ladda upp filen för att få en fil UID:

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

### Steg 2: Bifoga fil till post

Bifoga sedan den uppladdade filen till en post:

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

## Hantera filbilagor

### Lägga till enskilda filer

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

### Ta bort filer

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Massfiloperationer

Uppdatera flera filer på en gång med hjälp av customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Filuppladdningsindata parametrar

### UploadFileInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `file` | Upload! | ✅ Ja | Fil att ladda upp |
| `companyId` | String! | ✅ Ja | Företags-ID för fillagring |
| `projectId` | String | Nej | Projekt-ID för projektspecifika filer |

### Filhanteringsindata parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten |
| `customFieldId` | String! | ✅ Ja | ID för det filanpassade fältet |
| `fileUid` | String! | ✅ Ja | Unik identifierare för den uppladdade filen |

## Fil lagring och begränsningar

### Filstorleksbegränsningar

| Begränsningstyp | Storlek |
|-----------------|--------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Stödda filtyper

#### Bilder
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Videor
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Ljud
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Dokument
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Arkiv
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Kod/Text
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Lagringsarkitektur

- **Lagring**: AWS S3 med organiserad mappstruktur
- **Sökvägsformat**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Säkerhet**: Signerade URL:er för säker åtkomst
- **Backup**: Automatisk S3 redundans

## Svarsfält

### Fil Svar

| Fält | Typ | Beskrivning |
|------|-----|-------------|
| `id` | ID! | Databas-ID |
| `uid` | String! | Unik filidentifierare |
| `name` | String! | Original filnamn |
| `size` | Float! | Filstorlek i byte |
| `type` | String! | MIME-typ |
| `extension` | String! | Filändelse |
| `status` | FileStatus | PENDING eller CONFIRMED (nullable) |
| `shared` | Boolean! | Om filen är delad |
| `createdAt` | DateTime! | Uppladdningstidstämpel |

### TodoCustomFieldFile Svar

| Fält | Typ | Beskrivning |
|------|-----|-------------|
| `id` | ID! | Junction post-ID |
| `uid` | String! | Unik identifierare |
| `position` | Float! | Visningsordning |
| `file` | File! | Associerad filobjekt |
| `todoCustomField` | TodoCustomField! | Föräldranpassat fält |
| `createdAt` | DateTime! | När filen bifogades |

## Skapa poster med filer

När du skapar poster kan du bifoga filer med deras UIDs:

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

## Filvalidering och säkerhet

### Uppladdningsvalidering

- **MIME-typkontroll**: Validerar mot tillåtna typer
- **Filändelsevalidering**: Reserv för `application/octet-stream`
- **Storleksbegränsningar**: Tillämpas vid uppladdning
- **Filnamnssanering**: Tar bort specialtecken

### Åtkomstkontroll

- **Uppladdningsbehörigheter**: Projekt-/företagsmedlemskap krävs
- **Filassociation**: ADMIN, ÄGARE, MEDLEM, KLIENT roller
- **Filåtkomst**: Ärver från projekt-/företagsbehörigheter
- **Säkra URL:er**: Tidsbegränsade signerade URL:er för filåtkomst

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk behörighet |
|--------|------------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Felrespons

### Fil för stor
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

### Fil hittades inte
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

### Fält hittades inte
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

## Bästa praxis

### Filhantering
- Ladda upp filer innan de bifogas till poster
- Använd beskrivande filnamn
- Organisera filer efter projekt/ändamål
- Rensa bort oanvända filer regelbundet

### Prestanda
- Ladda upp filer i batchar när det är möjligt
- Använd lämpliga filformat för innehållstyp
- Komprimera stora filer innan uppladdning
- Tänk på krav för filförhandsvisning

### Säkerhet
- Validera filinnehåll, inte bara ändelser
- Använd viruskontroll för uppladdade filer
- Implementera korrekta åtkomstkontroller
- Övervaka filuppladdningsmönster

## Vanliga användningsfall

1. **Dokumenthantering**
   - Projektspecifikationer
   - Kontrakt och avtal
   - Mötesanteckningar och presentationer
   - Teknisk dokumentation

2. **Tillgångshantering**
   - Designfiler och mockups
   - Varumärkesresurser och logotyper
   - Marknadsföringsmaterial
   - Produktbilder

3. **Regelefterlevnad och poster**
   - Juridiska dokument
   - Revisionsspår
   - Certifikat och licenser
   - Finansiella poster

4. **Samarbete**
   - Delade resurser
   - Versionskontrollerade dokument
   - Feedback och anteckningar
   - Referensmaterial

## Integrationsfunktioner

### Med automatiseringar
- Utlösa åtgärder när filer läggs till/ tas bort
- Bearbeta filer baserat på typ eller metadata
- Skicka meddelanden för filändringar
- Arkivera filer baserat på villkor

### Med omslagsbilder
- Använd filfält som källor för omslagsbilder
- Automatisk bildbehandling och miniatyrer
- Dynamiska omslagsuppdateringar när filer ändras

### Med uppslag
- Referera till filer från andra poster
- Sammanställ filantal och storlekar
- Hitta poster efter filmetadata
- Korsreferens filbilagor

## Begränsningar

- Maximalt 256 MB per fil
- Beroende av S3 tillgänglighet
- Ingen inbyggd filversionshantering
- Ingen automatisk filkonvertering
- Begränsade möjligheter för filförhandsvisning
- Ingen realtids samarbetsredigering

## Relaterade resurser

- [Ladda upp filer API](/api/upload-files) - Filuppladdningsändpunkter
- [Översikt över anpassade fält](/api/custom-fields/list-custom-fields) - Allmänna koncept
- [Automatiserings-API](/api/automations) - Filbaserade automatiseringar
- [AWS S3 Dokumentation](https://docs.aws.amazon.com/s3/) - Lagringsbackend