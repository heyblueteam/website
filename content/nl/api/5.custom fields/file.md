---
title: Bestands Aangepast Veld
description: Maak bestandsvelden aan om documenten, afbeeldingen en andere bestanden aan records toe te voegen
---

Bestands aangepaste velden stellen je in staat om meerdere bestanden aan records toe te voegen. Bestanden worden veilig opgeslagen in AWS S3 met uitgebreide metadata-tracking, validatie van bestandstypen en juiste toegangscontroles.

## Basisvoorbeeld

Maak een eenvoudig bestandsveld aan:

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

## Geavanceerd Voorbeeld

Maak een bestandsveld met beschrijving:

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

## Invoervelden

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het bestandsveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `FILE` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

**Opmerking**: Aangepaste velden worden automatisch geassocieerd met het project op basis van de huidige projectcontext van de gebruiker. Geen `projectId` parameter is vereist.

## Bestandsuploadproces

### Stap 1: Upload Bestand

Upload eerst het bestand om een bestand UID te krijgen:

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

### Stap 2: Koppel Bestand aan Record

Koppel vervolgens het geüploade bestand aan een record:

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

## Beheren van Bestandsbijlagen

### Enkele Bestanden Toevoegen

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

### Bestanden Verwijderen

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Bulk Bestandsbewerkingen

Werk meerdere bestanden tegelijk bij met behulp van customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Bestandsupload Invoervelden

### UploadFileInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `file` | Upload! | ✅ Ja | Bestand om te uploaden |
| `companyId` | String! | ✅ Ja | Bedrijfs-ID voor bestandsopslag |
| `projectId` | String | Nee | Project-ID voor projectspecifieke bestanden |

### Bestandsbeheer Invoervelden

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van het record |
| `customFieldId` | String! | ✅ Ja | ID van het bestands aangepaste veld |
| `fileUid` | String! | ✅ Ja | Unieke identificatie van het geüploade bestand |

## Bestandsopslag en Limieten

### Bestandsformaten Limieten

| Limiet Type | Grootte |
|-------------|---------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Ondersteunde Bestandsformaten

#### Afbeeldingen
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Video's
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Audio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Documenten
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Archieven
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Code/Tekst
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Opslagarchitectuur

- **Opslag**: AWS S3 met georganiseerde mappenstructuur
- **Padformaat**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Beveiliging**: Ondertekende URL's voor veilige toegang
- **Back-up**: Automatische S3-redundantie

## Antwoordvelden

### Bestandsantwoord

| Veld | Type | Beschrijving |
|------|------|-------------|
| `id` | ID! | Database-ID |
| `uid` | String! | Unieke bestandsidentificatie |
| `name` | String! | Originele bestandsnaam |
| `size` | Float! | Bestandsomvang in bytes |
| `type` | String! | MIME-type |
| `extension` | String! | Bestandsextensie |
| `status` | FileStatus | PENDING of BEVESTIGD (nullable) |
| `shared` | Boolean! | Of het bestand gedeeld is |
| `createdAt` | DateTime! | Upload-timestamp |

### TodoCustomFieldFile Antwoord

| Veld | Type | Beschrijving |
|------|------|-------------|
| `id` | ID! | Koppelrecord-ID |
| `uid` | String! | Unieke identificatie |
| `position` | Float! | Weergavevolgorde |
| `file` | File! | Geassocieerd bestandsobject |
| `todoCustomField` | TodoCustomField! | Bovenliggend aangepast veld |
| `createdAt` | DateTime! | Wanneer het bestand werd toegevoegd |

## Records Aanmaken met Bestanden

Bij het aanmaken van records kun je bestanden toevoegen met hun UIDs:

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

## Bestandsvalidatie en Beveiliging

### Uploadvalidatie

- **MIME-type Controle**: Valideert tegen toegestane types
- **Bestandsextensie Validatie**: Back-up voor `application/octet-stream`
- **Grootte Limieten**: Afgedwongen bij uploadtijd
- **Bestandsnaam Sanitization**: Verwijdert speciale tekens

### Toegangscontrole

- **Uploadrechten**: Project-/bedrijfsleden vereist
- **Bestand Associatie**: ADMIN, EIGENAAR, LID, KLANT rollen
- **Bestandstoegang**: Geërfd van project-/bedrijfsrechten
- **Veilige URL's**: Tijdslimiet ondertekende URL's voor bestandstoegang

## Vereiste Rechten

| Actie | Vereiste Rechten |
|-------|------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Foutreacties

### Bestand Te Groot
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

### Bestand Niet Gevonden
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

### Veld Niet Gevonden
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

## Best Practices

### Bestandsbeheer
- Upload bestanden voordat je ze aan records toevoegt
- Gebruik beschrijvende bestandsnamen
- Organiseer bestanden per project/doel
- Maak ongebruikte bestanden periodiek schoon

### Prestaties
- Upload bestanden in batches wanneer mogelijk
- Gebruik geschikte bestandsformaten voor inhoudstype
- Comprimeer grote bestanden voor upload
- Houd rekening met bestandsvoorbeeldvereisten

### Beveiliging
- Valideer bestandsinhoud, niet alleen extensies
- Gebruik virus-scanning voor geüploade bestanden
- Implementeer juiste toegangscontroles
- Monitor bestandsuploadpatronen

## Veelvoorkomende Gebruikscases

1. **Documentbeheer**
   - Project specificaties
   - Contracten en overeenkomsten
   - Notulen en presentaties
   - Technische documentatie

2. **Assetbeheer**
   - Ontwerpbestanden en mockups
   - Merkmiddelen en logo's
   - Marketingmaterialen
   - Productafbeeldingen

3. **Compliance en Records**
   - Juridische documenten
   - Auditsporen
   - Certificaten en licenties
   - Financiële records

4. **Samenwerking**
   - Gedeelde middelen
   - Versiebeheerde documenten
   - Feedback en annotaties
   - Referentiematerialen

## Integratiefuncties

### Met Automatiseringen
- Trigger acties wanneer bestanden worden toegevoegd/verwijderd
- Verwerk bestanden op basis van type of metadata
- Stuur meldingen voor bestandswijzigingen
- Archiveer bestanden op basis van voorwaarden

### Met Coverafbeeldingen
- Gebruik bestandsvelden als bronnen voor coverafbeeldingen
- Automatische beeldverwerking en miniaturen
- Dynamische cover-updates wanneer bestanden veranderen

### Met Lookups
- Verwijs naar bestanden van andere records
- Agregeer bestandsaantallen en -groottes
- Vind records op basis van bestandsmetadata
- Kruisverwijs bestandsbijlagen

## Beperkingen

- Maximaal 256MB per bestand
- Afhankelijk van S3-beschikbaarheid
- Geen ingebouwde bestandsversiebeheer
- Geen automatische bestandsconversie
- Beperkte bestandsvoorbeeldmogelijkheden
- Geen realtime samenwerkingsbewerking

## Gerelateerde Bronnen

- [Upload Bestanden API](/api/upload-files) - Bestandsupload eindpunten
- [Aangepaste Velden Overzicht](/api/custom-fields/list-custom-fields) - Algemene concepten
- [Automatiseringen API](/api/automations) - Bestandsgebaseerde automatiseringen
- [AWS S3 Documentatie](https://docs.aws.amazon.com/s3/) - Opslag backend