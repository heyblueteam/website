---
title: Checkbox benutzerdefiniertes Feld
description: Erstellen Sie boolesche Checkboxfelder für Ja/Nein- oder Wahr/Falsch-Daten
---

Checkbox benutzerdefinierte Felder bieten eine einfache boolesche (wahr/falsch) Eingabe für Aufgaben. Sie sind perfekt für binäre Entscheidungen, Statusanzeigen oder um zu verfolgen, ob etwas abgeschlossen wurde.

## Einfaches Beispiel

Erstellen Sie ein einfaches Checkboxfeld:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Checkboxfeld mit Beschreibung und Validierung:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ Ja | Anzeigename der Checkbox |
| `type` | CustomFieldType! | ✅ Ja | Muss `CHECKBOX` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

## Festlegen von Checkboxwerten

Um einen Checkboxwert für eine Aufgabe festzulegen oder zu aktualisieren:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Um eine Checkbox zu deaktivieren:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID der zu aktualisierenden Aufgabe |
| `customFieldId` | String! | ✅ Ja | ID des benutzerdefinierten Checkboxfelds |
| `checked` | Boolean | Nein | Wahr, um zu aktivieren, falsch, um zu deaktivieren |

## Erstellen von Aufgaben mit Checkboxwerten

Beim Erstellen einer neuen Aufgabe mit Checkboxwerten:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Akzeptierte String-Werte

Beim Erstellen von Aufgaben müssen Checkboxwerte als Strings übergeben werden:

| String-Wert | Ergebnis |
|--------------|---------|
| `"true"` | ✅ Aktiviert (groß-/kleinschreibung beachten) |
| `"1"` | ✅ Aktiviert |
| `"checked"` | ✅ Aktiviert (groß-/kleinschreibung beachten) |
| Any other value | ❌ Deaktiviert |

**Hinweis**: Stringvergleiche während der Erstellung von Aufgaben sind groß-/kleinschreibungsempfindlich. Die Werte müssen genau mit `"true"`, `"1"` oder `"checked"` übereinstimmen, um einen aktivierten Zustand zu ergeben.

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | ID! | Eindeutige Kennung für den Feldwert |
| `uid` | String! | Alternative eindeutige Kennung |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `checked` | Boolean | Der Checkboxzustand (wahr/falsch/null) |
| `todo` | Todo! | Die Aufgabe, zu der dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

## Automatisierungsintegration

Checkboxfelder lösen verschiedene Automatisierungsereignisse basierend auf Zustandsänderungen aus:

| Aktion | Ereignis ausgelöst | Beschreibung |
|--------|-------------------|-------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Ausgelöst, wenn die Checkbox aktiviert wird |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Ausgelöst, wenn die Checkbox deaktiviert wird |

Dies ermöglicht es Ihnen, Automatisierungen zu erstellen, die auf Änderungen des Checkboxzustands reagieren, wie zum Beispiel:
- Benachrichtigungen senden, wenn Elemente genehmigt werden
- Aufgaben verschieben, wenn Überprüfungs-Checkboxen aktiviert sind
- Verwandte Felder basierend auf Checkboxzuständen aktualisieren

## Datenimport/-export

### Importieren von Checkboxwerten

Beim Importieren von Daten über CSV oder andere Formate:
- `"true"`, `"yes"` → Aktiviert (groß-/kleinschreibung nicht beachten)
- Jeder andere Wert (einschließlich `"false"`, `"no"`, `"0"`, leer) → Deaktiviert

### Exportieren von Checkboxwerten

Beim Exportieren von Daten:
- Aktivierte Kästchen werden als `"X"` exportiert
- Deaktivierte Kästchen werden als leerer String `""` exportiert

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|---------------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Fehlerantworten

### Ungültiger Werttyp
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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

## Best Practices

### Namenskonventionen
- Verwenden Sie klare, handlungsorientierte Namen: "Genehmigt", "Überprüft", "Ist abgeschlossen"
- Vermeiden Sie negative Namen, die Benutzer verwirren: Bevorzugen Sie "Ist aktiv" gegenüber "Ist inaktiv"
- Seien Sie spezifisch, was die Checkbox darstellt

### Wann man Checkboxen verwenden sollte
- **Binäre Entscheidungen**: Ja/Nein, Wahr/Falsch, Erledigt/Nicht erledigt
- **Statusanzeigen**: Genehmigt, Überprüft, Veröffentlicht
- **Feature-Flags**: Hat priorisierten Support, Benötigt Unterschrift
- **Einfache Verfolgung**: E-Mail gesendet, Rechnung bezahlt, Artikel versendet

### Wann man Checkboxen NICHT verwenden sollte
- Wenn Sie mehr als zwei Optionen benötigen (verwenden Sie stattdessen SELECT_SINGLE)
- Für numerische oder Textdaten (verwenden Sie NUMBER- oder TEXT-Felder)
- Wenn Sie verfolgen müssen, wer es aktiviert hat oder wann (verwenden Sie Audit-Protokolle)

## Häufige Anwendungsfälle

1. **Genehmigungs-Workflows**
   - "Manager genehmigt"
   - "Kundenfreigabe"
   - "Rechtsprüfung abgeschlossen"

2. **Aufgabenverwaltung**
   - "Ist blockiert"
   - "Bereit zur Überprüfung"
   - "Hohe Priorität"

3. **Qualitätskontrolle**
   - "QA bestanden"
   - "Dokumentation abgeschlossen"
   - "Tests geschrieben"

4. **Administrative Flags**
   - "Rechnung gesendet"
   - "Vertrag unterschrieben"
   - "Nachverfolgung erforderlich"

## Einschränkungen

- Checkboxfelder können nur wahr/falsch Werte speichern (kein Tri-State oder null nach der ersten Festlegung)
- Keine Standardwertkonfiguration (beginnt immer als null, bis sie festgelegt wird)
- Kann keine zusätzlichen Metadaten speichern, wie wer es aktiviert hat oder wann
- Keine bedingte Sichtbarkeit basierend auf anderen Feldwerten

## Verwandte Ressourcen

- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte zu benutzerdefinierten Feldern
- [Automatisierungs-API](/api/automations) - Erstellen Sie Automatisierungen, die durch Änderungen an Checkboxen ausgelöst werden