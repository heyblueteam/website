---
title: Campo Personalizado de Email
description: Crie campos de email para armazenar e validar endereços de email
---

Os campos personalizados de email permitem que você armazene endereços de email em registros com validação integrada. Eles são ideais para rastrear informações de contato, emails de responsáveis ou qualquer dado relacionado a email em seus projetos.

## Exemplo Básico

Crie um campo de email simples:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de email com descrição:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ Sim | Nome exibido do campo de email |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `EMAIL` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

## Definindo Valores de Email

Para definir ou atualizar um valor de email em um registro:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de email |
| `text` | String | Não | Endereço de email a ser armazenado |

## Criando Registros com Valores de Email

Ao criar um novo registro com valores de email:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## Campos de Resposta

### Resposta CustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | Identificador único para o campo personalizado |
| `name` | String! | Nome exibido do campo de email |
| `type` | CustomFieldType! | O tipo do campo (EMAIL) |
| `description` | String | Texto de ajuda para o campo |
| `value` | JSON | Contém o valor do email (veja abaixo) |
| `createdAt` | DateTime! | Quando o campo foi criado |
| `updatedAt` | DateTime! | Quando o campo foi modificado pela última vez |

**Importante**: Os valores de email são acessados através do campo `customField.value.text`, não diretamente na resposta.

## Consultando Valores de Email

Ao consultar registros com campos personalizados de email, acesse o email através do caminho `customField.value.text`:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

A resposta incluirá o email na estrutura aninhada:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## Validação de Email

### Validação de Formulário
Quando campos de email são usados em formulários, eles validam automaticamente o formato do email:
- Usa regras padrão de validação de email
- Remove espaços em branco da entrada
- Rejeita formatos de email inválidos

### Regras de Validação
- Deve conter um símbolo `@`
- Deve ter um formato de domínio válido
- Espaços em branco no início/fim são removidos automaticamente
- Formatos de email comuns são aceitos

### Exemplos de Email Válidos
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### Exemplos de Email Inválidos
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## Notas Importantes

### API Direta vs Formulários
- **Formulários**: Validação automática de email é aplicada
- **API Direta**: Sem validação - qualquer texto pode ser armazenado
- **Recomendação**: Use formulários para entrada do usuário para garantir a validação

### Formato de Armazenamento
- Endereços de email são armazenados como texto simples
- Sem formatação ou análise especial
- Sensibilidade a maiúsculas: campos personalizados de EMAIL são armazenados de forma sensível a maiúsculas (diferente de emails de autenticação de usuário que são normalizados para minúsculas)
- Sem limitações de comprimento máximo além das restrições do banco de dados (limite de 16MB)

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## Respostas de Erro

### Formato de Email Inválido (Somente Formulários)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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

## Melhores Práticas

### Entrada de Dados
- Sempre valide endereços de email em sua aplicação
- Use campos de email apenas para endereços de email reais
- Considere usar formulários para entrada do usuário para obter validação automática

### Qualidade dos Dados
- Remova espaços em branco antes de armazenar
- Considere normalização de maiúsculas (tipicamente minúsculas)
- Valide o formato de email antes de operações importantes

### Considerações de Privacidade
- Endereços de email são armazenados como texto simples
- Considere regulamentos de privacidade de dados (GDPR, CCPA)
- Implemente controles de acesso apropriados

## Casos de Uso Comuns

1. **Gerenciamento de Contatos**
   - Endereços de email de clientes
   - Informações de contato de fornecedores
   - Emails de membros da equipe
   - Detalhes de contato de suporte

2. **Gerenciamento de Projetos**
   - Emails de partes interessadas
   - Emails de contato para aprovação
   - Recipientes de notificações
   - Emails de colaboradores externos

3. **Suporte ao Cliente**
   - Endereços de email de clientes
   - Contatos de tickets de suporte
   - Contatos de escalonamento
   - Endereços de email para feedback

4. **Vendas e Marketing**
   - Endereços de email de leads
   - Listas de contatos de campanhas
   - Informações de contato de parceiros
   - Emails de fontes de referência

## Recursos de Integração

### Com Automações
- Acione ações quando campos de email forem atualizados
- Envie notificações para endereços de email armazenados
- Crie tarefas de acompanhamento com base em mudanças de email

### Com Pesquisas
- Referencie dados de email de outros registros
- Agregue listas de email de várias fontes
- Encontre registros por endereço de email

### Com Formulários
- Validação automática de email
- Verificação de formato de email
- Remoção de espaços em branco

## Limitações

- Sem verificação ou validação de email integrada além da verificação de formato
- Sem recursos de UI específicos para email (como links de email clicáveis)
- Armazenados como texto simples sem criptografia
- Sem capacidades de composição ou envio de email
- Sem armazenamento de metadados de email (nome exibido, etc.)
- Chamadas diretas da API ignoram validação (apenas formulários validam)

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para dados de texto não relacionados a email
- [Campos de URL](/api/custom-fields/url) - Para endereços de sites
- [Campos de Telefone](/api/custom-fields/phone) - Para números de telefone
- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais