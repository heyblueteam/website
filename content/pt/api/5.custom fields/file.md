---
title: Campo Personalizado de Arquivo
description: Crie campos de arquivo para anexar documentos, imagens e outros arquivos a registros
---

Os campos personalizados de arquivo permitem que você anexe vários arquivos a registros. Os arquivos são armazenados de forma segura no AWS S3 com rastreamento abrangente de metadados, validação de tipo de arquivo e controles de acesso adequados.

## Exemplo Básico

Crie um campo de arquivo simples:

```graphql
mutation CreateFileField {
  createCustomField(input: {
    name: "Attachments"
    type: FILE
  }) {
    id
    name
    type
  }
}
```

## Exemplo Avançado

Crie um campo de arquivo com descrição:

```graphql
mutation CreateDetailedFileField {
  createCustomField(input: {
    name: "Project Documents"
    type: FILE
    description: "Upload project-related documents, images, and files"
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

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `name` | String! | ✅ Sim | Nome de exibição do campo de arquivo |
| `type` | CustomFieldType! | ✅ Sim | Deve ser `FILE` |
| `description` | String | Não | Texto de ajuda exibido para os usuários |

**Nota**: Campos personalizados são automaticamente associados ao projeto com base no contexto do projeto atual do usuário. Nenhum `projectId` parâmetro é necessário.

## Processo de Upload de Arquivo

### Passo 1: Fazer Upload do Arquivo

Primeiro, faça o upload do arquivo para obter um UID de arquivo:

```graphql
mutation UploadFile {
  uploadFile(input: {
    file: $file  # File upload variable
    companyId: "company_123"
    projectId: "proj_123"
  }) {
    id
    uid
    name
    size
    type
    extension
    status
  }
}
```

### Passo 2: Anexar Arquivo ao Registro

Em seguida, anexe o arquivo enviado a um registro:

```graphql
mutation AttachFileToRecord {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "file_field_456"
    fileUid: "file_uid_from_upload"
  }) {
    id
    file {
      uid
      name
      size
      type
    }
  }
}
```

## Gerenciando Anexos de Arquivo

### Adicionando Arquivos Únicos

```graphql
mutation AddFileToField {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  }) {
    id
    position
    file {
      uid
      name
      size
      type
      extension
    }
  }
}
```

### Removendo Arquivos

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### Operações em Lote de Arquivos

Atualize vários arquivos de uma vez usando customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## Parâmetros de Entrada para Upload de Arquivo

### UploadFileInput

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `file` | Upload! | ✅ Sim | Arquivo a ser enviado |
| `companyId` | String! | ✅ Sim | ID da empresa para armazenamento de arquivo |
| `projectId` | String | Não | ID do projeto para arquivos específicos do projeto |

### Parâmetros de Entrada para Gerenciamento de Arquivos

| Parâmetro | Tipo | Necessário | Descrição |
|-----------|------|------------|-----------|
| `todoId` | String! | ✅ Sim | ID do registro |
| `customFieldId` | String! | ✅ Sim | ID do campo personalizado de arquivo |
| `fileUid` | String! | ✅ Sim | Identificador único do arquivo enviado |

## Armazenamento e Limites de Arquivo

### Limites de Tamanho de Arquivo

| Tipo de Limite | Tamanho |
|----------------|--------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### Tipos de Arquivo Suportados

#### Imagens
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### Vídeos
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### Áudio
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### Documentos
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### Arquivos
- `zip`, `rar`, `7z`, `tar`, `gz`

#### Código/Textos
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### Arquitetura de Armazenamento

- **Armazenamento**: AWS S3 com estrutura de pastas organizada
- **Formato de Caminho**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **Segurança**: URLs assinadas para acesso seguro
- **Backup**: Redundância automática do S3

## Campos de Resposta

### Resposta de Arquivo

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | ID do banco de dados |
| `uid` | String! | Identificador único do arquivo |
| `name` | String! | Nome original do arquivo |
| `size` | Float! | Tamanho do arquivo em bytes |
| `type` | String! | Tipo MIME |
| `extension` | String! | Extensão do arquivo |
| `status` | FileStatus | PENDENTE ou CONFIRMADO (nulo) |
| `shared` | Boolean! | Se o arquivo é compartilhado |
| `createdAt` | DateTime! | Timestamp de upload |

### Resposta TodoCustomFieldFile

| Campo | Tipo | Descrição |
|-------|------|-----------|
| `id` | ID! | ID do registro de junção |
| `uid` | String! | Identificador único |
| `position` | Float! | Ordem de exibição |
| `file` | File! | Objeto de arquivo associado |
| `todoCustomField` | TodoCustomField! | Campo personalizado pai |
| `createdAt` | DateTime! | Quando o arquivo foi anexado |

## Criando Registros com Arquivos

Ao criar registros, você pode anexar arquivos usando seus UIDs:

```graphql
mutation CreateRecordWithFiles {
  createTodo(input: {
    title: "Project deliverables"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "file_field_id"
      customFieldOptionIds: ["file_uid_1", "file_uid_2"]
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
      todoCustomFieldFiles {
        id
        position
        file {
          uid
          name
          size
          type
        }
      }
    }
  }
}
```

## Validação e Segurança de Arquivos

### Validação de Upload

- **Verificação de Tipo MIME**: Valida contra tipos permitidos
- **Validação de Extensão de Arquivo**: Fallback para `application/octet-stream`
- **Limites de Tamanho**: Aplicados no momento do upload
- **Sanitização de Nome de Arquivo**: Remove caracteres especiais

### Controle de Acesso

- **Permissões de Upload**: Membro do projeto/empresa necessário
- **Associação de Arquivo**: Funções ADMIN, OWNER, MEMBER, CLIENT
- **Acesso ao Arquivo**: Herdado das permissões do projeto/empresa
- **URLs Seguras**: URLs assinadas com tempo limitado para acesso ao arquivo

## Permissões Necessárias

| Ação | Permissão Necessária |
|------|---------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## Respostas de Erro

### Arquivo Muito Grande
```json
{
  "errors": [{
    "message": "File \"filename.pdf\": Size exceeds maximum limit of 256MB",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Arquivo Não Encontrado
```json
{
  "errors": [{
    "message": "File not found",
    "extensions": {
      "code": "FILE_NOT_FOUND"
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

## Melhores Práticas

### Gerenciamento de Arquivos
- Faça upload de arquivos antes de anexá-los a registros
- Use nomes de arquivos descritivos
- Organize arquivos por projeto/propósito
- Limpe arquivos não utilizados periodicamente

### Desempenho
- Faça upload de arquivos em lotes quando possível
- Use formatos de arquivo apropriados para o tipo de conteúdo
- Comprimir arquivos grandes antes do upload
- Considere os requisitos de visualização de arquivos

### Segurança
- Valide o conteúdo do arquivo, não apenas as extensões
- Use verificação de vírus para arquivos enviados
- Implemente controles de acesso adequados
- Monitore padrões de upload de arquivos

## Casos de Uso Comuns

1. **Gerenciamento de Documentos**
   - Especificações do projeto
   - Contratos e acordos
   - Notas de reuniões e apresentações
   - Documentação técnica

2. **Gerenciamento de Ativos**
   - Arquivos de design e mockups
   - Ativos de marca e logotipos
   - Materiais de marketing
   - Imagens de produtos

3. **Conformidade e Registros**
   - Documentos legais
   - Trilhas de auditoria
   - Certificados e licenças
   - Registros financeiros

4. **Colaboração**
   - Recursos compartilhados
   - Documentos com controle de versão
   - Feedback e anotações
   - Materiais de referência

## Recursos de Integração

### Com Automação
- Acionar ações quando arquivos são adicionados/removidos
- Processar arquivos com base em tipo ou metadados
- Enviar notificações para alterações de arquivos
- Arquivar arquivos com base em condições

### Com Imagens de Capa
- Usar campos de arquivo como fontes de imagem de capa
- Processamento automático de imagens e miniaturas
- Atualizações dinâmicas de capa quando arquivos mudam

### Com Pesquisas
- Referenciar arquivos de outros registros
- Agregar contagens e tamanhos de arquivos
- Encontrar registros por metadados de arquivo
- Referenciar anexos de arquivos

## Limitações

- Máximo de 256MB por arquivo
- Dependente da disponibilidade do S3
- Sem versionamento de arquivo embutido
- Sem conversão automática de arquivos
- Capacidades limitadas de visualização de arquivos
- Sem edição colaborativa em tempo real

## Recursos Relacionados

- [API de Upload de Arquivos](/api/upload-files) - Endpoints de upload de arquivos
- [Visão Geral de Campos Personalizados](/api/custom-fields/list-custom-fields) - Conceitos gerais
- [API de Automação](/api/automations) - Automação baseada em arquivos
- [Documentação do AWS S3](https://docs.aws.amazon.com/s3/) - Backend de armazenamento