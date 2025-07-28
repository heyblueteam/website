---
title: Pole niestandardowe pliku
description: Twórz pola plików, aby dołączać dokumenty, obrazy i inne pliki do rekordów
---

Pola niestandardowe plików pozwalają na dołączanie wielu plików do rekordów. Pliki są przechowywane w bezpieczny sposób w AWS S3 z kompleksowym śledzeniem metadanych, walidacją typu pliku i odpowiednimi kontrolami dostępu.

## Podstawowy przykład

Utwórz proste pole pliku:

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

## Zaawansowany przykład

Utwórz pole pliku z opisem:

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

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola pliku |
| `type` | CustomFieldType! | ✅ Tak | Musi być `FILE` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

**Uwaga**: Pola niestandardowe są automatycznie powiązane z projektem na podstawie aktualnego kontekstu projektu użytkownika. Żaden parametr `projectId` nie jest wymagany.

## Proces przesyłania plików

### Krok 1: Prześlij plik

Najpierw prześlij plik, aby uzyskać UID pliku:

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

### Krok 2: Dołącz plik do rekordu

Następnie dołącz przesłany plik do rekordu:

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

## Zarządzanie załącznikami plików

### Dodawanie pojedynczych plików

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

### Usuwanie plików

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Operacje zbiorowe na plikach

Zaktualizuj wiele plików jednocześnie, używając customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Parametry wejściowe przesyłania plików

### UploadFileInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `file` | Upload! | ✅ Tak | Plik do przesłania |
| `companyId` | String! | ✅ Tak | ID firmy do przechowywania plików |
| `projectId` | String | Nie | ID projektu dla plików specyficznych dla projektu |

### Parametry wejściowe zarządzania plikami

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego pliku |
| `fileUid` | String! | ✅ Tak | Unikalny identyfikator przesłanego pliku |

## Przechowywanie plików i limity

### Limity rozmiaru plików

| Typ limitu | Rozmiar |
|------------|---------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Obsługiwane typy plików

#### Obrazy
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Filmy
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Audio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Dokumenty
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Archiwa
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Kod/Tekst
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Architektura przechowywania

- **Przechowywanie**: AWS S3 z zorganizowaną strukturą folderów
- **Format ścieżki**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Bezpieczeństwo**: Podpisane URL-e dla bezpiecznego dostępu
- **Kopia zapasowa**: Automatyczna redundancja S3

## Pola odpowiedzi

### Odpowiedź pliku

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | ID bazy danych |
| `uid` | String! | Unikalny identyfikator pliku |
| `name` | String! | Oryginalna nazwa pliku |
| `size` | Float! | Rozmiar pliku w bajtach |
| `type` | String! | Typ MIME |
| `extension` | String! | Rozszerzenie pliku |
| `status` | FileStatus | OCZEKUJĄCE lub POTWIERDZONE (nullable) |
| `shared` | Boolean! | Czy plik jest udostępniony |
| `createdAt` | DateTime! | Znacznik czasu przesłania |

### Odpowiedź TodoCustomFieldFile

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | ID rekordu połączeniowego |
| `uid` | String! | Unikalny identyfikator |
| `position` | Float! | Kolejność wyświetlania |
| `file` | File! | Powiązany obiekt pliku |
| `todoCustomField` | TodoCustomField! | Rodzicielskie pole niestandardowe |
| `createdAt` | DateTime! | Kiedy plik został dołączony |

## Tworzenie rekordów z plikami

Podczas tworzenia rekordów możesz dołączać pliki, używając ich UID:

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

## Walidacja plików i bezpieczeństwo

### Walidacja przesyłania

- **Sprawdzanie typu MIME**: Walidacja w stosunku do dozwolonych typów
- **Walidacja rozszerzenia pliku**: Zapas dla `application/octet-stream`
- **Limity rozmiaru**: Egzekwowane w czasie przesyłania
- **Sanityzacja nazwy pliku**: Usuwa znaki specjalne

### Kontrola dostępu

- **Uprawnienia do przesyłania**: Wymagana przynależność do projektu/firma
- **Powiązanie pliku**: Role ADMIN, WŁAŚCICIEL, CZŁONEK, KLIENT
- **Dostęp do pliku**: Dziedziczone z uprawnień projektu/firma
- **Bezpieczne URL-e**: Podpisane URL-e z ograniczonym czasem dostępu do pliku

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Odpowiedzi błędów

### Plik za duży
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

### Plik nie znaleziony
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

### Pole nie znalezione
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

## Najlepsze praktyki

### Zarządzanie plikami
- Przesyłaj pliki przed dołączeniem do rekordów
- Używaj opisowych nazw plików
- Organizuj pliki według projektu/celu
- Okresowo sprzątaj nieużywane pliki

### Wydajność
- Przesyłaj pliki w partiach, gdy to możliwe
- Używaj odpowiednich formatów plików dla typu treści
- Kompresuj duże pliki przed przesłaniem
- Rozważ wymagania dotyczące podglądu plików

### Bezpieczeństwo
- Waliduj zawartość plików, a nie tylko rozszerzenia
- Używaj skanowania wirusów dla przesyłanych plików
- Wdrażaj odpowiednie kontrole dostępu
- Monitoruj wzorce przesyłania plików

## Typowe przypadki użycia

1. **Zarządzanie dokumentami**
   - Specyfikacje projektów
   - Umowy i kontrakty
   - Notatki i prezentacje ze spotkań
   - Dokumentacja techniczna

2. **Zarządzanie zasobami**
   - Pliki projektowe i makiety
   - Zasoby marki i logotypy
   - Materiały marketingowe
   - Obrazy produktów

3. **Zgodność i rejestry**
   - Dokumenty prawne
   - Ślady audytowe
   - Certyfikaty i licencje
   - Rejestry finansowe

4. **Współpraca**
   - Wspólne zasoby
   - Dokumenty z kontrolą wersji
   - Opinie i adnotacje
   - Materiały referencyjne

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje, gdy pliki są dodawane/usuwane
- Przetwarzaj pliki na podstawie typu lub metadanych
- Wysyłaj powiadomienia o zmianach w plikach
- Archiwizuj pliki na podstawie warunków

### Z obrazami okładkowymi
- Używaj pól plików jako źródeł obrazów okładkowych
- Automatyczne przetwarzanie obrazów i miniatur
- Dynamiczne aktualizacje okładek, gdy pliki się zmieniają

### Z wyszukiwaniami
- Odwołuj się do plików z innych rekordów
- Agreguj liczby i rozmiary plików
- Znajduj rekordy według metadanych plików
- Krzyżowo odwołuj się do załączników plików

## Ograniczenia

- Maksymalnie 256MB na plik
- Zależne od dostępności S3
- Brak wbudowanego wersjonowania plików
- Brak automatycznej konwersji plików
- Ograniczone możliwości podglądu plików
- Brak edytowania w czasie rzeczywistym

## Powiązane zasoby

- [API przesyłania plików](/api/upload-files) - Punkty końcowe przesyłania plików
- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne koncepcje
- [API automatyzacji](/api/automations) - Automatyzacje oparte na plikach
- [Dokumentacja AWS S3](https://docs.aws.amazon.com/s3/) - Zaplecze przechowywania