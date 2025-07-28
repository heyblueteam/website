---
title: 时间持续时间自定义字段
description: 创建计算的时间持续时间字段，以跟踪工作流中事件之间的时间
---

时间持续时间自定义字段会自动计算并显示工作流中两个事件之间的持续时间。它们非常适合跟踪处理时间、响应时间、周期时间或项目中的任何基于时间的指标。

## 基本示例

创建一个简单的时间持续时间字段，以跟踪任务完成所需的时间：

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## 高级示例

创建一个复杂的时间持续时间字段，以跟踪自定义字段更改之间的时间，并设定SLA目标：

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## 输入参数

### CreateCustomFieldInput (TIME_DURATION)

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 持续时间字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `TIME_DURATION` |
| `description` | String | 否 | 显示给用户的帮助文本 |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ 是 | 如何显示持续时间 |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ 是 | 开始事件配置 |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ 是 | 结束事件配置 |
| `timeDurationTargetTime` | Float | 否 | 用于SLA监控的目标持续时间（以秒为单位） |

### CustomFieldTimeDurationInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `type` | CustomFieldTimeDurationType! | ✅ 是 | 要跟踪的事件类型 |
| `condition` | CustomFieldTimeDurationCondition! | ✅ 是 | `FIRST` 或 `LAST` 发生 |
| `customFieldId` | String | Conditional | 对于 `TODO_CUSTOM_FIELD` 类型是必需的 |
| `customFieldOptionIds` | [String!] | Conditional | 对于选择字段更改是必需的 |
| `todoListId` | String | Conditional | 对于 `TODO_MOVED` 类型是必需的 |
| `tagId` | String | Conditional | 对于 `TODO_TAG_ADDED` 类型是必需的 |
| `assigneeId` | String | Conditional | 对于 `TODO_ASSIGNEE_ADDED` 类型是必需的 |

### CustomFieldTimeDurationType 值

| 值 | 描述 |
|----|------|
| `TODO_CREATED_AT` | 记录创建的时间 |
| `TODO_CUSTOM_FIELD` | 自定义字段值更改的时间 |
| `TODO_DUE_DATE` | 设置截止日期的时间 |
| `TODO_MARKED_AS_COMPLETE` | 记录被标记为完成的时间 |
| `TODO_MOVED` | 记录被移动到不同列表的时间 |
| `TODO_TAG_ADDED` | 记录被添加标签的时间 |
| `TODO_ASSIGNEE_ADDED` | 记录被分配人的时间 |

### CustomFieldTimeDurationCondition 值

| 值 | 描述 |
|----|------|
| `FIRST` | 使用事件的第一次发生 |
| `LAST` | 使用事件的最后一次发生 |

### CustomFieldTimeDurationDisplayType 值

| 值 | 描述 | 示例 |
|----|------|------|
| `FULL_DATE` | 天:小时:分钟:秒格式 | `"01:02:03:04"` |
| `FULL_DATE_STRING` | 完整单词书写 | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | 带单位的数字 | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | 仅以天为单位的持续时间 | `"2.5"` (2.5 days) |
| `HOURS` | 仅以小时为单位的持续时间 | `"60"` (60 hours) |
| `MINUTES` | 仅以分钟为单位的持续时间 | `"3600"` (3600 minutes) |
| `SECONDS` | 仅以秒为单位的持续时间 | `"216000"` (216000 seconds) |

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `number` | Float | 持续时间（以秒为单位） |
| `value` | Float | 数字的别名（持续时间，以秒为单位） |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后更新的时间 |

### CustomField 响应 (TIME_DURATION)

| 字段 | 类型 | 描述 |
|------|------|------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | 持续时间的显示格式 |
| `timeDurationStart` | CustomFieldTimeDuration | 开始事件配置 |
| `timeDurationEnd` | CustomFieldTimeDuration | 结束事件配置 |
| `timeDurationTargetTime` | Float | 目标持续时间（以秒为单位，用于SLA监控） |

## 持续时间计算

### 工作原理
1. **开始事件**：系统监控指定的开始事件
2. **结束事件**：系统监控指定的结束事件
3. **计算**：持续时间 = 结束时间 - 开始时间
4. **存储**：持续时间以数字形式存储（以秒为单位）
5. **显示**：根据 `timeDurationDisplay` 设置格式化

### 更新触发器
当以下情况发生时，持续时间值会自动重新计算：
- 记录被创建或更新
- 自定义字段值发生变化
- 标签被添加或移除
- 分配人被添加或移除
- 记录在列表之间移动
- 记录被标记为完成/未完成

## 读取持续时间值

### 查询持续时间字段
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### 格式化显示值
持续时间值会根据 `timeDurationDisplay` 设置自动格式化：

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## 常见配置示例

### 任务完成时间
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### 状态变化持续时间
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### 特定列表中的时间
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### 分配响应时间
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## 错误响应

### 配置无效
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### 找不到引用的字段
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

### 缺少必需的选项
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 重要说明

### 自动计算
- 持续时间字段是**只读**的 - 值会自动计算
- 您无法通过API手动设置持续时间值
- 计算通过后台作业异步进行
- 当触发事件发生时，值会自动更新

### 性能考虑
- 持续时间计算是排队并异步处理的
- 大量持续时间字段可能会影响性能
- 在设计持续时间字段时，请考虑触发事件的频率
- 使用特定条件以避免不必要的重新计算

### 空值
当以下情况发生时，持续时间字段将显示 `null`：
- 开始事件尚未发生
- 结束事件尚未发生
- 配置引用不存在的实体
- 计算遇到错误

## 最佳实践

### 配置设计
- 尽可能使用特定事件类型，而不是通用类型
- 根据您的工作流选择适当的 `FIRST` 与 `LAST` 条件
- 在部署之前使用示例数据测试持续时间计算
- 为团队成员记录您的持续时间字段逻辑

### 显示格式
- 对于最易读的格式，使用 `FULL_DATE_SUBSTRING`
- 对于紧凑且宽度一致的显示，使用 `FULL_DATE`
- 对于正式报告和文档，使用 `FULL_DATE_STRING`
- 对于简单的数字显示，使用 `DAYS`、`HOURS`、`MINUTES` 或 `SECONDS`
- 在选择格式时，请考虑您的UI空间限制

### 使用目标时间进行SLA监控
在使用 `timeDurationTargetTime` 时：
- 设定目标持续时间（以秒为单位）
- 将实际持续时间与目标进行比较，以确保符合SLA
- 在仪表板中使用，以突出逾期项目
- 示例：24小时响应SLA = 86400秒

### 工作流集成
- 设计持续时间字段以匹配您的实际业务流程
- 使用持续时间数据进行流程改进和优化
- 监控持续时间趋势以识别工作流瓶颈
- 如有需要，设置持续时间阈值的警报

## 常见用例

1. **流程性能**
   - 任务完成时间
   - 审核周期时间
   - 批准处理时间
   - 响应时间

2. **SLA监控**
   - 首次响应时间
   - 解决时间
   - 升级时间框架
   - 服务水平合规性

3. **工作流分析**
   - 瓶颈识别
   - 流程优化
   - 团队绩效指标
   - 质量保证时间

4. **项目管理**
   - 阶段持续时间
   - 里程碑时间
   - 资源分配时间
   - 交付时间框架

## 限制

- 持续时间字段是**只读**的，无法手动设置
- 值是异步计算的，可能不会立即可用
- 需要在您的工作流中设置适当的事件触发器
- 无法计算尚未发生的事件的持续时间
- 限于跟踪离散事件之间的时间（而不是连续时间跟踪）
- 没有内置的SLA警报或通知
- 无法将多个持续时间计算聚合到单个字段中

## 相关资源

- [数字字段](/api/custom-fields/number) - 用于手动数字值
- [日期字段](/api/custom-fields/date) - 用于特定日期跟踪
- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般概念
- [自动化](/api/automations) - 用于基于持续时间阈值触发操作