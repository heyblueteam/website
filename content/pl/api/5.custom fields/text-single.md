---
title: Pole niestandardowe tekstu jednoliniowego
description: Twórz pola tekstowe jednoliniowe dla krótkich wartości tekstowych, takich jak imiona, tytuły i etykiety
---

Pola niestandardowe tekstu jednoliniowego pozwalają na przechowywanie krótkich wartości tekstowych przeznaczonych do wprowadzania w jednej linii. Są idealne do imion, tytułów, etykiet lub jakichkolwiek danych tekstowych, które powinny być wyświetlane w jednej linii.

## Podstawowy przykład

Utwórz proste pole tekstowe jednoliniowe:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole tekstowe jednoliniowe z opisem:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola tekstowego |
| `type` | CustomFieldType! | ✅ Tak | Musi być `TEXT_SINGLE` |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

**Uwaga**: Kontekst projektu jest automatycznie określany na podstawie nagłówków autoryzacji. Nie jest wymagany parametr `projectId`.

## Ustawianie wartości tekstowych

Aby ustawić lub zaktualizować wartość tekstową jednoliniową w rekordzie:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID pola tekstowego niestandardowego |
| `text` | String | Nie | Treść tekstowa jednoliniowa do przechowania |

## Tworzenie rekordów z wartościami tekstowymi

Podczas tworzenia nowego rekordu z wartościami tekstowymi jednoliniowymi:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja pola niestandardowego (zawiera wartość tekstową) |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

**Ważne**: Wartości tekstowe są dostępne przez pole `customField.value.text`, a nie bezpośrednio w TodoCustomField.

## Zapytania o wartości tekstowe

Podczas zapytań o rekordy z polami tekstowymi niestandardowymi, uzyskaj dostęp do tekstu przez ścieżkę `customField.value.text`:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

Odpowiedź będzie zawierać tekst w zagnieżdżonej strukturze:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Walidacja tekstu

### Walidacja formularza
Gdy pola tekstowe jednoliniowe są używane w formularzach:
- Wiodące i końcowe białe znaki są automatycznie usuwane
- Walidacja wymagana jest stosowana, jeśli pole jest oznaczone jako wymagane
- Nie stosuje się żadnej konkretnej walidacji formatu

### Zasady walidacji
- Akceptuje dowolną zawartość tekstową, w tym znaki nowej linii (choć nie jest to zalecane)
- Brak ograniczeń długości znaków (do limitów bazy danych)
- Obsługuje znaki Unicode i symbole specjalne
- Znaki nowej linii są zachowywane, ale nie są przeznaczone dla tego typu pola

### Typowe przykłady tekstu
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Ważne uwagi

### Pojemność przechowywania
- Przechowywane przy użyciu typu MySQL `MediumText`
- Obsługuje do 16MB zawartości tekstowej
- Identyczna pojemność do pól tekstowych wieloliniowych
- Kodowanie UTF-8 dla znaków międzynarodowych

### Bezpośrednie API vs Formularze
- **Formularze**: Automatyczne usuwanie białych znaków i walidacja wymagana
- **Bezpośrednie API**: Tekst jest przechowywany dokładnie tak, jak podano
- **Zalecenie**: Używaj formularzy do wprowadzania danych przez użytkowników, aby zapewnić spójne formatowanie

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: Wprowadzanie tekstu jednoliniowego, idealne dla krótkich wartości
- **TEXT_MULTI**: Wprowadzanie tekstu w polu wieloliniowym, idealne dla dłuższej zawartości
- **Backend**: Oba używają identycznego przechowywania i walidacji
- **Frontend**: Różne komponenty UI do wprowadzania danych
- **Zamiar**: TEXT_SINGLE jest semantycznie przeznaczone dla wartości jednoliniowych

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|----------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## Odpowiedzi błędów

### Walidacja wymaganego pola (tylko formularze)
```json
{
  "errors": [{
    "message": "This field is required",
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
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

## Najlepsze praktyki

### Wytyczne dotyczące treści
- Zachowaj tekst zwięzły i odpowiedni do jednoliniowego wyświetlania
- Unikaj łamań linii dla zamierzonego wyświetlania jednoliniowego
- Używaj spójnego formatowania dla podobnych typów danych
- Rozważ limity znaków w oparciu o wymagania UI

### Wprowadzanie danych
- Podaj jasne opisy pól, aby prowadzić użytkowników
- Używaj formularzy do wprowadzania danych przez użytkowników, aby zapewnić walidację
- Waliduj format zawartości w swojej aplikacji, jeśli to konieczne
- Rozważ użycie rozwijanych list dla ustandaryzowanych wartości

### Rozważania dotyczące wydajności
- Pola tekstowe jednoliniowe są lekkie i wydajne
- Rozważ indeksowanie dla często przeszukiwanych pól
- Używaj odpowiednich szerokości wyświetlania w swoim UI
- Monitoruj długość zawartości w celach wyświetlania

## Filtrowanie i wyszukiwanie

### Wyszukiwanie zawierające
Pola tekstowe jednoliniowe obsługują wyszukiwanie podciągów:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Możliwości wyszukiwania
- Dopasowanie podciągów bez uwzględnienia wielkości liter
- Obsługuje częściowe dopasowanie słów
- Dopasowanie dokładnych wartości
- Brak wyszukiwania pełnotekstowego ani rankingowania

## Typowe przypadki użycia

1. **Identyfikatory i kody**
   - SKU produktów
   - Numery zamówień
   - Kody referencyjne
   - Numery wersji

2. **Imiona i tytuły**
   - Imiona klientów
   - Tytuły projektów
   - Nazwy produktów
   - Etykiety kategorii

3. **Krótkie opisy**
   - Krótkie podsumowania
   - Etykiety statusu
   - Wskaźniki priorytetu
   - Etykiety klasyfikacyjne

4. **Odniesienia zewnętrzne**
   - Numery biletów
   - Odniesienia do faktur
   - ID systemów zewnętrznych
   - Numery dokumentów

## Funkcje integracji

### Z wyszukiwaniami
- Odwołania do danych tekstowych z innych rekordów
- Znajdowanie rekordów według zawartości tekstowej
- Wyświetlanie powiązanych informacji tekstowych
- Agregowanie wartości tekstowych z wielu źródeł

### Z formularzami
- Automatyczne usuwanie białych znaków
- Walidacja wymaganych pól
- UI do wprowadzania tekstu jednoliniowego
- Wyświetlanie limitu znaków (jeśli skonfigurowane)

### Z importami/eksportami
- Bezpośrednie mapowanie kolumn CSV
- Automatyczne przypisywanie wartości tekstowych
- Obsługa importu danych zbiorczych
- Eksport do formatów arkuszy kalkulacyjnych

## Ograniczenia

### Ograniczenia automatyzacji
- Nie jest bezpośrednio dostępne jako pola wyzwalające automatyzację
- Nie można używać w aktualizacjach pól automatyzacji
- Można odwoływać się w warunkach automatyzacji
- Dostępne w szablonach e-mail i webhookach

### Ogólne ograniczenia
- Brak wbudowanego formatowania lub stylizacji tekstu
- Brak automatycznej walidacji poza wymaganymi polami
- Brak wbudowanego egzekwowania unikalności
- Brak kompresji zawartości dla bardzo dużego tekstu
- Brak wersjonowania lub śledzenia zmian
- Ograniczone możliwości wyszukiwania (brak wyszukiwania pełnotekstowego)

## Powiązane zasoby

- [Pola tekstowe wieloliniowe](/api/custom-fields/text-multi) - Dla dłuższej zawartości tekstowej
- [Pola e-mailowe](/api/custom-fields/email) - Dla adresów e-mail
- [Pola URL](/api/custom-fields/url) - Dla adresów stron internetowych
- [Pola unikalnych ID](/api/custom-fields/unique-id) - Dla automatycznie generowanych identyfikatorów
- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne koncepcje