---
title: Währungsbenutzerdefiniertes Feld
description: Erstellen Sie Währungsfelder, um monetäre Werte mit ordnungsgemäßer Formatierung und Validierung zu verfolgen
---

Währungsbenutzerdefinierte Felder ermöglichen es Ihnen, monetäre Werte mit zugehörigen Währungscodes zu speichern und zu verwalten. Das Feld unterstützt 72 verschiedene Währungen, einschließlich wichtiger Fiat-Währungen und Kryptowährungen, mit automatischer Formatierung und optionalen Min-/Max-Beschränkungen.

## Einfaches Beispiel

Erstellen Sie ein einfaches Währungsfeld:

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Fortgeschrittenes Beispiel

Erstellen Sie ein Währungsfeld mit Validierungsbeschränkungen:

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Währungsfeldes |
| `type` | CustomFieldType! | ✅ Ja | Muss `CURRENCY` sein |
| `currency` | String | Nein | Standard-Währungscode (3-Buchstaben-ISO-Code) |
| `min` | Float | Nein | Minimal erlaubter Wert (gespeichert, aber bei Aktualisierungen nicht durchgesetzt) |
| `max` | Float | Nein | Maximal erlaubter Wert (gespeichert, aber bei Aktualisierungen nicht durchgesetzt) |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

**Hinweis**: Der Projektkontext wird automatisch aus Ihrer Authentifizierung bestimmt. Sie müssen Zugriff auf das Projekt haben, in dem Sie das Feld erstellen.

## Währungswerte festlegen

Um einen Währungswert in einem Datensatz festzulegen oder zu aktualisieren:

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### SetTodoCustomFieldInput-Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des Währungsbenutzerdefinierten Feldes |
| `number` | Float! | ✅ Ja | Der monetäre Betrag |
| `currency` | String! | ✅ Ja | 3-Buchstaben-Währungscode |

## Datensätze mit Währungswerten erstellen

Beim Erstellen eines neuen Datensatzes mit Währungswerten:

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### Eingabeformat für die Erstellung

Beim Erstellen von Datensätzen werden Währungswerte anders übergeben:

| Parameter | Typ | Beschreibung |
|-----------|------|-------------|
| `customFieldId` | String! | ID des Währungsfeldes |
| `value` | String! | Betrag als Zeichenfolge (z. B. "1500.50") |
| `currency` | String! | 3-Buchstaben-Währungscode |

## Unterstützte Währungen

Blue unterstützt 72 Währungen, darunter 70 Fiat-Währungen und 2 Kryptowährungen:

### Fiat-Währungen

#### Amerika
| Währung | Code | Name |
|----------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Venezolanischer Bolívar Soberano |
| Bolivianischer Boliviano | `BOB` | Bolivianischer Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europa
| Währung | Code | Name |
|----------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Norwegische Krone | `NOK` | Norwegische Krone |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Asien-Pazifik
| Währung | Code | Name |
|----------|------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### Naher Osten & Afrika
| Währung | Code | Name |
|----------|------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Kryptowährungen
| Währung | Code |
|----------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Antwortfelder

### TodoCustomField-Antwort

| Feld | Typ | Beschreibung |
|-------|------|-------------|
| `id` | String! | Eindeutiger Bezeichner für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Feldes |
| `number` | Float | Der monetäre Betrag |
| `currency` | String | Der 3-Buchstaben-Währungscode |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

## Währungsformatierung

Das System formatiert Währungswerte automatisch basierend auf der Gebietsschema:

- **Symbolplatzierung**: Positioniert Währungssymbole korrekt (vor/nach)
- **Dezimaltrennzeichen**: Verwendet gebietsabhängige Trennzeichen (. oder ,)
- **Tausendertrennzeichen**: Wendet geeignete Gruppierung an
- **Dezimalstellen**: Zeigt 0-2 Dezimalstellen basierend auf dem Betrag an
- **Besondere Handhabung**: USD/CAD zeigen zur Klarheit das Währungscode-Präfix

### Formatierungsbeispiele

| Wert | Währung | Anzeige |
|-------|----------|---------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Validierung

### Betrag Validierung
- Muss eine gültige Zahl sein
- Min-/Max-Beschränkungen werden mit der Felddefinition gespeichert, aber nicht bei Wertaktualisierungen durchgesetzt
- Unterstützt bis zu 2 Dezimalstellen zur Anzeige (volle Präzision intern gespeichert)

### Währungscode Validierung
- Muss einer der 72 unterstützten Währungscodes sein
- Groß-/Kleinschreibung ist wichtig (verwenden Sie Großbuchstaben)
- Ungültige Codes geben einen Fehler zurück

## Integrationsfunktionen

### Formeln
Währungsfelder können in FORMEL-Benutzerdefinierten Feldern für Berechnungen verwendet werden:
- Summieren Sie mehrere Währungsfelder
- Berechnen Sie Prozentsätze
- Führen Sie arithmetische Operationen durch

### Währungsumrechnung
Verwenden Sie WÄHRUNGSUMRECHNUNG-Felder, um automatisch zwischen Währungen umzuwandeln (siehe [Währungsumrechnungsfelder](/api/custom-fields/currency-conversion))

### Automatisierungen
Währungswerte können Automatisierungen basierend auf auslösen:
- Betragsgrenzen
- Währungstyp
- Wertänderungen

## Erforderliche Berechtigungen

| Aktion | Erforderliche Berechtigung |
|--------|-------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Hinweis**: Während jedes Projektmitglied benutzerdefinierte Felder erstellen kann, hängt die Fähigkeit, Werte festzulegen, von den rollenbasierten Berechtigungen ab, die für jedes Feld konfiguriert sind.

## Fehlermeldungen

### Ungültiger Währungswert
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

Dieser Fehler tritt auf, wenn:
- Der Währungscode nicht einer der 72 unterstützten Codes ist
- Das Zahlenformat ungültig ist
- Der Wert nicht korrekt geparst werden kann

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

## Beste Praktiken

### Währungsauswahl
- Setzen Sie eine Standardwährung, die Ihrem Hauptmarkt entspricht
- Verwenden Sie ISO 4217-Währungscodes konsequent
- Berücksichtigen Sie den Standort des Benutzers bei der Auswahl von Standardwerten

### Wertbeschränkungen
- Setzen Sie angemessene Min-/Max-Werte, um Dateneingabefehler zu vermeiden
- Verwenden Sie 0 als Minimum für Felder, die nicht negativ sein sollten
- Berücksichtigen Sie Ihren Anwendungsfall bei der Festlegung von Maximalwerten

### Projekte mit mehreren Währungen
- Verwenden Sie eine konsistente Basiswährung für Berichterstattung
- Implementieren Sie WÄHRUNGSUMRECHNUNG-Felder für automatische Umwandlung
- Dokumentieren Sie, welche Währung für jedes Feld verwendet werden soll

## Häufige Anwendungsfälle

1. **Projektbudgetierung**
   - Verfolgung des Projektbudgets
   - Kostenschätzungen
   - Ausgabenverfolgung

2. **Verkäufe & Deals**
   - Deal-Werte
   - Vertragsbeträge
   - Umsatzverfolgung

3. **Finanzplanung**
   - Investitionsbeträge
   - Finanzierungsrunden
   - Finanzziele

4. **Internationales Geschäft**
   - Preisgestaltung in mehreren Währungen
   - Verfolgung von Devisen
   - Grenzüberschreitende Transaktionen

## Einschränkungen

- Maximal 2 Dezimalstellen zur Anzeige (obwohl mehr Präzision gespeichert wird)
- Keine integrierte Währungsumrechnung in Standard-WÄHRUNGSfeldern
- Währungen können nicht in einem einzigen Feldwert gemischt werden
- Keine automatischen Wechselkursaktualisierungen (verwenden Sie dafür WÄHRUNGSUMRECHNUNG)
- Währungssymbole sind nicht anpassbar

## Verwandte Ressourcen

- [Währungsumrechnungsfelder](/api/custom-fields/currency-conversion) - Automatische Währungsumrechnung
- [Zahlenfelder](/api/custom-fields/number) - Für nicht-monetäre numerische Werte
- [Formel-Felder](/api/custom-fields/formula) - Berechnen mit Währungswerten
- [Listenbenutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Abfragen aller benutzerdefinierten Felder in einem Projekt