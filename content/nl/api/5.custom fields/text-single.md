---
title: Aangepast Veld voor Enkelregelige Tekst
description: Maak enkelregelige tekstvelden voor korte tekstwaarden zoals namen, titels en labels
---

Aangepaste velden voor enkelregelige tekst stellen je in staat om korte tekstwaarden op te slaan die bedoeld zijn voor invoer op één regel. Ze zijn ideaal voor namen, titels, labels of andere tekstgegevens die op één regel moeten worden weergegeven.

## Basisvoorbeeld

Maak een eenvoudig enkelregelige tekstveld:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een enkelregelige tekstveld met beschrijving:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
  }) {
    id
    name
    type
    description
  }
}
```

## Invoervariabelen

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het tekstveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `TEXT_SINGLE` |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |

**Opmerking**: De projectcontext wordt automatisch bepaald aan de hand van je authenticatieheaders. Geen `projectId` parameter is nodig.

## Instellen van Tekstwaarden

Om een enkelregelige tekstwaarde op een record in te stellen of bij te werken:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het tekst aangepaste veld |
| `text` | String | Nee | Inhoud van de enkelregelige tekst om op te slaan |

## Records Maken met Tekstwaarden

Bij het maken van een nieuw record met enkelregelige tekstwaarden:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | ID! | Unieke identifier voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld (bevat de tekstwaarde) |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

**Belangrijk**: Tekstwaarden worden benaderd via het `customField.value.text` veld, niet direct op TodoCustomField.

## Opvragen van Tekstwaarden

Bij het opvragen van records met tekst aangepaste velden, krijg je toegang tot de tekst via het `customField.value.text` pad:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

Het antwoord bevat de tekst in de geneste structuur:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Tekstvalidatie

### Formulier Validatie
Wanneer enkelregelige tekstvelden in formulieren worden gebruikt:
- Voorafgaande en achterafgaande spaties worden automatisch verwijderd
- Vereiste validatie wordt toegepast als het veld als vereist is gemarkeerd
- Geen specifieke opmaakvalidatie wordt toegepast

### Validatieregels
- Accepteert elke tekenreeksinhoud, inclusief regelafbrekingen (hoewel niet aanbevolen)
- Geen beperking op het aantal tekens (tot de databasebeperkingen)
- Ondersteunt Unicode-tekens en speciale symbolen
- Regelafbrekingen worden behouden, maar zijn niet bedoeld voor dit type veld

### Typische Tekst Voorbeelden
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Belangrijke Opmerkingen

### Opslagcapaciteit
- Opgeslagen met MySQL `MediumText` type
- Ondersteunt tot 16MB aan tekstinhoud
- Identieke opslag aan meerregelige tekstvelden
- UTF-8 codering voor internationale tekens

### Directe API vs Formulieren
- **Formulieren**: Automatische verwijdering van spaties en vereiste validatie
- **Directe API**: Tekst wordt precies opgeslagen zoals opgegeven
- **Aanbeveling**: Gebruik formulieren voor gebruikersinvoer om consistente opmaak te waarborgen

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: Enkelregelige tekstinvoer, ideaal voor korte waarden
- **TEXT_MULTI**: Meerregelige tekstinvoer, ideaal voor langere inhoud
- **Backend**: Beide gebruiken identieke opslag en validatie
- **Frontend**: Verschillende UI-componenten voor gegevensinvoer
- **Intentie**: TEXT_SINGLE is semantisch bedoeld voor enkelregelige waarden

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## Foutantwoorden

### Vereiste Veldvalidatie (Alleen Formulieren)
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
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
      "code": "NOT_FOUND"
    }
  }]
}
```

## Beste Praktijken

### Inhoud Richtlijnen
- Houd tekst beknopt en geschikt voor enkelregelige invoer
- Vermijd regelafbrekingen voor beoogde weergave op één regel
- Gebruik consistente opmaak voor vergelijkbare gegevenstypen
- Overweeg tekenlimieten op basis van je UI-vereisten

### Gegevensinvoer
- Geef duidelijke veldbeschrijvingen om gebruikers te begeleiden
- Gebruik formulieren voor gebruikersinvoer om validatie te waarborgen
- Valideer de inhoudsopmaak in je applicatie indien nodig
- Overweeg het gebruik van dropdowns voor gestandaardiseerde waarden

### Prestatieoverwegingen
- Enkelregelige tekstvelden zijn lichtgewicht en performant
- Overweeg indexering voor vaak doorzochte velden
- Gebruik geschikte weergavebreedtes in je UI
- Houd de inhoudslengte in de gaten voor weergave doeleinden

## Filteren en Zoeken

### Bevat Zoekopdracht
Enkelregelige tekstvelden ondersteunen substring-zoekopdrachten:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Zoekmogelijkheden
- Hoofdletterongevoelige substring-overeenkomsten
- Ondersteunt gedeeltelijke woordovereenkomsten
- Exacte waardeovereenkomsten
- Geen full-text zoekopdracht of ranking

## Veelvoorkomende Gebruikscases

1. **Identificatoren en Codes**
   - Product SKU's
   - Ordernummers
   - Referentiecodes
   - Versienummers

2. **Namen en Titels**
   - Klantnamen
   - Projecttitels
   - Productnamen
   - Categorie-labels

3. **Korte Beschrijvingen**
   - Korte samenvattingen
   - Statuslabels
   - Prioriteitsindicatoren
   - Classificatietags

4. **Externe Referenties**
   - Ticketnummers
   - Factuurreferenties
   - ID's van externe systemen
   - Documentnummers

## Integratiefuncties

### Met Zoekopdrachten
- Referentie tekstgegevens van andere records
- Vind records op basis van tekstinhoud
- Toon gerelateerde tekstinformatie
- Agregeer tekstwaarden uit meerdere bronnen

### Met Formulieren
- Automatische verwijdering van spaties
- Validatie van vereiste velden
- Enkelregelige tekstinvoerveld UI
- Weergave van tekenlimieten (indien geconfigureerd)

### Met Import/Export
- Directe CSV-kolomtoewijzing
- Automatische toewijzing van tekstwaarden
- Ondersteuning voor bulkgegevensimport
- Exporteren naar spreadsheetformaten

## Beperkingen

### Automatiseringsbeperkingen
- Niet direct beschikbaar als automatisering trigger-velden
- Kunnen niet worden gebruikt in automatiseringsveldupdates
- Kunnen worden verwezen in automatiseringsvoorwaarden
- Beschikbaar in e-mailtemplates en webhooks

### Algemene Beperkingen
- Geen ingebouwde tekstopmaak of styling
- Geen automatische validatie buiten vereiste velden
- Geen ingebouwde uniciteitsafdwinging
- Geen inhoudscompressie voor zeer grote tekst
- Geen versiebeheer of wijzigingsregistratie
- Beperkte zoekmogelijkheden (geen full-text zoekopdracht)

## Gerelateerde Bronnen

- [Meerregelige Tekstvelden](/api/custom-fields/text-multi) - Voor langere tekstinhoud
- [E-mailvelden](/api/custom-fields/email) - Voor e-mailadressen
- [URL-velden](/api/custom-fields/url) - Voor website-adressen
- [Unieke ID-velden](/api/custom-fields/unique-id) - Voor automatisch gegenereerde identificatoren
- [Overzicht van Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten