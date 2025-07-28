---
title: Campo Personalizado de Texto em Várias Linhas
description: Crie campos de texto em várias linhas para conteúdos mais longos, como descrições, notas e comentários
---

Os campos personalizados de texto em várias linhas permitem armazenar conteúdos de texto mais longos com quebras de linha e formatação. Eles são ideais para descrições, notas, comentários ou qualquer dado textual que precise de várias linhas.

## Exemplo Básico

Crie um campo de texto simples em várias linhas:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de texto em várias linhas com descrição:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
    id
    name
    type
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de texto |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `TEXT_MULTI` |
| `description` | String | Não | Texto de ajuda exibido aos usuários |

**Nota:** O `projectId` é passado como um argumento separado para a mutação, não como parte do objeto de entrada. Alternativamente, o contexto do projeto pode ser determinado pelo cabeçalho `X-Bloo-Project-ID` na sua solicitação GraphQL.

## Definindo Valores de Texto

Para definir ou atualizar um valor de texto em várias linhas em um registro:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo de texto personalizado |
| `text` | String | Não | Conteúdo de texto em várias linhas a ser armazenado |

## Criando Registros com Valores de Texto

Ao criar um novo registro com valores de texto em várias linhas:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      text
    }
  }
}
```

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `text` | String | O conteúdo de texto em várias linhas armazenado |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

## Validação de Texto

### Validação de Formulário
Quando campos de texto em várias linhas são usados em formulários:
- Espaços em branco iniciais e finais são automaticamente removidos
- Validação obrigatória é aplicada se o campo for marcado como obrigatório
- Nenhuma validação de formato específica é aplicada

### Regras de Validação
- Aceita qualquer conteúdo de string, incluindo quebras de linha
- Sem limites de comprimento de caracteres (até os limites do banco de dados)
- Suporta caracteres Unicode e símbolos especiais
- Quebras de linha são preservadas no armazenamento

### Exemplos de Texto Válido
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## Notas Importantes

### Capacidade de Armazenamento
- Armazenado usando o tipo MySQL `MediumText`
- Suporta até 16MB de conteúdo de texto
- Quebras de linha e formatação são preservadas
- Codificação UTF-8 para caracteres internacionais

### API Direta vs Formulários
- **Formulários**: Remoção automática de espaços em branco e validação obrigatória
- **API Direta**: O texto é armazenado exatamente como fornecido
- **Recomendação**: Use formulários para entrada do usuário para garantir formatação consistente

### TEXT_MULTI vs TEXT_SINGLE
- **TEXT_MULTI**: Entrada de área de texto em várias linhas, ideal para conteúdos mais longos
- **TEXT_SINGLE**: Entrada de texto em uma linha, ideal para valores curtos
- **Backend**: Ambos os tipos são idênticos - mesmo campo de armazenamento, validação e processamento
- **Frontend**: Diferentes componentes de UI para entrada de dados (área de texto vs campo de entrada)
- **Importante**: A distinção entre TEXT_MULTI e TEXT_SINGLE existe puramente para fins de UI

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

## Respostas de Erro

### Validação de Campo Obrigatório (Somente Formulários)
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Campo Não Encontrado
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Melhores Práticas

### Organização do Conteúdo
- Use formatação consistente para conteúdo estruturado
- Considere usar uma sintaxe semelhante ao markdown para legibilidade
- Divida conteúdos longos em seções lógicas
- Use quebras de linha para melhorar a legibilidade

### Entrada de Dados
- Forneça descrições claras dos campos para orientar os usuários
- Use formulários para entrada do usuário para garantir validação
- Considere limites de caracteres com base no seu caso de uso
- Valide o formato do conteúdo em sua aplicação, se necessário

### Considerações de Desempenho
- Conteúdos de texto muito longos podem afetar o desempenho da consulta
- Considere paginação para exibir grandes campos de texto
- Considerações de índice para funcionalidade de busca
- Monitore o uso de armazenamento para campos com conteúdo grande

## Filtragem e Busca

### Busca por Contém
Campos de texto em várias linhas suportam busca por substring através de filtros de campo personalizado:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### Capacidades de Busca
- Correspondência de substring dentro de campos de texto usando o operador `CONTAINS`
- Busca sem diferenciação entre maiúsculas e minúsculas usando o operador `NCONTAINS`
- Correspondência exata usando o operador `IS`
- Correspondência negativa usando o operador `NOT`
- Busca em todas as linhas de texto
- Suporta correspondência de palavras parciais

## Casos de Uso Comuns

1. **Gerenciamento de Projetos**
   - Descrições de tarefas
   - Requisitos do projeto
   - Notas de reuniões
   - Atualizações de status

2. **Suporte ao Cliente**
   - Descrições de problemas
   - Notas de resolução
   - Feedback de clientes
   - Registros de comunicação

3. **Gerenciamento de Conteúdo**
   - Conteúdo de artigos
   - Descrições de produtos
   - Comentários de usuários
   - Detalhes de avaliações

4. **Documentação**
   - Descrições de processos
   - Instruções
   - Diretrizes
   - Materiais de referência

## Recursos de Integração

### Com Automação
- Acionar ações quando o conteúdo de texto mudar
- Extrair palavras-chave do conteúdo de texto
- Criar resumos ou notificações
- Processar conteúdo de texto com serviços externos

### Com Consultas
- Referenciar dados de texto de outros registros
- Agregar conteúdo de texto de várias fontes
- Encontrar registros pelo conteúdo de texto
- Exibir informações de texto relacionadas

### Com Formulários
- Remoção automática de espaços em branco
- Validação de campo obrigatório
- UI de área de texto em várias linhas
- Exibição de contagem de caracteres (se configurado)

## Limitações

- Sem formatação de texto embutida ou edição de texto rico
- Sem detecção ou conversão automática de links
- Sem verificação ortográfica ou validação gramatical
- Sem análise ou processamento de texto embutido
- Sem versionamento ou rastreamento de mudanças
- Capacidades de busca limitadas (sem busca de texto completo)
- Sem compressão de conteúdo para textos muito grandes

## Recursos Relacionados

- [Campos de Texto em Uma Linha](/api/custom-fields/text-single) - Para valores de texto curtos
- [Campos de Email](/api/custom-fields/email) - Para endereços de email
- [Campos de URL](/api/custom-fields/url) - Para endereços de sites
- [Visão Geral dos Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceitos gerais