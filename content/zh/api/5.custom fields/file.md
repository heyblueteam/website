---
title: 文件自定义字段
description: 创建文件字段以将文档、图像和其他文件附加到记录
---

文件自定义字段允许您将多个文件附加到记录。文件安全地存储在 AWS S3 中，并具有全面的元数据跟踪、文件类型验证和适当的访问控制。

## 基本示例

创建一个简单的文件字段：

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

## 高级示例

创建一个带描述的文件字段：

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

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 文件字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `FILE` |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：自定义字段会根据用户当前的项目上下文自动与项目关联。无需 `projectId` 参数。

## 文件上传过程

### 步骤 1：上传文件

首先，上传文件以获取文件 UID：

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

### 步骤 2：将文件附加到记录

然后将上传的文件附加到记录：

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

## 管理文件附件

### 添加单个文件

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

### 移除文件

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### 批量文件操作

使用 customFieldOptionIds 一次更新多个文件：

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## 文件上传输入参数

### UploadFileInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `file` | Upload! | ✅ 是 | 要上传的文件 |
| `companyId` | String! | ✅ 是 | 文件存储的公司 ID |
| `projectId` | String | 否 | 项目特定文件的项目 ID |

### 文件管理输入参数

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 记录的 ID |
| `customFieldId` | String! | ✅ 是 | 文件自定义字段的 ID |
| `fileUid` | String! | ✅ 是 | 上传文件的唯一标识符 |

## 文件存储和限制

### 文件大小限制

| 限制类型 | 大小 |
|------------|------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### 支持的文件类型

#### 图像
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### 视频
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### 音频
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### 文档
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### 压缩文件
- `zip`, `rar`, `7z`, `tar`, `gz`

#### 代码/文本
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### 存储架构

- **存储**：AWS S3，具有组织的文件夹结构
- **路径格式**： `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **安全性**：签名 URL 以确保安全访问
- **备份**：自动 S3 冗余

## 响应字段

### 文件响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | ID! | 数据库 ID |
| `uid` | String! | 唯一文件标识符 |
| `name` | String! | 原始文件名 |
| `size` | Float! | 文件大小（字节） |
| `type` | String! | MIME 类型 |
| `extension` | String! | 文件扩展名 |
| `status` | FileStatus | PENDING 或 CONFIRMED（可为空） |
| `shared` | Boolean! | 文件是否共享 |
| `createdAt` | DateTime! | 上传时间戳 |

### TodoCustomFieldFile 响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | ID! | 连接记录 ID |
| `uid` | String! | 唯一标识符 |
| `position` | Float! | 显示顺序 |
| `file` | File! | 关联的文件对象 |
| `todoCustomField` | TodoCustomField! | 父自定义字段 |
| `createdAt` | DateTime! | 文件附加的时间 |

## 使用文件创建记录

创建记录时，您可以使用它们的 UID 附加文件：

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

## 文件验证和安全性

### 上传验证

- **MIME 类型检查**：验证允许的类型
- **文件扩展名验证**：用于 `application/octet-stream` 的后备
- **大小限制**：在上传时强制执行
- **文件名清理**：移除特殊字符

### 访问控制

- **上传权限**：需要项目/公司成员资格
- **文件关联**：ADMIN、OWNER、MEMBER、CLIENT 角色
- **文件访问**：从项目/公司权限继承
- **安全 URL**：时间限制的签名 URL 用于文件访问

## 所需权限

| 操作 | 所需权限 |
|--------|-------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## 错误响应

### 文件过大
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

### 文件未找到
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

### 字段未找到
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

## 最佳实践

### 文件管理
- 在附加到记录之前上传文件
- 使用描述性文件名
- 按项目/目的组织文件
- 定期清理未使用的文件

### 性能
- 尽可能批量上传文件
- 使用适当的文件格式以匹配内容类型
- 在上传之前压缩大文件
- 考虑文件预览要求

### 安全性
- 验证文件内容，而不仅仅是扩展名
- 对上传的文件进行病毒扫描
- 实施适当的访问控制
- 监控文件上传模式

## 常见用例

1. **文档管理**
   - 项目规范
   - 合同和协议
   - 会议记录和演示文稿
   - 技术文档

2. **资产管理**
   - 设计文件和模型
   - 品牌资产和徽标
   - 营销材料
   - 产品图像

3. **合规性和记录**
   - 法律文件
   - 审计跟踪
   - 证书和许可证
   - 财务记录

4. **协作**
   - 共享资源
   - 版本控制的文档
   - 反馈和注释
   - 参考材料

## 集成功能

### 与自动化
- 在文件添加/移除时触发操作
- 根据类型或元数据处理文件
- 发送文件更改的通知
- 根据条件归档文件

### 与封面图像
- 使用文件字段作为封面图像源
- 自动图像处理和缩略图
- 当文件更改时动态更新封面

### 与查找
- 从其他记录引用文件
- 聚合文件数量和大小
- 根据文件元数据查找记录
- 交叉引用文件附件

## 限制

- 每个文件最大 256MB
- 依赖于 S3 可用性
- 无内置文件版本控制
- 无自动文件转换
- 有限的文件预览能力
- 无实时协作编辑

## 相关资源

- [上传文件 API](/api/upload-files) - 文件上传端点
- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般概念
- [自动化 API](/api/automations) - 基于文件的自动化
- [AWS S3 文档](https://docs.aws.amazon.com/s3/) - 存储后端