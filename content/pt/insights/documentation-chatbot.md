---
title: Por Que Construímos Nosso Próprio Chatbot de Documentação com IA
description: Construímos nosso próprio chatbot de documentação com IA que é treinado na documentação da plataforma Blue.
category: "Product Updates"
date: 2024-07-09
---



Na Blue, estamos sempre procurando maneiras de facilitar a vida de nossos clientes. Temos [documentação detalhada de cada recurso](https://documentation.blue.cc), [Vídeos no YouTube](https://www.youtube.com/@HeyBlueTeam), [Dicas e Truques](/insights/tips-tricks) e [vários canais de suporte](/support). 

Temos acompanhado de perto o desenvolvimento da IA (Inteligência Artificial), pois estamos muito envolvidos com [automação de gerenciamento de projetos](/platform/features/automations). Também lançamos recursos como [Auto Categorização com IA](/insights/ai-auto-categorization) e [Resumos com IA](/insights/ai-content-summarization) para facilitar o trabalho de nossos milhares de clientes. 

Uma coisa que está clara é que a IA veio para ficar e terá um efeito incrível na maioria das indústrias, e o gerenciamento de projetos não é exceção. Então, nos perguntamos como poderíamos aproveitar ainda mais a IA para ajudar em todo o ciclo de vida de um cliente, desde a descoberta, pré-vendas, integração e também com perguntas contínuas.

A resposta foi bastante clara: **Precisávamos de um chatbot de IA treinado em nossa documentação.**

Vamos encarar a realidade: *toda* organização provavelmente deveria ter um chatbot. Eles são ótimas maneiras para os clientes obterem respostas instantâneas a perguntas típicas, sem precisar vasculhar páginas de documentação densa ou seu site. A importância dos chatbots em sites de marketing modernos não pode ser subestimada. 

![](/insights/ai-chatbot-regular.png)

Para empresas de software especificamente, não se deve considerar o site de marketing como uma "coisa" separada — ele *é* parte do seu produto. Isso porque se encaixa na vida típica do cliente:

- **Conscientização** (Descoberta): É aqui que os potenciais clientes se deparam pela primeira vez com seu produto incrível. Seu chatbot pode ser o guia amigável deles, apontando para recursos e benefícios chave desde o início.
- **Consideração** (Educação): Agora eles estão curiosos e querem saber mais. Seu chatbot se torna o tutor pessoal deles, fornecendo informações adaptadas às suas necessidades e perguntas específicas.
- **Compra/Conversão**: Este é o momento da verdade - quando um prospect decide se arriscar e se tornar um cliente. Seu chatbot pode suavizar quaisquer contratempos de última hora, responder aquelas perguntas "só antes de eu comprar" e talvez até oferecer um bom negócio para fechar a venda.
- **Integração**: Eles compraram, e agora? Seu chatbot se transforma em um ajudante útil, guiando novos usuários pela configuração, mostrando como tudo funciona e garantindo que eles não se sintam perdidos no maravilhoso mundo do seu produto.
- **Retenção**: Manter os clientes felizes é o nome do jogo. Seu chatbot está disponível 24/7, pronto para solucionar problemas, oferecer dicas e truques, e garantir que seus clientes sintam o carinho.
- **Expansão**: Hora de subir de nível! Seu chatbot pode sugerir sutilmente novos recursos, upsells ou cross-sells que se alinhem com a forma como o cliente já está usando seu produto. É como ter um vendedor realmente inteligente e não insistente sempre à disposição.
- **Defesa**: Clientes felizes se tornam seus maiores torcedores. Seu chatbot pode encorajar usuários satisfeitos a espalhar a palavra, deixar avaliações ou participar de programas de referência. É como ter uma máquina de hype embutida diretamente no seu produto!

## Decisão de Construir ou Comprar

Uma vez que decidimos implementar um chatbot de IA, a próxima grande questão foi: construir ou comprar? Como uma equipe pequena focada em nosso produto principal, geralmente preferimos soluções "como serviço" ou plataformas populares de código aberto. Afinal, não estamos no negócio de reinventar a roda para cada parte de nossa pilha tecnológica.
Então, arregaçamos as mangas e mergulhamos no mercado, procurando soluções de chatbot de IA pagas e de código aberto. 

Nossos requisitos eram diretos, mas não negociáveis:

- **Experiência Sem Marca**: Este chatbot não é apenas um widget legal; ele vai para o nosso site de marketing e eventualmente para o nosso produto. Não estamos interessados em anunciar a marca de outra pessoa em nosso próprio espaço digital.
- **Ótima UX**: Para muitos potenciais clientes, este chatbot pode ser seu primeiro ponto de contato com a Blue. Ele define o tom para a percepção deles sobre nossa empresa. Vamos encarar a realidade: se não conseguimos implementar um chatbot adequado em nosso site, como podemos esperar que os clientes confiem em nós com seus projetos e processos críticos?
- **Custo Razoável**: Com uma grande base de usuários e planos de integrar o chatbot em nosso produto principal, precisávamos de uma solução que não quebrasse o banco à medida que o uso aumentasse. Idealmente, queríamos uma **opção BYOK (Bring Your Own Key)**. Isso nos permitiria usar nossa própria chave de serviço de IA, pagando apenas pelos custos variáveis diretos, em vez de um markup para um fornecedor terceirizado que não executa os modelos.
- **Compatível com a API de Assistentes da OpenAI**: Se fôssemos optar por um software de código aberto, não queríamos ter o trabalho de gerenciar um pipeline para ingestão de documentos, indexação, bancos de dados vetoriais e tudo isso. Queríamos usar a [API de Assistentes da OpenAI](https://platform.openai.com/docs/assistants/overview) que abstraísse toda a complexidade por trás de uma API. Honestamente — isso é muito bem feito. 
- **Escalabilidade**: Queremos ter este chatbot em vários lugares, com potencialmente dezenas de milhares de usuários por ano. Esperamos um uso significativo e não queremos ficar presos a uma solução que não possa escalar com nossas necessidades.

## Chatbots Comerciais de IA

Os que revisamos tendiam a ter uma UX melhor do que as soluções de código aberto — como infelizmente é muitas vezes o caso. Provavelmente haverá uma discussão separada a ser feita um dia sobre *por que* muitas soluções de código aberto ignoram ou subestimam a importância da UX. 

Vamos fornecer uma lista aqui, caso você esteja procurando algumas ofertas comerciais sólidas:

- **[Chatbase](https://chatbase.co):** O Chatbase permite que você construa um chatbot de IA personalizado treinado em sua base de conhecimento e o adicione ao seu site ou interaja com ele através de sua API. Oferece recursos como respostas confiáveis, geração de leads, análises avançadas e a capacidade de se conectar a várias fontes de dados. Para nós, isso parecia uma das ofertas comerciais mais polidas disponíveis. 
- **[DocsBot AI](https://docsbot.ai/):** O DocsBot AI cria bots personalizados do ChatGPT treinados em sua documentação e conteúdo para suporte, pré-vendas, pesquisa e mais. Ele fornece widgets embutíveis para adicionar facilmente o chatbot ao seu site, a capacidade de responder automaticamente a tickets de suporte e uma API poderosa para integração.
- **[CustomGPT.ai](https://customgpt.ai):** O CustomGPT.ai cria uma experiência de chatbot pessoal ao ingerir seus dados de negócios, incluindo conteúdo do site, helpdesk, bases de conhecimento, documentos e mais. Ele permite que leads façam perguntas e obtenham respostas instantâneas com base em seu conteúdo, sem precisar pesquisar. Curiosamente, eles também [afirmam superar a OpenAI em RAG (Geração Aumentada por Recuperação)!](https://customgpt.ai/customgpt-beats-open-ai-in-rag-benchmark/)
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Esta é uma oferta comercial interessante, porque também é um software de código aberto. Parece estar em um estágio inicial, e os preços não pareceram realistas ($27/mês para mensagens ilimitadas nunca funcionará comercialmente para eles).

Também analisamos [InterCom Fin](https://www.intercom.com/fin), que faz parte de seu software de suporte ao cliente. Isso significaria mudar do [HelpScout](https://wwww.helpscout.com), que usamos desde que começamos a Blue. Isso poderia ser possível, mas o InterCom Fin tem uma precificação insana que simplesmente o excluiu da consideração.

E esse é, na verdade, o problema com muitas das ofertas comerciais. O InterCom Fin cobra $0,99 por solicitação de suporte ao cliente tratada, e o Chatbase cobra $399/mês por 40.000 mensagens. Isso dá quase $5k por ano por um simples widget de chat. 

Considerando que os preços para inferência de IA estão caindo como loucos. A OpenAI reduziu seus preços de forma bastante dramática:

- O GPT-4 original (contexto de 8k) estava precificado em $0,03 por 1K tokens de prompt.
- O GPT-4 Turbo (contexto de 128k) estava precificado em $0,01 por 1K tokens de prompt, uma redução de 50% em relação ao GPT-4 original.
- O modelo GPT-4o está precificado em $0,005 por 1K tokens, o que é uma redução adicional de 50% em relação ao preço do GPT-4 Turbo.

Isso representa uma redução de 83% nos custos, e não esperamos que isso permaneça estagnado. 

Considerando que estávamos procurando uma solução escalável que seria usada por dezenas de milhares de usuários por ano com uma quantidade significativa de mensagens, faz sentido ir diretamente à fonte e pagar pelos custos da API diretamente, em vez de usar uma versão comercial que aumenta os custos.

## Chatbots de IA de Código Aberto

Como mencionado, as opções de código aberto que revisamos foram, em sua maioria, decepcionantes em relação ao requisito de "Ótima UX". 

Analisamos:

- **[Deepchat](https://deepchat.dev/)**: Este é um componente de chat independente de framework para serviços de IA, que se conecta a várias APIs de IA, incluindo OpenAI. Ele também tem a capacidade de os usuários baixarem um modelo de IA que roda diretamente no navegador. Brincamos com isso e conseguimos fazer uma versão funcionar, mas a API de Assistentes da OpenAI implementada parecia bastante bugada com vários problemas. No entanto, este é um projeto muito promissor, e seu playground é realmente bem feito. 
- **[OpenAssistantGPT](https://www.openassistantgpt.io/)**: Olhando para isso novamente de uma perspectiva de código aberto, isso exigiria que nós criássemos uma boa quantidade de infraestrutura, algo que não queríamos fazer, porque queríamos depender o máximo possível da API de Assistentes da OpenAI. 


## Construindo Nosso Próprio ChatBot

E assim, sem conseguir encontrar algo que atendesse a todos os nossos requisitos, decidimos construir nosso próprio chatbot de IA que pudesse interagir com a API de Assistentes da OpenAI. Isso, no final, acabou sendo relativamente indolor! 

Nosso site usa [Nuxt3](https://nuxt.com), [Vue3](https://vuejs.org/) (que é o mesmo framework da Plataforma Blue) e [TailwindUI](https://tailwindui.com/).

O primeiro passo foi criar uma API (Interface de Programação de Aplicativos) no Nuxt3 que pudesse "falar" com a API de Assistentes da OpenAI. Isso foi necessário, pois não queríamos fazer tudo no front-end, pois isso exporia nossa chave da API da OpenAI ao mundo, com potencial para abuso. 

Nossa API de backend atua como um intermediário seguro entre o navegador do usuário e a OpenAI. Aqui está o que ela faz:

- **Gerenciamento de Conversas:** Cria e gerencia "threads" para cada conversa. Pense em uma thread como uma sessão de chat única que lembra tudo o que você disse.
- **Manipulação de Mensagens:** Quando você envia uma mensagem, nossa API a adiciona à thread correta e pede ao assistente da OpenAI para elaborar uma resposta.
- **Espera Inteligente:** Em vez de fazer você encarar uma tela de carregamento, nossa API verifica com a OpenAI a cada segundo para ver se sua resposta está pronta. É como ter um garçom que fica de olho no seu pedido sem incomodar o chef a cada dois segundos.
- **Segurança em Primeiro Lugar:** Ao lidar com tudo isso no servidor, mantemos seus dados e nossas chaves de API seguras e protegidas.

Então, havia o front-end e a experiência do usuário. Como discutido anteriormente, isso era *crucial* porque não temos uma segunda chance de causar uma boa primeira impressão! 

Ao projetar nosso chatbot, prestamos atenção meticulosa à experiência do usuário, garantindo que cada interação seja suave, intuitiva e refletiva do compromisso da Blue com a qualidade. A interface do chatbot começa com um simples e elegante círculo azul, usando [HeroIcons para nossos ícones](https://heroicons.com/) (que usamos em todo o site da Blue) para atuar como nosso widget de abertura do chatbot. Essa escolha de design garante consistência visual e reconhecimento imediato da marca.


![](/insights/ai-chatbot-circle.png)

Entendemos que às vezes os usuários podem precisar de suporte adicional ou informações mais detalhadas. É por isso que incluímos links convenientes dentro da interface do chatbot. Um link de e-mail para suporte está prontamente disponível, permitindo que os usuários entrem em contato diretamente com nossa equipe se precisarem de assistência mais personalizada. Além disso, incorporamos um link para a documentação, proporcionando fácil acesso a recursos mais abrangentes para aqueles que desejam se aprofundar nas ofertas da Blue.

A experiência do usuário é ainda aprimorada por animações sutis de fade-in e fade-up ao abrir a janela do chatbot. Essas animações sutis adicionam um toque de sofisticação à interface, tornando a interação mais dinâmica e envolvente. Também implementamos um indicador de digitação, um pequeno, mas crucial recurso que informa os usuários que o chatbot está processando sua consulta e elaborando uma resposta. Esse sinal visual ajuda a gerenciar as expectativas do usuário e mantém uma sensação de comunicação ativa.

<video autoplay loop muted playsinline>
  <source src="/videos/ai-chatbot-animation.mp4" type="video/mp4">
</video>


Reconhecendo que algumas conversas podem exigir mais espaço na tela, adicionamos a capacidade de abrir a conversa em uma janela maior. Esse recurso é particularmente útil para trocas mais longas ou ao revisar informações detalhadas, dando aos usuários a flexibilidade de adaptar o chatbot às suas necessidades.

Nos bastidores, implementamos um processamento inteligente para otimizar as respostas do chatbot. Nosso sistema analisa automaticamente as respostas da IA para remover referências a nossos documentos internos, garantindo que as informações apresentadas sejam limpas, relevantes e focadas apenas em abordar a consulta do usuário.
Para melhorar a legibilidade e permitir uma comunicação mais nuançada, incorporamos suporte a markdown usando a biblioteca 'marked'. Esse recurso permite que nossa IA forneça texto ricamente formatado, incluindo ênfase em negrito e itálico, listas estruturadas e até mesmo trechos de código quando necessário. É como receber um mini-documento bem formatado e adaptado em resposta às suas perguntas.

Por último, mas certamente não menos importante, priorizamos a segurança em nossa implementação. Usando a biblioteca DOMPurify, sanitizamos o HTML gerado a partir da análise de markdown. Essa etapa crucial garante que quaisquer scripts ou códigos potencialmente prejudiciais sejam removidos antes que o conteúdo seja exibido para você. É nossa maneira de garantir que as informações úteis que você recebe sejam não apenas informativas, mas também seguras para consumo.


## Desenvolvimentos Futuros

Então, isso é apenas o começo, temos algumas coisas empolgantes no roteiro para esse recurso. 

Uma de nossas próximas funcionalidades é a capacidade de transmitir respostas em tempo real. Em breve, você verá as respostas do chatbot aparecerem caractere por caractere, tornando as conversas mais naturais e dinâmicas. É como assistir a IA pensar, criando uma experiência mais envolvente e interativa que mantém você informado a cada passo do caminho.

Para nossos valiosos usuários da Blue, estamos trabalhando na personalização. O chatbot reconhecerá quando você estiver logado, adaptando suas respostas com base nas informações da sua conta, histórico de uso e preferências. Imagine um chatbot que não apenas responde às suas perguntas, mas entende seu contexto específico dentro do ecossistema da Blue, fornecendo assistência mais relevante e personalizada.

Entendemos que você pode estar trabalhando em vários projetos ou ter várias consultas. É por isso que estamos desenvolvendo a capacidade de manter várias threads de conversa distintas com nosso chatbot. Esse recurso permitirá que você mude entre diferentes tópicos sem perder o contexto – assim como ter várias abas abertas em seu navegador.

Para tornar suas interações ainda mais produtivas, estamos criando um recurso que oferecerá perguntas de acompanhamento sugeridas com base em sua conversa atual. Isso ajudará você a explorar tópicos mais profundamente e descobrir informações relacionadas que você pode não ter pensado em perguntar, tornando cada sessão de chat mais abrangente e valiosa.

Estamos também empolgados em criar um conjunto de assistentes de IA especializados, cada um adaptado para necessidades específicas. Se você está procurando responder perguntas de pré-venda, configurar um novo projeto ou solucionar recursos avançados, você poderá escolher o assistente que melhor se adapta às suas necessidades atuais. É como ter uma equipe de especialistas da Blue ao seu alcance, cada um especializado em diferentes aspectos de nossa plataforma.

Por último, estamos trabalhando para permitir que você faça upload de capturas de tela diretamente para o chat. A IA analisará a imagem e fornecerá explicações ou etapas de solução de problemas com base no que vê. Esse recurso tornará mais fácil do que nunca obter ajuda com questões específicas que você encontra ao usar a Blue, conectando a informação visual à assistência textual.

## Conclusão

Esperamos que esta imersão em nosso processo de desenvolvimento de chatbot de IA tenha fornecido algumas percepções valiosas sobre nosso pensamento em desenvolvimento de produtos na Blue. Nossa jornada desde a identificação da necessidade de um chatbot até a construção de nossa própria solução demonstra como abordamos a tomada de decisões e a inovação.

![](/insights/ai-chatbot-modal.png)

Na Blue, pesamos cuidadosamente as opções de construir versus comprar, sempre com um olhar no que melhor servirá nossos usuários e se alinhará com nossos objetivos de longo prazo. Neste caso, identificamos uma lacuna significativa no mercado para um chatbot visualmente atraente e custo-efetivo que pudesse atender às nossas necessidades específicas. Embora geralmente defendamos a utilização de soluções existentes em vez de reinventar a roda, às vezes o melhor caminho a seguir é criar algo adaptado aos seus requisitos únicos.

Nossa decisão de construir nosso próprio chatbot não foi tomada levianamente. Foi o resultado de uma pesquisa de mercado minuciosa, uma compreensão clara de nossas necessidades e um compromisso em fornecer a melhor experiência possível para nossos usuários. Ao desenvolver internamente, conseguimos criar uma solução que não apenas atende às nossas necessidades atuais, mas também estabelece as bases para futuras melhorias e integrações.

Este projeto exemplifica nossa abordagem na Blue: não temos medo de arregaçar as mangas e construir algo do zero quando é a escolha certa para nosso produto e nossos usuários. É essa disposição de ir além que nos permite oferecer soluções inovadoras que realmente atendem às necessidades de nossos clientes.
Estamos empolgados com o futuro de nosso chatbot de IA e o valor que ele trará tanto para usuários potenciais quanto existentes da Blue. À medida que continuamos a refinar e expandir suas capacidades, permanecemos comprometidos em ultrapassar os limites do que é possível em gerenciamento de projetos e interação com o cliente.

Obrigado por nos acompanhar nesta jornada através de nosso processo de desenvolvimento. Esperamos que tenha dado a você um vislumbre da abordagem centrada no usuário que adotamos em todos os aspectos da Blue. Fique atento para mais atualizações enquanto continuamos a evoluir e aprimorar nossa plataforma para melhor atendê-lo.

Se você estiver interessado, pode encontrar o link para o código-fonte deste projeto aqui:

- **[ChatWidget](https://gitlab.com/bloohq/blue-website/-/blob/main/components/ChatWidget.vue)**: Este é um componente Vue que alimenta o widget de chat em si. 
- **[Chat API](https://gitlab.com/bloohq/blue-website/-/blob/main/server/api/chat.post.ts)**: Este é o middleware que funciona entre o componente de chat e a API de Assistentes da OpenAI.