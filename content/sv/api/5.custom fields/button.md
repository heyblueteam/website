---
title: Anpassad knappfält
description: Skapa interaktiva knappfält som utlöser automatiseringar när de klickas
---

Anpassade knappfält tillhandahåller interaktiva UI-element som utlöser automatiseringar när de klickas. Till skillnad från andra typer av anpassade fält som lagrar data, fungerar knappfält som åtgärdsutlösare för att utföra konfigurerade arbetsflöden.

## Grundläggande exempel

Skapa ett enkelt knappfält som utlöser en automatisering:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Avancerat exempel

Skapa en knapp med bekräftelsekrav:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Indata parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för knappen |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `BUTTON` |
| `projectId` | String! | ✅ Ja | Projekt-ID där fältet kommer att skapas |
| `buttonType` | String | Nej | Bekräftelsebeteende (se knapp typer nedan) |
| `buttonConfirmText` | String | Nej | Text som användare måste skriva för hård bekräftelse |
| `description` | String | Nej | Hjälptext som visas för användare |
| `required` | Boolean | Nej | Om fältet är obligatoriskt (standard är falskt) |
| `isActive` | Boolean | Nej | Om fältet är aktivt (standard är sant) |

### Knapptypfält

Fältet `buttonType` är en fri textsträng som kan användas av UI-klienter för att bestämma bekräftelsebeteende. Vanliga värden inkluderar:

- `""` (tom) - Ingen bekräftelse
- `"soft"` - Enkel bekräftelsedialog
- `"hard"` - Kräver att bekräftelsetext skrivs

**Obs**: Dessa är endast UI-hänvisningar. API:et validerar eller tvingar inte specifika värden.

## Utlösning av knappklick

För att utlösa ett knappklick och utföra associerade automatiseringar:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Klicka Indata parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för uppgiften som innehåller knappen |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade knappfältet |

### Viktigt: API-beteende

**Alla knappklick genom API:et utförs omedelbart** oavsett några `buttonType` eller `buttonConfirmText` inställningar. Dessa fält lagras för UI-klienter att implementera bekräftelsedialoger, men API:et självt:

- Validerar inte bekräftelsetext
- Tvingar inte några bekräftelsekrav
- Utför knappåtgärden omedelbart när den anropas

Bekräftelse är enbart en klientsidefunktion för UI-säkerhet.

### Exempel: Klicka på olika knapptyper

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Alla tre mutationer ovan kommer att utföra knappåtgärden omedelbart när de anropas genom API:et, vilket kringgår eventuella bekräftelsekrav.

## Svarsfält

### Svar för anpassat fält

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för det anpassade fältet |
| `name` | String! | Visningsnamn för knappen |
| `type` | CustomFieldType! | Alltid `BUTTON` för knappfält |
| `buttonType` | String | Inställning för bekräftelsebeteende |
| `buttonConfirmText` | String | Obligatorisk bekräftelsetext (om hård bekräftelse används) |
| `description` | String | Hjälptext för användare |
| `required` | Boolean! | Om fältet är obligatoriskt |
| `isActive` | Boolean! | Om fältet för närvarande är aktivt |
| `projectId` | String! | ID för projektet som detta fält tillhör |
| `createdAt` | DateTime! | När fältet skapades |
| `updatedAt` | DateTime! | När fältet senast ändrades |

## Hur knappfält fungerar

### Automatiseringsintegration

Knappfält är utformade för att fungera med Blues automatiseringssystem:

1. **Skapa knappfältet** med mutation ovan
2. **Konfigurera automatiseringar** som lyssnar efter `CUSTOM_FIELD_BUTTON_CLICKED` händelser
3. **Användare klickar på knappen** i UI
4. **Automatiseringar utför** de konfigurerade åtgärderna

### Händelseflöde

När en knapp klickas:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Ingen datalagring

Viktigt: Knappfält lagrar inga värdedata. De fungerar enbart som åtgärdsutlösare. Varje klick:
- Genererar en händelse
- Utlöser associerade automatiseringar
- Registrerar en åtgärd i uppgiftshistoriken
- Modifierar inte något fältvärde

## Obligatoriska behörigheter

Användare behöver lämpliga projektroller för att skapa och använda knappfält:

| Åtgärd | Obligatorisk roll |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Felmeddelanden

### Behörighet nekad
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Anpassat fält hittades inte
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

**Obs**: API:et returnerar inte specifika fel för saknade automatiseringar eller bekräftelseavvikelser.

## Bästa praxis

### Namngivningskonventioner
- Använd åtgärdsorienterade namn: "Skicka faktura", "Skapa rapport", "Meddela teamet"
- Var specifik om vad knappen gör
- Undvik generiska namn som "Knapp 1" eller "Klicka här"

### Bekräftelseinställningar
- Lämna `buttonType` tomt för säkra, reversibla åtgärder
- Ställ in `buttonType` för att föreslå bekräftelsebeteende till UI-klienter
- Använd `buttonConfirmText` för att specificera vad användare ska skriva i UI-bekräftelser
- Kom ihåg: Dessa är endast UI-hänvisningar - API-anrop utförs alltid omedelbart

### Automatiseringsdesign
- Håll knappåtgärder fokuserade på ett enda arbetsflöde
- Ge tydlig feedback om vad som hände efter klickning
- Överväg att lägga till beskrivningstext för att förklara knappens syfte

## Vanliga användningsfall

1. **Arbetsflödesövergångar**
   - "Markera som slutförd"
   - "Skicka för godkännande"
   - "Arkivera uppgift"

2. **Externa integrationer**
   - "Synkronisera till CRM"
   - "Generera faktura"
   - "Skicka e-postuppdatering"

3. **Batchoperationer**
   - "Uppdatera alla deluppgifter"
   - "Kopiera till projekt"
   - "Tillämpa mall"

4. **Rapporteringsåtgärder**
   - "Generera rapport"
   - "Exportera data"
   - "Skapa sammanfattning"

## Begränsningar

- Knappar kan inte lagra eller visa datavärden
- Varje knapp kan endast utlösa automatiseringar, inte direkta API-anrop (dock kan automatiseringar inkludera HTTP-begäran åtgärder för att anropa externa API:er eller Blues egna API:er)
- Knappens synlighet kan inte kontrolleras villkorsbundet
- Max en automatiseringsexekvering per klick (även om den automatiseringen kan utlösa flera åtgärder)

## Relaterade resurser

- [Automations API](/api/automations/index) - Konfigurera åtgärder som utlöses av knappar
- [Översikt över anpassade fält](/custom-fields/list-custom-fields) - Allmänna koncept för anpassade fält