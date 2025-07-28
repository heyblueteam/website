---
title: Campo Personalizado de Correo Electrónico
description: Crea campos de correo electrónico para almacenar y validar direcciones de correo electrónico
---

Los campos personalizados de correo electrónico te permiten almacenar direcciones de correo electrónico en registros con validación incorporada. Son ideales para rastrear información de contacto, correos electrónicos de asignados o cualquier dato relacionado con correos electrónicos en tus proyectos.

## Ejemplo Básico

Crea un campo de correo electrónico simple:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de correo electrónico con descripción:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de correo electrónico |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `EMAIL` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

## Estableciendo Valores de Correo Electrónico

Para establecer o actualizar un valor de correo electrónico en un registro:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de correo electrónico |
| `text` | String | No | Dirección de correo electrónico a almacenar |

## Creando Registros con Valores de Correo Electrónico

Al crear un nuevo registro con valores de correo electrónico:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Campos de Respuesta

### Respuesta de CustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | ID! | Identificador único para el campo personalizado |
| `name` | String! | Nombre para mostrar del campo de correo electrónico |
| `type` | CustomFieldType! | El tipo de campo (EMAIL) |
| `description` | String | Texto de ayuda para el campo |
| `value` | JSON | Contiene el valor del correo electrónico (ver abajo) |
| `createdAt` | DateTime! | Cuándo se creó el campo |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el campo |

**Importante**: Los valores de correo electrónico se acceden a través del `customField.value.text` campo, no directamente en la respuesta.

## Consultando Valores de Correo Electrónico

Al consultar registros con campos personalizados de correo electrónico, accede al correo electrónico a través de la ruta `customField.value.text`:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

La respuesta incluirá el correo electrónico en la estructura anidada:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## Validación de Correo Electrónico

### Validación de Formularios
Cuando se utilizan campos de correo electrónico en formularios, validan automáticamente el formato del correo electrónico:
- Utiliza reglas estándar de validación de correo electrónico
- Elimina espacios en blanco de la entrada
- Rechaza formatos de correo electrónico inválidos

### Reglas de Validación
- Debe contener un símbolo `@`
- Debe tener un formato de dominio válido
- Los espacios en blanco al principio y al final se eliminan automáticamente
- Se aceptan formatos de correo electrónico comunes

### Ejemplos de Correos Electrónicos Válidos
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Ejemplos de Correos Electrónicos Inválidos
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Notas Importantes

### API Directa vs Formularios
- **Formularios**: Se aplica validación automática de correo electrónico
- **API Directa**: Sin validación - se puede almacenar cualquier texto
- **Recomendación**: Utiliza formularios para la entrada del usuario para asegurar la validación

### Formato de Almacenamiento
- Las direcciones de correo electrónico se almacenan como texto plano
- Sin formato o análisis especial
- Sensibilidad a mayúsculas: los campos personalizados de EMAIL se almacenan de manera sensible a mayúsculas (a diferencia de los correos electrónicos de autenticación de usuarios que se normalizan a minúsculas)
- Sin limitaciones de longitud máxima más allá de las restricciones de la base de datos (límite de 16 MB)

## Permisos Requeridos

| Acción | Permiso Requerido |
|--------|-------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Respuestas de Error

### Formato de Correo Electrónico Inválido (Solo Formularios)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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

### Entrada de Datos
- Siempre valida las direcciones de correo electrónico en tu aplicación
- Utiliza campos de correo electrónico solo para direcciones de correo electrónico reales
- Considera usar formularios para la entrada del usuario para obtener validación automática

### Calidad de Datos
- Elimina espacios en blanco antes de almacenar
- Considera la normalización de mayúsculas (típicamente minúsculas)
- Valida el formato del correo electrónico antes de operaciones importantes

### Consideraciones de Privacidad
- Las direcciones de correo electrónico se almacenan como texto plano
- Considera las regulaciones de privacidad de datos (GDPR, CCPA)
- Implementa controles de acceso apropiados

## Casos de Uso Comunes

1. **Gestión de Contactos**
   - Direcciones de correo electrónico de clientes
   - Información de contacto de proveedores
   - Correos electrónicos de miembros del equipo
   - Detalles de contacto de soporte

2. **Gestión de Proyectos**
   - Correos electrónicos de interesados
   - Correos electrónicos de contacto para aprobación
   - Destinatarios de notificaciones
   - Correos electrónicos de colaboradores externos

3. **Soporte al Cliente**
   - Direcciones de correo electrónico de clientes
   - Contactos de tickets de soporte
   - Contactos de escalación
   - Direcciones de correo electrónico de retroalimentación

4. **Ventas y Marketing**
   - Direcciones de correo electrónico de leads
   - Listas de contactos de campañas
   - Información de contacto de socios
   - Correos electrónicos de fuentes de referencia

## Características de Integración

### Con Automatizaciones
- Disparar acciones cuando se actualizan los campos de correo electrónico
- Enviar notificaciones a direcciones de correo electrónico almacenadas
- Crear tareas de seguimiento basadas en cambios de correo electrónico

### Con Búsquedas
- Referenciar datos de correo electrónico de otros registros
- Agregar listas de correos electrónicos de múltiples fuentes
- Encontrar registros por dirección de correo electrónico

### Con Formularios
- Validación automática de correo electrónico
- Verificación de formato de correo electrónico
- Eliminación de espacios en blanco

## Limitaciones

- Sin verificación o validación de correo electrónico incorporada más allá de la verificación de formato
- Sin características de UI específicas de correo electrónico (como enlaces de correo electrónico clicables)
- Almacenado como texto plano sin cifrado
- Sin capacidades de composición o envío de correo electrónico
- Sin almacenamiento de metadatos de correo electrónico (nombre para mostrar, etc.)
- Las llamadas a la API directa omiten la validación (solo los formularios validan)

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para datos de texto no relacionados con correos electrónicos
- [Campos de URL](/api/custom-fields/url) - Para direcciones de sitios web
- [Campos de Teléfono](/api/custom-fields/phone) - Para números de teléfono
- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales