---
title: Categorização Automática com IA (Análise Técnica Aprofundada)
category: "Engineering"
description: Conheça os bastidores com a equipe de engenharia da Blue enquanto explicam como construíram um recurso de categorização e etiquetagem automática alimentado por IA.
date: 2024-12-07
---

Recentemente lançamos a [Categorização Automática com IA](/insights/ai-auto-categorization) para todos os usuários da Blue. Este é um recurso de IA incluído na assinatura principal da Blue, sem custos adicionais. Neste post, mergulhamos na engenharia por trás da criação deste recurso.

---
Na Blue, nossa abordagem para o desenvolvimento de recursos está enraizada em uma compreensão profunda das necessidades dos usuários e tendências de mercado, aliada ao compromisso de manter a simplicidade e facilidade de uso que definem nossa plataforma. Isso é o que impulsiona nosso [roadmap](/platform/roadmap), e o que nos [permitiu entregar consistentemente recursos todos os meses por anos](/platform/changelog).

A introdução da etiquetagem automática alimentada por IA na Blue é um exemplo perfeito dessa filosofia em ação. Antes de mergulharmos nos detalhes técnicos de como construímos este recurso, é crucial entender o problema que estávamos resolvendo e a consideração cuidadosa que entrou em seu desenvolvimento.

O cenário de gerenciamento de projetos está evoluindo rapidamente, com as capacidades de IA tornando-se cada vez mais centrais para as expectativas dos usuários. Nossos clientes, particularmente aqueles gerenciando [projetos](/platform) em grande escala com milhões de [registros](/platform/features/records), foram vocais sobre seu desejo por maneiras mais inteligentes e eficientes de organizar e categorizar seus dados.

No entanto, na Blue, não simplesmente adicionamos recursos porque são tendência ou solicitados. Nossa filosofia é que cada nova adição deve provar seu valor, com a resposta padrão sendo um firme *"não"* até que um recurso demonstre forte demanda e utilidade clara.

Para realmente entender a profundidade do problema e o potencial da etiquetagem automática com IA, conduzimos extensas entrevistas com clientes, focando em usuários de longa data que gerenciam projetos complexos e ricos em dados em múltiplos domínios.

Essas conversas revelaram um fio comum: *embora a etiquetagem fosse inestimável para organização e capacidade de busca, a natureza manual do processo estava se tornando um gargalo, especialmente para equipes lidando com altos volumes de registros.*

Mas vimos além de apenas resolver o ponto de dor imediato da etiquetagem manual.

Vislumbramos um futuro onde a etiquetagem alimentada por IA poderia se tornar a base para fluxos de trabalho mais inteligentes e automatizados.

O verdadeiro poder deste recurso, percebemos, estava em seu potencial de integração com nosso [sistema de automação de gerenciamento de projetos](/platform/features/automations). Imagine uma ferramenta de gerenciamento de projetos que não apenas categoriza informações inteligentemente, mas também usa essas categorias para rotear tarefas, acionar ações e adaptar fluxos de trabalho em tempo real.

Essa visão alinhava-se perfeitamente com nosso objetivo de manter a Blue simples, mas poderosa.

Além disso, reconhecemos o potencial de estender essa capacidade além dos limites de nossa plataforma. Ao desenvolver um sistema robusto de etiquetagem com IA, estávamos estabelecendo as bases para uma "API de categorização" que poderia funcionar imediatamente, potencialmente abrindo novos caminhos para como nossos usuários interagem e aproveitam a Blue em seus ecossistemas tecnológicos mais amplos.

Este recurso, portanto, não era apenas sobre adicionar uma caixa de seleção de IA à nossa lista de recursos.

Era sobre dar um passo significativo em direção a uma plataforma de gerenciamento de projetos mais inteligente e adaptativa, mantendo-se fiel à nossa filosofia central de simplicidade e foco no usuário.

Nas seções seguintes, mergulharemos nos desafios técnicos que enfrentamos para dar vida a essa visão, a arquitetura que projetamos para suportá-la e as soluções que implementamos. Também exploraremos as possibilidades futuras que este recurso abre, mostrando como uma adição cuidadosamente considerada pode pavimentar o caminho para mudanças transformadoras no gerenciamento de projetos.

---
## O Problema

Como discutido acima, a etiquetagem manual de registros de projeto pode ser demorada e inconsistente.

Nos propusemos a resolver isso aproveitando IA para sugerir automaticamente tags com base no conteúdo do registro.

Os principais desafios foram:

1. Escolher um modelo de IA apropriado
2. Processar eficientemente grandes volumes de registros
3. Garantir privacidade e segurança de dados
4. Integrar o recurso perfeitamente em nossa arquitetura existente

## Selecionando o Modelo de IA

Avaliamos várias plataformas de IA, incluindo [OpenAI](https://openai.com), modelos de código aberto no [HuggingFace](https://huggingface.co/) e [Replicate](https://replicate.com).

Nossos critérios incluíram:

- Custo-benefício
- Precisão na compreensão do contexto
- Capacidade de aderir a formatos de saída específicos
- Garantias de privacidade de dados

Após testes minuciosos, escolhemos o [GPT-3.5 Turbo](https://platform.openai.com/docs/models/gpt-3-5-turbo) da OpenAI. Embora o [GPT-4](https://softgist.com/the-ultimate-guide-to-prompt-engineering) pudesse oferecer melhorias marginais na precisão, nossos testes mostraram que o desempenho do GPT-3.5 era mais do que adequado para nossas necessidades de etiquetagem automática. O equilíbrio entre custo-benefício e fortes capacidades de categorização fez do GPT-3.5 a escolha ideal para este recurso.

O custo mais alto do GPT-4 nos teria forçado a oferecer o recurso como um complemento pago, conflitando com nosso objetivo de **incluir IA em nosso produto principal sem custo adicional para os usuários finais.**

No momento de nossa implementação, os preços para o GPT-3.5 Turbo são:

- $0.0005 por 1K tokens de entrada (ou $0.50 por 1M tokens de entrada)
- $0.0015 por 1K tokens de saída (ou $1.50 por 1M tokens de saída)

Vamos fazer algumas suposições sobre um registro médio na Blue:

- **Título**: ~10 tokens
- **Descrição**: ~50 tokens
- **2 comentários**: ~30 tokens cada
- **5 campos personalizados**: ~10 tokens cada
- **Nome da lista, data de vencimento e outros metadados**: ~20 tokens
- **Prompt do sistema e tags disponíveis**: ~50 tokens

Total de tokens de entrada por registro: 10 + 50 + (30 * 2) + (10 * 5) + 20 + 50 ≈ 240 tokens

Para a saída, vamos supor uma média de 3 tags sugeridas por registro, que podem totalizar cerca de 20 tokens de saída incluindo a formatação JSON.

Para 1 milhão de registros:

- Custo de entrada: (240 * 1.000.000 / 1.000.000) * $0.50 = $120
- Custo de saída: (20 * 1.000.000 / 1.000.000) * $1.50 = $30

**Custo total para etiquetagem automática de 1 milhão de registros: $120 + $30 = $150**

## Desempenho do GPT3.5 Turbo

A categorização é uma tarefa em que modelos de linguagem grande (LLMs) como o GPT-3.5 Turbo se destacam, tornando-os particularmente adequados para nosso recurso de etiquetagem automática. LLMs são treinados em vastas quantidades de dados de texto, permitindo-lhes entender contexto, semântica e relações entre conceitos. Essa ampla base de conhecimento permite que realizem tarefas de categorização com alta precisão em uma ampla gama de domínios.

Para nosso caso de uso específico de etiquetagem de gerenciamento de projetos, o GPT-3.5 Turbo demonstra várias forças principais:

- **Compreensão Contextual:** Pode captar o contexto geral de um registro de projeto, considerando não apenas palavras individuais, mas o significado transmitido por toda a descrição, comentários e outros campos.
- **Flexibilidade:** Pode se adaptar a vários tipos de projetos e indústrias sem exigir reprogramação extensiva.
- **Lidando com Ambiguidade:** Pode pesar múltiplos fatores para tomar decisões nuançadas.
- **Aprendendo com Exemplos:** Pode entender e aplicar rapidamente novos esquemas de categorização sem treinamento adicional.
- **Classificação Multi-rótulo:** Pode sugerir múltiplas tags relevantes para um único registro, o que foi crucial para nossos requisitos.

O GPT-3.5 Turbo também se destacou por sua confiabilidade em aderir ao nosso formato de saída JSON necessário, o que foi *crucial* para integração perfeita com nossos sistemas existentes. Modelos de código aberto, embora promissores, frequentemente adicionavam comentários extras ou desviavam do formato esperado, o que teria exigido pós-processamento adicional. Essa consistência no formato de saída foi um fator chave em nossa decisão, pois simplificou significativamente nossa implementação e reduziu potenciais pontos de falha.

Optar pelo GPT-3.5 Turbo com sua saída JSON consistente nos permitiu implementar uma solução mais direta, confiável e sustentável.

Se tivéssemos escolhido um modelo com formatação menos confiável, teríamos enfrentado uma cascata de complicações: a necessidade de lógica de análise robusta para lidar com vários formatos de saída, tratamento extensivo de erros para saídas inconsistentes, potenciais impactos de desempenho do processamento adicional, complexidade aumentada de testes para cobrir todas as variações de saída e uma maior carga de manutenção a longo prazo.

Erros de análise poderiam levar a etiquetagem incorreta, impactando negativamente a experiência do usuário. Ao evitar essas armadilhas, pudemos focar nossos esforços de engenharia em aspectos críticos como otimização de desempenho e design de interface do usuário, em vez de lutar com saídas de IA imprevisíveis.

## Arquitetura do Sistema

Nosso recurso de etiquetagem automática com IA é construído sobre uma arquitetura robusta e escalável, projetada para lidar com altos volumes de solicitações eficientemente, enquanto proporciona uma experiência de usuário perfeita. Como com todos os nossos sistemas, arquitetamos este recurso para suportar uma ordem de magnitude a mais de tráfego do que experimentamos atualmente. Essa abordagem, embora aparentemente sobre-engenheirada para as necessidades atuais, é uma prática recomendada que nos permite lidar perfeitamente com picos repentinos de uso e nos dá ampla margem para crescimento sem grandes revisões arquitetônicas. Caso contrário, teríamos que reengenheirar todos os nossos sistemas a cada 18 meses — algo que aprendemos da maneira difícil no passado!

Vamos detalhar os componentes e o fluxo de nosso sistema:

- **Interação do Usuário:** O processo começa quando um usuário pressiona o botão "Etiquetar automaticamente" na interface da Blue. Esta ação aciona o fluxo de trabalho de etiquetagem automática.
- **Chamada da API Blue:** A ação do usuário é traduzida em uma chamada de API para nosso backend Blue. Este endpoint de API é projetado para lidar com solicitações de etiquetagem automática.
- **Gerenciamento de Fila:** Em vez de processar a solicitação imediatamente, o que poderia levar a problemas de desempenho sob alta carga, adicionamos a solicitação de etiquetagem a uma fila. Usamos Redis para este mecanismo de fila, o que nos permite gerenciar a carga efetivamente e garantir a escalabilidade do sistema.
- **Serviço em Background:** Implementamos um serviço em background que monitora continuamente a fila por novas solicitações. Este serviço é responsável por processar solicitações enfileiradas.
- **Integração com API OpenAI:** O serviço em background prepara os dados necessários e faz chamadas de API para o modelo GPT-3.5 da OpenAI. É aqui que a etiquetagem alimentada por IA realmente ocorre. Enviamos dados relevantes do projeto e recebemos tags sugeridas em retorno.
- **Processamento de Resultados:** O serviço em background processa os resultados recebidos da OpenAI. Isso envolve analisar a resposta da IA e preparar os dados para aplicação ao projeto.
- **Aplicação de Tags:** Os resultados processados são usados para aplicar as novas tags aos itens relevantes no projeto. Este passo atualiza nosso banco de dados com as tags sugeridas pela IA.
- **Reflexão na Interface do Usuário:** Finalmente, as novas tags aparecem na visualização do projeto do usuário, completando o processo de etiquetagem automática do ponto de vista do usuário.

Esta arquitetura oferece vários benefícios-chave que melhoram tanto o desempenho do sistema quanto a experiência do usuário. Ao utilizar uma fila e processamento em background, alcançamos escalabilidade impressionante, permitindo-nos lidar com numerosas solicitações simultaneamente sem sobrecarregar nosso sistema ou atingir os limites de taxa da API OpenAI. Implementar esta arquitetura exigiu consideração cuidadosa de vários fatores para garantir desempenho e confiabilidade ótimos. Para gerenciamento de fila, escolhemos Redis, aproveitando sua velocidade e confiabilidade no manuseio de filas distribuídas.

Essa abordagem também contribui para a responsividade geral do recurso. Os usuários recebem feedback imediato de que sua solicitação está sendo processada, mesmo que a etiquetagem real leve algum tempo, criando uma sensação de interação em tempo real. A tolerância a falhas da arquitetura é outra vantagem crucial. Se qualquer parte do processo encontrar problemas, como interrupções temporárias da API OpenAI, podemos reintentar graciosamente ou lidar com a falha sem impactar todo o sistema.

Essa robustez, combinada com o aparecimento em tempo real das tags, melhora a experiência do usuário, dando a impressão de "mágica" de IA em ação.

## Dados e Prompts

Um passo crucial em nosso processo de etiquetagem automática é preparar os dados a serem enviados ao modelo GPT-3.5. Este passo exigiu consideração cuidadosa para equilibrar o fornecimento de contexto suficiente para etiquetagem precisa, mantendo eficiência e protegendo a privacidade do usuário. Aqui está uma visão detalhada de nosso processo de preparação de dados.

Para cada registro, compilamos as seguintes informações:

- **Nome da Lista**: Fornece contexto sobre a categoria ou fase mais ampla do projeto.
- **Título do Registro**: Frequentemente contém informações-chave sobre o propósito ou conteúdo do registro.
- **Campos Personalizados**: Incluímos [campos personalizados](/platform/features/custom-fields) baseados em texto e números, que frequentemente contêm informações cruciais específicas do projeto.
- **Descrição**: Normalmente contém as informações mais detalhadas sobre o registro.
- **Comentários**: Podem fornecer contexto adicional ou atualizações que podem ser relevantes para etiquetagem.
- **Data de Vencimento**: Informações temporais que podem influenciar a seleção de tags.

Interessantemente, não enviamos dados de tags existentes ao GPT-3.5, e fazemos isso para evitar enviesar o modelo.

O núcleo do nosso recurso de etiquetagem automática está em como interagimos com o modelo GPT-3.5 e processamos suas respostas. Esta seção de nosso pipeline exigiu design cuidadoso para garantir etiquetagem precisa, consistente e eficiente.

Usamos um prompt de sistema cuidadosamente elaborado para instruir a IA sobre sua tarefa. Aqui está uma análise de nosso prompt e a lógica por trás de cada componente:

```
You will be provided with record data, and your task is to choose the tags that are relevant to the record.
You can respond with an empty array if you are unsure.
Available tags: ${tags}.
Today: ${currentDate}.
Please respond in JSON using the following format:
{ "tags": ["tag-1", "tag-2"] }
```

- **Definição da Tarefa:** Declaramos claramente a tarefa da IA para garantir respostas focadas.
- **Tratamento de Incerteza:** Permitimos explicitamente respostas vazias, evitando etiquetagem forçada quando a IA está incerta.
- **Tags Disponíveis:** Fornecemos uma lista de tags válidas (${tags}) para restringir as escolhas da IA às tags existentes do projeto.
- **Data Atual:** Incluir ${currentDate} ajuda a IA a entender o contexto temporal, que pode ser crucial para certos tipos de projetos.
- **Formato de Resposta:** Especificamos um formato JSON para análise fácil e verificação de erros.

Este prompt é o resultado de testes e iterações extensivos. Descobrimos que ser explícito sobre a tarefa, opções disponíveis e formato de saída desejado melhorou significativamente a precisão e consistência das respostas da IA — simplicidade é chave!

A lista de tags disponíveis é gerada no lado do servidor e validada antes da inclusão no prompt. Implementamos limites rígidos de caracteres em nomes de tags para evitar prompts muito grandes.

Como mencionado acima, não tivemos problemas com o GPT-3.5 Turbo em obter de volta a resposta JSON pura no formato correto 100% do tempo.

Então, em resumo,

- Combinamos o prompt do sistema com os dados do registro preparados.
- Este prompt combinado é enviado ao modelo GPT-3.5 via API da OpenAI.
- Usamos uma configuração de temperatura de 0.3 para equilibrar criatividade e consistência nas respostas da IA.
- Nossa chamada de API inclui um parâmetro max_tokens para limitar o tamanho da resposta e controlar custos.

Uma vez que recebemos a resposta da IA, passamos por várias etapas para processar e aplicar as tags sugeridas:

* **Análise JSON**: Tentamos analisar a resposta como JSON. Se a análise falhar, registramos o erro e pulamos a etiquetagem para esse registro.
* **Validação de Schema**: Validamos o JSON analisado contra nosso schema esperado (um objeto com um array "tags"). Isso captura quaisquer desvios estruturais na resposta da IA.
* **Validação de Tags**: Cruzamos as tags sugeridas com nossa lista de tags válidas do projeto. Este passo filtra quaisquer tags que não existem no projeto, o que poderia ocorrer se a IA entendesse mal ou se as tags do projeto mudassem entre a criação do prompt e o processamento da resposta.
* **Deduplicação**: Removemos quaisquer tags duplicadas da sugestão da IA para evitar etiquetagem redundante.
* **Aplicação**: As tags validadas e deduplicadas são então aplicadas ao registro em nosso banco de dados.
* **Registro e Análise**: Registramos as tags finais aplicadas. Esses dados são valiosos para monitorar o desempenho do sistema e melhorá-lo ao longo do tempo.

## Desafios

Implementar etiquetagem automática alimentada por IA na Blue apresentou vários desafios únicos, cada um exigindo soluções inovadoras para garantir um recurso robusto, eficiente e amigável ao usuário.

### Desfazer Operação em Massa

O recurso de Etiquetagem com IA pode ser feito tanto em registros individuais quanto em massa. O problema com a operação em massa é que, se o usuário não gostar do resultado, teria que passar manualmente por milhares de registros e desfazer o trabalho da IA. Claramente, isso é inaceitável.

Para resolver isso, implementamos um sistema inovador de sessão de etiquetagem. Cada operação de etiquetagem em massa recebe um ID de sessão único, que é associado a todas as tags aplicadas durante essa sessão. Isso nos permite gerenciar eficientemente operações de desfazer simplesmente excluindo todas as tags associadas a um ID de sessão específico. Também removemos trilhas de auditoria relacionadas, garantindo que operações desfeitas não deixem rastros no sistema. Essa abordagem dá aos usuários a confiança para experimentar com etiquetagem de IA, sabendo que podem facilmente reverter mudanças se necessário.

### Privacidade de Dados

A privacidade de dados foi outro desafio crítico que enfrentamos.

Nossos usuários confiam em nós com seus dados de projeto, e era fundamental garantir que essas informações não fossem retidas ou usadas para treinamento de modelo pela OpenAI. Abordamos isso em múltiplas frentes.

Primeiro, formamos um acordo com a OpenAI que proíbe explicitamente o uso de nossos dados para treinamento de modelo. Além disso, a OpenAI exclui os dados após o processamento, fornecendo uma camada extra de proteção de privacidade.

Por nossa parte, tomamos a precaução de excluir informações sensíveis, como detalhes de responsáveis, dos dados enviados à IA, garantindo assim que nomes específicos de indivíduos não sejam enviados a terceiros junto com outros dados.

Essa abordagem abrangente nos permite aproveitar as capacidades de IA mantendo os mais altos padrões de privacidade e segurança de dados.

### Limites de Taxa e Captura de Erros

Uma de nossas principais preocupações era escalabilidade e limitação de taxa. Chamadas diretas de API para a OpenAI para cada registro teriam sido ineficientes e poderiam rapidamente atingir limites de taxa, especialmente para grandes projetos ou durante horários de pico de uso. Para abordar isso, desenvolvemos uma arquitetura de serviço em background que nos permite agrupar solicitações e implementar nosso próprio sistema de fila. Essa abordagem nos ajuda a gerenciar a frequência de chamadas de API e permite processamento mais eficiente de grandes volumes de registros, garantindo desempenho suave mesmo sob carga pesada.

A natureza das interações de IA significava que também tínhamos que nos preparar para erros ocasionais ou saídas inesperadas. Houve instâncias em que a IA poderia produzir JSON inválido ou saídas que não correspondiam ao nosso formato esperado. Para lidar com isso, implementamos tratamento robusto de erros e lógica de análise em todo o nosso sistema. Se a resposta da IA não for JSON válido ou não contiver a chave "tags" esperada, nosso sistema é projetado para tratá-la como se nenhuma tag fosse sugerida, em vez de tentar processar dados potencialmente corrompidos. Isso garante que, mesmo diante da imprevisibilidade da IA, nosso sistema permaneça estável e confiável.

## Desenvolvimentos Futuros

Acreditamos que recursos, e o produto Blue como um todo, nunca estão "prontos" — sempre há espaço para melhorias.

Havia alguns recursos que consideramos na construção inicial que não passaram pela fase de escopo, mas são interessantes notar, pois provavelmente implementaremos alguma versão deles no futuro.

O primeiro é adicionar descrição de tag. Isso permitiria aos usuários finais não apenas dar às tags um nome e uma cor, mas também uma descrição opcional. Isso também seria passado para a IA para ajudar a fornecer mais contexto e potencialmente melhorar a precisão.

Embora contexto adicional possa ser valioso, estamos atentos à complexidade potencial que pode introduzir. Há um delicado equilíbrio a ser alcançado entre fornecer informações úteis e sobrecarregar os usuários com muitos detalhes. À medida que desenvolvemos este recurso, nos concentraremos em encontrar esse ponto ideal onde o contexto adicionado melhora em vez de complicar a experiência do usuário.

Talvez o aprimoramento mais empolgante em nosso horizonte seja a integração da etiquetagem automática com IA com nosso [sistema de automação de gerenciamento de projetos](/platform/features/automations).

Isso significa que o recurso de etiquetagem com IA poderia ser tanto um gatilho quanto uma ação de uma automação. Isso poderia ser enorme, pois poderia transformar este recurso de categorização de IA relativamente simples em um sistema de roteamento baseado em IA para trabalho.

Imagine uma automação que declare:

Quando IA etiquetar um registro como "Crítico" -> Atribuir ao "Gerente" e Enviar um Email Personalizado

Isso significa que quando você etiquetar um registro com IA, se a IA decidir que é um problema crítico, então pode automaticamente atribuir ao gerente do projeto e enviar-lhe um email personalizado. Isso estende os [benefícios de nosso sistema de automação de gerenciamento de projetos](/platform/features/automations) de um sistema puramente baseado em regras para um verdadeiro sistema flexível de IA.

Ao explorar continuamente as fronteiras da IA no gerenciamento de projetos, pretendemos fornecer aos nossos usuários ferramentas que não apenas atendam às suas necessidades atuais, mas antecipem e moldem o futuro do trabalho. Como sempre, estaremos desenvolvendo esses recursos em estreita colaboração com nossa comunidade de usuários, garantindo que cada aprimoramento adicione valor real e prático ao processo de gerenciamento de projetos.

## Conclusão

Então é isso!

Este foi um recurso divertido de implementar, e um de nossos primeiros passos em IA, junto com o [Resumo de Conteúdo com IA](/insights/ai-content-summarization) que lançamos anteriormente. Sabemos que a IA terá um papel cada vez maior no gerenciamento de projetos no futuro, e mal podemos esperar para lançar mais recursos inovadores aproveitando LLMs (Large Language Models) avançados.

Houve bastante coisa para pensar ao implementar isso, e estamos especialmente empolgados sobre como podemos aproveitar este recurso no futuro com o [mecanismo de automação de gerenciamento de projetos](/insights/benefits-project-management-automation) existente da Blue.

Também esperamos que tenha sido uma leitura interessante e que lhe dê uma visão de como pensamos sobre a engenharia dos recursos que você usa todos os dias.
