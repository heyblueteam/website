---
title: Aangepast Nummer Veld
description: Maak nummervelden aan om numerieke waarden op te slaan met optionele min/max beperkingen en prefix-opmaak
---

Aangepaste nummervelden stellen je in staat om numerieke waarden voor records op te slaan. Ze ondersteunen validatiebeperkingen, decimale precisie en kunnen worden gebruikt voor hoeveelheden, scores, metingen of andere numerieke gegevens die geen speciale opmaak vereisen.

## Basis Voorbeeld

Maak een eenvoudig nummer veld aan:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een nummer veld aan met beperkingen en prefix:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het nummer veld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `NUMBER` |
| `projectId` | String! | ✅ Ja | ID van het project waarin het veld moet worden aangemaakt |
| `min` | Float | Nee | Minimum waarde beperking (alleen UI begeleiding) |
| `max` | Float | Nee | Maximum waarde beperking (alleen UI begeleiding) |
| `prefix` | String | Nee | Weergave prefix (bijv., "#", "~", "$") |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

## Instellen van Nummerwaarden

Nummervelden slaan decimale waarden op met optionele validatie:

### Eenvoudige Nummerwaarde

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Gehele Waarde

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het aangepaste nummer veld |
| `number` | Float | Nee | Numerieke waarde om op te slaan |

## Waarde Beperkingen

### Min/Max Beperkingen (UI Begeleiding)

**Belangrijk**: Min/max beperkingen worden opgeslagen maar NIET afgedwongen aan de serverzijde. Ze dienen als UI begeleiding voor frontend toepassingen.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Client-Side Validatie Vereist**: Frontend toepassingen moeten validatielogica implementeren om min/max beperkingen af te dwingen.

### Ondersteunde Waardetypen

| Type | Voorbeeld | Beschrijving |
|------|-----------|--------------|
| Integer | `42` | Gehele getallen |
| Decimal | `42.5` | Getallen met decimalen |
| Negative | `-10` | Negatieve waarden (indien geen min beperking) |
| Zero | `0` | Nulwaarde |

**Opmerking**: Min/max beperkingen worden NIET gevalideerd aan de serverzijde. Waarden buiten het opgegeven bereik worden geaccepteerd en opgeslagen.

## Records Aanmaken met Nummerwaarden

Bij het aanmaken van een nieuw record met nummerwaarden:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### Ondersteunde Invoerformaten

Bij het aanmaken van records, gebruik de `number` parameter (niet `value`) in de array van aangepaste velden:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Antwoord Velden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `number` | Float | De numerieke waarde |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

### CustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificatie voor de velddefinitie |
| `name` | String! | Weergavenaam van het veld |
| `type` | CustomFieldType! | Altijd `NUMBER` |
| `min` | Float | Minimum toegestane waarde |
| `max` | Float | Maximum toegestane waarde |
| `prefix` | String | Weergave prefix |
| `description` | String | Hulptekst |

**Opmerking**: Als de nummerwaarde niet is ingesteld, zal het `number` veld zijn `null`.

## Filteren en Queryen

Nummervelden ondersteunen uitgebreide numerieke filtering:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Ondersteunde Operators

| Operator | Beschrijving | Voorbeeld |
|----------|--------------|-----------|
| `EQ` | Gelijk aan | `number = 42` |
| `NE` | Niet gelijk aan | `number ≠ 42` |
| `GT` | Groter dan | `number > 42` |
| `GTE` | Groter dan of gelijk aan | `number ≥ 42` |
| `LT` | Kleiner dan | `number < 42` |
| `LTE` | Kleiner dan of gelijk aan | `number ≤ 42` |
| `IN` | In array | `number in [1, 2, 3]` |
| `NIN` | Niet in array | `number not in [1, 2, 3]` |
| `IS` | Is null/niet null | `number is null` |

### Bereik Filtering

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Weergave Opmaak

### Met Prefix

Als er een prefix is ingesteld, zal deze worden weergegeven:

| Waarde | Prefix | Weergave |
|--------|--------|----------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Decimale Precisie

Getallen behouden hun decimale precisie:

| Invoer | Opgeslagen | Weergegeven |
|--------|------------|-------------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Foutreacties

### Ongeldig Nummerformaat
```json
{
  "errors": [{
    "message": "Invalid number format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Veld Niet Gevonden
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

**Opmerking**: Min/max validatiefouten komen NIET voor aan de serverzijde. Beperkingsvalidatie moet worden geïmplementeerd in je frontend toepassing.

### Geen Nummer
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Beste Praktijken

### Beperkingsontwerp
- Stel realistische min/max waarden in voor UI begeleiding
- Implementeer client-side validatie om beperkingen af te dwingen
- Gebruik beperkingen om gebruikersfeedback in formulieren te geven
- Overweeg of negatieve waarden geldig zijn voor jouw gebruiksgeval

### Waarde Precisie
- Gebruik geschikte decimale precisie voor jouw behoeften
- Overweeg afronding voor weergave doeleinden
- Wees consistent met precisie over verwante velden

### Weergave Verbetering
- Gebruik betekenisvolle prefixes voor context
- Overweeg eenheden in veldnamen (bijv., "Gewicht (kg)")
- Geef duidelijke beschrijvingen voor validatieregels

## Veelvoorkomende Gebruikscases

1. **Scoringssystemen**
   - Prestatiebeoordelingen
   - Kwaliteitsscores
   - Prioriteitsniveaus
   - Klanttevredenheidsbeoordelingen

2. **Metingen**
   - Hoeveelheden en bedragen
   - Afmetingen en groottes
   - Duur (in numeriek formaat)
   - Capaciteiten en limieten

3. **Zakelijke Statistieken**
   - Omzetcijfers
   - Conversiepercentages
   - Budgetallocaties
   - Doelgetallen

4. **Technische Gegevens**
   - Versienummers
   - Configuratie waarden
   - Prestatiestatistieken
   - Drempelinstellingen

## Integratiefuncties

### Met Grafieken en Dashboards
- Gebruik NUMMER velden in grafiekberekeningen
- Maak numerieke visualisaties
- Volg trends in de tijd

### Met Automatiseringen
- Trigger acties op basis van nummerdrempels
- Werk verwante velden bij op basis van nummerwijzigingen
- Stuur meldingen voor specifieke waarden

### Met Zoekopdrachten
- Agregeer getallen van verwante records
- Bereken totalen en gemiddelden
- Vind min/max waarden over relaties

### Met Grafieken
- Maak numerieke visualisaties
- Volg trends in de tijd
- Vergelijk waarden tussen records

## Beperkingen

- **Geen server-side validatie** van min/max beperkingen
- **Client-side validatie vereist** voor handhaving van beperkingen
- Geen ingebouwde valuta-opmaak (gebruik in plaats daarvan TYPE VALUTA)
- Geen automatische procentteken (gebruik in plaats daarvan TYPE PERCENT)
- Geen eenheidsconversiemogelijkheden
- Decimale precisie beperkt door database Decimal type
- Geen wiskundige formule-evaluatie in het veld zelf

## Gerelateerde Bronnen

- [Overzicht Aangepaste Velden](/api/custom-fields/1.index) - Algemene concepten van aangepaste velden
- [Valuta Aangepast Veld](/api/custom-fields/currency) - Voor monetaire waarden
- [Percentage Aangepast Veld](/api/custom-fields/percent) - Voor percentage waarden
- [Automatiseringen API](/api/automations/1.index) - Maak nummer-gebaseerde automatiseringen