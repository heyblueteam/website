---
title: Pole niestandardowe telefonu
description: Twórz pola telefoniczne do przechowywania i walidacji numerów telefonów w formacie międzynarodowym
---

Pola niestandardowe telefonu pozwalają na przechowywanie numerów telefonów w rekordach z wbudowaną walidacją i międzynarodowym formatowaniem. Są idealne do śledzenia informacji kontaktowych, kontaktów awaryjnych lub wszelkich danych związanych z telefonem w Twoich projektach.

## Podstawowy przykład

Utwórz proste pole telefoniczne:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole telefoniczne z opisem:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola telefonicznego |
| `type` | CustomFieldType! | ✅ Tak | Musi być `PHONE` |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

**Uwaga**: Pola niestandardowe są automatycznie powiązane z projektem na podstawie aktualnego kontekstu projektu użytkownika. Żaden parametr `projectId` nie jest wymagany.

## Ustawianie wartości telefonu

Aby ustawić lub zaktualizować wartość telefonu w rekordzie:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do aktualizacji |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego telefonu |
| `text` | String | Nie | Numer telefonu z kodem kraju |
| `regionCode` | String | Nie | Kod kraju (automatycznie wykrywany) |

**Uwaga**: Chociaż `text` jest opcjonalny w schemacie, numer telefonu jest wymagany, aby pole miało sens. Podczas korzystania z `setTodoCustomField` nie jest przeprowadzana walidacja - możesz przechowywać dowolną wartość tekstową i regionCode. Automatyczne wykrywanie odbywa się tylko podczas tworzenia rekordu.

## Tworzenie rekordów z wartościami telefonu

Podczas tworzenia nowego rekordu z wartościami telefonu:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
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
| `text` | String | Sformatowany numer telefonu (format międzynarodowy) |
| `regionCode` | String | Kod kraju (np. "US", "GB", "CA") |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

## Walidacja numeru telefonu

**Ważne**: Walidacja i formatowanie numeru telefonu odbywa się tylko podczas tworzenia nowych rekordów za pomocą `createTodo`. Podczas aktualizacji istniejących wartości telefonów za pomocą `setTodoCustomField` nie jest przeprowadzana walidacja, a wartości są przechowywane tak, jak zostały podane.

### Akceptowane formaty (podczas tworzenia rekordu)
Numery telefonów muszą zawierać kod kraju w jednym z tych formatów:

- **Format E.164 (preferowany)**: `+12345678900`
- **Format międzynarodowy**: `+1 234 567 8900`
- **Międzynarodowy z interpunkcją**: `+1 (234) 567-8900`
- **Kod kraju z myślnikami**: `+1-234-567-8900`

**Uwaga**: Krajowe formaty bez kodu kraju (jak `(234) 567-8900`) będą odrzucane podczas tworzenia rekordu.

### Zasady walidacji (podczas tworzenia rekordu)
- Używa libphonenumber-js do analizy i walidacji
- Akceptuje różne międzynarodowe formaty numerów telefonów
- Automatycznie wykrywa kraj na podstawie numeru
- Formatuje numer w międzynarodowym formacie wyświetlania (np. `+1 234 567 8900`)
- Wyodrębnia i przechowuje kod kraju osobno (np. `US`)

### Przykłady poprawnych numerów telefonów
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Przykłady niepoprawnych numerów telefonów
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Format przechowywania

Podczas tworzenia rekordów z numerami telefonów:
- **text**: Przechowywane w formacie międzynarodowym (np. `+1 234 567 8900`) po walidacji
- **regionCode**: Przechowywane jako kod kraju ISO (np. `US`, `GB`, `CA`) automatycznie wykrywane

Podczas aktualizacji za pomocą `setTodoCustomField`:
- **text**: Przechowywane dokładnie tak, jak podano (bez formatowania)
- **regionCode**: Przechowywane dokładnie tak, jak podano (bez walidacji)

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|----------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Odpowiedzi błędów

### Niepoprawny format telefonu
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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

### Brak kodu kraju
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Najlepsze praktyki

### Wprowadzanie danych
- Zawsze dołączaj kod kraju w numerach telefonów
- Używaj formatu E.164 dla spójności
- Waliduj numery przed przechowywaniem dla ważnych operacji
- Weź pod uwagę preferencje regionalne dla formatowania wyświetlania

### Jakość danych
- Przechowuj numery w formacie międzynarodowym dla globalnej kompatybilności
- Używaj regionCode dla funkcji specyficznych dla kraju
- Waliduj numery telefonów przed krytycznymi operacjami (SMS, połączenia)
- Weź pod uwagę implikacje strefy czasowej dla czasu kontaktu

### Rozważania międzynarodowe
- Kod kraju jest automatycznie wykrywany i przechowywany
- Numery są formatowane w standardzie międzynarodowym
- Preferencje wyświetlania regionalnego mogą korzystać z regionCode
- Weź pod uwagę lokalne konwencje wybierania podczas wyświetlania

## Typowe przypadki użycia

1. **Zarządzanie kontaktami**
   - Numery telefonów klientów
   - Informacje kontaktowe dostawców
   - Numery telefonów członków zespołu
   - Szczegóły kontaktowe wsparcia

2. **Kontakty awaryjne**
   - Numery kontaktowe w sytuacjach awaryjnych
   - Informacje kontaktowe na dyżurze
   - Kontakty do odpowiedzi kryzysowej
   - Numery telefonów do eskalacji

3. **Wsparcie klienta**
   - Numery telefonów klientów
   - Numery telefonów do oddzwonienia wsparcia
   - Numery telefonów do weryfikacji
   - Numery telefonów do kontaktów follow-up

4. **Sprzedaż i marketing**
   - Numery telefonów leadów
   - Listy kontaktowe kampanii
   - Informacje kontaktowe partnerów
   - Numery telefonów źródeł poleceń

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje, gdy pola telefoniczne są aktualizowane
- Wysyłaj powiadomienia SMS na przechowywane numery telefonów
- Twórz zadania follow-up na podstawie zmian w telefonach
- Kieruj połączenia na podstawie danych o numerach telefonów

### Z wyszukiwaniami
- Odwołuj się do danych telefonicznych z innych rekordów
- Agreguj listy telefonów z wielu źródeł
- Znajduj rekordy według numeru telefonu
- Krzyżowo odniesienia informacje kontaktowe

### Z formularzami
- Automatyczna walidacja telefonów
- Sprawdzanie formatu międzynarodowego
- Wykrywanie kodu kraju
- Informacje zwrotne w czasie rzeczywistym dotyczące formatu

## Ograniczenia

- Wymaga kodu kraju dla wszystkich numerów
- Brak wbudowanych możliwości SMS lub połączeń
- Brak weryfikacji numerów telefonów poza sprawdzaniem formatu
- Brak przechowywania metadanych telefonu (operator, typ itp.)
- Krajowe numery formatów bez kodu kraju są odrzucane
- Brak automatycznego formatowania numerów telefonów w UI poza standardem międzynarodowym

## Powiązane zasoby

- [Pola tekstowe](/api/custom-fields/text-single) - Dla danych tekstowych, które nie są telefonami
- [Pola e-mailowe](/api/custom-fields/email) - Dla adresów e-mail
- [Pola URL](/api/custom-fields/url) - Dla adresów stron internetowych
- [Przegląd pól niestandardowych](/custom-fields/list-custom-fields) - Ogólne koncepcje