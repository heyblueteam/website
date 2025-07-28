---
title: 查找自定义字段
description: 创建自动从引用记录中提取数据的查找字段
---

查找自定义字段自动从通过[引用字段](/api/custom-fields/reference)引用的记录中提取数据，显示来自链接记录的信息，无需手动复制。当引用的数据发生变化时，它们会自动更新。

## 基本示例

创建一个查找字段以显示来自引用记录的标签：

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 高级示例

创建一个查找字段以从引用记录中提取自定义字段值：

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 查找字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ 是 | 查找配置 |
| `description` | String | 否 | 显示给用户的帮助文本 |

## 查找配置

### CustomFieldLookupOptionInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `referenceId` | String! | ✅ 是 | 用于提取数据的引用字段的 ID |
| `lookupId` | String | 否 | 要查找的特定自定义字段的 ID（对于 TODO_CUSTOM_FIELD 类型是必需的） |
| `lookupType` | CustomFieldLookupType! | ✅ 是 | 从引用记录中提取的数据类型 |

## 查找类型

### CustomFieldLookupType 值

| 类型 | 描述 | 返回 |
|------|------|------|
| `TODO_DUE_DATE` | 来自引用待办事项的到期日期 | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | 来自引用待办事项的创建日期 | Array of creation timestamps |
| `TODO_UPDATED_AT` | 来自引用待办事项的最后更新日期 | Array of update timestamps |
| `TODO_TAG` | 来自引用待办事项的标签 | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | 来自引用待办事项的指派人 | Array of user objects |
| `TODO_DESCRIPTION` | 来自引用待办事项的描述 | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | 来自引用待办事项的待办事项列表名称 | Array of list titles |
| `TODO_CUSTOM_FIELD` | 来自引用待办事项的自定义字段值 | Array of values based on the field type |

## 响应字段

### CustomField 响应（用于查找字段）

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段的唯一标识符 |
| `name` | String! | 查找字段的显示名称 |
| `type` | CustomFieldType! | 将是 `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | 查找配置和结果 |
| `createdAt` | DateTime! | 字段创建的时间 |
| `updatedAt` | DateTime! | 字段最后更新的时间 |

### CustomFieldLookupOption 结构

| 字段 | 类型 | 描述 |
|------|------|------|
| `lookupType` | CustomFieldLookupType! | 正在执行的查找类型 |
| `lookupResult` | JSON | 从引用记录中提取的数据 |
| `reference` | CustomField | 用作源的引用字段 |
| `lookup` | CustomField | 正在查找的特定字段（对于 TODO_CUSTOM_FIELD） |
| `parentCustomField` | CustomField | 父查找字段 |
| `parentLookup` | CustomField | 链中的父查找（用于嵌套查找） |

## 查找工作原理

1. **数据提取**：查找从通过引用字段链接的所有记录中提取特定数据
2. **自动更新**：当引用记录发生变化时，查找值会自动更新
3. **只读**：查找字段不能直接编辑 - 它们始终反映当前引用的数据
4. **无计算**：查找直接提取和显示数据，不进行聚合或计算

## TODO_CUSTOM_FIELD 查找

使用 `TODO_CUSTOM_FIELD` 类型时，必须使用 `lookupId` 参数指定要提取的自定义字段：

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

这将从所有引用记录中提取指定自定义字段的值。

## 查询查找数据

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## 示例查找结果

### 标签查找结果
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### 指派人查找结果
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### 自定义字段查找结果
结果根据正在查找的自定义字段类型而异。例如，货币字段查找可能返回：
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**重要**：用户必须对当前项目和引用项目具有查看权限才能查看查找结果。

## 错误响应

### 无效的引用字段
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

### 检测到循环查找
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### TODO_CUSTOM_FIELD 缺少查找 ID
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## 最佳实践

1. **清晰命名**：使用描述性名称，指示正在查找的数据
2. **适当类型**：选择与您的数据需求匹配的查找类型
3. **性能**：查找处理所有引用记录，因此要注意具有多个链接的引用字段
4. **权限**：确保用户对引用项目有访问权限，以便查找正常工作

## 常见用例

### 跨项目可见性
显示来自相关项目的标签、指派人或状态，而无需手动同步。

### 依赖跟踪
显示当前工作所依赖的任务的到期日期或完成状态。

### 资源概览
显示分配给引用任务的所有团队成员，以便进行资源规划。

### 状态汇总
收集来自相关任务的所有唯一状态，以便一目了然地查看项目健康状况。

## 限制

- 查找字段是只读的，不能直接编辑
- 无聚合函数（SUM、COUNT、AVG） - 查找仅提取数据
- 无过滤选项 - 所有引用记录均被包含
- 防止循环查找链以避免无限循环
- 结果反映当前数据并自动更新

## 相关资源

- [引用字段](/api/custom-fields/reference) - 创建指向记录的链接以作为查找源
- [自定义字段值](/api/custom-fields/custom-field-values) - 设置可编辑自定义字段的值
- [列出自定义字段](/api/custom-fields/list-custom-fields) - 查询项目中的所有自定义字段