---
title: 按钮自定义字段
description: 创建交互式按钮字段，点击时触发自动化
---

按钮自定义字段提供交互式 UI 元素，点击时触发自动化。与存储数据的其他自定义字段类型不同，按钮字段作为操作触发器来执行配置的工作流。

## 基本示例

创建一个简单的按钮字段，触发自动化：

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个需要确认的按钮：

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 按钮的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `BUTTON` |
| `projectId` | String! | ✅ 是 | 字段将被创建的项目 ID |
| `buttonType` | String | 否 | 确认行为（见下面的按钮类型） |
| `buttonConfirmText` | String | 否 | 用户必须输入的确认文本 |
| `description` | String | 否 | 显示给用户的帮助文本 |
| `required` | Boolean | 否 | 字段是否为必填（默认为 false） |
| `isActive` | Boolean | 否 | 字段是否处于激活状态（默认为 true） |

### 按钮类型字段

`buttonType` 字段是一个自由格式字符串，可以被 UI 客户端用于确定确认行为。常见值包括：

- `""`（空） - 无确认
- `"soft"` - 简单确认对话框
- `"hard"` - 需要输入确认文本

**注意**：这些仅是 UI 提示。API 不会验证或强制特定值。

## 触发按钮点击

要触发按钮点击并执行相关自动化：

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### 点击输入参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 包含按钮的任务 ID |
| `customFieldId` | String! | ✅ 是 | 按钮自定义字段的 ID |

### 重要：API 行为

**所有通过 API 的按钮点击立即执行**，无论任何 `buttonType` 或 `buttonConfirmText` 设置。这些字段是为 UI 客户端存储以实现确认对话框，但 API 本身：

- 不验证确认文本
- 不强制任何确认要求
- 在调用时立即执行按钮操作

确认纯粹是客户端 UI 安全功能。

### 示例：点击不同按钮类型

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

上述三个变更将在通过 API 调用时立即执行按钮操作，绕过任何确认要求。

## 响应字段

### 自定义字段响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 自定义字段的唯一标识符 |
| `name` | String! | 按钮的显示名称 |
| `type` | CustomFieldType! | 对于按钮字段始终是 `BUTTON` |
| `buttonType` | String | 确认行为设置 |
| `buttonConfirmText` | String | 所需确认文本（如果使用硬确认） |
| `description` | String | 用户的帮助文本 |
| `required` | Boolean! | 字段是否为必填 |
| `isActive` | Boolean! | 字段当前是否处于激活状态 |
| `projectId` | String! | 此字段所属项目的 ID |
| `createdAt` | DateTime! | 字段创建时间 |
| `updatedAt` | DateTime! | 字段最后修改时间 |

## 按钮字段的工作原理

### 自动化集成

按钮字段旨在与 Blue 的自动化系统配合使用：

1. **使用上述变更创建按钮字段**
2. **配置监听 `CUSTOM_FIELD_BUTTON_CLICKED` 事件的自动化**
3. **用户在 UI 中点击按钮**
4. **自动化执行配置的操作**

### 事件流

当按钮被点击时：

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### 无数据存储

重要提示：按钮字段不存储任何值数据。它们纯粹作为操作触发器。每次点击：
- 生成一个事件
- 触发相关的自动化
- 在任务历史中记录一个操作
- 不修改任何字段值

## 所需权限

用户需要适当的项目角色才能创建和使用按钮字段：

| 操作 | 所需角色 |
|------|----------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## 错误响应

### 权限被拒绝
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### 自定义字段未找到
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**注意**：API 不会返回缺少自动化或确认不匹配的具体错误。

## 最佳实践

### 命名约定
- 使用以行动为导向的名称：“发送发票”、“创建报告”、“通知团队”
- 明确按钮的功能
- 避免使用“按钮 1”或“点击这里”等通用名称

### 确认设置
- 对于安全、可逆的操作，保持 `buttonType` 为空
- 设置 `buttonType` 以向 UI 客户端建议确认行为
- 使用 `buttonConfirmText` 指定用户在 UI 确认中应输入的内容
- 记住：这些仅是 UI 提示 - API 调用始终立即执行

### 自动化设计
- 保持按钮操作专注于单一工作流
- 提供清晰的反馈，说明点击后发生了什么
- 考虑添加描述文本以解释按钮的目的

## 常见用例

1. **工作流转换**
   - “标记为完成”
   - “发送审批”
   - “归档任务”

2. **外部集成**
   - “同步到 CRM”
   - “生成发票”
   - “发送电子邮件更新”

3. **批量操作**
   - “更新所有子任务”
   - “复制到项目”
   - “应用模板”

4. **报告操作**
   - “生成报告”
   - “导出数据”
   - “创建摘要”

## 限制

- 按钮不能存储或显示数据值
- 每个按钮只能触发自动化，而不能直接调用 API（但是，自动化可以包含 HTTP 请求操作以调用外部 API 或 Blue 自己的 API）
- 按钮的可见性不能有条件控制
- 每次点击最多只能执行一个自动化（尽管该自动化可以触发多个操作）

## 相关资源

- [自动化 API](/api/automations/index) - 配置由按钮触发的操作
- [自定义字段概述](/custom-fields/list-custom-fields) - 一般自定义字段概念