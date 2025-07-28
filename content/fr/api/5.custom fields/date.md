---
title: Champ personnalisé de date
description: Créez des champs de date pour suivre des dates uniques ou des plages de dates avec prise en charge des fuseaux horaires
---

Les champs personnalisés de date vous permettent de stocker des dates uniques ou des plages de dates pour des enregistrements. Ils prennent en charge la gestion des fuseaux horaires, le formatage intelligent et peuvent être utilisés pour suivre les délais, les dates d'événements ou toute information basée sur le temps.

## Exemple de base

Créez un champ de date simple :

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de date d'échéance avec description :

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom affiché du champ de date |
| `type` | CustomFieldType! | ✅ Oui | Doit être `DATE` |
| `isDueDate` | Boolean | Non | Indique si ce champ représente une date d'échéance |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Les champs personnalisés sont automatiquement associés au projet en fonction du contexte de projet actuel de l'utilisateur. Aucun `projectId` n'est requis.

## Définir les valeurs de date

Les champs de date peuvent stocker soit une date unique, soit une plage de dates :

### Date unique

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Plage de dates

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Événement toute la journée

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de date |
| `startDate` | DateTime | Non | Date/heure de début au format ISO 8601 |
| `endDate` | DateTime | Non | Date/heure de fin au format ISO 8601 |
| `timezone` | String | Non | Identifiant de fuseau horaire (par exemple, "America/New_York") |

**Remarque** : Si seul `startDate` est fourni, `endDate` par défaut automatiquement à la même valeur.

## Formats de date

### Format ISO 8601
Toutes les dates doivent être fournies au format ISO 8601 :
- `2025-01-15T14:30:00Z` - Heure UTC
- `2025-01-15T14:30:00+05:00` - Avec décalage de fuseau horaire
- `2025-01-15T14:30:00.123Z` - Avec millisecondes

### Identifiants de fuseau horaire
Utilisez des identifiants de fuseau horaire standard :
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Si aucun fuseau horaire n'est fourni, le système par défaut au fuseau horaire détecté de l'utilisateur.

## Création d'enregistrements avec des valeurs de date

Lors de la création d'un nouvel enregistrement avec des valeurs de date :

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### Formats d'entrée pris en charge

Lors de la création d'enregistrements, les dates peuvent être fournies dans divers formats :

| Format | Exemple | Résultat |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | Identifiant unique pour la valeur du champ |
| `uid` | String! | Chaîne d'identifiant unique |
| `customField` | CustomField! | La définition du champ personnalisé (contient les valeurs de date) |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

**Important** : Les valeurs de date (`startDate`, `endDate`, `timezone`) sont accessibles par le champ `customField.value`, et non directement sur TodoCustomField.

### Structure de l'objet valeur

Les valeurs de date sont retournées par le champ `customField.value` sous forme d'objet JSON :

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**Remarque** : Le champ `value` est de type `CustomField`, et non de `TodoCustomField`.

## Interrogation des valeurs de date

Lors de l'interrogation d'enregistrements avec des champs personnalisés de date, accédez aux valeurs de date par le champ `customField.value` :

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

La réponse inclura les valeurs de date dans le champ `value` :

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## Intelligence d'affichage des dates

Le système formate automatiquement les dates en fonction de la plage :

| Scénario | Format d'affichage |
|----------|--------------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (aucune heure affichée) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Détection de toute la journée** : Les événements de 00:00 à 23:59 sont automatiquement détectés comme des événements de toute la journée.

## Gestion des fuseaux horaires

### Stockage
- Toutes les dates sont stockées en UTC dans la base de données
- Les informations de fuseau horaire sont préservées séparément
- La conversion se fait lors de l'affichage

### Meilleures pratiques
- Fournissez toujours un fuseau horaire pour plus de précision
- Utilisez des fuseaux horaires cohérents au sein d'un projet
- Tenez compte des emplacements des utilisateurs pour les équipes mondiales

### Fuseaux horaires courants

| Région | ID de fuseau horaire | Décalage UTC |
|--------|----------------------|--------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filtrage et interrogation

Les champs de date prennent en charge le filtrage complexe :

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### Vérification des champs de date vides

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### Opérateurs pris en charge

| Opérateur | Utilisation | Description |
|-----------|-------------|-------------|
| `EQ` | Avec dateRange | La date chevauche la plage spécifiée (toute intersection) |
| `NE` | Avec dateRange | La date ne chevauche pas la plage |
| `IS` | Avec `values: null` | Le champ de date est vide (startDate ou endDate est nul) |
| `NOT` | Avec `values: null` | Le champ de date a une valeur (les deux dates ne sont pas nulles) |

## Autorisations requises

| Action | Autorisation requise |
|--------|----------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Réponses d'erreur

### Format de date invalide
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
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
      "code": "NOT_FOUND"
    }
  }]
}
```


## Limitations

- Pas de prise en charge des dates récurrentes (utilisez des automatisations pour les événements récurrents)
- Impossible de définir une heure sans date
- Pas de calcul des jours ouvrés intégré
- Les plages de dates ne valident pas automatiquement fin > début
- La précision maximale est à la seconde (pas de stockage de millisecondes)

## Ressources connexes

- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux des champs personnalisés
- [API des automatisations](/api/automations/index) - Créez des automatisations basées sur des dates