---
title: Beoordelings Aangepast Veld
description: Maak beoordelingsvelden om numerieke beoordelingen op te slaan met configureerbare schalen en validatie
---

Beoordelings aangepaste velden stellen je in staat om numerieke beoordelingen in records op te slaan met configureerbare minimum- en maximumwaarden. Ze zijn ideaal voor prestatiebeoordelingen, tevredenheidsscores, prioriteitsniveaus of gegevens op basis van een numerieke schaal in je projecten.

## Basisvoorbeeld

Maak een eenvoudig beoordelingsveld met de standaard 0-5 schaal:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Geavanceerd Voorbeeld

Maak een beoordelingsveld met een aangepaste schaal en beschrijving:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Invoervelden

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het beoordelingsveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `RATING` |
| `projectId` | String! | ✅ Ja | Het project-ID waar dit veld zal worden aangemaakt |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |
| `min` | Float | Nee | Minimale beoordelingswaarde (geen standaard) |
| `max` | Float | Nee | Maximale beoordelingswaarde |

## Beoordelingswaarden Instellen

Om een beoordelingswaarde op een record in te stellen of bij te werken:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het beoordelings aangepaste veld |
| `value` | String! | ✅ Ja | Beoordelingswaarde als string (binnen het geconfigureerde bereik) |

## Records Maken met Beoordelingswaarden

Bij het maken van een nieuw record met beoordelingswaarden:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
    }
  }
}
```

## Responsvelden

### TodoCustomField Respons

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `value` | Float | De opgeslagen beoordelingswaarde (toegankelijk via customField.value) |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

**Opmerking**: De beoordelingswaarde wordt daadwerkelijk toegankelijk via `customField.value.number` in queries.

### CustomField Respons

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificatie voor het veld |
| `name` | String! | Weergavenaam van het beoordelingsveld |
| `type` | CustomFieldType! | Altijd `RATING` |
| `min` | Float | Minimale toegestane beoordelingswaarde |
| `max` | Float | Maximale toegestane beoordelingswaarde |
| `description` | String | Helptekst voor het veld |

## Beoordelingsvalidatie

### Waarde Beperkingen
- Beoordelingswaarden moeten numeriek zijn (Float type)
- Waarden moeten binnen het geconfigureerde min/max bereik liggen
- Als er geen minimum is opgegeven, is er geen standaardwaarde
- Maximale waarde is optioneel maar aanbevolen

### Validatieregels
**Belangrijk**: Validatie vindt alleen plaats bij het indienen van formulieren, niet bij het direct gebruiken van `setTodoCustomField`.

- Invoer wordt geparsed als een float-nummer (bij gebruik van formulieren)
- Moet groter zijn dan of gelijk aan de minimumwaarde (bij gebruik van formulieren)
- Moet kleiner zijn dan of gelijk aan de maximumwaarde (bij gebruik van formulieren)
- `setTodoCustomField` accepteert elke stringwaarde zonder validatie

### Geldige Beoordelingsvoorbeelden
Voor een veld met min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Ongeldige Beoordelingsvoorbeelden
Voor een veld met min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Configuratieopties

### Beoordelingsschaal Instellen
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Veelvoorkomende Beoordelingsschalen
- **1-5 Sterren**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 Prestatie**: `min: 1, max: 10`
- **0-100 Percentage**: `min: 0, max: 100`
- **Aangepaste Schaal**: Elke numerieke reeks

## Vereiste Machtigingen

Aangepaste veldbewerkingen volgen de standaard rolgebaseerde machtigingen:

| Actie | Vereiste Rol |
|-------|--------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Opmerking**: De specifieke rollen die vereist zijn, zijn afhankelijk van de aangepaste rolconfiguratie van je project en de machtigingen op veldniveau.

## Foutreacties

### Validatiefout (Alleen Formulieren)
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

**Belangrijk**: Validatie van de beoordelingswaarde (min/max beperkingen) vindt alleen plaats bij het indienen van formulieren, niet bij het direct gebruiken van `setTodoCustomField`.

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

### Schaalontwerp
- Gebruik consistente beoordelingsschalen voor vergelijkbare velden
- Houd rekening met de bekendheid van de gebruiker (1-5 sterren, 0-10 NPS)
- Stel geschikte minimumwaarden in (0 vs 1)
- Definieer een duidelijke betekenis voor elk beoordelingsniveau

### Gegevenskwaliteit
- Valideer beoordelingswaarden voordat je ze opslaat
- Gebruik decimale precisie op de juiste manier
- Overweeg afronding voor weergavedoeleinden
- Geef duidelijke richtlijnen over de betekenissen van beoordelingen

### Gebruikerservaring
- Toon beoordelingsschalen visueel (sterren, voortgangsbalken)
- Toon de huidige waarde en schaalgrenzen
- Geef context voor de betekenissen van beoordelingen
- Overweeg standaardwaarden voor nieuwe records

## Veelvoorkomende Gebruikscases

1. **Prestatiemanagement**
   - Beoordelingen van werknemersprestaties
   - Kwaliteitsscores van projecten
   - Beoordelingen van taakvoltooiing
   - Beoordelingen van vaardigheidsniveaus

2. **Klantfeedback**
   - Tevredenheidsbeoordelingen
   - Kwaliteitsscores van producten
   - Beoordelingen van service-ervaring
   - Net Promoter Score (NPS)

3. **Prioriteit en Belang**
   - Prioriteitsniveaus van taken
   - Urgentiebeoordelingen
   - Risicobeoordelingsscores
   - Impactbeoordelingen

4. **Kwaliteitsborging**
   - Beoordelingen van codebeoordelingen
   - Kwaliteitsscores van testen
   - Kwaliteit van documentatie
   - Beoordelingen van procesnaleving

## Integratiefuncties

### Met Automatiseringen
- Acties triggeren op basis van beoordelingsdrempels
- Meldingen verzenden voor lage beoordelingen
- Follow-up taken aanmaken voor hoge beoordelingen
- Werk routeren op basis van beoordelingswaarden

### Met Lookup
- Gemiddelde beoordelingen over records berekenen
- Records vinden op basis van beoordelingsbereiken
- Beoordelingsgegevens van andere records verwijzen
- Beoordelingsstatistieken aggregeren

### Met Blue Frontend
- Automatische bereikvalidatie in formuliercontexten
- Visuele beoordelingsinvoervelden
- Real-time validatiefeedback
- Ster- of schuifinvoermogelijkheden

## Activiteit Tracking

Wijzigingen in beoordelingsvelden worden automatisch gevolgd:
- Oude en nieuwe beoordelingswaarden worden gelogd
- Activiteit toont numerieke wijzigingen
- Timestamps voor alle beoordelingsupdates
- Gebruikersattributie voor wijzigingen

## Beperkingen

- Alleen numerieke waarden worden ondersteund
- Geen ingebouwde visuele beoordelingsweergave (sterren, enz.)
- Decimale precisie hangt af van databaseconfiguratie
- Geen opslag van beoordelingsmetadata (opmerkingen, context)
- Geen automatische aggregatie of statistieken voor beoordelingen
- Geen ingebouwde conversie van beoordelingen tussen schalen
- **Kritisch**: Min/max validatie werkt alleen in formulieren, niet via `setTodoCustomField`

## Gerelateerde Bronnen

- [Nummer Velden](/api/5.custom%20fields/number) - Voor algemene numerieke gegevens
- [Percentage Velden](/api/5.custom%20fields/percent) - Voor percentagewaarden
- [Selecteer Velden](/api/5.custom%20fields/select-single) - Voor discrete keuze beoordelingen
- [Overzicht van Aangepaste Velden](/api/5.custom%20fields/2.list-custom-fields) - Algemene concepten