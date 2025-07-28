---
title: Pole niestandardowe referencji
description: Twórz pola referencyjne, które łączą się z rekordami w innych projektach w celu nawiązywania relacji między projektami
---

Pola niestandardowe referencji umożliwiają tworzenie linków między rekordami w różnych projektach, co pozwala na nawiązywanie relacji między projektami i dzielenie się danymi. Stanowią potężny sposób na łączenie powiązanej pracy w strukturze projektowej Twojej organizacji.

## Podstawowy przykład

Utwórz proste pole referencyjne:

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## Zaawansowany przykład

Utwórz pole referencyjne z filtrowaniem i wieloma wyborami:

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola referencyjnego |
| `type` | CustomFieldType! | ✅ Tak | Musi być `REFERENCE` |
| `referenceProjectId` | String | Nie | ID projektu do odniesienia |
| `referenceMultiple` | Boolean | Nie | Zezwól na wybór wielu rekordów (domyślnie: fałsz) |
| `referenceFilter` | TodoFilterInput | Nie | Kryteria filtrowania dla odniesionych rekordów |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

**Uwaga**: Pola niestandardowe są automatycznie powiązane z projektem na podstawie bieżącego kontekstu projektu użytkownika.

## Konfiguracja referencji

### Pojedyncze vs wielokrotne odniesienia

**Pojedyncze odniesienie (domyślnie):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Użytkownicy mogą wybrać jeden rekord z projektu odniesienia
- Zwraca pojedynczy obiekt Todo

**Wielokrotne odniesienia:**
```graphql
{
  referenceMultiple: true
}
```
- Użytkownicy mogą wybrać wiele rekordów z projektu odniesienia
- Zwraca tablicę obiektów Todo

### Filtrowanie referencji

Użyj `referenceFilter`, aby ograniczyć, które rekordy mogą być wybierane:

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## Ustawianie wartości referencyjnych

### Pojedyncze odniesienie

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Wielokrotne odniesienia

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagane | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID pola niestandardowego referencji |
| `customFieldReferenceTodoIds` | [String!] | ✅ Tak | Tablica ID odniesionych rekordów |

## Tworzenie rekordów z odniesieniami

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## Pola odpowiedzi

### Odpowiedź TodoCustomField

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja pola referencyjnego |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

**Uwaga**: Odniesione todos są dostępne za pośrednictwem `customField.selectedTodos`, a nie bezpośrednio na TodoCustomField.

### Odniesione pola Todo

Każde odniesione Todo zawiera:

| Pole | Typ | Opis |
|------|-----|------|
| `id` | ID! | Unikalny identyfikator odniesionego rekordu |
| `title` | String! | Tytuł odniesionego rekordu |
| `status` | TodoStatus! | Bieżący status (AKTYWNY, ZREALIZOWANY itp.) |
| `description` | String | Opis odniesionego rekordu |
| `dueDate` | DateTime | Termin, jeśli jest ustawiony |
| `assignees` | [User!] | Przypisani użytkownicy |
| `tags` | [Tag!] | Powiązane tagi |
| `project` | Project! | Projekt zawierający odniesiony rekord |

## Zapytania o dane referencyjne

### Podstawowe zapytanie

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### Zaawansowane zapytanie z danymi zagnieżdżonymi

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## Wymagane uprawnienia

| Akcja | Wymagane uprawnienie |
|-------|---------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Ważne**: Użytkownicy muszą mieć uprawnienia do przeglądania odniesionego projektu, aby zobaczyć powiązane rekordy.

## Dostęp między projektami

### Widoczność projektu

- Użytkownicy mogą odnosić się tylko do rekordów z projektów, do których mają dostęp
- Odniesione rekordy respektują uprawnienia oryginalnego projektu
- Zmiany w odniesionych rekordach pojawiają się w czasie rzeczywistym
- Usunięcie odniesionych rekordów usuwa je z pól referencyjnych

### Dziedziczenie uprawnień

- Pola referencyjne dziedziczą uprawnienia z obu projektów
- Użytkownicy potrzebują dostępu do przeglądania odniesionego projektu
- Uprawnienia do edycji opierają się na zasadach bieżącego projektu
- Odniesione dane są tylko do odczytu w kontekście pola referencyjnego

## Odpowiedzi na błędy

### Nieprawidłowy projekt odniesienia

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### Rekord odniesiony nie znaleziony

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

### Odrzucono uprawnienia

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## Najlepsze praktyki

### Projektowanie pól

1. **Jasne nazewnictwo** - Używaj opisowych nazw, które wskazują na relację
2. **Odpowiednie filtrowanie** - Ustaw filtry, aby pokazywały tylko odpowiednie rekordy
3. **Rozważ uprawnienia** - Upewnij się, że użytkownicy mają dostęp do odniesionych projektów
4. **Dokumentuj relacje** - Podaj jasne opisy połączeń

### Rozważania dotyczące wydajności

1. **Ogranicz zakres odniesienia** - Użyj filtrów, aby zmniejszyć liczbę wybieralnych rekordów
2. **Unikaj głębokiego zagnieżdżania** - Nie twórz skomplikowanych łańcuchów odniesień
3. **Rozważ pamięć podręczną** - Odniesione dane są buforowane dla wydajności
4. **Monitoruj użycie** - Śledź, jak odniesienia są używane w projektach

### Integralność danych

1. **Zarządzaj usunięciami** - Planuj, co się stanie, gdy odniesione rekordy zostaną usunięte
2. **Waliduj uprawnienia** - Upewnij się, że użytkownicy mogą uzyskać dostęp do odniesionych projektów
3. **Aktualizuj zależności** - Rozważ wpływ przy zmianie odniesionych rekordów
4. **Ślady audytowe** - Śledź relacje odniesień dla zgodności

## Typowe przypadki użycia

### Zależności projektowe

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### Wymagania klienta

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### Alokacja zasobów

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### Kontrola jakości

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## Integracja z wyszukiwaniami

Pola referencyjne współpracują z [Polami wyszukiwania](/api/custom-fields/lookup), aby pobierać dane z odniesionych rekordów. Pola wyszukiwania mogą wyodrębniać wartości z rekordów wybranych w polach referencyjnych, ale są tylko ekstraktorami danych (nie obsługują funkcji agregacji, takich jak SUM).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## Ograniczenia

- Odniesione projekty muszą być dostępne dla użytkownika
- Zmiany w uprawnieniach projektu odniesionego wpływają na dostęp do pól referencyjnych
- Głębokie zagnieżdżanie odniesień może wpłynąć na wydajność
- Brak wbudowanej walidacji dla odniesień cyklicznych
- Brak automatycznego ograniczenia zapobiegającego odniesieniom w tym samym projekcie
- Walidacja filtrów nie jest egzekwowana podczas ustawiania wartości referencyjnych

## Powiązane zasoby

- [Pola wyszukiwania](/api/custom-fields/lookup) - Ekstrahowanie danych z odniesionych rekordów
- [API projektów](/api/projects) - Zarządzanie projektami, które zawierają odniesienia
- [API rekordów](/api/records) - Praca z rekordami, które mają odniesienia
- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne pojęcia