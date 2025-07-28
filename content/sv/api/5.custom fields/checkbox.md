---
title: Checkbox Anpassad Fält
description: Skapa booleska checkbox-fält för ja/nej eller sant/falskt data
---

Checkbox-anpassade fält ger en enkel boolesk (sant/falskt) inmatning för uppgifter. De är perfekta för binära val, statusindikatorer eller för att spåra om något har slutförts.

## Grundläggande Exempel

Skapa ett enkelt checkbox-fält:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Avancerat Exempel

Skapa ett checkbox-fält med beskrivning och validering:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ Ja | Visningsnamn för checkboxen |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `CHECKBOX` |
| `description` | String | Nej | Hjälptext som visas för användare |

## Ställa in Checkbox-värden

För att ställa in eller uppdatera ett checkbox-värde på en uppgift:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

För att avmarkera en checkbox:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för uppgiften som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för checkbox-anpassat fält |
| `checked` | Boolean | Nej | Sant för att markera, falskt för att avmarkera |

## Skapa Uppgifter med Checkbox-värden

När du skapar en ny uppgift med checkbox-värden:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Godkända Strängvärden

När du skapar uppgifter måste checkbox-värden skickas som strängar:

| Strängvärde | Resultat |
|--------------|---------|
| `"true"` | ✅ Markerad (skiftlägeskänslig) |
| `"1"` | ✅ Markerad |
| `"checked"` | ✅ Markerad (skiftlägeskänslig) |
| Any other value | ❌ Avmarkerad |

**Observera**: Strängjämförelser under uppgiftsskapande är skiftlägeskänsliga. Värdena måste exakt matcha `"true"`, `"1"`, eller `"checked"` för att resultera i ett markerat tillstånd.

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | ID! | Unik identifierare för fältvärdet |
| `uid` | String! | Alternativ unik identifierare |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `checked` | Boolean | Checkboxens tillstånd (sant/falskt/null) |
| `todo` | Todo! | Den uppgift som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

## Automatiseringsintegration

Checkbox-fält utlöser olika automatiseringsevenemang baserat på tillståndsändringar:

| Åtgärd | Händelse Utlöst | Beskrivning |
|--------|----------------|-------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Utlöst när checkboxen är markerad |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Utlöst när checkboxen är avmarkerad |

Detta gör att du kan skapa automatiseringar som svarar på förändringar i checkboxens tillstånd, såsom:
- Skicka meddelanden när objekt godkänns
- Flytta uppgifter när granskningscheckboxar är markerade
- Uppdatera relaterade fält baserat på checkboxens tillstånd

## Dataimport/Export

### Importera Checkbox-värden

Vid import av data via CSV eller andra format:
- `"true"`, `"yes"` → Markerad (skiftlägesokänslig)
- Något annat värde (inklusive `"false"`, `"no"`, `"0"`, tom) → Avmarkerad

### Exportera Checkbox-värden

Vid export av data:
- Markerade rutor exporteras som `"X"`
- Avmarkerade rutor exporteras som tom sträng `""`

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|-------------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Felrespons

### Ogiltig Värdetyp
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Bästa Praxis

### Namngivningskonventioner
- Använd tydliga, handlingsorienterade namn: "Godkänd", "Granskad", "Är Komplett"
- Undvik negativa namn som förvirrar användare: föredra "Är Aktiv" framför "Är Inaktiv"
- Var specifik om vad checkboxen representerar

### När man ska använda Checkboxar
- **Binära val**: Ja/Nej, Sant/Falskt, Gjort/Inte Gjort
- **Statusindikatorer**: Godkänd, Granskad, Publicerad
- **Funktioner**: Har Prioriterad Support, Kräver Signatur
- **Enkel spårning**: E-post Skickad, Faktura Betald, Objekt Skickat

### När man INTE ska använda Checkboxar
- När du behöver mer än två alternativ (använd SELECT_SINGLE istället)
- För numeriska eller textdata (använd NUMBER eller TEXT-fält)
- När du behöver spåra vem som markerade det eller när (använd revisionsloggar)

## Vanliga Användningsfall

1. **Godkännandearbetsflöden**
   - "Chef Godkänd"
   - "Klient Godkännande"
   - "Juridisk Granskning Komplett"

2. **Uppgiftshantering**
   - "Är Blockerad"
   - "Redo för Granskning"
   - "Hög Prioritet"

3. **Kvalitetskontroll**
   - "QA Godkänd"
   - "Dokumentation Komplett"
   - "Tester Skriven"

4. **Administrativa Flaggar**
   - "Faktura Skickad"
   - "Kontrakt Undertecknat"
   - "Uppföljning Krävs"

## Begränsningar

- Checkbox-fält kan endast lagra sant/falskt värden (ingen tri-state eller null efter första inställning)
- Ingen standardvärdeskonfiguration (börjar alltid som null tills den ställs in)
- Kan inte lagra ytterligare metadata som vem som markerade det eller när
- Ingen villkorlig synlighet baserat på andra fältvärden

## Relaterade Resurser

- [Översikt över Anpassade Fält](/api/custom-fields/list-custom-fields) - Allmänna koncept för anpassade fält
- [Automations API](/api/automations) - Skapa automatiseringar som utlöses av checkboxändringar