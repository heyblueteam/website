---
title: Buscar Campo Personalizado
description: Crear campos de búsqueda que extraen automáticamente datos de registros referenciados
---

Los campos personalizados de búsqueda extraen automáticamente datos de registros referenciados por [Campos de referencia](/api/custom-fields/reference), mostrando información de registros vinculados sin necesidad de copiar manualmente. Se actualizan automáticamente cuando cambian los datos referenciados.

## Ejemplo Básico

Crea un campo de búsqueda para mostrar etiquetas de registros referenciados:

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Ejemplo Avanzado

Crea un campo de búsqueda para extraer valores de campos personalizados de registros referenciados:

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de búsqueda |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Sí | Configuración de búsqueda |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

## Configuración de Búsqueda

### CustomFieldLookupOptionInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `referenceId` | String! | ✅ Sí | ID del campo de referencia del que se extraerán datos |
| `lookupId` | String | No | ID del campo personalizado específico a buscar (requerido para el tipo TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Sí | Tipo de datos a extraer de los registros referenciados |

## Tipos de Búsqueda

### Valores de CustomFieldLookupType

| Tipo | Descripción | Retorna |
|------|-------------|---------|
| `TODO_DUE_DATE` | Fechas de vencimiento de los todos referenciados | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Fechas de creación de los todos referenciados | Array of creation timestamps |
| `TODO_UPDATED_AT` | Fechas de última actualización de los todos referenciados | Array of update timestamps |
| `TODO_TAG` | Etiquetas de los todos referenciados | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Asignados de los todos referenciados | Array of user objects |
| `TODO_DESCRIPTION` | Descripciones de los todos referenciados | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Nombres de listas de todos de los todos referenciados | Array of list titles |
| `TODO_CUSTOM_FIELD` | Valores de campos personalizados de los todos referenciados | Array of values based on the field type |

## Campos de Respuesta

### Respuesta de CustomField (para campos de búsqueda)

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el campo |
| `name` | String! | Nombre para mostrar del campo de búsqueda |
| `type` | CustomFieldType! | Será `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Configuración y resultados de búsqueda |
| `createdAt` | DateTime! | Cuándo se creó el campo |
| `updatedAt` | DateTime! | Cuándo se actualizó por última vez el campo |

### Estructura de CustomFieldLookupOption

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | Tipo de búsqueda que se está realizando |
| `lookupResult` | JSON | Los datos extraídos de los registros referenciados |
| `reference` | CustomField | El campo de referencia que se utiliza como fuente |
| `lookup` | CustomField | El campo específico que se está buscando (para TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | El campo de búsqueda principal |
| `parentLookup` | CustomField | Búsqueda principal en cadena (para búsquedas anidadas) |

## Cómo Funcionan las Búsquedas

1. **Extracción de Datos**: Las búsquedas extraen datos específicos de todos los registros vinculados a través de un campo de referencia.
2. **Actualizaciones Automáticas**: Cuando cambian los registros referenciados, los valores de búsqueda se actualizan automáticamente.
3. **Solo Lectura**: Los campos de búsqueda no se pueden editar directamente; siempre reflejan los datos referenciados actuales.
4. **Sin Cálculos**: Las búsquedas extraen y muestran datos tal como están, sin agregaciones ni cálculos.

## Búsquedas de TODO_CUSTOM_FIELD

Al usar el tipo `TODO_CUSTOM_FIELD`, debes especificar qué campo personalizado extraer utilizando el parámetro `lookupId`:

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

Esto extrae los valores del campo personalizado especificado de todos los registros referenciados.

## Consultando Datos de Búsqueda

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## Ejemplo de Resultados de Búsqueda

### Resultado de Búsqueda de Etiquetas
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### Resultado de Búsqueda de Asignados
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### Resultado de Búsqueda de Campo Personalizado
Los resultados varían según el tipo de campo personalizado que se esté buscando. Por ejemplo, una búsqueda de campo de moneda podría devolver:
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Importante**: Los usuarios deben tener permisos de visualización tanto en el proyecto actual como en el proyecto referenciado para ver los resultados de búsqueda.

## Respuestas de Error

### Campo de Referencia Inválido
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

### Búsqueda Circular Detectada
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### ID de Búsqueda Faltante para TODO_CUSTOM_FIELD
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## Mejores Prácticas

1. **Nombres Claros**: Usa nombres descriptivos que indiquen qué datos se están buscando.
2. **Tipos Apropiados**: Elige el tipo de búsqueda que coincida con tus necesidades de datos.
3. **Rendimiento**: Las búsquedas procesan todos los registros referenciados, así que ten en cuenta los campos de referencia con muchos enlaces.
4. **Permisos**: Asegúrate de que los usuarios tengan acceso a los proyectos referenciados para que las búsquedas funcionen.

## Casos de Uso Comunes

### Visibilidad entre Proyectos
Muestra etiquetas, asignados o estados de proyectos relacionados sin sincronización manual.

### Seguimiento de Dependencias
Muestra fechas de vencimiento o estado de finalización de tareas de las que depende el trabajo actual.

### Resumen de Recursos
Muestra todos los miembros del equipo asignados a tareas referenciadas para la planificación de recursos.

### Agregación de Estado
Recoge todos los estados únicos de tareas relacionadas para ver la salud del proyecto de un vistazo.

## Limitaciones

- Los campos de búsqueda son de solo lectura y no se pueden editar directamente.
- No hay funciones de agregación (SUMA, CONTAR, PROMEDIO) - las búsquedas solo extraen datos.
- No hay opciones de filtrado - todos los registros referenciados están incluidos.
- Se evitan cadenas de búsqueda circulares para evitar bucles infinitos.
- Los resultados reflejan datos actuales y se actualizan automáticamente.

## Recursos Relacionados

- [Campos de Referencia](/api/custom-fields/reference) - Crea enlaces a registros para fuentes de búsqueda.
- [Valores de Campos Personalizados](/api/custom-fields/custom-field-values) - Establece valores en campos personalizados editables.
- [Listar Campos Personalizados](/api/custom-fields/list-custom-fields) - Consulta todos los campos personalizados en un proyecto.