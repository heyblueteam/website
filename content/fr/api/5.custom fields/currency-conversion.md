---
title: Champ personnalisé de conversion de devises
description: Créez des champs qui convertissent automatiquement les valeurs monétaires en utilisant des taux de change en temps réel
---

Les champs personnalisés de conversion de devises convertissent automatiquement les valeurs d'un champ de DEVISE source vers différentes devises cibles en utilisant des taux de change en temps réel. Ces champs se mettent à jour automatiquement chaque fois que la valeur de la devise source change.

Les taux de conversion sont fournis par l'[API Frankfurter](https://github.com/hakanensari/frankfurter), un service open-source qui suit les taux de change de référence publiés par la [Banque centrale européenne](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Cela garantit des conversions de devises précises, fiables et à jour pour vos besoins commerciaux internationaux.

## Exemple de base

Créez un champ de conversion de devises simple :

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## Exemple avancé

Créez un champ de conversion avec une date spécifique pour les taux historiques :

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## Processus de configuration complet

La configuration d'un champ de conversion de devises nécessite trois étapes :

### Étape 1 : Créer un champ de DEVISE source

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### Étape 2 : Créer le champ CURRENCY_CONVERSION

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### Étape 3 : Créer des options de conversion

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ de conversion |
| `type` | CustomFieldType! | ✅ Oui | Doit être `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | Non | ID du champ de DEVISE source à partir duquel convertir |
| `conversionDateType` | String | Non | Stratégie de date pour les taux de change (voir ci-dessous) |
| `conversionDate` | String | Non | Chaîne de date pour la conversion (basée sur conversionDateType) |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Les champs personnalisés sont automatiquement associés au projet en fonction du contexte de projet actuel de l'utilisateur. Aucun `projectId` n'est requis.

### Types de date de conversion

| Type | Description | Paramètre conversionDate |
|------|-------------|-------------------------|
| `currentDate` | Utilise des taux de change en temps réel | Non requis |
| `specificDate` | Utilise des taux d'une date fixe | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Utilise la date d'un autre champ | "todoDueDate" or DATE field ID |

## Création d'options de conversion

Les options de conversion définissent quelles paires de devises peuvent être converties :

### CreateCustomFieldOptionInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `customFieldId` | String! | ✅ Oui | ID du champ CURRENCY_CONVERSION |
| `title` | String! | ✅ Oui | Nom d'affichage pour cette option de conversion |
| `currencyConversionFrom` | String! | ✅ Oui | Code de la devise source ou "Tout" |
| `currencyConversionTo` | String! | ✅ Oui | Code de la devise cible |

### Utilisation de "Tout" comme source

La valeur spéciale "Tout" en tant que `currencyConversionFrom` crée une option de secours :

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

Cette option sera utilisée lorsqu'aucune correspondance spécifique de paire de devises n'est trouvée.

## Comment fonctionne la conversion automatique

1. **Mise à jour de la valeur** : Lorsqu'une valeur est définie dans le champ de DEVISE source
2. **Correspondance des options** : Le système trouve l'option de conversion correspondante basée sur la devise source
3. **Récupération du taux** : Récupère le taux de change de l'API Frankfurter
4. **Calcul** : Multiplie le montant source par le taux de change
5. **Stockage** : Enregistre la valeur convertie avec le code de la devise cible

### Flux d'exemple

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## Conversions basées sur la date

### Utilisation de la date actuelle

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

Les conversions se mettent à jour avec les taux de change actuels chaque fois que la valeur source change.

### Utilisation d'une date spécifique

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

Utilise toujours les taux de change de la date spécifiée.

### Utilisation de la date d'un champ

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

Utilise la date d'un autre champ (soit la date d'échéance d'une tâche, soit un champ personnalisé de DATE).

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ de conversion |
| `number` | Float | Le montant converti |
| `currency` | String | Le code de la devise cible |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été mise à jour pour la dernière fois |

## Source du taux de change

Blue utilise l'**API Frankfurter** pour les taux de change :
- API open-source hébergée par la Banque centrale européenne
- Mises à jour quotidiennes avec les taux de change officiels
- Prend en charge les taux historiques depuis 1999
- Gratuit et fiable pour un usage commercial

## Gestion des erreurs

### Échecs de conversion

Lorsque la conversion échoue (erreur API, devise invalide, etc.) :
- La valeur convertie est définie sur `0`
- La devise cible est toujours stockée
- Aucune erreur n'est signalée à l'utilisateur

### Scénarios courants

| Scénario | Résultat |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Aucune option correspondante | Uses "Any" option if available |
| Missing source value | Aucune conversion effectuée |

## Autorisations requises

La gestion des champs personnalisés nécessite un accès au niveau du projet :

| Rôle | Peut créer/met à jour des champs |
|------|-------------------------|
| `OWNER` | ✅ Oui |
| `ADMIN` | ✅ Oui |
| `MEMBER` | ❌ Non |
| `CLIENT` | ❌ Non |

Les permissions de visualisation pour les valeurs converties suivent les règles d'accès standard aux enregistrements.

## Meilleures pratiques

### Configuration des options
- Créez des paires de devises spécifiques pour les conversions courantes
- Ajoutez une option de secours "Tout" pour plus de flexibilité
- Utilisez des titres descriptifs pour les options

### Sélection de la stratégie de date
- Utilisez `currentDate` pour le suivi financier en direct
- Utilisez `specificDate` pour les rapports historiques
- Utilisez `fromDateField` pour des taux spécifiques aux transactions

### Considérations de performance
- Plusieurs champs de conversion se mettent à jour en parallèle
- Les appels API ne sont effectués que lorsque la valeur source change
- Les conversions de même devise évitent les appels API

## Cas d'utilisation courants

1. **Projets multi-devises**
   - Suivez les coûts du projet dans les devises locales
   - Rapportez le budget total dans la devise de l'entreprise
   - Comparez les valeurs entre les régions

2. **Ventes internationales**
   - Convertissez les valeurs des transactions dans la devise de rapport
   - Suivez les revenus dans plusieurs devises
   - Conversion historique pour les transactions closes

3. **Rapports financiers**
   - Conversions de devises à la fin de la période
   - États financiers consolidés
   - Budget vs. réel dans la devise locale

4. **Gestion des contrats**
   - Convertissez les valeurs des contrats à la date de signature
   - Suivez les calendriers de paiement dans plusieurs devises
   - Évaluation du risque de change

## Limitations

- Pas de support pour les conversions de cryptomonnaies
- Impossible de définir manuellement les valeurs converties (toujours calculées)
- Précision fixe de 2 décimales pour tous les montants convertis
- Pas de support pour les taux de change personnalisés
- Pas de mise en cache des taux de change (appel API frais pour chaque conversion)
- Dépend de la disponibilité de l'API Frankfurter

## Ressources connexes

- [Champs de devises](/api/custom-fields/currency) - Champs source pour les conversions
- [Champs de date](/api/custom-fields/date) - Pour les conversions basées sur la date
- [Champs de formule](/api/custom-fields/formula) - Calculs alternatifs
- [Aperçu des champs personnalisés](/custom-fields/list-custom-fields) - Concepts généraux