---
title: Campo Personalizado de Seleção Única
description: Crie campos de seleção única para permitir que os usuários escolham uma opção de uma lista predefinida
---

Campos personalizados de seleção única permitem que os usuários escolham exatamente uma opção de uma lista predefinida. Eles são ideais para campos de status, categorias, prioridades ou qualquer cenário em que apenas uma escolha deva ser feita a partir de um conjunto controlado de opções.

## Exemplo Básico

Crie um campo simples de seleção única:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de seleção única com opções predefinidas:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de seleção única |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `SELECT_SINGLE` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | Não | Opções iniciais para o campo |

### CreateCustomFieldOptionInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `title` | String! | ✅ Sim | Texto exibido para a opção |
| `color` | String | Não | Código de cor hexadecimal para a opção |

## Adicionando Opções a Campos Existentes

Adicione novas opções a um campo de seleção única existente:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## Definindo Valores de Seleção Única

Para definir a opção selecionada em um registro:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de seleção única |
| `customFieldOptionId` | String | Não | ID da opção a ser selecionada (preferido para seleção única) |
| `customFieldOptionIds` | [String!] | Não | Array de IDs de opções (usa o primeiro elemento para seleção única) |

## Consultando Valores de Seleção Única

Consulte o valor de seleção única de um registro:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

O campo `value` retorna um objeto JSON com os detalhes da opção selecionada.

## Criando Registros com Valores de Seleção Única

Ao criar um novo registro com valores de seleção única:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
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
| `value` | JSON | Contém o objeto da opção selecionada com id, título, cor, posição |
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
| `name` | String! | Nome exibido do campo de seleção única |
| `type` | CustomFieldType! | Sempre `SELECT_SINGLE` |
| `description` | String | Texto de ajuda para o campo |
| `customFieldOptions` | [CustomFieldOption!] | Todas as opções disponíveis |

## Formato do Valor

### Formato de Entrada
- **Parâmetro da API**: Use `customFieldOptionId` para ID de opção única
- **Alternativa**: Use `customFieldOptionIds` array (usa o primeiro elemento)
- **Limpar Seleção**: Omitir ambos os campos ou passar valores vazios

### Formato de Saída
- **Resposta GraphQL**: Objeto JSON no campo `value` contendo {id, título, cor, posição}
- **Registro de Atividade**: Título da opção como string
- **Dados de Automação**: Título da opção como string

## Comportamento de Seleção

### Seleção Exclusiva
- Definir uma nova opção remove automaticamente a seleção anterior
- Apenas uma opção pode ser selecionada por vez
- Definir `null` ou valor vazio limpa a seleção

### Lógica de Retorno
- Se o array `customFieldOptionIds` for fornecido, apenas a primeira opção é usada
- Isso garante compatibilidade com formatos de entrada de seleção múltipla
- Arrays vazios ou valores nulos limpam a seleção

## Gerenciando Opções

### Atualizar Propriedades da Opção
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### Deletar Opção
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**Nota**: Deletar uma opção a removerá de todos os registros onde foi selecionada.

### Reordenar Opções
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## Regras de Validação

### Validação de Opção
- O ID da opção fornecido deve existir
- A opção deve pertencer ao campo personalizado especificado
- Apenas uma opção pode ser selecionada (aplicado automaticamente)
- Valores nulos/vazios são válidos (sem seleção)

### Validação de Campo
- Deve ter pelo menos uma opção definida para ser utilizável
- Títulos de opções devem ser únicos dentro do campo
- Códigos de cor devem estar em formato hexadecimal válido (se fornecidos)

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|----------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## Respostas de Erro

### ID de Opção Inválido
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Não Foi Possível Analisar o Valor
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Melhores Práticas

### Design de Opção
- Use títulos de opção claros e descritivos
- Aplique codificação de cores significativa
- Mantenha listas de opções focadas e relevantes
- Ordene opções logicamente (por prioridade, frequência, etc.)

### Padrões de Campo de Status
- Use fluxos de trabalho de status consistentes em projetos
- Considere a progressão natural das opções
- Inclua estados finais claros (Concluído, Cancelado, etc.)
- Use cores que reflitam o significado da opção

### Gestão de Dados
- Revise e limpe opções não utilizadas periodicamente
- Use convenções de nomenclatura consistentes
- Considere o impacto da exclusão de opções em registros existentes
- Planeje atualizações e migrações de opções

## Casos de Uso Comuns

1. **Status e Fluxo de Trabalho**
   - Status da tarefa (A Fazer, Em Progresso, Concluído)
   - Status de aprovação (Pendente, Aprovado, Rejeitado)
   - Fase do projeto (Planejamento, Desenvolvimento, Testes, Lançado)
   - Status de resolução de problemas

2. **Classificação e Categorização**
   - Níveis de prioridade (Baixa, Média, Alta, Crítica)
   - Tipos de tarefa (Erro, Recurso, Melhoria, Documentação)
   - Categorias de projeto (Interno, Cliente, Pesquisa)
   - Atribuições de departamento

3. **Qualidade e Avaliação**
   - Status de revisão (Não Iniciado, Em Revisão, Aprovado)
   - Avaliações de qualidade (Ruim, Regular, Bom, Excelente)
   - Níveis de risco (Baixo, Médio, Alto)
   - Níveis de confiança

4. **Atribuição e Propriedade**
   - Atribuições de equipe
   - Propriedade de departamento
   - Atribuições baseadas em função
   - Atribuições regionais

## Recursos de Integração

### Com Automação
- Acione ações quando opções específicas forem selecionadas
- Roteie trabalho com base nas categorias selecionadas
- Envie notificações para mudanças de status
- Crie fluxos de trabalho condicionais com base nas seleções

### Com Pesquisas
- Filtre registros por opções selecionadas
- Referencie dados de opções de outros registros
- Crie relatórios com base nas seleções de opções
- Agrupe registros por valores selecionados

### Com Formulários
- Controles de entrada em dropdown
- Interfaces de botão de rádio
- Validação e filtragem de opções
- Exibição de campo condicional com base nas seleções

## Rastreamento de Atividade

Mudanças em campos de seleção única são rastreadas automaticamente:
- Mostra seleções de opções antigas e novas
- Exibe títulos de opções no registro de atividade
- Marcas de tempo para todas as mudanças de seleção
- Atribuição de usuário para modificações

## Diferenças em Relação à Seleção Múltipla

| Recurso | Seleção Única | Seleção Múltipla |
|---------|---------------|------------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## Limitações

- Apenas uma opção pode ser selecionada por vez
- Sem estrutura de opção hierárquica ou aninhada
- Opções são compartilhadas entre todos os registros que usam o campo
- Sem análises ou rastreamento de uso de opções embutidos
- Códigos de cor são apenas para exibição, sem impacto funcional
- Não é possível definir permissões diferentes por opção

## Recursos Relacionados

- [Campos de Seleção Múltipla](/api/custom-fields/select-multi) - Para seleções de múltiplas escolhas
- [Campos de Caixa de Seleção](/api/custom-fields/checkbox) - Para escolhas booleanas simples
- [Campos de Texto](/api/custom-fields/text-single) - Para entrada de texto livre
- [Visão Geral de Campos Personalizados](/api/custom-fields/1.index) - Conceitos gerais