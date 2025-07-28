---
title: Champ personnalisé d'email
description: Créez des champs d'email pour stocker et valider des adresses email
---

Les champs personnalisés d'email vous permettent de stocker des adresses email dans des enregistrements avec une validation intégrée. Ils sont idéaux pour suivre les informations de contact, les emails des assignés ou toute donnée liée aux emails dans vos projets.

## Exemple de base

Créez un champ d'email simple :

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ d'email avec description :

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ Oui | Nom d'affichage du champ d'email |
| `type` | CustomFieldType! | ✅ Oui | Doit être `EMAIL` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

## Définir des valeurs d'email

Pour définir ou mettre à jour une valeur d'email sur un enregistrement :

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé d'email |
| `text` | String | Non | Adresse email à stocker |

## Création d'enregistrements avec des valeurs d'email

Lors de la création d'un nouvel enregistrement avec des valeurs d'email :

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Champs de réponse

### Réponse CustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | Identifiant unique pour le champ personnalisé |
| `name` | String! | Nom d'affichage du champ d'email |
| `type` | CustomFieldType! | Le type de champ (EMAIL) |
| `description` | String | Texte d'aide pour le champ |
| `value` | JSON | Contient la valeur d'email (voir ci-dessous) |
| `createdAt` | DateTime! | Date de création du champ |
| `updatedAt` | DateTime! | Date de dernière modification du champ |

**Important** : Les valeurs d'email sont accessibles via le champ `customField.value.text`, et non directement dans la réponse.

## Interrogation des valeurs d'email

Lors de l'interrogation d'enregistrements avec des champs personnalisés d'email, accédez à l'email via le chemin `customField.value.text` :

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

La réponse inclura l'email dans la structure imbriquée :

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## Validation des emails

### Validation de formulaire
Lorsque les champs d'email sont utilisés dans des formulaires, ils valident automatiquement le format de l'email :
- Utilise des règles de validation d'email standard
- Supprime les espaces vides de l'entrée
- Rejette les formats d'email invalides

### Règles de validation
- Doit contenir un symbole `@`
- Doit avoir un format de domaine valide
- Les espaces vides en début/fin sont automatiquement supprimés
- Les formats d'email courants sont acceptés

### Exemples d'emails valides
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Exemples d'emails invalides
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Notes importantes

### API directe vs formulaires
- **Formulaires** : La validation automatique des emails est appliquée
- **API directe** : Pas de validation - tout texte peut être stocké
- **Recommandation** : Utilisez des formulaires pour les entrées utilisateur afin d'assurer la validation

### Format de stockage
- Les adresses email sont stockées en texte brut
- Pas de formatage ou d'analyse spéciale
- Sensibilité à la casse : les champs personnalisés d'email sont stockés de manière sensible à la casse (contrairement aux emails d'authentification utilisateur qui sont normalisés en minuscules)
- Pas de limitations de longueur maximale au-delà des contraintes de la base de données (limite de 16 Mo)

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Réponses d'erreur

### Format d'email invalide (formulaires uniquement)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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
      "code": "NOT_FOUND"
    }
  }]
}
```

## Meilleures pratiques

### Saisie des données
- Validez toujours les adresses email dans votre application
- Utilisez des champs d'email uniquement pour de véritables adresses email
- Envisagez d'utiliser des formulaires pour les entrées utilisateur afin d'obtenir une validation automatique

### Qualité des données
- Supprimez les espaces vides avant de stocker
- Envisagez la normalisation de la casse (typiquement en minuscules)
- Validez le format de l'email avant des opérations importantes

### Considérations de confidentialité
- Les adresses email sont stockées en texte brut
- Prenez en compte les réglementations sur la confidentialité des données (GDPR, CCPA)
- Mettez en œuvre des contrôles d'accès appropriés

## Cas d'utilisation courants

1. **Gestion des contacts**
   - Adresses email des clients
   - Informations de contact des fournisseurs
   - Emails des membres de l'équipe
   - Détails de contact du support

2. **Gestion de projet**
   - Emails des parties prenantes
   - Emails de contact pour approbation
   - Destinataires de notifications
   - Emails des collaborateurs externes

3. **Support client**
   - Adresses email des clients
   - Contacts des tickets de support
   - Contacts d'escalade
   - Adresses email de retour d'information

4. **Ventes et marketing**
   - Adresses email des prospects
   - Listes de contacts pour les campagnes
   - Informations de contact des partenaires
   - Emails des sources de référence

## Fonctionnalités d'intégration

### Avec des automatisations
- Déclenchez des actions lorsque les champs d'email sont mis à jour
- Envoyez des notifications aux adresses email stockées
- Créez des tâches de suivi basées sur les changements d'email

### Avec des recherches
- Référencez les données d'email d'autres enregistrements
- Agrégez des listes d'emails de plusieurs sources
- Trouvez des enregistrements par adresse email

### Avec des formulaires
- Validation automatique des emails
- Vérification du format d'email
- Suppression des espaces vides

## Limitations

- Pas de vérification ou de validation d'email intégrée au-delà de la vérification du format
- Pas de fonctionnalités UI spécifiques aux emails (comme des liens email cliquables)
- Stocké en texte brut sans cryptage
- Pas de capacités de composition ou d'envoi d'emails
- Pas de stockage de métadonnées d'email (nom d'affichage, etc.)
- Les appels API directs contournent la validation (seuls les formulaires valident)

## Ressources connexes

- [Champs de texte](/api/custom-fields/text-single) - Pour les données textuelles non-email
- [Champs d'URL](/api/custom-fields/url) - Pour les adresses de sites web
- [Champs de téléphone](/api/custom-fields/phone) - Pour les numéros de téléphone
- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux