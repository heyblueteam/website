---
title: Unikalne pole identyfikacyjne
description: Twórz automatycznie generowane pola unikalnych identyfikatorów z sekwencyjnym numerowaniem i niestandardowym formatowaniem
---

Pola unikalnych identyfikatorów automatycznie generują sekwencyjne, unikalne identyfikatory dla Twoich rekordów. Są idealne do tworzenia numerów biletów, identyfikatorów zamówień, numerów faktur lub jakiegokolwiek systemu identyfikatorów sekwencyjnych w Twoim przepływie pracy.

## Podstawowy przykład

Utwórz proste pole unikalnego identyfikatora z automatycznym sekwencjonowaniem:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Zaawansowany przykład

Utwórz sformatowane pole unikalnego identyfikatora z prefiksem i zerowym wypełnieniem:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput (UNIQUE_ID)

| Parametr | Typ | Wymagane | Opis |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola unikalnego identyfikatora |
| `type` | CustomFieldType! | ✅ Tak | Musi być `UNIQUE_ID` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |
| `useSequenceUniqueId` | Boolean | Nie | Włącz automatyczne sekwencjonowanie (domyślnie: fałsz) |
| `prefix` | String | Nie | Prefiks tekstowy dla generowanych identyfikatorów (np. "ZADANIE-") |
| `sequenceDigits` | Int | Nie | Liczba cyfr do zerowego wypełnienia |
| `sequenceStartingNumber` | Int | Nie | Początkowa liczba dla sekwencji |

## Opcje konfiguracji

### Automatyczne sekwencjonowanie (`useSequenceUniqueId`)
- **prawda**: Automatycznie generuje sekwencyjne identyfikatory, gdy rekordy są tworzone
- **fałsz** lub **nieokreślone**: Wymagana ręczna wpisywanie (działa jak pole tekstowe)

### Prefiks (`prefix`)
- Opcjonalny prefiks tekstowy dodawany do wszystkich generowanych identyfikatorów
- Przykłady: "ZADANIE-", "ZAM-", "BŁĄD-", "WYM-"
- Brak ograniczeń długości, ale zachowaj rozsądne dla wyświetlania

### Cyfry sekwencji (`sequenceDigits`)
- Liczba cyfr do zerowego wypełnienia numeru sekwencji
- Przykład: `sequenceDigits: 3` produkuje `001`, `002`, `003`
- Jeśli nie określono, nie stosuje się wypełnienia

### Początkowa liczba (`sequenceStartingNumber`)
- Pierwsza liczba w sekwencji
- Przykład: `sequenceStartingNumber: 1000` zaczyna się od 1000, 1001, 1002...
- Jeśli nie określono, zaczyna się od 1 (domyślne zachowanie)

## Format generowanego identyfikatora

Ostateczny format identyfikatora łączy wszystkie opcje konfiguracji:

```
{prefix}{paddedSequenceNumber}
```

### Przykłady formatów

| Konfiguracja | Generowane identyfikatory |
|---------------|---------------|
| Brak opcji | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Odczytywanie wartości unikalnych identyfikatorów

### Zapytanie rekordów z unikalnymi identyfikatorami
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Format odpowiedzi
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Automatyczne generowanie identyfikatorów

### Kiedy identyfikatory są generowane
- **Tworzenie rekordu**: Identyfikatory są automatycznie przypisywane, gdy nowe rekordy są tworzone
- **Dodawanie pola**: Podczas dodawania pola UNIQUE_ID do istniejących rekordów, zadanie w tle jest kolejkowane (implementacja robocza w toku)
- **Przetwarzanie w tle**: Generowanie identyfikatorów dla nowych rekordów odbywa się synchronicznie za pomocą wyzwalaczy bazy danych

### Proces generowania
1. **Wyzwalacz**: Nowy rekord jest tworzony lub pole UNIQUE_ID jest dodawane
2. **Wyszukiwanie sekwencji**: System znajduje następny dostępny numer sekwencji
3. **Przypisanie identyfikatora**: Numer sekwencji jest przypisywany do rekordu
4. **Aktualizacja licznika**: Licznik sekwencji jest inkrementowany dla przyszłych rekordów
5. **Formatowanie**: Identyfikator jest formatowany z prefiksem i wypełnieniem podczas wyświetlania

### Gwarancje unikalności
- **Ograniczenia bazy danych**: Unikalne ograniczenie na identyfikatory sekwencji w każdym polu
- **Operacje atomowe**: Generowanie sekwencji wykorzystuje blokady bazy danych, aby zapobiec duplikatom
- **Zakres projektu**: Sekwencje są niezależne dla każdego projektu
- **Ochrona przed warunkami wyścigu**: Równoległe żądania są obsługiwane bezpiecznie

## Tryb ręczny vs automatyczny

### Tryb automatyczny (`useSequenceUniqueId: true`)
- Identyfikatory są automatycznie generowane za pomocą wyzwalaczy bazy danych
- Gwarantowane jest sekwencyjne numerowanie
- Atomowe generowanie sekwencji zapobiega duplikatom
- Sformatowane identyfikatory łączą prefiks + wypełniony numer sekwencji

### Tryb ręczny (`useSequenceUniqueId: false` lub `undefined`)
- Działa jak zwykłe pole tekstowe
- Użytkownicy mogą wprowadzać niestandardowe wartości za pomocą `setTodoCustomField` z parametrem `text`
- Brak automatycznego generowania
- Brak egzekwowania unikalności poza ograniczeniami bazy danych

## Ustawianie wartości ręcznych (tylko tryb ręczny)

Gdy `useSequenceUniqueId` jest fałsz, możesz ustawić wartości ręcznie:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField (UNIQUE_ID)

| Pole | Typ | Opis |
|-------|------|-------------|
| `id` | String! | Unikalny identyfikator dla wartości pola |
| `customField` | CustomField! | Definicja pola niestandardowego |
| `sequenceId` | Int | Wygenerowany numer sekwencji (uzupełniony dla pól UNIQUE_ID) |
| `text` | String | Sformatowana wartość tekstowa (łączy prefiks + wypełnioną sekwencję) |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zaktualizowana |

### Odpowiedź CustomField (UNIQUE_ID)

| Pole | Typ | Opis |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | Czy automatyczne sekwencjonowanie jest włączone |
| `prefix` | String | Prefiks tekstowy dla generowanych identyfikatorów |
| `sequenceDigits` | Int | Liczba cyfr do zerowego wypełnienia |
| `sequenceStartingNumber` | Int | Początkowa liczba dla sekwencji |

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Odpowiedzi błędów

### Błąd konfiguracji pola
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Błąd uprawnień
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Ważne uwagi

### Automatycznie generowane identyfikatory
- **Tylko do odczytu**: Automatycznie generowane identyfikatory nie mogą być edytowane ręcznie
- **Trwałe**: Po przypisaniu identyfikatory sekwencji nie zmieniają się
- **Chronologiczne**: Identyfikatory odzwierciedlają kolejność tworzenia
- **Zakres**: Sekwencje są niezależne dla każdego projektu

### Rozważania dotyczące wydajności
- Generowanie identyfikatorów dla nowych rekordów odbywa się synchronicznie za pomocą wyzwalaczy bazy danych
- Generowanie sekwencji wykorzystuje `FOR UPDATE` blokady do operacji atomowych
- System zadań w tle istnieje, ale implementacja robocza jest w toku
- Rozważ liczby początkowe sekwencji dla projektów o dużym wolumenie

### Migracja i aktualizacje
- Dodanie automatycznego sekwencjonowania do istniejących rekordów kolejkowało zadanie w tle (implementacja robocza w toku)
- Zmiana ustawień sekwencji wpływa tylko na przyszłe rekordy
- Istniejące identyfikatory pozostają niezmienione po aktualizacjach konfiguracji
- Liczniki sekwencji kontynuują od aktualnej maksymalnej wartości

## Najlepsze praktyki

### Projektowanie konfiguracji
- Wybierz opisowe prefiksy, które nie będą kolidować z innymi systemami
- Użyj odpowiedniego wypełnienia cyfr dla oczekiwanego wolumenu
- Ustaw rozsądne liczby początkowe, aby uniknąć konfliktów
- Przetestuj konfigurację z przykładowymi danymi przed wdrożeniem

### Wytyczne dotyczące prefiksów
- Zachowaj prefiksy krótkie i łatwe do zapamiętania (2-5 znaków)
- Używaj wielkich liter dla spójności
- Dodaj separatory (myślniki, podkreślenia) dla czytelności
- Unikaj znaków specjalnych, które mogą powodować problemy w adresach URL lub systemach

### Planowanie sekwencji
- Oszacuj swój wolumen rekordów, aby wybrać odpowiednie wypełnienie cyfr
- Rozważ przyszły wzrost przy ustawianiu liczb początkowych
- Zaplanuj różne zakresy sekwencji dla różnych typów rekordów
- Udokumentuj swoje schematy identyfikatorów dla odniesienia zespołu

## Typowe przypadki użycia

1. **Systemy wsparcia**
   - Numery biletów: `TICK-001`, `TICK-002`
   - Identyfikatory spraw: `CASE-2024-001`
   - Wnioski wsparcia: `SUP-001`

2. **Zarządzanie projektami**
   - Identyfikatory zadań: `TASK-001`, `TASK-002`
   - Elementy sprintu: `SPRINT-001`
   - Numery dostaw: `DEL-001`

3. **Operacje biznesowe**
   - Numery zamówień: `ORD-2024-001`
   - Identyfikatory faktur: `INV-001`
   - Zamówienia zakupu: `PO-001`

4. **Zarządzanie jakością**
   - Raporty błędów: `BUG-001`
   - Identyfikatory przypadków testowych: `TEST-001`
   - Numery przeglądów: `REV-001`

## Funkcje integracji

### Z automatyzacjami
- Wyzwalaj akcje, gdy przypisywane są unikalne identyfikatory
- Użyj wzorców identyfikatorów w regułach automatyzacji
- Odwołuj się do identyfikatorów w szablonach e-mail i powiadomieniach

### Z wyszukiwaniami
- Odwołuj się do unikalnych identyfikatorów z innych rekordów
- Znajdź rekordy według unikalnego identyfikatora
- Wyświetl identyfikatory powiązanych rekordów

### Z raportowaniem
- Grupuj i filtruj według wzorców identyfikatorów
- Śledź trendy przypisywania identyfikatorów
- Monitoruj wykorzystanie sekwencji i luki

## Ograniczenia

- **Tylko sekwencyjnie**: Identyfikatory są przypisywane w kolejności chronologicznej
- **Brak luk**: Usunięte rekordy pozostawiają luki w sekwencjach
- **Brak ponownego użycia**: Numery sekwencji nigdy nie są ponownie używane
- **Zakres projektu**: Nie można dzielić sekwencji między projektami
- **Ograniczenia formatu**: Ograniczone opcje formatowania
- **Brak aktualizacji zbiorczych**: Nie można zbiorczo aktualizować istniejących identyfikatorów sekwencji
- **Brak niestandardowej logiki**: Nie można wdrażać niestandardowych reguł generowania identyfikatorów

## Powiązane zasoby

- [Pola tekstowe](/api/custom-fields/text-single) - Do ręcznych identyfikatorów tekstowych
- [Pola numeryczne](/api/custom-fields/number) - Do sekwencji numerycznych
- [Przegląd pól niestandardowych](/api/custom-fields/2.list-custom-fields) - Ogólne pojęcia
- [Automatyzacje](/api/automations) - Do reguł automatyzacji opartych na identyfikatorach