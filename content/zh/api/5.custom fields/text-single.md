---
title: 单行文本自定义字段
description: 创建用于短文本值（如名称、标题和标签）的单行文本字段
---

单行文本自定义字段允许您存储用于单行输入的短文本值。它们非常适合名称、标题、标签或任何应显示在单行上的文本数据。

## 基本示例

创建一个简单的单行文本字段：

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带描述的单行文本字段：

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
|------|------|------|------|
| `name` | String! | ✅ 是 | 文本字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `TEXT_SINGLE` |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：项目上下文是根据您的身份验证头自动确定的。无需 `projectId` 参数。

## 设置文本值

要在记录上设置或更新单行文本值：

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 文本自定义字段的 ID |
| `text` | String | 否 | 要存储的单行文本内容 |

## 使用文本值创建记录

创建带有单行文本值的新记录时：

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | ID! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义（包含文本值） |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

**重要**：文本值通过 `customField.value.text` 字段访问，而不是直接在 TodoCustomField 上。

## 查询文本值

查询具有文本自定义字段的记录时，通过 `customField.value.text` 路径访问文本：

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

响应将包含嵌套结构中的文本：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## 文本验证

### 表单验证
当单行文本字段用于表单时：
- 自动修剪前导和尾随空格
- 如果字段标记为必需，则应用必需验证
- 不应用特定格式验证

### 验证规则
- 接受任何字符串内容，包括换行符（尽管不推荐）
- 没有字符长度限制（受数据库限制）
- 支持 Unicode 字符和特殊符号
- 保留换行符，但不适用于此字段类型

### 典型文本示例
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## 重要注意事项

### 存储容量
- 使用 MySQL `MediumText` 类型存储
- 支持最多 16MB 的文本内容
- 与多行文本字段的存储相同
- UTF-8 编码用于国际字符

### 直接 API 与表单
- **表单**：自动修剪空格和必需验证
- **直接 API**：文本按提供的方式存储
- **建议**：使用表单进行用户输入，以确保一致的格式

### TEXT_SINGLE 与 TEXT_MULTI
- **TEXT_SINGLE**：单行文本输入，适合短值
- **TEXT_MULTI**：多行文本区域输入，适合较长内容
- **后端**：两者使用相同的存储和验证
- **前端**：用于数据输入的不同 UI 组件
- **意图**：TEXT_SINGLE 在语义上适合单行值

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## 错误响应

### 必需字段验证（仅限表单）
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

### 找不到字段
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

## 最佳实践

### 内容指南
- 保持文本简洁且适合单行
- 避免换行以便于单行显示
- 对于类似数据类型使用一致的格式
- 根据您的 UI 要求考虑字符限制

### 数据输入
- 提供清晰的字段描述以指导用户
- 使用表单进行用户输入以确保验证
- 如果需要，在您的应用程序中验证内容格式
- 考虑使用下拉菜单以标准化值

### 性能考虑
- 单行文本字段轻量且性能良好
- 考虑为经常搜索的字段建立索引
- 在您的 UI 中使用适当的显示宽度
- 监控内容长度以便于显示

## 过滤和搜索

### 包含搜索
单行文本字段支持子字符串搜索：

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### 搜索能力
- 不区分大小写的子字符串匹配
- 支持部分单词匹配
- 精确值匹配
- 不支持全文搜索或排名

## 常见用例

1. **标识符和代码**
   - 产品 SKU
   - 订单号
   - 参考代码
   - 版本号

2. **名称和标题**
   - 客户名称
   - 项目标题
   - 产品名称
   - 类别标签

3. **简短描述**
   - 简要摘要
   - 状态标签
   - 优先级指示
   - 分类标签

4. **外部引用**
   - 工单号
   - 发票参考
   - 外部系统 ID
   - 文档编号

## 集成功能

### 与查找
- 从其他记录引用文本数据
- 按文本内容查找记录
- 显示相关文本信息
- 从多个来源聚合文本值

### 与表单
- 自动修剪空格
- 必需字段验证
- 单行文本输入 UI
- 字符限制显示（如果配置）

### 与导入/导出
- 直接 CSV 列映射
- 自动文本值分配
- 批量数据导入支持
- 导出到电子表格格式

## 限制

### 自动化限制
- 不直接作为自动化触发字段可用
- 不能用于自动化字段更新
- 可以在自动化条件中引用
- 可用于电子邮件模板和 Webhook

### 一般限制
- 没有内置文本格式或样式
- 除必需字段外没有自动验证
- 没有内置唯一性强制
- 对于非常大的文本没有内容压缩
- 没有版本控制或变更跟踪
- 限制搜索能力（不支持全文搜索）

## 相关资源

- [多行文本字段](/api/custom-fields/text-multi) - 用于更长的文本内容
- [电子邮件字段](/api/custom-fields/email) - 用于电子邮件地址
- [URL 字段](/api/custom-fields/url) - 用于网站地址
- [唯一 ID 字段](/api/custom-fields/unique-id) - 用于自动生成的标识符
- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般概念