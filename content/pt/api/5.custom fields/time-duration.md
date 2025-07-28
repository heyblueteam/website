---
title: Campo Personalizado de Duração de Tempo
description: Crie campos de duração de tempo calculados que rastreiam o tempo entre eventos em seu fluxo de trabalho
---

Os campos personalizados de Duração de Tempo calculam e exibem automaticamente a duração entre dois eventos em seu fluxo de trabalho. Eles são ideais para rastrear tempos de processamento, tempos de resposta, tempos de ciclo ou quaisquer métricas baseadas em tempo em seus projetos.

## Exemplo Básico

Crie um campo de duração de tempo simples que rastreia quanto tempo as tarefas levam para serem concluídas:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## Exemplo Avançado

Crie um campo de duração de tempo complexo que rastreia o tempo entre mudanças de campo personalizado com um alvo de SLA:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput (TIME_DURATION)

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `name` | String! | ✅ Sim | Nome de exibição do campo de duração |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `TIME_DURATION` |
| `description` | String | Não | Texto de ajuda exibido aos usuários |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Sim | Como exibir a duração |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Sim | Configuração do evento de início |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Sim | Configuração do evento de término |
| `timeDurationTargetTime` | Float | Não | Duração alvo em segundos para monitoramento de SLA |

### CustomFieldTimeDurationInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-----------|
| `type` | CustomFieldTimeDurationType! | ✅ Sim | Tipo de evento a ser rastreado |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Sim | `FIRST` ou `LAST` ocorrência |
| `customFieldId` | String | Conditional | Obrigatório para tipo `TODO_CUSTOM_FIELD` |
| `customFieldOptionIds` | [String!] | Conditional | Obrigatório para alterações de campo de seleção |
| `todoListId` | String | Conditional | Obrigatório para tipo `TODO_MOVED` |
| `tagId` | String | Conditional | Obrigatório para tipo `TODO_TAG_ADDED` |
| `assigneeId` | String | Conditional | Obrigatório para tipo `TODO_ASSIGNEE_ADDED` |

### Valores de CustomFieldTimeDurationType

| Valor | Descrição |
|-------|-----------|
| `TODO_CREATED_AT` | Quando o registro foi criado |
| `TODO_CUSTOM_FIELD` | Quando um valor de campo personalizado mudou |
| `TODO_DUE_DATE` | Quando a data de vencimento foi definida |
| `TODO_MARKED_AS_COMPLETE` | Quando o registro foi marcado como completo |
| `TODO_MOVED` | Quando o registro foi movido para uma lista diferente |
| `TODO_TAG_ADDED` | Quando uma tag foi adicionada ao registro |
| `TODO_ASSIGNEE_ADDED` | Quando um responsável foi adicionado ao registro |

### Valores de CustomFieldTimeDurationCondition

| Valor | Descrição |
|-------|-----------|
| `FIRST` | Use a primeira ocorrência do evento |
| `LAST` | Use a última ocorrência do evento |

### Valores de CustomFieldTimeDurationDisplayType

| Valor | Descrição | Exemplo |
|-------|-----------|---------|
| `FULL_DATE` | Formato Dias:Horas:Minutos:Segundos | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Escrito por extenso | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Numérico com unidades | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Duração em dias apenas | `"2.5"` (2.5 days) |
| `HOURS` | Duração em horas apenas | `"60"` (60 hours) |
| `MINUTES` | Duração em minutos apenas | `"3600"` (3600 minutes) |
| `SECONDS` | Duração em segundos apenas | `"216000"` (216000 seconds) |

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `number` | Float | A duração em segundos |
| `value` | Float | Alias para número (duração em segundos) |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi atualizado pela última vez |

### Resposta CustomField (TIME_DURATION)

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Formato de exibição para a duração |
| `timeDurationStart` | CustomFieldTimeDuration | Configuração do evento de início |
| `timeDurationEnd` | CustomFieldTimeDuration | Configuração do evento de término |
| `timeDurationTargetTime` | Float | Duração alvo em segundos (para monitoramento de SLA) |

## Cálculo de Duração

### Como Funciona
1. **Evento de Início**: O sistema monitora o evento de início especificado
2. **Evento de Término**: O sistema monitora o evento de término especificado
3. **Cálculo**: Duração = Tempo de Término - Tempo de Início
4. **Armazenamento**: Duração armazenada em segundos como um número
5. **Exibição**: Formatada de acordo com a configuração `timeDurationDisplay`

### Gatilhos de Atualização
Os valores de duração são recalculados automaticamente quando:
- Registros são criados ou atualizados
- Valores de campos personalizados mudam
- Tags são adicionadas ou removidas
- Responsáveis são adicionados ou removidos
- Registros são movidos entre listas
- Registros são marcados como completos/incompletos

## Leitura de Valores de Duração

### Consultar Campos de Duração
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### Valores de Exibição Formatados
Os valores de duração são automaticamente formatados com base na configuração `timeDurationDisplay`:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## Exemplos Comuns de Configuração

### Tempo de Conclusão da Tarefa
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### Duração da Mudança de Status
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### Tempo em Lista Específica
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### Tempo de Resposta de Atribuição
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Respostas de Erro

### Configuração Inválida
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### Campo Referenciado Não Encontrado
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

### Opções Obrigatórias Ausentes
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Notas Importantes

### Cálculo Automático
- Campos de duração são **somente leitura** - os valores são calculados automaticamente
- Você não pode definir manualmente valores de duração via API
- Cálculos ocorrem de forma assíncrona através de trabalhos em segundo plano
- Os valores são atualizados automaticamente quando eventos de gatilho ocorrem

### Considerações de Desempenho
- Cálculos de duração são enfileirados e processados de forma assíncrona
- Um grande número de campos de duração pode impactar o desempenho
- Considere a frequência de eventos de gatilho ao projetar campos de duração
- Use condições específicas para evitar recalculos desnecessários

### Valores Nulos
Campos de duração mostrarão `null` quando:
- O evento de início ainda não ocorreu
- O evento de término ainda não ocorreu
- A configuração faz referência a entidades inexistentes
- O cálculo encontra um erro

## Melhores Práticas

### Design de Configuração
- Use tipos de eventos específicos em vez de genéricos quando possível
- Escolha condições apropriadas `FIRST` vs `LAST` com base em seu fluxo de trabalho
- Teste cálculos de duração com dados de exemplo antes da implantação
- Documente a lógica do seu campo de duração para membros da equipe

### Formatação de Exibição
- Use `FULL_DATE_SUBSTRING` para o formato mais legível
- Use `FULL_DATE` para exibição compacta e de largura consistente
- Use `FULL_DATE_STRING` para relatórios e documentos formais
- Use `DAYS`, `HOURS`, `MINUTES`, ou `SECONDS` para exibições numéricas simples
- Considere as limitações de espaço da sua interface ao escolher o formato

### Monitoramento de SLA com Tempo Alvo
Ao usar `timeDurationTargetTime`:
- Defina a duração alvo em segundos
- Compare a duração real com a alvo para conformidade com o SLA
- Use em painéis para destacar itens atrasados
- Exemplo: SLA de resposta de 24 horas = 86400 segundos

### Integração ao Fluxo de Trabalho
- Projete campos de duração para corresponder aos seus processos de negócios reais
- Use dados de duração para melhoria e otimização de processos
- Monitore tendências de duração para identificar gargalos no fluxo de trabalho
- Configure alertas para limites de duração se necessário

## Casos de Uso Comuns

1. **Desempenho do Processo**
   - Tempos de conclusão de tarefas
   - Tempos de ciclo de revisão
   - Tempos de processamento de aprovação
   - Tempos de resposta

2. **Monitoramento de SLA**
   - Tempo até a primeira resposta
   - Tempos de resolução
   - Prazos de escalonamento
   - Conformidade com o nível de serviço

3. **Análise de Fluxo de Trabalho**
   - Identificação de gargalos
   - Otimização de processos
   - Métricas de desempenho da equipe
   - Cronometragem de garantia de qualidade

4. **Gerenciamento de Projetos**
   - Durações de fase
   - Cronometragem de marcos
   - Tempo de alocação de recursos
   - Prazos de entrega

## Limitações

- Campos de duração são **somente leitura** e não podem ser definidos manualmente
- Os valores são calculados de forma assíncrona e podem não estar imediatamente disponíveis
- Requer que os gatilhos de eventos apropriados sejam configurados em seu fluxo de trabalho
- Não é possível calcular durações para eventos que ainda não ocorreram
- Limitado ao rastreamento de tempo entre eventos discretos (não rastreamento contínuo de tempo)
- Sem alertas ou notificações de SLA integrados
- Não é possível agregar múltiplos cálculos de duração em um único campo

## Recursos Relacionados

- [Campos Numéricos](/api/custom-fields/number) - Para valores numéricos manuais
- [Campos de Data](/api/custom-fields/date) - Para rastreamento de datas específicas
- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais
- [Automatizações](/api/automations) - Para acionar ações com base em limites de duração