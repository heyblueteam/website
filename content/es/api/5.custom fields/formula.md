---
title: Campo Personalizado de Fórmula
description: Crea campos calculados que computan automáticamente valores basados en otros datos
---

Los campos personalizados de fórmula se utilizan para cálculos en gráficos y tableros dentro de Blue. Definen funciones de agregación (SUMA, PROMEDIO, CONTAR, etc.) que operan sobre los datos del campo personalizado para mostrar métricas calculadas en gráficos. Las fórmulas no se calculan a nivel de cada tarea individual, sino que agregan datos a través de múltiples registros para fines de visualización.

## Ejemplo Básico

Crea un campo de fórmula para cálculos en gráficos:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Ejemplo Avanzado

Crea una fórmula de moneda con cálculos complejos:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de fórmula |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `FORMULA` |
| `projectId` | String! | ✅ Sí | El ID del proyecto donde se creará este campo |
| `formula` | JSON | No | Definición de la fórmula para cálculos en gráficos |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

### Estructura de la Fórmula

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Funciones Soportadas

### Funciones de Agregación de Gráficos

Los campos de fórmula soportan las siguientes funciones de agregación para cálculos en gráficos:

| Función | Descripción | Enum ChartFunction |
|----------|-------------|-------------------|
| `SUM` | Suma de todos los valores | `SUM` |
| `AVERAGE` | Promedio de valores numéricos | `AVERAGE` |
| `AVERAGEA` | Promedio excluyendo ceros y nulos | `AVERAGEA` |
| `COUNT` | Conteo de valores | `COUNT` |
| `COUNTA` | Conteo excluyendo ceros y nulos | `COUNTA` |
| `MAX` | Valor máximo | `MAX` |
| `MIN` | Valor mínimo | `MIN` |

**Nota**: Estas funciones se utilizan en el `display.function` campo y operan sobre datos agregados para visualizaciones en gráficos. No se admiten expresiones matemáticas complejas o cálculos a nivel de campo.

## Tipos de Visualización

### Visualización de Números

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Resultado: `1250.75`

### Visualización de Moneda

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Resultado: `$1,250.75`

### Visualización de Porcentaje

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Resultado: `87.5%`

## Edición de Campos de Fórmula

Actualiza campos de fórmula existentes:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Procesamiento de Fórmulas

### Contexto de Cálculo de Gráficos

Los campos de fórmula se procesan en el contexto de segmentos de gráficos y tableros:
- Los cálculos ocurren cuando se renderizan o actualizan los gráficos
- Los resultados se almacenan en `ChartSegment.formulaResult` como valores decimales
- El procesamiento se maneja a través de una cola dedicada de BullMQ llamada 'formula'
- Las actualizaciones se publican a los suscriptores del tablero para actualizaciones en tiempo real

### Formateo de Visualización

La función `getFormulaDisplayValue` formatea los resultados calculados según el tipo de visualización:
- **NÚMERO**: Se muestra como número simple con precisión opcional
- **PORCENTAJE**: Agrega sufijo % con precisión opcional  
- **MONEDA**: Formatea utilizando el código de moneda especificado

## Almacenamiento de Resultados de Fórmula

Los resultados se almacenan en el `formulaResult` campo:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Campos de Respuesta

### Respuesta TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo de fórmula |
| `number` | Float | Resultado numérico calculado |
| `formulaResult` | JSON | Resultado completo con formateo de visualización |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se calculó por última vez el valor |

## Contexto de Datos

### Fuente de Datos de Gráfico

Los campos de fórmula operan dentro del contexto de la fuente de datos del gráfico:
- Las fórmulas agregan valores de campos personalizados a través de tareas en un proyecto
- La función de agregación especificada en `display.function` determina el cálculo
- Los resultados se calculan utilizando funciones de agregación SQL (promedio, suma, conteo, etc.)
- Los cálculos se realizan a nivel de base de datos para mayor eficiencia

## Ejemplos Comunes de Fórmulas

### Presupuesto Total (Visualización de Gráfico)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Puntuación Promedio (Visualización de Gráfico)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Conteo de Tareas (Visualización de Gráfico)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Permisos Requeridos

Las operaciones de campo personalizado siguen los permisos estándar basados en roles:

| Acción | Rol Requerido |
|--------|---------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Nota**: Los roles específicos requeridos dependen de la configuración de roles personalizados de tu proyecto. No hay constantes de permisos especiales como CUSTOM_FIELDS_CREATE.

## Manejo de Errores

### Error de Validación
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

### Diseño de Fórmulas
- Usa nombres claros y descriptivos para los campos de fórmula
- Agrega descripciones que expliquen la lógica de cálculo
- Prueba las fórmulas con datos de muestra antes de la implementación
- Mantén las fórmulas simples y legibles

### Optimización del Rendimiento
- Evita dependencias de fórmulas profundamente anidadas
- Usa referencias de campo específicas en lugar de comodines
- Considera estrategias de almacenamiento en caché para cálculos complejos
- Monitorea el rendimiento de las fórmulas en proyectos grandes

### Calidad de Datos
- Valida los datos de origen antes de usarlos en fórmulas
- Maneja adecuadamente los valores vacíos o nulos
- Usa la precisión adecuada para los tipos de visualización
- Considera casos extremos en los cálculos

## Casos de Uso Comunes

1. **Seguimiento Financiero**
   - Cálculos de presupuesto
   - Estados de ganancias/pérdidas
   - Análisis de costos
   - Proyecciones de ingresos

2. **Gestión de Proyectos**
   - Porcentajes de finalización
   - Utilización de recursos
   - Cálculos de cronograma
   - Métricas de rendimiento

3. **Control de Calidad**
   - Puntuaciones promedio
   - Tasas de aprobación/reprobación
   - Métricas de calidad
   - Seguimiento de cumplimiento

4. **Inteligencia Empresarial**
   - Cálculos de KPI
   - Análisis de tendencias
   - Métricas comparativas
   - Valores de tablero

## Limitaciones

- Las fórmulas son solo para agregaciones de gráficos/tableros, no para cálculos a nivel de tarea
- Limitadas a las siete funciones de agregación soportadas (SUMA, PROMEDIO, etc.)
- No se permiten expresiones matemáticas complejas o cálculos de campo a campo
- No se pueden referenciar múltiples campos en una sola fórmula
- Los resultados solo son visibles en gráficos y tableros
- El `logic` campo es solo para texto de visualización, no para la lógica de cálculo real

## Recursos Relacionados

- [Campos Numéricos](/api/5.custom%20fields/number) - Para valores numéricos estáticos
- [Campos de Moneda](/api/5.custom%20fields/currency) - Para valores monetarios
- [Campos de Referencia](/api/5.custom%20fields/reference) - Para datos entre proyectos
- [Campos de Búsqueda](/api/5.custom%20fields/lookup) - Para datos agregados
- [Descripción General de Campos Personalizados](/api/5.custom%20fields/2.list-custom-fields) - Conceptos generales