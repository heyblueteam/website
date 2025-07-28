---
title: Criar Painel
description: Crie um novo painel para visualização de dados e relatórios no Blue
---

## Criar um Painel

A mutação `createDashboard` permite que você crie um novo painel dentro da sua empresa ou projeto. Os painéis são ferramentas de visualização poderosas que ajudam as equipes a acompanhar métricas, monitorar o progresso e tomar decisões baseadas em dados.

### Exemplo Básico

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### Painel Específico do Projeto

Crie um painel associado a um projeto específico:

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## Parâmetros de Entrada

### CreateDashboardInput

| Parâmetro | Tipo | Obrigatório | Descrição |
|-----------|------|-------------|-------------|
| `companyId` | String! | ✅ Sim | O ID da empresa onde o painel será criado |
| `title` | String! | ✅ Sim | O nome do painel. Deve ser uma string não vazia |
| `projectId` | String | Não | ID opcional de um projeto para associar a este painel |

## Campos de Resposta

A mutação retorna um objeto completo `Dashboard`:

| Campo | Tipo | Descrição |
|-------|------|-------------|
| `id` | String! | Identificador único para o painel criado |
| `title` | String! | O título do painel conforme fornecido |
| `companyId` | String! | A empresa à qual este painel pertence |
| `projectId` | String | O ID do projeto associado (se fornecido) |
| `project` | Project | O objeto do projeto associado (se projectId foi fornecido) |
| `createdBy` | User! | O usuário que criou o painel (você) |
| `dashboardUsers` | [DashboardUser!]! | Lista de usuários com acesso (inicialmente apenas o criador) |
| `createdAt` | DateTime! | Timestamp de quando o painel foi criado |
| `updatedAt` | DateTime! | Timestamp da última modificação (igual a createdAt para novos painéis) |

### Campos DashboardUser

Quando um painel é criado, o criador é automaticamente adicionado como um usuário do painel:

| Campo | Tipo | Descrição |
|-------|------|-------------|
| `id` | String! | Identificador único para a relação do usuário do painel |
| `user` | User! | O objeto do usuário com acesso ao painel |
| `role` | DashboardRole! | O papel do usuário (o criador tem acesso total) |
| `dashboard` | Dashboard! | Referência de volta ao painel |

## Permissões Necessárias

Qualquer usuário autenticado que pertença à empresa especificada pode criar painéis. Não há requisitos de função especiais.

| Status do Usuário | Pode Criar Painel |
|-------------------|-------------------|
| Company Member | ✅ Sim |
| Membro Não da Empresa | ❌ Não |
| Unauthenticated | ❌ Não |

## Respostas de Erro

### Empresa Inválida
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Usuário Não na Empresa
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### Projeto Inválido
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### Título Vazio
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Notas Importantes

- **Propriedade automática**: O usuário que cria o painel automaticamente se torna seu proprietário com permissões totais
- **Associação de projeto**: Se você fornecer um `projectId`, ele deve pertencer à mesma empresa
- **Permissões iniciais**: Apenas o criador tem acesso inicialmente. Use `editDashboard` para adicionar mais usuários
- **Requisitos de título**: Os títulos dos painéis devem ser strings não vazias. Não há requisito de exclusividade
- **Membresia da empresa**: Você deve ser membro da empresa para criar painéis nela

## Fluxo de Trabalho de Criação de Painel

1. **Crie o painel** usando esta mutação
2. **Configure gráficos e widgets** usando a interface do construtor de painéis
3. **Adicione membros da equipe** usando a mutação `editDashboard` com `dashboardUsers`
4. **Configure filtros e intervalos de datas** através da interface do painel
5. **Compartilhe ou incorpore** o painel usando seu ID exclusivo

## Casos de Uso

1. **Painéis executivos**: Crie visões gerais de alto nível das métricas da empresa
2. **Acompanhamento de projetos**: Construa painéis específicos de projetos para monitorar o progresso
3. **Desempenho da equipe**: Acompanhe a produtividade da equipe e métricas de realização
4. **Relatórios para clientes**: Crie painéis para relatórios voltados para clientes
5. **Monitoramento em tempo real**: Configure painéis para dados operacionais ao vivo

## Melhores Práticas

1. **Convenções de nomenclatura**: Use títulos claros e descritivos que indiquem o propósito do painel
2. **Associação de projeto**: Vincule painéis a projetos quando forem específicos de projetos
3. **Gerenciamento de acesso**: Adicione membros da equipe imediatamente após a criação para colaboração
4. **Organização**: Crie uma hierarquia de painéis usando padrões de nomenclatura consistentes

## Operações Relacionadas

- [Listar Painéis](/api/dashboards/) - Recupere todos os painéis de uma empresa ou projeto
- [Editar Painel](/api/dashboards/rename-dashboard) - Renomeie o painel ou gerencie usuários
- [Copiar Painel](/api/dashboards/copy-dashboard) - Duplicar um painel existente
- [Excluir Painel](/api/dashboards/delete-dashboard) - Remover um painel