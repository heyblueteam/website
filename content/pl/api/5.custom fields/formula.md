---
title: Pole niestandardowe formuły
description: Twórz pola obliczeniowe, które automatycznie obliczają wartości na podstawie innych danych
---

Pola niestandardowe formuły są używane do obliczeń w wykresach i pulpitach nawigacyjnych w Blue. Definiują funkcje agregacji (SUMA, ŚREDNIA, LICZBA itp.), które działają na danych pól niestandardowych, aby wyświetlać obliczone metryki w wykresach. Formuły nie są obliczane na poziomie pojedynczego zadania, lecz agregują dane z wielu rekordów w celach wizualizacji.

## Podstawowy przykład

Utwórz pole formuły do obliczeń wykresu:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Zaawansowany przykład

Utwórz formułę walutową złożonymi obliczeniami:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola formuły |
| `type` | CustomFieldType! | ✅ Tak | Musi być `FORMULA` |
| `projectId` | String! | ✅ Tak | ID projektu, w którym to pole zostanie utworzone |
| `formula` | JSON | Nie | Definicja formuły do obliczeń wykresów |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

### Struktura formuły

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Obsługiwane funkcje

### Funkcje agregacji wykresów

Pola formuły obsługują następujące funkcje agregacji do obliczeń wykresów:

| Funkcja | Opis | Enum ChartFunction |
|---------|------|-------------------|
| `SUM` | Suma wszystkich wartości | `SUM` |
| `AVERAGE` | Średnia wartości numerycznych | `AVERAGE` |
| `AVERAGEA` | Średnia z wyłączeniem zer i wartości null | `AVERAGEA` |
| `COUNT` | Liczba wartości | `COUNT` |
| `COUNTA` | Liczba z wyłączeniem zer i wartości null | `COUNTA` |
| `MAX` | Wartość maksymalna | `MAX` |
| `MIN` | Wartość minimalna | `MIN` |

**Uwaga**: Te funkcje są używane w polu `display.function` i działają na danych agregowanych do wizualizacji wykresów. Złożone wyrażenia matematyczne lub obliczenia na poziomie pól nie są obsługiwane.

## Typy wyświetlania

### Wyświetlanie liczby

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Wynik: `1250.75`

### Wyświetlanie waluty

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Wynik: `$1,250.75`

### Wyświetlanie procentów

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Wynik: `87.5%`

## Edytowanie pól formuły

Zaktualizuj istniejące pola formuły:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Przetwarzanie formuły

### Kontekst obliczeń wykresu

Pola formuły są przetwarzane w kontekście segmentów wykresu i pulpitów nawigacyjnych:
- Obliczenia odbywają się, gdy wykresy są renderowane lub aktualizowane
- Wyniki są przechowywane w `ChartSegment.formulaResult` jako wartości dziesiętne
- Przetwarzanie odbywa się przez dedykowaną kolejkę BullMQ o nazwie 'formula'
- Aktualizacje publikują do subskrybentów pulpitów nawigacyjnych w czasie rzeczywistym

### Formatowanie wyświetlania

Funkcja `getFormulaDisplayValue` formatuje obliczone wyniki w zależności od typu wyświetlania:
- **LICZBA**: Wyświetla jako zwykła liczba z opcjonalną precyzją
- **PROCENT**: Dodaje sufiks % z opcjonalną precyzją  
- **WALUTA**: Formatuje przy użyciu określonego kodu waluty

## Przechowywanie wyników formuły

Wyniki są przechowywane w polu `formulaResult`:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja pola formuły |
| `number` | Float | Obliczony wynik liczbowy |
| `formulaResult` | JSON | Pełny wynik z formatowaniem wyświetlania |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio obliczona |

## Kontekst danych

### Źródło danych wykresu

Pola formuły działają w kontekście źródła danych wykresu:
- Formuły agregują wartości pól niestandardowych w zadaniach w projekcie
- Funkcja agregacji określona w `display.function` określa obliczenia
- Wyniki są obliczane przy użyciu funkcji agregacji SQL (avg, sum, count itp.)
- Obliczenia są wykonywane na poziomie bazy danych dla efektywności

## Przykłady formuł

### Całkowity budżet (wyświetlanie wykresu)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Średni wynik (wyświetlanie wykresu)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Liczba zadań (wyświetlanie wykresu)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Wymagane uprawnienia

Operacje na polach niestandardowych podlegają standardowym uprawnieniom opartym na rolach:

| Akcja | Wymagana rola |
|-------|---------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Uwaga**: Wymagane konkretne role zależą od konfiguracji ról niestandardowych w Twoim projekcie. Nie ma specjalnych stałych uprawnień, takich jak CUSTOM_FIELDS_CREATE.

## Obsługa błędów

### Błąd walidacji
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Pole niestandardowe nie znalezione
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

### Projektowanie formuły
- Używaj jasnych, opisowych nazw dla pól formuły
- Dodaj opisy wyjaśniające logikę obliczeń
- Testuj formuły na przykładowych danych przed wdrożeniem
- Utrzymuj formuły proste i czytelne

### Optymalizacja wydajności
- Unikaj głęboko zagnieżdżonych zależności formuł
- Używaj konkretnych odniesień do pól zamiast symboli wieloznacznych
- Rozważ strategie pamięci podręcznej dla złożonych obliczeń
- Monitoruj wydajność formuł w dużych projektach

### Jakość danych
- Waliduj dane źródłowe przed użyciem w formułach
- Obsługuj puste lub wartości null w odpowiedni sposób
- Używaj odpowiedniej precyzji dla typów wyświetlania
- Rozważ przypadki brzegowe w obliczeniach

## Typowe przypadki użycia

1. **Śledzenie finansowe**
   - Obliczenia budżetu
   - Sprawozdania zysków/strat
   - Analiza kosztów
   - Prognozy przychodów

2. **Zarządzanie projektami**
   - Procenty ukończenia
   - Wykorzystanie zasobów
   - Obliczenia harmonogramu
   - Metryki wydajności

3. **Kontrola jakości**
   - Średnie wyniki
   - Wskaźniki zaliczeń/niezaliczeń
   - Metryki jakości
   - Śledzenie zgodności

4. **Inteligencja biznesowa**
   - Obliczenia KPI
   - Analiza trendów
   - Metryki porównawcze
   - Wartości pulpitów nawigacyjnych

## Ograniczenia

- Formuły są przeznaczone wyłącznie do agregacji wykresów/pulpitów nawigacyjnych, a nie do obliczeń na poziomie zadań
- Ograniczone do siedmiu obsługiwanych funkcji agregacji (SUMA, ŚREDNIA itp.)
- Brak złożonych wyrażeń matematycznych lub obliczeń między polami
- Nie można odwoływać się do wielu pól w jednej formule
- Wyniki są widoczne tylko w wykresach i pulpitach nawigacyjnych
- Pole `logic` jest przeznaczone tylko do tekstu wyświetlanego, a nie do rzeczywistej logiki obliczeniowej

## Powiązane zasoby

- [Pola liczbowe](/api/5.custom%20fields/number) - Dla statycznych wartości liczbowych
- [Pola walutowe](/api/5.custom%20fields/currency) - Dla wartości pieniężnych
- [Pola referencyjne](/api/5.custom%20fields/reference) - Dla danych międzyprojektowych
- [Pola wyszukiwania](/api/5.custom%20fields/lookup) - Dla danych agregowanych
- [Przegląd pól niestandardowych](/api/5.custom%20fields/2.list-custom-fields) - Ogólne koncepcje