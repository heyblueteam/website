---
title: Formule Aangepast Veld
description: Maak berekende velden die automatisch waarden berekenen op basis van andere gegevens
---

Formule aangepaste velden worden gebruikt voor grafiek- en dashboardberekeningen binnen Blue. Ze definiëren aggregatiefuncties (SOM, GEMIDDELDE, TELLEN, enz.) die werken op aangepaste veldgegevens om berekende statistieken in grafieken weer te geven. Formules worden niet berekend op het niveau van individuele taken, maar aggregeren gegevens over meerdere records voor visualisatiedoeleinden.

## Basisvoorbeeld

Maak een formuleveld voor grafiekberekeningen:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Geavanceerd Voorbeeld

Maak een valutafomule met complexe berekeningen:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het formuleveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `FORMULA` |
| `projectId` | String! | ✅ Ja | Het project-ID waar dit veld zal worden aangemaakt |
| `formula` | JSON | Nee | Formule-definitie voor grafiekberekeningen |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

### Formule Structuur

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Ondersteunde Functies

### Grafiek Aggregatiefuncties

Formulevelden ondersteunen de volgende aggregatiefuncties voor grafiekberekeningen:

| Functie | Beschrijving | ChartFunction Enum |
|----------|-------------|-------------------|
| `SUM` | Som van alle waarden | `SUM` |
| `AVERAGE` | Gemiddelde van numerieke waarden | `AVERAGE` |
| `AVERAGEA` | Gemiddelde zonder nullen en null-waarden | `AVERAGEA` |
| `COUNT` | Aantal waarden | `COUNT` |
| `COUNTA` | Aantal zonder nullen en null-waarden | `COUNTA` |
| `MAX` | Maximale waarde | `MAX` |
| `MIN` | Minimale waarde | `MIN` |

**Opmerking**: Deze functies worden gebruikt in het `display.function` veld en werken op geaggregeerde gegevens voor grafiekvisualisaties. Complexe wiskundige uitdrukkingen of berekeningen op veldniveau worden niet ondersteund.

## Weergavetypes

### Nummerweergave

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Resultaat: `1250.75`

### Valutaweergave

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Resultaat: `$1,250.75`

### Percentageweergave

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Resultaat: `87.5%`

## Bewerken van Formulevelden

Werk bestaande formulevelden bij:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Formuleverwerking

### Grafiekberekeningscontext

Formulevelden worden verwerkt in de context van grafieksegmenten en dashboards:
- Berekeningen vinden plaats wanneer grafieken worden weergegeven of bijgewerkt
- Resultaten worden opgeslagen in `ChartSegment.formulaResult` als decimale waarden
- Verwerking wordt afgehandeld via een speciale BullMQ-queue genaamd 'formula'
- Updates worden gepubliceerd naar dashboard-abonnees voor realtime-updates

### Weergaveformattering

De `getFormulaDisplayValue` functie formatteert de berekende resultaten op basis van het weergavetyp:
- **NUMMER**: Wordt weergegeven als een gewoon nummer met optionele precisie
- **PERCENTAGE**: Voegt % achtervoegsel toe met optionele precisie  
- **VALUTA**: Formatteert met de opgegeven valutacode

## Opslag van Formule Resultaten

Resultaten worden opgeslagen in het `formulaResult` veld:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het formuleveld |
| `number` | Float | Berekende numerieke resultaat |
| `formulaResult` | JSON | Volledig resultaat met weergaveformattering |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is berekend |

## Gegevenscontext

### Grafiekgegevensbron

Formulevelden werken binnen de context van de grafiekgegevensbron:
- Formules aggregeren aangepaste veldwaarden over taken in een project
- De aggregatiefunctie die is opgegeven in `display.function` bepaalt de berekening
- Resultaten worden berekend met behulp van SQL-aggregatiefuncties (gemiddelde, som, telling, enz.)
- Berekeningen worden op database-niveau uitgevoerd voor efficiëntie

## Veelvoorkomende Formule Voorbeelden

### Totaal Budget (Grafiekweergave)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Gemiddelde Score (Grafiekweergave)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Taak Telling (Grafiekweergave)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Vereiste Machtigingen

Bewerkingen van aangepaste velden volgen de standaard rolgebaseerde machtigingen:

| Actie | Vereiste Rol |
|--------|---------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Opmerking**: De specifieke vereiste rollen zijn afhankelijk van de aangepaste rolconfiguratie van uw project. Er zijn geen speciale machtigingsconstanten zoals CUSTOM_FIELDS_CREATE.

## Foutafhandeling

### Validatiefout
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

### Formule Ontwerp
- Gebruik duidelijke, beschrijvende namen voor formulevelden
- Voeg beschrijvingen toe die de berekeningslogica uitleggen
- Test formules met voorbeeldgegevens voordat u ze implementeert
- Houd formules eenvoudig en leesbaar

### Prestatieoptimalisatie
- Vermijd diep geneste formule-afhankelijkheden
- Gebruik specifieke veldverwijzingen in plaats van jokertekens
- Overweeg cachingstrategieën voor complexe berekeningen
- Houd de prestaties van formules in grote projecten in de gaten

### Gegevenskwaliteit
- Valideer brondgegevens voordat u deze in formules gebruikt
- Behandel lege of null-waarden op de juiste manier
- Gebruik de juiste precisie voor weergavetypes
- Overweeg randgevallen in berekeningen

## Veelvoorkomende Toepassingen

1. **Financiële Tracking**
   - Budgetberekeningen
   - Winst/verliesoverzichten
   - Kostenanalyse
   - Omzetprognoses

2. **Projectmanagement**
   - Voltooiingspercentages
   - Hulpbronnenbenutting
   - Tijdlijnberekeningen
   - Prestatiestatistieken

3. **Kwaliteitscontrole**
   - Gemiddelde scores
   - Slagings-/faalpercentages
   - Kwaliteitsstatistieken
   - Nalevingsbewaking

4. **Business Intelligence**
   - KPI-berekeningen
   - Trendanalyse
   - Vergelijkende statistieken
   - Dashboardwaarden

## Beperkingen

- Formules zijn alleen voor aggregaties van grafieken/dashboards, niet voor berekeningen op taakniveau
- Beperkt tot de zeven ondersteunde aggregatiefuncties (SOM, GEMIDDELDE, enz.)
- Geen complexe wiskundige uitdrukkingen of veld-tot-veld berekeningen
- Kan geen meerdere velden in een enkele formule verwijzen
- Resultaten zijn alleen zichtbaar in grafieken en dashboards
- Het `logic` veld is alleen voor weergavetekst, niet voor daadwerkelijke berekeningslogica

## Gerelateerde Bronnen

- [Nummer Velden](/api/5.custom%20fields/number) - Voor statische numerieke waarden
- [Valuta Velden](/api/5.custom%20fields/currency) - Voor monetaire waarden
- [Referentie Velden](/api/5.custom%20fields/reference) - Voor cross-projectgegevens
- [Lookup Velden](/api/5.custom%20fields/lookup) - Voor geaggregeerde gegevens
- [Overzicht Aangepaste Velden](/api/5.custom%20fields/2.list-custom-fields) - Algemene concepten