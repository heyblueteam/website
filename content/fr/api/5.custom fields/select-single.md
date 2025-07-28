---
title: Champ personnalisé à sélection unique
description: Créez des champs à sélection unique pour permettre aux utilisateurs de choisir une option parmi une liste prédéfinie
---

Les champs personnalisés à sélection unique permettent aux utilisateurs de choisir exactement une option parmi une liste prédéfinie. Ils sont idéaux pour les champs de statut, les catégories, les priorités ou tout scénario où une seule option doit être choisie parmi un ensemble contrôlé d'options.

## Exemple de base

Créez un champ à sélection unique simple :

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ à sélection unique avec des options prédéfinies :

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ à sélection unique |
| `type` | CustomFieldType! | ✅ Oui | Doit être `SELECT_SINGLE` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | Non | Options initiales pour le champ |

### CreateCustomFieldOptionInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `title` | String! | ✅ Oui | Texte d'affichage pour l'option |
| `color` | String | Non | Code couleur hexadécimal pour l'option |

## Ajout d'options à des champs existants

Ajoutez de nouvelles options à un champ à sélection unique existant :

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## Définir les valeurs à sélection unique

Pour définir l'option sélectionnée sur un enregistrement :

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé à sélection unique |
| `customFieldOptionId` | String | Non | ID de l'option à sélectionner (préféré pour la sélection unique) |
| `customFieldOptionIds` | [String!] | Non | Tableau d'IDs d'options (utilise le premier élément pour la sélection unique) |

## Interrogation des valeurs à sélection unique

Interrogez la valeur à sélection unique d'un enregistrement :

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

Le champ `value` renvoie un objet JSON avec les détails de l'option sélectionnée.

## Création d'enregistrements avec des valeurs à sélection unique

Lors de la création d'un nouvel enregistrement avec des valeurs à sélection unique :

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
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
| `value` | JSON | Contient l'objet de l'option sélectionnée avec id, titre, couleur, position |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Date de création de la valeur |
| `updatedAt` | DateTime! | Date de dernière modification de la valeur |

### Réponse CustomFieldOption

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour l'option |
| `title` | String! | Texte d'affichage pour l'option |
| `color` | String | Code couleur hexadécimal pour la représentation visuelle |
| `position` | Float | Ordre de tri pour l'option |
| `customField` | CustomField! | Le champ personnalisé auquel cette option appartient |

### Réponse CustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour le champ |
| `name` | String! | Nom d'affichage du champ à sélection unique |
| `type` | CustomFieldType! | Toujours `SELECT_SINGLE` |
| `description` | String | Texte d'aide pour le champ |
| `customFieldOptions` | [CustomFieldOption!] | Toutes les options disponibles |

## Format de valeur

### Format d'entrée
- **Paramètre API** : Utilisez `customFieldOptionId` pour l'ID d'option unique
- **Alternative** : Utilisez `customFieldOptionIds` tableau (prend le premier élément)
- **Effacer la sélection** : Omettez les deux champs ou passez des valeurs vides

### Format de sortie
- **Réponse GraphQL** : Objet JSON dans le champ `value` contenant {id, titre, couleur, position}
- **Journal d'activité** : Titre de l'option sous forme de chaîne
- **Données d'automatisation** : Titre de l'option sous forme de chaîne

## Comportement de sélection

### Sélection exclusive
- Définir une nouvelle option supprime automatiquement la sélection précédente
- Une seule option peut être sélectionnée à la fois
- Définir `null` ou une valeur vide efface la sélection

### Logique de secours
- Si le tableau `customFieldOptionIds` est fourni, seule la première option est utilisée
- Cela garantit la compatibilité avec les formats d'entrée à sélection multiple
- Les tableaux vides ou les valeurs nulles effacent la sélection

## Gestion des options

### Mettre à jour les propriétés de l'option
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### Supprimer une option
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**Remarque** : La suppression d'une option l'effacera de tous les enregistrements où elle a été sélectionnée.

### Réorganiser les options
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## Règles de validation

### Validation des options
- L'ID d'option fourni doit exister
- L'option doit appartenir au champ personnalisé spécifié
- Une seule option peut être sélectionnée (appliqué automatiquement)
- Les valeurs nulles/vides sont valides (aucune sélection)

### Validation des champs
- Doit avoir au moins une option définie pour être utilisable
- Les titres des options doivent être uniques dans le champ
- Les codes couleur doivent être au format hexadécimal valide (s'ils sont fournis)

## Autorisations requises

| Action | Autorisation requise |
|--------|---------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Réponses d'erreur

### ID d'option invalide
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### L'option n'appartient pas au champ
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

### Impossible d'analyser la valeur
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Meilleures pratiques

### Conception des options
- Utilisez des titres d'option clairs et descriptifs
- Appliquez un codage couleur significatif
- Gardez les listes d'options ciblées et pertinentes
- Classez les options de manière logique (par priorité, fréquence, etc.)

### Modèles de champ de statut
- Utilisez des flux de travail de statut cohérents à travers les projets
- Considérez la progression naturelle des options
- Incluez des états finaux clairs (Fait, Annulé, etc.)
- Utilisez des couleurs qui reflètent la signification de l'option

### Gestion des données
- Examinez et nettoyez périodiquement les options inutilisées
- Utilisez des conventions de nommage cohérentes
- Considérez l'impact de la suppression d'options sur les enregistrements existants
- Planifiez les mises à jour et les migrations d'options

## Cas d'utilisation courants

1. **Statut et flux de travail**
   - Statut de la tâche (À faire, En cours, Fait)
   - Statut d'approbation (En attente, Approuvé, Rejeté)
   - Phase du projet (Planification, Développement, Test, Publié)
   - Statut de résolution des problèmes

2. **Classification et catégorisation**
   - Niveaux de priorité (Faible, Moyen, Élevé, Critique)
   - Types de tâches (Bug, Fonctionnalité, Amélioration, Documentation)
   - Catégories de projet (Interne, Client, Recherche)
   - Attributions de département

3. **Qualité et évaluation**
   - Statut de révision (Non commencé, En révision, Approuvé)
   - Évaluations de qualité (Mauvais, Passable, Bon, Excellent)
   - Niveaux de risque (Faible, Moyen, Élevé)
   - Niveaux de confiance

4. **Attribution et propriété**
   - Attributions d'équipe
   - Propriété de département
   - Attributions basées sur les rôles
   - Attributions régionales

## Fonctionnalités d'intégration

### Avec des automatisations
- Déclenchez des actions lorsque des options spécifiques sont sélectionnées
- Dirigez le travail en fonction des catégories sélectionnées
- Envoyez des notifications pour les changements de statut
- Créez des flux de travail conditionnels basés sur les sélections

### Avec des recherches
- Filtrer les enregistrements par options sélectionnées
- Référencer les données d'option d'autres enregistrements
- Créer des rapports basés sur les sélections d'options
- Regrouper les enregistrements par valeurs sélectionnées

### Avec des formulaires
- Contrôles d'entrée déroulants
- Interfaces de boutons radio
- Validation et filtrage des options
- Affichage conditionnel des champs basé sur les sélections

## Suivi des activités

Les changements de champ à sélection unique sont automatiquement suivis :
- Affiche les anciennes et nouvelles sélections d'options
- Affiche les titres d'options dans le journal d'activité
- Horodatages pour tous les changements de sélection
- Attribution des utilisateurs pour les modifications

## Différences avec la sélection multiple

| Fonctionnalité | Sélection unique | Sélection multiple |
|----------------|------------------|-------------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Limitations

- Une seule option peut être sélectionnée à la fois
- Pas de structure d'options hiérarchique ou imbriquée
- Les options sont partagées entre tous les enregistrements utilisant le champ
- Pas d'analytique ou de suivi d'utilisation des options intégré
- Les codes couleur sont uniquement pour l'affichage, sans impact fonctionnel
- Impossible de définir des autorisations différentes par option

## Ressources connexes

- [Champs à sélection multiple](/api/custom-fields/select-multi) - Pour des sélections à choix multiples
- [Champs de case à cocher](/api/custom-fields/checkbox) - Pour des choix booléens simples
- [Champs de texte](/api/custom-fields/text-single) - Pour une saisie de texte libre
- [Aperçu des champs personnalisés](/api/custom-fields/1.index) - Concepts généraux