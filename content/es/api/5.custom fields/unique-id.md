---
title: Campo Personalizado de ID Único
description: Crea campos de identificador único auto-generados con numeración secuencial y formato personalizado
---

Los campos personalizados de ID único generan automáticamente identificadores únicos y secuenciales para tus registros. Son perfectos para crear números de ticket, IDs de pedidos, números de factura o cualquier sistema de identificador secuencial en tu flujo de trabajo.

## Ejemplo Básico

Crea un campo de ID único simple con auto-secuenciación:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Ejemplo Avanzado

Crea un campo de ID único formateado con prefijo y relleno con ceros:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Parámetros de Entrada

### CreateCustomFieldInput (UNIQUE_ID)

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de ID único |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `UNIQUE_ID` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |
| `useSequenceUniqueId` | Boolean | No | Habilitar auto-secuenciación (predeterminado: falso) |
| `prefix` | String | No | Prefijo de texto para los IDs generados (por ejemplo, "TAREA-") |
| `sequenceDigits` | Int | No | Número de dígitos para el relleno con ceros |
| `sequenceStartingNumber` | Int | No | Número inicial para la secuencia |

## Opciones de Configuración

### Auto-Sequenciación (`useSequenceUniqueId`)
- **true**: Genera automáticamente IDs secuenciales cuando se crean registros
- **false** o **undefined**: Se requiere entrada manual (funciona como un campo de texto)

### Prefijo (`prefix`)
- Prefijo de texto opcional añadido a todos los IDs generados
- Ejemplos: "TAREA-", "ORD-", "ERROR-", "REQ-"
- Sin límite de longitud, pero mantenlo razonable para la visualización

### Dígitos de Secuencia (`sequenceDigits`)
- Número de dígitos para el relleno con ceros del número de secuencia
- Ejemplo: `sequenceDigits: 3` produce `001`, `002`, `003`
- Si no se especifica, no se aplica relleno

### Número Inicial (`sequenceStartingNumber`)
- El primer número en la secuencia
- Ejemplo: `sequenceStartingNumber: 1000` comienza en 1000, 1001, 1002...
- Si no se especifica, comienza en 1 (comportamiento predeterminado)

## Formato de ID Generado

El formato final del ID combina todas las opciones de configuración:

```
{prefix}{paddedSequenceNumber}
```

### Ejemplos de Formato

| Configuración | IDs Generados |
|---------------|---------------|
| Sin opciones | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Lectura de Valores de ID Único

### Consultar Registros con IDs Únicos
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Formato de Respuesta
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Generación Automática de ID

### Cuándo se Generan los IDs
- **Creación de Registro**: Los IDs se asignan automáticamente cuando se crean nuevos registros
- **Adición de Campo**: Al agregar un campo UNIQUE_ID a registros existentes, se encola un trabajo en segundo plano (implementación del trabajador pendiente)
- **Procesamiento en Segundo Plano**: La generación de ID para nuevos registros ocurre de manera sincrónica a través de disparadores de base de datos

### Proceso de Generación
1. **Disparador**: Se crea un nuevo registro o se agrega un campo UNIQUE_ID
2. **Búsqueda de Secuencia**: El sistema encuentra el siguiente número de secuencia disponible
3. **Asignación de ID**: Se asigna el número de secuencia al registro
4. **Actualización del Contador**: El contador de secuencia se incrementa para futuros registros
5. **Formateo**: El ID se formatea con prefijo y relleno cuando se muestra

### Garantías de Unicidad
- **Restricciones de Base de Datos**: Restricción única en los IDs de secuencia dentro de cada campo
- **Operaciones Atómicas**: La generación de secuencia utiliza bloqueos de base de datos para prevenir duplicados
- **Alcance de Proyecto**: Las secuencias son independientes por proyecto
- **Protección contra Condiciones de Carrera**: Las solicitudes concurrentes se manejan de manera segura

## Modo Manual vs Automático

### Modo Automático (`useSequenceUniqueId: true`)
- Los IDs se generan automáticamente a través de disparadores de base de datos
- Se garantiza la numeración secuencial
- La generación de secuencia atómica previene duplicados
- Los IDs formateados combinan prefijo + número de secuencia con relleno

### Modo Manual (`useSequenceUniqueId: false` o `undefined`)
- Funciona como un campo de texto regular
- Los usuarios pueden ingresar valores personalizados a través de `setTodoCustomField` con el parámetro `text`
- No hay generación automática
- No hay aplicación de unicidad más allá de las restricciones de la base de datos

## Estableciendo Valores Manuales (Solo Modo Manual)

Cuando `useSequenceUniqueId` es falso, puedes establecer valores manualmente:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Campos de Respuesta

### Respuesta TodoCustomField (UNIQUE_ID)

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `sequenceId` | Int | El número de secuencia generado (poblado para campos UNIQUE_ID) |
| `text` | String | El valor de texto formateado (combina prefijo + secuencia con relleno) |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se actualizó por última vez el valor |

### Respuesta CustomField (UNIQUE_ID)

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | Si la auto-secuenciación está habilitada |
| `prefix` | String | Prefijo de texto para los IDs generados |
| `sequenceDigits` | Int | Número de dígitos para el relleno con ceros |
| `sequenceStartingNumber` | Int | Número inicial para la secuencia |

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Respuestas de Error

### Error de Configuración de Campo
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Error de Permiso
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Notas Importantes

### IDs Auto-Generados
- **Solo Lectura**: Los IDs auto-generados no pueden ser editados manualmente
- **Permanente**: Una vez asignados, los IDs de secuencia no cambian
- **Cronológico**: Los IDs reflejan el orden de creación
- **Alcance**: Las secuencias son independientes por proyecto

### Consideraciones de Rendimiento
- La generación de ID para nuevos registros es sincrónica a través de disparadores de base de datos
- La generación de secuencia utiliza `FOR UPDATE` bloqueos para operaciones atómicas
- Existe un sistema de trabajos en segundo plano, pero la implementación del trabajador está pendiente
- Considera los números iniciales de secuencia para proyectos de alto volumen

### Migración y Actualizaciones
- Agregar auto-secuenciación a registros existentes encola un trabajo en segundo plano (trabajador pendiente)
- Cambiar la configuración de secuencia afecta solo a registros futuros
- Los IDs existentes permanecen sin cambios cuando se actualiza la configuración
- Los contadores de secuencia continúan desde el máximo actual

## Mejores Prácticas

### Diseño de Configuración
- Elige prefijos descriptivos que no entren en conflicto con otros sistemas
- Usa un relleno de dígitos apropiado para tu volumen esperado
- Establece números iniciales razonables para evitar conflictos
- Prueba la configuración con datos de muestra antes del despliegue

### Directrices de Prefijo
- Mantén los prefijos cortos y memorables (2-5 caracteres)
- Usa mayúsculas para consistencia
- Incluye separadores (guiones, guiones bajos) para legibilidad
- Evita caracteres especiales que puedan causar problemas en URLs o sistemas

### Planificación de Secuencia
- Estima tu volumen de registros para elegir un relleno de dígitos apropiado
- Considera el crecimiento futuro al establecer números iniciales
- Planifica diferentes rangos de secuencia para diferentes tipos de registros
- Documenta tus esquemas de ID para referencia del equipo

## Casos de Uso Comunes

1. **Sistemas de Soporte**
   - Números de ticket: `TICK-001`, `TICK-002`
   - IDs de casos: `CASE-2024-001`
   - Solicitudes de soporte: `SUP-001`

2. **Gestión de Proyectos**
   - IDs de tareas: `TASK-001`, `TASK-002`
   - Elementos de sprint: `SPRINT-001`
   - Números de entregables: `DEL-001`

3. **Operaciones Comerciales**
   - Números de pedido: `ORD-2024-001`
   - IDs de factura: `INV-001`
   - Órdenes de compra: `PO-001`

4. **Gestión de Calidad**
   - Informes de errores: `BUG-001`
   - IDs de casos de prueba: `TEST-001`
   - Números de revisión: `REV-001`

## Características de Integración

### Con Automatizaciones
- Disparar acciones cuando se asignan IDs únicos
- Usar patrones de ID en reglas de automatización
- Referenciar IDs en plantillas de correo electrónico y notificaciones

### Con Búsquedas
- Referenciar IDs únicos de otros registros
- Encontrar registros por ID único
- Mostrar identificadores de registros relacionados

### Con Informes
- Agrupar y filtrar por patrones de ID
- Rastrear tendencias de asignación de ID
- Monitorear el uso de secuencias y huecos

## Limitaciones

- **Solo Secuencial**: Los IDs se asignan en orden cronológico
- **Sin Huecos**: Los registros eliminados dejan huecos en las secuencias
- **Sin Reutilización**: Los números de secuencia nunca se reutilizan
- **Alcance de Proyecto**: No se pueden compartir secuencias entre proyectos
- **Restricciones de Formato**: Opciones de formato limitadas
- **Sin Actualizaciones Masivas**: No se pueden actualizar masivamente los IDs de secuencia existentes
- **Sin Lógica Personalizada**: No se pueden implementar reglas personalizadas de generación de ID

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para identificadores de texto manuales
- [Campos Numéricos](/api/custom-fields/number) - Para secuencias numéricas
- [Descripción General de Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceptos generales
- [Automatizaciones](/api/automations) - Para reglas de automatización basadas en ID