---
title: Telefonbenutzerdefiniertes Feld
description: Erstellen Sie Telefonfelder, um Telefonnummern mit internationalem Format zu speichern und zu validieren
---

Telefonbenutzerdefinierte Felder ermöglichen es Ihnen, Telefonnummern in Datensätzen mit integrierter Validierung und internationalem Format zu speichern. Sie sind ideal für die Verfolgung von Kontaktinformationen, Notfallkontakten oder anderen telefonbezogenen Daten in Ihren Projekten.

## Basisbeispiel

Erstellen Sie ein einfaches Telefonfeld:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Telefonfeld mit Beschreibung:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ Ja | Anzeigename des Telefonfelds |
| `type` | CustomFieldType! | ✅ Ja | Muss `PHONE` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Benutzerdefinierte Felder werden automatisch mit dem Projekt basierend auf dem aktuellen Projektkontext des Benutzers verknüpft. Kein `projectId` Parameter ist erforderlich.

## Telefonwerte festlegen

Um einen Telefonwert in einem Datensatz festzulegen oder zu aktualisieren:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Telefonbenutzerdefinierten Feldes |
| `text` | String | Nein | Telefonnummer mit Ländervorwahl |
| `regionCode` | String | Nein | Ländervorwahl (automatisch erkannt) |

**Hinweis**: Während `text` im Schema optional ist, ist eine Telefonnummer erforderlich, damit das Feld sinnvoll ist. Bei Verwendung von `setTodoCustomField` erfolgt keine Validierung - Sie können jeden Textwert und regionCode speichern. Die automatische Erkennung erfolgt nur während der Erstellung des Datensatzes.

## Datensätze mit Telefonwerten erstellen

Beim Erstellen eines neuen Datensatzes mit Telefonwerten:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
    }
  }
}
```

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Identifikator für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `text` | String | Die formatierte Telefonnummer (internationales Format) |
| `regionCode` | String | Die Ländervorwahl (z.B. "US", "GB", "CA") |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

## Telefonnummernvalidierung

**Wichtig**: Die Validierung und Formatierung von Telefonnummern erfolgt nur beim Erstellen neuer Datensätze über `createTodo`. Bei der Aktualisierung bestehender Telefonwerte mit `setTodoCustomField` erfolgt keine Validierung und die Werte werden wie angegeben gespeichert.

### Akzeptierte Formate (Während der Datensatz Erstellung)
Telefonnummern müssen eine Ländervorwahl in einem dieser Formate enthalten:

- **E.164-Format (bevorzugt)**: `+12345678900`
- **Internationales Format**: `+1 234 567 8900`
- **International mit Interpunktion**: `+1 (234) 567-8900`
- **Ländervorwahl mit Bindestrichen**: `+1-234-567-8900`

**Hinweis**: Nationale Formate ohne Ländervorwahl (wie `(234) 567-8900`) werden während der Erstellung des Datensatzes abgelehnt.

### Validierungsregeln (Während der Datensatz Erstellung)
- Verwendet libphonenumber-js zum Parsen und Validieren
- Akzeptiert verschiedene internationale Telefonnummernformate
- Erkennt automatisch das Land aus der Nummer
- Formatiert die Nummer im internationalen Anzeigeformat (z.B. `+1 234 567 8900`)
- Extrahiert und speichert die Ländervorwahl separat (z.B. `US`)

### Gültige Telefonnummernbeispiele
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Ungültige Telefonnummernbeispiele
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Speicherformat

Beim Erstellen von Datensätzen mit Telefonnummern:
- **text**: Wird im internationalen Format gespeichert (z.B. `+1 234 567 8900`) nach der Validierung
- **regionCode**: Wird als ISO-Ländervorwahl gespeichert (z.B. `US`, `GB`, `CA`) automatisch erkannt

Beim Aktualisieren über `setTodoCustomField`:
- **text**: Wird genau wie angegeben gespeichert (keine Formatierung)
- **regionCode**: Wird genau wie angegeben gespeichert (keine Validierung)

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Fehlerantworten

### Ungültiges Telefonformat
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
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

### Fehlende Ländervorwahl
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Best Practices

### Dateneingabe
- Immer die Ländervorwahl in Telefonnummern einfügen
- E.164-Format für Konsistenz verwenden
- Nummern vor der Speicherung für wichtige Vorgänge validieren
- Regionale Präferenzen für die Anzeigeformatierung berücksichtigen

### Datenqualität
- Nummern im internationalen Format für globale Kompatibilität speichern
- regionCode für länderspezifische Funktionen verwenden
- Telefonnummern vor kritischen Vorgängen (SMS, Anrufe) validieren
- Zeitunterschiede bei der Kontaktzeit berücksichtigen

### Internationale Überlegungen
- Ländervorwahl wird automatisch erkannt und gespeichert
- Nummern werden im internationalen Standard formatiert
- Regionale Anzeigepräferenzen können regionCode verwenden
- Lokale Wählgewohnheiten bei der Anzeige berücksichtigen

## Häufige Anwendungsfälle

1. **Kontaktverwaltung**
   - Telefonnummern von Kunden
   - Kontaktinformationen von Anbietern
   - Telefonnummern von Teammitgliedern
   - Kontaktdaten für den Support

2. **Notfallkontakte**
   - Notfalltelefonnummern
   - Kontaktinformationen für Bereitschaftsdienste
   - Kontakte für Krisenreaktionen
   - Eskalationstelefonnummern

3. **Kundensupport**
   - Telefonnummern von Kunden
   - Rückrufnummern für den Support
   - Verifizierungsnummern
   - Nachverfolgungskontaktnummern

4. **Vertrieb & Marketing**
   - Telefonnummern von Leads
   - Kontaktlisten für Kampagnen
   - Kontaktinformationen von Partnern
   - Telefonnummern von Empfehlungsquellen

## Integrationsfunktionen

### Mit Automatisierungen
- Aktionen auslösen, wenn Telefonfelder aktualisiert werden
- SMS-Benachrichtigungen an gespeicherte Telefonnummern senden
- Nachverfolgungsaufgaben basierend auf Telefonänderungen erstellen
- Anrufe basierend auf Telefonnummerndaten weiterleiten

### Mit Nachschlagen
- Telefonnummerndaten aus anderen Datensätzen referenzieren
- Telefonlisten aus mehreren Quellen aggregieren
- Datensätze nach Telefonnummer finden
- Kontaktinformationen abgleichen

### Mit Formularen
- Automatische Telefonnummernvalidierung
- Überprüfung des internationalen Formats
- Erkennung der Ländervorwahl
- Echtzeit-Formatfeedback

## Einschränkungen

- Erfordert die Ländervorwahl für alle Nummern
- Keine integrierten SMS- oder Anruffunktionen
- Keine Telefonnummernüberprüfung über die Formatprüfung hinaus
- Keine Speicherung von Telefonnummernmetadaten (Netzbetreiber, Typ usw.)
- Nationale Formatnummern ohne Ländervorwahl werden abgelehnt
- Keine automatische Telefonnummernformatierung in der Benutzeroberfläche über den internationalen Standard hinaus

## Verwandte Ressourcen

- [Textfelder](/api/custom-fields/text-single) - Für nicht-telefonbezogene Textdaten
- [E-Mail-Felder](/api/custom-fields/email) - Für E-Mail-Adressen
- [URL-Felder](/api/custom-fields/url) - Für Webseitenadressen
- [Übersicht über benutzerdefinierte Felder](/custom-fields/list-custom-fields) - Allgemeine Konzepte