---
title: Unieke ID Aangepast Veld
description: Maak automatisch gegenereerde unieke identificatievelden met sequentiële nummering en aangepaste opmaak
---

Unieke ID aangepaste velden genereren automatisch sequentiële, unieke identificatoren voor uw records. Ze zijn perfect voor het maken van ticketnummers, order-ID's, factuurnummers of elk sequentieel identificatiesysteem in uw workflow.

## Basis Voorbeeld

Maak een eenvoudig uniek ID-veld met automatische nummering:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Geavanceerd Voorbeeld

Maak een geformatteerd uniek ID-veld met een voorvoegsel en nul-padding:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Invoergegevens

### CreateCustomFieldInput (UNIQUE_ID)

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|---------|--------------|
| `name` | String! | ✅ Ja | Weergavenaam van het unieke ID-veld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `UNIQUE_ID` |
| `description` | String | Nee | Helptekst die aan gebruikers wordt getoond |
| `useSequenceUniqueId` | Boolean | Nee | Automatische nummering inschakelen (standaard: false) |
| `prefix` | String | Nee | Tekstvoorvoegsel voor gegenereerde ID's (bijv. "TAKEN-") |
| `sequenceDigits` | Int | Nee | Aantal cijfers voor nul-padding |
| `sequenceStartingNumber` | Int | Nee | Startnummer voor de reeks |

## Configuratieopties

### Automatische Nummering (`useSequenceUniqueId`)
- **true**: Genereert automatisch sequentiële ID's wanneer records worden aangemaakt
- **false** of **onbepaald**: Handmatige invoer vereist (werkt als een tekstveld)

### Voorvoegsel (`prefix`)
- Optioneel tekstvoorvoegsel toegevoegd aan alle gegenereerde ID's
- Voorbeelden: "TAKEN-", "ORD-", "BUG-", "REQ-"
- Geen lengtebeperking, maar houd het redelijk voor weergave

### Sequentiecijfers (`sequenceDigits`)
- Aantal cijfers voor nul-padding van het sequentienummer
- Voorbeeld: `sequenceDigits: 3` produceert `001`, `002`, `003`
- Als niet gespecificeerd, wordt er geen padding toegepast

### Startnummer (`sequenceStartingNumber`)
- Het eerste nummer in de reeks
- Voorbeeld: `sequenceStartingNumber: 1000` begint bij 1000, 1001, 1002...
- Als niet gespecificeerd, begint het bij 1 (standaardgedrag)

## Geproduceerd ID-formaat

Het uiteindelijke ID-formaat combineert alle configuratieopties:

```
{prefix}{paddedSequenceNumber}
```

### Voorbeeldformaten

| Configuratie | Geproduceerde ID's |
|---------------|--------------------|
| Geen opties | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Lezen van unieke ID-waarden

### Query Records met Unieke ID's
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Antwoordformaat
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Automatische ID-generatie

### Wanneer ID's worden gegenereerd
- **Recordcreatie**: ID's worden automatisch toegewezen wanneer nieuwe records worden aangemaakt
- **Veldtoevoeging**: Bij het toevoegen van een UNIQUE_ID-veld aan bestaande records wordt een achtergrondtaak in de wachtrij geplaatst (implementatie van de worker in afwachting)
- **Achtergrondverwerking**: ID-generatie voor nieuwe records gebeurt synchronisch via database-triggers

### Generatieproces
1. **Trigger**: Nieuw record wordt aangemaakt of UNIQUE_ID-veld wordt toegevoegd
2. **Sequentie-opzoeking**: Systeem vindt het volgende beschikbare sequentienummer
3. **ID-toewijzing**: Sequentienummer wordt toegewezen aan het record
4. **Tellerupdate**: Sequentieteller wordt verhoogd voor toekomstige records
5. **Opmaak**: ID wordt geformatteerd met voorvoegsel en padding wanneer weergegeven

### Uniciteitsgaranties
- **Databasebeperkingen**: Unieke beperking op sequentie-ID's binnen elk veld
- **Atomische bewerkingen**: Sequentiegeneratie gebruikt databasevergrendelingen om duplicaten te voorkomen
- **Projectafbakening**: Sequenties zijn onafhankelijk per project
- **Raceconditie-bescherming**: Gelijktijdige verzoeken worden veilig afgehandeld

## Handmatige versus Automatische Modus

### Automatische Modus (`useSequenceUniqueId: true`)
- ID's worden automatisch gegenereerd via database-triggers
- Sequentiële nummering is gegarandeerd
- Atomische sequentiegeneratie voorkomt duplicaten
- Geformatteerde ID's combineren voorvoegsel + gepadded sequentienummer

### Handmatige Modus (`useSequenceUniqueId: false` of `undefined`)
- Werkt als een regulier tekstveld
- Gebruikers kunnen aangepaste waarden invoeren via `setTodoCustomField` met `text` parameter
- Geen automatische generatie
- Geen uniciteitsafdwinging buiten databasebeperkingen

## Handmatige Waarden Instellen (Alleen Handmatige Modus)

Wanneer `useSequenceUniqueId` false is, kunt u waarden handmatig instellen:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Antwoordvelden

### TodoCustomField Antwoord (UNIQUE_ID)

| Veld | Type | Beschrijving |
|------|------|--------------|
| `id` | String! | Unieke identificator voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `sequenceId` | Int | Het gegenereerde sequentienummer (gevuld voor UNIQUE_ID-velden) |
| `text` | String | De geformatteerde tekstwaarde (combineert voorvoegsel + gepadded sequentie) |
| `todo` | Todo! | Het record waar deze waarde bij hoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is bijgewerkt |

### CustomField Antwoord (UNIQUE_ID)

| Veld | Type | Beschrijving |
|------|------|--------------|
| `useSequenceUniqueId` | Boolean | Of automatische nummering is ingeschakeld |
| `prefix` | String | Tekstvoorvoegsel voor gegenereerde ID's |
| `sequenceDigits` | Int | Aantal cijfers voor nul-padding |
| `sequenceStartingNumber` | Int | Startnummer voor de reeks |

## Vereiste Machtigingen

| Actie | Vereiste Machtiging |
|-------|---------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Foutreacties

### Fout bij Veldconfiguratie
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Fout bij Machtiging
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Belangrijke Notities

### Automatisch Gegenereerde ID's
- **Alleen-lezen**: Automatisch gegenereerde ID's kunnen niet handmatig worden bewerkt
- **Permanent**: Eenmaal toegewezen, veranderen sequentie-ID's niet
- **Chronologisch**: ID's weerspiegelen de volgorde van aanmaak
- **Afgebakend**: Sequenties zijn onafhankelijk per project

### Prestatieoverwegingen
- ID-generatie voor nieuwe records is synchronisch via database-triggers
- Sequentiegeneratie gebruikt `FOR UPDATE` vergrendelingen voor atomische bewerkingen
- Er bestaat een achtergrondtaaksysteem, maar de implementatie van de worker is in afwachting
- Overweeg startnummers voor sequenties voor projecten met een hoog volume

### Migratie en Updates
- Het toevoegen van automatische nummering aan bestaande records plaatst een achtergrondtaak in de wachtrij (worker in afwachting)
- Wijzigingen in sequentie-instellingen hebben alleen invloed op toekomstige records
- Bestaande ID's blijven onveranderd bij configuratie-updates
- Sequentietellers gaan door vanaf de huidige maximumwaarde

## Beste Praktijken

### Configuratieontwerp
- Kies beschrijvende voorvoegsels die niet in conflict komen met andere systemen
- Gebruik geschikte cijfer-padding voor uw verwachte volume
- Stel redelijke startnummers in om conflicten te vermijden
- Test configuratie met voorbeeldgegevens voordat u deze implementeert

### Voorvoegselrichtlijnen
- Houd voorvoegsels kort en gemakkelijk te onthouden (2-5 tekens)
- Gebruik hoofdletters voor consistentie
- Voeg scheidingstekens (koppelteken, onderstreping) toe voor leesbaarheid
- Vermijd speciale tekens die problemen kunnen veroorzaken in URL's of systemen

### Sequentieplanning
- Schat uw recordvolume om geschikte cijfer-padding te kiezen
- Overweeg toekomstige groei bij het instellen van startnummers
- Plan voor verschillende sequentiebereiken voor verschillende recordtypes
- Documenteer uw ID-schema's voor teamreferentie

## Veelvoorkomende Gebruikscases

1. **Ondersteuningssystemen**
   - Ticketnummers: `TICK-001`, `TICK-002`
   - Zaak-ID's: `CASE-2024-001`
   - Ondersteuningsverzoeken: `SUP-001`

2. **Projectmanagement**
   - Taak-ID's: `TASK-001`, `TASK-002`
   - Sprintitems: `SPRINT-001`
   - Leverbare nummers: `DEL-001`

3. **Bedrijfsvoering**
   - Ordernummers: `ORD-2024-001`
   - Factuur-ID's: `INV-001`
   - Inkooporders: `PO-001`

4. **Kwaliteitsbeheer**
   - Bugrapporten: `BUG-001`
   - Testgeval-ID's: `TEST-001`
   - Beoordelingsnummers: `REV-001`

## Integratiefuncties

### Met Automatiseringen
- Trigger acties wanneer unieke ID's worden toegewezen
- Gebruik ID-patronen in automatiseringsregels
- Verwijs naar ID's in e-mailsjablonen en meldingen

### Met Opzoekingen
- Verwijs naar unieke ID's van andere records
- Vind records op unieke ID
- Toon gerelateerde recordidentificatoren

### Met Rapportage
- Groepeer en filter op ID-patronen
- Volg trends in ID-toewijzing
- Monitor sequentiegebruik en hiaten

## Beperkingen

- **Alleen Sequentieel**: ID's worden toegewezen in chronologische volgorde
- **Geen Hiaten**: Verwijderde records laten hiaten in sequenties
- **Geen Hergebruik**: Sequentienummers worden nooit hergebruikt
- **Projectafgebakend**: Sequenties kunnen niet worden gedeeld tussen projecten
- **Formatbeperkingen**: Beperkte opmaakopties
- **Geen Bulkupdates**: Bestaande sequentie-ID's kunnen niet in bulk worden bijgewerkt
- **Geen Aangepaste Logica**: Aangepaste ID-generatieregels kunnen niet worden geïmplementeerd

## Gerelateerde Bronnen

- [Tekstvelden](/api/custom-fields/text-single) - Voor handmatige tekstidentificatoren
- [Nummervelden](/api/custom-fields/number) - Voor numerieke sequenties
- [Overzicht Aangepaste Velden](/api/custom-fields/2.list-custom-fields) - Algemene concepten
- [Automatiseringen](/api/automations) - Voor ID-gebaseerde automatiseringsregels