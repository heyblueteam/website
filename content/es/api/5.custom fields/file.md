---
title: Campo Personalizado de Archivo
description: Crea campos de archivo para adjuntar documentos, imágenes y otros archivos a los registros
---

Los campos personalizados de archivo te permiten adjuntar múltiples archivos a los registros. Los archivos se almacenan de forma segura en AWS S3 con un seguimiento completo de metadatos, validación de tipos de archivo y controles de acceso adecuados.

## Ejemplo Básico

Crea un campo de archivo simple:

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

## Ejemplo Avanzado

Crea un campo de archivo con descripción:

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

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de archivo |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `FILE` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: Los campos personalizados se asocian automáticamente con el proyecto basado en el contexto del proyecto actual del usuario. No se requiere ningún parámetro `projectId`.

## Proceso de Carga de Archivos

### Paso 1: Subir Archivo

Primero, sube el archivo para obtener un UID de archivo:

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

### Paso 2: Adjuntar Archivo al Registro

Luego, adjunta el archivo subido a un registro:

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

## Gestión de Archivos Adjuntos

### Agregar Archivos Únicos

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

### Eliminar Archivos

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Operaciones de Archivos en Masa

Actualiza múltiples archivos a la vez utilizando customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Parámetros de Entrada para Carga de Archivos

### UploadFileInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `file` | Upload! | ✅ Sí | Archivo a subir |
| `companyId` | String! | ✅ Sí | ID de la empresa para el almacenamiento de archivos |
| `projectId` | String | No | ID del proyecto para archivos específicos del proyecto |

### Parámetros de Entrada para Gestión de Archivos

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de archivo |
| `fileUid` | String! | ✅ Sí | Identificador único del archivo subido |

## Almacenamiento de Archivos y Límites

### Límites de Tamaño de Archivo

| Tipo de Límite | Tamaño |
|----------------|--------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Tipos de Archivo Soportados

#### Imágenes
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Videos
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Audio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Documentos
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Archivos
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Código/Textos
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Arquitectura de Almacenamiento

- **Almacenamiento**: AWS S3 con estructura de carpetas organizada
- **Formato de Ruta**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Seguridad**: URLs firmadas para acceso seguro
- **Respaldo**: Redundancia automática de S3

## Campos de Respuesta

### Respuesta de Archivo

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | ID de la base de datos |
| `uid` | String! | Identificador único del archivo |
| `name` | String! | Nombre de archivo original |
| `size` | Float! | Tamaño del archivo en bytes |
| `type` | String! | Tipo MIME |
| `extension` | String! | Extensión del archivo |
| `status` | FileStatus | PENDIENTE o CONFIRMADO (nullable) |
| `shared` | Boolean! | Indica si el archivo está compartido |
| `createdAt` | DateTime! | Marca de tiempo de carga |

### Respuesta de TodoCustomFieldFile

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | ID del registro de unión |
| `uid` | String! | Identificador único |
| `position` | Float! | Orden de visualización |
| `file` | File! | Objeto de archivo asociado |
| `todoCustomField` | TodoCustomField! | Campo personalizado padre |
| `createdAt` | DateTime! | Cuándo se adjuntó el archivo |

## Creando Registros con Archivos

Al crear registros, puedes adjuntar archivos utilizando sus UIDs:

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

## Validación y Seguridad de Archivos

### Validación de Carga

- **Verificación de Tipo MIME**: Valida contra tipos permitidos
- **Validación de Extensión de Archivo**: Respaldo para `application/octet-stream`
- **Límites de Tamaño**: Aplicados en el momento de la carga
- **Saneamiento de Nombres de Archivo**: Elimina caracteres especiales

### Control de Acceso

- **Permisos de Carga**: Se requiere membresía en el proyecto/empresa
- **Asociación de Archivos**: Roles ADMIN, OWNER, MEMBER, CLIENT
- **Acceso a Archivos**: Heredado de permisos de proyecto/empresa
- **URLs Seguras**: URLs firmadas con límite de tiempo para acceso a archivos

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Respuestas de Error

### Archivo Demasiado Grande
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

### Archivo No Encontrado
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

### Campo No Encontrado
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

## Mejores Prácticas

### Gestión de Archivos
- Sube archivos antes de adjuntarlos a los registros
- Usa nombres de archivo descriptivos
- Organiza archivos por proyecto/propósito
- Limpia archivos no utilizados periódicamente

### Rendimiento
- Sube archivos en lotes cuando sea posible
- Usa formatos de archivo apropiados para el tipo de contenido
- Comprime archivos grandes antes de la carga
- Considera los requisitos de vista previa de archivos

### Seguridad
- Valida el contenido del archivo, no solo las extensiones
- Usa escaneo de virus para archivos subidos
- Implementa controles de acceso adecuados
- Monitorea patrones de carga de archivos

## Casos de Uso Comunes

1. **Gestión de Documentos**
   - Especificaciones del proyecto
   - Contratos y acuerdos
   - Notas de reuniones y presentaciones
   - Documentación técnica

2. **Gestión de Activos**
   - Archivos de diseño y maquetas
   - Activos de marca y logotipos
   - Materiales de marketing
   - Imágenes de productos

3. **Cumplimiento y Registros**
   - Documentos legales
   - Rutas de auditoría
   - Certificados y licencias
   - Registros financieros

4. **Colaboración**
   - Recursos compartidos
   - Documentos controlados por versiones
   - Comentarios y anotaciones
   - Materiales de referencia

## Características de Integración

### Con Automatizaciones
- Dispara acciones cuando se agregan/eliminan archivos
- Procesa archivos según tipo o metadatos
- Envía notificaciones por cambios en archivos
- Archiva archivos según condiciones

### Con Imágenes de Portada
- Usa campos de archivo como fuentes de imágenes de portada
- Procesamiento automático de imágenes y miniaturas
- Actualizaciones dinámicas de portada cuando cambian los archivos

### Con Búsquedas
- Referencia archivos desde otros registros
- Agrega conteos y tamaños de archivos
- Encuentra registros por metadatos de archivos
- Referencia cruzada de archivos adjuntos

## Limitaciones

- Máximo de 256MB por archivo
- Dependiente de la disponibilidad de S3
- Sin versionado de archivos incorporado
- Sin conversión automática de archivos
- Capacidades limitadas de vista previa de archivos
- Sin edición colaborativa en tiempo real

## Recursos Relacionados

- [API de Carga de Archivos](/api/upload-files) - Puntos finales de carga de archivos
- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales
- [API de Automatizaciones](/api/automations) - Automatizaciones basadas en archivos
- [Documentación de AWS S3](https://docs.aws.amazon.com/s3/) - Backend de almacenamiento