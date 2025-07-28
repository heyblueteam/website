---
title: Utwórz pulpit nawigacyjny
description: Utwórz nowy pulpit nawigacyjny do wizualizacji danych i raportowania w Blue
---

## Utwórz pulpit nawigacyjny

Mutacja `createDashboard` pozwala na utworzenie nowego pulpitu nawigacyjnego w Twojej firmie lub projekcie. Pulpity nawigacyjne to potężne narzędzia wizualizacyjne, które pomagają zespołom śledzić metryki, monitorować postępy i podejmować decyzje oparte na danych.

### Podstawowy przykład

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### Pulpit nawigacyjny specyficzny dla projektu

Utwórz pulpit nawigacyjny powiązany z konkretnym projektem:

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## Parametry wejściowe

### CreateDashboardInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `companyId` | String! | ✅ Tak | ID firmy, w której zostanie utworzony pulpit nawigacyjny |
| `title` | String! | ✅ Tak | Nazwa pulpitu nawigacyjnego. Musi być niepustym ciągiem |
| `projectId` | String | Nie | Opcjonalne ID projektu do powiązania z tym pulpitem nawigacyjnym |

## Pola odpowiedzi

Mutacja zwraca kompletny obiekt `Dashboard`:

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator utworzonego pulpitu nawigacyjnego |
| `title` | String! | Tytuł pulpitu nawigacyjnego podany przez użytkownika |
| `companyId` | String! | Firma, do której należy ten pulpit nawigacyjny |
| `projectId` | String | Powiązane ID projektu (jeśli podano) |
| `project` | Project | Powiązany obiekt projektu (jeśli podano projectId) |
| `createdBy` | User! | Użytkownik, który utworzył pulpit nawigacyjny (ty) |
| `dashboardUsers` | [DashboardUser!]! | Lista użytkowników z dostępem (początkowo tylko twórca) |
| `createdAt` | DateTime! | Znacznik czasu utworzenia pulpitu nawigacyjnego |
| `updatedAt` | DateTime! | Znacznik czasu ostatniej modyfikacji (taki sam jak createdAt dla nowych pulpitów nawigacyjnych) |

### Pola DashboardUser

Gdy pulpit nawigacyjny jest tworzony, twórca jest automatycznie dodawany jako użytkownik pulpitu nawigacyjnego:

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator relacji użytkownika pulpitu nawigacyjnego |
| `user` | User! | Obiekt użytkownika z dostępem do pulpitu nawigacyjnego |
| `role` | DashboardRole! | Rola użytkownika (twórca ma pełny dostęp) |
| `dashboard` | Dashboard! | Odniesienie do pulpitu nawigacyjnego |

## Wymagane uprawnienia

Każdy uwierzytelniony użytkownik, który należy do określonej firmy, może tworzyć pulpity nawigacyjne. Nie ma specjalnych wymagań dotyczących ról.

| Status użytkownika | Może utworzyć pulpit nawigacyjny |
|--------------------|----------------------------------|
| Company Member | ✅ Tak |
| Nieczłonek firmy | ❌ Nie |
| Unauthenticated | ❌ Nie |

## Odpowiedzi błędów

### Nieprawidłowa firma
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Użytkownik nie w firmie
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Nieprawidłowy projekt
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Pusty tytuł
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Ważne uwagi

- **Automatyczne posiadanie**: Użytkownik tworzący pulpit nawigacyjny automatycznie staje się jego właścicielem z pełnymi uprawnieniami
- **Powiązanie z projektem**: Jeśli podasz `projectId`, musi on należeć do tej samej firmy
- **Początkowe uprawnienia**: Tylko twórca ma początkowo dostęp. Użyj `editDashboard`, aby dodać więcej użytkowników
- **Wymagania dotyczące tytułu**: Tytuły pulpitów nawigacyjnych muszą być niepustymi ciągami. Nie ma wymogu unikalności
- **Członkostwo w firmie**: Musisz być członkiem firmy, aby tworzyć w niej pulpity nawigacyjne

## Proces tworzenia pulpitu nawigacyjnego

1. **Utwórz pulpit nawigacyjny** za pomocą tej mutacji
2. **Skonfiguruj wykresy i widżety** za pomocą interfejsu budowy pulpitu nawigacyjnego
3. **Dodaj członków zespołu** za pomocą mutacji `editDashboard` z `dashboardUsers`
4. **Skonfiguruj filtry i zakresy dat** za pomocą interfejsu pulpitu nawigacyjnego
5. **Udostępnij lub osadź** pulpit nawigacyjny za pomocą jego unikalnego ID

## Przykłady użycia

1. **Pulpity nawigacyjne dla kierownictwa**: Twórz przeglądy metryk firmy na wysokim poziomie
2. **Śledzenie projektów**: Buduj pulpity nawigacyjne specyficzne dla projektów, aby monitorować postępy
3. **Wydajność zespołu**: Śledź produktywność zespołu i metryki osiągnięć
4. **Raportowanie dla klientów**: Twórz pulpity nawigacyjne do raportów skierowanych do klientów
5. **Monitorowanie w czasie rzeczywistym**: Ustawiaj pulpity nawigacyjne dla danych operacyjnych na żywo

## Najlepsze praktyki

1. **Konwencje nazewnictwa**: Używaj jasnych, opisowych tytułów, które wskazują na cel pulpitu nawigacyjnego
2. **Powiązanie z projektem**: Łącz pulpity nawigacyjne z projektami, gdy są specyficzne dla projektów
3. **Zarządzanie dostępem**: Dodawaj członków zespołu natychmiast po utworzeniu w celu współpracy
4. **Organizacja**: Twórz hierarchię pulpitu nawigacyjnego, używając spójnych wzorców nazewnictwa

## Powiązane operacje

- [Lista pulpitów nawigacyjnych](/api/dashboards/) - Pobierz wszystkie pulpity nawigacyjne dla firmy lub projektu
- [Edytuj pulpit nawigacyjny](/api/dashboards/rename-dashboard) - Zmień nazwę pulpitu nawigacyjnego lub zarządzaj użytkownikami
- [Kopiuj pulpit nawigacyjny](/api/dashboards/copy-dashboard) - Duplikuj istniejący pulpit nawigacyjny
- [Usuń pulpit nawigacyjny](/api/dashboards/delete-dashboard) - Usuń pulpit nawigacyjny