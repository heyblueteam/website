---
title: Campo Personalizado de ID Único
description: Crie campos de identificador único gerados automaticamente com numeração sequencial e formatação personalizada
---

Os campos personalizados de ID único geram automaticamente identificadores sequenciais e únicos para seus registros. Eles são perfeitos para criar números de ticket, IDs de pedido, números de fatura ou qualquer sistema de identificador sequencial em seu fluxo de trabalho.

## Exemplo Básico

Crie um campo de ID único simples com auto-sequenciamento:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## Exemplo Avançado

Crie um campo de ID único formatado com prefixo e preenchimento com zeros:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput (UNIQUE_ID)

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de ID único |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `UNIQUE_ID` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |
| `useSequenceUniqueId` | Boolean | Não | Ativar auto-sequenciamento (padrão: falso) |
| `prefix` | String | Não | Texto prefixo para IDs gerados (ex: "TAREFA-") |
| `sequenceDigits` | Int | Não | Número de dígitos para preenchimento com zeros |
| `sequenceStartingNumber` | Int | Não | Número inicial para a sequência |

## Opções de Configuração

### Auto-Sequenciamento (`useSequenceUniqueId`)
- **true**: Gera automaticamente IDs sequenciais quando os registros são criados
- **false** ou **undefined**: Entrada manual necessária (funciona como um campo de texto)

### Prefixo (`prefix`)
- Prefixo de texto opcional adicionado a todos os IDs gerados
- Exemplos: "TAREFA-", "PED-", "BUG-", "REQ-"
- Sem limite de comprimento, mas mantenha razoável para exibição

### Dígitos da Sequência (`sequenceDigits`)
- Número de dígitos para preenchimento com zeros do número da sequência
- Exemplo: `sequenceDigits: 3` produz `001`, `002`, `003`
- Se não especificado, nenhum preenchimento é aplicado

### Número Inicial (`sequenceStartingNumber`)
- O primeiro número na sequência
- Exemplo: `sequenceStartingNumber: 1000` começa em 1000, 1001, 1002...
- Se não especificado, começa em 1 (comportamento padrão)

## Formato do ID Gerado

O formato final do ID combina todas as opções de configuração:

```
{prefix}{paddedSequenceNumber}
```

### Exemplos de Formato

| Configuração | IDs Gerados |
|---------------|-------------|
| Sem opções | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## Lendo Valores de ID Único

### Consultar Registros com IDs Únicos
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### Formato de Resposta
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## Geração Automática de IDs

### Quando os IDs São Gerados
- **Criação de Registro**: IDs são atribuídos automaticamente quando novos registros são criados
- **Adição de Campo**: Ao adicionar um campo UNIQUE_ID a registros existentes, um trabalho em segundo plano é enfileirado (implementação do trabalhador pendente)
- **Processamento em Segundo Plano**: A geração de ID para novos registros ocorre de forma síncrona via gatilhos de banco de dados

### Processo de Geração
1. **Gatilho**: Novo registro é criado ou campo UNIQUE_ID é adicionado
2. **Busca de Sequência**: O sistema encontra o próximo número de sequência disponível
3. **Atribuição de ID**: O número da sequência é atribuído ao registro
4. **Atualização do Contador**: O contador de sequência é incrementado para registros futuros
5. **Formatação**: O ID é formatado com prefixo e preenchimento ao ser exibido

### Garantias de Exclusividade
- **Restrições de Banco de Dados**: Restrição única nos IDs de sequência dentro de cada campo
- **Operações Atômicas**: A geração de sequência utiliza bloqueios de banco de dados para evitar duplicatas
- **Escopo de Projeto**: Sequências são independentes por projeto
- **Proteção contra Condição de Corrida**: Solicitações concorrentes são tratadas de forma segura

## Modo Manual vs Automático

### Modo Automático (`useSequenceUniqueId: true`)
- IDs são gerados automaticamente via gatilhos de banco de dados
- A numeração sequencial é garantida
- A geração de sequência atômica previne duplicatas
- IDs formatados combinam prefixo + número da sequência preenchido

### Modo Manual (`useSequenceUniqueId: false` ou `undefined`)
- Funciona como um campo de texto regular
- Os usuários podem inserir valores personalizados via `setTodoCustomField` com o parâmetro `text`
- Sem geração automática
- Sem imposição de exclusividade além das restrições do banco de dados

## Definindo Valores Manuais (Somente Modo Manual)

Quando `useSequenceUniqueId` é falso, você pode definir valores manualmente:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## Campos de Resposta

### Resposta TodoCustomField (UNIQUE_ID)

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `sequenceId` | Int | O número da sequência gerado (preenchido para campos UNIQUE_ID) |
| `text` | String | O valor de texto formatado (combina prefixo + sequência preenchida) |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi atualizado pela última vez |

### Resposta CustomField (UNIQUE_ID)

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `useSequenceUniqueId` | Boolean | Se o auto-sequenciamento está ativado |
| `prefix` | String | Texto prefixo para IDs gerados |
| `sequenceDigits` | Int | Número de dígitos para preenchimento com zeros |
| `sequenceStartingNumber` | Int | Número inicial para a sequência |

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## Respostas de Erro

### Erro de Configuração de Campo
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Erro de Permissão
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## Notas Importantes

### IDs Gerados Automaticamente
- **Somente Leitura**: IDs gerados automaticamente não podem ser editados manualmente
- **Permanente**: Uma vez atribuídos, os IDs de sequência não mudam
- **Cronológico**: IDs refletem a ordem de criação
- **Escopo**: Sequências são independentes por projeto

### Considerações de Desempenho
- A geração de ID para novos registros é síncrona via gatilhos de banco de dados
- A geração de sequência utiliza `FOR UPDATE` bloqueios para operações atômicas
- Um sistema de trabalho em segundo plano existe, mas a implementação do trabalhador está pendente
- Considere números iniciais de sequência para projetos de alto volume

### Migração e Atualizações
- Adicionar auto-sequenciamento a registros existentes enfileira um trabalho em segundo plano (trabalhador pendente)
- Alterar configurações de sequência afeta apenas registros futuros
- IDs existentes permanecem inalterados quando as configurações são atualizadas
- Contadores de sequência continuam a partir do máximo atual

## Melhores Práticas

### Design de Configuração
- Escolha prefixos descritivos que não conflitem com outros sistemas
- Use preenchimento de dígitos apropriado para seu volume esperado
- Defina números iniciais razoáveis para evitar conflitos
- Teste a configuração com dados de exemplo antes da implantação

### Diretrizes de Prefixo
- Mantenha os prefixos curtos e memoráveis (2-5 caracteres)
- Use maiúsculas para consistência
- Inclua separadores (hífens, sublinhados) para legibilidade
- Evite caracteres especiais que possam causar problemas em URLs ou sistemas

### Planejamento de Sequência
- Estime seu volume de registros para escolher o preenchimento de dígitos apropriado
- Considere o crescimento futuro ao definir números iniciais
- Planeje diferentes intervalos de sequência para diferentes tipos de registros
- Documente seus esquemas de ID para referência da equipe

## Casos de Uso Comuns

1. **Sistemas de Suporte**
   - Números de ticket: `TICK-001`, `TICK-002`
   - IDs de caso: `CASE-2024-001`
   - Solicitações de suporte: `SUP-001`

2. **Gestão de Projetos**
   - IDs de tarefa: `TASK-001`, `TASK-002`
   - Itens de sprint: `SPRINT-001`
   - Números de entregáveis: `DEL-001`

3. **Operações Comerciais**
   - Números de pedido: `ORD-2024-001`
   - IDs de fatura: `INV-001`
   - Ordens de compra: `PO-001`

4. **Gestão de Qualidade**
   - Relatórios de bugs: `BUG-001`
   - IDs de casos de teste: `TEST-001`
   - Números de revisão: `REV-001`

## Recursos de Integração

### Com Automação
- Acione ações quando IDs únicos forem atribuídos
- Use padrões de ID em regras de automação
- Referencie IDs em modelos de e-mail e notificações

### Com Consultas
- Referencie IDs únicos de outros registros
- Encontre registros por ID único
- Exiba identificadores de registros relacionados

### Com Relatórios
- Agrupe e filtre por padrões de ID
- Acompanhe tendências de atribuição de ID
- Monitore o uso e as lacunas da sequência

## Limitações

- **Somente Sequencial**: IDs são atribuídos em ordem cronológica
- **Sem Lacunas**: Registros excluídos deixam lacunas nas sequências
- **Sem Reutilização**: Números de sequência nunca são reutilizados
- **Escopo de Projeto**: Não é possível compartilhar sequências entre projetos
- **Restrições de Formato**: Opções de formatação limitadas
- **Sem Atualizações em Massa**: Não é possível atualizar em massa IDs de sequência existentes
- **Sem Lógica Personalizada**: Não é possível implementar regras de geração de ID personalizadas

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para identificadores de texto manuais
- [Campos Numéricos](/api/custom-fields/number) - Para sequências numéricas
- [Visão Geral de Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceitos gerais
- [Automatizações](/api/automations) - Para regras de automação baseadas em ID