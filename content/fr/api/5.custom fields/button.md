---
title: Champ personnalisé de bouton
description: Créez des champs de bouton interactifs qui déclenchent des automatisations lorsqu'ils sont cliqués
---

Les champs personnalisés de bouton fournissent des éléments d'interface utilisateur interactifs qui déclenchent des automatisations lorsqu'ils sont cliqués. Contrairement à d'autres types de champs personnalisés qui stockent des données, les champs de bouton servent de déclencheurs d'action pour exécuter des flux de travail configurés.

## Exemple de base

Créez un champ de bouton simple qui déclenche une automatisation :

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un bouton avec des exigences de confirmation :

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du bouton |
| `type` | CustomFieldType! | ✅ Oui | Doit être `BUTTON` |
| `projectId` | String! | ✅ Oui | ID du projet où le champ sera créé |
| `buttonType` | String | Non | Comportement de confirmation (voir Types de bouton ci-dessous) |
| `buttonConfirmText` | String | Non | Texte que les utilisateurs doivent taper pour une confirmation stricte |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |
| `required` | Boolean | Non | Indique si le champ est requis (par défaut à faux) |
| `isActive` | Boolean | Non | Indique si le champ est actif (par défaut à vrai) |

### Champ de type de bouton

Le champ `buttonType` est une chaîne libre qui peut être utilisée par les clients de l'interface utilisateur pour déterminer le comportement de confirmation. Les valeurs courantes incluent :

- `""` (vide) - Pas de confirmation
- `"soft"` - Dialogue de confirmation simple
- `"hard"` - Exiger la saisie du texte de confirmation

**Remarque** : Ce ne sont que des indications pour l'interface utilisateur. L'API ne valide ni n'impose des valeurs spécifiques.

## Déclenchement des clics sur les boutons

Pour déclencher un clic sur un bouton et exécuter les automatisations associées :

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Paramètres d'entrée de clic

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de la tâche contenant le bouton |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de bouton |

### Important : Comportement de l'API

**Tous les clics sur les boutons via l'API s'exécutent immédiatement** indépendamment de tout paramètre `buttonType` ou `buttonConfirmText`. Ces champs sont stockés pour que les clients de l'interface utilisateur mettent en œuvre des dialogues de confirmation, mais l'API elle-même :

- Ne valide pas le texte de confirmation
- N'impose pas d'exigences de confirmation
- Exécute l'action du bouton immédiatement lorsqu'elle est appelée

La confirmation est purement une fonctionnalité de sécurité côté client de l'interface utilisateur.

### Exemple : Clic sur différents types de boutons

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Les trois mutations ci-dessus exécuteront l'action du bouton immédiatement lorsqu'elles sont appelées via l'API, contournant ainsi toutes les exigences de confirmation.

## Champs de réponse

### Réponse de champ personnalisé

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour le champ personnalisé |
| `name` | String! | Nom d'affichage du bouton |
| `type` | CustomFieldType! | Toujours `BUTTON` pour les champs de bouton |
| `buttonType` | String | Paramètre de comportement de confirmation |
| `buttonConfirmText` | String | Texte de confirmation requis (si utilisation de la confirmation stricte) |
| `description` | String | Texte d'aide pour les utilisateurs |
| `required` | Boolean! | Indique si le champ est requis |
| `isActive` | Boolean! | Indique si le champ est actuellement actif |
| `projectId` | String! | ID du projet auquel appartient ce champ |
| `createdAt` | DateTime! | Date de création du champ |
| `updatedAt` | DateTime! | Date de dernière modification du champ |

## Comment fonctionnent les champs de bouton

### Intégration d'automatisation

Les champs de bouton sont conçus pour fonctionner avec le système d'automatisation de Blue :

1. **Créez le champ de bouton** en utilisant la mutation ci-dessus
2. **Configurez les automatisations** qui écoutent les événements `CUSTOM_FIELD_BUTTON_CLICKED`
3. **Les utilisateurs cliquent sur le bouton** dans l'interface utilisateur
4. **Les automatisations exécutent** les actions configurées

### Flux d'événements

Lorsqu'un bouton est cliqué :

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Pas de stockage de données

Important : Les champs de bouton ne stockent aucune donnée de valeur. Ils servent uniquement de déclencheurs d'action. Chaque clic :
- Génère un événement
- Déclenche des automatisations associées
- Enregistre une action dans l'historique des tâches
- Ne modifie aucune valeur de champ

## Autorisations requises

Les utilisateurs ont besoin de rôles de projet appropriés pour créer et utiliser des champs de bouton :

| Action | Rôle requis |
|--------|-------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Réponses d'erreur

### Permission refusée
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Champ personnalisé non trouvé
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

**Remarque** : L'API ne renvoie pas d'erreurs spécifiques pour les automatisations manquantes ou les incohérences de confirmation.

## Meilleures pratiques

### Conventions de nommage
- Utilisez des noms orientés action : "Envoyer la facture", "Créer un rapport", "Notifier l'équipe"
- Soyez spécifique sur ce que fait le bouton
- Évitez les noms génériques comme "Bouton 1" ou "Cliquez ici"

### Paramètres de confirmation
- Laissez `buttonType` vide pour des actions sûres et réversibles
- Définissez `buttonType` pour suggérer un comportement de confirmation aux clients de l'interface utilisateur
- Utilisez `buttonConfirmText` pour spécifier ce que les utilisateurs doivent taper dans les confirmations de l'interface utilisateur
- Rappelez-vous : Ce ne sont que des indications pour l'interface utilisateur - les appels API s'exécutent toujours immédiatement

### Conception d'automatisation
- Gardez les actions des boutons concentrées sur un seul flux de travail
- Fournissez un retour clair sur ce qui s'est passé après le clic
- Envisagez d'ajouter un texte de description pour expliquer l'objectif du bouton

## Cas d'utilisation courants

1. **Transitions de flux de travail**
   - "Marquer comme complet"
   - "Envoyer pour approbation"
   - "Archiver la tâche"

2. **Intégrations externes**
   - "Synchroniser avec le CRM"
   - "Générer une facture"
   - "Envoyer une mise à jour par e-mail"

3. **Opérations par lots**
   - "Mettre à jour toutes les sous-tâches"
   - "Copier vers des projets"
   - "Appliquer un modèle"

4. **Actions de reporting**
   - "Générer un rapport"
   - "Exporter des données"
   - "Créer un résumé"

## Limitations

- Les boutons ne peuvent pas stocker ou afficher des valeurs de données
- Chaque bouton ne peut déclencher que des automatisations, pas des appels API directs (cependant, les automatisations peuvent inclure des actions de requête HTTP pour appeler des API externes ou les propres API de Blue)
- La visibilité des boutons ne peut pas être contrôlée de manière conditionnelle
- Maximum d'une exécution d'automatisation par clic (bien que cette automatisation puisse déclencher plusieurs actions)

## Ressources connexes

- [API des automatisations](/api/automations/index) - Configurez les actions déclenchées par les boutons
- [Aperçu des champs personnalisés](/custom-fields/list-custom-fields) - Concepts généraux des champs personnalisés