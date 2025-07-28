---
title: Valutaomvandling Anpassat Fält
description: Skapa fält som automatiskt omvandlar valutavärden med hjälp av realtids växelkurser
---

Valutaomvandling anpassade fält omvandlar automatiskt värden från ett källvaluta fält till olika målvalutor med hjälp av realtids växelkurser. Dessa fält uppdateras automatiskt när källvaluta värdet ändras.

Växelkurserna tillhandahålls av [Frankfurter API](https://github.com/hakanensari/frankfurter), en öppen källkodstjänst som spårar referensväxelkurser publicerade av [Europeiska centralbanken](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Detta säkerställer noggranna, pålitliga och aktuella valutaomvandlingar för dina internationella affärsbehov.

## Grundläggande Exempel

Skapa ett enkelt valutaomvandlingsfält:

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

## Avancerat Exempel

Skapa ett omvandlingsfält med ett specifikt datum för historiska kurser:

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

## Fullständig Installationsprocess

Att ställa in ett valutaomvandlingsfält kräver tre steg:

### Steg 1: Skapa ett Källvaluta Fält

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

### Steg 2: Skapa CURRENCY_CONVERSION Fältet

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

### Steg 3: Skapa Omvandlingsalternativ

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

## Inmatningsparametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för omvandlingsfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | Nej | ID för källvaluta fältet att omvandla från |
| `conversionDateType` | String | Nej | Dato-strategi för växelkurser (se nedan) |
| `conversionDate` | String | Nej | Datumsträng för omvandling (baserat på conversionDateType) |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: Anpassade fält är automatiskt kopplade till projektet baserat på användarens aktuella projektkontext. Ingen `projectId` parameter krävs.

### Omvandlingsdatumstyper

| Typ | Beskrivning | conversionDate Parameter |
|------|-------------|-------------------------|
| `currentDate` | Använder realtids växelkurser | Inte nödvändig |
| `specificDate` | Använder kurser från ett fast datum | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Använder datum från ett annat fält | "todoDueDate" or DATE field ID |

## Skapa Omvandlingsalternativ

Omvandlingsalternativ definierar vilka valutapar som kan omvandlas:

### CreateCustomFieldOptionInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `customFieldId` | String! | ✅ Ja | ID för CURRENCY_CONVERSION fältet |
| `title` | String! | ✅ Ja | Visningsnamn för detta omvandlingsalternativ |
| `currencyConversionFrom` | String! | ✅ Ja | Källvaluta kod eller "Any" |
| `currencyConversionTo` | String! | ✅ Ja | Målvaluta kod |

### Använda "Any" som Källa

Det speciella värdet "Any" som `currencyConversionFrom` skapar ett fallback-alternativ:

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

Detta alternativ kommer att användas när ingen specifik valuta-par matchning hittas.

## Hur Automatisk Omvandling Fungerar

1. **Värdeuppdatering**: När ett värde sätts i källvaluta fältet
2. **Alternativmatchning**: Systemet hittar matchande omvandlingsalternativ baserat på källvaluta
3. **Hämtning av kurs**: Hämtar växelkurs från Frankfurter API
4. **Beräkning**: Multiplicerar källbeloppet med växelkursen
5. **Lagring**: Sparar det omvandlade värdet med målvaluta kod

### Exempel Flöde

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

## Datumbaserade Omvandlingar

### Använda Aktuellt Datum

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

Omvandlingar uppdateras med aktuella växelkurser varje gång källvärdet ändras.

### Använda Specifikt Datum

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

Använder alltid växelkurser från det angivna datumet.

### Använda Datum från Fält

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

Använder datumet från ett annat fält (antingen todo förfallodatum eller ett DATUM anpassat fält).

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Definitionen av omvandlingsfältet |
| `number` | Float | Det omvandlade beloppet |
| `currency` | String | Målvaluta kod |
| `todo` | Todo! | Den post detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast uppdaterades |

## Växelkurser Källa

Blue använder **Frankfurter API** för växelkurser:
- Öppen källkod API som hostas av Europeiska centralbanken
- Uppdateras dagligen med officiella växelkurser
- Stöder historiska kurser tillbaka till 1999
- Gratis och pålitlig för affärsanvändning

## Felhantering

### Omvandlingsfel

När omvandlingen misslyckas (API-fel, ogiltig valuta, etc.):
- Det omvandlade värdet sätts till `0`
- Målvalutan lagras fortfarande
- Inga felmeddelanden ges till användaren

### Vanliga Scenarier

| Scenario | Resultat |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Inga matchande alternativ | Uses "Any" option if available |
| Missing source value | Ingen omvandling utförd |

## Nödvändiga Behörigheter

Hantera anpassade fält kräver projekt-nivå åtkomst:

| Roll | Kan Skapa/Uppdatera Fält |
|------|-------------------------|
| `OWNER` | ✅ Ja |
| `ADMIN` | ✅ Ja |
| `MEMBER` | ❌ Nej |
| `CLIENT` | ❌ Nej |

Visningsbehörigheter för omvandlade värden följer standardregler för poståtkomst.

## Bästa Praxis

### Alternativ Konfiguration
- Skapa specifika valutapar för vanliga omvandlingar
- Lägg till ett "Any" fallback-alternativ för flexibilitet
- Använd beskrivande titlar för alternativ

### Val av Datumstrategi
- Använd `currentDate` för live finansiell spårning
- Använd `specificDate` för historisk rapportering
- Använd `fromDateField` för transaktionsspecifika kurser

### Prestandaöverväganden
- Flera omvandlingsfält uppdateras parallellt
- API-anrop görs endast när källvärdet ändras
- Samma-valuta omvandlingar hoppar över API-anrop

## Vanliga Användningsfall

1. **Flervaluta Projekt**
   - Spåra projektkostnader i lokala valutor
   - Rapportera total budget i företagsvaluta
   - Jämföra värden över regioner

2. **Internationell Försäljning**
   - Omvandla affärsvärden till rapporteringsvaluta
   - Spåra intäkter i flera valutor
   - Historisk omvandling för avslutade affärer

3. **Finansiell Rapportering**
   - Periodslut valutaomvandlingar
   - Konsoliderade finansiella rapporter
   - Budget vs. verklighet i lokal valuta

4. **Kontrakthantering**
   - Omvandla kontraktsvärden vid signeringsdatum
   - Spåra betalningsplaner i flera valutor
   - Valutariskbedömning

## Begränsningar

- Ingen support för kryptovalutaomvandlingar
- Kan inte ställa in omvandlade värden manuellt (alltid beräknade)
- Fast 2 decimalers precision för alla omvandlade belopp
- Ingen support för anpassade växelkurser
- Ingen caching av växelkurser (färskt API-anrop för varje omvandling)
- Beror på Frankfurter API tillgänglighet

## Relaterade Resurser

- [Valutafält](/api/custom-fields/currency) - Källfält för omvandlingar
- [Datumfält](/api/custom-fields/date) - För datumbaserade omvandlingar
- [Formelfält](/api/custom-fields/formula) - Alternativa beräkningar
- [Översikt över Anpassade Fält](/custom-fields/list-custom-fields) - Allmänna koncept