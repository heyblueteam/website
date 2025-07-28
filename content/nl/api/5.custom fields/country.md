---
title: Land Aangepast Veld
description: Maak landselectievelden met ISO-landcodevalidatie
---

Landaangepaste velden stellen je in staat om landinformatie voor records op te slaan en te beheren. Het veld ondersteunt zowel landnamen als ISO Alpha-2 landcodes.

**Belangrijk**: De validatie en conversie van landen verschilt aanzienlijk tussen mutaties:
- **createTodo**: Valideert en converteert automatisch landnamen naar ISO-codes
- **setTodoCustomField**: Accepteert elke waarde zonder validatie

## Basisvoorbeeld

Maak een eenvoudig landveld:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een landveld met beschrijving:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het landveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `COUNTRY` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

**Opmerking**: De `projectId` wordt niet doorgegeven in de invoer, maar wordt bepaald door de GraphQL-context (typisch uit aanvraagheaders of authenticatie).

## Instellen van Landwaarden

Landvelden slaan gegevens op in twee databasevelden:
- **`countryCodes`**: Slaat ISO Alpha-2 landcodes op als een door komma's gescheiden string in de database (teruggegeven als array via API)
- **`text`**: Slaat weergavetekst of landnamen op als een string

### Begrijpen van de Parameters

De `setTodoCustomField` mutatie accepteert twee optionele parameters voor landvelden:

| Parameter | Type | Vereist | Beschrijving | Wat het doet |
|-----------|------|----------|-------------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt | - |
| `customFieldId` | String! | ✅ Ja | ID van het land aangepaste veld | - |
| `countryCodes` | [String!] | Nee | Array van ISO Alpha-2 landcodes | Stored in the `countryCodes` field |
| `text` | String | Nee | Weergavetekst of landnamen | Stored in the `text` field |

**Belangrijk**: 
- In `setTodoCustomField`: Beide parameters zijn optioneel en worden onafhankelijk opgeslagen
- In `createTodo`: Het systeem stelt automatisch beide velden in op basis van jouw invoer (je kunt ze niet onafhankelijk beheren)

### Optie 1: Alleen Landcodes Gebruiken

Sla gevalideerde ISO-codes op zonder weergavetekst:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Resultaat: `countryCodes` = `["US"]`, `text` = `null`

### Optie 2: Alleen Tekst Gebruiken

Sla weergavetekst op zonder gevalideerde codes:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Resultaat: `countryCodes` = `null`, `text` = `"United States"`

**Opmerking**: Bij gebruik van `setTodoCustomField` vindt er geen validatie plaats, ongeacht welke parameter je gebruikt. De waarden worden exact opgeslagen zoals opgegeven.

### Optie 3: Beide Gebruiken (Aanbevolen)

Sla zowel gevalideerde codes als weergavetekst op:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Resultaat: `countryCodes` = `["US"]`, `text` = `"United States"`

### Meerdere Landen

Sla meerdere landen op met behulp van arrays:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Records Maken met Landwaarden

Bij het maken van records valideert en converteert de `createTodo` mutatie **automatisch** landwaarden. Dit is de enige mutatie die landvalidatie uitvoert:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      text
      countryCodes
    }
  }
}
```

### Geaccepteerde Invoerformaten

| Invoertype | Voorbeeld | Resultaat |
|------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Niet ondersteund** - behandeld als enkele ongeldige waarde |
| Mixed format | `"United States, CA"` | **Niet ondersteund** - behandeld als enkele ongeldige waarde |

## Responsvelden

### TodoCustomField Respons

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `text` | String | Weergavetekst (landnamen) |
| `countryCodes` | [String!] | Array van ISO Alpha-2 landcodes |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

## Landnormen

Blue gebruikt de **ISO 3166-1 Alpha-2** standaard voor landcodes:

- Tweeletterige landcodes (bijv. US, GB, FR, DE)
- Validatie met behulp van de `i18n-iso-countries` bibliotheek **vindt alleen plaats in createTodo**
- Ondersteunt alle officieel erkende landen

### Voorbeeld Landcodes

| Land | ISO Code |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Voor de complete officiële lijst van ISO 3166-1 alpha-2 landcodes, bezoek de [ISO Online Browsing Platform](https://www.iso.org/obp/ui/#search/code/).

## Validatie

**Validatie vindt alleen plaats in de `createTodo` mutatie**:

1. **Geldige ISO-code**: Accepteert elke geldige ISO Alpha-2 code
2. **Landnaam**: Converteert automatisch erkende landnamen naar codes
3. **Ongeldige invoer**: Gooit `CustomFieldValueParseError` voor niet-herkende waarden

**Opmerking**: De `setTodoCustomField` mutatie voert GEEN validatie uit en accepteert elke stringwaarde.

### Foutvoorbeeld

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Integratiefuncties

### Lookup Velden
Landvelden kunnen worden verwezen door LOOKUP aangepaste velden, waardoor je landgegevens uit gerelateerde records kunt ophalen.

### Automatiseringen
Gebruik landwaarden in automatiseringsvoorwaarden:
- Filter acties op specifieke landen
- Stuur meldingen op basis van land
- Routeer taken op basis van geografische regio's

### Formulieren
Landvelden in formulieren valideren automatisch gebruikersinvoer en converteren landnamen naar codes.

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Foutreacties

### Ongeldige Landwaarde
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Veldtype Mismatch
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Beste Praktijken

### Invoerafhandeling
- Gebruik `createTodo` voor automatische validatie en conversie
- Gebruik `setTodoCustomField` voorzichtig, aangezien het validatie omzeilt
- Overweeg om invoer in je applicatie te valideren voordat je `setTodoCustomField` gebruikt
- Toon volledige landnamen in de UI voor duidelijkheid

### Gegevenskwaliteit
- Valideer landinvoer op het invoerpunt
- Gebruik consistente formaten in je systeem
- Overweeg regionale groeperingen voor rapportage

### Meerdere Landen
- Gebruik array-ondersteuning in `setTodoCustomField` voor meerdere landen
- Meerdere landen in `createTodo` worden **niet ondersteund** via het waardeveld
- Sla landcodes op als array in `setTodoCustomField` voor juiste verwerking

## Veelvoorkomende Gebruikscases

1. **Klantbeheer**
   - Locatie van het hoofdkantoor van de klant
   - Verzendbestemmingen
   - Belastingjurisdicties

2. **Projecttracking**
   - Projectlocatie
   - Locaties van teamleden
   - Marktdoelen

3. **Naleving & Juridisch**
   - Regelgevende jurisdicties
   - Gegevensresidentievereisten
   - Exportcontroles

4. **Verkoop & Marketing**
   - Territoriale toewijzingen
   - Marktsegmentatie
   - Campagne-targeting

## Beperkingen

- Ondersteunt alleen ISO 3166-1 Alpha-2 codes (2-letterige codes)
- Geen ingebouwde ondersteuning voor landonderverdelingen (staten/provincies)
- Geen automatische landvlag-iconen (alleen tekstgebaseerd)
- Kan historische landcodes niet valideren
- Geen ingebouwde regio- of continentgroepering
- **Validatie werkt alleen in `createTodo`, niet in `setTodoCustomField`**
- **Meerdere landen worden niet ondersteund in `createTodo` waardeveld**
- **Landcodes worden opgeslagen als een door komma's gescheiden string, niet als echte array**

## Gerelateerde Bronnen

- [Overzicht van Aangepaste Velden](/custom-fields/list-custom-fields) - Algemene concepten van aangepaste velden
- [Lookup Velden](/api/custom-fields/lookup) - Verwijs landgegevens uit andere records
- [Formulieren API](/api/forms) - Inclusief landvelden in aangepaste formulieren