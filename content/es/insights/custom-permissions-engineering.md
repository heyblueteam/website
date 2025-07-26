---
title: Creando el Motor de Permisos Personalizados de Blue
description: Acompáñanos tras bambalinas con el equipo de ingeniería de Blue mientras explican cómo construir una función de auto-categorización y etiquetado impulsada por IA.
category: "Engineering"
date: 2024-07-25
---


La gestión efectiva de proyectos y procesos es crucial para organizaciones de todos los tamaños.

En Blue, [hemos hecho de nuestra misión](/about) organizar el trabajo del mundo construyendo la mejor plataforma de gestión de proyectos del planeta: simple, poderosa, flexible y asequible para todos.

Esto significa que nuestra plataforma debe adaptarse a las necesidades únicas de cada equipo. Hoy, estamos emocionados de revelar uno de nuestras características más poderosas: Permisos Personalizados.

Las herramientas de gestión de proyectos son la columna vertebral de los flujos de trabajo modernos, albergando datos sensibles, comunicaciones cruciales y planes estratégicos. Como tal, la capacidad de controlar finamente el acceso a esta información no es solo un lujo, es una necesidad.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>

Los permisos personalizados juegan un papel crítico en las plataformas B2B SaaS, especialmente en las herramientas de gestión de proyectos, donde el equilibrio entre colaboración y seguridad puede determinar el éxito de un proyecto.

Pero aquí es donde Blue adopta un enfoque diferente: **creemos que las características de nivel empresarial no deberían reservarse solo para presupuestos de gran empresa.**

En una era donde la IA empodera a pequeños equipos para operar a escalas sin precedentes, ¿por qué la seguridad robusta y la personalización deberían estar fuera de alcance?

En esta mirada tras bambalinas, exploraremos cómo desarrollamos nuestra función de Permisos Personalizados, desafiando el statu quo de los niveles de precios SaaS y llevando opciones de seguridad poderosas y flexibles a empresas de todos los tamaños.

Ya seas una startup con grandes sueños o un jugador establecido que busca optimizar sus procesos, los permisos personalizados pueden habilitar nuevos casos de uso que nunca supiste que eran posibles.

## Entendiendo los Permisos de Usuario Personalizados

Antes de sumergirnos en nuestro viaje de desarrollo de permisos personalizados para Blue, tomemos un momento para entender qué son los permisos de usuario personalizados y por qué son tan cruciales en el software de gestión de proyectos.

Los permisos de usuario personalizados se refieren a la capacidad de adaptar los derechos de acceso para usuarios individuales o grupos dentro de un sistema de software. En lugar de depender de roles predefinidos con conjuntos fijos de permisos, los permisos personalizados permiten a los administradores crear perfiles de acceso altamente específicos que se alinean perfectamente con la estructura y las necesidades de flujo de trabajo de su organización.

En el contexto de software de gestión de proyectos como Blue, los permisos personalizados incluyen:

* **Control de acceso granular**: Determinar quién puede ver, editar o eliminar tipos específicos de datos del proyecto.
* **Restricciones basadas en características**: Habilitar o deshabilitar ciertas características para usuarios o equipos particulares.
* **Niveles de sensibilidad de datos**: Establecer diferentes niveles de acceso a información sensible dentro de los proyectos.
* **Permisos específicos de flujo de trabajo**: Alinear las capacidades del usuario con etapas o aspectos específicos de su flujo de trabajo de proyecto.

La importancia de los permisos personalizados en la gestión de proyectos no puede ser subestimada:

* **Seguridad mejorada**: Al proporcionar a los usuarios solo el acceso que necesitan, se reduce el riesgo de violaciones de datos o cambios no autorizados.
* **Mejora de la conformidad**: Los permisos personalizados ayudan a las organizaciones a cumplir con requisitos regulatorios específicos de la industria al controlar el acceso a los datos.
* **Colaboración optimizada**: Los equipos pueden trabajar de manera más eficiente cuando cada miembro tiene el nivel adecuado de acceso para desempeñar su rol sin restricciones innecesarias o privilegios abrumadores.
* **Flexibilidad para organizaciones complejas**: A medida que las empresas crecen y evolucionan, los permisos personalizados permiten que el software se adapte a las estructuras organizativas y procesos cambiantes.

## Llegando al SÍ

[Hemos escrito antes](/insights/value-proposition-blue) que cada característica en Blue debe ser un **SÍ** rotundo antes de decidir construirla. No tenemos el lujo de contar con cientos de ingenieros y desperdiciar tiempo y dinero construyendo cosas que los clientes no necesitan.

Y así, el camino para implementar permisos personalizados en Blue no fue una línea recta. Como muchas características poderosas, comenzó con una necesidad clara de nuestros usuarios y evolucionó a través de una cuidadosa consideración y planificación.

Durante años, nuestros clientes habían estado solicitando un control más granular sobre los permisos de usuario. A medida que organizaciones de todos los tamaños comenzaron a manejar proyectos cada vez más complejos y sensibles, las limitaciones de nuestro control de acceso basado en roles estándar se hicieron evidentes.

Pequeñas startups que trabajan con clientes externos, empresas medianas con procesos de aprobación intrincados y grandes empresas con estrictos requisitos de cumplimiento expresaron la misma necesidad:

Más flexibilidad en la gestión del acceso de los usuarios.

A pesar de la clara demanda, inicialmente dudamos en sumergirnos en el desarrollo de permisos personalizados.

¿Por qué?

¡Entendimos la complejidad involucrada!

Los permisos personalizados tocan cada parte de un sistema de gestión de proyectos, desde la interfaz de usuario hasta la estructura de la base de datos. Sabíamos que implementar esta característica requeriría cambios significativos en nuestra arquitectura central y una cuidadosa consideración de las implicaciones de rendimiento.

A medida que examinamos el panorama, notamos que muy pocos de nuestros competidores habían intentado implementar un motor de permisos personalizados tan poderoso como el que nuestros clientes estaban solicitando. Aquellos que lo hicieron a menudo lo reservaron para sus planes empresariales de nivel más alto.

Se hizo claro por qué: el esfuerzo de desarrollo es sustancial y las apuestas son altas.

Implementar permisos personalizados incorrectamente podría introducir errores críticos o vulnerabilidades de seguridad, comprometiendo potencialmente todo el sistema. Esta realización subrayó la magnitud del desafío que estábamos considerando.

### Desafiando el Statu Quo

Sin embargo, a medida que continuamos creciendo y evolucionando, llegamos a una realización crucial:

**El modelo SaaS tradicional de reservar características poderosas para clientes empresariales ya no tiene sentido en el panorama empresarial actual.**

En 2024, con el poder de la IA y herramientas avanzadas, los pequeños equipos pueden operar a una escala y complejidad que rivalizan con organizaciones mucho más grandes. Una startup podría estar manejando datos sensibles de clientes en múltiples países. Una pequeña agencia de marketing podría estar gestionando docenas de proyectos de clientes con diferentes requisitos de confidencialidad. Estas empresas necesitan el mismo nivel de seguridad y personalización que *cualquier* gran empresa.

Nos preguntamos: ¿Por qué debería el tamaño de la fuerza laboral o el presupuesto de una empresa determinar su capacidad para mantener sus datos seguros y sus procesos eficientes?

### Características de Nivel Empresarial para Todos

Esta realización nos llevó a una filosofía central que ahora impulsa gran parte de nuestro desarrollo en Blue: Las características de nivel empresarial deberían ser accesibles para empresas de todos los tamaños.

Creemos que:

- **La seguridad no debería ser un lujo.** Cada empresa, independientemente de su tamaño, merece las herramientas para proteger sus datos y procesos.
- **La flexibilidad impulsa la innovación.** Al dar a todos nuestros usuarios herramientas poderosas, les permitimos crear flujos de trabajo y sistemas que impulsan sus industrias hacia adelante.
- **El crecimiento no debería requerir cambios en la plataforma.** A medida que nuestros clientes crecen, sus herramientas deberían crecer sin problemas con ellos.

Con esta mentalidad, decidimos abordar el desafío de los permisos personalizados de frente, comprometidos a hacerlo disponible para todos nuestros usuarios, no solo para aquellos en planes de nivel superior.

Esta decisión nos llevó por un camino de diseño cuidadoso, desarrollo iterativo y retroalimentación continua de los usuarios que finalmente condujo a la característica de permisos personalizados de la que estamos orgullosos de ofrecer hoy.

En la siguiente sección, profundizaremos en cómo abordamos el proceso de diseño y desarrollo para dar vida a esta característica compleja.

### Diseño y Desarrollo

Cuando decidimos abordar los permisos personalizados, rápidamente nos dimos cuenta de que estábamos enfrentando una tarea titánica.

A primera vista, "permisos personalizados" puede sonar sencillo, pero es una característica engañosamente compleja que toca cada aspecto de nuestro sistema.

El desafío era desalentador: necesitábamos implementar permisos en cascada, permitir ediciones sobre la marcha, realizar cambios significativos en el esquema de la base de datos y asegurar una funcionalidad fluida en todo nuestro ecosistema: aplicaciones web, Mac, Windows, iOS y Android, así como nuestra API y webhooks.

La complejidad era suficiente para hacer que incluso los desarrolladores más experimentados se detuvieran.

Nuestro enfoque se centró en dos principios clave:

1. Desglosar la característica en versiones manejables
2. Adoptar un envío incremental.

Ante la complejidad de los permisos personalizados a gran escala, nos hicimos una pregunta crucial:

> ¿Cuál sería la versión más simple posible de esta característica?

Este enfoque se alinea con el principio ágil de entregar un Producto Mínimamente Viable (MVP) e iterar en función de la retroalimentación.

Nuestra respuesta fue refrescantemente sencilla:

1. Introducir un interruptor para ocultar la pestaña de actividad del proyecto
2. Agregar otro interruptor para ocultar la pestaña de formularios

**Eso fue todo.**

Sin campanas ni silbatos, sin matrices de permisos complejas, solo dos simples interruptores de encendido/apagado.

Si bien puede parecer poco impresionante a primera vista, este enfoque ofreció varias ventajas significativas:

* **Implementación Rápida**: Estos simples interruptores podían desarrollarse y probarse rápidamente, lo que nos permitió poner una versión básica de permisos personalizados en manos de los usuarios rápidamente.
* **Valor Claro para el Usuario**: Incluso con solo estas dos opciones, estábamos proporcionando un valor tangible. Algunos equipos podrían querer ocultar el feed de actividad de los clientes, mientras que otros podrían necesitar restringir el acceso a formularios para ciertos grupos de usuarios.
* **Fundación para el Crecimiento**: Este inicio simple sentó las bases para permisos más complejos. Nos permitió configurar la infraestructura básica para permisos personalizados sin quedar atrapados en la complejidad desde el principio.
* **Retroalimentación de Usuarios**: Al lanzar esta versión simple, pudimos recopilar retroalimentación del mundo real sobre cómo los usuarios interactuaban con los permisos personalizados, informando nuestro desarrollo futuro.
* **Aprendizaje Técnico**: Esta implementación inicial brindó a nuestro equipo de desarrollo experiencia práctica en la modificación de permisos a través de nuestra plataforma, preparándonos para iteraciones más complejas.

Y, sabes, es realmente bastante humilde tener una gran visión para algo y luego lanzar algo que es un porcentaje tan pequeño de esa visión.

Después de lanzar estos primeros dos interruptores, decidimos abordar algo más sofisticado. Optamos por dos nuevos permisos de rol de usuario personalizados.

El primero fue la capacidad de limitar a los usuarios a solo ver registros que han sido específicamente asignados a ellos. Esto es muy útil si tienes un cliente en un proyecto y solo quieres que vea registros que están específicamente asignados a él en lugar de todo lo que estás trabajando para él.

El segundo fue una opción para que los administradores del proyecto bloquearan grupos de usuarios de poder invitar a otros usuarios. Esto es bueno si tienes un proyecto sensible que deseas asegurar que permanezca en una base de "necesidad de ver".

Una vez que lanzamos esto, ganamos más confianza y para nuestra tercera versión abordamos permisos a nivel de columna, lo que significa poder decidir qué campos personalizados puede ver o editar un grupo de usuarios específico.

Esto es extremadamente poderoso. Imagina que tienes un proyecto de CRM y tienes datos allí que no solo están relacionados con los montos que el cliente pagará, sino también con tus costos y márgenes de beneficio. Puede que no desees que tus campos de costo y el campo de fórmula de margen del proyecto sean visibles para el personal junior, y los permisos personalizados te permiten bloquear esos campos para que no se muestren.

A continuación, pasamos a crear permisos basados en listas, donde los administradores del proyecto pueden decidir si un grupo de usuarios puede ver, editar y eliminar una lista específica. Si ocultan una lista, todos los registros dentro de esa lista también se vuelven ocultos, lo cual es genial porque significa que puedes ocultar ciertas partes de tu proceso de tus miembros del equipo o clientes.

Este es el resultado final:

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Consideraciones Técnicas

En el corazón de la arquitectura técnica de Blue se encuentra GraphQL, una elección fundamental que ha influido significativamente en nuestra capacidad para implementar características complejas como los permisos personalizados. Pero antes de profundizar en los detalles, retrocedamos un paso y entendamos qué es GraphQL y cómo se diferencia del enfoque más tradicional de API REST.
GraphQL vs API REST: Una Explicación Accesible

Imagina que estás en un restaurante. Con una API REST, es como pedir de un menú fijo. Pides un plato específico (endpoint) y obtienes todo lo que viene con él, ya sea que lo quieras todo o no. Si deseas personalizar tu comida, es posible que necesites hacer múltiples pedidos (llamadas a la API) o pedir un plato especialmente preparado (endpoint personalizado).

GraphQL, por otro lado, es como tener una conversación con un chef que puede preparar cualquier cosa. Le dices al chef exactamente qué ingredientes quieres (campos de datos) y en qué cantidades. El chef luego prepara un plato que es precisamente lo que pediste, ni más ni menos. Esto es esencialmente lo que hace GraphQL: permite al cliente pedir exactamente los datos que necesita y el servidor proporciona solo eso.

### Un Almuerzo Importante

Aproximadamente seis semanas después del desarrollo inicial de Blue, nuestro ingeniero principal y CEO salió a almorzar.

¿El tema de discusión?

Si cambiar de APIs REST a GraphQL. Esta no era una decisión que se pudiera tomar a la ligera: adoptar GraphQL significaría descartar seis semanas de trabajo inicial.

En el camino de regreso a la oficina, el CEO planteó una pregunta crucial al ingeniero principal: "¿Nos arrepentiremos de no haber hecho esto dentro de cinco años?"

La respuesta se volvió clara: GraphQL era el camino a seguir.

Reconocimos el potencial de esta tecnología desde el principio, viendo cómo podría apoyar nuestra visión de una plataforma de gestión de proyectos flexible y poderosa.

Nuestra previsión al adoptar GraphQL dio sus frutos cuando se trató de implementar permisos personalizados. Con una API REST, habríamos necesitado un endpoint diferente para cada posible configuración de permisos personalizados, un enfoque que rápidamente se volvería engorroso y difícil de mantener.

GraphQL, sin embargo, nos permite manejar los permisos personalizados de manera dinámica. Así es como funciona:

- **Verificaciones de Permisos en Tiempo Real**: Cuando un cliente hace una solicitud, nuestro servidor GraphQL puede verificar los permisos del usuario directamente desde nuestra base de datos.
- **Recuperación de Datos Precisa**: Basado en estos permisos, GraphQL devuelve solo los datos solicitados que se ajustan a los derechos de acceso del usuario.
- **Consultas Flexibles**: A medida que cambian los permisos, no necesitamos crear nuevos endpoints ni alterar los existentes. La misma consulta GraphQL puede adaptarse a diferentes configuraciones de permisos.
- **Obtención Eficiente de Datos**: GraphQL permite a los clientes solicitar exactamente lo que necesitan. Esto significa que no estamos sobrecargando datos, lo que podría exponer información a la que el usuario no debería tener acceso.

Esta flexibilidad es crucial para una característica tan compleja como los permisos personalizados. Permite ofrecer un control granular *sin* sacrificar el rendimiento o la mantenibilidad.

## Desafíos

Implementar permisos personalizados en Blue trajo consigo una serie de desafíos, cada uno de los cuales nos empujó a innovar y refinar nuestro enfoque. La optimización del rendimiento rápidamente emergió como una preocupación crítica. A medida que agregamos más verificaciones de permisos granulares, corríamos el riesgo de ralentizar nuestro sistema, especialmente para proyectos grandes con muchos usuarios y configuraciones de permisos complejas. Para abordar esto, implementamos una estrategia de caché de múltiples niveles, optimizamos nuestras consultas de base de datos y aprovechamos la capacidad de GraphQL para solicitar solo los datos necesarios. Este enfoque nos permitió mantener tiempos de respuesta rápidos incluso a medida que los proyectos escalaban y la complejidad de los permisos crecía.

La interfaz de usuario para los permisos personalizados presentó otro obstáculo significativo. Necesitábamos hacer que la interfaz fuera intuitiva y manejable para los administradores, incluso a medida que agregábamos más opciones y aumentábamos la complejidad del sistema.

Nuestra solución involucró múltiples rondas de pruebas de usuarios y diseño iterativo.

Introdujimos una matriz visual de permisos que permitió a los administradores ver y modificar rápidamente los permisos a través de diferentes roles y áreas del proyecto.

Asegurar la consistencia entre plataformas presentó su propio conjunto de desafíos. Necesitábamos implementar permisos personalizados de manera uniforme en nuestras aplicaciones web, de escritorio y móviles, cada una con su interfaz y consideraciones de experiencia de usuario únicas. Esto fue particularmente complicado para nuestras aplicaciones móviles, que debían ocultar y mostrar dinámicamente características basadas en los permisos del usuario. Abordamos esto centralizando nuestra lógica de permisos en la capa de API, asegurando que todas las plataformas recibieran datos de permisos consistentes.

Luego, desarrollamos un marco de interfaz de usuario flexible que pudiera adaptarse a estos cambios de permisos en tiempo real, proporcionando una experiencia fluida independientemente de la plataforma utilizada.

La educación y adopción de usuarios presentaron el último obstáculo en nuestro viaje de permisos personalizados. Introducir una característica tan poderosa significaba que necesitábamos ayudar a nuestros usuarios a entender y aprovechar efectivamente los permisos personalizados.

Inicialmente lanzamos los permisos personalizados a un subconjunto de nuestra base de usuarios, monitoreando cuidadosamente sus experiencias y recopilando información. Este enfoque nos permitió refinar la característica y nuestros materiales educativos basados en el uso del mundo real antes de lanzarlo a toda nuestra base de usuarios.

El lanzamiento por fases resultó invaluable, ayudándonos a identificar y abordar problemas menores y puntos de confusión de los usuarios que no habíamos anticipado, lo que finalmente condujo a una característica más pulida y fácil de usar para todos nuestros usuarios.

Este enfoque de lanzar a un subconjunto de usuarios, así como nuestro típico período de "Beta" de 2-3 semanas en nuestra Beta pública, nos ayuda a dormir tranquilos por la noche. :)

## Mirando Hacia Adelante

Como con todas las características, nada está *"terminado"*.

Nuestra visión a largo plazo para la característica de permisos personalizados se extiende a etiquetas, filtros de campos personalizados, navegación de proyectos personalizable y controles de comentarios.

Vamos a profundizar en cada aspecto.

### Permisos de Etiquetas

Creemos que sería increíble poder crear permisos basados en si un registro tiene una o más etiquetas. El caso de uso más obvio sería crear un rol de usuario personalizado llamado "Clientes" y permitir que solo los usuarios en ese rol vean registros que tengan la etiqueta "Clientes".

Esto te brinda una vista rápida de si un registro puede o no ser visto por tus clientes.

Esto podría volverse aún más poderoso con combinadores AND/OR, donde puedes especificar reglas más complejas. Por ejemplo, podrías configurar una regla que permita el acceso a registros etiquetados tanto "Clientes" COMO "Público", o registros etiquetados ya sea "Interno" O "Confidencial". Este nivel de flexibilidad permitiría configuraciones de permisos increíblemente matizadas, atendiendo incluso a las estructuras organizativas y flujos de trabajo más complejos.

Las aplicaciones potenciales son vastas. Los gerentes de proyecto podrían segregar fácilmente información sensible, los equipos de ventas podrían tener acceso automático a datos relevantes de clientes, y los colaboradores externos podrían integrarse sin problemas en partes específicas de un proyecto sin arriesgarse a exponer información interna sensible.

### Filtros de Campos Personalizados

Nuestra visión para los Filtros de Campos Personalizados representa un avance significativo en el control de acceso granular. Esta característica empoderará a los administradores de proyectos para definir qué registros pueden ver grupos específicos de usuarios en función de los valores de campos personalizados. Se trata de crear límites dinámicos y basados en datos para el acceso a la información.

Imagina poder configurar permisos así:

- Mostrar solo registros donde el desplegable "Estado del Proyecto" esté configurado como "Público"
- Restringir la visibilidad a elementos donde el campo de selección múltiple "Departamento" incluya "Marketing"
- Permitir acceso a tareas donde la casilla de verificación "Prioridad" esté marcada
- Mostrar proyectos donde el campo numérico "Presupuesto" esté por encima de un cierto umbral

### Navegación de Proyectos Personalizable

Esto es simplemente una extensión de los interruptores que ya tenemos. En lugar de solo tener interruptores para "actividad" y "formularios", queremos extender eso a cada parte de la navegación del proyecto. De esta manera, los administradores del proyecto pueden crear interfaces enfocadas y eliminar herramientas que no necesitan.

### Controles de Comentarios

En el futuro, queremos ser creativos en cómo permitimos que nuestros clientes decidan quién puede y no puede ver comentarios. Podríamos permitir múltiples áreas de comentarios con pestañas bajo un registro, y cada una puede ser visible o no visible para diferentes grupos de usuarios.

Además, también podríamos permitir una característica donde solo los comentarios donde un usuario es *específicamente* mencionado son visibles, y nada más. Esto permitiría a los equipos que tienen clientes en proyectos asegurarse de que solo los comentarios que desean que los clientes vean sean visibles.

## Conclusión

Así que ahí lo tenemos, ¡así es como abordamos la construcción de una de las características más interesantes y poderosas! [Como puedes ver en nuestra herramienta de comparación de gestión de proyectos](/compare), muy pocos sistemas de gestión de proyectos tienen una configuración de matriz de permisos tan poderosa, y los que lo hacen lo reservan para sus planes empresariales más caros, haciéndolo inaccesible para una empresa pequeña o mediana típica.

Con Blue, tienes *todas* las características disponibles con nuestro plan: ¡no creemos que las características de nivel empresarial deban reservarse para clientes empresariales!