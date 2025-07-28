---
title: 多选自定义字段
description: 创建多选字段，以允许用户从预定义列表中选择多个选项
---

多选自定义字段允许用户从预定义列表中选择多个选项。它们非常适合用于类别、标签、技能、特性或任何需要从受控选项集中进行多重选择的场景。

## 基本示例

创建一个简单的多选字段：

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个多选字段，然后单独添加选项：

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 多选字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `SELECT_MULTI` |
| `description` | String | 否 | 显示给用户的帮助文本 |
| `projectId` | String! | ✅ 是 | 此字段的项目 ID |

### CreateCustomFieldOptionInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `customFieldId` | String! | ✅ 是 | 自定义字段的 ID |
| `title` | String! | ✅ 是 | 选项的显示文本 |
| `color` | String | 否 | 选项的颜色（任何字符串） |
| `position` | Float | 否 | 选项的排序顺序 |

## 向现有字段添加选项

向现有多选字段添加新选项：

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## 设置多选值

要在记录上设置多个选定选项：

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 多选自定义字段的 ID |
| `customFieldOptionIds` | [String!] | ✅ 是 | 要选择的选项 ID 数组 |

## 使用多选值创建记录

创建新记录时使用多选值：

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
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
| `selectedOptions` | [CustomFieldOption!] | 选定选项的数组 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

### CustomFieldOption 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 选项的唯一标识符 |
| `title` | String! | 选项的显示文本 |
| `color` | String | 视觉表示的十六进制颜色代码 |
| `position` | Float | 选项的排序顺序 |
| `customField` | CustomField! | 此选项所属的自定义字段 |

### CustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段的唯一标识符 |
| `name` | String! | 多选字段的显示名称 |
| `type` | CustomFieldType! | 始终是 `SELECT_MULTI` |
| `description` | String | 字段的帮助文本 |
| `customFieldOptions` | [CustomFieldOption!] | 所有可用选项 |

## 值格式

### 输入格式
- **API 参数**: 选项 ID 数组 (`["option1", "option2", "option3"]`)
- **字符串格式**: 逗号分隔的选项 ID (`"option1,option2,option3"`)

### 输出格式
- **GraphQL 响应**: CustomFieldOption 对象数组
- **活动日志**: 逗号分隔的选项标题
- **自动化数据**: 选项标题数组

## 管理选项

### 更新选项属性
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
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

### 重新排序选项
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## 验证规则

### 选项验证
- 所有提供的选项 ID 必须存在
- 选项必须属于指定的自定义字段
- 只有 SELECT_MULTI 字段可以选择多个选项
- 空数组是有效的（没有选择）

### 字段验证
- 必须定义至少一个选项才能使用
- 选项标题在字段内必须唯一
- 颜色字段接受任何字符串值（不进行十六进制验证）

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## 错误响应

### 无效选项 ID
```json
{
  "errors": [{
    "message": "Custom field option not found",
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### 非多选字段上的多个选项
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 最佳实践

### 选项设计
- 使用描述性、简洁的选项标题
- 应用一致的颜色编码方案
- 保持选项列表可管理（通常 3-20 个选项）
- 逻辑排序选项（按字母顺序、按频率等）

### 数据管理
- 定期审查和清理未使用的选项
- 在项目之间使用一致的命名约定
- 创建字段时考虑选项的可重用性
- 计划选项更新和迁移

### 用户体验
- 提供清晰的字段描述
- 使用颜色提高视觉区分度
- 将相关选项分组在一起
- 考虑常见情况的默认选择

## 常见用例

1. **项目管理**
   - 任务类别和标签
   - 优先级级别和类型
   - 团队成员分配
   - 状态指示器

2. **内容管理**
   - 文章类别和主题
   - 内容类型和格式
   - 发布渠道
   - 审批工作流

3. **客户支持**
   - 问题类别和类型
   - 受影响的产品或服务
   - 解决方法
   - 客户细分

4. **产品开发**
   - 特性类别
   - 技术要求
   - 测试环境
   - 发布渠道

## 集成功能

### 与自动化
- 在选择特定选项时触发操作
- 根据选择的类别路由工作
- 对高优先级选择发送通知
- 根据选项组合创建后续任务

### 与查找
- 按选择的选项过滤记录
- 聚合选项选择的数据
- 从其他记录引用选项数据
- 根据选项组合创建报告

### 与表单
- 多选输入控件
- 选项验证和过滤
- 动态加载选项
- 条件字段显示

## 活动跟踪

多选字段的更改会自动跟踪：
- 显示添加和删除的选项
- 在活动日志中显示选项标题
- 所有选择更改的时间戳
- 用户归属修改

## 限制

- 选项的最大实际限制取决于 UI 性能
- 无层次或嵌套选项结构
- 选项在使用该字段的所有记录之间共享
- 没有内置的选项分析或使用跟踪
- 颜色字段接受任何字符串（不进行十六进制验证）
- 不能为每个选项设置不同的权限
- 选项必须单独创建，而不能与字段创建一起进行
- 没有专用的重新排序变更（使用 editCustomFieldOption 和位置）

## 相关资源

- [单选字段](/api/custom-fields/select-single) - 用于单选选择
- [复选框字段](/api/custom-fields/checkbox) - 用于简单布尔选择
- [文本字段](/api/custom-fields/text-single) - 用于自由格式文本输入
- [自定义字段概述](/api/custom-fields/2.list-custom-fields) - 一般概念