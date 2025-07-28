---
title: 邮件自定义字段
description: 创建邮件字段以存储和验证电子邮件地址
---

邮件自定义字段允许您在记录中存储电子邮件地址，并具有内置的验证功能。它们非常适合跟踪联系信息、受让人电子邮件或您项目中的任何与电子邮件相关的数据。

## 基本示例

创建一个简单的电子邮件字段：

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带描述的电子邮件字段：

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ 是 | 电子邮件字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `EMAIL` |
| `description` | String | 否 | 显示给用户的帮助文本 |

## 设置电子邮件值

要在记录上设置或更新电子邮件值：

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 电子邮件自定义字段的 ID |
| `text` | String | 否 | 要存储的电子邮件地址 |

## 使用电子邮件值创建记录

在使用电子邮件值创建新记录时：

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## 响应字段

### CustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | ID! | 自定义字段的唯一标识符 |
| `name` | String! | 电子邮件字段的显示名称 |
| `type` | CustomFieldType! | 字段类型 (EMAIL) |
| `description` | String | 字段的帮助文本 |
| `value` | JSON | 包含电子邮件值 (见下文) |
| `createdAt` | DateTime! | 字段创建的时间 |
| `updatedAt` | DateTime! | 字段最后修改的时间 |

**重要**: 电子邮件值通过 `customField.value.text` 字段访问，而不是直接在响应中。

## 查询电子邮件值

在查询具有电子邮件自定义字段的记录时，通过 `customField.value.text` 路径访问电子邮件：

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

响应将包括嵌套结构中的电子邮件：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## 电子邮件验证

### 表单验证
当电子邮件字段在表单中使用时，它们会自动验证电子邮件格式：
- 使用标准电子邮件验证规则
- 修剪输入中的空格
- 拒绝无效的电子邮件格式

### 验证规则
- 必须包含 `@` 符号
- 必须具有有效的域格式
- 自动删除前导/尾随空格
- 接受常见电子邮件格式

### 有效电子邮件示例
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### 无效电子邮件示例
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## 重要说明

### 直接 API 与表单
- **表单**: 应用自动电子邮件验证
- **直接 API**: 无验证 - 可以存储任何文本
- **建议**: 使用表单进行用户输入以确保验证

### 存储格式
- 电子邮件地址以纯文本形式存储
- 无特殊格式或解析
- 大小写敏感: EMAIL 自定义字段以区分大小写存储（与用户身份验证电子邮件不同，后者规范化为小写）
- 除数据库约束外，没有最大长度限制（16MB 限制）

## 所需权限

| 操作 | 所需权限 |
|------|----------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## 错误响应

### 无效的电子邮件格式（仅限表单）
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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
      "code": "NOT_FOUND"
    }
  }]
}
```

## 最佳实践

### 数据输入
- 始终在您的应用程序中验证电子邮件地址
- 仅将电子邮件字段用于实际的电子邮件地址
- 考虑使用表单进行用户输入以获得自动验证

### 数据质量
- 存储前修剪空格
- 考虑大小写规范化（通常为小写）
- 在重要操作之前验证电子邮件格式

### 隐私考虑
- 电子邮件地址以纯文本形式存储
- 考虑数据隐私法规（GDPR，CCPA）
- 实施适当的访问控制

## 常见用例

1. **联系管理**
   - 客户电子邮件地址
   - 供应商联系信息
   - 团队成员电子邮件
   - 支持联系详情

2. **项目管理**
   - 利益相关者电子邮件
   - 批准联系电子邮件
   - 通知接收者
   - 外部合作者电子邮件

3. **客户支持**
   - 客户电子邮件地址
   - 支持票据联系人
   - 升级联系人
   - 反馈电子邮件地址

4. **销售与营销**
   - 潜在客户电子邮件地址
   - 活动联系人列表
   - 合作伙伴联系信息
   - 推荐来源电子邮件

## 集成功能

### 与自动化
- 当电子邮件字段更新时触发操作
- 向存储的电子邮件地址发送通知
- 根据电子邮件更改创建后续任务

### 与查找
- 从其他记录引用电子邮件数据
- 从多个来源聚合电子邮件列表
- 按电子邮件地址查找记录

### 与表单
- 自动电子邮件验证
- 电子邮件格式检查
- 空格修剪

## 限制

- 除格式检查外，没有内置的电子邮件验证或验证功能
- 没有电子邮件特定的 UI 功能（如可点击的电子邮件链接）
- 以纯文本形式存储，无加密
- 没有电子邮件撰写或发送功能
- 不存储电子邮件元数据（显示名称等）
- 直接 API 调用绕过验证（仅表单进行验证）

## 相关资源

- [文本字段](/api/custom-fields/text-single) - 用于非电子邮件文本数据
- [URL 字段](/api/custom-fields/url) - 用于网站地址
- [电话字段](/api/custom-fields/phone) - 用于电话号码
- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般概念