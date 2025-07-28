---
title: Campo Personalizado de Referência
description: Crie campos de referência que se conectam a registros em outros projetos para relacionamentos entre projetos
---

Os campos personalizados de referência permitem que você crie links entre registros em diferentes projetos, possibilitando relacionamentos entre projetos e compartilhamento de dados. Eles fornecem uma maneira poderosa de conectar trabalhos relacionados na estrutura de projetos da sua organização.

## Exemplo Básico

Crie um campo de referência simples:

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## Exemplo Avançado

Crie um campo de referência com filtragem e seleção múltipla:

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de referência |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `REFERENCE` |
| `referenceProjectId` | String | Não | ID do projeto a ser referenciado |
| `referenceMultiple` | Boolean | Não | Permitir seleção de múltiplos registros (padrão: falso) |
| `referenceFilter` | TodoFilterInput | Não | Critérios de filtragem para registros referenciados |
| `description` | String | Não | Texto de ajuda exibido aos usuários |

**Nota**: Campos personalizados são automaticamente associados ao projeto com base no contexto do projeto atual do usuário.

## Configuração de Referência

### Referências Únicas vs Múltiplas

**Referência Única (padrão):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Os usuários podem selecionar um registro do projeto referenciado
- Retorna um único objeto Todo

**Referências Múltiplas:**
```graphql
{
  referenceMultiple: true
}
```
- Os usuários podem selecionar múltiplos registros do projeto referenciado
- Retorna um array de objetos Todo

### Filtragem de Referência

Use `referenceFilter` para limitar quais registros podem ser selecionados:

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## Definindo Valores de Referência

### Referência Única

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Referências Múltiplas

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### Parâmetros SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de referência |
| `customFieldReferenceTodoIds` | [String!] | ✅ Sim | Array de IDs de registros referenciados |

## Criando Registros com Referências

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
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
| `customField` | CustomField! | A definição do campo de referência |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

**Nota**: Todos os todos referenciados são acessados via `customField.selectedTodos`, não diretamente no TodoCustomField.

### Campos Todo Referenciados

Cada Todo referenciado inclui:

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | Identificador único do registro referenciado |
| `title` | String! | Título do registro referenciado |
| `status` | TodoStatus! | Status atual (ATIVO, CONCLUÍDO, etc.) |
| `description` | String | Descrição do registro referenciado |
| `dueDate` | DateTime | Data de vencimento, se definida |
| `assignees` | [User!] | Usuários atribuídos |
| `tags` | [Tag!] | Tags associadas |
| `project` | Project! | Projeto contendo o registro referenciado |

## Consultando Dados de Referência

### Consulta Básica

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### Consulta Avançada com Dados Aninhados

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Importante**: Os usuários devem ter permissões de visualização no projeto referenciado para ver os registros vinculados.

## Acesso entre Projetos

### Visibilidade do Projeto

- Os usuários só podem referenciar registros de projetos aos quais têm acesso
- Os registros referenciados respeitam as permissões do projeto original
- Alterações nos registros referenciados aparecem em tempo real
- Excluir registros referenciados os remove dos campos de referência

### Herança de Permissões

- Os campos de referência herdam permissões de ambos os projetos
- Os usuários precisam de acesso de visualização ao projeto referenciado
- As permissões de edição são baseadas nas regras do projeto atual
- Os dados referenciados são somente leitura no contexto do campo de referência

## Respostas de Erro

### Projeto de Referência Inválido

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### Registro Referenciado Não Encontrado

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

### Permissão Negada

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## Melhores Práticas

### Design de Campo

1. **Nomenclatura clara** - Use nomes descritivos que indiquem o relacionamento
2. **Filtragem apropriada** - Defina filtros para mostrar apenas registros relevantes
3. **Considere as permissões** - Garanta que os usuários tenham acesso aos projetos referenciados
4. **Documente relacionamentos** - Forneça descrições claras da conexão

### Considerações de Desempenho

1. **Limite o escopo de referência** - Use filtros para reduzir o número de registros selecionáveis
2. **Evite aninhamento profundo** - Não crie cadeias complexas de referências
3. **Considere o cache** - Os dados referenciados são armazenados em cache para desempenho
4. **Monitore o uso** - Acompanhe como as referências estão sendo usadas entre projetos

### Integridade dos Dados

1. **Lide com exclusões** - Planeje quando registros referenciados forem excluídos
2. **Valide permissões** - Garanta que os usuários possam acessar projetos referenciados
3. **Atualize dependências** - Considere o impacto ao alterar registros referenciados
4. **Auditorias** - Acompanhe relacionamentos de referência para conformidade

## Casos de Uso Comuns

### Dependências de Projetos

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### Requisitos do Cliente

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### Alocação de Recursos

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### Garantia de Qualidade

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## Integração com Pesquisas

Os campos de referência funcionam com [Campos de Pesquisa](/api/custom-fields/lookup) para extrair dados de registros referenciados. Os campos de pesquisa podem extrair valores de registros selecionados em campos de referência, mas são apenas extratores de dados (nenhuma função de agregação como SOMA é suportada).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## Limitações

- Projetos referenciados devem ser acessíveis ao usuário
- Alterações nas permissões do projeto referenciado afetam o acesso ao campo de referência
- Aninhamento profundo de referências pode impactar o desempenho
- Sem validação embutida para referências circulares
- Sem restrição automática impedindo referências de mesmo projeto
- A validação de filtro não é aplicada ao definir valores de referência

## Recursos Relacionados

- [Campos de Pesquisa](/api/custom-fields/lookup) - Extraia dados de registros referenciados
- [API de Projetos](/api/projects) - Gerenciando projetos que contêm referências
- [API de Registros](/api/records) - Trabalhando com registros que têm referências
- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais