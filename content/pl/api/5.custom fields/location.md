---
title: Niestandardowe pole lokalizacji
description: Twórz pola lokalizacji, aby przechowywać współrzędne geograficzne dla rekordów
---

Niestandardowe pola lokalizacji przechowują współrzędne geograficzne (szerokość i długość geograficzną) dla rekordów. Obsługują precyzyjne przechowywanie współrzędnych, zapytania geospołeczne oraz efektywne filtrowanie oparte na lokalizacji.

## Podstawowy przykład

Utwórz proste pole lokalizacji:

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

## Zaawansowany przykład

Utwórz pole lokalizacji z opisem:

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

## Parametry wejściowe

### CreateCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `name` | String! | ✅ Tak | Nazwa wyświetlana pola lokalizacji |
| `type` | CustomFieldType! | ✅ Tak | Musi być `LOCATION` |
| `description` | String | Nie | Tekst pomocy wyświetlany użytkownikom |

## Ustawianie wartości lokalizacji

Pola lokalizacji przechowują współrzędne szerokości i długości geograficznej:

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

### Parametry SetTodoCustomFieldInput

| Parametr | Typ | Wymagany | Opis |
|----------|-----|----------|------|
| `todoId` | String! | ✅ Tak | ID rekordu do zaktualizowania |
| `customFieldId` | String! | ✅ Tak | ID niestandardowego pola lokalizacji |
| `latitude` | Float | Nie | Współrzędna szerokości (-90 do 90) |
| `longitude` | Float | Nie | Współrzędna długości (-180 do 180) |

**Uwaga**: Chociaż oba parametry są opcjonalne w schemacie, obie współrzędne są wymagane dla ważnej lokalizacji. Jeśli podano tylko jedną, lokalizacja będzie nieważna.

## Walidacja współrzędnych

### Ważne zakresy

| Współrzędna | Zakres | Opis |
|-------------|--------|------|
| Latitude | -90 to 90 | Pozycja północna/południowa |
| Longitude | -180 to 180 | Pozycja wschodnia/zachodnia |

### Przykładowe współrzędne

| Lokalizacja | Szerokość | Długość |
|-------------|-----------|---------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Tworzenie rekordów z wartościami lokalizacji

Podczas tworzenia nowego rekordu z danymi lokalizacji:

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

### Format wejściowy do tworzenia

Podczas tworzenia rekordów, wartości lokalizacji używają formatu oddzielonego przecinkami:

| Format | Przykład | Opis |
|--------|----------|------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Standardowy format współrzędnych |
| `"51.5074,-0.1278"` | London coordinates | Brak spacji wokół przecinka |
| `"-33.8688,151.2093"` | Sydney coordinates | Dozwolone wartości ujemne |

## Pola odpowiedzi

### TodoCustomField Response

| Pole | Typ | Opis |
|------|-----|------|
| `id` | String! | Unikalny identyfikator wartości pola |
| `customField` | CustomField! | Definicja niestandardowego pola |
| `latitude` | Float | Współrzędna szerokości |
| `longitude` | Float | Współrzędna długości |
| `todo` | Todo! | Rekord, do którego należy ta wartość |
| `createdAt` | DateTime! | Kiedy wartość została utworzona |
| `updatedAt` | DateTime! | Kiedy wartość została ostatnio zmodyfikowana |

## Ważne ograniczenia

### Brak wbudowanego geokodowania

Pola lokalizacji przechowują tylko współrzędne - nie zawierają:
- Konwersji adresu na współrzędne
- Geokodowania odwrotnego (współrzędne na adres)
- Walidacji lub wyszukiwania adresów
- Integracji z usługami mapowymi
- Wyszukiwania nazw miejsc

### Wymagane usługi zewnętrzne

Aby uzyskać funkcjonalność adresu, będziesz musiał zintegrować usługi zewnętrzne:
- **Google Maps API** do geokodowania
- **OpenStreetMap Nominatim** do bezpłatnego geokodowania
- **MapBox** do mapowania i geokodowania
- **Here API** do usług lokalizacyjnych

### Przykład integracji

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

## Wymagane uprawnienia

| Akcja | Wymagana rola |
|-------|---------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Odpowiedzi błędów

### Nieprawidłowe współrzędne
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

### Nieprawidłowa długość
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

## Najlepsze praktyki

### Zbieranie danych
- Używaj współrzędnych GPS dla precyzyjnych lokalizacji
- Waliduj współrzędne przed przechowywaniem
- Rozważ potrzeby precyzji współrzędnych (6 miejsc po przecinku ≈ 10 cm dokładności)
- Przechowuj współrzędne w stopniach dziesiętnych (nie w stopniach/minutach/sekundach)

### Doświadczenie użytkownika
- Zapewnij interfejsy mapowe do wyboru współrzędnych
- Wyświetlaj podglądy lokalizacji podczas wyświetlania współrzędnych
- Waliduj współrzędne po stronie klienta przed wywołaniami API
- Rozważ implikacje strefy czasowej dla danych lokalizacyjnych

### Wydajność
- Używaj indeksów przestrzennych do efektywnych zapytań
- Ogranicz precyzję współrzędnych do potrzebnej dokładności
- Rozważ buforowanie dla często używanych lokalizacji
- Grupuj aktualizacje lokalizacji, gdy to możliwe

## Typowe przypadki użycia

1. **Operacje w terenie**
   - Lokalizacje sprzętu
   - Adresy wezwania serwisowego
   - Miejsca inspekcji
   - Lokalizacje dostaw

2. **Zarządzanie wydarzeniami**
   - Miejsca wydarzeń
   - Lokalizacje spotkań
   - Miejsca konferencyjne
   - Miejsca warsztatów

3. **Śledzenie aktywów**
   - Pozycje sprzętu
   - Lokalizacje obiektów
   - Śledzenie pojazdów
   - Lokalizacje zapasów

4. **Analiza geograficzna**
   - Obszary zasięgu usług
   - Rozkład klientów
   - Analiza rynku
   - Zarządzanie terytoriami

## Funkcje integracji

### Z wyszukiwaniami
- Odwołuj się do danych lokalizacji z innych rekordów
- Znajduj rekordy według bliskości geograficznej
- Agreguj dane oparte na lokalizacji
- Krzyżowo odniesienia współrzędnych

### Z automatyzacjami
- Wyzwalaj akcje na podstawie zmian lokalizacji
- Twórz powiadomienia geofencingowe
- Aktualizuj powiązane rekordy, gdy zmieniają się lokalizacje
- Generuj raporty oparte na lokalizacji

### Z formułami
- Obliczaj odległości między lokalizacjami
- Określaj centra geograficzne
- Analizuj wzorce lokalizacji
- Twórz metryki oparte na lokalizacji

## Ograniczenia

- Brak wbudowanego geokodowania lub konwersji adresów
- Brak interfejsu mapowego
- Wymaga zewnętrznych usług dla funkcjonalności adresów
- Ograniczone tylko do przechowywania współrzędnych
- Brak automatycznej walidacji lokalizacji poza sprawdzaniem zakresu

## Powiązane zasoby

- [Przegląd pól niestandardowych](/api/custom-fields/list-custom-fields) - Ogólne pojęcia
- [Google Maps API](https://developers.google.com/maps) - Usługi geokodowania
- [OpenStreetMap Nominatim](https://nominatim.org/) - Bezpłatne geokodowanie
- [MapBox API](https://docs.mapbox.com/) - Usługi mapowania i geokodowania