---
title: Campo Personalizado de Número
description: Crie campos numéricos para armazenar valores numéricos com restrições de min/max opcionais e formatação de prefixo
---

Os campos personalizados de número permitem que você armazene valores numéricos para registros. Eles suportam restrições de validação, precisão decimal e podem ser usados para quantidades, pontuações, medições ou qualquer dado numérico que não exija formatação especial.

## Exemplo Básico

Crie um campo numérico simples:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo numérico com restrições e prefixo:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo numérico |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `NUMBER` |
| `projectId` | String! | ✅ Sim | ID do projeto para criar o campo |
| `min` | Float | Não | Restrição de valor mínimo (apenas orientação da UI) |
| `max` | Float | Não | Restrição de valor máximo (apenas orientação da UI) |
| `prefix` | String | Não | Prefixo de exibição (por exemplo, "#", "~", "$") |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

## Definindo Valores Numéricos

Os campos numéricos armazenam valores decimais com validação opcional:

### Valor Numérico Simples

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### Valor Inteiro

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### Parâmetros SetTodoCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de número |
| `number` | Float | Não | Valor numérico a ser armazenado |

## Restrições de Valor

### Restrições de Min/Max (Orientação da UI)

**Importante**: Restrições de min/max são armazenadas, mas NÃO são aplicadas no lado do servidor. Elas servem como orientação da UI para aplicações frontend.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**Validação do Lado do Cliente Necessária**: Aplicações frontend devem implementar lógica de validação para aplicar restrições de min/max.

### Tipos de Valor Suportados

| Tipo | Exemplo | Descrição |
|------|---------|-----------|
| Integer | `42` | Números inteiros |
| Decimal | `42.5` | Números com casas decimais |
| Negative | `-10` | Valores negativos (se não houver restrição de min) |
| Zero | `0` | Valor zero |

**Nota**: Restrições de min/max NÃO são validadas no lado do servidor. Valores fora do intervalo especificado serão aceitos e armazenados.

## Criando Registros com Valores Numéricos

Ao criar um novo registro com valores numéricos:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### Formatos de Entrada Suportados

Ao criar registros, use o parâmetro `number` (não `value`) no array de campos personalizados:

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `number` | Float | O valor numérico |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

### Resposta CustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para a definição do campo |
| `name` | String! | Nome exibido do campo |
| `type` | CustomFieldType! | Sempre `NUMBER` |
| `min` | Float | Valor mínimo permitido |
| `max` | Float | Valor máximo permitido |
| `prefix` | String | Prefixo de exibição |
| `description` | String | Texto de ajuda |

**Nota**: Se o valor numérico não estiver definido, o campo `number` será `null`.

## Filtragem e Consulta

Os campos numéricos suportam filtragem numérica abrangente:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### Operadores Suportados

| Operador | Descrição | Exemplo |
|----------|-----------|---------|
| `EQ` | Igual a | `number = 42` |
| `NE` | Diferente de | `number ≠ 42` |
| `GT` | Maior que | `number > 42` |
| `GTE` | Maior ou igual | `number ≥ 42` |
| `LT` | Menor que | `number < 42` |
| `LTE` | Menor ou igual | `number ≤ 42` |
| `IN` | No array | `number in [1, 2, 3]` |
| `NIN` | Não no array | `number not in [1, 2, 3]` |
| `IS` | É nulo/não nulo | `number is null` |

### Filtragem por Intervalo

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## Formatação de Exibição

### Com Prefixo

Se um prefixo estiver definido, ele será exibido:

| Valor | Prefixo | Exibição |
|-------|---------|----------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### Precisão Decimal

Os números mantêm sua precisão decimal:

| Entrada | Armazenado | Exibido |
|---------|------------|---------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|----------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## Respostas de Erro

### Formato de Número Inválido
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**Nota**: Erros de validação de min/max NÃO ocorrem no lado do servidor. A validação de restrições deve ser implementada em sua aplicação frontend.

### Não é um Número
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Melhores Práticas

### Design de Restrições
- Defina valores de min/max realistas para orientação da UI
- Implemente validação do lado do cliente para aplicar restrições
- Use restrições para fornecer feedback ao usuário em formulários
- Considere se valores negativos são válidos para seu caso de uso

### Precisão de Valor
- Use precisão decimal apropriada para suas necessidades
- Considere arredondar para fins de exibição
- Seja consistente com a precisão entre campos relacionados

### Melhoria de Exibição
- Use prefixos significativos para contexto
- Considere unidades nos nomes dos campos (por exemplo, "Peso (kg)")
- Forneça descrições claras para regras de validação

## Casos de Uso Comuns

1. **Sistemas de Pontuação**
   - Avaliações de desempenho
   - Pontuações de qualidade
   - Níveis de prioridade
   - Avaliações de satisfação do cliente

2. **Medições**
   - Quantidades e valores
   - Dimensões e tamanhos
   - Durações (em formato numérico)
   - Capacidades e limites

3. **Métricas de Negócios**
   - Números de receita
   - Taxas de conversão
   - Alocações orçamentárias
   - Números-alvo

4. **Dados Técnicos**
   - Números de versão
   - Valores de configuração
   - Métricas de desempenho
   - Configurações de limite

## Recursos de Integração

### Com Gráficos e Painéis
- Use campos NUMÉRICOS em cálculos de gráficos
- Crie visualizações numéricas
- Acompanhe tendências ao longo do tempo

### Com Automação
- Acione ações com base em limites numéricos
- Atualize campos relacionados com base em alterações numéricas
- Envie notificações para valores específicos

### Com Pesquisas
- Agregue números de registros relacionados
- Calcule totais e médias
- Encontre valores min/max em relacionamentos

### Com Gráficos
- Crie visualizações numéricas
- Acompanhe tendências ao longo do tempo
- Compare valores entre registros

## Limitações

- **Sem validação no lado do servidor** das restrições de min/max
- **Validação do lado do cliente necessária** para aplicação de restrições
- Sem formatação de moeda embutida (use o tipo MOEDA em vez disso)
- Sem símbolo de porcentagem automático (use o tipo PORCENTAGEM em vez disso)
- Sem capacidades de conversão de unidades
- Precisão decimal limitada pelo tipo Decimal do banco de dados
- Sem avaliação de fórmulas matemáticas no próprio campo

## Recursos Relacionados

- [Visão Geral de Campos Personalizados](/api/custom-fields/1.index) - Conceitos gerais de campos personalizados
- [Campo Personalizado de Moeda](/api/custom-fields/currency) - Para valores monetários
- [Campo Personalizado de Porcentagem](/api/custom-fields/percent) - Para valores percentuais
- [API de Automação](/api/automations/1.index) - Crie automações baseadas em números