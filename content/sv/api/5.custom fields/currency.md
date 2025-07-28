---
title: Valuta Anpassat Fält
description: Skapa valutafält för att spåra monetära värden med korrekt formatering och validering
---

Valuta anpassade fält gör att du kan lagra och hantera monetära värden med tillhörande valutakoder. Fältet stöder 72 olika valutor inklusive stora fiat-valutor och kryptovalutor, med automatisk formatering och valfria min/max begränsningar.

## Grundläggande Exempel

Skapa ett enkelt valutafält:

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

## Avancerat Exempel

Skapa ett valutafält med valideringsbegränsningar:

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

## Indata Parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för valutafältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `CURRENCY` |
| `currency` | String | Nej | Standard valutakod (3-bokstavs ISO kod) |
| `min` | Float | Nej | Minimi tillåtet värde (lagras men inte tillämpas vid uppdateringar) |
| `max` | Float | Nej | Maximalt tillåtet värde (lagras men inte tillämpas vid uppdateringar) |
| `description` | String | Nej | Hjälptext som visas för användare |

**Obs**: Projektkontexten bestäms automatiskt från din autentisering. Du måste ha tillgång till projektet där du skapar fältet.

## Ställa in Valuta Värden

För att ställa in eller uppdatera ett valutavärde på en post:

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

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade valutafältet |
| `number` | Float! | ✅ Ja | Det monetära beloppet |
| `currency` | String! | ✅ Ja | 3-bokstavs valutakod |

## Skapa Poster med Valuta Värden

När du skapar en ny post med valutavärden:

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

### Indataformat för Skapande

När du skapar poster, passerar valutavärden på ett annat sätt:

| Parameter | Typ | Beskrivning |
|-----------|------|-------------|
| `customFieldId` | String! | ID för valutafältet |
| `value` | String! | Belopp som en sträng (t.ex. "1500.50") |
| `currency` | String! | 3-bokstavs valutakod |

## Stödda Valutor

Blue stöder 72 valutor inklusive 70 fiat-valutor och 2 kryptovalutor:

### Fiat Valutor

#### Amerika
| Valuta | Kod | Namn |
|--------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Venezuelansk Bolívar Soberano |
| Boliviansk Boliviano | `BOB` | Boliviansk Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europa
| Valuta | Kod | Namn |
|--------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Norsk Krona | `NOK` | Norsk Krona |
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

#### Asien-Stillahavsområdet
| Valuta | Kod | Namn |
|--------|------|------|
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

#### Mellanöstern & Afrika
| Valuta | Kod | Namn |
|--------|------|------|
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

### Kryptovalutor
| Valuta | Kod |
|--------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `number` | Float | Det monetära beloppet |
| `currency` | String | Den 3-bokstavs valutakoden |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

## Valutaformatering

Systemet formaterar automatiskt valutavärden baserat på lokal:

- **Symbolplacering**: Korrekt placerar valutasymboler (före/efter)
- **Decimalavskiljare**: Använder lokal-specifika avskiljare (. eller ,)
- **Tusentalsavskiljare**: Tillämpa lämplig gruppering
- **Decimaler**: Visar 0-2 decimaler baserat på beloppet
- **Speciell hantering**: USD/CAD visar valutakod prefix för tydlighet

### Formaterings Exempel

| Värde | Valuta | Visning |
|-------|--------|---------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Validering

### Beloppsvalidering
- Måste vara ett giltigt nummer
- Min/max begränsningar lagras med fältdefinitionen men tillämpas inte vid värdeuppdateringar
- Stöder upp till 2 decimaler för visning (full precision lagras internt)

### Valutakod Validering
- Måste vara en av de 72 stödda valutakoderna
- Skiftlägeskänslig (använd versaler)
- Ogiltiga koder returnerar ett fel

## Integrationsfunktioner

### Formler
Valutafält kan användas i FORMULA anpassade fält för beräkningar:
- Summera flera valutafält
- Beräkna procent
- Utföra aritmetiska operationer

### Valutaomvandling
Använd CURRENCY_CONVERSION fält för att automatiskt omvandla mellan valutor (se [Valutaomvandlingsfält](/api/custom-fields/currency-conversion))

### Automatiseringar
Valutavärden kan utlösa automatiseringar baserat på:
- Beloppsgränser
- Valutatyp
- Värdeförändringar

## Nödvändiga Behörigheter

| Åtgärd | Nödvändig Behörighet |
|--------|----------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Obs**: Även om alla projektmedlemmar kan skapa anpassade fält, beror möjligheten att ställa in värden på rollbaserade behörigheter som är konfigurerade för varje fält.

## Fel Svar

### Ogiltigt Valutavärde
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

Detta fel uppstår när:
- Valutakoden inte är en av de 72 stödda koderna
- Nummerformatet är ogiltigt
- Värdet kan inte tolkas korrekt

### Anpassat Fält Hittades Inte
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

## Bästa Praxis

### Val av Valuta
- Ställ in en standardvaluta som matchar din primära marknad
- Använd ISO 4217 valutakoder konsekvent
- Tänk på användarens plats när du väljer standarder

### Värdebegränsningar
- Ställ in rimliga min/max värden för att förhindra datainmatningsfel
- Använd 0 som minimum för fält som inte ska vara negativa
- Tänk på ditt användningsfall när du sätter maximala värden

### Multi-Valuta Projekt
- Använd konsekvent basvaluta för rapportering
- Implementera CURRENCY_CONVERSION fält för automatisk omvandling
- Dokumentera vilken valuta som ska användas för varje fält

## Vanliga Användningsfall

1. **Projektbudgetering**
   - Spåra projektbudget
   - Kostnadsuppskattningar
   - Utgiftsspårning

2. **Försäljning & Affärer**
   - Affärsvärden
   - Kontraktsbelopp
   - Intäktsuppföljning

3. **Finansiell Planering**
   - Investeringsbelopp
   - Finansieringsrundor
   - Finansiella mål

4. **Internationell Verksamhet**
   - Multi-valuta prissättning
   - Valutaomvandling
   - Gränsöverskridande transaktioner

## Begränsningar

- Maximalt 2 decimaler för visning (även om mer precision lagras)
- Ingen inbyggd valutaomvandling i standard CURRENCY fält
- Kan inte blanda valutor i ett enda fältvärde
- Inga automatiska växelkursuppdateringar (använd CURRENCY_CONVERSION för detta)
- Valutasymboler är inte anpassningsbara

## Relaterade Resurser

- [Valutaomvandlingsfält](/api/custom-fields/currency-conversion) - Automatisk valutaomvandling
- [Nummerfält](/api/custom-fields/number) - För icke-monetära numeriska värden
- [Formelfält](/api/custom-fields/formula) - Beräkna med valutavärden
- [Lista Anpassade Fält](/api/custom-fields/list-custom-fields) - Fråga alla anpassade fält i ett projekt