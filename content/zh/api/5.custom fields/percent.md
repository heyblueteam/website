---
title: 百分比自定义字段
description: 创建百分比字段以存储数值，并自动处理 % 符号和显示格式
---

百分比自定义字段允许您为记录存储百分比值。它们自动处理输入和显示的 % 符号，同时在内部存储原始数值。非常适合完成率、成功率或任何基于百分比的指标。

## 基本示例

创建一个简单的百分比字段：

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带描述的百分比字段：

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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
| `name` | String! | ✅ 是 | 百分比字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `PERCENT` |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：项目上下文会根据您的身份验证头自动确定。无需 `projectId` 参数。

**注意**：PERCENT 字段不支持最小/最大约束或像 NUMBER 字段那样的前缀格式。

## 设置百分比值

百分比字段存储数值，并自动处理 % 符号：

### 带百分比符号

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### 直接数值

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 百分比自定义字段的 ID |
| `number` | Float | 否 | 数值百分比（例如，75.5 表示 75.5%） |

## 值存储和显示

### 存储格式
- **内部存储**：原始数值（例如，75.5）
- **数据库**：存储为 `Decimal` 在 `number` 列中
- **GraphQL**：返回为 `Float` 类型

### 显示格式
- **用户界面**：客户端应用程序必须附加 % 符号（例如，“75.5%”）
- **图表**：当输出类型为 PERCENTAGE 时显示 % 符号
- **API 响应**：原始数值不带 % 符号（例如，75.5）

## 使用百分比值创建记录

创建新记录时使用百分比值：

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### 支持的输入格式

| 格式 | 示例 | 结果 |
|------|------|------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**注意**：输入中的 % 符号会自动去除，并在显示时重新添加。

## 查询百分比值

查询具有百分比自定义字段的记录时，通过 `customField.value.number` 路径访问值：

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

响应将包括作为原始数字的百分比：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | ID! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义（包含百分比值） |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

**重要**：百分比值通过 `customField.value.number` 字段访问。存储的值不包含 % 符号，客户端应用程序必须在显示时添加。

## 过滤和查询

百分比字段支持与 NUMBER 字段相同的过滤：

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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

### 支持的运算符

| 运算符 | 描述 | 示例 |
|--------|------|------|
| `EQ` | 等于 | `percentage = 75` |
| `NE` | 不等于 | `percentage ≠ 75` |
| `GT` | 大于 | `percentage > 75` |
| `GTE` | 大于或等于 | `percentage ≥ 75` |
| `LT` | 小于 | `percentage < 75` |
| `LTE` | 小于或等于 | `percentage ≤ 75` |
| `IN` | 列表中的值 | `percentage in [50, 75, 100]` |
| `NIN` | 不在列表中的值 | `percentage not in [0, 25]` |
| `IS` | 使用 `values: null` 检查 null | `percentage is null` |
| `NOT` | 使用 `values: null` 检查非 null | `percentage is not null` |

### 范围过滤

对于范围过滤，使用多个运算符：

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## 百分比值范围

### 常见范围

| 范围 | 描述 | 用例 |
|------|------|------|
| `0-100` | 标准百分比 | Completion rates, success rates |
| `0-∞` | 无限百分比 | Growth rates, performance metrics |
| `-∞-∞` | 任何值 | Change rates, variance |

### 示例值

| 输入 | 存储 | 显示 |
|------|------|------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## 图表聚合

百分比字段支持在仪表板图表和报告中的聚合。可用的函数包括：

- `AVERAGE` - 平均百分比值
- `COUNT` - 具有值的记录数量
- `MIN` - 最低百分比值
- `MAX` - 最高百分比值 
- `SUM` - 所有百分比值的总和

这些聚合在创建图表和仪表板时可用，而不是在直接的 GraphQL 查询中。

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## 错误响应

### 无效的百分比格式
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

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

### 值输入
- 允许用户输入带或不带 % 符号
- 验证合理的范围以适应您的用例
- 提供关于 100% 代表什么的清晰上下文

### 显示
- 始终在用户界面中显示 % 符号
- 使用适当的小数精度
- 考虑使用颜色编码来表示范围（红色/黄色/绿色）

### 数据解释
- 记录在您的上下文中 100% 的含义
- 适当处理超过 100% 的值
- 考虑负值是否有效

## 常见用例

1. **项目管理**
   - 任务完成率
   - 项目进度
   - 资源利用率
   - 决策速度

2. **绩效跟踪**
   - 成功率
   - 错误率
   - 效率指标
   - 质量评分

3. **财务指标**
   - 增长率
   - 利润率
   - 折扣金额
   - 变化百分比

4. **分析**
   - 转化率
   - 点击率
   - 参与度指标
   - 绩效指标

## 集成功能

### 使用公式
- 在计算中引用 PERCENT 字段
- 公式输出中自动格式化 % 符号
- 与其他数值字段结合

### 使用自动化
- 根据百分比阈值触发操作
- 针对里程碑百分比发送通知
- 根据完成率更新状态

### 使用查找
- 从相关记录中聚合百分比
- 计算平均成功率
- 查找表现最高/最低的项目

### 使用图表
- 创建基于百分比的可视化
- 跟踪进度
- 比较绩效指标

## 与 NUMBER 字段的区别

### 有什么不同
- **输入处理**：自动去除 % 符号
- **显示**：自动添加 % 符号
- **约束**：没有最小/最大验证
- **格式**：不支持前缀

### 有什么相同
- **存储**：相同的数据库列和类型
- **过滤**：相同的查询运算符
- **聚合**：相同的聚合函数
- **权限**：相同的权限模型

## 限制

- 无最小/最大值约束
- 无前缀格式选项
- 不会自动验证 0-100% 范围
- 不支持百分比格式之间的转换（例如，0.75 ↔ 75%）
- 允许超过 100% 的值

## 相关资源

- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般自定义字段概念
- [数字自定义字段](/api/custom-fields/number) - 用于原始数值
- [自动化 API](/api/automations/index) - 创建基于百分比的自动化