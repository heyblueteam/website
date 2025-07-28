---
title: 引用自定义字段
description: 创建引用字段，链接到其他项目中的记录，以实现跨项目关系
---

引用自定义字段允许您在不同项目之间创建记录链接，从而实现跨项目关系和数据共享。它们提供了一种强大的方式来连接您组织项目结构中相关的工作。

## 基本示例

创建一个简单的引用字段：

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## 高级示例

创建一个具有过滤和多重选择的引用字段：

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 引用字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `REFERENCE` |
| `referenceProjectId` | String | 否 | 要引用的项目的 ID |
| `referenceMultiple` | Boolean | 否 | 允许多条记录选择（默认：false） |
| `referenceFilter` | TodoFilterInput | 否 | 引用记录的过滤条件 |
| `description` | String | 否 | 显示给用户的帮助文本 |

**注意**：自定义字段会根据用户当前的项目上下文自动与项目关联。

## 引用配置

### 单一引用与多重引用

**单一引用（默认）：**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- 用户可以从引用项目中选择一条记录
- 返回单个 Todo 对象

**多重引用：**
```graphql
{
  referenceMultiple: true
}
```
- 用户可以从引用项目中选择多条记录
- 返回 Todo 对象数组

### 引用过滤

使用 `referenceFilter` 来限制可以选择的记录：

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## 设置引用值

### 单一引用

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### 多重引用

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 引用自定义字段的 ID |
| `customFieldReferenceTodoIds` | [String!] | ✅ 是 | 引用记录 ID 的数组 |

## 创建带有引用的记录

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | ID! | 字段值的唯一标识符 |
| `customField` | CustomField! | 引用字段定义 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

**注意**：引用的 todos 通过 `customField.selectedTodos` 访问，而不是直接在 TodoCustomField 上访问。

### 引用的 Todo 字段

每个引用的 Todo 包含：

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | ID! | 引用记录的唯一标识符 |
| `title` | String! | 引用记录的标题 |
| `status` | TodoStatus! | 当前状态（ACTIVE, COMPLETED 等） |
| `description` | String | 引用记录的描述 |
| `dueDate` | DateTime | 如果设置了，截止日期 |
| `assignees` | [User!] | 指定的用户 |
| `tags` | [Tag!] | 关联标签 |
| `project` | Project! | 包含引用记录的项目 |

## 查询引用数据

### 基本查询

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### 带有嵌套数据的高级查询

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## 所需权限

| 操作 | 所需权限 |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**重要**：用户必须对引用项目具有查看权限才能查看链接的记录。

## 跨项目访问

### 项目可见性

- 用户只能引用他们有访问权限的项目中的记录
- 引用记录遵循原项目的权限
- 对引用记录的更改实时显示
- 删除引用记录会将其从引用字段中移除

### 权限继承

- 引用字段从两个项目继承权限
- 用户需要对引用项目具有查看权限
- 编辑权限基于当前项目的规则
- 在引用字段的上下文中，引用数据为只读

## 错误响应

### 无效的引用项目

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### 找不到引用记录

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

### 权限被拒绝

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## 最佳实践

### 字段设计

1. **清晰命名** - 使用描述性名称来指示关系
2. **适当过滤** - 设置过滤器以仅显示相关记录
3. **考虑权限** - 确保用户对引用项目有访问权限
4. **记录关系** - 提供连接的清晰描述

### 性能考虑

1. **限制引用范围** - 使用过滤器减少可选择记录的数量
2. **避免深层嵌套** - 不要创建复杂的引用链
3. **考虑缓存** - 引用数据会被缓存以提高性能
4. **监控使用情况** - 跟踪引用在项目中的使用情况

### 数据完整性

1. **处理删除** - 计划处理引用记录被删除的情况
2. **验证权限** - 确保用户可以访问引用项目
3. **更新依赖关系** - 在更改引用记录时考虑影响
4. **审计跟踪** - 跟踪引用关系以确保合规性

## 常见用例

### 项目依赖

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### 客户需求

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### 资源分配

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### 质量保证

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## 与查找的集成

引用字段与 [查找字段](/api/custom-fields/lookup) 一起工作，以从引用记录中提取数据。查找字段可以从在引用字段中选择的记录中提取值，但它们仅是数据提取器（不支持像 SUM 这样的聚合函数）。

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## 限制

- 引用项目必须对用户可访问
- 对引用项目权限的更改会影响引用字段的访问
- 深层嵌套的引用可能会影响性能
- 对循环引用没有内置验证
- 不会自动限制防止同一项目引用
- 设置引用值时不强制执行过滤验证

## 相关资源

- [查找字段](/api/custom-fields/lookup) - 从引用记录中提取数据
- [项目 API](/api/projects) - 管理包含引用的项目
- [记录 API](/api/records) - 处理具有引用的记录
- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般概念