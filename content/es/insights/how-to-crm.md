---
title: Cómo configurar Blue como un CRM
description: Aprende a configurar Blue para rastrear a tus clientes y oportunidades de manera sencilla.
category: "Best Practices"
date: 2024-08-11
---


## Introducción

Una de las principales ventajas de usar Blue no es utilizarlo para un caso de uso *específico*, sino usarlo *en* múltiples casos de uso. Esto significa que no tienes que pagar por múltiples herramientas y también tienes un lugar donde puedes cambiar fácilmente entre tus diversos proyectos y procesos, como contratación, ventas, marketing y más.

Al ayudar a miles de clientes a configurarse en Blue a lo largo de los años, hemos notado que la parte difícil no es *configurar* Blue en sí, sino pensar en los procesos y aprovechar al máximo nuestra plataforma.

Las partes clave son pensar en el flujo de trabajo paso a paso para cada uno de tus procesos comerciales que deseas rastrear, así como los detalles específicos de los datos que deseas capturar y cómo esto se traduce en los campos personalizados que configuras.

Hoy, te guiaremos a través de la creación de [un sistema de CRM de ventas fácil de usar pero poderoso](/solutions/use-case/sales-crm) con una base de datos de clientes que está vinculada a un pipeline de oportunidades. Todos estos datos fluirán hacia un panel donde podrás ver datos en tiempo real sobre tus ventas totales, ventas pronosticadas y más.

## Base de Datos de Clientes

Lo primero que debes hacer es configurar un nuevo proyecto para almacenar los datos de tus clientes. Estos datos se referenciarán en otro proyecto donde rastreas oportunidades de ventas específicas.

La razón por la que separamos tu información de cliente de las oportunidades es que no se mapean uno a uno.

Un cliente puede tener múltiples oportunidades o proyectos.

Por ejemplo, si eres una agencia de marketing y diseño, puedes inicialmente involucrarte con un cliente para su branding, y luego hacer un proyecto separado para su sitio web, y luego otro para la gestión de sus redes sociales.

Todas estas serían oportunidades de ventas separadas que requieren su propio seguimiento y propuestas, pero todas están vinculadas a ese único cliente.

La ventaja de separar tu base de datos de clientes en un proyecto separado es que si actualizas algún detalle en tu base de datos de clientes, todas tus oportunidades tendrán automáticamente los nuevos datos, lo que significa que ahora tienes una única fuente de verdad en tu negocio. ¡No tienes que volver a editar todo manualmente!

Así que, lo primero que debes decidir es si vas a centrarte en la empresa o en la persona.

Esta decisión realmente depende de lo que estás vendiendo y a quién se lo vendes. Si vendes principalmente a empresas, entonces probablemente querrás que el nombre del registro sea el nombre de la empresa. Sin embargo, si vendes principalmente a individuos (es decir, eres un entrenador personal de salud o un consultor de marca personal), entonces lo más probable es que adoptes un enfoque centrado en la persona.

Así que el campo del nombre del registro será el nombre de la empresa o el nombre de la persona, dependiendo de tu elección. La razón de esto es que significa que puedes identificar fácilmente a un cliente de un vistazo, solo mirando tu tablero o base de datos.

A continuación, necesitas considerar qué información deseas capturar como parte de tu base de datos de clientes. Estos se convertirán en tus campos personalizados.

Los sospechosos habituales aquí son:

- Correo electrónico
- Número de teléfono
- Sitio web
- Dirección
- Fuente (es decir, ¿de dónde provino este cliente por primera vez?)
- Categoría

En Blue, también puedes eliminar cualquier campo predeterminado que no necesites. Para esta base de datos de clientes, normalmente recomendamos que elimines la fecha de vencimiento, el asignado, las dependencias y las listas de verificación. Puede que desees mantener nuestro campo de descripción predeterminado disponible en caso de que tengas notas generales sobre ese cliente que no sean específicas de ninguna oportunidad de venta.

Recomendamos que mantengas el campo "Referencia por", ya que esto será útil más adelante. Una vez que configuremos nuestra base de datos de oportunidades, podremos ver cada registro de ventas que está vinculado a este cliente en particular aquí.

En términos de listas, normalmente vemos que nuestros clientes simplemente lo mantienen simple y tienen una lista llamada "Clientes" y lo dejan así. Es mejor usar etiquetas o campos personalizados para la categorización.

Lo que es genial aquí es que una vez que tengas esto configurado, puedes importar fácilmente tus datos de otros sistemas o hojas de Excel a Blue a través de nuestra función de importación CSV, y también puedes crear un formulario para que nuevos clientes potenciales envíen sus detalles para que puedas **capturarlos automáticamente** en tu base de datos.

## Base de Datos de Oportunidades

Ahora que tenemos nuestra base de datos de clientes, necesitamos crear otro proyecto para capturar nuestras oportunidades de ventas reales. Puedes llamar a este proyecto "CRM de Ventas" u "Oportunidades".

### Listas como Pasos del Proceso

Para configurar tu proceso de ventas, necesitas pensar en cuáles son los pasos habituales que sigue una oportunidad desde el momento en que recibes una solicitud de un cliente hasta obtener un contrato firmado.

Cada lista en tu proyecto será un paso en tu proceso.

Independientemente de tu proceso específico, habrá algunas listas comunes que TODOS los CRM de ventas deberían tener:

- No Calificado — Todas las solicitudes entrantes, donde aún no has calificado a un cliente.
- Cerrado Ganado — Todas las oportunidades que ganaste y convertiste en ventas.
- Cerrado Perdido — Todas las oportunidades donde cotizaste a un cliente y no aceptó.
- N/A — Aquí colocas todas las oportunidades que no ganaste, pero que tampoco fueron "perdidas". Podría ser las que rechazaste, las que el cliente, por cualquier razón, te ignoró, y así sucesivamente.

En términos de pensar en tu proceso comercial de CRM de ventas, debes considerar el nivel de granularidad que deseas. No recomendamos tener 20 o 30 columnas, esto típicamente se vuelve confuso y te impide ver el panorama general.

Sin embargo, también es importante no hacer cada proceso demasiado amplio, ya que de lo contrario los tratos se quedarán "atascados" en una etapa específica durante semanas o meses, incluso cuando en realidad están avanzando. Aquí hay un enfoque típico recomendado:

- **No Calificado**: Todas las solicitudes entrantes, donde aún no has calificado a un cliente.
- **Calificación**: Aquí es donde tomas la oportunidad y comienzas el proceso de entender si es una buena opción para tu empresa.
- **Escribiendo Propuesta**: Aquí es donde comienzas a convertir la oportunidad en una propuesta para tu empresa. Este es un documento que enviarías al cliente.
- **Propuesta Enviada**: Aquí es donde has enviado la propuesta al cliente y estás esperando una respuesta.
- **Negociaciones**: Aquí es donde estás en el proceso de finalizar el trato.
- **Contrato Enviado para Firma**: Aquí es donde solo estás esperando que el cliente firme el contrato.
- **Cerrado Ganado**: Aquí es donde has ganado el trato y ahora estás trabajando en el proyecto.
- **Cerrado Perdido**: Aquí es donde has cotizado al cliente, pero no han aceptado los términos.
- **N/A**: Aquí colocas todas las oportunidades que no ganaste, pero que tampoco fueron "perdidas". Podría ser las que rechazaste, las que el cliente, por cualquier razón, te ignoró, y así sucesivamente.

### Etiquetas como Categorías de Servicio
Hablemos ahora de las etiquetas.

Recomendamos que uses etiquetas para los diferentes tipos de servicios que ofreces. Así que, volviendo a nuestro ejemplo de agencia de marketing y diseño, puedes tener etiquetas para "branding", "sitio web", "SEO", "Gestión de Facebook", y así sucesivamente.

Las ventajas aquí son que puedes filtrar fácilmente por servicio con un clic, lo que puede darte una breve visión general de qué servicios son más populares, y esto también puede informar futuras contrataciones, ya que típicamente diferentes servicios requieren diferentes miembros del equipo.

### Campos Personalizados del CRM de Ventas

A continuación, necesitamos considerar qué campos personalizados queremos tener.

Los típicos que vemos utilizados son:

- **Monto**: Este es un campo de moneda para el monto del proyecto.
- **Costo**: Tu costo esperado para cumplir con la venta, también un campo de moneda.
- **Beneficio**: Un campo de fórmula para calcular el beneficio basado en los campos de monto y costo.
- **URL de Propuesta**: Esto puede incluir un enlace a un documento de Google o Word en línea de tu propuesta, para que puedas hacer clic y revisarlo fácilmente.
- **Archivos Recibidos**: Este puede ser un campo personalizado de archivo donde puedes dejar cualquier archivo recibido del cliente, como materiales de investigación, NDAs, y así sucesivamente.
- **Contratos**: Otro campo personalizado de archivo donde puedes agregar contratos firmados para su custodia.
- **Nivel de Confianza**: Un campo personalizado de estrellas con 5 estrellas, indicando cuán seguro estás de ganar esta oportunidad en particular. ¡Esto puede usarse más adelante en el panel para pronósticos!
- **Fecha de Cierre Esperada**: Un campo de fecha para estimar cuándo es probable que se cierre el trato.
- **Cliente**: Un campo de referencia que vincula a la persona de contacto principal en la base de datos de clientes.
- **Nombre del Cliente**: Un campo de búsqueda que extrae el nombre del cliente del registro vinculado en la base de datos de clientes.
- **Correo Electrónico del Cliente**: Un campo de búsqueda que extrae el correo electrónico del cliente del registro vinculado en la base de datos de clientes.
- **Fuente del Trato**: Un campo desplegable para rastrear de dónde provino la oportunidad (por ejemplo, referencia, sitio web, llamada en frío, feria comercial).
- **Razón de Pérdida**: Un campo desplegable (para tratos cerrados perdidos) para categorizar por qué se perdió la oportunidad.
- **Tamaño del Cliente**: Un campo desplegable para categorizar a los clientes por tamaño (por ejemplo, pequeño, mediano, gran empresa).

Nuevamente, realmente **depende de ti** decidir exactamente qué campos deseas tener. Una advertencia: es fácil al configurar agregar muchos y muchos campos a tu CRM de ventas de datos que te gustaría capturar. Sin embargo, debes ser realista en términos de la disciplina y el compromiso de tiempo. No tiene sentido tener 30 campos en tu CRM de ventas si el 90% de los registros no tendrán datos en ellos.

Lo grandioso de los campos personalizados es que se integran bien en [Permisos Personalizados](/platform/features/user-permissions). Esto significa que puedes decidir exactamente qué campos los miembros de tu equipo pueden ver o editar. Por ejemplo, puede que desees ocultar información de costos y beneficios del personal junior.

### Automatizaciones

[Las Automatizaciones del CRM de Ventas](/platform/features/automations) son una característica poderosa en Blue que puede agilizar tu proceso de ventas, garantizar consistencia y ahorrar tiempo en tareas repetitivas. Al configurar automatizaciones inteligentes, puedes mejorar la efectividad de tu CRM de ventas y permitir que tu equipo se concentre en lo que más importa: cerrar tratos. Aquí hay algunas automatizaciones clave a considerar para tu CRM de ventas:

- **Asignación de Nuevos Leads**: Asigna automáticamente nuevos leads a representantes de ventas según criterios predefinidos como ubicación, tamaño del trato o industria. Esto asegura un seguimiento rápido y una distribución equilibrada de la carga de trabajo.
- **Recordatorios de Seguimiento**: Configura recordatorios automáticos para que los representantes de ventas sigan a los prospectos después de un cierto período de inactividad. Esto ayuda a prevenir que los leads se pierdan.
- **Notificaciones de Progresión de Etapas**: Notifica a los miembros relevantes del equipo cuando un trato pasa a una nueva etapa en el pipeline. Esto mantiene a todos informados sobre el progreso y permite intervenciones oportunas si es necesario.
- **Alertas de Envejecimiento de Tratos**: Crea alertas para tratos que han estado en una etapa particular durante más tiempo del esperado. Esto ayuda a identificar tratos estancados que pueden necesitar atención adicional.

## Vinculando Clientes y Tratos

Una de las características más poderosas de Blue para crear un sistema CRM efectivo es la capacidad de vincular tu base de datos de clientes con tus oportunidades de ventas. Esta conexión te permite mantener una única fuente de verdad para la información del cliente mientras rastreas múltiples tratos asociados con cada cliente. Vamos a explorar cómo configurar esto utilizando campos personalizados de Referencia y Búsqueda.

### Configurando el Campo de Referencia

1. En tu proyecto de Oportunidades (o CRM de Ventas), crea un nuevo campo personalizado.
2. Elige el tipo de campo "Referencia".
3. Selecciona tu proyecto de Base de Datos de Clientes como la fuente para la referencia.
4. Configura el campo para permitir selección única (ya que cada oportunidad está típicamente asociada con un cliente).
5. Nombra este campo algo como "Cliente" o "Empresa Asociada".

Ahora, al crear o editar una oportunidad, podrás seleccionar el cliente asociado de un menú desplegable poblado con registros de tu Base de Datos de Clientes.

### Mejorando con Campos de Búsqueda

Una vez que hayas establecido la conexión de referencia, puedes usar campos de Búsqueda para traer información relevante del cliente directamente a tu vista de oportunidades. Aquí te mostramos cómo:

1. En tu proyecto de Oportunidades, crea un nuevo campo personalizado.
2. Elige el tipo de campo "Búsqueda".
3. Selecciona el campo de Referencia que acabas de crear ("Cliente" o "Empresa Asociada") como la fuente.
4. Elige qué información del cliente deseas mostrar. Podrías considerar campos como: Correo Electrónico, Número de Teléfono, Categoría del Cliente, Gerente de Cuenta.

Repite este proceso para cada pieza de información del cliente que desees mostrar en tu vista de oportunidades.

Los beneficios de esto son:

- **Única Fuente de Verdad**: Actualiza la información del cliente una vez en la Base de Datos de Clientes, y se refleja automáticamente en todas las oportunidades vinculadas.
- **Eficiencia**: Accede rápidamente a los detalles relevantes del cliente mientras trabajas en oportunidades sin cambiar entre proyectos.
- **Integridad de Datos**: Reduce errores de entrada manual de datos al extraer automáticamente la información del cliente.
- **Visión Holística**: Ve fácilmente todas las oportunidades asociadas con un cliente utilizando el campo "Referencia por" en tu Base de Datos de Clientes.

### Consejo Avanzado: Buscar una Búsqueda

Blue ofrece una característica avanzada llamada "Buscar una Búsqueda" que puede ser increíblemente útil para configuraciones de CRM complejas. Esta característica te permite crear conexiones a través de múltiples proyectos, lo que te permite acceder a información tanto de tu Base de Datos de Clientes como del proyecto de Oportunidades en un tercer proyecto.

Por ejemplo, supongamos que tienes un espacio de trabajo de "Proyectos" donde gestionas el trabajo real para tus clientes. Quieres que este espacio de trabajo tenga acceso tanto a los detalles del cliente como a la información de oportunidades. Aquí te mostramos cómo puedes configurar esto:

Primero, crea un campo de Referencia en tu espacio de trabajo de Proyectos que vincule al proyecto de Oportunidades. Esto establece la conexión inicial. A continuación, crea campos de Búsqueda basados en esta Referencia para extraer detalles específicos de las oportunidades, como el valor del trato o la fecha de cierre esperada.

El verdadero poder viene en el siguiente paso: puedes crear campos de Búsqueda adicionales que lleguen a través de la Referencia de la oportunidad a la Base de Datos de Clientes. Esto te permite extraer información del cliente, como detalles de contacto o estado de la cuenta, directamente en tu espacio de trabajo de Proyectos.

Esta cadena de conexiones te da una vista integral en tu espacio de trabajo de Proyectos, combinando datos de tus oportunidades y base de datos de clientes. Es una forma poderosa de asegurarte de que tus equipos de proyecto tengan toda la información relevante al alcance de la mano sin necesidad de cambiar entre diferentes proyectos.

### Mejores Prácticas para Sistemas CRM Vinculados

Mantén tu Base de Datos de Clientes como la única fuente de verdad para toda la información del cliente. Siempre que necesites actualizar los detalles del cliente, hazlo primero en la Base de Datos de Clientes. Esto asegura que la información permanezca consistente en todos los proyectos vinculados.

Al crear campos de Referencia y Búsqueda, utiliza nombres claros y significativos. Esto ayuda a mantener la claridad, especialmente a medida que tu sistema se vuelve más complejo.

Revisa regularmente tu configuración para asegurarte de que estás extrayendo la información más relevante. A medida que las necesidades de tu negocio evolucionen, es posible que necesites agregar nuevos campos de Búsqueda o eliminar aquellos que ya no son útiles. Las revisiones periódicas ayudan a mantener tu sistema optimizado y efectivo.

Considera aprovechar las características de automatización de Blue para mantener tus datos sincronizados y actualizados en todos los proyectos. Por ejemplo, podrías configurar una automatización para notificar a los miembros relevantes del equipo cuando se actualice información clave del cliente en la Base de Datos de Clientes.

Al implementar efectivamente estas estrategias y hacer pleno uso de los campos de Referencia y Búsqueda, puedes crear un poderoso sistema CRM interconectado en Blue. Este sistema te proporcionará una visión integral de 360 grados de tus relaciones con los clientes y tu pipeline de ventas, lo que permitirá una toma de decisiones más informada y operaciones más fluidas en toda tu organización.

## Paneles

Los paneles son un componente crucial de cualquier sistema CRM efectivo, proporcionando información instantánea sobre tu rendimiento de ventas y relaciones con los clientes. La función de panel de Blue es particularmente poderosa porque te permite combinar datos en tiempo real de múltiples proyectos simultáneamente, dándote una visión integral de tus operaciones de ventas.

Al configurar tu panel de CRM en Blue, considera incluir varias métricas clave. El pipeline generado por mes muestra el valor total de nuevas oportunidades añadidas a tu pipeline, ayudándote a rastrear la capacidad de tu equipo para generar nuevos negocios. Las ventas por mes muestran tus tratos cerrados reales, permitiéndote monitorear el rendimiento de tu equipo en convertir oportunidades en ventas.

Introducir el concepto de descuentos en el pipeline puede llevar a pronósticos más precisos. Por ejemplo, podrías contar el 90% del valor de los tratos en la etapa de "Contrato Enviado para Firma", pero solo el 50% de los tratos en la etapa de "Propuesta Enviada". Este enfoque ponderado proporciona un pronóstico de ventas más realista.

Rastrear nuevas oportunidades por mes te ayuda a monitorear el número de nuevos tratos potenciales que ingresan a tu pipeline, lo cual es un buen indicador de los esfuerzos de prospección de tu equipo de ventas. Desglosar las ventas por tipo puede ayudarte a identificar tus ofertas más exitosas. Si configuras un proyecto de seguimiento de facturas vinculado a tus oportunidades, también puedes rastrear los ingresos reales en tu panel, proporcionando una imagen completa desde la oportunidad hasta el efectivo.

Blue ofrece varias características poderosas para ayudarte a crear un panel de CRM informativo e interactivo. La plataforma proporciona tres tipos principales de gráficos: tarjetas de estadísticas, gráficos de pastel y gráficos de barras. Las tarjetas de estadísticas son ideales para mostrar métricas clave como el valor total del pipeline o el número de oportunidades activas. Los gráficos de pastel son perfectos para mostrar la composición de tus ventas por tipo o la distribución de tratos en diferentes etapas. Los gráficos de barras son excelentes para comparar métricas a lo largo del tiempo, como ventas mensuales o nuevas oportunidades.

Las sofisticadas capacidades de filtrado de Blue te permiten segmentar tus datos por proyecto, lista, etiqueta y período de tiempo. Esto es particularmente útil para profundizar en aspectos específicos de tus datos de ventas o comparar el rendimiento entre diferentes equipos o productos. La plataforma consolida inteligentemente listas y etiquetas con el mismo nombre a través de proyectos, lo que permite un análisis fluido entre proyectos. Esto es invaluable para una configuración de CRM donde podrías tener proyectos separados para clientes, oportunidades y facturas.

La personalización es una fortaleza clave de la función de panel de Blue. La funcionalidad de arrastrar y soltar y la flexibilidad de visualización te permiten crear un panel que se adapte perfectamente a tus necesidades. Puedes reorganizar fácilmente gráficos y elegir la visualización más apropiada para cada métrica. 
Si bien los paneles son actualmente solo para uso interno, puedes compartirlos fácilmente con los miembros del equipo, otorgando permisos de vista o edición. Esto asegura que todos en tu equipo de ventas tengan acceso a la información que necesitan.

Al aprovechar estas características e incluir las métricas clave que hemos discutido, puedes crear un panel de CRM integral en Blue que proporcione información en tiempo real sobre tu rendimiento de ventas, salud del pipeline y crecimiento general del negocio. Este panel se convertirá en una herramienta invaluable para tomar decisiones basadas en datos y mantener a todo tu equipo alineado en tus objetivos y progreso de ventas.

## Conclusión

Configurar un CRM de ventas integral en Blue es una forma poderosa de agilizar tu proceso de ventas y obtener información valiosa sobre tus relaciones con los clientes y el rendimiento de tu negocio. Al seguir los pasos descritos en esta guía, has creado un sistema robusto que integra información del cliente, oportunidades de ventas y métricas de rendimiento en una plataforma única y cohesiva.

Comenzamos creando una base de datos de clientes, estableciendo una única fuente de verdad para toda tu información del cliente. Esta base permite mantener registros precisos y actualizados para todos tus clientes y prospectos. Luego, construimos sobre esto con una base de datos de oportunidades, permitiéndote rastrear y gestionar tu pipeline de ventas de manera efectiva.

Una de las principales fortalezas de usar Blue para tu CRM es la capacidad de vincular estas bases de datos utilizando campos de referencia y búsqueda. Esta integración crea un sistema dinámico donde las actualizaciones a la información del cliente se reflejan instantáneamente en todas las oportunidades relacionadas, asegurando la consistencia de los datos y ahorrando tiempo en actualizaciones manuales.
Exploramos cómo aprovechar las poderosas características de automatización de Blue para agilizar tu flujo de trabajo, desde la asignación de nuevos leads hasta el envío de recordatorios de seguimiento. Estas automatizaciones ayudan a garantizar que no se pierdan oportunidades y que tu equipo pueda concentrarse en actividades de alto valor en lugar de tareas administrativas.

Finalmente, profundizamos en la creación de paneles que proporcionan información instantánea sobre tu rendimiento de ventas. Al combinar datos de tus bases de datos de clientes y oportunidades, estos paneles ofrecen una visión integral de tu pipeline de ventas, tratos cerrados y salud general del negocio.

Recuerda, la clave para obtener el máximo provecho de tu CRM es el uso constante y la refinación regular. Anima a tu equipo a adoptar completamente el sistema, revisa regularmente tus procesos y automatizaciones, y continúa explorando nuevas formas de aprovechar las características de Blue para apoyar tus esfuerzos de ventas.

Con esta configuración de CRM de ventas en Blue, estás bien equipado para cultivar relaciones con los clientes, cerrar más tratos y llevar tu negocio hacia adelante.