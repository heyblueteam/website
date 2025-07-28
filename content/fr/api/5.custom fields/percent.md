---
title: Champ personnalisé de pourcentage
description: Créez des champs de pourcentage pour stocker des valeurs numériques avec gestion automatique du symbole % et formatage d'affichage
---

Les champs personnalisés de pourcentage vous permettent de stocker des valeurs de pourcentage pour des enregistrements. Ils gèrent automatiquement le symbole % pour l'entrée et l'affichage, tout en stockant la valeur numérique brute en interne. Parfait pour les taux d'achèvement, les taux de réussite ou toute métrique basée sur un pourcentage.

## Exemple de base

Créez un champ de pourcentage simple :

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de pourcentage avec description :

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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
| `name` | String! | ✅ Oui | Nom d'affichage du champ de pourcentage |
| `type` | CustomFieldType! | ✅ Oui | Doit être `PERCENT` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Le contexte du projet est automatiquement déterminé à partir de vos en-têtes d'authentification. Aucun paramètre `projectId` n'est nécessaire.

**Remarque** : Les champs PERCENT ne prennent pas en charge les contraintes min/max ou le formatage de préfixe comme les champs NUMBER.

## Définir des valeurs de pourcentage

Les champs de pourcentage stockent des valeurs numériques avec gestion automatique du symbole % :

### Avec symbole de pourcentage

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Valeur numérique directe

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### Paramètres SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de pourcentage |
| `number` | Float | Non | Valeur de pourcentage numérique (par exemple, 75.5 pour 75.5%) |

## Stockage et affichage des valeurs

### Format de stockage
- **Stockage interne** : Valeur numérique brute (par exemple, 75.5)
- **Base de données** : Stocké en tant que `Decimal` dans la colonne `number`
- **GraphQL** : Renvoyé en tant que type `Float`

### Format d'affichage
- **Interface utilisateur** : Les applications clientes doivent ajouter le symbole % (par exemple, "75.5%")
- **Graphiques** : Affiche avec le symbole % lorsque le type de sortie est PERCENTAGE
- **Réponses API** : Valeur numérique brute sans symbole % (par exemple, 75.5)

## Création d'enregistrements avec des valeurs de pourcentage

Lors de la création d'un nouvel enregistrement avec des valeurs de pourcentage :

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Formats d'entrée pris en charge

| Format | Exemple | Résultat |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Remarque** : Le symbole % est automatiquement supprimé de l'entrée et ajouté lors de l'affichage.

## Interrogation des valeurs de pourcentage

Lors de l'interrogation d'enregistrements avec des champs personnalisés de pourcentage, accédez à la valeur via le chemin `customField.value.number` :

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

La réponse inclura le pourcentage sous forme de nombre brut :

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé (contient la valeur de pourcentage) |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

**Important** : Les valeurs de pourcentage sont accessibles via le champ `customField.value.number`. Le symbole % n'est pas inclus dans les valeurs stockées et doit être ajouté par les applications clientes pour l'affichage.

## Filtrage et interrogation

Les champs de pourcentage prennent en charge le même filtrage que les champs NUMBER :

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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
| `EQ` | Égal à | `percentage = 75` |
| `NE` | Différent de | `percentage ≠ 75` |
| `GT` | Supérieur à | `percentage > 75` |
| `GTE` | Supérieur ou égal à | `percentage ≥ 75` |
| `LT` | Inférieur à | `percentage < 75` |
| `LTE` | Inférieur ou égal à | `percentage ≤ 75` |
| `IN` | Valeur dans la liste | `percentage in [50, 75, 100]` |
| `NIN` | Valeur non dans la liste | `percentage not in [0, 25]` |
| `IS` | Vérifier si nul avec `values: null` | `percentage is null` |
| `NOT` | Vérifier si non nul avec `values: null` | `percentage is not null` |

### Filtrage par plage

Pour le filtrage par plage, utilisez plusieurs opérateurs :

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Plages de valeurs de pourcentage

### Plages courantes

| Plage | Description | Cas d'utilisation |
|-------|-------------|------------------|
| `0-100` | Pourcentage standard | Completion rates, success rates |
| `0-∞` | Pourcentage illimité | Growth rates, performance metrics |
| `-∞-∞` | Toute valeur | Change rates, variance |

### Valeurs d'exemple

| Entrée | Stocké | Affichage |
|--------|--------|-----------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Agrégation de graphiques

Les champs de pourcentage prennent en charge l'agrégation dans les graphiques et rapports de tableau de bord. Les fonctions disponibles incluent :

- `AVERAGE` - Valeur moyenne de pourcentage
- `COUNT` - Nombre d'enregistrements avec des valeurs
- `MIN` - Valeur de pourcentage la plus basse
- `MAX` - Valeur de pourcentage la plus élevée 
- `SUM` - Total de toutes les valeurs de pourcentage

Ces agrégations sont disponibles lors de la création de graphiques et de tableaux de bord, pas dans des requêtes GraphQL directes.

## Permissions requises

| Action | Permission requise |
|--------|--------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Réponses d'erreur

### Format de pourcentage invalide
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

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

### Saisie de valeur
- Permettre aux utilisateurs d'entrer avec ou sans symbole %
- Valider des plages raisonnables pour votre cas d'utilisation
- Fournir un contexte clair sur ce que représente 100%

### Affichage
- Toujours afficher le symbole % dans les interfaces utilisateur
- Utiliser une précision décimale appropriée
- Envisager un code couleur pour les plages (rouge/jaune/vert)

### Interprétation des données
- Documenter ce que signifie 100% dans votre contexte
- Gérer les valeurs supérieures à 100% de manière appropriée
- Considérer si les valeurs négatives sont valides

## Cas d'utilisation courants

1. **Gestion de projet**
   - Taux d'achèvement des tâches
   - Avancement du projet
   - Utilisation des ressources
   - Vitesse de sprint

2. **Suivi de performance**
   - Taux de réussite
   - Taux d'erreur
   - Métriques d'efficacité
   - Scores de qualité

3. **Métriques financières**
   - Taux de croissance
   - Marges bénéficiaires
   - Montants de remise
   - Pourcentages de changement

4. **Analytique**
   - Taux de conversion
   - Taux de clics
   - Métriques d'engagement
   - Indicateurs de performance

## Fonctionnalités d'intégration

### Avec des formules
- Référencer les champs PERCENT dans les calculs
- Formatage automatique du symbole % dans les sorties de formule
- Combiner avec d'autres champs numériques

### Avec des automatisations
- Déclencher des actions basées sur des seuils de pourcentage
- Envoyer des notifications pour des pourcentages de jalons
- Mettre à jour le statut en fonction des taux d'achèvement

### Avec des recherches
- Agréger des pourcentages à partir d'enregistrements liés
- Calculer des taux de réussite moyens
- Trouver les éléments les plus performants/les moins performants

### Avec des graphiques
- Créer des visualisations basées sur des pourcentages
- Suivre les progrès au fil du temps
- Comparer les métriques de performance

## Différences par rapport aux champs NUMBER

### Qu'est-ce qui est différent
- **Gestion de l'entrée** : Supprime automatiquement le symbole %
- **Affichage** : Ajoute automatiquement le symbole %
- **Contraintes** : Pas de validation min/max
- **Formatage** : Pas de support de préfixe

### Qu'est-ce qui est le même
- **Stockage** : Même colonne de base de données et type
- **Filtrage** : Même opérateurs de requête
- **Agrégation** : Même fonctions d'agrégation
- **Permissions** : Même modèle de permission

## Limitations

- Pas de contraintes de valeur min/max
- Pas d'options de formatage de préfixe
- Pas de validation automatique de la plage 0-100%
- Pas de conversion entre formats de pourcentage (par exemple, 0.75 ↔ 75%)
- Les valeurs supérieures à 100% sont autorisées

## Ressources connexes

- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux des champs personnalisés
- [Champ personnalisé numérique](/api/custom-fields/number) - Pour des valeurs numériques brutes
- [API d'automatisations](/api/automations/index) - Créer des automatisations basées sur des pourcentages