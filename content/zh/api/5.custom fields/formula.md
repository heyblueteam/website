---
title: 公式自定义字段
description: 创建基于其他数据自动计算值的计算字段
---

公式自定义字段用于 Blue 中的图表和仪表板计算。它们定义在自定义字段数据上操作的聚合函数（SUM、AVERAGE、COUNT 等），以在图表中显示计算的指标。公式不是在单个待办事项级别计算的，而是跨多个记录聚合数据以便于可视化。

## 基本示例

为图表计算创建一个公式字段：

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## 高级示例

创建一个具有复杂计算的货币公式：

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 公式字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `FORMULA` |
| `projectId` | String! | ✅ 是 | 此字段将被创建的项目 ID |
| `formula` | JSON | 否 | 图表计算的公式定义 |
| `description` | String | 否 | 显示给用户的帮助文本 |

### 公式结构

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## 支持的函数

### 图表聚合函数

公式字段支持以下聚合函数用于图表计算：

| 函数 | 描述 | ChartFunction 枚举 |
|------|------|-------------------|
| `SUM` | 所有值的总和 | `SUM` |
| `AVERAGE` | 数值的平均值 | `AVERAGE` |
| `AVERAGEA` | 排除零和空值的平均值 | `AVERAGEA` |
| `COUNT` | 值的计数 | `COUNT` |
| `COUNTA` | 排除零和空值的计数 | `COUNTA` |
| `MAX` | 最大值 | `MAX` |
| `MIN` | 最小值 | `MIN` |

**注意**：这些函数用于 `display.function` 字段，并在图表可视化中对聚合数据进行操作。复杂的数学表达式或字段级计算不受支持。

## 显示类型

### 数字显示

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

结果： `1250.75`

### 货币显示

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

结果： `$1,250.75`

### 百分比显示

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

结果： `87.5%`

## 编辑公式字段

更新现有的公式字段：

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## 公式处理

### 图表计算上下文

公式字段在图表段和仪表板的上下文中处理：
- 计算在图表渲染或更新时发生
- 结果以十进制值存储在 `ChartSegment.formulaResult` 中
- 处理通过名为 'formula' 的专用 BullMQ 队列进行
- 更新发布到仪表板订阅者以实现实时更新

### 显示格式

`getFormulaDisplayValue` 函数根据显示类型格式化计算结果：
- **NUMBER**：以普通数字显示，带可选精度
- **PERCENTAGE**：添加 % 后缀，带可选精度  
- **CURRENCY**：使用指定的货币代码格式化

## 公式结果存储

结果存储在 `formulaResult` 字段中：

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 公式字段定义 |
| `number` | Float | 计算的数值结果 |
| `formulaResult` | JSON | 带显示格式的完整结果 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后计算的时间 |

## 数据上下文

### 图表数据源

公式字段在图表数据源上下文中操作：
- 公式在项目中的待办事项之间聚合自定义字段值
- `display.function` 中指定的聚合函数决定计算
- 使用 SQL 聚合函数（avg、sum、count 等）计算结果
- 计算在数据库级别执行以提高效率

## 常见公式示例

### 总预算（图表显示）

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### 平均分数（图表显示）

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### 任务计数（图表显示）

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## 所需权限

自定义字段操作遵循标准基于角色的权限：

| 操作 | 所需角色 |
|------|----------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**注意**：所需的具体角色取决于您项目的自定义角色配置。没有像 CUSTOM_FIELDS_CREATE 这样的特殊权限常量。

## 错误处理

### 验证错误
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### 自定义字段未找到
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

## 最佳实践

### 公式设计
- 使用清晰、描述性的名称为公式字段命名
- 添加描述以解释计算逻辑
- 在部署之前使用示例数据测试公式
- 保持公式简单易读

### 性能优化
- 避免深层嵌套的公式依赖
- 使用特定字段引用而不是通配符
- 考虑复杂计算的缓存策略
- 在大型项目中监控公式性能

### 数据质量
- 在使用公式之前验证源数据
- 适当地处理空值或 null 值
- 为显示类型使用适当的精度
- 考虑计算中的边缘情况

## 常见用例

1. **财务跟踪**
   - 预算计算
   - 利润/损失报表
   - 成本分析
   - 收入预测

2. **项目管理**
   - 完成百分比
   - 资源利用率
   - 时间线计算
   - 绩效指标

3. **质量控制**
   - 平均分数
   - 通过/未通过率
   - 质量指标
   - 合规跟踪

4. **商业智能**
   - KPI 计算
   - 趋势分析
   - 比较指标
   - 仪表板值

## 限制

- 公式仅用于图表/仪表板聚合，而不是待办事项级别的计算
- 限于七个支持的聚合函数（SUM、AVERAGE 等）
- 不支持复杂的数学表达式或字段到字段的计算
- 不能在单个公式中引用多个字段
- 结果仅在图表和仪表板中可见
- `logic` 字段仅用于显示文本，而不是实际的计算逻辑

## 相关资源

- [数字字段](/api/5.custom%20fields/number) - 用于静态数值
- [货币字段](/api/5.custom%20fields/currency) - 用于货币值
- [引用字段](/api/5.custom%20fields/reference) - 用于跨项目数据
- [查找字段](/api/5.custom%20fields/lookup) - 用于聚合数据
- [自定义字段概述](/api/5.custom%20fields/2.list-custom-fields) - 一般概念