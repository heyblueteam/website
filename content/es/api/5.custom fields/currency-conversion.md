---
title: Campo Personalizado de Conversión de Moneda
description: Crea campos que convierten automáticamente valores de moneda utilizando tasas de cambio en tiempo real
---

Los campos personalizados de Conversión de Moneda convierten automáticamente los valores de un campo de MONEDA de origen a diferentes monedas de destino utilizando tasas de cambio en tiempo real. Estos campos se actualizan automáticamente cada vez que cambia el valor de la moneda de origen.

Las tasas de conversión son proporcionadas por la [API de Frankfurter](https://github.com/hakanensari/frankfurter), un servicio de código abierto que rastrea las tasas de cambio de referencia publicadas por el [Banco Central Europeo](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Esto asegura conversiones de moneda precisas, confiables y actualizadas para tus necesidades comerciales internacionales.

## Ejemplo Básico

Crea un campo de conversión de moneda simple:

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## Ejemplo Avanzado

Crea un campo de conversión con una fecha específica para tasas históricas:

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## Proceso Completo de Configuración

Configurar un campo de conversión de moneda requiere tres pasos:

### Paso 1: Crear un Campo de MONEDA de Origen

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### Paso 2: Crear el Campo CURRENCY_CONVERSION

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### Paso 3: Crear Opciones de Conversión

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de conversión |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | No | ID del campo de MONEDA de origen del cual convertir |
| `conversionDateType` | String | No | Estrategia de fecha para las tasas de cambio (ver abajo) |
| `conversionDate` | String | No | Cadena de fecha para la conversión (basada en conversionDateType) |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: Los campos personalizados se asocian automáticamente con el proyecto basado en el contexto del proyecto actual del usuario. No se requiere el parámetro `projectId`.

### Tipos de Fecha de Conversión

| Tipo | Descripción | Parámetro conversionDate |
|------|-------------|-------------------------|
| `currentDate` | Usa tasas de cambio en tiempo real | No requerido |
| `specificDate` | Usa tasas de una fecha fija | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Usa la fecha de otro campo | "todoDueDate" or DATE field ID |

## Creando Opciones de Conversión

Las opciones de conversión definen qué pares de monedas pueden ser convertidos:

### CreateCustomFieldOptionInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ Sí | ID del campo CURRENCY_CONVERSION |
| `title` | String! | ✅ Sí | Nombre para mostrar de esta opción de conversión |
| `currencyConversionFrom` | String! | ✅ Sí | Código de moneda de origen o "Cualquiera" |
| `currencyConversionTo` | String! | ✅ Sí | Código de moneda de destino |

### Usando "Cualquiera" como Origen

El valor especial "Cualquiera" como `currencyConversionFrom` crea una opción de respaldo:

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

Esta opción se utilizará cuando no se encuentre una coincidencia específica de par de monedas.

## Cómo Funciona la Conversión Automática

1. **Actualización de Valor**: Cuando se establece un valor en el campo de MONEDA de origen
2. **Coincidencia de Opción**: El sistema encuentra la opción de conversión coincidente basada en la moneda de origen
3. **Obtención de Tasa**: Recupera la tasa de cambio de la API de Frankfurter
4. **Cálculo**: Multiplica la cantidad de origen por la tasa de cambio
5. **Almacenamiento**: Guarda el valor convertido con el código de moneda de destino

### Flujo de Ejemplo

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## Conversiones Basadas en Fecha

### Usando la Fecha Actual

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

Las conversiones se actualizan con las tasas de cambio actuales cada vez que cambia el valor de origen.

### Usando una Fecha Específica

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

Siempre utiliza las tasas de cambio de la fecha especificada.

### Usando la Fecha de un Campo

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

Utiliza la fecha de otro campo (ya sea la fecha de vencimiento de una tarea o un campo personalizado de FECHA).

## Campos de Respuesta

### Respuesta TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo de conversión |
| `number` | Float | La cantidad convertida |
| `currency` | String | El código de moneda de destino |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se actualizó por última vez el valor |

## Fuente de Tasa de Cambio

Blue utiliza la **API de Frankfurter** para las tasas de cambio:
- API de código abierto alojada por el Banco Central Europeo
- Se actualiza diariamente con tasas de cambio oficiales
- Soporta tasas históricas desde 1999
- Gratuita y confiable para uso comercial

## Manejo de Errores

### Fallos de Conversión

Cuando la conversión falla (error de API, moneda no válida, etc.):
- El valor convertido se establece en `0`
- La moneda de destino aún se almacena
- No se lanza ningún error al usuario

### Escenarios Comunes

| Escenario | Resultado |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Sin opción coincidente | Uses "Any" option if available |
| Missing source value | No se realizó ninguna conversión |

## Permisos Requeridos

La gestión de campos personalizados requiere acceso a nivel de proyecto:

| Rol | Puede Crear/Actualizar Campos |
|------|-------------------------|
| `OWNER` | ✅ Sí |
| `ADMIN` | ✅ Sí |
| `MEMBER` | ❌ No |
| `CLIENT` | ❌ No |

Los permisos de visualización para los valores convertidos siguen las reglas estándar de acceso a registros.

## Mejores Prácticas

### Configuración de Opciones
- Crea pares de monedas específicos para conversiones comunes
- Agrega una opción de respaldo "Cualquiera" para flexibilidad
- Usa títulos descriptivos para las opciones

### Selección de Estrategia de Fecha
- Usa `currentDate` para seguimiento financiero en vivo
- Usa `specificDate` para informes históricos
- Usa `fromDateField` para tasas específicas de transacciones

### Consideraciones de Rendimiento
- Múltiples campos de conversión se actualizan en paralelo
- Se realizan llamadas a la API solo cuando cambia el valor de origen
- Las conversiones de la misma moneda omiten las llamadas a la API

## Casos de Uso Comunes

1. **Proyectos Multimoneda**
   - Rastrear costos del proyecto en monedas locales
   - Informar el presupuesto total en la moneda de la empresa
   - Comparar valores entre regiones

2. **Ventas Internacionales**
   - Convertir valores de acuerdos a la moneda de informe
   - Rastrear ingresos en múltiples monedas
   - Conversión histórica para acuerdos cerrados

3. **Informes Financieros**
   - Conversiones de moneda al final del período
   - Estados financieros consolidados
   - Presupuesto vs. real en moneda local

4. **Gestión de Contratos**
   - Convertir valores de contratos en la fecha de firma
   - Rastrear cronogramas de pago en múltiples monedas
   - Evaluación de riesgo cambiario

## Limitaciones

- No se admite la conversión de criptomonedas
- No se pueden establecer valores convertidos manualmente (siempre se calculan)
- Precisión fija de 2 decimales para todos los montos convertidos
- No se admiten tasas de cambio personalizadas
- No hay almacenamiento en caché de tasas de cambio (llamada a la API fresca para cada conversión)
- Depende de la disponibilidad de la API de Frankfurter

## Recursos Relacionados

- [Campos de Moneda](/api/custom-fields/currency) - Campos de origen para conversiones
- [Campos de Fecha](/api/custom-fields/date) - Para conversiones basadas en fecha
- [Campos de Fórmula](/api/custom-fields/formula) - Cálculos alternativos
- [Descripción General de Campos Personalizados](/custom-fields/list-custom-fields) - Conceptos generales