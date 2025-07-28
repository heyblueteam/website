---
title: Telefonanpassat fält
description: Skapa telefonfält för att lagra och validera telefonnummer med internationell formatering
---

Telefonanpassade fält gör att du kan lagra telefonnummer i poster med inbyggd validering och internationell formatering. De är idealiska för att spåra kontaktinformation, nödkontakter eller annan telefonrelaterad data i dina projekt.

## Grundläggande exempel

Skapa ett enkelt telefonfält:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Avancerat exempel

Skapa ett telefonfält med beskrivning:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ Ja | Visningsnamn för telefonfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `PHONE` |
| `description` | String | Nej | Hjälptext som visas för användare |

**Notera**: Anpassade fält är automatiskt kopplade till projektet baserat på användarens aktuella projektkontext. Ingen `projectId` parameter krävs.

## Ställa in telefonvärden

För att ställa in eller uppdatera ett telefonvärde på en post:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det telefonanpassade fältet |
| `text` | String | Nej | Telefonnummer med landskod |
| `regionCode` | String | Nej | Landskod (automatiskt upptäckt) |

**Notera**: Medan `text` är valfri i schemat, krävs ett telefonnummer för att fältet ska vara meningsfullt. När `setTodoCustomField` används, utförs ingen validering - du kan lagra vilket textvärde som helst och regionkod. Den automatiska upptäckten sker endast under postskapande.

## Skapa poster med telefonvärden

När du skapar en ny post med telefonvärden:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
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
| `text` | String | Det formaterade telefonnumret (internationellt format) |
| `regionCode` | String | Landskoden (t.ex. "US", "GB", "CA") |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

## Validering av telefonnummer

**Viktigt**: Validering och formatering av telefonnummer sker endast när nya poster skapas via `createTodo`. När befintliga telefonvärden uppdateras med `setTodoCustomField`, utförs ingen validering och värdena lagras som angivna.

### Godkända format (under postskapande)
Telefonnummer måste inkludera en landskod i ett av dessa format:

- **E.164-format (föredraget)**: `+12345678900`
- **Internationellt format**: `+1 234 567 8900`
- **Internationellt med interpunktion**: `+1 (234) 567-8900`
- **Landskod med bindestreck**: `+1-234-567-8900`

**Notera**: Nationella format utan landskod (som `(234) 567-8900`) kommer att avvisas under postskapande.

### Valideringsregler (under postskapande)
- Använder libphonenumber-js för analys och validering
- Accepterar olika internationella telefonnummerformat
- Upptäcker automatiskt landet från numret
- Formaterar nummer i internationellt visningsformat (t.ex. `+1 234 567 8900`)
- Extraherar och lagrar landskod separat (t.ex. `US`)

### Giltiga telefonexempel
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Ogiltiga telefonexempel
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Lagringsformat

När du skapar poster med telefonnummer:
- **text**: Lagrade i internationellt format (t.ex. `+1 234 567 8900`) efter validering
- **regionCode**: Lagrade som ISO landskod (t.ex. `US`, `GB`, `CA`) automatiskt upptäckta

När du uppdaterar via `setTodoCustomField`:
- **text**: Lagrade exakt som angivet (ingen formatering)
- **regionCode**: Lagrade exakt som angivet (ingen validering)

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk behörighet |
|--------|------------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Felrespons

### Ogiltigt telefonformat
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
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

### Saknad landskod
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Bästa praxis

### Datainmatning
- Inkludera alltid landskod i telefonnummer
- Använd E.164-format för konsekvens
- Validera nummer innan lagring för viktiga operationer
- Tänk på regionala preferenser för visningsformatering

### Datakvalitet
- Lagra nummer i internationellt format för global kompatibilitet
- Använd regionCode för landspecifika funktioner
- Validera telefonnummer innan kritiska operationer (SMS, samtal)
- Tänk på tidszonsimplikationer för kontaktens timing

### Internationella överväganden
- Landskod upptäckts automatiskt och lagras
- Nummer formateras i internationell standard
- Regionala visningspreferenser kan använda regionCode
- Tänk på lokala ringsystem när du visar

## Vanliga användningsfall

1. **Kontaktledning**
   - Klienttelefonnummer
   - Leverantörskontaktinformation
   - Teammedlemmars telefonnummer
   - Supportkontaktuppgifter

2. **Nödkontakter**
   - Nödkontaktnummer
   - Kontaktinformation för jour
   - Kontakter för krisrespons
   - Eskaleringstelefonnummer

3. **Kundsupport**
   - Kundtelefonnummer
   - Supportåteruppringningsnummer
   - Verifieringstelefonnummer
   - Uppföljningskontakt nummer

4. **Försäljning och marknadsföring**
   - Leadtelefonnummer
   - Kampanjkontaktlistor
   - Partnerkontaktinformation
   - Remisskällor telefoner

## Integrationsfunktioner

### Med automatiseringar
- Utlös åtgärder när telefonfält uppdateras
- Skicka SMS-notifikationer till lagrade telefonnummer
- Skapa uppföljningsuppgifter baserat på telefonändringar
- Rutta samtal baserat på telefonnumrets data

### Med uppslag
- Referera telefondata från andra poster
- Sammanställ telefonlistor från flera källor
- Hitta poster efter telefonnummer
- Korskontrollera kontaktinformation

### Med formulär
- Automatisk telefonvalidering
- Kontroll av internationellt format
- Upptäck landskod
- Realtidsformatåterkoppling

## Begränsningar

- Kräver landskod för alla nummer
- Inga inbyggda SMS- eller samtalsfunktioner
- Ingen telefonnummerverifiering utöver formatkontroll
- Ingen lagring av telefonmetadata (operatör, typ, etc.)
- Nationella formatnummer utan landskod avvisas
- Ingen automatisk telefonnummerformatering i UI utöver internationell standard

## Relaterade resurser

- [Textfält](/api/custom-fields/text-single) - För icke-telefontextdata
- [E-postfält](/api/custom-fields/email) - För e-postadresser
- [URL-fält](/api/custom-fields/url) - För webbplatsadresser
- [Översikt över anpassade fält](/custom-fields/list-custom-fields) - Allmänna koncept