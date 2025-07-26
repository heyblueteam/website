---
title: Creando listas de verificación reutilizables usando automatizaciones
category: "Best Practices"
description: Aprenda cómo crear automatizaciones de gestión de proyectos para listas de verificación reutilizables.
date: 2024-07-08
---

En muchos proyectos y procesos, es posible que necesite usar la misma lista de verificación en múltiples registros o tareas.

Sin embargo, no es muy eficiente reescribir manualmente la lista de verificación cada vez que desea agregarla a un registro. ¡Aquí es donde puede aprovechar las [poderosas automatizaciones de gestión de proyectos](/platform/features/automations) para hacer esto automáticamente por usted!

Como recordatorio, las automatizaciones en Blue requieren dos elementos clave:

1. Un Disparador — Qué debe suceder para iniciar la automatización. Esto puede ser cuando un registro recibe una etiqueta específica, cuando se mueve a una posición específica
2. Una o más Acciones — En este caso, sería la creación automática de una o más listas de verificación.

Comencemos primero con la acción, luego discutiremos los posibles disparadores que puede usar.

## Acción de Automatización de Lista de Verificación

Puede crear una nueva automatización, y puede configurar una o más listas de verificación para que se creen, como en el ejemplo a continuación:

![](/insights/checklist-automation.png)

Estas serían las listas de verificación que desea que se creen cada vez que realiza la acción.

## Disparadores de Automatización de Lista de Verificación

Hay varias formas en que puede disparar la creación de sus listas de verificación reutilizables. Aquí hay algunas opciones populares:

- **Agregar una Etiqueta Específica:** Puede configurar la automatización para que se dispare cuando se agregue una etiqueta particular a un registro. Por ejemplo, cuando se agregue la etiqueta "Nuevo Proyecto", podría crear automáticamente su lista de verificación de iniciación de proyecto.
- **Asignación de Registro:** La creación de la lista de verificación puede dispararse cuando un registro se asigna a un individuo específico o a cualquier persona. Esto es útil para listas de verificación de incorporación o procedimientos específicos de tareas.
- **Mover a una Lista Específica:** Cuando un registro se mueve a una lista particular en su tablero de proyecto, puede disparar la creación de una lista de verificación relevante. Por ejemplo, mover un elemento a una lista de "Aseguramiento de Calidad" podría disparar una lista de verificación de QA.
- **Campo de Casilla de Verificación Personalizado:** Cree un campo de casilla de verificación personalizado y configure la automatización para que se dispare cuando esta casilla esté marcada. Esto le da control manual sobre cuándo agregar la lista de verificación.
- **Campos Personalizados de Selección Única o Múltiple:** Puede crear un campo personalizado de selección única o múltiple con varias opciones. Cada opción puede vincularse a una plantilla de lista de verificación específica a través de automatizaciones separadas. Esto permite un control más granular y la capacidad de tener múltiples plantillas de listas de verificación listas para diferentes escenarios.

Para mejorar el control sobre quién puede disparar estas automatizaciones, puede ocultar estos campos personalizados de ciertos usuarios usando roles de usuario personalizados. Esto asegura que solo los administradores de proyecto u otro personal autorizado pueda disparar estas opciones.

Recuerde, la clave para el uso efectivo de listas de verificación reutilizables con automatizaciones es diseñar sus disparadores cuidadosamente. Considere el flujo de trabajo de su equipo, los tipos de proyectos que maneja, y quién debería tener la capacidad de iniciar diferentes procesos. Con automatizaciones bien planificadas, puede agilizar significativamente su gestión de proyectos y asegurar consistencia en sus operaciones.

## Recursos Útiles

- [Documentación de Automatización de Gestión de Proyectos](https://documentation.blue.cc/automations)
- [Documentación de Roles de Usuario Personalizados](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Documentación de Campos Personalizados](https://documentation.blue.cc/custom-fields)