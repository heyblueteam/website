---
title: URL Aangepast Veld
description: Maak URL-velden aan om website-adressen en links op te slaan
---

URL aangepaste velden stellen je in staat om website-adressen en links in je records op te slaan. Ze zijn ideaal voor het bijhouden van projectwebsites, referentielinks, documentatie-URL's of andere webgebaseerde bronnen die verband houden met je werk.

## Basis Voorbeeld

Maak een eenvoudig URL-veld aan:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een URL-veld met beschrijving:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het URL-veld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `URL` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

**Opmerking:** De `projectId` wordt als een apart argument aan de mutatie doorgegeven, niet als onderdeel van het invoerobject.

## URL-waarden Instellen

Om een URL-waarde op een record in te stellen of bij te werken:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het URL-aangepaste veld |
| `text` | String! | ✅ Ja | URL-adres om op te slaan |

## Records Maken met URL-waarden

Bij het maken van een nieuw record met URL-waarden:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
|------|------|--------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `text` | String | Het opgeslagen URL-adres |
| `todo` | Todo! | Het record waar deze waarde bij hoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

## URL-validatie

### Huidige Implementatie
- **Directe API**: Momenteel wordt er geen URL-formaatvalidatie afgedwongen
- **Formulieren**: URL-validatie is gepland maar momenteel niet actief
- **Opslag**: Elke tekenreeks kan in URL-velden worden opgeslagen

### Geplande Validatie
Toekomstige versies zullen bevatten:
- HTTP/HTTPS protocolvalidatie
- Controle op geldig URL-formaat
- Validatie van domeinnamen
- Automatische toevoeging van protocolprefix

### Aanbevolen URL-formaten
Hoewel momenteel niet afgedwongen, gebruik deze standaardformaten:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Belangrijke Opmerkingen

### Opslagformaat
- URL's worden als platte tekst zonder wijziging opgeslagen
- Geen automatische toevoeging van protocol (http://, https://)
- Hoofdlettergevoeligheid behouden zoals ingevoerd
- Geen URL-encoding/decoding uitgevoerd

### Directe API vs Formulieren
- **Formulieren**: Geplande URL-validatie (momenteel niet actief)
- **Directe API**: Geen validatie - elke tekst kan worden opgeslagen
- **Aanbeveling**: Valideer URL's in je applicatie voordat je ze opslaat

### URL vs Tekstvelden
- **URL**: Semantisch bedoeld voor webadressen
- **TEXT_SINGLE**: Algemeen enkel-regel tekst
- **Backend**: Momenteel identieke opslag en validatie
- **Frontend**: Verschillende UI-componenten voor gegevensinvoer

## Vereiste Machtigingen

Aangepaste veldbewerkingen gebruiken rolgebaseerde machtigingen:

| Actie | Vereiste Rol |
|-------|--------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Opmerking:** Machtigingen worden gecontroleerd op basis van gebruikersrollen in het project, niet op specifieke machtigingsconstanten.

## Foutreacties

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

## Beste Praktijken

### URL-formaatstandaarden
- Voeg altijd een protocol toe (http:// of https://)
- Gebruik HTTPS waar mogelijk voor beveiliging
- Test URL's voordat je ze opslaat om ervoor te zorgen dat ze toegankelijk zijn
- Overweeg het gebruik van verkorte URL's voor weergave doeleinden

### Gegevenskwaliteit
- Valideer URL's in je applicatie voordat je ze opslaat
- Controleer op veelvoorkomende typfouten (ontbrekende protocollen, onjuiste domeinen)
- Standaardiseer URL-formaten binnen je organisatie
- Overweeg de toegankelijkheid en beschikbaarheid van URL's

### Beveiligingsoverwegingen
- Wees voorzichtig met door gebruikers opgegeven URL's
- Valideer domeinen als je beperkt tot specifieke sites
- Overweeg URL-scanning op kwaadaardige inhoud
- Gebruik HTTPS-URL's bij het omgaan met gevoelige gegevens

## Filteren en Zoeken

### Bevat Zoekopdracht
URL-velden ondersteunen substring-zoekopdrachten:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Zoekmogelijkheden
- Hoofdletterongevoelige substring-overeenkomsten
- Gedeeltelijke domeinovereenkomsten
- Pad- en parameterzoekopdrachten
- Geen protocol-specifieke filtering

## Veelvoorkomende Gebruikscases

1. **Projectmanagement**
   - Projectwebsites
   - Documentatielinks
   - Repository-URL's
   - Demo-sites

2. **Contentbeheer**
   - Referentiematerialen
   - Bronnenlinks
   - Mediaresources
   - Externe artikelen

3. **Klantenservice**
   - Klantwebsites
   - Ondersteuningsdocumentatie
   - Kennisbankartikelen
   - Video-tutorials

4. **Verkoop & Marketing**
   - Bedrijfswebsites
   - Productpagina's
   - Marketingmaterialen
   - Social media-profielen

## Integratiefuncties

### Met Zoekopdrachten
- Referentie-URL's van andere records
- Vind records op domein of URL-patroon
- Toon gerelateerde webbronnen
- Agregeer links van meerdere bronnen

### Met Formulieren
- URL-specifieke invoercomponenten
- Geplande validatie voor juist URL-formaat
- Linkvoorbeeldmogelijkheden (frontend)
- Klikbare URL-weergave

### Met Rapportage
- Volg het gebruik en de patronen van URL's
- Monitor gebroken of ontoegankelijke links
- Categoriseer op domein of protocol
- Exporteer URL-lijsten voor analyse

## Beperkingen

### Huidige Beperkingen
- Geen actieve URL-formaatvalidatie
- Geen automatische protocoltoevoeging
- Geen linkverificatie of toegankelijkheidscontrole
- Geen URL-verkorting of -uitbreiding
- Geen favicon of voorbeeldgeneratie

### Automatiseringsbeperkingen
- Niet beschikbaar als automatiseringstrigger-velden
- Kunnen niet worden gebruikt in automatiseringsveldupdates
- Kunnen worden verwezen in automatiseringsvoorwaarden
- Beschikbaar in e-mailtemplates en webhooks

### Algemene Beperkingen
- Geen ingebouwde linkvoorbeeldfunctionaliteit
- Geen automatische URL-verkorting
- Geen kliktracking of analytics
- Geen URL-vervalcontrole
- Geen kwaadaardige URL-scanning

## Toekomstige Verbeteringen

### Geplande Functies
- HTTP/HTTPS protocolvalidatie
- Aangepaste regex-validatiepatronen
- Automatische toevoeging van protocolprefix
- URL-toegankelijkheidscontrole

### Potentiële Verbeteringen
- Linkvoorbeeldgeneratie
- Favicon-weergave
- Integratie van URL-verkorting
- Kliktrackingmogelijkheden
- Detectie van gebroken links

## Gerelateerde Bronnen

- [Tekstvelden](/api/custom-fields/text-single) - Voor niet-URL tekstgegevens
- [E-mailvelden](/api/custom-fields/email) - Voor e-mailadressen
- [Overzicht van Aangepaste Velden](/api/custom-fields/2.list-custom-fields) - Algemene concepten

## Migratie van Tekstvelden

Als je migreert van tekstvelden naar URL-velden:

1. **Maak een URL-veld** met dezelfde naam en configuratie
2. **Exporteer bestaande tekstwaarden** om te verifiëren of ze geldige URL's zijn
3. **Werk records bij** om het nieuwe URL-veld te gebruiken
4. **Verwijder het oude tekstveld** na succesvolle migratie
5. **Werk applicaties bij** om URL-specifieke UI-componenten te gebruiken

### Migratie Voorbeeld
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```