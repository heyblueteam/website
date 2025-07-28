---
title: Pole niestandardowe waluty
description: Tworzenie pól walutowych do śledzenia wartości pieniężnych z odpowiednim formatowaniem i walidacją
---

Pola niestandardowe waluty pozwalają na przechowywanie i zarządzanie wartościami pieniężnymi z powiązanymi kodami walut. Pole obsługuje 72 różne waluty, w tym główne waluty fiat i kryptowaluty, z automatycznym formatowaniem i opcjonalnymi ograniczeniami min/max.

## Podstawowy przykład

Utwórz proste pole walutowe:

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Zaawansowany przykład

Utwórz pole walutowe z ograniczeniami walidacyjnymi:

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola walutowego |
| `type` | CustomFieldType! | ✅ Tak | Musi być `CURRENCY` |
| `currency` | String | Nie | Domyślny kod waluty (3-literowy kod ISO) |
| `min` | Float | Nie | Minimalna dozwolona wartość (przechowywana, ale nie egzekwowana przy aktualizacjach) |
| `max` | Float | Nie | Maksymalna dozwolona wartość (przechowywana, ale nie egzekwowana przy aktualizacjach) |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

**Uwaga**: Kontekst projektu jest automatycznie określany na podstawie Twojej autoryzacji. Musisz mieć dostęp do projektu, w którym tworzysz pole.

## Ustawianie wartości waluty

Aby ustawić lub zaktualizować wartość waluty w rekordzie:

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola walutowego |
| `number` | Float! | ✅ Tak | Kwota pieniężna |
| `currency` | String! | ✅ Tak | 3-literowy kod waluty |

## Tworzenie rekordów z wartościami walutowymi

Podczas tworzenia nowego rekordu z wartościami walutowymi:

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### Format wejściowy dla Create

Podczas tworzenia rekordów wartości walutowe są przekazywane w inny sposób:

| Parametr | Typ | Opis |
|----------|-----|------|
| `customFieldId` | String! | ID pola walutowego |
| `value` | String! | Kwota jako ciąg (np. "1500.50") |
| `currency` | String! | 3-literowy kod waluty |

## Obsługiwane waluty

Blue obsługuje 72 waluty, w tym 70 walut fiat i 2 kryptowaluty:

### Waluty fiat

#### Ameryki
| Waluta | Kod | Nazwa |
|--------|-----|-------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Wenezuelski Bolívar Soberano |
| Bolivijski Boliviano | `BOB` | Bolivijski Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europa
| Waluta | Kod | Nazwa |
|--------|-----|-------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Korona norweska | `NOK` | Korona norweska |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Azja-Pacyfik
| Waluta | Kod | Nazwa |
|--------|-----|-------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### Bliski Wschód i Afryka
| Waluta | Kod | Nazwa |
|--------|-----|-------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Kryptowaluty
| Waluta | Kod |
|--------|-----|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja niestandardowego pola |
| `number` | Float | Kwota pieniężna |
| `currency` | String | 3-literowy kod waluty |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

## Formatowanie waluty

System automatycznie formatuje wartości walutowe na podstawie lokalizacji:

- **Umiejscowienie symbolu**: Poprawnie umieszcza symbole walut (przed/po)
- **Separator dziesiętny**: Używa separatorów specyficznych dla lokalizacji (. lub ,)
- **Separatory tysięcy**: Zastosowuje odpowiednie grupowanie
- **Miejsca dziesiętne**: Wyświetla 0-2 miejsca dziesiętne w zależności od kwoty
- **Specjalne traktowanie**: USD/CAD pokazują prefiks kodu waluty dla jasności

### Przykłady formatowania

| Wartość | Waluta | Wyświetlanie |
|---------|--------|--------------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Walidacja

### Walidacja kwoty
- Musi być poprawną liczbą
- Ograniczenia min/max są przechowywane z definicją pola, ale nie egzekwowane podczas aktualizacji wartości
- Obsługuje do 2 miejsc dziesiętnych do wyświetlania (pełna precyzja przechowywana wewnętrznie)

### Walidacja kodu waluty
- Musi być jednym z 72 obsługiwanych kodów walut
- Wrażliwe na wielkość liter (używaj wielkich liter)
- Nieprawidłowe kody zwracają błąd

## Funkcje integracji

### Formuły
Pola walutowe mogą być używane w niestandardowych polach FORMULA do obliczeń:
- Sumowanie wielu pól walutowych
- Obliczanie procentów
- Wykonywanie operacji arytmetycznych

### Konwersja walut
Użyj pól CURRENCY_CONVERSION, aby automatycznie konwertować między walutami (zobacz [Pola konwersji walut](/api/custom-fields/currency-conversion))

### Automatyzacje
Wartości walutowe mogą wywoływać automatyzacje na podstawie:
- Progów kwotowych
- Typu waluty
- Zmian wartości

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|----------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Uwaga**: Chociaż każdy członek projektu może tworzyć niestandardowe pola, możliwość ustawiania wartości zależy od uprawnień opartych na roli skonfigurowanych dla każdego pola.

## Odpowiedzi na błędy

### Nieprawidłowa wartość waluty
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

Ten błąd występuje, gdy:
- Kod waluty nie jest jednym z 72 obsługiwanych kodów
- Format liczby jest nieprawidłowy
- Wartość nie może być poprawnie sparsowana

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

### Wybór waluty
- Ustaw domyślną walutę, która odpowiada Twojemu głównemu rynkowi
- Używaj kodów walut ISO 4217 konsekwentnie
- Weź pod uwagę lokalizację użytkownika przy wyborze domyślnych ustawień

### Ograniczenia wartości
- Ustaw rozsądne wartości min/max, aby zapobiec błędom wprowadzania danych
- Użyj 0 jako minimum dla pól, które nie powinny być ujemne
- Weź pod uwagę swój przypadek użycia przy ustawianiu maksymalnych wartości

### Projekty wielowalutowe
- Użyj spójnej waluty bazowej do raportowania
- Wdrażaj pola CURRENCY_CONVERSION do automatycznej konwersji
- Dokumentuj, która waluta powinna być używana dla każdego pola

## Typowe przypadki użycia

1. **Budżetowanie projektu**
   - Śledzenie budżetu projektu
   - Szacowanie kosztów
   - Śledzenie wydatków

2. **Sprzedaż i umowy**
   - Wartości umów
   - Kwoty kontraktów
   - Śledzenie przychodów

3. **Planowanie finansowe**
   - Kwoty inwestycji
   - Rundy finansowania
   - Cele finansowe

4. **Międzynarodowy biznes**
   - Ceny w wielu walutach
   - Śledzenie wymiany walut
   - Transakcje transgraniczne

## Ograniczenia

- Maksymalnie 2 miejsca dziesiętne do wyświetlania (choć więcej precyzji jest przechowywane)
- Brak wbudowanej konwersji walut w standardowych polach CURRENCY
- Nie można łączyć walut w jednej wartości pola
- Brak automatycznych aktualizacji kursów wymiany (użyj CURRENCY_CONVERSION do tego)
- Symbole walut nie są dostosowywalne

## Powiązane zasoby

- [Pola konwersji walut](/api/custom-fields/currency-conversion) - Automatyczna konwersja walut
- [Pola liczbowe](/api/custom-fields/number) - Dla wartości numerycznych niepieniężnych
- [Pola formuł](/api/custom-fields/formula) - Obliczanie z wartościami walutowymi
- [Niestandardowe pola listy](/api/custom-fields/list-custom-fields) - Zapytaj o wszystkie niestandardowe pola w projekcie