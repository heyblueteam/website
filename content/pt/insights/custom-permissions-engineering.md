---
title: Criando o Motor de Permissões Personalizadas do Blue
description: Vá aos bastidores com a equipe de engenharia do Blue enquanto eles explicam como construir um recurso de categorização e etiquetagem automática impulsionado por IA.
category: "Engineering"
date: 2024-07-25
---



A gestão eficaz de projetos e processos é crucial para organizações de todos os tamanhos.

No Blue, [fizemos da nossa missão](/about) organizar o trabalho do mundo, construindo a melhor plataforma de gestão de projetos do planeta—simples, poderosa, flexível e acessível a todos.

Isso significa que nossa plataforma deve se adaptar às necessidades únicas de cada equipe. Hoje, estamos empolgados em revelar um dos nossos recursos mais poderosos: Permissões Personalizadas.

As ferramentas de gestão de projetos são a espinha dorsal dos fluxos de trabalho modernos, abrigando dados sensíveis, comunicações cruciais e planos estratégicos. Assim, a capacidade de controlar finamente o acesso a essas informações não é apenas um luxo—é uma necessidade.

<video autoplay loop muted playsinline>
  <source src="/videos/user-roles.mp4" type="video/mp4">
</video>


As permissões personalizadas desempenham um papel crítico nas plataformas B2B SaaS, especialmente nas ferramentas de gestão de projetos, onde o equilíbrio entre colaboração e segurança pode determinar o sucesso de um projeto.

Mas aqui é onde o Blue adota uma abordagem diferente: **acreditamos que recursos de nível empresarial não devem ser reservados apenas para orçamentos de grandes empresas.**

Em uma era em que a IA capacita pequenas equipes a operar em escalas sem precedentes, por que a segurança robusta e a personalização deveriam estar fora de alcance?

Nesta visão dos bastidores, exploraremos como desenvolvemos nosso recurso de Permissões Personalizadas, desafiando o status quo dos níveis de preços SaaS e trazendo opções de segurança poderosas e flexíveis para empresas de todos os tamanhos.

Se você é uma startup com grandes sonhos ou um jogador estabelecido buscando otimizar seus processos, as permissões personalizadas podem possibilitar novos casos de uso que você nunca soube que eram possíveis.

## Entendendo as Permissões de Usuário Personalizadas

Antes de mergulharmos em nossa jornada de desenvolvimento de permissões personalizadas para o Blue, vamos reservar um momento para entender o que são as permissões de usuário personalizadas e por que elas são tão cruciais em softwares de gestão de projetos.

As permissões de usuário personalizadas referem-se à capacidade de adaptar os direitos de acesso para usuários individuais ou grupos dentro de um sistema de software. Em vez de depender de funções predefinidas com conjuntos fixos de permissões, as permissões personalizadas permitem que os administradores criem perfis de acesso altamente específicos que se alinham perfeitamente com a estrutura e as necessidades de fluxo de trabalho de sua organização.

No contexto de softwares de gestão de projetos como o Blue, as permissões personalizadas incluem:

* **Controle de acesso granular**: Determinar quem pode visualizar, editar ou excluir tipos específicos de dados do projeto.
* **Restrições baseadas em recursos**: Habilitar ou desabilitar certos recursos para usuários ou equipes específicas.
* **Níveis de sensibilidade de dados**: Definir níveis variados de acesso a informações sensíveis dentro dos projetos.
* **Permissões específicas de fluxo de trabalho**: Alinhar as capacidades dos usuários com estágios ou aspectos específicos do fluxo de trabalho do seu projeto.

A importância das permissões personalizadas na gestão de projetos não pode ser subestimada:

* **Segurança aprimorada**: Ao fornecer aos usuários apenas o acesso de que precisam, você reduz o risco de vazamentos de dados ou alterações não autorizadas.
* **Conformidade melhorada**: As permissões personalizadas ajudam as organizações a atender aos requisitos regulatórios específicos do setor, controlando o acesso aos dados.
* **Colaboração otimizada**: As equipes podem trabalhar de forma mais eficiente quando cada membro tem o nível certo de acesso para desempenhar seu papel sem restrições desnecessárias ou privilégios excessivos.
* **Flexibilidade para organizações complexas**: À medida que as empresas crescem e evoluem, as permissões personalizadas permitem que o software se adapte a estruturas organizacionais e processos em mudança.

## Chegando ao SIM

[Já escrevemos antes](/insights/value-proposition-blue) que cada recurso no Blue precisa ser um **SIM** firme antes de decidirmos construí-lo. Não temos o luxo de centenas de engenheiros e de desperdiçar tempo e dinheiro construindo coisas que os clientes não precisam.

E assim, o caminho para implementar permissões personalizadas no Blue não foi uma linha reta. Como muitos recursos poderosos, começou com uma necessidade clara de nossos usuários e evoluiu através de uma consideração e planejamento cuidadosos.

Durante anos, nossos clientes solicitaram um controle mais granular sobre as permissões dos usuários. À medida que organizações de todos os tamanhos começaram a lidar com projetos cada vez mais complexos e sensíveis, as limitações do nosso controle de acesso baseado em funções padrão se tornaram evidentes.

Pequenas startups trabalhando com clientes externos, empresas de médio porte com processos de aprovação intrincados e grandes empresas com requisitos de conformidade rigorosos expressaram a mesma necessidade:

Mais flexibilidade na gestão do acesso dos usuários.

Apesar da demanda clara, inicialmente hesitamos em mergulhar no desenvolvimento de permissões personalizadas.

Por quê?

Entendemos a complexidade envolvida!

As permissões personalizadas tocam todas as partes de um sistema de gestão de projetos, desde a interface do usuário até a estrutura do banco de dados. Sabíamos que implementar esse recurso exigiria mudanças significativas em nossa arquitetura central e uma consideração cuidadosa das implicações de desempenho.

Ao analisarmos o cenário, percebemos que muito poucos de nossos concorrentes haviam tentado implementar um motor de permissões personalizadas poderoso como o que nossos clientes estavam solicitando. Aqueles que o fizeram frequentemente o reservavam para seus planos empresariais de mais alto nível.

Ficou claro o porquê: o esforço de desenvolvimento é substancial e os riscos são altos.

Implementar permissões personalizadas incorretamente poderia introduzir bugs críticos ou vulnerabilidades de segurança, comprometendo potencialmente todo o sistema. Essa percepção destacou a magnitude do desafio que estávamos considerando.

### Desafiando o Status Quo

No entanto, à medida que continuamos a crescer e evoluir, chegamos a uma realização crucial:

**O modelo SaaS tradicional de reservar recursos poderosos para clientes empresariais não faz mais sentido no cenário de negócios atual.**

Em 2024, com o poder da IA e ferramentas avançadas, pequenas equipes podem operar em uma escala e complexidade que rivalizam com organizações muito maiores. Uma startup pode estar lidando com dados sensíveis de clientes em vários países. Uma pequena agência de marketing pode estar gerenciando dezenas de projetos de clientes com diferentes requisitos de confidencialidade. Esses negócios precisam do mesmo nível de segurança e personalização que *qualquer* grande empresa.

Perguntamo-nos: Por que o tamanho da força de trabalho ou do orçamento de uma empresa deve determinar sua capacidade de manter seus dados seguros e seus processos eficientes?

### Recursos de Nível Empresarial para Todos

Essa realização nos levou a uma filosofia central que agora orienta grande parte do nosso desenvolvimento no Blue: Recursos de nível empresarial devem ser acessíveis a empresas de todos os tamanhos.

Acreditamos que:

- **A segurança não deve ser um luxo.** Toda empresa, independentemente do tamanho, merece as ferramentas para proteger seus dados e processos.
- **A flexibilidade impulsiona a inovação.** Ao fornecer a todos os nossos usuários ferramentas poderosas, capacitamos eles a criar fluxos de trabalho e sistemas que impulsionam suas indústrias para frente.
- **O crescimento não deve exigir mudanças na plataforma.** À medida que nossos clientes crescem, suas ferramentas devem crescer com eles de forma contínua.

Com essa mentalidade, decidimos enfrentar o desafio das permissões personalizadas de frente, comprometidos em torná-las disponíveis para todos os nossos usuários, não apenas para aqueles em planos de nível superior.

Essa decisão nos colocou em um caminho de design cuidadoso, desenvolvimento iterativo e feedback contínuo dos usuários que, em última análise, levou ao recurso de permissões personalizadas que temos orgulho de oferecer hoje.

Na próxima seção, vamos mergulhar em como abordamos o processo de design e desenvolvimento para trazer esse recurso complexo à vida.

### Design e Desenvolvimento

Quando decidimos enfrentar as permissões personalizadas, rapidamente percebemos que estávamos diante de uma tarefa colossal.

À primeira vista, "permissões personalizadas" pode parecer simples, mas é um recurso enganosamente complexo que toca todos os aspectos do nosso sistema.

O desafio era assustador: precisávamos implementar permissões em cascata, permitir edições em tempo real, fazer mudanças significativas no esquema do banco de dados e garantir funcionalidade contínua em todo o nosso ecossistema – aplicativos web, Mac, Windows, iOS e Android, assim como nossa API e webhooks.

A complexidade era suficiente para fazer até mesmo os desenvolvedores mais experientes hesitarem.

Nossa abordagem se concentrou em dois princípios-chave:

1. Dividir o recurso em versões gerenciáveis
2. Adotar o envio incremental.

Diante da complexidade das permissões personalizadas em grande escala, fizemos a nós mesmos uma pergunta crucial:

> Qual seria a versão mais simples possível desse recurso?

Essa abordagem se alinha ao princípio ágil de entregar um Produto Mínimo Viável (MVP) e iterar com base no feedback.

Nossa resposta foi refrescantemente simples:

1. Introduzir um interruptor para ocultar a aba de atividade do projeto
2. Adicionar outro interruptor para ocultar a aba de formulários

**Era isso.**

Sem sinos e apitos, sem matrizes de permissões complexas—apenas dois simples interruptores liga/desliga.

Embora possa parecer decepcionante à primeira vista, essa abordagem ofereceu várias vantagens significativas:

* **Implementação Rápida**: Esses interruptores simples poderiam ser desenvolvidos e testados rapidamente, permitindo que obtivéssemos uma versão básica das permissões personalizadas nas mãos dos usuários rapidamente.
* **Valor Claro para o Usuário**: Mesmo com apenas essas duas opções, estávamos fornecendo valor tangível. Algumas equipes podem querer ocultar o feed de atividades de clientes, enquanto outras podem precisar restringir o acesso a formulários para certos grupos de usuários.
* **Fundação para Crescimento**: Esse início simples lançou as bases para permissões mais complexas. Permitimos que configurássemos a infraestrutura básica para permissões personalizadas sem nos perdermos na complexidade desde o início.
* **Feedback dos Usuários**: Ao lançar essa versão simples, pudemos coletar feedback do mundo real sobre como os usuários interagiam com as permissões personalizadas, informando nosso desenvolvimento futuro.
* **Aprendizado Técnico**: Essa implementação inicial deu à nossa equipe de desenvolvimento experiência prática na modificação de permissões em nossa plataforma, preparando-nos para iterações mais complexas.

E você sabe, é realmente bastante humilde ter uma grande visão para algo e, em seguida, lançar algo que é uma porcentagem tão pequena dessa visão.

Depois de lançar esses primeiros dois interruptores, decidimos enfrentar algo mais sofisticado. Chegamos a duas novas permissões de função de usuário personalizadas.

A primeira foi a capacidade de limitar os usuários a visualizar apenas registros que foram especificamente atribuídos a eles. Isso é muito útil se você tem um cliente em um projeto e deseja que ele veja apenas registros que são especificamente atribuídos a ele, em vez de tudo que você está trabalhando para ele.

A segunda foi uma opção para administradores de projetos bloquearem grupos de usuários de poder convidar outros usuários. Isso é bom se você tem um projeto sensível que deseja garantir que permaneça em uma base de "necessidade de ver".

Uma vez que lançamos isso, ganhamos mais confiança e, para nossa terceira versão, abordamos permissões em nível de coluna, o que significa ser capaz de decidir quais campos personalizados um grupo específico de usuários pode visualizar ou editar.

Isso é extremamente poderoso. Imagine que você tem um projeto de CRM, e você tem dados lá que estão relacionados não apenas aos valores que o cliente pagará, mas também ao seu custo e margens de lucro. Você pode não querer que seus campos de custo e o campo da fórmula de margem do projeto sejam visíveis para funcionários juniores, e as permissões personalizadas permitem que você bloqueie esses campos para que não sejam mostrados.

Em seguida, passamos a criar permissões baseadas em listas, onde administradores de projetos podem decidir se um grupo de usuários pode visualizar, editar e excluir uma lista específica. Se eles ocultarem uma lista, todos os registros dentro dessa lista também ficam ocultos, o que é ótimo porque significa que você pode esconder certas partes do seu processo de membros da equipe ou clientes.

Esse é o resultado final:

<video autoplay loop muted playsinline>
  <source src="/videos/custom-user-roles.mp4" type="video/mp4">
</video>

## Considerações Técnicas

No coração da arquitetura técnica do Blue está o GraphQL, uma escolha crucial que influenciou significativamente nossa capacidade de implementar recursos complexos como permissões personalizadas. Mas antes de mergulharmos nos detalhes, vamos dar um passo atrás e entender o que é o GraphQL e como ele difere da abordagem mais tradicional da API REST.
GraphQL vs API REST: Uma Explicação Acessível

Imagine que você está em um restaurante. Com uma API REST, é como pedir de um menu fixo. Você pede um prato específico (endpoint) e recebe tudo que vem com ele, quer você queira tudo ou não. Se você quiser personalizar sua refeição, pode precisar fazer vários pedidos (chamadas de API) ou pedir um prato especialmente preparado (endpoint personalizado).

O GraphQL, por outro lado, é como ter uma conversa com um chef que pode preparar qualquer coisa. Você diz ao chef exatamente quais ingredientes deseja (campos de dados) e em quais quantidades. O chef então prepara um prato que é exatamente o que você pediu - nem mais, nem menos. Isso é essencialmente o que o GraphQL faz - permite que o cliente peça exatamente os dados de que precisa, e o servidor fornece apenas isso.

### Um Almoço Importante

Cerca de seis semanas após o desenvolvimento inicial do Blue, nosso engenheiro chefe e CEO saíram para almoçar.

O tópico da discussão?

Se deveríamos mudar de APIs REST para GraphQL. Essa não era uma decisão a ser tomada levianamente - adotar o GraphQL significaria descartar seis semanas de trabalho inicial.

Na caminhada de volta para o escritório, o CEO fez uma pergunta crucial ao engenheiro chefe: "Nós nos arrependeríamos de não fazer isso daqui a cinco anos?"

A resposta se tornou clara: o GraphQL era o caminho a seguir.

Reconhecemos o potencial dessa tecnologia desde o início, vendo como poderia apoiar nossa visão de uma plataforma de gestão de projetos flexível e poderosa.

Nossa previsão na adoção do GraphQL rendeu dividendos quando se tratou de implementar permissões personalizadas. Com uma API REST, precisaríamos de um endpoint diferente para cada configuração possível de permissões personalizadas - uma abordagem que rapidamente se tornaria difícil de gerenciar e manter.

O GraphQL, no entanto, nos permite lidar com permissões personalizadas de forma dinâmica. Veja como funciona:

- **Verificações de Permissão em Tempo Real**: Quando um cliente faz uma solicitação, nosso servidor GraphQL pode verificar as permissões do usuário diretamente em nosso banco de dados.
- **Recuperação Precisa de Dados**: Com base nessas permissões, o GraphQL retorna apenas os dados solicitados que se encaixam nos direitos de acesso do usuário.
- **Consultas Flexíveis**: À medida que as permissões mudam, não precisamos criar novos endpoints ou alterar os existentes. A mesma consulta GraphQL pode se adaptar a diferentes configurações de permissão.
- **Busca de Dados Eficiente**: O GraphQL permite que os clientes solicitem exatamente o que precisam. Isso significa que não estamos buscando dados em excesso, o que poderia expor informações que o usuário não deveria acessar.

Essa flexibilidade é crucial para um recurso tão complexo quanto as permissões personalizadas. Ela nos permite oferecer controle granular *sem* sacrificar o desempenho ou a manutenibilidade.

## Desafios

Implementar permissões personalizadas no Blue trouxe seus desafios, cada um nos empurrando a inovar e refinar nossa abordagem. A otimização de desempenho rapidamente se tornou uma preocupação crítica. À medida que adicionamos mais verificações de permissão granulares, corremos o risco de desacelerar nosso sistema, especialmente para grandes projetos com muitos usuários e configurações de permissão complexas. Para resolver isso, implementamos uma estratégia de cache em múltiplas camadas, otimizamos nossas consultas de banco de dados e aproveitamos a capacidade do GraphQL de solicitar apenas os dados necessários. Essa abordagem nos permitiu manter tempos de resposta rápidos, mesmo à medida que os projetos escalavam e a complexidade das permissões crescia.

A interface do usuário para permissões personalizadas apresentou outro obstáculo significativo. Precisávamos tornar a interface intuitiva e gerenciável para os administradores, mesmo à medida que adicionávamos mais opções e aumentávamos a complexidade do sistema.

Nossa solução envolveu várias rodadas de testes com usuários e design iterativo.

Introduzimos uma matriz visual de permissões que permitia aos administradores visualizar e modificar rapidamente as permissões em diferentes funções e áreas do projeto.

Garantir a consistência entre plataformas apresentou seu próprio conjunto de desafios. Precisávamos implementar permissões personalizadas de forma uniforme em nossos aplicativos web, desktop e móveis, cada um com sua interface e considerações de experiência do usuário únicas. Isso foi particularmente complicado para nossos aplicativos móveis, que precisavam ocultar e mostrar dinamicamente recursos com base nas permissões do usuário. Abordamos isso centralizando nossa lógica de permissões na camada da API, garantindo que todas as plataformas recebessem dados de permissões consistentes.

Em seguida, desenvolvemos uma estrutura de UI flexível que pudesse se adaptar a essas mudanças de permissão em tempo real, proporcionando uma experiência contínua, independentemente da plataforma utilizada.

A educação e adoção dos usuários apresentaram o último obstáculo em nossa jornada de permissões personalizadas. Introduzir um recurso tão poderoso significava que precisávamos ajudar nossos usuários a entender e aproveitar efetivamente as permissões personalizadas.

Inicialmente, lançamos as permissões personalizadas para um subconjunto de nossa base de usuários, monitorando cuidadosamente suas experiências e coletando insights. Essa abordagem nos permitiu refinar o recurso e nossos materiais educacionais com base no uso do mundo real antes de lançá-lo para toda a nossa base de usuários.

O lançamento em fases provou ser inestimável, ajudando-nos a identificar e resolver problemas menores e pontos de confusão dos usuários que não havíamos antecipado, levando, em última análise, a um recurso mais polido e amigável para todos os nossos usuários.

Essa abordagem de lançar para um subconjunto de usuários, assim como nosso típico período de "Beta" de 2 a 3 semanas em nosso Beta público, nos ajuda a dormir à noite. :)

## Olhando para o Futuro

Como acontece com todos os recursos, nada está *"pronto"*.

Nossa visão de longo prazo para o recurso de permissões personalizadas se estende por tags, filtros de campos personalizados, navegação de projetos personalizável e controles de comentários.

Vamos mergulhar em cada aspecto.

### Permissões de Tags

Achamos que seria incrível poder criar permissões com base em se um registro tem uma ou mais tags. O caso de uso mais óbvio seria criar uma função de usuário personalizada chamada "Clientes" e permitir que apenas usuários nessa função vissem registros que tivessem a tag "Clientes".

Isso lhe dá uma visão rápida de se um registro pode ou não ser visto por seus clientes.

Isso poderia se tornar ainda mais poderoso com combinadores AND/OR, onde você pode especificar regras mais complexas. Por exemplo, você poderia configurar uma regra que permite acesso a registros etiquetados tanto como "Clientes" QUANTO "Público", ou registros etiquetados como "Interno" OU "Confidencial". Esse nível de flexibilidade permitiria configurações de permissão incrivelmente nuançadas, atendendo até mesmo as estruturas organizacionais e fluxos de trabalho mais complexos.

As aplicações potenciais são vastas. Gerentes de projeto poderiam facilmente segregar informações sensíveis, equipes de vendas poderiam ter acesso automático a dados relevantes de clientes, e colaboradores externos poderiam ser integrados de forma transparente em partes específicas de um projeto sem risco de exposição a informações internas sensíveis.

### Filtros de Campos Personalizados

Nossa visão para Filtros de Campos Personalizados representa um salto significativo no controle de acesso granular. Esse recurso capacitará administradores de projetos a definir quais registros grupos específicos de usuários podem ver com base nos valores de campos personalizados. Trata-se de criar limites dinâmicos e orientados por dados para o acesso à informação.

Imagine ser capaz de configurar permissões assim:

- Mostrar apenas registros onde o dropdown "Status do Projeto" está definido como "Público"
- Restringir a visibilidade para itens onde o campo de múltipla seleção "Departamento" inclui "Marketing"
- Permitir acesso a tarefas onde a caixa de seleção "Prioridade" está marcada
- Exibir projetos onde o campo numérico "Orçamento" está acima de um certo limite

### Navegação de Projetos Personalizável

Isso é simplesmente uma extensão dos interruptores que já temos. Em vez de apenas ter interruptores para "atividade" e "formulários", queremos estender isso a cada parte da navegação do projeto. Dessa forma, os administradores de projetos podem criar interfaces focadas e remover ferramentas que não precisam.

### Controles de Comentários

No futuro, queremos ser criativos em como permitimos que nossos clientes decidam quem pode e não pode ver comentários. Podemos permitir várias áreas de comentários em abas sob um registro, e cada uma pode ser visível ou não visível para diferentes grupos de usuários.

Além disso, também podemos permitir um recurso onde apenas comentários em que um usuário é *especificamente* mencionado são visíveis, e nada mais. Isso permitiria que equipes que têm clientes em projetos garantissem que apenas comentários que desejam que os clientes vejam sejam visíveis.

## Conclusão

Então, aí está, assim abordamos a construção de um dos recursos mais interessantes e poderosos! [Como você pode ver em nossa ferramenta de comparação de gestão de projetos](/compare), muito poucos sistemas de gestão de projetos têm uma configuração de matriz de permissões tão poderosa, e os que têm reservam isso para seus planos empresariais mais caros, tornando-o inacessível para uma empresa pequena ou média típica.

Com o Blue, você tem *todos* os recursos disponíveis com nosso plano — não acreditamos que recursos de nível empresarial devam ser reservados para clientes empresariais!