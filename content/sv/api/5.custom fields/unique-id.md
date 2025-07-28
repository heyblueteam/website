---
title: Unik ID Anpassat Fält
description: Skapa automatiskt genererade unika identifieringsfält med sekventiell numrering och anpassad formatering
---

Unika ID anpassade fält genererar automatiskt sekventiella, unika identifierare för dina poster. De är perfekta för att skapa biljett nummer, order ID, faktura nummer, eller vilket sekventiellt identifieringssystem som helst i ditt arbetsflöde.

## Grundläggande Exempel

Skapa ett enkelt unikt ID-fält med automatisk sekvensering:

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

## Avancerat Exempel

Skapa ett formaterat unikt ID-fält med prefix och noll-padding:

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

## Inmatningsparametrar

### CreateCustomFieldInput (UNIQUE_ID)

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för det unika ID-fältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `UNIQUE_ID` |
| `description` | String | Nej | Hjälptext som visas för användare |
| `useSequenceUniqueId` | Boolean | Nej | Aktivera automatisk sekvensering (standard: false) |
| `prefix` | String | Nej | Textprefix för genererade ID (t.ex. "TASK-") |
| `sequenceDigits` | Int | Nej | Antal siffror för noll-padding |
| `sequenceStartingNumber` | Int | Nej | Startnummer för sekvensen |

## Konfigurationsalternativ

### Automatisk Sekvensering (`useSequenceUniqueId`)
- **true**: Genererar automatiskt sekventiella ID när poster skapas
- **false** eller **undefined**: Manuell inmatning krävs (fungerar som ett textfält)

### Prefix (`prefix`)
- Valfritt textprefix som läggs till alla genererade ID
- Exempel: "TASK-", "ORD-", "BUG-", "REQ-"
- Ingen längdgräns, men håll rimlig för visning

### Sekvenssiffror (`sequenceDigits`)
- Antal siffror för noll-padding av sekvensnumret
- Exempel: `sequenceDigits: 3` ger `001`, `002`, `003`
- Om inte angivet, tillämpas ingen padding

### Startnummer (`sequenceStartingNumber`)
- Det första numret i sekvensen
- Exempel: `sequenceStartingNumber: 1000` börjar på 1000, 1001, 1002...
- Om inte angivet, börjar det på 1 (standardbeteende)

## Genererat ID Format

Det slutliga ID-formatet kombinerar alla konfigurationsalternativ:

```
{prefix}{paddedSequenceNumber}
```

### Exempel på Format

| Konfiguration | Genererade ID |
|---------------|---------------|
| Inga alternativ | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Läsning av Unika ID Värden

### Fråga Poster med Unika ID
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

### Svar Format
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

## Automatisk ID Generering

### När ID Genereras
- **Postskapande**: ID tilldelas automatiskt när nya poster skapas
- **Fält Tillägg**: När ett UNIQUE_ID-fält läggs till befintliga poster, köas ett bakgrundsjobb (arbetsimplementation pending)
- **Bakgrundsbehandling**: ID-generering för nya poster sker synkront via databasutlösare

### Genereringsprocess
1. **Utlösare**: Ny post skapas eller UNIQUE_ID-fält läggs till
2. **Sekvenssökning**: Systemet hittar nästa tillgängliga sekvensnummer
3. **ID Tilldelning**: Sekvensnummer tilldelas posten
4. **Räknaruppdatering**: Sekvensräknaren ökas för framtida poster
5. **Formatering**: ID formateras med prefix och padding när det visas

### Unikhetsgarantier
- **Databasbegränsningar**: Unik begränsning på sekvens-ID inom varje fält
- **Atomära Operationer**: Sekvensgenerering använder databaslås för att förhindra dubbletter
- **Projektavgränsning**: Sekvenser är oberoende per projekt
- **Skydd mot Tävlingstillstånd**: Samtidiga förfrågningar hanteras säkert

## Manuell vs Automatisk Läge

### Automatisk Läge (`useSequenceUniqueId: true`)
- ID genereras automatiskt via databasutlösare
- Sekventiell numrering garanteras
- Atomär sekvensgenerering förhindrar dubbletter
- Formaterade ID kombinerar prefix + paddat sekvensnummer

### Manuell Läge (`useSequenceUniqueId: false` eller `undefined`)
- Fungerar som ett vanligt textfält
- Användare kan ange anpassade värden via `setTodoCustomField` med `text` parameter
- Ingen automatisk generering
- Ingen unikhetskontroll utöver databasbegränsningar

## Inställning av Manuella Värden (Endast Manuell Läge)

När `useSequenceUniqueId` är false, kan du ställa in värden manuellt:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Svarsfält

### TodoCustomField Svar (UNIQUE_ID)

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `sequenceId` | Int | Det genererade sekvensnumret (fylls i för UNIQUE_ID-fält) |
| `text` | String | Det formaterade textvärdet (kombinerar prefix + paddad sekvens) |
| `todo` | Todo! | Den post detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast uppdaterades |

### CustomField Svar (UNIQUE_ID)

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | Om automatisk sekvensering är aktiverad |
| `prefix` | String | Textprefix för genererade ID |
| `sequenceDigits` | Int | Antal siffror för noll-padding |
| `sequenceStartingNumber` | Int | Startnummer för sekvensen |

## Obligatoriska Behörigheter

| Åtgärd | Obligatorisk Behörighet |
|--------|-----------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Fel Svar

### Fältkonfigurationsfel
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

### Behörighetsfel
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

## Viktiga Anteckningar

### Automatiskt Genererade ID
- **Skrivskyddade**: Automatiskt genererade ID kan inte redigeras manuellt
- **Permanent**: När de har tilldelats, ändras sekvens-ID inte
- **Kronologiska**: ID återspeglar skapelseordning
- **Avgränsade**: Sekvenser är oberoende per projekt

### Prestandaöverväganden
- ID-generering för nya poster är synkron via databasutlösare
- Sekvensgenerering använder `FOR UPDATE` lås för atomära operationer
- Bakgrundsjobbssystem finns men arbetsimplementation är pending
- Tänk på sekvensstartnummer för projekt med hög volym

### Migration och Uppdateringar
- Tillägg av automatisk sekvensering till befintliga poster köar bakgrundsjobb (arbetspending)
- Ändring av sekvensinställningar påverkar endast framtida poster
- Befintliga ID förblir oförändrade när konfiguration uppdateras
- Sekvensräknare fortsätter från nuvarande maximum

## Bästa Praxis

### Konfigurationsdesign
- Välj beskrivande prefix som inte kommer i konflikt med andra system
- Använd lämplig siffra padding för din förväntade volym
- Ställ in rimliga startnummer för att undvika konflikter
- Testa konfigurationen med exempeldata innan distribution

### Prefixriktlinjer
- Håll prefix korta och minnesvärda (2-5 tecken)
- Använd versaler för konsekvens
- Inkludera avgränsare (bindestreck, understreck) för läsbarhet
- Undvik specialtecken som kan orsaka problem i URL:er eller system

### Sekvensplanering
- Uppskatta din postvolym för att välja lämplig siffra padding
- Tänk på framtida tillväxt när du ställer in startnummer
- Planera för olika sekvensintervall för olika posttyper
- Dokumentera dina ID-scheman för teamreferens

## Vanliga Användningsfall

1. **Support System**
   - Biljettnummer: `TICK-001`, `TICK-002`
   - Ärende-ID: `CASE-2024-001`
   - Supportförfrågningar: `SUP-001`

2. **Projektledning**
   - Uppgifts-ID: `TASK-001`, `TASK-002`
   - Sprintobjekt: `SPRINT-001`
   - Leveransnummer: `DEL-001`

3. **Affärsverksamhet**
   - Ordernummer: `ORD-2024-001`
   - Faktura-ID: `INV-001`
   - Inköpsorder: `PO-001`

4. **Kvalitetshantering**
   - Buggrapporter: `BUG-001`
   - Testfall-ID: `TEST-001`
   - Granskningsnummer: `REV-001`

## Integrationsfunktioner

### Med Automatiseringar
- Utlös åtgärder när unika ID tilldelas
- Använd ID-mönster i automatiseringsregler
- Referera till ID i e-postmallar och meddelanden

### Med Uppslag
- Referera till unika ID från andra poster
- Hitta poster efter unikt ID
- Visa relaterade postidentifierare

### Med Rapportering
- Gruppera och filtrera efter ID-mönster
- Spåra trender i ID-tilldelning
- Övervaka sekvensanvändning och luckor

## Begränsningar

- **Endast Sekventiella**: ID tilldelas i kronologisk ordning
- **Inga Luckor**: Raderade poster lämnar luckor i sekvenser
- **Ingen Återanvändning**: Sekvensnummer återanvänds aldrig
- **Projektavgränsade**: Kan inte dela sekvenser över projekt
- **Formatbegränsningar**: Begränsade formateringsalternativ
- **Inga Massuppdateringar**: Kan inte massuppdatera befintliga sekvens-ID
- **Ingen Anpassad Logik**: Kan inte implementera anpassade ID-genereringsregler

## Relaterade Resurser

- [Textfält](/api/custom-fields/text-single) - För manuella textidentifierare
- [Nummerfält](/api/custom-fields/number) - För numeriska sekvenser
- [Översikt över Anpassade Fält](/api/custom-fields/2.list-custom-fields) - Allmänna koncept
- [Automatiseringar](/api/automations) - För ID-baserade automatiseringsregler