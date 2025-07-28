---
title: Campo Personalizado de País
description: Crie campos de seleção de país com validação de código de país ISO
---

Os campos personalizados de país permitem que você armazene e gerencie informações de país para registros. O campo suporta tanto nomes de países quanto códigos de país ISO Alpha-2.

**Importante**: O comportamento de validação e conversão de países difere significativamente entre as mutações:
- **createTodo**: Valida e converte automaticamente nomes de países para códigos ISO
- **setTodoCustomField**: Aceita qualquer valor sem validação

## Exemplo Básico

Crie um campo de país simples:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de país com descrição:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome de exibição do campo de país |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `COUNTRY` |
| `description` | String | Não | Texto de ajuda exibido aos usuários |

**Nota**: O `projectId` não é passado na entrada, mas é determinado pelo contexto GraphQL (normalmente a partir dos cabeçalhos de solicitação ou autenticação).

## Definindo Valores de País

Os campos de país armazenam dados em dois campos de banco de dados:
- **`countryCodes`**: Armazena códigos de país ISO Alpha-2 como uma string separada por vírgulas no banco de dados (retornado como array via API)
- **`text`**: Armazena texto de exibição ou nomes de países como uma string

### Compreendendo os Parâmetros

A mutação `setTodoCustomField` aceita dois parâmetros opcionais para campos de país:

| Parâmetro | Tipo | Obrigatório | Descrição | O que faz |
|-----------|------|-------------|-----------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado | - |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de país | - |
| `countryCodes` | [String!] | Não | Array de códigos de país ISO Alpha-2 | Stored in the `countryCodes` field |
| `text` | String | Não | Texto de exibição ou nomes de países | Stored in the `text` field |

**Importante**: 
- Em `setTodoCustomField`: Ambos os parâmetros são opcionais e armazenados de forma independente
- Em `createTodo`: O sistema define automaticamente ambos os campos com base na sua entrada (você não pode controlá-los de forma independente)

### Opção 1: Usando Apenas Códigos de País

Armazene códigos ISO validados sem texto de exibição:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

Resultado: `countryCodes` = `["US"]`, `text` = `null`

### Opção 2: Usando Apenas Texto

Armazene texto de exibição sem códigos validados:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

Resultado: `countryCodes` = `null`, `text` = `"United States"`

**Nota**: Ao usar `setTodoCustomField`, nenhuma validação ocorre, independentemente de qual parâmetro você use. Os valores são armazenados exatamente como fornecidos.

### Opção 3: Usando Ambos (Recomendado)

Armazene tanto códigos validados quanto texto de exibição:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

Resultado: `countryCodes` = `["US"]`, `text` = `"United States"`

### Vários Países

Armazene vários países usando arrays:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## Criando Registros com Valores de País

Ao criar registros, a mutação `createTodo` **valida e converte automaticamente** os valores de país. Esta é a única mutação que realiza a validação de país:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      text
      countryCodes
    }
  }
}
```

### Formatos de Entrada Aceitos

| Tipo de Entrada | Exemplo | Resultado |
|-----------------|---------|----------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **Não suportado** - tratado como um único valor inválido |
| Mixed format | `"United States, CA"` | **Não suportado** - tratado como um único valor inválido |

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `text` | String | Texto de exibição (nomes de países) |
| `countryCodes` | [String!] | Array de códigos de país ISO Alpha-2 |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

## Padrões de País

Blue utiliza o padrão **ISO 3166-1 Alpha-2** para códigos de país:

- Códigos de país de duas letras (ex: US, GB, FR, DE)
- A validação usando a biblioteca `i18n-iso-countries` **ocorre apenas em createTodo**
- Suporta todos os países oficialmente reconhecidos

### Exemplos de Códigos de País

| País | Código ISO |
|------|------------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

Para a lista oficial completa de códigos de país ISO 3166-1 alpha-2, visite a [Plataforma de Navegação Online ISO](https://www.iso.org/obp/ui/#search/code/).

## Validação

**A validação ocorre apenas na mutação `createTodo`**:

1. **Código ISO Válido**: Aceita qualquer código ISO Alpha-2 válido
2. **Nome do País**: Converte automaticamente nomes de países reconhecidos em códigos
3. **Entrada Inválida**: Lança `CustomFieldValueParseError` para valores não reconhecidos

**Nota**: A mutação `setTodoCustomField` não realiza NENHUMA validação e aceita qualquer valor de string.

### Exemplo de Erro

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Recursos de Integração

### Campos de Pesquisa
Os campos de país podem ser referenciados por campos personalizados de PESQUISA, permitindo que você extraia dados de país de registros relacionados.

### Automação
Use valores de país em condições de automação:
- Filtrar ações por países específicos
- Enviar notificações com base no país
- Roteirizar tarefas com base em regiões geográficas

### Formulários
Os campos de país em formulários validam automaticamente a entrada do usuário e convertem nomes de países em códigos.

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## Respostas de Erro

### Valor de País Inválido
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### Incompatibilidade de Tipo de Campo
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## Melhores Práticas

### Manipulação de Entrada
- Use `createTodo` para validação e conversão automáticas
- Use `setTodoCustomField` com cuidado, pois ignora a validação
- Considere validar entradas em sua aplicação antes de usar `setTodoCustomField`
- Exiba os nomes completos dos países na interface do usuário para clareza

### Qualidade dos Dados
- Valide as entradas de país no ponto de entrada
- Use formatos consistentes em todo o seu sistema
- Considere agrupamentos regionais para relatórios

### Vários Países
- Use o suporte a arrays em `setTodoCustomField` para múltiplos países
- Múltiplos países em `createTodo` **não são suportados** via campo de valor
- Armazene códigos de país como array em `setTodoCustomField` para um tratamento adequado

## Casos de Uso Comuns

1. **Gerenciamento de Clientes**
   - Localização da sede do cliente
   - Destinos de envio
   - Jurisdições fiscais

2. **Rastreamento de Projetos**
   - Localização do projeto
   - Localizações dos membros da equipe
   - Alvos de mercado

3. **Conformidade e Legal**
   - Jurisdições regulatórias
   - Requisitos de residência de dados
   - Controles de exportação

4. **Vendas e Marketing**
   - Atribuições de território
   - Segmentação de mercado
   - Direcionamento de campanhas

## Limitações

- Suporta apenas códigos ISO 3166-1 Alpha-2 (códigos de 2 letras)
- Sem suporte embutido para subdivisões de países (estados/províncias)
- Sem ícones de bandeira de país automáticos (apenas baseados em texto)
- Não pode validar códigos de país históricos
- Sem agrupamento embutido de regiões ou continentes
- **A validação funciona apenas em `createTodo`, não em `setTodoCustomField`**
- **Múltiplos países não suportados no campo de valor `createTodo`**
- **Códigos de país armazenados como string separada por vírgulas, não como array verdadeiro**

## Recursos Relacionados

- [Visão Geral de Campos Personalizados](/custom-fields/list-custom-fields) - Conceitos gerais de campos personalizados
- [Campos de Pesquisa](/api/custom-fields/lookup) - Referenciar dados de país de outros registros
- [API de Formulários](/api/forms) - Incluir campos de país em formulários personalizados