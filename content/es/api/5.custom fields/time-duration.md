---
title: Campo Personalizado de Duración de Tiempo
description: Crea campos de duración de tiempo calculados que rastrean el tiempo entre eventos en tu flujo de trabajo
---

Los campos personalizados de duración de tiempo calculan y muestran automáticamente la duración entre dos eventos en tu flujo de trabajo. Son ideales para rastrear tiempos de procesamiento, tiempos de respuesta, tiempos de ciclo o cualquier métrica basada en el tiempo en tus proyectos.

## Ejemplo Básico

Crea un campo de duración de tiempo simple que rastree cuánto tiempo tardan las tareas en completarse:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Ejemplo Avanzado

Crea un campo de duración de tiempo complejo que rastree el tiempo entre cambios de campo personalizado con un objetivo de SLA:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput (TIME_DURATION)

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de duración |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `TIME_DURATION` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Sí | Cómo mostrar la duración |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Sí | Configuración del evento de inicio |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Sí | Configuración del evento de finalización |
| `timeDurationTargetTime` | Float | No | Duración objetivo en segundos para el monitoreo de SLA |

### CustomFieldTimeDurationInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ Sí | Tipo de evento a rastrear |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Sí | `FIRST` o `LAST` ocurrencia |
| `customFieldId` | String | Conditional | Requerido para el tipo `TODO_CUSTOM_FIELD` |
| `customFieldOptionIds` | [String!] | Conditional | Requerido para cambios en campos de selección |
| `todoListId` | String | Conditional | Requerido para el tipo `TODO_MOVED` |
| `tagId` | String | Conditional | Requerido para el tipo `TODO_TAG_ADDED` |
| `assigneeId` | String | Conditional | Requerido para el tipo `TODO_ASSIGNEE_ADDED` |

### Valores de CustomFieldTimeDurationType

| Valor | Descripción |
|-------|-------------|
| `TODO_CREATED_AT` | Cuando se creó el registro |
| `TODO_CUSTOM_FIELD` | Cuando cambió un valor de campo personalizado |
| `TODO_DUE_DATE` | Cuando se estableció la fecha de vencimiento |
| `TODO_MARKED_AS_COMPLETE` | Cuando se marcó el registro como completo |
| `TODO_MOVED` | Cuando se movió el registro a una lista diferente |
| `TODO_TAG_ADDED` | Cuando se agregó una etiqueta al registro |
| `TODO_ASSIGNEE_ADDED` | Cuando se agregó un asignado al registro |

### Valores de CustomFieldTimeDurationCondition

| Valor | Descripción |
|-------|-------------|
| `FIRST` | Usar la primera ocurrencia del evento |
| `LAST` | Usar la última ocurrencia del evento |

### Valores de CustomFieldTimeDurationDisplayType

| Valor | Descripción | Ejemplo |
|-------|-------------|---------|
| `FULL_DATE` | Formato Días:Horas:Minutos:Segundos | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Escrito en palabras completas | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numérico con unidades | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Duración solo en días | `"2.5"` (2.5 days) |
| `HOURS` | Duración solo en horas | `"60"` (60 hours) |
| `MINUTES` | Duración solo en minutos | `"3600"` (3600 minutes) |
| `SECONDS` | Duración solo en segundos | `"216000"` (216000 seconds) |

## Campos de Respuesta

### Respuesta TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `number` | Float | La duración en segundos |
| `value` | Float | Alias para número (duración en segundos) |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se actualizó por última vez el valor |

### Respuesta CustomField (TIME_DURATION)

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Formato de visualización para la duración |
| `timeDurationStart` | CustomFieldTimeDuration | Configuración del evento de inicio |
| `timeDurationEnd` | CustomFieldTimeDuration | Configuración del evento de finalización |
| `timeDurationTargetTime` | Float | Duración objetivo en segundos (para monitoreo de SLA) |

## Cálculo de Duración

### Cómo Funciona
1. **Evento de Inicio**: El sistema monitorea el evento de inicio especificado
2. **Evento de Finalización**: El sistema monitorea el evento de finalización especificado
3. **Cálculo**: Duración = Hora de Finalización - Hora de Inicio
4. **Almacenamiento**: Duración almacenada en segundos como un número
5. **Visualización**: Formateada de acuerdo con la configuración `timeDurationDisplay`

### Disparadores de Actualización
Los valores de duración se recalculan automáticamente cuando:
- Se crean o actualizan registros
- Cambian los valores de campos personalizados
- Se agregan o eliminan etiquetas
- Se agregan o eliminan asignados
- Se mueven registros entre listas
- Se marcan registros como completos/incompletos

## Lectura de Valores de Duración

### Consultar Campos de Duración
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Valores de Visualización Formateados
Los valores de duración se formatean automáticamente según la configuración `timeDurationDisplay`:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Ejemplos Comunes de Configuración

### Tiempo de Finalización de Tareas
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Duración del Cambio de Estado
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Tiempo en Lista Específica
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Tiempo de Respuesta de Asignación
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Respuestas de Error

### Configuración Inválida
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Campo Referenciado No Encontrado
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

### Opciones Requeridas Faltantes
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Notas Importantes

### Cálculo Automático
- Los campos de duración son **solo lectura** - los valores se calculan automáticamente
- No puedes establecer manualmente los valores de duración a través de la API
- Los cálculos ocurren de manera asíncrona a través de trabajos en segundo plano
- Los valores se actualizan automáticamente cuando ocurren eventos de activación

### Consideraciones de Rendimiento
- Los cálculos de duración se colocan en cola y se procesan de manera asíncrona
- Un gran número de campos de duración puede afectar el rendimiento
- Considera la frecuencia de los eventos de activación al diseñar campos de duración
- Usa condiciones específicas para evitar recalculos innecesarios

### Valores Nulos
Los campos de duración mostrarán `null` cuando:
- El evento de inicio aún no ha ocurrido
- El evento de finalización aún no ha ocurrido
- La configuración hace referencia a entidades no existentes
- El cálculo encuentra un error

## Mejores Prácticas

### Diseño de Configuración
- Usa tipos de eventos específicos en lugar de genéricos cuando sea posible
- Elige condiciones apropiadas `FIRST` vs `LAST` según tu flujo de trabajo
- Prueba los cálculos de duración con datos de muestra antes de la implementación
- Documenta la lógica de tu campo de duración para los miembros del equipo

### Formato de Visualización
- Usa `FULL_DATE_SUBSTRING` para el formato más legible
- Usa `FULL_DATE` para una visualización compacta y de ancho consistente
- Usa `FULL_DATE_STRING` para informes y documentos formales
- Usa `DAYS`, `HOURS`, `MINUTES`, o `SECONDS` para visualizaciones numéricas simples
- Considera las limitaciones de espacio de tu interfaz de usuario al elegir el formato

### Monitoreo de SLA con Tiempo Objetivo
Al usar `timeDurationTargetTime`:
- Establece la duración objetivo en segundos
- Compara la duración real con la duración objetivo para el cumplimiento de SLA
- Usa en paneles para resaltar elementos atrasados
- Ejemplo: SLA de respuesta de 24 horas = 86400 segundos

### Integración en el Flujo de Trabajo
- Diseña campos de duración para que coincidan con tus procesos comerciales reales
- Usa datos de duración para la mejora y optimización de procesos
- Monitorea las tendencias de duración para identificar cuellos de botella en el flujo de trabajo
- Configura alertas para umbrales de duración si es necesario

## Casos de Uso Comunes

1. **Rendimiento del Proceso**
   - Tiempos de finalización de tareas
   - Tiempos de ciclo de revisión
   - Tiempos de procesamiento de aprobación
   - Tiempos de respuesta

2. **Monitoreo de SLA**
   - Tiempo hasta la primera respuesta
   - Tiempos de resolución
   - Plazos de escalación
   - Cumplimiento del nivel de servicio

3. **Analítica de Flujo de Trabajo**
   - Identificación de cuellos de botella
   - Optimización de procesos
   - Métricas de rendimiento del equipo
   - Tiempos de aseguramiento de calidad

4. **Gestión de Proyectos**
   - Duraciones de fase
   - Tiempos de hitos
   - Tiempo de asignación de recursos
   - Plazos de entrega

## Limitaciones

- Los campos de duración son **solo lectura** y no se pueden establecer manualmente
- Los valores se calculan de manera asíncrona y pueden no estar disponibles de inmediato
- Requiere que se configuren correctamente los disparadores de eventos en tu flujo de trabajo
- No se pueden calcular duraciones para eventos que no han ocurrido
- Limitado a rastrear el tiempo entre eventos discretos (no seguimiento de tiempo continuo)
- Sin alertas o notificaciones de SLA integradas
- No se pueden agregar múltiples cálculos de duración en un solo campo

## Recursos Relacionados

- [Campos Numéricos](/api/custom-fields/number) - Para valores numéricos manuales
- [Campos de Fecha](/api/custom-fields/date) - Para el seguimiento de fechas específicas
- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales
- [Automatizaciones](/api/automations) - Para activar acciones basadas en umbrales de duración