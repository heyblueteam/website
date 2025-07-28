---
title: Campo Personalizado de Número
description: Crea campos numéricos para almacenar valores numéricos con restricciones opcionales de min/max y formato de prefijo
---

Los campos personalizados de número te permiten almacenar valores numéricos para registros. Soportan restricciones de validación, precisión decimal, y pueden ser utilizados para cantidades, puntuaciones, mediciones, o cualquier dato numérico que no requiera un formato especial.

## Ejemplo Básico

Crea un campo numérico simple:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo numérico con restricciones y prefijo:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre a mostrar del campo numérico |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `NUMBER` |
| `projectId` | String! | ✅ Sí | ID del proyecto en el que crear el campo |
| `min` | Float | No | Restricción de valor mínimo (solo guía de UI) |
| `max` | Float | No | Restricción de valor máximo (solo guía de UI) |
| `prefix` | String | No | Prefijo de visualización (por ejemplo, "#", "~", "$") |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

## Estableciendo Valores Numéricos

Los campos numéricos almacenan valores decimales con validación opcional:

### Valor Numérico Simple

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Valor Entero

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado numérico |
| `number` | Float | No | Valor numérico a almacenar |

## Restricciones de Valor

### Restricciones de Min/Max (Guía de UI)

**Importante**: Las restricciones de min/max se almacenan pero NO se aplican del lado del servidor. Sirven como guía de UI para aplicaciones frontend.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Validación del Lado del Cliente Requerida**: Las aplicaciones frontend deben implementar la lógica de validación para hacer cumplir las restricciones de min/max.

### Tipos de Valor Soportados

| Tipo | Ejemplo | Descripción |
|------|---------|-------------|
| Integer | `42` | Números enteros |
| Decimal | `42.5` | Números con decimales |
| Negative | `-10` | Valores negativos (si no hay restricción mínima) |
| Zero | `0` | Valor cero |

**Nota**: Las restricciones de min/max NO se validan del lado del servidor. Los valores fuera del rango especificado serán aceptados y almacenados.

## Creando Registros con Valores Numéricos

Al crear un nuevo registro con valores numéricos:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
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
        prefix
      }
      number
      value
    }
  }
}
```

### Formatos de Entrada Soportados

Al crear registros, utiliza el parámetro `number` (no `value`) en el arreglo de campos personalizados:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Campos de Respuesta

### Respuesta TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `number` | Float | El valor numérico |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

### Respuesta CustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para la definición del campo |
| `name` | String! | Nombre a mostrar del campo |
| `type` | CustomFieldType! | Siempre `NUMBER` |
| `min` | Float | Valor mínimo permitido |
| `max` | Float | Valor máximo permitido |
| `prefix` | String | Prefijo de visualización |
| `description` | String | Texto de ayuda |

**Nota**: Si el valor numérico no está establecido, el campo `number` será `null`.

## Filtrado y Consulta

Los campos numéricos soportan un filtrado numérico completo:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
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

### Operadores Soportados

| Operador | Descripción | Ejemplo |
|----------|-------------|---------|
| `EQ` | Igual a | `number = 42` |
| `NE` | No igual a | `number ≠ 42` |
| `GT` | Mayor que | `number > 42` |
| `GTE` | Mayor o igual | `number ≥ 42` |
| `LT` | Menor que | `number < 42` |
| `LTE` | Menor o igual | `number ≤ 42` |
| `IN` | En array | `number in [1, 2, 3]` |
| `NIN` | No en array | `number not in [1, 2, 3]` |
| `IS` | Es nulo/no nulo | `number is null` |

### Filtrado por Rango

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Formato de Visualización

### Con Prefijo

Si se establece un prefijo, se mostrará:

| Valor | Prefijo | Visualización |
|-------|--------|---------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Precisión Decimal

Los números mantienen su precisión decimal:

| Entrada | Almacenado | Mostrado |
|-------|--------|-----------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|--------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Respuestas de Error

### Formato de Número Inválido
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Nota**: Los errores de validación de min/max NO ocurren del lado del servidor. La validación de restricciones debe implementarse en tu aplicación frontend.

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

### Diseño de Restricciones
- Establecer valores de min/max realistas para la guía de UI
- Implementar validación del lado del cliente para hacer cumplir las restricciones
- Usar restricciones para proporcionar retroalimentación al usuario en formularios
- Considerar si los valores negativos son válidos para tu caso de uso

### Precisión de Valor
- Usar la precisión decimal adecuada para tus necesidades
- Considerar el redondeo para fines de visualización
- Ser consistente con la precisión en campos relacionados

### Mejora de Visualización
- Usar prefijos significativos para el contexto
- Considerar unidades en los nombres de los campos (por ejemplo, "Peso (kg)")
- Proporcionar descripciones claras para las reglas de validación

## Casos de Uso Comunes

1. **Sistemas de Puntuación**
   - Calificaciones de rendimiento
   - Puntuaciones de calidad
   - Niveles de prioridad
   - Calificaciones de satisfacción del cliente

2. **Mediciones**
   - Cantidades y montos
   - Dimensiones y tamaños
   - Duraciones (en formato numérico)
   - Capacidades y límites

3. **Métricas Empresariales**
   - Cifras de ingresos
   - Tasas de conversión
   - Asignaciones presupuestarias
   - Números objetivo

4. **Datos Técnicos**
   - Números de versión
   - Valores de configuración
   - Métricas de rendimiento
   - Configuraciones de umbral

## Características de Integración

### Con Gráficos y Tableros
- Usar campos NUMERO en cálculos de gráficos
- Crear visualizaciones numéricas
- Rastrear tendencias a lo largo del tiempo

### Con Automatizaciones
- Activar acciones basadas en umbrales numéricos
- Actualizar campos relacionados basados en cambios numéricos
- Enviar notificaciones para valores específicos

### Con Búsquedas
- Agregar números de registros relacionados
- Calcular totales y promedios
- Encontrar valores min/max a través de relaciones

### Con Gráficos
- Crear visualizaciones numéricas
- Rastrear tendencias a lo largo del tiempo
- Comparar valores entre registros

## Limitaciones

- **Sin validación del lado del servidor** de restricciones de min/max
- **Validación del lado del cliente requerida** para hacer cumplir las restricciones
- Sin formato de moneda incorporado (usar el tipo CURRENCY en su lugar)
- Sin símbolo de porcentaje automático (usar el tipo PERCENT en su lugar)
- Sin capacidades de conversión de unidades
- Precisión decimal limitada por el tipo Decimal de la base de datos
- Sin evaluación de fórmulas matemáticas en el campo mismo

## Recursos Relacionados

- [Descripción General de Campos Personalizados](/api/custom-fields/1.index) - Conceptos generales de campos personalizados
- [Campo Personalizado de Moneda](/api/custom-fields/currency) - Para valores monetarios
- [Campo Personalizado de Porcentaje](/api/custom-fields/percent) - Para valores porcentuales
- [API de Automatizaciones](/api/automations/1.index) - Crear automatizaciones basadas en números