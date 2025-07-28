---
title: Campo Personalizado de Data
description: Crie campos de data para rastrear datas únicas ou intervalos de datas com suporte a fuso horário
---

Os campos personalizados de data permitem que você armazene datas únicas ou intervalos de datas para registros. Eles suportam o tratamento de fusos horários, formatação inteligente e podem ser usados para rastrear prazos, datas de eventos ou qualquer informação baseada em tempo.

## Exemplo Básico

Crie um campo de data simples:

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de data de vencimento com descrição:

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de data |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `DATE` |
| `isDueDate` | Boolean | Não | Se este campo representa uma data de vencimento |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

**Nota**: Campos personalizados são automaticamente associados ao projeto com base no contexto do projeto atual do usuário. Nenhum `projectId` é necessário.

## Definindo Valores de Data

Os campos de data podem armazenar uma única data ou um intervalo de datas:

### Data Única

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Intervalo de Datas

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Evento de Dia Todo

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### Parâmetros SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de data |
| `startDate` | DateTime | Não | Data/hora de início no formato ISO 8601 |
| `endDate` | DateTime | Não | Data/hora de término no formato ISO 8601 |
| `timezone` | String | Não | Identificador de fuso horário (por exemplo, "America/New_York") |

**Nota**: Se apenas `startDate` for fornecido, `endDate` automaticamente será definido para o mesmo valor.

## Formatos de Data

### Formato ISO 8601
Todas as datas devem ser fornecidas no formato ISO 8601:
- `2025-01-15T14:30:00Z` - hora UTC
- `2025-01-15T14:30:00+05:00` - Com deslocamento de fuso horário
- `2025-01-15T14:30:00.123Z` - Com milissegundos

### Identificadores de Fuso Horário
Use identificadores de fuso horário padrão:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Se nenhum fuso horário for fornecido, o sistema usará o fuso horário detectado do usuário.

## Criando Registros com Valores de Data

Ao criar um novo registro com valores de data:

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### Formatos de Entrada Suportados

Ao criar registros, as datas podem ser fornecidas em vários formatos:

| Formato | Exemplo | Resultado |
|---------|---------|----------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | Identificador único para o valor do campo |
| `uid` | String! | String de identificador único |
| `customField` | CustomField! | A definição do campo personalizado (contém os valores de data) |
| `todo` | Todo! | O registro ao qual esse valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

**Importante**: Valores de data (`startDate`, `endDate`, `timezone`) são acessados através do campo `customField.value`, não diretamente no TodoCustomField.

### Estrutura do Objeto de Valor

Valores de data são retornados através do campo `customField.value` como um objeto JSON:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**Nota**: O campo `value` é do tipo `CustomField`, não do `TodoCustomField`.

## Consultando Valores de Data

Ao consultar registros com campos personalizados de data, acesse os valores de data através do campo `customField.value`:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

A resposta incluirá os valores de data no campo `value`:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## Inteligência de Exibição de Data

O sistema formata automaticamente as datas com base no intervalo:

| Cenário | Formato de Exibição |
|---------|---------------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (sem hora exibida) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Detecção de dia todo**: Eventos de 00:00 a 23:59 são automaticamente detectados como eventos de dia todo.

## Tratamento de Fuso Horário

### Armazenamento
- Todas as datas são armazenadas em UTC no banco de dados
- As informações de fuso horário são preservadas separadamente
- A conversão ocorre na exibição

### Melhores Práticas
- Sempre forneça o fuso horário para precisão
- Use fusos horários consistentes dentro de um projeto
- Considere as localizações dos usuários para equipes globais

### Fusos Horários Comuns

| Região | ID do Fuso Horário | Deslocamento UTC |
|--------|--------------------|------------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Filtragem e Consulta

Os campos de data suportam filtragem complexa:

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### Verificando Campos de Data Vazios

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### Operadores Suportados

| Operador | Uso | Descrição |
|----------|-----|-----------|
| `EQ` | Com dateRange | A data se sobrepõe ao intervalo especificado (qualquer interseção) |
| `NE` | Com dateRange | A data não se sobrepõe ao intervalo |
| `IS` | Com `values: null` | O campo de data está vazio (startDate ou endDate é nulo) |
| `NOT` | Com `values: null` | O campo de data tem um valor (ambas as datas não são nulas) |

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Respostas de Erro

### Formato de Data Inválido
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
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
      "code": "NOT_FOUND"
    }
  }]
}
```


## Limitações

- Sem suporte a datas recorrentes (use automações para eventos recorrentes)
- Não é possível definir hora sem data
- Sem cálculo embutido de dias úteis
- Intervalos de datas não validam automaticamente fim > início
- A precisão máxima é até o segundo (sem armazenamento de milissegundos)

## Recursos Relacionados

- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais de campos personalizados
- [API de Automações](/api/automations/index) - Crie automações baseadas em datas