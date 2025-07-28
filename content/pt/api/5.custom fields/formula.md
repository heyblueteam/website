---
title: Campo Personalizado de Fórmula
description: Crie campos calculados que computam automaticamente valores com base em outros dados
---

Os campos personalizados de fórmula são usados para cálculos de gráficos e painéis dentro do Blue. Eles definem funções de agregação (SOMA, MÉDIA, CONTAGEM, etc.) que operam em dados de campos personalizados para exibir métricas calculadas em gráficos. As fórmulas não são calculadas no nível de cada tarefa, mas sim agregam dados de vários registros para fins de visualização.

## Exemplo Básico

Crie um campo de fórmula para cálculos de gráficos:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## Exemplo Avançado

Crie uma fórmula de moeda com cálculos complexos:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de fórmula |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `FORMULA` |
| `projectId` | String! | ✅ Sim | O ID do projeto onde este campo será criado |
| `formula` | JSON | Não | Definição da fórmula para cálculos de gráficos |
| `description` | String | Não | Texto de ajuda exibido aos usuários |

### Estrutura da Fórmula

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## Funções Suportadas

### Funções de Agregação de Gráficos

Os campos de fórmula suportam as seguintes funções de agregação para cálculos de gráficos:

| Função | Descrição | Enum ChartFunction |
|--------|-----------|--------------------|
| `SUM` | Soma de todos os valores | `SUM` |
| `AVERAGE` | Média de valores numéricos | `AVERAGE` |
| `AVERAGEA` | Média excluindo zeros e nulos | `AVERAGEA` |
| `COUNT` | Contagem de valores | `COUNT` |
| `COUNTA` | Contagem excluindo zeros e nulos | `COUNTA` |
| `MAX` | Valor máximo | `MAX` |
| `MIN` | Valor mínimo | `MIN` |

**Nota**: Essas funções são usadas no campo `display.function` e operam em dados agregados para visualizações de gráficos. Expressões matemáticas complexas ou cálculos em nível de campo não são suportados.

## Tipos de Exibição

### Exibição de Número

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

Resultado: `1250.75`

### Exibição de Moeda

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

Resultado: `$1,250.75`

### Exibição de Porcentagem

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

Resultado: `87.5%`

## Edição de Campos de Fórmula

Atualize campos de fórmula existentes:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## Processamento de Fórmulas

### Contexto de Cálculo de Gráficos

Os campos de fórmula são processados no contexto de segmentos de gráficos e painéis:
- Os cálculos ocorrem quando os gráficos são renderizados ou atualizados
- Os resultados são armazenados em `ChartSegment.formulaResult` como valores decimais
- O processamento é gerenciado através de uma fila dedicada do BullMQ chamada 'formula'
- As atualizações são publicadas para assinantes do painel para atualizações em tempo real

### Formatação de Exibição

A função `getFormulaDisplayValue` formata os resultados calculados com base no tipo de exibição:
- **NÚMERO**: Exibe como número simples com precisão opcional
- **PORCENTAGEM**: Adiciona sufixo % com precisão opcional  
- **MOEDA**: Formata usando o código de moeda especificado

## Armazenamento de Resultados da Fórmula

Os resultados são armazenados no campo `formulaResult`:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo de fórmula |
| `number` | Float | Resultado numérico calculado |
| `formulaResult` | JSON | Resultado completo com formatação de exibição |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi calculado pela última vez |

## Contexto de Dados

### Fonte de Dados do Gráfico

Os campos de fórmula operam dentro do contexto da fonte de dados do gráfico:
- As fórmulas agregam valores de campos personalizados em tarefas em um projeto
- A função de agregação especificada em `display.function` determina o cálculo
- Os resultados são computados usando funções de agregação SQL (avg, sum, count, etc.)
- Os cálculos são realizados no nível do banco de dados para eficiência

## Exemplos Comuns de Fórmulas

### Orçamento Total (Exibição de Gráfico)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### Pontuação Média (Exibição de Gráfico)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### Contagem de Tarefas (Exibição de Gráfico)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## Permissões Necessárias

As operações de campo personalizado seguem as permissões padrão baseadas em função:

| Ação | Função Necessária |
|------|-------------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**Nota**: Os papéis específicos necessários dependem da configuração de papéis personalizados do seu projeto. Não existem constantes de permissão especiais como CUSTOM_FIELDS_CREATE.

## Tratamento de Erros

### Erro de Validação
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Campo Personalizado Não Encontrado
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

## Melhores Práticas

### Design de Fórmulas
- Use nomes claros e descritivos para campos de fórmula
- Adicione descrições explicando a lógica de cálculo
- Teste fórmulas com dados de exemplo antes da implantação
- Mantenha as fórmulas simples e legíveis

### Otimização de Desempenho
- Evite dependências de fórmulas profundamente aninhadas
- Use referências de campo específicas em vez de curingas
- Considere estratégias de cache para cálculos complexos
- Monitore o desempenho da fórmula em grandes projetos

### Qualidade dos Dados
- Valide os dados de origem antes de usá-los em fórmulas
- Lide com valores vazios ou nulos de forma apropriada
- Use precisão apropriada para tipos de exibição
- Considere casos extremos em cálculos

## Casos de Uso Comuns

1. **Rastreamento Financeiro**
   - Cálculos de orçamento
   - Demonstrações de lucro/prejuízo
   - Análise de custos
   - Projeções de receita

2. **Gerenciamento de Projetos**
   - Percentuais de conclusão
   - Utilização de recursos
   - Cálculos de cronograma
   - Métricas de desempenho

3. **Controle de Qualidade**
   - Pontuações médias
   - Taxas de aprovação/reprovação
   - Métricas de qualidade
   - Rastreamento de conformidade

4. **Inteligência de Negócios**
   - Cálculos de KPI
   - Análise de tendências
   - Métricas comparativas
   - Valores de painel

## Limitações

- As fórmulas são apenas para agregações de gráficos/painéis, não para cálculos em nível de tarefa
- Limitadas às sete funções de agregação suportadas (SOMA, MÉDIA, etc.)
- Sem expressões matemáticas complexas ou cálculos de campo para campo
- Não é possível referenciar múltiplos campos em uma única fórmula
- Os resultados são visíveis apenas em gráficos e painéis
- O campo `logic` é apenas para texto de exibição, não para lógica de cálculo real

## Recursos Relacionados

- [Campos Numéricos](/api/5.custom%20fields/number) - Para valores numéricos estáticos
- [Campos de Moeda](/api/5.custom%20fields/currency) - Para valores monetários
- [Campos de Referência](/api/5.custom%20fields/reference) - Para dados entre projetos
- [Campos de Pesquisa](/api/5.custom%20fields/lookup) - Para dados agregados
- [Visão Geral de Campos Personalizados](/api/5.custom%20fields/2.list-custom-fields) - Conceitos gerais