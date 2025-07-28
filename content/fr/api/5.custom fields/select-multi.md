---
title: Champ personnalisé à sélection multiple
description: Créez des champs à sélection multiple pour permettre aux utilisateurs de choisir plusieurs options dans une liste prédéfinie
---

Les champs personnalisés à sélection multiple permettent aux utilisateurs de choisir plusieurs options dans une liste prédéfinie. Ils sont idéaux pour les catégories, les étiquettes, les compétences, les fonctionnalités ou tout scénario où plusieurs sélections sont nécessaires à partir d'un ensemble contrôlé d'options.

## Exemple de base

Créez un simple champ à sélection multiple :

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ à sélection multiple, puis ajoutez des options séparément :

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ à sélection multiple |
| `type` | CustomFieldType! | ✅ Oui | Doit être `SELECT_MULTI` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |
| `projectId` | String! | ✅ Oui | ID du projet pour ce champ |

### CreateCustomFieldOptionInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé |
| `title` | String! | ✅ Oui | Texte d'affichage pour l'option |
| `color` | String | Non | Couleur pour l'option (n'importe quelle chaîne) |
| `position` | Float | Non | Ordre de tri pour l'option |

## Ajout d'options à des champs existants

Ajoutez de nouvelles options à un champ à sélection multiple existant :

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## Définir des valeurs à sélection multiple

Pour définir plusieurs options sélectionnées sur un enregistrement :

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### Paramètres SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé à sélection multiple |
| `customFieldOptionIds` | [String!] | ✅ Oui | Tableau des IDs d'options à sélectionner |

## Création d'enregistrements avec des valeurs à sélection multiple

Lors de la création d'un nouvel enregistrement avec des valeurs à sélection multiple :

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
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
| `selectedOptions` | [CustomFieldOption!] | Tableau des options sélectionnées |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

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
| `name` | String! | Nom d'affichage du champ à sélection multiple |
| `type` | CustomFieldType! | Toujours `SELECT_MULTI` |
| `description` | String | Texte d'aide pour le champ |
| `customFieldOptions` | [CustomFieldOption!] | Toutes les options disponibles |

## Format de valeur

### Format d'entrée
- **Paramètre API** : Tableau des IDs d'options (`["option1", "option2", "option3"]`)
- **Format de chaîne** : IDs d'options séparés par des virgules (`"option1,option2,option3"`)

### Format de sortie
- **Réponse GraphQL** : Tableau d'objets CustomFieldOption
- **Journal d'activité** : Titres d'options séparés par des virgules
- **Données d'automatisation** : Tableau des titres d'options

## Gestion des options

### Mettre à jour les propriétés de l'option
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
  }) {
    id
    title
    color
  }
}
```

### Supprimer l'option
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### Réorganiser les options
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## Règles de validation

### Validation des options
- Tous les IDs d'options fournis doivent exister
- Les options doivent appartenir au champ personnalisé spécifié
- Seuls les champs SELECT_MULTI peuvent avoir plusieurs options sélectionnées
- Un tableau vide est valide (aucune sélection)

### Validation du champ
- Doit avoir au moins une option définie pour être utilisable
- Les titres des options doivent être uniques dans le champ
- Le champ de couleur accepte n'importe quelle valeur de chaîne (pas de validation hexadécimale)

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Réponses d'erreur

### ID d'option invalide
```json
{
  "errors": [{
    "message": "Custom field option not found",
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Plusieurs options sur un champ non-multi
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Meilleures pratiques

### Conception des options
- Utilisez des titres d'options descriptifs et concis
- Appliquez des schémas de codage couleur cohérents
- Gardez les listes d'options gérables (typiquement 3-20 options)
- Ordre logique des options (alphabétiquement, par fréquence, etc.)

### Gestion des données
- Examinez et nettoyez périodiquement les options inutilisées
- Utilisez des conventions de nommage cohérentes à travers les projets
- Considérez la réutilisabilité des options lors de la création de champs
- Planifiez les mises à jour et les migrations d'options

### Expérience utilisateur
- Fournissez des descriptions claires des champs
- Utilisez des couleurs pour améliorer la distinction visuelle
- Regroupez les options connexes
- Envisagez des sélections par défaut pour les cas courants

## Cas d'utilisation courants

1. **Gestion de projet**
   - Catégories et étiquettes de tâches
   - Niveaux et types de priorité
   - Assignations de membres d'équipe
   - Indicateurs de statut

2. **Gestion de contenu**
   - Catégories et sujets d'articles
   - Types et formats de contenu
   - Canaux de publication
   - Flux de travail d'approbation

3. **Support client**
   - Catégories et types de problèmes
   - Produits ou services concernés
   - Méthodes de résolution
   - Segments de clients

4. **Développement de produit**
   - Catégories de fonctionnalités
   - Exigences techniques
   - Environnements de test
   - Canaux de publication

## Fonctionnalités d'intégration

### Avec les automatisations
- Déclencher des actions lorsque des options spécifiques sont sélectionnées
- Acheminer le travail en fonction des catégories sélectionnées
- Envoyer des notifications pour des sélections de haute priorité
- Créer des tâches de suivi basées sur des combinaisons d'options

### Avec les recherches
- Filtrer les enregistrements par options sélectionnées
- Agréger des données à travers les sélections d'options
- Référencer des données d'options à partir d'autres enregistrements
- Créer des rapports basés sur des combinaisons d'options

### Avec les formulaires
- Contrôles d'entrée à sélection multiple
- Validation et filtrage des options
- Chargement dynamique des options
- Affichage conditionnel des champs

## Suivi des activités

Les changements de champ à sélection multiple sont automatiquement suivis :
- Montre les options ajoutées et supprimées
- Affiche les titres des options dans le journal d'activité
- Horodatages pour tous les changements de sélection
- Attribution utilisateur pour les modifications

## Limitations

- La limite pratique maximale d'options dépend des performances de l'interface utilisateur
- Pas de structure d'options hiérarchique ou imbriquée
- Les options sont partagées entre tous les enregistrements utilisant le champ
- Pas d'analytique ou de suivi d'utilisation des options intégré
- Le champ de couleur accepte n'importe quelle chaîne (pas de validation hexadécimale)
- Impossible de définir des permissions différentes par option
- Les options doivent être créées séparément, pas en ligne avec la création de champ
- Pas de mutation de réorganisation dédiée (utilisez editCustomFieldOption avec position)

## Ressources connexes

- [Champs à sélection unique](/api/custom-fields/select-single) - Pour des sélections à choix unique
- [Champs à cases à cocher](/api/custom-fields/checkbox) - Pour des choix booléens simples
- [Champs de texte](/api/custom-fields/text-single) - Pour une saisie de texte libre
- [Aperçu des champs personnalisés](/api/custom-fields/2.list-custom-fields) - Concepts généraux