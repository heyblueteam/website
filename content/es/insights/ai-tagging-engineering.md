---
title: Categorización Automática con IA (Análisis Profundo de Ingeniería)
category: "Engineering"
description: Vaya detrás de escena con el equipo de ingeniería de Blue mientras explican cómo construyeron una función de categorización automática y etiquetado impulsada por IA.
date: 2024-12-07
---

Recientemente lanzamos la [Categorización Automática con IA](/insights/ai-auto-categorization) para todos los usuarios de Blue. Esta es una función de IA que está incluida en la suscripción principal de Blue, sin costos adicionales. En esta publicación, profundizamos en la ingeniería detrás de hacer realidad esta función.

---
En Blue, nuestro enfoque para el desarrollo de funciones está arraigado en una comprensión profunda de las necesidades de los usuarios y las tendencias del mercado, junto con un compromiso de mantener la simplicidad y facilidad de uso que define nuestra plataforma. Esto es lo que impulsa nuestra [hoja de ruta](/platform/roadmap), y lo que nos ha [permitido entregar consistentemente funciones cada mes durante años](/platform/changelog).

La introducción del etiquetado automático impulsado por IA en Blue es un ejemplo perfecto de esta filosofía en acción. Antes de sumergirnos en los detalles técnicos de cómo construimos esta función, es crucial entender el problema que estábamos resolviendo y la cuidadosa consideración que se puso en su desarrollo.

El panorama de la gestión de proyectos está evolucionando rápidamente, con las capacidades de IA volviéndose cada vez más centrales para las expectativas de los usuarios. Nuestros clientes, particularmente aquellos que gestionan [proyectos](/platform) a gran escala con millones de [registros](/platform/features/records), habían sido muy vocales sobre su deseo de formas más inteligentes y eficientes de organizar y categorizar sus datos.

Sin embargo, en Blue, no simplemente agregamos funciones porque están de moda o son solicitadas. Nuestra filosofía es que cada nueva adición debe demostrar su valor, con la respuesta predeterminada siendo un firme *"no"* hasta que una función demuestre una fuerte demanda y utilidad clara.

Para comprender verdaderamente la profundidad del problema y el potencial del etiquetado automático con IA, realizamos extensas entrevistas con clientes, enfocándonos en usuarios de largo plazo que gestionan proyectos complejos y ricos en datos en múltiples dominios.

Estas conversaciones revelaron un hilo común: *aunque el etiquetado era invaluable para la organización y la capacidad de búsqueda, la naturaleza manual del proceso se estaba convirtiendo en un cuello de botella, especialmente para los equipos que manejan grandes volúmenes de registros.*

Pero vimos más allá de solo resolver el punto de dolor inmediato del etiquetado manual.

Imaginamos un futuro donde el etiquetado impulsado por IA podría convertirse en la base para flujos de trabajo más inteligentes y automatizados.

El verdadero poder de esta función, nos dimos cuenta, radicaba en su potencial integración con nuestro [sistema de automatización de gestión de proyectos](/platform/features/automations). Imagine una herramienta de gestión de proyectos que no solo categoriza información de manera inteligente, sino que también usa esas categorías para enrutar tareas, activar acciones y adaptar flujos de trabajo en tiempo real.

Esta visión se alineó perfectamente con nuestro objetivo de mantener Blue simple pero poderoso.

Además, reconocimos el potencial de extender esta capacidad más allá de los confines de nuestra plataforma. Al desarrollar un sistema robusto de etiquetado con IA, estábamos sentando las bases para una "API de categorización" que podría funcionar de inmediato, potencialmente abriendo nuevas vías sobre cómo nuestros usuarios interactúan y aprovechan Blue en sus ecosistemas tecnológicos más amplios.

Esta función, por lo tanto, no se trataba solo de agregar una casilla de verificación de IA a nuestra lista de funciones.

Se trataba de dar un paso significativo hacia una plataforma de gestión de proyectos más inteligente y adaptativa mientras permanecíamos fieles a nuestra filosofía central de simplicidad y centricidad en el usuario.

En las siguientes secciones, profundizaremos en los desafíos técnicos que enfrentamos al dar vida a esta visión, la arquitectura que diseñamos para soportarla y las soluciones que implementamos. También exploraremos las posibilidades futuras que esta función abre, mostrando cómo una adición cuidadosamente considerada puede allanar el camino para cambios transformadores en la gestión de proyectos.

---
## El Problema

Como se discutió anteriormente, el etiquetado manual de registros de proyectos puede consumir mucho tiempo y ser inconsistente.

Nos propusimos resolver esto aprovechando la IA para sugerir automáticamente etiquetas basadas en el contenido del registro.

Los principales desafíos fueron:

1. Elegir un modelo de IA apropiado
2. Procesar eficientemente grandes volúmenes de registros
3. Garantizar la privacidad y seguridad de los datos
4. Integrar la función sin problemas en nuestra arquitectura existente

## Selección del Modelo de IA

Evaluamos varias plataformas de IA, incluyendo [OpenAI](https://openai.com), modelos de código abierto en [HuggingFace](https://huggingface.co/), y [Replicate](https://replicate.com).

Nuestros criterios incluyeron:

- Costo-efectividad
- Precisión en la comprensión del contexto
- Capacidad para adherirse a formatos de salida específicos
- Garantías de privacidad de datos

Después de pruebas exhaustivas, elegimos [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo) de OpenAI. Aunque [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) podría ofrecer mejoras marginales en precisión, nuestras pruebas mostraron que el rendimiento de GPT-3.5 era más que adecuado para nuestras necesidades de etiquetado automático. El equilibrio entre costo-efectividad y fuertes capacidades de categorización hizo de GPT-3.5 la elección ideal para esta función.

El costo más alto de GPT-4 nos habría obligado a ofrecer la función como un complemento pagado, entrando en conflicto con nuestro objetivo de **incluir la IA dentro de nuestro producto principal sin costo adicional para los usuarios finales.**

Al momento de nuestra implementación, los precios para GPT-3.5 Turbo son:

- $0.0005 por 1K tokens de entrada (o $0.50 por 1M tokens de entrada)
- $0.0015 por 1K tokens de salida (o $1.50 por 1M tokens de salida)

Hagamos algunas suposiciones sobre un registro promedio en Blue:

- **Título**: ~10 tokens
- **Descripción**: ~50 tokens
- **2 comentarios**: ~30 tokens cada uno
- **5 campos personalizados**: ~10 tokens cada uno
- **Nombre de lista, fecha de vencimiento y otros metadatos**: ~20 tokens
- **Indicación del sistema y etiquetas disponibles**: ~50 tokens

Total de tokens de entrada por registro: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 tokens

Para la salida, supongamos un promedio de 3 etiquetas sugeridas por registro, lo que podría totalizar alrededor de 20 tokens de salida incluyendo el formato JSON.

Para 1 millón de registros:

- Costo de entrada: (240 * 1,000,000 / 1,000,000) * $0.50 = $120
- Costo de salida: (20 * 1,000,000 / 1,000,000) * $1.50 = $30

**Costo total para el etiquetado automático de 1 millón de registros: $120 + $30 = $150**

## Rendimiento de GPT3.5 Turbo

La categorización es una tarea en la que los modelos de lenguaje grandes (LLMs) como GPT-3.5 Turbo sobresalen, haciéndolos particularmente adecuados para nuestra función de etiquetado automático. Los LLMs están entrenados en vastas cantidades de datos de texto, permitiéndoles entender el contexto, la semántica y las relaciones entre conceptos. Esta amplia base de conocimiento les permite realizar tareas de categorización con alta precisión en una amplia gama de dominios.

Para nuestro caso de uso específico de etiquetado de gestión de proyectos, GPT-3.5 Turbo demuestra varias fortalezas clave:

- **Comprensión Contextual:** Puede captar el contexto general de un registro de proyecto, considerando no solo palabras individuales sino el significado transmitido por toda la descripción, comentarios y otros campos.
- **Flexibilidad:** Puede adaptarse a varios tipos de proyectos e industrias sin requerir reprogramación extensa.
- **Manejo de Ambigüedad:** Puede sopesar múltiples factores para tomar decisiones matizadas.
- **Aprendizaje de Ejemplos:** Puede comprender y aplicar rápidamente nuevos esquemas de categorización sin entrenamiento adicional.
- **Clasificación Multi-etiqueta:** Puede sugerir múltiples etiquetas relevantes para un solo registro, lo cual fue crucial para nuestros requisitos.

GPT-3.5 Turbo también destacó por su confiabilidad en adherirse a nuestro formato de salida JSON requerido, lo cual fue *crucial* para la integración sin problemas con nuestros sistemas existentes. Los modelos de código abierto, aunque prometedores, a menudo agregaban comentarios adicionales o se desviaban del formato esperado, lo que habría requerido procesamiento adicional. Esta consistencia en el formato de salida fue un factor clave en nuestra decisión, ya que simplificó significativamente nuestra implementación y redujo los puntos potenciales de falla.

Optar por GPT-3.5 Turbo con su salida JSON consistente nos permitió implementar una solución más directa, confiable y mantenible.

Si hubiéramos elegido un modelo con formato menos confiable, habríamos enfrentado una cascada de complicaciones: la necesidad de lógica de análisis robusta para manejar varios formatos de salida, manejo extensivo de errores para salidas inconsistentes, impactos potenciales en el rendimiento del procesamiento adicional, mayor complejidad de pruebas para cubrir todas las variaciones de salida, y una mayor carga de mantenimiento a largo plazo.

Los errores de análisis podrían llevar a un etiquetado incorrecto, impactando negativamente la experiencia del usuario. Al evitar estos obstáculos, pudimos enfocar nuestros esfuerzos de ingeniería en aspectos críticos como la optimización del rendimiento y el diseño de la interfaz de usuario, en lugar de luchar con salidas de IA impredecibles.

## Arquitectura del Sistema

Nuestra función de etiquetado automático con IA está construida sobre una arquitectura robusta y escalable diseñada para manejar grandes volúmenes de solicitudes de manera eficiente mientras proporciona una experiencia de usuario sin problemas. Como con todos nuestros sistemas, hemos arquitecturado esta función para soportar un orden de magnitud más de tráfico del que experimentamos actualmente. Este enfoque, aunque aparentemente sobreingeniería para las necesidades actuales, es una mejor práctica que nos permite manejar sin problemas picos repentinos en el uso y nos da amplio margen para el crecimiento sin revisiones arquitectónicas importantes. De lo contrario, tendríamos que reingenierar todos nuestros sistemas cada 18 meses, ¡algo que hemos aprendido por las malas en el pasado!

Desglosemos los componentes y el flujo de nuestro sistema:

- **Interacción del Usuario:** El proceso comienza cuando un usuario presiona el botón "Autoetiquetar" en la interfaz de Blue. Esta acción activa el flujo de trabajo de etiquetado automático.
- **Llamada API de Blue:** La acción del usuario se traduce en una llamada API a nuestro backend de Blue. Este endpoint API está diseñado para manejar solicitudes de etiquetado automático.
- **Gestión de Cola:** En lugar de procesar la solicitud inmediatamente, lo que podría llevar a problemas de rendimiento bajo carga alta, agregamos la solicitud de etiquetado a una cola. Usamos Redis para este mecanismo de cola, lo que nos permite gestionar la carga de manera efectiva y garantizar la escalabilidad del sistema.
- **Servicio en Segundo Plano:** Implementamos un servicio en segundo plano que monitorea continuamente la cola en busca de nuevas solicitudes. Este servicio es responsable de procesar las solicitudes en cola.
- **Integración con API de OpenAI:** El servicio en segundo plano prepara los datos necesarios y realiza llamadas API al modelo GPT-3.5 de OpenAI. Aquí es donde ocurre el etiquetado real impulsado por IA. Enviamos datos relevantes del proyecto y recibimos etiquetas sugeridas a cambio.
- **Procesamiento de Resultados:** El servicio en segundo plano procesa los resultados recibidos de OpenAI. Esto implica analizar la respuesta de la IA y preparar los datos para su aplicación al proyecto.
- **Aplicación de Etiquetas:** Los resultados procesados se utilizan para aplicar las nuevas etiquetas a los elementos relevantes en el proyecto. Este paso actualiza nuestra base de datos con las etiquetas sugeridas por la IA.
- **Reflexión en la Interfaz de Usuario:** Finalmente, las nuevas etiquetas aparecen en la vista del proyecto del usuario, completando el proceso de etiquetado automático desde la perspectiva del usuario.

Esta arquitectura ofrece varios beneficios clave que mejoran tanto el rendimiento del sistema como la experiencia del usuario. Al utilizar una cola y procesamiento en segundo plano, hemos logrado una escalabilidad impresionante, permitiéndonos manejar numerosas solicitudes simultáneamente sin abrumar nuestro sistema o alcanzar los límites de velocidad de la API de OpenAI. Implementar esta arquitectura requirió una cuidadosa consideración de varios factores para garantizar un rendimiento y confiabilidad óptimos. Para la gestión de colas, elegimos Redis, aprovechando su velocidad y confiabilidad en el manejo de colas distribuidas.

Este enfoque también contribuye a la capacidad de respuesta general de la función. Los usuarios reciben retroalimentación inmediata de que su solicitud está siendo procesada, incluso si el etiquetado real toma algo de tiempo, creando una sensación de interacción en tiempo real. La tolerancia a fallas de la arquitectura es otra ventaja crucial. Si alguna parte del proceso encuentra problemas, como interrupciones temporales de la API de OpenAI, podemos reintentar con gracia o manejar la falla sin impactar todo el sistema.

Esta robustez, combinada con la aparición en tiempo real de las etiquetas, mejora la experiencia del usuario, dando la impresión de "magia" de IA en funcionamiento.

## Datos y Prompts

Un paso crucial en nuestro proceso de etiquetado automático es preparar los datos para enviar al modelo GPT-3.5. Este paso requirió una consideración cuidadosa para equilibrar el proporcionar suficiente contexto para un etiquetado preciso mientras se mantiene la eficiencia y se protege la privacidad del usuario. Aquí hay una mirada detallada a nuestro proceso de preparación de datos.

Para cada registro, compilamos la siguiente información:

- **Nombre de Lista**: Proporciona contexto sobre la categoría o fase más amplia del proyecto.
- **Título del Registro**: A menudo contiene información clave sobre el propósito o contenido del registro.
- **Campos Personalizados**: Incluimos [campos personalizados](/platform/features/custom-fields) basados en texto y números, que a menudo contienen información crucial específica del proyecto.
- **Descripción**: Típicamente contiene la información más detallada sobre el registro.
- **Comentarios**: Pueden proporcionar contexto adicional o actualizaciones que podrían ser relevantes para el etiquetado.
- **Fecha de Vencimiento**: Información temporal que podría influir en la selección de etiquetas.

Curiosamente, no enviamos datos de etiquetas existentes a GPT-3.5, y hacemos esto para evitar sesgar el modelo.

El núcleo de nuestra función de etiquetado automático radica en cómo interactuamos con el modelo GPT-3.5 y procesamos sus respuestas. Esta sección de nuestro pipeline requirió un diseño cuidadoso para garantizar un etiquetado preciso, consistente y eficiente.

Usamos un prompt del sistema cuidadosamente elaborado para instruir a la IA sobre su tarea. Aquí hay un desglose de nuestro prompt y la justificación detrás de cada componente:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Definición de Tarea:** Establecemos claramente la tarea de la IA para garantizar respuestas enfocadas.
- **Manejo de Incertidumbre:** Permitimos explícitamente respuestas vacías, evitando el etiquetado forzado cuando la IA no está segura.
- **Etiquetas Disponibles:** Proporcionamos una lista de etiquetas válidas (${tags}) para restringir las opciones de la IA a las etiquetas existentes del proyecto.
- **Fecha Actual:** Incluir ${currentDate} ayuda a la IA a entender el contexto temporal, lo cual puede ser crucial para ciertos tipos de proyectos.
- **Formato de Respuesta:** Especificamos un formato JSON para facilitar el análisis y la verificación de errores.

Este prompt es el resultado de pruebas e iteraciones extensivas. Descubrimos que ser explícito sobre la tarea, las opciones disponibles y el formato de salida deseado mejoró significativamente la precisión y consistencia de las respuestas de la IA: ¡la simplicidad es clave!

La lista de etiquetas disponibles se genera del lado del servidor y se valida antes de incluirla en el prompt. Implementamos límites estrictos de caracteres en los nombres de etiquetas para evitar prompts de tamaño excesivo.

Como se mencionó anteriormente, no tuvimos problemas con GPT-3.5 Turbo para obtener la respuesta JSON pura en el formato correcto el 100% de las veces.

Entonces, en resumen,

- Combinamos el prompt del sistema con los datos preparados del registro.
- Este prompt combinado se envía al modelo GPT-3.5 a través de la API de OpenAI.
- Usamos una configuración de temperatura de 0.3 para equilibrar la creatividad y la consistencia en las respuestas de la IA.
- Nuestra llamada API incluye un parámetro max_tokens para limitar el tamaño de la respuesta y controlar los costos.

Una vez que recibimos la respuesta de la IA, pasamos por varios pasos para procesar y aplicar las etiquetas sugeridas:

* **Análisis JSON**: Intentamos analizar la respuesta como JSON. Si el análisis falla, registramos el error y omitimos el etiquetado para ese registro.
* **Validación de Esquema**: Validamos el JSON analizado contra nuestro esquema esperado (un objeto con un array "tags"). Esto captura cualquier desviación estructural en la respuesta de la IA.
* **Validación de Etiquetas**: Cruzamos las etiquetas sugeridas con nuestra lista de etiquetas válidas del proyecto. Este paso filtra cualquier etiqueta que no exista en el proyecto, lo cual podría ocurrir si la IA malinterpretó o si las etiquetas del proyecto cambiaron entre la creación del prompt y el procesamiento de la respuesta.
* **Deduplicación**: Eliminamos cualquier etiqueta duplicada de la sugerencia de la IA para evitar el etiquetado redundante.
* **Aplicación**: Las etiquetas validadas y deduplicadas se aplican luego al registro en nuestra base de datos.
* **Registro y Análisis**: Registramos las etiquetas finales aplicadas. Estos datos son valiosos para monitorear el rendimiento del sistema y mejorarlo con el tiempo.

## Desafíos

Implementar el etiquetado automático impulsado por IA en Blue presentó varios desafíos únicos, cada uno requiriendo soluciones innovadoras para garantizar una función robusta, eficiente y fácil de usar.

### Deshacer Operación Masiva

La función de Etiquetado con IA se puede realizar tanto en registros individuales como en masa. El problema con la operación masiva es que si al usuario no le gusta el resultado, tendría que revisar manualmente miles de registros y deshacer el trabajo de la IA. Claramente, eso es inaceptable.

Para resolver esto, implementamos un sistema innovador de sesión de etiquetado. A cada operación de etiquetado masivo se le asigna un ID de sesión único, que se asocia con todas las etiquetas aplicadas durante esa sesión. Esto nos permite gestionar eficientemente las operaciones de deshacer simplemente eliminando todas las etiquetas asociadas con un ID de sesión particular. También eliminamos los rastros de auditoría relacionados, asegurando que las operaciones deshechas no dejen rastro en el sistema. Este enfoque les da a los usuarios la confianza para experimentar con el etiquetado con IA, sabiendo que pueden revertir fácilmente los cambios si es necesario.

### Privacidad de Datos

La privacidad de datos fue otro desafío crítico que enfrentamos.

Nuestros usuarios confían en nosotros con los datos de sus proyectos, y era primordial garantizar que esta información no fuera retenida o utilizada para el entrenamiento de modelos por OpenAI. Abordamos esto en múltiples frentes.

Primero, formamos un acuerdo con OpenAI que prohíbe explícitamente el uso de nuestros datos para el entrenamiento de modelos. Además, OpenAI elimina los datos después del procesamiento, proporcionando una capa adicional de protección de privacidad.

Por nuestra parte, tomamos la precaución de excluir información sensible, como detalles del asignado, de los datos enviados a la IA, así que esto garantiza que los nombres específicos de individuos no se envíen a terceros junto con otros datos.

Este enfoque integral nos permite aprovechar las capacidades de IA mientras mantenemos los más altos estándares de privacidad y seguridad de datos.

### Límites de Velocidad y Captura de Errores

Una de nuestras principales preocupaciones era la escalabilidad y la limitación de velocidad. Las llamadas API directas a OpenAI para cada registro habrían sido ineficientes y podrían alcanzar rápidamente los límites de velocidad, especialmente para proyectos grandes o durante los momentos de mayor uso. Para abordar esto, desarrollamos una arquitectura de servicio en segundo plano que nos permite agrupar solicitudes e implementar nuestro propio sistema de colas. Este enfoque nos ayuda a gestionar la frecuencia de llamadas API y permite un procesamiento más eficiente de grandes volúmenes de registros, garantizando un rendimiento fluido incluso bajo carga pesada.

La naturaleza de las interacciones con IA significaba que también teníamos que prepararnos para errores ocasionales o salidas inesperadas. Hubo instancias donde la IA podría producir JSON inválido o salidas que no coincidían con nuestro formato esperado. Para manejar esto, implementamos un manejo robusto de errores y lógica de análisis en todo nuestro sistema. Si la respuesta de la IA no es JSON válido o no contiene la clave "tags" esperada, nuestro sistema está diseñado para tratarlo como si no se sugirieran etiquetas, en lugar de intentar procesar datos potencialmente corruptos. Esto garantiza que incluso ante la imprevisibilidad de la IA, nuestro sistema permanezca estable y confiable.

## Desarrollos Futuros

Creemos que las funciones, y el producto Blue en su conjunto, nunca está "terminado": siempre hay espacio para mejorar.

Hubo algunas funciones que consideramos en la construcción inicial que no pasaron la fase de alcance, pero son interesantes de notar ya que probablemente implementaremos alguna versión de ellas en el futuro.

La primera es agregar descripción de etiquetas. Esto permitiría a los usuarios finales no solo darle a las etiquetas un nombre y un color, sino también una descripción opcional. Esto también se pasaría a la IA para ayudar a proporcionar más contexto y potencialmente mejorar la precisión.

Si bien el contexto adicional podría ser valioso, somos conscientes de la complejidad potencial que podría introducir. Hay un equilibrio delicado que lograr entre proporcionar información útil y abrumar a los usuarios con demasiados detalles. A medida que desarrollamos esta función, nos enfocaremos en encontrar ese punto ideal donde el contexto agregado mejore en lugar de complicar la experiencia del usuario.

Quizás la mejora más emocionante en nuestro horizonte es la integración del etiquetado automático con IA con nuestro [sistema de automatización de gestión de proyectos](/platform/features/automations).

Esto significa que la función de etiquetado con IA podría ser un disparador o una acción de una automatización. Esto podría ser enorme ya que podría convertir esta función de categorización de IA bastante simple en un sistema de enrutamiento basado en IA para el trabajo.

Imagine una automatización que establece:

Cuando la IA etiqueta un registro como "Crítico" → Asignar a "Gerente" y Enviar un Correo Electrónico Personalizado

Esto significa que cuando usted etiqueta un registro con IA, si la IA decide que es un problema crítico, entonces puede asignar automáticamente al gerente del proyecto y enviarle un correo electrónico personalizado. Esto extiende los [beneficios de nuestro sistema de automatización de gestión de proyectos](/platform/features/automations) de un sistema puramente basado en reglas a un verdadero sistema flexible de IA.

Al explorar continuamente las fronteras de la IA en la gestión de proyectos, nuestro objetivo es proporcionar a nuestros usuarios herramientas que no solo satisfagan sus necesidades actuales sino que anticipen y den forma al futuro del trabajo. Como siempre, desarrollaremos estas funciones en estrecha colaboración con nuestra comunidad de usuarios, asegurando que cada mejora agregue valor real y práctico al proceso de gestión de proyectos.

## Conclusión

¡Eso es todo!

Esta fue una función divertida de implementar, y uno de nuestros primeros pasos en IA, junto con el [Resumen de Contenido con IA](/insights/ai-content-summarization) que lanzamos anteriormente. Sabemos que la IA va a desempeñar un papel cada vez más importante en la gestión de proyectos en el futuro, y no podemos esperar para lanzar más funciones innovadoras aprovechando LLMs avanzados (Modelos de Lenguaje Grande).

Hubo bastante en qué pensar mientras implementábamos esto, y estamos especialmente emocionados sobre cómo podemos aprovechar esta función en el futuro con el [motor de automatización de gestión de proyectos](/insights/benefits-project-management-automation) existente de Blue.

También esperamos que haya sido una lectura interesante, y que le dé una visión de cómo pensamos sobre la ingeniería de las funciones que usted usa todos los días.
