---
title: Pole niestandardowe oceny
description: Twórz pola oceny, aby przechowywać numeryczne oceny z konfigurowalnymi skalami i walidacją
---

Pola niestandardowe oceny pozwalają na przechowywanie numerycznych ocen w rekordach z konfigurowalnymi wartościami minimalnymi i maksymalnymi. Są idealne do ocen wydajności, wyników satysfakcji, poziomów priorytetu lub wszelkich danych opartych na skali numerycznej w Twoich projektach.

## Podstawowy przykład

Utwórz proste pole oceny z domyślną skalą 0-5:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Zaawansowany przykład

Utwórz pole oceny z niestandardową skalą i opisem:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola oceny |
| `type` | CustomFieldType! | ✅ Tak | Musi być `RATING` |
| `projectId` | String! | ✅ Tak | ID projektu, w którym to pole zostanie utworzone |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |
| `min` | Float | Nie | Minimalna wartość oceny (brak domyślnej) |
| `max` | Float | Nie | Maksymalna wartość oceny |

## Ustawianie wartości ocen

Aby ustawić lub zaktualizować wartość oceny w rekordzie:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola oceny |
| `value` | String! | ✅ Tak | Wartość oceny jako ciąg (w ramach skonfigurowanego zakresu) |

## Tworzenie rekordów z wartościami ocen

Podczas tworzenia nowego rekordu z wartościami ocen:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja niestandardowego pola |
| `value` | Float | Przechowywana wartość oceny (dostępna przez customField.value) |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

**Uwaga**: Wartość oceny jest faktycznie dostępna przez `customField.value.number` w zapytaniach.

### Odpowiedź CustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator pola |
| `name` | String! | Nazwa wyświetlana pola oceny |
| `type` | CustomFieldType! | Zawsze `RATING` |
| `min` | Float | Minimalna dozwolona wartość oceny |
| `max` | Float | Maksymalna dozwolona wartość oceny |
| `description` | String | Tekst pomocy dla pola |

## Walidacja ocen

### Ograniczenia wartości
- Wartości ocen muszą być numeryczne (typ Float)
- Wartości muszą mieścić się w skonfigurowanym zakresie min/max
- Jeśli nie określono minimum, nie ma wartości domyślnej
- Wartość maksymalna jest opcjonalna, ale zalecana

### Zasady walidacji
**Ważne**: Walidacja występuje tylko podczas przesyłania formularzy, a nie podczas używania `setTodoCustomField` bezpośrednio.

- Wprowadzenie jest analizowane jako liczba zmiennoprzecinkowa (podczas korzystania z formularzy)
- Musi być większe lub równe wartości minimalnej (podczas korzystania z formularzy)
- Musi być mniejsze lub równe wartości maksymalnej (podczas korzystania z formularzy)
- `setTodoCustomField` akceptuje dowolną wartość ciągu bez walidacji

### Przykłady prawidłowych ocen
Dla pola z min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Przykłady nieprawidłowych ocen
Dla pola z min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Opcje konfiguracyjne

### Ustawienie skali ocen
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Typowe skale ocen
- **1-5 gwiazdek**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 wydajności**: `min: 1, max: 10`
- **0-100 procent**: `min: 0, max: 100`
- **Niestandardowa skala**: Dowolny zakres numeryczny

## Wymagane uprawnienia

Operacje na polach niestandardowych podlegają standardowym uprawnieniom opartym na rolach:

| Akcja | Wymagana rola |
|-------|---------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Uwaga**: Wymagane konkretne role zależą od konfiguracji ról niestandardowych w Twoim projekcie oraz uprawnień na poziomie pól.

## Odpowiedzi na błędy

### Błąd walidacji (tylko formularze)
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

**Ważne**: Walidacja wartości ocen (ograniczenia min/max) występuje tylko podczas przesyłania formularzy, a nie podczas używania `setTodoCustomField` bezpośrednio.

### Niestandardowe pole nie znalezione
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

## Najlepsze praktyki

### Projektowanie skali
- Używaj spójnych skal ocen w podobnych polach
- Weź pod uwagę znajomość użytkowników (1-5 gwiazdek, 0-10 NPS)
- Ustaw odpowiednie wartości minimalne (0 vs 1)
- Zdefiniuj jasne znaczenie dla każdego poziomu oceny

### Jakość danych
- Waliduj wartości ocen przed ich przechowywaniem
- Używaj odpowiedniej precyzji dziesiętnej
- Rozważ zaokrąglanie do celów wyświetlania
- Zapewnij jasne wskazówki dotyczące znaczenia ocen

### Doświadczenie użytkownika
- Wyświetlaj skale ocen wizualnie (gwiazdki, paski postępu)
- Pokaż bieżącą wartość i limity skali
- Zapewnij kontekst dla znaczenia ocen
- Rozważ wartości domyślne dla nowych rekordów

## Typowe przypadki użycia

1. **Zarządzanie wydajnością**
   - Oceny wydajności pracowników
   - Wyniki jakości projektu
   - Oceny ukończenia zadań
   - Oceny poziomu umiejętności

2. **Opinie klientów**
   - Oceny satysfakcji
   - Wyniki jakości produktu
   - Oceny doświadczenia obsługi
   - Net Promoter Score (NPS)

3. **Priorytet i ważność**
   - Poziomy priorytetu zadań
   - Oceny pilności
   - Wyniki oceny ryzyka
   - Oceny wpływu

4. **Zapewnienie jakości**
   - Oceny przeglądów kodu
   - Wyniki jakości testowania
   - Jakość dokumentacji
   - Oceny przestrzegania procesów

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje na podstawie progów ocen
- Wysyłaj powiadomienia o niskich ocenach
- Twórz zadania do śledzenia dla wysokich ocen
- Kieruj pracę na podstawie wartości ocen

### Z wyszukiwaniami
- Obliczaj średnie oceny w rekordach
- Znajdź rekordy według zakresów ocen
- Odwołuj się do danych ocen z innych rekordów
- Agreguj statystyki ocen

### Z interfejsem Blue
- Automatyczna walidacja zakresu w kontekście formularzy
- Wizualne kontrolki wejściowe ocen
- Informacje zwrotne o walidacji w czasie rzeczywistym
- Opcje wejściowe w postaci gwiazdek lub suwaków

## Śledzenie aktywności

Zmiany w polach ocen są automatycznie śledzone:
- Stare i nowe wartości ocen są rejestrowane
- Aktywność pokazuje zmiany numeryczne
- Znaczniki czasowe dla wszystkich aktualizacji ocen
- Przypisanie użytkownika do zmian

## Ograniczenia

- Obsługiwane są tylko wartości numeryczne
- Brak wbudowanego wizualnego wyświetlania ocen (gwiazdki itp.)
- Precyzja dziesiętna zależy od konfiguracji bazy danych
- Brak przechowywania metadanych ocen (komentarze, kontekst)
- Brak automatycznej agregacji ocen lub statystyk
- Brak wbudowanej konwersji ocen między skalami
- **Krytyczne**: Walidacja min/max działa tylko w formularzach, a nie za pomocą `setTodoCustomField`

## Powiązane zasoby

- [Pola liczbowе](/api/5.custom%20fields/number) - Do ogólnych danych numerycznych
- [Pola procentowe](/api/5.custom%20fields/percent) - Do wartości procentowych
- [Pola wyboru](/api/5.custom%20fields/select-single) - Do ocen opartych na wyborze dyskretnym
- [Przegląd pól niestandardowych](/api/5.custom%20fields/2.list-custom-fields) - Ogólne pojęcia