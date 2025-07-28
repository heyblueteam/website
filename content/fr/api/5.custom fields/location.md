---
title: Champ personnalisé de localisation
description: Créez des champs de localisation pour stocker des coordonnées géographiques pour des enregistrements
---

Les champs personnalisés de localisation stockent des coordonnées géographiques (latitude et longitude) pour des enregistrements. Ils prennent en charge le stockage précis des coordonnées, les requêtes géospatiales et le filtrage efficace basé sur la localisation.

## Exemple de base

Créez un champ de localisation simple :

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

## Exemple avancé

Créez un champ de localisation avec description :

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

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ de localisation |
| `type` | CustomFieldType! | ✅ Oui | Doit être `LOCATION` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

## Définir les valeurs de localisation

Les champs de localisation stockent les coordonnées de latitude et de longitude :

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

### Paramètres SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de localisation |
| `latitude` | Float | Non | Coordonnée de latitude (-90 à 90) |
| `longitude` | Float | Non | Coordonnée de longitude (-180 à 180) |

**Remarque** : Bien que les deux paramètres soient optionnels dans le schéma, les deux coordonnées sont requises pour une localisation valide. Si une seule est fournie, la localisation sera invalide.

## Validation des coordonnées

### Plages valides

| Coordonnée | Plage | Description |
|------------|-------|-------------|
| Latitude | -90 to 90 | Position Nord/Sud |
| Longitude | -180 to 180 | Position Est/Ouest |

### Exemples de coordonnées

| Localisation | Latitude | Longitude |
|--------------|----------|-----------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Création d'enregistrements avec des valeurs de localisation

Lors de la création d'un nouvel enregistrement avec des données de localisation :

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

### Format d'entrée pour la création

Lors de la création d'enregistrements, les valeurs de localisation utilisent un format séparé par des virgules :

| Format | Exemple | Description |
|--------|---------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Format de coordonnées standard |
| `"51.5074,-0.1278"` | London coordinates | Pas d'espaces autour de la virgule |
| `"-33.8688,151.2093"` | Sydney coordinates | Valeurs négatives autorisées |

## Champs de réponse

### TodoCustomField Response

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `latitude` | Float | Coordonnée de latitude |
| `longitude` | Float | Coordonnée de longitude |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Date de création de la valeur |
| `updatedAt` | DateTime! | Date de dernière modification de la valeur |

## Limitations importantes

### Pas de géocodage intégré

Les champs de localisation ne stockent que des coordonnées - ils ne comprennent **pas** :
- Conversion d'adresse en coordonnées
- Géocodage inversé (coordonnées en adresse)
- Validation ou recherche d'adresse
- Intégration avec des services de cartographie
- Recherche de nom de lieu

### Services externes requis

Pour la fonctionnalité d'adresse, vous devrez intégrer des services externes :
- **Google Maps API** pour le géocodage
- **OpenStreetMap Nominatim** pour le géocodage gratuit
- **MapBox** pour la cartographie et le géocodage
- **Here API** pour les services de localisation

### Exemple d'intégration

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

## Autorisations requises

| Action | Rôle requis |
|--------|-------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Réponses d'erreur

### Coordonnées invalides
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

### Longitude invalide
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

## Meilleures pratiques

### Collecte de données
- Utilisez des coordonnées GPS pour des emplacements précis
- Validez les coordonnées avant de les stocker
- Tenez compte des besoins de précision des coordonnées (6 décimales ≈ 10 cm de précision)
- Stockez les coordonnées en degrés décimaux (pas en degrés/minutes/secondes)

### Expérience utilisateur
- Fournissez des interfaces de carte pour la sélection de coordonnées
- Affichez des aperçus de localisation lors de l'affichage des coordonnées
- Validez les coordonnées côté client avant les appels API
- Tenez compte des implications de fuseau horaire pour les données de localisation

### Performance
- Utilisez des index spatiaux pour des requêtes efficaces
- Limitez la précision des coordonnées à la précision nécessaire
- Envisagez la mise en cache pour les emplacements fréquemment consultés
- Regroupez les mises à jour de localisation lorsque cela est possible

## Cas d'utilisation courants

1. **Opérations sur le terrain**
   - Emplacements des équipements
   - Adresses des appels de service
   - Sites d'inspection
   - Lieux de livraison

2. **Gestion d'événements**
   - Lieux d'événements
   - Lieux de réunion
   - Sites de conférence
   - Lieux d'ateliers

3. **Suivi des actifs**
   - Positions des équipements
   - Emplacements des installations
   - Suivi des véhicules
   - Lieux d'inventaire

4. **Analyse géographique**
   - Zones de couverture de service
   - Distribution des clients
   - Analyse de marché
   - Gestion de territoire

## Fonctionnalités d'intégration

### Avec des recherches
- Référencez les données de localisation d'autres enregistrements
- Trouvez des enregistrements par proximité géographique
- Agrégez des données basées sur la localisation
- Croisez les coordonnées

### Avec des automatisations
- Déclenchez des actions en fonction des changements de localisation
- Créez des notifications géorepérées
- Mettez à jour les enregistrements associés lorsque les emplacements changent
- Générez des rapports basés sur la localisation

### Avec des formules
- Calculez les distances entre les emplacements
- Déterminez les centres géographiques
- Analysez les modèles de localisation
- Créez des métriques basées sur la localisation

## Limitations

- Pas de géocodage intégré ou de conversion d'adresse
- Pas d'interface de cartographie fournie
- Nécessite des services externes pour la fonctionnalité d'adresse
- Limité au stockage de coordonnées uniquement
- Pas de validation automatique de la localisation au-delà de la vérification de la plage

## Ressources connexes

- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux
- [Google Maps API](https://developers.google.com/maps) - Services de géocodage
- [OpenStreetMap Nominatim](https://nominatim.org/) - Géocodage gratuit
- [MapBox API](https://docs.mapbox.com/) - Services de cartographie et de géocodage