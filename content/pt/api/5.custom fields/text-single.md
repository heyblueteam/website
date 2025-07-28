---
title: Campo Personalizado de Texto em Uma Linha
description: Crie campos de texto em uma linha para valores de texto curtos, como nomes, títulos e rótulos
---

Os campos personalizados de texto em uma linha permitem que você armazene valores de texto curtos destinados à entrada em uma linha. Eles são ideais para nomes, títulos, rótulos ou qualquer dado de texto que deve ser exibido em uma única linha.

## Exemplo Básico

Crie um simples campo de texto em uma linha:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de texto em uma linha com descrição:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
  }) {
    id
    name
    type
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Sim | Nome de exibição do campo de texto |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `TEXT_SINGLE` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

**Nota**: O contexto do projeto é determinado automaticamente a partir dos seus cabeçalhos de autenticação. Nenhum parâmetro `projectId` é necessário.

## Definindo Valores de Texto

Para definir ou atualizar um valor de texto em uma linha em um registro:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo de texto personalizado |
| `text` | String | Não | Conteúdo de texto em uma linha a ser armazenado |

## Criando Registros com Valores de Texto

Ao criar um novo registro com valores de texto em uma linha:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado (contém o valor de texto) |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

**Importante**: Os valores de texto são acessados através do campo `customField.value.text`, não diretamente no TodoCustomField.

## Consultando Valores de Texto

Ao consultar registros com campos personalizados de texto, acesse o texto através do caminho `customField.value.text`:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

A resposta incluirá o texto na estrutura aninhada:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## Validação de Texto

### Validação de Formulário
Quando campos de texto em uma linha são usados em formulários:
- Espaços em branco no início e no final são automaticamente removidos
- A validação de obrigatoriedade é aplicada se o campo for marcado como obrigatório
- Nenhuma validação de formato específico é aplicada

### Regras de Validação
- Aceita qualquer conteúdo de string, incluindo quebras de linha (embora não recomendado)
- Sem limites de comprimento de caracteres (até os limites do banco de dados)
- Suporta caracteres Unicode e símbolos especiais
- Quebras de linha são preservadas, mas não são destinadas a este tipo de campo

### Exemplos Típicos de Texto
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## Notas Importantes

### Capacidade de Armazenamento
- Armazenado usando o tipo MySQL `MediumText`
- Suporta até 16MB de conteúdo de texto
- Armazenamento idêntico aos campos de texto em múltiplas linhas
- Codificação UTF-8 para caracteres internacionais

### API Direta vs Formulários
- **Formulários**: Remoção automática de espaços em branco e validação obrigatória
- **API Direta**: O texto é armazenado exatamente como fornecido
- **Recomendação**: Use formulários para entrada de usuário para garantir formatação consistente

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: Entrada de texto em uma linha, ideal para valores curtos
- **TEXT_MULTI**: Entrada de área de texto em múltiplas linhas, ideal para conteúdo mais longo
- **Backend**: Ambos usam armazenamento e validação idênticos
- **Frontend**: Diferentes componentes de UI para entrada de dados
- **Intenção**: TEXT_SINGLE é semanticamente destinado a valores de uma linha

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|----------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

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
      "code": "NOT_FOUND"
    }
  }]
}
```

## Melhores Práticas

### Diretrizes de Conteúdo
- Mantenha o texto conciso e apropriado para uma linha
- Evite quebras de linha para exibição em uma única linha
- Use formatação consistente para tipos de dados semelhantes
- Considere limites de caracteres com base nos requisitos da sua UI

### Entrada de Dados
- Forneça descrições claras dos campos para orientar os usuários
- Use formulários para entrada de usuário para garantir validação
- Valide o formato do conteúdo em sua aplicação, se necessário
- Considere usar dropdowns para valores padronizados

### Considerações de Desempenho
- Campos de texto em uma linha são leves e eficientes
- Considere indexar campos frequentemente pesquisados
- Use larguras de exibição apropriadas em sua UI
- Monitore o comprimento do conteúdo para fins de exibição

## Filtragem e Pesquisa

### Pesquisa de Contém
Campos de texto em uma linha suportam pesquisa de substring:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### Capacidades de Pesquisa
- Correspondência de substring sem diferenciação entre maiúsculas e minúsculas
- Suporta correspondência de palavras parciais
- Correspondência de valores exatos
- Sem pesquisa de texto completo ou classificação

## Casos de Uso Comuns

1. **Identificadores e Códigos**
   - SKUs de produtos
   - Números de pedidos
   - Códigos de referência
   - Números de versão

2. **Nomes e Títulos**
   - Nomes de clientes
   - Títulos de projetos
   - Nomes de produtos
   - Rótulos de categoria

3. **Descrições Curtas**
   - Resumos breves
   - Rótulos de status
   - Indicadores de prioridade
   - Tags de classificação

4. **Referências Externas**
   - Números de bilhetes
   - Referências de faturas
   - IDs de sistemas externos
   - Números de documentos

## Recursos de Integração

### Com Pesquisas
- Referenciar dados de texto de outros registros
- Encontrar registros por conteúdo de texto
- Exibir informações de texto relacionadas
- Agregar valores de texto de múltiplas fontes

### Com Formulários
- Remoção automática de espaços em branco
- Validação de campo obrigatório
- UI de entrada de texto em uma linha
- Exibição de limite de caracteres (se configurado)

### Com Importações/Exportações
- Mapeamento direto de colunas CSV
- Atribuição automática de valores de texto
- Suporte à importação de dados em massa
- Exportar para formatos de planilhas

## Limitações

### Restrições de Automação
- Não disponível diretamente como campos de gatilho de automação
- Não pode ser usado em atualizações de campos de automação
- Pode ser referenciado em condições de automação
- Disponível em modelos de e-mail e webhooks

### Limitações Gerais
- Sem formatação ou estilo de texto embutido
- Sem validação automática além de campos obrigatórios
- Sem aplicação de unicidade embutida
- Sem compressão de conteúdo para texto muito grande
- Sem versionamento ou rastreamento de alterações
- Capacidades de pesquisa limitadas (sem pesquisa de texto completo)

## Recursos Relacionados

- [Campos de Texto em Múltiplas Linhas](/api/custom-fields/text-multi) - Para conteúdo de texto mais longo
- [Campos de E-mail](/api/custom-fields/email) - Para endereços de e-mail
- [Campos de URL](/api/custom-fields/url) - Para endereços de sites
- [Campos de ID Único](/api/custom-fields/unique-id) - Para identificadores gerados automaticamente
- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais