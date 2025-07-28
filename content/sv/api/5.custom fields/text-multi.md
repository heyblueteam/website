---
title: Fler-radig text anpassat fält
description: Skapa fler-radiga textfält för längre innehåll som beskrivningar, anteckningar och kommentarer
---

Fler-radiga text anpassade fält gör att du kan lagra längre textinnehåll med radbrytningar och formatering. De är idealiska för beskrivningar, anteckningar, kommentarer eller vilken textdata som helst som behöver flera rader.

## Grundläggande exempel

Skapa ett enkelt fler-radigt textfält:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Avancerat exempel

Skapa ett fler-radigt textfält med beskrivning:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `type` | CustomFieldType! | ✅ Ja | Måste vara `TEXT_MULTI` |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera:** `projectId` skickas som ett separat argument till mutation, inte som en del av inmatningsobjektet. Alternativt kan projektkontexten bestämmas från `X-Bloo-Project-ID`-huvudet i din GraphQL-förfrågan.

## Ställa in textvärden

För att ställa in eller uppdatera ett fler-radigt textvärde på en post:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput-parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade textfältet |
| `text` | String | Nej | Fler-radigt textinnehåll att lagra |

## Skapa poster med textvärden

När du skapar en ny post med fler-radiga textvärden:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
    }
  }
}
```

## Svarsfält

### TodoCustomField-svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `text` | String | Det lagrade fler-radiga textinnehållet |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

## Textvalidering

### Formulärvalidering
När fler-radiga textfält används i formulär:
- Ledande och avslutande mellanslag trimmas automatiskt
- Obligatorisk validering tillämpas om fältet är markerat som obligatoriskt
- Ingen specifik formatvalidering tillämpas

### Valideringsregler
- Accepterar vilket stränginnehåll som helst inklusive radbrytningar
- Inga teckenlängdsbegränsningar (upp till databasbegränsningar)
- Stöder Unicode-tecken och specialsymboler
- Radbrytningar bevaras i lagring

### Giltiga textexempel
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Viktiga anteckningar

### Lagringskapacitet
- Lagrade med MySQL `MediumText` typ
- Stöder upp till 16MB textinnehåll
- Radbrytningar och formatering bevaras
- UTF-8-kodning för internationella tecken

### Direkt API vs Formulär
- **Formulär**: Automatisk trimning av mellanslag och obligatorisk validering
- **Direkt API**: Text lagras exakt som den tillhandahålls
- **Rekommendation**: Använd formulär för användarinmatning för att säkerställa konsekvent formatering

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI**: Fler-radigt textområde, idealiskt för längre innehåll
- **TEXT_SINGLE**: En-radigt textfält, idealiskt för korta värden
- **Backend**: Båda typerna är identiska - samma lagringsfält, validering och bearbetning
- **Frontend**: Olika UI-komponenter för datainmatning (textområde vs inmatningsfält)
- **Viktigt**: Åtskillnaden mellan TEXT_MULTI och TEXT_SINGLE finns enbart för UI-ändamål

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk behörighet |
|--------|------------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Bästa praxis

### Innehållsorganisation
- Använd konsekvent formatering för strukturerat innehåll
- Överväg att använda markdown-liknande syntax för läsbarhet
- Dela upp långt innehåll i logiska sektioner
- Använd radbrytningar för att förbättra läsbarheten

### Datainmatning
- Ge tydliga fälbeskrivningar för att vägleda användare
- Använd formulär för användarinmatning för att säkerställa validering
- Överväg teckenbegränsningar baserat på ditt användningsfall
- Validera innehållsformatet i din applikation om det behövs

### Prestandaöverväganden
- Mycket långt textinnehåll kan påverka frågeprestanda
- Överväg paginering för att visa stora textfält
- Indexöverväganden för sökfunktionalitet
- Övervaka lagringsanvändning för fält med stort innehåll

## Filtrering och sökning

### Innehåller sökning
Fler-radiga textfält stöder delsträngssökning genom anpassade fältfilter:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Sökfunktioner
- Delsträngsmatchning inom textfält med `CONTAINS`-operatorn
- Skiftlägesokänslig sökning med `NCONTAINS`-operatorn
- Exakt matchning med `IS`-operatorn
- Negativ matchning med `NOT`-operatorn
- Sökningar över alla textlinjer
- Stöder delvis ordmatchning

## Vanliga användningsfall

1. **Projektledning**
   - Uppgiftsbeskrivningar
   - Projektkrav
   - Mötesanteckningar
   - Statusuppdateringar

2. **Kundsupport**
   - Problembeskrivningar
   - Lösningsanteckningar
   - Kundfeedback
   - Kommunikationsloggar

3. **Innehållshantering**
   - Artikelinnehåll
   - Produktbeskrivningar
   - Användarkommentarer
   - Recensionsdetaljer

4. **Dokumentation**
   - Processbeskrivningar
   - Instruktioner
   - Riktlinjer
   - Referensmaterial

## Integrationsfunktioner

### Med automatiseringar
- Utlösa åtgärder när textinnehållet ändras
- Extrahera nyckelord från textinnehåll
- Skapa sammanfattningar eller meddelanden
- Bearbeta textinnehåll med externa tjänster

### Med uppslag
- Referera till textdata från andra poster
- Sammanställa textinnehåll från flera källor
- Hitta poster efter textinnehåll
- Visa relaterad textinformation

### Med formulär
- Automatisk trimning av mellanslag
- Validering av obligatoriska fält
- Fler-radigt textområde UI
- Teckenantalvisning (om konfigurerad)

## Begränsningar

- Ingen inbyggd textformatering eller rik textredigering
- Ingen automatisk länkdetektering eller konvertering
- Ingen stavningskontroll eller grammatikvalidering
- Ingen inbyggd textanalys eller bearbetning
- Ingen versionshantering eller ändringsspårning
- Begränsade sökmöjligheter (ingen fulltextsökning)
- Ingen innehållskomprimering för mycket stor text

## Relaterade resurser

- [En-radiga textfält](/api/custom-fields/text-single) - För korta textvärden
- [E-postfält](/api/custom-fields/email) - För e-postadresser
- [URL-fält](/api/custom-fields/url) - För webbplatsadresser
- [Översikt över anpassade fält](/api/custom-fields/2.list-custom-fields) - Allmänna koncept