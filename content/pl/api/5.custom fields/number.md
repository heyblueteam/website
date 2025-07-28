---
title: Pole niestandardowe liczby
description: Twórz pola liczbowo, aby przechowywać wartości numeryczne z opcjonalnymi ograniczeniami min/max i formatowaniem prefiksów
---

Pola niestandardowe liczby pozwalają na przechowywanie wartości numerycznych dla rekordów. Obsługują ograniczenia walidacji, precyzję dziesiętną i mogą być używane do ilości, wyników, pomiarów lub jakichkolwiek danych numerycznych, które nie wymagają specjalnego formatowania.

## Podstawowy przykład

Utwórz proste pole liczbowe:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole liczbowe z ograniczeniami i prefiksem:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola liczbowego |
| `type` | CustomFieldType! | ✅ Tak | Musi być `NUMBER` |
| `projectId` | String! | ✅ Tak | ID projektu, w którym ma być utworzone pole |
| `min` | Float | Nie | Ograniczenie wartości minimalnej (tylko wskazówki UI) |
| `max` | Float | Nie | Ograniczenie wartości maksymalnej (tylko wskazówki UI) |
| `prefix` | String | Nie | Prefiks wyświetlania (np. "#", "~", "$") |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

## Ustawianie wartości liczbowych

Pola liczbowe przechowują wartości dziesiętne z opcjonalną walidacją:

### Prosta wartość liczby

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Wartość całkowita

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego liczby |
| `number` | Float | Nie | Wartość numeryczna do przechowania |

## Ograniczenia wartości

### Ograniczenia min/max (wskazówki UI)

**Ważne**: Ograniczenia min/max są przechowywane, ale NIE są egzekwowane po stronie serwera. Służą jako wskazówki UI dla aplikacji frontendowych.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Wymagana walidacja po stronie klienta**: Aplikacje frontendowe muszą implementować logikę walidacji, aby egzekwować ograniczenia min/max.

### Obsługiwane typy wartości

| Typ | Przykład | Opis |
|------|---------|-------------|
| Integer | `42` | Liczby całkowite |
| Decimal | `42.5` | Liczby z miejscami dziesiętnymi |
| Negative | `-10` | Wartości ujemne (jeśli brak ograniczenia min) |
| Zero | `0` | Wartość zero |

**Uwaga**: Ograniczenia min/max NIE są walidowane po stronie serwera. Wartości spoza określonego zakresu będą akceptowane i przechowywane.

## Tworzenie rekordów z wartościami liczbowymi

Podczas tworzenia nowego rekordu z wartościami liczbowymi:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### Obsługiwane formaty wejściowe

Podczas tworzenia rekordów użyj parametru `number` (nie `value`) w tablicy pól niestandardowych:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|-------|------|-------------|
| `id` | String! | Unikalny identyfikator dla wartości pola |
| `customField` | CustomField! | Definicja pola niestandardowego |
| `number` | Float | Wartość numeryczna |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość była ostatnio modyfikowana |

### Odpowiedź CustomField

| Pole | Typ | Opis |
|-------|------|-------------|
| `id` | String! | Unikalny identyfikator dla definicji pola |
| `name` | String! | Nazwa wyświetlana pola |
| `type` | CustomFieldType! | Zawsze `NUMBER` |
| `min` | Float | Minimalna dozwolona wartość |
| `max` | Float | Maksymalna dozwolona wartość |
| `prefix` | String | Prefiks wyświetlania |
| `description` | String | Tekst pomocy |

**Uwaga**: Jeśli wartość liczby nie jest ustawiona, pole `number` będzie `null`.

## Filtrowanie i zapytania

Pola liczbowe obsługują kompleksowe filtrowanie numeryczne:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
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
|----------|-------------|---------|
| `EQ` | Równy | `number = 42` |
| `NE` | Nie równy | `number ≠ 42` |
| `GT` | Większy niż | `number > 42` |
| `GTE` | Większy lub równy | `number ≥ 42` |
| `LT` | Mniejszy niż | `number < 42` |
| `LTE` | Mniejszy lub równy | `number ≤ 42` |
| `IN` | W tablicy | `number in [1, 2, 3]` |
| `NIN` | Nie w tablicy | `number not in [1, 2, 3]` |
| `IS` | Jest null/nie jest null | `number is null` |

### Filtrowanie zakresu

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Formatowanie wyświetlania

### Z prefiksem

Jeśli ustawiony jest prefiks, będzie wyświetlany:

| Wartość | Prefiks | Wyświetlanie |
|-------|--------|---------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Precyzja dziesiętna

Liczby zachowują swoją precyzję dziesiętną:

| Wejście | Przechowywane | Wyświetlane |
|-------|--------|-----------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|--------|--------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Odpowiedzi błędów

### Nieprawidłowy format liczby
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Uwaga**: Błędy walidacji min/max nie występują po stronie serwera. Walidacja ograniczeń musi być zaimplementowana w Twojej aplikacji frontendowej.

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

### Projektowanie ograniczeń
- Ustaw realistyczne wartości min/max jako wskazówki UI
- Implementuj walidację po stronie klienta, aby egzekwować ograniczenia
- Używaj ograniczeń, aby zapewnić użytkownikom informacje zwrotne w formularzach
- Rozważ, czy wartości ujemne są ważne dla Twojego przypadku użycia

### Precyzja wartości
- Używaj odpowiedniej precyzji dziesiętnej dla swoich potrzeb
- Rozważ zaokrąglanie w celach wyświetlania
- Bądź konsekwentny w precyzji w powiązanych polach

### Udoskonalenie wyświetlania
- Używaj znaczących prefiksów dla kontekstu
- Rozważ jednostki w nazwach pól (np. "Waga (kg)")
- Podawaj jasne opisy dla reguł walidacji

## Typowe przypadki użycia

1. **Systemy oceniania**
   - Oceny wydajności
   - Oceny jakości
   - Poziomy priorytetu
   - Oceny satysfakcji klientów

2. **Pomiar**
   - Ilości i kwoty
   - Wymiary i rozmiary
   - Czas trwania (w formacie numerycznym)
   - Pojemności i limity

3. **Metryki biznesowe**
   - Wartości przychodów
   - Wskaźniki konwersji
   - Alokacje budżetowe
   - Liczby docelowe

4. **Dane techniczne**
   - Numery wersji
   - Wartości konfiguracyjne
   - Metryki wydajności
   - Ustawienia progowe

## Funkcje integracji

### Z wykresami i pulpitami
- Używaj pól NUMER w obliczeniach wykresów
- Twórz wizualizacje numeryczne
- Śledź trendy w czasie

### Z automatyzacjami
- Wyzwalaj akcje na podstawie progów liczbowych
- Aktualizuj powiązane pola na podstawie zmian liczbowych
- Wysyłaj powiadomienia dla określonych wartości

### Z wyszukiwaniami
- Agreguj liczby z powiązanych rekordów
- Obliczaj sumy i średnie
- Znajduj wartości min/max w relacjach

### Z wykresami
- Twórz wizualizacje numeryczne
- Śledź trendy w czasie
- Porównuj wartości między rekordami

## Ograniczenia

- **Brak walidacji po stronie serwera** ograniczeń min/max
- **Wymagana walidacja po stronie klienta** dla egzekwowania ograniczeń
- Brak wbudowanego formatowania walutowego (użyj zamiast tego typu WALUTA)
- Brak automatycznego symbolu procentowego (użyj zamiast tego typu PROCENT)
- Brak możliwości konwersji jednostek
- Precyzja dziesiętna ograniczona przez typ Decimal bazy danych
- Brak oceny formuły matematycznej w samym polu

## Powiązane zasoby

- [Przegląd pól niestandardowych](/api/custom-fields/1.index) - Ogólne koncepcje pól niestandardowych
- [Pole niestandardowe waluty](/api/custom-fields/currency) - Dla wartości pieniężnych
- [Pole niestandardowe procentowe](/api/custom-fields/percent) - Dla wartości procentowych
- [API automatyzacji](/api/automations/1.index) - Twórz automatyzacje oparte na liczbach