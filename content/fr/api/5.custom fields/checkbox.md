---
title: Champ personnalisé de case à cocher
description: Créez des champs de case à cocher booléens pour des données oui/non ou vrai/faux
---

Les champs personnalisés de case à cocher fournissent une entrée booléenne simple (vrai/faux) pour les tâches. Ils sont parfaits pour les choix binaires, les indicateurs de statut ou le suivi de l'achèvement d'une tâche.

## Exemple de base

Créez un champ de case à cocher simple :

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de case à cocher avec description et validation :

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ Oui | Nom d'affichage de la case à cocher |
| `type` | CustomFieldType! | ✅ Oui | Doit être `CHECKBOX` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

## Définir les valeurs de case à cocher

Pour définir ou mettre à jour une valeur de case à cocher sur une tâche :

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Pour décocher une case à cocher :

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de la tâche à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de case à cocher |
| `checked` | Boolean | Non | Vrai pour cocher, faux pour décocher |

## Création de tâches avec des valeurs de case à cocher

Lors de la création d'une nouvelle tâche avec des valeurs de case à cocher :

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Valeurs de chaîne acceptées

Lors de la création de tâches, les valeurs de case à cocher doivent être passées sous forme de chaînes :

| Valeur de chaîne | Résultat |
|------------------|----------|
| `"true"` | ✅ Coché (sensible à la casse) |
| `"1"` | ✅ Coché |
| `"checked"` | ✅ Coché (sensible à la casse) |
| Any other value | ❌ Décoché |

**Remarque** : Les comparaisons de chaînes lors de la création de tâches sont sensibles à la casse. Les valeurs doivent correspondre exactement à `"true"`, `"1"`, ou `"checked"` pour aboutir à un état coché.

## Champs de réponse

### TodoCustomField Response

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | Identifiant unique pour la valeur du champ |
| `uid` | String! | Identifiant unique alternatif |
| `customField` | CustomField! | La définition du champ personnalisé |
| `checked` | Boolean | L'état de la case à cocher (vrai/faux/null) |
| `todo` | Todo! | La tâche à laquelle cette valeur appartient |
| `createdAt` | DateTime! | Date de création de la valeur |
| `updatedAt` | DateTime! | Date de dernière modification de la valeur |

## Intégration d'automatisation

Les champs de case à cocher déclenchent différents événements d'automatisation en fonction des changements d'état :

| Action | Événement déclenché | Description |
|--------|---------------------|-------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Déclenché lorsque la case à cocher est cochée |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Déclenché lorsque la case à cocher est décochée |

Cela vous permet de créer des automatisations qui répondent aux changements d'état de la case à cocher, telles que :
- Envoi de notifications lorsque des éléments sont approuvés
- Déplacement de tâches lorsque des cases à cocher de révision sont cochées
- Mise à jour des champs associés en fonction des états des cases à cocher

## Importation/Exportation de données

### Importation de valeurs de case à cocher

Lors de l'importation de données via CSV ou d'autres formats :
- `"true"`, `"yes"` → Coché (insensible à la casse)
- Toute autre valeur (y compris `"false"`, `"no"`, `"0"`, vide) → Décoché

### Exportation de valeurs de case à cocher

Lors de l'exportation de données :
- Les cases cochées s'exportent sous forme de `"X"`
- Les cases décochées s'exportent sous forme de chaîne vide `""`

## Permissions requises

| Action | Permission requise |
|--------|--------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Réponses d'erreur

### Type de valeur invalide
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Meilleures pratiques

### Conventions de nommage
- Utilisez des noms clairs et orientés vers l'action : "Approuvé", "Révisé", "Est complet"
- Évitez les noms négatifs qui peuvent confondre les utilisateurs : préférez "Est actif" à "Est inactif"
- Soyez spécifique sur ce que représente la case à cocher

### Quand utiliser des cases à cocher
- **Choix binaires** : Oui/Non, Vrai/Faux, Fait/Non fait
- **Indicateurs de statut** : Approuvé, Révisé, Publié
- **Drapeaux de fonctionnalité** : A un support prioritaire, Nécessite une signature
- **Suivi simple** : Email envoyé, Facture payée, Article expédié

### Quand NE PAS utiliser des cases à cocher
- Lorsque vous avez besoin de plus de deux options (utilisez SELECT_SINGLE à la place)
- Pour des données numériques ou textuelles (utilisez des champs NUMBER ou TEXT)
- Lorsque vous devez suivre qui l'a coché ou quand (utilisez des journaux d'audit)

## Cas d'utilisation courants

1. **Flux de travail d'approbation**
   - "Approuvé par le manager"
   - "Validation du client"
   - "Révision juridique terminée"

2. **Gestion des tâches**
   - "Est bloqué"
   - "Prêt pour révision"
   - "Haute priorité"

3. **Contrôle de qualité**
   - "QA réussi"
   - "Documentation complète"
   - "Tests écrits"

4. **Drapeaux administratifs**
   - "Facture envoyée"
   - "Contrat signé"
   - "Suivi requis"

## Limitations

- Les champs de case à cocher ne peuvent stocker que des valeurs vrai/faux (pas de tri-état ou null après le premier réglage)
- Pas de configuration de valeur par défaut (commence toujours comme null jusqu'à ce qu'il soit défini)
- Ne peut pas stocker de métadonnées supplémentaires comme qui l'a coché ou quand
- Pas de visibilité conditionnelle basée sur d'autres valeurs de champ

## Ressources connexes

- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux des champs personnalisés
- [API d'automatisations](/api/automations) - Créez des automatisations déclenchées par des changements de case à cocher