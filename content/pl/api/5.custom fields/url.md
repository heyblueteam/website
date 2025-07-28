---
title: Pole niestandardowe URL
description: Twórz pola URL do przechowywania adresów stron internetowych i linków
---

Pola niestandardowe URL pozwalają na przechowywanie adresów stron internetowych i linków w Twoich rekordach. Są idealne do śledzenia stron internetowych projektów, linków referencyjnych, adresów URL dokumentacji lub wszelkich zasobów internetowych związanych z Twoją pracą.

## Podstawowy przykład

Utwórz proste pole URL:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Zaawansowany przykład

Utwórz pole URL z opisem:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola URL |
| `type` | CustomFieldType! | ✅ Tak | Musi być `URL` |
| `description` | String | Nie | Tekst pomocniczy wyświetlany użytkownikom |

**Uwaga:** `projectId` jest przekazywane jako osobny argument do mutacji, a nie jako część obiektu wejściowego.

## Ustawianie wartości URL

Aby ustawić lub zaktualizować wartość URL w rekordzie:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola URL |
| `text` | String! | ✅ Tak | Adres URL do przechowania |

## Tworzenie rekordów z wartościami URL

Podczas tworzenia nowego rekordu z wartościami URL:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
| `text` | String | Przechowywany adres URL |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

## Walidacja URL

### Aktualna implementacja
- **Bezpośrednie API**: Obecnie nie wymusza się walidacji formatu URL
- **Formularze**: Walidacja URL jest planowana, ale obecnie nieaktywna
- **Przechowywanie**: W polach URL można przechowywać dowolną wartość tekstową

### Planowana walidacja
Przyszłe wersje będą zawierać:
- Walidację protokołu HTTP/HTTPS
- Sprawdzanie poprawności formatu URL
- Walidację nazwy domeny
- Automatyczne dodawanie prefiksu protokołu

### Zalecane formaty URL
Choć obecnie nie są wymuszane, używaj tych standardowych formatów:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Ważne uwagi

### Format przechowywania
- URL są przechowywane jako zwykły tekst bez modyfikacji
- Brak automatycznego dodawania protokołu (http://, https://)
- Wrażliwość na wielkość liter zachowana zgodnie z wprowadzonym tekstem
- Nie wykonuje się kodowania/odkodowywania URL

### Bezpośrednie API vs Formularze
- **Formularze**: Planowana walidacja URL (obecnie nieaktywna)
- **Bezpośrednie API**: Brak walidacji - można przechowywać dowolny tekst
- **Rekomendacja**: Waliduj URL w swojej aplikacji przed przechowaniem

### Pola URL vs Pola tekstowe
- **URL**: Semantycznie przeznaczone dla adresów internetowych
- **TEXT_SINGLE**: Ogólny tekst w jednej linii
- **Backend**: Aktualnie identyczne przechowywanie i walidacja
- **Frontend**: Różne komponenty UI do wprowadzania danych

## Wymagane uprawnienia

Operacje na polach niestandardowych wykorzystują uprawnienia oparte na rolach:

| Akcja | Wymagana rola |
|-------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Uwaga:** Uprawnienia są sprawdzane na podstawie ról użytkowników w projekcie, a nie konkretnych stałych uprawnień.

## Odpowiedzi błędów

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

## Najlepsze praktyki

### Standardy formatu URL
- Zawsze dołączaj protokół (http:// lub https://)
- Używaj HTTPS, gdy to możliwe, dla bezpieczeństwa
- Testuj URL przed przechowaniem, aby upewnić się, że są dostępne
- Rozważ użycie skróconych URL do celów wyświetlania

### Jakość danych
- Waliduj URL w swojej aplikacji przed przechowaniem
- Sprawdzaj typowe błędy (brakujące protokoły, niepoprawne domeny)
- Standaryzuj formaty URL w całej organizacji
- Rozważ dostępność i dostępność URL

### Rozważania dotyczące bezpieczeństwa
- Bądź ostrożny z URL dostarczonymi przez użytkowników
- Waliduj domeny, jeśli ograniczasz do konkretnych witryn
- Rozważ skanowanie URL w poszukiwaniu złośliwej zawartości
- Używaj URL HTTPS przy obsłudze danych wrażliwych

## Filtrowanie i wyszukiwanie

### Wyszukiwanie zawierające
Pola URL obsługują wyszukiwanie podciągów:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Możliwości wyszukiwania
- Niedostrzeganie wielkości liter w dopasowywaniu podciągów
- Częściowe dopasowywanie domen
- Wyszukiwanie ścieżek i parametrów
- Brak filtrowania specyficznego dla protokołu

## Typowe przypadki użycia

1. **Zarządzanie projektami**
   - Strony internetowe projektów
   - Linki do dokumentacji
   - Adresy URL repozytoriów
   - Strony demo

2. **Zarządzanie treścią**
   - Materiały referencyjne
   - Linki źródłowe
   - Zasoby multimedialne
   - Artykuły zewnętrzne

3. **Wsparcie klienta**
   - Strony internetowe klientów
   - Dokumentacja wsparcia
   - Artykuły bazy wiedzy
   - Samouczki wideo

4. **Sprzedaż i marketing**
   - Strony internetowe firm
   - Strony produktów
   - Materiały marketingowe
   - Profile w mediach społecznościowych

## Funkcje integracji

### Z wyszukiwaniami
- Odwołania do URL z innych rekordów
- Znajdowanie rekordów według domeny lub wzoru URL
- Wyświetlanie powiązanych zasobów internetowych
- Agregowanie linków z wielu źródeł

### Z formularzami
- Specyficzne komponenty wejściowe dla URL
- Planowana walidacja dla poprawnego formatu URL
- Możliwości podglądu linków (frontend)
- Wyświetlanie klikalnych URL

### Z raportowaniem
- Śledzenie użycia URL i wzorców
- Monitorowanie uszkodzonych lub niedostępnych linków
- Kategoryzowanie według domeny lub protokołu
- Eksportowanie list URL do analizy

## Ograniczenia

### Aktualne ograniczenia
- Brak aktywnej walidacji formatu URL
- Brak automatycznego dodawania protokołu
- Brak weryfikacji linków lub sprawdzania dostępności
- Brak skracania lub rozszerzania URL
- Brak generowania favicon lub podglądów

### Ograniczenia automatyzacji
- Nie są dostępne jako pola wyzwalające automatyzację
- Nie mogą być używane w aktualizacjach pól automatyzacji
- Mogą być odwoływane w warunkach automatyzacji
- Dostępne w szablonach e-mail i webhookach

### Ogólne ograniczenia
- Brak wbudowanej funkcjonalności podglądu linków
- Brak automatycznego skracania URL
- Brak śledzenia kliknięć lub analityki
- Brak sprawdzania wygaśnięcia URL
- Brak skanowania złośliwych URL

## Przyszłe ulepszenia

### Planowane funkcje
- Walidacja protokołu HTTP/HTTPS
- Niestandardowe wzory walidacji regex
- Automatyczne dodawanie prefiksu protokołu
- Sprawdzanie dostępności URL

### Potencjalne ulepszenia
- Generowanie podglądów linków
- Wyświetlanie favicon
- Integracja skracania URL
- Możliwości śledzenia kliknięć
- Wykrywanie uszkodzonych linków

## Powiązane zasoby

- [Pola tekstowe](/api/custom-fields/text-single) - Dla danych tekstowych, które nie są URL
- [Pola e-mailowe](/api/custom-fields/email) - Dla adresów e-mail
- [Przegląd pól niestandardowych](/api/custom-fields/2.list-custom-fields) - Ogólne koncepcje

## Migracja z pól tekstowych

Jeśli migrujesz z pól tekstowych do pól URL:

1. **Utwórz pole URL** o tej samej nazwie i konfiguracji
2. **Eksportuj istniejące wartości tekstowe** aby zweryfikować, że są poprawnymi URL
3. **Zaktualizuj rekordy** aby używały nowego pola URL
4. **Usuń stare pole tekstowe** po udanej migracji
5. **Zaktualizuj aplikacje** aby używały komponentów UI specyficznych dla URL

### Przykład migracji
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```