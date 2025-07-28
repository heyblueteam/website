---
title: Campo Personalizado de Ubicación
description: Crea campos de ubicación para almacenar coordenadas geográficas para registros
---

Los campos personalizados de ubicación almacenan coordenadas geográficas (latitud y longitud) para registros. Soportan almacenamiento preciso de coordenadas, consultas geoespaciales y filtrado eficiente basado en la ubicación.

## Ejemplo Básico

Crea un campo de ubicación simple:

```graphql
mutation CreateLocationField {
  createCustomField(input: {
    name: "Meeting Location"
    type: LOCATION
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Ejemplo Avanzado

Crea un campo de ubicación con descripción:

```graphql
mutation CreateDetailedLocationField {
  createCustomField(input: {
    name: "Office Location"
    type: LOCATION
    projectId: "proj_123"
    description: "Primary office location coordinates"
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
| `name` | String! | ✅ Sí | Nombre para mostrar del campo de ubicación |
| `type` | CustomFieldType! | ✅ Sí | Debe ser `LOCATION` |
| `description` | String | No | Texto de ayuda mostrado a los usuarios |

## Estableciendo Valores de Ubicación

Los campos de ubicación almacenan coordenadas de latitud y longitud:

```graphql
mutation SetLocationValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    latitude: 40.7128
    longitude: -74.0060
  })
}
```

### Parámetros de SetTodoCustomFieldInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sí | ID del registro a actualizar |
| `customFieldId` | String! | ✅ Sí | ID del campo personalizado de ubicación |
| `latitude` | Float | No | Coordenada de latitud (-90 a 90) |
| `longitude` | Float | No | Coordenada de longitud (-180 a 180) |

**Nota**: Aunque ambos parámetros son opcionales en el esquema, se requieren ambas coordenadas para una ubicación válida. Si solo se proporciona una, la ubicación será inválida.

## Validación de Coordenadas

### Rangos Válidos

| Coordenada | Rango | Descripción |
|------------|-------|-------------|
| Latitude | -90 to 90 | Posición Norte/Sur |
| Longitude | -180 to 180 | Posición Este/Oeste |

### Ejemplo de Coordenadas

| Ubicación | Latitud | Longitud |
|----------|----------|-----------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## Creando Registros con Valores de Ubicación

Al crear un nuevo registro con datos de ubicación:

```graphql
mutation CreateRecordWithLocation {
  createTodo(input: {
    title: "Site Visit"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "location_field_id"
      value: "40.7128,-74.0060"  # Format: "latitude,longitude"
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
      latitude
      longitude
    }
  }
}
```

### Formato de Entrada para Creación

Al crear registros, los valores de ubicación utilizan un formato separado por comas:

| Formato | Ejemplo | Descripción |
|--------|---------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | Formato de coordenada estándar |
| `"51.5074,-0.1278"` | London coordinates | Sin espacios alrededor de la coma |
| `"-33.8688,151.2093"` | Sydney coordinates | Se permiten valores negativos |

## Campos de Respuesta

### Respuesta de TodoCustomField

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el valor del campo |
| `customField` | CustomField! | La definición del campo personalizado |
| `latitude` | Float | Coordenada de latitud |
| `longitude` | Float | Coordenada de longitud |
| `todo` | Todo! | El registro al que pertenece este valor |
| `createdAt` | DateTime! | Cuándo se creó el valor |
| `updatedAt` | DateTime! | Cuándo se modificó por última vez el valor |

## Limitaciones Importantes

### Sin Geocodificación Incorporada

Los campos de ubicación solo almacenan coordenadas - **no** incluyen:
- Conversión de dirección a coordenadas
- Geocodificación inversa (coordenadas a dirección)
- Validación o búsqueda de direcciones
- Integración con servicios de mapas
- Búsqueda de nombres de lugares

### Servicios Externos Requeridos

Para la funcionalidad de dirección, necesitarás integrar servicios externos:
- **Google Maps API** para geocodificación
- **OpenStreetMap Nominatim** para geocodificación gratuita
- **MapBox** para mapeo y geocodificación
- **Here API** para servicios de ubicación

### Ejemplo de Integración

```javascript
// Client-side geocoding example (not part of Blue API)
async function geocodeAddress(address) {
  const response = await fetch(
    `https://maps.googleapis.com/maps/api/geocode/json?address=${encodeURIComponent(address)}&key=${API_KEY}`
  );
  const data = await response.json();
  
  if (data.results.length > 0) {
    const { lat, lng } = data.results[0].geometry.location;
    
    // Now set the location field in Blue
    await setTodoCustomField({
      todoId: "todo_123",
      customFieldId: "location_field_456",
      latitude: lat,
      longitude: lng
    });
  }
}
```

## Permisos Requeridos

| Acción | Rol Requerido |
|--------|---------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## Respuestas de Error

### Coordenadas Inválidas
```json
{
  "errors": [{
    "message": "Invalid coordinates: latitude must be between -90 and 90",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Longitud Inválida
```json
{
  "errors": [{
    "message": "Invalid coordinates: longitude must be between -180 and 180",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Mejores Prácticas

### Recolección de Datos
- Usa coordenadas GPS para ubicaciones precisas
- Valida las coordenadas antes de almacenarlas
- Considera las necesidades de precisión de las coordenadas (6 decimales ≈ 10cm de precisión)
- Almacena coordenadas en grados decimales (no en grados/minutos/segundos)

### Experiencia del Usuario
- Proporciona interfaces de mapa para la selección de coordenadas
- Muestra vistas previas de ubicación al mostrar coordenadas
- Valida las coordenadas del lado del cliente antes de las llamadas a la API
- Considera las implicaciones de la zona horaria para los datos de ubicación

### Rendimiento
- Usa índices espaciales para consultas eficientes
- Limita la precisión de las coordenadas a la precisión necesaria
- Considera el almacenamiento en caché para ubicaciones de acceso frecuente
- Agrupa actualizaciones de ubicación cuando sea posible

## Casos de Uso Comunes

1. **Operaciones de Campo**
   - Ubicaciones de equipos
   - Direcciones de llamadas de servicio
   - Sitios de inspección
   - Ubicaciones de entrega

2. **Gestión de Eventos**
   - Lugares de eventos
   - Ubicaciones de reuniones
   - Sitios de conferencias
   - Ubicaciones de talleres

3. **Seguimiento de Activos**
   - Posiciones de equipos
   - Ubicaciones de instalaciones
   - Seguimiento de vehículos
   - Ubicaciones de inventario

4. **Análisis Geográfico**
   - Áreas de cobertura de servicio
   - Distribución de clientes
   - Análisis de mercado
   - Gestión de territorios

## Características de Integración

### Con Búsquedas
- Referencia datos de ubicación de otros registros
- Encuentra registros por proximidad geográfica
- Agrega datos basados en la ubicación
- Referencia cruzada de coordenadas

### Con Automatizaciones
- Dispara acciones basadas en cambios de ubicación
- Crea notificaciones geocercadas
- Actualiza registros relacionados cuando cambian las ubicaciones
- Genera informes basados en la ubicación

### Con Fórmulas
- Calcula distancias entre ubicaciones
- Determina centros geográficos
- Analiza patrones de ubicación
- Crea métricas basadas en la ubicación

## Limitaciones

- Sin geocodificación incorporada o conversión de direcciones
- No se proporciona interfaz de mapeo
- Requiere servicios externos para funcionalidad de dirección
- Limitado solo al almacenamiento de coordenadas
- Sin validación automática de ubicación más allá de la verificación de rango

## Recursos Relacionados

- [Descripción General de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceptos generales
- [Google Maps API](https://developers.google.com/maps) - Servicios de geocodificación
- [OpenStreetMap Nominatim](https://nominatim.org/) - Geocodificación gratuita
- [MapBox API](https://docs.mapbox.com/) - Servicios de mapeo y geocodificación