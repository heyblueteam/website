---
title: Tidsduration Anpassat Fält
description: Skapa beräknade tidsdurationfält som spårar tiden mellan händelser i ditt arbetsflöde
---

Tidsduration anpassade fält beräknar och visar automatiskt varaktigheten mellan två händelser i ditt arbetsflöde. De är idealiska för att spåra bearbetningstider, svarstider, cykeltider eller andra tidsbaserade mätvärden i dina projekt.

## Grundläggande Exempel

Skapa ett enkelt tidsdurationfält som spårar hur lång tid uppgifter tar att slutföra:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Avancerat Exempel

Skapa ett komplext tidsdurationfält som spårar tiden mellan ändringar av anpassade fält med ett SLA-mål:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Inmatningsparametrar

### CreateCustomFieldInput (TIME_DURATION)

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för tidsdurationfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `TIME_DURATION` |
| `description` | String | Nej | Hjälptext som visas för användare |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Ja | Hur varaktigheten ska visas |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Ja | Konfiguration för start-händelse |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Ja | Konfiguration för slut-händelse |
| `timeDurationTargetTime` | Float | Nej | Målvaraktighet i sekunder för SLA-övervakning |

### CustomFieldTimeDurationInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ Ja | Typ av händelse att spåra |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Ja | `FIRST` eller `LAST` förekomst |
| `customFieldId` | String | Conditional | Obligatorisk för `TODO_CUSTOM_FIELD` typ |
| `customFieldOptionIds` | [String!] | Conditional | Obligatorisk för val av fältändringar |
| `todoListId` | String | Conditional | Obligatorisk för `TODO_MOVED` typ |
| `tagId` | String | Conditional | Obligatorisk för `TODO_TAG_ADDED` typ |
| `assigneeId` | String | Conditional | Obligatorisk för `TODO_ASSIGNEE_ADDED` typ |

### CustomFieldTimeDurationType Värden

| Värde | Beskrivning |
|-------|-------------|
| `TODO_CREATED_AT` | När posten skapades |
| `TODO_CUSTOM_FIELD` | När ett anpassat fältvärde ändrades |
| `TODO_DUE_DATE` | När förfallodatumet ställdes in |
| `TODO_MARKED_AS_COMPLETE` | När posten markerades som slutförd |
| `TODO_MOVED` | När posten flyttades till en annan lista |
| `TODO_TAG_ADDED` | När en tagg lades till posten |
| `TODO_ASSIGNEE_ADDED` | När en tilldelad person lades till posten |

### CustomFieldTimeDurationCondition Värden

| Värde | Beskrivning |
|-------|-------------|
| `FIRST` | Använd den första förekomsten av händelsen |
| `LAST` | Använd den senaste förekomsten av händelsen |

### CustomFieldTimeDurationDisplayType Värden

| Värde | Beskrivning | Exempel |
|-------|-------------|---------|
| `FULL_DATE` | Dagar:Timmar:Minuter:Sekunder format | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Skrivet ut i fullständiga ord | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numeriskt med enheter | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Varaktighet i dagar endast | `"2.5"` (2.5 days) |
| `HOURS` | Varaktighet i timmar endast | `"60"` (60 hours) |
| `MINUTES` | Varaktighet i minuter endast | `"3600"` (3600 minutes) |
| `SECONDS` | Varaktighet i sekunder endast | `"216000"` (216000 seconds) |

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `number` | Float | Varaktigheten i sekunder |
| `value` | Float | Alias för nummer (varaktighet i sekunder) |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast uppdaterades |

### CustomField Svar (TIME_DURATION)

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Visningsformat för varaktigheten |
| `timeDurationStart` | CustomFieldTimeDuration | Konfiguration för start-händelse |
| `timeDurationEnd` | CustomFieldTimeDuration | Konfiguration för slut-händelse |
| `timeDurationTargetTime` | Float | Målvaraktighet i sekunder (för SLA-övervakning) |

## Varaktighetsberäkning

### Hur Det Fungerar
1. **Start-Händelse**: Systemet övervakar den angivna start-händelsen
2. **Slut-Händelse**: Systemet övervakar den angivna slut-händelsen
3. **Beräkning**: Varaktighet = Sluttid - Starttid
4. **Lagring**: Varaktighet lagras i sekunder som ett nummer
5. **Visning**: Formateras enligt `timeDurationDisplay` inställningen

### Uppdateringsutlösare
Varaktighetsvärden beräknas automatiskt om:
- Poster skapas eller uppdateras
- Värden för anpassade fält ändras
- Taggar läggs till eller tas bort
- Tilldelade personer läggs till eller tas bort
- Poster flyttas mellan listor
- Poster markeras som slutförda/inte slutförda

## Läsa Varaktighetsvärden

### Fråga Varaktighetsfält
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Formaterade Visningsvärden
Varaktighetsvärden formateras automatiskt baserat på `timeDurationDisplay` inställningen:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Vanliga Konfigurations Exempel

### Uppgiftens Slutförande Tid
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Statusändrings Varaktighet
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Tid i Specifik Lista
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Tilldelnings Svarstid
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|-------------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Fel Svar

### Ogiltig Konfiguration
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Refererad Fält Hittades Inte
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

### Saknade Obligatoriska Alternativ
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Viktiga Anteckningar

### Automatisk Beräkning
- Varaktighetsfält är **skrivskyddade** - värden beräknas automatiskt
- Du kan inte manuellt ställa in varaktighetsvärden via API
- Beräkningar sker asynkront via bakgrundsjobb
- Värden uppdateras automatiskt när utlösande händelser inträffar

### Prestanda Överväganden
- Varaktighetsberäkningar köas och bearbetas asynkront
- Stora mängder varaktighetsfält kan påverka prestandan
- Tänk på frekvensen av utlösande händelser när du designar varaktighetsfält
- Använd specifika villkor för att undvika onödiga omberäkningar

### Null Värden
Varaktighetsfält kommer att visa `null` när:
- Start-händelsen ännu inte har inträffat
- Slut-händelsen ännu inte har inträffat
- Konfigurationen refererar till icke-existerande enheter
- Beräkningen stöter på ett fel

## Bästa Praxis

### Konfigurationsdesign
- Använd specifika händelsetyper snarare än generiska när det är möjligt
- Välj lämpliga `FIRST` vs `LAST` villkor baserat på ditt arbetsflöde
- Testa varaktighetsberäkningar med exempeldata innan distribution
- Dokumentera din logik för varaktighetsfält för teammedlemmar

### Visningsformat
- Använd `FULL_DATE_SUBSTRING` för mest läsbart format
- Använd `FULL_DATE` för kompakt, konsekvent breddvisning
- Använd `FULL_DATE_STRING` för formella rapporter och dokument
- Använd `DAYS`, `HOURS`, `MINUTES`, eller `SECONDS` för enkla numeriska visningar
- Tänk på dina UI-ytbegränsningar när du väljer format

### SLA Övervakning med Mål Tids
När du använder `timeDurationTargetTime`:
- Ställ in målvaraktigheten i sekunder
- Jämför faktisk varaktighet mot mål för SLA-efterlevnad
- Använd i instrumentpaneler för att lyfta fram försenade objekt
- Exempel: 24-timmars svar SLA = 86400 sekunder

### Arbetsflödesintegration
- Designa varaktighetsfält för att matcha dina faktiska affärsprocesser
- Använd varaktighetsdata för processförbättring och optimering
- Övervaka varaktighetstrender för att identifiera flaskhalsar i arbetsflödet
- Ställ in varningar för varaktighetsgränser om det behövs

## Vanliga Användningsfall

1. **Processprestanda**
   - Tider för uppgiftsfullföljande
   - Granskning cykeltider
   - Godkännande bearbetningstider
   - Svarstider

2. **SLA Övervakning**
   - Tid till första svar
   - Lösningstider
   - Eskaleringstidsramar
   - Efterlevnad av servicenivå

3. **Arbetsflödesanalys**
   - Identifiering av flaskhalsar
   - Processoptimering
   - Teamprestandamätningar
   - Kvalitetssäkrings tidsmätning

4. **Projektledning**
   - Fasadurationer
   - Tidsramar för milstolpar
   - Resursallokeringstid
   - Leveranstidsramar

## Begränsningar

- Varaktighetsfält är **skrivskyddade** och kan inte ställas in manuellt
- Värden beräknas asynkront och kan vara otillgängliga omedelbart
- Kräver att korrekta händelseutlösare är inställda i ditt arbetsflöde
- Kan inte beräkna varaktigheter för händelser som inte har inträffat
- Begränsat till att spåra tid mellan diskreta händelser (inte kontinuerlig tidsövervakning)
- Inga inbyggda SLA-varningar eller meddelanden
- Kan inte aggregera flera varaktighetsberäkningar till ett enda fält

## Relaterade Resurser

- [Nummerfält](/api/custom-fields/number) - För manuella numeriska värden
- [Datumfält](/api/custom-fields/date) - För specifik datumspårning
- [Översikt över Anpassade Fält](/api/custom-fields/list-custom-fields) - Allmänna koncept
- [Automatiseringar](/api/automations) - För att utlösa åtgärder baserat på varaktighetsgränser