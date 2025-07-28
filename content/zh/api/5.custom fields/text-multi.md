---
title: 多行文本自定义字段
description: 创建用于较长内容（如描述、备注和评论）的多行文本字段
---

多行文本自定义字段允许您存储带有换行和格式的较长文本内容。它们非常适合描述、备注、评论或任何需要多行的文本数据。

## 基本示例

创建一个简单的多行文本字段：

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带描述的多行文本字段：

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `type` | CustomFieldType! | ✅ 是 | 必须是 `TEXT_MULTI` |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意：** `projectId` 作为单独的参数传递给变更，而不是作为输入对象的一部分。或者，可以从您的 GraphQL 请求中的 `X-Bloo-Project-ID` 头中确定项目上下文。

## 设置文本值

要在记录上设置或更新多行文本值：

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 文本自定义字段的 ID |
| `text` | String | 否 | 要存储的多行文本内容 |

## 使用文本值创建记录

创建带有多行文本值的新记录时：

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `text` | String | 存储的多行文本内容 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

## 文本验证

### 表单验证
当多行文本字段在表单中使用时：
- 前导和尾随空格会自动修剪
- 如果字段标记为必需，则应用必需验证
- 不应用特定格式验证

### 验证规则
- 接受任何字符串内容，包括换行符
- 没有字符长度限制（最多到数据库限制）
- 支持 Unicode 字符和特殊符号
- 换行符在存储中被保留

### 有效文本示例
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## 重要说明

### 存储容量
- 使用 MySQL `MediumText` 类型存储
- 支持最多 16MB 的文本内容
- 换行符和格式被保留
- 对于国际字符使用 UTF-8 编码

### 直接 API 与表单
- **表单**：自动修剪空格和必需验证
- **直接 API**：文本按提供的方式存储
- **建议**：使用表单进行用户输入，以确保一致的格式

### TEXT_MULTI 与 TEXT_SINGLE
- **TEXT_MULTI**：多行文本区域输入，适合较长内容
- **TEXT_SINGLE**：单行文本输入，适合短值
- **后端**：两种类型是相同的 - 相同的存储字段、验证和处理
- **前端**：不同的数据输入 UI 组件（文本区域与输入字段）
- **重要**：TEXT_MULTI 和 TEXT_SINGLE 之间的区别纯粹是为了 UI 目的

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

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

### 内容组织
- 对结构化内容使用一致的格式
- 考虑使用类似 markdown 的语法以提高可读性
- 将长内容分成逻辑部分
- 使用换行符以提高可读性

### 数据输入
- 提供清晰的字段描述以指导用户
- 使用表单进行用户输入以确保验证
- 根据您的用例考虑字符限制
- 如有需要，在您的应用程序中验证内容格式

### 性能考虑
- 非常长的文本内容可能会影响查询性能
- 考虑为显示大型文本字段进行分页
- 搜索功能的索引考虑
- 监控大型内容字段的存储使用情况

## 过滤和搜索

### 包含搜索
多行文本字段支持通过自定义字段过滤器进行子字符串搜索：

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### 搜索能力
- 使用 `CONTAINS` 操作符在文本字段中进行子字符串匹配
- 使用 `NCONTAINS` 操作符进行不区分大小写的搜索
- 使用 `IS` 操作符进行精确匹配
- 使用 `NOT` 操作符进行负匹配
- 在所有文本行中进行搜索
- 支持部分单词匹配

## 常见用例

1. **项目管理**
   - 任务描述
   - 项目需求
   - 会议记录
   - 状态更新

2. **客户支持**
   - 问题描述
   - 解决方案备注
   - 客户反馈
   - 通信日志

3. **内容管理**
   - 文章内容
   - 产品描述
   - 用户评论
   - 评论详情

4. **文档**
   - 过程描述
   - 指令
   - 指南
   - 参考材料

## 集成功能

### 与自动化
- 当文本内容更改时触发操作
- 从文本内容中提取关键词
- 创建摘要或通知
- 使用外部服务处理文本内容

### 与查找
- 引用来自其他记录的文本数据
- 从多个来源聚合文本内容
- 按文本内容查找记录
- 显示相关文本信息

### 与表单
- 自动修剪空格
- 必需字段验证
- 多行文本区域 UI
- 字符计数显示（如果配置）

## 限制

- 没有内置文本格式或富文本编辑
- 没有自动链接检测或转换
- 没有拼写检查或语法验证
- 没有内置文本分析或处理
- 没有版本控制或更改跟踪
- 限制搜索能力（没有全文搜索）
- 对于非常大的文本没有内容压缩

## 相关资源

- [单行文本字段](/api/custom-fields/text-single) - 用于短文本值
- [电子邮件字段](/api/custom-fields/email) - 用于电子邮件地址
- [URL 字段](/api/custom-fields/url) - 用于网站地址
- [自定义字段概述](/api/custom-fields/2.list-custom-fields) - 一般概念