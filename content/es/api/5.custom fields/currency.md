---
title: Campo Personalizado de Moneda
description: Crea campos de moneda para rastrear valores monetarios con el formato y la validación adecuados
---

Los campos personalizados de moneda te permiten almacenar y gestionar valores monetarios con códigos de moneda asociados. El campo admite 72 monedas diferentes, incluidas las principales monedas fiduciarias y criptomonedas, con formato automático y restricciones opcionales de mínimo/máximo.

## Ejemplo Básico

Crea un campo de moneda simple:

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Ejemplo Avanzado

Crea un campo de moneda con restricciones de validación:

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de moneda |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `CURRENCY` |
| `currency` | String | No | Código de moneda predeterminado (código ISO de 3 letras) |
| `min` | Float | No | Valor mínimo permitido (almacenado pero no aplicado en actualizaciones) |
| `max` | Float | No | Valor máximo permitido (almacenado pero no aplicado en actualizaciones) |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: El contexto del proyecto se determina automáticamente a partir de tu autenticación. Debes tener acceso al proyecto donde estás creando el campo.

## Estableciendo Valores de Moneda

Para establecer o actualizar un valor de moneda en un registro:

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de moneda |
| `number` | Float! | ✅ Sí | La cantidad monetaria |
| `currency` | String! | ✅ Sí | Código de moneda de 3 letras |

## Creando Registros con Valores de Moneda

Al crear un nuevo registro con valores de moneda:

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### Formato de Entrada para Crear

Al crear registros, los valores de moneda se pasan de manera diferente:

| Parámetro | Tipo | Descripción |
|-----------|------|-------------|
| `customFieldId` | String! | ID del campo de moneda |
| `value` | String! | Cantidad como una cadena (por ejemplo, "1500.50") |
| `currency` | String! | Código de moneda de 3 letras |

## Monedas Soportadas

Blue admite 72 monedas, incluidas 70 monedas fiduciarias y 2 criptomonedas:

### Monedas Fiduciarias

#### Américas
| Moneda | Código | Nombre |
|--------|--------|--------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Bolívar Soberano Venezolano |
| Boliviano Boliviano | `BOB` | Boliviano Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europa
| Moneda | Código | Nombre |
|--------|--------|--------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Corona Noruega | `NOK` | Corona Noruega |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Asia-Pacífico
| Moneda | Código | Nombre |
|--------|--------|--------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### Medio Oriente y África
| Moneda | Código | Nombre |
|--------|--------|--------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Criptomonedas
| Moneda | Código |
|--------|--------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `number` | Float | La cantidad monetaria |
| `currency` | String | El código de moneda de 3 letras |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

## Formateo de Moneda

El sistema formatea automáticamente los valores de moneda según la configuración regional:

- **Ubicación del símbolo**: Coloca correctamente los símbolos de moneda (antes/después)
- **Separadores decimales**: Utiliza separadores específicos de la configuración regional (. o ,)
- **Separadores de miles**: Aplica agrupamiento apropiado
- **Decimales**: Muestra de 0 a 2 decimales según la cantidad
- **Manejo especial**: USD/CAD muestra el prefijo del código de moneda para mayor claridad

### Ejemplos de Formateo

| Valor | Moneda | Visualización |
|-------|--------|---------------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Validación

### Validación de Cantidad
- Debe ser un número válido
- Las restricciones de mínimo/máximo se almacenan con la definición del campo pero no se aplican durante las actualizaciones de valor
- Admite hasta 2 decimales para visualización (la precisión completa se almacena internamente)

### Validación del Código de Moneda
- Debe ser uno de los 72 códigos de moneda admitidos
- Sensible a mayúsculas (usar mayúsculas)
- Códigos no válidos devuelven un error

## Características de Integración

### Fórmulas
Los campos de moneda se pueden usar en campos personalizados de FORMULA para cálculos:
- Sumar múltiples campos de moneda
- Calcular porcentajes
- Realizar operaciones aritméticas

### Conversión de Moneda
Usa campos de CONVERSIÓN_DE_MONEDA para convertir automáticamente entre monedas (ver [Campos de Conversión de Moneda](/api/custom-fields/currency-conversion))

### Automatizaciones
Los valores de moneda pueden activar automatizaciones basadas en:
- Umbrales de cantidad
- Tipo de moneda
- Cambios de valor

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Nota**: Si bien cualquier miembro del proyecto puede crear campos personalizados, la capacidad de establecer valores depende de los permisos basados en roles configurados para cada campo.

## Respuestas de Error

### Valor de Moneda Inválido
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

Este error ocurre cuando:
- El código de moneda no es uno de los 72 códigos admitidos
- El formato del número es inválido
- El valor no se puede analizar correctamente

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

### Selección de Moneda
- Establece una moneda predeterminada que coincida con tu mercado principal
- Usa códigos de moneda ISO 4217 de manera consistente
- Considera la ubicación del usuario al elegir predeterminados

### Restricciones de Valor
- Establece valores mínimos/máximos razonables para prevenir errores de entrada de datos
- Usa 0 como mínimo para campos que no deberían ser negativos
- Considera tu caso de uso al establecer máximos

### Proyectos Multimoneda
- Usa una moneda base consistente para informes
- Implementa campos de CONVERSIÓN_DE_MONEDA para conversión automática
- Documenta qué moneda debe usarse para cada campo

## Casos de Uso Comunes

1. **Presupuestación de Proyectos**
   - Seguimiento del presupuesto del proyecto
   - Estimaciones de costos
   - Seguimiento de gastos

2. **Ventas y Negocios**
   - Valores de negocios
   - Montos de contratos
   - Seguimiento de ingresos

3. **Planificación Financiera**
   - Montos de inversión
   - Rondas de financiamiento
   - Objetivos financieros

4. **Negocios Internacionales**
   - Precios en múltiples monedas
   - Seguimiento de cambios extranjeros
   - Transacciones transfronterizas

## Limitaciones

- Máximo de 2 decimales para visualización (aunque se almacena más precisión)
- No hay conversión de moneda integrada en campos de CURRENCY estándar
- No se pueden mezclar monedas en un solo valor de campo
- No hay actualizaciones automáticas de tasas de cambio (usa CONVERSIÓN_DE_MONEDA para esto)
- Los símbolos de moneda no son personalizables

## Recursos Relacionados

- [Campos de Conversión de Moneda](/api/custom-fields/currency-conversion) - Conversión automática de moneda
- [Campos Numéricos](/api/custom-fields/number) - Para valores numéricos no monetarios
- [Campos de Fórmula](/api/custom-fields/formula) - Calcular con valores de moneda
- [Campos Personalizados de Lista](/api/custom-fields/list-custom-fields) - Consultar todos los campos personalizados en un proyecto