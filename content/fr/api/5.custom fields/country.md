---
title: Champ personnalisé de pays
description: Créez des champs de sélection de pays avec validation du code de pays ISO
---

Les champs personnalisés de pays vous permettent de stocker et de gérer des informations sur les pays pour les enregistrements. Le champ prend en charge à la fois les noms de pays et les codes de pays ISO Alpha-2.

**Important** : Le comportement de validation et de conversion des pays diffère considérablement entre les mutations :
- **createTodo** : Valide et convertit automatiquement les noms de pays en codes ISO
- **setTodoCustomField** : Accepte toute valeur sans validation

## Exemple de base

Créez un champ de pays simple :

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de pays avec description :

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ de pays |
| `type` | CustomFieldType! | ✅ Oui | Doit être `COUNTRY` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Le `projectId` n'est pas passé dans l'entrée mais est déterminé par le contexte GraphQL (généralement à partir des en-têtes de requête ou de l'authentification).

## Définition des valeurs de pays

Les champs de pays stockent des données dans deux champs de base de données :
- **`countryCodes`** : Stocke les codes de pays ISO Alpha-2 sous forme de chaîne séparée par des virgules dans la base de données (retournée sous forme de tableau via l'API)
- **`text`** : Stocke le texte d'affichage ou les noms de pays sous forme de chaîne

### Comprendre les paramètres

La mutation `setTodoCustomField` accepte deux paramètres optionnels pour les champs de pays :

| Paramètre | Type | Requis | Description | Ce qu'il fait |
|-----------|------|--------|-------------|----------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour | - |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de pays | - |
| `countryCodes` | [String!] | Non | Tableau de codes de pays ISO Alpha-2 | Stored in the `countryCodes` field |
| `text` | String | Non | Texte d'affichage ou noms de pays | Stored in the `text` field |

**Important** : 
- Dans `setTodoCustomField` : Les deux paramètres sont optionnels et stockés indépendamment
- Dans `createTodo` : Le système définit automatiquement les deux champs en fonction de votre entrée (vous ne pouvez pas les contrôler indépendamment)

### Option 1 : Utiliser uniquement des codes de pays

Stockez des codes ISO validés sans texte d'affichage :

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Résultat : `countryCodes` = `["US"]`, `text` = `null`

### Option 2 : Utiliser uniquement du texte

Stockez du texte d'affichage sans codes validés :

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Résultat : `countryCodes` = `null`, `text` = `"United States"`

**Remarque** : Lors de l'utilisation de `setTodoCustomField`, aucune validation n'a lieu, quel que soit le paramètre utilisé. Les valeurs sont stockées exactement comme fournies.

### Option 3 : Utiliser les deux (recommandé)

Stockez à la fois des codes validés et du texte d'affichage :

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Résultat : `countryCodes` = `["US"]`, `text` = `"United States"`

### Plusieurs pays

Stockez plusieurs pays en utilisant des tableaux :

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Création d'enregistrements avec des valeurs de pays

Lors de la création d'enregistrements, la mutation `createTodo` **valide et convertit automatiquement** les valeurs de pays. C'est la seule mutation qui effectue une validation des pays :

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      countryCodes
    }
  }
}
```

### Formats d'entrée acceptés

| Type d'entrée | Exemple | Résultat |
|---------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Non pris en charge** - traité comme une seule valeur invalide |
| Mixed format | `"United States, CA"` | **Non pris en charge** - traité comme une seule valeur invalide |

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `text` | String | Texte d'affichage (noms de pays) |
| `countryCodes` | [String!] | Tableau de codes de pays ISO Alpha-2 |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Date de création de la valeur |
| `updatedAt` | DateTime! | Date de dernière modification de la valeur |

## Normes de pays

Blue utilise la norme **ISO 3166-1 Alpha-2** pour les codes de pays :

- Codes de pays à deux lettres (par exemple, US, GB, FR, DE)
- La validation utilisant la bibliothèque `i18n-iso-countries` **n'a lieu que dans createTodo**
- Prend en charge tous les pays officiellement reconnus

### Exemples de codes de pays

| Pays | Code ISO |
|------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Pour la liste officielle complète des codes de pays ISO 3166-1 alpha-2, visitez la [Plateforme de navigation en ligne de l'ISO](https://www.iso.org/obp/ui/#search/code/).

## Validation

**La validation n'a lieu que dans la mutation `createTodo`** :

1. **Code ISO valide** : Accepte tout code ISO Alpha-2 valide
2. **Nom de pays** : Convertit automatiquement les noms de pays reconnus en codes
3. **Entrée invalide** : Lance `CustomFieldValueParseError` pour les valeurs non reconnues

**Remarque** : La mutation `setTodoCustomField` ne réalise AUCUNE validation et accepte toute valeur de chaîne.

### Exemple d'erreur

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Fonctionnalités d'intégration

### Champs de recherche
Les champs de pays peuvent être référencés par des champs personnalisés de RECHERCHE, vous permettant d'extraire des données sur les pays à partir d'enregistrements liés.

### Automatisations
Utilisez les valeurs de pays dans les conditions d'automatisation :
- Filtrer les actions par pays spécifiques
- Envoyer des notifications en fonction du pays
- Acheminer les tâches en fonction des régions géographiques

### Formulaires
Les champs de pays dans les formulaires valident automatiquement l'entrée des utilisateurs et convertissent les noms de pays en codes.

## Autorisations requises

| Action | Autorisation requise |
|--------|---------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Réponses d'erreur

### Valeur de pays invalide
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Incompatibilité de type de champ
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Meilleures pratiques

### Gestion des entrées
- Utilisez `createTodo` pour une validation et une conversion automatiques
- Utilisez `setTodoCustomField` avec précaution car cela contourne la validation
- Envisagez de valider les entrées dans votre application avant d'utiliser `setTodoCustomField`
- Affichez les noms complets des pays dans l'interface utilisateur pour plus de clarté

### Qualité des données
- Validez les entrées de pays au point d'entrée
- Utilisez des formats cohérents dans votre système
- Envisagez des regroupements régionaux pour les rapports

### Plusieurs pays
- Utilisez le support des tableaux dans `setTodoCustomField` pour plusieurs pays
- Plusieurs pays dans `createTodo` ne sont **pas pris en charge** via le champ de valeur
- Stockez les codes de pays sous forme de tableau dans `setTodoCustomField` pour un traitement approprié

## Cas d'utilisation courants

1. **Gestion des clients**
   - Emplacement du siège social du client
   - Destinations d'expédition
   - Juridictions fiscales

2. **Suivi de projet**
   - Emplacement du projet
   - Emplacements des membres de l'équipe
   - Cibles de marché

3. **Conformité et juridique**
   - Juridictions réglementaires
   - Exigences de résidence des données
   - Contrôles à l'exportation

4. **Ventes et marketing**
   - Assignations de territoire
   - Segmentation de marché
   - Ciblage de campagne

## Limitations

- Ne prend en charge que les codes ISO 3166-1 Alpha-2 (codes à 2 lettres)
- Pas de support intégré pour les subdivisions de pays (états/provinces)
- Pas d'icônes de drapeau de pays automatiques (uniquement basées sur du texte)
- Impossible de valider les codes de pays historiques
- Pas de regroupement de région ou de continent intégré
- **La validation ne fonctionne que dans `createTodo`, pas dans `setTodoCustomField`**
- **Plusieurs pays non pris en charge dans le champ de valeur `createTodo`**
- **Les codes de pays sont stockés sous forme de chaîne séparée par des virgules, pas de véritable tableau**

## Ressources connexes

- [Aperçu des champs personnalisés](/custom-fields/list-custom-fields) - Concepts généraux des champs personnalisés
- [Champs de recherche](/api/custom-fields/lookup) - Référencez les données sur les pays à partir d'autres enregistrements
- [API des formulaires](/api/forms) - Incluez des champs de pays dans des formulaires personnalisés