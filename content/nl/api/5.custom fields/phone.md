---
title: Telefoon Aangepast Veld
description: Maak telefoonvelden aan om telefoonnummers op te slaan en te valideren met internationale opmaak
---

Telefoon aangepaste velden stellen je in staat om telefoonnummers op te slaan in records met ingebouwde validatie en internationale opmaak. Ze zijn ideaal voor het bijhouden van contactinformatie, noodcontacten of andere telefoon gerelateerde gegevens in je projecten.

## Basis Voorbeeld

Maak een eenvoudig telefoonveld aan:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een telefoonveld met beschrijving:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
  }) {
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
| `name` | String! | ✅ Ja | Weergavenaam van het telefoonveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `PHONE` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

**Opmerking**: Aangepaste velden worden automatisch gekoppeld aan het project op basis van de huidige projectcontext van de gebruiker. Geen `projectId` parameter is vereist.

## Telefoonwaarden Instellen

Om een telefoonwaarde op een record in te stellen of bij te werken:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het telefoon aangepaste veld |
| `text` | String | Nee | Telefoonnummer met landcode |
| `regionCode` | String | Nee | Landcode (automatisch gedetecteerd) |

**Opmerking**: Hoewel `text` optioneel is in het schema, is een telefoonnummer vereist voor het veld om betekenisvol te zijn. Bij gebruik van `setTodoCustomField` wordt er geen validatie uitgevoerd - je kunt elke tekstwaarde en regionCode opslaan. De automatische detectie vindt alleen plaats tijdens het aanmaken van records.

## Records Maken met Telefoonwaarden

Bij het maken van een nieuw record met telefoonwaarden:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
    }
  }
}
```

## Respons Velden

### TodoCustomField Respons

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `text` | String | Het geformatteerde telefoonnummer (internationale opmaak) |
| `regionCode` | String | De landcode (bijv. "US", "GB", "CA") |
| `todo` | Todo! | Het record waar deze waarde bij hoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

## Telefoonnummer Validatie

**Belangrijk**: Telefoonnummer validatie en opmaak vindt alleen plaats bij het aanmaken van nieuwe records via `createTodo`. Bij het bijwerken van bestaande telefoonwaarden met `setTodoCustomField` wordt er geen validatie uitgevoerd en worden de waarden opgeslagen zoals opgegeven.

### Geaccepteerde Indelingen (Tijdens Recordcreatie)
Telefoonnummers moeten een landcode bevatten in een van deze indelingen:

- **E.164 indeling (voorkeur)**: `+12345678900`
- **Internationale indeling**: `+1 234 567 8900`
- **Internationaal met interpunctie**: `+1 (234) 567-8900`
- **Landcode met streepjes**: `+1-234-567-8900`

**Opmerking**: Nationale indelingen zonder landcode (zoals `(234) 567-8900`) worden afgewezen tijdens het aanmaken van records.

### Validatieregels (Tijdens Recordcreatie)
- Gebruikt libphonenumber-js voor parsing en validatie
- Accepteert verschillende internationale telefoonnummer indelingen
- Detecteert automatisch het land op basis van het nummer
- Formatteert nummer in internationale weergave-indeling (bijv. `+1 234 567 8900`)
- Extraheert en slaat landcode apart op (bijv. `US`)

### Geldige Telefoon Voorbeelden
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Ongeldige Telefoon Voorbeelden
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Opslag Indeling

Bij het aanmaken van records met telefoonnummers:
- **tekst**: Opgeslagen in internationale indeling (bijv. `+1 234 567 8900`) na validatie
- **regionCode**: Opgeslagen als ISO landcode (bijv. `US`, `GB`, `CA`) automatisch gedetecteerd

Bij het bijwerken via `setTodoCustomField`:
- **tekst**: Opgeslagen precies zoals opgegeven (geen opmaak)
- **regionCode**: Opgeslagen precies zoals opgegeven (geen validatie)

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|--------|-------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Foutreacties

### Ongeldig Telefoonformaat
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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

### Ontbrekende Landcode
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Beste Praktijken

### Gegevensinvoer
- Zorg ervoor dat je altijd een landcode opneemt in telefoonnummers
- Gebruik E.164 indeling voor consistentie
- Valideer nummers voordat je ze opslaat voor belangrijke operaties
- Houd rekening met regionale voorkeuren voor weergave-opmaak

### Gegevenskwaliteit
- Sla nummers op in internationale indeling voor wereldwijde compatibiliteit
- Gebruik regionCode voor land specifieke functies
- Valideer telefoonnummers voordat je kritieke operaties uitvoert (SMS, oproepen)
- Houd rekening met tijdzone-implicaties voor contactmomenten

### Internationale Overwegingen
- Landcode wordt automatisch gedetecteerd en opgeslagen
- Nummers worden geformatteerd in internationale standaard
- Regionale weergavevoorkeuren kunnen gebruik maken van regionCode
- Houd rekening met lokale belconventies bij het weergeven

## Veelvoorkomende Gebruikscases

1. **Contactbeheer**
   - Telefoonnummers van klanten
   - Contactinformatie van leveranciers
   - Telefoonnummers van teamleden
   - Ondersteuningscontactgegevens

2. **Noodcontacten**
   - Noodcontactnummers
   - On-call contactinformatie
   - Crisisresponscontacten
   - Escalatie telefoonnummers

3. **Klantenservice**
   - Telefoonnummers van klanten
   - Terugbelformulieren voor ondersteuning
   - Verificatie telefoonnummers
   - Follow-up contactnummers

4. **Verkoop & Marketing**
   - Telefoonnummers van leads
   - Campagne contactlijsten
   - Contactinformatie van partners
   - Telefoons van verwijzingsbronnen

## Integratiefuncties

### Met Automatiseringen
- Acties triggeren wanneer telefoonvelden worden bijgewerkt
- SMS-meldingen verzenden naar opgeslagen telefoonnummers
- Follow-up taken creëren op basis van telefoonwijzigingen
- Oproepen routeren op basis van telefoonnummert gegevens

### Met Opzoekingen
- Telefoon gegevens van andere records refereren
- Telefoonlijsten aggregeren uit meerdere bronnen
- Records vinden op telefoonnummer
- Contactinformatie kruisverwijzen

### Met Formulieren
- Automatische telefoonvalidatie
- Controle van internationale indeling
- Detectie van landcode
- Real-time feedback over indeling

## Beperkingen

- Vereist landcode voor alle nummers
- Geen ingebouwde SMS- of belmogelijkheden
- Geen verificatie van telefoonnummers buiten formaatcontrole
- Geen opslag van telefoonmetadata (provider, type, enz.)
- Nationale formatnummers zonder landcode worden afgewezen
- Geen automatische telefoonnummeropmaak in de UI buiten de internationale standaard

## Gerelateerde Bronnen

- [Tekstvelden](/api/custom-fields/text-single) - Voor niet-telefoon tekstgegevens
- [E-mailvelden](/api/custom-fields/email) - Voor e-mailadressen
- [URL-velden](/api/custom-fields/url) - Voor website-adressen
- [Overzicht van Aangepaste Velden](/custom-fields/list-custom-fields) - Algemene concepten