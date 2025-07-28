---
title: Enkeltval Anpassat Fält
description: Skapa enkeltvalsfält för att låta användare välja ett alternativ från en fördefinierad lista
---

Enkeltval anpassade fält tillåter användare att välja exakt ett alternativ från en fördefinierad lista. De är idealiska för statusfält, kategorier, prioriteringar eller något scenario där endast ett val bör göras från en kontrollerad uppsättning alternativ.

## Grundläggande Exempel

Skapa ett enkelt enkeltvalsfält:

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

## Avancerat Exempel

Skapa ett enkeltvalsfält med fördefinierade alternativ:

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

## Inmatningsparametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för enkeltvalsfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `SELECT_SINGLE` |
| `description` | String | Nej | Hjälptext som visas för användare |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | Nej | Initiala alternativ för fältet |

### CreateCustomFieldOptionInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `title` | String! | ✅ Ja | Visningstext för alternativet |
| `color` | String | Nej | Hex färgkod för alternativet |

## Lägga till Alternativ till Befintliga Fält

Lägg till nya alternativ till ett befintligt enkeltvalsfält:

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

## Ställa In Enkeltvals Värden

För att ställa in det valda alternativet på en post:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det enkeltvalda anpassade fältet |
| `customFieldOptionId` | String | Nej | ID för alternativet som ska väljas (föredras för enkeltval) |
| `customFieldOptionIds` | [String!] | Nej | Array av alternativ-ID:n (använder första elementet för enkeltval) |

## Fråga Enkeltvals Värden

Fråga en posts enkeltvals värde:

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

Fältet `value` returnerar ett JSON-objekt med detaljer om det valda alternativet.

## Skapa Poster med Enkeltvals Värden

När du skapar en ny post med enkeltvals värden:

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

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `value` | JSON | Innehåller det valda alternativobjektet med id, titel, färg, position |
| `todo` | Todo! | Den post detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast modifierades |

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
| `name` | String! | Visningsnamn för enkeltvalsfältet |
| `type` | CustomFieldType! | Alltid `SELECT_SINGLE` |
| `description` | String | Hjälptext för fältet |
| `customFieldOptions` | [CustomFieldOption!] | Alla tillgängliga alternativ |

## Värdeformat

### Inmatningsformat
- **API Parameter**: Använd `customFieldOptionId` för enkelt alternativ-ID
- **Alternativ**: Använd `customFieldOptionIds` array (tar första elementet)
- **Rensa Val**: Utelämna båda fälten eller skicka tomma värden

### Utdataformat
- **GraphQL Svar**: JSON-objekt i `value` fältet som innehåller {id, titel, färg, position}
- **Aktivitetslogg**: Alternativtitel som sträng
- **Automationsdata**: Alternativtitel som sträng

## Urvalsbeteende

### Exklusivt Urval
- Att ställa in ett nytt alternativ tar automatiskt bort det tidigare valet
- Endast ett alternativ kan väljas åt gången
- Att ställa in `null` eller tomt värde rensar valet

### Återfall Logik
- Om `customFieldOptionIds` array tillhandahålls, används endast det första alternativet
- Detta säkerställer kompatibilitet med multi-val inmatningsformat
- Tomma arrayer eller null-värden rensar valet

## Hantera Alternativ

### Uppdatera Alternativsegenskaper
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

### Ta Bort Alternativ
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**Notera**: Att ta bort ett alternativ kommer att rensa det från alla poster där det valdes.

### Omordna Alternativ
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

## Valideringsregler

### Alternativ Validering
- Det angivna alternativ-ID:t måste existera
- Alternativet måste tillhöra det angivna anpassade fältet
- Endast ett alternativ kan väljas (tillämpas automatiskt)
- Null/tomma värden är giltiga (ingen val)

### Fält Validering
- Måste ha minst ett alternativ definierat för att vara användbart
- Alternativstitlar måste vara unika inom fältet
- Färgkoder måste vara i giltigt hex-format (om tillhandahålls)

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|-------------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Fel Svar

### Ogiltigt Alternativ-ID
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

### Alternativ Tillhör Inte Fältet
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Kan Inte Tolka Värde
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

## Bästa Praxis

### Alternativdesign
- Använd tydliga, beskrivande alternativtitlar
- Tillämpa meningsfull färgkodning
- Håll alternativlistor fokuserade och relevanta
- Ordna alternativ logiskt (efter prioritet, frekvens, etc.)

### Statusfältmönster
- Använd konsekventa statusarbetsflöden över projekt
- Tänk på den naturliga progressionen av alternativ
- Inkludera tydliga slutliga tillstånd (Färdig, Avbruten, etc.)
- Använd färger som återspeglar alternativets betydelse

### Databehandling
- Granska och rensa oanvända alternativ periodiskt
- Använd konsekventa namngivningskonventioner
- Tänk på påverkan av alternativborttagning på befintliga poster
- Planera för alternativuppdateringar och migreringar

## Vanliga Användningsfall

1. **Status och Arbetsflöde**
   - Uppgiftsstatus (Att Göra, Pågående, Färdig)
   - Godkännandestatus (Väntande, Godkänd, Avvisad)
   - Projektfas (Planering, Utveckling, Testning, Släppt)
   - Problemlösningsstatus

2. **Klassificering och Kategorisering**
   - Prioritetsnivåer (Låg, Medel, Hög, Kritisk)
   - Uppgiftstyper (Bug, Funktion, Förbättring, Dokumentation)
   - Projektkategorier (Intern, Klient, Forskning)
   - Avdelningsuppdrag

3. **Kvalitet och Bedömning**
   - Granskningsstatus (Inte Påbörjad, Under Granskning, Godkänd)
   - Kvalitetsbetyg (Dålig, Acceptabel, Bra, Utmärkt)
   - Risknivåer (Låg, Medel, Hög)
   - Förtroendenivåer

4. **Tilldelning och Ägarskap**
   - Teamuppdrag
   - Avdelningse ägarskap
   - Rollbaserade uppdrag
   - Regionala uppdrag

## Integrationsfunktioner

### Med Automatiseringar
- Utlösa åtgärder när specifika alternativ väljs
- Styra arbete baserat på valda kategorier
- Skicka meddelanden för statusändringar
- Skapa villkorliga arbetsflöden baserat på val

### Med Uppslag
- Filtrera poster efter valda alternativ
- Referera alternativdata från andra poster
- Skapa rapporter baserat på alternativval
- Gruppera poster efter valda värden

### Med Formulär
- Dropdown-inmatningskontroller
- Radioknappgränssnitt
- Alternativvalidering och filtrering
- Villkorlig fältvisning baserat på val

## Aktivitetsövervakning

Ändringar i enkeltvalsfält spåras automatiskt:
- Visar gamla och nya alternativval
- Visar alternativtitlar i aktivitetsloggen
- Tidsstämplar för alla valändringar
- Användarattribution för modifieringar

## Skillnader från Multi-Select

| Funktion | Enkeltval | Multi-Select |
|---------|-----------|--------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Begränsningar

- Endast ett alternativ kan väljas åt gången
- Ingen hierarkisk eller nästlad alternativstruktur
- Alternativ delas över alla poster som använder fältet
- Ingen inbyggd alternativanalys eller användningsspårning
- Färgkoder är endast för visning, ingen funktionell påverkan
- Kan inte ställa in olika behörigheter per alternativ

## Relaterade Resurser

- [Multi-Select Fält](/api/custom-fields/select-multi) - För flervalsval
- [Checkbox Fält](/api/custom-fields/checkbox) - För enkla booleanval
- [Textfält](/api/custom-fields/text-single) - För fri textinmatning
- [Översikt över Anpassade Fält](/api/custom-fields/1.index) - Allmänna koncept