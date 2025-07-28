---
title: Tijdduur Aangepast Veld
description: Maak berekende tijdduurvelden die de tijd tussen gebeurtenissen in uw workflow bijhouden
---

Tijdduur aangepaste velden berekenen en tonen automatisch de duur tussen twee gebeurtenissen in uw workflow. Ze zijn ideaal voor het bijhouden van verwerkingstijden, responstijden, cyclustijden of andere tijdgebaseerde statistieken in uw projecten.

## Basisvoorbeeld

Maak een eenvoudig tijdduurveld dat bijhoudt hoe lang taken duren om te voltooien:

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

## Geavanceerd Voorbeeld

Maak een complex tijdduurveld dat de tijd bijhoudt tussen wijzigingen in aangepaste velden met een SLA-doel:

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

## Invoergegevens

### CreateCustomFieldInput (TIME_DURATION)

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het duurveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `TIME_DURATION` |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Ja | Hoe de duur moet worden weergegeven |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Ja | Start evenementconfiguratie |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Ja | Eind evenementconfiguratie |
| `timeDurationTargetTime` | Float | Nee | Doelduur in seconden voor SLA-monitoring |

### CustomFieldTimeDurationInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `type` | CustomFieldTimeDurationType! | ✅ Ja | Type gebeurtenis om bij te houden |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Ja | `FIRST` of `LAST` gebeurtenis |
| `customFieldId` | String | Conditional | Vereist voor `TODO_CUSTOM_FIELD` type |
| `customFieldOptionIds` | [String!] | Conditional | Vereist voor select veldwijzigingen |
| `todoListId` | String | Conditional | Vereist voor `TODO_MOVED` type |
| `tagId` | String | Conditional | Vereist voor `TODO_TAG_ADDED` type |
| `assigneeId` | String | Conditional | Vereist voor `TODO_ASSIGNEE_ADDED` type |

### CustomFieldTimeDurationType Waarden

| Waarde | Beschrijving |
|--------|--------------|
| `TODO_CREATED_AT` | Wanneer het record is aangemaakt |
| `TODO_CUSTOM_FIELD` | Wanneer een aangepaste veldwaarde is gewijzigd |
| `TODO_DUE_DATE` | Wanneer de vervaldatum is ingesteld |
| `TODO_MARKED_AS_COMPLETE` | Wanneer het record als voltooid is gemarkeerd |
| `TODO_MOVED` | Wanneer het record naar een andere lijst is verplaatst |
| `TODO_TAG_ADDED` | Wanneer een tag aan het record is toegevoegd |
| `TODO_ASSIGNEE_ADDED` | Wanneer een toegewezen persoon aan het record is toegevoegd |

### CustomFieldTimeDurationCondition Waarden

| Waarde | Beschrijving |
|--------|--------------|
| `FIRST` | Gebruik de eerste gebeurtenis |
| `LAST` | Gebruik de laatste gebeurtenis |

### CustomFieldTimeDurationDisplayType Waarden

| Waarde | Beschrijving | Voorbeeld |
|--------|--------------|-----------|
| `FULL_DATE` | Dagen:Uren:Minuten:Seconden formaat | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Volledig uitgeschreven | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numeriek met eenheden | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Duur in dagen alleen | `"2.5"` (2.5 days) |
| `HOURS` | Duur in uren alleen | `"60"` (60 hours) |
| `MINUTES` | Duur in minuten alleen | `"3600"` (3600 minutes) |
| `SECONDS` | Duur in seconden alleen | `"216000"` (216000 seconds) |

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `number` | Float | De duur in seconden |
| `value` | Float | Alias voor nummer (duur in seconden) |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is bijgewerkt |

### CustomField Antwoord (TIME_DURATION)

| Veld | Type | Beschrijving |
|------|------|--------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Weergaveformaat voor de duur |
| `timeDurationStart` | CustomFieldTimeDuration | Start evenementconfiguratie |
| `timeDurationEnd` | CustomFieldTimeDuration | Eind evenementconfiguratie |
| `timeDurationTargetTime` | Float | Doelduur in seconden (voor SLA-monitoring) |

## Duurberekening

### Hoe het Werkt
1. **Start Evenement**: Systeem monitort het opgegeven startevenement
2. **Eind Evenement**: Systeem monitort het opgegeven eindevenement
3. **Berekening**: Duur = Eindtijd - Starttijd
4. **Opslag**: Duur opgeslagen in seconden als een getal
5. **Weergave**: Geformatteerd volgens `timeDurationDisplay` instelling

### Update Triggers
Duurwaarden worden automatisch herberekend wanneer:
- Records worden aangemaakt of bijgewerkt
- Aangepaste veldwaarden veranderen
- Tags worden toegevoegd of verwijderd
- Toegewezen personen worden toegevoegd of verwijderd
- Records worden verplaatst tussen lijsten
- Records worden gemarkeerd als voltooid/onvoltooid

## Duurwaarden Lezen

### Vraag Duurvelden
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

### Geformatteerde Weergavewaarden
Duurwaarden worden automatisch geformatteerd op basis van de `timeDurationDisplay` instelling:

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

## Veelvoorkomende Configuratie Voorbeelden

### Taak Voltooiingstijd
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

### Statuswijzigingsduur
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

### Tijd in Specifieke Lijst
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

### Toewijzingsresponstijd
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

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Foutreacties

### Ongeldige Configuratie
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

### Verwezen Veld Niet Gevonden
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

### Ontbrekende Vereiste Opties
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

## Belangrijke Notities

### Automatische Berekening
- Duurvelden zijn **alleen-lezen** - waarden worden automatisch berekend
- U kunt geen duurwaarden handmatig instellen via de API
- Berekeningen gebeuren asynchroon via achtergrondtaken
- Waarden worden automatisch bijgewerkt wanneer triggergebeurtenissen plaatsvinden

### Prestatieoverwegingen
- Duurberekeningen worden in de wachtrij gezet en asynchroon verwerkt
- Grote aantallen duurvelden kunnen de prestaties beïnvloeden
- Houd rekening met de frequentie van triggergebeurtenissen bij het ontwerpen van duurvelden
- Gebruik specifieke voorwaarden om onnodige herberekeningen te voorkomen

### Null Waarden
Duurvelden tonen `null` wanneer:
- Het startevenement nog niet heeft plaatsgevonden
- Het eindevenement nog niet heeft plaatsgevonden
- Configuratie verwijst naar niet-bestaande entiteiten
- Berekening tegen een fout aanloopt

## Beste Praktijken

### Configuratieontwerp
- Gebruik specifieke gebeurtenistypen in plaats van algemene wanneer mogelijk
- Kies geschikte `FIRST` vs `LAST` voorwaarden op basis van uw workflow
- Test duurberekeningen met voorbeeldgegevens voordat u deze implementeert
- Documenteer uw logica voor het duurveld voor teamleden

### Weergaveformattering
- Gebruik `FULL_DATE_SUBSTRING` voor de meest leesbare indeling
- Gebruik `FULL_DATE` voor een compacte, consistente breedteweergave
- Gebruik `FULL_DATE_STRING` voor formele rapporten en documenten
- Gebruik `DAYS`, `HOURS`, `MINUTES`, of `SECONDS` voor eenvoudige numerieke weergaven
- Houd rekening met uw UI-ruimtebeperkingen bij het kiezen van een formaat

### SLA-monitoring met Doeltijd
Bij gebruik van `timeDurationTargetTime`:
- Stel de doelduur in seconden in
- Vergelijk de werkelijke duur met de doelduur voor SLA-naleving
- Gebruik in dashboards om achterstallige items te markeren
- Voorbeeld: 24-uurs responssla = 86400 seconden

### Workflow-integratie
- Ontwerp duurvelden om overeen te komen met uw werkelijke bedrijfsprocessen
- Gebruik duurgegevens voor procesverbetering en optimalisatie
- Monitor duurtrends om knelpunten in de workflow te identificeren
- Stel indien nodig waarschuwingen in voor duurgrenzen

## Veelvoorkomende Gebruikscases

1. **Procesprestaties**
   - Taak voltooiingstijden
   - Beoordelingscyclustijden
   - Goedkeuringsverwerkingstijden
   - Responstijden

2. **SLA-monitoring**
   - Tijd tot eerste reactie
   - Oplostijden
   - Escalatie tijdframes
   - Naleving van serviceniveaus

3. **Workflow-analyse**
   - Knelpuntidentificatie
   - Procesoptimalisatie
   - Team prestatiemetrics
   - Kwaliteitsborgingstiming

4. **Projectbeheer**
   - Faseduren
   - Mijlpunt timing
   - Tijd voor resourceallocatie
   - Leveringstijdframes

## Beperkingen

- Duurvelden zijn **alleen-lezen** en kunnen niet handmatig worden ingesteld
- Waarden worden asynchroon berekend en zijn mogelijk niet onmiddellijk beschikbaar
- Vereist dat de juiste triggergebeurtenissen zijn ingesteld in uw workflow
- Kan geen duur berekenen voor gebeurtenissen die nog niet hebben plaatsgevonden
- Beperkt tot het bijhouden van tijd tussen discrete gebeurtenissen (niet continue tijdregistratie)
- Geen ingebouwde SLA-waarschuwingen of meldingen
- Kan meerdere duurcalculaties niet samenvoegen in één veld

## Gerelateerde Bronnen

- [Nummer Velden](/api/custom-fields/number) - Voor handmatige numerieke waarden
- [Datum Velden](/api/custom-fields/date) - Voor specifieke datumregistratie
- [Overzicht van Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten
- [Automatiseringen](/api/automations) - Voor het triggeren van acties op basis van duurgrenzen