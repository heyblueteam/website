---
title: Pole niestandardowe typu Checkbox
description: Tworzenie pól checkbox dla danych typu tak/nie lub prawda/fałsz
---

Pola niestandardowe typu checkbox zapewniają prosty boolean (prawda/fałsz) do zadań. Są idealne do wyborów binarnych, wskaźników statusu lub śledzenia, czy coś zostało ukończone.

## Podstawowy przykład

Utwórz proste pole checkbox:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole checkbox z opisem i walidacją:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola checkbox |
| `type` | CustomFieldType! | ✅ Tak | Musi być `CHECKBOX` |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

## Ustawianie wartości checkbox

Aby ustawić lub zaktualizować wartość checkbox w zadaniu:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Aby odznaczyć checkbox:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID zadania do aktualizacji |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola checkbox |
| `checked` | Boolean | Nie | Prawda, aby zaznaczyć, fałsz, aby odznaczyć |

## Tworzenie zadań z wartościami checkbox

Podczas tworzenia nowego zadania z wartościami checkbox:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Akceptowane wartości ciągów

Podczas tworzenia zadań wartości checkbox muszą być przekazywane jako ciągi:

| Wartość ciągu | Wynik |
|----------------|-------|
| `"true"` | ✅ Zaznaczone (wrażliwe na wielkość liter) |
| `"1"` | ✅ Zaznaczone |
| `"checked"` | ✅ Zaznaczone (wrażliwe na wielkość liter) |
| Any other value | ❌ Odznaczone |

**Uwaga**: Porównania ciągów podczas tworzenia zadań są wrażliwe na wielkość liter. Wartości muszą dokładnie pasować do `"true"`, `"1"` lub `"checked"`, aby uzyskać stan zaznaczony.

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | Unikalny identyfikator dla wartości pola |
| `uid` | String! | Alternatywny unikalny identyfikator |
| `customField` | CustomField! | Definicja pola niestandardowego |
| `checked` | Boolean | Stan checkboxa (prawda/fałsz/null) |
| `todo` | Todo! | Zadanie, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

## Integracja automatyzacji

Pola checkbox wyzwalają różne zdarzenia automatyzacji na podstawie zmian stanu:

| Akcja | Wyzwalane zdarzenie | Opis |
|-------|---------------------|------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Wyzwalane, gdy checkbox jest zaznaczony |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Wyzwalane, gdy checkbox jest odznaczony |

Pozwala to na tworzenie automatyzacji, które reagują na zmiany stanu checkboxa, takie jak:
- Wysyłanie powiadomień, gdy przedmioty są zatwierdzane
- Przenoszenie zadań, gdy checkboxy przeglądu są zaznaczone
- Aktualizowanie powiązanych pól na podstawie stanów checkboxów

## Import/eksport danych

### Importowanie wartości checkbox

Podczas importowania danych za pomocą CSV lub innych formatów:
- `"true"`, `"yes"` → Zaznaczone (niezależnie od wielkości liter)
- Każda inna wartość (w tym `"false"`, `"no"`, `"0"`, pusta) → Odznaczone

### Eksportowanie wartości checkbox

Podczas eksportowania danych:
- Zaznaczone pola eksportują jako `"X"`
- Odznaczone pola eksportują jako pusty ciąg `""`

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|----------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Odpowiedzi na błędy

### Nieprawidłowy typ wartości
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Najlepsze praktyki

### Konwencje nazewnictwa
- Używaj jasnych, zorientowanych na działanie nazw: "Zatwierdzone", "Przejrzane", "Ukończone"
- Unikaj negatywnych nazw, które mogą mylić użytkowników: preferuj "Aktywne" zamiast "Nieaktywne"
- Bądź konkretny co do tego, co reprezentuje checkbox

### Kiedy używać checkboxów
- **Wybory binarne**: Tak/Nie, Prawda/Fałsz, Ukończone/Nieukończone
- **Wskaźniki statusu**: Zatwierdzone, Przejrzane, Opublikowane
- **Flagi funkcji**: Ma Priorytetowe Wsparcie, Wymaga Podpisu
- **Proste śledzenie**: Email Wysłany, Faktura Opłacona, Przedmiot Wysłany

### Kiedy NIE używać checkboxów
- Gdy potrzebujesz więcej niż dwóch opcji (użyj SELECT_SINGLE zamiast tego)
- Dla danych numerycznych lub tekstowych (użyj pól NUMBER lub TEXT)
- Gdy musisz śledzić, kto zaznaczył lub kiedy (użyj dzienników audytowych)

## Typowe przypadki użycia

1. **Workflow zatwierdzania**
   - "Zatwierdzone przez menedżera"
   - "Podpis klienta"
   - "Przegląd prawny zakończony"

2. **Zarządzanie zadaniami**
   - "Jest zablokowane"
   - "Gotowe do przeglądu"
   - "Wysoki priorytet"

3. **Kontrola jakości**
   - "QA Zatwierdzone"
   - "Dokumentacja ukończona"
   - "Testy napisane"

4. **Flagi administracyjne**
   - "Faktura wysłana"
   - "Umowa podpisana"
   - "Wymagana kontynuacja"

## Ograniczenia

- Pola checkbox mogą przechowywać tylko wartości prawda/fałsz (brak stanu trójdrożnego lub null po początkowym ustawieniu)
- Brak konfiguracji wartości domyślnej (zawsze zaczyna się jako null, aż do ustawienia)
- Nie można przechowywać dodatkowych metadanych, takich jak kto zaznaczył lub kiedy
- Brak warunkowej widoczności na podstawie wartości innych pól

## Powiązane zasoby

- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne pojęcia dotyczące pól niestandardowych
- [API automatyzacji](/api/automations) - Tworzenie automatyzacji wyzwalanych przez zmiany checkboxów