---
title: URL Anpassat Fält
description: Skapa URL-fält för att lagra webbplatsadresser och länkar
---

URL-anpassade fält gör att du kan lagra webbplatsadresser och länkar i dina poster. De är idealiska för att spåra projektwebbplatser, referenslänkar, dokumentations-URL:er eller andra webbaserade resurser relaterade till ditt arbete.

## Grundläggande Exempel

Skapa ett enkelt URL-fält:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Avancerat Exempel

Skapa ett URL-fält med beskrivning:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
    }
  ) {
    id
    name
    type
    description
  }
}
```

## Indata Parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för URL-fältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `URL` |
| `description` | String | Nej | Hjälptext som visas för användare |

**Obs:** `projectId` skickas som ett separat argument till mutation, inte som en del av indataobjektet.

## Ställa in URL-värden

För att ställa in eller uppdatera ett URL-värde på en post:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade URL-fältet |
| `text` | String! | ✅ Ja | URL-adress att lagra |

## Skapa Poster med URL-värden

När du skapar en ny post med URL-värden:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `text` | String | Den lagrade URL-adressen |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

## URL Validering

### Aktuell Implementering
- **Direkt API**: Ingen URL-formatvalidering tillämpas för närvarande
- **Formulär**: URL-validering är planerad men inte aktiv för närvarande
- **Lagring**: Vilken strängvärde som helst kan lagras i URL-fält

### Planerad Validering
Framtida versioner kommer att inkludera:
- HTTP/HTTPS protokollvalidering
- Validering av giltigt URL-format
- Validering av domännamn
- Automatisk tillägg av protokollprefix

### Rekommenderade URL-format
Även om det för närvarande inte tillämpas, använd dessa standardformat:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Viktiga Noteringar

### Lagringsformat
- URL:er lagras som vanlig text utan modifiering
- Inget automatiskt protokolltillägg (http://, https://)
- Skiftlägeskänslighet bevaras som inmatat
- Ingen URL-kodning/avkodning utförs

### Direkt API vs Formulär
- **Formulär**: Planerad URL-validering (inte aktiv för närvarande)
- **Direkt API**: Ingen validering - vilken text som helst kan lagras
- **Rekommendation**: Validera URL:er i din applikation innan du lagrar

### URL vs Textfält
- **URL**: Semantiskt avsett för webbadresser
- **TEXT_SINGLE**: Allmän enradig text
- **Backend**: För närvarande identisk lagring och validering
- **Frontend**: Olika UI-komponenter för datainmatning

## Obligatoriska Behörigheter

Anpassade fältoperationer använder rollbaserade behörigheter:

| Åtgärd | Obligatorisk Roll |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Obs:** Behörigheter kontrolleras baserat på användarroller i projektet, inte specifika behörighetskonstanter.

## Fel Svar

### Fält Inte Hittat
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

### Obligatorisk Fältvalidering (Endast Formulär)
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

## Bästa Praxis

### URL-formatstandarder
- Inkludera alltid protokoll (http:// eller https://)
- Använd HTTPS när det är möjligt för säkerhet
- Testa URL:er innan du lagrar för att säkerställa att de är tillgängliga
- Överväg att använda förkortade URL:er för visningsändamål

### Datakvalitet
- Validera URL:er i din applikation innan du lagrar
- Kontrollera vanliga stavfel (saknade protokoll, felaktiga domäner)
- Standardisera URL-format över din organisation
- Överväg URL-tillgänglighet och tillgänglighet

### Säkerhetsöverväganden
- Var försiktig med användartillhandahållna URL:er
- Validera domäner om du begränsar till specifika webbplatser
- Överväg URL-skanning för skadligt innehåll
- Använd HTTPS-URL:er när du hanterar känslig data

## Filtrering och Sökning

### Innehåller Sök
URL-fält stöder delsträngsökning:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Sökfunktioner
- Skiftlägesokänslig delsträngsmatchning
- Partiell domänmatchning
- Sökning av sökväg och parametrar
- Ingen protokollspecifik filtrering

## Vanliga Användningsfall

1. **Projektledning**
   - Projektwebbplatser
   - Dokumentationslänkar
   - Repositories URL:er
   - Demowebbplatser

2. **Innehållshantering**
   - Referensmaterial
   - Källor länkar
   - Medieresurser
   - Externa artiklar

3. **Kundsupport**
   - Kundwebbplatser
   - Supportdokumentation
   - Artiklar i kunskapsbasen
   - Videohandledningar

4. **Försäljning & Marknadsföring**
   - Företagswebbplatser
   - Produktsidor
   - Marknadsföringsmaterial
   - Sociala medieprofiler

## Integrationsfunktioner

### Med Uppslag
- Referens-URL:er från andra poster
- Hitta poster efter domän eller URL-mönster
- Visa relaterade webbresurser
- Aggregatlänkar från flera källor

### Med Formulär
- URL-specifika inmatningskomponenter
- Planerad validering för korrekt URL-format
- Länkförhandsgranskning (frontend)
- Klickbara URL-visningar

### Med Rapportering
- Spåra URL-användning och mönster
- Övervaka brutna eller otillgängliga länkar
- Kategorisera efter domän eller protokoll
- Exportera URL-listor för analys

## Begränsningar

### Aktuella Begränsningar
- Ingen aktiv URL-formatvalidering
- Inget automatiskt protokolltillägg
- Ingen länkverifiering eller tillgänglighetskontroll
- Ingen URL-förkortning eller expansion
- Ingen favicon eller förhandsgranskning generering

### Automatiseringsbegränsningar
- Inte tillgänglig som automatiseringstriggerfält
- Kan inte användas i automatiseringsfältuppdateringar
- Kan refereras i automatiseringsvillkor
- Tillgänglig i e-postmallar och webhooks

### Allmänna Begränsningar
- Ingen inbyggd länkförhandsgranskningsfunktionalitet
- Ingen automatisk URL-förkortning
- Ingen klickspårning eller analys
- Ingen URL-utgångskontroll
- Ingen skanning av skadliga URL:er

## Framtida Förbättringar

### Planerade Funktioner
- HTTP/HTTPS protokollvalidering
- Anpassade regex-valideringsmönster
- Automatisk protokollprefix-tillägg
- URL-tillgänglighetskontroll

### Potentiella Förbättringar
- Länkförhandsgranskning generering
- Favicon visning
- URL-förkortningsintegration
- Klickspårningsfunktioner
- Upptäckte brutna länkar

## Relaterade Resurser

- [Textfält](/api/custom-fields/text-single) - För icke-URL textdata
- [E-postfält](/api/custom-fields/email) - För e-postadresser
- [Översikt över Anpassade Fält](/api/custom-fields/2.list-custom-fields) - Allmänna koncept

## Migrering från Textfält

Om du migrerar från textfält till URL-fält:

1. **Skapa URL-fält** med samma namn och konfiguration
2. **Exportera befintliga textvärden** för att verifiera att de är giltiga URL:er
3. **Uppdatera poster** för att använda det nya URL-fältet
4. **Ta bort det gamla textfältet** efter framgångsrik migrering
5. **Uppdatera applikationer** för att använda URL-specifika UI-komponenter

### Migreringsexempel
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```