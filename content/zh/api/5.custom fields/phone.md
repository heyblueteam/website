---
title: 电话自定义字段
description: 创建电话字段以存储和验证国际格式的电话号码
---

电话自定义字段允许您在记录中存储电话号码，并内置验证和国际格式。它们非常适合跟踪联系信息、紧急联系人或您项目中的任何与电话相关的数据。

## 基本示例

创建一个简单的电话字段：

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带描述的电话字段：

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ 是 | 电话字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `PHONE` |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：自定义字段会根据用户当前的项目上下文自动与项目关联。无需 `projectId` 参数。

## 设置电话值

要在记录上设置或更新电话值：

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 电话自定义字段的 ID |
| `text` | String | 否 | 带国家代码的电话号码 |
| `regionCode` | String | 否 | 国家代码（自动检测） |

**注意**：虽然 `text` 在架构中是可选的，但电话号码是使字段有意义所必需的。当使用 `setTodoCustomField` 时，不会执行验证 - 您可以存储任何文本值和 regionCode。自动检测仅在记录创建期间发生。

## 创建带有电话值的记录

在创建带有电话值的新记录时：

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      text
      regionCode
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
| `text` | String | 格式化的电话号码（国际格式） |
| `regionCode` | String | 国家代码（例如，“US”，“GB”，“CA”） |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

## 电话号码验证

**重要**：电话号码验证和格式化仅在通过 `createTodo` 创建新记录时发生。当使用 `setTodoCustomField` 更新现有电话值时，不会执行验证，值将按提供的方式存储。

### 接受的格式（在记录创建期间）
电话号码必须包含国家代码，格式如下：

- **E.164 格式（首选）**： `+12345678900`
- **国际格式**： `+1 234 567 8900`
- **带标点的国际格式**： `+1 (234) 567-8900`
- **带破折号的国家代码**： `+1-234-567-8900`

**注意**：没有国家代码的国家格式（如 `(234) 567-8900`）将在记录创建期间被拒绝。

### 验证规则（在记录创建期间）
- 使用 libphonenumber-js 进行解析和验证
- 接受各种国际电话号码格式
- 自动从号码中检测国家
- 以国际显示格式格式化号码（例如， `+1 234 567 8900`）
- 单独提取并存储国家代码（例如， `US`）

### 有效电话示例
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### 无效电话示例
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## 存储格式

在创建带有电话号码的记录时：
- **text**：在验证后以国际格式存储（例如， `+1 234 567 8900`）
- **regionCode**：以 ISO 国家代码存储（例如， `US`， `GB`， `CA`）自动检测

通过 `setTodoCustomField` 更新时：
- **text**：按提供的方式存储（无格式）
- **regionCode**：按提供的方式存储（无验证）

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## 错误响应

### 无效的电话号码格式
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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

### 缺少国家代码
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 最佳实践

### 数据输入
- 始终在电话号码中包含国家代码
- 使用 E.164 格式以保持一致性
- 在存储重要操作之前验证号码
- 考虑显示格式的区域偏好

### 数据质量
- 以国际格式存储号码以确保全球兼容性
- 使用 regionCode 进行国家特定功能
- 在关键操作（短信、电话）之前验证电话号码
- 考虑联系时间的时区影响

### 国际考虑
- 国家代码会自动检测并存储
- 号码以国际标准格式进行格式化
- 区域显示偏好可以使用 regionCode
- 显示时考虑当地拨号惯例

## 常见用例

1. **联系管理**
   - 客户电话号码
   - 供应商联系信息
   - 团队成员电话号码
   - 支持联系详情

2. **紧急联系人**
   - 紧急联系电话
   - 值班联系信息
   - 危机响应联系人
   - 升级电话号码

3. **客户支持**
   - 客户电话号码
   - 支持回电号码
   - 验证电话号码
   - 跟进联系号码

4. **销售与营销**
   - 潜在客户电话号码
   - 活动联系名单
   - 合作伙伴联系信息
   - 推荐来源电话

## 集成功能

### 与自动化
- 在电话字段更新时触发操作
- 向存储的电话号码发送短信通知
- 根据电话变更创建后续任务
- 根据电话号码数据路由电话

### 与查找
- 引用其他记录中的电话号码数据
- 从多个来源聚合电话号码列表
- 按电话号码查找记录
- 交叉引用联系信息

### 与表单
- 自动电话号码验证
- 国际格式检查
- 国家代码检测
- 实时格式反馈

## 限制

- 所有号码都需要国家代码
- 没有内置的短信或拨打功能
- 除格式检查外，不进行电话号码验证
- 不存储电话号码元数据（运营商、类型等）
- 没有国家代码的国家格式号码会被拒绝
- 除国际标准外，UI 中没有自动电话号码格式化

## 相关资源

- [文本字段](/api/custom-fields/text-single) - 用于非电话文本数据
- [电子邮件字段](/api/custom-fields/email) - 用于电子邮件地址
- [网址字段](/api/custom-fields/url) - 用于网站地址
- [自定义字段概述](/custom-fields/list-custom-fields) - 一般概念