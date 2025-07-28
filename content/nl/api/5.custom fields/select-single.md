---
title: Enkele-selectie Aangepast Veld
description: Maak enkele-selectie velden zodat gebruikers één optie kunnen kiezen uit een vooraf gedefinieerde lijst
---

Enkele-selectie aangepaste velden stellen gebruikers in staat om precies één optie te kiezen uit een vooraf gedefinieerde lijst. Ze zijn ideaal voor statusvelden, categorieën, prioriteiten of elke situatie waarin slechts één keuze moet worden gemaakt uit een gecontroleerde set opties.

## Basisvoorbeeld

Maak een eenvoudig enkele-selectie veld:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een enkele-selectie veld met vooraf gedefinieerde opties:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het enkele-selectie veld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `SELECT_SINGLE` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | Nee | Initiële opties voor het veld |

### CreateCustomFieldOptionInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `title` | String! | ✅ Ja | Weergavetekst voor de optie |
| `color` | String | Nee | Hex-kleurcode voor de optie |

## Opties Toevoegen aan Bestaande Velden

Voeg nieuwe opties toe aan een bestaand enkele-selectie veld:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## Instellen van Enkele-Selectie Waarden

Om de geselecteerde optie op een record in te stellen:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van het bij te werken record |
| `customFieldId` | String! | ✅ Ja | ID van het enkele-selectie aangepaste veld |
| `customFieldOptionId` | String | Nee | ID van de optie om te selecteren (voorkeur voor enkele-selectie) |
| `customFieldOptionIds` | [String!] | Nee | Array van optie-ID's (gebruikt het eerste element voor enkele-selectie) |

## Opvragen van Enkele-Selectie Waarden

Vraag de enkele-selectiewaarde van een record op:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

Het `value` veld retourneert een JSON-object met de details van de geselecteerde optie.

## Records Maken met Enkele-Selectie Waarden

Bij het maken van een nieuw record met enkele-selectiewaarden:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
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
| `value` | JSON | Bevat het geselecteerde optieobject met id, titel, kleur, positie |
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
| `name` | String! | Weergavenaam van het enkele-selectie veld |
| `type` | CustomFieldType! | Altijd `SELECT_SINGLE` |
| `description` | String | Hulptekst voor het veld |
| `customFieldOptions` | [CustomFieldOption!] | Alle beschikbare opties |

## Waardeformaat

### Invoerformaat
- **API-parameter**: Gebruik `customFieldOptionId` voor enkele optie-ID
- **Alternatief**: Gebruik `customFieldOptionIds` array (neemt het eerste element)
- **Selectie wissen**: Laat beide velden weg of geef lege waarden door

### Uitvoerformaat
- **GraphQL Antwoord**: JSON-object in `value` veld met {id, titel, kleur, positie}
- **Activiteitenlog**: Optietitel als string
- **Automatiseringsgegevens**: Optietitel als string

## Selectiegedrag

### Exclusieve Selectie
- Het instellen van een nieuwe optie verwijdert automatisch de vorige selectie
- Slechts één optie kan tegelijk worden geselecteerd
- Het instellen van `null` of een lege waarde wist de selectie

### Terugvallogica
- Als `customFieldOptionIds` array wordt opgegeven, wordt alleen de eerste optie gebruikt
- Dit zorgt voor compatibiliteit met multi-select invoerformaten
- Lege arrays of null-waarden wissen de selectie

## Beheren van Opties

### Optie-eigenschappen Bijwerken
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
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

**Opmerking**: Het verwijderen van een optie wist deze uit alle records waar deze was geselecteerd.

### Opties Herordenen
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## Validatieregels

### Optievalidatie
- De opgegeven optie-ID moet bestaan
- Optie moet behoren tot het opgegeven aangepaste veld
- Slechts één optie kan worden geselecteerd (automatisch afgedwongen)
- Null/lege waarden zijn geldig (geen selectie)

### Veldvalidatie
- Moet ten minste één optie gedefinieerd hebben om bruikbaar te zijn
- Optietitels moeten uniek zijn binnen het veld
- Kleurcodes moeten geldig hex-formaat zijn (indien opgegeven)

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Foutantwoorden

### Ongeldige Optie-ID
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Waarde Kan Niet Geparsed Worden
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

## Beste Praktijken

### Optieontwerp
- Gebruik duidelijke, beschrijvende optietitels
- Pas betekenisvolle kleurcodering toe
- Houd optielijsten gefocust en relevant
- Orden opties logisch (op prioriteit, frequentie, enz.)

### Statusveldpatronen
- Gebruik consistente statusworkflows in projecten
- Overweeg de natuurlijke voortgang van opties
- Inclusief duidelijke eindtoestanden (Voltooid, Geannuleerd, enz.)
- Gebruik kleuren die de betekenis van de optie weerspiegelen

### Gegevensbeheer
- Beoordeel en ruim ongebruikte opties periodiek op
- Gebruik consistente naamgevingsconventies
- Overweeg de impact van optie-verwijdering op bestaande records
- Plan voor optie-updates en migraties

## Veelvoorkomende Gebruikscases

1. **Status en Workflow**
   - Taakstatus (Te Doen, In Behandeling, Voltooid)
   - Goedkeuringsstatus (In Afwachting, Goedgekeurd, Geweigerd)
   - Projectfase (Planning, Ontwikkeling, Testen, Vrijgegeven)
   - Probleemoplossingsstatus

2. **Classificatie en Categorisatie**
   - Prioriteitsniveaus (Laag, Gemiddeld, Hoog, Kritiek)
   - Taaktypen (Bug, Kenmerk, Verbetering, Documentatie)
   - Projectcategorieën (Intern, Klant, Onderzoek)
   - Afdelingstoewijzingen

3. **Kwaliteit en Beoordeling**
   - Beoordelingsstatus (Niet Begonnen, In Beoordeling, Goedgekeurd)
   - Kwaliteitsbeoordelingen (Slecht, Voldoende, Goed, Uitstekend)
   - Risiconiveaus (Laag, Gemiddeld, Hoog)
   - Vertrouwensniveaus

4. **Toewijzing en Eigenaarschap**
   - Teamtoewijzingen
   - Afdelingseigenaarschap
   - Rolgebaseerde toewijzingen
   - Regionale toewijzingen

## Integratiefuncties

### Met Automatiseringen
- Acties triggeren wanneer specifieke opties worden geselecteerd
- Werk routeren op basis van geselecteerde categorieën
- Meldingen verzenden voor statuswijzigingen
- Voorwaardelijke workflows creëren op basis van selecties

### Met Opzoekingen
- Records filteren op geselecteerde opties
- Optiegegevens uit andere records refereren
- Rapporten opstellen op basis van optie-selecties
- Records groeperen op geselecteerde waarden

### Met Formulieren
- Dropdown-invoervelden
- Radioknopinterfaces
- Optievalidatie en filtering
- Voorwaardelijke veldweergave op basis van selecties

## Activiteit Tracking

Veranderingen in enkele-selectie velden worden automatisch gevolgd:
- Toont oude en nieuwe optie-selecties
- Toont optietitels in de activiteitenlog
- Tijdstempels voor alle selectie wijzigingen
- Gebruikersattribuut voor aanpassingen

## Verschillen met Multi-Select

| Kenmerk | Enkele-Selectie | Multi-Select |
|---------|---------------|--------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Beperkingen

- Slechts één optie kan tegelijk worden geselecteerd
- Geen hiërarchische of geneste optie-structuur
- Opties worden gedeeld over alle records die het veld gebruiken
- Geen ingebouwde optie-analyse of gebruiksregistratie
- Kleurcodes zijn alleen voor weergave, geen functionele impact
- Geen verschillende machtigingen per optie instelbaar

## Gerelateerde Bronnen

- [Multi-Select Velden](/api/custom-fields/select-multi) - Voor meerdere-keuze selecties
- [Checkbox Velden](/api/custom-fields/checkbox) - Voor eenvoudige boolean keuzes
- [Tekstvelden](/api/custom-fields/text-single) - Voor vrije tekstinvoer
- [Overzicht van Aangepaste Velden](/api/custom-fields/1.index) - Algemene concepten