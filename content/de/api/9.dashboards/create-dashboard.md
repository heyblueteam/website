---
title: Dashboard erstellen
description: Erstellen Sie ein neues Dashboard zur Datenvisualisierung und Berichterstattung in Blue
---

## Ein Dashboard erstellen

Die `createDashboard` Mutation ermöglicht es Ihnen, ein neues Dashboard innerhalb Ihres Unternehmens oder Projekts zu erstellen. Dashboards sind leistungsstarke Visualisierungstools, die Teams helfen, Kennzahlen zu verfolgen, den Fortschritt zu überwachen und datengestützte Entscheidungen zu treffen.

### Einfaches Beispiel

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

### Projektbezogenes Dashboard

Erstellen Sie ein Dashboard, das mit einem bestimmten Projekt verknüpft ist:

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

## Eingabeparameter

### CreateDashboardInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `companyId` | String! | ✅ Ja | Die ID des Unternehmens, in dem das Dashboard erstellt wird |
| `title` | String! | ✅ Ja | Der Name des Dashboards. Muss eine nicht leere Zeichenfolge sein |
| `projectId` | String | Nein | Optionale ID eines Projekts, das mit diesem Dashboard verknüpft werden soll |

## Antwortfelder

Die Mutation gibt ein vollständiges `Dashboard` Objekt zurück:

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutige Kennung für das erstellte Dashboard |
| `title` | String! | Der Dashboard-Titel wie angegeben |
| `companyId` | String! | Das Unternehmen, dem dieses Dashboard gehört |
| `projectId` | String | Die zugehörige Projekt-ID (falls angegeben) |
| `project` | Project | Das zugehörige Projektobjekt (falls projectId angegeben wurde) |
| `createdBy` | User! | Der Benutzer, der das Dashboard erstellt hat (Sie) |
| `dashboardUsers` | [DashboardUser!]! | Liste der Benutzer mit Zugriff (anfangs nur der Ersteller) |
| `createdAt` | DateTime! | Zeitstempel, wann das Dashboard erstellt wurde |
| `updatedAt` | DateTime! | Zeitstempel der letzten Änderung (gleich wie createdAt für neue Dashboards) |

### DashboardUser-Felder

Wenn ein Dashboard erstellt wird, wird der Ersteller automatisch als Dashboard-Benutzer hinzugefügt:

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutige Kennung für die Dashboard-Benutzerbeziehung |
| `user` | User! | Das Benutzerobjekt mit Zugriff auf das Dashboard |
| `role` | DashboardRole! | Die Rolle des Benutzers (Ersteller erhält vollen Zugriff) |
| `dashboard` | Dashboard! | Rückverweis auf das Dashboard |

## Erforderliche Berechtigungen

Jeder authentifizierte Benutzer, der zum angegebenen Unternehmen gehört, kann Dashboards erstellen. Es gibt keine speziellen Rollenanforderungen.

| Benutzerstatus | Kann Dashboard erstellen |
|----------------|-------------------------|
| Company Member | ✅ Ja |
| Nicht-Unternehmensmitglied | ❌ Nein |
| Unauthenticated | ❌ Nein |

## Fehlerantworten

### Ungültiges Unternehmen
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

### Benutzer nicht im Unternehmen
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

### Ungültiges Projekt
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

### Leerer Titel
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

## Wichtige Hinweise

- **Automatischer Besitz**: Der Benutzer, der das Dashboard erstellt, wird automatisch dessen Eigentümer mit vollen Berechtigungen
- **Projektzuordnung**: Wenn Sie eine `projectId` angeben, muss sie zum gleichen Unternehmen gehören
- **Anfängliche Berechtigungen**: Nur der Ersteller hat anfänglich Zugriff. Verwenden Sie `editDashboard`, um weitere Benutzer hinzuzufügen
- **Titelanforderungen**: Dashboard-Titel müssen nicht leere Zeichenfolgen sein. Es gibt keine Einzigartigkeitsanforderung
- **Unternehmensmitgliedschaft**: Sie müssen Mitglied des Unternehmens sein, um darin Dashboards zu erstellen

## Workflow zur Dashboard-Erstellung

1. **Erstellen Sie das Dashboard** mit dieser Mutation
2. **Konfigurieren Sie Diagramme und Widgets** mit der Dashboard-Builder-Benutzeroberfläche
3. **Fügen Sie Teammitglieder hinzu** mit der `editDashboard` Mutation mit `dashboardUsers`
4. **Richten Sie Filter und Datumsbereiche** über die Dashboard-Oberfläche ein
5. **Teilen oder betten** Sie das Dashboard mit seiner eindeutigen ID

## Anwendungsfälle

1. **Executive Dashboards**: Erstellen Sie hochrangige Übersichten über Unternehmenskennzahlen
2. **Projektverfolgung**: Erstellen Sie projektbezogene Dashboards zur Überwachung des Fortschritts
3. **Teamleistung**: Verfolgen Sie die Produktivität und Leistungskennzahlen des Teams
4. **Kundenberichterstattung**: Erstellen Sie Dashboards für kundenorientierte Berichte
5. **Echtzeitüberwachung**: Richten Sie Dashboards für Live-Betriebsdaten ein

## Best Practices

1. **Benennungskonventionen**: Verwenden Sie klare, beschreibende Titel, die den Zweck des Dashboards angeben
2. **Projektzuordnung**: Verknüpfen Sie Dashboards mit Projekten, wenn sie projektspezifisch sind
3. **Zugriffsverwaltung**: Fügen Sie Teammitglieder sofort nach der Erstellung zur Zusammenarbeit hinzu
4. **Organisation**: Erstellen Sie eine Dashboard-Hierarchie mit konsistenten Benennungsschemata

## Verwandte Operationen

- [Dashboards auflisten](/api/dashboards/) - Alle Dashboards für ein Unternehmen oder Projekt abrufen
- [Dashboard bearbeiten](/api/dashboards/rename-dashboard) - Dashboard umbenennen oder Benutzer verwalten
- [Dashboard kopieren](/api/dashboards/copy-dashboard) - Ein bestehendes Dashboard duplizieren
- [Dashboard löschen](/api/dashboards/delete-dashboard) - Ein Dashboard entfernen