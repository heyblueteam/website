---
title: Flerval Anpassat Fält
description: Skapa flervalsfält för att låta användare välja flera alternativ från en fördefinierad lista
---

Flervalsanpassade fält gör det möjligt för användare att välja flera alternativ från en fördefinierad lista. De är idealiska för kategorier, taggar, färdigheter, funktioner eller något scenario där flera val behövs från en kontrollerad uppsättning alternativ.

## Grundläggande Exempel

Skapa ett enkelt flervalsfält:

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

## Avancerat Exempel

Skapa ett flervalsfält och lägg sedan till alternativ separat:

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

## Inmatningsparametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för flervalsfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `SELECT_MULTI` |
| `description` | String | Nej | Hjälptext som visas för användare |
| `projectId` | String! | ✅ Ja | ID för projektet för detta fält |

### CreateCustomFieldOptionInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ Ja | ID för det anpassade fältet |
| `title` | String! | ✅ Ja | Visningstext för alternativet |
| `color` | String | Nej | Färg för alternativet (valfri sträng) |
| `position` | Float | Nej | Sorteringsordning för alternativet |

## Lägga till Alternativ till Befintliga Fält

Lägg till nya alternativ till ett befintligt flervalsfält:

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

## Ställa in Flervalsvärden

För att ställa in flera valda alternativ på en post:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det flervalsanpassade fältet |
| `customFieldOptionIds` | [String!] | ✅ Ja | Array av alternativ-ID:n som ska väljas |

## Skapa Poster med Flervalsvärden

När du skapar en ny post med flervalsvärden:

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

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `selectedOptions` | [CustomFieldOption!] | Array av valda alternativ |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

### CustomFieldOption Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för alternativet |
| `title` | String! | Visningstext för alternativet |
| `color` | String | Hex färgkod för visuell representation |
| `position` | Float | Sorteringsordning för alternativet |
| `customField` | CustomField! | Det anpassade fältet som detta alternativ tillhör |

### CustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältet |
| `name` | String! | Visningsnamn för flervalsfältet |
| `type` | CustomFieldType! | Alltid `SELECT_MULTI` |
| `description` | String | Hjälptext för fältet |
| `customFieldOptions` | [CustomFieldOption!] | Alla tillgängliga alternativ |

## Värdeformat

### Inmatningsformat
- **API Parameter**: Array av alternativ-ID:n (`["option1", "option2", "option3"]`)
- **Strängformat**: Komma-separerade alternativ-ID:n (`"option1,option2,option3"`)

### Utdataformat
- **GraphQL Svar**: Array av CustomFieldOption objekt
- **Aktivitetslogg**: Komma-separerade alternativtitlar
- **Automationsdata**: Array av alternativtitlar

## Hantering av Alternativ

### Uppdatera Alternativsegenskaper
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

### Ta Bort Alternativ
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### Omordna Alternativ
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

## Valideringsregler

### Alternativsvalidering
- Alla angivna alternativ-ID:n måste existera
- Alternativen måste tillhöra det angivna anpassade fältet
- Endast SELECT_MULTI fält kan ha flera alternativ valda
- Tom array är giltig (inga val)

### Fältvalidering
- Måste ha minst ett alternativ definierat för att vara användbart
- Alternativstitlar måste vara unika inom fältet
- Färgfältet accepterar vilket strängvärde som helst (ingen hexvalidering)

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Felrespons

### Ogiltigt Alternativ-ID
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

### Alternativ Tillhör Inte Fält
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

### Fält Inte Hittat
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

### Flera Alternativ på Icke-Multi Fält
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

## Bästa Praxis

### Alternativsdesign
- Använd beskrivande, koncisa alternativstitlar
- Tillämpa konsekventa färgkodningsscheman
- Håll alternativlistor hanterbara (vanligtvis 3-20 alternativ)
- Ordna alternativ logiskt (alfabetiskt, efter frekvens, etc.)

### Datamanagement
- Granska och rensa oanvända alternativ periodiskt
- Använd konsekventa namngivningskonventioner över projekt
- Tänk på alternativens återanvändbarhet när du skapar fält
- Planera för alternativuppdateringar och migreringar

### Användarupplevelse
- Ge tydliga fälbeskrivningar
- Använd färger för att förbättra visuell distinktion
- Gruppera relaterade alternativ tillsammans
- Tänk på standardval för vanliga fall

## Vanliga Användningsfall

1. **Projektledning**
   - Uppgiftskategorier och taggar
   - Prioritetsnivåer och typer
   - Tilldelningar av teammedlemmar
   - Statusindikatorer

2. **Innehållshantering**
   - Artikelkategorier och ämnen
   - Innehållstyper och format
   - Publikationskanaler
   - Godkännandearbetsflöden

3. **Kundsupport**
   - Probleminriktningar och typer
   - Påverkade produkter eller tjänster
   - Lösningsmetoder
   - Kundsegment

4. **Produktutveckling**
   - Funktionskategorier
   - Tekniska krav
   - Testmiljöer
   - Utgivningskanaler

## Integrationsfunktioner

### Med Automatiseringar
- Utlös åtgärder när specifika alternativ väljs
- Rutta arbete baserat på valda kategorier
- Skicka meddelanden för högprioriterade val
- Skapa uppföljningsuppgifter baserat på alternativkombinationer

### Med Sökningar
- Filtrera poster efter valda alternativ
- Aggregatdata över alternativval
- Referera till alternativdata från andra poster
- Skapa rapporter baserat på alternativkombinationer

### Med Formulär
- Flervalsinmatningskontroller
- Alternativvalidering och filtrering
- Dynamisk alternativladdning
- Villkorlig fältdisplay

## Aktivitetsövervakning

Ändringar i flervalsfält spåras automatiskt:
- Visar tillagda och borttagna alternativ
- Visar alternativtitlar i aktivitetsloggen
- Tidsstämplar för alla valändringar
- Användarattribution för modifieringar

## Begränsningar

- Maximalt praktiskt antal alternativ beror på UI-prestanda
- Ingen hierarkisk eller nästlad alternativstruktur
- Alternativ delas över alla poster som använder fältet
- Ingen inbyggd alternativanalys eller användningsspårning
- Färgfältet accepterar vilken sträng som helst (ingen hexvalidering)
- Kan inte ställa in olika behörigheter per alternativ
- Alternativ måste skapas separat, inte inline med fältets skapande
- Ingen dedikerad omordningsmutation (använd editCustomFieldOption med position)

## Relaterade Resurser

- [Enkelvalsfält](/api/custom-fields/select-single) - För enskilda val
- [Checkboxfält](/api/custom-fields/checkbox) - För enkla booleska val
- [Textfält](/api/custom-fields/text-single) - För fri textinmatning
- [Översikt över Anpassade Fält](/api/custom-fields/2.list-custom-fields) - Allmänna begrepp