---
title: Anpassat fält för plats
description: Skapa platsfält för att lagra geografiska koordinater för poster
---

Anpassade platsfält lagrar geografiska koordinater (latitud och longitud) för poster. De stödjer exakt lagring av koordinater, geospatiala frågor och effektiv filtrering baserat på plats.

## Grundläggande exempel

Skapa ett enkelt platsfält:

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

## Avancerat exempel

Skapa ett platsfält med beskrivning:

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

## Indata parametrar

### CreateCustomFieldInput

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Visningsnamn för platsfältet |
| `type` | CustomFieldType! | ✅ Ja | Måste vara `LOCATION` |
| `description` | String | Nej | Hjälptext som visas för användare |

## Ställa in platsvärden

Platsfält lagrar latitud- och longitudkoordinater:

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

### SetTodoCustomFieldInput parametrar

| Parameter | Typ | Obligatorisk | Beskrivning |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID för posten som ska uppdateras |
| `customFieldId` | String! | ✅ Ja | ID för det anpassade platsfältet |
| `latitude` | Float | Nej | Latitudkoordinat (-90 till 90) |
| `longitude` | Float | Nej | Longitudkoordinat (-180 till 180) |

**Notera**: Även om båda parametrarna är valfria i schemat, krävs båda koordinaterna för en giltig plats. Om endast en tillhandahålls, kommer platsen att vara ogiltig.

## Koordinatvalidering

### Giltiga intervall

| Koordinat | Intervall | Beskrivning |
|------------|-----------|-------------|
| Latitude | -90 to 90 | Nord/Syd position |
| Longitude | -180 to 180 | Öst/Väst position |

### Exempelkoordinater

| Plats | Latitud | Longitud |
|----------|----------|-----------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Skapa poster med platsvärden

När du skapar en ny post med platsdata:

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

### Indataformat för skapande

När du skapar poster använder platsvärden ett kommaseparerat format:

| Format | Exempel | Beskrivning |
|--------|---------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Standardkoordinatformat |
| `"51.5074,-0.1278"` | London coordinates | Inga mellanslag runt kommatecknet |
| `"-33.8688,151.2093"` | Sydney coordinates | Negativa värden tillåtna |

## Svarsfält

### TodoCustomField Svar

| Fält | Typ | Beskrivning |
|-------|------|-------------|
| `id` | String! | Unik identifierare för fältvärdet |
| `customField` | CustomField! | Den anpassade fältdefinitionen |
| `latitude` | Float | Latitudkoordinat |
| `longitude` | Float | Longitudkoordinat |
| `todo` | Todo! | Den post som detta värde tillhör |
| `createdAt` | DateTime! | När värdet skapades |
| `updatedAt` | DateTime! | När värdet senast ändrades |

## Viktiga begränsningar

### Ingen inbyggd geokodning

Platsfält lagrar endast koordinater - de inkluderar **inte**:
- Adress-till-koordinater konvertering
- Omvänd geokodning (koordinater-till-adress)
- Adressvalidering eller sökning
- Integration med karttjänster
- Platsnamnssökning

### Externa tjänster krävs

För adressfunktionalitet måste du integrera externa tjänster:
- **Google Maps API** för geokodning
- **OpenStreetMap Nominatim** för gratis geokodning
- **MapBox** för kartläggning och geokodning
- **Here API** för plats tjänster

### Exempel på integration

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

## Obligatoriska behörigheter

| Åtgärd | Obligatorisk roll |
|--------|-------------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Felrespons

### Ogiltiga koordinater
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

### Ogiltig longitud
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

## Bästa praxis

### Datainsamling
- Använd GPS-koordinater för exakta platser
- Validera koordinater innan de lagras
- Tänk på behovet av koordinatprecision (6 decimaler ≈ 10 cm noggrannhet)
- Lagra koordinater i decimalgrader (inte grader/minuter/sekunder)

### Användarupplevelse
- Tillhandahåll kartgränssnitt för koordinatval
- Visa platsförhandsvisningar när du visar koordinater
- Validera koordinater på klientsidan innan API-anrop
- Tänk på tidszonsimplikationer för platsdata

### Prestanda
- Använd rumsliga index för effektiva frågor
- Begränsa koordinatprecisionen till nödvändig noggrannhet
- Tänk på caching för ofta åtkomna platser
- Batcha platsuppdateringar när det är möjligt

## Vanliga användningsfall

1. **Fältoperationer**
   - Utrustningsplatser
   - Serviceanropsadresser
   - Inspektionsplatser
   - Leveransplatser

2. **Evenemangshantering**
   - Evenemangsplatser
   - Mötesplatser
   - Konferensplatser
   - Workshopplatser

3. **Tillgångsspårning**
   - Utrustningspositioner
   - Anläggningsplatser
   - Fordonsspårning
   - Lagerplatser

4. **Geografisk analys**
   - Tjänsteområden
   - Kunddistribution
   - Marknadsanalys
   - Territoriehantering

## Integrationsfunktioner

### Med uppslag
- Referera platsdata från andra poster
- Hitta poster efter geografisk närhet
- Aggregat platsbaserad data
- Korsreferenskoordinater

### Med automatiseringar
- Utlös åtgärder baserat på platsändringar
- Skapa geofenced-notifikationer
- Uppdatera relaterade poster när platser ändras
- Generera platsbaserade rapporter

### Med formler
- Beräkna avstånd mellan platser
- Bestäm geografiska centra
- Analysera platsmönster
- Skapa platsbaserade mätvärden

## Begränsningar

- Ingen inbyggd geokodning eller adresskonvertering
- Ingen kartgränssnitt tillhandahålls
- Kräver externa tjänster för adressfunktionalitet
- Begränsat till lagring av koordinater endast
- Ingen automatisk platsvalidering utöver intervalls kontroll

## Relaterade resurser

- [Översikt över anpassade fält](/api/custom-fields/list-custom-fields) - Allmänna koncept
- [Google Maps API](https://developers.google.com/maps) - Geokodningstjänster
- [OpenStreetMap Nominatim](https://nominatim.org/) - Gratis geokodning
- [MapBox API](https://docs.mapbox.com/) - Kartläggning och geokodningstjänster