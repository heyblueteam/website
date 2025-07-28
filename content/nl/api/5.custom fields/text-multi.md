---
title: Aangepast Veld voor Meerdere Regels Tekst
description: Maak velden voor meerdere regels tekst voor langere inhoud zoals beschrijvingen, notities en opmerkingen
---

Aangepaste velden voor meerdere regels tekst stellen je in staat om langere tekstinhoud met regelafbrekingen en opmaak op te slaan. Ze zijn ideaal voor beschrijvingen, notities, opmerkingen of elke tekstgegevens die meerdere regels nodig heeft.

## Basisvoorbeeld

Maak een eenvoudig veld voor meerdere regels tekst:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een veld voor meerdere regels tekst met beschrijving:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
    id
    name
    type
    description
  }
}
```

## Invoervelden

### CreateCustomFieldInput

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het tekstveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `TEXT_MULTI` |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |

**Opmerking:** De `projectId` wordt als een apart argument aan de mutatie doorgegeven, niet als onderdeel van het invoerobject. Alternatief kan de projectcontext worden bepaald vanuit de `X-Bloo-Project-ID` header in je GraphQL-verzoek.

## Instellen van Tekstwaarden

Om een waarde voor meerdere regels tekst in een record in te stellen of bij te werken:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het aangepaste tekstveld |
| `text` | String | Nee | Inhoud van meerdere regels tekst om op te slaan |

## Records Maken met Tekstwaarden

Bij het maken van een nieuw record met waarden voor meerdere regels tekst:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
    }
  }
}
```

## Antwoordvelden

### TodoCustomField Antwoord

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `text` | String | De opgeslagen inhoud van meerdere regels tekst |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

## Tekstvalidatie

### Formuliervalidatie
Wanneer velden voor meerdere regels tekst in formulieren worden gebruikt:
- Voorafgaande en volgende spaties worden automatisch verwijderd
- Vereiste validatie wordt toegepast als het veld als vereist is gemarkeerd
- Er wordt geen specifieke opmaakvalidatie toegepast

### Validatieregels
- Accepteert elke tekenreeksinhoud, inclusief regelafbrekingen
- Geen tekenlengtebeperkingen (tot databasebeperkingen)
- Ondersteunt Unicode-tekens en speciale symbolen
- Regelafbrekingen worden bewaard in opslag

### Geldige Tekst Voorbeelden
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Belangrijke Opmerkingen

### Opslagcapaciteit
- Opgeslagen met MySQL `MediumText` type
- Ondersteunt tot 16MB aan tekstinhoud
- Regelafbrekingen en opmaak worden bewaard
- UTF-8 codering voor internationale tekens

### Directe API vs Formulieren
- **Formulieren**: Automatische verwijdering van spaties en vereiste validatie
- **Directe API**: Tekst wordt precies opgeslagen zoals opgegeven
- **Aanbeveling**: Gebruik formulieren voor gebruikersinvoer om consistente opmaak te waarborgen

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI**: Invoer voor meerdere regels tekst, ideaal voor langere inhoud
- **TEXT_SINGLE**: Invoer voor één regel tekst, ideaal voor korte waarden
- **Backend**: Beide typen zijn identiek - zelfde opslagveld, validatie en verwerking
- **Frontend**: Verschillende UI-componenten voor gegevensinvoer (textarea vs invoerveld)
- **Belangrijk**: Het onderscheid tussen TEXT_MULTI en TEXT_SINGLE bestaat puur voor UI-doeleinden

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Best Practices

### Inhoudsorganisatie
- Gebruik consistente opmaak voor gestructureerde inhoud
- Overweeg het gebruik van markdown-achtige syntaxis voor leesbaarheid
- Verdeel lange inhoud in logische secties
- Gebruik regelafbrekingen om de leesbaarheid te verbeteren

### Gegevensinvoer
- Geef duidelijke veldbeschrijvingen om gebruikers te begeleiden
- Gebruik formulieren voor gebruikersinvoer om validatie te waarborgen
- Overweeg tekenlimieten op basis van je gebruiksdoel
- Valideer de inhoudsopmaak in je applicatie indien nodig

### Prestatieoverwegingen
- Zeer lange tekstinhoud kan de query-prestaties beïnvloeden
- Overweeg paginering voor het weergeven van grote tekstvelden
- Indexoverwegingen voor zoekfunctionaliteit
- Houd het opslaggebruik in de gaten voor velden met grote inhoud

## Filteren en Zoeken

### Bevat Zoekopdracht
Velden voor meerdere regels tekst ondersteunen substring-zoekopdrachten via aangepaste veldfilters:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Zoekmogelijkheden
- Substring-matching binnen tekstvelden met behulp van `CONTAINS` operator
- Hoofdletterongevoelige zoekopdracht met behulp van `NCONTAINS` operator
- Exacte overeenkomst met behulp van `IS` operator
- Negatieve overeenkomst met behulp van `NOT` operator
- Doorzoekt alle regels tekst
- Ondersteunt gedeeltelijke woordovereenkomsten

## Veelvoorkomende Gebruiksscenario's

1. **Projectbeheer**
   - Taakbeschrijvingen
   - Projectvereisten
   - Vergadernotities
   - Statusupdates

2. **Klantenservice**
   - Probleembeschrijvingen
   - Oplossingsnotities
   - Klantfeedback
   - Communicatielogs

3. **Inhoudsbeheer**
   - Artikelinhoud
   - Productbeschrijvingen
   - Gebruikersopmerkingen
   - Beoordelingsdetails

4. **Documentatie**
   - Procesbeschrijvingen
   - Instructies
   - Richtlijnen
   - Referentiemateriaal

## Integratiefuncties

### Met Automatiseringen
- Acties triggeren wanneer tekstinhoud verandert
- Sleutelwoorden uit tekstinhoud extraheren
- Samenvattingen of meldingen maken
- Tekstinhoud verwerken met externe diensten

### Met Zoekopdrachten
- Tekstgegevens van andere records refereren
- Tekstinhoud uit meerdere bronnen aggregeren
- Records vinden op basis van tekstinhoud
- Gerelateerde tekstinformatie weergeven

### Met Formulieren
- Automatische verwijdering van spaties
- Validatie van vereiste velden
- UI voor meerdere regels tekst
- Teken telling weergeven (indien geconfigureerd)

## Beperkingen

- Geen ingebouwde tekstopmaak of rijke tekstbewerking
- Geen automatische linkdetectie of conversie
- Geen spellingscontrole of grammaticale validatie
- Geen ingebouwde tekstanalyse of verwerking
- Geen versiebeheer of wijzigingsregistratie
- Beperkte zoekmogelijkheden (geen full-text zoekopdracht)
- Geen inhoudscompressie voor zeer grote tekst

## Gerelateerde Bronnen

- [Enkele Regels Tekstvelden](/api/custom-fields/text-single) - Voor korte tekstwaarden
- [E-mailvelden](/api/custom-fields/email) - Voor e-mailadressen
- [URL-velden](/api/custom-fields/url) - Voor website-adressen
- [Overzicht van Aangepaste Velden](/api/custom-fields/2.list-custom-fields) - Algemene concepten