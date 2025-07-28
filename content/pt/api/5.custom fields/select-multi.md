---
title: Campo Personalizado de Múltipla Seleção
description: Crie campos de múltipla seleção para permitir que os usuários escolham várias opções de uma lista predefinida
---

Os campos personalizados de múltipla seleção permitem que os usuários escolham várias opções de uma lista predefinida. Eles são ideais para categorias, tags, habilidades, recursos ou qualquer cenário em que múltiplas seleções sejam necessárias a partir de um conjunto controlado de opções.

## Exemplo Básico

Crie um campo simples de múltipla seleção:

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de múltipla seleção e, em seguida, adicione opções separadamente:

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de múltipla seleção |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `SELECT_MULTI` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |
| `projectId` | String! | ✅ Sim | ID do projeto para este campo |

### CreateCustomFieldOptionInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado |
| `title` | String! | ✅ Sim | Texto exibido para a opção |
| `color` | String | Não | Cor para a opção (qualquer string) |
| `position` | Float | Não | Ordem de classificação para a opção |

## Adicionando Opções a Campos Existentes

Adicione novas opções a um campo de múltipla seleção existente:

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## Definindo Valores de Múltipla Seleção

Para definir várias opções selecionadas em um registro:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### Parâmetros SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de múltipla seleção |
| `customFieldOptionIds` | [String!] | ✅ Sim | Array de IDs de opções a selecionar |

## Criando Registros com Valores de Múltipla Seleção

Ao criar um novo registro com valores de múltipla seleção:

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
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
| `selectedOptions` | [CustomFieldOption!] | Array de opções selecionadas |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

### Resposta CustomFieldOption

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para a opção |
| `title` | String! | Texto exibido para a opção |
| `color` | String | Código de cor hexadecimal para representação visual |
| `position` | Float | Ordem de classificação para a opção |
| `customField` | CustomField! | O campo personalizado ao qual esta opção pertence |

### Resposta CustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o campo |
| `name` | String! | Nome exibido do campo de múltipla seleção |
| `type` | CustomFieldType! | Sempre `SELECT_MULTI` |
| `description` | String | Texto de ajuda para o campo |
| `customFieldOptions` | [CustomFieldOption!] | Todas as opções disponíveis |

## Formato de Valor

### Formato de Entrada
- **Parâmetro da API**: Array de IDs de opções (`["option1", "option2", "option3"]`)
- **Formato de String**: IDs de opções separados por vírgula (`"option1,option2,option3"`)

### Formato de Saída
- **Resposta GraphQL**: Array de objetos CustomFieldOption
- **Registro de Atividade**: Títulos de opções separados por vírgula
- **Dados de Automação**: Array de títulos de opções

## Gerenciando Opções

### Atualizar Propriedades da Opção
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
  }) {
    id
    title
    color
  }
}
```

### Excluir Opção
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### Reordenar Opções
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## Regras de Validação

### Validação de Opção
- Todos os IDs de opções fornecidos devem existir
- As opções devem pertencer ao campo personalizado especificado
- Apenas campos SELECT_MULTI podem ter várias opções selecionadas
- Array vazio é válido (sem seleções)

### Validação de Campo
- Deve ter pelo menos uma opção definida para ser utilizável
- Os títulos das opções devem ser exclusivos dentro do campo
- O campo de cor aceita qualquer valor de string (sem validação hexadecimal)

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## Respostas de Erro

### ID de Opção Inválido
```json
{
  "errors": [{
    "message": "Custom field option not found",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### Opção Não Pertence ao Campo
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Múltiplas Opções em Campo Não Múltiplo
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Melhores Práticas

### Design de Opção
- Use títulos de opções descritivos e concisos
- Aplique esquemas de codificação de cores consistentes
- Mantenha listas de opções gerenciáveis (tipicamente de 3 a 20 opções)
- Ordene opções de forma lógica (alfabeticamente, por frequência, etc.)

### Gerenciamento de Dados
- Revise e limpe opções não utilizadas periodicamente
- Use convenções de nomenclatura consistentes entre projetos
- Considere a reutilização de opções ao criar campos
- Planeje atualizações e migrações de opções

### Experiência do Usuário
- Forneça descrições claras dos campos
- Use cores para melhorar a distinção visual
- Agrupe opções relacionadas
- Considere seleções padrão para casos comuns

## Casos de Uso Comuns

1. **Gerenciamento de Projetos**
   - Categorias e tags de tarefas
   - Níveis e tipos de prioridade
   - Atribuições de membros da equipe
   - Indicadores de status

2. **Gerenciamento de Conteúdo**
   - Categorias e tópicos de artigos
   - Tipos e formatos de conteúdo
   - Canais de publicação
   - Fluxos de trabalho de aprovação

3. **Suporte ao Cliente**
   - Categorias e tipos de problemas
   - Produtos ou serviços afetados
   - Métodos de resolução
   - Segmentos de clientes

4. **Desenvolvimento de Produtos**
   - Categorias de recursos
   - Requisitos técnicos
   - Ambientes de teste
   - Canais de lançamento

## Recursos de Integração

### Com Automação
- Acionar ações quando opções específicas são selecionadas
- Roteirizar trabalho com base em categorias selecionadas
- Enviar notificações para seleções de alta prioridade
- Criar tarefas de acompanhamento com base em combinações de opções

### Com Pesquisas
- Filtrar registros por opções selecionadas
- Agregar dados com base em seleções de opções
- Referenciar dados de opções de outros registros
- Criar relatórios com base em combinações de opções

### Com Formulários
- Controles de entrada de múltipla seleção
- Validação e filtragem de opções
- Carregamento dinâmico de opções
- Exibição condicional de campos

## Rastreamento de Atividades

As alterações no campo de múltipla seleção são rastreadas automaticamente:
- Mostra opções adicionadas e removidas
- Exibe títulos de opções no registro de atividade
- Registra horários para todas as alterações de seleção
- Atribuição de usuário para modificações

## Limitações

- O limite prático máximo de opções depende do desempenho da interface do usuário
- Sem estrutura de opção hierárquica ou aninhada
- As opções são compartilhadas entre todos os registros que usam o campo
- Sem análises ou rastreamento de uso de opções incorporados
- O campo de cor aceita qualquer string (sem validação hexadecimal)
- Não é possível definir permissões diferentes por opção
- As opções devem ser criadas separadamente, não inline com a criação do campo
- Sem mutação dedicada para reordenar (use editCustomFieldOption com posição)

## Recursos Relacionados

- [Campos de Seleção Única](/api/custom-fields/select-single) - Para seleções de escolha única
- [Campos de Caixa de Seleção](/api/custom-fields/checkbox) - Para escolhas booleanas simples
- [Campos de Texto](/api/custom-fields/text-single) - Para entrada de texto livre
- [Visão Geral de Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceitos gerais