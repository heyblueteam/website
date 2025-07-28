---
title: Pole niestandardowe z wieloma wyborami
description: Twórz pola z wieloma wyborami, aby umożliwić użytkownikom wybór wielu opcji z predefiniowanej listy
---

Pola niestandardowe z wieloma wyborami pozwalają użytkownikom na wybór wielu opcji z predefiniowanej listy. Są idealne do kategorii, tagów, umiejętności, funkcji lub jakiegokolwiek scenariusza, w którym potrzebne są wielokrotne wybory z kontrolowanego zestawu opcji.

## Podstawowy przykład

Utwórz proste pole z wieloma wyborami:

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole z wieloma wyborami, a następnie dodaj opcje osobno:

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola z wieloma wyborami |
| `type` | CustomFieldType! | ✅ Tak | Musi być `SELECT_MULTI` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |
| `projectId` | String! | ✅ Tak | ID projektu dla tego pola |

### CreateCustomFieldOptionInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego |
| `title` | String! | ✅ Tak | Tekst wyświetlany dla opcji |
| `color` | String | Nie | Kolor dla opcji (dowolny ciąg) |
| `position` | Float | Nie | Kolejność sortowania dla opcji |

## Dodawanie opcji do istniejących pól

Dodaj nowe opcje do istniejącego pola z wieloma wyborami:

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## Ustawianie wartości z wieloma wyborami

Aby ustawić wiele wybranych opcji w rekordzie:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego z wieloma wyborami |
| `customFieldOptionIds` | [String!] | ✅ Tak | Tablica ID opcji do wybrania |

## Tworzenie rekordów z wartościami z wieloma wyborami

Podczas tworzenia nowego rekordu z wartościami z wieloma wyborami:

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja pola niestandardowego |
| `selectedOptions` | [CustomFieldOption!] | Tablica wybranych opcji |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

### Odpowiedź CustomFieldOption

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator opcji |
| `title` | String! | Tekst wyświetlany dla opcji |
| `color` | String | Kod koloru hex dla wizualnej reprezentacji |
| `position` | Float | Kolejność sortowania dla opcji |
| `customField` | CustomField! | Pole niestandardowe, do którego należy ta opcja |

### Odpowiedź CustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator pola |
| `name` | String! | Nazwa wyświetlana pola z wieloma wyborami |
| `type` | CustomFieldType! | Zawsze `SELECT_MULTI` |
| `description` | String | Tekst pomocy dla pola |
| `customFieldOptions` | [CustomFieldOption!] | Wszystkie dostępne opcje |

## Format wartości

### Format wejściowy
- **Parametr API**: Tablica ID opcji (`["option1", "option2", "option3"]`)
- **Format ciągu**: ID opcji oddzielone przecinkami (`"option1,option2,option3"`)

### Format wyjściowy
- **Odpowiedź GraphQL**: Tablica obiektów CustomFieldOption
- **Dziennik aktywności**: Tytuły opcji oddzielone przecinkami
- **Dane automatyzacji**: Tablica tytułów opcji

## Zarządzanie opcjami

### Aktualizacja właściwości opcji
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
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

### Zmiana kolejności opcji
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## Zasady walidacji

### Walidacja opcji
- Wszystkie podane ID opcji muszą istnieć
- Opcje muszą należeć do określonego pola niestandardowego
- Tylko pola SELECT_MULTI mogą mieć wybrane wiele opcji
- Pusta tablica jest ważna (brak wyborów)

### Walidacja pola
- Musi mieć zdefiniowaną przynajmniej jedną opcję, aby mogło być używane
- Tytuły opcji muszą być unikalne w obrębie pola
- Pole koloru akceptuje dowolną wartość ciągu (brak walidacji hex)

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Odpowiedzi błędów

### Nieprawidłowe ID opcji
```json
{
  "errors": [{
    "message": "Custom field option not found",
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Wiele opcji w polu, które nie jest wielokrotne
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Najlepsze praktyki

### Projektowanie opcji
- Używaj opisowych, zwięzłych tytułów opcji
- Stosuj spójne schematy kolorów
- Utrzymuj listy opcji w rozsądnych granicach (zwykle 3-20 opcji)
- Porządkuj opcje logicznie (alfabetycznie, według częstotliwości itp.)

### Zarządzanie danymi
- Okresowo przeglądaj i porządkuj nieużywane opcje
- Używaj spójnych konwencji nazewniczych w projektach
- Rozważ możliwość ponownego wykorzystania opcji przy tworzeniu pól
- Planuj aktualizacje i migracje opcji

### Doświadczenie użytkownika
- Zapewnij jasne opisy pól
- Używaj kolorów, aby poprawić wizualną różnicę
- Grupuj powiązane opcje razem
- Rozważ domyślne wybory dla typowych przypadków

## Typowe przypadki użycia

1. **Zarządzanie projektami**
   - Kategorie i tagi zadań
   - Poziomy i typy priorytetów
   - Przypisania członków zespołu
   - Wskaźniki statusu

2. **Zarządzanie treścią**
   - Kategorie i tematy artykułów
   - Typy i formaty treści
   - Kanały publikacji
   - Procesy zatwierdzania

3. **Wsparcie klienta**
   - Kategorie i typy problemów
   - Dotknięte produkty lub usługi
   - Metody rozwiązania
   - Segmenty klientów

4. **Rozwój produktu**
   - Kategorie funkcji
   - Wymagania techniczne
   - Środowiska testowe
   - Kanały wydania

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje, gdy wybrane są określone opcje
- Kieruj pracą na podstawie wybranych kategorii
- Wysyłaj powiadomienia dla wyborów o wysokim priorytecie
- Twórz zadania do śledzenia na podstawie kombinacji opcji

### Z wyszukiwaniami
- Filtruj rekordy według wybranych opcji
- Agreguj dane w oparciu o wybory opcji
- Odwołuj się do danych opcji z innych rekordów
- Twórz raporty na podstawie kombinacji opcji

### Z formularzami
- Kontrolki wejściowe z wieloma wyborami
- Walidacja i filtrowanie opcji
- Dynamiczne ładowanie opcji
- Warunkowe wyświetlanie pól

## Śledzenie aktywności

Zmiany w polu z wieloma wyborami są automatycznie śledzone:
- Pokazuje dodane i usunięte opcje
- Wyświetla tytuły opcji w dzienniku aktywności
- Znaczniki czasowe dla wszystkich zmian wyboru
- Przypisanie użytkownika do modyfikacji

## Ograniczenia

- Maksymalny praktyczny limit opcji zależy od wydajności interfejsu użytkownika
- Brak hierarchicznej lub zagnieżdżonej struktury opcji
- Opcje są współdzielone we wszystkich rekordach korzystających z pola
- Brak wbudowanej analityki opcji lub śledzenia użycia
- Pole koloru akceptuje dowolny ciąg (brak walidacji hex)
- Nie można ustawić różnych uprawnień dla każdej opcji
- Opcje muszą być tworzone osobno, a nie w linii z tworzeniem pola
- Brak dedykowanej mutacji zmiany kolejności (użyj editCustomFieldOption z pozycją)

## Powiązane zasoby

- [Pola z pojedynczym wyborem](/api/custom-fields/select-single) - Do wyborów jednolitych
- [Pola wyboru](/api/custom-fields/checkbox) - Do prostych wyborów boolowskich
- [Pola tekstowe](/api/custom-fields/text-single) - Do wprowadzania tekstu w formie wolnej
- [Przegląd pól niestandardowych](/api/custom-fields/2.list-custom-fields) - Ogólne koncepcje