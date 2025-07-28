---
title: Campo Personalizado de Conversão de Moeda
description: Crie campos que convertem automaticamente valores de moeda usando taxas de câmbio em tempo real
---

Os campos personalizados de Conversão de Moeda convertem automaticamente valores de um campo de MOEDA de origem para diferentes moedas-alvo usando taxas de câmbio em tempo real. Esses campos são atualizados automaticamente sempre que o valor da moeda de origem muda.

As taxas de conversão são fornecidas pela [Frankfurter API](https://github.com/hakanensari/frankfurter), um serviço de código aberto que rastreia as taxas de câmbio de referência publicadas pelo [Banco Central Europeu](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Isso garante conversões de moeda precisas, confiáveis e atualizadas para as necessidades do seu negócio internacional.

## Exemplo Básico

Crie um campo simples de conversão de moeda:

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## Exemplo Avançado

Crie um campo de conversão com uma data específica para taxas históricas:

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## Processo Completo de Configuração

Configurar um campo de conversão de moeda requer três etapas:

### Etapa 1: Crie um Campo de MOEDA de Origem

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### Etapa 2: Crie o Campo CURRENCY_CONVERSION

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### Etapa 3: Crie Opções de Conversão

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de conversão |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | Não | ID do campo de MOEDA de origem para converter |
| `conversionDateType` | String | Não | Estratégia de data para taxas de câmbio (veja abaixo) |
| `conversionDate` | String | Não | String de data para conversão (baseada em conversionDateType) |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

**Nota**: Campos personalizados são automaticamente associados ao projeto com base no contexto atual do projeto do usuário. Nenhum `projectId` parâmetro é necessário.

### Tipos de Data de Conversão

| Tipo | Descrição | Parâmetro conversionDate |
|------|-----------|--------------------------|
| `currentDate` | Usa taxas de câmbio em tempo real | Não necessário |
| `specificDate` | Usa taxas de uma data fixa | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Usa data de outro campo | "todoDueDate" or DATE field ID |

## Criando Opções de Conversão

As opções de conversão definem quais pares de moedas podem ser convertidos:

### CreateCustomFieldOptionInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `customFieldId` | String! | ✅ Sim | ID do campo CURRENCY_CONVERSION |
| `title` | String! | ✅ Sim | Nome exibido para esta opção de conversão |
| `currencyConversionFrom` | String! | ✅ Sim | Código da moeda de origem ou "Qualquer" |
| `currencyConversionTo` | String! | ✅ Sim | Código da moeda-alvo |

### Usando "Qualquer" como Origem

O valor especial "Qualquer" como `currencyConversionFrom` cria uma opção de fallback:

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

Esta opção será usada quando nenhuma correspondência específica de par de moedas for encontrada.

## Como a Conversão Automática Funciona

1. **Atualização de Valor**: Quando um valor é definido no campo de MOEDA de origem
2. **Correspondência de Opção**: O sistema encontra a opção de conversão correspondente com base na moeda de origem
3. **Busca de Taxa**: Recupera a taxa de câmbio da Frankfurter API
4. **Cálculo**: Multiplica o valor de origem pela taxa de câmbio
5. **Armazenamento**: Salva o valor convertido com o código da moeda-alvo

### Fluxo de Exemplo

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## Conversões Baseadas em Data

### Usando a Data Atual

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

As conversões são atualizadas com as taxas de câmbio atuais sempre que o valor de origem muda.

### Usando uma Data Específica

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

Sempre usa as taxas de câmbio da data especificada.

### Usando Data de Campo

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

Usa a data de outro campo (ou a data de vencimento da tarefa ou um campo personalizado de DATA).

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo de conversão |
| `number` | Float | O valor convertido |
| `currency` | String | O código da moeda-alvo |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi atualizado pela última vez |

## Fonte da Taxa de Câmbio

Blue usa a **Frankfurter API** para taxas de câmbio:
- API de código aberto hospedada pelo Banco Central Europeu
- Atualizações diárias com taxas de câmbio oficiais
- Suporta taxas históricas desde 1999
- Gratuita e confiável para uso comercial

## Tratamento de Erros

### Falhas de Conversão

Quando a conversão falha (erro da API, moeda inválida, etc.):
- O valor convertido é definido como `0`
- A moeda-alvo ainda é armazenada
- Nenhum erro é exibido para o usuário

### Cenários Comuns

| Cenário | Resultado |
|---------|----------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Nenhuma opção correspondente | Uses "Any" option if available |
| Missing source value | Nenhuma conversão realizada |

## Permissões Necessárias

A gestão de campos personalizados requer acesso a nível de projeto:

| Função | Pode Criar/Atualizar Campos |
|--------|-----------------------------|
| `OWNER` | ✅ Sim |
| `ADMIN` | ✅ Sim |
| `MEMBER` | ❌ Não |
| `CLIENT` | ❌ Não |

As permissões de visualização para valores convertidos seguem as regras padrão de acesso a registros.

## Melhores Práticas

### Configuração de Opções
- Crie pares de moedas específicos para conversões comuns
- Adicione uma opção de fallback "Qualquer" para flexibilidade
- Use títulos descritivos para as opções

### Seleção de Estratégia de Data
- Use `currentDate` para rastreamento financeiro ao vivo
- Use `specificDate` para relatórios históricos
- Use `fromDateField` para taxas específicas de transações

### Considerações de Desempenho
- Vários campos de conversão são atualizados em paralelo
- Chamadas à API são feitas apenas quando o valor de origem muda
- Conversões de mesma moeda pulam chamadas à API

## Casos de Uso Comuns

1. **Projetos Multimoeda**
   - Rastrear custos do projeto em moedas locais
   - Relatar orçamento total na moeda da empresa
   - Comparar valores entre regiões

2. **Vendas Internacionais**
   - Converter valores de negócios para a moeda de relatório
   - Rastrear receita em várias moedas
   - Conversão histórica para negócios fechados

3. **Relatórios Financeiros**
   - Conversões de moeda no final do período
   - Demonstrações financeiras consolidadas
   - Orçamento vs. real na moeda local

4. **Gestão de Contratos**
   - Converter valores de contratos na data da assinatura
   - Rastrear cronogramas de pagamento em várias moedas
   - Avaliação de risco cambial

## Limitações

- Sem suporte para conversões de criptomoeda
- Não é possível definir valores convertidos manualmente (sempre calculados)
- Precisão fixa de 2 casas decimais para todos os valores convertidos
- Sem suporte para taxas de câmbio personalizadas
- Sem cache de taxas de câmbio (chamada à API fresca para cada conversão)
- Depende da disponibilidade da Frankfurter API

## Recursos Relacionados

- [Campos de Moeda](/api/custom-fields/currency) - Campos de origem para conversões
- [Campos de Data](/api/custom-fields/date) - Para conversões baseadas em data
- [Campos de Fórmula](/api/custom-fields/formula) - Cálculos alternativos
- [Visão Geral de Campos Personalizados](/custom-fields/list-custom-fields) - Conceitos gerais