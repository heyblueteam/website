---
title: Mehrzeiliges Textbenutzerfeld
description: Erstellen Sie mehrzeilige Textfelder für längere Inhalte wie Beschreibungen, Notizen und Kommentare
---

Mehrzeilige Textbenutzerfelder ermöglichen es Ihnen, längere Textinhalte mit Zeilenumbrüchen und Formatierungen zu speichern. Sie sind ideal für Beschreibungen, Notizen, Kommentare oder beliebige Textdaten, die mehrere Zeilen benötigen.

## Einfaches Beispiel

Erstellen Sie ein einfaches mehrzeiliges Textfeld:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein mehrzeiliges Textfeld mit Beschreibung:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
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
| `name` | String! | ✅ Ja | Anzeigename des Textfelds |
| `type` | CustomFieldType! | ✅ Ja | Muss `TEXT_MULTI` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis:** Das `projectId` wird als separates Argument an die Mutation übergeben, nicht als Teil des Eingabeobjekts. Alternativ kann der Projektkontext aus dem `X-Bloo-Project-ID`-Header in Ihrer GraphQL-Anfrage bestimmt werden.

## Textwerte festlegen

Um einen mehrzeiligen Textwert in einem Datensatz festzulegen oder zu aktualisieren:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Textbenutzerfelds |
| `text` | String | Nein | Mehrzeiliger Textinhalt, der gespeichert werden soll |

## Datensätze mit Textwerten erstellen

Beim Erstellen eines neuen Datensatzes mit mehrzeiligen Textwerten:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
| `id` | String! | Eindeutiger Bezeichner für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Felds |
| `text` | String | Der gespeicherte mehrzeilige Textinhalt |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

## Textvalidierung

### Formularvalidierung
Wenn mehrzeilige Textfelder in Formularen verwendet werden:
- Vorangestellter und nachgestellter Leerraum wird automatisch entfernt
- Erforderliche Validierung wird angewendet, wenn das Feld als erforderlich markiert ist
- Es wird keine spezifische Formatvalidierung angewendet

### Validierungsregeln
- Akzeptiert beliebigen Stringinhalt, einschließlich Zeilenumbrüche
- Keine Zeichenlängenbeschränkungen (bis zu den Datenbankgrenzen)
- Unterstützt Unicode-Zeichen und spezielle Symbole
- Zeilenumbrüche werden im Speicher beibehalten

### Gültige Textbeispiele
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Wichtige Hinweise

### Speicherkapazität
- Wird im MySQL `MediumText`-Typ gespeichert
- Unterstützt bis zu 16 MB Textinhalt
- Zeilenumbrüche und Formatierungen werden beibehalten
- UTF-8-Codierung für internationale Zeichen

### Direkte API vs. Formulare
- **Formulare**: Automatische Leerraumtrimmung und erforderliche Validierung
- **Direkte API**: Text wird genau so gespeichert, wie er bereitgestellt wird
- **Empfehlung**: Verwenden Sie Formulare für Benutzereingaben, um konsistente Formatierungen sicherzustellen

### TEXT_MULTI vs. TEXT_SINGLE
- **TEXT_MULTI**: Mehrzeilige Textarea-Eingabe, ideal für längere Inhalte
- **TEXT_SINGLE**: Einzeilige Texteingabe, ideal für kurze Werte
- **Backend**: Beide Typen sind identisch - dasselbe Speicherfeld, dieselbe Validierung und Verarbeitung
- **Frontend**: Unterschiedliche UI-Komponenten für die Dateneingabe (Textarea vs. Eingabefeld)
- **Wichtig**: Der Unterschied zwischen TEXT_MULTI und TEXT_SINGLE besteht rein zu UI-Zwecken

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|---------------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Best Practices

### Inhaltsorganisation
- Verwenden Sie konsistente Formatierungen für strukturierte Inhalte
- Erwägen Sie die Verwendung einer markdown-ähnlichen Syntax für die Lesbarkeit
- Teilen Sie lange Inhalte in logische Abschnitte auf
- Verwenden Sie Zeilenumbrüche, um die Lesbarkeit zu verbessern

### Dateneingabe
- Geben Sie klare Feldbeschreibungen an, um die Benutzer zu leiten
- Verwenden Sie Formulare für Benutzereingaben, um die Validierung sicherzustellen
- Berücksichtigen Sie Zeichenbeschränkungen basierend auf Ihrem Anwendungsfall
- Validieren Sie das Inhaltsformat in Ihrer Anwendung, falls erforderlich

### Leistungsüberlegungen
- Sehr lange Textinhalte können die Abfrageleistung beeinträchtigen
- Erwägen Sie die Paginierung zur Anzeige großer Textfelder
- Indexüberlegungen für die Suchfunktionalität
- Überwachen Sie die Speichernutzung für Felder mit großen Inhalten

## Filterung und Suche

### Enthält-Suche
Mehrzeilige Textfelder unterstützen die Teilstring-Suche über benutzerdefinierte Feldfilter:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Suchfunktionen
- Teilstring-Abgleich innerhalb von Textfeldern mit dem `CONTAINS`-Operator
- Groß-/Kleinschreibung-empfindliche Suche mit dem `NCONTAINS`-Operator
- Exakte Übereinstimmung mit dem `IS`-Operator
- Negative Übereinstimmung mit dem `NOT`-Operator
- Sucht über alle Zeilen des Textes
- Unterstützt partielle Wortübereinstimmungen

## Häufige Anwendungsfälle

1. **Projektmanagement**
   - Aufgabenbeschreibungen
   - Projektanforderungen
   - Besprechungsnotizen
   - Statusaktualisierungen

2. **Kundensupport**
   - Problembeschreibungen
   - Lösungsnotizen
   - Kundenfeedback
   - Kommunikationsprotokolle

3. **Inhaltsverwaltung**
   - Artikelinhalt
   - Produktbeschreibungen
   - Benutzerkommentare
   - Bewertungsdetails

4. **Dokumentation**
   - Prozessbeschreibungen
   - Anleitungen
   - Richtlinien
   - Referenzmaterialien

## Integrationsfunktionen

### Mit Automatisierungen
- Aktionen auslösen, wenn sich der Textinhalt ändert
- Schlüsselwörter aus dem Textinhalt extrahieren
- Zusammenfassungen oder Benachrichtigungen erstellen
- Textinhalt mit externen Diensten verarbeiten

### Mit Nachschlägen
- Textdaten aus anderen Datensätzen referenzieren
- Textinhalte aus mehreren Quellen aggregieren
- Datensätze nach Textinhalt finden
- Verwandte Textinformationen anzeigen

### Mit Formularen
- Automatische Leerraumtrimmung
- Validierung erforderlicher Felder
- Mehrzeilige Textarea-Benutzeroberfläche
- Zeichenanzahl anzeigen (falls konfiguriert)

## Einschränkungen

- Keine integrierte Textformatierung oder Rich-Text-Bearbeitung
- Keine automatische Linkerkennung oder -konvertierung
- Keine Rechtschreibprüfung oder Grammatikvalidierung
- Keine integrierte Textanalyse oder -verarbeitung
- Keine Versionierung oder Änderungsverfolgung
- Eingeschränkte Suchmöglichkeiten (keine Volltextsuche)
- Keine Inhaltskompression für sehr großen Text

## Verwandte Ressourcen

- [Einzeilige Textfelder](/api/custom-fields/text-single) - Für kurze Textwerte
- [E-Mail-Felder](/api/custom-fields/email) - Für E-Mail-Adressen
- [URL-Felder](/api/custom-fields/url) - Für Webseitenadressen
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/2.list-custom-fields) - Allgemeine Konzepte