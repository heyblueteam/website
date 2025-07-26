---
title:  Como usamos o Blue para construir o Blue. 
description: Aprenda como usamos nossa própria plataforma de gerenciamento de projetos para construir nossa plataforma de gerenciamento de projetos!
category: "CEO Blog"
date: 2024-08-07
---



Você está prestes a receber um tour interno de como o Blue constrói o Blue.

No Blue, consumimos nossa própria comida de cachorro.

Isso significa que usamos o Blue para *construir* o Blue.

Esse termo que soa estranho, frequentemente referido como "dogfooding", é frequentemente atribuído a Paul Maritz, um gerente da Microsoft na década de 1980. Ele supostamente enviou um e-mail com o assunto *"Comendo nossa própria comida de cachorro"* para encorajar os funcionários da Microsoft a usarem os produtos da empresa.

A ideia de usar suas próprias ferramentas para construir suas ferramentas é que isso leva a um ciclo de feedback positivo.

A ideia de usar suas próprias ferramentas para construir suas ferramentas leva a um ciclo de feedback positivo, criando inúmeros benefícios:

- **Ajuda-nos a identificar rapidamente problemas de usabilidade do mundo real.** Como usamos o Blue diariamente, encontramos os mesmos desafios que nossos usuários podem enfrentar, permitindo que os abordemos proativamente.
- **Acelera a descoberta de bugs.** O uso interno geralmente revela bugs antes que eles cheguem aos nossos clientes, melhorando a qualidade geral do produto.
- **Aumenta nossa empatia pelos usuários finais.** Nossa equipe ganha experiência em primeira mão sobre os pontos fortes e fracos do Blue, ajudando-nos a tomar decisões mais centradas no usuário.
- **Impulsiona uma cultura de qualidade dentro da nossa organização.** Quando todos usam o produto, há um interesse compartilhado em sua excelência.
- **Fomenta a inovação.** O uso regular frequentemente gera ideias para novos recursos ou melhorias, mantendo o Blue na vanguarda.


[Já falamos antes sobre por que não temos uma equipe de teste dedicada](/insights/open-beta) e esta é mais uma razão.

Se houver bugs em nosso sistema, quase sempre os encontramos em nosso uso diário constante da plataforma. E isso também cria uma função de pressão para corrigi-los, já que os acharemos muito irritantes, pois provavelmente somos um dos principais usuários do Blue!

Essa abordagem demonstra nosso compromisso com o produto. Ao confiar no Blue, mostramos aos nossos clientes que realmente acreditamos no que estamos construindo. Não é apenas um produto que vendemos – é uma ferramenta da qual dependemos todos os dias.

## Processo Principal

Temos um projeto no Blue, apropriadamente chamado "Produto".

**Tudo** relacionado ao nosso desenvolvimento de produto é rastreado aqui. Feedback dos clientes, bugs, ideias de recursos, trabalho em andamento, e assim por diante. A ideia de ter um projeto onde rastreamos tudo é que isso [promove um melhor trabalho em equipe.](/insights/great-teamwork)

Cada registro é um recurso ou parte de um recurso. É assim que passamos de "não seria legal se..." para "confira este novo recurso incrível!"

O projeto tem as seguintes listas:

- **Ideias/Feedback**: Esta é uma lista de ideias da equipe ou feedback de clientes com base em chamadas ou trocas de e-mail. Sinta-se à vontade para adicionar quaisquer ideias aqui! Nesta lista, ainda não decidimos que construiremos algum desses recursos, mas revisamos regularmente em busca de ideias que queremos explorar mais a fundo.
- **Backlog (Longo Prazo)**: É aqui que os recursos da lista de Ideias/Feedback vão se decidirmos que seriam uma boa adição ao Blue.
- **{Trimestre Atual}**: Isso é tipicamente estruturado como "Qx AAAA" e mostra nossas prioridades trimestrais.
- **Bugs**: Esta é uma lista de bugs conhecidos relatados pela equipe ou clientes. Bugs adicionados aqui terão automaticamente a tag "Bug" adicionada.
- **Especificações**: Esses recursos estão atualmente sendo especificados. Nem todo recurso requer uma especificação ou design; isso depende do tamanho esperado do recurso e do nível de confiança que temos em relação a casos extremos e complexidade.
- **Backlog de Design**: Este é o backlog para os designers; sempre que eles terminam algo que está em andamento, podem escolher qualquer item desta lista.
- **Design em Andamento**: Estes são os recursos atuais que os designers estão projetando.
- **Revisão de Design**: É aqui que estão os recursos cujos designs estão sendo revisados.
- **Backlog (Curto Prazo)**: Esta é uma lista de recursos que provavelmente começaremos a trabalhar nas próximas semanas. É aqui que ocorrem as atribuições. O CEO e o Chefe de Engenharia decidem quais recursos são atribuídos a qual engenheiro com base na experiência anterior e na carga de trabalho. [Os membros da equipe podem então puxar esses itens para o Em Andamento](/insights/push-vs-pull-kanban) assim que concluírem seu trabalho atual.
- **Em Andamento**: Estes são recursos que estão atualmente sendo desenvolvidos.
- **Revisão de Código**: Uma vez que um recurso termina o desenvolvimento, ele passa por uma revisão de código. Então, ele será movido de volta para "Em Andamento" se ajustes forem necessários ou implantado no ambiente de Desenvolvimento.
- **Dev**: Estes são todos os recursos atualmente no ambiente de Desenvolvimento. Outros membros da equipe e certos clientes podem revisar esses recursos.
- **Beta**: Estes são todos os recursos atualmente no [ambiente Beta](https://beta.app.blue.cc). Muitos clientes usam isso como sua plataforma diária do Blue e também fornecerão feedback.
- **Produção**: Quando um recurso chega à produção, é considerado concluído.

Às vezes, ao desenvolver um recurso, percebemos que certos sub-recursos são mais difíceis de implementar do que inicialmente esperado, e podemos optar por não fazê-los na versão inicial que implantamos para os clientes. Nesse caso, podemos criar um novo registro com um nome seguindo o formato "{NomeDoRecurso} V2" e incluir todos os sub-recursos como itens de checklist.

## Tags

- **Móvel**: Isso significa que o recurso é específico para nossos aplicativos iOS, Android ou iPad.
- **{NomeDoClienteEnterprise}**: Um recurso está sendo especificamente construído para um cliente enterprise. O rastreamento é importante, pois geralmente há acordos comerciais adicionais para cada recurso.
- **Bug**: Isso significa que este é um bug que requer correção.
- **Fast-Track**: Isso significa que esta é uma Mudança Fast-Track que não precisa passar pelo ciclo completo de lançamento, conforme descrito acima.
- **Principal**: Este é um desenvolvimento de recurso importante. Geralmente é reservado para grandes trabalhos de infraestrutura, grandes atualizações de dependência e novos módulos significativos dentro do Blue.
- **IA**: Este recurso contém um componente de inteligência artificial.
- **Segurança**: Isso significa que uma implicação de segurança deve ser revisada ou um patch é necessário.

A tag fast-track é particularmente interessante. Isso é reservado para atualizações menores, menos complexas que não requerem nosso ciclo completo de lançamento e que queremos enviar aos clientes dentro de 24-48 horas.

Mudanças fast-track são tipicamente ajustes menores que podem melhorar significativamente a experiência do usuário sem alterar a funcionalidade principal. Pense em corrigir erros de digitação na interface do usuário, ajustar o preenchimento dos botões ou adicionar novos ícones para melhor orientação visual. Essas são as mudanças que, embora pequenas, podem fazer uma grande diferença na forma como os usuários percebem e interagem com nosso produto. Elas também são irritantes se demorarem para serem enviadas!

Nosso processo fast-track é simples.

Começa com a criação de um novo branch a partir do principal, implementando as mudanças e, em seguida, criando solicitações de mesclagem para cada branch de destino - Dev, Beta e Produção. Geramos um link de pré-visualização para revisão, garantindo que mesmo essas pequenas mudanças atendam aos nossos padrões de qualidade. Uma vez aprovadas, as mudanças são mescladas simultaneamente em todos os branches, mantendo nossos ambientes em sincronia.

## Campos Personalizados

Não temos muitos campos personalizados em nosso projeto de Produto.

- **Especificações**: Isso vincula a um documento do Blue que contém a especificação para aquele recurso específico. Isso nem sempre é feito, pois depende da complexidade do recurso.
- **MR**: Este é o link para a Solicitação de Mesclagem no [Gitlab](https://gitlab.com) onde hospedamos nosso código.
- **Link de Pré-visualização**: Para recursos que mudam principalmente o front-end, podemos criar uma URL única que tenha essas mudanças para cada commit, para que possamos revisar facilmente as alterações.
- **Líder**: Este campo nos diz qual engenheiro sênior está liderando a revisão de código. Isso garante que cada recurso receba a atenção especializada que merece e que sempre haja uma pessoa de referência clara para perguntas ou preocupações.

## Checklists

Durante nossas demonstrações semanais, colocaremos o feedback discutido em uma checklist chamada "feedback" e haverá também outra checklist que contém o principal [WBS (Estrutura de Divisão do Trabalho)](/insights/simple-work-breakdown-structure) do recurso, para que possamos facilmente identificar o que está feito e o que ainda precisa ser feito.

## Conclusão

E é isso!

Achamos que às vezes as pessoas ficam surpresas com a simplicidade do nosso processo, mas acreditamos que processos simples são frequentemente muito superiores a processos excessivamente complexos que você não consegue entender facilmente.

Essa simplicidade é intencional. Ela nos permite permanecer ágeis, responder rapidamente às necessidades dos clientes e manter toda a nossa equipe alinhada.

Ao usar o Blue para construir o Blue, não estamos apenas desenvolvendo um produto – estamos vivenciando-o.

Então, da próxima vez que você estiver usando o Blue, lembre-se: você não está apenas usando um produto que construímos. Você está usando um produto do qual dependemos pessoalmente todos os dias.

E isso faz toda a diferença.