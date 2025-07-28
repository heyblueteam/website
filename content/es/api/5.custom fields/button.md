---
title: Campo Personalizado de Botón
description: Crea campos de botón interactivos que activan automatizaciones al hacer clic
---

Los campos personalizados de botón proporcionan elementos de interfaz de usuario interactivos que activan automatizaciones al hacer clic. A diferencia de otros tipos de campos personalizados que almacenan datos, los campos de botón sirven como disparadores de acción para ejecutar flujos de trabajo configurados.

## Ejemplo Básico

Crea un campo de botón simple que activa una automatización:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un botón con requisitos de confirmación:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del botón |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `BUTTON` |
| `projectId` | String! | ✅ Sí | ID del proyecto donde se creará el campo |
| `buttonType` | String | No | Comportamiento de confirmación (ver Tipos de Botón a continuación) |
| `buttonConfirmText` | String | No | Texto que los usuarios deben escribir para confirmación dura |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |
| `required` | Boolean | No | Si el campo es requerido (por defecto es falso) |
| `isActive` | Boolean | No | Si el campo está activo (por defecto es verdadero) |

### Campo de Tipo de Botón

El campo `buttonType` es una cadena de texto libre que puede ser utilizada por los clientes de la interfaz de usuario para determinar el comportamiento de confirmación. Los valores comunes incluyen:

- `""` (vacío) - Sin confirmación
- `"soft"` - Diálogo de confirmación simple
- `"hard"` - Requiere escribir texto de confirmación

**Nota**: Estos son solo indicios para la interfaz de usuario. La API no valida ni impone valores específicos.

## Activando Clics en Botones

Para activar un clic en un botón y ejecutar automatizaciones asociadas:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Parámetros de Entrada para Clic

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID de la tarea que contiene el botón |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de botón |

### Importante: Comportamiento de la API

**Todos los clics en botones a través de la API se ejecutan inmediatamente** independientemente de cualquier configuración de `buttonType` o `buttonConfirmText`. Estos campos se almacenan para que los clientes de la interfaz de usuario implementen cuadros de diálogo de confirmación, pero la API en sí:

- No valida el texto de confirmación
- No impone ningún requisito de confirmación
- Ejecuta la acción del botón inmediatamente cuando se llama

La confirmación es puramente una característica de seguridad de la interfaz de usuario del lado del cliente.

### Ejemplo: Haciendo Clic en Diferentes Tipos de Botones

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Las tres mutaciones anteriores ejecutarán la acción del botón inmediatamente cuando se llamen a través de la API, omitiendo cualquier requisito de confirmación.

## Campos de Respuesta

### Respuesta de Campo Personalizado

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el campo personalizado |
| `name` | String! | Nombre para mostrar del botón |
| `type` | CustomFieldType! | Siempre `BUTTON` para campos de botón |
| `buttonType` | String | Configuración del comportamiento de confirmación |
| `buttonConfirmText` | String | Texto de confirmación requerido (si se usa confirmación dura) |
| `description` | String | Texto de ayuda para los usuarios |
| `required` | Boolean! | Si el campo es requerido |
| `isActive` | Boolean! | Si el campo está actualmente activo |
| `projectId` | String! | ID del proyecto al que pertenece este campo |
| `createdAt` | DateTime! | Cuándo se creó el campo |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el campo |

## Cómo Funcionan los Campos de Botón

### Integración de Automatización

Los campos de botón están diseñados para trabajar con el sistema de automatización de Blue:

1. **Crea el campo de botón** utilizando la mutación anterior
2. **Configura automatizaciones** que escuchen eventos `CUSTOM_FIELD_BUTTON_CLICKED`
3. **Los usuarios hacen clic en el botón** en la interfaz de usuario
4. **Las automatizaciones ejecutan** las acciones configuradas

### Flujo de Eventos

Cuando se hace clic en un botón:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### No Almacenamiento de Datos

Importante: Los campos de botón no almacenan ningún valor de datos. Sirven puramente como disparadores de acción. Cada clic:
- Genera un evento
- Activa automatizaciones asociadas
- Registra una acción en el historial de tareas
- No modifica ningún valor de campo

## Permisos Requeridos

Los usuarios necesitan roles de proyecto apropiados para crear y usar campos de botón:

| Acción | Rol Requerido |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Respuestas de Error

### Permiso Denegado
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Campo Personalizado No Encontrado
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

**Nota**: La API no devuelve errores específicos por automatizaciones faltantes o desajustes de confirmación.

## Mejores Prácticas

### Convenciones de Nomenclatura
- Usa nombres orientados a la acción: "Enviar Factura", "Crear Informe", "Notificar al Equipo"
- Sé específico sobre lo que hace el botón
- Evita nombres genéricos como "Botón 1" o "Hacer Clic Aquí"

### Configuraciones de Confirmación
- Deja `buttonType` vacío para acciones seguras y reversibles
- Establece `buttonType` para sugerir comportamiento de confirmación a los clientes de la interfaz de usuario
- Usa `buttonConfirmText` para especificar lo que los usuarios deben escribir en las confirmaciones de la interfaz de usuario
- Recuerda: Estos son solo indicios para la interfaz de usuario - las llamadas a la API siempre se ejecutan inmediatamente

### Diseño de Automatización
- Mantén las acciones del botón enfocadas en un solo flujo de trabajo
- Proporciona retroalimentación clara sobre lo que sucedió después de hacer clic
- Considera agregar texto descriptivo para explicar el propósito del botón

## Casos de Uso Comunes

1. **Transiciones de Flujo de Trabajo**
   - "Marcar como Completo"
   - "Enviar para Aprobación"
   - "Archivar Tarea"

2. **Integraciones Externas**
   - "Sincronizar con CRM"
   - "Generar Factura"
   - "Enviar Actualización por Correo Electrónico"

3. **Operaciones por Lotes**
   - "Actualizar Todas las Subtareas"
   - "Copiar a Proyectos"
   - "Aplicar Plantilla"

4. **Acciones de Reporte**
   - "Generar Informe"
   - "Exportar Datos"
   - "Crear Resumen"

## Limitaciones

- Los botones no pueden almacenar ni mostrar valores de datos
- Cada botón solo puede activar automatizaciones, no llamadas directas a la API (sin embargo, las automatizaciones pueden incluir acciones de solicitud HTTP para llamar a APIs externas o a las propias APIs de Blue)
- La visibilidad del botón no puede ser controlada condicionalmente
- Máximo de una ejecución de automatización por clic (aunque esa automatización puede activar múltiples acciones)

## Recursos Relacionados

- [API de Automatizaciones](/api/automations/index) - Configura acciones activadas por botones
- [Descripción General de Campos Personalizados](/custom-fields/list-custom-fields) - Conceptos generales de campos personalizados