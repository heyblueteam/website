---
title: Champ personnalisé de téléphone
description: Créez des champs de téléphone pour stocker et valider les numéros de téléphone avec un format international
---

Les champs personnalisés de téléphone vous permettent de stocker des numéros de téléphone dans des enregistrements avec une validation intégrée et un format international. Ils sont idéaux pour suivre les informations de contact, les contacts d'urgence ou toute donnée liée au téléphone dans vos projets.

## Exemple de base

Créez un champ de téléphone simple :

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de téléphone avec une description :

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ Oui | Nom d'affichage du champ de téléphone |
| `type` | CustomFieldType! | ✅ Oui | Doit être `PHONE` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Les champs personnalisés sont automatiquement associés au projet en fonction du contexte de projet actuel de l'utilisateur. Aucun paramètre `projectId` n'est requis.

## Définir des valeurs de téléphone

Pour définir ou mettre à jour une valeur de téléphone dans un enregistrement :

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de téléphone |
| `text` | String | Non | Numéro de téléphone avec code pays |
| `regionCode` | String | Non | Code pays (détecté automatiquement) |

**Remarque** : Bien que `text` soit optionnel dans le schéma, un numéro de téléphone est requis pour que le champ ait un sens. Lors de l'utilisation de `setTodoCustomField`, aucune validation n'est effectuée - vous pouvez stocker n'importe quelle valeur texte et regionCode. La détection automatique n'a lieu que lors de la création de l'enregistrement.

## Création d'enregistrements avec des valeurs de téléphone

Lors de la création d'un nouvel enregistrement avec des valeurs de téléphone :

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
    }
  }
}
```

## Champs de réponse

### TodoCustomField Response

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `text` | String | Le numéro de téléphone formaté (format international) |
| `regionCode` | String | Le code pays (par exemple, "US", "GB", "CA") |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

## Validation des numéros de téléphone

**Important** : La validation et le formatage des numéros de téléphone n'ont lieu que lors de la création de nouveaux enregistrements via `createTodo`. Lors de la mise à jour des valeurs de téléphone existantes en utilisant `setTodoCustomField`, aucune validation n'est effectuée et les valeurs sont stockées telles qu'elles sont fournies.

### Formats acceptés (lors de la création d'enregistrement)
Les numéros de téléphone doivent inclure un code pays dans l'un de ces formats :

- **Format E.164 (préféré)** : `+12345678900`
- **Format international** : `+1 234 567 8900`
- **International avec ponctuation** : `+1 (234) 567-8900`
- **Code pays avec tirets** : `+1-234-567-8900`

**Remarque** : Les formats nationaux sans code pays (comme `(234) 567-8900`) seront rejetés lors de la création de l'enregistrement.

### Règles de validation (lors de la création d'enregistrement)
- Utilise libphonenumber-js pour l'analyse et la validation
- Accepte divers formats de numéros de téléphone internationaux
- Détecte automatiquement le pays à partir du numéro
- Formate le numéro au format d'affichage international (par exemple, `+1 234 567 8900`)
- Extrait et stocke le code pays séparément (par exemple, `US`)

### Exemples de numéros de téléphone valides
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Exemples de numéros de téléphone invalides
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Format de stockage

Lors de la création d'enregistrements avec des numéros de téléphone :
- **text** : Stocké au format international (par exemple, `+1 234 567 8900`) après validation
- **regionCode** : Stocké comme code pays ISO (par exemple, `US`, `GB`, `CA`) détecté automatiquement

Lors de la mise à jour via `setTodoCustomField` :
- **text** : Stocké exactement tel que fourni (sans formatage)
- **regionCode** : Stocké exactement tel que fourni (sans validation)

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Réponses d'erreur

### Format de téléphone invalide
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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

### Code pays manquant
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Meilleures pratiques

### Saisie des données
- Incluez toujours le code pays dans les numéros de téléphone
- Utilisez le format E.164 pour la cohérence
- Validez les numéros avant de les stocker pour des opérations importantes
- Tenez compte des préférences régionales pour le format d'affichage

### Qualité des données
- Stockez les numéros au format international pour une compatibilité mondiale
- Utilisez regionCode pour des fonctionnalités spécifiques au pays
- Validez les numéros de téléphone avant des opérations critiques (SMS, appels)
- Tenez compte des implications de fuseau horaire pour le timing des contacts

### Considérations internationales
- Le code pays est détecté et stocké automatiquement
- Les numéros sont formatés selon la norme internationale
- Les préférences d'affichage régionales peuvent utiliser regionCode
- Tenez compte des conventions de numérotation locales lors de l'affichage

## Cas d'utilisation courants

1. **Gestion des contacts**
   - Numéros de téléphone des clients
   - Informations de contact des fournisseurs
   - Numéros de téléphone des membres de l'équipe
   - Détails de contact du support

2. **Contacts d'urgence**
   - Numéros de contact d'urgence
   - Informations de contact d'astreinte
   - Contacts pour la réponse aux crises
   - Numéros de téléphone d'escalade

3. **Support client**
   - Numéros de téléphone des clients
   - Numéros de rappel du support
   - Numéros de téléphone de vérification
   - Numéros de contact pour le suivi

4. **Ventes et marketing**
   - Numéros de téléphone des prospects
   - Listes de contacts de campagne
   - Informations de contact des partenaires
   - Téléphones des sources de référence

## Fonctionnalités d'intégration

### Avec des automatisations
- Déclenchez des actions lorsque les champs de téléphone sont mis à jour
- Envoyez des notifications SMS aux numéros de téléphone stockés
- Créez des tâches de suivi basées sur les changements de téléphone
- Acheminez les appels en fonction des données de numéro de téléphone

### Avec des recherches
- Référencez les données téléphoniques d'autres enregistrements
- Agrégez des listes de téléphones provenant de plusieurs sources
- Trouvez des enregistrements par numéro de téléphone
- Croisez les informations de contact

### Avec des formulaires
- Validation automatique des téléphones
- Vérification du format international
- Détection du code pays
- Retour d'information en temps réel sur le format

## Limitations

- Nécessite un code pays pour tous les numéros
- Pas de capacités SMS ou d'appel intégrées
- Pas de vérification des numéros de téléphone au-delà de la vérification du format
- Pas de stockage des métadonnées téléphoniques (opérateur, type, etc.)
- Les numéros au format national sans code pays sont rejetés
- Pas de formatage automatique des numéros de téléphone dans l'interface utilisateur au-delà de la norme internationale

## Ressources connexes

- [Champs de texte](/api/custom-fields/text-single) - Pour les données textuelles non téléphoniques
- [Champs d'email](/api/custom-fields/email) - Pour les adresses email
- [Champs d'URL](/api/custom-fields/url) - Pour les adresses de site web
- [Aperçu des champs personnalisés](/custom-fields/list-custom-fields) - Concepts généraux