---
title: Formel Anpassat Fält
description: Skapa beräknade fält som automatiskt beräknar värden baserat på annan data
---

Formel anpassade fält används för diagram- och instrumentpanelberäkningar inom Blue. De definierar aggregationsfunktioner (SUM, MEDEL, ANTAL, etc.) som arbetar på data från anpassade fält för att visa beräknade mätvärden i diagram. Formler beräknas inte på den individuella todo-nivån utan aggregerar data över flera poster för visualiseringsändamål.

## Grundläggande Exempel

Skapa ett formelfält för diagramberäkningar:

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

## Avancerat Exempel

Skapa en valutaformel med komplexa beräkningar:

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

## Inmatningsparametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för formelfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `FORMULA` |
| `projectId` | String! | ✅ Ja | Projekt-ID där detta fält kommer att skapas |
| `formula` | JSON | Nej | Formeldefinition för diagramberäkningar |
| `description` | String | Nej | Hjälptext som visas för användare |

### Formelstruktur

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

## Stödda Funktioner

### Diagramaggregationsfunktioner

Formelfält stöder följande aggregationsfunktioner för diagramberäkningar:

| Funktion | Beskrivning | ChartFunction Enum |
|----------|-------------|-------------------|
| `SUM` | Summa av alla värden | `SUM` |
| `AVERAGE` | Medelvärde av numeriska värden | `AVERAGE` |
| `AVERAGEA` | Medelvärde utan nollor och null-värden | `AVERAGEA` |
| `COUNT` | Antal värden | `COUNT` |
| `COUNTA` | Antal utan nollor och null-värden | `COUNTA` |
| `MAX` | Maximalt värde | `MAX` |
| `MIN` | Minimi värde | `MIN` |

**Notera**: Dessa funktioner används i `display.function` fältet och arbetar på aggregerad data för diagramvisualiseringar. Komplexa matematiska uttryck eller fält-till-fält beräkningar stöds inte.

## Visningstyper

### Nummervisning

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Resultat: `1250.75`

### Valuta Visning

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

Resultat: `$1,250.75`

### Procentvisning

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Resultat: `87.5%`

## Redigera Formel Fält

Uppdatera befintliga formelfält:

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

## Formelbehandling

### Diagramberäkningskontext

Formelfält bearbetas i kontexten av diagramsegment och instrumentpaneler:
- Beräkningar sker när diagram renderas eller uppdateras
- Resultat lagras i `ChartSegment.formulaResult` som decimalvärden
- Bearbetning hanteras genom en dedikerad BullMQ-kö som heter 'formel'
- Uppdateringar publiceras till instrumentpanelens prenumeranter för realtidsuppdateringar

### Visningsformatering

Funktionen `getFormulaDisplayValue` formaterar de beräknade resultaten baserat på visningstyp:
- **NUMMER**: Visas som ett vanligt nummer med valfri precision
- **PROCENT**: Lägger till % suffix med valfri precision  
- **VALUTA**: Formaterar med den angivna valutakoden

## Formelresultatlager

Resultat lagras i `formulaResult` fältet:

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

## Svarfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Formelfältets definition |
| `number` | Float | Beräknat numeriskt resultat |
| `formulaResult` | JSON | Fullständigt resultat med visningsformatering |
| `todo` | Todo! | Den post detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast beräknades |

## Datakontext

### Diagramdatakälla

Formelfält fungerar inom diagramdatakällans kontext:
- Formler aggregerar värden från anpassade fält över todos i ett projekt
- Aggregationsfunktionen som anges i `display.function` bestämmer beräkningen
- Resultat beräknas med hjälp av SQL-aggregationsfunktioner (medel, summa, antal, etc.)
- Beräkningar utförs på databasnivå för effektivitet

## Vanliga Formel Exempel

### Total Budget (Diagramvisning)

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

### Genomsnittligt Resultat (Diagramvisning)

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

### Uppgiftsantal (Diagramvisning)

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

## Obligatoriska Behörigheter

Operationer med anpassade fält följer standard rollbaserade behörigheter:

| Åtgärd | Obligatorisk Roll |
|--------|-------------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Notera**: De specifika roller som krävs beror på din projekts anpassade rollkonfiguration. Det finns inga speciella behörighetskonstanter som CUSTOM_FIELDS_CREATE.

## Felhantering

### Valideringsfel
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

### Anpassat Fält Inte Hittat
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

### Formel Design
- Använd tydliga, beskrivande namn för formelfält
- Lägg till beskrivningar som förklarar beräkningslogiken
- Testa formler med exempeldata innan distribution
- Håll formler enkla och läsbara

### Prestandaoptimering
- Undvik djupt nästlade formelberoenden
- Använd specifika fältreferenser istället för jokertecken
- Överväg cachingstrategier för komplexa beräkningar
- Övervaka formelns prestanda i stora projekt

### Datakvalitet
- Validera källdata innan användning i formler
- Hantera tomma eller null-värden på lämpligt sätt
- Använd lämplig precision för visningstyper
- Överväg kantfall i beräkningar

## Vanliga Användningsfall

1. **Finansiell Spårning**
   - Budgetberäkningar
   - Resultat-/förlustrapporter
   - Kostnadsanalys
   - Intäktsprognoser

2. **Projektledning**
   - Slutförandeprocent
   - Resursutnyttjande
   - Tidslinjeberäkningar
   - Prestationsmått

3. **Kvalitetskontroll**
   - Genomsnittliga poäng
   - Godkännande/underkännande räntor
   - Kvalitetsmått
   - Efterlevnadsspårning

4. **Affärsintelligens**
   - KPI-beräkningar
   - Trendanalys
   - Jämförande mått
   - Instrumentpanelvärden

## Begränsningar

- Formler är endast för diagram/instrumentpanelaggregat, inte todo-nivåberäkningar
- Begränsat till de sju stödda aggregationsfunktionerna (SUM, MEDEL, etc.)
- Inga komplexa matematiska uttryck eller fält-till-fält beräkningar
- Kan inte referera till flera fält i en enda formel
- Resultat är endast synliga i diagram och instrumentpaneler
- Fältet `logic` är endast för visningstext, inte faktisk beräkningslogik

## Relaterade Resurser

- [Nummerfält](/api/5.custom%20fields/number) - För statiska numeriska värden
- [Valutafält](/api/5.custom%20fields/currency) - För monetära värden
- [Referensfält](/api/5.custom%20fields/reference) - För data över projekt
- [Uppslagsfält](/api/5.custom%20fields/lookup) - För aggregerad data
- [Översikt över Anpassade Fält](/api/5.custom%20fields/2.list-custom-fields) - Allmänna begrepp