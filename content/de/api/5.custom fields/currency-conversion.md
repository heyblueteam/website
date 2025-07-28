---
title: Währungsumrechnungsbenutzerfeld
description: Erstellen Sie Felder, die Währungswerte automatisch mit Echtzeit-Wechselkursen umrechnen
---

Währungsumrechnungsbenutzerfelder rechnen automatisch Werte aus einem Quellwährungsfeld in verschiedene Zielwährungen um, indem sie Echtzeit-Wechselkurse verwenden. Diese Felder werden automatisch aktualisiert, wenn sich der Wert der Quellwährung ändert.

Die Umrechnungskurse werden von der [Frankfurter API](https://github.com/hakanensari/frankfurter) bereitgestellt, einem Open-Source-Dienst, der Referenzwechselkurse verfolgt, die von der [Europäischen Zentralbank](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html) veröffentlicht werden. Dies gewährleistet genaue, zuverlässige und aktuelle Währungsumrechnungen für Ihre internationalen Geschäftsbedürfnisse.

## Einfaches Beispiel

Erstellen Sie ein einfaches Währungsumrechnungsfeld:

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Umrechnungsfeld mit einem bestimmten Datum für historische Kurse:

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## Vollständiger Einrichtungsprozess

Die Einrichtung eines Währungsumrechnungsfelds erfordert drei Schritte:

### Schritt 1: Erstellen Sie ein Quellwährungsfeld

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### Schritt 2: Erstellen Sie das WÄHRUNGSUMRECHNUNGSFELD

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### Schritt 3: Erstellen Sie Umrechnungsoptionen

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Umrechnungsfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `CURRENCY_CONVERSION` sein |
| `currencyFieldId` | String | Nein | ID des Quellwährungsfelds, von dem umgerechnet werden soll |
| `conversionDateType` | String | Nein | Datumsstrategie für Wechselkurse (siehe unten) |
| `conversionDate` | String | Nein | Datumszeichenfolge für die Umrechnung (basierend auf conversionDateType) |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Benutzerfelder sind automatisch mit dem Projekt verknüpft, basierend auf dem aktuellen Projektkontext des Benutzers. Kein `projectId`-Parameter ist erforderlich.

### Umrechnungsdatentypen

| Typ | Beschreibung | conversionDate-Parameter |
|------|-------------|-------------------------|
| `currentDate` | Verwendet Echtzeit-Wechselkurse | Nicht erforderlich |
| `specificDate` | Verwendet Kurse von einem festen Datum | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Verwendet Datum aus einem anderen Feld | "todoDueDate" or DATE field ID |

## Erstellen von Umrechnungsoptionen

Umrechnungsoptionen definieren, welche Währungspaare umgerechnet werden können:

### CreateCustomFieldOptionInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `customFieldId` | String! | ✅ Ja | ID des WÄHRUNGSUMRECHNUNGSFELDES |
| `title` | String! | ✅ Ja | Anzeigename für diese Umrechnungsoption |
| `currencyConversionFrom` | String! | ✅ Ja | Quellwährungs-Code oder "Any" |
| `currencyConversionTo` | String! | ✅ Ja | Zielwährungs-Code |

### Verwendung von "Any" als Quelle

Der spezielle Wert "Any" als `currencyConversionFrom` erstellt eine Fallback-Option:

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

Diese Option wird verwendet, wenn keine spezifische Übereinstimmung für das Währungspaar gefunden wird.

## So funktioniert die automatische Umrechnung

1. **Wertaktualisierung**: Wenn ein Wert im Quellwährungsfeld festgelegt wird
2. **Optionenabgleich**: Das System findet die passende Umrechnungsoption basierend auf der Quellwährung
3. **Kursabfrage**: Ruft den Wechselkurs von der Frankfurter API ab
4. **Berechnung**: Multipliziert den Quellbetrag mit dem Wechselkurs
5. **Speicherung**: Speichert den umgerechneten Wert mit dem Zielwährungscode

### Beispielablauf

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## Datumsgestützte Umrechnungen

### Verwendung des aktuellen Datums

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

Umrechnungen werden mit den aktuellen Wechselkursen aktualisiert, jedes Mal, wenn sich der Quellwert ändert.

### Verwendung eines bestimmten Datums

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

Verwendet immer die Wechselkurse vom angegebenen Datum.

### Verwendung des Datums aus einem Feld

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

Verwendet das Datum aus einem anderen Feld (entweder Fälligkeitsdatum oder ein DATUM-Benutzerfeld).

## Antwortfelder

### TodoCustomField-Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des Umrechnungsfeldes |
| `number` | Float | Der umgerechnete Betrag |
| `currency` | String | Der Zielwährungscode |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt aktualisiert wurde |

## Wechselkursquelle

Blue verwendet die **Frankfurter API** für Wechselkurse:
- Open-Source-API, die von der Europäischen Zentralbank gehostet wird
- Tägliche Aktualisierungen mit offiziellen Wechselkursen
- Unterstützt historische Kurse bis 1999
- Kostenlos und zuverlässig für geschäftliche Zwecke

## Fehlerbehandlung

### Umrechnungsfehler

Wenn die Umrechnung fehlschlägt (API-Fehler, ungültige Währung usw.):
- Der umgerechnete Wert wird auf `0` gesetzt
- Die Zielwährung wird weiterhin gespeichert
- Es wird kein Fehler an den Benutzer ausgegeben

### Häufige Szenarien

| Szenario | Ergebnis |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Keine passende Option | Uses "Any" option if available |
| Missing source value | Keine Umrechnung durchgeführt |

## Erforderliche Berechtigungen

Die Verwaltung von Benutzerfeldern erfordert Projektzugriff:

| Rolle | Kann Felder erstellen/aktualisieren |
|------|-------------------------|
| `OWNER` | ✅ Ja |
| `ADMIN` | ✅ Ja |
| `MEMBER` | ❌ Nein |
| `CLIENT` | ❌ Nein |

Die Berechtigungen zum Anzeigen von umgerechneten Werten folgen den Standardzugriffsregeln für Datensätze.

## Best Practices

### Optionen konfigurieren
- Erstellen Sie spezifische Währungspaare für häufige Umrechnungen
- Fügen Sie eine "Any"-Fallback-Option für Flexibilität hinzu
- Verwenden Sie beschreibende Titel für Optionen

### Auswahl der Datumsstrategie
- Verwenden Sie `currentDate` für die Echtzeit-Finanzverfolgung
- Verwenden Sie `specificDate` für historische Berichterstattung
- Verwenden Sie `fromDateField` für transaktionsspezifische Kurse

### Leistungsüberlegungen
- Mehrere Umrechnungsfelder werden parallel aktualisiert
- API-Aufrufe werden nur durchgeführt, wenn sich der Quellwert ändert
- Umrechnungen in derselben Währung überspringen API-Aufrufe

## Häufige Anwendungsfälle

1. **Multi-Währungsprojekte**
   - Verfolgen Sie Projektkosten in lokalen Währungen
   - Berichten Sie über das Gesamtbudget in Unternehmenswährung
   - Vergleichen Sie Werte über Regionen hinweg

2. **Internationale Verkäufe**
   - Wandeln Sie Vertragswerte in die Berichtswährung um
   - Verfolgen Sie Einnahmen in mehreren Währungen
   - Historische Umrechnung für abgeschlossene Verträge

3. **Finanzberichterstattung**
   - Währungsumrechnungen zum Periodenende
   - Konsolidierte Finanzberichte
   - Budget vs. Ist in lokaler Währung

4. **Vertragsmanagement**
   - Wandeln Sie Vertragswerte zum Zeitpunkt der Unterzeichnung um
   - Verfolgen Sie Zahlungspläne in mehreren Währungen
   - Bewertung des Währungsrisikos

## Einschränkungen

- Keine Unterstützung für Kryptowährungsumrechnungen
- Umgerechnete Werte können nicht manuell festgelegt werden (immer berechnet)
- Feste Genauigkeit von 2 Dezimalstellen für alle umgerechneten Beträge
- Keine Unterstützung für benutzerdefinierte Wechselkurse
- Kein Caching von Wechselkursen (frischer API-Aufruf für jede Umrechnung)
- Abhängig von der Verfügbarkeit der Frankfurter API

## Verwandte Ressourcen

- [Währungsfelder](/api/custom-fields/currency) - Quellfelder für Umrechnungen
- [Datumsfelder](/api/custom-fields/date) - Für datumsbasierte Umrechnungen
- [Formel-Felder](/api/custom-fields/formula) - Alternative Berechnungen
- [Überblick über benutzerdefinierte Felder](/custom-fields/list-custom-fields) - Allgemeine Konzepte