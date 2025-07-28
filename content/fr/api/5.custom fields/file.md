---
title: Champ personnalisé de fichier
description: Créez des champs de fichier pour attacher des documents, des images et d'autres fichiers aux enregistrements
---

Les champs personnalisés de fichier vous permettent d'attacher plusieurs fichiers aux enregistrements. Les fichiers sont stockés en toute sécurité dans AWS S3 avec un suivi complet des métadonnées, une validation des types de fichiers et des contrôles d'accès appropriés.

## Exemple de base

Créez un champ de fichier simple :

```graphql
mutation CreateFileField {
  createCustomField(input: {
    name: "Attachments"
    type: FILE
  }) {
    id
    name
    type
  }
}
```

## Exemple avancé

Créez un champ de fichier avec une description :

```graphql
mutation CreateDetailedFileField {
  createCustomField(input: {
    name: "Project Documents"
    type: FILE
    description: "Upload project-related documents, images, and files"
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
| `name` | String! | ✅ Oui | Nom d'affichage du champ de fichier |
| `type` | CustomFieldType! | ✅ Oui | Doit être `FILE` |
| `description` | String | Non | Texte d'aide affiché aux utilisateurs |

**Remarque** : Les champs personnalisés sont automatiquement associés au projet en fonction du contexte de projet actuel de l'utilisateur. Aucun `projectId` n'est requis.

## Processus de téléchargement de fichiers

### Étape 1 : Télécharger le fichier

Tout d'abord, téléchargez le fichier pour obtenir un UID de fichier :

```graphql
mutation UploadFile {
  uploadFile(input: {
    file: $file  # File upload variable
    companyId: "company_123"
    projectId: "proj_123"
  }) {
    id
    uid
    name
    size
    type
    extension
    status
  }
}
```

### Étape 2 : Attacher le fichier à l'enregistrement

Ensuite, attachez le fichier téléchargé à un enregistrement :

```graphql
mutation AttachFileToRecord {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "file_field_456"
    fileUid: "file_uid_from_upload"
  }) {
    id
    file {
      uid
      name
      size
      type
    }
  }
}
```

## Gestion des pièces jointes de fichiers

### Ajout de fichiers uniques

```graphql
mutation AddFileToField {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  }) {
    id
    position
    file {
      uid
      name
      size
      type
      extension
    }
  }
}
```

### Suppression de fichiers

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Opérations de fichiers en masse

Mettez à jour plusieurs fichiers à la fois en utilisant customFieldOptionIds :

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Paramètres d'entrée pour le téléchargement de fichiers

### UploadFileInput

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `file` | Upload! | ✅ Oui | Fichier à télécharger |
| `companyId` | String! | ✅ Oui | ID de l'entreprise pour le stockage de fichiers |
| `projectId` | String | Non | ID du projet pour les fichiers spécifiques au projet |

### Paramètres d'entrée pour la gestion des fichiers

| Paramètre | Type | Requis | Description |
|-----------|------|--------|-------------|
| `todoId` | String! | ✅ Oui | ID de l'enregistrement |
| `customFieldId` | String! | ✅ Oui | ID du champ personnalisé de fichier |
| `fileUid` | String! | ✅ Oui | Identifiant unique du fichier téléchargé |

## Stockage de fichiers et limites

### Limites de taille de fichier

| Type de limite | Taille |
|----------------|-------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Types de fichiers pris en charge

#### Images
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Vidéos
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Audio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Documents
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Archives
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Code/Texte
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Architecture de stockage

- **Stockage** : AWS S3 avec une structure de dossier organisée
- **Format de chemin** : `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Sécurité** : URL signées pour un accès sécurisé
- **Sauvegarde** : Redondance S3 automatique

## Champs de réponse

### Réponse de fichier

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | ID de la base de données |
| `uid` | String! | Identifiant unique du fichier |
| `name` | String! | Nom de fichier original |
| `size` | Float! | Taille du fichier en octets |
| `type` | String! | Type MIME |
| `extension` | String! | Extension de fichier |
| `status` | FileStatus | EN ATTENTE ou CONFIRMÉ (nullable) |
| `shared` | Boolean! | Indique si le fichier est partagé |
| `createdAt` | DateTime! | Horodatage de téléchargement |

### Réponse TodoCustomFieldFile

| Champ | Type | Description |
|-------|------|-------------|
| `id` | ID! | ID de l'enregistrement de jonction |
| `uid` | String! | Identifiant unique |
| `position` | Float! | Ordre d'affichage |
| `file` | File! | Objet de fichier associé |
| `todoCustomField` | TodoCustomField! | Champ personnalisé parent |
| `createdAt` | DateTime! | Date à laquelle le fichier a été attaché |

## Création d'enregistrements avec des fichiers

Lors de la création d'enregistrements, vous pouvez attacher des fichiers en utilisant leurs UIDs :

```graphql
mutation CreateRecordWithFiles {
  createTodo(input: {
    title: "Project deliverables"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "file_field_id"
      customFieldOptionIds: ["file_uid_1", "file_uid_2"]
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
      todoCustomFieldFiles {
        id
        position
        file {
          uid
          name
          size
          type
        }
      }
    }
  }
}
```

## Validation et sécurité des fichiers

### Validation de téléchargement

- **Vérification du type MIME** : Valide par rapport aux types autorisés
- **Validation de l'extension de fichier** : Solution de secours pour `application/octet-stream`
- **Limites de taille** : Appliquées au moment du téléchargement
- **Assainissement des noms de fichiers** : Supprime les caractères spéciaux

### Contrôle d'accès

- **Permissions de téléchargement** : Adhésion au projet/à l'entreprise requise
- **Association de fichiers** : Rôles ADMIN, PROPRIÉTAIRE, MEMBRE, CLIENT
- **Accès aux fichiers** : Hérité des permissions du projet/de l'entreprise
- **URLs sécurisées** : URLs signées à durée limitée pour l'accès aux fichiers

## Permissions requises

| Action | Permission requise |
|--------|-------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Réponses d'erreur

### Fichier trop volumineux
```json
{
  "errors": [{
    "message": "File \"filename.pdf\": Size exceeds maximum limit of 256MB",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Fichier non trouvé
```json
{
  "errors": [{
    "message": "File not found",
    "extensions": {
      "code": "FILE_NOT_FOUND"
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

## Meilleures pratiques

### Gestion des fichiers
- Téléchargez les fichiers avant de les attacher aux enregistrements
- Utilisez des noms de fichiers descriptifs
- Organisez les fichiers par projet/objectif
- Nettoyez régulièrement les fichiers inutilisés

### Performance
- Téléchargez des fichiers par lots lorsque cela est possible
- Utilisez des formats de fichiers appropriés pour le type de contenu
- Compressez les fichiers volumineux avant le téléchargement
- Prenez en compte les exigences de prévisualisation des fichiers

### Sécurité
- Validez le contenu des fichiers, pas seulement les extensions
- Utilisez un scan de virus pour les fichiers téléchargés
- Mettez en œuvre des contrôles d'accès appropriés
- Surveillez les modèles de téléchargement de fichiers

## Cas d'utilisation courants

1. **Gestion de documents**
   - Spécifications de projet
   - Contrats et accords
   - Notes de réunion et présentations
   - Documentation technique

2. **Gestion des actifs**
   - Fichiers de conception et maquettes
   - Actifs de marque et logos
   - Matériaux marketing
   - Images de produits

3. **Conformité et dossiers**
   - Documents juridiques
   - Pistes d'audit
   - Certificats et licences
   - Dossiers financiers

4. **Collaboration**
   - Ressources partagées
   - Documents sous contrôle de version
   - Retours et annotations
   - Matériaux de référence

## Fonctionnalités d'intégration

### Avec des automatisations
- Déclencher des actions lorsque des fichiers sont ajoutés/supprimés
- Traiter des fichiers en fonction du type ou des métadonnées
- Envoyer des notifications pour les changements de fichiers
- Archiver des fichiers en fonction des conditions

### Avec des images de couverture
- Utilisez des champs de fichiers comme sources d'images de couverture
- Traitement d'image automatique et vignettes
- Mises à jour dynamiques de la couverture lorsque les fichiers changent

### Avec des recherches
- Référencez des fichiers à partir d'autres enregistrements
- Agrégez les comptes et tailles de fichiers
- Trouvez des enregistrements par métadonnées de fichiers
- Faites des références croisées aux pièces jointes de fichiers

## Limitations

- Taille maximale de 256 Mo par fichier
- Dépend de la disponibilité de S3
- Pas de versioning de fichier intégré
- Pas de conversion automatique de fichiers
- Capacités de prévisualisation de fichiers limitées
- Pas d'édition collaborative en temps réel

## Ressources connexes

- [API de téléchargement de fichiers](/api/upload-files) - Points de terminaison de téléchargement de fichiers
- [Aperçu des champs personnalisés](/api/custom-fields/list-custom-fields) - Concepts généraux
- [API d'automatisations](/api/automations) - Automatisations basées sur des fichiers
- [Documentation AWS S3](https://docs.aws.amazon.com/s3/) - Backend de stockage