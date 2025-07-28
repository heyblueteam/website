---
title: Betygsättningsanpassat fält
description: Skapa betygsfält för att lagra numeriska betyg med konfigurerbara skalor och validering
---

Betygsanpassade fält gör att du kan lagra numeriska betyg i poster med konfigurerbara minimi- och maximivärden. De är idealiska för prestationsbetyg, nöjdhetsbetyg, prioriteringsnivåer eller vilken som helst databaserad på numeriska skalor i dina projekt.

## Grundläggande exempel

Skapa ett enkelt betygsfält med standard 0-5 skala:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Avancerat exempel

Skapa ett betygsfält med anpassad skala och beskrivning:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Indata parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för betygsfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `RATING` |
| `projectId` | String! | ✅ Ja | Projekt-ID där detta fält kommer att skapas |
| `description` | String | Nej | Hjälptext som visas för användare |
| `min` | Float | Nej | Minimi betygsvärde (ingen standard) |
| `max` | Float | Nej | Maximalt betygsvärde |

## Ställa in betygsvärden

För att ställa in eller uppdatera ett betygsvärde på en post:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade betygsfältet |
| `value` | String! | ✅ Ja | Betygsvärde som sträng (inom det konfigurerade intervallet) |

## Skapa poster med betygsvärden

När du skapar en ny post med betygsvärden:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
      }
      value
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
| `value` | Float | Det lagrade betygsvärdet (åtkomligt via customField.value) |
| `todo` | Todo! | Den post detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

**Obs**: Betygsvärdet nås faktiskt via `customField.value.number` i frågor.

### CustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältet |
| `name` | String! | Visningsnamn för betygsfältet |
| `type` | CustomFieldType! | Alltid `RATING` |
| `min` | Float | Minimi tillåtna betygsvärde |
| `max` | Float | Maximalt tillåtet betygsvärde |
| `description` | String | Hjälptext för fältet |

## Betygsvalidering

### Värdebegränsningar
- Betygsvärden måste vara numeriska (Float-typ)
- Värden måste ligga inom det konfigurerade min/max-intervallet
- Om inget minimum anges finns det inget standardvärde
- Maximalt värde är valfritt men rekommenderas

### Valideringsregler
**Viktigt**: Validering sker endast när formulär skickas, inte när `setTodoCustomField` används direkt.

- Indata tolkas som ett flyttal (när formulär används)
- Måste vara större än eller lika med minimi värdet (när formulär används)
- Måste vara mindre än eller lika med maximivärdet (när formulär används)
- `setTodoCustomField` accepterar vilket strängvärde som helst utan validering

### Giltiga betygsexempel
För ett fält med min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Ogiltiga betygsexempel
För ett fält med min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Konfigurationsalternativ

### Betygsskala inställning
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Vanliga betygsskalor
- **1-5 Stjärnor**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 Prestanda**: `min: 1, max: 10`
- **0-100 Procent**: `min: 0, max: 100`
- **Anpassad skala**: Valfritt numeriskt intervall

## Obligatoriska behörigheter

Anpassade fältoperationer följer standard rollbaserade behörigheter:

| Åtgärd | Obligatorisk roll |
|--------|-------------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Obs**: De specifika roller som krävs beror på din projekts anpassade rollkonfiguration och fältbehörigheter.

## Felmeddelanden

### Valideringsfel (Endast formulär)
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

**Viktigt**: Validering av betygsvärden (min/max begränsningar) sker endast när formulär skickas, inte när `setTodoCustomField` används direkt.

### Anpassat fält hittades inte
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

## Bästa praxis

### Skala design
- Använd konsekventa betygsskalor över liknande fält
- Tänk på användarens bekantskap (1-5 stjärnor, 0-10 NPS)
- Sätt lämpliga minimi värden (0 vs 1)
- Definiera tydlig betydelse för varje betygsnivå

### Datakvalitet
- Validera betygsvärden innan lagring
- Använd decimalprecision på rätt sätt
- Tänk på avrundning för visningsändamål
- Ge tydlig vägledning om betygens betydelser

### Användarupplevelse
- Visa betygsskalor visuellt (stjärnor, progressionsfält)
- Visa aktuellt värde och skala begränsningar
- Ge kontext för betygens betydelser
- Tänk på standardvärden för nya poster

## Vanliga användningsfall

1. **Prestandahantering**
   - Anställdas prestationsbetyg
   - Projektkvalitetsbetyg
   - Uppgiftsfullföljande betyg
   - Bedömningar av färdighetsnivåer

2. **Kundfeedback**
   - Nöjdhetsbetyg
   - Produktkvalitetsbetyg
   - Tjänsteupplevelsebetyg
   - Net Promoter Score (NPS)

3. **Prioritet och betydelse**
   - Uppgiftens prioriteringsnivåer
   - Brådskande betyg
   - Riskbedömningsbetyg
   - Påverkningsbetyg

4. **Kvalitetssäkring**
   - Kodgranskningsbetyg
   - Testkvalitetsbetyg
   - Dokumentationskvalitet
   - Processöverensstämmelsebetyg

## Integrationsfunktioner

### Med automatiseringar
- Utlösa åtgärder baserat på betygströsklar
- Skicka meddelanden för låga betyg
- Skapa uppföljningsuppgifter för höga betyg
- Rutta arbete baserat på betygsvärden

### Med uppslagningar
- Beräkna genomsnittliga betyg över poster
- Hitta poster efter betygsområden
- Referera till betygsdata från andra poster
- Aggregat betygsstatistik

### Med Blue frontend
- Automatisk intervallvalidering i formulärsammanhang
- Visuella betygsinmatningskontroller
- Realtidsvalideringsfeedback
- Stjärn- eller glidkontroller för inmatning

## Aktivitetsspårning

Ändringar i betygsfält spåras automatiskt:
- Gamla och nya betygsvärden loggas
- Aktiviteten visar numeriska förändringar
- Tidsstämplar för alla betygsuppdateringar
- Användarattribution för ändringar

## Begränsningar

- Endast numeriska värden stöds
- Ingen inbyggd visuell betygsvisning (stjärnor, etc.)
- Decimalprecision beror på databasens konfiguration
- Ingen lagring av betygsmetadata (kommentarer, kontext)
- Ingen automatisk betygsaggregat eller statistik
- Ingen inbyggd betygskonvertering mellan skalor
- **Kritiskt**: Min/max validering fungerar endast i formulär, inte via `setTodoCustomField`

## Relaterade resurser

- [Nummerfält](/api/5.custom%20fields/number) - För allmänna numeriska data
- [Procentfält](/api/5.custom%20fields/percent) - För procentvärden
- [Väljfält](/api/5.custom%20fields/select-single) - För diskreta valbetyg
- [Översikt över anpassade fält](/api/5.custom%20fields/2.list-custom-fields) - Allmänna koncept