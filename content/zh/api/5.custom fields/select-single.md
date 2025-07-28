---
title: 单选自定义字段
description: 创建单选字段以允许用户从预定义列表中选择一个选项
---

单选自定义字段允许用户从预定义列表中选择一个选项。它们非常适合状态字段、类别、优先级或任何只需从受控选项集中做出一个选择的场景。

## 基本示例

创建一个简单的单选字段：

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个具有预定义选项的单选字段：

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 单选字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `SELECT_SINGLE` |
| `description` | String | 否 | 显示给用户的帮助文本 |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | 否 | 字段的初始选项 |

### CreateCustomFieldOptionInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `title` | String! | ✅ 是 | 选项的显示文本 |
| `color` | String | 否 | 选项的十六进制颜色代码 |

## 向现有字段添加选项

向现有单选字段添加新选项：

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## 设置单选值

要在记录上设置选定的选项：

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 单选自定义字段的 ID |
| `customFieldOptionId` | String | 否 | 要选择的选项的 ID（单选时首选） |
| `customFieldOptionIds` | [String!] | 否 | 选项 ID 的数组（单选时使用第一个元素） |

## 查询单选值

查询记录的单选值：

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

`value` 字段返回一个 JSON 对象，包含所选选项的详细信息。

## 使用单选值创建记录

创建带有单选值的新记录时：

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
    }
  }
}
```

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `value` | JSON | 包含所选选项对象，具有 id、title、color、position |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

### CustomFieldOption 响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 选项的唯一标识符 |
| `title` | String! | 选项的显示文本 |
| `color` | String | 用于视觉表示的十六进制颜色代码 |
| `position` | Float | 选项的排序顺序 |
| `customField` | CustomField! | 此选项所属的自定义字段 |

### CustomField 响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 字段的唯一标识符 |
| `name` | String! | 单选字段的显示名称 |
| `type` | CustomFieldType! | 始终是 `SELECT_SINGLE` |
| `description` | String | 字段的帮助文本 |
| `customFieldOptions` | [CustomFieldOption!] | 所有可用选项 |

## 值格式

### 输入格式
- **API 参数**: 使用 `customFieldOptionId` 作为单个选项 ID
- **替代**: 使用 `customFieldOptionIds` 数组（取第一个元素）
- **清除选择**: 省略两个字段或传递空值

### 输出格式
- **GraphQL 响应**: `value` 字段中的 JSON 对象，包含 {id, title, color, position}
- **活动日志**: 选项标题作为字符串
- **自动化数据**: 选项标题作为字符串

## 选择行为

### 独占选择
- 设置新选项会自动移除先前的选择
- 一次只能选择一个选项
- 设置 `null` 或空值会清除选择

### 回退逻辑
- 如果提供了 `customFieldOptionIds` 数组，则仅使用第一个选项
- 这确保与多选输入格式的兼容性
- 空数组或 null 值会清除选择

## 管理选项

### 更新选项属性
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### 删除选项
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**注意**: 删除选项将从所有选择了该选项的记录中清除它。

### 重新排序选项
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## 验证规则

### 选项验证
- 提供的选项 ID 必须存在
- 选项必须属于指定的自定义字段
- 只能选择一个选项（自动强制执行）
- Null/空值是有效的（无选择）

### 字段验证
- 必须定义至少一个选项才能可用
- 选项标题在字段内必须唯一
- 颜色代码必须是有效的十六进制格式（如果提供）

## 所需权限

| 操作 | 所需权限 |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## 错误响应

### 无效选项 ID
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### 选项不属于字段
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### 无法解析值
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 最佳实践

### 选项设计
- 使用清晰、描述性的选项标题
- 应用有意义的颜色编码
- 保持选项列表集中且相关
- 按逻辑顺序排列选项（按优先级、频率等）

### 状态字段模式
- 在项目中使用一致的状态工作流程
- 考虑选项的自然进展
- 包含明确的最终状态（完成、已取消等）
- 使用反映选项含义的颜色

### 数据管理
- 定期审查并清理未使用的选项
- 使用一致的命名约定
- 考虑选项删除对现有记录的影响
- 计划选项更新和迁移

## 常见用例

1. **状态和工作流程**
   - 任务状态（待办、进行中、完成）
   - 审批状态（待处理、已批准、已拒绝）
   - 项目阶段（规划、开发、测试、发布）
   - 问题解决状态

2. **分类和归类**
   - 优先级级别（低、中、高、紧急）
   - 任务类型（缺陷、功能、增强、文档）
   - 项目类别（内部、客户、研究）
   - 部门分配

3. **质量和评估**
   - 审查状态（未开始、审核中、已批准）
   - 质量评级（差、一般、好、优秀）
   - 风险级别（低、中、高）
   - 信心级别

4. **分配和所有权**
   - 团队分配
   - 部门所有权
   - 基于角色的分配
   - 区域分配

## 集成功能

### 与自动化
- 在选择特定选项时触发操作
- 根据选择的类别路由工作
- 发送状态更改通知
- 基于选择创建条件工作流程

### 与查找
- 根据选择的选项过滤记录
- 从其他记录引用选项数据
- 基于选项选择创建报告
- 按选择的值分组记录

### 与表单
- 下拉输入控件
- 单选按钮界面
- 选项验证和过滤
- 基于选择的条件字段显示

## 活动跟踪

单选字段的更改会自动跟踪：
- 显示旧的和新的选项选择
- 在活动日志中显示选项标题
- 所有选择更改的时间戳
- 修改的用户归属

## 与多选的区别

| 特性 | 单选 | 多选 |
|---------|---------------|--------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## 限制

- 一次只能选择一个选项
- 没有层次或嵌套选项结构
- 选项在使用该字段的所有记录中共享
- 没有内置的选项分析或使用跟踪
- 颜色代码仅用于显示，没有功能影响
- 不能为每个选项设置不同的权限

## 相关资源

- [多选字段](/api/custom-fields/select-multi) - 用于多项选择
- [复选框字段](/api/custom-fields/checkbox) - 用于简单布尔选择
- [文本字段](/api/custom-fields/text-single) - 用于自由格式文本输入
- [自定义字段概述](/api/custom-fields/1.index) - 一般概念