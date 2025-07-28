---
title: Einziges ID benutzerdefiniertes Feld
description: Erstellen Sie automatisch generierte eindeutige Identifikatorfelder mit fortlaufender Nummerierung und benutzerdefinierter Formatierung
---

Eindeutige ID benutzerdefinierte Felder generieren automatisch fortlaufende, eindeutige Identifikatoren für Ihre Datensätze. Sie sind perfekt für die Erstellung von Ticketnummern, Bestell-IDs, Rechnungsnummern oder jedem fortlaufenden Identifikatorsystem in Ihrem Arbeitsablauf.

## Einfaches Beispiel

Erstellen Sie ein einfaches eindeutiges ID-Feld mit automatischer Sequenzierung:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein formatiertes eindeutiges ID-Feld mit Präfix und Nullauffüllung:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput (EINZIG_ID)

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des eindeutigen ID-Feldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `UNIQUE_ID` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |
| `useSequenceUniqueId` | Boolean | Nein | Automatische Sequenzierung aktivieren (Standard: falsch) |
| `prefix` | String | Nein | Textpräfix für generierte IDs (z. B. "TASK-") |
| `sequenceDigits` | Int | Nein | Anzahl der Ziffern für die Nullauffüllung |
| `sequenceStartingNumber` | Int | Nein | Startnummer für die Sequenz |

## Konfigurationsoptionen

### Automatische Sequenzierung (`useSequenceUniqueId`)
- **true**: Generiert automatisch fortlaufende IDs, wenn Datensätze erstellt werden
- **false** oder **undefined**: Manuelle Eingabe erforderlich (funktioniert wie ein Textfeld)

### Präfix (`prefix`)
- Optionales Textpräfix, das allen generierten IDs hinzugefügt wird
- Beispiele: "TASK-", "ORD-", "BUG-", "REQ-"
- Keine Längenbeschränkung, aber vernünftig für die Anzeige halten

### Sequenzziffern (`sequenceDigits`)
- Anzahl der Ziffern für die Nullauffüllung der Sequenznummer
- Beispiel: `sequenceDigits: 3` erzeugt `001`, `002`, `003`
- Wenn nicht angegeben, wird keine Auffüllung angewendet

### Startnummer (`sequenceStartingNumber`)
- Die erste Nummer in der Sequenz
- Beispiel: `sequenceStartingNumber: 1000` beginnt bei 1000, 1001, 1002...
- Wenn nicht angegeben, beginnt es bei 1 (Standardverhalten)

## Generiertes ID-Format

Das endgültige ID-Format kombiniert alle Konfigurationsoptionen:

```
{prefix}{paddedSequenceNumber}
```

### Formatbeispiele

| Konfiguration | Generierte IDs |
|---------------|---------------|
| Keine Optionen | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Lesen von eindeutigen ID-Werten

### Abfrage von Datensätzen mit eindeutigen IDs
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Antwortformat
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Automatische ID-Generierung

### Wann IDs generiert werden
- **Datensatz Erstellung**: IDs werden automatisch zugewiesen, wenn neue Datensätze erstellt werden
- **Feld Hinzufügen**: Beim Hinzufügen eines EINZIG_ID-Feldes zu bestehenden Datensätzen wird ein Hintergrundjob in die Warteschlange gestellt (Implementierung des Arbeiters ausstehend)
- **Hintergrundverarbeitung**: Die ID-Generierung für neue Datensätze erfolgt synchron über Datenbankauslöser

### Generierungsprozess
1. **Auslöser**: Neuer Datensatz wird erstellt oder EINZIG_ID-Feld wird hinzugefügt
2. **Sequenzsuche**: Das System findet die nächste verfügbare Sequenznummer
3. **ID-Zuweisung**: Die Sequenznummer wird dem Datensatz zugewiesen
4. **Zähleraktualisierung**: Der Sequenzzähler wird für zukünftige Datensätze erhöht
5. **Formatierung**: Die ID wird mit Präfix und Auffüllung formatiert, wenn sie angezeigt wird

### Einzigartigkeit Garantien
- **Datenbankbeschränkungen**: Eindeutige Einschränkung auf Sequenz-IDs innerhalb jedes Feldes
- **Atomare Operationen**: Die Sequenzgenerierung verwendet Datenbanksperren, um Duplikate zu verhindern
- **Projektbereich**: Sequenzen sind pro Projekt unabhängig
- **Schutz vor Wettlaufbedingungen**: Gleichzeitige Anfragen werden sicher behandelt

## Manueller vs. Automatischer Modus

### Automatischer Modus (`useSequenceUniqueId: true`)
- IDs werden automatisch über Datenbankauslöser generiert
- Fortlaufende Nummerierung ist garantiert
- Atomare Sequenzgenerierung verhindert Duplikate
- Formatierte IDs kombinieren Präfix + gepolsterte Sequenznummer

### Manueller Modus (`useSequenceUniqueId: false` oder `undefined`)
- Funktioniert wie ein reguläres Textfeld
- Benutzer können benutzerdefinierte Werte über `setTodoCustomField` mit dem `text`-Parameter eingeben
- Keine automatische Generierung
- Keine Durchsetzung der Eindeutigkeit über die Datenbankbeschränkungen hinaus

## Manuelle Werte festlegen (nur im manuellen Modus)

Wenn `useSequenceUniqueId` falsch ist, können Sie Werte manuell festlegen:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Antwortfelder

### TodoCustomField Antwort (EINZIG_ID)

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Identifikator für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `sequenceId` | Int | Die generierte Sequenznummer (gefüllt für EINZIG_ID-Felder) |
| `text` | String | Der formatierte Textwert (kombiniert Präfix + gepolsterte Sequenz) |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt aktualisiert wurde |

### CustomField Antwort (EINZIG_ID)

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | Ob die automatische Sequenzierung aktiviert ist |
| `prefix` | String | Textpräfix für generierte IDs |
| `sequenceDigits` | Int | Anzahl der Ziffern für die Nullauffüllung |
| `sequenceStartingNumber` | Int | Startnummer für die Sequenz |

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Fehlerantworten

### Fehler bei der Feldkonfiguration
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Berechtigungsfehler
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Wichtige Hinweise

### Automatisch generierte IDs
- **Schreibgeschützt**: Automatisch generierte IDs können nicht manuell bearbeitet werden
- **Permanent**: Einmal zugewiesene Sequenz-IDs ändern sich nicht
- **Chronologisch**: IDs spiegeln die Erstellungsreihenfolge wider
- **Bereichsgebunden**: Sequenzen sind pro Projekt unabhängig

### Leistungsüberlegungen
- Die ID-Generierung für neue Datensätze erfolgt synchron über Datenbankauslöser
- Die Sequenzgenerierung verwendet `FOR UPDATE`-Sperren für atomare Operationen
- Ein Hintergrundjobsystem existiert, aber die Implementierung des Arbeiters steht noch aus
- Berücksichtigen Sie die Startnummern der Sequenzen für Projekte mit hohem Volumen

### Migration und Updates
- Das Hinzufügen der automatischen Sequenzierung zu bestehenden Datensätzen stellt einen Hintergrundjob in die Warteschlange (Arbeiter ausstehend)
- Änderungen an den Sequenzeinstellungen betreffen nur zukünftige Datensätze
- Bestehende IDs bleiben unverändert, wenn Konfigurationsupdates durchgeführt werden
- Sequenzzähler setzen sich vom aktuellen Maximum fort

## Best Practices

### Konfigurationsdesign
- Wählen Sie beschreibende Präfixe, die nicht mit anderen Systemen in Konflikt stehen
- Verwenden Sie angemessene Ziffernauffüllungen für Ihr erwartetes Volumen
- Setzen Sie angemessene Startnummern, um Konflikte zu vermeiden
- Testen Sie die Konfiguration mit Beispieldaten vor der Bereitstellung

### Präfixrichtlinien
- Halten Sie Präfixe kurz und einprägsam (2-5 Zeichen)
- Verwenden Sie Großbuchstaben für Konsistenz
- Fügen Sie Trennzeichen (Bindestriche, Unterstriche) für die Lesbarkeit hinzu
- Vermeiden Sie Sonderzeichen, die Probleme in URLs oder Systemen verursachen könnten

### Sequenzplanung
- Schätzen Sie Ihr Datensatzvolumen, um die geeignete Ziffernauffüllung zu wählen
- Berücksichtigen Sie zukünftiges Wachstum bei der Festlegung von Startnummern
- Planen Sie unterschiedliche Sequenzbereiche für verschiedene Datensatztypen
- Dokumentieren Sie Ihre ID-Schemata zur Referenz für das Team

## Häufige Anwendungsfälle

1. **Support-Systeme**
   - Ticketnummern: `TICK-001`, `TICK-002`
   - Fall-IDs: `CASE-2024-001`
   - Supportanfragen: `SUP-001`

2. **Projektmanagement**
   - Aufgaben-IDs: `TASK-001`, `TASK-002`
   - Sprint-Elemente: `SPRINT-001`
   - Liefernummern: `DEL-001`

3. **Geschäftsbetrieb**
   - Bestellnummern: `ORD-2024-001`
   - Rechnungs-IDs: `INV-001`
   - Bestellungen: `PO-001`

4. **Qualitätsmanagement**
   - Fehlerberichte: `BUG-001`
   - Testfall-IDs: `TEST-001`
   - Überprüfungsnummern: `REV-001`

## Integrationsfunktionen

### Mit Automatisierungen
- Aktionen auslösen, wenn eindeutige IDs zugewiesen werden
- ID-Muster in Automatisierungsregeln verwenden
- IDs in E-Mail-Vorlagen und Benachrichtigungen referenzieren

### Mit Nachschlägen
- Eindeutige IDs aus anderen Datensätzen referenzieren
- Datensätze nach eindeutiger ID finden
- Verwandte Datensatzidentifikatoren anzeigen

### Mit Berichterstattung
- Nach ID-Mustern gruppieren und filtern
- Trends bei der ID-Zuweisung verfolgen
- Sequenznutzung und -lücken überwachen

## Einschränkungen

- **Nur sequenziell**: IDs werden in chronologischer Reihenfolge zugewiesen
- **Keine Lücken**: Gelöschte Datensätze hinterlassen Lücken in den Sequenzen
- **Keine Wiederverwendung**: Sequenznummern werden niemals wiederverwendet
- **Projektbereich**: Sequenzen können nicht über Projekte hinweg geteilt werden
- **Formatbeschränkungen**: Begrenzte Formatierungsoptionen
- **Keine Massenupdates**: Bestehende Sequenz-IDs können nicht massenhaft aktualisiert werden
- **Keine benutzerdefinierte Logik**: Benutzerdefinierte ID-Generierungsregeln können nicht implementiert werden

## Verwandte Ressourcen

- [Textfelder](/api/custom-fields/text-single) - Für manuelle Textidentifikatoren
- [Zahlenfelder](/api/custom-fields/number) - Für numerische Sequenzen
- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/2.list-custom-fields) - Allgemeine Konzepte
- [Automatisierungen](/api/automations) - Für ID-basierte Automatisierungsregeln