---
title: Pole niestandardowe kraju
description: Tworzenie pól wyboru kraju z walidacją kodu kraju ISO
---

Pola niestandardowe kraju pozwalają na przechowywanie i zarządzanie informacjami o krajach dla rekordów. Pole obsługuje zarówno nazwy krajów, jak i kody krajów ISO Alpha-2.

**Ważne**: Zachowanie walidacji i konwersji krajów różni się znacznie między mutacjami:
- **createTodo**: Automatycznie waliduje i konwertuje nazwy krajów na kody ISO
- **setTodoCustomField**: Akceptuje dowolną wartość bez walidacji

## Podstawowy przykład

Utwórz proste pole kraju:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole kraju z opisem:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola kraju |
| `type` | CustomFieldType! | ✅ Tak | Musi być `COUNTRY` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

**Uwaga**: `projectId` nie jest przekazywany w wejściu, ale jest określany przez kontekst GraphQL (zazwyczaj z nagłówków żądania lub uwierzytelnienia).

## Ustawianie wartości kraju

Pola krajowe przechowują dane w dwóch polach bazy danych:
- **`countryCodes`**: Przechowuje kody krajów ISO Alpha-2 jako ciąg oddzielony przecinkami w bazie danych (zwracane jako tablica przez API)
- **`text`**: Przechowuje tekst wyświetlany lub nazwy krajów jako ciąg

### Zrozumienie parametrów

Mutacja `setTodoCustomField` akceptuje dwa opcjonalne parametry dla pól krajowych:

| Parametr | Typ | Wymagany | Opis | Co robi |
|----------|-----|----------|------|---------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania | - |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola kraju | - |
| `countryCodes` | [String!] | Nie | Tablica kodów krajów ISO Alpha-2 | Stored in the `countryCodes` field |
| `text` | String | Nie | Tekst wyświetlany lub nazwy krajów | Stored in the `text` field |

**Ważne**: 
- W `setTodoCustomField`: Oba parametry są opcjonalne i przechowywane niezależnie
- W `createTodo`: System automatycznie ustawia oba pola na podstawie twojego wejścia (nie możesz nimi zarządzać niezależnie)

### Opcja 1: Używanie tylko kodów krajów

Przechowuj zwalidowane kody ISO bez tekstu wyświetlanego:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Wynik: `countryCodes` = `["US"]`, `text` = `null`

### Opcja 2: Używanie tylko tekstu

Przechowuj tekst wyświetlany bez zwalidowanych kodów:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Wynik: `countryCodes` = `null`, `text` = `"United States"`

**Uwaga**: Używając `setTodoCustomField`, nie zachodzi walidacja niezależnie od tego, który parametr używasz. Wartości są przechowywane dokładnie tak, jak podano.

### Opcja 3: Używanie obu (zalecane)

Przechowuj zarówno zwalidowane kody, jak i tekst wyświetlany:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Wynik: `countryCodes` = `["US"]`, `text` = `"United States"`

### Wiele krajów

Przechowuj wiele krajów, używając tablic:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Tworzenie rekordów z wartościami krajów

Podczas tworzenia rekordów, mutacja `createTodo` **automatycznie waliduje i konwertuje** wartości krajów. To jest jedyna mutacja, która wykonuje walidację krajów:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      text
      countryCodes
    }
  }
}
```

### Akceptowane formaty wejściowe

| Typ wejściowy | Przykład | Wynik |
|---------------|----------|-------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Nieobsługiwane** - traktowane jako pojedyncza nieprawidłowa wartość |
| Mixed format | `"United States, CA"` | **Nieobsługiwane** - traktowane jako pojedyncza nieprawidłowa wartość |

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator dla wartości pola |
| `customField` | CustomField! | Definicja niestandardowego pola |
| `text` | String | Tekst wyświetlany (nazwy krajów) |
| `countryCodes` | [String!] | Tablica kodów krajów ISO Alpha-2 |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

## Standardy krajowe

Blue używa standardu **ISO 3166-1 Alpha-2** dla kodów krajów:

- Kody krajów składające się z dwóch liter (np. US, GB, FR, DE)
- Walidacja przy użyciu biblioteki `i18n-iso-countries` **zachodzi tylko w createTodo**
- Obsługuje wszystkie oficjalnie uznawane kraje

### Przykładowe kody krajów

| Kraj | Kod ISO |
|------|---------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Aby uzyskać pełną oficjalną listę kodów krajów ISO 3166-1 alpha-2, odwiedź [ISO Online Browsing Platform](https://www.iso.org/obp/ui/#search/code/).

## Walidacja

**Walidacja zachodzi tylko w mutacji `createTodo`**:

1. **Prawidłowy kod ISO**: Akceptuje każdy prawidłowy kod ISO Alpha-2
2. **Nazwa kraju**: Automatycznie konwertuje rozpoznane nazwy krajów na kody
3. **Nieprawidłowe wejście**: Zgłasza `CustomFieldValueParseError` dla nierozpoznanych wartości

**Uwaga**: Mutacja `setTodoCustomField` nie wykonuje żadnej walidacji i akceptuje dowolną wartość tekstową.

### Przykład błędu

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Funkcje integracji

### Pola wyszukiwania
Pola krajowe mogą być odniesione przez niestandardowe pola WYSZUKIWANIA, co pozwala na pobieranie danych krajowych z powiązanych rekordów.

### Automatyzacje
Używaj wartości krajów w warunkach automatyzacji:
- Filtruj działania według konkretnych krajów
- Wysyłaj powiadomienia w zależności od kraju
- Kieruj zadania w zależności od regionów geograficznych

### Formularze
Pola krajowe w formularzach automatycznie walidują dane wprowadzone przez użytkownika i konwertują nazwy krajów na kody.

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Odpowiedzi błędów

### Nieprawidłowa wartość kraju
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Niezgodność typu pola
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Najlepsze praktyki

### Obsługa wejścia
- Używaj `createTodo` do automatycznej walidacji i konwersji
- Używaj `setTodoCustomField` ostrożnie, ponieważ omija walidację
- Rozważ walidację wejść w swojej aplikacji przed użyciem `setTodoCustomField`
- Wyświetlaj pełne nazwy krajów w interfejsie użytkownika dla jasności

### Jakość danych
- Waliduj dane krajowe w punkcie wejścia
- Używaj spójnych formatów w całym systemie
- Rozważ grupowanie regionalne do raportowania

### Wiele krajów
- Używaj wsparcia tablic w `setTodoCustomField` dla wielu krajów
- Wiele krajów w `createTodo` **nie jest obsługiwane** przez pole wartości
- Przechowuj kody krajów jako tablicę w `setTodoCustomField` dla prawidłowego przetwarzania

## Typowe przypadki użycia

1. **Zarządzanie klientami**
   - Lokalizacja siedziby klienta
   - Miejsca wysyłki
   - Jurysdykcje podatkowe

2. **Śledzenie projektów**
   - Lokalizacja projektu
   - Lokalizacje członków zespołu
   - Cele rynkowe

3. **Zgodność i prawo**
   - Jurysdykcje regulacyjne
   - Wymagania dotyczące miejsca przechowywania danych
   - Kontrole eksportowe

4. **Sprzedaż i marketing**
   - Przydziały terytorialne
   - Segmentacja rynku
   - Targetowanie kampanii

## Ograniczenia

- Obsługuje tylko kody ISO 3166-1 Alpha-2 (kody składające się z 2 liter)
- Brak wbudowanego wsparcia dla podziałów krajowych (stanów/prowincji)
- Brak automatycznych ikon flag krajowych (tylko tekstowe)
- Nie można walidować historycznych kodów krajów
- Brak wbudowanego grupowania regionów lub kontynentów
- **Walidacja działa tylko w `createTodo`, a nie w `setTodoCustomField`**
- **Wiele krajów nie jest obsługiwane w polu wartości `createTodo`**
- **Kody krajów przechowywane jako ciąg oddzielony przecinkami, a nie jako prawdziwa tablica**

## Powiązane zasoby

- [Przegląd pól niestandardowych](/custom-fields/list-custom-fields) - Ogólne koncepcje pól niestandardowych
- [Pola wyszukiwania](/api/custom-fields/lookup) - Odniesienie danych krajowych z innych rekordów
- [API formularzy](/api/forms) - Uwzględnienie pól krajowych w niestandardowych formularzach