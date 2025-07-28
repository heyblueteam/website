---
title: Champ personnalisé d'évaluation
description: Créez des champs d'évaluation pour stocker des évaluations numériques avec des échelles et des validations configurables
---

Les champs personnalisés d'évaluation vous permettent de stocker des évaluations numériques dans des enregistrements avec des valeurs minimales et maximales configurables. Ils sont idéaux pour les évaluations de performance, les scores de satisfaction, les niveaux de priorité ou toute donnée basée sur une échelle numérique dans vos projets.

## Exemple de base

Créez un champ d'évaluation simple avec une échelle par défaut de 0 à 5 :

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Exemple avancé

Créez un champ d'évaluation avec une échelle et une description personnalisées :

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ d'évaluation |
| `type` | CustomFieldType! | ✅ Oui | Doit être `RATING` |
| `projectId` | String! | ✅ Oui | L'ID du projet où ce champ sera créé |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |
| `min` | Float | Non | Valeur minimale d'évaluation (pas de valeur par défaut) |
| `max` | Float | Non | Valeur maximale d'évaluation |

## Définir les valeurs d'évaluation

Pour définir ou mettre à jour une valeur d'évaluation sur un enregistrement :

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé d'évaluation |
| `value` | String! | ✅ Oui | Valeur d'évaluation sous forme de chaîne (dans la plage configurée) |

## Création d'enregistrements avec des valeurs d'évaluation

Lors de la création d'un nouvel enregistrement avec des valeurs d'évaluation :

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
    }
  }
}
```

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `value` | Float | La valeur d'évaluation stockée (accessible via customField.value) |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

**Remarque** : La valeur d'évaluation est en fait accessible via `customField.value.number` dans les requêtes.

### Réponse CustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour le champ |
| `name` | String! | Nom d'affichage du champ d'évaluation |
| `type` | CustomFieldType! | Toujours `RATING` |
| `min` | Float | Valeur minimale d'évaluation autorisée |
| `max` | Float | Valeur maximale d'évaluation autorisée |
| `description` | String | Texte d'aide pour le champ |

## Validation des évaluations

### Contraintes de valeur
- Les valeurs d'évaluation doivent être numériques (type Float)
- Les valeurs doivent être dans la plage min/max configurée
- Si aucune valeur minimale n'est spécifiée, il n'y a pas de valeur par défaut
- La valeur maximale est facultative mais recommandée

### Règles de validation
**Important** : La validation n'a lieu que lors de la soumission des formulaires, pas lors de l'utilisation de `setTodoCustomField` directement.

- L'entrée est analysée comme un nombre à virgule flottante (lors de l'utilisation de formulaires)
- Doit être supérieur ou égal à la valeur minimale (lors de l'utilisation de formulaires)
- Doit être inférieur ou égal à la valeur maximale (lors de l'utilisation de formulaires)
- `setTodoCustomField` accepte toute valeur de chaîne sans validation

### Exemples d'évaluation valides
Pour un champ avec min=1, max=5 :
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Exemples d'évaluation invalides
Pour un champ avec min=1, max=5 :
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Options de configuration

### Configuration de l'échelle d'évaluation
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Échelles d'évaluation courantes
- **1-5 Étoiles** : `min: 1, max: 5`
- **0-10 NPS** : `min: 0, max: 10`
- **1-10 Performance** : `min: 1, max: 10`
- **0-100 Pourcentage** : `min: 0, max: 100`
- **Échelle personnalisée** : Toute plage numérique

## Autorisations requises

Les opérations sur les champs personnalisés suivent les autorisations basées sur les rôles standard :

| Action | Rôle requis |
|--------|-------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Remarque** : Les rôles spécifiques requis dépendent de la configuration des rôles personnalisés de votre projet et des autorisations au niveau des champs.

## Réponses d'erreur

### Erreur de validation (uniquement pour les formulaires)
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

**Important** : La validation de la valeur d'évaluation (contraintes min/max) n'a lieu que lors de la soumission des formulaires, pas lors de l'utilisation de `setTodoCustomField` directement.

### Champ personnalisé introuvable
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

## Meilleures pratiques

### Conception de l'échelle
- Utilisez des échelles d'évaluation cohérentes à travers des champs similaires
- Tenez compte de la familiarité des utilisateurs (1-5 étoiles, 0-10 NPS)
- Définissez des valeurs minimales appropriées (0 contre 1)
- Définissez une signification claire pour chaque niveau d'évaluation

### Qualité des données
- Validez les valeurs d'évaluation avant de les stocker
- Utilisez la précision décimale de manière appropriée
- Envisagez d'arrondir à des fins d'affichage
- Fournissez des indications claires sur les significations des évaluations

### Expérience utilisateur
- Affichez visuellement les échelles d'évaluation (étoiles, barres de progression)
- Montrez la valeur actuelle et les limites de l'échelle
- Fournissez un contexte pour les significations des évaluations
- Envisagez des valeurs par défaut pour les nouveaux enregistrements

## Cas d'utilisation courants

1. **Gestion de la performance**
   - Évaluations de la performance des employés
   - Scores de qualité des projets
   - Évaluations de l'achèvement des tâches
   - Évaluations des niveaux de compétence

2. **Retour d'information des clients**
   - Évaluations de satisfaction
   - Scores de qualité des produits
   - Évaluations de l'expérience de service
   - Score Net Promoter (NPS)

3. **Priorité et importance**
   - Niveaux de priorité des tâches
   - Évaluations d'urgence
   - Scores d'évaluation des risques
   - Évaluations d'impact

4. **Assurance qualité**
   - Évaluations de révision de code
   - Scores de qualité des tests
   - Qualité de la documentation
   - Évaluations de conformité aux processus

## Fonctionnalités d'intégration

### Avec des automatisations
- Déclenchez des actions basées sur des seuils d'évaluation
- Envoyez des notifications pour des évaluations faibles
- Créez des tâches de suivi pour des évaluations élevées
- Dirigez le travail en fonction des valeurs d'évaluation

### Avec des recherches
- Calculez les évaluations moyennes à travers les enregistrements
- Trouvez des enregistrements par plages d'évaluation
- Référencez les données d'évaluation d'autres enregistrements
- Agrégez les statistiques d'évaluation

### Avec l'interface utilisateur de Blue
- Validation automatique des plages dans les contextes de formulaire
- Contrôles d'entrée d'évaluation visuels
- Retour d'information en temps réel sur la validation
- Options d'entrée par étoiles ou curseur

## Suivi des activités

Les changements de champ d'évaluation sont automatiquement suivis :
- Les anciennes et nouvelles valeurs d'évaluation sont enregistrées
- L'activité montre les changements numériques
- Horodatages pour toutes les mises à jour d'évaluation
- Attribution des utilisateurs pour les changements

## Limitations

- Seules les valeurs numériques sont prises en charge
- Pas d'affichage visuel intégré des évaluations (étoiles, etc.)
- La précision décimale dépend de la configuration de la base de données
- Pas de stockage de métadonnées d'évaluation (commentaires, contexte)
- Pas d'agrégation ou de statistiques d'évaluation automatiques
- Pas de conversion d'évaluation intégrée entre les échelles
- **Critique** : La validation min/max ne fonctionne que dans les formulaires, pas via `setTodoCustomField`

## Ressources connexes

- [Champs numériques](/api/5.custom%20fields/number) - Pour les données numériques générales
- [Champs de pourcentage](/api/5.custom%20fields/percent) - Pour les valeurs en pourcentage
- [Champs de sélection](/api/5.custom%20fields/select-single) - Pour les évaluations de choix discrets
- [Aperçu des champs personnalisés](/api/5.custom%20fields/2.list-custom-fields) - Concepts généraux