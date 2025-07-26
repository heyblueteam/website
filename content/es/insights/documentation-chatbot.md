---
title: Por qué construimos nuestro propio chatbot de documentación con IA
description: Creamos nuestro propio chatbot de documentación con IA que está entrenado en la documentación de la plataforma Blue.
category: "Product Updates"
date: 2024-07-09
---


En Blue, siempre estamos buscando formas de facilitar la vida a nuestros clientes. Tenemos [documentación detallada de cada función](https://documentation.blue.cc), [videos en YouTube](https://www.youtube.com/@HeyBlueTeam), [Consejos y trucos](/insights/tips-tricks) y [varios canales de soporte](/support).

Hemos estado observando de cerca el desarrollo de la IA (Inteligencia Artificial) ya que estamos muy interesados en [automatizaciones de gestión de proyectos](/platform/features/automations). También lanzamos funciones como [Clasificación Automática con IA](/insights/ai-auto-categorization) y [Resúmenes con IA](/insights/ai-content-summarization) para facilitar el trabajo a nuestros miles de clientes.

Una cosa que está clara es que la IA ha llegado para quedarse, y tendrá un efecto increíble en la mayoría de las industrias, y la gestión de proyectos no es la excepción. Así que nos preguntamos cómo podríamos aprovechar aún más la IA para ayudar en todo el ciclo de vida de un cliente, desde el descubrimiento, las preventas, la incorporación y también con preguntas continuas.

La respuesta fue bastante clara: **Necesitábamos un chatbot de IA entrenado en nuestra documentación.**

Seamos realistas: *cada* organización probablemente debería tener un chatbot. Son una excelente manera para que los clientes obtengan respuestas instantáneas a preguntas típicas, sin tener que buscar en páginas de documentación densa o en su sitio web. La importancia de los chatbots en los sitios web de marketing modernos no puede ser subestimada.

![](/insights/ai-chatbot-regular.png)

Para las empresas de software específicamente, no se debe considerar el sitio web de marketing como una "cosa" separada: *es* parte de su producto. Esto se debe a que encaja en la vida típica del cliente:

- **Conciencia** (Descubrimiento): Aquí es donde los clientes potenciales se topan por primera vez con su increíble producto. Su chatbot puede ser su guía amigable, señalando características y beneficios clave desde el principio.
- **Consideración** (Educación): Ahora tienen curiosidad y quieren aprender más. Su chatbot se convierte en su tutor personal, proporcionando información adaptada a sus necesidades y preguntas específicas.
- **Compra/Conversión**: Este es el momento de la verdad: cuando un prospecto decide dar el paso y convertirse en cliente. Su chatbot puede suavizar cualquier contratiempo de última hora, responder esas preguntas "justo antes de comprar" y tal vez incluso ofrecer un buen trato para cerrar el trato.
- **Incorporación**: Han comprado, ¿y ahora qué? Su chatbot se transforma en un compañero útil, guiando a los nuevos usuarios a través de la configuración, mostrándoles cómo funciona y asegurándose de que no se sientan perdidos en el mundo de su producto.
- **Retención**: Mantener a los clientes felices es el nombre del juego. Su chatbot está disponible 24/7, listo para resolver problemas, ofrecer consejos y trucos, y asegurarse de que sus clientes sientan el cariño.
- **Expansión**: ¡Es hora de subir de nivel! Su chatbot puede sugerir sutilmente nuevas características, ventas adicionales o ventas cruzadas que se alineen con la forma en que el cliente ya está utilizando su producto. Es como tener un vendedor realmente inteligente y no insistente siempre a la espera.
- **Defensa**: Los clientes felices se convierten en sus mayores animadores. Su chatbot puede alentar a los usuarios satisfechos a correr la voz, dejar reseñas o participar en programas de referencia. ¡Es como tener una máquina de entusiasmo integrada directamente en su producto!

## Decisión de Construir vs Comprar

Una vez que decidimos implementar un chatbot de IA, la siguiente gran pregunta fue: ¿construir o comprar? Como un equipo pequeño enfocado en nuestro producto principal, generalmente preferimos soluciones "como servicio" o plataformas de código abierto populares. Después de todo, no estamos en el negocio de reinventar la rueda para cada parte de nuestra pila tecnológica. 
Así que nos arremangamos y nos sumergimos en el mercado, buscando tanto soluciones de chatbot de IA de pago como de código abierto.

Nuestros requisitos eran sencillos, pero no negociables:

- **Experiencia sin Marca**: Este chatbot no es solo un widget bonito; se integrará en nuestro sitio web de marketing y eventualmente en nuestro producto. No estamos interesados en publicitar la marca de otra persona en nuestro propio espacio digital.
- **Gran UX**: Para muchos clientes potenciales, este chatbot podría ser su primer punto de contacto con Blue. Establece el tono para su percepción de nuestra empresa. Seamos realistas: si no podemos implementar un chatbot adecuado en nuestro sitio web, ¿cómo podemos esperar que los clientes confíen en nosotros con sus proyectos y procesos críticos?
- **Costo Razonable**: Con una gran base de usuarios y planes para integrar el chatbot en nuestro producto principal, necesitábamos una solución que no rompiera el banco a medida que el uso aumentara. Idealmente, queríamos una **opción BYOK (Bring Your Own Key)**. Esto nos permitiría usar nuestra propia clave de servicio de OpenAI u otro servicio de IA, pagando solo los costos variables directos en lugar de un recargo a un proveedor externo que no ejecuta realmente los modelos.
- **Compatible con OpenAI Assistants API**: Si íbamos a optar por un software de código abierto, no queríamos tener la molestia de gestionar una canalización para la ingestión de documentos, indexación, bases de datos vectoriales y todo eso. Queríamos usar la [OpenAI Assistants API](https://platform.openai.com/docs/assistants/overview) que abstraería toda la complejidad detrás de una API. Honestamente, esto está muy bien hecho.
- **Escalabilidad**: Queremos tener este chatbot en múltiples lugares, con potencialmente decenas de miles de usuarios al año. Esperamos un uso significativo y no queremos estar atados a una solución que no pueda escalar con nuestras necesidades.

## Chatbots de IA Comerciales

Los que revisamos tendían a tener una mejor UX que las soluciones de código abierto, como suele ser el caso. Probablemente haya una discusión separada que se deba tener algún día sobre *por qué* muchas soluciones de código abierto ignoran o minimizan la importancia de la UX.

Proporcionaremos una lista aquí, en caso de que esté buscando algunas ofertas comerciales sólidas:

- **[Chatbase](https://chatbase.co):** Chatbase le permite construir un chatbot de IA personalizado entrenado en su base de conocimientos y agregarlo a su sitio web o interactuar con él a través de su API. Ofrece características como respuestas confiables, generación de leads, análisis avanzados y la capacidad de conectarse a múltiples fuentes de datos. Para nosotros, esto se sintió como una de las ofertas comerciales más pulidas que existen.
- **[DocsBot AI](https://docsbot.ai/):** DocsBot AI crea bots de ChatGPT personalizados entrenados en su documentación y contenido para soporte, preventas, investigación y más. Proporciona widgets embebibles para agregar fácilmente el chatbot a su sitio web, la capacidad de responder automáticamente a tickets de soporte y una poderosa API para integración.
- **[CustomGPT.ai](https://customgpt.ai):** CustomGPT.ai crea una experiencia de chatbot personal al ingerir sus datos comerciales, incluidos contenido del sitio web, helpdesk, bases de conocimientos, documentos y más. Permite a los leads hacer preguntas y obtener respuestas instantáneas basadas en su contenido, sin necesidad de buscar. Curiosamente, también [afirman superar a OpenAI en RAG (Generación Aumentada por Recuperación)!](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Esta es una oferta comercial interesante, porque *también* es software de código abierto. Parece estar en una etapa temprana, y el precio no se siente realista ($27/mes por mensajes ilimitados nunca funcionará comercialmente para ellos).

También miramos [InterCom Fin](https://www.intercom.com/fin) que es parte de su software de soporte al cliente. Esto habría significado cambiar de [HelpScout](https://wwww.helpscout.com) que hemos utilizado desde que comenzamos Blue. Esto podría haber sido posible, pero InterCom Fin tiene algunos precios locos que simplemente lo excluyeron de consideración.

Y este es en realidad el problema con muchas de las ofertas comerciales. InterCom Fin cobra $0.99 por cada solicitud de soporte al cliente manejada, y ChatBase cobra $399/mes por 40,000 mensajes. Eso es casi $5k al año por un simple widget de chat.

Considerando que los precios de la inferencia de IA están cayendo como locos. OpenAI redujo sus precios de manera bastante dramática:

- El GPT-4 original (contexto de 8k) estaba a $0.03 por 1K tokens de prompt.
- El GPT-4 Turbo (contexto de 128k) estaba a $0.01 por 1K tokens de prompt, una reducción del 50% respecto al GPT-4 original.
- El modelo GPT-4o está a $0.005 por 1K tokens, lo que representa una reducción adicional del 50% respecto al precio del GPT-4 Turbo.

Eso es una reducción del 83% en costos, y no esperamos que se mantenga estancado.

Considerando que estábamos buscando una solución escalable que sería utilizada por decenas de miles de usuarios al año con una cantidad significativa de mensajes, tiene sentido ir directamente a la fuente y pagar los costos de la API directamente, no usar una versión comercial que aumenta los costos.

## Chatbots de IA de Código Abierto

Como se mencionó, las opciones de código abierto que revisamos fueron en su mayoría decepcionantes en cuanto al requisito de "Gran UX".

Miramos:

- **[Deepchat](https://deepchat.dev/)**: Este es un componente de chat independiente del marco para servicios de IA, que se conecta a varias API de IA, incluida OpenAI. También tiene la capacidad de que los usuarios descarguen un modelo de IA que se ejecuta directamente en el navegador. Jugamos con esto y logramos hacer que una versión funcionara, pero la API de OpenAI Assistants implementada se sintió bastante defectuosa con varios problemas. Sin embargo, este es un proyecto muy prometedor, y su playground está muy bien hecho.
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Mirando esto nuevamente desde una perspectiva de código abierto, esto requeriría que levantáramos bastante infraestructura, algo que no queríamos hacer, porque queríamos depender lo más posible de la API de Assistants de OpenAI.

## Construyendo Nuestro Propio ChatBot

Y así, sin poder encontrar algo que coincidiera con todos nuestros requisitos, decidimos construir nuestro propio chatbot de IA que pudiera interactuar con la API de OpenAI Assistants. ¡Esto resultó ser relativamente indoloro!

Nuestro sitio web utiliza [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (que es el mismo marco que la Plataforma Blue), y [TailwindUI](https://tailwindui.com/).

El primer paso fue crear una API (Interfaz de Programación de Aplicaciones) en Nuxt3 que pueda "hablar" con la API de OpenAI Assistants. Esto era necesario ya que no queríamos hacer todo en el front-end, ya que esto expondría nuestra clave de API de OpenAI al mundo, con el potencial de abuso.

Nuestra API de backend actúa como un intermediario seguro entre el navegador del usuario y OpenAI. Esto es lo que hace:

- **Gestión de Conversaciones:** Crea y gestiona "hilos" para cada conversación. Piensa en un hilo como una sesión de chat única que recuerda todo lo que has dicho.
- **Manejo de Mensajes:** Cuando envías un mensaje, nuestra API lo agrega al hilo correcto y le pide al asistente de OpenAI que elabore una respuesta.
- **Espera Inteligente:** En lugar de hacerte mirar una pantalla de carga, nuestra API consulta a OpenAI cada segundo para ver si tu respuesta está lista. Es como tener un camarero que mantiene un ojo en tu pedido sin molestar al chef cada dos segundos.
- **Seguridad Primero:** Al manejar todo esto en el servidor, mantenemos tus datos y nuestras claves de API a salvo.

Luego, estaba el front-end y la experiencia del usuario. Como se discutió anteriormente, esto era *críticamente* importante, porque no tenemos una segunda oportunidad para causar una buena primera impresión.

Al diseñar nuestro chatbot, prestamos meticulosa atención a la experiencia del usuario, asegurándonos de que cada interacción sea fluida, intuitiva y refleje el compromiso de Blue con la calidad. La interfaz del chatbot comienza con un simple y elegante círculo azul, utilizando [HeroIcons para nuestros íconos](https://heroicons.com/) (que usamos en todo el sitio web de Blue) para actuar como nuestro widget de apertura del chatbot. Esta elección de diseño asegura consistencia visual y reconocimiento inmediato de la marca.

![](/insights/ai-chatbot-circle.png)

Entendemos que a veces los usuarios pueden necesitar soporte adicional o información más detallada. Por eso hemos incluido enlaces convenientes dentro de la interfaz del chatbot. Un enlace de correo electrónico para soporte está fácilmente disponible, permitiendo a los usuarios comunicarse directamente con nuestro equipo si necesitan asistencia más personalizada. Además, hemos incorporado un enlace a la documentación, proporcionando fácil acceso a recursos más completos para aquellos que deseen profundizar en las ofertas de Blue.

La experiencia del usuario se mejora aún más con animaciones de desvanecimiento sutiles al abrir la ventana del chatbot. Estas animaciones sutiles añaden un toque de sofisticación a la interfaz, haciendo que la interacción se sienta más dinámica y atractiva. También hemos implementado un indicador de escritura, una pequeña pero crucial característica que permite a los usuarios saber que el chatbot está procesando su consulta y elaborando una respuesta. Esta señal visual ayuda a gestionar las expectativas del usuario y mantiene un sentido de comunicación activa.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>

Reconociendo que algunas conversaciones pueden requerir más espacio en la pantalla, hemos añadido la capacidad de abrir la conversación en una ventana más grande. Esta característica es particularmente útil para intercambios más largos o al revisar información detallada, dando a los usuarios la flexibilidad de adaptar el chatbot a sus necesidades.

Detrás de escena, hemos implementado un procesamiento inteligente para optimizar las respuestas del chatbot. Nuestro sistema analiza automáticamente las respuestas de la IA para eliminar referencias a nuestros documentos internos, asegurando que la información presentada sea clara, relevante y centrada únicamente en abordar la consulta del usuario. 
Para mejorar la legibilidad y permitir una comunicación más matizada, hemos incorporado soporte para markdown utilizando la biblioteca 'marked'. Esta característica permite que nuestra IA proporcione texto ricamente formateado, incluyendo énfasis en negrita y cursiva, listas estructuradas e incluso fragmentos de código cuando sea necesario. Es como recibir un mini-documento bien formateado y adaptado en respuesta a tus preguntas.

Por último, pero no menos importante, hemos priorizado la seguridad en nuestra implementación. Usando la biblioteca DOMPurify, sanitizamos el HTML generado a partir del análisis de markdown. Este paso crucial asegura que cualquier script o código potencialmente dañino sea eliminado antes de que el contenido se muestre. Es nuestra forma de garantizar que la información útil que recibes no solo sea informativa, sino también segura de consumir.

## Desarrollos Futuros

Así que esto es solo el comienzo, tenemos algunas cosas emocionantes en la hoja de ruta para esta función.

Una de nuestras próximas características es la capacidad de transmitir respuestas en tiempo real. Pronto, verás las respuestas del chatbot aparecer carácter por carácter, haciendo que las conversaciones se sientan más naturales y dinámicas. Es como ver a la IA pensar, creando una experiencia más atractiva e interactiva que te mantiene informado en cada paso del camino.

Para nuestros valiosos usuarios de Blue, estamos trabajando en la personalización. El chatbot reconocerá cuando estés conectado, adaptando sus respuestas según tu información de cuenta, historial de uso y preferencias. Imagina un chatbot que no solo responde tus preguntas, sino que entiende tu contexto específico dentro del ecosistema de Blue, proporcionando asistencia más relevante y personalizada.

Entendemos que podrías estar trabajando en múltiples proyectos o tener varias consultas. Por eso estamos desarrollando la capacidad de mantener varios hilos de conversación distintos con nuestro chatbot. Esta característica te permitirá cambiar entre diferentes temas sin perder el contexto, como tener múltiples pestañas abiertas en tu navegador.

Para hacer tus interacciones aún más productivas, estamos creando una función que ofrecerá preguntas de seguimiento sugeridas basadas en tu conversación actual. Esto te ayudará a explorar temas más a fondo y descubrir información relacionada que quizás no habías pensado en preguntar, haciendo que cada sesión de chat sea más completa y valiosa.

También estamos emocionados por crear un conjunto de asistentes de IA especializados, cada uno adaptado a necesidades específicas. Ya sea que estés buscando responder preguntas de preventa, configurar un nuevo proyecto o solucionar características avanzadas, podrás elegir el asistente que mejor se adapte a tus necesidades actuales. Es como tener un equipo de expertos de Blue al alcance de tu mano, cada uno especializado en diferentes aspectos de nuestra plataforma.

Por último, estamos trabajando en permitirte subir capturas de pantalla directamente al chat. La IA analizará la imagen y proporcionará explicaciones o pasos de solución de problemas según lo que vea. Esta característica hará que sea más fácil que nunca obtener ayuda con problemas específicos que encuentres mientras usas Blue, cerrando la brecha entre la información visual y la asistencia textual.

## Conclusión

Esperamos que esta profunda inmersión en nuestro proceso de desarrollo de chatbot de IA haya proporcionado algunas ideas valiosas sobre nuestro pensamiento en el desarrollo de productos en Blue. Nuestro viaje desde identificar la necesidad de un chatbot hasta construir nuestra propia solución muestra cómo abordamos la toma de decisiones y la innovación.

![](/insights/ai-chatbot-modal.png)

En Blue, sopesamos cuidadosamente las opciones de construir frente a comprar, siempre con un ojo en lo que mejor servirá a nuestros usuarios y se alineará con nuestros objetivos a largo plazo. En este caso, identificamos una brecha significativa en el mercado para un chatbot rentable pero visualmente atractivo que pudiera satisfacer nuestras necesidades específicas. Si bien generalmente abogamos por aprovechar soluciones existentes en lugar de reinventar la rueda, a veces el mejor camino a seguir es crear algo adaptado a tus requisitos únicos.

Nuestra decisión de construir nuestro propio chatbot no se tomó a la ligera. Fue el resultado de una investigación de mercado exhaustiva, una comprensión clara de nuestras necesidades y un compromiso de proporcionar la mejor experiencia posible para nuestros usuarios. Al desarrollar internamente, pudimos crear una solución que no solo satisface nuestras necesidades actuales, sino que también sienta las bases para futuras mejoras e integraciones.

Este proyecto ejemplifica nuestro enfoque en Blue: no tenemos miedo de arremangarnos y construir algo desde cero cuando es la opción correcta para nuestro producto y nuestros usuarios. Es esta disposición a ir más allá lo que nos permite ofrecer soluciones innovadoras que realmente satisfacen las necesidades de nuestros clientes.
Estamos emocionados por el futuro de nuestro chatbot de IA y el valor que aportará tanto a los usuarios potenciales como a los existentes de Blue. A medida que continuamos refinando y expandiendo sus capacidades, seguimos comprometidos a empujar los límites de lo que es posible en la gestión de proyectos y la interacción con el cliente.

Gracias por acompañarnos en este viaje a través de nuestro proceso de desarrollo. Esperamos que te haya dado un vistazo a la atención al usuario que tomamos en cada aspecto de Blue. Mantente atento a más actualizaciones a medida que continuamos evolucionando y mejorando nuestra plataforma para servirte mejor.

Si estás interesado, puedes encontrar el enlace al código fuente de este proyecto aquí:

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)**: Este es un componente de Vue que alimenta el widget de chat en sí.
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)**: Este es el middleware que funciona entre el componente de chat y la API de OpenAI Assistants.