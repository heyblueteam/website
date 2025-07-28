---
title: 国家自定义字段
description: 创建带有ISO国家代码验证的国家选择字段
---

国家自定义字段允许您存储和管理记录的国家信息。该字段支持国家名称和ISO Alpha-2国家代码。

**重要**：国家验证和转换行为在不同的变更中有显著差异：
- **createTodo**：自动验证并将国家名称转换为ISO代码
- **setTodoCustomField**：接受任何值而不进行验证

## 基本示例

创建一个简单的国家字段：

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带有描述的国家字段：

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 国家字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `COUNTRY` |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：`projectId`未在输入中传递，而是由GraphQL上下文确定（通常来自请求头或身份验证）。

## 设置国家值

国家字段在两个数据库字段中存储数据：
- **`countryCodes`**：在数据库中以逗号分隔的字符串形式存储ISO Alpha-2国家代码（通过API返回为数组）
- **`text`**：以字符串形式存储显示文本或国家名称

### 理解参数

`setTodoCustomField`变更接受两个可选参数用于国家字段：

| 参数 | 类型 | 必需 | 描述 | 功能 |
|-----------|------|----------|-------------|--------------|
| `todoId` | String! | ✅ 是 | 要更新的记录的ID | - |
| `customFieldId` | String! | ✅ 是 | 国家自定义字段的ID | - |
| `countryCodes` | [String!] | 否 | ISO Alpha-2国家代码的数组 | Stored in the `countryCodes` field |
| `text` | String | 否 | 显示文本或国家名称 | Stored in the `text` field |

**重要**：
- 在 `setTodoCustomField`：两个参数都是可选的，并且独立存储
- 在 `createTodo`：系统会根据您的输入自动设置这两个字段（您无法独立控制它们）

### 选项1：仅使用国家代码

存储经过验证的ISO代码而不显示文本：

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

结果： `countryCodes` = `["US"]`, `text` = `null`

### 选项2：仅使用文本

存储显示文本而不经过验证的代码：

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

结果： `countryCodes` = `null`, `text` = `"United States"`

**注意**：使用 `setTodoCustomField` 时，无论您使用哪个参数，都不会进行验证。值将按提供的方式存储。

### 选项3：同时使用（推荐）

存储经过验证的代码和显示文本：

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

结果： `countryCodes` = `["US"]`, `text` = `"United States"`

### 多个国家

使用数组存储多个国家：

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## 使用国家值创建记录

在创建记录时，`createTodo`变更**自动验证并转换**国家值。这是唯一执行国家验证的变更：

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      countryCodes
    }
  }
}
```

### 接受的输入格式

| 输入类型 | 示例 | 结果 |
|------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **不支持** - 被视为单个无效值 |
| Mixed format | `"United States, CA"` | **不支持** - 被视为单个无效值 |

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `text` | String | 显示文本（国家名称） |
| `countryCodes` | [String!] | ISO Alpha-2国家代码的数组 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

## 国家标准

Blue使用**ISO 3166-1 Alpha-2**标准来表示国家代码：

- 两字母国家代码（例如，US，GB，FR，DE）
- 使用 `i18n-iso-countries` 库进行验证**仅在createTodo中发生**
- 支持所有正式认可的国家

### 示例国家代码

| 国家 | ISO代码 |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

要查看完整的ISO 3166-1 Alpha-2国家代码官方列表，请访问[ISO在线浏览平台](https://www.iso.org/obp/ui/#search/code/)。

## 验证

**验证仅在 `createTodo` 变更中发生**：

1. **有效的ISO代码**：接受任何有效的ISO Alpha-2代码
2. **国家名称**：自动将已识别的国家名称转换为代码
3. **无效输入**：对于无法识别的值抛出 `CustomFieldValueParseError`

**注意**：`setTodoCustomField` 变更不执行任何验证，并接受任何字符串值。

### 错误示例

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 集成功能

### 查找字段
国家字段可以通过查找自定义字段进行引用，允许您从相关记录中提取国家数据。

### 自动化
在自动化条件中使用国家值：
- 按特定国家过滤操作
- 根据国家发送通知
- 根据地理区域路由任务

### 表单
表单中的国家字段会自动验证用户输入并将国家名称转换为代码。

## 所需权限

| 操作 | 所需权限 |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## 错误响应

### 无效的国家值
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 字段类型不匹配
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## 最佳实践

### 输入处理
- 使用 `createTodo` 进行自动验证和转换
- 小心使用 `setTodoCustomField`，因为它绕过验证
- 考虑在您的应用程序中验证输入，然后再使用 `setTodoCustomField`
- 在用户界面中显示完整的国家名称以提高清晰度

### 数据质量
- 在输入点验证国家输入
- 在系统中使用一致的格式
- 考虑进行区域分组以便于报告

### 多个国家
- 在 `setTodoCustomField` 中使用数组支持多个国家
- 在 `createTodo` 中的多个国家**不支持**通过值字段
- 在 `setTodoCustomField` 中将国家代码存储为数组以便于正确处理

## 常见用例

1. **客户管理**
   - 客户总部位置
   - 运输目的地
   - 税务管辖区

2. **项目跟踪**
   - 项目位置
   - 团队成员位置
   - 市场目标

3. **合规与法律**
   - 监管管辖区
   - 数据驻留要求
   - 出口控制

4. **销售与市场**
   - 领土分配
   - 市场细分
   - 活动目标

## 限制

- 仅支持ISO 3166-1 Alpha-2代码（两字母代码）
- 不支持国家细分（州/省）
- 不支持自动国家旗帜图标（仅基于文本）
- 无法验证历史国家代码
- 不支持内置区域或大陆分组
- **验证仅在 `createTodo` 中有效，而在 `setTodoCustomField` 中无效**
- **在 `createTodo` 值字段中不支持多个国家**
- **国家代码存储为逗号分隔的字符串，而不是真正的数组**

## 相关资源

- [自定义字段概述](/custom-fields/list-custom-fields) - 一般自定义字段概念
- [查找字段](/api/custom-fields/lookup) - 从其他记录引用国家数据
- [表单API](/api/forms) - 在自定义表单中包含国家字段