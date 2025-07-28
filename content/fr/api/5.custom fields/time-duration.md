---
title: Champ personnalisé de durée
description: Créez des champs de durée calculée qui suivent le temps entre les événements de votre flux de travail
---

Les champs personnalisés de durée calculent et affichent automatiquement la durée entre deux événements dans votre flux de travail. Ils sont idéaux pour suivre les temps de traitement, les temps de réponse, les temps de cycle ou toute métrique basée sur le temps dans vos projets.

## Exemple de base

Créez un champ de durée simple qui suit combien de temps les tâches prennent pour être complétées :

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Exemple avancé

Créez un champ de durée complexe qui suit le temps entre les changements de champs personnalisés avec un objectif SLA :

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput (TIME_DURATION)

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ de durée |
| `type` | CustomFieldType! | ✅ Oui | Doit être `TIME_DURATION` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Oui | Comment afficher la durée |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Oui | Configuration de l'événement de début |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Oui | Configuration de l'événement de fin |
| `timeDurationTargetTime` | Float | Non | Durée cible en secondes pour le suivi SLA |

### CustomFieldTimeDurationInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ Oui | Type d'événement à suivre |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Oui | `FIRST` ou `LAST` occurrence |
| `customFieldId` | String | Conditional | Requis pour le type `TODO_CUSTOM_FIELD` |
| `customFieldOptionIds` | [String!] | Conditional | Requis pour les changements de champ sélectionné |
| `todoListId` | String | Conditional | Requis pour le type `TODO_MOVED` |
| `tagId` | String | Conditional | Requis pour le type `TODO_TAG_ADDED` |
| `assigneeId` | String | Conditional | Requis pour le type `TODO_ASSIGNEE_ADDED` |

### Valeurs de CustomFieldTimeDurationType

| Valeur | Description |
|--------|-------------|
| `TODO_CREATED_AT` | Lorsque l'enregistrement a été créé |
| `TODO_CUSTOM_FIELD` | Lorsque la valeur d'un champ personnalisé a changé |
| `TODO_DUE_DATE` | Lorsque la date d'échéance a été fixée |
| `TODO_MARKED_AS_COMPLETE` | Lorsque l'enregistrement a été marqué comme complet |
| `TODO_MOVED` | Lorsque l'enregistrement a été déplacé vers une autre liste |
| `TODO_TAG_ADDED` | Lorsque une étiquette a été ajoutée à l'enregistrement |
| `TODO_ASSIGNEE_ADDED` | Lorsque un assigné a été ajouté à l'enregistrement |

### Valeurs de CustomFieldTimeDurationCondition

| Valeur | Description |
|--------|-------------|
| `FIRST` | Utiliser la première occurrence de l'événement |
| `LAST` | Utiliser la dernière occurrence de l'événement |

### Valeurs de CustomFieldTimeDurationDisplayType

| Valeur | Description | Exemple |
|--------|-------------|---------|
| `FULL_DATE` | Format Jours:Heures:Minutes:Secondes | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Écrit en toutes lettres | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numérique avec unités | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Durée en jours uniquement | `"2.5"` (2.5 days) |
| `HOURS` | Durée en heures uniquement | `"60"` (60 hours) |
| `MINUTES` | Durée en minutes uniquement | `"3600"` (3600 minutes) |
| `SECONDS` | Durée en secondes uniquement | `"216000"` (216000 seconds) |

## Champs de réponse

### TodoCustomField Réponse

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `number` | Float | La durée en secondes |
| `value` | Float | Alias pour le nombre (durée en secondes) |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Lorsque la valeur a été créée |
| `updatedAt` | DateTime! | Lorsque la valeur a été mise à jour pour la dernière fois |

### CustomField Réponse (TIME_DURATION)

| Champ | Type | Description |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Format d'affichage pour la durée |
| `timeDurationStart` | CustomFieldTimeDuration | Configuration de l'événement de début |
| `timeDurationEnd` | CustomFieldTimeDuration | Configuration de l'événement de fin |
| `timeDurationTargetTime` | Float | Durée cible en secondes (pour le suivi SLA) |

## Calcul de la durée

### Comment ça fonctionne
1. **Événement de début** : Le système surveille l'événement de début spécifié
2. **Événement de fin** : Le système surveille l'événement de fin spécifié
3. **Calcul** : Durée = Heure de fin - Heure de début
4. **Stockage** : Durée stockée en secondes sous forme de nombre
5. **Affichage** : Formatée selon le paramètre `timeDurationDisplay`

### Déclencheurs de mise à jour
Les valeurs de durée sont automatiquement recalculées lorsque :
- Des enregistrements sont créés ou mis à jour
- Les valeurs des champs personnalisés changent
- Des étiquettes sont ajoutées ou supprimées
- Des assignés sont ajoutés ou supprimés
- Des enregistrements sont déplacés entre des listes
- Des enregistrements sont marqués comme complets/incomplets

## Lecture des valeurs de durée

### Interroger les champs de durée
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Valeurs d'affichage formatées
Les valeurs de durée sont automatiquement formatées en fonction du paramètre `timeDurationDisplay` :

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Exemples de configuration courants

### Temps de complétion des tâches
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Durée de changement de statut
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Temps dans une liste spécifique
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Temps de réponse d'assignation
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Réponses d'erreur

### Configuration invalide
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Champ référencé non trouvé
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Options requises manquantes
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Notes importantes

### Calcul automatique
- Les champs de durée sont **en lecture seule** - les valeurs sont automatiquement calculées
- Vous ne pouvez pas définir manuellement les valeurs de durée via l'API
- Les calculs se font de manière asynchrone via des tâches en arrière-plan
- Les valeurs se mettent à jour automatiquement lorsque des événements déclencheurs se produisent

### Considérations de performance
- Les calculs de durée sont mis en file d'attente et traités de manière asynchrone
- Un grand nombre de champs de durée peut affecter les performances
- Considérez la fréquence des événements déclencheurs lors de la conception des champs de durée
- Utilisez des conditions spécifiques pour éviter des recalculs inutiles

### Valeurs nulles
Les champs de durée afficheront `null` lorsque :
- L'événement de début ne s'est pas encore produit
- L'événement de fin ne s'est pas encore produit
- La configuration fait référence à des entités inexistantes
- Le calcul rencontre une erreur

## Meilleures pratiques

### Conception de la configuration
- Utilisez des types d'événements spécifiques plutôt que génériques lorsque cela est possible
- Choisissez des conditions appropriées `FIRST` vs `LAST` en fonction de votre flux de travail
- Testez les calculs de durée avec des données d'exemple avant le déploiement
- Documentez votre logique de champ de durée pour les membres de l'équipe

### Formatage d'affichage
- Utilisez `FULL_DATE_SUBSTRING` pour le format le plus lisible
- Utilisez `FULL_DATE` pour un affichage compact et de largeur constante
- Utilisez `FULL_DATE_STRING` pour des rapports et documents formels
- Utilisez `DAYS`, `HOURS`, `MINUTES`, ou `SECONDS` pour des affichages numériques simples
- Considérez les contraintes d'espace de votre interface utilisateur lors du choix du format

### Suivi SLA avec temps cible
Lors de l'utilisation de `timeDurationTargetTime` :
- Définissez la durée cible en secondes
- Comparez la durée réelle par rapport à la cible pour la conformité SLA
- Utilisez dans les tableaux de bord pour mettre en évidence les éléments en retard
- Exemple : SLA de réponse de 24 heures = 86400 secondes

### Intégration au flux de travail
- Concevez des champs de durée pour correspondre à vos processus métier réels
- Utilisez les données de durée pour l'amélioration et l'optimisation des processus
- Surveillez les tendances de durée pour identifier les goulets d'étranglement du flux de travail
- Configurez des alertes pour les seuils de durée si nécessaire

## Cas d'utilisation courants

1. **Performance des processus**
   - Temps de complétion des tâches
   - Temps de cycle de révision
   - Temps de traitement des approbations
   - Temps de réponse

2. **Suivi SLA**
   - Temps jusqu'à la première réponse
   - Temps de résolution
   - Délais d'escalade
   - Conformité au niveau de service

3. **Analyse du flux de travail**
   - Identification des goulets d'étranglement
   - Optimisation des processus
   - Métriques de performance de l'équipe
   - Timing de l'assurance qualité

4. **Gestion de projet**
   - Durées des phases
   - Timing des jalons
   - Temps d'allocation des ressources
   - Délais de livraison

## Limitations

- Les champs de durée sont **en lecture seule** et ne peuvent pas être définis manuellement
- Les valeurs sont calculées de manière asynchrone et peuvent ne pas être immédiatement disponibles
- Nécessite que des déclencheurs d'événements appropriés soient configurés dans votre flux de travail
- Ne peut pas calculer les durées pour des événements qui ne se sont pas produits
- Limité à suivre le temps entre des événements discrets (pas de suivi du temps continu)
- Pas d'alertes ou de notifications SLA intégrées
- Ne peut pas agréger plusieurs calculs de durée dans un seul champ

## Ressources connexes

- [Champs numériques](/api/custom-fields/number) - Pour des valeurs numériques manuelles
- [Champs de date](/api/custom-fields/date) - Pour le suivi de dates spécifiques
- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux
- [Automatisations](/api/automations) - Pour déclencher des actions basées sur des seuils de durée