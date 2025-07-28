---
title: Locatie Aangepast Veld
description: Maak locatievelden om geografische coördinaten voor records op te slaan
---

Locatie aangepaste velden slaan geografische coördinaten (breedte- en lengtegraad) voor records op. Ze ondersteunen nauwkeurige coördinatenopslag, geospatiale queries en efficiënte locatiegebaseerde filtering.

## Basisvoorbeeld

Maak een eenvoudig locatieveld:

```graphql
mutation CreateLocationField {
  createCustomField(input: {
    name: "Meeting Location"
    type: LOCATION
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Geavanceerd Voorbeeld

Maak een locatieveld met beschrijving:

```graphql
mutation CreateDetailedLocationField {
  createCustomField(input: {
    name: "Office Location"
    type: LOCATION
    projectId: "proj_123"
    description: "Primary office location coordinates"
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
|-----------|------|----------|-------------|
| `name` | String! | ✅ Ja | Weergavenaam van het locatieveld |
| `type` | CustomFieldType! | ✅ Ja | Moet zijn `LOCATION` |
| `description` | String | Nee | Hulptekst die aan gebruikers wordt getoond |

## Instellen van Locatiewaarden

Locatievelden slaan breedte- en lengtegraadcoördinaten op:

```graphql
mutation SetLocationValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    latitude: 40.7128
    longitude: -74.0060
  })
}
```

### SetTodoCustomFieldInput Parameters

| Parameter | Type | Vereist | Beschrijving |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Ja | ID van het record dat moet worden bijgewerkt |
| `customFieldId` | String! | ✅ Ja | ID van het locatie aangepaste veld |
| `latitude` | Float | Nee | Breedtegraadcoördinaat (-90 tot 90) |
| `longitude` | Float | Nee | Lengtegraadcoördinaat (-180 tot 180) |

**Opmerking**: Hoewel beide parameters optioneel zijn in het schema, zijn beide coördinaten vereist voor een geldige locatie. Als er slechts één wordt opgegeven, is de locatie ongeldig.

## Coördinatenvalidatie

### Geldige Bereiken

| Coördinaat | Bereik | Beschrijving |
|------------|-------|-------------|
| Latitude | -90 to 90 | Noord/Zuid positie |
| Longitude | -180 to 180 | Oost/West positie |

### Voorbeeld Coördinaten

| Locatie | Breedtegraad | Lengtegraad |
|----------|----------|-----------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Records Maken met Locatiewaarden

Bij het maken van een nieuw record met locatiegegevens:

```graphql
mutation CreateRecordWithLocation {
  createTodo(input: {
    title: "Site Visit"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "location_field_id"
      value: "40.7128,-74.0060"  # Format: "latitude,longitude"
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
      latitude
      longitude
    }
  }
}
```

### Invoervormaat voor Creatie

Bij het maken van records gebruiken locatiewaarden een komma-gescheiden formaat:

| Formaat | Voorbeeld | Beschrijving |
|--------|---------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Standaard coördinatenformaat |
| `"51.5074,-0.1278"` | London coordinates | Geen spaties rond de komma |
| `"-33.8688,151.2093"` | Sydney coordinates | Negatieve waarden toegestaan |

## Responsvelden

### TodoCustomField Respons

| Veld | Type | Beschrijving |
|-------|------|-------------|
| `id` | String! | Unieke identificatie voor de veldwaarde |
| `customField` | CustomField! | De definitie van het aangepaste veld |
| `latitude` | Float | Breedtegraadcoördinaat |
| `longitude` | Float | Lengtegraadcoördinaat |
| `todo` | Todo! | Het record waartoe deze waarde behoort |
| `createdAt` | DateTime! | Wanneer de waarde is aangemaakt |
| `updatedAt` | DateTime! | Wanneer de waarde voor het laatst is gewijzigd |

## Belangrijke Beperkingen

### Geen Ingebouwde Geocodering

Locatievelden slaan alleen coördinaten op - ze bevatten **geen**:
- Adres-naar-coördinaten conversie
- Omgekeerde geocodering (coördinaten-naar-adres)
- Adresvalidatie of -zoekopdracht
- Integratie met kaartdiensten
- Plaatsnaam opzoeken

### Externe Diensten Vereist

Voor adresfunctionaliteit moet je externe diensten integreren:
- **Google Maps API** voor geocodering
- **OpenStreetMap Nominatim** voor gratis geocodering
- **MapBox** voor kaarten en geocodering
- **Here API** voor locatie-diensten

### Voorbeeldintegratie

```javascript
// Client-side geocoding example (not part of Blue API)
async function geocodeAddress(address) {
  const response = await fetch(
    `https://maps.googleapis.com/maps/api/geocode/json?address=${encodeURIComponent(address)}&key=${API_KEY}`
  );
  const data = await response.json();
  
  if (data.results.length > 0) {
    const { lat, lng } = data.results[0].geometry.location;
    
    // Now set the location field in Blue
    await setTodoCustomField({
      todoId: "todo_123",
      customFieldId: "location_field_456",
      latitude: lat,
      longitude: lng
    });
  }
}
```

## Vereiste Machtigingen

| Actie | Vereiste Rol |
|--------|---------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Foutreacties

### Ongeldige Coördinaten
```json
{
  "errors": [{
    "message": "Invalid coordinates: latitude must be between -90 and 90",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Ongeldige Lengtegraad
```json
{
  "errors": [{
    "message": "Invalid coordinates: longitude must be between -180 and 180",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Beste Praktijken

### Gegevensverzameling
- Gebruik GPS-coördinaten voor nauwkeurige locaties
- Valideer coördinaten voordat je ze opslaat
- Houd rekening met de precisiebehoeften van coördinaten (6 decimalen ≈ 10 cm nauwkeurigheid)
- Sla coördinaten op in decimale graden (niet in graden/minuten/seconden)

### Gebruikerservaring
- Bied kaartinterfaces voor coördinatenselectie
- Toon locatievoorbeelden bij het weergeven van coördinaten
- Valideer coördinaten aan de clientzijde voordat je API-aanroepen doet
- Houd rekening met tijdzone-implicaties voor locatiegegevens

### Prestaties
- Gebruik ruimtelijke indexen voor efficiënte queries
- Beperk de precisie van coördinaten tot de benodigde nauwkeurigheid
- Overweeg caching voor vaak geraadpleegde locaties
- Groepeer locatie-updates waar mogelijk

## Veelvoorkomende Gebruiksscenario's

1. **Veldoperaties**
   - Apparatuurlocaties
   - Service-aanroepadressen
   - Inspectieplaatsen
   - Leveringslocaties

2. **Evenementbeheer**
   - Evenementlocaties
   - Vergaderlocaties
   - Conferentieplaatsen
   - Werkplaatslocaties

3. **Activa Tracking**
   - Apparatuurposities
   - Faciliteitlocaties
   - Voertuigtracking
   - Voorraadlocaties

4. **Geografische Analyse**
   - Service-dekking gebieden
   - Klantverdeling
   - Marktanalyse
   - Territoriumbeheer

## Integratiefuncties

### Met Opzoekingen
- Verwijs naar locatiegegevens van andere records
- Vind records op geografische nabijheid
- Aggregatie van locatiegebaseerde gegevens
- Coördinaten kruisverwijzen

### Met Automatiseringen
- Trigger acties op basis van locatieveranderingen
- Maak geofenced meldingen
- Werk gerelateerde records bij wanneer locaties veranderen
- Genereer locatiegebaseerde rapporten

### Met Formules
- Bereken afstanden tussen locaties
- Bepaal geografische centra
- Analyseer locatiepatronen
- Maak locatiegebaseerde statistieken

## Beperkingen

- Geen ingebouwde geocodering of adresconversie
- Geen kaartinterface voorzien
- Vereist externe diensten voor adresfunctionaliteit
- Beperkt tot alleen coördinatenopslag
- Geen automatische locatievalidatie buiten bereikcontroles

## Gerelateerde Bronnen

- [Overzicht Aangepaste Velden](/api/custom-fields/list-custom-fields) - Algemene concepten
- [Google Maps API](https://developers.google.com/maps) - Geocoderingdiensten
- [OpenStreetMap Nominatim](https://nominatim.org/) - Gratis geocodering
- [MapBox API](https://docs.mapbox.com/) - Kaart- en geocoderingdiensten