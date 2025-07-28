---
title: E-mail Aangepast Veld
description: Maak e-mailvelden om e-mailadressen op te slaan en te valideren
---

E-mail aangepaste velden stellen je in staat om e-mailadressen op te slaan in records met ingebouwde validatie. Ze zijn ideaal voor het bijhouden van contactinformatie, e-mailadressen van toegewezen personen of andere e-mailgerelateerde gegevens in je projecten.

## Basisvoorbeeld

Maak een eenvoudig e-mailveld:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een e-mailveld met beschrijving:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ Ja | Weergavenaam van het e-mailveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `EMAIL` |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |

## E-mailwaarden Instellen

Om een e-mailwaarde in een record in te stellen of bij te werken:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het e-mail aangepaste veld |
| `text` | String | Nee | E-mailadres om op te slaan |

## Records Maken met E-mailwaarden

Bij het maken van een nieuw record met e-mailwaarden:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Antwoordvelden

### CustomField Antwoord

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | ID! | Unieke identificatie voor het aangepaste veld |
| `name` | String! | Weergavenaam van het e-mailveld |
| `type` | CustomFieldType! | Het veldtype (EMAIL) |
| `description` | String | Helptekst voor het veld |
| `value` | JSON | Bevat de e-mailwaarde (zie hieronder) |
| `createdAt` | DateTime! | Wanneer het veld is aangemaakt |
| `updatedAt` | DateTime! | Wanneer het veld voor het laatst is gewijzigd |

**Belangrijk**: E-mailwaarden worden benaderd via het `customField.value.text` veld, niet direct op de respons.

## E-mailwaarden Opvragen

Bij het opvragen van records met e-mail aangepaste velden, krijg je toegang tot de e-mail via het `customField.value.text` pad:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

De respons bevat de e-mail in de geneste structuur:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## E-mailvalidatie

### Formulier Validatie
Wanneer e-mailvelden in formulieren worden gebruikt, valideren ze automatisch het e-mailformaat:
- Gebruikt standaard e-mailvalidatieregels
- Verwijdert spaties uit invoer
- Weigert ongeldige e-mailformaten

### Validatieregels
- Moet een `@` symbool bevatten
- Moet een geldig domeinformaat hebben
- Voorafgaande/achterafgaande spaties worden automatisch verwijderd
- Veelvoorkomende e-mailformaten worden geaccepteerd

### Geldige E-mail Voorbeelden
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Ongeldige E-mail Voorbeelden
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Belangrijke Opmerkingen

### Directe API vs Formulieren
- **Formulieren**: Automatische e-mailvalidatie wordt toegepast
- **Directe API**: Geen validatie - elke tekst kan worden opgeslagen
- **Aanbeveling**: Gebruik formulieren voor gebruikersinvoer om validatie te waarborgen

### Opslagformaat
- E-mailadressen worden opgeslagen als platte tekst
- Geen speciale opmaak of parsing
- Hoofdlettergevoeligheid: EMAIL aangepaste velden worden hoofdlettergevoelig opgeslagen (in tegenstelling tot e-mails voor gebruikersauthenticatie die worden genormaliseerd naar kleine letters)
- Geen maximale lengtebeperkingen behalve databasebeperkingen (16 MB limiet)

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Foutreacties

### Ongeldig E-mailformaat (Alleen Formulieren)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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

### Gegevensinvoer
- Valideer altijd e-mailadressen in je applicatie
- Gebruik e-mailvelden alleen voor daadwerkelijke e-mailadressen
- Overweeg om formulieren te gebruiken voor gebruikersinvoer om automatische validatie te krijgen

### Gegevenskwaliteit
- Verwijder spaties voordat je opslaat
- Overweeg hoofdletternormalisatie (typisch kleine letters)
- Valideer e-mailformaat voordat je belangrijke bewerkingen uitvoert

### Privacy Overwegingen
- E-mailadressen worden opgeslagen als platte tekst
- Houd rekening met gegevensprivacyregelgeving (GDPR, CCPA)
- Implementeer geschikte toegangscontroles

## Veelvoorkomende Gebruiksscenario's

1. **Contactbeheer**
   - E-mailadressen van klanten
   - Contactinformatie van leveranciers
   - E-mailadressen van teamleden
   - Contactgegevens voor ondersteuning

2. **Projectbeheer**
   - E-mailadressen van belanghebbenden
   - E-mailadressen voor goedkeuring
   - Ontvangers van meldingen
   - E-mailadressen van externe samenwerkers

3. **Klantenservice**
   - E-mailadressen van klanten
   - Contacten voor ondersteuningsverzoeken
   - Escalatiecontacten
   - Feedback e-mailadressen

4. **Verkoop & Marketing**
   - E-mailadressen van leads
   - Campagne contactlijsten
   - Contactinformatie van partners
   - E-mailadressen van verwijzingsbronnen

## Integratiefuncties

### Met Automatiseringen
- Acties triggeren wanneer e-mailvelden worden bijgewerkt
- Meldingen verzenden naar opgeslagen e-mailadressen
- Volgtaken creëren op basis van e-mailwijzigingen

### Met Opzoekingen
- E-mailgegevens uit andere records refereren
- E-maillijsten uit meerdere bronnen aggregeren
- Records vinden op basis van e-mailadres

### Met Formulieren
- Automatische e-mailvalidatie
- E-mailformaatcontrole
- Verwijdering van spaties

## Beperkingen

- Geen ingebouwde e-mailverificatie of validatie buiten formaatcontrole
- Geen e-mail specifieke UI-functies (zoals klikbare e-maillinks)
- Opgeslagen als platte tekst zonder encryptie
- Geen mogelijkheden voor e-mailopstelling of -verzending
- Geen opslag van e-mailmetadata (weergavenaam, enz.)
- Directe API-aanroepen omzeilen validatie (alleen formulieren valideren)

## Gerelateerde Bronnen

- [Tekstvelden](/api/custom-fields/text-single) - Voor niet-e-mail tekstgegevens
- [URL-velden](/api/custom-fields/url) - Voor website-adressen
- [Telefoonvelden](/api/custom-fields/phone) - Voor telefoonnummers
- [Overzicht van Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten