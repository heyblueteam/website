---
title: Champ personnalisé de texte sur une seule ligne
description: Créez des champs de texte sur une seule ligne pour des valeurs de texte courtes telles que des noms, des titres et des étiquettes
---

Les champs personnalisés de texte sur une seule ligne vous permettent de stocker des valeurs de texte courtes destinées à une saisie sur une seule ligne. Ils sont idéaux pour les noms, les titres, les étiquettes ou toute donnée textuelle qui doit être affichée sur une seule ligne.

## Exemple de base

Créez un simple champ de texte sur une seule ligne :

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de texte sur une seule ligne avec une description :

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
| `name` | String! | ✅ Oui | Nom d'affichage du champ de texte |
| `type` | CustomFieldType! | ✅ Oui | Doit être `TEXT_SINGLE` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Le contexte du projet est automatiquement déterminé à partir de vos en-têtes d'authentification. Aucun paramètre `projectId` n'est nécessaire.

## Définir des valeurs textuelles

Pour définir ou mettre à jour une valeur de texte sur une seule ligne sur un enregistrement :

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### Paramètres SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ de texte personnalisé |
| `text` | String | Non | Contenu de texte sur une seule ligne à stocker |

## Création d'enregistrements avec des valeurs textuelles

Lors de la création d'un nouvel enregistrement avec des valeurs de texte sur une seule ligne :

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé (contient la valeur textuelle) |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Date de création de la valeur |
| `updatedAt` | DateTime! | Date de dernière modification de la valeur |

**Important** : Les valeurs textuelles sont accessibles via le champ `customField.value.text`, et non directement sur TodoCustomField.

## Interrogation des valeurs textuelles

Lors de l'interrogation d'enregistrements avec des champs de texte personnalisés, accédez au texte via le chemin `customField.value.text` :

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

La réponse inclura le texte dans la structure imbriquée :

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Validation du texte

### Validation de formulaire
Lorsque les champs de texte sur une seule ligne sont utilisés dans des formulaires :
- Les espaces vides au début et à la fin sont automatiquement supprimés
- La validation requise est appliquée si le champ est marqué comme requis
- Aucune validation de format spécifique n'est appliquée

### Règles de validation
- Accepte tout contenu de chaîne, y compris les sauts de ligne (bien que non recommandé)
- Pas de limites de longueur de caractère (jusqu'aux limites de la base de données)
- Prend en charge les caractères Unicode et les symboles spéciaux
- Les sauts de ligne sont préservés mais ne sont pas destinés à ce type de champ

### Exemples typiques de texte
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Remarques importantes

### Capacité de stockage
- Stocké en utilisant le type MySQL `MediumText`
- Prend en charge jusqu'à 16 Mo de contenu textuel
- Stockage identique aux champs de texte multi-lignes
- Encodage UTF-8 pour les caractères internationaux

### API directe vs formulaires
- **Formulaires** : Suppression automatique des espaces vides et validation requise
- **API directe** : Le texte est stocké exactement tel que fourni
- **Recommandation** : Utilisez des formulaires pour la saisie utilisateur afin d'assurer un formatage cohérent

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE** : Saisie de texte sur une seule ligne, idéale pour des valeurs courtes
- **TEXT_MULTI** : Saisie de texte multi-lignes, idéale pour un contenu plus long
- **Backend** : Les deux utilisent un stockage et une validation identiques
- **Frontend** : Différents composants UI pour la saisie de données
- **Intention** : TEXT_SINGLE est sémantiquement destiné aux valeurs sur une seule ligne

## Autorisations requises

| Action | Autorisation requise |
|--------|----------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

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

### Champ introuvable
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

## Meilleures pratiques

### Directives de contenu
- Gardez le texte concis et approprié pour une seule ligne
- Évitez les sauts de ligne pour un affichage prévu sur une seule ligne
- Utilisez un formatage cohérent pour des types de données similaires
- Considérez les limites de caractères en fonction de vos exigences UI

### Saisie de données
- Fournissez des descriptions claires des champs pour guider les utilisateurs
- Utilisez des formulaires pour la saisie utilisateur afin d'assurer la validation
- Validez le format du contenu dans votre application si nécessaire
- Envisagez d'utiliser des listes déroulantes pour des valeurs standardisées

### Considérations de performance
- Les champs de texte sur une seule ligne sont légers et performants
- Envisagez l'indexation pour les champs fréquemment recherchés
- Utilisez des largeurs d'affichage appropriées dans votre UI
- Surveillez la longueur du contenu à des fins d'affichage

## Filtrage et recherche

### Recherche par contenu
Les champs de texte sur une seule ligne prennent en charge la recherche de sous-chaînes :

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Capacités de recherche
- Correspondance de sous-chaînes insensible à la casse
- Prend en charge la correspondance de mots partiels
- Correspondance exacte des valeurs
- Pas de recherche en texte intégral ou de classement

## Cas d'utilisation courants

1. **Identifiants et codes**
   - SKU de produit
   - Numéros de commande
   - Codes de référence
   - Numéros de version

2. **Noms et titres**
   - Noms de clients
   - Titres de projet
   - Noms de produits
   - Étiquettes de catégorie

3. **Descriptions courtes**
   - Résumés brefs
   - Étiquettes de statut
   - Indicateurs de priorité
   - Étiquettes de classification

4. **Références externes**
   - Numéros de ticket
   - Références de facture
   - ID de systèmes externes
   - Numéros de document

## Fonctionnalités d'intégration

### Avec des recherches
- Référencez des données textuelles d'autres enregistrements
- Trouvez des enregistrements par contenu textuel
- Affichez des informations textuelles connexes
- Agrégez des valeurs textuelles de plusieurs sources

### Avec des formulaires
- Suppression automatique des espaces vides
- Validation de champ requise
- UI de saisie de texte sur une seule ligne
- Affichage de la limite de caractères (si configuré)

### Avec des importations/exportations
- Mappage direct des colonnes CSV
- Attribution automatique des valeurs textuelles
- Prise en charge de l'importation de données en masse
- Exportation vers des formats de feuille de calcul

## Limitations

### Restrictions d'automatisation
- Pas directement disponible en tant que champs de déclenchement d'automatisation
- Ne peut pas être utilisé dans les mises à jour de champs d'automatisation
- Peut être référencé dans les conditions d'automatisation
- Disponible dans les modèles d'e-mail et les webhooks

### Limitations générales
- Pas de formatage ou de style de texte intégré
- Pas de validation automatique au-delà des champs requis
- Pas d'application de l'unicité intégrée
- Pas de compression de contenu pour des textes très longs
- Pas de versionnage ou de suivi des modifications
- Capacités de recherche limitées (pas de recherche en texte intégral)

## Ressources connexes

- [Champs de texte multi-lignes](/api/custom-fields/text-multi) - Pour un contenu textuel plus long
- [Champs d'e-mail](/api/custom-fields/email) - Pour les adresses e-mail
- [Champs d'URL](/api/custom-fields/url) - Pour les adresses de site web
- [Champs d'ID uniques](/api/custom-fields/unique-id) - Pour les identifiants générés automatiquement
- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux