---
title: 货币转换自定义字段
description: 创建自动使用实时汇率转换货币值的字段
---

货币转换自定义字段自动将源货币字段中的值转换为不同的目标货币，使用实时汇率。这些字段会在源货币值变化时自动更新。

转换汇率由 [Frankfurter API](https://github.com/hakanensari/frankfurter) 提供，这是一个开源服务，跟踪由 [欧洲中央银行](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html) 发布的参考汇率。这确保了您国际业务需求的准确、可靠和最新的货币转换。

## 基本示例

创建一个简单的货币转换字段：

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## 高级示例

创建一个具有特定日期的历史汇率转换字段：

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## 完整设置过程

设置货币转换字段需要三个步骤：

### 步骤 1：创建源货币字段

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### 步骤 2：创建货币转换字段

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### 步骤 3：创建转换选项

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 转换字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | 否 | 要转换的源货币字段的 ID |
| `conversionDateType` | String | 否 | 汇率的日期策略（见下文） |
| `conversionDate` | String | 否 | 转换的日期字符串（基于 conversionDateType） |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：自定义字段会根据用户当前的项目上下文自动与项目关联。无需 `projectId` 参数。

### 转换日期类型

| 类型 | 描述 | conversionDate 参数 |
|------|-------------|-------------------------|
| `currentDate` | 使用实时汇率 | 不需要 |
| `specificDate` | 使用固定日期的汇率 | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | 使用来自其他字段的日期 | "todoDueDate" or DATE field ID |

## 创建转换选项

转换选项定义可以转换哪些货币对：

### CreateCustomFieldOptionInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ 是 | 货币转换字段的 ID |
| `title` | String! | ✅ 是 | 此转换选项的显示名称 |
| `currencyConversionFrom` | String! | ✅ 是 | 源货币代码或“任何” |
| `currencyConversionTo` | String! | ✅ 是 | 目标货币代码 |

### 使用“任何”作为源

特殊值“任何”作为 `currencyConversionFrom` 创建一个后备选项：

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

当没有找到特定货币对匹配时，将使用此选项。

## 自动转换的工作原理

1. **值更新**：当在源货币字段中设置值时
2. **选项匹配**：系统根据源货币找到匹配的转换选项
3. **汇率获取**：从 Frankfurter API 检索汇率
4. **计算**：将源金额乘以汇率
5. **存储**：使用目标货币代码保存转换后的值

### 示例流程

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## 基于日期的转换

### 使用当前日期

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

每次源值变化时，转换会使用当前汇率进行更新。

### 使用特定日期

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

始终使用指定日期的汇率。

### 使用字段中的日期

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

使用来自其他字段的日期（无论是待办事项到期日还是日期自定义字段）。

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 转换字段定义 |
| `number` | Float | 转换后的金额 |
| `currency` | String | 目标货币代码 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后更新的时间 |

## 汇率来源

Blue 使用 **Frankfurter API** 进行汇率查询：
- 由欧洲中央银行托管的开源 API
- 每天更新官方汇率
- 支持追溯到 1999 年的历史汇率
- 免费且可靠，适用于商业用途

## 错误处理

### 转换失败

当转换失败（API 错误、无效货币等）时：
- 转换后的值设置为 `0`
- 目标货币仍然被存储
- 不会向用户抛出错误

### 常见场景

| 场景 | 结果 |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| 没有匹配选项 | Uses "Any" option if available |
| Missing source value | 未执行转换 |

## 所需权限

自定义字段管理需要项目级访问权限：

| 角色 | 可以创建/更新字段 |
|------|-------------------------|
| `OWNER` | ✅ 是 |
| `ADMIN` | ✅ 是 |
| `MEMBER` | ❌ 否 |
| `CLIENT` | ❌ 否 |

转换值的查看权限遵循标准记录访问规则。

## 最佳实践

### 选项配置
- 为常见转换创建特定货币对
- 添加“任何”后备选项以提高灵活性
- 为选项使用描述性标题

### 日期策略选择
- 使用 `currentDate` 进行实时财务跟踪
- 使用 `specificDate` 进行历史报告
- 使用 `fromDateField` 进行交易特定汇率

### 性能考虑
- 多个转换字段并行更新
- 仅在源值变化时进行 API 调用
- 相同货币的转换跳过 API 调用

## 常见用例

1. **多货币项目**
   - 以当地货币跟踪项目成本
   - 以公司货币报告总预算
   - 比较各地区的值

2. **国际销售**
   - 将交易值转换为报告货币
   - 以多种货币跟踪收入
   - 对已关闭交易进行历史转换

3. **财务报告**
   - 期末货币转换
   - 合并财务报表
   - 以当地货币进行预算与实际对比

4. **合同管理**
   - 在签署日期转换合同值
   - 以多种货币跟踪付款计划
   - 货币风险评估

## 限制

- 不支持加密货币转换
- 不能手动设置转换值（始终计算得出）
- 所有转换金额固定为 2 位小数精度
- 不支持自定义汇率
- 不缓存汇率（每次转换都需新 API 调用）
- 依赖于 Frankfurter API 的可用性

## 相关资源

- [货币字段](/api/custom-fields/currency) - 转换的源字段
- [日期字段](/api/custom-fields/date) - 用于基于日期的转换
- [公式字段](/api/custom-fields/formula) - 替代计算
- [自定义字段概述](/custom-fields/list-custom-fields) - 一般概念