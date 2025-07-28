---
title: Champ personnalisé d'ID unique
description: Créez des champs d'identifiant unique auto-générés avec numérotation séquentielle et formatage personnalisé
---

Les champs personnalisés d'ID unique génèrent automatiquement des identifiants uniques et séquentiels pour vos enregistrements. Ils sont parfaits pour créer des numéros de ticket, des ID de commande, des numéros de facture ou tout système d'identifiant séquentiel dans votre flux de travail.

## Exemple de base

Créez un simple champ d'ID unique avec auto-séquençage :

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Exemple avancé

Créez un champ d'ID unique formaté avec préfixe et remplissage de zéros :

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput (UNIQUE_ID)

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ d'ID unique |
| `type` | CustomFieldType! | ✅ Oui | Doit être `UNIQUE_ID` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |
| `useSequenceUniqueId` | Boolean | Non | Activer l'auto-séquençage (par défaut : faux) |
| `prefix` | String | Non | Préfixe de texte pour les ID générés (ex. : "TÂCHE-") |
| `sequenceDigits` | Int | Non | Nombre de chiffres pour le remplissage de zéros |
| `sequenceStartingNumber` | Int | Non | Nombre de départ pour la séquence |

## Options de configuration

### Auto-Séquençage (`useSequenceUniqueId`)
- **true** : Génère automatiquement des ID séquentiels lors de la création d'enregistrements
- **false** ou **undefined** : Saisie manuelle requise (fonctionne comme un champ de texte)

### Préfixe (`prefix`)
- Préfixe de texte optionnel ajouté à tous les ID générés
- Exemples : "TÂCHE-", "ORD-", "BUG-", "REQ-"
- Pas de limite de longueur, mais gardez-le raisonnable pour l'affichage

### Chiffres de séquence (`sequenceDigits`)
- Nombre de chiffres pour le remplissage de zéros du numéro de séquence
- Exemple : `sequenceDigits: 3` produit `001`, `002`, `003`
- Si non spécifié, aucun remplissage n'est appliqué

### Nombre de départ (`sequenceStartingNumber`)
- Le premier numéro de la séquence
- Exemple : `sequenceStartingNumber: 1000` commence à 1000, 1001, 1002...
- Si non spécifié, commence à 1 (comportement par défaut)

## Format d'ID généré

Le format final de l'ID combine toutes les options de configuration :

```
{prefix}{paddedSequenceNumber}
```

### Exemples de format

| Configuration | IDs générés |
|---------------|-------------|
| Aucune option | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Lecture des valeurs d'ID unique

### Interroger des enregistrements avec des ID uniques
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Format de réponse
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Génération automatique d'ID

### Quand les ID sont générés
- **Création d'enregistrement** : Les ID sont automatiquement attribués lors de la création de nouveaux enregistrements
- **Ajout de champ** : Lors de l'ajout d'un champ UNIQUE_ID à des enregistrements existants, un travail en arrière-plan est mis en file d'attente (implémentation du travail en attente)
- **Traitement en arrière-plan** : La génération d'ID pour les nouveaux enregistrements se fait de manière synchrone via des déclencheurs de base de données

### Processus de génération
1. **Déclencheur** : Un nouvel enregistrement est créé ou un champ UNIQUE_ID est ajouté
2. **Recherche de séquence** : Le système trouve le prochain numéro de séquence disponible
3. **Attribution d'ID** : Le numéro de séquence est attribué à l'enregistrement
4. **Mise à jour du compteur** : Le compteur de séquence est incrémenté pour les futurs enregistrements
5. **Formatage** : L'ID est formaté avec préfixe et remplissage lors de l'affichage

### Garanties d'unicité
- **Contraintes de base de données** : Contrainte d'unicité sur les ID de séquence dans chaque champ
- **Opérations atomiques** : La génération de séquence utilise des verrous de base de données pour éviter les doublons
- **Portée de projet** : Les séquences sont indépendantes par projet
- **Protection contre les conditions de concurrence** : Les demandes simultanées sont traitées en toute sécurité

## Mode manuel vs automatique

### Mode automatique (`useSequenceUniqueId: true`)
- Les ID sont générés automatiquement via des déclencheurs de base de données
- La numérotation séquentielle est garantie
- La génération de séquence atomique empêche les doublons
- Les ID formatés combinent préfixe + numéro de séquence rempli

### Mode manuel (`useSequenceUniqueId: false` ou `undefined`)
- Fonctionne comme un champ de texte normal
- Les utilisateurs peuvent saisir des valeurs personnalisées via `setTodoCustomField` avec le paramètre `text`
- Pas de génération automatique
- Pas d'application de l'unicité au-delà des contraintes de base de données

## Définir des valeurs manuelles (Mode manuel uniquement)

Lorsque `useSequenceUniqueId` est faux, vous pouvez définir des valeurs manuellement :

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Champs de réponse

### Réponse TodoCustomField (UNIQUE_ID)

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `sequenceId` | Int | Le numéro de séquence généré (rempli pour les champs UNIQUE_ID) |
| `text` | String | La valeur de texte formatée (combine préfixe + séquence remplie) |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été mise à jour pour la dernière fois |

### Réponse CustomField (UNIQUE_ID)

| Champ | Type | Description |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | Indique si l'auto-séquençage est activé |
| `prefix` | String | Préfixe de texte pour les ID générés |
| `sequenceDigits` | Int | Nombre de chiffres pour le remplissage de zéros |
| `sequenceStartingNumber` | Int | Nombre de départ pour la séquence |

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Réponses d'erreur

### Erreur de configuration de champ
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Erreur de permission
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

## Remarques importantes

### IDs auto-générés
- **Lecture seule** : Les IDs auto-générés ne peuvent pas être modifiés manuellement
- **Permanent** : Une fois attribués, les IDs de séquence ne changent pas
- **Chronologique** : Les IDs reflètent l'ordre de création
- **Scoped** : Les séquences sont indépendantes par projet

### Considérations de performance
- La génération d'ID pour de nouveaux enregistrements est synchrone via des déclencheurs de base de données
- La génération de séquence utilise des verrous `FOR UPDATE` pour des opérations atomiques
- Un système de travail en arrière-plan existe mais l'implémentation du travail est en attente
- Considérez les numéros de départ de séquence pour les projets à fort volume

### Migration et mises à jour
- Ajouter l'auto-séquençage à des enregistrements existants met en file d'attente un travail en arrière-plan (travail en attente)
- Changer les paramètres de séquence n'affecte que les futurs enregistrements
- Les IDs existants restent inchangés lors des mises à jour de configuration
- Les compteurs de séquence continuent à partir du maximum actuel

## Meilleures pratiques

### Conception de configuration
- Choisissez des préfixes descriptifs qui ne seront pas en conflit avec d'autres systèmes
- Utilisez un remplissage de chiffres approprié pour votre volume attendu
- Définissez des numéros de départ raisonnables pour éviter les conflits
- Testez la configuration avec des données d'exemple avant le déploiement

### Directives sur les préfixes
- Gardez les préfixes courts et mémorables (2-5 caractères)
- Utilisez des majuscules pour la cohérence
- Incluez des séparateurs (traits d'union, soulignés) pour la lisibilité
- Évitez les caractères spéciaux qui pourraient causer des problèmes dans les URLs ou les systèmes

### Planification de séquence
- Estimez votre volume d'enregistrements pour choisir un remplissage de chiffres approprié
- Considérez la croissance future lors de la définition des numéros de départ
- Planifiez des plages de séquence différentes pour différents types d'enregistrements
- Documentez vos schémas d'ID pour référence de l'équipe

## Cas d'utilisation courants

1. **Systèmes de support**
   - Numéros de ticket : `TICK-001`, `TICK-002`
   - ID de cas : `CASE-2024-001`
   - Demandes de support : `SUP-001`

2. **Gestion de projet**
   - IDs de tâche : `TASK-001`, `TASK-002`
   - Éléments de sprint : `SPRINT-001`
   - Numéros de livrables : `DEL-001`

3. **Opérations commerciales**
   - Numéros de commande : `ORD-2024-001`
   - IDs de facture : `INV-001`
   - Commandes d'achat : `PO-001`

4. **Gestion de la qualité**
   - Rapports de bogues : `BUG-001`
   - IDs de cas de test : `TEST-001`
   - Numéros de révision : `REV-001`

## Fonctionnalités d'intégration

### Avec des automatisations
- Déclencher des actions lorsque des ID uniques sont attribués
- Utiliser des modèles d'ID dans les règles d'automatisation
- Référencer des IDs dans des modèles d'e-mail et des notifications

### Avec des recherches
- Référencer des IDs uniques à partir d'autres enregistrements
- Trouver des enregistrements par ID unique
- Afficher les identifiants d'enregistrements liés

### Avec des rapports
- Regrouper et filtrer par modèles d'ID
- Suivre les tendances d'attribution d'ID
- Surveiller l'utilisation et les lacunes de séquence

## Limitations

- **Séquentiel uniquement** : Les ID sont attribués dans l'ordre chronologique
- **Pas de lacunes** : Les enregistrements supprimés laissent des lacunes dans les séquences
- **Pas de réutilisation** : Les numéros de séquence ne sont jamais réutilisés
- **Portée projet** : Ne peut pas partager des séquences entre projets
- **Contraintes de format** : Options de formatage limitées
- **Pas de mises à jour en masse** : Impossible de mettre à jour en masse les IDs de séquence existants
- **Pas de logique personnalisée** : Impossible de mettre en œuvre des règles de génération d'ID personnalisées

## Ressources connexes

- [Champs de texte](/api/custom-fields/text-single) - Pour des identifiants de texte manuels
- [Champs numériques](/api/custom-fields/number) - Pour des séquences numériques
- [Aperçu des champs personnalisés](/api/custom-fields/2.list-custom-fields) - Concepts généraux
- [Automatisations](/api/automations) - Pour des règles d'automatisation basées sur des ID