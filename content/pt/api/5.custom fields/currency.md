---
title: Campo Personalizado de Moeda
description: Crie campos de moeda para rastrear valores monetários com formatação e validação adequadas
---

Os campos personalizados de moeda permitem que você armazene e gerencie valores monetários com códigos de moeda associados. O campo suporta 72 moedas diferentes, incluindo as principais moedas fiduciárias e criptomoedas, com formatação automática e restrições de mínimo/máximo opcionais.

## Exemplo Básico

Crie um campo de moeda simples:

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## Exemplo Avançado

Crie um campo de moeda com restrições de validação:

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de moeda |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `CURRENCY` |
| `currency` | String | Não | Código de moeda padrão (código ISO de 3 letras) |
| `min` | Float | Não | Valor mínimo permitido (armazenado, mas não aplicado em atualizações) |
| `max` | Float | Não | Valor máximo permitido (armazenado, mas não aplicado em atualizações) |
| `description` | String | Não | Texto de ajuda exibido aos usuários |

**Nota**: O contexto do projeto é determinado automaticamente a partir da sua autenticação. Você deve ter acesso ao projeto onde está criando o campo.

## Definindo Valores de Moeda

Para definir ou atualizar um valor de moeda em um registro:

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de moeda |
| `number` | Float! | ✅ Sim | O valor monetário |
| `currency` | String! | ✅ Sim | Código de moeda de 3 letras |

## Criando Registros com Valores de Moeda

Ao criar um novo registro com valores de moeda:

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### Formato de Entrada para Criação

Ao criar registros, os valores de moeda são passados de maneira diferente:

| Parâmetro | Tipo | Descrição |
|-----------|------|-----------|
| `customFieldId` | String! | ID do campo de moeda |
| `value` | String! | Valor como uma string (ex.: "1500.50") |
| `currency` | String! | Código de moeda de 3 letras |

## Moedas Suportadas

Blue suporta 72 moedas, incluindo 70 moedas fiduciárias e 2 criptomoedas:

### Moedas Fiduciárias

#### Américas
| Moeda | Código | Nome |
|-------|--------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Bolívar Soberano Venezuelano |
| Boliviano Boliviano | `BOB` | Boliviano Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### Europa
| Moeda | Código | Nome |
|-------|--------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Coroa Norueguesa | `NOK` | Coroa Norueguesa |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### Ásia-Pacífico
| Moeda | Código | Nome |
|-------|--------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### Oriente Médio e África
| Moeda | Código | Nome |
|-------|--------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### Criptomoedas
| Moeda | Código |
|-------|--------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `number` | Float | O valor monetário |
| `currency` | String | O código de moeda de 3 letras |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

## Formatação de Moeda

O sistema formata automaticamente os valores de moeda com base na localidade:

- **Posicionamento do símbolo**: Posiciona corretamente os símbolos de moeda (antes/depois)
- **Separadores decimais**: Usa separadores específicos da localidade (. ou ,)
- **Separadores de milhar**: Aplica agrupamento apropriado
- **Casas decimais**: Mostra de 0 a 2 casas decimais com base no valor
- **Tratamento especial**: USD/CAD mostram o prefixo do código de moeda para clareza

### Exemplos de Formatação

| Valor | Moeda | Exibição |
|-------|-------|----------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## Validação

### Validação de Montante
- Deve ser um número válido
- Restrições de mínimo/máximo são armazenadas com a definição do campo, mas não aplicadas durante as atualizações de valor
- Suporta até 2 casas decimais para exibição (precisão total armazenada internamente)

### Validação do Código de Moeda
- Deve ser um dos 72 códigos de moeda suportados
- Sensível a maiúsculas (use letras maiúsculas)
- Códigos inválidos retornam um erro

## Recursos de Integração

### Fórmulas
Campos de moeda podem ser usados em campos personalizados de FÓRMULA para cálculos:
- Somar múltiplos campos de moeda
- Calcular porcentagens
- Realizar operações aritméticas

### Conversão de Moeda
Use campos de CONVERSÃO_DE_MOEDA para converter automaticamente entre moedas (veja [Campos de Conversão de Moeda](/api/custom-fields/currency-conversion))

### Automação
Valores de moeda podem acionar automações com base em:
- Limites de montante
- Tipo de moeda
- Mudanças de valor

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|----------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**Nota**: Embora qualquer membro do projeto possa criar campos personalizados, a capacidade de definir valores depende das permissões baseadas em funções configuradas para cada campo.

## Respostas de Erro

### Valor de Moeda Inválido
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

Esse erro ocorre quando:
- O código da moeda não é um dos 72 códigos suportados
- O formato do número é inválido
- O valor não pode ser analisado corretamente

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

### Seleção de Moeda
- Defina uma moeda padrão que corresponda ao seu mercado principal
- Use códigos de moeda ISO 4217 de forma consistente
- Considere a localização do usuário ao escolher padrões

### Restrições de Valor
- Defina valores mínimos/máximos razoáveis para evitar erros de entrada de dados
- Use 0 como mínimo para campos que não devem ser negativos
- Considere seu caso de uso ao definir máximos

### Projetos Multi-Moeda
- Use uma moeda base consistente para relatórios
- Implemente campos de CONVERSÃO_DE_MOEDA para conversão automática
- Documente qual moeda deve ser usada para cada campo

## Casos de Uso Comuns

1. **Orçamento de Projetos**
   - Rastreamento de orçamento de projetos
   - Estimativas de custo
   - Rastreamento de despesas

2. **Vendas e Negócios**
   - Valores de negócios
   - Montantes de contratos
   - Rastreamento de receita

3. **Planejamento Financeiro**
   - Montantes de investimento
   - Rodadas de financiamento
   - Metas financeiras

4. **Negócios Internacionais**
   - Preços em múltiplas moedas
   - Rastreamento de câmbio
   - Transações transfronteiriças

## Limitações

- Máximo de 2 casas decimais para exibição (embora mais precisão seja armazenada)
- Sem conversão de moeda embutida nos campos de MOEDA padrão
- Não é possível misturar moedas em um único valor de campo
- Sem atualizações automáticas da taxa de câmbio (use CONVERSÃO_DE_MOEDA para isso)
- Símbolos de moeda não são personalizáveis

## Recursos Relacionados

- [Campos de Conversão de Moeda](/api/custom-fields/currency-conversion) - Conversão automática de moeda
- [Campos Numéricos](/api/custom-fields/number) - Para valores numéricos não monetários
- [Campos de Fórmula](/api/custom-fields/formula) - Calcular com valores de moeda
- [Campos Personalizados de Lista](/api/custom-fields/list-custom-fields) - Consultar todos os campos personalizados em um projeto