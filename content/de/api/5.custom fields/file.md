---
title: Datei benutzerdefiniertes Feld
description: Erstellen Sie Dateifelder, um Dokumente, Bilder und andere Dateien an Datensätze anzuhängen
---

Datei benutzerdefinierte Felder ermöglichen es Ihnen, mehrere Dateien an Datensätze anzuhängen. Dateien werden sicher in AWS S3 mit umfassender Metadatenverfolgung, Dateitypvalidierung und ordnungsgemäßen Zugriffskontrollen gespeichert.

## Einfaches Beispiel

Erstellen Sie ein einfaches Dateifeld:

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

## Fortgeschrittenes Beispiel

Erstellen Sie ein Dateifeld mit Beschreibung:

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

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Dateifelds |
| `type` | CustomFieldType! | ✅ Ja | Muss `FILE` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Benutzerdefinierte Felder sind automatisch mit dem Projekt basierend auf dem aktuellen Projektkontext des Benutzers verknüpft. Kein `projectId` Parameter ist erforderlich.

## Datei-Upload-Prozess

### Schritt 1: Datei hochladen

Laden Sie zuerst die Datei hoch, um eine Datei-UID zu erhalten:

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

### Schritt 2: Datei an Datensatz anhängen

Hängen Sie dann die hochgeladene Datei an einen Datensatz an:

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

## Verwaltung von Dateianhängen

### Hinzufügen einzelner Dateien

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

### Entfernen von Dateien

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Massen-Dateioperationen

Aktualisieren Sie mehrere Dateien gleichzeitig mit customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Datei-Upload-Eingabeparameter

### UploadFileInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `file` | Upload! | ✅ Ja | Datei zum Hochladen |
| `companyId` | String! | ✅ Ja | Unternehmens-ID für die Dateispeicherung |
| `projectId` | String | Nein | Projekt-ID für projektspezifische Dateien |

### Datei-Verwaltungs-Eingabeparameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Datei benutzerdefinierten Feldes |
| `fileUid` | String! | ✅ Ja | Eindeutiger Identifikator der hochgeladenen Datei |

## Dateispeicherung und -limits

### Dateigrößenlimits

| Limittyp | Größe |
|----------|------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Unterstützte Dateitypen

#### Bilder
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Videos
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Audio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Dokumente
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Archive
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Code/Text
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Speicherarchitektur

- **Speicher**: AWS S3 mit organisierter Ordnerstruktur
- **Pfadformat**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Sicherheit**: Signierte URLs für sicheren Zugriff
- **Backup**: Automatische S3-Redundanz

## Antwortfelder

### Datei-Antwort

| Feld | Typ | Beschreibung |
|------|------|-------------|
| `id` | ID! | Datenbank-ID |
| `uid` | String! | Eindeutiger Dateiidentifikator |
| `name` | String! | Ursprünglicher Dateiname |
| `size` | Float! | Dateigröße in Bytes |
| `type` | String! | MIME-Typ |
| `extension` | String! | Dateierweiterung |
| `status` | FileStatus | PENDING oder CONFIRMED (nullable) |
| `shared` | Boolean! | Ob die Datei geteilt ist |
| `createdAt` | DateTime! | Upload-Zeitstempel |

### TodoCustomFieldFile Antwort

| Feld | Typ | Beschreibung |
|------|------|-------------|
| `id` | ID! | Verknüpfungsdatensatz-ID |
| `uid` | String! | Eindeutiger Identifikator |
| `position` | Float! | Anzeigeordnung |
| `file` | File! | Zugehöriges Dateiobjekt |
| `todoCustomField` | TodoCustomField! | Übergeordnetes benutzerdefiniertes Feld |
| `createdAt` | DateTime! | Wann die Datei angehängt wurde |

## Erstellen von Datensätzen mit Dateien

Beim Erstellen von Datensätzen können Sie Dateien mit ihren UIDs anhängen:

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

## Dateiüberprüfung und Sicherheit

### Upload-Validierung

- **MIME-Typ-Prüfung**: Validiert gegen erlaubte Typen
- **Dateierweiterungsvalidierung**: Fallback für `application/octet-stream`
- **Größenlimits**: Bei Upload-Zeit durchgesetzt
- **Dateinamen-Säuberung**: Entfernt Sonderzeichen

### Zugriffskontrolle

- **Upload-Berechtigungen**: Projekt-/Unternehmensmitgliedschaft erforderlich
- **Dateiverknüpfung**: ADMIN, OWNER, MEMBER, CLIENT Rollen
- **Dateizugriff**: Erbt von Projekt-/Unternehmensberechtigungen
- **Sichere URLs**: Zeitlich begrenzte signierte URLs für den Dateizugriff

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|---------------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Fehlermeldungen

### Datei zu groß
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

### Datei nicht gefunden
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

### Feld nicht gefunden
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

## Beste Praktiken

### Datei-Verwaltung
- Dateien vor dem Anhängen an Datensätze hochladen
- Beschreibende Dateinamen verwenden
- Dateien nach Projekt/Zweck organisieren
- Unbenutzte Dateien regelmäßig bereinigen

### Leistung
- Dateien nach Möglichkeit in Chargen hochladen
- Geeignete Dateiformate für den Inhaltstyp verwenden
- Große Dateien vor dem Upload komprimieren
- Anforderungen an die Dateivorschau berücksichtigen

### Sicherheit
- Datei-Inhalte validieren, nicht nur Erweiterungen
- Virenscanning für hochgeladene Dateien verwenden
- Angemessene Zugriffskontrollen implementieren
- Datei-Upload-Muster überwachen

## Häufige Anwendungsfälle

1. **Dokumentenverwaltung**
   - Projektspezifikationen
   - Verträge und Vereinbarungen
   - Besprechungsnotizen und Präsentationen
   - Technische Dokumentation

2. **Asset-Management**
   - Entwurfsdateien und Mockups
   - Markenassets und Logos
   - Marketingmaterialien
   - Produktbilder

3. **Compliance und Aufzeichnungen**
   - Rechtliche Dokumente
   - Prüfprotokolle
   - Zertifikate und Lizenzen
   - Finanzunterlagen

4. **Zusammenarbeit**
   - Gemeinsame Ressourcen
   - Versionskontrollierte Dokumente
   - Feedback und Anmerkungen
   - Referenzmaterialien

## Integrationsfunktionen

### Mit Automatisierungen
- Aktionen auslösen, wenn Dateien hinzugefügt/entfernt werden
- Dateien basierend auf Typ oder Metadaten verarbeiten
- Benachrichtigungen für Dateiänderungen senden
- Dateien basierend auf Bedingungen archivieren

### Mit Coverbildern
- Dateifelder als Quellen für Coverbilder verwenden
- Automatische Bildverarbeitung und Thumbnails
- Dynamische Cover-Updates, wenn Dateien sich ändern

### Mit Nachschlägen
- Dateien aus anderen Datensätzen referenzieren
- Dateizahlen und -größen aggregieren
- Datensätze nach Dateimetadaten finden
- Dateianhänge kreuzreferenzieren

## Einschränkungen

- Maximale 256 MB pro Datei
- Abhängig von der Verfügbarkeit von S3
- Keine integrierte Dateiversionierung
- Keine automatische Datei-Konvertierung
- Eingeschränkte Dateivorschau-Funktionen
- Keine Echtzeit-Kollaboration

## Verwandte Ressourcen

- [Dateien hochladen API](/api/upload-files) - Endpunkte zum Datei-Upload
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte
- [Automatisierungen API](/api/automations) - Datei-basierte Automatisierungen
- [AWS S3 Dokumentation](https://docs.aws.amazon.com/s3/) - Speicher-Backend