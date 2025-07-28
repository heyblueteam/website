---
title: Niestandardowe pole e-mail
description: Twórz pola e-mail do przechowywania i walidacji adresów e-mail
---

Niestandardowe pola e-mail pozwalają na przechowywanie adresów e-mail w rekordach z wbudowaną walidacją. Są idealne do śledzenia informacji kontaktowych, adresów e-mail przypisanych użytkowników lub wszelkich danych związanych z e-mailem w Twoich projektach.

## Podstawowy przykład

Utwórz proste pole e-mail:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole e-mail z opisem:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola e-mail |
| `type` | CustomFieldType! | ✅ Tak | Musi być `EMAIL` |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

## Ustawianie wartości e-mail

Aby ustawić lub zaktualizować wartość e-mail w rekordzie:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola e-mail |
| `text` | String | Nie | Adres e-mail do przechowania |

## Tworzenie rekordów z wartościami e-mail

Podczas tworzenia nowego rekordu z wartościami e-mail:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź CustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | Unikalny identyfikator dla niestandardowego pola |
| `name` | String! | Nazwa wyświetlana pola e-mail |
| `type` | CustomFieldType! | Typ pola (EMAIL) |
| `description` | String | Tekst pomocniczy dla pola |
| `value` | JSON | Zawiera wartość e-mail (patrz poniżej) |
| `createdAt` | DateTime! | Kiedy pole zostało utworzone |
| `updatedAt` | DateTime! | Kiedy pole zostało ostatnio zmodyfikowane |

**Ważne**: Wartości e-mail są dostępne przez pole `customField.value.text`, a nie bezpośrednio w odpowiedzi.

## Zapytania o wartości e-mail

Podczas zapytań o rekordy z niestandardowymi polami e-mail, uzyskaj dostęp do e-maila przez ścieżkę `customField.value.text`:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

Odpowiedź będzie zawierać e-mail w zagnieżdżonej strukturze:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## Walidacja e-mail

### Walidacja formularza
Gdy pola e-mail są używane w formularzach, automatycznie walidują format e-mail:
- Używa standardowych zasad walidacji e-mail
- Usuwa białe znaki z wejścia
- Odrzuca nieprawidłowe formaty e-mail

### Zasady walidacji
- Musi zawierać symbol `@`
- Musi mieć prawidłowy format domeny
- Białe znaki na początku/końcu są automatycznie usuwane
- Akceptowane są powszechne formaty e-mail

### Przykłady prawidłowych e-maili
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Przykłady nieprawidłowych e-maili
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Ważne uwagi

### Bezpośrednie API vs formularze
- **Formularze**: Automatyczna walidacja e-mail jest stosowana
- **Bezpośrednie API**: Brak walidacji - można przechować dowolny tekst
- **Zalecenie**: Używaj formularzy do wprowadzania danych przez użytkowników, aby zapewnić walidację

### Format przechowywania
- Adresy e-mail są przechowywane jako zwykły tekst
- Brak specjalnego formatowania lub analizy
- Wrażliwość na wielkość liter: niestandardowe pola e-mail są przechowywane z uwzględnieniem wielkości liter (w przeciwieństwie do e-maili autoryzacji użytkowników, które są normalizowane do małych liter)
- Brak ograniczeń długości poza ograniczeniami bazy danych (limit 16MB)

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Odpowiedzi błędów

### Nieprawidłowy format e-mail (tylko formularze)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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

### Wprowadzanie danych
- Zawsze waliduj adresy e-mail w swojej aplikacji
- Używaj pól e-mail tylko dla rzeczywistych adresów e-mail
- Rozważ użycie formularzy do wprowadzania danych przez użytkowników, aby uzyskać automatyczną walidację

### Jakość danych
- Usuń białe znaki przed przechowaniem
- Rozważ normalizację wielkości liter (zwykle małe litery)
- Waliduj format e-mail przed ważnymi operacjami

### Rozważania dotyczące prywatności
- Adresy e-mail są przechowywane jako zwykły tekst
- Rozważ przepisy dotyczące prywatności danych (GDPR, CCPA)
- Wprowadź odpowiednie kontrole dostępu

## Typowe przypadki użycia

1. **Zarządzanie kontaktami**
   - Adresy e-mail klientów
   - Informacje kontaktowe dostawców
   - Adresy e-mail członków zespołu
   - Szczegóły kontaktowe wsparcia

2. **Zarządzanie projektami**
   - Adresy e-mail interesariuszy
   - Adresy e-mail do zatwierdzeń
   - Odbiorcy powiadomień
   - Adresy e-mail zewnętrznych współpracowników

3. **Wsparcie klienta**
   - Adresy e-mail klientów
   - Kontakty do zgłoszeń wsparcia
   - Kontakty do eskalacji
   - Adresy e-mail do opinii

4. **Sprzedaż i marketing**
   - Adresy e-mail leadów
   - Listy kontaktowe kampanii
   - Informacje kontaktowe partnerów
   - Adresy e-mail źródeł poleceń

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje, gdy pola e-mail są aktualizowane
- Wysyłaj powiadomienia na zapisane adresy e-mail
- Twórz zadania do śledzenia na podstawie zmian e-mail

### Z wyszukiwaniami
- Odwołuj się do danych e-mail z innych rekordów
- Agreguj listy e-mail z wielu źródeł
- Znajdź rekordy według adresu e-mail

### Z formularzami
- Automatyczna walidacja e-mail
- Sprawdzanie formatu e-mail
- Usuwanie białych znaków

## Ograniczenia

- Brak wbudowanej weryfikacji lub walidacji e-mail poza sprawdzaniem formatu
- Brak funkcji UI specyficznych dla e-mail (np. klikalne linki e-mail)
- Przechowywane jako zwykły tekst bez szyfrowania
- Brak możliwości tworzenia lub wysyłania e-maili
- Brak przechowywania metadanych e-mail (nazwa wyświetlana itp.)
- Bezpośrednie wywołania API omijają walidację (tylko formularze walidują)

## Powiązane zasoby

- [Pola tekstowe](/api/custom-fields/text-single) - Do danych tekstowych, które nie są e-mailem
- [Pola URL](/api/custom-fields/url) - Do adresów stron internetowych
- [Pola telefoniczne](/api/custom-fields/phone) - Do numerów telefonów
- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne pojęcia