---
title: Preguntas frecuentes sobre la seguridad de Blue
description: Esta es una lista de las preguntas más frecuentes sobre los protocolos y prácticas de seguridad en Blue.
category: "FAQ"
date: 2024-07-19
---


Nuestra misión es organizar el trabajo del mundo construyendo la mejor plataforma de gestión de proyectos del planeta.

Central para lograr esta misión es asegurar que nuestra plataforma sea segura, confiable y digna de confianza. Entendemos que para ser su única fuente de verdad, Blue debe proteger sus datos comerciales sensibles contra amenazas externas, pérdida de datos y tiempo de inactividad.

Esto significa que tomamos la seguridad en serio en Blue.

Cuando pensamos en seguridad, consideramos un enfoque holístico que se centra en tres áreas clave:

1.  **Seguridad de Infraestructura y Red**: Asegura que nuestros sistemas físicos y virtuales estén protegidos contra amenazas externas y accesos no autorizados.
2.  **Seguridad del Software**: Se centra en la seguridad del código en sí, incluyendo prácticas de codificación segura, revisiones de código regulares y gestión de vulnerabilidades.
3.  **Seguridad de la Plataforma**: Incluye las características dentro de Blue, como [controles de acceso sofisticados](/platform/features/user-permissions), asegurando que los proyectos sean privados por defecto, y otras medidas para proteger los datos y la privacidad del usuario.

## ¿Qué tan escalable es Blue?

Esta es una pregunta importante, ya que desea un sistema que pueda *crecer* con usted. No querrá tener que cambiar su plataforma de gestión de proyectos y procesos en seis o doce meses.

Elegimos proveedores de plataformas con cuidado, para asegurarnos de que puedan manejar las cargas de trabajo exigentes de nuestros clientes. Utilizamos servicios en la nube de algunos de los principales proveedores de la nube del mundo que impulsan empresas como [Spotify](https://spotify.com) y [Netflix](https://netflix.com), que tienen varios órdenes de magnitud más tráfico que nosotros.

Los principales proveedores de la nube que utilizamos son:

- **[Cloudflare](https://cloudflare.com)**: Gestionamos nuestro DNS (Servicio de Nombre de Dominio) a través de Cloudflare, así como nuestro sitio web de marketing que funciona en [Cloudflare Pages](https://pages.cloudflare.com/).
- **[Amazon Web Services](https://aws.amazon.com/)**: Utilizamos AWS para nuestra base de datos, que es [Aurora](https://aws.amazon.com/rds/aurora/), para almacenamiento de archivos a través de [Simple Storage Service (S3)](https://aws.amazon.com/s3/), y también para enviar correos electrónicos a través de [Simple Email Service (SES)](https://aws.amazon.com/ses/).
- **[Render](https://render.com)**: Utilizamos Render para nuestros servidores front-end, servidores de aplicaciones/API, nuestros servicios en segundo plano, sistema de colas y base de datos Redis. Curiosamente, Render está construido *sobre* AWS.

## ¿Qué tan seguros son los archivos en Blue?

Comencemos con el almacenamiento de datos. Nuestros archivos están alojados en [AWS S3](https://aws.amazon.com/s3/), que es el almacenamiento de objetos en la nube más popular del mundo, con escalabilidad, disponibilidad de datos, seguridad y rendimiento líderes en la industria.

Tenemos un 99.99% de disponibilidad de archivos y un 99.999999999% de alta durabilidad.

Desglosamos lo que esto significa.

La disponibilidad se refiere a la cantidad de tiempo que los datos están operativos y accesibles. La disponibilidad de archivos del 99.99% significa que podemos esperar que los archivos no estén disponibles por no más de aproximadamente 8.76 horas al año.

La durabilidad se refiere a la probabilidad de que los datos permanezcan intactos y no corruptos con el tiempo. Este nivel de durabilidad significa que podemos esperar no perder más de un archivo de cada 10 mil millones de archivos subidos, gracias a la extensa redundancia y replicación de datos en múltiples centros de datos.

Utilizamos [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/) para mover automáticamente archivos a diferentes clases de almacenamiento según la frecuencia de acceso. Basándonos en los patrones de actividad de cientos de miles de proyectos, notamos que la mayoría de los archivos se acceden en un patrón que se asemeja a una curva de retroceso exponencial. Esto significa que la mayoría de los archivos se acceden con mucha frecuencia en los primeros días, y luego se accede a ellos cada vez con menos frecuencia. Esto nos permite mover archivos más antiguos a un almacenamiento más lento, pero significativamente más barato, sin afectar la experiencia del usuario de manera significativa.

Los ahorros de costos por esto son significativos. S3 Standard-Infrequent Access (S3 Standard-IA) es aproximadamente 1.84 veces más barato que S3 Standard. Esto significa que por cada dólar que habríamos gastado en S3 Standard, solo gastamos alrededor de 54 centavos en S3 Standard-IA por la misma cantidad de datos almacenados.

| Característica           | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Costo de Almacenamiento   | $0.023 - $0.021 por GB  | $0.0125 por GB        |
| Costo de Solicitud (PUT, etc.) | $0.005 por 1,000 solicitudes | $0.01 por 1,000 solicitudes |
| Costo de Solicitud (GET) | $0.0004 por 1,000 solicitudes | $0.001 por 1,000 solicitudes |
| Costo de Recuperación de Datos | $0.00 por GB            | $0.01 por GB          |

Los archivos que sube a través de Blue están encriptados tanto en tránsito como en reposo. Los datos transferidos hacia y desde Amazon S3 están asegurados utilizando [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), protegiendo contra [intercepciones](https://es.wikipedia.org/wiki/Eavesdropping) y [ataques de intermediario](https://es.wikipedia.org/wiki/Man-in-the-middle_attack). Para la encriptación en reposo, Amazon S3 utiliza Server-Side Encryption (SSE-S3), que encripta automáticamente todas las nuevas cargas con cifrado AES-256, con Amazon gestionando las claves de cifrado. Esto asegura que sus datos permanezcan seguros durante todo su ciclo de vida.

## ¿Qué pasa con los datos que no son archivos?

Nuestra base de datos está impulsada por [AWS Aurora](https://aws.amazon.com/rds/aurora/), un servicio de base de datos relacional moderno que garantiza un alto rendimiento, disponibilidad y seguridad para sus datos.

Los datos en Aurora están encriptados tanto en tránsito como en reposo. Utilizamos SSL (AES-256) para asegurar las conexiones entre su instancia de base de datos y su aplicación, protegiendo los datos durante la transferencia. Para la encriptación en reposo, Aurora utiliza claves gestionadas a través del Servicio de Gestión de Claves de AWS (KMS), asegurando que todos los datos almacenados, incluidos los respaldos automáticos, instantáneas y réplicas, estén encriptados y protegidos.

Aurora cuenta con un sistema de almacenamiento distribuido, tolerante a fallos y auto-reparador. Este sistema está desacoplado de los recursos de cómputo y puede escalar automáticamente hasta 128 TiB por instancia de base de datos. Los datos se replican en tres [Zonas de Disponibilidad](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZs), proporcionando resiliencia contra la pérdida de datos y asegurando alta disponibilidad. En caso de un fallo en la base de datos, Aurora reduce los tiempos de recuperación a menos de 60 segundos, asegurando una mínima interrupción.

Blue respalda continuamente nuestra base de datos en Amazon S3, permitiendo la recuperación en un punto en el tiempo. Esto significa que podemos restaurar la base de datos maestra de Blue a cualquier momento específico dentro de los últimos cinco minutos, asegurando que sus datos siempre sean recuperables. También tomamos instantáneas regulares de la base de datos para períodos de retención de respaldo más largos.

Como un servicio completamente gestionado, Aurora automatiza tareas administrativas que consumen tiempo, como aprovisionamiento de hardware, configuración de bases de datos, parches y respaldos. Esto reduce la carga operativa y asegura que nuestra base de datos esté siempre actualizada con los últimos parches de seguridad y mejoras de rendimiento.

Si somos más eficientes, podemos trasladar nuestros ahorros de costos a nuestros clientes con nuestra [tarificación líder en la industria](/pricing).

Aurora cumple con varios estándares de la industria, como HIPAA, GDPR y SOC 2, asegurando que sus prácticas de gestión de datos cumplan con estrictos requisitos regulatorios. Auditorías de seguridad regulares e integración con [Amazon GuardDuty](https://aws.amazon.com/guardduty/) ayudan a detectar y mitigar posibles amenazas de seguridad.

## ¿Cómo asegura Blue la seguridad de inicio de sesión?

Blue utiliza [enlaces mágicos a través de correo electrónico](https://documentation.blue.cc/user-management/magic-links) para proporcionar acceso seguro y conveniente a su cuenta, eliminando la necesidad de contraseñas tradicionales.

Este enfoque mejora significativamente la seguridad al mitigar amenazas comunes asociadas con inicios de sesión basados en contraseñas. Al eliminar las contraseñas, los enlaces mágicos protegen contra ataques de phishing y robo de contraseñas, *ya que no hay contraseña que robar o explotar.*

Cada enlace mágico es válido solo para una sesión de inicio de sesión, reduciendo el riesgo de acceso no autorizado. Además, estos enlaces caducan después de 15 minutos, asegurando que cualquier enlace no utilizado no pueda ser explotado, mejorando aún más la seguridad.

La conveniencia que ofrecen los enlaces mágicos también es notable. Los enlaces mágicos proporcionan una experiencia de inicio de sesión sin complicaciones, permitiéndole acceder a su cuenta *sin* la necesidad de recordar contraseñas complejas.

Esto no solo simplifica el proceso de inicio de sesión, sino que también previene brechas de seguridad que ocurren cuando las contraseñas se reutilizan en múltiples servicios. Muchos usuarios tienden a usar la misma contraseña en varias plataformas, lo que significa que una brecha de seguridad en un servicio podría comprometer sus cuentas en otros servicios, incluido Blue. Al usar enlaces mágicos, la seguridad de Blue no depende de las prácticas de seguridad de otros servicios, proporcionando una capa de protección más robusta e independiente para nuestros usuarios.

Cuando solicita iniciar sesión en su cuenta de Blue, se envía una URL de inicio de sesión única a su correo electrónico. Hacer clic en este enlace lo registrará instantáneamente en su cuenta. El enlace está diseñado para caducar después de un solo uso o después de 15 minutos, lo que ocurra primero, añadiendo una capa adicional de seguridad. Al usar enlaces mágicos, Blue asegura que su proceso de inicio de sesión sea tanto seguro como fácil de usar, proporcionando tranquilidad y conveniencia.

## ¿Cómo puedo verificar la confiabilidad y el tiempo de actividad de Blue?

En Blue, estamos comprometidos a mantener un alto nivel de confiabilidad y transparencia para nuestros usuarios. Para proporcionar visibilidad sobre el rendimiento de nuestra plataforma, ofrecemos una [página de estado del sistema dedicada](https://status.blue.cc) que también está vinculada desde nuestro pie de página en cada página de nuestro sitio web.

![](/insights/status-page.png)

Esta página muestra nuestros datos históricos de tiempo de actividad, lo que le permite ver cuán consistentemente nuestros servicios han estado disponibles a lo largo del tiempo. Además, la página de estado incluye informes de incidentes detallados, proporcionando transparencia sobre cualquier problema pasado, su impacto y los pasos que hemos tomado para resolverlos y prevenir futuras ocurrencias.