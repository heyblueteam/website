---
title: 复选框自定义字段
description: 创建用于是/否或真/假的布尔复选框字段
---

复选框自定义字段为任务提供了简单的布尔（真/假）输入。它们非常适合二元选择、状态指示器或跟踪某项任务是否已完成。

## 基本示例

创建一个简单的复选框字段：

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带有描述和验证的复选框字段：

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ 是 | 复选框的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `CHECKBOX` |
| `description` | String | 否 | 显示给用户的帮助文本 |

## 设置复选框值

要设置或更新任务上的复选框值：

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

要取消选中复选框：

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的任务的 ID |
| `customFieldId` | String! | ✅ 是 | 复选框自定义字段的 ID |
| `checked` | Boolean | 否 | true 选中，false 取消选中 |

## 创建带有复选框值的任务

创建新任务时带有复选框值：

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### 接受的字符串值

创建任务时，复选框值必须作为字符串传递：

| 字符串值 | 结果 |
|----------|------|
| `"true"` | ✅ 已选中（区分大小写） |
| `"1"` | ✅ 已选中 |
| `"checked"` | ✅ 已选中（区分大小写） |
| Any other value | ❌ 未选中 |

**注意**：任务创建期间的字符串比较是区分大小写的。值必须完全匹配 `"true"`、`"1"` 或 `"checked"` 才能结果为选中状态。

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | ID! | 字段值的唯一标识符 |
| `uid` | String! | 备用唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `checked` | Boolean | 复选框状态（真/假/空） |
| `todo` | Todo! | 此值所属的任务 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

## 自动化集成

复选框字段根据状态变化触发不同的自动化事件：

| 操作 | 触发事件 | 描述 |
|------|----------|------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | 复选框被选中时触发 |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | 复选框被取消选中时触发 |

这使您能够创建响应复选框状态变化的自动化，例如：
- 当项目被批准时发送通知
- 当审核复选框被选中时移动任务
- 根据复选框状态更新相关字段

## 数据导入/导出

### 导入复选框值

通过 CSV 或其他格式导入数据时：
- `"true"`、`"yes"` → 已选中（不区分大小写）
- 任何其他值（包括 `"false"`、`"no"`、`"0"`、空） → 未选中

### 导出复选框值

导出数据时：
- 已选中的复选框导出为 `"X"`
- 未选中的复选框导出为空字符串 `""`

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## 错误响应

### 无效值类型
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## 最佳实践

### 命名约定
- 使用清晰、以行动为导向的名称：“已批准”、“已审核”、“已完成”
- 避免混淆用户的负面名称：更喜欢“处于活动状态”而不是“处于非活动状态”
- 明确复选框所代表的内容

### 何时使用复选框
- **二元选择**：是/否，真/假，完成/未完成
- **状态指示器**：已批准，已审核，已发布
- **功能标志**：具有优先支持，需要签名
- **简单跟踪**：电子邮件已发送，发票已支付，物品已发货

### 何时不使用复选框
- 当您需要超过两个选项时（使用 SELECT_SINGLE 代替）
- 对于数字或文本数据（使用 NUMBER 或 TEXT 字段）
- 当您需要跟踪谁选中它或何时选中时（使用审计日志）

## 常见用例

1. **审批工作流**
   - “经理已批准”
   - “客户签字”
   - “法律审查完成”

2. **任务管理**
   - “被阻塞”
   - “准备审核”
   - “高优先级”

3. **质量控制**
   - “QA 通过”
   - “文档完成”
   - “测试已编写”

4. **行政标志**
   - “发票已发送”
   - “合同已签署”
   - “需要跟进”

## 限制

- 复选框字段只能存储真/假值（在初始设置后没有三态或空值）
- 无法配置默认值（始终在设置之前为 null）
- 无法存储额外的元数据，例如谁选中它或何时选中
- 无法根据其他字段值的条件可见性

## 相关资源

- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般自定义字段概念
- [自动化 API](/api/automations) - 创建由复选框更改触发的自动化