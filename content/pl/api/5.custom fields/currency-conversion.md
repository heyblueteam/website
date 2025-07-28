---
title: Niestandardowe pole konwersji walut
description: Twórz pola, które automatycznie konwertują wartości walutowe przy użyciu kursów wymiany w czasie rzeczywistym
---

Niestandardowe pola konwersji walut automatycznie konwertują wartości z pola CURRENCY źródłowego na różne waluty docelowe, korzystając z kursów wymiany w czasie rzeczywistym. Pola te aktualizują się automatycznie, gdy tylko zmienia się wartość waluty źródłowej.

Kursy konwersji są dostarczane przez [Frankfurter API](https://github.com/hakanensari/frankfurter), otwartą usługę, która śledzi referencyjne kursy wymiany publikowane przez [Europejski Bank Centralny](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Zapewnia to dokładne, wiarygodne i aktualne konwersje walutowe dla Twoich międzynarodowych potrzeb biznesowych.

## Podstawowy przykład

Utwórz proste pole konwersji walut:

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## Zaawansowany przykład

Utwórz pole konwersji z określoną datą dla kursów historycznych:

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## Proces pełnej konfiguracji

Ustawienie pola konwersji walut wymaga trzech kroków:

### Krok 1: Utwórz pole CURRENCY źródłowe

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### Krok 2: Utwórz pole CURRENCY_CONVERSION

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### Krok 3: Utwórz opcje konwersji

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola konwersji |
| `type` | CustomFieldType! | ✅ Tak | Musi być `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | Nie | ID źródłowego pola CURRENCY, z którego ma być dokonana konwersja |
| `conversionDateType` | String | Nie | Strategia daty dla kursów wymiany (patrz poniżej) |
| `conversionDate` | String | Nie | Ciąg daty dla konwersji (na podstawie conversionDateType) |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

**Uwaga**: Niestandardowe pola są automatycznie powiązane z projektem na podstawie aktualnego kontekstu projektu użytkownika. Żaden parametr `projectId` nie jest wymagany.

### Typy dat konwersji

| Typ | Opis | Parametr conversionDate |
|-----|------|-------------------------|
| `currentDate` | Używa kursów wymiany w czasie rzeczywistym | Nie wymagany |
| `specificDate` | Używa kursów z ustalonej daty | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Używa daty z innego pola | "todoDueDate" or DATE field ID |

## Tworzenie opcji konwersji

Opcje konwersji definiują, które pary walutowe mogą być konwertowane:

### CreateCustomFieldOptionInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `customFieldId` | String! | ✅ Tak | ID pola CURRENCY_CONVERSION |
| `title` | String! | ✅ Tak | Nazwa wyświetlana dla tej opcji konwersji |
| `currencyConversionFrom` | String! | ✅ Tak | Kod waluty źródłowej lub "Dowolna" |
| `currencyConversionTo` | String! | ✅ Tak | Kod waluty docelowej |

### Używanie "Dowolna" jako źródła

Specjalna wartość "Dowolna" jako `currencyConversionFrom` tworzy opcję zapasową:

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

Ta opcja będzie używana, gdy nie zostanie znalezione konkretne dopasowanie pary walutowej.

## Jak działa automatyczna konwersja

1. **Aktualizacja wartości**: Gdy wartość jest ustawiana w źródłowym polu CURRENCY
2. **Dopasowanie opcji**: System znajduje dopasowaną opcję konwersji na podstawie waluty źródłowej
3. **Pobieranie kursu**: Pobiera kurs wymiany z Frankfurter API
4. **Obliczenie**: Mnoży kwotę źródłową przez kurs wymiany
5. **Przechowywanie**: Zapisuje przekonwertowaną wartość z kodem waluty docelowej

### Przykład przepływu

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## Konwersje oparte na dacie

### Używanie bieżącej daty

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

Konwersje aktualizują się z bieżącymi kursami wymiany za każdym razem, gdy zmienia się wartość źródłowa.

### Używanie określonej daty

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

Zawsze używa kursów wymiany z określonej daty.

### Używanie daty z pola

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

Używa daty z innego pola (czy to daty wykonania zadania, czy pola DATE).

## Pola odpowiedzi

### TodoCustomField Response

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja pola konwersji |
| `number` | Float | Przekonwertowana kwota |
| `currency` | String | Kod waluty docelowej |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zaktualizowana |

## Źródło kursu wymiany

Blue korzysta z **Frankfurter API** do kursów wymiany:
- Otwarte API hostowane przez Europejski Bank Centralny
- Aktualizacje codziennie z oficjalnymi kursami wymiany
- Obsługuje kursy historyczne od 1999 roku
- Darmowe i wiarygodne do użytku biznesowego

## Obsługa błędów

### Niepowodzenia konwersji

Gdy konwersja nie powiedzie się (błąd API, nieprawidłowa waluta itp.):
- Przekonwertowana wartość jest ustawiana na `0`
- Kod waluty docelowej jest nadal przechowywany
- Żaden błąd nie jest zgłaszany użytkownikowi

### Typowe scenariusze

| Scenariusz | Wynik |
|------------|-------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Brak dopasowanej opcji | Uses "Any" option if available |
| Missing source value | Nie wykonano konwersji |

## Wymagane uprawnienia

Zarządzanie niestandardowymi polami wymaga dostępu na poziomie projektu:

| Rola | Może tworzyć/aktualizować pola |
|------|-------------------------------|
| `OWNER` | ✅ Tak |
| `ADMIN` | ✅ Tak |
| `MEMBER` | ❌ Nie |
| `CLIENT` | ❌ Nie |

Uprawnienia do przeglądania wartości przekonwertowanych podlegają standardowym zasadom dostępu do rekordów.

## Najlepsze praktyki

### Konfiguracja opcji
- Twórz konkretne pary walutowe dla powszechnych konwersji
- Dodaj opcję zapasową "Dowolna" dla elastyczności
- Używaj opisowych tytułów dla opcji

### Wybór strategii daty
- Używaj `currentDate` do bieżącego śledzenia finansowego
- Używaj `specificDate` do raportowania historycznego
- Używaj `fromDateField` do kursów specyficznych dla transakcji

### Rozważania dotyczące wydajności
- Wiele pól konwersji aktualizuje się równolegle
- Wywołania API są wykonywane tylko wtedy, gdy zmienia się wartość źródłowa
- Konwersje w tej samej walucie pomijają wywołania API

## Typowe przypadki użycia

1. **Projekty wielowalutowe**
   - Śledzenie kosztów projektu w lokalnych walutach
   - Raportowanie całkowitego budżetu w walucie firmy
   - Porównywanie wartości w różnych regionach

2. **Sprzedaż międzynarodowa**
   - Konwersja wartości transakcji na walutę raportową
   - Śledzenie przychodów w wielu walutach
   - Konwersja historyczna dla zamkniętych transakcji

3. **Raportowanie finansowe**
   - Konwersje walut na koniec okresu
   - Skonsolidowane sprawozdania finansowe
   - Budżet w porównaniu do rzeczywistości w lokalnej walucie

4. **Zarządzanie umowami**
   - Konwersja wartości umów w dniu podpisania
   - Śledzenie harmonogramów płatności w wielu walutach
   - Ocena ryzyka walutowego

## Ograniczenia

- Brak wsparcia dla konwersji kryptowalut
- Nie można ręcznie ustawiać wartości przekonwertowanych (zawsze obliczane)
- Stała precyzja 2 miejsc po przecinku dla wszystkich przekonwertowanych kwot
- Brak wsparcia dla niestandardowych kursów wymiany
- Brak buforowania kursów wymiany (świeże wywołanie API dla każdej konwersji)
- Zależy od dostępności Frankfurter API

## Powiązane zasoby

- [Pola walutowe](/api/custom-fields/currency) - Pola źródłowe dla konwersji
- [Pola daty](/api/custom-fields/date) - Do konwersji opartych na dacie
- [Pola formuły](/api/custom-fields/formula) - Alternatywne obliczenia
- [Przegląd pól niestandardowych](/custom-fields/list-custom-fields) - Ogólne koncepcje