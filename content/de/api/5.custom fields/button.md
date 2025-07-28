---
title: Benutzerdefiniertes Schaltflächenfeld
description: Erstellen Sie interaktive Schaltflächenfelder, die Automatisierungen auslösen, wenn sie angeklickt werden
---

Benutzerdefinierte Schaltflächenfelder bieten interaktive UI-Elemente, die Automatisierungen auslösen, wenn sie angeklickt werden. Im Gegensatz zu anderen benutzerdefinierten Feldtypen, die Daten speichern, dienen Schaltflächenfelder als Aktionsauslöser zur Ausführung konfigurierter Workflows.

## Einfaches Beispiel

Erstellen Sie ein einfaches Schaltflächenfeld, das eine Automatisierung auslöst:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie eine Schaltfläche mit Bestätigungsanforderungen:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename der Schaltfläche |
| `type` | CustomFieldType! | ✅ Ja | Muss `BUTTON` sein |
| `projectId` | String! | ✅ Ja | Projekt-ID, in dem das Feld erstellt wird |
| `buttonType` | String | Nein | Bestätigungsverhalten (siehe Schaltflächentypen unten) |
| `buttonConfirmText` | String | Nein | Text, den Benutzer für die harte Bestätigung eingeben müssen |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |
| `required` | Boolean | Nein | Ob das Feld erforderlich ist (Standardwert ist falsch) |
| `isActive` | Boolean | Nein | Ob das Feld aktiv ist (Standardwert ist wahr) |

### Schaltflächentypfeld

Das `buttonType` Feld ist ein Freitext, der von UI-Clients verwendet werden kann, um das Bestätigungsverhalten zu bestimmen. Häufige Werte sind:

- `""` (leer) - Keine Bestätigung
- `"soft"` - Einfaches Bestätigungsdialogfeld
- `"hard"` - Erfordert die Eingabe des Bestätigungstexts

**Hinweis**: Dies sind nur UI-Hinweise. Die API validiert oder erzwingt keine spezifischen Werte.

## Auslösen von Schaltflächenklicks

Um einen Schaltflächenklick auszulösen und die zugehörigen Automatisierungen auszuführen:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Klick-Eingabeparameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID der Aufgabe, die die Schaltfläche enthält |
| `customFieldId` | String! | ✅ Ja | ID des benutzerdefinierten Schaltflächenfeldes |

### Wichtig: API-Verhalten

**Alle Schaltflächenklicks über die API werden sofort ausgeführt**, unabhängig von `buttonType` oder `buttonConfirmText` Einstellungen. Diese Felder werden für UI-Clients gespeichert, um Bestätigungsdialoge zu implementieren, aber die API selbst:

- Validiert keinen Bestätigungstext
- Erzwingt keine Bestätigungsanforderungen
- Führt die Schaltflächenaktion sofort aus, wenn sie aufgerufen wird

Die Bestätigung ist rein eine Sicherheitsfunktion der UI-Seite.

### Beispiel: Klicken auf verschiedene Schaltflächentypen

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Alle drei Mutationen oben führen die Schaltflächenaktion sofort aus, wenn sie über die API aufgerufen werden, und umgehen dabei alle Bestätigungsanforderungen.

## Antwortfelder

### Benutzerdefinierte Feldantwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für das benutzerdefinierte Feld |
| `name` | String! | Anzeigename der Schaltfläche |
| `type` | CustomFieldType! | Immer `BUTTON` für Schaltflächenfelder |
| `buttonType` | String | Einstellung des Bestätigungsverhaltens |
| `buttonConfirmText` | String | Erforderlicher Bestätigungstext (wenn harte Bestätigung verwendet wird) |
| `description` | String | Hilfetext für Benutzer |
| `required` | Boolean! | Ob das Feld erforderlich ist |
| `isActive` | Boolean! | Ob das Feld derzeit aktiv ist |
| `projectId` | String! | ID des Projekts, zu dem dieses Feld gehört |
| `createdAt` | DateTime! | Wann das Feld erstellt wurde |
| `updatedAt` | DateTime! | Wann das Feld zuletzt bearbeitet wurde |

## Wie Schaltflächenfelder funktionieren

### Automatisierungsintegration

Schaltflächenfelder sind so konzipiert, dass sie mit Blues Automatisierungssystem arbeiten:

1. **Erstellen Sie das Schaltflächenfeld** mit der obigen Mutation
2. **Konfigurieren Sie Automatisierungen**, die auf `CUSTOM_FIELD_BUTTON_CLICKED` Ereignisse hören
3. **Benutzer klicken auf die Schaltfläche** in der UI
4. **Automatisierungen führen** die konfigurierten Aktionen aus

### Ereignisfluss

Wenn eine Schaltfläche angeklickt wird:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Keine Datenspeicherung

Wichtig: Schaltflächenfelder speichern keine Wertdaten. Sie dienen rein als Aktionsauslöser. Jeder Klick:
- Generiert ein Ereignis
- Löst zugehörige Automatisierungen aus
- Protokolliert eine Aktion in der Aufgabenhistorie
- Ändert keinen Feldwert

## Erforderliche Berechtigungen

Benutzer benötigen geeignete Projektrollen, um Schaltflächenfelder zu erstellen und zu verwenden:

| Aktion | Erforderliche Rolle |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Fehlerantworten

### Berechtigung verweigert
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Benutzerdefiniertes Feld nicht gefunden
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

**Hinweis**: Die API gibt keine spezifischen Fehler für fehlende Automatisierungen oder Bestätigungsabweichungen zurück.

## Best Practices

### Namenskonventionen
- Verwenden Sie handlungsorientierte Namen: "Rechnung senden", "Bericht erstellen", "Team benachrichtigen"
- Seien Sie spezifisch, was die Schaltfläche tut
- Vermeiden Sie allgemeine Namen wie "Schaltfläche 1" oder "Hier klicken"

### Bestätigungseinstellungen
- Lassen Sie `buttonType` leer für sichere, umkehrbare Aktionen
- Setzen Sie `buttonType`, um das Bestätigungsverhalten für UI-Clients vorzuschlagen
- Verwenden Sie `buttonConfirmText`, um anzugeben, was Benutzer in UI-Bestätigungen eingeben sollten
- Denken Sie daran: Dies sind nur UI-Hinweise - API-Aufrufe werden immer sofort ausgeführt

### Automatisierungsdesign
- Halten Sie Schaltflächenaktionen auf einen einzigen Workflow fokussiert
- Geben Sie klares Feedback darüber, was nach dem Klicken passiert ist
- Erwägen Sie, Beschreibungstexte hinzuzufügen, um den Zweck der Schaltfläche zu erklären

## Häufige Anwendungsfälle

1. **Workflow-Übergänge**
   - "Als abgeschlossen markieren"
   - "Zur Genehmigung senden"
   - "Aufgabe archivieren"

2. **Externe Integrationen**
   - "Mit CRM synchronisieren"
   - "Rechnung erstellen"
   - "E-Mail-Update senden"

3. **Batch-Operationen**
   - "Alle Unteraufgaben aktualisieren"
   - "In Projekte kopieren"
   - "Vorlage anwenden"

4. **Berichtshandlungen**
   - "Bericht erstellen"
   - "Daten exportieren"
   - "Zusammenfassung erstellen"

## Einschränkungen

- Schaltflächen können keine Datenwerte speichern oder anzeigen
- Jede Schaltfläche kann nur Automatisierungen auslösen, keine direkten API-Aufrufe (Automatisierungen können jedoch HTTP-Anforderungsaktionen enthalten, um externe APIs oder Blues eigene APIs aufzurufen)
- Die Sichtbarkeit von Schaltflächen kann nicht bedingt gesteuert werden
- Maximal eine Automatisierungsausführung pro Klick (obwohl diese Automatisierung mehrere Aktionen auslösen kann)

## Verwandte Ressourcen

- [Automatisierungs-API](/api/automations/index) - Konfigurieren Sie Aktionen, die durch Schaltflächen ausgelöst werden
- [Überblick über benutzerdefinierte Felder](/custom-fields/list-custom-fields) - Allgemeine Konzepte zu benutzerdefinierten Feldern