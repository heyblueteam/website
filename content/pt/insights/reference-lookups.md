---
title: Referência e campos personalizados de busca
description: Crie projetos interconectados no Blue sem esforço, transformando-o em uma única fonte de verdade para o seu negócio com os novos Campos de Referência e Busca.
category: "Product Updates"
date: 2023-11-01
---



Os projetos no Blue já são uma maneira poderosa de gerenciar os dados do seu negócio e avançar no trabalho.

Hoje, estamos dando o próximo passo lógico e permitindo que você interconecte seus dados *entre* projetos para a máxima flexibilidade e poder.

Interconectar projetos dentro do Blue transforma-o em uma única fonte de verdade para o seu negócio. Essa capacidade permite a criação de um conjunto de dados abrangente e interconectado, possibilitando um fluxo de dados contínuo e uma visibilidade aprimorada entre os projetos. Ao vincular projetos, as equipes podem alcançar uma visão unificada das operações, melhorando a tomada de decisões e a eficiência operacional.

## Um exemplo

Considere a ACME Company, que usa os campos personalizados de Referência e Busca do Blue para criar um ecossistema de dados interconectados em seus projetos de Clientes, Vendas e Inventário. Os registros de clientes no projeto de Clientes estão vinculados via campos de Referência a transações de vendas no projeto de Vendas. Essa vinculação permite que os campos de Busca puxem detalhes associados ao cliente, como números de telefone e status de conta, diretamente para cada registro de venda. Além disso, os itens de inventário vendidos são exibidos no registro de venda por meio de um campo de Busca que referencia os dados de Quantidade Vendida do projeto de Inventário. Por fim, as retiradas de inventário estão conectadas às vendas relevantes por meio de um campo de Referência no Inventário, apontando de volta para os registros de Vendas. Essa configuração fornece total visibilidade sobre qual venda acionou a remoção do inventário, criando uma visão integrada de 360 graus entre os projetos.

## Como Funcionam os Campos de Referência

Os campos personalizados de Referência permitem que você crie relacionamentos entre registros em diferentes projetos no Blue. Ao criar um campo de Referência, o Administrador do Projeto seleciona o projeto específico que fornecerá a lista de registros de referência. As opções de configuração incluem:

* **Seleção Única**: Permite escolher um registro de referência.
* **Seleção Múltipla**: Permite escolher vários registros de referência.
* **Filtragem**: Defina filtros para permitir que os usuários selecionem apenas registros que correspondam aos critérios de filtro.

Uma vez configurado, os usuários podem selecionar registros específicos no menu suspenso dentro do campo de Referência, estabelecendo um link entre os projetos.

## Estendendo campos de referência usando buscas

Os campos personalizados de Busca são usados para importar dados de registros em outros projetos, criando visibilidade unidirecional. Eles são sempre somente leitura e estão conectados a um campo personalizado de Referência específico. Quando um usuário seleciona um ou mais registros usando um campo personalizado de Referência, o campo personalizado de Busca mostrará dados desses registros. As buscas podem exibir dados como:

* Criado em
* Atualizado em
* Data de Vencimento
* Descrição
* Lista
* Tag
* Responsável
* Qualquer campo personalizado suportado do registro referenciado — incluindo outros campos de busca!

Por exemplo, imagine um cenário onde você tem três projetos: **Projeto A** é um projeto de vendas, **Projeto B** é um projeto de gerenciamento de inventário, e **Projeto C** é um projeto de relacionamento com o cliente. No Projeto A, você tem um campo personalizado de Referência que vincula registros de vendas aos registros de clientes correspondentes no Projeto C. No Projeto B, você tem um campo personalizado de Busca que importa informações do Projeto A, como a quantidade vendida. Dessa forma, quando um registro de venda é criado no Projeto A, as informações do cliente associadas a essa venda são automaticamente puxadas do Projeto C, e a quantidade vendida é automaticamente puxada do Projeto B. Isso permite que você mantenha todas as informações relevantes em um só lugar e visualize sem precisar criar dados duplicados ou atualizar manualmente registros entre os projetos.

Um exemplo da vida real disso é uma empresa de e-commerce que usa o Blue para gerenciar suas vendas, inventário e relacionamentos com clientes. Em seu projeto de **Vendas**, eles têm um campo personalizado de Referência que vincula cada registro de venda ao correspondente registro de **Cliente** em seu projeto de **Clientes**. Em seu projeto de **Inventário**, eles têm um campo personalizado de Busca que importa informações do projeto de Vendas, como a quantidade vendida, e as exibe no registro do item de inventário. Isso permite que eles vejam facilmente quais vendas estão impulsionando as remoções de inventário e mantenham seus níveis de inventário atualizados sem precisar atualizar manualmente registros entre os projetos.

## Conclusão

Imagine um mundo onde seus dados de projeto não estão isolados, mas fluem livremente entre os projetos, fornecendo insights abrangentes e impulsionando a eficiência. Esse é o poder dos campos de Referência e Busca do Blue. Ao permitir conexões de dados contínuas e fornecer visibilidade em tempo real entre os projetos, esses recursos transformam a forma como as equipes colaboram e tomam decisões. Seja gerenciando relacionamentos com clientes, rastreando vendas ou supervisionando inventário, os campos de Referência e Busca no Blue capacitam sua equipe a trabalhar de forma mais inteligente, rápida e eficaz. Mergulhe no mundo interconectado do Blue e veja sua produtividade disparar.

[Confira a documentação](https://documentation.blue.cc/custom-fields/reference) ou [inscreva-se e experimente você mesmo.](https://app.blue.cc)