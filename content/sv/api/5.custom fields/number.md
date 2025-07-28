---
title: Nummer Anpassat Fält
description: Skapa nummerfält för att lagra numeriska värden med valfria min/max begränsningar och prefixformatering
---

Nummeranpassade fält gör att du kan lagra numeriska värden för poster. De stödjer valideringsbegränsningar, decimalprecision och kan användas för kvantiteter, poäng, mätningar eller vilken numerisk data som helst som inte kräver speciell formatering.

## Grundläggande Exempel

Skapa ett enkelt nummerfält:

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

## Avancerat Exempel

Skapa ett nummerfält med begränsningar och prefix:

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

## Indata Parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för nummerfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `NUMBER` |
| `projectId` | String! | ✅ Ja | ID för projektet där fältet ska skapas |
| `min` | Float | Nej | Minimi värde begränsning (endast UI-vägledning) |
| `max` | Float | Nej | Maximivärde begränsning (endast UI-vägledning) |
| `prefix` | String | Nej | Visningsprefix (t.ex. "#", "~", "$") |
| `description` | String | Nej | Hjälptext som visas för användare |

## Ställa in Nummer Värden

Nummerfält lagrar decimalvärden med valfri validering:

### Enkelt Nummer Värde

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Heltal Värde

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för nummeranpassat fält |
| `number` | Float | Nej | Numeriskt värde att lagra |

## Värde Begränsningar

### Min/Max Begränsningar (UI Vägledning)

**Viktigt**: Min/max begränsningar lagras men TILLÄMPAS INTE på serversidan. De fungerar som UI-vägledning för frontend-applikationer.

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

**Klient-Side Validering Krävs**: Frontend-applikationer måste implementera valideringslogik för att tillämpa min/max begränsningar.

### Stödda Värdetyper

| Typ | Exempel | Beskrivning |
|------|---------|-------------|
| Integer | `42` | Heltal |
| Decimal | `42.5` | Tal med decimaler |
| Negative | `-10` | Negativa värden (om ingen min begränsning) |
| Zero | `0` | Nollvärde |

**Notera**: Min/max begränsningar valideras INTE på serversidan. Värden utanför det angivna intervallet kommer att accepteras och lagras.

## Skapa Poster med Nummer Värden

När du skapar en ny post med nummer värden:

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

### Stödda Indataformat

När du skapar poster, använd `number` parametern (inte `value`) i arrayen för anpassade fält:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `number` | Float | Det numeriska värdet |
| `todo` | Todo! | Den post detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

### CustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältdefinitionen |
| `name` | String! | Visningsnamn för fältet |
| `type` | CustomFieldType! | Alltid `NUMBER` |
| `min` | Float | Minimi tillåtna värde |
| `max` | Float | Maximalt tillåtna värde |
| `prefix` | String | Visningsprefix |
| `description` | String | Hjälptext |

**Notera**: Om nummervärdet inte är inställt, kommer `number` fältet att vara `null`.

## Filtrering och Frågor

Nummerfält stödjer omfattande numerisk filtrering:

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

### Stödda Operatörer

| Operatör | Beskrivning | Exempel |
|----------|-------------|---------|
| `EQ` | Lika med | `number = 42` |
| `NE` | Inte lika med | `number ≠ 42` |
| `GT` | Större än | `number > 42` |
| `GTE` | Större än eller lika med | `number ≥ 42` |
| `LT` | Mindre än | `number < 42` |
| `LTE` | Mindre än eller lika med | `number ≤ 42` |
| `IN` | I array | `number in [1, 2, 3]` |
| `NIN` | Inte i array | `number not in [1, 2, 3]` |
| `IS` | Är null/inte null | `number is null` |

### Intervall Filtrering

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

## Visningsformatering

### Med Prefix

Om ett prefix är inställt, kommer det att visas:

| Värde | Prefix | Visning |
|-------|--------|---------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Decimal Precision

Nummer behåller sin decimalprecision:

| Indata | Lagrad | Visad |
|-------|--------|-----------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|-------------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Fel Svar

### Ogiltigt Nummerformat
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

### Fält Inte Hittat
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

**Notera**: Min/max valideringsfel inträffar INTE på serversidan. Begränsningsvalidering måste implementeras i din frontend-applikation.

### Inte Ett Nummer
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

### Begränsningsdesign
- Sätt realistiska min/max värden för UI-vägledning
- Implementera klient-sida validering för att tillämpa begränsningar
- Använd begränsningar för att ge användarfeedback i formulär
- Överväg om negativa värden är giltiga för ditt användningsfall

### Värde Precision
- Använd lämplig decimalprecision för dina behov
- Överväg avrundning för visningsändamål
- Var konsekvent med precision över relaterade fält

### Visningsförbättring
- Använd meningsfulla prefix för kontext
- Överväg enheter i fältnamn (t.ex. "Vikt (kg)")
- Ge tydliga beskrivningar för valideringsregler

## Vanliga Användningsfall

1. **Poängsystem**
   - Prestandabetyg
   - Kvalitetspoäng
   - Prioritetsnivåer
   - Kundnöjdhetsbetyg

2. **Mätningar**
   - Kvantiteter och belopp
   - Dimensioner och storlekar
   - Tider (i numeriskt format)
   - Kapaciteter och gränser

3. **Affärsmetrik**
   - Intäktsfigurer
   - Konverteringsgrader
   - Budgetallokeringar
   - Målantal

4. **Teknisk Data**
   - Versionsnummer
   - Konfigurationsvärden
   - Prestandamått
   - Tröskelinställningar

## Integrationsfunktioner

### Med Diagram och Instrumentpaneler
- Använd NUMMER fält i diagramberäkningar
- Skapa numeriska visualiseringar
- Spåra trender över tid

### Med Automatiseringar
- Utlös åtgärder baserat på nummertrösklar
- Uppdatera relaterade fält baserat på nummerändringar
- Skicka meddelanden för specifika värden

### Med Uppslag
- Aggregatnummer från relaterade poster
- Beräkna totalsummor och genomsnitt
- Hitta min/max värden över relationer

### Med Diagram
- Skapa numeriska visualiseringar
- Spåra trender över tid
- Jämför värden över poster

## Begränsningar

- **Ingen server-sida validering** av min/max begränsningar
- **Klient-sida validering krävs** för att tillämpa begränsningar
- Ingen inbyggd valutaformatering (använd VALUTA typ istället)
- Ingen automatisk procentsymbol (använd PROCENT typ istället)
- Ingen enhetskonverteringskapacitet
- Decimalprecision begränsad av databasens Decimal typ
- Ingen matematisk formelutvärdering i fältet självt

## Relaterade Resurser

- [Översikt över Anpassade Fält](/api/custom-fields/1.index) - Allmänna koncept för anpassade fält
- [Valuta Anpassat Fält](/api/custom-fields/currency) - För monetära värden
- [Procent Anpassat Fält](/api/custom-fields/percent) - För procentvärden
- [Automatiseringar API](/api/automations/1.index) - Skapa nummerbaserade automatiseringar