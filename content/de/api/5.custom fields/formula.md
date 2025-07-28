---
title: Formel benutzerdefiniertes Feld
description: Erstellen Sie berechnete Felder, die automatisch Werte basierend auf anderen Daten berechnen
---

Formel benutzerdefinierte Felder werden für Diagramm- und Dashboard-Berechnungen innerhalb von Blue verwendet. Sie definieren Aggregationsfunktionen (SUMME, DURCHSCHNITT, ANZAHL usw.), die auf benutzerdefinierten Felddaten arbeiten, um berechnete Kennzahlen in Diagrammen anzuzeigen. Formeln werden nicht auf der Ebene einzelner Todos berechnet, sondern aggregieren Daten über mehrere Datensätze hinweg zu Visualisierungszwecken.

## Grundlegendes Beispiel

Erstellen Sie ein Formel-Feld für Diagrammberechnungen:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie eine Währungsformel mit komplexen Berechnungen:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Formel-Feldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `FORMULA` sein |
| `projectId` | String! | ✅ Ja | Die Projekt-ID, in der dieses Feld erstellt wird |
| `formula` | JSON | Nein | Formeldefinition für Diagrammberechnungen |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

### Formelstruktur

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Unterstützte Funktionen

### Diagramm-Aggregationsfunktionen

Formel-Felder unterstützen die folgenden Aggregationsfunktionen für Diagrammberechnungen:

| Funktion | Beschreibung | ChartFunction Enum |
|----------|-------------|-------------------|
| `SUM` | Summe aller Werte | `SUM` |
| `AVERAGE` | Durchschnitt der numerischen Werte | `AVERAGE` |
| `AVERAGEA` | Durchschnitt ohne Nullen und Nullwerte | `AVERAGEA` |
| `COUNT` | Anzahl der Werte | `COUNT` |
| `COUNTA` | Anzahl ohne Nullen und Nullwerte | `COUNTA` |
| `MAX` | Maximalwert | `MAX` |
| `MIN` | Minimalwert | `MIN` |

**Hinweis**: Diese Funktionen werden im `display.function` Feld verwendet und arbeiten mit aggregierten Daten für Diagrammvisualisierungen. Komplexe mathematische Ausdrücke oder Berechnungen auf Feldebene werden nicht unterstützt.

## Anzeigearten

### Zahlenanzeige

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Ergebnis: `1250.75`

### Währungsanzeige

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Ergebnis: `$1,250.75`

### Prozentanzeige

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Ergebnis: `87.5%`

## Bearbeiten von Formel-Feldern

Aktualisieren Sie vorhandene Formel-Felder:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Formelverarbeitung

### Kontext der Diagrammberechnung

Formel-Felder werden im Kontext von Diagrammsegmenten und Dashboards verarbeitet:
- Berechnungen erfolgen, wenn Diagramme gerendert oder aktualisiert werden
- Ergebnisse werden im `ChartSegment.formulaResult` als Dezimalwerte gespeichert
- Die Verarbeitung erfolgt über eine spezielle BullMQ-Warteschlange mit dem Namen 'formel'
- Updates werden an Dashboard-Abonnenten für Echtzeitaktualisierungen veröffentlicht

### Anzeigeformatierung

Die `getFormulaDisplayValue` Funktion formatiert die berechneten Ergebnisse basierend auf dem Anzeigetyp:
- **ZAHL**: Wird als einfache Zahl mit optionaler Genauigkeit angezeigt
- **PROZENT**: Fügt ein %-Suffix mit optionaler Genauigkeit hinzu  
- **WÄHRUNG**: Formatiert mit dem angegebenen Währungscode

## Speicherung der Formel Ergebnisse

Ergebnisse werden im `formulaResult` Feld gespeichert:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des Formel-Feldes |
| `number` | Float | Berechnetes numerisches Ergebnis |
| `formulaResult` | JSON | Vollständiges Ergebnis mit Anzeigeformatierung |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt berechnet wurde |

## Datenkontext

### Diagramm-Datenquelle

Formel-Felder arbeiten im Kontext der Diagramm-Datenquelle:
- Formeln aggregieren benutzerdefinierte Feldwerte über Todos in einem Projekt
- Die in `display.function` angegebene Aggregationsfunktion bestimmt die Berechnung
- Ergebnisse werden mit SQL-Aggregatfunktionen (avg, sum, count usw.) berechnet
- Berechnungen werden auf Datenbankebene zur Effizienz durchgeführt

## Häufige Formelbeispiele

### Gesamtes Budget (Diagrammanzeige)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Durchschnittliche Punktzahl (Diagrammanzeige)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Aufgabenanzahl (Diagrammanzeige)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Erforderliche Berechtigungen

Benutzerdefinierte Feldoperationen folgen den standardmäßigen rollenbasierten Berechtigungen:

| Aktion | Erforderliche Rolle |
|--------|---------------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Hinweis**: Die spezifischen erforderlichen Rollen hängen von der benutzerdefinierten Rollenkonfiguration Ihres Projekts ab. Es gibt keine speziellen Berechtigungskonstanten wie CUSTOM_FIELDS_CREATE.

## Fehlerbehandlung

### Validierungsfehler
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Benutzerdefiniertes Feld nicht gefunden
```json
{
  "errors": [{
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Best Practices

### Formelgestaltung
- Verwenden Sie klare, beschreibende Namen für Formel-Felder
- Fügen Sie Beschreibungen hinzu, die die Berechnungslogik erklären
- Testen Sie Formeln mit Beispieldaten vor der Bereitstellung
- Halten Sie Formeln einfach und lesbar

### Leistungsoptimierung
- Vermeiden Sie tief verschachtelte Formelabhängigkeiten
- Verwenden Sie spezifische Feldreferenzen anstelle von Platzhaltern
- Berücksichtigen Sie Caching-Strategien für komplexe Berechnungen
- Überwachen Sie die Formel-Leistung in großen Projekten

### Datenqualität
- Validieren Sie Quelldaten, bevor Sie sie in Formeln verwenden
- Gehen Sie angemessen mit leeren oder Nullwerten um
- Verwenden Sie die geeignete Genauigkeit für Anzeigearten
- Berücksichtigen Sie Randfälle in Berechnungen

## Häufige Anwendungsfälle

1. **Finanzverfolgung**
   - Budgetberechnungen
   - Gewinn-/Verlustrechnungen
   - Kostenanalysen
   - Umsatzprognosen

2. **Projektmanagement**
   - Abschlussprozentsätze
   - Ressourcenauslastung
   - Zeitplanberechnungen
   - Leistungskennzahlen

3. **Qualitätskontrolle**
   - Durchschnittliche Punktzahlen
   - Bestehen/Nichtbestehen-Raten
   - Qualitätskennzahlen
   - Compliance-Überwachung

4. **Business Intelligence**
   - KPI-Berechnungen
   - Trendanalysen
   - Vergleichskennzahlen
   - Dashboard-Werte

## Einschränkungen

- Formeln sind nur für Diagramm-/Dashboard-Aggregationen gedacht, nicht für Berechnungen auf Todo-Ebene
- Beschränkt auf die sieben unterstützten Aggregationsfunktionen (SUMME, DURCHSCHNITT usw.)
- Keine komplexen mathematischen Ausdrücke oder Feld-zu-Feld-Berechnungen
- Es können keine mehreren Felder in einer einzigen Formel referenziert werden
- Ergebnisse sind nur in Diagrammen und Dashboards sichtbar
- Das `logic` Feld dient nur für Anzeigetexte, nicht für tatsächliche Berechnungslogik

## Verwandte Ressourcen

- [Zahlenfelder](/api/5.custom%20fields/number) - Für statische numerische Werte
- [Währungsfelder](/api/5.custom%20fields/currency) - Für Geldwerte
- [Referenzfelder](/api/5.custom%20fields/reference) - Für projektübergreifende Daten
- [Lookup-Felder](/api/5.custom%20fields/lookup) - Für aggregierte Daten
- [Überblick über benutzerdefinierte Felder](/api/5.custom%20fields/2.list-custom-fields) - Allgemeine Konzepte