---
title: Campo Personalizado de Selección Múltiple
description: Crea campos de selección múltiple para permitir a los usuarios elegir múltiples opciones de una lista predefinida
---

Los campos personalizados de selección múltiple permiten a los usuarios elegir múltiples opciones de una lista predefinida. Son ideales para categorías, etiquetas, habilidades, características o cualquier escenario donde se necesiten múltiples selecciones de un conjunto controlado de opciones.

## Ejemplo Básico

Crea un campo de selección múltiple simple:

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de selección múltiple y luego agrega opciones por separado:

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de selección múltiple |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `SELECT_MULTI` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |
| `projectId` | String! | ✅ Sí | ID del proyecto para este campo |

### CreateCustomFieldOptionInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado |
| `title` | String! | ✅ Sí | Texto para mostrar de la opción |
| `color` | String | No | Color para la opción (cualquier cadena) |
| `position` | Float | No | Orden de clasificación para la opción |

## Agregar Opciones a Campos Existentes

Agrega nuevas opciones a un campo de selección múltiple existente:

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## Establecer Valores de Selección Múltiple

Para establecer múltiples opciones seleccionadas en un registro:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de selección múltiple |
| `customFieldOptionIds` | [String!] | ✅ Sí | Array de IDs de opciones a seleccionar |

## Crear Registros con Valores de Selección Múltiple

Al crear un nuevo registro con valores de selección múltiple:

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
    }
  }
}
```

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `selectedOptions` | [CustomFieldOption!] | Array de opciones seleccionadas |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

### Respuesta de CustomFieldOption

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para la opción |
| `title` | String! | Texto para mostrar de la opción |
| `color` | String | Código de color hexadecimal para representación visual |
| `position` | Float | Orden de clasificación para la opción |
| `customField` | CustomField! | El campo personalizado al que pertenece esta opción |

### Respuesta de CustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el campo |
| `name` | String! | Nombre para mostrar del campo de selección múltiple |
| `type` | CustomFieldType! | Siempre `SELECT_MULTI` |
| `description` | String | Texto de ayuda para el campo |
| `customFieldOptions` | [CustomFieldOption!] | Todas las opciones disponibles |

## Formato de Valor

### Formato de Entrada
- **Parámetro de API**: Array de IDs de opciones (`["option1", "option2", "option3"]`)
- **Formato de Cadena**: IDs de opciones separados por comas (`"option1,option2,option3"`)

### Formato de Salida
- **Respuesta de GraphQL**: Array de objetos CustomFieldOption
- **Registro de Actividad**: Títulos de opciones separados por comas
- **Datos de Automatización**: Array de títulos de opciones

## Gestión de Opciones

### Actualizar Propiedades de Opción
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
  }) {
    id
    title
    color
  }
}
```

### Eliminar Opción
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### Reordenar Opciones
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## Reglas de Validación

### Validación de Opción
- Todos los IDs de opción proporcionados deben existir
- Las opciones deben pertenecer al campo personalizado especificado
- Solo los campos SELECT_MULTI pueden tener múltiples opciones seleccionadas
- Array vacío es válido (sin selecciones)

### Validación de Campo
- Debe tener al menos una opción definida para ser utilizable
- Los títulos de las opciones deben ser únicos dentro del campo
- El campo de color acepta cualquier valor de cadena (sin validación hexadecimal)

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Respuestas de Error

### ID de Opción Inválido
```json
{
  "errors": [{
    "message": "Custom field option not found",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### La Opción No Pertenece al Campo
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Campo No Encontrado
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Múltiples Opciones en un Campo No Múltiple
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Mejores Prácticas

### Diseño de Opción
- Usa títulos de opción descriptivos y concisos
- Aplica esquemas de codificación de colores consistentes
- Mantén las listas de opciones manejables (típicamente de 3 a 20 opciones)
- Ordena las opciones lógicamente (alfabéticamente, por frecuencia, etc.)

### Gestión de Datos
- Revisa y limpia periódicamente las opciones no utilizadas
- Usa convenciones de nomenclatura consistentes en todos los proyectos
- Considera la reutilización de opciones al crear campos
- Planifica actualizaciones y migraciones de opciones

### Experiencia del Usuario
- Proporciona descripciones claras del campo
- Usa colores para mejorar la distinción visual
- Agrupa opciones relacionadas
- Considera selecciones predeterminadas para casos comunes

## Casos de Uso Comunes

1. **Gestión de Proyectos**
   - Categorías y etiquetas de tareas
   - Niveles y tipos de prioridad
   - Asignaciones de miembros del equipo
   - Indicadores de estado

2. **Gestión de Contenidos**
   - Categorías y temas de artículos
   - Tipos y formatos de contenido
   - Canales de publicación
   - Flujos de trabajo de aprobación

3. **Soporte al Cliente**
   - Categorías y tipos de problemas
   - Productos o servicios afectados
   - Métodos de resolución
   - Segmentos de clientes

4. **Desarrollo de Productos**
   - Categorías de características
   - Requisitos técnicos
   - Entornos de prueba
   - Canales de lanzamiento

## Características de Integración

### Con Automatizaciones
- Disparar acciones cuando se seleccionan opciones específicas
- Dirigir el trabajo según las categorías seleccionadas
- Enviar notificaciones para selecciones de alta prioridad
- Crear tareas de seguimiento basadas en combinaciones de opciones

### Con Búsquedas
- Filtrar registros por opciones seleccionadas
- Agregar datos a través de selecciones de opciones
- Referenciar datos de opciones de otros registros
- Crear informes basados en combinaciones de opciones

### Con Formularios
- Controles de entrada de selección múltiple
- Validación y filtrado de opciones
- Carga dinámica de opciones
- Visualización condicional de campos

## Seguimiento de Actividades

Los cambios en los campos de selección múltiple se rastrean automáticamente:
- Muestra opciones agregadas y eliminadas
- Muestra títulos de opciones en el registro de actividad
- Tiempos de cambio para todas las selecciones
- Atribución de usuario para modificaciones

## Limitaciones

- El límite práctico máximo de opciones depende del rendimiento de la interfaz de usuario
- No hay estructura de opción jerárquica o anidada
- Las opciones se comparten entre todos los registros que utilizan el campo
- No hay análisis de opciones o seguimiento de uso incorporado
- El campo de color acepta cualquier cadena (sin validación hexadecimal)
- No se pueden establecer diferentes permisos por opción
- Las opciones deben crearse por separado, no en línea con la creación del campo
- No hay mutación de reordenamiento dedicada (usa editCustomFieldOption con posición)

## Recursos Relacionados

- [Campos de Selección Única](/api/custom-fields/select-single) - Para selecciones de una sola opción
- [Campos de Casilla de Verificación](/api/custom-fields/checkbox) - Para elecciones booleanas simples
- [Campos de Texto](/api/custom-fields/text-single) - Para entrada de texto libre
- [Resumen de Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceptos generales