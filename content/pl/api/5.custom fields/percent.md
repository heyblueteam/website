---
title: Pole niestandardowe procentowe
description: Twórz pola procentowe do przechowywania wartości numerycznych z automatycznym obsługiwaniem symbolu % i formatowaniem wyświetlania
---

Pola niestandardowe procentowe pozwalają na przechowywanie wartości procentowych dla rekordów. Automatycznie obsługują symbol % dla wprowadzania i wyświetlania, jednocześnie przechowując wewnętrzną wartość numeryczną. Idealne do wskaźników ukończenia, wskaźników sukcesu lub wszelkich metryk opartych na procentach.

## Podstawowy przykład

Utwórz proste pole procentowe:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole procentowe z opisem:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola procentowego |
| `type` | CustomFieldType! | ✅ Tak | Musi być `PERCENT` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

**Uwaga**: Kontekst projektu jest automatycznie określany na podstawie nagłówków uwierzytelniania. Nie jest wymagany parametr `projectId`.

**Uwaga**: Pola PERCENT nie obsługują ograniczeń min/max ani formatowania prefiksów jak pola NUMBER.

## Ustawianie wartości procentowych

Pola procentowe przechowują wartości numeryczne z automatycznym obsługiwaniem symbolu %:

### Z symbolem procentowym

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Bezpośrednia wartość numeryczna

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego procentowego |
| `number` | Float | Nie | Wartość procentowa (np. 75.5 dla 75.5%) |

## Przechowywanie i wyświetlanie wartości

### Format przechowywania
- **Przechowywanie wewnętrzne**: Surowa wartość numeryczna (np. 75.5)
- **Baza danych**: Przechowywane jako `Decimal` w kolumnie `number`
- **GraphQL**: Zwracane jako typ `Float`

### Format wyświetlania
- **Interfejs użytkownika**: Aplikacje klienckie muszą dodać symbol % (np. "75.5%")
- **Wykresy**: Wyświetlane z symbolem % gdy typ wyjścia to PERCENTAGE
- **Odpowiedzi API**: Surowa wartość numeryczna bez symbolu % (np. 75.5)

## Tworzenie rekordów z wartościami procentowymi

Podczas tworzenia nowego rekordu z wartościami procentowymi:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Obsługiwane formaty wejściowe

| Format | Przykład | Wynik |
|--------|----------|-------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Uwaga**: Symbol % jest automatycznie usuwany z wejścia i dodawany ponownie podczas wyświetlania.

## Zapytania o wartości procentowe

Podczas zapytań o rekordy z niestandardowymi polami procentowymi, uzyskaj wartość przez ścieżkę `customField.value.number`:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

Odpowiedź będzie zawierać procent jako surową liczbę:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | Unikalny identyfikator dla wartości pola |
| `customField` | CustomField! | Definicja pola niestandardowego (zawiera wartość procentową) |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

**Ważne**: Wartości procentowe są uzyskiwane przez pole `customField.value.number`. Symbol % nie jest uwzględniany w przechowywanych wartościach i musi być dodany przez aplikacje klienckie do wyświetlania.

## Filtrowanie i zapytania

Pola procentowe obsługują te same filtry co pola NUMBER:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Obsługiwane operatory

| Operator | Opis | Przykład |
|----------|------|---------|
| `EQ` | Równy | `percentage = 75` |
| `NE` | Nie równy | `percentage ≠ 75` |
| `GT` | Większy niż | `percentage > 75` |
| `GTE` | Większy lub równy | `percentage ≥ 75` |
| `LT` | Mniejszy niż | `percentage < 75` |
| `LTE` | Mniejszy lub równy | `percentage ≤ 75` |
| `IN` | Wartość w liście | `percentage in [50, 75, 100]` |
| `NIN` | Wartość nie w liście | `percentage not in [0, 25]` |
| `IS` | Sprawdź null z `values: null` | `percentage is null` |
| `NOT` | Sprawdź, czy nie null z `values: null` | `percentage is not null` |

### Filtrowanie zakresu

Do filtrowania zakresu użyj wielu operatorów:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Zakresy wartości procentowych

### Typowe zakresy

| Zakres | Opis | Przykład użycia |
|--------|------|-----------------|
| `0-100` | Standardowy procent | Completion rates, success rates |
| `0-∞` | Nieograniczony procent | Growth rates, performance metrics |
| `-∞-∞` | Dowolna wartość | Change rates, variance |

### Przykładowe wartości

| Wejście | Przechowywane | Wyświetlane |
|---------|---------------|-------------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Agregacja wykresów

Pola procentowe obsługują agregację w wykresach i raportach na pulpicie. Dostępne funkcje obejmują:

- `AVERAGE` - Średnia wartość procentowa
- `COUNT` - Liczba rekordów z wartościami
- `MIN` - Najniższa wartość procentowa
- `MAX` - Najwyższa wartość procentowa 
- `SUM` - Suma wszystkich wartości procentowych

Te agregacje są dostępne podczas tworzenia wykresów i pulpitów, a nie w bezpośrednich zapytaniach GraphQL.

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|----------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Odpowiedzi błędów

### Nieprawidłowy format procentowy
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Nie jest liczbą
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Najlepsze praktyki

### Wprowadzanie wartości
- Pozwól użytkownikom wprowadzać z lub bez symbolu %
- Waliduj rozsądne zakresy dla swojego przypadku użycia
- Zapewnij jasny kontekst, co oznacza 100%

### Wyświetlanie
- Zawsze pokazuj symbol % w interfejsach użytkownika
- Używaj odpowiedniej precyzji dziesiętnej
- Rozważ kodowanie kolorami dla zakresów (czerwony/żółty/zielony)

### Interpretacja danych
- Udokumentuj, co oznacza 100% w twoim kontekście
- Odpowiednio obsługuj wartości powyżej 100%
- Rozważ, czy wartości ujemne są ważne

## Typowe przypadki użycia

1. **Zarządzanie projektami**
   - Wskaźniki ukończenia zadań
   - Postęp projektu
   - Wykorzystanie zasobów
   - Prędkość sprintu

2. **Śledzenie wydajności**
   - Wskaźniki sukcesu
   - Wskaźniki błędów
   - Metryki efektywności
   - Wyniki jakości

3. **Metryki finansowe**
   - Wskaźniki wzrostu
   - Marże zysku
   - Kwoty rabatów
   - Procenty zmian

4. **Analiza**
   - Wskaźniki konwersji
   - Wskaźniki klikalności
   - Metryki zaangażowania
   - Wskaźniki wydajności

## Funkcje integracji

### Z formułami
- Odwołuj się do pól PERCENT w obliczeniach
- Automatyczne formatowanie symbolu % w wynikach formuł
- Łącz z innymi polami numerycznymi

### Z automatyzacjami
- Uruchamiaj akcje na podstawie progów procentowych
- Wysyłaj powiadomienia dla procentów kamieni milowych
- Aktualizuj status na podstawie wskaźników ukończenia

### Z wyszukiwaniami
- Agreguj procenty z powiązanych rekordów
- Obliczaj średnie wskaźniki sukcesu
- Znajdź najwyżej/najniżej oceniane elementy

### Z wykresami
- Twórz wizualizacje oparte na procentach
- Śledź postęp w czasie
- Porównuj metryki wydajności

## Różnice w stosunku do pól NUMBER

### Co jest inne
- **Obsługa wejścia**: Automatycznie usuwa symbol %
- **Wyświetlanie**: Automatycznie dodaje symbol %
- **Ograniczenia**: Brak walidacji min/max
- **Formatowanie**: Brak wsparcia dla prefiksów

### Co jest takie samo
- **Przechowywanie**: Ta sama kolumna i typ bazy danych
- **Filtrowanie**: Te same operatory zapytań
- **Agregacja**: Te same funkcje agregacji
- **Uprawnienia**: Ten sam model uprawnień

## Ograniczenia

- Brak ograniczeń wartości min/max
- Brak opcji formatowania prefiksów
- Brak automatycznej walidacji zakresu 0-100%
- Brak konwersji między formatami procentowymi (np. 0.75 ↔ 75%)
- Wartości powyżej 100% są dozwolone

## Powiązane zasoby

- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne koncepcje pól niestandardowych
- [Pole niestandardowe liczby](/api/custom-fields/number) - Dla surowych wartości numerycznych
- [API automatyzacji](/api/automations/index) - Twórz automatyzacje oparte na procentach