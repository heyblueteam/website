---
title: Wyszukiwanie pola niestandardowego
description: Tworzenie pól wyszukiwania, które automatycznie pobierają dane z odniesionych rekordów
---

Pola niestandardowe wyszukiwania automatycznie pobierają dane z rekordów odniesionych przez [Pola odniesienia](/api/custom-fields/reference), wyświetlając informacje z powiązanych rekordów bez ręcznego kopiowania. Aktualizują się automatycznie, gdy zmieniają się dane odniesione.

## Podstawowy przykład

Utwórz pole wyszukiwania, aby wyświetlić tagi z odniesionych rekordów:

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Zaawansowany przykład

Utwórz pole wyszukiwania, aby wyodrębnić wartości pól niestandardowych z odniesionych rekordów:

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola wyszukiwania |
| `type` | CustomFieldType! | ✅ Tak | Musi być `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Tak | Konfiguracja wyszukiwania |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

## Konfiguracja wyszukiwania

### CustomFieldLookupOptionInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `referenceId` | String! | ✅ Tak | ID pola odniesienia, z którego pobierane są dane |
| `lookupId` | String | Nie | ID konkretnego pola niestandardowego do wyszukiwania (wymagane dla typu TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Tak | Typ danych do wyodrębnienia z odniesionych rekordów |

## Typy wyszukiwania

### Wartości CustomFieldLookupType

| Typ | Opis | Zwraca |
|-----|------|--------|
| `TODO_DUE_DATE` | Daty wykonania z odniesionych zadań | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Daty utworzenia z odniesionych zadań | Array of creation timestamps |
| `TODO_UPDATED_AT` | Daty ostatniej aktualizacji z odniesionych zadań | Array of update timestamps |
| `TODO_TAG` | Tagi z odniesionych zadań | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Osoby przypisane z odniesionych zadań | Array of user objects |
| `TODO_DESCRIPTION` | Opisy z odniesionych zadań | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Nazwy list zadań z odniesionych zadań | Array of list titles |
| `TODO_CUSTOM_FIELD` | Wartości pól niestandardowych z odniesionych zadań | Array of values based on the field type |

## Pola odpowiedzi

### Odpowiedź CustomField (dla pól wyszukiwania)

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator pola |
| `name` | String! | Nazwa wyświetlana pola wyszukiwania |
| `type` | CustomFieldType! | Będzie `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Konfiguracja wyszukiwania i wyniki |
| `createdAt` | DateTime! | Kiedy pole zostało utworzone |
| `updatedAt` | DateTime! | Kiedy pole zostało ostatnio zaktualizowane |

### Struktura CustomFieldLookupOption

| Pole | Typ | Opis |
|------|-----|------|
| `lookupType` | CustomFieldLookupType! | Typ wyszukiwania, które jest wykonywane |
| `lookupResult` | JSON | Wyodrębnione dane z odniesionych rekordów |
| `reference` | CustomField | Pole odniesienia używane jako źródło |
| `lookup` | CustomField | Konkretne pole, które jest wyszukiwane (dla TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Rodzicielskie pole wyszukiwania |
| `parentLookup` | CustomField | Rodzic w łańcuchu wyszukiwania (dla zagnieżdżonych wyszukiwań) |

## Jak działają wyszukiwania

1. **Wyodrębnianie danych**: Wyszukiwania wyodrębniają konkretne dane ze wszystkich rekordów powiązanych przez pole odniesienia
2. **Automatyczne aktualizacje**: Gdy zmieniają się rekordy odniesione, wartości wyszukiwania aktualizują się automatycznie
3. **Tylko do odczytu**: Pola wyszukiwania nie mogą być edytowane bezpośrednio - zawsze odzwierciedlają aktualne dane odniesione
4. **Brak obliczeń**: Wyszukiwania wyodrębniają i wyświetlają dane w takiej formie, w jakiej są, bez agregacji lub obliczeń

## Wyszukiwania TODO_CUSTOM_FIELD

Podczas używania typu `TODO_CUSTOM_FIELD` musisz określić, które pole niestandardowe wyodrębnić, używając parametru `lookupId`:

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

To wyodrębnia wartości określonego pola niestandardowego ze wszystkich odniesionych rekordów.

## Zapytania o dane wyszukiwania

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## Przykładowe wyniki wyszukiwania

### Wynik wyszukiwania tagu
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### Wynik wyszukiwania osoby przypisanej
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### Wynik wyszukiwania pola niestandardowego
Wyniki różnią się w zależności od typu pola niestandardowego, które jest wyszukiwane. Na przykład, wyszukiwanie pola walutowego może zwrócić:
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Ważne**: Użytkownicy muszą mieć uprawnienia do wyświetlania zarówno w bieżącym projekcie, jak i w projekcie odniesionym, aby zobaczyć wyniki wyszukiwania.

## Odpowiedzi błędów

### Nieprawidłowe pole odniesienia
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

### Wykryto cykliczne wyszukiwanie
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Brak ID wyszukiwania dla TODO_CUSTOM_FIELD
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## Najlepsze praktyki

1. **Jasne nazewnictwo**: Używaj opisowych nazw, które wskazują, jakie dane są wyszukiwane
2. **Odpowiednie typy**: Wybierz typ wyszukiwania, który odpowiada Twoim potrzebom danych
3. **Wydajność**: Wyszukiwania przetwarzają wszystkie odniesione rekordy, więc bądź ostrożny w przypadku pól odniesienia z wieloma linkami
4. **Uprawnienia**: Upewnij się, że użytkownicy mają dostęp do odniesionych projektów, aby wyszukiwania działały

## Typowe przypadki użycia

### Widoczność między projektami
Wyświetlaj tagi, osoby przypisane lub statusy z powiązanych projektów bez ręcznej synchronizacji.

### Śledzenie zależności
Pokaż daty wykonania lub status ukończenia zadań, od których zależy bieżąca praca.

### Przegląd zasobów
Wyświetl wszystkich członków zespołu przypisanych do odniesionych zadań w celu planowania zasobów.

### Agregacja statusów
Zbieraj wszystkie unikalne statusy z powiązanych zadań, aby zobaczyć zdrowie projektu na pierwszy rzut oka.

## Ograniczenia

- Pola wyszukiwania są tylko do odczytu i nie mogą być edytowane bezpośrednio
- Brak funkcji agregacji (SUMA, LICZBA, ŚREDNIA) - wyszukiwania tylko wyodrębniają dane
- Brak opcji filtrowania - wszystkie odniesione rekordy są uwzględnione
- Cykliczne łańcuchy wyszukiwania są zapobiegane, aby uniknąć nieskończonych pętli
- Wyniki odzwierciedlają aktualne dane i aktualizują się automatycznie

## Powiązane zasoby

- [Pola odniesienia](/api/custom-fields/reference) - Tworzenie linków do rekordów jako źródła wyszukiwania
- [Wartości pól niestandardowych](/api/custom-fields/custom-field-values) - Ustawianie wartości w edytowalnych polach niestandardowych
- [Lista pól niestandardowych](/api/custom-fields/list-custom-fields) - Zapytanie o wszystkie pola niestandardowe w projekcie