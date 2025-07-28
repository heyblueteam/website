---
title: Valuta Aangepast Veld
description: Maak valutavelden aan om monetaire waarden bij te houden met de juiste opmaak en validatie
---

Valuta aangepaste velden stellen je in staat om monetaire waarden op te slaan en te beheren met bijbehorende valutacodes. Het veld ondersteunt 72 verschillende valuta's, waaronder belangrijke fiat-valuta's en cryptocurrencies, met automatische opmaak en optionele min/max-beperkingen.

## Basisvoorbeeld

Maak een eenvoudig valutaveld aan:

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

## Geavanceerd Voorbeeld

Maak een valutaveld aan met validatiebeperkingen:

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

## Invoervelden

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het valutaveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `CURRENCY` |
| `currency` | String | Nee | Standaard valutacode (3-letterige ISO-code) |
| `min` | Float | Nee | Minimum toegestane waarde (opgeslagen maar niet afgedwongen bij updates) |
| `max` | Float | Nee | Maximum toegestane waarde (opgeslagen maar niet afgedwongen bij updates) |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

**Opmerking**: De projectcontext wordt automatisch bepaald op basis van je authenticatie. Je moet toegang hebben tot het project waarin je het veld aanmaakt.

## Instellen van Valutawaarden

Om een valutawaarde op een record in te stellen of bij te werken:

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

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het valuta aangepaste veld |
| `number` | Float! | ✅ Ja | Het monetaire bedrag |
| `currency` | String! | ✅ Ja | 3-letterige valutacode |

## Records Aanmaken met Valutawaarden

Bij het aanmaken van een nieuw record met valutawaarden:

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

### Invoervormaat voor Aanmaak

Bij het aanmaken van records worden valutawaarden anders doorgegeven:

| Parameter | Type | Beschrijving |
|-----------|------|-------------|
| `customFieldId` | String! | ID van het valutaveld |
| `value` | String! | Bedrag als een string (bijv. "1500.50") |
| `currency` | String! | 3-letterige valutacode |

## Ondersteunde Valuta's

Blue ondersteunt 72 valuta's, waaronder 70 fiat-valuta's en 2 cryptocurrencies:

### Fiat-valuta's

#### Amerika
| Valuta | Code | Naam |
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
| Venezuelan Bolívar | `VES` | Venezolaanse Bolívar Soberano |
| Bolivian Boliviano | `BOB` | Bolivian Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europa
| Valuta | Code | Naam |
|--------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Noorse Kroon | `NOK` | Noorse Kroon |
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

#### Azië-Pacific
| Valuta | Code | Naam |
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

#### Midden-Oosten & Afrika
| Valuta | Code | Naam |
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

### Cryptocurrencies
| Valuta | Code |
|--------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|-------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `number` | Float | Het monetaire bedrag |
| `currency` | String | De 3-letterige valutacode |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

## Valuta-opmaak

Het systeem formatteert valutawaarden automatisch op basis van de locale:

- **Plaatsing van symbolen**: Plaatst valutatekens correct (voor/achter)
- **Decimale scheidingstekens**: Gebruikt locale-specifieke scheidingstekens (. of ,)
- **Duizendtallen scheidingstekens**: Past geschikte groepering toe
- **Decimale plaatsen**: Toont 0-2 decimalen op basis van het bedrag
- **Speciale behandeling**: USD/CAD toont valutacode-prefix voor duidelijkheid

### Opmaakvoorbeelden

| Waarde | Valuta | Weergave |
|--------|--------|----------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Validatie

### Bedrag Validatie
- Moet een geldig getal zijn
- Min/max-beperkingen worden opgeslagen met de velddefinitie maar niet afgedwongen tijdens waarde-updates
- Ondersteunt tot 2 decimalen voor weergave (volledige precisie intern opgeslagen)

### Valutacode Validatie
- Moet een van de 72 ondersteunde valutacodes zijn
- Hoofdlettergevoelig (gebruik hoofdletters)
- Ongeldige codes geven een foutmelding terug

## Integratiefuncties

### Formules
Valutavelden kunnen worden gebruikt in FORMULE aangepaste velden voor berekeningen:
- Sommeer meerdere valutavelden
- Bereken percentages
- Voer rekenkundige bewerkingen uit

### Valutaconversie
Gebruik VALUTA_CONVERSIE velden om automatisch tussen valuta's te converteren (zie [Valutaconversie Velden](/api/custom-fields/currency-conversion))

### Automatiseringen
Valutawaarden kunnen automatiseringen activeren op basis van:
- Bedragdrempels
- Valutatypes
- Waarde wijzigingen

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Opmerking**: Hoewel elk projectlid aangepaste velden kan aanmaken, hangt de mogelijkheid om waarden in te stellen af van rolgebaseerde machtigingen die voor elk veld zijn geconfigureerd.

## Foutantwoorden

### Ongeldige Valutawaarde
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

Deze fout treedt op wanneer:
- De valutacode is niet een van de 72 ondersteunde codes
- Het getalformaat is ongeldig
- De waarde kan niet correct worden geparsed

### Aangepast Veld Niet Gevonden
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

## Beste Praktijken

### Valutakeuze
- Stel een standaardvaluta in die overeenkomt met je primaire markt
- Gebruik ISO 4217 valutacodes consistent
- Houd rekening met de locatie van de gebruiker bij het kiezen van standaardwaarden

### Waarde Beperkingen
- Stel redelijke min/max waarden in om invoerfouten te voorkomen
- Gebruik 0 als minimum voor velden die niet negatief mogen zijn
- Houd rekening met je gebruiksdoel bij het instellen van maximumwaarden

### Multi-Valuta Projecten
- Gebruik een consistente basisvaluta voor rapportage
- Implementeer VALUTA_CONVERSIE velden voor automatische conversie
- Documenteer welke valuta voor elk veld moet worden gebruikt

## Veelvoorkomende Gebruikscases

1. **Projectbegroting**
   - Projectbegroting bijhouden
   - Kostenschattingen
   - Uitgaven bijhouden

2. **Verkoop & Deals**
   - Dealwaarden
   - Contractbedragen
   - Omzet bijhouden

3. **Financiële Planning**
   - Investeringsbedragen
   - Financieringsrondes
   - Financiële doelstellingen

4. **Internationale Zaken**
   - Multi-valuta prijsstelling
   - Valutawisselkoersen bijhouden
   - Grensoverschrijdende transacties

## Beperkingen

- Maximaal 2 decimalen voor weergave (hoewel meer precisie wordt opgeslagen)
- Geen ingebouwde valutaconversie in standaard VALUTA velden
- Kan geen valuta's mengen in een enkele veldwaarde
- Geen automatische wisselkoersupdates (gebruik VALUTA_CONVERSIE hiervoor)
- Valutasymbolen zijn niet aanpasbaar

## Gerelateerde Bronnen

- [Valutaconversie Velden](/api/custom-fields/currency-conversion) - Automatische valutaconversie
- [Nummer Velden](/api/custom-fields/number) - Voor niet-monetaire numerieke waarden
- [Formule Velden](/api/custom-fields/formula) - Berekenen met valutawaarden
- [Lijst Aangepaste Velden](/api/custom-fields/list-custom-fields) - Vraag alle aangepaste velden in een project op