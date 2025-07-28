---
title: Campo Personalizado de Botão
description: Crie campos de botão interativos que acionam automações quando clicados
---

Os campos personalizados de botão fornecem elementos de interface do usuário interativos que acionam automações quando clicados. Ao contrário de outros tipos de campos personalizados que armazenam dados, os campos de botão servem como gatilhos de ação para executar fluxos de trabalho configurados.

## Exemplo Básico

Crie um campo de botão simples que aciona uma automação:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um botão com requisitos de confirmação:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Requerido | Descrição |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sim | Nome exibido do botão |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `BUTTON` |
| `projectId` | String! | ✅ Sim | ID do projeto onde o campo será criado |
| `buttonType` | String | Não | Comportamento de confirmação (veja Tipos de Botão abaixo) |
| `buttonConfirmText` | String | Não | Texto que os usuários devem digitar para confirmação rígida |
| `description` | String | Não | Texto de ajuda exibido para os usuários |
| `required` | Boolean | Não | Se o campo é obrigatório (padrão é falso) |
| `isActive` | Boolean | Não | Se o campo está ativo (padrão é verdadeiro) |

### Campo de Tipo de Botão

O `buttonType` é uma string de formato livre que pode ser usada pelos clientes da interface do usuário para determinar o comportamento de confirmação. Valores comuns incluem:

- `""` (vazio) - Sem confirmação
- `"soft"` - Diálogo de confirmação simples
- `"hard"` - Exigir digitação do texto de confirmação

**Nota**: Estes são apenas dicas de interface do usuário. A API não valida ou impõe valores específicos.

## Acionando Cliques no Botão

Para acionar um clique no botão e executar automações associadas:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### Parâmetros de Entrada do Clique

| Parâmetro | Tipo | Requerido | Descrição |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sim | ID da tarefa contendo o botão |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado do botão |

### Importante: Comportamento da API

**Todos os cliques de botão através da API são executados imediatamente** independentemente de quaisquer configurações de `buttonType` ou `buttonConfirmText`. Esses campos são armazenados para que os clientes da interface do usuário implementem diálogos de confirmação, mas a API em si:

- Não valida o texto de confirmação
- Não impõe quaisquer requisitos de confirmação
- Executa a ação do botão imediatamente quando chamada

A confirmação é puramente um recurso de segurança da interface do usuário do lado do cliente.

### Exemplo: Clicando em Diferentes Tipos de Botão

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

Todas as três mutações acima executarão a ação do botão imediatamente quando chamadas através da API, ignorando quaisquer requisitos de confirmação.

## Campos de Resposta

### Resposta do Campo Personalizado

| Campo | Tipo | Descrição |
|-------|------|-------------|
| `id` | String! | Identificador único para o campo personalizado |
| `name` | String! | Nome exibido do botão |
| `type` | CustomFieldType! | Sempre `BUTTON` para campos de botão |
| `buttonType` | String | Configuração do comportamento de confirmação |
| `buttonConfirmText` | String | Texto de confirmação obrigatório (se usando confirmação rígida) |
| `description` | String | Texto de ajuda para os usuários |
| `required` | Boolean! | Se o campo é obrigatório |
| `isActive` | Boolean! | Se o campo está atualmente ativo |
| `projectId` | String! | ID do projeto ao qual este campo pertence |
| `createdAt` | DateTime! | Quando o campo foi criado |
| `updatedAt` | DateTime! | Quando o campo foi modificado pela última vez |

## Como os Campos de Botão Funcionam

### Integração de Automação

Os campos de botão são projetados para funcionar com o sistema de automação do Blue:

1. **Crie o campo de botão** usando a mutação acima
2. **Configure automações** que escutam eventos `CUSTOM_FIELD_BUTTON_CLICKED`
3. **Os usuários clicam no botão** na interface do usuário
4. **As automações executam** as ações configuradas

### Fluxo de Eventos

Quando um botão é clicado:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### Sem Armazenamento de Dados

Importante: Campos de botão não armazenam nenhum valor de dados. Eles servem puramente como gatilhos de ação. Cada clique:
- Gera um evento
- Aciona automações associadas
- Registra uma ação no histórico da tarefa
- Não modifica nenhum valor de campo

## Permissões Necessárias

Os usuários precisam de funções de projeto apropriadas para criar e usar campos de botão:

| Ação | Função Necessária |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## Respostas de Erro

### Permissão Negada
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Campo Personalizado Não Encontrado
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

**Nota**: A API não retorna erros específicos para automações ausentes ou discrepâncias de confirmação.

## Melhores Práticas

### Convenções de Nomenclatura
- Use nomes orientados à ação: "Enviar Fatura", "Criar Relatório", "Notificar Equipe"
- Seja específico sobre o que o botão faz
- Evite nomes genéricos como "Botão 1" ou "Clique Aqui"

### Configurações de Confirmação
- Deixe `buttonType` vazio para ações seguras e reversíveis
- Defina `buttonType` para sugerir comportamento de confirmação aos clientes da interface do usuário
- Use `buttonConfirmText` para especificar o que os usuários devem digitar nas confirmações da interface do usuário
- Lembre-se: Estas são apenas dicas de interface do usuário - chamadas de API sempre executam imediatamente

### Design de Automação
- Mantenha as ações do botão focadas em um único fluxo de trabalho
- Forneça feedback claro sobre o que aconteceu após o clique
- Considere adicionar texto descritivo para explicar o propósito do botão

## Casos de Uso Comuns

1. **Transições de Fluxo de Trabalho**
   - "Marcar como Completo"
   - "Enviar para Aprovação"
   - "Arquivar Tarefa"

2. **Integrações Externas**
   - "Sincronizar com CRM"
   - "Gerar Fatura"
   - "Enviar Atualização por Email"

3. **Operações em Lote**
   - "Atualizar Todas as Subtarefas"
   - "Copiar para Projetos"
   - "Aplicar Modelo"

4. **Ações de Relatório**
   - "Gerar Relatório"
   - "Exportar Dados"
   - "Criar Resumo"

## Limitações

- Botões não podem armazenar ou exibir valores de dados
- Cada botão só pode acionar automações, não chamadas diretas da API (no entanto, automações podem incluir ações de solicitação HTTP para chamar APIs externas ou as próprias APIs do Blue)
- A visibilidade do botão não pode ser controlada condicionalmente
- Máximo de uma execução de automação por clique (embora essa automação possa acionar várias ações)

## Recursos Relacionados

- [API de Automação](/api/automations/index) - Configure ações acionadas por botões
- [Visão Geral dos Campos Personalizados](/custom-fields/list-custom-fields) - Conceitos gerais de campos personalizados