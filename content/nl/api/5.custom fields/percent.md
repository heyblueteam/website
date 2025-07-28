---
title: Percent Aangepast Veld
description: Maak procentuele velden om numerieke waarden op te slaan met automatische % symboolverwerking en weergaveformattering
---

Percent aangepaste velden stellen je in staat om procentuele waarden voor records op te slaan. Ze verwerken automatisch het % symbool voor invoer en weergave, terwijl ze de ruwe numerieke waarde intern opslaan. Perfect voor voltooiingspercentages, succespercentages of andere op percentages gebaseerde statistieken.

## Basisvoorbeeld

Maak een eenvoudig percent veld:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een percent veld met beschrijving:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
  }) {
    id
    name
    type
    description
  }
}
```

## Invoervariabelen

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het percent veld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `PERCENT` |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |

**Opmerking**: De projectcontext wordt automatisch bepaald aan de hand van je authenticatieheaders. Geen `projectId` parameter is nodig.

**Opmerking**: PERCENT velden ondersteunen geen min/max beperkingen of prefixformattering zoals NUMBER velden.

## Percentwaarden Instellen

Percent velden slaan numerieke waarden op met automatische % symboolverwerking:

### Met Percent Symbool

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Directe Numerieke Waarde

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het percent aangepast veld |
| `number` | Float | Nee | Numerieke percentage waarde (bijv. 75.5 voor 75.5%) |

## Waarde Opslag en Weergave

### Opslagformaat
- **Interne opslag**: Ruwe numerieke waarde (bijv. 75.5)
- **Database**: Opgeslagen als `Decimal` in `number` kolom
- **GraphQL**: Teruggegeven als `Float` type

### Weergaveformaat
- **Gebruikersinterface**: Klantapplicaties moeten % symbool toevoegen (bijv. "75.5%")
- **Grafieken**: Toont met % symbool wanneer het uitvoertype PERCENTAGE is
- **API-antwoorden**: Ruwe numerieke waarde zonder % symbool (bijv. 75.5)

## Records Maken met Percentwaarden

Bij het maken van een nieuw record met percentwaarden:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Ondersteunde Invoervormen

| Formaat | Voorbeeld | Resultaat |
|---------|-----------|-----------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Opmerking**: Het % symbool wordt automatisch verwijderd van de invoer en opnieuw toegevoegd tijdens de weergave.

## Percentwaarden Opvragen

Bij het opvragen van records met percent aangepaste velden, krijg je de waarde via het `customField.value.number` pad:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

De reactie bevat het percentage als een ruwe waarde:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | ID! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld (bevat de percentwaarde) |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

**Belangrijk**: Percentwaarden worden benaderd via het `customField.value.number` veld. Het % symbool is niet inbegrepen in opgeslagen waarden en moet door klantapplicaties worden toegevoegd voor weergave.

## Filteren en Opvragen

Percent velden ondersteunen dezelfde filtering als NUMBER velden:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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
| `EQ` | Gelijk aan | `percentage = 75` |
| `NE` | Niet gelijk aan | `percentage ≠ 75` |
| `GT` | Groter dan | `percentage > 75` |
| `GTE` | Groter dan of gelijk aan | `percentage ≥ 75` |
| `LT` | Kleiner dan | `percentage < 75` |
| `LTE` | Kleiner dan of gelijk aan | `percentage ≤ 75` |
| `IN` | Waarde in lijst | `percentage in [50, 75, 100]` |
| `NIN` | Waarde niet in lijst | `percentage not in [0, 25]` |
| `IS` | Controleer op null met `values: null` | `percentage is null` |
| `NOT` | Controleer op niet null met `values: null` | `percentage is not null` |

### Bereikfiltering

Voor bereikfiltering, gebruik meerdere operators:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Percentage Waarde Bereiken

### Veelvoorkomende Bereiken

| Bereik | Beschrijving | Toepassing |
|--------|--------------|------------|
| `0-100` | Standaardpercentage | Completion rates, success rates |
| `0-∞` | Onbeperkt percentage | Growth rates, performance metrics |
| `-∞-∞` | Elke waarde | Change rates, variance |

### Voorbeeldwaarden

| Invoer | Opgeslagen | Weergave |
|--------|------------|----------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Grafiekaggregatie

Percent velden ondersteunen aggregatie in dashboardgrafieken en rapporten. Beschikbare functies zijn onder andere:

- `AVERAGE` - Gemiddelde procentuele waarde
- `COUNT` - Aantal records met waarden
- `MIN` - Laagste procentuele waarde
- `MAX` - Hoogste procentuele waarde 
- `SUM` - Totaal van alle procentuele waarden

Deze aggregaties zijn beschikbaar bij het maken van grafieken en dashboards, niet in directe GraphQL-queries.

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Foutantwoorden

### Ongeldig Percentageformaat
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

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

### Waarde-invoer
- Sta gebruikers toe om in te voeren met of zonder % symbool
- Valideer redelijke bereiken voor jouw gebruiksgeval
- Geef duidelijke context over wat 100% vertegenwoordigt

### Weergave
- Toon altijd % symbool in gebruikersinterfaces
- Gebruik geschikte decimale precisie
- Overweeg kleurcodering voor bereiken (rood/geel/groen)

### Gegevensinterpretatie
- Documenteer wat 100% betekent in jouw context
- Behandel waarden boven 100% op de juiste manier
- Overweeg of negatieve waarden geldig zijn

## Veelvoorkomende Toepassingen

1. **Projectmanagement**
   - Taakvoltooiingspercentages
   - Projectvoortgang
   - Hulpbronnenbenutting
   - Sprint snelheid

2. **Prestatie Tracking**
   - Succespercentages
   - Foutpercentages
   - Efficiëntiemetingen
   - Kwaliteitsscores

3. **Financiële Statistieken**
   - Groeipercentages
   - Winstmarges
   - Kortingbedragen
   - Veranderpercentages

4. **Analytics**
   - Conversiepercentages
   - Klikfrequenties
   - Betrokkenheidsstatistieken
   - Prestatie-indicatoren

## Integratiefuncties

### Met Formules
- Verwijs naar PERCENT velden in berekeningen
- Automatische % symboolformattering in formule-uitvoer
- Combineer met andere numerieke velden

### Met Automatiseringen
- Trigger acties op basis van percentage drempels
- Stuur meldingen voor mijlpaalpercentages
- Werk status bij op basis van voltooiingspercentages

### Met Zoekopdrachten
- Aggregatie van percentages uit gerelateerde records
- Bereken gemiddelde succespercentages
- Vind de hoogste/laagste presterende items

### Met Grafieken
- Maak procentuele visualisaties
- Volg voortgang in de tijd
- Vergelijk prestatiestatistieken

## Verschillen met NUMBER Velden

### Wat is Anders
- **Invoerhandling**: Verwijdert automatisch % symbool
- **Weergave**: Voegt automatisch % symbool toe
- **Beperkingen**: Geen min/max validatie
- **Formattering**: Geen prefixondersteuning

### Wat is Hetzelfde
- **Opslag**: Zelfde databasekolom en type
- **Filteren**: Zelfde query-operators
- **Aggregatie**: Zelfde aggregatiefuncties
- **Machtigingen**: Zelfde machtigingsmodel

## Beperkingen

- Geen min/max waarde beperkingen
- Geen prefix formatteringsopties
- Geen automatische validatie van 0-100% bereik
- Geen conversie tussen percentageformaten (bijv. 0.75 ↔ 75%)
- Waarden boven 100% zijn toegestaan

## Gerelateerde Bronnen

- [Overzicht van Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten van aangepaste velden
- [Nummer Aangepast Veld](/api/custom-fields/number) - Voor ruwe numerieke waarden
- [Automatiseringen API](/api/automations/index) - Maak percentage-gebaseerde automatiseringen