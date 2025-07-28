---
title: Créer un tableau de bord
description: Créer un nouveau tableau de bord pour la visualisation et le reporting des données dans Blue
---

## Créer un tableau de bord

La mutation `createDashboard` vous permet de créer un nouveau tableau de bord au sein de votre entreprise ou projet. Les tableaux de bord sont des outils de visualisation puissants qui aident les équipes à suivre les métriques, à surveiller les progrès et à prendre des décisions basées sur les données.

### Exemple de base

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### Tableau de bord spécifique au projet

Créez un tableau de bord associé à un projet spécifique :

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## Paramètres d'entrée

### CreateDashboardInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `companyId` | String! | ✅ Oui | L'ID de l'entreprise où le tableau de bord sera créé |
| `title` | String! | ✅ Oui | Le nom du tableau de bord. Doit être une chaîne non vide |
| `projectId` | String | Non | ID optionnel d'un projet à associer à ce tableau de bord |

## Champs de réponse

La mutation renvoie un objet complet `Dashboard` :

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour le tableau de bord créé |
| `title` | String! | Le titre du tableau de bord tel que fourni |
| `companyId` | String! | L'entreprise à laquelle appartient ce tableau de bord |
| `projectId` | String | L'ID du projet associé (si fourni) |
| `project` | Project | L'objet projet associé (si projectId a été fourni) |
| `createdBy` | User! | L'utilisateur qui a créé le tableau de bord (vous) |
| `dashboardUsers` | [DashboardUser!]! | Liste des utilisateurs ayant accès (initialement juste le créateur) |
| `createdAt` | DateTime! | Horodatage de la création du tableau de bord |
| `updatedAt` | DateTime! | Horodatage de la dernière modification (identique à createdAt pour les nouveaux tableaux de bord) |

### Champs DashboardUser

Lorsqu'un tableau de bord est créé, le créateur est automatiquement ajouté en tant qu'utilisateur du tableau de bord :

| Champ | Type | Description |
|-------|------|-------------|
| `id` | String! | Identifiant unique pour la relation utilisateur du tableau de bord |
| `user` | User! | L'objet utilisateur ayant accès au tableau de bord |
| `role` | DashboardRole! | Le rôle de l'utilisateur (le créateur obtient un accès complet) |
| `dashboard` | Dashboard! | Référence au tableau de bord |

## Permissions requises

Tout utilisateur authentifié appartenant à l'entreprise spécifiée peut créer des tableaux de bord. Il n'y a pas d'exigences de rôle spéciales.

| Statut de l'utilisateur | Peut créer un tableau de bord |
|------------------------|-------------------------------|
| Company Member | ✅ Oui |
| Membre non-entreprise  | ❌ Non |
| Unauthenticated | ❌ Non |

## Réponses d'erreur

### Entreprise invalide
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Utilisateur non dans l'entreprise
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Projet invalide
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Titre vide
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Notes importantes

- **Propriété automatique** : L'utilisateur créant le tableau de bord devient automatiquement son propriétaire avec des permissions complètes
- **Association de projet** : Si vous fournissez un `projectId`, il doit appartenir à la même entreprise
- **Permissions initiales** : Seul le créateur a accès initialement. Utilisez `editDashboard` pour ajouter d'autres utilisateurs
- **Exigences de titre** : Les titres des tableaux de bord doivent être des chaînes non vides. Il n'y a pas d'exigence d'unicité
- **Appartenance à l'entreprise** : Vous devez être membre de l'entreprise pour créer des tableaux de bord dans celle-ci

## Flux de création de tableau de bord

1. **Créer le tableau de bord** en utilisant cette mutation
2. **Configurer les graphiques et les widgets** à l'aide de l'interface de création de tableau de bord
3. **Ajouter des membres de l'équipe** en utilisant la mutation `editDashboard` avec `dashboardUsers`
4. **Configurer les filtres et les plages de dates** via l'interface du tableau de bord
5. **Partager ou intégrer** le tableau de bord en utilisant son ID unique

## Cas d'utilisation

1. **Tableaux de bord exécutifs** : Créer des aperçus de haut niveau des métriques de l'entreprise
2. **Suivi de projet** : Construire des tableaux de bord spécifiques aux projets pour surveiller les progrès
3. **Performance de l'équipe** : Suivre la productivité de l'équipe et les métriques de réussite
4. **Reporting client** : Créer des tableaux de bord pour des rapports destinés aux clients
5. **Surveillance en temps réel** : Configurer des tableaux de bord pour des données opérationnelles en direct

## Meilleures pratiques

1. **Conventions de nommage** : Utilisez des titres clairs et descriptifs qui indiquent l'objectif du tableau de bord
2. **Association de projet** : Liez les tableaux de bord aux projets lorsqu'ils sont spécifiques à un projet
3. **Gestion des accès** : Ajoutez des membres de l'équipe immédiatement après la création pour la collaboration
4. **Organisation** : Créez une hiérarchie de tableaux de bord en utilisant des modèles de nommage cohérents

## Opérations connexes

- [Lister les tableaux de bord](/api/dashboards/) - Récupérer tous les tableaux de bord pour une entreprise ou un projet
- [Modifier le tableau de bord](/api/dashboards/rename-dashboard) - Renommer le tableau de bord ou gérer les utilisateurs
- [Copier le tableau de bord](/api/dashboards/copy-dashboard) - Dupliquer un tableau de bord existant
- [Supprimer le tableau de bord](/api/dashboards/delete-dashboard) - Retirer un tableau de bord