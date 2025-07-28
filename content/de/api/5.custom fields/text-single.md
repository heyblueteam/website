---
title: Einzelzeilen-Textbenutzerfeld
description: Erstellen Sie Einzelzeilen-Textfelder für kurze Textwerte wie Namen, Titel und Beschriftungen
---

Einzelzeilen-Textbenutzerfelder ermöglichen es Ihnen, kurze Textwerte zu speichern, die für die Eingabe in einer Zeile vorgesehen sind. Sie sind ideal für Namen, Titel, Beschriftungen oder beliebige Textdaten, die auf einer einzigen Zeile angezeigt werden sollen.

## Einfaches Beispiel

Erstellen Sie ein einfaches Einzelzeilen-Textfeld:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Einzelzeilen-Textfeld mit Beschreibung:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
  }) {
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
| `name` | String! | ✅ Ja | Anzeigename des Textfelds |
| `type` | CustomFieldType! | ✅ Ja | Muss sein `TEXT_SINGLE` |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Der Projektkontext wird automatisch aus Ihren Authentifizierungsheadern bestimmt. Kein `projectId` Parameter ist erforderlich.

## Textwerte festlegen

Um einen Einzelzeilen-Textwert in einem Datensatz festzulegen oder zu aktualisieren:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Textbenutzerfelds |
| `text` | String | Nein | Einzelzeilen-Textinhalt, der gespeichert werden soll |

## Datensätze mit Textwerten erstellen

Beim Erstellen eines neuen Datensatzes mit Einzelzeilen-Textwerten:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | ID! | Eindeutiger Bezeichner für den Feldwert |
| `customField` | CustomField! | Die Definition des Benutzerfelds (enthält den Textwert) |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

**Wichtig**: Textwerte werden über das `customField.value.text` Feld abgerufen, nicht direkt auf TodoCustomField.

## Textwerte abfragen

Beim Abfragen von Datensätzen mit Textbenutzerfeldern greifen Sie über den `customField.value.text` Pfad auf den Text zu:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

Die Antwort enthält den Text in der geschachtelten Struktur:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Textvalidierung

### Formularvalidierung
Wenn Einzelzeilen-Textfelder in Formularen verwendet werden:
- Vor- und Nachlaufende Leerzeichen werden automatisch entfernt
- Die erforderliche Validierung wird angewendet, wenn das Feld als erforderlich markiert ist
- Es wird keine spezifische Formatvalidierung angewendet

### Validierungsregeln
- Akzeptiert beliebige Zeichenfolgeninhalte, einschließlich Zeilenumbrüche (obwohl nicht empfohlen)
- Keine Zeichenlängenbeschränkungen (bis zu den Datenbankgrenzen)
- Unterstützt Unicode-Zeichen und Sonderzeichen
- Zeilenumbrüche werden beibehalten, sind jedoch nicht für diesen Feldtyp vorgesehen

### Typische Textbeispiele
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Wichtige Hinweise

### Speicherkapazität
- Wird mit MySQL `MediumText` Typ gespeichert
- Unterstützt bis zu 16 MB Textinhalt
- Identische Speicherung wie mehrzeilige Textfelder
- UTF-8-Codierung für internationale Zeichen

### Direkte API vs. Formulare
- **Formulare**: Automatische Leerzeichenentfernung und erforderliche Validierung
- **Direkte API**: Text wird genau so gespeichert, wie er bereitgestellt wird
- **Empfehlung**: Verwenden Sie Formulare für Benutzereingaben, um eine konsistente Formatierung sicherzustellen

### TEXT_SINGLE vs. TEXT_MULTI
- **TEXT_SINGLE**: Einzelzeilen-Textinput, ideal für kurze Werte
- **TEXT_MULTI**: Mehrzeilige Textarea-Eingabe, ideal für längere Inhalte
- **Backend**: Beide verwenden identische Speicherung und Validierung
- **Frontend**: Unterschiedliche UI-Komponenten für die Dateneingabe
- **Absicht**: TEXT_SINGLE ist semantisch für Einzelzeilenwerte gedacht

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|---------------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## Fehlermeldungen

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

## Beste Praktiken

### Inhaltsrichtlinien
- Halten Sie den Text prägnant und geeignet für Einzelzeilen
- Vermeiden Sie Zeilenumbrüche für die beabsichtigte Einzelzeilenanzeige
- Verwenden Sie eine konsistente Formatierung für ähnliche Datentypen
- Berücksichtigen Sie Zeichenbeschränkungen basierend auf Ihren UI-Anforderungen

### Dateneingabe
- Stellen Sie klare Feldbeschreibungen zur Verfügung, um Benutzer zu leiten
- Verwenden Sie Formulare für Benutzereingaben, um die Validierung sicherzustellen
- Validieren Sie das Inhaltsformat in Ihrer Anwendung, falls erforderlich
- Erwägen Sie die Verwendung von Dropdowns für standardisierte Werte

### Leistungsüberlegungen
- Einzelzeilen-Textfelder sind leichtgewichtig und leistungsfähig
- Erwägen Sie Indizes für häufig durchsucht Felder
- Verwenden Sie angemessene Anzeigegrößen in Ihrer UI
- Überwachen Sie die Inhaltslänge für Anzeigezwecke

## Filterung und Suche

### Enthält-Suche
Einzelzeilen-Textfelder unterstützen die Suche nach Teilzeichenfolgen:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Suchfähigkeiten
- Groß-/Kleinschreibung wird bei der Teilzeichenfolgenübereinstimmung ignoriert
- Unterstützt partielle Wortübereinstimmungen
- Exakte Wertübereinstimmung
- Keine Volltextsuche oder -bewertung

## Häufige Anwendungsfälle

1. **Identifikatoren und Codes**
   - Produkt-SKUs
   - Bestellnummern
   - Referenzcodes
   - Versionsnummern

2. **Namen und Titel**
   - Kundennamen
   - Projekttitel
   - Produktnamen
   - Kategorienbeschriftungen

3. **Kurze Beschreibungen**
   - Kurze Zusammenfassungen
   - Statusbeschriftungen
   - Prioritätsindikatoren
   - Klassifizierungstags

4. **Externe Referenzen**
   - Ticketnummern
   - Rechnungsreferenzen
   - IDs externer Systeme
   - Dokumentnummern

## Integrationsfunktionen

### Mit Nachschlägen
- Referenzieren Sie Textdaten aus anderen Datensätzen
- Finden Sie Datensätze nach Textinhalt
- Anzeigen verwandter Textinformationen
- Aggregieren Sie Textwerte aus mehreren Quellen

### Mit Formularen
- Automatische Leerzeichenentfernung
- Validierung erforderlicher Felder
- Benutzeroberfläche für Einzelzeilen-Textinput
- Anzeige von Zeichenbeschränkungen (wenn konfiguriert)

### Mit Importen/Exporten
- Direkte CSV-Spaltenzuordnung
- Automatische Zuweisung von Textwerten
- Unterstützung für den Massenimport von Daten
- Export in Tabellenformat

## Einschränkungen

### Automatisierungsbeschränkungen
- Nicht direkt als Automatisierungsauslöserfelder verfügbar
- Können nicht in Automatisierungsfeldaktualisierungen verwendet werden
- Können in Automatisierungsbedingungen referenziert werden
- Verfügbar in E-Mail-Vorlagen und Webhooks

### Allgemeine Einschränkungen
- Keine integrierte Textformatierung oder -gestaltung
- Keine automatische Validierung über erforderliche Felder hinaus
- Keine integrierte Einzigartigkeitsdurchsetzung
- Keine Inhaltskompression für sehr großen Text
- Keine Versionierung oder Änderungsverfolgung
- Eingeschränkte Suchfähigkeiten (keine Volltextsuche)

## Verwandte Ressourcen

- [Mehrzeilige Textfelder](/api/custom-fields/text-multi) - Für längere Textinhalte
- [E-Mail-Felder](/api/custom-fields/email) - Für E-Mail-Adressen
- [URL-Felder](/api/custom-fields/url) - Für Webseitenadressen
- [Eindeutige ID-Felder](/api/custom-fields/unique-id) - Für automatisch generierte Identifikatoren
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte