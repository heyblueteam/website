---
title: Pesquisa em tempo real
description: A Blue apresenta um novo mecanismo de busca ultrarrápido que retorna resultados em todos os seus projetos em milissegundos, permitindo que você mude de contexto em um piscar de olhos.
category: "Product Updates"
date: 2024-03-01
---



Estamos empolgados em anunciar o lançamento do nosso novo mecanismo de busca, projetado para revolucionar a forma como você encontra informações dentro da Blue. Uma funcionalidade de busca eficiente é crucial para um gerenciamento de projetos sem interrupções, e nossa atualização mais recente garante que você possa acessar seus dados mais rápido do que nunca.

Nosso novo mecanismo de busca permite que você procure por todos os comentários, arquivos, registros, campos personalizados, descrições e listas de verificação. Se você precisa encontrar um comentário específico feito em um projeto, localizar rapidamente um arquivo ou pesquisar um registro ou campo em particular, nosso mecanismo de busca fornece resultados ultrarrápidos.

À medida que as ferramentas se aproximam de uma responsividade de 50-100ms, elas tendem a desaparecer e se misturar ao fundo, proporcionando uma experiência do usuário quase invisível. Para contextualizar, um piscar de olhos humano leva aproximadamente 60-120ms, então 50ms é, na verdade, mais rápido do que um piscar de olhos! Esse nível de responsividade permite que você interaja com a Blue sem nem perceber que ela está lá, liberando você para se concentrar no trabalho real em mãos. Ao aproveitar esse nível de desempenho, nosso novo mecanismo de busca garante que você possa acessar rapidamente as informações de que precisa, sem que isso interfira no seu fluxo de trabalho.

Para alcançar nosso objetivo de busca ultrarrápida, utilizamos as mais recentes tecnologias de código aberto. Nosso mecanismo de busca é construído sobre o MeiliSearch, um popular serviço de busca como serviço que utiliza processamento de linguagem natural e busca vetorial para encontrar rapidamente resultados relevantes. Além disso, implementamos armazenamento em memória, que nos permite armazenar dados frequentemente acessados na RAM, reduzindo o tempo necessário para retornar os resultados da busca. Essa combinação de MeiliSearch e armazenamento em memória permite que nosso mecanismo de busca forneça resultados em milissegundos, tornando possível que você encontre rapidamente o que precisa sem ter que pensar na tecnologia subjacente.

A nova barra de busca está convenientemente localizada na barra de navegação, permitindo que você comece a buscar imediatamente. Para uma experiência de busca mais detalhada, basta pressionar a tecla Tab enquanto busca para acessar a página de busca completa. Além disso, você pode ativar rapidamente a função de busca de qualquer lugar usando o atalho CMD/Ctrl+K, facilitando ainda mais encontrar o que você precisa.

<video autoplay loop muted playsinline>
  <source src="/videos/search-demo.mp4" type="video/mp4">
</video>


## Desenvolvimentos Futuros

Isso é apenas o começo. Agora que temos uma infraestrutura de busca de próxima geração, podemos fazer algumas coisas realmente interessantes no futuro.

A próxima novidade será a busca semântica, que é uma melhoria significativa em relação à busca típica por palavras-chave. Permita-nos explicar.

Esse recurso permitirá que o mecanismo de busca entenda o contexto de suas consultas. Por exemplo, buscar por "mar" irá recuperar documentos relevantes mesmo que a frase exata não seja utilizada. Você pode estar pensando "mas eu digitei 'oceano' em vez disso!" - e você está certo. O mecanismo de busca também entenderá a similaridade entre "mar" e "oceano", e retornará documentos relevantes mesmo que a frase exata não seja utilizada. Esse recurso é particularmente útil ao buscar documentos que contêm termos técnicos, siglas ou apenas palavras comuns que têm várias variações ou erros de digitação.

Outro recurso que está por vir é a capacidade de buscar imagens pelo seu conteúdo. Para alcançar isso, estaremos processando cada imagem em seu projeto, criando uma incorporação para cada uma. Em termos de alto nível, uma incorporação é um conjunto matemático de coordenadas que corresponde ao significado de uma imagem. Isso significa que todas as imagens podem ser pesquisadas com base no que contêm, independentemente do nome do arquivo ou dos metadados. Imagine buscar por "fluxograma" e encontrar todas as imagens relacionadas a fluxogramas, *independentemente dos seus nomes de arquivo.*