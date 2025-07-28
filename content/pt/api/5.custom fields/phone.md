---
title: Campo Personalizado de Telefone
description: Crie campos de telefone para armazenar e validar números de telefone com formatação internacional
---

Os campos personalizados de telefone permitem que você armazene números de telefone em registros com validação integrada e formatação internacional. Eles são ideais para rastrear informações de contato, contatos de emergência ou qualquer dado relacionado a telefone em seus projetos.

## Exemplo Básico

Crie um campo de telefone simples:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de telefone com descrição:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ Sim | Nome exibido do campo de telefone |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `PHONE` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

**Nota**: Campos personalizados são automaticamente associados ao projeto com base no contexto do projeto atual do usuário. Nenhum parâmetro `projectId` é necessário.

## Definindo Valores de Telefone

Para definir ou atualizar um valor de telefone em um registro:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de telefone |
| `text` | String | Não | Número de telefone com código do país |
| `regionCode` | String | Não | Código do país (detectado automaticamente) |

**Nota**: Embora `text` seja opcional no esquema, um número de telefone é necessário para que o campo tenha significado. Ao usar `setTodoCustomField`, nenhuma validação é realizada - você pode armazenar qualquer valor de texto e regionCode. A detecção automática ocorre apenas durante a criação do registro.

## Criando Registros com Valores de Telefone

Ao criar um novo registro com valores de telefone:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
    }
  }
}
```

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `text` | String | O número de telefone formatado (formato internacional) |
| `regionCode` | String | O código do país (por exemplo, "US", "GB", "CA") |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

## Validação de Números de Telefone

**Importante**: A validação e formatação de números de telefone ocorrem apenas ao criar novos registros via `createTodo`. Ao atualizar valores de telefone existentes usando `setTodoCustomField`, nenhuma validação é realizada e os valores são armazenados conforme fornecidos.

### Formatos Aceitos (Durante a Criação do Registro)
Os números de telefone devem incluir um código do país em um destes formatos:

- **Formato E.164 (preferido)**: `+12345678900`
- **Formato internacional**: `+1 234 567 8900`
- **Internacional com pontuação**: `+1 (234) 567-8900`
- **Código do país com traços**: `+1-234-567-8900`

**Nota**: Formatos nacionais sem código do país (como `(234) 567-8900`) serão rejeitados durante a criação do registro.

### Regras de Validação (Durante a Criação do Registro)
- Usa libphonenumber-js para análise e validação
- Aceita vários formatos internacionais de números de telefone
- Detecta automaticamente o país a partir do número
- Formata o número em formato de exibição internacional (por exemplo, `+1 234 567 8900`)
- Extrai e armazena o código do país separadamente (por exemplo, `US`)

### Exemplos de Telefone Válidos
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### Exemplos de Telefone Inválidos
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## Formato de Armazenamento

Ao criar registros com números de telefone:
- **text**: Armazenado em formato internacional (por exemplo, `+1 234 567 8900`) após validação
- **regionCode**: Armazenado como código do país ISO (por exemplo, `US`, `GB`, `CA`) detectado automaticamente

Ao atualizar via `setTodoCustomField`:
- **text**: Armazenado exatamente como fornecido (sem formatação)
- **regionCode**: Armazenado exatamente como fornecido (sem validação)

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## Respostas de Erro

### Formato de Telefone Inválido
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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

### Código do País Ausente
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## Melhores Práticas

### Entrada de Dados
- Sempre inclua o código do país nos números de telefone
- Use o formato E.164 para consistência
- Valide os números antes de armazená-los para operações importantes
- Considere preferências regionais para formatação de exibição

### Qualidade dos Dados
- Armazene números em formato internacional para compatibilidade global
- Use regionCode para recursos específicos do país
- Valide números de telefone antes de operações críticas (SMS, chamadas)
- Considere implicações de fuso horário para o tempo de contato

### Considerações Internacionais
- O código do país é detectado e armazenado automaticamente
- Os números são formatados no padrão internacional
- Preferências de exibição regionais podem usar regionCode
- Considere convenções de discagem locais ao exibir

## Casos de Uso Comuns

1. **Gerenciamento de Contatos**
   - Números de telefone de clientes
   - Informações de contato de fornecedores
   - Números de telefone de membros da equipe
   - Detalhes de contato de suporte

2. **Contatos de Emergência**
   - Números de contato de emergência
   - Informações de contato de plantão
   - Contatos de resposta a crises
   - Números de telefone de escalonamento

3. **Suporte ao Cliente**
   - Números de telefone de clientes
   - Números de retorno de chamada de suporte
   - Números de telefone de verificação
   - Números de contato para acompanhamento

4. **Vendas e Marketing**
   - Números de telefone de leads
   - Listas de contatos de campanhas
   - Informações de contato de parceiros
   - Telefones de fontes de referência

## Recursos de Integração

### Com Automação
- Acione ações quando campos de telefone forem atualizados
- Envie notificações SMS para números de telefone armazenados
- Crie tarefas de acompanhamento com base em alterações de telefone
- Roteie chamadas com base em dados de número de telefone

### Com Pesquisas
- Referencie dados de telefone de outros registros
- Agregue listas de telefone de várias fontes
- Encontre registros por número de telefone
- Faça referência cruzada a informações de contato

### Com Formulários
- Validação automática de telefone
- Verificação de formato internacional
- Detecção de código do país
- Feedback de formato em tempo real

## Limitações

- Requer código do país para todos os números
- Sem capacidades integradas de SMS ou chamadas
- Sem verificação de número de telefone além da verificação de formato
- Sem armazenamento de metadados de telefone (operadora, tipo, etc.)
- Números em formato nacional sem código do país são rejeitados
- Sem formatação automática de número de telefone na interface além do padrão internacional

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para dados de texto não relacionados a telefone
- [Campos de Email](/api/custom-fields/email) - Para endereços de email
- [Campos de URL](/api/custom-fields/url) - Para endereços de sites
- [Visão Geral de Campos Personalizados](/custom-fields/list-custom-fields) - Conceitos gerais