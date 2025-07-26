---
title: Criando checklists reutilizáveis usando automações
description: Aprenda como criar automações de gerenciamento de projetos para checklists reutilizáveis.
category: "Best Practices"
date: 2024-07-08
---



Em muitos projetos e processos, pode ser necessário usar o mesmo checklist em vários registros ou tarefas.

No entanto, não é muito eficiente reescrever manualmente o checklist toda vez que você deseja adicioná-lo a um registro. É aqui que você pode aproveitar as [poderosas automações de gerenciamento de projetos](/platform/features/automations) para fazer isso automaticamente para você!

Como lembrete, as automações no Blue requerem duas coisas principais:

1. Um Gatilho — O que deve acontecer para iniciar a automação. Isso pode ser quando um registro recebe uma tag específica, quando ele se move para uma lista específica.
2. Uma ou mais Ações — Neste caso, seria a criação automática de um ou mais checklists.

Vamos começar pela ação primeiro, depois discutir os possíveis gatilhos que você pode usar.

## Ação de Automação de Checklist

Você pode criar uma nova automação e configurar um ou mais checklists para serem criados, conforme o exemplo abaixo:

![](/insights/checklist-automation.png)

Esses seriam os checklists que você deseja que sejam criados toda vez que você realizar a ação.

## Gatilhos de Automação de Checklist

Existem várias maneiras de você acionar a criação de seus checklists reutilizáveis. Aqui estão algumas opções populares:

- **Adicionar uma Tag Específica:** Você pode configurar a automação para ser acionada quando uma tag particular é adicionada a um registro. Por exemplo, quando a tag "Novo Projeto" é adicionada, ela pode automaticamente criar seu checklist de iniciação de projeto.
- **Atribuição de Registro:** A criação do checklist pode ser acionada quando um registro é atribuído a um indivíduo específico ou a qualquer pessoa. Isso é útil para checklists de integração ou procedimentos específicos de tarefas.
- **Mover para uma Lista Específica:** Quando um registro é movido para uma lista particular no seu quadro de projeto, isso pode acionar a criação de um checklist relevante. Por exemplo, mover um item para uma lista de "Garantia de Qualidade" pode acionar um checklist de QA.
- **Campo de Caixa de Seleção Personalizado:** Crie um campo de caixa de seleção personalizado e configure a automação para ser acionada quando essa caixa for marcada. Isso lhe dá controle manual sobre quando adicionar o checklist.
- **Campos Personalizados de Seleção Única ou Múltipla:** Você pode criar um campo personalizado de seleção única ou múltipla com várias opções. Cada opção pode ser vinculada a um modelo de checklist específico por meio de automações separadas. Isso permite um controle mais granular e a capacidade de ter vários modelos de checklist prontos para diferentes cenários.

Para aumentar o controle sobre quem pode acionar essas automações, você pode ocultar esses campos personalizados de certos usuários usando funções de usuário personalizadas. Isso garante que apenas administradores de projeto ou outros pessoal autorizado possam acionar essas opções.

Lembre-se, a chave para o uso eficaz de checklists reutilizáveis com automações é projetar seus gatilhos de forma cuidadosa. Considere o fluxo de trabalho da sua equipe, os tipos de projetos que você gerencia e quem deve ter a capacidade de iniciar diferentes processos. Com automações bem planejadas, você pode agilizar significativamente seu gerenciamento de projetos e garantir consistência em suas operações.

## Recursos Úteis

- [Documentação de Automação de Gerenciamento de Projetos](https://documentation.blue.cc/automations)
- [Documentação de Funções de Usuário Personalizadas](https://documentation.blue.cc/user-management/roles/custom-user-roles)
- [Documentação de Campos Personalizados](https://documentation.blue.cc/custom-fields)