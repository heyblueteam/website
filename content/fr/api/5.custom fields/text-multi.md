---
title: Champ personnalisé de texte multi-lignes
description: Créez des champs de texte multi-lignes pour un contenu plus long comme des descriptions, des notes et des commentaires
---

Les champs personnalisés de texte multi-lignes vous permettent de stocker un contenu textuel plus long avec des sauts de ligne et un formatage. Ils sont idéaux pour des descriptions, des notes, des commentaires ou toute donnée textuelle nécessitant plusieurs lignes.

## Exemple de base

Créez un simple champ de texte multi-lignes :

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de texte multi-lignes avec description :

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `name` | String! | ✅ Oui | Nom d'affichage du champ de texte |
| `type` | CustomFieldType! | ✅ Oui | Doit être `TEXT_MULTI` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque :** Le `projectId` est passé comme un argument séparé à la mutation, et non comme partie de l'objet d'entrée. Alternativement, le contexte du projet peut être déterminé à partir de l'en-tête `X-Bloo-Project-ID` dans votre requête GraphQL.

## Définir des valeurs textuelles

Pour définir ou mettre à jour une valeur de texte multi-lignes sur un enregistrement :

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ de texte personnalisé |
| `text` | String | Non | Contenu de texte multi-lignes à stocker |

## Création d'enregistrements avec des valeurs textuelles

Lors de la création d'un nouvel enregistrement avec des valeurs de texte multi-lignes :

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
| `text` | String | Le contenu de texte multi-lignes stocké |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

## Validation du texte

### Validation de formulaire
Lorsque des champs de texte multi-lignes sont utilisés dans des formulaires :
- Les espaces vides en début et en fin sont automatiquement supprimés
- La validation requise est appliquée si le champ est marqué comme requis
- Aucune validation de format spécifique n'est appliquée

### Règles de validation
- Accepte tout contenu de chaîne y compris les sauts de ligne
- Pas de limites de longueur de caractère (jusqu'aux limites de la base de données)
- Prend en charge les caractères Unicode et les symboles spéciaux
- Les sauts de ligne sont préservés dans le stockage

### Exemples de texte valide
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Remarques importantes

### Capacité de stockage
- Stocké en utilisant le type `MediumText` de MySQL
- Prend en charge jusqu'à 16 Mo de contenu textuel
- Les sauts de ligne et le formatage sont préservés
- Encodage UTF-8 pour les caractères internationaux

### API directe vs formulaires
- **Formulaires** : Suppression automatique des espaces vides et validation requise
- **API directe** : Le texte est stocké exactement tel que fourni
- **Recommandation** : Utilisez des formulaires pour l'entrée utilisateur afin d'assurer un formatage cohérent

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI** : Saisie de texte multi-lignes, idéale pour un contenu plus long
- **TEXT_SINGLE** : Saisie de texte sur une seule ligne, idéale pour des valeurs courtes
- **Backend** : Les deux types sont identiques - même champ de stockage, validation et traitement
- **Frontend** : Différents composants UI pour la saisie de données (zone de texte vs champ de saisie)
- **Important** : La distinction entre TEXT_MULTI et TEXT_SINGLE existe uniquement à des fins d'interface utilisateur

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

## Réponses d'erreur

### Validation de champ requis (formulaires uniquement)
```json
{
  "errors": [{
    "message": "This field is required",
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
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Meilleures pratiques

### Organisation du contenu
- Utilisez un formatage cohérent pour un contenu structuré
- Envisagez d'utiliser une syntaxe similaire à Markdown pour la lisibilité
- Divisez un long contenu en sections logiques
- Utilisez des sauts de ligne pour améliorer la lisibilité

### Saisie de données
- Fournissez des descriptions de champ claires pour guider les utilisateurs
- Utilisez des formulaires pour l'entrée utilisateur afin d'assurer la validation
- Envisagez des limites de caractères en fonction de votre cas d'utilisation
- Validez le format du contenu dans votre application si nécessaire

### Considérations de performance
- Un contenu textuel très long peut affecter les performances des requêtes
- Envisagez la pagination pour afficher de grands champs de texte
- Considérations d'index pour la fonctionnalité de recherche
- Surveillez l'utilisation du stockage pour les champs avec un contenu volumineux

## Filtrage et recherche

### Recherche par contenu
Les champs de texte multi-lignes prennent en charge la recherche de sous-chaînes via des filtres de champ personnalisé :

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Capacités de recherche
- Correspondance de sous-chaîne dans les champs de texte en utilisant l'opérateur `CONTAINS`
- Recherche insensible à la casse en utilisant l'opérateur `NCONTAINS`
- Correspondance exacte en utilisant l'opérateur `IS`
- Correspondance négative en utilisant l'opérateur `NOT`
- Recherches à travers toutes les lignes de texte
- Prend en charge la correspondance de mots partiels

## Cas d'utilisation courants

1. **Gestion de projet**
   - Descriptions de tâches
   - Exigences du projet
   - Notes de réunion
   - Mises à jour de statut

2. **Support client**
   - Descriptions de problèmes
   - Notes de résolution
   - Retours clients
   - Journaux de communication

3. **Gestion de contenu**
   - Contenu d'article
   - Descriptions de produits
   - Commentaires d'utilisateurs
   - Détails des avis

4. **Documentation**
   - Descriptions de processus
   - Instructions
   - Directives
   - Matériaux de référence

## Fonctionnalités d'intégration

### Avec des automatisations
- Déclencher des actions lorsque le contenu textuel change
- Extraire des mots-clés du contenu textuel
- Créer des résumés ou des notifications
- Traiter le contenu textuel avec des services externes

### Avec des recherches
- Référencer des données textuelles d'autres enregistrements
- Agréger du contenu textuel de plusieurs sources
- Trouver des enregistrements par contenu textuel
- Afficher des informations textuelles connexes

### Avec des formulaires
- Suppression automatique des espaces vides
- Validation des champs requis
- UI de zone de texte multi-lignes
- Affichage du nombre de caractères (si configuré)

## Limitations

- Pas de formatage de texte intégré ou d'édition de texte enrichi
- Pas de détection ou de conversion automatique de liens
- Pas de vérification orthographique ou de validation grammaticale
- Pas d'analyse ou de traitement de texte intégré
- Pas de versioning ou de suivi des modifications
- Capacités de recherche limitées (pas de recherche en texte intégral)
- Pas de compression de contenu pour des textes très volumineux

## Ressources connexes

- [Champs de texte sur une seule ligne](/api/custom-fields/text-single) - Pour des valeurs de texte courtes
- [Champs d'email](/api/custom-fields/email) - Pour des adresses email
- [Champs d'URL](/api/custom-fields/url) - Pour des adresses de sites web
- [Aperçu des champs personnalisés](/api/custom-fields/2.list-custom-fields) - Concepts généraux