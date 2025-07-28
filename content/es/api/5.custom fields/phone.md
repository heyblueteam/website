---
title: Campo Personalizado de Teléfono
description: Crea campos de teléfono para almacenar y validar números de teléfono con formato internacional
---

Los campos personalizados de teléfono te permiten almacenar números de teléfono en registros con validación incorporada y formato internacional. Son ideales para rastrear información de contacto, contactos de emergencia o cualquier dato relacionado con teléfonos en tus proyectos.

## Ejemplo Básico

Crea un campo de teléfono simple:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de teléfono con descripción:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ Sí | Nombre de visualización del campo de teléfono |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `PHONE` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota**: Los campos personalizados se asocian automáticamente con el proyecto en función del contexto del proyecto actual del usuario. No se requiere el parámetro `projectId`.

## Estableciendo Valores de Teléfono

Para establecer o actualizar un valor de teléfono en un registro:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de teléfono |
| `text` | String | No | Número de teléfono con código de país |
| `regionCode` | String | No | Código de país (detectado automáticamente) |

**Nota**: Aunque `text` es opcional en el esquema, se requiere un número de teléfono para que el campo tenga sentido. Al usar `setTodoCustomField`, no se realiza ninguna validación: puedes almacenar cualquier valor de texto y regionCode. La detección automática solo ocurre durante la creación del registro.

## Creando Registros con Valores de Teléfono

Al crear un nuevo registro con valores de teléfono:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
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
| `text` | String | El número de teléfono formateado (formato internacional) |
| `regionCode` | String | El código de país (por ejemplo, "US", "GB", "CA") |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

## Validación de Números de Teléfono

**Importante**: La validación y el formato de números de teléfono solo ocurren al crear nuevos registros a través de `createTodo`. Al actualizar valores de teléfono existentes usando `setTodoCustomField`, no se realiza ninguna validación y los valores se almacenan tal como se proporcionan.

### Formatos Aceptados (Durante la Creación del Registro)
Los números de teléfono deben incluir un código de país en uno de estos formatos:

- **Formato E.164 (preferido)**: `+12345678900`
- **Formato internacional**: `+1 234 567 8900`
- **Internacional con puntuación**: `+1 (234) 567-8900`
- **Código de país con guiones**: `+1-234-567-8900`

**Nota**: Los formatos nacionales sin código de país (como `(234) 567-8900`) serán rechazados durante la creación del registro.

### Reglas de Validación (Durante la Creación del Registro)
- Utiliza libphonenumber-js para el análisis y la validación
- Acepta varios formatos de números de teléfono internacionales
- Detecta automáticamente el país a partir del número
- Formatea el número en formato de visualización internacional (por ejemplo, `+1 234 567 8900`)
- Extrae y almacena el código de país por separado (por ejemplo, `US`)

### Ejemplos de Teléfonos Válidos
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Ejemplos de Teléfonos Inválidos
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Formato de Almacenamiento

Al crear registros con números de teléfono:
- **texto**: Almacenado en formato internacional (por ejemplo, `+1 234 567 8900`) después de la validación
- **regionCode**: Almacenado como código de país ISO (por ejemplo, `US`, `GB`, `CA`) detectado automáticamente

Al actualizar a través de `setTodoCustomField`:
- **texto**: Almacenado exactamente como se proporciona (sin formato)
- **regionCode**: Almacenado exactamente como se proporciona (sin validación)

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Respuestas de Error

### Formato de Teléfono Inválido
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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

### Código de País Faltante
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Mejores Prácticas

### Entrada de Datos
- Siempre incluye el código de país en los números de teléfono
- Usa el formato E.164 para consistencia
- Valida los números antes de almacenarlos para operaciones importantes
- Considera las preferencias regionales para el formato de visualización

### Calidad de Datos
- Almacena números en formato internacional para compatibilidad global
- Usa regionCode para características específicas del país
- Valida los números de teléfono antes de operaciones críticas (SMS, llamadas)
- Considera las implicaciones de la zona horaria para el tiempo de contacto

### Consideraciones Internacionales
- El código de país se detecta y almacena automáticamente
- Los números se formatean en estándar internacional
- Las preferencias de visualización regional pueden usar regionCode
- Considera las convenciones de marcado local al mostrar

## Casos de Uso Comunes

1. **Gestión de Contactos**
   - Números de teléfono de clientes
   - Información de contacto de proveedores
   - Números de teléfono de miembros del equipo
   - Detalles de contacto de soporte

2. **Contactos de Emergencia**
   - Números de contacto de emergencia
   - Información de contacto de guardia
   - Contactos para respuesta a crisis
   - Números de teléfono de escalación

3. **Soporte al Cliente**
   - Números de teléfono de clientes
   - Números de devolución de llamada de soporte
   - Números de teléfono de verificación
   - Números de contacto para seguimiento

4. **Ventas y Marketing**
   - Números de teléfono de prospectos
   - Listas de contacto de campañas
   - Información de contacto de socios
   - Teléfonos de fuentes de referencia

## Características de Integración

### Con Automatizaciones
- Disparar acciones cuando se actualizan los campos de teléfono
- Enviar notificaciones SMS a números de teléfono almacenados
- Crear tareas de seguimiento basadas en cambios de teléfono
- Dirigir llamadas según los datos del número de teléfono

### Con Consultas
- Referenciar datos de teléfono de otros registros
- Agregar listas de teléfonos de múltiples fuentes
- Encontrar registros por número de teléfono
- Hacer referencia cruzada a información de contacto

### Con Formularios
- Validación automática de teléfonos
- Verificación de formato internacional
- Detección de código de país
- Retroalimentación de formato en tiempo real

## Limitaciones

- Requiere código de país para todos los números
- No tiene capacidades integradas de SMS o llamadas
- No hay verificación de número de teléfono más allá de la verificación de formato
- No se almacena metadatos de teléfono (operador, tipo, etc.)
- Los números en formato nacional sin código de país son rechazados
- No hay formateo automático de números de teléfono en la interfaz más allá del estándar internacional

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para datos de texto que no sean teléfonos
- [Campos de Correo Electrónico](/api/custom-fields/email) - Para direcciones de correo electrónico
- [Campos de URL](/api/custom-fields/url) - Para direcciones de sitios web
- [Descripción General de Campos Personalizados](/custom-fields/list-custom-fields) - Conceptos generales