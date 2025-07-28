---
title: Campo Personalizado de Texto de Una Línea
description: Crea campos de texto de una línea para valores de texto cortos como nombres, títulos y etiquetas
---

Los campos personalizados de texto de una línea te permiten almacenar valores de texto cortos destinados a una entrada de una sola línea. Son ideales para nombres, títulos, etiquetas o cualquier dato de texto que deba mostrarse en una sola línea.

## Ejemplo Básico

Crea un campo de texto de una línea simple:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de texto de una línea con descripción:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de texto |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `TEXT_SINGLE` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: El contexto del proyecto se determina automáticamente a partir de tus encabezados de autenticación. No se necesita el parámetro `projectId`.

## Estableciendo Valores de Texto

Para establecer o actualizar un valor de texto de una línea en un registro:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo de texto personalizado |
| `text` | String | No | Contenido de texto de una línea a almacenar |

## Creando Registros con Valores de Texto

Al crear un nuevo registro con valores de texto de una línea:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado (contiene el valor de texto) |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

**Importante**: Los valores de texto se acceden a través del campo `customField.value.text`, no directamente en TodoCustomField.

## Consultando Valores de Texto

Al consultar registros con campos de texto personalizados, accede al texto a través de la ruta `customField.value.text`:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

La respuesta incluirá el texto en la estructura anidada:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Validación de Texto

### Validación de Formulario
Cuando se utilizan campos de texto de una línea en formularios:
- Los espacios en blanco al principio y al final se eliminan automáticamente
- Se aplica validación requerida si el campo está marcado como obligatorio
- No se aplica validación de formato específica

### Reglas de Validación
- Acepta cualquier contenido de cadena, incluidos los saltos de línea (aunque no se recomienda)
- Sin límites de longitud de caracteres (hasta los límites de la base de datos)
- Soporta caracteres Unicode y símbolos especiales
- Los saltos de línea se preservan pero no están destinados para este tipo de campo

### Ejemplos Típicos de Texto
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Notas Importantes

### Capacidad de Almacenamiento
- Almacenado utilizando el tipo `MediumText` de MySQL
- Soporta hasta 16MB de contenido de texto
- Almacenamiento idéntico a los campos de texto de múltiples líneas
- Codificación UTF-8 para caracteres internacionales

### API Directa vs Formularios
- **Formularios**: Recorte automático de espacios en blanco y validación requerida
- **API Directa**: El texto se almacena exactamente como se proporciona
- **Recomendación**: Usa formularios para la entrada del usuario para asegurar un formato consistente

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: Entrada de texto de una línea, ideal para valores cortos
- **TEXT_MULTI**: Entrada de área de texto de múltiples líneas, ideal para contenido más largo
- **Backend**: Ambos utilizan almacenamiento y validación idénticos
- **Frontend**: Diferentes componentes de UI para la entrada de datos
- **Intención**: TEXT_SINGLE está semánticamente destinado a valores de una línea

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## Respuestas de Error

### Validación de Campo Requerido (Solo Formularios)
```json
{
  "errors": [{
    "message": "This field is required",
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
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

## Mejores Prácticas

### Directrices de Contenido
- Mantén el texto conciso y apropiado para una línea
- Evita los saltos de línea para la visualización destinada a una sola línea
- Usa un formato consistente para tipos de datos similares
- Considera los límites de caracteres según tus requisitos de UI

### Entrada de Datos
- Proporciona descripciones claras de los campos para guiar a los usuarios
- Usa formularios para la entrada del usuario para asegurar la validación
- Valida el formato del contenido en tu aplicación si es necesario
- Considera usar menús desplegables para valores estandarizados

### Consideraciones de Rendimiento
- Los campos de texto de una línea son ligeros y eficientes
- Considera la indexación para campos buscados con frecuencia
- Usa anchos de visualización apropiados en tu UI
- Monitorea la longitud del contenido para fines de visualización

## Filtrado y Búsqueda

### Búsqueda de Contiene
Los campos de texto de una línea soportan la búsqueda de subcadenas:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Capacidades de Búsqueda
- Coincidencia de subcadenas sin distinción entre mayúsculas y minúsculas
- Soporta coincidencia de palabras parciales
- Coincidencia de valores exactos
- Sin búsqueda de texto completo ni clasificación

## Casos de Uso Comunes

1. **Identificadores y Códigos**
   - SKU de productos
   - Números de pedido
   - Códigos de referencia
   - Números de versión

2. **Nombres y Títulos**
   - Nombres de clientes
   - Títulos de proyectos
   - Nombres de productos
   - Etiquetas de categoría

3. **Descripciones Cortas**
   - Resúmenes breves
   - Etiquetas de estado
   - Indicadores de prioridad
   - Etiquetas de clasificación

4. **Referencias Externas**
   - Números de tickets
   - Referencias de facturas
   - IDs de sistemas externos
   - Números de documentos

## Características de Integración

### Con Búsquedas
- Referencia de datos de texto de otros registros
- Encuentra registros por contenido de texto
- Muestra información de texto relacionada
- Agrega valores de texto de múltiples fuentes

### Con Formularios
- Recorte automático de espacios en blanco
- Validación de campo requerida
- UI de entrada de texto de una línea
- Visualización de límite de caracteres (si está configurado)

### Con Importaciones/Exportaciones
- Mapeo directo de columnas CSV
- Asignación automática de valores de texto
- Soporte para importación de datos en masa
- Exportar a formatos de hoja de cálculo

## Limitaciones

### Restricciones de Automatización
- No disponible directamente como campos de activación de automatización
- No se puede usar en actualizaciones de campos de automatización
- Puede ser referenciado en condiciones de automatización
- Disponible en plantillas de correo electrónico y webhooks

### Limitaciones Generales
- Sin formato o estilo de texto incorporado
- Sin validación automática más allá de los campos requeridos
- Sin imposición de unicidad incorporada
- Sin compresión de contenido para texto muy grande
- Sin versionado o seguimiento de cambios
- Capacidades de búsqueda limitadas (sin búsqueda de texto completo)

## Recursos Relacionados

- [Campos de Texto de Múltiples Líneas](/api/custom-fields/text-multi) - Para contenido de texto más largo
- [Campos de Correo Electrónico](/api/custom-fields/email) - Para direcciones de correo electrónico
- [Campos de URL](/api/custom-fields/url) - Para direcciones de sitios web
- [Campos de ID Único](/api/custom-fields/unique-id) - Para identificadores generados automáticamente
- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales