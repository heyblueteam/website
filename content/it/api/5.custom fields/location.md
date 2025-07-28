---
title: Campo Personalizzato di Posizione
description: Crea campi di posizione per memorizzare coordinate geografiche per i record
---

I campi personalizzati di posizione memorizzano coordinate geografiche (latitudine e longitudine) per i record. Supportano la memorizzazione precisa delle coordinate, le query geospaziali e il filtraggio basato sulla posizione in modo efficiente.

## Esempio di Base

Crea un semplice campo di posizione:

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

## Esempio Avanzato

Crea un campo di posizione con descrizione:

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

## Parametri di Input

### CreateCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sì | Nome visualizzato del campo di posizione |
| `type` | CustomFieldType! | ✅ Sì | Deve essere `LOCATION` |
| `description` | String | No | Testo di aiuto mostrato agli utenti |

## Impostazione dei Valori di Posizione

I campi di posizione memorizzano coordinate di latitudine e longitudine:

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

### Parametri di SetTodoCustomFieldInput

| Parametro | Tipo | Richiesto | Descrizione |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sì | ID del record da aggiornare |
| `customFieldId` | String! | ✅ Sì | ID del campo personalizzato di posizione |
| `latitude` | Float | No | Coordinata di latitudine (-90 a 90) |
| `longitude` | Float | No | Coordinata di longitudine (-180 a 180) |

**Nota**: Sebbene entrambi i parametri siano facoltativi nello schema, entrambe le coordinate sono necessarie per una posizione valida. Se ne viene fornita solo una, la posizione sarà non valida.

## Validazione delle Coordinate

### Intervalli Validi

| Coordinata | Intervallo | Descrizione |
|------------|------------|-------------|
| Latitude | -90 to 90 | Posizione Nord/Sud |
| Longitude | -180 to 180 | Posizione Est/Ovest |

### Coordinate di Esempio

| Posizione | Latitudine | Longitudine |
|-----------|------------|-------------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Creazione di Record con Valori di Posizione

Quando si crea un nuovo record con dati di posizione:

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

### Formato di Input per la Creazione

Quando si creano record, i valori di posizione utilizzano un formato separato da virgole:

| Formato | Esempio | Descrizione |
|---------|---------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Formato standard delle coordinate |
| `"51.5074,-0.1278"` | London coordinates | Nessuno spazio attorno alla virgola |
| `"-33.8688,151.2093"` | Sydney coordinates | Valori negativi consentiti |

## Campi di Risposta

### TodoCustomField Risposta

| Campo | Tipo | Descrizione |
|-------|------|-------------|
| `id` | String! | Identificatore unico per il valore del campo |
| `customField` | CustomField! | La definizione del campo personalizzato |
| `latitude` | Float | Coordinata di latitudine |
| `longitude` | Float | Coordinata di longitudine |
| `todo` | Todo! | Il record a cui appartiene questo valore |
| `createdAt` | DateTime! | Quando è stato creato il valore |
| `updatedAt` | DateTime! | Quando è stato modificato per l'ultima volta il valore |

## Limitazioni Importanti

### Nessuna Geocodifica Integrata

I campi di posizione memorizzano solo coordinate - non includono:
- Conversione indirizzo-coordinata
- Geocodifica inversa (coordinata-indirizzo)
- Validazione o ricerca degli indirizzi
- Integrazione con servizi di mappatura
- Ricerca di nomi di luoghi

### Servizi Esterni Richiesti

Per la funzionalità degli indirizzi, è necessario integrare servizi esterni:
- **Google Maps API** per la geocodifica
- **OpenStreetMap Nominatim** per geocodifica gratuita
- **MapBox** per mappatura e geocodifica
- **Here API** per servizi di localizzazione

### Esempio di Integrazione

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

## Permessi Richiesti

| Azione | Ruolo Richiesto |
|--------|------------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Risposte di Errore

### Coordinate Non Valide
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

### Longitudine Non Valida
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

## Migliori Pratiche

### Raccolta Dati
- Utilizzare coordinate GPS per posizioni precise
- Validare le coordinate prima di memorizzarle
- Considerare le esigenze di precisione delle coordinate (6 decimali ≈ 10 cm di accuratezza)
- Memorizzare le coordinate in gradi decimali (non in gradi/minuti/secondi)

### Esperienza Utente
- Fornire interfacce mappa per la selezione delle coordinate
- Mostrare anteprime della posizione quando si visualizzano le coordinate
- Validare le coordinate lato client prima delle chiamate API
- Considerare le implicazioni del fuso orario per i dati di posizione

### Prestazioni
- Utilizzare indici spaziali per query efficienti
- Limitare la precisione delle coordinate all'accuratezza necessaria
- Considerare la memorizzazione nella cache per posizioni frequentemente accessibili
- Aggiornare le posizioni in batch quando possibile

## Casi d'Uso Comuni

1. **Operazioni sul Campo**
   - Posizioni delle attrezzature
   - Indirizzi delle chiamate di servizio
   - Siti di ispezione
   - Posizioni di consegna

2. **Gestione Eventi**
   - Luoghi degli eventi
   - Posizioni degli incontri
   - Siti delle conferenze
   - Luoghi dei workshop

3. **Tracciamento delle Risorse**
   - Posizioni delle attrezzature
   - Posizioni delle strutture
   - Tracciamento dei veicoli
   - Posizioni dell'inventario

4. **Analisi Geografica**
   - Aree di copertura del servizio
   - Distribuzione dei clienti
   - Analisi di mercato
   - Gestione del territorio

## Funzionalità di Integrazione

### Con Ricerche
- Riferire i dati di posizione da altri record
- Trovare record per prossimità geografica
- Aggregare dati basati sulla posizione
- Incrociare coordinate

### Con Automazioni
- Attivare azioni basate su cambiamenti di posizione
- Creare notifiche geofence
- Aggiornare record correlati quando cambiano le posizioni
- Generare report basati sulla posizione

### Con Formule
- Calcolare distanze tra posizioni
- Determinare centri geografici
- Analizzare modelli di posizione
- Creare metriche basate sulla posizione

## Limitazioni

- Nessuna geocodifica integrata o conversione degli indirizzi
- Nessuna interfaccia di mappatura fornita
- Richiede servizi esterni per la funzionalità degli indirizzi
- Limitato solo alla memorizzazione delle coordinate
- Nessuna validazione automatica della posizione oltre il controllo dell'intervallo

## Risorse Correlate

- [Panoramica dei Campi Personalizzati](/api/custom-fields/list-custom-fields) - Concetti generali
- [Google Maps API](https://developers.google.com/maps) - Servizi di geocodifica
- [OpenStreetMap Nominatim](https://nominatim.org/) - Geocodifica gratuita
- [MapBox API](https://docs.mapbox.com/) - Servizi di mappatura e geocodifica