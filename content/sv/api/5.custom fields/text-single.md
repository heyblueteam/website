---
title: Anpassat fält för enradig text
description: Skapa enradiga textfält för korta textvärden som namn, titlar och etiketter
---

Enradiga textfält gör att du kan lagra korta textvärden avsedda för inmatning på en rad. De är idealiska för namn, titlar, etiketter eller vilken textdata som helst som ska visas på en enda rad.

## Grundläggande exempel

Skapa ett enkelt enradigt textfält:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Avancerat exempel

Skapa ett enradigt textfält med beskrivning:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
| `name` | String! | ✅ Ja | Visningsnamn för textfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `TEXT_SINGLE` |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: Projektkontexten bestäms automatiskt från dina autentiseringshuvuden. Ingen `projectId` parameter behövs.

## Ställa in textvärden

För att ställa in eller uppdatera ett enradigt textvärde på en post:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade textfältet |
| `text` | String | Nej | Innehåll av enradig text att lagra |

## Skapa poster med textvärden

När du skapar en ny post med enradiga textvärden:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Svarsfält

### TodoCustomField svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | ID! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen (innehåller textvärdet) |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast modifierades |

**Viktigt**: Textvärden nås genom `customField.value.text` fältet, inte direkt på TodoCustomField.

## Fråga textvärden

När du frågar poster med anpassade textfält, nå texten genom `customField.value.text` vägen:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

Svaret kommer att inkludera texten i den inbäddade strukturen:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Textvalidering

### Formvalidering
När enradiga textfält används i formulär:
- Ledande och avslutande blanksteg trimmas automatiskt
- Obligatorisk validering tillämpas om fältet är markerat som obligatoriskt
- Ingen specifik formatvalidering tillämpas

### Valideringsregler
- Accepterar vilket stränginnehåll som helst inklusive radbrytningar (även om det inte rekommenderas)
- Inga teckengränser (upp till databasens gränser)
- Stöder Unicode-tecken och specialsymboler
- Radbrytningar bevaras men är inte avsedda för denna fälttyp

### Typiska textexempel
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Viktiga anteckningar

### Lagringskapacitet
- Lagrade med MySQL `MediumText` typ
- Stöder upp till 16MB textinnehåll
- Identisk lagring som fleradiga textfält
- UTF-8-kodning för internationella tecken

### Direkt API vs formulär
- **Formulär**: Automatisk trimning av blanksteg och obligatorisk validering
- **Direkt API**: Text lagras exakt som den anges
- **Rekommendation**: Använd formulär för användarinmatning för att säkerställa konsekvent formatering

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: Enradig textinmatning, idealisk för korta värden
- **TEXT_MULTI**: Fleradig textinmatning, idealisk för längre innehåll
- **Backend**: Båda använder identisk lagring och validering
- **Frontend**: Olika UI-komponenter för datainmatning
- **Avsikt**: TEXT_SINGLE är semantiskt avsett för enradiga värden

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk behörighet |
|--------|-------------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## Felrespons

### Obligatorisk fältvalidering (endast formulär)
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Fält hittades inte
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

## Bästa praxis

### Innehållsriktlinjer
- Håll texten kort och lämplig för enradig visning
- Undvik radbrytningar för avsedd enradig visning
- Använd konsekvent formatering för liknande datatyper
- Tänk på teckengränser baserat på dina UI-krav

### Datainmatning
- Ge tydliga fälbeskrivningar för att vägleda användare
- Använd formulär för användarinmatning för att säkerställa validering
- Validera innehållsformat i din applikation om det behövs
- Överväg att använda rullgardinsmenyer för standardiserade värden

### Prestandaöverväganden
- Enradiga textfält är lätta och presterar bra
- Överväg indexering för ofta sökta fält
- Använd lämpliga visningsbredder i din UI
- Övervaka innehållslängd för visningsändamål

## Filtrering och sökning

### Innehåller sökning
Enradiga textfält stöder delsträngsökning:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Sökfunktioner
- Skiftlägesokänslig delsträngsmatchning
- Stöder partiell ordmatchning
- Exakt värdematchning
- Ingen fulltextsökning eller rankning

## Vanliga användningsfall

1. **Identifierare och koder**
   - Produkt-SKU:er
   - Ordernummer
   - Referenskoder
   - Versionsnummer

2. **Namn och titlar**
   - Klientnamn
   - Projekttitlar
   - Produktnamn
   - Kategorietiketter

3. **Korta beskrivningar**
   - Kortfattade sammanfattningar
   - Statusetiketter
   - Prioritetsindikatorer
   - Klassificeringstaggar

4. **Externa referenser**
   - Biljettnummer
   - Fakturareferenser
   - ID:n för externa system
   - Dokumentnummer

## Integrationsfunktioner

### Med uppslag
- Referens textdata från andra poster
- Hitta poster efter textinnehåll
- Visa relaterad textinformation
- Aggregat textvärden från flera källor

### Med formulär
- Automatisk trimning av blanksteg
- Obligatorisk fältvalidering
- Enradig textinmatnings-UI
- Teckengränsvisning (om konfigurerad)

### Med importer/exporter
- Direkt CSV-kolumnmappning
- Automatisk tilldelning av textvärden
- Stöd för bulkdataimport
- Export till kalkylbladsformat

## Begränsningar

### Automatiseringsbegränsningar
- Inte direkt tillgängliga som automatiseringstriggerfält
- Kan inte användas i automatiseringsfältuppdateringar
- Kan refereras i automatiseringsvillkor
- Tillgängliga i e-postmallar och webbhooks

### Allmänna begränsningar
- Ingen inbyggd textformatering eller stil
- Ingen automatisk validering utöver obligatoriska fält
- Ingen inbyggd unikhetskontroll
- Ingen innehållskomprimering för mycket stor text
- Ingen versionshantering eller ändringsspårning
- Begränsade sökfunktioner (ingen fulltextsökning)

## Relaterade resurser

- [Fleradiga textfält](/api/custom-fields/text-multi) - För längre textinnehåll
- [E-postfält](/api/custom-fields/email) - För e-postadresser
- [URL-fält](/api/custom-fields/url) - För webbadresser
- [Unika ID-fält](/api/custom-fields/unique-id) - För automatiskt genererade identifierare
- [Översikt över anpassade fält](/api/custom-fields/list-custom-fields) - Allmänna koncept