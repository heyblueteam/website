---
title: Länderbenutzerdefiniertes Feld
description: Erstellen Sie Länder-Auswahlfelder mit ISO-Ländercode-Validierung
---

Länderbenutzerdefinierte Felder ermöglichen es Ihnen, Länderinformationen für Datensätze zu speichern und zu verwalten. Das Feld unterstützt sowohl Ländernamen als auch ISO Alpha-2-Ländercodes.

**Wichtig**: Das Verhalten der Ländervalidierung und -konvertierung unterscheidet sich erheblich zwischen den Mutationen:
- **createTodo**: Validiert und konvertiert automatisch Ländernamen in ISO-Codes
- **setTodoCustomField**: Akzeptiert jeden Wert ohne Validierung

## Einfaches Beispiel

Erstellen Sie ein einfaches Länderfeld:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Länderfeld mit Beschreibung:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Länderfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `COUNTRY` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Der `projectId` wird nicht in der Eingabe übergeben, sondern vom GraphQL-Kontext bestimmt (typischerweise aus Anfrage-Headern oder Authentifizierung).

## Festlegen von Länderwerten

Länderfelder speichern Daten in zwei Datenbankfeldern:
- **`countryCodes`**: Speichert ISO Alpha-2-Ländercodes als kommagetrennte Zeichenfolge in der Datenbank (als Array über die API zurückgegeben)
- **`text`**: Speichert Anzeigetext oder Ländernamen als Zeichenfolge

### Verständnis der Parameter

Die `setTodoCustomField`-Mutation akzeptiert zwei optionale Parameter für Länderfelder:

| Parameter | Typ | Erforderlich | Beschreibung | Was es tut |
|-----------|------|--------------|-------------|--------------|
| `todoId` | String! | ✅ Ja | ID des Datensatzes, der aktualisiert werden soll | - |
| `customFieldId` | String! | ✅ Ja | ID des benutzerdefinierten Länderfeldes | - |
| `countryCodes` | [String!] | Nein | Array von ISO Alpha-2-Ländercodes | Stored in the `countryCodes` field |
| `text` | String | Nein | Anzeigetext oder Ländernamen | Stored in the `text` field |

**Wichtig**: 
- In `setTodoCustomField`: Beide Parameter sind optional und werden unabhängig gespeichert
- In `createTodo`: Das System setzt automatisch beide Felder basierend auf Ihrer Eingabe (Sie können sie nicht unabhängig steuern)

### Option 1: Nur Ländercodes verwenden

Speichern Sie validierte ISO-Codes ohne Anzeigetext:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Ergebnis: `countryCodes` = `["US"]`, `text` = `null`

### Option 2: Nur Text verwenden

Speichern Sie Anzeigetext ohne validierte Codes:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Ergebnis: `countryCodes` = `null`, `text` = `"United States"`

**Hinweis**: Bei der Verwendung von `setTodoCustomField` erfolgt keine Validierung, unabhängig davon, welchen Parameter Sie verwenden. Die Werte werden genau so gespeichert, wie sie bereitgestellt werden.

### Option 3: Beide verwenden (Empfohlen)

Speichern Sie sowohl validierte Codes als auch Anzeigetext:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Ergebnis: `countryCodes` = `["US"]`, `text` = `"United States"`

### Mehrere Länder

Speichern Sie mehrere Länder mithilfe von Arrays:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Erstellen von Datensätzen mit Länderwerten

Beim Erstellen von Datensätzen validiert und konvertiert die `createTodo`-Mutation **automatisch** Länderwerte. Dies ist die einzige Mutation, die eine Ländervalidierung durchführt:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      countryCodes
    }
  }
}
```

### Akzeptierte Eingabeformate

| Eingabetyp | Beispiel | Ergebnis |
|------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Nicht unterstützt** - wird als einzelner ungültiger Wert behandelt |
| Mixed format | `"United States, CA"` | **Nicht unterstützt** - wird als einzelner ungültiger Wert behandelt |

## Antwortfelder

### TodoCustomField-Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `text` | String | Anzeigetext (Ländernamen) |
| `countryCodes` | [String!] | Array von ISO Alpha-2-Ländercodes |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

## Länderstandards

Blue verwendet den **ISO 3166-1 Alpha-2**-Standard für Ländercodes:

- Zweibuchstabencodes für Länder (z. B. US, GB, FR, DE)
- Die Validierung mit der `i18n-iso-countries`-Bibliothek **erfolgt nur in createTodo**
- Unterstützt alle offiziell anerkannten Länder

### Beispiel-Ländercodes

| Land | ISO-Code |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Für die vollständige offizielle Liste der ISO 3166-1 Alpha-2-Ländercodes besuchen Sie die [ISO Online Browsing Platform](https://www.iso.org/obp/ui/#search/code/).

## Validierung

**Die Validierung erfolgt nur in der `createTodo`-Mutation**:

1. **Gültiger ISO-Code**: Akzeptiert jeden gültigen ISO Alpha-2-Code
2. **Ländernamen**: Konvertiert automatisch erkannte Ländernamen in Codes
3. **Ungültige Eingabe**: Wirft `CustomFieldValueParseError` für nicht erkannte Werte

**Hinweis**: Die `setTodoCustomField`-Mutation führt KEINE Validierung durch und akzeptiert jeden Zeichenfolgenwert.

### Fehlerbeispiel

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Integrationsfunktionen

### Nachschlagefelder
Länderfelder können von LOOKUP-Benutzerdefinierten Feldern referenziert werden, sodass Sie Länderdaten aus verwandten Datensätzen abrufen können.

### Automatisierungen
Verwenden Sie Länderwerte in Automatisierungsbedingungen:
- Aktionen nach bestimmten Ländern filtern
- Benachrichtigungen basierend auf dem Land senden
- Aufgaben basierend auf geografischen Regionen zuordnen

### Formulare
Länderfelder in Formularen validieren automatisch die Benutzereingabe und konvertieren Ländernamen in Codes.

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Fehlerantworten

### Ungültiger Länderwert
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Feldtyp-Mismatch
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Best Practices

### Eingabeverarbeitung
- Verwenden Sie `createTodo` für automatische Validierung und Konvertierung
- Verwenden Sie `setTodoCustomField` vorsichtig, da es die Validierung umgeht
- Ziehen Sie in Betracht, Eingaben in Ihrer Anwendung zu validieren, bevor Sie `setTodoCustomField` verwenden
- Zeigen Sie vollständige Ländernamen in der Benutzeroberfläche zur Klarheit an

### Datenqualität
- Validieren Sie Länder-Eingaben am Eingabepunkt
- Verwenden Sie konsistente Formate in Ihrem System
- Berücksichtigen Sie regionale Gruppierungen für Berichterstattung

### Mehrere Länder
- Verwenden Sie die Array-Unterstützung in `setTodoCustomField` für mehrere Länder
- Mehrere Länder in `createTodo` werden **nicht unterstützt** über das Wertfeld
- Speichern Sie Ländercodes als Array in `setTodoCustomField` für eine ordnungsgemäße Verarbeitung

## Häufige Anwendungsfälle

1. **Kundenmanagement**
   - Standort der Unternehmenszentrale
   - Versandziele
   - Steuerjurisdiktionen

2. **Projektverfolgung**
   - Projektstandort
   - Standorte der Teammitglieder
   - Marktziele

3. **Compliance & Recht**
   - Regulierungsbehörden
   - Anforderungen an den Datenaufenthalt
   - Exportkontrollen

4. **Vertrieb & Marketing**
   - Gebietszuteilungen
   - Marktsegmentierung
   - Kampagnenzielsetzung

## Einschränkungen

- Unterstützt nur ISO 3166-1 Alpha-2-Codes (2-Buchstaben-Codes)
- Keine integrierte Unterstützung für Länderunterteilungen (Bundesstaaten/Provinzen)
- Keine automatischen Länderflaggen-Icons (nur textbasiert)
- Kann historische Ländercodes nicht validieren
- Keine integrierte Region- oder Kontinentgruppierung
- **Die Validierung funktioniert nur in `createTodo`, nicht in `setTodoCustomField`**
- **Mehrere Länder werden im `createTodo`-Wertfeld nicht unterstützt**
- **Ländercodes werden als kommagetrennte Zeichenfolge gespeichert, nicht als echtes Array**

## Verwandte Ressourcen

- [Übersicht über benutzerdefinierte Felder](/custom-fields/list-custom-fields) - Allgemeine Konzepte zu benutzerdefinierten Feldern
- [Nachschlagefelder](/api/custom-fields/lookup) - Länderdaten aus anderen Datensätzen referenzieren
- [Formulare API](/api/forms) - Länderfelder in benutzerdefinierten Formularen einfügen