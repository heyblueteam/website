---
title: Champ personnalisé de référence
description: Créez des champs de référence qui lient des enregistrements dans d'autres projets pour des relations inter-projets
---

Les champs personnalisés de référence vous permettent de créer des liens entre des enregistrements dans différents projets, permettant des relations inter-projets et le partage de données. Ils offrent un moyen puissant de connecter des travaux connexes à travers la structure des projets de votre organisation.

## Exemple de base

Créez un champ de référence simple :

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## Exemple avancé

Créez un champ de référence avec filtrage et sélection multiple :

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom affiché du champ de référence |
| `type` | CustomFieldType! | ✅ Oui | Doit être `REFERENCE` |
| `referenceProjectId` | String | Non | ID du projet à référencer |
| `referenceMultiple` | Boolean | Non | Autoriser la sélection de plusieurs enregistrements (par défaut : faux) |
| `referenceFilter` | TodoFilterInput | Non | Critères de filtrage pour les enregistrements référencés |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Les champs personnalisés sont automatiquement associés au projet en fonction du contexte de projet actuel de l'utilisateur.

## Configuration de référence

### Références uniques vs multiples

**Référence unique (par défaut) :**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Les utilisateurs peuvent sélectionner un enregistrement dans le projet référencé
- Renvoie un seul objet Todo

**Références multiples :**
```graphql
{
  referenceMultiple: true
}
```
- Les utilisateurs peuvent sélectionner plusieurs enregistrements dans le projet référencé
- Renvoie un tableau d'objets Todo

### Filtrage de référence

Utilisez `referenceFilter` pour limiter les enregistrements pouvant être sélectionnés :

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## Définition des valeurs de référence

### Référence unique

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Références multiples

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### Paramètres SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de référence |
| `customFieldReferenceTodoIds` | [String!] | ✅ Oui | Tableau des IDs des enregistrements référencés |

## Création d'enregistrements avec des références

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## Champs de réponse

### TodoCustomField Response

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ de référence |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Date de création de la valeur |
| `updatedAt` | DateTime! | Date de dernière modification de la valeur |

**Remarque** : Les todos référencés sont accessibles via `customField.selectedTodos`, et non directement sur TodoCustomField.

### Champs Todo référencés

Chaque Todo référencé comprend :

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | Identifiant unique de l'enregistrement référencé |
| `title` | String! | Titre de l'enregistrement référencé |
| `status` | TodoStatus! | Statut actuel (ACTIF, TERMINÉ, etc.) |
| `description` | String | Description de l'enregistrement référencé |
| `dueDate` | DateTime | Date d'échéance si définie |
| `assignees` | [User!] | Utilisateurs assignés |
| `tags` | [Tag!] | Étiquettes associées |
| `project` | Project! | Projet contenant l'enregistrement référencé |

## Interrogation des données de référence

### Requête de base

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### Requête avancée avec données imbriquées

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Important** : Les utilisateurs doivent avoir des permissions de visualisation sur le projet référencé pour voir les enregistrements liés.

## Accès inter-projets

### Visibilité du projet

- Les utilisateurs ne peuvent référencer que des enregistrements provenant de projets auxquels ils ont accès
- Les enregistrements référencés respectent les permissions du projet d'origine
- Les modifications apportées aux enregistrements référencés apparaissent en temps réel
- La suppression d'enregistrements référencés les retire des champs de référence

### Héritage des permissions

- Les champs de référence héritent des permissions des deux projets
- Les utilisateurs ont besoin d'un accès en lecture au projet référencé
- Les permissions d'édition sont basées sur les règles du projet actuel
- Les données référencées sont en lecture seule dans le contexte du champ de référence

## Réponses d'erreur

### Projet de référence invalide

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### Enregistrement référencé non trouvé

```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Permission refusée

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## Meilleures pratiques

### Conception des champs

1. **Nommage clair** - Utilisez des noms descriptifs qui indiquent la relation
2. **Filtrage approprié** - Définissez des filtres pour afficher uniquement les enregistrements pertinents
3. **Considérer les permissions** - Assurez-vous que les utilisateurs ont accès aux projets référencés
4. **Documenter les relations** - Fournissez des descriptions claires de la connexion

### Considérations de performance

1. **Limiter la portée de référence** - Utilisez des filtres pour réduire le nombre d'enregistrements sélectionnables
2. **Éviter les imbrications profondes** - Ne créez pas de chaînes complexes de références
3. **Considérer le cache** - Les données référencées sont mises en cache pour des performances optimales
4. **Surveiller l'utilisation** - Suivez comment les références sont utilisées à travers les projets

### Intégrité des données

1. **Gérer les suppressions** - Planifiez quand les enregistrements référencés sont supprimés
2. **Valider les permissions** - Assurez-vous que les utilisateurs peuvent accéder aux projets référencés
3. **Mettre à jour les dépendances** - Considérez l'impact lors du changement d'enregistrements référencés
4. **Pistes de vérification** - Suivez les relations de référence pour la conformité

## Cas d'utilisation courants

### Dépendances de projet

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### Exigences des clients

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### Allocation des ressources

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### Assurance qualité

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## Intégration avec les recherches

Les champs de référence fonctionnent avec [Champs de recherche](/api/custom-fields/lookup) pour extraire des données des enregistrements référencés. Les champs de recherche peuvent extraire des valeurs des enregistrements sélectionnés dans les champs de référence, mais ils ne sont que des extracteurs de données (aucune fonction d'agrégation comme SUM n'est prise en charge).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## Limitations

- Les projets référencés doivent être accessibles à l'utilisateur
- Les modifications des permissions du projet référencé affectent l'accès aux champs de référence
- L'imbrication profonde des références peut impacter la performance
- Pas de validation intégrée pour les références circulaires
- Pas de restriction automatique empêchant les références au même projet
- La validation des filtres n'est pas appliquée lors de la définition des valeurs de référence

## Ressources connexes

- [Champs de recherche](/api/custom-fields/lookup) - Extraire des données des enregistrements référencés
- [API Projets](/api/projects) - Gestion des projets contenant des références
- [API Enregistrements](/api/records) - Travailler avec des enregistrements ayant des références
- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux