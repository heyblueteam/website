---
title: Pole niestandardowe tekstu wieloliniowego
description: Tworzenie pól tekstowych wieloliniowych dla dłuższych treści, takich jak opisy, notatki i komentarze
---

Pola niestandardowe tekstu wieloliniowego pozwalają na przechowywanie dłuższych treści tekstowych z łamaniem linii i formatowaniem. Są idealne do opisów, notatek, komentarzy lub wszelkich danych tekstowych, które wymagają wielu linii.

## Podstawowy przykład

Utwórz proste pole tekstowe wieloliniowe:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole tekstowe wieloliniowe z opisem:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `type` | CustomFieldType! | ✅ Tak | Musi być `TEXT_MULTI` |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

**Uwaga:** `projectId` jest przekazywane jako osobny argument do mutacji, a nie jako część obiektu wejściowego. Alternatywnie, kontekst projektu można określić z nagłówka `X-Bloo-Project-ID` w żądaniu GraphQL.

## Ustawianie wartości tekstowych

Aby ustawić lub zaktualizować wartość tekstu wieloliniowego w rekordzie:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID pola tekstowego niestandardowego |
| `text` | String | Nie | Treść tekstu wieloliniowego do przechowania |

## Tworzenie rekordów z wartościami tekstowymi

Podczas tworzenia nowego rekordu z wartościami tekstu wieloliniowego:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
| `text` | String | Przechowywana treść tekstu wieloliniowego |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

## Walidacja tekstu

### Walidacja formularza
Gdy pola tekstowe wieloliniowe są używane w formularzach:
- Wiodące i końcowe białe znaki są automatycznie usuwane
- Walidacja wymagana jest stosowana, jeśli pole jest oznaczone jako wymagane
- Nie stosuje się żadnej specyficznej walidacji formatu

### Zasady walidacji
- Akceptuje dowolną treść tekstową, w tym łamanie linii
- Brak ograniczeń długości znaków (do limitów bazy danych)
- Obsługuje znaki Unicode i symbole specjalne
- Łamanie linii jest zachowywane w przechowywaniu

### Przykłady poprawnego tekstu
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Ważne uwagi

### Pojemność przechowywania
- Przechowywane przy użyciu typu MySQL `MediumText`
- Obsługuje do 16 MB treści tekstowej
- Łamanie linii i formatowanie są zachowywane
- Kodowanie UTF-8 dla znaków międzynarodowych

### API bezpośrednie vs formularze
- **Formularze**: Automatyczne usuwanie białych znaków i walidacja wymagana
- **API bezpośrednie**: Tekst jest przechowywany dokładnie tak, jak podano
- **Zalecenie**: Używaj formularzy do wprowadzania danych przez użytkowników, aby zapewnić spójne formatowanie

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI**: Wprowadzenie tekstu wieloliniowego, idealne do dłuższych treści
- **TEXT_SINGLE**: Wprowadzenie tekstu jednoliniowego, idealne do krótkich wartości
- **Backend**: Oba typy są identyczne - to samo pole przechowywania, walidacja i przetwarzanie
- **Frontend**: Różne komponenty UI do wprowadzania danych (textarea vs pole wejściowe)
- **Ważne**: Rozróżnienie między TEXT_MULTI a TEXT_SINGLE istnieje wyłącznie w celach UI

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|----------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

## Odpowiedzi błędów

### Walidacja pola wymaganego (tylko formularze)
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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Najlepsze praktyki

### Organizacja treści
- Używaj spójnego formatowania dla zorganizowanej treści
- Rozważ użycie składni podobnej do markdown dla czytelności
- Dziel długie treści na logiczne sekcje
- Używaj łamania linii, aby poprawić czytelność

### Wprowadzanie danych
- Podaj jasne opisy pól, aby prowadzić użytkowników
- Używaj formularzy do wprowadzania danych przez użytkowników, aby zapewnić walidację
- Rozważ ograniczenia znaków w zależności od przypadku użycia
- Waliduj format treści w swojej aplikacji, jeśli to konieczne

### Rozważania dotyczące wydajności
- Bardzo długie treści tekstowe mogą wpływać na wydajność zapytań
- Rozważ paginację do wyświetlania dużych pól tekstowych
- Rozważania dotyczące indeksowania dla funkcjonalności wyszukiwania
- Monitoruj wykorzystanie pamięci dla pól z dużą zawartością

## Filtrowanie i wyszukiwanie

### Wyszukiwanie zawierające
Pola tekstowe wieloliniowe obsługują wyszukiwanie podciągów za pomocą filtrów pól niestandardowych:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Możliwości wyszukiwania
- Dopasowanie podciągu w polach tekstowych przy użyciu operatora `CONTAINS`
- Wyszukiwanie bez uwzględnienia wielkości liter przy użyciu operatora `NCONTAINS`
- Dokładne dopasowanie przy użyciu operatora `IS`
- Dopasowanie negatywne przy użyciu operatora `NOT`
- Wyszukiwanie w obrębie wszystkich linii tekstu
- Obsługuje częściowe dopasowanie słów

## Typowe przypadki użycia

1. **Zarządzanie projektami**
   - Opisy zadań
   - Wymagania projektowe
   - Notatki ze spotkań
   - Aktualizacje statusu

2. **Wsparcie klienta**
   - Opisy problemów
   - Notatki dotyczące rozwiązań
   - Opinie klientów
   - Dzienniki komunikacji

3. **Zarządzanie treścią**
   - Treść artykułów
   - Opisy produktów
   - Komentarze użytkowników
   - Szczegóły recenzji

4. **Dokumentacja**
   - Opisy procesów
   - Instrukcje
   - Wytyczne
   - Materiały referencyjne

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje, gdy treść tekstowa się zmienia
- Ekstrahuj słowa kluczowe z treści tekstowej
- Twórz podsumowania lub powiadomienia
- Przetwarzaj treść tekstową za pomocą usług zewnętrznych

### Z wyszukiwaniami
- Odwołuj się do danych tekstowych z innych rekordów
- Agreguj treści tekstowe z wielu źródeł
- Znajduj rekordy według treści tekstowej
- Wyświetlaj powiązane informacje tekstowe

### Z formularzami
- Automatyczne usuwanie białych znaków
- Walidacja pól wymaganych
- UI dla wieloliniowego pola tekstowego
- Wyświetlanie liczby znaków (jeśli skonfigurowane)

## Ograniczenia

- Brak wbudowanego formatowania tekstu lub edytora tekstu bogatego
- Brak automatycznego wykrywania lub konwersji linków
- Brak sprawdzania pisowni lub walidacji gramatycznej
- Brak wbudowanej analizy lub przetwarzania tekstu
- Brak wersjonowania lub śledzenia zmian
- Ograniczone możliwości wyszukiwania (brak pełnotekstowego wyszukiwania)
- Brak kompresji treści dla bardzo dużego tekstu

## Powiązane zasoby

- [Pola tekstowe jednoliniowe](/api/custom-fields/text-single) - Dla krótkich wartości tekstowych
- [Pola e-mailowe](/api/custom-fields/email) - Dla adresów e-mail
- [Pola URL](/api/custom-fields/url) - Dla adresów stron internetowych
- [Przegląd pól niestandardowych](/api/custom-fields/2.list-custom-fields) - Ogólne pojęcia