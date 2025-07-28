---
title: Campo Personalizado de Porcentaje
description: Crea campos de porcentaje para almacenar valores numéricos con manejo automático del símbolo % y formato de visualización
---

Los campos personalizados de porcentaje te permiten almacenar valores porcentuales para registros. Manejan automáticamente el símbolo % para la entrada y la visualización, mientras almacenan el valor numérico bruto internamente. Perfecto para tasas de finalización, tasas de éxito o cualquier métrica basada en porcentajes.

## Ejemplo Básico

Crea un campo de porcentaje simple:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de porcentaje con descripción:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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
| `name` | String! | ✅ Sí | Nombre de visualización del campo de porcentaje |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `PERCENT` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: El contexto del proyecto se determina automáticamente a partir de tus encabezados de autenticación. No se necesita el parámetro `projectId`.

**Nota**: Los campos PERCENT no admiten restricciones de min/max ni formato de prefijo como los campos NUMBER.

## Estableciendo Valores de Porcentaje

Los campos de porcentaje almacenan valores numéricos con manejo automático del símbolo %:

### Con Símbolo de Porcentaje

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Valor Numérico Directo

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de porcentaje |
| `number` | Float | No | Valor porcentual numérico (p. ej., 75.5 para 75.5%) |

## Almacenamiento y Visualización de Valores

### Formato de Almacenamiento
- **Almacenamiento interno**: Valor numérico bruto (p. ej., 75.5)
- **Base de datos**: Almacenado como `Decimal` en la columna `number`
- **GraphQL**: Devuelto como tipo `Float`

### Formato de Visualización
- **Interfaz de usuario**: Las aplicaciones cliente deben agregar el símbolo % (p. ej., "75.5%")
- **Gráficos**: Se muestra con el símbolo % cuando el tipo de salida es PERCENTAGE
- **Respuestas de API**: Valor numérico bruto sin el símbolo % (p. ej., 75.5)

## Creando Registros con Valores de Porcentaje

Al crear un nuevo registro con valores de porcentaje:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Formatos de Entrada Admitidos

| Formato | Ejemplo | Resultado |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Nota**: El símbolo % se elimina automáticamente de la entrada y se agrega nuevamente durante la visualización.

## Consultando Valores de Porcentaje

Al consultar registros con campos personalizados de porcentaje, accede al valor a través de la ruta `customField.value.number`:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

La respuesta incluirá el porcentaje como un número bruto:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Campos de Respuesta

### Respuesta TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado (contiene el valor porcentual) |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

**Importante**: Los valores porcentuales se acceden a través del campo `customField.value.number`. El símbolo % no se incluye en los valores almacenados y debe ser agregado por las aplicaciones cliente para la visualización.

## Filtrado y Consulta

Los campos de porcentaje admiten el mismo filtrado que los campos NUMBER:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Operadores Admitidos

| Operador | Descripción | Ejemplo |
|----------|-------------|---------|
| `EQ` | Igual a | `percentage = 75` |
| `NE` | No igual a | `percentage ≠ 75` |
| `GT` | Mayor que | `percentage > 75` |
| `GTE` | Mayor o igual | `percentage ≥ 75` |
| `LT` | Menor que | `percentage < 75` |
| `LTE` | Menor o igual | `percentage ≤ 75` |
| `IN` | Valor en lista | `percentage in [50, 75, 100]` |
| `NIN` | Valor no en lista | `percentage not in [0, 25]` |
| `IS` | Comprobar si es nulo con `values: null` | `percentage is null` |
| `NOT` | Comprobar si no es nulo con `values: null` | `percentage is not null` |

### Filtrado por Rango

Para el filtrado por rango, usa múltiples operadores:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Rangos de Valores de Porcentaje

### Rangos Comunes

| Rango | Descripción | Caso de Uso |
|-------|-------------|----------|
| `0-100` | Porcentaje estándar | Completion rates, success rates |
| `0-∞` | Porcentaje ilimitado | Growth rates, performance metrics |
| `-∞-∞` | Cualquier valor | Change rates, variance |

### Valores de Ejemplo

| Entrada | Almacenado | Visualización |
|-------|--------|---------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Agregación de Gráficos

Los campos de porcentaje admiten la agregación en gráficos de panel y reportes. Las funciones disponibles incluyen:

- `AVERAGE` - Valor porcentual medio
- `COUNT` - Número de registros con valores
- `MIN` - Valor porcentual más bajo
- `MAX` - Valor porcentual más alto 
- `SUM` - Total de todos los valores porcentuales

Estas agregaciones están disponibles al crear gráficos y paneles, no en consultas GraphQL directas.

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Respuestas de Error

### Formato de Porcentaje Inválido
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### No es un Número
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Mejores Prácticas

### Entrada de Valor
- Permitir a los usuarios ingresar con o sin símbolo %
- Validar rangos razonables para tu caso de uso
- Proporcionar un contexto claro sobre lo que representa el 100%

### Visualización
- Siempre mostrar el símbolo % en las interfaces de usuario
- Usar una precisión decimal apropiada
- Considerar la codificación de colores para rangos (rojo/amarillo/verde)

### Interpretación de Datos
- Documentar lo que significa el 100% en tu contexto
- Manejar valores superiores al 100% de manera apropiada
- Considerar si los valores negativos son válidos

## Casos de Uso Comunes

1. **Gestión de Proyectos**
   - Tasas de finalización de tareas
   - Progreso del proyecto
   - Utilización de recursos
   - Velocidad de sprint

2. **Seguimiento de Rendimiento**
   - Tasas de éxito
   - Tasas de error
   - Métricas de eficiencia
   - Puntuaciones de calidad

3. **Métricas Financieras**
   - Tasas de crecimiento
   - Márgenes de beneficio
   - Montos de descuento
   - Porcentajes de cambio

4. **Analítica**
   - Tasas de conversión
   - Tasas de clics
   - Métricas de participación
   - Indicadores de rendimiento

## Características de Integración

### Con Fórmulas
- Referenciar campos PERCENT en cálculos
- Formato automático del símbolo % en salidas de fórmulas
- Combinar con otros campos numéricos

### Con Automatizaciones
- Activar acciones basadas en umbrales porcentuales
- Enviar notificaciones para porcentajes de hitos
- Actualizar estado basado en tasas de finalización

### Con Búsquedas
- Agregar porcentajes de registros relacionados
- Calcular tasas de éxito promedio
- Encontrar elementos de mejor/peor rendimiento

### Con Gráficos
- Crear visualizaciones basadas en porcentajes
- Rastrear el progreso a lo largo del tiempo
- Comparar métricas de rendimiento

## Diferencias con Campos NUMBER

### Qué es Diferente
- **Manejo de entrada**: Elimina automáticamente el símbolo %
- **Visualización**: Agrega automáticamente el símbolo %
- **Restricciones**: Sin validación de min/max
- **Formato**: Sin soporte de prefijo

### Qué es Igual
- **Almacenamiento**: Misma columna y tipo de base de datos
- **Filtrado**: Mismos operadores de consulta
- **Agregación**: Mismas funciones de agregación
- **Permisos**: Mismo modelo de permisos

## Limitaciones

- Sin restricciones de valores min/max
- Sin opciones de formato de prefijo
- Sin validación automática del rango 0-100%
- Sin conversión entre formatos de porcentaje (p. ej., 0.75 ↔ 75%)
- Se permiten valores superiores al 100%

## Recursos Relacionados

- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales de campos personalizados
- [Campo Personalizado de Número](/api/custom-fields/number) - Para valores numéricos brutos
- [API de Automatizaciones](/api/automations/index) - Crear automatizaciones basadas en porcentajes