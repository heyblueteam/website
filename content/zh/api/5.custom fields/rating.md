---
title: 评分自定义字段
description: 创建评分字段以存储具有可配置尺度和验证的数字评分
---

评分自定义字段允许您在记录中存储数字评分，并可配置最小值和最大值。它们非常适合用于绩效评分、满意度评分、优先级级别或您项目中的任何基于数字尺度的数据。

## 基本示例

创建一个默认0-5尺度的简单评分字段：

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## 高级示例

创建一个具有自定义尺度和描述的评分字段：

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 评分字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `RATING` |
| `projectId` | String! | ✅ 是 | 此字段将被创建的项目 ID |
| `description` | String | 否 | 显示给用户的帮助文本 |
| `min` | Float | 否 | 最小评分值（无默认值） |
| `max` | Float | 否 | 最大评分值 |

## 设置评分值

要在记录上设置或更新评分值：

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 评分自定义字段的 ID |
| `value` | String! | ✅ 是 | 评分值作为字符串（在配置范围内） |

## 创建具有评分值的记录

创建新记录时包含评分值：

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
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
| `value` | Float | 存储的评分值（通过 customField.value 访问） |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

**注意**：评分值实际上通过 `customField.value.number` 在查询中访问。

### CustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段的唯一标识符 |
| `name` | String! | 评分字段的显示名称 |
| `type` | CustomFieldType! | 始终是 `RATING` |
| `min` | Float | 允许的最小评分值 |
| `max` | Float | 允许的最大评分值 |
| `description` | String | 字段的帮助文本 |

## 评分验证

### 值约束
- 评分值必须是数字（浮点类型）
- 值必须在配置的最小/最大范围内
- 如果未指定最小值，则没有默认值
- 最大值是可选的，但建议设置

### 验证规则
**重要**：验证仅在提交表单时发生，而不是在直接使用 `setTodoCustomField` 时。

- 输入被解析为浮点数（使用表单时）
- 必须大于或等于最小值（使用表单时）
- 必须小于或等于最大值（使用表单时）
- `setTodoCustomField` 接受任何字符串值而不进行验证

### 有效评分示例
对于最小值=1，最大值=5 的字段：
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### 无效评分示例
对于最小值=1，最大值=5 的字段：
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## 配置选项

### 评分尺度设置
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### 常见评分尺度
- **1-5 星**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 绩效**: `min: 1, max: 10`
- **0-100 百分比**: `min: 0, max: 100`
- **自定义尺度**: 任何数字范围

## 所需权限

自定义字段操作遵循标准的基于角色的权限：

| 操作 | 所需角色 |
|------|----------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**注意**：所需的具体角色取决于您项目的自定义角色配置和字段级权限。

## 错误响应

### 验证错误（仅限表单）
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

**重要**：评分值验证（最小/最大约束）仅在提交表单时发生，而不是在直接使用 `setTodoCustomField` 时。

### 找不到自定义字段
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

### 尺度设计
- 在类似字段中使用一致的评分尺度
- 考虑用户熟悉度（1-5 星，0-10 NPS）
- 设置适当的最小值（0 与 1）
- 为每个评分级别定义明确的含义

### 数据质量
- 在存储之前验证评分值
- 适当使用小数精度
- 考虑用于显示目的的四舍五入
- 提供关于评分含义的明确指导

### 用户体验
- 以视觉方式显示评分尺度（星星、进度条）
- 显示当前值和尺度限制
- 提供评分含义的上下文
- 考虑新记录的默认值

## 常见用例

1. **绩效管理**
   - 员工绩效评分
   - 项目质量评分
   - 任务完成评分
   - 技能水平评估

2. **客户反馈**
   - 满意度评分
   - 产品质量评分
   - 服务体验评分
   - 净推荐值（NPS）

3. **优先级和重要性**
   - 任务优先级级别
   - 紧急评分
   - 风险评估评分
   - 影响评分

4. **质量保证**
   - 代码审查评分
   - 测试质量评分
   - 文档质量
   - 过程遵循评分

## 集成功能

### 与自动化
- 根据评分阈值触发操作
- 发送低评分通知
- 为高评分创建后续任务
- 根据评分值分配工作

### 与查找
- 计算记录的平均评分
- 按评分范围查找记录
- 从其他记录引用评分数据
- 汇总评分统计信息

### 与 Blue 前端
- 在表单上下文中自动范围验证
- 视觉评分输入控件
- 实时验证反馈
- 星星或滑块输入选项

## 活动跟踪

评分字段的更改会自动跟踪：
- 旧评分值和新评分值被记录
- 活动显示数字变化
- 所有评分更新的时间戳
- 用户归属更改

## 限制

- 仅支持数字值
- 没有内置的视觉评分显示（星星等）
- 小数精度取决于数据库配置
- 不存储评分元数据（评论、上下文）
- 不进行自动评分聚合或统计
- 不支持不同尺度之间的评分转换
- **关键**：最小/最大验证仅在表单中有效，而不通过 `setTodoCustomField`

## 相关资源

- [数字字段](/api/5.custom%20fields/number) - 用于一般数字数据
- [百分比字段](/api/5.custom%20fields/percent) - 用于百分比值
- [选择字段](/api/5.custom%20fields/select-single) - 用于离散选择评分
- [自定义字段概述](/api/5.custom%20fields/2.list-custom-fields) - 一般概念