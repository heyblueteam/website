---
title: Pole niestandardowe przycisku
description: Twórz interaktywne pola przycisków, które uruchamiają automatyzacje po kliknięciu
---

Pola niestandardowe przycisku zapewniają interaktywne elementy UI, które uruchamiają automatyzacje po kliknięciu. W przeciwieństwie do innych typów pól niestandardowych, które przechowują dane, pola przycisków służą jako wyzwalacze akcji do wykonania skonfigurowanych przepływów pracy.

## Podstawowy przykład

Utwórz proste pole przycisku, które uruchamia automatyzację:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz przycisk z wymaganiami potwierdzenia:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana przycisku |
| `type` | CustomFieldType! | ✅ Tak | Musi być `BUTTON` |
| `projectId` | String! | ✅ Tak | ID projektu, w którym pole zostanie utworzone |
| `buttonType` | String | Nie | Zachowanie potwierdzenia (zobacz Typy przycisków poniżej) |
| `buttonConfirmText` | String | Nie | Tekst, który użytkownicy muszą wpisać w celu twardego potwierdzenia |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |
| `required` | Boolean | Nie | Czy pole jest wymagane (domyślnie fałsz) |
| `isActive` | Boolean | Nie | Czy pole jest aktywne (domyślnie prawda) |

### Typ pola przycisku

Pole `buttonType` jest ciągiem w formie wolnej, który może być używany przez klientów UI do określenia zachowania potwierdzenia. Powszechne wartości to:

- `""` (puste) - Brak potwierdzenia
- `"soft"` - Prosty dialog potwierdzenia
- `"hard"` - Wymaga wpisania tekstu potwierdzenia

**Uwaga**: To są tylko wskazówki UI. API nie weryfikuje ani nie egzekwuje konkretnych wartości.

## Uruchamianie kliknięć przycisku

Aby uruchomić kliknięcie przycisku i wykonać związane automatyzacje:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Parametry wejściowe kliknięcia

| Parametr | Typ | Wymagane | Opis |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Tak | ID zadania zawierającego przycisk |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego przycisku |

### Ważne: Zachowanie API

**Wszystkie kliknięcia przycisków przez API wykonują się natychmiast** niezależnie od jakichkolwiek ustawień `buttonType` lub `buttonConfirmText`. Te pola są przechowywane dla klientów UI, aby wdrożyć dialogi potwierdzenia, ale samo API:

- Nie weryfikuje tekstu potwierdzenia
- Nie egzekwuje żadnych wymagań dotyczących potwierdzenia
- Wykonuje akcję przycisku natychmiast po wywołaniu

Potwierdzenie jest czysto funkcją bezpieczeństwa po stronie klienta UI.

### Przykład: Klikanie różnych typów przycisków

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Wszystkie trzy mutacje powyżej wykonają akcję przycisku natychmiast po wywołaniu przez API, omijając jakiekolwiek wymagania dotyczące potwierdzenia.

## Pola odpowiedzi

### Odpowiedź pola niestandardowego

| Pole | Typ | Opis |
|-------|------|-------------|
| `id` | String! | Unikalny identyfikator dla pola niestandardowego |
| `name` | String! | Nazwa wyświetlana przycisku |
| `type` | CustomFieldType! | Zawsze `BUTTON` dla pól przycisków |
| `buttonType` | String | Ustawienie zachowania potwierdzenia |
| `buttonConfirmText` | String | Wymagany tekst potwierdzenia (jeśli używasz twardego potwierdzenia) |
| `description` | String | Tekst pomocy dla użytkowników |
| `required` | Boolean! | Czy pole jest wymagane |
| `isActive` | Boolean! | Czy pole jest obecnie aktywne |
| `projectId` | String! | ID projektu, do którego należy to pole |
| `createdAt` | DateTime! | Kiedy pole zostało utworzone |
| `updatedAt` | DateTime! | Kiedy pole zostało ostatnio zmodyfikowane |

## Jak działają pola przycisków

### Integracja z automatyzacją

Pola przycisków są zaprojektowane do współpracy z systemem automatyzacji Blue:

1. **Utwórz pole przycisku** za pomocą powyższej mutacji
2. **Skonfiguruj automatyzacje**, które nasłuchują na zdarzenia `CUSTOM_FIELD_BUTTON_CLICKED`
3. **Użytkownicy klikają przycisk** w UI
4. **Automatyzacje wykonują** skonfigurowane akcje

### Przepływ zdarzeń

Gdy przycisk zostanie kliknięty:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Brak przechowywania danych

Ważne: Pola przycisków nie przechowują żadnych wartości danych. Służą wyłącznie jako wyzwalacze akcji. Każde kliknięcie:
- Generuje zdarzenie
- Uruchamia związane automatyzacje
- Rejestruje akcję w historii zadania
- Nie modyfikuje żadnej wartości pola

## Wymagane uprawnienia

Użytkownicy potrzebują odpowiednich ról projektowych, aby tworzyć i używać pól przycisków:

| Akcja | Wymagana rola |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Odpowiedzi błędów

### Odrzucone uprawnienia
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Pole niestandardowe nie znalezione
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

**Uwaga**: API nie zwraca konkretnych błędów dla brakujących automatyzacji lub niezgodności potwierdzenia.

## Najlepsze praktyki

### Konwencje nazewnictwa
- Używaj nazw zorientowanych na akcje: "Wyślij fakturę", "Utwórz raport", "Powiadom zespół"
- Bądź konkretny co do tego, co robi przycisk
- Unikaj ogólnych nazw, takich jak "Przycisk 1" lub "Kliknij tutaj"

### Ustawienia potwierdzenia
- Pozostaw `buttonType` puste dla bezpiecznych, odwracalnych działań
- Ustaw `buttonType`, aby zasugerować zachowanie potwierdzenia klientom UI
- Użyj `buttonConfirmText`, aby określić, co użytkownicy powinni wpisać w potwierdzeniach UI
- Pamiętaj: To są tylko wskazówki UI - wywołania API zawsze wykonują się natychmiast

### Projektowanie automatyzacji
- Skup się na pojedynczym przepływie pracy w akcjach przycisków
- Zapewnij jasną informację zwrotną na temat tego, co się wydarzyło po kliknięciu
- Rozważ dodanie tekstu opisu, aby wyjaśnić cel przycisku

## Typowe przypadki użycia

1. **Przejścia w przepływie pracy**
   - "Oznacz jako zakończone"
   - "Wyślij do zatwierdzenia"
   - "Archiwizuj zadanie"

2. **Integracje zewnętrzne**
   - "Synchronizuj z CRM"
   - "Generuj fakturę"
   - "Wyślij aktualizację e-mail"

3. **Operacje wsadowe**
   - "Aktualizuj wszystkie podzadania"
   - "Kopiuj do projektów"
   - "Zastosuj szablon"

4. **Akcje raportowe**
   - "Generuj raport"
   - "Eksportuj dane"
   - "Utwórz podsumowanie"

## Ograniczenia

- Przycisk nie może przechowywać ani wyświetlać wartości danych
- Każdy przycisk może tylko uruchamiać automatyzacje, a nie bezpośrednie wywołania API (jednak automatyzacje mogą zawierać akcje żądań HTTP do wywoływania zewnętrznych API lub własnych API Blue)
- Widoczność przycisku nie może być kontrolowana warunkowo
- Maksymalnie jedno wykonanie automatyzacji na kliknięcie (choć ta automatyzacja może uruchomić wiele akcji)

## Powiązane zasoby

- [API automatyzacji](/api/automations/index) - Konfiguruj akcje uruchamiane przez przyciski
- [Przegląd pól niestandardowych](/custom-fields/list-custom-fields) - Ogólne koncepcje pól niestandardowych