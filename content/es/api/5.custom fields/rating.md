---
title: Campo Personalizado de Calificación
description: Crea campos de calificación para almacenar calificaciones numéricas con escalas y validaciones configurables
---

Los campos personalizados de calificación te permiten almacenar calificaciones numéricas en registros con valores mínimos y máximos configurables. Son ideales para calificaciones de rendimiento, puntajes de satisfacción, niveles de prioridad o cualquier dato basado en escalas numéricas en tus proyectos.

## Ejemplo Básico

Crea un campo de calificación simple con una escala predeterminada de 0-5:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Ejemplo Avanzado

Crea un campo de calificación con escala y descripción personalizadas:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de calificación |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `RATING` |
| `projectId` | String! | ✅ Sí | El ID del proyecto donde se creará este campo |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |
| `min` | Float | No | Valor mínimo de calificación (sin valor predeterminado) |
| `max` | Float | No | Valor máximo de calificación |

## Estableciendo Valores de Calificación

Para establecer o actualizar un valor de calificación en un registro:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput Parámetros

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de calificación |
| `value` | String! | ✅ Sí | Valor de calificación como cadena (dentro del rango configurado) |

## Creando Registros con Valores de Calificación

Al crear un nuevo registro con valores de calificación:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
      }
      value
    }
  }
}
```

## Campos de Respuesta

### Respuesta TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `value` | Float | El valor de calificación almacenado (accedido a través de customField.value) |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

**Nota**: El valor de calificación se accede realmente a través de `customField.value.number` en las consultas.

### Respuesta CustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el campo |
| `name` | String! | Nombre para mostrar del campo de calificación |
| `type` | CustomFieldType! | Siempre `RATING` |
| `min` | Float | Valor mínimo de calificación permitido |
| `max` | Float | Valor máximo de calificación permitido |
| `description` | String | Texto de ayuda para el campo |

## Validación de Calificaciones

### Restricciones de Valor
- Los valores de calificación deben ser numéricos (tipo Float)
- Los valores deben estar dentro del rango mínimo/máximo configurado
- Si no se especifica un mínimo, no hay valor predeterminado
- El valor máximo es opcional pero recomendado

### Reglas de Validación
**Importante**: La validación solo ocurre al enviar formularios, no al usar `setTodoCustomField` directamente.

- La entrada se analiza como un número de punto flotante (al usar formularios)
- Debe ser mayor o igual al valor mínimo (al usar formularios)
- Debe ser menor o igual al valor máximo (al usar formularios)
- `setTodoCustomField` acepta cualquier valor de cadena sin validación

### Ejemplos de Calificación Válidos
Para un campo con min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Ejemplos de Calificación Inválidos
Para un campo con min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Opciones de Configuración

### Configuración de Escala de Calificación
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Escalas de Calificación Comunes
- **1-5 Estrellas**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 Rendimiento**: `min: 1, max: 10`
- **0-100 Porcentaje**: `min: 0, max: 100`
- **Escala Personalizada**: Cualquier rango numérico

## Permisos Requeridos

Las operaciones de campo personalizado siguen permisos estándar basados en roles:

| Acción | Rol Requerido |
|--------|---------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Nota**: Los roles específicos requeridos dependen de la configuración de roles personalizados de tu proyecto y los permisos a nivel de campo.

## Respuestas de Error

### Error de Validación (Solo Formularios)
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

**Importante**: La validación del valor de calificación (restricciones min/max) solo ocurre al enviar formularios, no al usar `setTodoCustomField` directamente.

### Campo Personalizado No Encontrado
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

## Mejores Prácticas

### Diseño de Escala
- Usa escalas de calificación consistentes en campos similares
- Considera la familiaridad del usuario (1-5 estrellas, 0-10 NPS)
- Establece valores mínimos apropiados (0 vs 1)
- Define un significado claro para cada nivel de calificación

### Calidad de Datos
- Valida los valores de calificación antes de almacenarlos
- Usa precisión decimal de manera apropiada
- Considera el redondeo para fines de visualización
- Proporciona orientación clara sobre los significados de calificación

### Experiencia del Usuario
- Muestra escalas de calificación visualmente (estrellas, barras de progreso)
- Muestra el valor actual y los límites de la escala
- Proporciona contexto para los significados de calificación
- Considera valores predeterminados para nuevos registros

## Casos de Uso Comunes

1. **Gestión del Rendimiento**
   - Calificaciones de rendimiento de empleados
   - Puntajes de calidad de proyectos
   - Calificaciones de finalización de tareas
   - Evaluaciones de nivel de habilidad

2. **Retroalimentación del Cliente**
   - Calificaciones de satisfacción
   - Puntajes de calidad del producto
   - Calificaciones de experiencia de servicio
   - Puntaje del Promotor Neto (NPS)

3. **Prioridad e Importancia**
   - Niveles de prioridad de tareas
   - Calificaciones de urgencia
   - Puntajes de evaluación de riesgos
   - Calificaciones de impacto

4. **Aseguramiento de Calidad**
   - Calificaciones de revisión de código
   - Puntajes de calidad de pruebas
   - Calidad de la documentación
   - Calificaciones de adherencia a procesos

## Características de Integración

### Con Automatizaciones
- Disparar acciones basadas en umbrales de calificación
- Enviar notificaciones para calificaciones bajas
- Crear tareas de seguimiento para calificaciones altas
- Dirigir el trabajo basado en valores de calificación

### Con Búsquedas
- Calcular calificaciones promedio a través de registros
- Encontrar registros por rangos de calificación
- Referenciar datos de calificación de otros registros
- Agregar estadísticas de calificación

### Con Blue Frontend
- Validación de rango automática en contextos de formularios
- Controles de entrada de calificación visual
- Retroalimentación de validación en tiempo real
- Opciones de entrada de estrellas o deslizadores

## Seguimiento de Actividades

Los cambios en el campo de calificación se rastrean automáticamente:
- Se registran los valores de calificación antiguos y nuevos
- La actividad muestra cambios numéricos
- Tiempos de registro para todas las actualizaciones de calificación
- Atribución de usuario para cambios

## Limitaciones

- Solo se admiten valores numéricos
- No hay visualización de calificación incorporada (estrellas, etc.)
- La precisión decimal depende de la configuración de la base de datos
- No hay almacenamiento de metadatos de calificación (comentarios, contexto)
- No hay agregación automática de calificaciones o estadísticas
- No hay conversión de calificación incorporada entre escalas
- **Crítico**: La validación min/max solo funciona en formularios, no a través de `setTodoCustomField`

## Recursos Relacionados

- [Campos Numéricos](/api/5.custom%20fields/number) - Para datos numéricos generales
- [Campos de Porcentaje](/api/5.custom%20fields/percent) - Para valores porcentuales
- [Campos de Selección](/api/5.custom%20fields/select-single) - Para calificaciones de elección discreta
- [Descripción General de Campos Personalizados](/api/5.custom%20fields/2.list-custom-fields) - Conceptos generales