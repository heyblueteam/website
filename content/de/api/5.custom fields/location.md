---
title: Standortbenutzerdefiniertes Feld
description: Erstellen Sie Standortfelder, um geografische Koordinaten für Datensätze zu speichern
---

Standortbenutzerdefinierte Felder speichern geografische Koordinaten (Breiten- und Längengrad) für Datensätze. Sie unterstützen die präzise Speicherung von Koordinaten, raumbezogene Abfragen und effizientes standortbasiertes Filtern.

## Einfaches Beispiel

Erstellen Sie ein einfaches Standortfeld:

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

## Fortgeschrittenes Beispiel

Erstellen Sie ein Standortfeld mit Beschreibung:

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

## Eingabeparameter

### CreateCustomFieldInput

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Ja | Anzeigename des Standortfelds |
| `type` | CustomFieldType! | ✅ Ja | Muss `LOCATION` sein |
| `description` | String | Nein | Hilfetext, der den Benutzern angezeigt wird |

## Standortwerte festlegen

Standortfelder speichern Breiten- und Längengradkoordinaten:

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

### SetTodoCustomFieldInput Parameter

| Parameter | Typ | Erforderlich | Beschreibung |
|-----------|------|--------------|-------------|
| `todoId` | String! | ✅ Ja | ID des zu aktualisierenden Datensatzes |
| `customFieldId` | String! | ✅ Ja | ID des benutzerdefinierten Standortfelds |
| `latitude` | Float | Nein | Breitengradkoordinate (-90 bis 90) |
| `longitude` | Float | Nein | Längengradkoordinate (-180 bis 180) |

**Hinweis**: Während beide Parameter im Schema optional sind, sind beide Koordinaten für einen gültigen Standort erforderlich. Wenn nur eine bereitgestellt wird, ist der Standort ungültig.

## Koordinatenvalidierung

### Gültige Bereiche

| Koordinate | Bereich | Beschreibung |
|------------|---------|-------------|
| Latitude | -90 to 90 | Nord/Süd-Position |
| Longitude | -180 to 180 | Ost/West-Position |

### Beispielkoordinaten

| Standort | Breitengrad | Längengrad |
|----------|-------------|------------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Erstellen von Datensätzen mit Standortwerten

Beim Erstellen eines neuen Datensatzes mit Standortdaten:

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

### Eingabeformat für die Erstellung

Beim Erstellen von Datensätzen verwenden Standortwerte ein durch Kommas getrenntes Format:

| Format | Beispiel | Beschreibung |
|--------|----------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Standardkoordinatenformat |
| `"51.5074,-0.1278"` | London coordinates | Keine Leerzeichen um das Komma |
| `"-33.8688,151.2093"` | Sydney coordinates | Negative Werte erlaubt |

## Antwortfelder

### TodoCustomField Antwort

| Feld | Typ | Beschreibung |
|------|-----|-------------|
| `id` | String! | Eindeutige Kennung für den Feldwert |
| `customField` | CustomField! | Die Definition des benutzerdefinierten Felds |
| `latitude` | Float | Breitengradkoordinate |
| `longitude` | Float | Längengradkoordinate |
| `todo` | Todo! | Der Datensatz, zu dem dieser Wert gehört |
| `createdAt` | DateTime! | Wann der Wert erstellt wurde |
| `updatedAt` | DateTime! | Wann der Wert zuletzt geändert wurde |

## Wichtige Einschränkungen

### Keine integrierte Geokodierung

Standortfelder speichern nur Koordinaten - sie beinhalten **nicht**:
- Adress-zu-Koordinaten-Konvertierung
- Rückwärtsgeokodierung (Koordinaten-zu-Adresse)
- Adressvalidierung oder -suche
- Integration mit Kartendiensten
- Ortsnamenabfrage

### Externe Dienste erforderlich

Für Adressfunktionen müssen Sie externe Dienste integrieren:
- **Google Maps API** für Geokodierung
- **OpenStreetMap Nominatim** für kostenlose Geokodierung
- **MapBox** für Karten- und Geokodierungsdienste
- **Here API** für Standortdienste

### Beispielintegration

```javascript
// Client-side geocoding example (not part of Blue API)
async function geocodeAddress(address) {
  const response = await fetch(
    `&key=${API_KEY}`
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

## Erforderliche Berechtigungen

| Aktion | Erforderliche Rolle |
|--------|---------------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Fehlermeldungen

### Ungültige Koordinaten
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

### Ungültiger Längengrad
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

## Beste Praktiken

### Datensammlung
- Verwenden Sie GPS-Koordinaten für präzise Standorte
- Validieren Sie Koordinaten vor der Speicherung
- Berücksichtigen Sie die Anforderungen an die Koordinatenpräzision (6 Dezimalstellen ≈ 10 cm Genauigkeit)
- Speichern Sie Koordinaten in Dezimalgraden (nicht in Grad/Minuten/Sekunden)

### Benutzererfahrung
- Stellen Sie Kartenoberflächen für die Koordinatenauswahl bereit
- Zeigen Sie Standortvorschauen beim Anzeigen von Koordinaten an
- Validieren Sie Koordinaten clientseitig vor API-Aufrufen
- Berücksichtigen Sie die Zeitzonenauswirkungen auf Standortdaten

### Leistung
- Verwenden Sie räumliche Indizes für effiziente Abfragen
- Begrenzen Sie die Koordinatenpräzision auf die erforderliche Genauigkeit
- Berücksichtigen Sie das Caching für häufig abgerufene Standorte
- Batch-Updates von Standorten, wenn möglich

## Häufige Anwendungsfälle

1. **Feldoperationen**
   - Standort von Geräten
   - Adressen von Serviceanrufen
   - Inspektionsstandorte
   - Lieferstandorte

2. **Veranstaltungsmanagement**
   - Veranstaltungsorte
   - Besprechungsstandorte
   - Konferenzstandorte
   - Workshopstandorte

3. **Asset-Tracking**
   - Positionen von Geräten
   - Standorte von Einrichtungen
   - Fahrzeugverfolgung
   - Lagerstandorte

4. **Geografische Analyse**
   - Serviceabdeckungsgebiete
   - Kundenverteilung
   - Marktanalyse
   - Gebietsmanagement

## Integrationsfunktionen

### Mit Nachschlägen
- Verweisen Sie auf Standortdaten aus anderen Datensätzen
- Finden Sie Datensätze nach geografischer Nähe
- Aggregieren Sie standortbasierte Daten
- Querverweisen von Koordinaten

### Mit Automatisierungen
- Auslösen von Aktionen basierend auf Standortänderungen
- Erstellen von geofenced Benachrichtigungen
- Aktualisieren von verwandten Datensätzen, wenn sich Standorte ändern
- Generieren von standortbasierten Berichten

### Mit Formeln
- Berechnen von Entfernungen zwischen Standorten
- Bestimmen geografischer Zentren
- Analysieren von Standortmustern
- Erstellen von standortbasierten Metriken

## Einschränkungen

- Keine integrierte Geokodierung oder Adresskonvertierung
- Keine bereitgestellte Kartenoberfläche
- Erfordert externe Dienste für Adressfunktionen
- Beschränkt auf die Speicherung von Koordinaten
- Keine automatische Standortvalidierung über die Bereichsprüfung hinaus

## Verwandte Ressourcen

- [Übersicht über benutzerdefinierte Felder](/api/custom-fields/list-custom-fields) - Allgemeine Konzepte
- [Google Maps API](https://developers.google.com/maps) - Geokodierungsdienste
- [OpenStreetMap Nominatim](https://nominatim.org/) - Kostenlose Geokodierung
- [MapBox API](https://docs.mapbox.com/) - Karten- und Geokodierungsdienste