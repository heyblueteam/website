---
title: Champ personnalisé de nombre
description: Créez des champs numériques pour stocker des valeurs numériques avec des contraintes min/max optionnelles et un formatage de préfixe
---

Les champs personnalisés de nombre vous permettent de stocker des valeurs numériques pour des enregistrements. Ils prennent en charge des contraintes de validation, une précision décimale, et peuvent être utilisés pour des quantités, des scores, des mesures ou toute donnée numérique qui ne nécessite pas de formatage spécial.

## Exemple de base

Créez un champ numérique simple :

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ numérique avec des contraintes et un préfixe :

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ numérique |
| `type` | CustomFieldType! | ✅ Oui | Doit être `NUMBER` |
| `projectId` | String! | ✅ Oui | ID du projet dans lequel créer le champ |
| `min` | Float | Non | Contrainte de valeur minimale (guidage UI uniquement) |
| `max` | Float | Non | Contrainte de valeur maximale (guidage UI uniquement) |
| `prefix` | String | Non | Préfixe d'affichage (ex : "#", "~", "$") |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

## Définir des valeurs numériques

Les champs numériques stockent des valeurs décimales avec une validation optionnelle :

### Valeur numérique simple

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Valeur entière

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de nombre |
| `number` | Float | Non | Valeur numérique à stocker |

## Contraintes de valeur

### Contraintes Min/Max (Guidage UI)

**Important** : Les contraintes min/max sont stockées mais NE SONT PAS appliquées côté serveur. Elles servent de guidage UI pour les applications frontend.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Validation côté client requise** : Les applications frontend doivent implémenter une logique de validation pour appliquer les contraintes min/max.

### Types de valeurs pris en charge

| Type | Exemple | Description |
|------|---------|-------------|
| Integer | `42` | Nombres entiers |
| Decimal | `42.5` | Nombres avec des décimales |
| Negative | `-10` | Valeurs négatives (si aucune contrainte min) |
| Zero | `0` | Valeur zéro |

**Remarque** : Les contraintes min/max NE SONT PAS validées côté serveur. Les valeurs en dehors de la plage spécifiée seront acceptées et stockées.

## Création d'enregistrements avec des valeurs numériques

Lors de la création d'un nouvel enregistrement avec des valeurs numériques :

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### Formats d'entrée pris en charge

Lors de la création d'enregistrements, utilisez le paramètre `number` (pas `value`) dans le tableau des champs personnalisés :

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `number` | Float | La valeur numérique |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

### Réponse CustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la définition du champ |
| `name` | String! | Nom d'affichage du champ |
| `type` | CustomFieldType! | Toujours `NUMBER` |
| `min` | Float | Valeur minimale autorisée |
| `max` | Float | Valeur maximale autorisée |
| `prefix` | String | Préfixe d'affichage |
| `description` | String | Texte d'aide |

**Remarque** : Si la valeur numérique n'est pas définie, le champ `number` sera `null`.

## Filtrage et requêtes

Les champs numériques prennent en charge un filtrage numérique complet :

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Opérateurs pris en charge

| Opérateur | Description | Exemple |
|-----------|-------------|---------|
| `EQ` | Égal à | `number = 42` |
| `NE` | Différent de | `number ≠ 42` |
| `GT` | Supérieur à | `number > 42` |
| `GTE` | Supérieur ou égal | `number ≥ 42` |
| `LT` | Inférieur à | `number < 42` |
| `LTE` | Inférieur ou égal | `number ≤ 42` |
| `IN` | Dans le tableau | `number in [1, 2, 3]` |
| `NIN` | Pas dans le tableau | `number not in [1, 2, 3]` |
| `IS` | Est nul/n'est pas nul | `number is null` |

### Filtrage par plage

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Formatage d'affichage

### Avec préfixe

Si un préfixe est défini, il sera affiché :

| Valeur | Préfixe | Affichage |
|--------|---------|-----------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Précision décimale

Les nombres conservent leur précision décimale :

| Entrée | Stocké | Affiché |
|--------|--------|---------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Permissions requises

| Action | Permission requise |
|--------|--------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Réponses d'erreur

### Format de nombre invalide
```json
{
  "errors": [{
    "message": "Invalid number format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Champ non trouvé
```json
{
  "errors": [{
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Remarque** : Les erreurs de validation min/max ne se produisent PAS côté serveur. La validation des contraintes doit être mise en œuvre dans votre application frontend.

### Pas un nombre
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Meilleures pratiques

### Conception des contraintes
- Définissez des valeurs min/max réalistes pour le guidage UI
- Implémentez une validation côté client pour appliquer les contraintes
- Utilisez des contraintes pour fournir des retours d'information aux utilisateurs dans les formulaires
- Considérez si les valeurs négatives sont valides pour votre cas d'utilisation

### Précision des valeurs
- Utilisez une précision décimale appropriée pour vos besoins
- Considérez l'arrondi à des fins d'affichage
- Soyez cohérent avec la précision à travers les champs connexes

### Amélioration de l'affichage
- Utilisez des préfixes significatifs pour le contexte
- Considérez les unités dans les noms de champs (ex : "Poids (kg)")
- Fournissez des descriptions claires pour les règles de validation

## Cas d'utilisation courants

1. **Systèmes de notation**
   - Évaluations de performance
   - Scores de qualité
   - Niveaux de priorité
   - Évaluations de satisfaction client

2. **Mesures**
   - Quantités et montants
   - Dimensions et tailles
   - Durées (au format numérique)
   - Capacités et limites

3. **Métriques commerciales**
   - Chiffres d'affaires
   - Taux de conversion
   - Allocations budgétaires
   - Chiffres cibles

4. **Données techniques**
   - Numéros de version
   - Valeurs de configuration
   - Métriques de performance
   - Paramètres de seuil

## Fonctionnalités d'intégration

### Avec des graphiques et des tableaux de bord
- Utilisez des champs NUMÉRIQUES dans les calculs de graphiques
- Créez des visualisations numériques
- Suivez les tendances au fil du temps

### Avec des automatisations
- Déclenchez des actions basées sur des seuils numériques
- Mettez à jour des champs connexes en fonction des changements numériques
- Envoyez des notifications pour des valeurs spécifiques

### Avec des recherches
- Agrégez des nombres à partir d'enregistrements connexes
- Calculez des totaux et des moyennes
- Trouvez des valeurs min/max à travers les relations

### Avec des graphiques
- Créez des visualisations numériques
- Suivez les tendances au fil du temps
- Comparez des valeurs à travers des enregistrements

## Limitations

- **Pas de validation côté serveur** des contraintes min/max
- **Validation côté client requise** pour l'application des contraintes
- Pas de formatage monétaire intégré (utilisez le type MONNAIE à la place)
- Pas de symbole de pourcentage automatique (utilisez le type POURCENT à la place)
- Pas de capacités de conversion d'unités
- Précision décimale limitée par le type Decimal de la base de données
- Pas d'évaluation de formules mathématiques dans le champ lui-même

## Ressources connexes

- [Aperçu des champs personnalisés](/api/custom-fields/1.index) - Concepts généraux des champs personnalisés
- [Champ personnalisé de monnaie](/api/custom-fields/currency) - Pour les valeurs monétaires
- [Champ personnalisé de pourcentage](/api/custom-fields/percent) - Pour les valeurs en pourcentage
- [API des automatisations](/api/automations/1.index) - Créez des automatisations basées sur des nombres