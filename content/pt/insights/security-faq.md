---
title: FAQ sobre Segurança do Blue
description: Esta é uma lista das perguntas mais frequentes sobre os protocolos e práticas de segurança no Blue.
category: "FAQ"
date: 2024-07-19
---



Nossa missão é organizar o trabalho do mundo construindo a melhor plataforma de gerenciamento de projetos do planeta.

Central para alcançar essa missão é garantir que nossa plataforma seja segura, confiável e digna de confiança. Entendemos que, para ser sua única fonte de verdade, o Blue deve proteger seus dados empresariais sensíveis contra ameaças externas, perda de dados e tempo de inatividade.

Isso significa que levamos a segurança a sério no Blue.

Quando pensamos em segurança, consideramos uma abordagem holística que se concentra em três áreas principais:

1.  **Segurança da Infraestrutura e da Rede**: Garante que nossos sistemas físicos e virtuais estejam protegidos contra ameaças externas e acesso não autorizado.
2.  **Segurança do Software**: Foca na segurança do próprio código, incluindo práticas de codificação seguras, revisões regulares de código e gerenciamento de vulnerabilidades.
3.  **Segurança da Plataforma**: Inclui os recursos dentro do Blue, como [controles de acesso sofisticados](/platform/features/user-permissions), garantindo que os projetos sejam privados por padrão, e outras medidas para proteger os dados e a privacidade do usuário.


## Quão escalável é o Blue?

Esta é uma pergunta importante, pois você deseja um sistema que possa *crescer* com você. Você não quer ter que trocar sua plataforma de gerenciamento de projetos e processos em seis ou doze meses.

Escolhemos provedores de plataforma com cuidado, para garantir que eles possam lidar com as cargas de trabalho exigentes de nossos clientes. Usamos serviços em nuvem de alguns dos principais provedores de nuvem do mundo que atendem empresas como [Spotify](https://spotify.com) e [Netflix](https://netflix.com), que têm várias ordens de magnitude mais tráfego do que nós.

Os principais provedores de nuvem que usamos são:

- **[Cloudflare](https://cloudflare.com)**: Gerenciamos nosso DNS (Serviço de Nome de Domínio) via Cloudflare, assim como nosso site de marketing que roda em [Cloudflare Pages](https://pages.cloudflare.com/).
- **[Amazon Web Services](https://aws.amazon.com/)**: Usamos AWS para nosso banco de dados, que é [Aurora](https://aws.amazon.com/rds/aurora/), para armazenamento de arquivos via [Serviço de Armazenamento Simples (S3)](https://aws.amazon.com/s3/), e também para enviar e-mails via [Serviço de E-mail Simples (SES)](https://aws.amazon.com/ses/)
- **[Render](https://render.com)**: Usamos Render para nossos servidores front-end, servidores de aplicação/API, nossos serviços em segundo plano, sistema de filas e banco de dados Redis. Curiosamente, o Render é na verdade construído *sobre* o AWS! 


## Quão seguros são os arquivos no Blue?

Vamos começar com o armazenamento de dados. Nossos arquivos estão hospedados no [AWS S3](https://aws.amazon.com/s3/), que é o armazenamento de objetos em nuvem mais popular do mundo, com escalabilidade, disponibilidade de dados, segurança e desempenho líderes do setor.

Temos 99,99% de disponibilidade de arquivos e 99,999999999% de alta durabilidade.

Vamos detalhar o que isso significa.

A disponibilidade refere-se à quantidade de tempo que os dados estão operacionais e acessíveis. A disponibilidade de arquivos de 99,99% significa que podemos esperar que os arquivos fiquem indisponíveis por no máximo aproximadamente 8,76 horas por ano.

A durabilidade refere-se à probabilidade de que os dados permaneçam intactos e não corrompidos ao longo do tempo. Esse nível de durabilidade significa que podemos esperar perder no máximo um arquivo a cada 10 bilhões de arquivos enviados, graças à extensa redundância e replicação de dados em vários data centers.

Usamos [S3 Intelligent-Tiering](https://aws.amazon.com/s3/storage-classes/intelligent-tiering/) para mover automaticamente arquivos para diferentes classes de armazenamento com base na frequência de acesso. Com base nos padrões de atividade de centenas de milhares de projetos, notamos que a maioria dos arquivos é acessada em um padrão que se assemelha a uma curva de retrocesso exponencial. Isso significa que a maioria dos arquivos é acessada com muita frequência nos primeiros dias e, em seguida, é acessada cada vez menos frequentemente. Isso nos permite mover arquivos mais antigos para um armazenamento mais lento, mas significativamente mais barato, sem impactar a experiência do usuário de maneira significativa.

As economias de custo para isso são significativas. O S3 Standard-Infrequent Access (S3 Standard-IA) é aproximadamente 1,84 vezes mais barato que o S3 Standard. Isso significa que, para cada dólar que gastaríamos no S3 Standard, gastamos apenas cerca de 54 centavos no S3 Standard-IA para a mesma quantidade de dados armazenados.

| Recurso                  | S3 Standard             | S3 Standard-IA       |
|--------------------------|-------------------------|-----------------------|
| Custo de Armazenamento   | $0.023 - $0.021 por GB  | $0.0125 por GB        |
| Custo de Requisição (PUT, etc.) | $0.005 por 1.000 requisições | $0.01 por 1.000 requisições |
| Custo de Requisição (GET)       | $0.0004 por 1.000 requisições | $0.001 por 1.000 requisições |
| Custo de Recuperação de Dados      | $0.00 por GB            | $0.01 por GB          |


Os arquivos que você envia através do Blue são criptografados tanto em trânsito quanto em repouso. Os dados transferidos para e do Amazon S3 são protegidos usando [Transport Layer Security (TLS)](https://www.internetsociety.org/deploy360/tls/basics), protegendo contra [escuta clandestina](https://en.wikipedia.org/wiki/Network_eavesdropping) e [ataques de intermediário](https://en.wikipedia.org/wiki/Man-in-the-middle_attack). Para a criptografia em repouso, o Amazon S3 usa Criptografia do Lado do Servidor (SSE-S3), que criptografa automaticamente todos os novos uploads com criptografia AES-256, com a Amazon gerenciando as chaves de criptografia. Isso garante que seus dados permaneçam seguros durante todo o seu ciclo de vida.

## E quanto aos dados não relacionados a arquivos?

Nosso banco de dados é alimentado pelo [AWS Aurora](https://aws.amazon.com/rds/aurora/), um serviço moderno de banco de dados relacional que garante alto desempenho, disponibilidade e segurança para seus dados.

Os dados no Aurora são criptografados tanto em trânsito quanto em repouso. Usamos SSL (AES-256) para proteger as conexões entre sua instância de banco de dados e sua aplicação, protegendo os dados durante a transferência. Para a criptografia em repouso, o Aurora usa chaves gerenciadas pelo AWS Key Management Service (KMS), garantindo que todos os dados armazenados, incluindo backups automáticos, snapshots e réplicas, sejam criptografados e protegidos.

O Aurora possui um sistema de armazenamento distribuído, tolerante a falhas e auto-reparável. Este sistema é desacoplado dos recursos de computação e pode escalar automaticamente até 128 TiB por instância de banco de dados. Os dados são replicados em três [Zonas de Disponibilidade](https://aws.amazon.com/about-aws/global-infrastructure/regions_az/) (AZs), proporcionando resiliência contra perda de dados e garantindo alta disponibilidade. Em caso de falha do banco de dados, o Aurora reduz os tempos de recuperação para menos de 60 segundos, garantindo mínima interrupção.

O Blue faz backup continuamente de nosso banco de dados para o Amazon S3, permitindo recuperação em um ponto no tempo. Isso significa que podemos restaurar o banco de dados mestre do Blue para qualquer momento específico dentro dos últimos cinco minutos, garantindo que seus dados estejam sempre recuperáveis. Também tiramos snapshots regulares do banco de dados para períodos de retenção de backup mais longos.

Como um serviço totalmente gerenciado, o Aurora automatiza tarefas administrativas que consomem tempo, como provisionamento de hardware, configuração de banco de dados, correções e backups. Isso reduz a sobrecarga operacional e garante que nosso banco de dados esteja sempre atualizado com os últimos patches de segurança e melhorias de desempenho.

Se formos mais eficientes, podemos repassar nossas economias de custos para nossos clientes com nossos [preços líderes do setor](/pricing).

O Aurora está em conformidade com vários padrões da indústria, como HIPAA, GDPR e SOC 2, garantindo que suas práticas de gerenciamento de dados atendam a requisitos regulatórios rigorosos. Auditorias de segurança regulares e integração com [Amazon GuardDuty](https://aws.amazon.com/guardduty/) ajudam a detectar e mitigar potenciais ameaças à segurança.

## Como o Blue garante a segurança do login?

O Blue usa [links mágicos via e-mail](https://documentation.blue.cc/user-management/magic-links) para fornecer acesso seguro e conveniente à sua conta, eliminando a necessidade de senhas tradicionais.

Essa abordagem melhora significativamente a segurança, mitigando ameaças comuns associadas a logins baseados em senhas. Ao eliminar senhas, os links mágicos protegem contra ataques de phishing e roubo de senhas, *já que não há senha para roubar ou explorar.*

Cada link mágico é válido para apenas uma sessão de login, reduzindo o risco de acesso não autorizado. Além disso, esses links expiram após 15 minutos, garantindo que quaisquer links não utilizados não possam ser explorados, aumentando ainda mais a segurança.

A conveniência oferecida pelos links mágicos também é notável. Os links mágicos proporcionam uma experiência de login sem complicações, permitindo que você acesse sua conta *sem* a necessidade de lembrar senhas complexas.

Isso não apenas simplifica o processo de login, mas também previne violações de segurança que ocorrem quando senhas são reutilizadas em vários serviços. Muitos usuários tendem a usar a mesma senha em várias plataformas, o que significa que uma violação de segurança em um serviço poderia comprometer suas contas em outros serviços, incluindo o Blue. Ao usar links mágicos, a segurança do Blue não depende das práticas de segurança de outros serviços, proporcionando uma camada de proteção mais robusta e independente para nossos usuários.

Quando você solicita o login em sua conta do Blue, uma URL de login exclusiva é enviada para seu e-mail. Clicar neste link fará com que você entre instantaneamente em sua conta. O link é projetado para expirar após um único uso ou após 15 minutos, o que ocorrer primeiro, adicionando uma camada extra de segurança. Ao usar links mágicos, o Blue garante que seu processo de login seja seguro e amigável, proporcionando tranquilidade e conveniência.

## Como posso verificar a confiabilidade e o tempo de atividade do Blue?

No Blue, estamos comprometidos em manter um alto nível de confiabilidade e transparência para nossos usuários. Para fornecer visibilidade sobre o desempenho de nossa plataforma, oferecemos uma [página de status do sistema dedicada](https://status.blue.cc), que também está vinculada ao nosso rodapé em todas as páginas do nosso site.

![](/insights/status-page.png)

Esta página exibe nossos dados históricos de tempo de atividade, permitindo que você veja com que consistência nossos serviços estiveram disponíveis ao longo do tempo. Além disso, a página de status inclui relatórios detalhados de incidentes, proporcionando transparência sobre quaisquer problemas passados, seu impacto e as etapas que tomamos para resolvê-los e prevenir ocorrências futuras.