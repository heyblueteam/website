---
title: E-post Anpassad Fält
description: Skapa e-postfält för att lagra och validera e-postadresser
---

E-post anpassade fält gör att du kan lagra e-postadresser i poster med inbyggd validering. De är idealiska för att spåra kontaktinformation, tilldelade e-postadresser eller annan e-postrelaterad data i dina projekt.

## Grundläggande Exempel

Skapa ett enkelt e-postfält:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Avancerat Exempel

Skapa ett e-postfält med beskrivning:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ Ja | Visningsnamn för e-postfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `EMAIL` |
| `description` | String | Nej | Hjälptext som visas för användare |

## Ställa in E-postvärden

För att ställa in eller uppdatera ett e-postvärde på en post:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### SetTodoCustomFieldInput Parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade e-postfältet |
| `text` | String | Nej | E-postadress att lagra |

## Skapa Poster med E-postvärden

När du skapar en ny post med e-postvärden:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Svarsfält

### CustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | ID! | Unik identifierare för det anpassade fältet |
| `name` | String! | Visningsnamn för e-postfältet |
| `type` | CustomFieldType! | Fälttypen (EMAIL) |
| `description` | String | Hjälptext för fältet |
| `value` | JSON | Innehåller e-postvärdet (se nedan) |
| `createdAt` | DateTime! | När fältet skapades |
| `updatedAt` | DateTime! | När fältet senast ändrades |

**Viktigt**: E-postvärden nås genom `customField.value.text` fältet, inte direkt på svaret.

## Fråga E-postvärden

När du frågar poster med e-post anpassade fält, nå e-posten genom `customField.value.text` vägen:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

Svaret kommer att inkludera e-posten i den nästlade strukturen:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## E-postvalidering

### Formulärvalidering
När e-postfält används i formulär, validerar de automatiskt e-postformatet:
- Använder standard e-postvalideringsregler
- Tar bort mellanslag från inmatning
- Avvisar ogiltiga e-postformat

### Valideringsregler
- Måste innehålla en `@` symbol
- Måste ha ett giltigt domänformat
- Ledande/efterföljande mellanslag tas automatiskt bort
- Vanliga e-postformat accepteras

### Giltiga E-postexempel
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Ogiltiga E-postexempel
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Viktiga Anteckningar

### Direkt API vs Formulär
- **Formulär**: Automatisk e-postvalidering tillämpas
- **Direkt API**: Ingen validering - vilken text som helst kan lagras
- **Rekommendation**: Använd formulär för användarinmatning för att säkerställa validering

### Lagringsformat
- E-postadresser lagras som ren text
- Ingen speciell formatering eller analys
- Skiftlägeskänslighet: EMAIL anpassade fält lagras skiftlägeskänsligt (till skillnad från användarautentiseringse-post som normaliseras till gemener)
- Inga maximala längdbegränsningar utöver databasbegränsningar (16MB gräns)

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|-----------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Felrespons

### Ogiltigt E-postformat (Endast Formulär)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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
      "code": "NOT_FOUND"
    }
  }]
}
```

## Bästa Praxis

### Datainmatning
- Validera alltid e-postadresser i din applikation
- Använd e-postfält endast för faktiska e-postadresser
- Överväg att använda formulär för användarinmatning för att få automatisk validering

### Datakvalitet
- Ta bort mellanslag innan lagring
- Överväg skiftlägesnormalisering (vanligtvis gemener)
- Validera e-postformat innan viktiga operationer

### Integritetsöverväganden
- E-postadresser lagras som ren text
- Överväg dataskyddsregler (GDPR, CCPA)
- Implementera lämpliga åtkomstkontroller

## Vanliga Användningsfall

1. **Kontaktledning**
   - Klienters e-postadresser
   - Leverantörers kontaktinformation
   - Teammedlemmars e-postadresser
   - Supportkontaktuppgifter

2. **Projektledning**
   - Intressenters e-postadresser
   - Godkännande kontakt-e-postadresser
   - Notifikationsmottagare
   - Externa samarbetspartners e-postadresser

3. **Kundsupport**
   - Kunders e-postadresser
   - Supportärende kontakter
   - Eskaleringskontakter
   - Feedback-e-postadresser

4. **Försäljning & Marknadsföring**
   - Lead-e-postadresser
   - Kampanjkontaktlistor
   - Partnerkontaktinformation
   - Remisskällor e-postadresser

## Integrationsfunktioner

### Med Automatiseringar
- Utlösa åtgärder när e-postfält uppdateras
- Skicka notifikationer till lagrade e-postadresser
- Skapa uppföljningsuppgifter baserat på e-poständringar

### Med Uppslag
- Referera e-postdata från andra poster
- Sammanställ e-postlistor från flera källor
- Hitta poster efter e-postadress

### Med Formulär
- Automatisk e-postvalidering
- E-postformatkontroll
- Ta bort mellanslag

## Begränsningar

- Ingen inbyggd e-postverifiering eller validering utöver formatkontroll
- Inga e-postspecifika UI-funktioner (som klickbara e-postlänkar)
- Lagrade som ren text utan kryptering
- Inga funktioner för att komponera eller skicka e-post
- Ingen lagring av e-postmetadata (visningsnamn, etc.)
- Direkta API-anrop kringgår validering (endast formulär validerar)

## Relaterade Resurser

- [Textfält](/api/custom-fields/text-single) - För icke-e-post textdata
- [URL-fält](/api/custom-fields/url) - För webbplatsadresser
- [Telefonfält](/api/custom-fields/phone) - För telefonnummer
- [Översikt över Anpassade Fält](/api/custom-fields/list-custom-fields) - Allmänna koncept