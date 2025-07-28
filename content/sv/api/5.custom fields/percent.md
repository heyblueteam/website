---
title: Procent Anpassat Fält
description: Skapa procentfält för att lagra numeriska värden med automatisk hantering av % symbol och visningsformat
---

Procentanpassade fält gör att du kan lagra procentvärden för poster. De hanterar automatiskt % symbolen för inmatning och visning, samtidigt som de lagrar det råa numeriska värdet internt. Perfekt för slutförandegrader, framgångsgrader eller andra procentbaserade mätvärden.

## Grundläggande Exempel

Skapa ett enkelt procentfält:

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

## Avancerat Exempel

Skapa ett procentfält med beskrivning:

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

## Inmatningsparametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för procentfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `PERCENT` |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: Projektkontext bestäms automatiskt från dina autentiseringhuvuden. Ingen `projectId` parameter behövs.

**Notera**: PERCENT fält stöder inte min/max begränsningar eller prefixformattering som NUMBER fält.

## Ställa in Procentvärden

Procentfält lagrar numeriska värden med automatisk hantering av % symbol:

### Med Procent Symbol

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

### Direkt Numeriskt Värde

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

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det procentanpassade fältet |
| `number` | Float | Nej | Numeriskt procentvärde (t.ex. 75.5 för 75.5%) |

## Värdelagring och Visning

### Lagringsformat
- **Intern lagring**: Rått numeriskt värde (t.ex. 75.5)
- **Databas**: Lagrad som `Decimal` i `number` kolumn
- **GraphQL**: Återlämnas som `Float` typ

### Visningsformat
- **Användargränssnitt**: Klientapplikationer måste lägga till % symbol (t.ex. "75.5%")
- **Diagram**: Visas med % symbol när utgångstypen är PERCENTAGE
- **API-svar**: Rått numeriskt värde utan % symbol (t.ex. 75.5)

## Skapa Poster med Procentvärden

När du skapar en ny post med procentvärden:

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

### Stödda Inmatningsformat

| Format | Exempel | Resultat |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Notera**: % symbolen tas automatiskt bort från inmatningen och läggs tillbaka vid visning.

## Fråga Procentvärden

När du frågar poster med procentanpassade fält, få åtkomst till värdet genom `customField.value.number` vägen:

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

Svaret kommer att inkludera procenten som ett rått nummer:

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

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | ID! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen (innehåller procentvärdet) |
| `todo` | Todo! | Den post detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

**Viktigt**: Procentvärden nås genom `customField.value.number` fältet. % symbolen ingår inte i lagrade värden och måste läggas till av klientapplikationer för visning.

## Filtrering och Fråga

Procentfält stöder samma filtrering som NUMBER fält:

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

### Stödda Operatörer

| Operatör | Beskrivning | Exempel |
|----------|-------------|---------|
| `EQ` | Lika med | `percentage = 75` |
| `NE` | Inte lika med | `percentage ≠ 75` |
| `GT` | Större än | `percentage > 75` |
| `GTE` | Större än eller lika med | `percentage ≥ 75` |
| `LT` | Mindre än | `percentage < 75` |
| `LTE` | Mindre än eller lika med | `percentage ≤ 75` |
| `IN` | Värde i lista | `percentage in [50, 75, 100]` |
| `NIN` | Värde inte i lista | `percentage not in [0, 25]` |
| `IS` | Kontrollera för null med `values: null` | `percentage is null` |
| `NOT` | Kontrollera för inte null med `values: null` | `percentage is not null` |

### Område Filtrering

För områdefiltrering, använd flera operatörer:

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

## Procentvärdesområden

### Vanliga Områden

| Område | Beskrivning | Användningsfall |
|-------|-------------|-----------------|
| `0-100` | Standardprocent | Completion rates, success rates |
| `0-∞` | Obegränsad procent | Growth rates, performance metrics |
| `-∞-∞` | Vilket värde som helst | Change rates, variance |

### Exempelvärden

| Inmatning | Lagrad | Visning |
|-----------|--------|---------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Diagramaggregat

Procentfält stöder aggregation i instrumentpaneldiagram och rapporter. Tillgängliga funktioner inkluderar:

- `AVERAGE` - Medelprocentvärde
- `COUNT` - Antal poster med värden
- `MIN` - Lägsta procentvärde
- `MAX` - Högsta procentvärde 
- `SUM` - Totalt av alla procentvärden

Dessa aggregationer är tillgängliga vid skapande av diagram och instrumentpaneler, inte i direkta GraphQL-frågor.

## Krävs Behörigheter

| Åtgärd | Krävs Behörighet |
|--------|-------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Fel Svar

### Ogiltigt Procentformat
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

### Inte ett Nummer
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

## Bästa Praxis

### Värdeinmatning
- Låt användare ange med eller utan % symbol
- Validera rimliga intervall för ditt användningsfall
- Ge tydlig kontext om vad 100% representerar

### Visning
- Visa alltid % symbol i användargränssnitt
- Använd lämplig decimalprecision
- Överväg färgkodning för intervall (röd/gul/grön)

### Data Tolkning
- Dokumentera vad 100% betyder i din kontext
- Hantera värden över 100% på lämpligt sätt
- Överväg om negativa värden är giltiga

## Vanliga Användningsfall

1. **Projektledning**
   - Uppgiftsfullföljandegrader
   - Projektframsteg
   - Resursutnyttjande
   - Sprinthastighet

2. **Prestandaövervakning**
   - Framgångsgrader
   - Felgrader
   - Effektivitetsmått
   - Kvalitetspoäng

3. **Finansiella Mått**
   - Tillväxtgrader
   - Vinstmarginaler
   - Rabatter
   - Förändringsprocent

4. **Analys**
   - Konverteringsgrader
   - Klickfrekvenser
   - Engagemangsmått
   - Prestandaindikatorer

## Integrationsfunktioner

### Med Formler
- Referera till PERCENT fält i beräkningar
- Automatisk % symbolformatering i formelutdata
- Kombinera med andra numeriska fält

### Med Automatiseringar
- Utlösa åtgärder baserat på procenttrösklar
- Skicka meddelanden för milstolpsprocent
- Uppdatera status baserat på slutförandegrader

### Med Uppslag
- Aggregatprocent från relaterade poster
- Beräkna genomsnittliga framgångsgrader
- Hitta högst/lägst presterande objekt

### Med Diagram
- Skapa procentbaserade visualiseringar
- Spåra framsteg över tid
- Jämför prestandamått

## Skillnader från NUMBER Fält

### Vad som är Olika
- **Inmatningshantering**: Tar automatiskt bort % symbol
- **Visning**: Lägger automatiskt till % symbol
- **Begränsningar**: Ingen min/max validering
- **Formatering**: Ingen prefixstöd

### Vad som är Detsamma
- **Lagring**: Samma databas kolumn och typ
- **Filtrering**: Samma frågeoperatörer
- **Aggregation**: Samma aggregationsfunktioner
- **Behörigheter**: Samma behörighetsmodell

## Begränsningar

- Inga min/max värdebegränsningar
- Inga prefixformateringsalternativ
- Ingen automatisk validering av 0-100% intervall
- Ingen konvertering mellan procentformat (t.ex. 0.75 ↔ 75%)
- Värden över 100% är tillåtna

## Relaterade Resurser

- [Översikt över Anpassade Fält](/api/custom-fields/list-custom-fields) - Allmänna koncept för anpassade fält
- [Nummer Anpassat Fält](/api/custom-fields/number) - För råa numeriska värden
- [Automatiserings API](/api/automations/index) - Skapa procentbaserade automatiseringar