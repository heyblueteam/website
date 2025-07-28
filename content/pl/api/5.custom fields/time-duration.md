---
title: Niestandardowe pole czasu trwania
description: Twórz obliczane pola czasu trwania, które śledzą czas między zdarzeniami w Twoim przepływie pracy
---

Niestandardowe pola czasu trwania automatycznie obliczają i wyświetlają czas trwania między dwoma zdarzeniami w Twoim przepływie pracy. Są idealne do śledzenia czasów przetwarzania, czasów odpowiedzi, czasów cyklu lub jakichkolwiek metryk opartych na czasie w Twoich projektach.

## Podstawowy przykład

Utwórz proste pole czasu trwania, które śledzi, jak długo trwa wykonanie zadań:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Zaawansowany przykład

Utwórz złożone pole czasu trwania, które śledzi czas między zmianami pól niestandardowych z celem SLA:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput (TIME_DURATION)

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola czasu trwania |
| `type` | CustomFieldType! | ✅ Tak | Musi być `TIME_DURATION` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Tak | Jak wyświetlać czas trwania |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Tak | Konfiguracja zdarzenia początkowego |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Tak | Konfiguracja zdarzenia końcowego |
| `timeDurationTargetTime` | Float | Nie | Docelowy czas trwania w sekundach dla monitorowania SLA |

### CustomFieldTimeDurationInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `type` | CustomFieldTimeDurationType! | ✅ Tak | Typ zdarzenia do śledzenia |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Tak | `FIRST` lub `LAST` wystąpienie |
| `customFieldId` | String | Conditional | Wymagane dla `TODO_CUSTOM_FIELD` typu |
| `customFieldOptionIds` | [String!] | Conditional | Wymagane dla zmian pól wyboru |
| `todoListId` | String | Conditional | Wymagane dla `TODO_MOVED` typu |
| `tagId` | String | Conditional | Wymagane dla `TODO_TAG_ADDED` typu |
| `assigneeId` | String | Conditional | Wymagane dla `TODO_ASSIGNEE_ADDED` typu |

### Wartości CustomFieldTimeDurationType

| Wartość | Opis |
|---------|------|
| `TODO_CREATED_AT` | Kiedy rekord został utworzony |
| `TODO_CUSTOM_FIELD` | Kiedy zmieniła się wartość pola niestandardowego |
| `TODO_DUE_DATE` | Kiedy ustawiono termin |
| `TODO_MARKED_AS_COMPLETE` | Kiedy rekord został oznaczony jako ukończony |
| `TODO_MOVED` | Kiedy rekord został przeniesiony do innej listy |
| `TODO_TAG_ADDED` | Kiedy tag został dodany do rekordu |
| `TODO_ASSIGNEE_ADDED` | Kiedy przypisany został dodany do rekordu |

### Wartości CustomFieldTimeDurationCondition

| Wartość | Opis |
|---------|------|
| `FIRST` | Użyj pierwszego wystąpienia zdarzenia |
| `LAST` | Użyj ostatniego wystąpienia zdarzenia |

### Wartości CustomFieldTimeDurationDisplayType

| Wartość | Opis | Przykład |
|---------|------|---------|
| `FULL_DATE` | Format Dni:Godziny:Minuty:Sekundy | `"01:02:03:04"` |
| `FULL_DATE_STRING` | W pełni zapisane słowami | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Liczbowo z jednostkami | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Czas trwania tylko w dniach | `"2.5"` (2.5 days) |
| `HOURS` | Czas trwania tylko w godzinach | `"60"` (60 hours) |
| `MINUTES` | Czas trwania tylko w minutach | `"3600"` (3600 minutes) |
| `SECONDS` | Czas trwania tylko w sekundach | `"216000"` (216000 seconds) |

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja pola niestandardowego |
| `number` | Float | Czas trwania w sekundach |
| `value` | Float | Alias dla liczby (czas trwania w sekundach) |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zaktualizowana |

### Odpowiedź CustomField (TIME_DURATION)

| Pole | Typ | Opis |
|------|-----|------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Format wyświetlania dla czasu trwania |
| `timeDurationStart` | CustomFieldTimeDuration | Konfiguracja zdarzenia początkowego |
| `timeDurationEnd` | CustomFieldTimeDuration | Konfiguracja zdarzenia końcowego |
| `timeDurationTargetTime` | Float | Docelowy czas trwania w sekundach (do monitorowania SLA) |

## Obliczanie czasu trwania

### Jak to działa
1. **Zdarzenie początkowe**: System monitoruje określone zdarzenie początkowe
2. **Zdarzenie końcowe**: System monitoruje określone zdarzenie końcowe
3. **Obliczenie**: Czas trwania = Czas zakończenia - Czas rozpoczęcia
4. **Przechowywanie**: Czas trwania przechowywany w sekundach jako liczba
5. **Wyświetlanie**: Formatowane zgodnie z ustawieniem `timeDurationDisplay`

### Wyzwalacze aktualizacji
Wartości czasu trwania są automatycznie przeliczane, gdy:
- Rekordy są tworzone lub aktualizowane
- Wartości pól niestandardowych się zmieniają
- Tagii są dodawane lub usuwane
- Przypisania są dodawane lub usuwane
- Rekordy są przenoszone między listami
- Rekordy są oznaczane jako ukończone/niedokończone

## Odczytywanie wartości czasu trwania

### Zapytanie o pola czasu trwania
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Sformatowane wartości wyświetlania
Wartości czasu trwania są automatycznie formatowane na podstawie ustawienia `timeDurationDisplay`:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Przykłady powszechnej konfiguracji

### Czas zakończenia zadania
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Czas zmiany statusu
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Czas w określonej liście
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Czas odpowiedzi na przypisanie
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Odpowiedzi błędów

### Nieprawidłowa konfiguracja
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Nie znaleziono pola referencyjnego
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

### Brak wymaganych opcji
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Ważne uwagi

### Automatyczne obliczenie
- Pola czasu trwania są **tylko do odczytu** - wartości są automatycznie obliczane
- Nie możesz ręcznie ustawiać wartości czasu trwania za pomocą API
- Obliczenia odbywają się asynchronicznie za pomocą zadań w tle
- Wartości aktualizują się automatycznie, gdy wystąpią zdarzenia wyzwalające

### Rozważania dotyczące wydajności
- Obliczenia czasu trwania są kolejkowane i przetwarzane asynchronicznie
- Duża liczba pól czasu trwania może wpłynąć na wydajność
- Rozważ częstotliwość zdarzeń wyzwalających podczas projektowania pól czasu trwania
- Używaj specyficznych warunków, aby uniknąć niepotrzebnych przeliczeń

### Wartości null
Pola czasu trwania będą wyświetlać `null`, gdy:
- Zdarzenie początkowe jeszcze się nie zdarzyło
- Zdarzenie końcowe jeszcze się nie zdarzyło
- Konfiguracja odnosi się do nieistniejących jednostek
- Obliczenie napotka błąd

## Najlepsze praktyki

### Projektowanie konfiguracji
- Używaj specyficznych typów zdarzeń zamiast ogólnych, gdy to możliwe
- Wybierz odpowiednie `FIRST` w porównaniu do `LAST` warunki w zależności od Twojego przepływu pracy
- Testuj obliczenia czasu trwania z próbnymi danymi przed wdrożeniem
- Dokumentuj logikę swojego pola czasu trwania dla członków zespołu

### Formatowanie wyświetlania
- Używaj `FULL_DATE_SUBSTRING` dla najbardziej czytelnego formatu
- Używaj `FULL_DATE` dla kompaktowego, spójnego wyświetlania szerokości
- Używaj `FULL_DATE_STRING` dla formalnych raportów i dokumentów
- Używaj `DAYS`, `HOURS`, `MINUTES` lub `SECONDS` dla prostych wyświetleń liczbowych
- Rozważ ograniczenia przestrzeni UI przy wyborze formatu

### Monitorowanie SLA z docelowym czasem
Podczas korzystania z `timeDurationTargetTime`:
- Ustaw docelowy czas trwania w sekundach
- Porównaj rzeczywisty czas trwania z celem w celu zgodności z SLA
- Używaj w pulpitach nawigacyjnych, aby wyróżnić przeterminowane elementy
- Przykład: 24-godzinny SLA na odpowiedź = 86400 sekund

### Integracja z przepływem pracy
- Projektuj pola czasu trwania, aby odpowiadały rzeczywistym procesom biznesowym
- Używaj danych czasu trwania do poprawy i optymalizacji procesów
- Monitoruj trendy czasu trwania, aby zidentyfikować wąskie gardła w przepływie pracy
- Ustaw alerty dla progów czasu trwania, jeśli to konieczne

## Powszechne przypadki użycia

1. **Wydajność procesów**
   - Czas zakończenia zadań
   - Czas cyklu przeglądów
   - Czas przetwarzania zatwierdzeń
   - Czas odpowiedzi

2. **Monitorowanie SLA**
   - Czas do pierwszej odpowiedzi
   - Czas rozwiązywania problemów
   - Ramy czasowe eskalacji
   - Zgodność z poziomem usług

3. **Analiza przepływu pracy**
   - Identyfikacja wąskich gardeł
   - Optymalizacja procesów
   - Metryki wydajności zespołu
   - Czas zapewnienia jakości

4. **Zarządzanie projektami**
   - Czas trwania faz
   - Czas realizacji kamieni milowych
   - Czas alokacji zasobów
   - Ramy czasowe dostawy

## Ograniczenia

- Pola czasu trwania są **tylko do odczytu** i nie mogą być ustawiane ręcznie
- Wartości są obliczane asynchronicznie i mogą nie być natychmiast dostępne
- Wymaga odpowiednich wyzwalaczy zdarzeń, aby zostały skonfigurowane w Twoim przepływie pracy
- Nie można obliczyć czasów trwania dla zdarzeń, które się jeszcze nie zdarzyły
- Ograniczone do śledzenia czasu między odrębnymi zdarzeniami (nie ciągłe śledzenie czasu)
- Brak wbudowanych alertów lub powiadomień SLA
- Nie można agregować wielu obliczeń czasu trwania w jednym polu

## Powiązane zasoby

- [Pola liczbowe](/api/custom-fields/number) - Do ręcznych wartości liczbowych
- [Pola daty](/api/custom-fields/date) - Do śledzenia konkretnych dat
- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne koncepcje
- [Automatyzacje](/api/automations) - Do wyzwalania akcji na podstawie progów czasu trwania