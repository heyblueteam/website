---
title: Campo Personalizado de Casilla de Verificación
description: Crea campos de casilla de verificación booleanos para datos de sí/no o verdadero/falso
---

Los campos personalizados de casilla de verificación proporcionan una entrada booleana simple (verdadero/falso) para tareas. Son perfectos para elecciones binarias, indicadores de estado o para rastrear si algo ha sido completado.

## Ejemplo Básico

Crea un campo de casilla de verificación simple:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de casilla de verificación con descripción y validación:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ Sí | Nombre para mostrar de la casilla de verificación |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `CHECKBOX` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

## Estableciendo Valores de Casilla de Verificación

Para establecer o actualizar un valor de casilla de verificación en una tarea:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Para desmarcar una casilla de verificación:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID de la tarea a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de casilla de verificación |
| `checked` | Boolean | No | Verdadero para marcar, falso para desmarcar |

## Creando Tareas con Valores de Casilla de Verificación

Al crear una nueva tarea con valores de casilla de verificación:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Valores de Cadena Aceptados

Al crear tareas, los valores de casilla de verificación deben ser pasados como cadenas:

| Valor de Cadena | Resultado |
|-----------------|----------|
| `"true"` | ✅ Marcada (sensible a mayúsculas) |
| `"1"` | ✅ Marcada |
| `"checked"` | ✅ Marcada (sensible a mayúsculas) |
| Any other value | ❌ Desmarcada |

**Nota**: Las comparaciones de cadenas durante la creación de tareas son sensibles a mayúsculas. Los valores deben coincidir exactamente con `"true"`, `"1"`, o `"checked"` para resultar en un estado marcado.

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | Identificador único para el valor del campo |
| `uid` | String! | Identificador único alternativo |
| `customField` | CustomField! | La definición del campo personalizado |
| `checked` | Boolean | El estado de la casilla de verificación (verdadero/falso/nulo) |
| `todo` | Todo! | La tarea a la que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo fue modificado por última vez el valor |

## Integración de Automatización

Los campos de casilla de verificación desencadenan diferentes eventos de automatización basados en cambios de estado:

| Acción | Evento Desencadenado | Descripción |
|--------|---------------------|-------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Desencadenado cuando la casilla de verificación está marcada |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Desencadenado cuando la casilla de verificación está desmarcada |

Esto te permite crear automatizaciones que respondan a cambios de estado de la casilla de verificación, como:
- Enviar notificaciones cuando los elementos son aprobados
- Mover tareas cuando las casillas de verificación de revisión están marcadas
- Actualizar campos relacionados según los estados de las casillas de verificación

## Importación/Exportación de Datos

### Importando Valores de Casilla de Verificación

Al importar datos a través de CSV u otros formatos:
- `"true"`, `"yes"` → Marcada (sin distinción entre mayúsculas y minúsculas)
- Cualquier otro valor (incluyendo `"false"`, `"no"`, `"0"`, vacío) → Desmarcada

### Exportando Valores de Casilla de Verificación

Al exportar datos:
- Las casillas marcadas se exportan como `"X"`
- Las casillas desmarcadas se exportan como cadena vacía `""`

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Respuestas de Error

### Tipo de Valor Inválido
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Mejores Prácticas

### Convenciones de Nomenclatura
- Usa nombres claros y orientados a la acción: "Aprobado", "Revisado", "Está Completo"
- Evita nombres negativos que confundan a los usuarios: prefiere "Está Activo" sobre "Está Inactivo"
- Sé específico sobre lo que representa la casilla de verificación

### Cuándo Usar Casillas de Verificación
- **Elecciones binarias**: Sí/No, Verdadero/Falso, Hecho/No Hecho
- **Indicadores de estado**: Aprobado, Revisado, Publicado
- **Banderas de características**: Tiene Soporte Prioritario, Requiere Firma
- **Seguimiento simple**: Correo Enviado, Factura Pagada, Artículo Enviado

### Cuándo NO Usar Casillas de Verificación
- Cuando necesitas más de dos opciones (usa SELECT_SINGLE en su lugar)
- Para datos numéricos o de texto (usa campos NUMBER o TEXT)
- Cuando necesitas rastrear quién la marcó o cuándo (usa registros de auditoría)

## Casos de Uso Comunes

1. **Flujos de Trabajo de Aprobación**
   - "Aprobado por el Gerente"
   - "Firma del Cliente"
   - "Revisión Legal Completa"

2. **Gestión de Tareas**
   - "Está Bloqueada"
   - "Listo para Revisión"
   - "Alta Prioridad"

3. **Control de Calidad**
   - "QA Aprobado"
   - "Documentación Completa"
   - "Pruebas Escritas"

4. **Banderas Administrativas**
   - "Factura Enviada"
   - "Contrato Firmado"
   - "Seguimiento Requerido"

## Limitaciones

- Los campos de casilla de verificación solo pueden almacenar valores verdadero/falso (sin estado tri-estable o nulo después de la configuración inicial)
- No hay configuración de valor predeterminado (siempre comienza como nulo hasta que se establece)
- No se puede almacenar metadatos adicionales como quién la marcó o cuándo
- No hay visibilidad condicional basada en otros valores de campo

## Recursos Relacionados

- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales de campos personalizados
- [API de Automatizaciones](/api/automations) - Crea automatizaciones desencadenadas por cambios en las casillas de verificación