---
title: Campo Personalizado de Pesquisa
description: Crie campos de pesquisa que puxam automaticamente dados de registros referenciados
---

Os campos personalizados de pesquisa puxam automaticamente dados de registros referenciados por [Campos de Referência](/api/custom-fields/reference), exibindo informações de registros vinculados sem cópias manuais. Eles se atualizam automaticamente quando os dados referenciados mudam.

## Exemplo Básico

Crie um campo de pesquisa para exibir tags de registros referenciados:

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Exemplo Avançado

Crie um campo de pesquisa para extrair valores de campos personalizados de registros referenciados:

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de pesquisa |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Sim | Configuração de pesquisa |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

## Configuração de Pesquisa

### CustomFieldLookupOptionInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `referenceId` | String! | ✅ Sim | ID do campo de referência para puxar dados |
| `lookupId` | String | Não | ID do campo personalizado específico a ser pesquisado (obrigatório para o tipo TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Sim | Tipo de dado a ser extraído de registros referenciados |

## Tipos de Pesquisa

### Valores de CustomFieldLookupType

| Tipo | Descrição | Retorna |
|------|-----------|---------|
| `TODO_DUE_DATE` | Datas de vencimento de tarefas referenciadas | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Datas de criação de tarefas referenciadas | Array of creation timestamps |
| `TODO_UPDATED_AT` | Datas da última atualização de tarefas referenciadas | Array of update timestamps |
| `TODO_TAG` | Tags de tarefas referenciadas | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Atribuições de tarefas referenciadas | Array of user objects |
| `TODO_DESCRIPTION` | Descrições de tarefas referenciadas | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Nomes de listas de tarefas referenciadas | Array of list titles |
| `TODO_CUSTOM_FIELD` | Valores de campos personalizados de tarefas referenciadas | Array of values based on the field type |

## Campos de Resposta

### Resposta de CustomField (para campos de pesquisa)

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o campo |
| `name` | String! | Nome exibido do campo de pesquisa |
| `type` | CustomFieldType! | Será `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Configuração e resultados da pesquisa |
| `createdAt` | DateTime! | Quando o campo foi criado |
| `updatedAt` | DateTime! | Quando o campo foi atualizado pela última vez |

### Estrutura de CustomFieldLookupOption

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `lookupType` | CustomFieldLookupType! | Tipo de pesquisa que está sendo realizada |
| `lookupResult` | JSON | Os dados extraídos de registros referenciados |
| `reference` | CustomField | O campo de referência que está sendo usado como fonte |
| `lookup` | CustomField | O campo específico que está sendo pesquisado (para TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | O campo de pesquisa pai |
| `parentLookup` | CustomField | Pesquisa pai na cadeia (para pesquisas aninhadas) |

## Como Funcionam as Pesquisas

1. **Extração de Dados**: As pesquisas extraem dados específicos de todos os registros vinculados através de um campo de referência
2. **Atualizações Automáticas**: Quando os registros referenciados mudam, os valores de pesquisa se atualizam automaticamente
3. **Somente Leitura**: Campos de pesquisa não podem ser editados diretamente - eles sempre refletem os dados referenciados atuais
4. **Sem Cálculos**: As pesquisas extraem e exibem dados como estão, sem agregações ou cálculos

## Pesquisas TODO_CUSTOM_FIELD

Ao usar o tipo `TODO_CUSTOM_FIELD`, você deve especificar qual campo personalizado extrair usando o parâmetro `lookupId`:

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

Isso extrai os valores do campo personalizado especificado de todos os registros referenciados.

## Consultando Dados de Pesquisa

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## Resultados de Pesquisa de Exemplo

### Resultado da Pesquisa de Tag
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### Resultado da Pesquisa de Atribuição
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### Resultado da Pesquisa de Campo Personalizado
Os resultados variam com base no tipo de campo personalizado que está sendo pesquisado. Por exemplo, uma pesquisa de campo de moeda pode retornar:
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Importante**: Os usuários devem ter permissões de visualização tanto no projeto atual quanto no projeto referenciado para ver os resultados da pesquisa.

## Respostas de Erro

### Campo de Referência Inválido
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

### Pesquisa Circular Detectada
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### ID de Pesquisa Ausente para TODO_CUSTOM_FIELD
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## Melhores Práticas

1. **Nomenclatura Clara**: Use nomes descritivos que indiquem quais dados estão sendo pesquisados
2. **Tipos Apropriados**: Escolha o tipo de pesquisa que corresponda às suas necessidades de dados
3. **Desempenho**: As pesquisas processam todos os registros referenciados, então tenha cuidado com campos de referência com muitos links
4. **Permissões**: Garanta que os usuários tenham acesso aos projetos referenciados para que as pesquisas funcionem

## Casos de Uso Comuns

### Visibilidade entre Projetos
Exiba tags, atribuições ou status de projetos relacionados sem sincronização manual.

### Rastreamento de Dependências
Mostre datas de vencimento ou status de conclusão de tarefas das quais o trabalho atual depende.

### Visão Geral de Recursos
Exiba todos os membros da equipe atribuídos a tarefas referenciadas para planejamento de recursos.

### Agregação de Status
Colete todos os status únicos de tarefas relacionadas para ver a saúde do projeto de relance.

## Limitações

- Campos de pesquisa são somente leitura e não podem ser editados diretamente
- Sem funções de agregação (SOMA, CONTAR, MÉDIA) - as pesquisas apenas extraem dados
- Sem opções de filtragem - todos os registros referenciados são incluídos
- Cadeias de pesquisa circulares são prevenidas para evitar loops infinitos
- Os resultados refletem os dados atuais e se atualizam automaticamente

## Recursos Relacionados

- [Campos de Referência](/api/custom-fields/reference) - Crie links para registros como fontes de pesquisa
- [Valores de Campos Personalizados](/api/custom-fields/custom-field-values) - Defina valores em campos personalizados editáveis
- [Listar Campos Personalizados](/api/custom-fields/list-custom-fields) - Consultar todos os campos personalizados em um projeto