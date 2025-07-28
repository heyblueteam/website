---
title: Campo Personalizado de Texto Multilínea
description: Crea campos de texto multilínea para contenido más extenso como descripciones, notas y comentarios
---

Los campos personalizados de texto multilínea te permiten almacenar contenido de texto más largo con saltos de línea y formato. Son ideales para descripciones, notas, comentarios o cualquier dato de texto que necesite múltiples líneas.

## Ejemplo Básico

Crea un campo de texto multilínea simple:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de texto multilínea con descripción:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `type` | CustomFieldType! | ✅ Sí | Debe ser `TEXT_MULTI` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota:** El `projectId` se pasa como un argumento separado a la mutación, no como parte del objeto de entrada. Alternativamente, el contexto del proyecto se puede determinar a partir del encabezado `X-Bloo-Project-ID` en tu solicitud GraphQL.

## Estableciendo Valores de Texto

Para establecer o actualizar un valor de texto multilínea en un registro:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo de texto personalizado |
| `text` | String | No | Contenido de texto multilínea a almacenar |

## Creando Registros con Valores de Texto

Al crear un nuevo registro con valores de texto multilínea:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
    }
  }
}
```

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `text` | String | El contenido de texto multilínea almacenado |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

## Validación de Texto

### Validación de Formularios
Cuando se utilizan campos de texto multilínea en formularios:
- Los espacios en blanco al principio y al final se recortan automáticamente
- Se aplica validación requerida si el campo está marcado como obligatorio
- No se aplica validación de formato específica

### Reglas de Validación
- Acepta cualquier contenido de cadena, incluidos los saltos de línea
- Sin límites de longitud de caracteres (hasta los límites de la base de datos)
- Soporta caracteres Unicode y símbolos especiales
- Los saltos de línea se conservan en el almacenamiento

### Ejemplos de Texto Válido
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Notas Importantes

### Capacidad de Almacenamiento
- Almacenado utilizando el tipo `MediumText` de MySQL
- Soporta hasta 16MB de contenido de texto
- Se conservan los saltos de línea y el formato
- Codificación UTF-8 para caracteres internacionales

### API Directa vs Formularios
- **Formularios**: Recorte automático de espacios en blanco y validación requerida
- **API Directa**: El texto se almacena exactamente como se proporciona
- **Recomendación**: Usa formularios para la entrada del usuario para asegurar un formato consistente

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI**: Entrada de área de texto multilínea, ideal para contenido más largo
- **TEXT_SINGLE**: Entrada de texto de una sola línea, ideal para valores cortos
- **Backend**: Ambos tipos son idénticos: mismo campo de almacenamiento, validación y procesamiento
- **Frontend**: Diferentes componentes de UI para la entrada de datos (área de texto vs campo de entrada)
- **Importante**: La distinción entre TEXT_MULTI y TEXT_SINGLE existe puramente por razones de UI

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Mejores Prácticas

### Organización del Contenido
- Usa un formato consistente para contenido estructurado
- Considera usar una sintaxis similar a markdown para mejorar la legibilidad
- Divide contenido largo en secciones lógicas
- Usa saltos de línea para mejorar la legibilidad

### Entrada de Datos
- Proporciona descripciones claras de los campos para guiar a los usuarios
- Usa formularios para la entrada del usuario para asegurar la validación
- Considera límites de caracteres según tu caso de uso
- Valida el formato del contenido en tu aplicación si es necesario

### Consideraciones de Rendimiento
- Contenido de texto muy largo puede afectar el rendimiento de las consultas
- Considera la paginación para mostrar campos de texto grandes
- Consideraciones de índice para la funcionalidad de búsqueda
- Monitorea el uso de almacenamiento para campos con contenido grande

## Filtrado y Búsqueda

### Búsqueda de Contiene
Los campos de texto multilínea soportan la búsqueda de subcadenas a través de filtros de campo personalizados:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Capacidades de Búsqueda
- Coincidencia de subcadenas dentro de campos de texto usando el operador `CONTAINS`
- Búsqueda sin distinción de mayúsculas usando el operador `NCONTAINS`
- Coincidencia exacta usando el operador `IS`
- Coincidencia negativa usando el operador `NOT`
- Búsquedas a través de todas las líneas de texto
- Soporta coincidencias de palabras parciales

## Casos de Uso Comunes

1. **Gestión de Proyectos**
   - Descripciones de tareas
   - Requisitos del proyecto
   - Notas de reuniones
   - Actualizaciones de estado

2. **Soporte al Cliente**
   - Descripciones de problemas
   - Notas de resolución
   - Comentarios de clientes
   - Registros de comunicación

3. **Gestión de Contenido**
   - Contenido de artículos
   - Descripciones de productos
   - Comentarios de usuarios
   - Detalles de reseñas

4. **Documentación**
   - Descripciones de procesos
   - Instrucciones
   - Directrices
   - Materiales de referencia

## Características de Integración

### Con Automatizaciones
- Activar acciones cuando cambia el contenido de texto
- Extraer palabras clave del contenido de texto
- Crear resúmenes o notificaciones
- Procesar contenido de texto con servicios externos

### Con Búsquedas
- Referenciar datos de texto de otros registros
- Agregar contenido de texto de múltiples fuentes
- Encontrar registros por contenido de texto
- Mostrar información de texto relacionada

### Con Formularios
- Recorte automático de espacios en blanco
- Validación de campo requerido
- UI de área de texto multilínea
- Visualización del conteo de caracteres (si está configurado)

## Limitaciones

- Sin formato de texto integrado o edición de texto enriquecido
- Sin detección o conversión automática de enlaces
- Sin verificación ortográfica o validación gramatical
- Sin análisis o procesamiento de texto integrado
- Sin versionado o seguimiento de cambios
- Capacidades de búsqueda limitadas (sin búsqueda de texto completo)
- Sin compresión de contenido para texto muy grande

## Recursos Relacionados

- [Campos de Texto de Una Línea](/api/custom-fields/text-single) - Para valores de texto cortos
- [Campos de Correo Electrónico](/api/custom-fields/email) - Para direcciones de correo electrónico
- [Campos de URL](/api/custom-fields/url) - Para direcciones de sitios web
- [Descripción General de Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceptos generales