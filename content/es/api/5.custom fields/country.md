---
title: Campo Personalizado de País
description: Crear campos de selección de país con validación de código de país ISO
---

Los campos personalizados de país te permiten almacenar y gestionar información de país para los registros. El campo admite tanto nombres de países como códigos de país ISO Alpha-2.

**Importante**: El comportamiento de validación y conversión de países difiere significativamente entre las mutaciones:
- **createTodo**: Valida y convierte automáticamente los nombres de países a códigos ISO
- **setTodoCustomField**: Acepta cualquier valor sin validación

## Ejemplo Básico

Crea un campo de país simple:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de país con descripción:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de país |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `COUNTRY` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: El `projectId` no se pasa en la entrada, sino que se determina por el contexto de GraphQL (típicamente de los encabezados de la solicitud o autenticación).

## Estableciendo Valores de País

Los campos de país almacenan datos en dos campos de base de datos:
- **`countryCodes`**: Almacena códigos de país ISO Alpha-2 como una cadena separada por comas en la base de datos (devuelto como un array a través de la API)
- **`text`**: Almacena texto de visualización o nombres de países como una cadena

### Entendiendo los Parámetros

La mutación `setTodoCustomField` acepta dos parámetros opcionales para campos de país:

| Parámetro | Tipo | Requerido | Descripción | Qué hace |
|-----------|------|----------|-------------|--------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar | - |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de país | - |
| `countryCodes` | [String!] | No | Array de códigos de país ISO Alpha-2 | Stored in the `countryCodes` field |
| `text` | String | No | Texto de visualización o nombres de países | Stored in the `text` field |

**Importante**: 
- En `setTodoCustomField`: Ambos parámetros son opcionales y se almacenan de forma independiente
- En `createTodo`: El sistema establece automáticamente ambos campos según tu entrada (no puedes controlarlos de forma independiente)

### Opción 1: Usando Solo Códigos de País

Almacena códigos ISO validados sin texto de visualización:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Resultado: `countryCodes` = `["US"]`, `text` = `null`

### Opción 2: Usando Solo Texto

Almacena texto de visualización sin códigos validados:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Resultado: `countryCodes` = `null`, `text` = `"United States"`

**Nota**: Al usar `setTodoCustomField`, no se realiza ninguna validación independientemente de qué parámetro uses. Los valores se almacenan exactamente como se proporcionan.

### Opción 3: Usando Ambos (Recomendado)

Almacena tanto códigos validados como texto de visualización:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Resultado: `countryCodes` = `["US"]`, `text` = `"United States"`

### Múltiples Países

Almacena múltiples países usando arrays:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Creando Registros con Valores de País

Al crear registros, la mutación `createTodo` **valida y convierte automáticamente** los valores de país. Esta es la única mutación que realiza la validación de país:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      text
      countryCodes
    }
  }
}
```

### Formatos de Entrada Aceptados

| Tipo de Entrada | Ejemplo | Resultado |
|----------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **No soportado** - tratado como un único valor no válido |
| Mixed format | `"United States, CA"` | **No soportado** - tratado como un único valor no válido |

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `text` | String | Texto de visualización (nombres de países) |
| `countryCodes` | [String!] | Array de códigos de país ISO Alpha-2 |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

## Estándares de País

Blue utiliza el estándar **ISO 3166-1 Alpha-2** para códigos de país:

- Códigos de país de dos letras (por ejemplo, US, GB, FR, DE)
- La validación usando la biblioteca `i18n-iso-countries` **solo ocurre en createTodo**
- Soporta todos los países oficialmente reconocidos

### Ejemplo de Códigos de País

| País | Código ISO |
|------|------------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Para la lista oficial completa de códigos de país ISO 3166-1 alpha-2, visita la [Plataforma de Navegación en Línea de ISO](https://www.iso.org/obp/ui/#search/code/).

## Validación

**La validación solo ocurre en la mutación `createTodo`**:

1. **Código ISO Válido**: Acepta cualquier código ISO Alpha-2 válido
2. **Nombre de País**: Convierte automáticamente nombres de países reconocidos a códigos
3. **Entrada No Válida**: Lanza `CustomFieldValueParseError` para valores no reconocidos

**Nota**: La mutación `setTodoCustomField` NO realiza ninguna validación y acepta cualquier valor de cadena.

### Ejemplo de Error

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Características de Integración

### Campos de Búsqueda
Los campos de país pueden ser referenciados por campos personalizados de BÚSQUEDA, lo que te permite extraer datos de país de registros relacionados.

### Automatizaciones
Usa valores de país en condiciones de automatización:
- Filtra acciones por países específicos
- Envía notificaciones basadas en el país
- Dirige tareas según regiones geográficas

### Formularios
Los campos de país en formularios validan automáticamente la entrada del usuario y convierten nombres de países a códigos.

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Respuestas de Error

### Valor de País No Válido
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Desajuste de Tipo de Campo
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Mejores Prácticas

### Manejo de Entrada
- Usa `createTodo` para validación y conversión automáticas
- Usa `setTodoCustomField` con cuidado ya que omite la validación
- Considera validar las entradas en tu aplicación antes de usar `setTodoCustomField`
- Muestra nombres completos de países en la interfaz de usuario para mayor claridad

### Calidad de Datos
- Valida las entradas de país en el punto de entrada
- Usa formatos consistentes en todo tu sistema
- Considera agrupaciones regionales para informes

### Múltiples Países
- Usa el soporte de arrays en `setTodoCustomField` para múltiples países
- Múltiples países en `createTodo` **no son soportados** a través del campo de valor
- Almacena códigos de país como array en `setTodoCustomField` para un manejo adecuado

## Casos de Uso Comunes

1. **Gestión de Clientes**
   - Ubicación de la sede del cliente
   - Destinos de envío
   - Jurisdicciones fiscales

2. **Seguimiento de Proyectos**
   - Ubicación del proyecto
   - Ubicaciones de los miembros del equipo
   - Objetivos de mercado

3. **Cumplimiento y Legal**
   - Jurisdicciones regulatorias
   - Requisitos de residencia de datos
   - Controles de exportación

4. **Ventas y Marketing**
   - Asignaciones de territorio
   - Segmentación de mercado
   - Objetivos de campaña

## Limitaciones

- Solo soporta códigos ISO 3166-1 Alpha-2 (códigos de 2 letras)
- Sin soporte incorporado para subdivisiones de países (estados/provincias)
- Sin íconos de bandera de país automáticos (solo texto)
- No se pueden validar códigos de país históricos
- Sin agrupación regional o de continente incorporada
- **La validación solo funciona en `createTodo`, no en `setTodoCustomField`**
- **Múltiples países no soportados en el campo de valor de `createTodo`**
- **Códigos de país almacenados como cadena separada por comas, no como un verdadero array**

## Recursos Relacionados

- [Descripción General de Campos Personalizados](/custom-fields/list-custom-fields) - Conceptos generales de campos personalizados
- [Campos de Búsqueda](/api/custom-fields/lookup) - Referencia de datos de país de otros registros
- [API de Formularios](/api/forms) - Incluir campos de país en formularios personalizados