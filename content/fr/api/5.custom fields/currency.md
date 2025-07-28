---
title: Champ personnalisé de devise
description: Créez des champs de devise pour suivre les valeurs monétaires avec un formatage et une validation appropriés
---

Les champs personnalisés de devise vous permettent de stocker et de gérer des valeurs monétaires avec des codes de devise associés. Le champ prend en charge 72 devises différentes, y compris les principales devises fiat et les cryptomonnaies, avec un formatage automatique et des contraintes min/max optionnelles.

## Exemple de base

Créez un champ de devise simple :

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Exemple avancé

Créez un champ de devise avec des contraintes de validation :

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ de devise |
| `type` | CustomFieldType! | ✅ Oui | Doit être `CURRENCY` |
| `currency` | String | Non | Code de devise par défaut (code ISO à 3 lettres) |
| `min` | Float | Non | Valeur minimale autorisée (stockée mais non appliquée lors des mises à jour) |
| `max` | Float | Non | Valeur maximale autorisée (stockée mais non appliquée lors des mises à jour) |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Le contexte du projet est automatiquement déterminé à partir de votre authentification. Vous devez avoir accès au projet où vous créez le champ.

## Définir des valeurs de devise

Pour définir ou mettre à jour une valeur de devise sur un enregistrement :

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### Paramètres de SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de devise |
| `number` | Float! | ✅ Oui | Le montant monétaire |
| `currency` | String! | ✅ Oui | Code de devise à 3 lettres |

## Création d'enregistrements avec des valeurs de devise

Lors de la création d'un nouvel enregistrement avec des valeurs de devise :

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### Format d'entrée pour la création

Lors de la création d'enregistrements, les valeurs de devise sont transmises différemment :

| Paramètre | Type | Description |
|-----------|------|-------------|
| `customFieldId` | String! | ID du champ de devise |
| `value` | String! | Montant sous forme de chaîne (par exemple, "1500.50") |
| `currency` | String! | Code de devise à 3 lettres |

## Devises prises en charge

Blue prend en charge 72 devises, y compris 70 devises fiat et 2 cryptomonnaies :

### Devises fiat

#### Amériques
| Devise | Code | Nom |
|--------|------|-----|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Bolívar Soberano vénézuélien |
| Bolivien Boliviano | `BOB` | Bolivien Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europe
| Devise | Code | Nom |
|--------|------|-----|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Couronne norvégienne | `NOK` | Couronne norvégienne |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Asie-Pacifique
| Devise | Code | Nom |
|--------|------|-----|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### Moyen-Orient et Afrique
| Devise | Code | Nom |
|--------|------|-----|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Cryptomonnaies
| Devise | Code |
|--------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Champs de réponse

### TodoCustomField Response

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `number` | Float | Le montant monétaire |
| `currency` | String | Le code de devise à 3 lettres |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été modifiée pour la dernière fois |

## Formatage des devises

Le système formate automatiquement les valeurs de devise en fonction de la locale :

- **Placement des symboles** : Positionne correctement les symboles de devise (avant/après)
- **Séparateurs décimaux** : Utilise des séparateurs spécifiques à la locale (. ou ,)
- **Séparateurs de milliers** : Applique un groupement approprié
- **Chiffres décimaux** : Affiche 0 à 2 chiffres décimaux en fonction du montant
- **Gestion spéciale** : USD/CAD affichent un préfixe de code de devise pour plus de clarté

### Exemples de formatage

| Valeur | Devise | Affichage |
|--------|--------|-----------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Validation

### Validation des montants
- Doit être un nombre valide
- Les contraintes min/max sont stockées avec la définition du champ mais ne sont pas appliquées lors des mises à jour de valeur
- Prend en charge jusqu'à 2 chiffres décimaux pour l'affichage (précision complète stockée en interne)

### Validation du code de devise
- Doit être l'un des 72 codes de devise pris en charge
- Sensible à la casse (utiliser des majuscules)
- Les codes invalides renvoient une erreur

## Fonctionnalités d'intégration

### Formules
Les champs de devise peuvent être utilisés dans des champs personnalisés de FORMULE pour des calculs :
- Somme de plusieurs champs de devise
- Calculs de pourcentages
- Effectuer des opérations arithmétiques

### Conversion de devise
Utilisez des champs de CONVERSION_DEVISE pour convertir automatiquement entre les devises (voir [Champs de conversion de devise](/api/custom-fields/currency-conversion))

### Automatisations
Les valeurs de devise peuvent déclencher des automatisations basées sur :
- Seuils de montant
- Type de devise
- Changements de valeur

## Autorisations requises

| Action | Autorisation requise |
|--------|---------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Remarque** : Bien que tout membre du projet puisse créer des champs personnalisés, la capacité à définir des valeurs dépend des autorisations basées sur les rôles configurées pour chaque champ.

## Réponses d'erreur

### Valeur de devise invalide
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

Cette erreur se produit lorsque :
- Le code de devise n'est pas l'un des 72 codes pris en charge
- Le format du nombre est invalide
- La valeur ne peut pas être analysée correctement

### Champ personnalisé introuvable
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

## Meilleures pratiques

### Sélection de la devise
- Définissez une devise par défaut qui correspond à votre marché principal
- Utilisez les codes de devise ISO 4217 de manière cohérente
- Tenez compte de la localisation de l'utilisateur lors du choix des valeurs par défaut

### Contraintes de valeur
- Définissez des valeurs min/max raisonnables pour éviter les erreurs de saisie de données
- Utilisez 0 comme minimum pour les champs qui ne devraient pas être négatifs
- Tenez compte de votre cas d'utilisation lors de la définition des maximums

### Projets multi-devises
- Utilisez une devise de base cohérente pour les rapports
- Implémentez des champs de CONVERSION_DEVISE pour la conversion automatique
- Documentez quelle devise doit être utilisée pour chaque champ

## Cas d'utilisation courants

1. **Budgétisation de projet**
   - Suivi du budget du projet
   - Estimations des coûts
   - Suivi des dépenses

2. **Ventes et affaires**
   - Valeurs des affaires
   - Montants des contrats
   - Suivi des revenus

3. **Planification financière**
   - Montants d'investissement
   - Tours de financement
   - Objectifs financiers

4. **Affaires internationales**
   - Tarification multi-devises
   - Suivi des changes
   - Transactions transfrontalières

## Limitations

- Maximum de 2 chiffres décimaux pour l'affichage (bien que plus de précision soit stockée)
- Pas de conversion de devise intégrée dans les champs de DEVISE standard
- Impossible de mélanger des devises dans une seule valeur de champ
- Pas de mises à jour automatiques des taux de change (utilisez CONVERSION_DEVISE pour cela)
- Les symboles de devise ne sont pas personnalisables

## Ressources connexes

- [Champs de conversion de devise](/api/custom-fields/currency-conversion) - Conversion automatique de devise
- [Champs numériques](/api/custom-fields/number) - Pour des valeurs numériques non monétaires
- [Champs de formule](/api/custom-fields/formula) - Calculer avec des valeurs de devise
- [Champs personnalisés de liste](/api/custom-fields/list-custom-fields) - Interroger tous les champs personnalisés dans un projet