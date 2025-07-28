---
title: Campo Personalizado de Checkbox
description: Crie campos de checkbox booleanos para dados de sim/não ou verdadeiro/falso
---

Os campos personalizados de checkbox fornecem uma entrada booleana simples (verdadeiro/falso) para tarefas. Eles são perfeitos para escolhas binárias, indicadores de status ou para rastrear se algo foi concluído.

## Exemplo Básico

Crie um campo de checkbox simples:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de checkbox com descrição e validação:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ Sim | Nome exibido do checkbox |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `CHECKBOX` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

## Definindo Valores de Checkbox

Para definir ou atualizar um valor de checkbox em uma tarefa:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

Para desmarcar um checkbox:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Sim | ID da tarefa a ser atualizada |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de checkbox |
| `checked` | Boolean | Não | Verdadeiro para marcar, falso para desmarcar |

## Criando Tarefas com Valores de Checkbox

Ao criar uma nova tarefa com valores de checkbox:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### Valores de String Aceitos

Ao criar tarefas, os valores de checkbox devem ser passados como strings:

| Valor da String | Resultado |
|------------------|----------|
| `"true"` | ✅ Marcado (case-sensitive) |
| `"1"` | ✅ Marcado |
| `"checked"` | ✅ Marcado (case-sensitive) |
| Any other value | ❌ Desmarcado |

**Nota**: Comparações de strings durante a criação de tarefas são sensíveis a maiúsculas e minúsculas. Os valores devem corresponder exatamente a `"true"`, `"1"`, ou `"checked"` para resultar em um estado marcado.

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | Identificador único para o valor do campo |
| `uid` | String! | Identificador único alternativo |
| `customField` | CustomField! | A definição do campo personalizado |
| `checked` | Boolean | O estado do checkbox (verdadeiro/falso/nulo) |
| `todo` | Todo! | A tarefa à qual esse valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

## Integração de Automação

Os campos de checkbox acionam diferentes eventos de automação com base em mudanças de estado:

| Ação | Evento Acionado | Descrição |
|------|----------------|-----------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | Acionado quando o checkbox é marcado |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | Acionado quando o checkbox é desmarcado |

Isso permite que você crie automações que respondem a mudanças de estado do checkbox, como:
- Enviar notificações quando itens são aprovados
- Mover tarefas quando checkboxes de revisão são marcados
- Atualizar campos relacionados com base nos estados dos checkboxes

## Importação/Exportação de Dados

### Importando Valores de Checkbox

Ao importar dados via CSV ou outros formatos:
- `"true"`, `"yes"` → Marcado (case-insensitive)
- Qualquer outro valor (incluindo `"false"`, `"no"`, `"0"`, vazio) → Desmarcado

### Exportando Valores de Checkbox

Ao exportar dados:
- Caixas marcadas exportam como `"X"`
- Caixas desmarcadas exportam como string vazia `""`

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|----------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## Respostas de Erro

### Tipo de Valor Inválido
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
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

### Convenções de Nomenclatura
- Use nomes claros e orientados para a ação: "Aprovado", "Revisado", "Está Completo"
- Evite nomes negativos que confundam os usuários: prefira "Está Ativo" em vez de "Está Inativo"
- Seja específico sobre o que o checkbox representa

### Quando Usar Checkboxes
- **Escolhas binárias**: Sim/Não, Verdadeiro/Falso, Feito/Não Feito
- **Indicadores de status**: Aprovado, Revisado, Publicado
- **Flags de recurso**: Tem Suporte Prioritário, Requer Assinatura
- **Rastreamento simples**: Email Enviado, Fatura Paga, Item Enviado

### Quando NÃO Usar Checkboxes
- Quando você precisa de mais de duas opções (use SELECT_SINGLE em vez disso)
- Para dados numéricos ou de texto (use campos NUMBER ou TEXT)
- Quando você precisa rastrear quem marcou ou quando (use logs de auditoria)

## Casos de Uso Comuns

1. **Fluxos de Trabalho de Aprovação**
   - "Aprovado pelo Gerente"
   - "Assinatura do Cliente"
   - "Revisão Legal Completa"

2. **Gerenciamento de Tarefas**
   - "Está Bloqueado"
   - "Pronto para Revisão"
   - "Alta Prioridade"

3. **Controle de Qualidade**
   - "QA Aprovado"
   - "Documentação Completa"
   - "Testes Escritos"

4. **Flags Administrativas**
   - "Fatura Enviada"
   - "Contrato Assinado"
   - "Acompanhamento Necessário"

## Limitações

- Campos de checkbox podem armazenar apenas valores verdadeiro/falso (sem tri-state ou nulo após a configuração inicial)
- Sem configuração de valor padrão (sempre começa como nulo até ser definido)
- Não é possível armazenar metadados adicionais, como quem marcou ou quando
- Sem visibilidade condicional com base em outros valores de campo

## Recursos Relacionados

- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais de campos personalizados
- [API de Automação](/api/automations) - Crie automações acionadas por mudanças de checkbox