---
title: Campo Personalizado de URL
description: Crea campos de URL para almacenar direcciones de sitios web y enlaces
---

Los campos personalizados de URL te permiten almacenar direcciones de sitios web y enlaces en tus registros. Son ideales para rastrear sitios web de proyectos, enlaces de referencia, URLs de documentación o cualquier recurso basado en la web relacionado con tu trabajo.

## Ejemplo Básico

Crea un campo de URL simple:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de URL con descripción:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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
| `name` | String! | ✅ Sí | Nombre a mostrar del campo de URL |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `URL` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

**Nota:** El `projectId` se pasa como un argumento separado a la mutación, no como parte del objeto de entrada.

## Estableciendo Valores de URL

Para establecer o actualizar un valor de URL en un registro:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de URL |
| `text` | String! | ✅ Sí | Dirección URL a almacenar |

## Creando Registros con Valores de URL

Al crear un nuevo registro con valores de URL:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
| `text` | String | La dirección URL almacenada |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

## Validación de URL

### Implementación Actual
- **API Directa**: No se aplica actualmente validación de formato de URL
- **Formularios**: La validación de URL está planificada pero no está activa actualmente
- **Almacenamiento**: Cualquier valor de cadena puede almacenarse en campos de URL

### Validación Planificada
Las versiones futuras incluirán:
- Validación de protocolo HTTP/HTTPS
- Verificación de formato de URL válido
- Validación de nombre de dominio
- Adición automática de prefijo de protocolo

### Formatos de URL Recomendados
Aunque no se aplica actualmente, utiliza estos formatos estándar:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Notas Importantes

### Formato de Almacenamiento
- Las URL se almacenan como texto sin formato sin modificación
- No se añade automáticamente el protocolo (http://, https://)
- La sensibilidad a mayúsculas se conserva tal como se ingresó
- No se realiza codificación/decodificación de URL

### API Directa vs Formularios
- **Formularios**: Validación de URL planificada (no activa actualmente)
- **API Directa**: Sin validación - se puede almacenar cualquier texto
- **Recomendación**: Valida las URL en tu aplicación antes de almacenarlas

### Campos de URL vs Texto
- **URL**: Semánticamente destinado a direcciones web
- **TEXT_SINGLE**: Texto general de una sola línea
- **Backend**: Almacenamiento y validación actualmente idénticos
- **Frontend**: Diferentes componentes de UI para la entrada de datos

## Permisos Requeridos

Las operaciones de campo personalizado utilizan permisos basados en roles:

| Acción | Rol Requerido |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Nota:** Los permisos se verifican en función de los roles de usuario en el proyecto, no en constantes de permiso específicas.

## Respuestas de Error

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

## Mejores Prácticas

### Estándares de Formato de URL
- Siempre incluye el protocolo (http:// o https://)
- Usa HTTPS cuando sea posible por seguridad
- Prueba las URL antes de almacenarlas para asegurarte de que sean accesibles
- Considera usar URL acortadas para fines de visualización

### Calidad de Datos
- Valida las URL en tu aplicación antes de almacenarlas
- Verifica errores comunes (protocolos faltantes, dominios incorrectos)
- Estandariza los formatos de URL en toda tu organización
- Considera la accesibilidad y disponibilidad de las URL

### Consideraciones de Seguridad
- Ten cuidado con las URL proporcionadas por los usuarios
- Valida los dominios si restringes a sitios específicos
- Considera el escaneo de URL para contenido malicioso
- Usa URL HTTPS al manejar datos sensibles

## Filtrado y Búsqueda

### Búsqueda Contiene
Los campos de URL admiten la búsqueda de subcadenas:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Capacidades de Búsqueda
- Coincidencia de subcadenas sin distinción entre mayúsculas y minúsculas
- Coincidencia parcial de dominios
- Búsqueda de rutas y parámetros
- Sin filtrado específico de protocolo

## Casos de Uso Comunes

1. **Gestión de Proyectos**
   - Sitios web de proyectos
   - Enlaces de documentación
   - URLs de repositorios
   - Sitios de demostración

2. **Gestión de Contenidos**
   - Materiales de referencia
   - Enlaces de origen
   - Recursos multimedia
   - Artículos externos

3. **Soporte al Cliente**
   - Sitios web de clientes
   - Documentación de soporte
   - Artículos de la base de conocimientos
   - Tutoriales en video

4. **Ventas y Marketing**
   - Sitios web de empresas
   - Páginas de productos
   - Materiales de marketing
   - Perfiles de redes sociales

## Características de Integración

### Con Búsquedas
- Referenciar URL de otros registros
- Encontrar registros por dominio o patrón de URL
- Mostrar recursos web relacionados
- Agregar enlaces de múltiples fuentes

### Con Formularios
- Componentes de entrada específicos para URL
- Validación planificada para el formato adecuado de URL
- Capacidades de vista previa de enlaces (frontend)
- Visualización de URL clicables

### Con Reportes
- Rastrear el uso y patrones de URL
- Monitorear enlaces rotos o inaccesibles
- Categorizar por dominio o protocolo
- Exportar listas de URL para análisis

## Limitaciones

### Limitaciones Actuales
- Sin validación activa de formato de URL
- Sin adición automática de protocolo
- Sin verificación de enlaces o comprobación de accesibilidad
- Sin acortamiento o expansión de URL
- Sin generación de favicon o vista previa

### Restricciones de Automatización
- No disponible como campos de activación de automatización
- No se puede usar en actualizaciones de campos de automatización
- Puede referenciarse en condiciones de automatización
- Disponible en plantillas de correo electrónico y webhooks

### Restricciones Generales
- Sin funcionalidad de vista previa de enlace incorporada
- Sin acortamiento automático de URL
- Sin seguimiento de clics o análisis
- Sin comprobación de expiración de URL
- Sin escaneo de URL maliciosas

## Mejoras Futuras

### Características Planificadas
- Validación de protocolo HTTP/HTTPS
- Patrones de validación regex personalizados
- Adición automática de prefijo de protocolo
- Comprobación de accesibilidad de URL

### Mejoras Potenciales
- Generación de vista previa de enlaces
- Visualización de favicon
- Integración de acortamiento de URL
- Capacidades de seguimiento de clics
- Detección de enlaces rotos

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para datos de texto no URL
- [Campos de Correo Electrónico](/api/custom-fields/email) - Para direcciones de correo electrónico
- [Descripción General de Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceptos generales

## Migración de Campos de Texto

Si estás migrando de campos de texto a campos de URL:

1. **Crea un campo de URL** con el mismo nombre y configuración
2. **Exporta los valores de texto existentes** para verificar que sean URLs válidas
3. **Actualiza los registros** para usar el nuevo campo de URL
4. **Elimina el antiguo campo de texto** después de una migración exitosa
5. **Actualiza las aplicaciones** para usar componentes de UI específicos de URL

### Ejemplo de Migración
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```