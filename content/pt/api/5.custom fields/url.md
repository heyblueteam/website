---
title: Campo Personalizado de URL
description: Crie campos de URL para armazenar endereços de sites e links
---

Os campos personalizados de URL permitem que você armazene endereços de sites e links em seus registros. Eles são ideais para rastrear sites de projetos, links de referência, URLs de documentação ou quaisquer recursos baseados na web relacionados ao seu trabalho.

## Exemplo Básico

Crie um campo de URL simples:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de URL com descrição:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
    }
  ) {
    id
    name
    type
    description
  }
}
```

## Parâmetros de Entrada

### CreateCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Sim | Nome exibido do campo de URL |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `URL` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

**Nota:** O `projectId` é passado como um argumento separado para a mutação, não como parte do objeto de entrada.

## Definindo Valores de URL

Para definir ou atualizar um valor de URL em um registro:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### Parâmetros de SetTodoCustomFieldInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro a ser atualizado |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de URL |
| `text` | String! | ✅ Sim | Endereço URL a ser armazenado |

## Criando Registros com Valores de URL

Ao criar um novo registro com valores de URL:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
| `text` | String | O endereço URL armazenado |
| `todo` | Todo! | O registro ao qual este valor pertence |
| `createdAt` | DateTime! | Quando o valor foi criado |
| `updatedAt` | DateTime! | Quando o valor foi modificado pela última vez |

## Validação de URL

### Implementação Atual
- **API Direta**: Nenhuma validação de formato de URL está atualmente em vigor
- **Formulários**: A validação de URL está planejada, mas não está atualmente ativa
- **Armazenamento**: Qualquer valor de string pode ser armazenado em campos de URL

### Validação Planejada
Versões futuras incluirão:
- Validação de protocolo HTTP/HTTPS
- Verificação de formato de URL válido
- Validação de nome de domínio
- Adição automática de prefixo de protocolo

### Formatos de URL Recomendados
Embora não estejam atualmente em vigor, use estes formatos padrão:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## Notas Importantes

### Formato de Armazenamento
- URLs são armazenadas como texto simples sem modificação
- Nenhuma adição automática de protocolo (http://, https://)
- Sensibilidade a maiúsculas e minúsculas preservada conforme digitado
- Nenhuma codificação/decodificação de URL realizada

### API Direta vs Formulários
- **Formulários**: Validação de URL planejada (não atualmente ativa)
- **API Direta**: Sem validação - qualquer texto pode ser armazenado
- **Recomendação**: Valide URLs em seu aplicativo antes de armazenar

### Campos de URL vs Texto
- **URL**: Semânticamente destinado a endereços da web
- **TEXT_SINGLE**: Texto geral de uma linha
- **Backend**: Armazenamento e validação atualmente idênticos
- **Frontend**: Diferentes componentes de UI para entrada de dados

## Permissões Necessárias

Operações de campo personalizado usam permissões baseadas em funções:

| Ação | Função Necessária |
|------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**Nota:** As permissões são verificadas com base nas funções dos usuários no projeto, não em constantes de permissão específicas.

## Respostas de Erro

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

### Validação de Campo Necessário (Somente Formulários)
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## Melhores Práticas

### Padrões de Formato de URL
- Sempre inclua o protocolo (http:// ou https://)
- Use HTTPS sempre que possível para segurança
- Teste URLs antes de armazenar para garantir que sejam acessíveis
- Considere usar URLs encurtadas para fins de exibição

### Qualidade de Dados
- Valide URLs em seu aplicativo antes de armazenar
- Verifique erros comuns de digitação (protocolos ausentes, domínios incorretos)
- Padronize formatos de URL em toda a sua organização
- Considere a acessibilidade e disponibilidade da URL

### Considerações de Segurança
- Tenha cautela com URLs fornecidas por usuários
- Valide domínios se restringindo a sites específicos
- Considere a verificação de URLs para conteúdo malicioso
- Use URLs HTTPS ao lidar com dados sensíveis

## Filtragem e Pesquisa

### Pesquisa por Contém
Campos de URL suportam pesquisa por substring:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### Capacidades de Pesquisa
- Correspondência de substring sem diferenciação entre maiúsculas e minúsculas
- Correspondência parcial de domínio
- Pesquisa de caminho e parâmetro
- Sem filtragem específica de protocolo

## Casos de Uso Comuns

1. **Gerenciamento de Projetos**
   - Sites de projetos
   - Links de documentação
   - URLs de repositórios
   - Sites de demonstração

2. **Gerenciamento de Conteúdo**
   - Materiais de referência
   - Links de origem
   - Recursos de mídia
   - Artigos externos

3. **Suporte ao Cliente**
   - Sites de clientes
   - Documentação de suporte
   - Artigos da base de conhecimento
   - Tutoriais em vídeo

4. **Vendas e Marketing**
   - Sites de empresas
   - Páginas de produtos
   - Materiais de marketing
   - Perfis de redes sociais

## Recursos de Integração

### Com Pesquisas
- URLs de referência de outros registros
- Encontrar registros por domínio ou padrão de URL
- Exibir recursos web relacionados
- Agregar links de várias fontes

### Com Formulários
- Componentes de entrada específicos de URL
- Validação planejada para formato de URL adequado
- Capacidades de visualização de links (frontend)
- Exibição de URL clicável

### Com Relatórios
- Rastrear uso e padrões de URL
- Monitorar links quebrados ou inacessíveis
- Categorizar por domínio ou protocolo
- Exportar listas de URLs para análise

## Limitações

### Limitações Atuais
- Sem validação de formato de URL ativa
- Sem adição automática de protocolo
- Sem verificação de link ou verificação de acessibilidade
- Sem encurtamento ou expansão de URL
- Sem geração de favicon ou visualização

### Restrições de Automação
- Não disponível como campos de gatilho de automação
- Não pode ser usado em atualizações de campo de automação
- Pode ser referenciado em condições de automação
- Disponível em modelos de email e webhooks

### Restrições Gerais
- Sem funcionalidade de visualização de link embutida
- Sem encurtamento automático de URL
- Sem rastreamento de cliques ou análises
- Sem verificação de expiração de URL
- Sem verificação de URL maliciosas

## Melhorias Futuras

### Recursos Planejados
- Validação de protocolo HTTP/HTTPS
- Padrões de validação regex personalizados
- Adição automática de prefixo de protocolo
- Verificação de acessibilidade de URL

### Melhorias Potenciais
- Geração de visualização de link
- Exibição de favicon
- Integração de encurtamento de URL
- Capacidades de rastreamento de cliques
- Detecção de links quebrados

## Recursos Relacionados

- [Campos de Texto](/api/custom-fields/text-single) - Para dados de texto não-URL
- [Campos de Email](/api/custom-fields/email) - Para endereços de email
- [Visão Geral de Campos Personalizados](/api/custom-fields/2.list-custom-fields) - Conceitos gerais

## Migração de Campos de Texto

Se você está migrando de campos de texto para campos de URL:

1. **Crie um campo de URL** com o mesmo nome e configuração
2. **Exporte os valores de texto existentes** para verificar se são URLs válidas
3. **Atualize os registros** para usar o novo campo de URL
4. **Exclua o campo de texto antigo** após a migração bem-sucedida
5. **Atualize os aplicativos** para usar componentes de UI específicos de URL

### Exemplo de Migração
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```