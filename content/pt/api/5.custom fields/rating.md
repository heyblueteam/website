---
title: Campo Personalizado de Avaliação
description: Crie campos de avaliação para armazenar classificações numéricas com escalas e validações configuráveis
---

Os campos personalizados de avaliação permitem que você armazene classificações numéricas em registros com valores mínimos e máximos configuráveis. Eles são ideais para classificações de desempenho, pontuações de satisfação, níveis de prioridade ou qualquer dado baseado em escala numérica em seus projetos.

## Exemplo Básico

Crie um campo de avaliação simples com escala padrão de 0-5:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## Exemplo Avançado

Crie um campo de avaliação com escala e descrição personalizadas:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Requerido | Descrição |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Sim | Nome exibido do campo de avaliação |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `RATING` |
| `projectId` | String! | ✅ Sim | O ID do projeto onde este campo será criado |
| `description` | String | Não | Texto de ajuda exibido aos usuários |
| `min` | Float | Não | Valor mínimo de avaliação (sem padrão) |
| `max` | Float | Não | Valor máximo de avaliação |

## Definindo Valores de Avaliação

Para definir ou atualizar um valor de avaliação em um registro:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Requerido | Descrição |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de avaliação |
| `value` | String! | ✅ Sim | Valor de avaliação como string (dentro do intervalo configurado) |

## Criando Registros com Valores de Avaliação

Ao criar um novo registro com valores de avaliação:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
    }
  }
}
```

## Campos de Resposta

### Resposta TodoCustomField

| Campo | Tipo | Descrição |
|-------|------|-------------|
| `id` | String! | Identificador único para o valor do campo |
| `customField` | CustomField! | A definição do campo personalizado |
| `value` | Float | O valor de avaliação armazenado (acessado via customField.value) |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

**Nota**: O valor de avaliação é acessado via `customField.value.number` em consultas.

### Resposta CustomField

| Campo | Tipo | Descrição |
|-------|------|-------------|
| `id` | String! | Identificador único para o campo |
| `name` | String! | Nome exibido do campo de avaliação |
| `type` | CustomFieldType! | Sempre `RATING` |
| `min` | Float | Valor mínimo de avaliação permitido |
| `max` | Float | Valor máximo de avaliação permitido |
| `description` | String | Texto de ajuda para o campo |

## Validação de Avaliação

### Restrições de Valor
- Os valores de avaliação devem ser numéricos (tipo Float)
- Os valores devem estar dentro do intervalo mínimo/máximo configurado
- Se nenhum mínimo for especificado, não há valor padrão
- O valor máximo é opcional, mas recomendado

### Regras de Validação
**Importante**: A validação ocorre apenas ao enviar formulários, não ao usar `setTodoCustomField` diretamente.

- A entrada é analisada como um número float (ao usar formulários)
- Deve ser maior ou igual ao valor mínimo (ao usar formulários)
- Deve ser menor ou igual ao valor máximo (ao usar formulários)
- `setTodoCustomField` aceita qualquer valor de string sem validação

### Exemplos de Avaliação Válidos
Para um campo com min=1, max=5:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### Exemplos de Avaliação Inválidos
Para um campo com min=1, max=5:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## Opções de Configuração

### Configuração da Escala de Avaliação
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### Escalas de Avaliação Comuns
- **1-5 Estrelas**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 Desempenho**: `min: 1, max: 10`
- **0-100 Porcentagem**: `min: 0, max: 100`
- **Escala Personalizada**: Qualquer intervalo numérico

## Permissões Requeridas

As operações de campo personalizado seguem as permissões padrão baseadas em função:

| Ação | Função Requerida |
|--------|---------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**Nota**: Os papéis específicos exigidos dependem da configuração de papéis personalizados do seu projeto e das permissões em nível de campo.

## Respostas de Erro

### Erro de Validação (Apenas Formulários)
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

**Importante**: A validação do valor de avaliação (restrições min/max) ocorre apenas ao enviar formulários, não ao usar `setTodoCustomField` diretamente.

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

### Design da Escala
- Use escalas de avaliação consistentes em campos semelhantes
- Considere a familiaridade do usuário (1-5 estrelas, 0-10 NPS)
- Defina valores mínimos apropriados (0 vs 1)
- Defina um significado claro para cada nível de avaliação

### Qualidade dos Dados
- Valide os valores de avaliação antes de armazenar
- Use precisão decimal de forma apropriada
- Considere arredondar para fins de exibição
- Forneça orientações claras sobre os significados das avaliações

### Experiência do Usuário
- Exiba escalas de avaliação visualmente (estrelas, barras de progresso)
- Mostre o valor atual e os limites da escala
- Forneça contexto para os significados das avaliações
- Considere valores padrão para novos registros

## Casos de Uso Comuns

1. **Gestão de Desempenho**
   - Avaliações de desempenho de funcionários
   - Pontuações de qualidade de projetos
   - Avaliações de conclusão de tarefas
   - Avaliações de nível de habilidade

2. **Feedback do Cliente**
   - Avaliações de satisfação
   - Pontuações de qualidade do produto
   - Avaliações de experiência de serviço
   - Net Promoter Score (NPS)

3. **Prioridade e Importância**
   - Níveis de prioridade de tarefas
   - Avaliações de urgência
   - Pontuações de avaliação de risco
   - Avaliações de impacto

4. **Garantia de Qualidade**
   - Avaliações de revisão de código
   - Pontuações de qualidade de testes
   - Qualidade da documentação
   - Avaliações de aderência a processos

## Recursos de Integração

### Com Automação
- Acionar ações com base em limites de avaliação
- Enviar notificações para avaliações baixas
- Criar tarefas de acompanhamento para avaliações altas
- Roteirizar trabalho com base em valores de avaliação

### Com Pesquisas
- Calcular médias de avaliações entre registros
- Encontrar registros por intervalos de avaliação
- Referenciar dados de avaliação de outros registros
- Agregar estatísticas de avaliação

### Com Frontend Blue
- Validação automática de intervalo em contextos de formulário
- Controles de entrada de avaliação visual
- Feedback de validação em tempo real
- Opções de entrada de estrela ou slider

## Rastreamento de Atividades

As alterações no campo de avaliação são rastreadas automaticamente:
- Valores de avaliação antigos e novos são registrados
- A atividade mostra alterações numéricas
- Carimbos de data/hora para todas as atualizações de avaliação
- Atribuição de usuário para alterações

## Limitações

- Apenas valores numéricos são suportados
- Sem exibição visual de avaliação integrada (estrelas, etc.)
- A precisão decimal depende da configuração do banco de dados
- Sem armazenamento de metadados de avaliação (comentários, contexto)
- Sem agregação ou estatísticas de avaliação automáticas
- Sem conversão de avaliação integrada entre escalas
- **Crítico**: A validação min/max funciona apenas em formulários, não via `setTodoCustomField`

## Recursos Relacionados

- [Campos Numéricos](/api/5.custom%20fields/number) - Para dados numéricos gerais
- [Campos de Porcentagem](/api/5.custom%20fields/percent) - Para valores percentuais
- [Campos de Seleção](/api/5.custom%20fields/select-single) - Para classificações de escolha discreta
- [Visão Geral de Campos Personalizados](/api/5.custom%20fields/2.list-custom-fields) - Conceitos gerais