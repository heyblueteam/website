---
title: Valuta Conversie Aangepast Veld
description: Maak velden die automatisch valutawaarden converteren met behulp van realtime wisselkoersen
---

Valuta conversie aangepaste velden converteren automatisch waarden van een bron VALUTA-veld naar verschillende doelvaluta's met behulp van realtime wisselkoersen. Deze velden worden automatisch bijgewerkt wanneer de waarde van de bronvaluta verandert.

De conversiekoersen worden geleverd door de [Frankfurter API](https://github.com/hakanensari/frankfurter), een open-source service die referentiewisselkoersen bijhoudt die zijn gepubliceerd door de [Europese Centrale Bank](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Dit zorgt voor nauwkeurige, betrouwbare en actuele valutaconversies voor uw internationale zakelijke behoeften.

## Basisvoorbeeld

Maak een eenvoudig valutaconversieveld:

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

## Geavanceerd Voorbeeld

Maak een conversieveld met een specifieke datum voor historische tarieven:

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

## Volledige Installatieproces

Het instellen van een valutaconversieveld vereist drie stappen:

### Stap 1: Maak een Bron VALUTA Veld

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

### Stap 2: Maak het VALUTA_CONVERSIE Veld

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

### Stap 3: Maak Conversieopties

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

## Invoervelden

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het conversieveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | Nee | ID van het bron VALUTA-veld om van te converteren |
| `conversionDateType` | String | Nee | Datumstrategie voor wisselkoersen (zie hieronder) |
| `conversionDate` | String | Nee | Datumstring voor conversie (gebaseerd op conversionDateType) |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |

**Opmerking**: Aangepaste velden zijn automatisch gekoppeld aan het project op basis van de huidige projectcontext van de gebruiker. Geen `projectId` parameter is vereist.

### Conversiedatatype

| Type | Beschrijving | conversionDate Parameter |
|------|--------------|-------------------------|
| `currentDate` | Gebruikt realtime wisselkoersen | Niet vereist |
| `specificDate` | Gebruikt tarieven van een vaste datum | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Gebruikt datum van een ander veld | "todoDueDate" or DATE field ID |

## Conversieopties maken

Conversieopties definiëren welke valutaparen kunnen worden geconverteerd:

### CreateCustomFieldOptionInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `customFieldId` | String! | ✅ Ja | ID van het VALUTA_CONVERSIE veld |
| `title` | String! | ✅ Ja | Weergavenaam voor deze conversieoptie |
| `currencyConversionFrom` | String! | ✅ Ja | Bronvalutacode of "Any" |
| `currencyConversionTo` | String! | ✅ Ja | Doelvalutacode |

### "Any" gebruiken als Bron

De speciale waarde "Any" als `currencyConversionFrom` creëert een fallback-optie:

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

Deze optie zal worden gebruikt wanneer er geen specifieke valutapaarovereenkomst wordt gevonden.

## Hoe Automatische Conversie Werkt

1. **Waarde-update**: Wanneer een waarde wordt ingesteld in het bron VALUTA-veld
2. **Optie-overeenkomst**: Systeem vindt overeenkomende conversieoptie op basis van bronvaluta
3. **Tarief ophalen**: Haalt wisselkoers op van Frankfurter API
4. **Berekening**: Vermenigvuldigt het bronbedrag met de wisselkoers
5. **Opslag**: Slaat de geconverteerde waarde op met de doelvalutacode

### Voorbeeldstroom

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

## Datum-gebaseerde Conversies

### Huidige Datum Gebruiken

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

Conversies worden bijgewerkt met de huidige wisselkoersen telkens wanneer de bronwaarde verandert.

### Specifieke Datum Gebruiken

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

Gebruikt altijd wisselkoersen van de opgegeven datum.

### Datum van Veld Gebruiken

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

Gebruikt de datum van een ander veld (ofwel todo vervaldatum of een DATUM aangepast veld).

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het conversieveld |
| `number` | Float | Het geconverteerde bedrag |
| `currency` | String | De doelvalutacode |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is bijgewerkt |

## Wisselkoersbron

Blue gebruikt de **Frankfurter API** voor wisselkoersen:
- Open-source API gehost door de Europese Centrale Bank
- Dagelijks bijgewerkt met officiële wisselkoersen
- Ondersteunt historische tarieven terug tot 1999
- Gratis en betrouwbaar voor zakelijk gebruik

## Foutafhandeling

### Conversiefouten

Wanneer conversie mislukt (API-fout, ongeldige valuta, enz.):
- De geconverteerde waarde wordt ingesteld op `0`
- De doelvaluta wordt nog steeds opgeslagen
- Er wordt geen fout aan de gebruiker weergegeven

### Veelvoorkomende Scenario's

| Scenario | Resultaat |
|----------|-----------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Geen overeenkomende optie | Uses "Any" option if available |
| Missing source value | Geen conversie uitgevoerd |

## Vereiste Machtigingen

Beheer van aangepaste velden vereist toegang op projectniveau:

| Rol | Kan Velden Aanmaken/Bijwerken |
|-----|-------------------------------|
| `OWNER` | ✅ Ja |
| `ADMIN` | ✅ Ja |
| `MEMBER` | ❌ Nee |
| `CLIENT` | ❌ Nee |

Bekijk machtigingen voor geconverteerde waarden volgen de standaard recordtoegangsregels.

## Beste Praktijken

### Optieconfiguratie
- Maak specifieke valutaparen voor veelvoorkomende conversies
- Voeg een "Any" fallback-optie toe voor flexibiliteit
- Gebruik beschrijvende titels voor opties

### Selectie van Datumstrategie
- Gebruik `currentDate` voor live financiële tracking
- Gebruik `specificDate` voor historische rapportage
- Gebruik `fromDateField` voor transactie-specifieke tarieven

### Prestatieoverwegingen
- Meerdere conversievelden worden parallel bijgewerkt
- API-aanroepen worden alleen gedaan wanneer de bronwaarde verandert
- Conversies met dezelfde valuta overslaan API-aanroepen

## Veelvoorkomende Gebruiksscenario's

1. **Multi-Valuta Projecten**
   - Volg projectkosten in lokale valuta's
   - Rapporteer totale budget in bedrijfsvaluta
   - Vergelijk waarden tussen regio's

2. **Internationale Verkoop**
   - Converteer dealwaarden naar rapportagevaluta
   - Volg inkomsten in meerdere valuta's
   - Historische conversie voor gesloten deals

3. **Financiële Rapportage**
   - Einde-periode valutaconversies
   - Geconsolideerde financiële overzichten
   - Budget versus werkelijke waarde in lokale valuta

4. **Contractbeheer**
   - Converteer contractwaarden op de datum van ondertekening
   - Volg betalingsschema's in meerdere valuta's
   - Valutarisico-evaluatie

## Beperkingen

- Geen ondersteuning voor cryptocurrency-conversies
- Geconverteerde waarden kunnen niet handmatig worden ingesteld (altijd berekend)
- Vaste precisie van 2 decimalen voor alle geconverteerde bedragen
- Geen ondersteuning voor aangepaste wisselkoersen
- Geen caching van wisselkoersen (verse API-aanroep voor elke conversie)
- Afhankelijk van de beschikbaarheid van de Frankfurter API

## Gerelateerde Bronnen

- [Valuta Velden](/api/custom-fields/currency) - Bronvelden voor conversies
- [Datum Velden](/api/custom-fields/date) - Voor datum-gebaseerde conversies
- [Formule Velden](/api/custom-fields/formula) - Alternatieve berekeningen
- [Overzicht van Aangepaste Velden](/custom-fields/list-custom-fields) - Algemene concepten