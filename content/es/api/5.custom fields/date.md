---
title: Campo Personalizado de Fecha
description: Crea campos de fecha para rastrear fechas únicas o rangos de fechas con soporte de zona horaria
---

Los campos personalizados de fecha te permiten almacenar fechas únicas o rangos de fechas para registros. Soportan el manejo de zonas horarias, formateo inteligente y pueden ser utilizados para rastrear plazos, fechas de eventos o cualquier información basada en el tiempo.

## Ejemplo Básico

Crea un campo de fecha simple:

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de fecha de vencimiento con descripción:

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de fecha |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `DATE` |
| `isDueDate` | Boolean | No | Si este campo representa una fecha de vencimiento |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: Los campos personalizados se asocian automáticamente con el proyecto basado en el contexto del proyecto actual del usuario. No se requiere ningún parámetro `projectId`.

## Estableciendo Valores de Fecha

Los campos de fecha pueden almacenar una fecha única o un rango de fechas:

### Fecha Única

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Rango de Fechas

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Evento Todo el Día

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de fecha |
| `startDate` | DateTime | No | Fecha/hora de inicio en formato ISO 8601 |
| `endDate` | DateTime | No | Fecha/hora de finalización en formato ISO 8601 |
| `timezone` | String | No | Identificador de zona horaria (por ejemplo, "America/New_York") |

**Nota**: Si solo se proporciona `startDate`, `endDate` automáticamente se establece en el mismo valor.

## Formatos de Fecha

### Formato ISO 8601
Todas las fechas deben ser proporcionadas en formato ISO 8601:
- `2025-01-15T14:30:00Z` - Tiempo UTC
- `2025-01-15T14:30:00+05:00` - Con desplazamiento de zona horaria
- `2025-01-15T14:30:00.123Z` - Con milisegundos

### Identificadores de Zona Horaria
Utiliza identificadores de zona horaria estándar:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Si no se proporciona ninguna zona horaria, el sistema utiliza la zona horaria detectada del usuario.

## Creando Registros con Valores de Fecha

Al crear un nuevo registro con valores de fecha:

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### Formatos de Entrada Soportados

Al crear registros, las fechas pueden ser proporcionadas en varios formatos:

| Formato | Ejemplo | Resultado |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Campos de Respuesta

### Respuesta TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | Identificador único para el valor del campo |
| `uid` | String! | Cadena de identificador único |
| `customField` | CustomField! | La definición del campo personalizado (contiene los valores de fecha) |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

**Importante**: Los valores de fecha (`startDate`, `endDate`, `timezone`) se acceden a través del campo `customField.value`, no directamente en TodoCustomField.

### Estructura del Objeto de Valor

Los valores de fecha se devuelven a través del campo `customField.value` como un objeto JSON:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**Nota**: El campo `value` es del tipo `CustomField`, no de `TodoCustomField`.

## Consultando Valores de Fecha

Al consultar registros con campos personalizados de fecha, accede a los valores de fecha a través del campo `customField.value`:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

La respuesta incluirá los valores de fecha en el campo `value`:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## Inteligencia de Visualización de Fechas

El sistema formatea automáticamente las fechas según el rango:

| Escenario | Formato de Visualización |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (sin hora mostrada) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Detección de todo el día**: Los eventos de 00:00 a 23:59 se detectan automáticamente como eventos de todo el día.

## Manejo de Zonas Horarias

### Almacenamiento
- Todas las fechas se almacenan en UTC en la base de datos
- La información de la zona horaria se conserva por separado
- La conversión ocurre en la visualización

### Mejores Prácticas
- Siempre proporciona la zona horaria para mayor precisión
- Usa zonas horarias consistentes dentro de un proyecto
- Considera las ubicaciones de los usuarios para equipos globales

### Zonas Horarias Comunes

| Región | ID de Zona Horaria | Desplazamiento UTC |
|--------|-------------------|--------------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filtrado y Consulta

Los campos de fecha soportan filtrado complejo:

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### Comprobando Campos de Fecha Vacíos

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### Operadores Soportados

| Operador | Uso | Descripción |
|----------|-------|-------------|
| `EQ` | Con dateRange | La fecha se superpone con el rango especificado (cualquier intersección) |
| `NE` | Con dateRange | La fecha no se superpone con el rango |
| `IS` | Con `values: null` | El campo de fecha está vacío (startDate o endDate es nulo) |
| `NOT` | Con `values: null` | El campo de fecha tiene un valor (ambas fechas no son nulas) |

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Respuestas de Error

### Formato de Fecha Inválido
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
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
      "code": "NOT_FOUND"
    }
  }]
}
```


## Limitaciones

- No hay soporte para fechas recurrentes (usa automatizaciones para eventos recurrentes)
- No se puede establecer la hora sin la fecha
- No hay cálculo de días laborables incorporado
- Los rangos de fechas no validan automáticamente fin > inicio
- La precisión máxima es hasta el segundo (sin almacenamiento de milisegundos)

## Recursos Relacionados

- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales de campos personalizados
- [API de Automatizaciones](/api/automations/index) - Crear automatizaciones basadas en fechas