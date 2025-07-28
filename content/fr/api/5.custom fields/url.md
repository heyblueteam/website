---
title: Champ personnalisé URL
description: Créez des champs URL pour stocker des adresses de sites Web et des liens
---

Les champs personnalisés URL vous permettent de stocker des adresses de sites Web et des liens dans vos enregistrements. Ils sont idéaux pour suivre les sites Web de projets, les liens de référence, les URL de documentation ou toute ressource en ligne liée à votre travail.

## Exemple de base

Créez un champ URL simple :

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ URL avec une description :

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
    }
  ) {
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
| `name` | String! | ✅ Oui | Nom d'affichage du champ URL |
| `type` | CustomFieldType! | ✅ Oui | Doit être `URL` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque :** Le `projectId` est passé comme un argument séparé à la mutation, et non comme partie de l'objet d'entrée.

## Définir des valeurs URL

Pour définir ou mettre à jour une valeur URL sur un enregistrement :

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### Paramètres SetTodoCustomFieldInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement à mettre à jour |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé URL |
| `text` | String! | ✅ Oui | Adresse URL à stocker |

## Création d'enregistrements avec des valeurs URL

Lors de la création d'un nouvel enregistrement avec des valeurs URL :

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
    }
  }
}
```

## Champs de réponse

### TodoCustomField Réponse

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la valeur du champ |
| `customField` | CustomField! | La définition du champ personnalisé |
| `text` | String | L'adresse URL stockée |
| `todo` | Todo! | L'enregistrement auquel cette valeur appartient |
| `createdAt` | DateTime! | Date de création de la valeur |
| `updatedAt` | DateTime! | Date de dernière modification de la valeur |

## Validation des URL

### Mise en œuvre actuelle
- **API directe** : Aucune validation de format d'URL n'est actuellement appliquée
- **Formulaires** : La validation des URL est prévue mais n'est pas actuellement active
- **Stockage** : Toute valeur de chaîne peut être stockée dans les champs URL

### Validation prévue
Les versions futures incluront :
- Validation du protocole HTTP/HTTPS
- Vérification du format d'URL valide
- Validation du nom de domaine
- Ajout automatique de préfixe de protocole

### Formats d'URL recommandés
Bien que non actuellement appliqués, utilisez ces formats standards :

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Remarques importantes

### Format de stockage
- Les URL sont stockées en texte brut sans modification
- Aucun ajout automatique de protocole (http://, https://)
- La sensibilité à la casse est préservée telle qu'entré
- Aucune encodage/décodage d'URL effectué

### API directe vs Formulaires
- **Formulaires** : Validation des URL prévue (non actuellement active)
- **API directe** : Aucune validation - tout texte peut être stocké
- **Recommandation** : Validez les URL dans votre application avant de les stocker

### Champs URL vs Champs texte
- **URL** : Destiné sémantiquement aux adresses Web
- **TEXT_SINGLE** : Texte général sur une seule ligne
- **Backend** : Stockage et validation actuellement identiques
- **Frontend** : Composants UI différents pour la saisie de données

## Permissions requises

Les opérations sur les champs personnalisés utilisent des permissions basées sur les rôles :

| Action | Rôle requis |
|--------|-------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Remarque :** Les permissions sont vérifiées en fonction des rôles des utilisateurs dans le projet, et non des constantes de permission spécifiques.

## Réponses d'erreur

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

### Validation de champ requis (Formulaires uniquement)
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

## Meilleures pratiques

### Normes de format d'URL
- Incluez toujours le protocole (http:// ou https://)
- Utilisez HTTPS lorsque cela est possible pour la sécurité
- Testez les URL avant de les stocker pour vous assurer qu'elles sont accessibles
- Envisagez d'utiliser des URL raccourcies à des fins d'affichage

### Qualité des données
- Validez les URL dans votre application avant de les stocker
- Vérifiez les fautes de frappe courantes (protocoles manquants, domaines incorrects)
- Standardisez les formats d'URL dans votre organisation
- Considérez l'accessibilité et la disponibilité des URL

### Considérations de sécurité
- Soyez prudent avec les URL fournies par les utilisateurs
- Validez les domaines si vous limitez à des sites spécifiques
- Envisagez de scanner les URL pour un contenu malveillant
- Utilisez des URL HTTPS lors du traitement de données sensibles

## Filtrage et recherche

### Recherche par contient
Les champs URL prennent en charge la recherche par sous-chaîne :

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Capacités de recherche
- Correspondance de sous-chaîne insensible à la casse
- Correspondance de domaine partiel
- Recherche de chemin et de paramètre
- Aucun filtrage spécifique au protocole

## Cas d'utilisation courants

1. **Gestion de projet**
   - Sites Web de projet
   - Liens de documentation
   - URL de dépôt
   - Sites de démonstration

2. **Gestion de contenu**
   - Matériaux de référence
   - Liens sources
   - Ressources multimédias
   - Articles externes

3. **Support client**
   - Sites Web des clients
   - Documentation de support
   - Articles de base de connaissances
   - Tutoriels vidéo

4. **Ventes et marketing**
   - Sites Web d'entreprise
   - Pages produits
   - Matériaux marketing
   - Profils de médias sociaux

## Fonctionnalités d'intégration

### Avec des recherches
- Référencer des URL d'autres enregistrements
- Trouver des enregistrements par domaine ou motif d'URL
- Afficher des ressources Web connexes
- Agréger des liens de plusieurs sources

### Avec des formulaires
- Composants d'entrée spécifiques aux URL
- Validation prévue pour un format d'URL approprié
- Capacités d'aperçu de lien (frontend)
- Affichage d'URL cliquable

### Avec des rapports
- Suivre l'utilisation des URL et les motifs
- Surveiller les liens brisés ou inaccessibles
- Catégoriser par domaine ou protocole
- Exporter des listes d'URL pour analyse

## Limitations

### Limitations actuelles
- Aucune validation active du format d'URL
- Aucun ajout automatique de protocole
- Aucune vérification de lien ou de vérification d'accessibilité
- Aucun raccourcissement ou expansion d'URL
- Aucune génération de favicon ou d'aperçu

### Restrictions d'automatisation
- Non disponible en tant que champs de déclenchement d'automatisation
- Ne peut pas être utilisé dans les mises à jour de champs d'automatisation
- Peut être référencé dans les conditions d'automatisation
- Disponible dans les modèles d'e-mail et les webhooks

### Contraintes générales
- Aucune fonctionnalité d'aperçu de lien intégrée
- Aucun raccourcissement automatique d'URL
- Aucune traçabilité des clics ou analyse
- Aucune vérification d'expiration d'URL
- Aucun scan d'URL malveillantes

## Améliorations futures

### Fonctionnalités prévues
- Validation du protocole HTTP/HTTPS
- Modèles de validation regex personnalisés
- Ajout automatique de préfixe de protocole
- Vérification de l'accessibilité des URL

### Améliorations potentielles
- Génération d'aperçu de lien
- Affichage de favicon
- Intégration de raccourcissement d'URL
- Capacités de suivi des clics
- Détection de liens brisés

## Ressources connexes

- [Champs texte](/api/custom-fields/text-single) - Pour les données texte non-URL
- [Champs e-mail](/api/custom-fields/email) - Pour les adresses e-mail
- [Aperçu des champs personnalisés](/api/custom-fields/2.list-custom-fields) - Concepts généraux

## Migration depuis les champs texte

Si vous migrez des champs texte vers des champs URL :

1. **Créez un champ URL** avec le même nom et la même configuration
2. **Exportez les valeurs texte existantes** pour vérifier qu'elles sont des URL valides
3. **Mettez à jour les enregistrements** pour utiliser le nouveau champ URL
4. **Supprimez l'ancien champ texte** après une migration réussie
5. **Mettez à jour les applications** pour utiliser des composants UI spécifiques aux URL

### Exemple de migration
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```