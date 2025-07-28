---
title: E-Mail Benutzerdefiniertes Feld
description: Erstellen Sie E-Mail-Felder, um E-Mail-Adressen zu speichern und zu validieren
---

Benutzerdefinierte E-Mail-Felder ermöglichen es Ihnen, E-Mail-Adressen in Datensätzen mit integrierter Validierung zu speichern. Sie sind ideal, um Kontaktinformationen, E-Mail-Adressen von Zuweisungen oder alle E-Mail-bezogenen Daten in Ihren Projekten zu verfolgen.

## Grundlegendes Beispiel

Erstellen Sie ein einfaches E-Mail-Feld:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein E-Mail-Feld mit Beschreibung:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des E-Mail-Feldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `EMAIL` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

## E-Mail-Werte festlegen

Um einen E-Mail-Wert in einem Datensatz festzulegen oder zu aktualisieren:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des benutzerdefinierten E-Mail-Feldes |
| `text` | String | Nein | E-Mail-Adresse, die gespeichert werden soll |

## Datensätze mit E-Mail-Werten erstellen

Beim Erstellen eines neuen Datensatzes mit E-Mail-Werten:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Antwortfelder

### CustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | ID! | Eindeutiger Identifikator für das benutzerdefinierte Feld |
| `name` | String! | Anzeigename des E-Mail-Feldes |
| `type` | CustomFieldType! | Der Feldtyp (EMAIL) |
| `description` | String | Hilfetext für das Feld |
| `value` | JSON | Enthält den E-Mail-Wert (siehe unten) |
| `createdAt` | DateTime! | Wann das Feld erstellt wurde |
| `updatedAt` | DateTime! | Wann das Feld zuletzt geändert wurde |

**Wichtig**: E-Mail-Werte werden über das `customField.value.text` Feld abgerufen, nicht direkt in der Antwort.

## Abfragen von E-Mail-Werten

Beim Abfragen von Datensätzen mit benutzerdefinierten E-Mail-Feldern greifen Sie über den `customField.value.text` Pfad auf die E-Mail zu:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

Die Antwort enthält die E-Mail in der verschachtelten Struktur:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## E-Mail-Validierung

### Formularvalidierung
Wenn E-Mail-Felder in Formularen verwendet werden, validieren sie automatisch das E-Mail-Format:
- Verwendet Standard-E-Mail-Validierungsregeln
- Entfernt Leerzeichen aus der Eingabe
- Lehnt ungültige E-Mail-Formate ab

### Validierungsregeln
- Muss ein `@` Symbol enthalten
- Muss ein gültiges Domainformat haben
- Führende/nachfolgende Leerzeichen werden automatisch entfernt
- Häufige E-Mail-Formate werden akzeptiert

### Gültige E-Mail-Beispiele
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Ungültige E-Mail-Beispiele
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Wichtige Hinweise

### Direkte API vs. Formulare
- **Formulare**: Automatische E-Mail-Validierung wird angewendet
- **Direkte API**: Keine Validierung - es kann beliebiger Text gespeichert werden
- **Empfehlung**: Verwenden Sie Formulare für Benutzereingaben, um die Validierung sicherzustellen

### Speicherformat
- E-Mail-Adressen werden als Klartext gespeichert
- Keine spezielle Formatierung oder Analyse
- Groß- und Kleinschreibung: E-Mail-Benutzerdefinierte Felder werden groß- und kleinschreibungssensitiv gespeichert (im Gegensatz zu Benutzer-Authentifizierungs-E-Mails, die in Kleinbuchstaben normalisiert werden)
- Keine maximalen Längenbeschränkungen über die Datenbankbeschränkungen hinaus (16 MB Limit)

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Fehlermeldungen

### Ungültiges E-Mail-Format (nur Formulare)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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

### Dateneingabe
- Validieren Sie E-Mail-Adressen immer in Ihrer Anwendung
- Verwenden Sie E-Mail-Felder nur für tatsächliche E-Mail-Adressen
- Ziehen Sie in Betracht, Formulare für Benutzereingaben zu verwenden, um eine automatische Validierung zu erhalten

### Datenqualität
- Entfernen Sie Leerzeichen, bevor Sie speichern
- Ziehen Sie eine Normalisierung der Groß- und Kleinschreibung in Betracht (typischerweise Kleinbuchstaben)
- Validieren Sie das E-Mail-Format vor wichtigen Operationen

### Datenschutzüberlegungen
- E-Mail-Adressen werden als Klartext gespeichert
- Berücksichtigen Sie Datenschutzbestimmungen (DSGVO, CCPA)
- Implementieren Sie geeignete Zugriffskontrollen

## Häufige Anwendungsfälle

1. **Kontaktverwaltung**
   - E-Mail-Adressen von Kunden
   - Kontaktinformationen von Anbietern
   - E-Mails von Teammitgliedern
   - Kontaktdaten für den Support

2. **Projektmanagement**
   - E-Mails von Stakeholdern
   - E-Mails für Genehmigungskontakte
   - Benachrichtigungsempfänger
   - E-Mails von externen Mitarbeitern

3. **Kundensupport**
   - E-Mail-Adressen von Kunden
   - Kontakte für Support-Tickets
   - Eskalationskontakte
   - Feedback-E-Mail-Adressen

4. **Vertrieb & Marketing**
   - E-Mail-Adressen von Leads
   - Kontaktlisten für Kampagnen
   - Kontaktinformationen von Partnern
   - E-Mails von Empfehlungsquellen

## Integrationsfunktionen

### Mit Automatisierungen
- Aktionen auslösen, wenn E-Mail-Felder aktualisiert werden
- Benachrichtigungen an gespeicherte E-Mail-Adressen senden
- Nachverfolgungsaufgaben basierend auf E-Mail-Änderungen erstellen

### Mit Nachschlägen
- E-Mail-Daten aus anderen Datensätzen referenzieren
- E-Mail-Listen aus mehreren Quellen aggregieren
- Datensätze nach E-Mail-Adresse finden

### Mit Formularen
- Automatische E-Mail-Validierung
- Überprüfung des E-Mail-Formats
- Entfernen von Leerzeichen

## Einschränkungen

- Keine integrierte E-Mail-Verifizierung oder -Validierung über die Formatprüfung hinaus
- Keine E-Mail-spezifischen UI-Funktionen (wie klickbare E-Mail-Links)
- Als Klartext ohne Verschlüsselung gespeichert
- Keine E-Mail-Komposition oder -Versandfähigkeiten
- Keine Speicherung von E-Mail-Metadaten (Anzeigename usw.)
- Direkte API-Aufrufe umgehen die Validierung (nur Formulare validieren)

## Verwandte Ressourcen

- [Textfelder](/api/custom-fields/text-single) - Für nicht-E-Mail-Textdaten
- [URL-Felder](/api/custom-fields/url) - Für Webadressen
- [Telefonfelder](/api/custom-fields/phone) - Für Telefonnummern
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte