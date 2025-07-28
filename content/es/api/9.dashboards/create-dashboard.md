---
title: Crear Tablero
description: Crea un nuevo tablero para visualización de datos e informes en Blue
---

## Crear un Tablero

La mutación `createDashboard` te permite crear un nuevo tablero dentro de tu empresa o proyecto. Los tableros son herramientas de visualización poderosas que ayudan a los equipos a rastrear métricas, monitorear el progreso y tomar decisiones basadas en datos.

### Ejemplo Básico

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### Tablero Específico del Proyecto

Crea un tablero asociado con un proyecto específico:

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## Parámetros de Entrada

### CreateDashboardInput

| Parámetro | Tipo | Requerido | Descripción |
|-----------|------|----------|-------------|
| `companyId` | String! | ✅ Sí | El ID de la empresa donde se creará el tablero |
| `title` | String! | ✅ Sí | El nombre del tablero. Debe ser una cadena no vacía |
| `projectId` | String | No | ID opcional de un proyecto para asociar con este tablero |

## Campos de Respuesta

La mutación devuelve un objeto completo `Dashboard`:

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para el tablero creado |
| `title` | String! | El título del tablero tal como se proporcionó |
| `companyId` | String! | La empresa a la que pertenece este tablero |
| `projectId` | String | El ID del proyecto asociado (si se proporcionó) |
| `project` | Project | El objeto del proyecto asociado (si se proporcionó projectId) |
| `createdBy` | User! | El usuario que creó el tablero (tú) |
| `dashboardUsers` | [DashboardUser!]! | Lista de usuarios con acceso (inicialmente solo el creador) |
| `createdAt` | DateTime! | Marca de tiempo de cuándo se creó el tablero |
| `updatedAt` | DateTime! | Marca de tiempo de la última modificación (igual que createdAt para nuevos tableros) |

### Campos de DashboardUser

Cuando se crea un tablero, el creador se agrega automáticamente como un usuario del tablero:

| Campo | Tipo | Descripción |
|-------|------|-------------|
| `id` | String! | Identificador único para la relación del usuario del tablero |
| `user` | User! | El objeto del usuario con acceso al tablero |
| `role` | DashboardRole! | El rol del usuario (el creador obtiene acceso completo) |
| `dashboard` | Dashboard! | Referencia de vuelta al tablero |

## Permisos Requeridos

Cualquier usuario autenticado que pertenezca a la empresa especificada puede crear tableros. No hay requisitos de rol especiales.

| Estado del Usuario | Puede Crear Tablero |
|--------------------|---------------------|
| Company Member | ✅ Sí |
| No Miembro de la Empresa | ❌ No |
| Unauthenticated | ❌ No |

## Respuestas de Error

### Empresa Inválida
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Usuario No en la Empresa
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Proyecto Inválido
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Título Vacío
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Notas Importantes

- **Propiedad automática**: El usuario que crea el tablero se convierte automáticamente en su propietario con permisos completos
- **Asociación de proyecto**: Si proporcionas un `projectId`, debe pertenecer a la misma empresa
- **Permisos iniciales**: Solo el creador tiene acceso inicialmente. Usa `editDashboard` para agregar más usuarios
- **Requisitos de título**: Los títulos de los tableros deben ser cadenas no vacías. No hay requisito de unicidad
- **Membresía de empresa**: Debes ser miembro de la empresa para crear tableros en ella

## Flujo de Trabajo de Creación de Tableros

1. **Crea el tablero** usando esta mutación
2. **Configura gráficos y widgets** usando la interfaz de construcción de tableros
3. **Agrega miembros del equipo** usando la mutación `editDashboard` con `dashboardUsers`
4. **Configura filtros y rangos de fechas** a través de la interfaz del tablero
5. **Comparte o incrusta** el tablero usando su ID único

## Casos de Uso

1. **Tableros ejecutivos**: Crea resúmenes de alto nivel de las métricas de la empresa
2. **Seguimiento de proyectos**: Construye tableros específicos de proyectos para monitorear el progreso
3. **Rendimiento del equipo**: Rastrear la productividad del equipo y métricas de logros
4. **Informes para clientes**: Crea tableros para informes dirigidos a clientes
5. **Monitoreo en tiempo real**: Configura tableros para datos operativos en vivo

## Mejores Prácticas

1. **Convenciones de nomenclatura**: Usa títulos claros y descriptivos que indiquen el propósito del tablero
2. **Asociación de proyectos**: Vincula tableros a proyectos cuando sean específicos de un proyecto
3. **Gestión de acceso**: Agrega miembros del equipo inmediatamente después de la creación para colaborar
4. **Organización**: Crea una jerarquía de tableros usando patrones de nomenclatura consistentes

## Operaciones Relacionadas

- [Listar Tableros](/api/dashboards/) - Recupera todos los tableros para una empresa o proyecto
- [Editar Tablero](/api/dashboards/rename-dashboard) - Cambiar el nombre del tablero o gestionar usuarios
- [Copiar Tablero](/api/dashboards/copy-dashboard) - Duplicar un tablero existente
- [Eliminar Tablero](/api/dashboards/delete-dashboard) - Eliminar un tablero