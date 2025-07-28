---
title: 数字自定义字段
description: 创建数字字段以存储数值，具有可选的最小/最大约束和前缀格式
---

数字自定义字段允许您为记录存储数值。它们支持验证约束、小数精度，并可用于数量、分数、测量或任何不需要特殊格式的数值数据。

## 基本示例

创建一个简单的数字字段：

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带有约束和前缀的数字字段：

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 数字字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `NUMBER` |
| `projectId` | String! | ✅ 是 | 创建字段的项目 ID |
| `min` | Float | 否 | 最小值约束（仅供 UI 指导） |
| `max` | Float | 否 | 最大值约束（仅供 UI 指导） |
| `prefix` | String | 否 | 显示前缀（例如，“#”，“~”，“$”） |
| `description` | String | 否 | 显示给用户的帮助文本 |

## 设置数字值

数字字段存储小数值，并具有可选的验证：

### 简单数字值

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### 整数值

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录 ID |
| `customFieldId` | String! | ✅ 是 | 数字自定义字段的 ID |
| `number` | Float | 否 | 要存储的数值 |

## 值约束

### 最小/最大约束（UI 指导）

**重要**：最小/最大约束被存储但不在服务器端强制执行。它们作为前端应用程序的 UI 指导。

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**需要客户端验证**：前端应用程序必须实现验证逻辑以强制执行最小/最大约束。

### 支持的值类型

| 类型 | 示例 | 描述 |
|------|------|------|
| Integer | `42` | 整数 |
| Decimal | `42.5` | 带小数位的数字 |
| Negative | `-10` | 负值（如果没有最小约束） |
| Zero | `0` | 零值 |

**注意**：最小/最大约束不在服务器端验证。超出指定范围的值将被接受并存储。

## 创建带有数字值的记录

在创建带有数字值的新记录时：

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
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
        prefix
      }
      number
      value
    }
  }
}
```

### 支持的输入格式

在创建记录时，在自定义字段数组中使用 `number` 参数（而不是 `value`）：

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `number` | Float | 数值 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

### CustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段定义的唯一标识符 |
| `name` | String! | 字段的显示名称 |
| `type` | CustomFieldType! | 始终是 `NUMBER` |
| `min` | Float | 允许的最小值 |
| `max` | Float | 允许的最大值 |
| `prefix` | String | 显示前缀 |
| `description` | String | 帮助文本 |

**注意**：如果未设置数字值，`number` 字段将为 `null`。

## 过滤和查询

数字字段支持全面的数字过滤：

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### 支持的操作符

| 操作符 | 描述 | 示例 |
|--------|------|------|
| `EQ` | 等于 | `number = 42` |
| `NE` | 不等于 | `number ≠ 42` |
| `GT` | 大于 | `number > 42` |
| `GTE` | 大于或等于 | `number ≥ 42` |
| `LT` | 小于 | `number < 42` |
| `LTE` | 小于或等于 | `number ≤ 42` |
| `IN` | 在数组中 | `number in [1, 2, 3]` |
| `NIN` | 不在数组中 | `number not in [1, 2, 3]` |
| `IS` | 是空/非空 | `number is null` |

### 范围过滤

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## 显示格式

### 带前缀

如果设置了前缀，它将被显示：

| 值 | 前缀 | 显示 |
|----|------|------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### 小数精度

数字保持其小数精度：

| 输入 | 存储 | 显示 |
|------|------|------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## 错误响应

### 无效的数字格式
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**注意**：最小/最大验证错误不会在服务器端发生。约束验证必须在您的前端应用程序中实现。

### 不是数字
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 最佳实践

### 约束设计
- 为 UI 指导设置现实的最小/最大值
- 实施客户端验证以强制执行约束
- 使用约束在表单中提供用户反馈
- 考虑负值是否适用于您的用例

### 值精度
- 根据需要使用适当的小数精度
- 考虑出于显示目的进行四舍五入
- 在相关字段之间保持精度一致

### 显示增强
- 使用有意义的前缀提供上下文
- 考虑在字段名称中使用单位（例如，“重量（kg）”）
- 为验证规则提供清晰的描述

## 常见用例

1. **评分系统**
   - 性能评级
   - 质量分数
   - 优先级级别
   - 客户满意度评级

2. **测量**
   - 数量和金额
   - 尺寸和大小
   - 持续时间（以数字格式）
   - 容量和限制

3. **业务指标**
   - 收入数字
   - 转化率
   - 预算分配
   - 目标数字

4. **技术数据**
   - 版本号
   - 配置值
   - 性能指标
   - 阈值设置

## 集成功能

### 与图表和仪表板
- 在图表计算中使用数字字段
- 创建数值可视化
- 跟踪趋势变化

### 与自动化
- 根据数字阈值触发操作
- 根据数字变化更新相关字段
- 对特定值发送通知

### 与查找
- 从相关记录中汇总数字
- 计算总数和平均值
- 查找关系中的最小/最大值

### 与图表
- 创建数值可视化
- 跟踪趋势变化
- 比较记录之间的值

## 限制

- **不进行服务器端验证** 的最小/最大约束
- **需要客户端验证** 以强制执行约束
- 无内置货币格式（请使用货币类型）
- 无自动百分号（请使用百分比类型）
- 无单位转换功能
- 小数精度受数据库小数类型限制
- 字段本身不进行数学公式评估

## 相关资源

- [自定义字段概述](/api/custom-fields/1.index) - 一般自定义字段概念
- [货币自定义字段](/api/custom-fields/currency) - 用于货币值
- [百分比自定义字段](/api/custom-fields/percent) - 用于百分比值
- [自动化 API](/api/automations/1.index) - 创建基于数字的自动化