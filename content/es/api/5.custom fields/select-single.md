---
title: Campo Personalizado de Selección Única
description: Crea campos de selección única para permitir a los usuarios elegir una opción de una lista predefinida
---

Los campos personalizados de selección única permiten a los usuarios elegir exactamente una opción de una lista predefinida. Son ideales para campos de estado, categorías, prioridades o cualquier escenario en el que solo se deba hacer una elección de un conjunto controlado de opciones.

## Ejemplo Básico

Crea un campo de selección única simple:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de selección única con opciones predefinidas:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de selección única |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `SELECT_SINGLE` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | No | Opciones iniciales para el campo |

### CreateCustomFieldOptionInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `title` | String! | ✅ Sí | Texto para mostrar de la opción |
| `color` | String | No | Código de color hexadecimal para la opción |

## Agregar Opciones a Campos Existentes

Agrega nuevas opciones a un campo de selección única existente:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## Establecer Valores de Selección Única

Para establecer la opción seleccionada en un registro:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de selección única |
| `customFieldOptionId` | String | No | ID de la opción a seleccionar (preferido para selección única) |
| `customFieldOptionIds` | [String!] | No | Array de IDs de opciones (usa el primer elemento para selección única) |

## Consultar Valores de Selección Única

Consulta el valor de selección única de un registro:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

El campo `value` devuelve un objeto JSON con los detalles de la opción seleccionada.

## Crear Registros con Valores de Selección Única

Al crear un nuevo registro con valores de selección única:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
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
| `value` | JSON | Contiene el objeto de opción seleccionada con id, título, color, posición |
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
| `name` | String! | Nombre para mostrar del campo de selección única |
| `type` | CustomFieldType! | Siempre `SELECT_SINGLE` |
| `description` | String | Texto de ayuda para el campo |
| `customFieldOptions` | [CustomFieldOption!] | Todas las opciones disponibles |

## Formato de Valor

### Formato de Entrada
- **Parámetro de API**: Usa `customFieldOptionId` para ID de opción única
- **Alternativa**: Usa `customFieldOptionIds` array (toma el primer elemento)
- **Limpiar Selección**: Omite ambos campos o pasa valores vacíos

### Formato de Salida
- **Respuesta de GraphQL**: Objeto JSON en `value` campo que contiene {id, título, color, posición}
- **Registro de Actividad**: Título de la opción como cadena
- **Datos de Automatización**: Título de la opción como cadena

## Comportamiento de Selección

### Selección Exclusiva
- Establecer una nueva opción elimina automáticamente la selección anterior
- Solo se puede seleccionar una opción a la vez
- Establecer `null` o un valor vacío limpia la selección

### Lógica de Respaldo
- Si se proporciona un array `customFieldOptionIds`, solo se usa la primera opción
- Esto asegura compatibilidad con formatos de entrada de selección múltiple
- Arrays vacíos o valores nulos limpian la selección

## Gestión de Opciones

### Actualizar Propiedades de Opción
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
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

**Nota**: Eliminar una opción la eliminará de todos los registros donde fue seleccionada.

### Reordenar Opciones
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## Reglas de Validación

### Validación de Opción
- El ID de opción proporcionado debe existir
- La opción debe pertenecer al campo personalizado especificado
- Solo se puede seleccionar una opción (se aplica automáticamente)
- Valores nulos/vacíos son válidos (sin selección)

### Validación de Campo
- Debe tener al menos una opción definida para ser utilizable
- Los títulos de opción deben ser únicos dentro del campo
- Los códigos de color deben ser un formato hexadecimal válido (si se proporcionan)

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Respuestas de Error

### ID de Opción Inválido
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### No se Puede Analizar el Valor
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

## Mejores Prácticas

### Diseño de Opción
- Usa títulos de opción claros y descriptivos
- Aplica codificación de color significativa
- Mantén las listas de opciones enfocadas y relevantes
- Ordena las opciones lógicamente (por prioridad, frecuencia, etc.)

### Patrones de Campo de Estado
- Usa flujos de trabajo de estado consistentes en proyectos
- Considera la progresión natural de las opciones
- Incluye estados finales claros (Hecho, Cancelado, etc.)
- Usa colores que reflejen el significado de la opción

### Gestión de Datos
- Revisa y limpia periódicamente las opciones no utilizadas
- Usa convenciones de nomenclatura consistentes
- Considera el impacto de la eliminación de opciones en registros existentes
- Planifica actualizaciones y migraciones de opciones

## Casos de Uso Comunes

1. **Estado y Flujo de Trabajo**
   - Estado de tarea (Por Hacer, En Progreso, Hecho)
   - Estado de aprobación (Pendiente, Aprobado, Rechazado)
   - Fase del proyecto (Planificación, Desarrollo, Pruebas, Lanzado)
   - Estado de resolución de problemas

2. **Clasificación y Categorización**
   - Niveles de prioridad (Bajo, Medio, Alto, Crítico)
   - Tipos de tarea (Error, Característica, Mejora, Documentación)
   - Categorías de proyecto (Interno, Cliente, Investigación)
   - Asignaciones de departamento

3. **Calidad y Evaluación**
   - Estado de revisión (No Comenzado, En Revisión, Aprobado)
   - Calificaciones de calidad (Pobre, Regular, Buena, Excelente)
   - Niveles de riesgo (Bajo, Medio, Alto)
   - Niveles de confianza

4. **Asignación y Propiedad**
   - Asignaciones de equipo
   - Propiedad de departamento
   - Asignaciones basadas en roles
   - Asignaciones regionales

## Características de Integración

### Con Automatizaciones
- Dispara acciones cuando se seleccionan opciones específicas
- Dirige el trabajo según las categorías seleccionadas
- Envía notificaciones por cambios de estado
- Crea flujos de trabajo condicionales basados en selecciones

### Con Búsquedas
- Filtra registros por opciones seleccionadas
- Referencia datos de opción de otros registros
- Crea informes basados en selecciones de opciones
- Agrupa registros por valores seleccionados

### Con Formularios
- Controles de entrada desplegables
- Interfaces de botones de opción
- Validación y filtrado de opciones
- Visualización condicional de campos según selecciones

## Seguimiento de Actividades

Los cambios en el campo de selección única se rastrean automáticamente:
- Muestra selecciones de opciones antiguas y nuevas
- Muestra títulos de opciones en el registro de actividad
- Tiempos para todos los cambios de selección
- Atribución de usuario para modificaciones

## Diferencias con Selección Múltiple

| Característica | Selección Única | Selección Múltiple |
|----------------|------------------|--------------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Limitaciones

- Solo se puede seleccionar una opción a la vez
- No hay estructura de opción jerárquica o anidada
- Las opciones se comparten entre todos los registros que utilizan el campo
- No hay análisis de opciones o seguimiento de uso incorporado
- Los códigos de color son solo para visualización, sin impacto funcional
- No se pueden establecer diferentes permisos por opción

## Recursos Relacionados

- [Campos de Selección Múltiple](/api/custom-fields/select-multi) - Para selecciones de múltiples opciones
- [Campos de Casilla de Verificación](/api/custom-fields/checkbox) - Para elecciones booleanas simples
- [Campos de Texto](/api/custom-fields/text-single) - Para entrada de texto libre
- [Descripción General de Campos Personalizados](/api/custom-fields/1.index) - Conceptos generales