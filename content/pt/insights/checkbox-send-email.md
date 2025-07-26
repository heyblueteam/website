---
title: Automação de gerenciamento de projetos — e-mails para partes interessadas.
description: Muitas vezes, você quer ter controle sobre suas automações de gerenciamento de projetos
category: "Product Updates"
date: 2024-07-08
---


Nós já cobrimos como [criar automações de e-mail antes.](/insights/email-automations)

No entanto, muitas vezes há partes interessadas em projetos que só precisam ser alertadas quando há algo *realmente* importante.

Não seria bom se houvesse uma automação de gerenciamento de projetos onde você, como gerente de projeto, pudesse ter controle sobre *exatamente* quando notificar uma parte interessada chave com o pressionar de um botão?

Bem, acontece que com o Blue, você pode fazer precisamente isso!

Hoje vamos aprender como criar uma automação de gerenciamento de projetos realmente útil:

Uma caixa de seleção que notifica automaticamente uma ou mais partes interessadas chave, fornecendo a elas todo o contexto importante sobre o que você está notificando. Como um ponto bônus, também aprenderemos como restringir essa capacidade para que apenas certos membros do seu projeto possam acionar essa notificação por e-mail.

Isso vai parecer algo assim quando você terminar:

![](/insights/checkbox-email-automation.png)

E apenas marcando essa caixa de seleção, você poderá acionar uma automação de gerenciamento de projetos para enviar um e-mail de notificação personalizado para as partes interessadas.

Vamos passo a passo.

## 1. Crie seu campo personalizado de caixa de seleção

Isso é muito fácil, você pode conferir nossa [documentação detalhada](https://documentation.blue.cc/custom-fields/introduction#creating-custom-fields) sobre como criar campos personalizados.

Certifique-se de nomear este campo algo óbvio que você se lembrará, como “notificar a gerência” ou “notificar partes interessadas”.

## 2. Crie seu gatilho de automação de gerenciamento de projetos.

Na visualização de registros do seu projeto, clique no pequeno robô no canto superior direito para abrir as configurações de automação:

<video autoplay loop muted playsinline>
  <source src="/videos/notify-stakeholders-automation-setup.mp4" type="video/mp4">
</video>

## 3. Crie sua ação de automação de gerenciamento de projetos.

Neste caso, nossa ação será enviar uma notificação por e-mail personalizada para um ou mais endereços de e-mail. É bom notar aqui que essas pessoas **não** precisam estar no Blue para receber esses e-mails, você pode enviar e-mails para *qualquer* endereço de e-mail.

Você pode aprender mais em nosso [guia de documentação detalhada sobre como configurar automações de e-mail](https://documentation.blue.cc/automations/actions/email-automations)

Seu resultado final deve parecer algo assim:

![](/insights/email-automation-example.png)

## 4. Bônus: Restringir acesso à caixa de seleção.

Você pode usar [funções de usuário personalizadas no Blue](/platform/features/user-permissions) para restringir o acesso aos campos personalizados de caixa de seleção, garantindo que apenas membros autorizados da equipe possam acionar notificações por e-mail.

O Blue permite que Administradores de Projetos definam funções e atribuam permissões a grupos de usuários. Este sistema é crucial para manter o controle sobre quem pode interagir com elementos específicos do seu projeto, incluindo campos personalizados como a caixa de seleção de notificação.

1. Navegue até a seção de Gerenciamento de Usuários no Blue e selecione "Funções de Usuário Personalizadas."
2. Crie uma nova função fornecendo um nome descritivo e uma descrição opcional.
3. Dentro das permissões da função, localize a seção de Acesso a Campos Personalizados.
4. Especifique se a função pode visualizar ou editar o campo personalizado da caixa de seleção. Por exemplo, restrinja o acesso de edição a funções como "Administrador de Projeto" enquanto permite que uma função personalizada recém-criada gerencie este campo.
5. Atribua a função recém-criada aos usuários ou grupos de usuários apropriados. Isso garante que apenas os indivíduos designados tenham a capacidade de interagir com a caixa de seleção de notificação.

[Leia mais em nosso site oficial de documentação.](https://documentation.blue.cc/user-management/roles/custom-user-roles)

Ao implementar essas funções personalizadas, você melhora a segurança e a integridade dos seus processos de gerenciamento de projetos. Apenas membros autorizados da equipe podem acionar notificações por e-mail críticas, garantindo que as partes interessadas recebam atualizações importantes sem alertas desnecessários.

## Conclusão

Ao implementar a automação de gerenciamento de projetos descrita acima, você ganha controle preciso sobre quando e como notificar partes interessadas chave. Essa abordagem garante que atualizações importantes sejam comunicadas efetivamente, sem sobrecarregar suas partes interessadas com informações desnecessárias. Utilizando os recursos de campo personalizado e automação do Blue, você pode simplificar seu processo de gerenciamento de projetos, melhorar a comunicação e manter um alto nível de eficiência.

Com apenas uma simples caixa de seleção, você pode acionar notificações por e-mail personalizadas adaptadas às necessidades do seu projeto, garantindo que as pessoas certas sejam informadas no momento certo. Além disso, a capacidade de restringir essa funcionalidade a membros específicos da equipe adiciona uma camada extra de controle e segurança.

Comece a aproveitar esse recurso poderoso no Blue hoje para manter suas partes interessadas informadas e seus projetos funcionando sem problemas. Para passos mais detalhados e opções de personalização adicionais, consulte os links de documentação fornecidos. Boa automação!