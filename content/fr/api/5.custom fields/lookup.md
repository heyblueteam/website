---
title: Champ de recherche personnalisé
description: Créez des champs de recherche qui extraient automatiquement des données des enregistrements référencés
---

Les champs de recherche personnalisés extraient automatiquement des données des enregistrements référencés par [Champs de référence](/api/custom-fields/reference), affichant des informations des enregistrements liés sans copie manuelle. Ils se mettent à jour automatiquement lorsque les données référencées changent.

## Exemple de base

Créez un champ de recherche pour afficher des étiquettes des enregistrements référencés :

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Exemple avancé

Créez un champ de recherche pour extraire des valeurs de champs personnalisés des enregistrements référencés :

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ de recherche |
| `type` | CustomFieldType! | ✅ Oui | Doit être `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Oui | Configuration de recherche |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

## Configuration de recherche

### CustomFieldLookupOptionInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `referenceId` | String! | ✅ Oui | ID du champ de référence à partir duquel extraire des données |
| `lookupId` | String | Non | ID du champ personnalisé spécifique à rechercher (requis pour le type TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Oui | Type de données à extraire des enregistrements référencés |

## Types de recherche

### Valeurs de CustomFieldLookupType

| Type | Description | Renvoie |
|------|-------------|---------|
| `TODO_DUE_DATE` | Dates d'échéance des todos référencés | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Dates de création des todos référencés | Array of creation timestamps |
| `TODO_UPDATED_AT` | Dates de dernière mise à jour des todos référencés | Array of update timestamps |
| `TODO_TAG` | Étiquettes des todos référencés | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Assignés des todos référencés | Array of user objects |
| `TODO_DESCRIPTION` | Descriptions des todos référencés | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Noms des listes de todos référencés | Array of list titles |
| `TODO_CUSTOM_FIELD` | Valeurs de champs personnalisés des todos référencés | Array of values based on the field type |

## Champs de réponse

### Réponse CustomField (pour les champs de recherche)

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour le champ |
| `name` | String! | Nom d'affichage du champ de recherche |
| `type` | CustomFieldType! | Sera `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Configuration de recherche et résultats |
| `createdAt` | DateTime! | Quand le champ a été créé |
| `updatedAt` | DateTime! | Quand le champ a été mis à jour pour la dernière fois |

### Structure CustomFieldLookupOption

| Champ | Type | Description |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | Type de recherche effectuée |
| `lookupResult` | JSON | Les données extraites des enregistrements référencés |
| `reference` | CustomField | Le champ de référence utilisé comme source |
| `lookup` | CustomField | Le champ spécifique recherché (pour TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Le champ de recherche parent |
| `parentLookup` | CustomField | Recherche parent dans la chaîne (pour les recherches imbriquées) |

## Comment fonctionnent les recherches

1. **Extraction de données** : Les recherches extraient des données spécifiques de tous les enregistrements liés par un champ de référence
2. **Mises à jour automatiques** : Lorsque les enregistrements référencés changent, les valeurs de recherche se mettent à jour automatiquement
3. **Lecture seule** : Les champs de recherche ne peuvent pas être modifiés directement - ils reflètent toujours les données référencées actuelles
4. **Pas de calculs** : Les recherches extraient et affichent les données telles quelles, sans agrégations ni calculs

## Recherches TODO_CUSTOM_FIELD

Lors de l'utilisation du type `TODO_CUSTOM_FIELD`, vous devez spécifier quel champ personnalisé extraire en utilisant le paramètre `lookupId` :

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

Cela extrait les valeurs du champ personnalisé spécifié de tous les enregistrements référencés.

## Interrogation des données de recherche

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## Exemples de résultats de recherche

### Résultat de recherche d'étiquettes
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### Résultat de recherche d'assigné
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### Résultat de recherche de champ personnalisé
Les résultats varient en fonction du type de champ personnalisé recherché. Par exemple, une recherche de champ de devise pourrait renvoyer :
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## Autorisations requises

| Action | Autorisation requise |
|--------|---------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Important** : Les utilisateurs doivent avoir des autorisations de visualisation sur le projet actuel et le projet référencé pour voir les résultats de recherche.

## Réponses d'erreur

### Champ de référence invalide
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

### Recherche circulaire détectée
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### ID de recherche manquant pour TODO_CUSTOM_FIELD
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## Meilleures pratiques

1. **Nommage clair** : Utilisez des noms descriptifs qui indiquent quelles données sont recherchées
2. **Types appropriés** : Choisissez le type de recherche qui correspond à vos besoins en données
3. **Performance** : Les recherches traitent tous les enregistrements référencés, donc faites attention aux champs de référence avec de nombreux liens
4. **Autorisations** : Assurez-vous que les utilisateurs ont accès aux projets référencés pour que les recherches fonctionnent

## Cas d'utilisation courants

### Visibilité inter-projets
Affichez des étiquettes, des assignés ou des statuts des projets liés sans synchronisation manuelle.

### Suivi des dépendances
Montrez les dates d'échéance ou l'état d'achèvement des tâches dont dépend le travail actuel.

### Vue d'ensemble des ressources
Affichez tous les membres de l'équipe assignés aux tâches référencées pour la planification des ressources.

### Agrégation des statuts
Collectez tous les statuts uniques des tâches liées pour voir la santé du projet d'un coup d'œil.

## Limitations

- Les champs de recherche sont en lecture seule et ne peuvent pas être modifiés directement
- Pas de fonctions d'agrégation (SOMME, COMPTE, MOY) - les recherches n'extraient que des données
- Pas d'options de filtrage - tous les enregistrements référencés sont inclus
- Les chaînes de recherche circulaires sont empêchées pour éviter les boucles infinies
- Les résultats reflètent les données actuelles et se mettent à jour automatiquement

## Ressources connexes

- [Champs de référence](/api/custom-fields/reference) - Créez des liens vers des enregistrements pour les sources de recherche
- [Valeurs de champs personnalisés](/api/custom-fields/custom-field-values) - Définissez des valeurs sur des champs personnalisés modifiables
- [Lister les champs personnalisés](/api/custom-fields/list-custom-fields) - Interrogez tous les champs personnalisés dans un projet