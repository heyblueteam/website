---
title: Campo Personalizado de Referencia
description: Crea campos de referencia que vinculan a registros en otros proyectos para relaciones entre proyectos
---

Los campos personalizados de referencia te permiten crear enlaces entre registros en diferentes proyectos, habilitando relaciones entre proyectos y compartición de datos. Proporcionan una forma poderosa de conectar trabajos relacionados a través de la estructura de proyectos de tu organización.

## Ejemplo Básico

Crea un campo de referencia simple:

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## Ejemplo Avanzado

Crea un campo de referencia con filtrado y selección múltiple:

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre a mostrar del campo de referencia |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `REFERENCE` |
| `referenceProjectId` | String | No | ID del proyecto a referenciar |
| `referenceMultiple` | Boolean | No | Permitir selección de múltiples registros (predeterminado: falso) |
| `referenceFilter` | TodoFilterInput | No | Criterios de filtrado para los registros referenciados |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: Los campos personalizados se asocian automáticamente con el proyecto basado en el contexto del proyecto actual del usuario.

## Configuración de Referencia

### Referencias Únicas vs Múltiples

**Referencia Única (predeterminado):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Los usuarios pueden seleccionar un registro del proyecto referenciado
- Devuelve un único objeto Todo

**Referencias Múltiples:**
```graphql
{
  referenceMultiple: true
}
```
- Los usuarios pueden seleccionar múltiples registros del proyecto referenciado
- Devuelve un array de objetos Todo

### Filtrado de Referencias

Usa `referenceFilter` para limitar qué registros pueden ser seleccionados:

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## Estableciendo Valores de Referencia

### Referencia Única

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Referencias Múltiples

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de referencia |
| `customFieldReferenceTodoIds` | [String!] | ✅ Sí | Array de IDs de registros referenciados |

## Creando Registros con Referencias

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo de referencia |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

**Nota**: Los todos referenciados se acceden a través de `customField.selectedTodos`, no directamente en TodoCustomField.

### Campos de Todo Referenciados

Cada Todo referenciado incluye:

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | Identificador único del registro referenciado |
| `title` | String! | Título del registro referenciado |
| `status` | TodoStatus! | Estado actual (ACTIVO, COMPLETADO, etc.) |
| `description` | String | Descripción del registro referenciado |
| `dueDate` | DateTime | Fecha de vencimiento si está establecida |
| `assignees` | [User!] | Usuarios asignados |
| `tags` | [Tag!] | Etiquetas asociadas |
| `project` | Project! | Proyecto que contiene el registro referenciado |

## Consultando Datos de Referencia

### Consulta Básica

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### Consulta Avanzada con Datos Anidados

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Importante**: Los usuarios deben tener permisos de visualización en el proyecto referenciado para ver los registros vinculados.

## Acceso Entre Proyectos

### Visibilidad del Proyecto

- Los usuarios solo pueden referenciar registros de proyectos a los que tienen acceso
- Los registros referenciados respetan los permisos del proyecto original
- Los cambios en los registros referenciados aparecen en tiempo real
- Eliminar registros referenciados los elimina de los campos de referencia

### Herencia de Permisos

- Los campos de referencia heredan permisos de ambos proyectos
- Los usuarios necesitan acceso de visualización al proyecto referenciado
- Los permisos de edición se basan en las reglas del proyecto actual
- Los datos referenciados son de solo lectura en el contexto del campo de referencia

## Respuestas de Error

### Proyecto de Referencia Inválido

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### Registro Referenciado No Encontrado

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

### Permiso Denegado

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## Mejores Prácticas

### Diseño de Campo

1. **Nombres claros** - Usa nombres descriptivos que indiquen la relación
2. **Filtrado apropiado** - Establece filtros para mostrar solo registros relevantes
3. **Considerar permisos** - Asegúrate de que los usuarios tengan acceso a los proyectos referenciados
4. **Documentar relaciones** - Proporciona descripciones claras de la conexión

### Consideraciones de Rendimiento

1. **Limitar el alcance de referencia** - Usa filtros para reducir el número de registros seleccionables
2. **Evitar anidamientos profundos** - No crees cadenas complejas de referencias
3. **Considerar el almacenamiento en caché** - Los datos referenciados se almacenan en caché para mejorar el rendimiento
4. **Monitorear el uso** - Rastrear cómo se están utilizando las referencias entre proyectos

### Integridad de Datos

1. **Manejar eliminaciones** - Planificar para cuando se eliminen registros referenciados
2. **Validar permisos** - Asegurarse de que los usuarios puedan acceder a los proyectos referenciados
3. **Actualizar dependencias** - Considerar el impacto al cambiar registros referenciados
4. **Rastros de auditoría** - Rastrear relaciones de referencia para cumplir con la normativa

## Casos de Uso Comunes

### Dependencias del Proyecto

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### Requisitos del Cliente

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### Asignación de Recursos

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### Aseguramiento de Calidad

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## Integración con Búsquedas

Los campos de referencia funcionan con [Campos de Búsqueda](/api/custom-fields/lookup) para extraer datos de registros referenciados. Los campos de búsqueda pueden extraer valores de registros seleccionados en campos de referencia, pero son solo extractores de datos (no se admiten funciones de agregación como SUM).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## Limitaciones

- Los proyectos referenciados deben ser accesibles para el usuario
- Los cambios en los permisos del proyecto referenciado afectan el acceso al campo de referencia
- El anidamiento profundo de referencias puede afectar el rendimiento
- No hay validación incorporada para referencias circulares
- No hay restricción automática que impida referencias del mismo proyecto
- La validación de filtros no se aplica al establecer valores de referencia

## Recursos Relacionados

- [Campos de Búsqueda](/api/custom-fields/lookup) - Extraer datos de registros referenciados
- [API de Proyectos](/api/projects) - Gestión de proyectos que contienen referencias
- [API de Registros](/api/records) - Trabajar con registros que tienen referencias
- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales