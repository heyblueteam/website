---
title: Multi-Select Aangepast Veld
description: Maak multi-selectvelden om gebruikers in staat te stellen meerdere opties te kiezen uit een vooraf gedefinieerde lijst
---

Multi-select aangepaste velden stellen gebruikers in staat om meerdere opties te kiezen uit een vooraf gedefinieerde lijst. Ze zijn ideaal voor categorieën, tags, vaardigheden, functies of elke situatie waarin meerdere selecties nodig zijn uit een gecontroleerde set opties.

## Basisvoorbeeld

Maak een eenvoudig multi-selectveld:

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een multi-selectveld en voeg vervolgens opties apart toe:

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het multi-selectveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `SELECT_MULTI` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |
| `projectId` | String! | ✅ Ja | ID van het project voor dit veld |

### CreateCustomFieldOptionInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ Ja | ID van het aangepaste veld |
| `title` | String! | ✅ Ja | Weergavetekst voor de optie |
| `color` | String | Nee | Kleur voor de optie (elke string) |
| `position` | Float | Nee | Sorteervolgorde voor de optie |

## Opties Toevoegen aan Bestaande Velden

Voeg nieuwe opties toe aan een bestaand multi-selectveld:

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## Multi-Select Waarden Instellen

Om meerdere geselecteerde opties op een record in te stellen:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het multi-select aangepaste veld |
| `customFieldOptionIds` | [String!] | ✅ Ja | Array van optie-ID's om te selecteren |

## Records Maken met Multi-Select Waarden

Bij het maken van een nieuw record met multi-select waarden:

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
    }
  }
}
```

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `selectedOptions` | [CustomFieldOption!] | Array van geselecteerde opties |
| `todo` | Todo! | Het record waar deze waarde bij hoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

### CustomFieldOption Antwoord

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de optie |
| `title` | String! | Weergavetekst voor de optie |
| `color` | String | Hex-kleurcode voor visuele weergave |
| `position` | Float | Sorteervolgorde voor de optie |
| `customField` | CustomField! | Het aangepaste veld waar deze optie bij hoort |

### CustomField Antwoord

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor het veld |
| `name` | String! | Weergavenaam van het multi-selectveld |
| `type` | CustomFieldType! | Altijd `SELECT_MULTI` |
| `description` | String | Hulptekst voor het veld |
| `customFieldOptions` | [CustomFieldOption!] | Alle beschikbare opties |

## Waardeformaat

### Invoerformaat
- **API-parameter**: Array van optie-ID's (`["option1", "option2", "option3"]`)
- **Stringformaat**: Komma-gescheiden optie-ID's (`"option1,option2,option3"`)

### Uitvoerformaat
- **GraphQL Antwoord**: Array van CustomFieldOption-objecten
- **Activiteitenlog**: Komma-gescheiden optietitels
- **Automatiseringsgegevens**: Array van optietitels

## Beheren van Opties

### Optie-eigenschappen Bijwerken
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
  }) {
    id
    title
    color
  }
}
```

### Optie Verwijderen
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### Opties Herordenen
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## Validatieregels

### Optievalidatie
- Alle opgegeven optie-ID's moeten bestaan
- Opties moeten behoren tot het opgegeven aangepaste veld
- Alleen SELECT_MULTI-velden kunnen meerdere opties geselecteerd hebben
- Lege array is geldig (geen selecties)

### Veldvalidatie
- Moet ten minste één optie gedefinieerd hebben om bruikbaar te zijn
- Optietitels moeten uniek zijn binnen het veld
- Kleurveld accepteert elke stringwaarde (geen hex-validatie)

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Foutantwoorden

### Ongeldige Optie-ID
```json
{
  "errors": [{
    "message": "Custom field option not found",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### Optie Behoort Niet tot Veld
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Veld Niet Gevonden
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Meerdere Opties op Niet-Multi Veld
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Best Practices

### Optieontwerp
- Gebruik beschrijvende, beknopte optietitels
- Pas consistente kleurcodering toe
- Houd optielijsten beheersbaar (typisch 3-20 opties)
- Orden opties logisch (alfabetisch, op frequentie, enz.)

### Gegevensbeheer
- Beoordeel en ruim ongebruikte opties periodiek op
- Gebruik consistente naamgevingsconventies in projecten
- Overweeg herbruikbaarheid van opties bij het maken van velden
- Plan voor optie-updates en migraties

### Gebruikerservaring
- Bied duidelijke veldbeschrijvingen
- Gebruik kleuren om visuele onderscheid te verbeteren
- Groepeer gerelateerde opties samen
- Overweeg standaardselecties voor veelvoorkomende gevallen

## Veelvoorkomende Gebruikscases

1. **Projectbeheer**
   - Taakcategorieën en tags
   - Prioriteitsniveaus en types
   - Toewijzingen van teamleden
   - Statusindicatoren

2. **Inhoudsbeheer**
   - Artikelcategorieën en onderwerpen
   - Inhoudstypen en -formaten
   - Publicatiekanalen
   - Goedkeuringsworkflows

3. **Klantenondersteuning**
   - Probleemcategorieën en -types
   - Aangetaste producten of diensten
   - Oplossingsmethoden
   - Klantsegmenten

4. **Productontwikkeling**
   - Functiecategorieën
   - Technische vereisten
   - Testomgevingen
   - Releasekanalen

## Integratiefuncties

### Met Automatiseringen
- Acties triggeren wanneer specifieke opties zijn geselecteerd
- Werk routeren op basis van geselecteerde categorieën
- Meldingen verzenden voor hoogprioritaire selecties
- Volgacties creëren op basis van optiecombinaties

### Met Lookup
- Records filteren op geselecteerde opties
- Gegevens aggregeren over optie-selecties
- Optiegegevens van andere records refereren
- Rapporten creëren op basis van optiecombinaties

### Met Formulieren
- Multi-select invoerbesturingselementen
- Optievalidatie en filtering
- Dynamisch laden van opties
- Voorwaardelijke veldweergave

## Activiteit Tracking

Wijzigingen in multi-selectvelden worden automatisch bijgehouden:
- Toont toegevoegde en verwijderde opties
- Toont optietitels in de activiteitenlog
- Tijdstempels voor alle selectie wijzigingen
- Gebruikersattributie voor aanpassingen

## Beperkingen

- Maximale praktische limiet van opties hangt af van UI-prestaties
- Geen hiërarchische of geneste optie-structuur
- Opties worden gedeeld over alle records die het veld gebruiken
- Geen ingebouwde optie-analyse of gebruiktracking
- Kleurveld accepteert elke string (geen hex-validatie)
- Geen verschillende machtigingen per optie instellen
- Opties moeten apart worden gemaakt, niet inline met veldcreatie
- Geen speciale herordening mutatie (gebruik editCustomFieldOption met positie)

## Gerelateerde Bronnen

- [Single-Select Velden](/api/custom-fields/select-single) - Voor enkelkeuze selecties
- [Checkbox Velden](/api/custom-fields/checkbox) - Voor eenvoudige boolean keuzes
- [Tekstvelden](/api/custom-fields/text-single) - Voor vrije tekstinvoer
- [Overzicht van Aangepaste Velden](/api/custom-fields/2.list-custom-fields) - Algemene concepten