---
title: URL自定义字段
description: 创建URL字段以存储网站地址和链接
---

URL自定义字段允许您在记录中存储网站地址和链接。它们非常适合跟踪项目网站、参考链接、文档URL或与您的工作相关的任何基于网络的资源。

## 基本示例

创建一个简单的URL字段：

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## 高级示例

创建一个带描述的URL字段：

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
    }
  ) {
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
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | URL字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `URL` |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意：** `projectId` 作为单独的参数传递给变更，而不是作为输入对象的一部分。

## 设置URL值

要在记录上设置或更新URL值：

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput参数

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的记录的ID |
| `customFieldId` | String! | ✅ 是 | URL自定义字段的ID |
| `text` | String! | ✅ 是 | 要存储的URL地址 |

## 创建带有URL值的记录

创建新记录时带有URL值：

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
    }
  }
}
```

## 响应字段

### TodoCustomField响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `text` | String | 存储的URL地址 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

## URL验证

### 当前实现
- **直接API**：当前不强制执行URL格式验证
- **表单**：计划进行URL验证，但目前未激活
- **存储**：可以在URL字段中存储任何字符串值

### 计划中的验证
未来版本将包括：
- HTTP/HTTPS协议验证
- 有效URL格式检查
- 域名验证
- 自动协议前缀添加

### 推荐的URL格式
虽然当前不强制执行，但请使用这些标准格式：

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## 重要说明

### 存储格式
- URL以纯文本形式存储，不进行修改
- 不自动添加协议 (http://, https://)
- 保留输入时的大小写敏感性
- 不执行URL编码/解码

### 直接API与表单
- **表单**：计划进行URL验证（目前未激活）
- **直接API**：没有验证 - 可以存储任何文本
- **建议**：在存储之前在您的应用程序中验证URL

### URL与文本字段
- **URL**：语义上用于网页地址
- **TEXT_SINGLE**：一般单行文本
- **后端**：当前存储和验证相同
- **前端**：用于数据输入的不同UI组件

## 所需权限

自定义字段操作使用基于角色的权限：

| 操作 | 所需角色 |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**注意：** 权限是根据项目中的用户角色检查的，而不是特定的权限常量。

## 错误响应

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

### 必需字段验证（仅限表单）
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 最佳实践

### URL格式标准
- 始终包含协议（http://或https://）
- 尽可能使用HTTPS以提高安全性
- 在存储之前测试URL以确保它们可访问
- 考虑使用缩短的URL以便于显示

### 数据质量
- 在存储之前在您的应用程序中验证URL
- 检查常见的拼写错误（缺少协议、错误的域名）
- 在您的组织中标准化URL格式
- 考虑URL的可访问性和可用性

### 安全考虑
- 对用户提供的URL保持谨慎
- 如果限制到特定网站，请验证域名
- 考虑对恶意内容进行URL扫描
- 在处理敏感数据时使用HTTPS URL

## 过滤和搜索

### 包含搜索
URL字段支持子字符串搜索：

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### 搜索能力
- 不区分大小写的子字符串匹配
- 部分域名匹配
- 路径和参数搜索
- 无协议特定过滤

## 常见用例

1. **项目管理**
   - 项目网站
   - 文档链接
   - 存储库URL
   - 演示网站

2. **内容管理**
   - 参考材料
   - 来源链接
   - 媒体资源
   - 外部文章

3. **客户支持**
   - 客户网站
   - 支持文档
   - 知识库文章
   - 视频教程

4. **销售与市场营销**
   - 公司网站
   - 产品页面
   - 营销材料
   - 社交媒体资料

## 集成功能

### 与查找
- 从其他记录引用URL
- 按域名或URL模式查找记录
- 显示相关的网络资源
- 从多个来源聚合链接

### 与表单
- URL特定输入组件
- 计划进行适当URL格式的验证
- 链接预览功能（前端）
- 可点击的URL显示

### 与报告
- 跟踪URL使用情况和模式
- 监控损坏或无法访问的链接
- 按域名或协议分类
- 导出URL列表以进行分析

## 限制

### 当前限制
- 没有活动的URL格式验证
- 没有自动协议添加
- 没有链接验证或可访问性检查
- 没有URL缩短或扩展
- 没有favicon或预览生成

### 自动化限制
- 不可用作自动化触发字段
- 不能用于自动化字段更新
- 可以在自动化条件中引用
- 可在电子邮件模板和Webhook中使用

### 一般约束
- 没有内置链接预览功能
- 没有自动URL缩短
- 没有点击跟踪或分析
- 没有URL过期检查
- 没有恶意URL扫描

## 未来增强

### 计划功能
- HTTP/HTTPS协议验证
- 自定义正则表达式验证模式
- 自动协议前缀添加
- URL可访问性检查

### 潜在改进
- 链接预览生成
- favicon显示
- URL缩短集成
- 点击跟踪功能
- 损坏链接检测

## 相关资源

- [文本字段](/api/custom-fields/text-single) - 用于非URL文本数据
- [电子邮件字段](/api/custom-fields/email) - 用于电子邮件地址
- [自定义字段概述](/api/custom-fields/2.list-custom-fields) - 一般概念

## 从文本字段迁移

如果您正在从文本字段迁移到URL字段：

1. **创建URL字段**，使用相同的名称和配置
2. **导出现有文本值**以验证它们是否有效的URL
3. **更新记录**以使用新的URL字段
4. **成功迁移后删除旧文本字段**
5. **更新应用程序**以使用URL特定的UI组件

### 迁移示例
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```