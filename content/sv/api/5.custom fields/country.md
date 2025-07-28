---
title: Anpassat fält för land
description: Skapa fält för landsval med validering av ISO landskoder
---

Anpassade fält för land gör att du kan lagra och hantera landinformation för poster. Fältet stöder både landsnamn och ISO Alpha-2 landskoder.

**Viktigt**: Validering och konverteringsbeteende för länder skiljer sig avsevärt mellan mutationer:
- **createTodo**: Validerar och konverterar automatiskt landsnamn till ISO-koder
- **setTodoCustomField**: Accepterar vilket värde som helst utan validering

## Grundläggande exempel

Skapa ett enkelt landsfält:

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

## Avancerat exempel

Skapa ett landsfält med beskrivning:

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

## Inmatningsparametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för landsfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `COUNTRY` |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: `projectId` skickas inte i inmatningen utan bestäms av GraphQL-kontexten (vanligtvis från begärningshuvuden eller autentisering).

## Ställa in landsvärden

Landsfält lagrar data i två databasfält:
- **`countryCodes`**: Lagrar ISO Alpha-2 landskoder som en kommaseparerad sträng i databasen (återges som array via API)
- **`text`**: Lagrar visningstext eller landsnamn som en sträng

### Förstå parametrarna

Mutation `setTodoCustomField` accepterar två valfria parametrar för landsfält:

| Parameter | Typ | Obligatorisk | Beskrivning | Vad den gör |
|-----------|------|--------------|-------------|--------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras | - |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade landsfältet | - |
| `countryCodes` | [String!] | Nej | Array av ISO Alpha-2 landskoder | Stored in the `countryCodes` field |
| `text` | String | Nej | Visningstext eller landsnamn | Stored in the `text` field |

**Viktigt**: 
- I `setTodoCustomField`: Båda parametrarna är valfria och lagras oberoende
- I `createTodo`: Systemet ställer automatiskt in båda fälten baserat på din inmatning (du kan inte styra dem oberoende)

### Alternativ 1: Använda endast landskoder

Lagra validerade ISO-koder utan visningstext:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Resultat: `countryCodes` = `["US"]`, `text` = `null`

### Alternativ 2: Använda endast text

Lagra visningstext utan validerade koder:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Resultat: `countryCodes` = `null`, `text` = `"United States"`

**Notera**: När du använder `setTodoCustomField` sker ingen validering oavsett vilken parameter du använder. Värdena lagras exakt som de anges.

### Alternativ 3: Använda båda (Rekommenderas)

Lagra både validerade koder och visningstext:

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

Resultat: `countryCodes` = `["US"]`, `text` = `"United States"`

### Flera länder

Lagra flera länder med hjälp av arrayer:

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

## Skapa poster med landsvärden

När du skapar poster utför mutation `createTodo` **automatiskt validering och konvertering** av landsvärden. Detta är den enda mutation som utför validering av länder:

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

### Accepterade inmatningsformat

| Inmatningstyp | Exempel | Resultat |
|----------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Ej stöds** - behandlas som ett ogiltigt värde |
| Mixed format | `"United States, CA"` | **Ej stöds** - behandlas som ett ogiltigt värde |

## Svarsfält

### TodoCustomField-svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `text` | String | Visningstext (landsnamn) |
| `countryCodes` | [String!] | Array av ISO Alpha-2 landskoder |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast modifierades |

## Landsstandarder

Blue använder **ISO 3166-1 Alpha-2**-standarden för landskoder:

- Tvåbokstaviga landskoder (t.ex. US, GB, FR, DE)
- Validering med hjälp av `i18n-iso-countries`-biblioteket **sker endast i createTodo**
- Stöder alla officiellt erkända länder

### Exempel på landskoder

| Land | ISO-kod |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

För den kompletta officiella listan över ISO 3166-1 alpha-2 landskoder, besök [ISO Online Browsing Platform](https://www.iso.org/obp/ui/#search/code/).

## Validering

**Validering sker endast i `createTodo`-mutation**:

1. **Giltig ISO-kod**: Accepterar vilken giltig ISO Alpha-2-kod som helst
2. **Landsnamn**: Konverterar automatiskt erkända landsnamn till koder
3. **Ogiltig inmatning**: Kastade `CustomFieldValueParseError` för oigenkända värden

**Notera**: `setTodoCustomField`-mutation utför INGEN validering och accepterar vilket strängvärde som helst.

### Exempel på fel

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

## Integrationsfunktioner

### Uppslagsfält
Landsfält kan refereras av LOOKUP-anpassade fält, vilket gör att du kan hämta landsdata från relaterade poster.

### Automatiseringar
Använd landsvärden i automatiseringsvillkor:
- Filtrera åtgärder efter specifika länder
- Skicka meddelanden baserat på land
- Rätta uppgifter baserat på geografiska regioner

### Formulär
Landsfält i formulär validerar automatiskt användarinmatning och konverterar landsnamn till koder.

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk behörighet |
|--------|-------------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Felrespons

### Ogiltigt landsvärde
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

### Typkonflikt för fält
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

## Bästa praxis

### Inmatningshantering
- Använd `createTodo` för automatisk validering och konvertering
- Använd `setTodoCustomField` med försiktighet eftersom det kringgår validering
- Överväg att validera inmatningar i din applikation innan du använder `setTodoCustomField`
- Visa fullständiga landsnamn i UI för tydlighet

### Datakvalitet
- Validera landsinmatningar vid inmatningspunkten
- Använd konsekventa format över ditt system
- Överväg regionala grupperingar för rapportering

### Flera länder
- Använd arraystöd i `setTodoCustomField` för flera länder
- Flera länder i `createTodo` stöds **inte** via värdefältet
- Lagra landskoder som array i `setTodoCustomField` för korrekt hantering

## Vanliga användningsfall

1. **Kundhantering**
   - Kundens huvudkontor
   - Fraktdestinationer
   - Skattejurisdiktioner

2. **Projektspårning**
   - Projektplats
   - Plats för teammedlemmar
   - Marknadsmål

3. **Efterlevnad och juridik**
   - Reglerande jurisdiktioner
   - Krav på datalagring
   - Exportkontroller

4. **Försäljning och marknadsföring**
   - Territorietilldelningar
   - Marknadssegmentering
   - Kampanjinriktning

## Begränsningar

- Stöder endast ISO 3166-1 Alpha-2-koder (2-bokstavskoder)
- Ingen inbyggd support för landsunderavdelningar (delstater/provinser)
- Inga automatiska landsflaggsikoner (endast textbaserade)
- Kan inte validera historiska landskoder
- Ingen inbyggd region- eller kontinentgruppering
- **Validering fungerar endast i `createTodo`, inte i `setTodoCustomField`**
- **Flera länder stöds inte i `createTodo` värdefält**
- **Landskoder lagras som kommaseparerad sträng, inte som verklig array**

## Relaterade resurser

- [Översikt över anpassade fält](/custom-fields/list-custom-fields) - Allmänna koncept för anpassade fält
- [Uppslagsfält](/api/custom-fields/lookup) - Referera till landsdata från andra poster
- [Formulär-API](/api/forms) - Inkludera landsfält i anpassade formulär