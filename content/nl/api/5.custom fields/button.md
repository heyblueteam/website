---
title: Aangepast Knopveld
description: Maak interactieve knopvelden die automatiseringen activeren wanneer erop geklikt wordt
---

Aangepaste knopvelden bieden interactieve UI-elementen die automatiseringen activeren wanneer erop geklikt wordt. In tegenstelling tot andere soorten aangepaste velden die gegevens opslaan, dienen knopvelden als actie-triggers om geconfigureerde workflows uit te voeren.

## Basisvoorbeeld

Maak een eenvoudig knopveld dat een automatisering activeert:

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

## Geavanceerd Voorbeeld

Maak een knop met bevestigingsvereisten:

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

## Invoergegevens

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van de knop |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `BUTTON` |
| `projectId` | String! | ✅ Ja | Project-ID waar het veld zal worden aangemaakt |
| `buttonType` | String | Nee | Bevestigingsgedrag (zie Knoptypen hieronder) |
| `buttonConfirmText` | String | Nee | Tekst die gebruikers moeten typen voor harde bevestiging |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |
| `required` | Boolean | Nee | Of het veld vereist is (standaard is false) |
| `isActive` | Boolean | Nee | Of het veld actief is (standaard is true) |

### Knoptype Veld

Het `buttonType` veld is een vrije tekststring die door UI-clients kan worden gebruikt om het bevestigingsgedrag te bepalen. Veelvoorkomende waarden zijn:

- `""` (leeg) - Geen bevestiging
- `"soft"` - Eenvoudige bevestigingsdialoog
- `"hard"` - Vereist het typen van bevestigingstekst

**Opmerking**: Dit zijn alleen UI-hints. De API valideert of handhaaft geen specifieke waarden.

## Activeren van Knopklikken

Om een knopklik te activeren en bijbehorende automatiseringen uit te voeren:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Klik Invoergegevens

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van de taak die de knop bevat |
| `customFieldId` | String! | ✅ Ja | ID van het aangepaste knopveld |

### Belangrijk: API Gedrag

**Alle knopklikken via de API worden onmiddellijk uitgevoerd** ongeacht `buttonType` of `buttonConfirmText` instellingen. Deze velden worden opgeslagen voor UI-clients om bevestigingsdialoogvensters te implementeren, maar de API zelf:

- Valideert geen bevestigingstekst
- Handhaaft geen bevestigingsvereisten
- Voert de knopactie onmiddellijk uit wanneer deze wordt aangeroepen

Bevestiging is puur een client-side UI veiligheidsfunctie.

### Voorbeeld: Verschillende Knoptypen Klikken

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

Alle drie de mutaties hierboven zullen de knopactie onmiddellijk uitvoeren wanneer ze via de API worden aangeroepen, waarbij eventuele bevestigingsvereisten worden omzeild.

## Respons Velden

### Aangepaste Veld Respons

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor het aangepaste veld |
| `name` | String! | Weergavenaam van de knop |
| `type` | CustomFieldType! | Altijd `BUTTON` voor knopvelden |
| `buttonType` | String | Instelling voor bevestigingsgedrag |
| `buttonConfirmText` | String | Vereiste bevestigingstekst (indien harde bevestiging wordt gebruikt) |
| `description` | String | Helptekst voor gebruikers |
| `required` | Boolean! | Of het veld vereist is |
| `isActive` | Boolean! | Of het veld momenteel actief is |
| `projectId` | String! | ID van het project waartoe dit veld behoort |
| `createdAt` | DateTime! | Wanneer het veld is aangemaakt |
| `updatedAt` | DateTime! | Wanneer het veld voor het laatst is gewijzigd |

## Hoe Knopvelden Werken

### Automatiseringsintegratie

Knopvelden zijn ontworpen om samen te werken met het automatiseringssysteem van Blue:

1. **Maak het knopveld** met behulp van de bovenstaande mutatie
2. **Configureer automatiseringen** die luisteren naar `CUSTOM_FIELD_BUTTON_CLICKED` evenementen
3. **Gebruikers klikken op de knop** in de UI
4. **Automatiseringen voeren** de geconfigureerde acties uit

### Evenementstroom

Wanneer een knop wordt geklikt:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Geen Gegevensopslag

Belangrijk: Knopvelden slaan geen waardegegevens op. Ze dienen puur als actie-triggers. Elke klik:
- Genereert een evenement
- Activeert bijbehorende automatiseringen
- Registreert een actie in de taakgeschiedenis
- Wijzigt geen veldwaarde

## Vereiste Machtigingen

Gebruikers hebben geschikte projectrollen nodig om knopvelden te maken en te gebruiken:

| Actie | Vereiste Rol |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Foutreacties

### Toegang Geweigerd
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

### Aangepast Veld Niet Gevonden
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

**Opmerking**: De API retourneert geen specifieke fouten voor ontbrekende automatiseringen of bevestigingsmismatches.

## Beste Praktijken

### Naamgevingsconventies
- Gebruik actiegerichte namen: "Factuur Verzenden", "Rapport Maken", "Team Notificeren"
- Wees specifiek over wat de knop doet
- Vermijd generieke namen zoals "Knop 1" of "Klik Hier"

### Bevestigingsinstellingen
- Laat `buttonType` leeg voor veilige, omkeerbare acties
- Stel `buttonType` in om bevestigingsgedrag aan UI-clients voor te stellen
- Gebruik `buttonConfirmText` om aan te geven wat gebruikers moeten typen in UI-bevestigingen
- Vergeet niet: Dit zijn alleen UI-hints - API-aanroepen worden altijd onmiddellijk uitgevoerd

### Automatiseringsontwerp
- Houd knopacties gericht op een enkele workflow
- Geef duidelijke feedback over wat er is gebeurd na het klikken
- Overweeg om beschrijvingstekst toe te voegen om het doel van de knop uit te leggen

## Veelvoorkomende Gebruikscases

1. **Workflow Overgangen**
   - "Markeer als Voltooid"
   - "Verzend voor Goedkeuring"
   - "Archiveren Taak"

2. **Externe Integraties**
   - "Synchroniseer naar CRM"
   - "Genereer Factuur"
   - "Verzend E-mail Update"

3. **Batchbewerkingen**
   - "Werk Alle Subtaken Bij"
   - "Kopieer naar Projecten"
   - "Pas Sjabloon Toe"

4. **Rapportage Acties**
   - "Genereer Rapport"
   - "Exporteer Gegevens"
   - "Maak Samenvatting"

## Beperkingen

- Knoppen kunnen geen gegevenswaarden opslaan of weergeven
- Elke knop kan alleen automatiseringen activeren, geen directe API-aanroepen (automatiseringen kunnen echter HTTP-verzoekacties bevatten om externe API's of Blue's eigen API's aan te roepen)
- De zichtbaarheid van knoppen kan niet voorwaardelijk worden gecontroleerd
- Maximaal één automatiseringsuitvoering per klik (hoewel die automatisering meerdere acties kan activeren)

## Gerelateerde Bronnen

- [Automatiserings-API](/api/automations/index) - Configureer acties die door knoppen worden geactiveerd
- [Overzicht Aangepaste Velden](/custom-fields/list-custom-fields) - Algemene concepten van aangepaste velden