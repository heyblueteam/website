---
title: Checkbox Aangepast Veld
description: Maak boolean checkboxvelden voor ja/nee of waar/onwaar gegevens
---

Checkbox aangepaste velden bieden een eenvoudige boolean (waar/onwaar) invoer voor taken. Ze zijn perfect voor binaire keuzes, statusindicatoren of om bij te houden of iets is voltooid.

## Basisvoorbeeld

Maak een eenvoudig checkboxveld:

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

## Geavanceerd Voorbeeld

Maak een checkboxveld met beschrijving en validatie:

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

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van de checkbox |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `CHECKBOX` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

## Checkboxwaarden Instellen

Om een checkboxwaarde op een taak in te stellen of bij te werken:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Om een checkbox uit te vinken:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van de taak om bij te werken |
| `customFieldId` | String! | ✅ Ja | ID van het checkbox aangepaste veld |
| `checked` | Boolean | Nee | Waar om aan te vinken, onwaar om uit te vinken |

## Taken Maken met Checkboxwaarden

Bij het maken van een nieuwe taak met checkboxwaarden:

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

### Geaccepteerde Stringwaarden

Bij het maken van taken moeten checkboxwaarden als strings worden doorgegeven:

| Stringwaarde | Resultaat |
|--------------|-----------|
| `"true"` | ✅ Aangevinkt (hoofdlettergevoelig) |
| `"1"` | ✅ Aangevinkt |
| `"checked"` | ✅ Aangevinkt (hoofdlettergevoelig) |
| Any other value | ❌ Niet aangevinkt |

**Opmerking**: Stringvergelijkingen tijdens het maken van taken zijn hoofdlettergevoelig. De waarden moeten exact overeenkomen met `"true"`, `"1"`, of `"checked"` om een aangevinkt staat te krijgen.

## Responsvelden

### TodoCustomField Respons

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | ID! | Unieke identificatie voor de veldwaarde |
| `uid` | String! | Alternatieve unieke identificatie |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `checked` | Boolean | De checkboxstatus (waar/onwaar/nul) |
| `todo` | Todo! | De taak waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

## Automatiseringsintegratie

Checkboxvelden triggeren verschillende automatiseringsevents op basis van statuswijzigingen:

| Actie | Geactiveerd Event | Beschrijving |
|-------|------------------|--------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Geactiveerd wanneer de checkbox is aangevinkt |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Geactiveerd wanneer de checkbox is uitgevinkt |

Dit stelt je in staat om automatiseringen te creëren die reageren op wijzigingen in de checkboxstatus, zoals:
- Het verzenden van meldingen wanneer items zijn goedgekeurd
- Het verplaatsen van taken wanneer beoordelingscheckboxen zijn aangevinkt
- Het bijwerken van gerelateerde velden op basis van checkboxstatussen

## Gegevensimport/-export

### Checkboxwaarden Importeren

Bij het importeren van gegevens via CSV of andere formaten:
- `"true"`, `"yes"` → Aangevinkt (hoofdletterongevoelig)
- Elke andere waarde (inclusief `"false"`, `"no"`, `"0"`, leeg) → Niet aangevinkt

### Checkboxwaarden Exporteren

Bij het exporteren van gegevens:
- Aangevinkte vakken worden geëxporteerd als `"X"`
- Niet aangevinkte vakken worden geëxporteerd als lege string `""`

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Foutreacties

### Ongeldig Waartype
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

### Veld Niet Gevonden
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

## Beste Praktijken

### Naamgevingsconventies
- Gebruik duidelijke, actiegerichte namen: "Goedgekeurd", "Beoordeeld", "Is Voltooid"
- Vermijd negatieve namen die gebruikers verwarren: geef de voorkeur aan "Is Actief" boven "Is Inactief"
- Wees specifiek over wat de checkbox vertegenwoordigt

### Wanneer Checkboxen te Gebruiken
- **Binaire keuzes**: Ja/Nee, Waar/Onwaar, Voltooid/Niet Voltooid
- **Statusindicatoren**: Goedgekeurd, Beoordeeld, Gepubliceerd
- **Functievlaggen**: Heeft Prioriteitsondersteuning, Vereist Handtekening
- **Eenvoudige tracking**: E-mail Verzonden, Factuur Betaald, Item Verzonden

### Wanneer GEEN Checkboxen te Gebruiken
- Wanneer je meer dan twee opties nodig hebt (gebruik SELECT_SINGLE in plaats daarvan)
- Voor numerieke of tekstgegevens (gebruik NUMMER of TEKST velden)
- Wanneer je moet bijhouden wie het heeft aangevinkt of wanneer (gebruik auditlogs)

## Veelvoorkomende Gebruikscases

1. **Goedkeuringsworkflows**
   - "Manager Goedgekeurd"
   - "Client Handtekening"
   - "Juridische Beoordeling Voltooid"

2. **Taakbeheer**
   - "Is Geblokkeerd"
   - "Klaar voor Beoordeling"
   - "Hoge Prioriteit"

3. **Kwaliteitscontrole**
   - "QA Geslaagd"
   - "Documentatie Voltooid"
   - "Tests Geschreven"

4. **Administratieve Vlaggen**
   - "Factuur Verzonden"
   - "Contract Ondertekend"
   - "Opvolging Vereist"

## Beperkingen

- Checkboxvelden kunnen alleen waar/onwaar waarden opslaan (geen tri-state of nul na de eerste instelling)
- Geen standaardwaardeconfiguratie (begint altijd als nul totdat ingesteld)
- Kan geen aanvullende metadata opslaan zoals wie het heeft aangevinkt of wanneer
- Geen voorwaardelijke zichtbaarheid op basis van andere veldwaarden

## Gerelateerde Bronnen

- [Overzicht Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten van aangepaste velden
- [Automatiseringen API](/api/automations) - Maak automatiseringen die worden geactiveerd door wijzigingen in checkboxen