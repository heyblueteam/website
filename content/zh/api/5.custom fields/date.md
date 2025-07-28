---
title: 日期自定义字段
description: 创建日期字段以跟踪单个日期或带时区支持的日期范围
---

日期自定义字段允许您为记录存储单个日期或日期范围。它们支持时区处理、智能格式化，并可用于跟踪截止日期、事件日期或任何基于时间的信息。

## 基本示例

创建一个简单的日期字段：

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带描述的截止日期字段：

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 日期字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `DATE` |
| `isDueDate` | Boolean | 否 | 此字段是否表示截止日期 |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：自定义字段会根据用户当前的项目上下文自动与项目关联。无需 `projectId` 参数。

## 设置日期值

日期字段可以存储单个日期或日期范围：

### 单个日期

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### 日期范围

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### 全天事件

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 日期自定义字段的 ID |
| `startDate` | DateTime | 否 | ISO 8601 格式的开始日期/时间 |
| `endDate` | DateTime | 否 | ISO 8601 格式的结束日期/时间 |
| `timezone` | String | 否 | 时区标识符（例如，“America/New_York”） |

**注意**：如果仅提供 `startDate`，`endDate` 将自动默认为相同值。

## 日期格式

### ISO 8601 格式
所有日期必须以 ISO 8601 格式提供：
- `2025-01-15T14:30:00Z` - UTC 时间
- `2025-01-15T14:30:00+05:00` - 带时区偏移
- `2025-01-15T14:30:00.123Z` - 带毫秒

### 时区标识符
使用标准时区标识符：
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

如果未提供时区，系统将默认为用户检测到的时区。

## 使用日期值创建记录

创建新记录时使用日期值：

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### 支持的输入格式

创建记录时，可以以各种格式提供日期：

| 格式 | 示例 | 结果 |
|------|------|------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | ID! | 字段值的唯一标识符 |
| `uid` | String! | 唯一标识符字符串 |
| `customField` | CustomField! | 自定义字段定义（包含日期值） |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

**重要**：日期值（`startDate`，`endDate`，`timezone`）通过 `customField.value` 字段访问，而不是直接在 TodoCustomField 上访问。

### 值对象结构

日期值通过 `customField.value` 字段作为 JSON 对象返回：

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**注意**：`value` 字段位于 `CustomField` 类型上，而不是 `TodoCustomField` 上。

## 查询日期值

查询带有日期自定义字段的记录时，通过 `customField.value` 字段访问日期值：

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

响应将包括 `value` 字段中的日期值：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## 日期显示智能

系统会根据范围自动格式化日期：

| 场景 | 显示格式 |
|------|----------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025`（不显示时间） |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**全天检测**：从 00:00 到 23:59 的事件会自动被检测为全天事件。

## 时区处理

### 存储
- 所有日期都以 UTC 存储在数据库中
- 时区信息单独保留
- 转换在显示时发生

### 最佳实践
- 始终提供时区以确保准确性
- 在项目中使用一致的时区
- 考虑全球团队的用户位置

### 常见时区

| 区域 | 时区 ID | UTC 偏移 |
|------|---------|----------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## 过滤和查询

日期字段支持复杂过滤：

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### 检查空日期字段

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### 支持的操作符

| 操作符 | 用法 | 描述 |
|--------|------|------|
| `EQ` | 与 dateRange 一起使用 | 日期与指定范围重叠（任何交集） |
| `NE` | 与 dateRange 一起使用 | 日期不与范围重叠 |
| `IS` | 与 `values: null` 一起使用 | 日期字段为空（startDate 或 endDate 为 null） |
| `NOT` | 与 `values: null` 一起使用 | 日期字段有值（两个日期都不为 null） |

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## 错误响应

### 无效的日期格式
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
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


## 限制

- 不支持重复日期（使用自动化处理重复事件）
- 无法在没有日期的情况下设置时间
- 没有内置的工作日计算
- 日期范围不会自动验证结束 > 开始
- 最大精度为秒（不存储毫秒）

## 相关资源

- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般自定义字段概念
- [自动化 API](/api/automations/index) - 创建基于日期的自动化