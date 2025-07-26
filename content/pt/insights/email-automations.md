---
title: Como criar automações de e-mail personalizadas
description: Notificações de e-mail personalizadas são um recurso incrivelmente poderoso no Blue que pode ajudar a manter o trabalho avançando e garantir que a comunicação esteja em piloto automático.
category: "Product Updates"
---



As automações de e-mail no Blue são uma [automação poderosa de gerenciamento de projetos](/platform/features/automations) para agilizar a comunicação, garantir [ótima colaboração](/insights/great-teamwork) e manter os projetos em andamento. Ao aproveitar os dados armazenados em seus registros, você pode enviar automaticamente e-mails personalizados quando certos gatilhos ocorrerem, como a criação de um novo registro ou uma tarefa se tornar atrasada.

Neste artigo, exploraremos como configurar e usar automações de e-mail no Blue.

## Configurando Automação de E-mail.

Criar uma automação de e-mail no Blue é um processo simples. Primeiro, selecione o gatilho que iniciará o e-mail automatizado. Alguns gatilhos comuns incluem:

- Um novo registro é criado
- Uma tag é adicionada a um registro
- Um registro é movido para outra lista

Em seguida, configure os detalhes do e-mail, incluindo:

- Nome do remetente e endereço de resposta
- Endereço de destino (pode ser estático ou extraído dinamicamente de um campo de e-mail personalizado)
- Endereços de CC ou CCO (opcional)

![](/insights/email-automations-image.webp)

Um dos principais benefícios das automações de e-mail no Blue é a capacidade de personalizar o conteúdo usando tags de mesclagem. Ao personalizar o assunto e o corpo do e-mail, você pode inserir tags de mesclagem que fazem referência a dados específicos do registro, como o nome do registro ou valores de campos personalizados. Basta usar a sintaxe {chave} para inserir as tags de mesclagem.

Você também pode incluir anexos de arquivos arrastando e soltando-os no e-mail ou usando o ícone de anexo. Arquivos de campos personalizados de arquivo podem ser anexados automaticamente se estiverem abaixo de 10MB.

Antes de finalizar sua automação de e-mail, é recomendável enviar um e-mail de teste para você ou um colega para garantir que tudo esteja funcionando conforme o esperado.

## Casos de Uso e Exemplos

As automações de e-mail no Blue podem ser usadas para uma variedade de propósitos. Aqui estão alguns exemplos:

1. Enviar um e-mail de confirmação quando um cliente enviar uma solicitação por meio de um formulário de entrada. Defina o gatilho para enviar um e-mail quando um novo registro for criado através do formulário e certifique-se de incluir um campo de e-mail no formulário para capturar o endereço do cliente.
2. Notificar um responsável quando uma nova tarefa de alta prioridade for criada. Defina o gatilho para enviar um e-mail quando uma tag de “Prioridade” for adicionada a um registro e use a tag de mesclagem {Responsável} para enviar automaticamente o e-mail ao usuário designado.
3. Enviar uma pesquisa a um cliente após um chamado de suporte ser marcado como resolvido. Defina o gatilho para enviar um e-mail quando um registro for marcado como concluído e movido para a lista “Concluído”. Inclua o e-mail do cliente em um campo personalizado e forneça informações detalhadas sobre o problema resolvido no corpo do e-mail.
4. Automatizar um programa de recrutamento enviando e-mails de confirmação aos candidatos. Defina o gatilho para enviar um e-mail quando uma candidatura for enviada através de um formulário e adicionada à lista “Recebido”. Capture o e-mail do candidato no formulário e use-o para enviar uma resposta de agradecimento.

## Benefícios das Automação de E-mail

As automações de e-mail no Blue oferecem vários benefícios importantes:

- Comunicação personalizada por meio do uso de tags de mesclagem e dados de campos personalizados
- Notificações automáticas que reduzem o trabalho manual e garantem atualizações pontuais
- Fluxos de trabalho estruturados e orientados por dados que impulsionam os projetos com base nos dados dos registros

## Conclusão 

As automações de e-mail no Blue são uma ferramenta valiosa para agilizar a comunicação e manter os projetos no caminho certo. Ao aproveitar gatilhos, tags de mesclagem e dados de campos personalizados, você pode criar e-mails automatizados e personalizados que aumentam a produtividade de sua equipe e garantem que atualizações importantes nunca sejam perdidas. Com uma ampla gama de casos de uso e configuração fácil, as automações de e-mail são um recurso indispensável para qualquer usuário do Blue que deseja otimizar seu fluxo de trabalho.