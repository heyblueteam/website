---
title: 创建仪表板
description: 在Blue中创建一个新的数据可视化和报告仪表板
---

## 创建仪表板

`createDashboard` 变更允许您在您的公司或项目中创建一个新的仪表板。仪表板是强大的可视化工具，帮助团队跟踪指标、监控进展并做出数据驱动的决策。

### 基本示例

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### 项目特定仪表板

创建一个与特定项目相关联的仪表板：

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## 输入参数

### CreateDashboardInput

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `companyId` | String! | ✅ 是 | 将要创建仪表板的公司的ID |
| `title` | String! | ✅ 是 | 仪表板的名称。必须是非空字符串 |
| `projectId` | String | 否 | 可选的与此仪表板关联的项目ID |

## 响应字段

该变更返回一个完整的 `Dashboard` 对象：

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 创建的仪表板的唯一标识符 |
| `title` | String! | 提供的仪表板标题 |
| `companyId` | String! | 此仪表板所属的公司 |
| `projectId` | String | 关联的项目ID（如果提供） |
| `project` | Project | 关联的项目对象（如果提供了projectId） |
| `createdBy` | User! | 创建仪表板的用户（您） |
| `dashboardUsers` | [DashboardUser!]! | 具有访问权限的用户列表（最初只有创建者） |
| `createdAt` | DateTime! | 创建仪表板的时间戳 |
| `updatedAt` | DateTime! | 最后修改的时间戳（对于新仪表板与创建时间相同） |

### DashboardUser 字段

当创建仪表板时，创建者会自动作为仪表板用户添加：

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 仪表板用户关系的唯一标识符 |
| `user` | User! | 具有访问仪表板权限的用户对象 |
| `role` | DashboardRole! | 用户的角色（创建者获得完全访问权限） |
| `dashboard` | Dashboard! | 回溯到仪表板的引用 |

## 所需权限

任何属于指定公司的经过身份验证的用户都可以创建仪表板。没有特殊的角色要求。

| 用户状态 | 可以创建仪表板 |
|-------------|-------------------|
| Company Member | ✅ 是 |
| 非公司成员 | ❌ 否 |
| Unauthenticated | ❌ 否 |

## 错误响应

### 无效公司
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### 用户不在公司
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### 无效项目
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### 标题为空
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 重要说明

- **自动所有权**：创建仪表板的用户自动成为其所有者，拥有完全权限
- **项目关联**：如果您提供了一个 `projectId`，它必须属于同一公司
- **初始权限**：最初只有创建者有访问权限。使用 `editDashboard` 添加更多用户
- **标题要求**：仪表板标题必须是非空字符串。没有唯一性要求
- **公司成员资格**：您必须是公司的成员才能在其中创建仪表板

## 仪表板创建工作流程

1. **使用此变更创建仪表板**
2. **使用仪表板构建器UI配置图表和小部件**
3. **使用 `editDashboard` 变更与 `dashboardUsers` 添加团队成员**
4. **通过仪表板界面设置过滤器和日期范围**
5. **使用其唯一ID共享或嵌入仪表板**

## 用例

1. **高管仪表板**：创建公司指标的高层概述
2. **项目跟踪**：构建特定项目的仪表板以监控进展
3. **团队绩效**：跟踪团队生产力和成就指标
4. **客户报告**：为客户面对的报告创建仪表板
5. **实时监控**：设置用于实时操作数据的仪表板

## 最佳实践

1. **命名约定**：使用清晰、描述性的标题，指示仪表板的目的
2. **项目关联**：在项目特定时将仪表板链接到项目
3. **访问管理**：创建后立即添加团队成员以便协作
4. **组织**：使用一致的命名模式创建仪表板层次结构

## 相关操作

- [列出仪表板](/api/dashboards/) - 检索公司或项目的所有仪表板
- [编辑仪表板](/api/dashboards/rename-dashboard) - 重命名仪表板或管理用户
- [复制仪表板](/api/dashboards/copy-dashboard) - 复制现有仪表板
- [删除仪表板](/api/dashboards/delete-dashboard) - 移除仪表板