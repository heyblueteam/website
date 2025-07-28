---
title: Campo Personalizado de Percentual
description: Crie campos de percentual para armazenar valores numéricos com manuseio automático do símbolo % e formatação de exibição
---

Os campos personalizados de percentual permitem que você armazene valores percentuais para registros. Eles lidam automaticamente com o símbolo % para entrada e exibição, enquanto armazenam o valor numérico bruto internamente. Perfeito para taxas de conclusão, taxas de sucesso ou quaisquer métricas baseadas em percentual.

## Exemplo Básico

Crie um campo percentual simples:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo percentual com descrição:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome de exibição do campo percentual |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `PERCENT` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

**Nota**: O contexto do projeto é determinado automaticamente a partir dos seus cabeçalhos de autenticação. Nenhum parâmetro `projectId` é necessário.

**Nota**: Campos PERCENT não suportam restrições de min/max ou formatação de prefixo como campos NUMBER.

## Definindo Valores Percentuais

Os campos percentuais armazenam valores numéricos com manuseio automático do símbolo %:

### Com Símbolo de Percentual

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### Valor Numérico Direto

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### Parâmetros SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado percentual |
| `number` | Float | Não | Valor percentual numérico (por exemplo, 75.5 para 75.5%) |

## Armazenamento e Exibição de Valores

### Formato de Armazenamento
- **Armazenamento interno**: Valor numérico bruto (por exemplo, 75.5)
- **Banco de dados**: Armazenado como `Decimal` na coluna `number`
- **GraphQL**: Retornado como tipo `Float`

### Formato de Exibição
- **Interface do usuário**: Aplicações cliente devem adicionar o símbolo % (por exemplo, "75.5%")
- **Gráficos**: Exibe com o símbolo % quando o tipo de saída é PERCENTAGE
- **Respostas da API**: Valor numérico bruto sem o símbolo % (por exemplo, 75.5)

## Criando Registros com Valores Percentuais

Ao criar um novo registro com valores percentuais:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### Formatos de Entrada Suportados

| Formato | Exemplo | Resultado |
|---------|---------|----------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**Nota**: O símbolo % é automaticamente removido da entrada e adicionado novamente durante a exibição.

## Consultando Valores Percentuais

Ao consultar registros com campos personalizados percentuais, acesse o valor através do caminho `customField.value.number`:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

A resposta incluirá a porcentagem como um número bruto:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado (contém o valor percentual) |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

**Importante**: Valores percentuais são acessados através do campo `customField.value.number`. O símbolo % não está incluído nos valores armazenados e deve ser adicionado por aplicações cliente para exibição.

## Filtragem e Consulta

Os campos percentuais suportam a mesma filtragem que os campos NUMBER:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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
| `EQ` | Igual a | `percentage = 75` |
| `NE` | Diferente de | `percentage ≠ 75` |
| `GT` | Maior que | `percentage > 75` |
| `GTE` | Maior ou igual | `percentage ≥ 75` |
| `LT` | Menor que | `percentage < 75` |
| `LTE` | Menor ou igual | `percentage ≤ 75` |
| `IN` | Valor na lista | `percentage in [50, 75, 100]` |
| `NIN` | Valor não na lista | `percentage not in [0, 25]` |
| `IS` | Verificar nulo com `values: null` | `percentage is null` |
| `NOT` | Verificar não nulo com `values: null` | `percentage is not null` |

### Filtragem por Intervalo

Para filtragem por intervalo, use múltiplos operadores:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## Intervalos de Valores Percentuais

### Intervalos Comuns

| Intervalo | Descrição | Caso de Uso |
|-----------|-----------|--------------|
| `0-100` | Percentual padrão | Completion rates, success rates |
| `0-∞` | Percentual ilimitado | Growth rates, performance metrics |
| `-∞-∞` | Qualquer valor | Change rates, variance |

### Valores de Exemplo

| Entrada | Armazenado | Exibição |
|---------|------------|----------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## Agregação de Gráficos

Os campos percentuais suportam agregação em gráficos e relatórios de painel. As funções disponíveis incluem:

- `AVERAGE` - Valor percentual médio
- `COUNT` - Número de registros com valores
- `MIN` - Menor valor percentual
- `MAX` - Maior valor percentual 
- `SUM` - Total de todos os valores percentuais

Essas agregações estão disponíveis ao criar gráficos e painéis, não em consultas diretas GraphQL.

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## Respostas de Erro

### Formato de Percentual Inválido
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

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

### Entrada de Valor
- Permita que os usuários insiram com ou sem o símbolo %
- Valide intervalos razoáveis para seu caso de uso
- Forneça contexto claro sobre o que 100% representa

### Exibição
- Sempre mostre o símbolo % nas interfaces do usuário
- Use precisão decimal apropriada
- Considere a codificação de cores para intervalos (vermelho/amarelo/verde)

### Interpretação de Dados
- Documente o que 100% significa em seu contexto
- Lide com valores acima de 100% de forma apropriada
- Considere se valores negativos são válidos

## Casos de Uso Comuns

1. **Gerenciamento de Projetos**
   - Taxas de conclusão de tarefas
   - Progresso do projeto
   - Utilização de recursos
   - Velocidade do sprint

2. **Rastreamento de Desempenho**
   - Taxas de sucesso
   - Taxas de erro
   - Métricas de eficiência
   - Pontuações de qualidade

3. **Métricas Financeiras**
   - Taxas de crescimento
   - Margens de lucro
   - Valores de desconto
   - Percentuais de mudança

4. **Análise**
   - Taxas de conversão
   - Taxas de cliques
   - Métricas de engajamento
   - Indicadores de desempenho

## Recursos de Integração

### Com Fórmulas
- Referencie campos PERCENT em cálculos
- Formatação automática do símbolo % nas saídas de fórmulas
- Combine com outros campos numéricos

### Com Automação
- Acione ações com base em limites percentuais
- Envie notificações para percentuais de marcos
- Atualize status com base em taxas de conclusão

### Com Consultas
- Agregue percentuais de registros relacionados
- Calcule taxas de sucesso médias
- Encontre itens com melhor/pior desempenho

### Com Gráficos
- Crie visualizações baseadas em percentual
- Acompanhe o progresso ao longo do tempo
- Compare métricas de desempenho

## Diferenças em Relação aos Campos NUMBER

### O que é Diferente
- **Manuseio de entrada**: Remove automaticamente o símbolo %
- **Exibição**: Adiciona automaticamente o símbolo %
- **Restrições**: Sem validação de min/max
- **Formatação**: Sem suporte a prefixo

### O que é Igual
- **Armazenamento**: Mesma coluna e tipo de banco de dados
- **Filtragem**: Mesmos operadores de consulta
- **Agregação**: Mesmas funções de agregação
- **Permissões**: Mesmo modelo de permissão

## Limitações

- Sem restrições de valores min/max
- Sem opções de formatação de prefixo
- Sem validação automática do intervalo de 0-100%
- Sem conversão entre formatos percentuais (por exemplo, 0.75 ↔ 75%)
- Valores acima de 100% são permitidos

## Recursos Relacionados

- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais de campos personalizados
- [Campo Personalizado Numérico](/api/custom-fields/number) - Para valores numéricos brutos
- [API de Automação](/api/automations/index) - Crie automações baseadas em percentual