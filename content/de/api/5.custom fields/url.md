---
title: URL Benutzerdefiniertes Feld
description: Erstellen Sie URL-Felder, um Webseitenadressen und Links zu speichern
---

URL-Benutzerdefinierte Felder ermöglichen es Ihnen, Webseitenadressen und Links in Ihren Datensätzen zu speichern. Sie sind ideal zum Verfolgen von Projektwebseiten, Referenzlinks, Dokumentations-URLs oder anderen webbasierten Ressourcen, die mit Ihrer Arbeit in Zusammenhang stehen.

## Einfaches Beispiel

Erstellen Sie ein einfaches URL-Feld:

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

## Fortgeschrittenes Beispiel

Erstellen Sie ein URL-Feld mit Beschreibung:

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

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des URL-Feldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `URL` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis:** Die `projectId` wird als separates Argument an die Mutation übergeben, nicht als Teil des Eingabeobjekts.

## URL-Werte festlegen

Um einen URL-Wert in einem Datensatz festzulegen oder zu aktualisieren:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des URL-Benutzerdefinierten Feldes |
| `text` | String! | ✅ Ja | URL-Adresse, die gespeichert werden soll |

## Datensätze mit URL-Werten erstellen

Beim Erstellen eines neuen Datensatzes mit URL-Werten:

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

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `text` | String | Die gespeicherte URL-Adresse |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

## URL-Validierung

### Aktuelle Implementierung
- **Direkte API**: Derzeit wird keine URL-Formatvalidierung durchgesetzt
- **Formulare**: URL-Validierung ist geplant, aber derzeit nicht aktiv
- **Speicherung**: Jeder Stringwert kann in URL-Feldern gespeichert werden

### Geplante Validierung
Zukünftige Versionen werden Folgendes umfassen:
- HTTP/HTTPS-Protokollvalidierung
- Überprüfung des gültigen URL-Formats
- Validierung des Domainnamens
- Automatische Protokollpräfixaddition

### Empfohlene URL-Formate
Obwohl derzeit nicht durchgesetzt, verwenden Sie diese Standardformate:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Wichtige Hinweise

### Speicherformat
- URLs werden als Klartext ohne Modifikation gespeichert
- Keine automatische Protokolladdition (http://, https://)
- Groß-/Kleinschreibung wird wie eingegeben beibehalten
- Keine URL-Codierung/-Decodierung durchgeführt

### Direkte API vs Formulare
- **Formulare**: Geplante URL-Validierung (derzeit nicht aktiv)
- **Direkte API**: Keine Validierung - jeder Text kann gespeichert werden
- **Empfehlung**: Validieren Sie URLs in Ihrer Anwendung, bevor Sie sie speichern

### URL vs Textfelder
- **URL**: Semantisch für Webadressen vorgesehen
- **TEXT_SINGLE**: Allgemeiner einzeiliger Text
- **Backend**: Derzeit identische Speicherung und Validierung
- **Frontend**: Unterschiedliche UI-Komponenten für die Dateneingabe

## Erforderliche Berechtigungen

Benutzerdefinierte Feldoperationen verwenden rollenbasierte Berechtigungen:

| Aktion | Erforderliche Rolle |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Hinweis:** Berechtigungen werden basierend auf den Benutzerrollen im Projekt überprüft, nicht auf spezifischen Berechtigungskonstanten.

## Fehlerantworten

### Feld nicht gefunden
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

### Erforderliche Feldvalidierung (nur Formulare)
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

## Best Practices

### URL-Formatstandards
- Immer Protokoll einschließen (http:// oder https://)
- Verwenden Sie HTTPS, wenn möglich, für Sicherheit
- Testen Sie URLs, bevor Sie sie speichern, um sicherzustellen, dass sie zugänglich sind
- Ziehen Sie in Betracht, verkürzte URLs für Anzeigezwecke zu verwenden

### Datenqualität
- Validieren Sie URLs in Ihrer Anwendung, bevor Sie sie speichern
- Überprüfen Sie auf häufige Tippfehler (fehlende Protokolle, falsche Domains)
- Standardisieren Sie URL-Formate in Ihrer Organisation
- Berücksichtigen Sie die Zugänglichkeit und Verfügbarkeit von URLs

### Sicherheitsüberlegungen
- Seien Sie vorsichtig mit von Benutzern bereitgestellten URLs
- Validieren Sie Domains, wenn Sie auf bestimmte Websites beschränken
- Ziehen Sie in Betracht, URLs auf bösartigen Inhalt zu scannen
- Verwenden Sie HTTPS-URLs, wenn Sie mit sensiblen Daten umgehen

## Filterung und Suche

### Enthält-Suche
URL-Felder unterstützen die Suche nach Teilstrings:

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

### Suchfähigkeiten
- Groß-/Kleinschreibung ignorierende Teilstringübereinstimmung
- Teilweise Domainübereinstimmung
- Pfad- und Parameter-Suche
- Keine protokollspezifische Filterung

## Häufige Anwendungsfälle

1. **Projektmanagement**
   - Projektwebseiten
   - Dokumentationslinks
   - Repository-URLs
   - Demoseiten

2. **Inhaltsverwaltung**
   - Referenzmaterialien
   - Quelllinks
   - Medienressourcen
   - Externe Artikel

3. **Kundensupport**
   - Kundenwebseiten
   - Unterstützungsdokumentation
   - Wissensdatenbankartikel
   - Video-Tutorials

4. **Vertrieb & Marketing**
   - Unternehmenswebseiten
   - Produktseiten
   - Marketingmaterialien
   - Profile in sozialen Medien

## Integrationsmerkmale

### Mit Nachschlägen
- Referenz-URLs aus anderen Datensätzen
- Finden Sie Datensätze nach Domain oder URL-Muster
- Anzeigen verwandter Webressourcen
- Aggregieren Sie Links aus mehreren Quellen

### Mit Formularen
- URL-spezifische Eingabekomponenten
- Geplante Validierung für das richtige URL-Format
- Linkvorschau-Funktionen (Frontend)
- Klickbare URL-Anzeige

### Mit Berichterstattung
- Verfolgen Sie die Nutzung und Muster von URLs
- Überwachen Sie defekte oder nicht zugängliche Links
- Kategorisieren nach Domain oder Protokoll
- Exportieren Sie URL-Listen zur Analyse

## Einschränkungen

### Aktuelle Einschränkungen
- Keine aktive URL-Formatvalidierung
- Keine automatische Protokolladdition
- Keine Linküberprüfung oder Zugänglichkeitsprüfung
- Keine URL-Verkürzung oder -Erweiterung
- Keine Favicon- oder Vorschau-Generierung

### Automatisierungseinschränkungen
- Nicht verfügbar als Automatisierungsauslöserfelder
- Können nicht in Automatisierungsfeldaktualisierungen verwendet werden
- Können in Automatisierungsbedingungen referenziert werden
- Verfügbar in E-Mail-Vorlagen und Webhooks

### Allgemeine Einschränkungen
- Keine integrierte Linkvorschaufunktionalität
- Keine automatische URL-Verkürzung
- Keine Klickverfolgung oder Analytik
- Keine URL-Ablaufprüfung
- Kein Scannen auf bösartige URLs

## Zukünftige Verbesserungen

### Geplante Funktionen
- HTTP/HTTPS-Protokollvalidierung
- Benutzerdefinierte Regex-Validierungsmuster
- Automatische Protokollpräfixaddition
- Überprüfung der URL-Zugänglichkeit

### Potenzielle Verbesserungen
- Generierung von Linkvorschauen
- Anzeige von Favicons
- Integration von URL-Verkürzungen
- Klickverfolgungsfunktionen
- Erkennung defekter Links

## Verwandte Ressourcen

- [Textfelder](/api/custom-fields/text-single) - Für nicht-URL-Textdaten
- [E-Mail-Felder](/api/custom-fields/email) - Für E-Mail-Adressen
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/2.list-custom-fields) - Allgemeine Konzepte

## Migration von Textfeldern

Wenn Sie von Textfeldern zu URL-Feldern migrieren:

1. **Erstellen Sie ein URL-Feld** mit demselben Namen und derselben Konfiguration
2. **Exportieren Sie vorhandene Textwerte**, um zu überprüfen, ob sie gültige URLs sind
3. **Aktualisieren Sie Datensätze**, um das neue URL-Feld zu verwenden
4. **Löschen Sie das alte Textfeld** nach erfolgreicher Migration
5. **Aktualisieren Sie Anwendungen**, um URL-spezifische UI-Komponenten zu verwenden

### Migrationsbeispiel
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