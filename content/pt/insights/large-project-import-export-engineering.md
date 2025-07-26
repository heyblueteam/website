---
title: Escalando Importações e Exportações CSV para 250.000+ Registros
description: Descubra como a Blue escalou importações e exportações CSV em 10x usando Rust e arquitetura escalável, além de escolhas tecnológicas estratégicas em B2B SaaS.
category: "Engineering"
date: 2024-07-18
---


Na Blue, estamos [constantemente ultrapassando os limites](/platform/roadmap) do que é possível em software de gerenciamento de projetos. Ao longo dos anos, lançamos [centenas de recursos](/platform/changelog).

Nossa mais recente conquista de engenharia?

Uma reformulação completa de nosso sistema de [importação](https://documentation.blue.cc/integrations/csv-import) e [exportação](https://documentation.blue.cc/integrations/csv-export) CSV, melhorando drasticamente o desempenho e a escalabilidade.

Este post leva você aos bastidores de como enfrentamos esse desafio, as tecnologias que empregamos e os resultados impressionantes que alcançamos.

A coisa mais interessante aqui é que tivemos que sair do nosso [stack tecnológico](https://sop.blue.cc/product/technology-stack) típico para alcançar os resultados que desejávamos. Esta é uma decisão que deve ser tomada com cuidado, pois a repercussão a longo prazo pode ser severa em termos de dívida técnica e sobrecarga de manutenção a longo prazo.

<video autoplay loop muted playsinline>
  <source src="/videos/import-export-video.mp4" type="video/mp4">
</video>

## Escalando para Necessidades Empresariais

Nossa jornada começou com um pedido de um cliente corporativo na indústria de eventos. Este cliente usa a Blue como seu hub central para gerenciar vastas listas de eventos, locais e palestrantes, integrando-a perfeitamente ao seu site.

Para eles, a Blue não é apenas uma ferramenta — é a única fonte de verdade para toda a sua operação.

Embora sempre fiquemos orgulhosos ao ouvir que os clientes nos utilizam para necessidades tão críticas, também há uma grande quantidade de responsabilidade do nosso lado para garantir um sistema rápido e confiável.

À medida que este cliente escalava suas operações, enfrentou um obstáculo significativo: **importar e exportar grandes arquivos CSV contendo de 100.000 a 200.000+ registros.**

Isso estava além da capacidade do nosso sistema na época. Na verdade, nosso sistema anterior de importação/exportação já estava lutando com importações e exportações contendo mais de 10.000 a 20.000 registros! Portanto, 200.000+ registros estavam fora de questão.

Os usuários experimentavam tempos de espera frustrantemente longos e, em alguns casos, as importações ou exportações *falhavam em completar totalmente.* Isso afetava significativamente suas operações, pois dependiam de importações e exportações diárias para gerenciar certos aspectos de suas operações.

> Multi-inquilino é uma arquitetura onde uma única instância de software atende a múltiplos clientes (inquilinos). Embora seja eficiente, requer um gerenciamento cuidadoso de recursos para garantir que as ações de um inquilino não impactem negativamente os outros.

E essa limitação não estava afetando apenas este cliente em particular.

Devido à nossa arquitetura multi-inquilino — onde múltiplos clientes compartilham a mesma infraestrutura — uma única importação ou exportação intensiva em recursos poderia potencialmente desacelerar as operações de outros usuários, o que na prática frequentemente acontecia.

Como de costume, fizemos uma análise de construir versus comprar, para entender se deveríamos gastar tempo para atualizar nosso próprio sistema ou comprar um sistema de outra pessoa. Olhamos para várias possibilidades.

O fornecedor que se destacou foi um provedor de SaaS chamado [Flatfile](https://flatfile.com/). O sistema e as capacidades deles pareciam exatamente o que precisávamos.

Mas, após revisar seu [preço](https://flatfile.com/pricing/), decidimos que isso acabaria sendo uma solução extremamente cara para uma aplicação da nossa escala — *$2/arquivo começa a somar rapidamente!* — e era melhor estender nosso motor de importação/exportação CSV embutido.

Para enfrentar esse desafio, tomamos uma decisão ousada: introduzir Rust em nosso stack tecnológico primário de Javascript. Esta linguagem de programação de sistemas, conhecida por seu desempenho e segurança, era a ferramenta perfeita para nossas necessidades críticas de análise de CSV e mapeamento de dados.

Aqui está como abordamos a solução.

### Introduzindo Serviços em Segundo Plano

A base de nossa solução foi a introdução de serviços em segundo plano para lidar com tarefas intensivas em recursos. Essa abordagem nos permitiu descarregar o processamento pesado de nosso servidor principal, melhorando significativamente o desempenho geral do sistema.

Nossa arquitetura de serviços em segundo plano foi projetada com escalabilidade em mente. Como todos os componentes de nossa infraestrutura, esses serviços se autoescalam com base na demanda.

Isso significa que durante os períodos de pico, quando várias grandes importações ou exportações estão sendo processadas simultaneamente, o sistema automaticamente aloca mais recursos para lidar com a carga aumentada. Por outro lado, durante períodos mais tranquilos, ele reduz a escala para otimizar o uso de recursos.

Essa arquitetura escalável de serviços em segundo plano beneficiou a Blue não apenas para importações e exportações CSV. Com o tempo, movemos um número substancial de recursos para serviços em segundo plano para aliviar a carga de nossos servidores principais:

- **[Cálculos de Fórmula](https://documentation.blue.cc/custom-fields/formula)**: Descarrega operações matemáticas complexas para garantir atualizações rápidas de campos derivados sem impactar o desempenho do servidor principal.
- **[Painel/Gráficos](/platform/features/dashboards)**: Processa grandes conjuntos de dados em segundo plano para gerar visualizações atualizadas sem desacelerar a interface do usuário.
- **[Índice de Pesquisa](https://documentation.blue.cc/projects/search)**: Atualiza continuamente o índice de pesquisa em segundo plano, garantindo resultados de pesquisa rápidos e precisos sem impactar o desempenho do sistema.
- **[Cópia de Projetos](https://documentation.blue.cc/projects/copying-projects)**: Lida com a replicação de grandes projetos complexos em segundo plano, permitindo que os usuários continuem trabalhando enquanto a cópia está sendo criada.
- **[Automatizações de Gerenciamento de Projetos](/platform/features/automations)**: Executa fluxos de trabalho automatizados definidos pelo usuário em segundo plano, garantindo ações pontuais sem bloquear outras operações.
- **[Registros Repetidos](https://documentation.blue.cc/records/repeat)**: Gera tarefas ou eventos recorrentes em segundo plano, mantendo a precisão do cronograma sem sobrecarregar o aplicativo principal.
- **[Campos Personalizados de Duração de Tempo](https://documentation.blue.cc/custom-fields/duration)**: Calcula e atualiza continuamente a diferença de tempo entre dois eventos na Blue, fornecendo dados de duração em tempo real sem impactar a capacidade de resposta do sistema.

## Novo Módulo Rust para Análise de Dados

O coração de nossa solução de processamento CSV é um módulo Rust personalizado. Embora isso tenha marcado nossa primeira incursão fora de nosso stack tecnológico principal de Javascript, a decisão de usar Rust foi impulsionada por seu desempenho excepcional em operações concorrentes e tarefas de processamento de arquivos.

Os pontos fortes do Rust alinham-se perfeitamente com as demandas de análise de CSV e mapeamento de dados. Suas abstrações de custo zero permitem programação de alto nível sem sacrificar o desempenho, enquanto seu modelo de propriedade garante segurança de memória sem a necessidade de coleta de lixo. Esses recursos tornam o Rust particularmente apto a lidar com grandes conjuntos de dados de forma eficiente e segura.

Para a análise de CSV, aproveitamos a crate csv do Rust, que oferece leitura e escrita de dados CSV de alto desempenho. Combinamos isso com lógica de mapeamento de dados personalizada para garantir uma integração perfeita com as estruturas de dados da Blue.

A curva de aprendizado para Rust foi íngreme, mas gerenciável. Nossa equipe dedicou cerca de duas semanas para um aprendizado intensivo sobre isso.

As melhorias foram impressionantes:

![](/insights/import-export.png)

Nosso novo sistema pode processar a mesma quantidade de registros que nosso antigo sistema poderia processar em 15 minutos em cerca de 30 segundos.

## Interação entre Servidor Web e Banco de Dados

Para o componente do servidor web de nossa implementação em Rust, escolhemos o Rocket como nosso framework. O Rocket se destacou por sua combinação de desempenho e recursos amigáveis ao desenvolvedor. Sua tipagem estática e verificação em tempo de compilação alinham-se bem com os princípios de segurança do Rust, ajudando-nos a detectar problemas potenciais no início do processo de desenvolvimento.

No front do banco de dados, optamos pelo SQLx. Esta biblioteca SQL assíncrona para Rust oferece várias vantagens que a tornaram ideal para nossas necessidades:

- SQL seguro por tipo: O SQLx nos permite escrever SQL bruto com consultas verificadas em tempo de compilação, garantindo segurança de tipo sem sacrificar desempenho.
- Suporte assíncrono: Isso se alinha bem com o Rocket e nossa necessidade de operações de banco de dados eficientes e não bloqueantes.
- Independente de banco de dados: Embora usemos principalmente [AWS Aurora](https://aws.amazon.com/rds/aurora/), que é compatível com MySQL, o suporte do SQLx para múltiplos bancos de dados nos dá flexibilidade para o futuro, caso decidamos mudar.

## Otimização de Lotes

Nossa jornada para a configuração de lotes ideal foi uma de rigorosos testes e análise cuidadosa. Realizamos extensos benchmarks com várias combinações de transações concorrentes e tamanhos de lotes, medindo não apenas a velocidade bruta, mas também a utilização de recursos e a estabilidade do sistema.

O processo envolveu a criação de conjuntos de dados de teste de tamanhos e complexidades variadas, simulando padrões de uso do mundo real. Em seguida, executamos esses conjuntos de dados em nosso sistema, ajustando o número de transações concorrentes e o tamanho do lote para cada execução.

Após analisar os resultados, descobrimos que processar 5 transações concorrentes com um tamanho de lote de 500 registros proporcionava o melhor equilíbrio entre velocidade e utilização de recursos. Essa configuração nos permite manter uma alta taxa de transferência sem sobrecarregar nosso banco de dados ou consumir memória excessiva.

Curiosamente, descobrimos que aumentar a concorrência além de 5 transações não resultava em ganhos significativos de desempenho e, às vezes, levava a um aumento na contenção do banco de dados. Da mesma forma, tamanhos de lote maiores melhoravam a velocidade bruta, mas à custa de maior uso de memória e tempos de resposta mais longos para importações/exportações de pequeno a médio porte.

## Exportações CSV via Links de Email

A última parte de nossa solução aborda o desafio de entregar grandes arquivos exportados aos usuários. Em vez de fornecer um download direto de nosso aplicativo web, o que poderia levar a problemas de timeout e aumentar a carga do servidor, implementamos um sistema de links de download enviados por email.

Quando um usuário inicia uma grande exportação, nosso sistema processa o pedido em segundo plano. Uma vez completo, em vez de manter a conexão aberta ou armazenar o arquivo em nossos servidores web, fazemos o upload do arquivo para um local de armazenamento temporário e seguro. Em seguida, geramos um link de download único e seguro e o enviamos por email ao usuário.

Esses links de download são válidos por 2 horas, equilibrando conveniência para o usuário e segurança da informação. Esse prazo dá aos usuários ampla oportunidade de recuperar seus dados, garantindo que informações sensíveis não fiquem acessíveis indefinidamente.

A segurança desses links de download foi uma prioridade em nosso design. Cada link é:

- Único e gerado aleatoriamente, tornando praticamente impossível adivinhá-lo
- Válido por apenas 2 horas
- Criptografado em trânsito, garantindo a segurança dos dados durante o download

Essa abordagem oferece vários benefícios:

- Reduz a carga em nossos servidores web, já que eles não precisam lidar com downloads de arquivos grandes diretamente
- Melhora a experiência do usuário, especialmente para usuários com conexões de internet mais lentas que poderiam enfrentar problemas de timeout do navegador com downloads diretos
- Fornece uma solução mais confiável para exportações muito grandes que poderiam exceder os limites típicos de timeout da web

O feedback dos usuários sobre esse recurso tem sido extremamente positivo, com muitos apreciando a flexibilidade que ele oferece na gestão de grandes exportações de dados.

## Exportando Dados Filtrados

A outra melhoria óbvia foi permitir que os usuários exportassem apenas os dados que já estavam filtrados em sua visualização de projeto. Isso significa que, se houver uma tag ativa "prioridade", apenas os registros que possuem essa tag acabariam na exportação CSV. Isso significa menos tempo manipulando dados no Excel para filtrar coisas que não são importantes e também nos ajuda a reduzir o número de linhas a processar.

## Olhando para o Futuro

Embora não tenhamos planos imediatos de expandir nosso uso de Rust, este projeto nos mostrou o potencial dessa tecnologia para operações críticas de desempenho. É uma opção empolgante que agora temos em nosso conjunto de ferramentas para futuras necessidades de otimização. Essa reformulação de importação e exportação CSV se alinha perfeitamente com o compromisso da Blue com a escalabilidade.

Estamos dedicados a fornecer uma plataforma que cresce com nossos clientes, lidando com suas crescentes necessidades de dados sem comprometer o desempenho.

A decisão de introduzir Rust em nosso stack tecnológico não foi tomada levianamente. Isso levantou uma questão importante que muitas equipes de engenharia enfrentam: Quando é apropriado sair do seu stack tecnológico principal e quando você deve permanecer com ferramentas familiares?

Não há uma resposta única, mas na Blue, desenvolvemos uma estrutura para tomar essas decisões cruciais:

- **Abordagem Focada no Problema:** Sempre começamos definindo claramente o problema que estamos tentando resolver. Neste caso, precisávamos melhorar drasticamente o desempenho das importações e exportações CSV para grandes conjuntos de dados.
- **Exaurindo Soluções Existentes:** Antes de olhar para fora de nosso stack principal, exploramos minuciosamente o que pode ser alcançado com nossas tecnologias existentes. Isso muitas vezes envolve perfilamento, otimização e repensar nossa abordagem dentro de restrições familiares.
- **Quantificando o Ganho Potencial:** Se estamos considerando uma nova tecnologia, precisamos ser capazes de articular claramente e, idealmente, quantificar os benefícios. Para nosso projeto CSV, projetamos melhorias de ordem de magnitude na velocidade de processamento.
- **Avaliação dos Custos:** Introduzir uma nova tecnologia não se trata apenas do projeto imediato. Consideramos os custos a longo prazo:
  - Curva de aprendizado para a equipe
  - Manutenção e suporte contínuos
  - Complicações potenciais na implantação e operações
  - Impacto na contratação e composição da equipe
- **Contenção e Integração:** Se introduzirmos uma nova tecnologia, buscamos contê-la a uma parte específica e bem definida de nosso sistema. Também garantimos que temos um plano claro de como ela se integrará com nosso stack existente.
- **Preparação para o Futuro:** Consideramos se essa escolha tecnológica abre oportunidades futuras ou se pode nos colocar em uma situação difícil.

Um dos principais riscos de adotar novas tecnologias com frequência é acabar com o que chamamos de *"zoológico tecnológico"* - um ecossistema fragmentado onde diferentes partes de sua aplicação são escritas em diferentes linguagens ou frameworks, exigindo uma ampla gama de habilidades especializadas para manter.

## Conclusão

Este projeto exemplifica a abordagem da Blue em engenharia: *não temos medo de sair da nossa zona de conforto e adotar novas tecnologias quando isso significa oferecer uma experiência significativamente melhor para nossos usuários.*

Ao reimaginar nosso processo de importação e exportação CSV, não apenas resolvemos uma necessidade premente para um cliente corporativo, mas melhoramos a experiência para todos os nossos usuários que lidam com grandes conjuntos de dados.

À medida que continuamos a ultrapassar os limites do que é possível em [software de gerenciamento de projetos](/solutions/use-case/project-management), estamos empolgados para enfrentar mais desafios como este.

Fique atento para mais [aprofundamentos na engenharia que alimenta a Blue!](/insights/engineering-blog)