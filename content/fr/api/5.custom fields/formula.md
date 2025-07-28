---
title: Champ personnalisé de formule
description: Créez des champs calculés qui calculent automatiquement des valeurs en fonction d'autres données
---

Les champs personnalisés de formule sont utilisés pour les calculs de graphiques et de tableaux de bord au sein de Blue. Ils définissent des fonctions d'agrégation (SOMME, MOYENNE, COMPTE, etc.) qui opèrent sur les données des champs personnalisés pour afficher des métriques calculées dans les graphiques. Les formules ne sont pas calculées au niveau individuel des tâches, mais agrègent plutôt des données à travers plusieurs enregistrements à des fins de visualisation.

## Exemple de base

Créez un champ de formule pour les calculs de graphiques :

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Exemple avancé

Créez une formule de devise avec des calculs complexes :

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Paramètres d'entrée

### CreateCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `name` | String! | ✅ Oui | Nom d'affichage du champ de formule |
| `type` | CustomFieldType! | ✅ Oui | Doit être `FORMULA` |
| `projectId` | String! | ✅ Oui | L'ID du projet où ce champ sera créé |
| `formula` | JSON | Non | Définition de la formule pour les calculs de graphiques |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

### Structure de la formule

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Fonctions prises en charge

### Fonctions d'agrégation de graphique

Les champs de formule prennent en charge les fonctions d'agrégation suivantes pour les calculs de graphiques :

| Fonction | Description | Enum ChartFunction |
|----------|-------------|-------------------|
| `SUM` | Somme de toutes les valeurs | `SUM` |
| `AVERAGE` | Moyenne des valeurs numériques | `AVERAGE` |
| `AVERAGEA` | Moyenne excluant les zéros et les nuls | `AVERAGEA` |
| `COUNT` | Compte des valeurs | `COUNT` |
| `COUNTA` | Compte excluant les zéros et les nuls | `COUNTA` |
| `MAX` | Valeur maximale | `MAX` |
| `MIN` | Valeur minimale | `MIN` |

**Remarque** : Ces fonctions sont utilisées dans le champ `display.function` et opèrent sur des données agrégées pour les visualisations de graphiques. Les expressions mathématiques complexes ou les calculs au niveau des champs ne sont pas pris en charge.

## Types d'affichage

### Affichage numérique

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Résultat : `1250.75`

### Affichage de devise

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Résultat : `$1,250.75`

### Affichage de pourcentage

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Résultat : `87.5%`

## Édition des champs de formule

Mettez à jour les champs de formule existants :

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Traitement des formules

### Contexte de calcul de graphique

Les champs de formule sont traités dans le contexte des segments de graphique et des tableaux de bord :
- Les calculs se produisent lorsque les graphiques sont rendus ou mis à jour
- Les résultats sont stockés dans `ChartSegment.formulaResult` sous forme de valeurs décimales
- Le traitement est géré par une file d'attente BullMQ dédiée nommée 'formula'
- Les mises à jour sont publiées aux abonnés du tableau de bord pour des mises à jour en temps réel

### Formatage d'affichage

La fonction `getFormulaDisplayValue` formate les résultats calculés en fonction du type d'affichage :
- **NUMÉRIQUE** : Affiche comme un nombre simple avec une précision optionnelle
- **POURCENTAGE** : Ajoute un suffixe % avec une précision optionnelle  
- **DEVISE** : Formate en utilisant le code de devise spécifié

## Stockage des résultats de formule

Les résultats sont stockés dans le champ `formulaResult` :

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Champs de réponse

### Réponse TodoCustomField

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ de formule |
| `number` | Float | Résultat numérique calculé |
| `formulaResult` | JSON | Résultat complet avec formatage d'affichage |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Quand la valeur a été créée |
| `updatedAt` | DateTime! | Quand la valeur a été calculée pour la dernière fois |

## Contexte des données

### Source de données de graphique

Les champs de formule fonctionnent dans le contexte de la source de données de graphique :
- Les formules agrègent les valeurs des champs personnalisés à travers les tâches dans un projet
- La fonction d'agrégation spécifiée dans `display.function` détermine le calcul
- Les résultats sont calculés à l'aide de fonctions d'agrégation SQL (avg, sum, count, etc.)
- Les calculs sont effectués au niveau de la base de données pour plus d'efficacité

## Exemples de formules courantes

### Budget total (Affichage de graphique)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Score moyen (Affichage de graphique)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Compte des tâches (Affichage de graphique)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Permissions requises

Les opérations sur les champs personnalisés suivent les permissions basées sur les rôles standard :

| Action | Rôle requis |
|--------|-------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Remarque** : Les rôles spécifiques requis dépendent de la configuration des rôles personnalisés de votre projet. Il n'existe pas de constantes de permission spéciales comme CUSTOM_FIELDS_CREATE.

## Gestion des erreurs

### Erreur de validation
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Champ personnalisé non trouvé
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

### Conception de formule
- Utilisez des noms clairs et descriptifs pour les champs de formule
- Ajoutez des descriptions expliquant la logique de calcul
- Testez les formules avec des données d'exemple avant le déploiement
- Gardez les formules simples et lisibles

### Optimisation des performances
- Évitez les dépendances de formule profondément imbriquées
- Utilisez des références de champ spécifiques plutôt que des jokers
- Envisagez des stratégies de mise en cache pour des calculs complexes
- Surveillez les performances des formules dans de grands projets

### Qualité des données
- Validez les données sources avant de les utiliser dans les formules
- Gérez les valeurs vides ou nulles de manière appropriée
- Utilisez une précision appropriée pour les types d'affichage
- Envisagez des cas extrêmes dans les calculs

## Cas d'utilisation courants

1. **Suivi financier**
   - Calculs de budget
   - États de profits/pertes
   - Analyse des coûts
   - Projections de revenus

2. **Gestion de projet**
   - Pourcentages d'achèvement
   - Utilisation des ressources
   - Calculs de calendrier
   - Indicateurs de performance

3. **Contrôle de qualité**
   - Scores moyens
   - Taux de réussite/échec
   - Indicateurs de qualité
   - Suivi de conformité

4. **Intelligence d'affaires**
   - Calculs d'indicateurs clés de performance (KPI)
   - Analyse des tendances
   - Indicateurs comparatifs
   - Valeurs de tableau de bord

## Limitations

- Les formules ne sont destinées qu'aux agrégations de graphiques/tableaux de bord, pas aux calculs au niveau des tâches
- Limitées aux sept fonctions d'agrégation prises en charge (SOMME, MOYENNE, etc.)
- Pas d'expressions mathématiques complexes ou de calculs champ à champ
- Ne peut pas faire référence à plusieurs champs dans une seule formule
- Les résultats ne sont visibles que dans les graphiques et les tableaux de bord
- Le champ `logic` est uniquement destiné au texte d'affichage, pas à la logique de calcul réelle

## Ressources connexes

- [Champs numériques](/api/5.custom%20fields/number) - Pour des valeurs numériques statiques
- [Champs de devise](/api/5.custom%20fields/currency) - Pour des valeurs monétaires
- [Champs de référence](/api/5.custom%20fields/reference) - Pour des données inter-projets
- [Champs de recherche](/api/5.custom%20fields/lookup) - Pour des données agrégées
- [Aperçu des champs personnalisés](/api/5.custom%20fields/2.list-custom-fields) - Concepts généraux