---
title: Pole niestandardowe z pojedynczym wyborem
description: Twórz pola z pojedynczym wyborem, aby umożliwić użytkownikom wybór jednej opcji z predefiniowanej listy
---

Pola niestandardowe z pojedynczym wyborem pozwalają użytkownikom wybrać dokładnie jedną opcję z predefiniowanej listy. Są idealne do pól statusu, kategorii, priorytetów lub w każdej sytuacji, w której należy dokonać tylko jednego wyboru z kontrolowanego zestawu opcji.

## Podstawowy przykład

Utwórz proste pole z pojedynczym wyborem:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole z pojedynczym wyborem z predefiniowanymi opcjami:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola z pojedynczym wyborem |
| `type` | CustomFieldType! | ✅ Tak | Musi być `SELECT_SINGLE` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | Nie | Opcje początkowe dla pola |

### CreateCustomFieldOptionInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `title` | String! | ✅ Tak | Tekst wyświetlany dla opcji |
| `color` | String | Nie | Kod koloru hex dla opcji |

## Dodawanie opcji do istniejących pól

Dodaj nowe opcje do istniejącego pola z pojedynczym wyborem:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## Ustawianie wartości pojedynczego wyboru

Aby ustawić wybraną opcję w rekordzie:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola z pojedynczym wyborem |
| `customFieldOptionId` | String | Nie | ID opcji do wybrania (preferowane dla pojedynczego wyboru) |
| `customFieldOptionIds` | [String!] | Nie | Tablica ID opcji (używa pierwszego elementu dla pojedynczego wyboru) |

## Zapytanie o wartości pojedynczego wyboru

Zapytaj o wartość pojedynczego wyboru rekordu:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

Pole `value` zwraca obiekt JSON z szczegółami wybranej opcji.

## Tworzenie rekordów z wartościami pojedynczego wyboru

Podczas tworzenia nowego rekordu z wartościami pojedynczego wyboru:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator dla wartości pola |
| `customField` | CustomField! | Definicja pola niestandardowego |
| `value` | JSON | Zawiera obiekt wybranej opcji z id, tytułem, kolorem, pozycją |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

### Odpowiedź CustomFieldOption

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator dla opcji |
| `title` | String! | Tekst wyświetlany dla opcji |
| `color` | String | Kod koloru hex dla reprezentacji wizualnej |
| `position` | Float | Kolejność sortowania dla opcji |
| `customField` | CustomField! | Pole niestandardowe, do którego należy ta opcja |

### Odpowiedź CustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator dla pola |
| `name` | String! | Nazwa wyświetlana pola z pojedynczym wyborem |
| `type` | CustomFieldType! | Zawsze `SELECT_SINGLE` |
| `description` | String | Tekst pomocy dla pola |
| `customFieldOptions` | [CustomFieldOption!] | Wszystkie dostępne opcje |

## Format wartości

### Format wejściowy
- **Parametr API**: Użyj `customFieldOptionId` dla ID pojedynczej opcji
- **Alternatywa**: Użyj `customFieldOptionIds` tablicy (bierze pierwszy element)
- **Czyszczenie wyboru**: Pomiń oba pola lub przekaż puste wartości

### Format wyjściowy
- **Odpowiedź GraphQL**: Obiekt JSON w `value` polu zawierający {id, tytuł, kolor, pozycja}
- **Dziennik aktywności**: Tytuł opcji jako ciąg
- **Dane automatyzacji**: Tytuł opcji jako ciąg

## Zachowanie wyboru

### Wyłączny wybór
- Ustawienie nowej opcji automatycznie usuwa poprzedni wybór
- Tylko jedna opcja może być wybrana w danym czasie
- Ustawienie `null` lub pustej wartości czyści wybór

### Logika zapasowa
- Jeśli podano tablicę `customFieldOptionIds`, używana jest tylko pierwsza opcja
- Zapewnia to zgodność z formatami wejściowymi dla wielu wyborów
- Puste tablice lub wartości null czyszczą wybór

## Zarządzanie opcjami

### Aktualizacja właściwości opcji
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### Usunięcie opcji
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**Uwaga**: Usunięcie opcji spowoduje jej usunięcie ze wszystkich rekordów, w których była wybrana.

### Zmiana kolejności opcji
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## Zasady walidacji

### Walidacja opcji
- Podane ID opcji musi istnieć
- Opcja musi należeć do określonego pola niestandardowego
- Tylko jedna opcja może być wybrana (wymuszone automatycznie)
- Wartości null/puste są ważne (brak wyboru)

### Walidacja pola
- Musi mieć zdefiniowaną co najmniej jedną opcję, aby mogło być używane
- Tytuły opcji muszą być unikalne w obrębie pola
- Kody kolorów muszą być w poprawnym formacie hex (jeśli podano)

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Odpowiedzi na błędy

### Nieprawidłowe ID opcji
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### Opcja nie należy do pola
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

### Nie można zanalizować wartości
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Najlepsze praktyki

### Projektowanie opcji
- Używaj jasnych, opisowych tytułów opcji
- Stosuj znaczące kodowanie kolorów
- Utrzymuj listy opcji skoncentrowane i odpowiednie
- Porządkuj opcje logicznie (według priorytetu, częstotliwości itp.)

### Wzorce pól statusu
- Używaj spójnych przepływów pracy statusów w projektach
- Rozważ naturalny postęp opcji
- Uwzględnij jasne stany końcowe (Zrobione, Anulowane itp.)
- Używaj kolorów, które odzwierciedlają znaczenie opcji

### Zarządzanie danymi
- Okresowo przeglądaj i porządkuj nieużywane opcje
- Używaj spójnych konwencji nazewnictwa
- Rozważ wpływ usunięcia opcji na istniejące rekordy
- Planuj aktualizacje i migracje opcji

## Typowe przypadki użycia

1. **Status i przepływ pracy**
   - Status zadania (Do zrobienia, W toku, Zrobione)
   - Status zatwierdzenia (Oczekujące, Zatwierdzone, Odrzucone)
   - Faza projektu (Planowanie, Rozwój, Testowanie, Wydane)
   - Status rozwiązania problemu

2. **Klasyfikacja i kategoryzacja**
   - Poziomy priorytetu (Niski, Średni, Wysoki, Krytyczny)
   - Typy zadań (Błąd, Funkcja, Ulepszenie, Dokumentacja)
   - Kategorie projektów (Wewnętrzne, Klient, Badania)
   - Przypisania do działów

3. **Jakość i ocena**
   - Status przeglądu (Nie rozpoczęto, W przeglądzie, Zatwierdzone)
   - Oceny jakości (Słaba, Przyzwoita, Dobra, Doskonała)
   - Poziomy ryzyka (Niski, Średni, Wysoki)
   - Poziomy pewności

4. **Przypisanie i własność**
   - Przypisania zespołowe
   - Własność działu
   - Przypisania oparte na rolach
   - Przypisania regionalne

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje, gdy wybrane są określone opcje
- Kieruj pracą na podstawie wybranych kategorii
- Wysyłaj powiadomienia o zmianach statusu
- Twórz warunkowe przepływy pracy na podstawie wyborów

### Z wyszukiwaniami
- Filtruj rekordy według wybranych opcji
- Odwołuj się do danych opcji z innych rekordów
- Twórz raporty na podstawie wyborów opcji
- Grupuj rekordy według wybranych wartości

### Z formularzami
- Kontrole wejściowe w rozwijanych listach
- Interfejsy przycisków radiowych
- Walidacja i filtrowanie opcji
- Warunkowe wyświetlanie pól na podstawie wyborów

## Śledzenie aktywności

Zmiany w polach z pojedynczym wyborem są automatycznie śledzone:
- Pokazuje stare i nowe wybory opcji
- Wyświetla tytuły opcji w dzienniku aktywności
- Znaczniki czasowe dla wszystkich zmian wyboru
- Przypisanie użytkownika do modyfikacji

## Różnice w porównaniu do wyboru wielokrotnego

| Cecha | Pojedynczy wybór | Wybór wielokrotny |
|-------|------------------|-------------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Ograniczenia

- Tylko jedna opcja może być wybrana w danym czasie
- Brak hierarchicznej lub zagnieżdżonej struktury opcji
- Opcje są współdzielone we wszystkich rekordach korzystających z pola
- Brak wbudowanej analityki opcji lub śledzenia użycia
- Kody kolorów są tylko do wyświetlania, nie mają wpływu na funkcjonalność
- Nie można ustawić różnych uprawnień dla każdej opcji

## Powiązane zasoby

- [Pola wielokrotnego wyboru](/api/custom-fields/select-multi) - Dla wyborów wielokrotnych
- [Pola wyboru](/api/custom-fields/checkbox) - Dla prostych wyborów boolean
- [Pola tekstowe](/api/custom-fields/text-single) - Dla swobodnego wprowadzania tekstu
- [Przegląd pól niestandardowych](/api/custom-fields/1.index) - Ogólne pojęcia