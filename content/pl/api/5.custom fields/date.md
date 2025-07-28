---
title: Pole niestandardowe daty
description: Tworzenie pól daty do śledzenia pojedynczych dat lub zakresów dat z obsługą stref czasowych
---

Pola niestandardowe daty pozwalają na przechowywanie pojedynczych dat lub zakresów dat dla rekordów. Obsługują zarządzanie strefami czasowymi, inteligentne formatowanie i mogą być używane do śledzenia terminów, dat wydarzeń lub wszelkich informacji opartych na czasie.

## Podstawowy przykład

Utwórz proste pole daty:

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole daty z terminem i opisem:

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola daty |
| `type` | CustomFieldType! | ✅ Tak | Musi być `DATE` |
| `isDueDate` | Boolean | Nie | Czy to pole reprezentuje termin |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

**Uwaga**: Pola niestandardowe są automatycznie powiązane z projektem na podstawie bieżącego kontekstu projektu użytkownika. Żaden parametr `projectId` nie jest wymagany.

## Ustawianie wartości daty

Pola daty mogą przechowywać pojedynczą datę lub zakres dat:

### Pojedyncza data

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Zakres dat

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Wydarzenie całodniowe

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego daty |
| `startDate` | DateTime | Nie | Data/godzina początkowa w formacie ISO 8601 |
| `endDate` | DateTime | Nie | Data/godzina końcowa w formacie ISO 8601 |
| `timezone` | String | Nie | Identyfikator strefy czasowej (np. "America/New_York") |

**Uwaga**: Jeśli tylko `startDate` jest podane, `endDate` automatycznie domyślnie przyjmuje tę samą wartość.

## Format dat

### Format ISO 8601
Wszystkie daty muszą być podawane w formacie ISO 8601:
- `2025-01-15T14:30:00Z` - czas UTC
- `2025-01-15T14:30:00+05:00` - z przesunięciem strefy czasowej
- `2025-01-15T14:30:00.123Z` - z milisekundami

### Identyfikatory stref czasowych
Używaj standardowych identyfikatorów stref czasowych:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Jeśli nie podano strefy czasowej, system domyślnie przyjmuje wykrytą strefę czasową użytkownika.

## Tworzenie rekordów z wartościami daty

Podczas tworzenia nowego rekordu z wartościami daty:

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### Obsługiwane formaty wejściowe

Podczas tworzenia rekordów daty mogą być podawane w różnych formatach:

| Format | Przykład | Wynik |
|--------|----------|-------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | Unikalny identyfikator dla wartości pola |
| `uid` | String! | Unikalny identyfikator w postaci ciągu |
| `customField` | CustomField! | Definicja pola niestandardowego (zawiera wartości dat) |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

**Ważne**: Wartości dat (`startDate`, `endDate`, `timezone`) są dostępne przez pole `customField.value`, a nie bezpośrednio w TodoCustomField.

### Struktura obiektu wartości

Wartości dat są zwracane przez pole `customField.value` jako obiekt JSON:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**Uwaga**: Pole `value` jest typu `CustomField`, a nie `TodoCustomField`.

## Zapytania o wartości dat

Podczas zapytań o rekordy z polami niestandardowymi dat, uzyskaj dostęp do wartości dat przez pole `customField.value`:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

Odpowiedź będzie zawierać wartości dat w polu `value`:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## Inteligencja wyświetlania dat

System automatycznie formatuje daty w zależności od zakresu:

| Scenariusz | Format wyświetlania |
|------------|---------------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (bez wyświetlania czasu) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Wykrywanie całodniowe**: Wydarzenia od 00:00 do 23:59 są automatycznie wykrywane jako wydarzenia całodniowe.

## Obsługa stref czasowych

### Przechowywanie
- Wszystkie daty są przechowywane w UTC w bazie danych
- Informacje o strefie czasowej są przechowywane osobno
- Konwersja następuje przy wyświetlaniu

### Najlepsze praktyki
- Zawsze podawaj strefę czasową dla dokładności
- Używaj spójnych stref czasowych w ramach projektu
- Weź pod uwagę lokalizacje użytkowników w zespołach globalnych

### Powszechne strefy czasowe

| Region | ID strefy czasowej | Przesunięcie UTC |
|--------|--------------------|------------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filtrowanie i zapytania

Pola dat wspierają złożone filtrowanie:

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### Sprawdzanie pustych pól dat

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### Obsługiwane operatory

| Operator | Użycie | Opis |
|----------|--------|------|
| `EQ` | Z dateRange | Data pokrywa się z określonym zakresem (jakiekolwiek przecięcie) |
| `NE` | Z dateRange | Data nie pokrywa się z zakresem |
| `IS` | Z `values: null` | Pole daty jest puste (startDate lub endDate jest null) |
| `NOT` | Z `values: null` | Pole daty ma wartość (oba daty nie są null) |

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Odpowiedzi błędów

### Nieprawidłowy format daty
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
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
      "code": "NOT_FOUND"
    }
  }]
}
```


## Ograniczenia

- Brak wsparcia dla dat cyklicznych (użyj automatyzacji dla wydarzeń cyklicznych)
- Nie można ustawić czasu bez daty
- Brak wbudowanego obliczania dni roboczych
- Zakresy dat nie walidują automatycznie końca > początku
- Maksymalna precyzja to sekunda (brak przechowywania milisekund)

## Powiązane zasoby

- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne koncepcje pól niestandardowych
- [API automatyzacji](/api/automations/index) - Tworzenie automatyzacji opartych na dacie