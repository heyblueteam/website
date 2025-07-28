---
title: 創建儀表板
description: 在 Blue 中創建一個新的數據可視化和報告儀表板
---

## 創建儀表板

`createDashboard` 變異允許您在您的公司或項目中創建一個新的儀表板。儀表板是強大的可視化工具，幫助團隊跟踪指標、監控進度並做出基於數據的決策。

### 基本範例

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

### 特定項目的儀表板

創建與特定項目相關聯的儀表板：

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

## 輸入參數

### CreateDashboardInput

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `companyId` | String! | ✅ 是 | 將創建儀表板的公司的 ID |
| `title` | String! | ✅ 是 | 儀表板的名稱。必須是非空字符串 |
| `projectId` | String | 否 | 與此儀表板關聯的可選項目 ID |

## 回應字段

該變異返回一個完整的 `Dashboard` 對象：

| 字段 | 類型 | 描述 |
|-------|------|-------------|
| `id` | String! | 創建的儀表板的唯一標識符 |
| `title` | String! | 提供的儀表板標題 |
| `companyId` | String! | 此儀表板所屬的公司 |
| `projectId` | String | 關聯的項目 ID（如果提供） |
| `project` | Project | 關聯的項目對象（如果提供了 projectId） |
| `createdBy` | User! | 創建儀表板的用戶（您） |
| `dashboardUsers` | [DashboardUser!]! | 具有訪問權限的用戶列表（最初僅限創建者） |
| `createdAt` | DateTime! | 儀表板創建的時間戳 |
| `updatedAt` | DateTime! | 最後修改的時間戳（對於新儀表板與 createdAt 相同） |

### DashboardUser 字段

當創建儀表板時，創建者會自動添加為儀表板用戶：

| 字段 | 類型 | 描述 |
|-------|------|-------------|
| `id` | String! | 儀表板用戶關係的唯一標識符 |
| `user` | User! | 具有訪問儀表板的用戶對象 |
| `role` | DashboardRole! | 用戶的角色（創建者獲得完全訪問權限） |
| `dashboard` | Dashboard! | 參考回儀表板 |

## 所需權限

任何屬於指定公司的經過身份驗證的用戶都可以創建儀表板。沒有特殊的角色要求。

| 用戶狀態 | 可以創建儀表板 |
|-------------|-------------------|
| Company Member | ✅ 是 |
| 非公司成員 | ❌ 否 |
| Unauthenticated | ❌ 否 |

## 錯誤回應

### 無效的公司
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

### 用戶不在公司中
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

### 無效的項目
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

### 標題為空
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

## 重要說明

- **自動擁有權**：創建儀表板的用戶自動成為其擁有者，並擁有完全權限
- **項目關聯**：如果您提供 `projectId`，則必須屬於同一公司
- **初始權限**：最初只有創建者具有訪問權限。使用 `editDashboard` 添加更多用戶
- **標題要求**：儀表板標題必須是非空字符串。沒有唯一性要求
- **公司成員資格**：您必須是該公司的成員才能在其中創建儀表板

## 儀表板創建工作流程

1. **使用此變異創建儀表板**
2. **使用儀表板構建器 UI 配置圖表和小部件**
3. **使用 `editDashboard` 變異添加團隊成員，並提供 `dashboardUsers`**
4. **通過儀表板界面設置過濾器和日期範圍**
5. **使用其唯一 ID 分享或嵌入儀表板**

## 使用案例

1. **高管儀表板**：創建公司指標的高層次概覽
2. **項目跟踪**：構建特定項目的儀表板以監控進度
3. **團隊績效**：跟踪團隊生產力和成就指標
4. **客戶報告**：為面向客戶的報告創建儀表板
5. **實時監控**：設置實時操作數據的儀表板

## 最佳實踐

1. **命名慣例**：使用清晰、描述性的標題來指示儀表板的目的
2. **項目關聯**：當儀表板是特定於項目時，將其鏈接到項目
3. **訪問管理**：在創建後立即添加團隊成員以便協作
4. **組織**：使用一致的命名模式創建儀表板層次結構

## 相關操作

- [列出儀表板](/api/dashboards/) - 檢索公司或項目的所有儀表板
- [編輯儀表板](/api/dashboards/rename-dashboard) - 重命名儀表板或管理用戶
- [複製儀表板](/api/dashboards/copy-dashboard) - 複製現有的儀表板
- [刪除儀表板](/api/dashboards/delete-dashboard) - 移除儀表板