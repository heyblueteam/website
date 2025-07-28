---
title: 參考自訂欄位
description: 創建參考欄位，鏈接到其他專案中的記錄以實現跨專案關係
---

參考自訂欄位允許您在不同專案中的記錄之間創建鏈接，使跨專案關係和數據共享成為可能。它們提供了一種強大的方式來連接貴組織專案結構中的相關工作。

## 基本範例

創建一個簡單的參考欄位：

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

## 進階範例

創建一個具有過濾和多重選擇的參考欄位：

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

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 參考欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `REFERENCE` |
| `referenceProjectId` | String | 否 | 要參考的專案 ID |
| `referenceMultiple` | Boolean | 否 | 允許多個記錄選擇（預設：false） |
| `referenceFilter` | TodoFilterInput | 否 | 參考記錄的過濾條件 |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：自訂欄位會根據用戶當前的專案上下文自動與專案關聯。

## 參考配置

### 單一與多重參考

**單一參考（預設）：**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- 用戶可以從參考專案中選擇一個記錄
- 返回單個 Todo 對象

**多重參考：**
```graphql
{
  referenceMultiple: true
}
```
- 用戶可以從參考專案中選擇多個記錄
- 返回 Todo 對象的數組

### 參考過濾

使用 `referenceFilter` 來限制可以選擇的記錄：

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

## 設定參考值

### 單一參考

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### 多重參考

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

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 參考自訂欄位的 ID |
| `customFieldReferenceTodoIds` | [String!] | ✅ 是 | 參考記錄 ID 的數組 |

## 創建帶有參考的記錄

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

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | ID! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 參考欄位定義 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

**注意**：參考的 todos 是通過 `customField.selectedTodos` 訪問的，而不是直接在 TodoCustomField 上。

### 參考的 Todo 欄位

每個參考的 Todo 包含：

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | ID! | 參考記錄的唯一標識符 |
| `title` | String! | 參考記錄的標題 |
| `status` | TodoStatus! | 當前狀態（ACTIVE, COMPLETED 等） |
| `description` | String | 參考記錄的描述 |
| `dueDate` | DateTime | 如果設置了到期日 |
| `assignees` | [User!] | 指派的用戶 |
| `tags` | [Tag!] | 關聯的標籤 |
| `project` | Project! | 包含參考記錄的專案 |

## 查詢參考數據

### 基本查詢

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

### 帶有嵌套數據的進階查詢

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

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**重要**：用戶必須對參考專案具有查看權限才能查看鏈接的記錄。

## 跨專案訪問

### 專案可見性

- 用戶只能參考他們有權訪問的專案中的記錄
- 參考記錄遵循原始專案的權限
- 對參考記錄的更改會即時顯示
- 刪除參考記錄會將其從參考欄位中移除

### 權限繼承

- 參考欄位從兩個專案繼承權限
- 用戶需要對參考專案具有查看權限
- 編輯權限基於當前專案的規則
- 在參考欄位的上下文中，參考數據為只讀

## 錯誤回應

### 無效的參考專案

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

### 找不到參考記錄

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

### 權限被拒絕

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

## 最佳實踐

### 欄位設計

1. **清晰命名** - 使用描述性名稱來指示關係
2. **適當過濾** - 設置過濾器以僅顯示相關記錄
3. **考慮權限** - 確保用戶可以訪問參考專案
4. **記錄關係** - 提供清晰的連接描述

### 性能考量

1. **限制參考範圍** - 使用過濾器減少可選記錄的數量
2. **避免深層嵌套** - 不要創建複雜的參考鏈
3. **考慮緩存** - 參考數據會被緩存以提高性能
4. **監控使用情況** - 跟踪參考在專案中的使用情況

### 數據完整性

1. **處理刪除** - 計劃參考記錄被刪除時的情況
2. **驗證權限** - 確保用戶可以訪問參考專案
3. **更新依賴** - 考慮更改參考記錄時的影響
4. **審計跟蹤** - 跟踪參考關係以符合合規要求

## 常見用例

### 專案依賴

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

### 客戶需求

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

### 資源分配

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

### 質量保證

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

## 與查找的整合

參考欄位與 [查找欄位](/api/custom-fields/lookup) 一起工作，以從參考記錄中提取數據。查找欄位可以從在參考欄位中選擇的記錄中提取值，但它們僅是數據提取器（不支持像 SUM 這樣的聚合函數）。

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

- 參考專案必須對用戶可訪問
- 對參考專案權限的更改會影響參考欄位的訪問
- 深層嵌套的參考可能會影響性能
- 沒有內建的循環參考驗證
- 沒有自動限制防止同專案參考
- 設置參考值時不強制執行過濾驗證

## 相關資源

- [查找欄位](/api/custom-fields/lookup) - 從參考記錄中提取數據
- [專案 API](/api/projects) - 管理包含參考的專案
- [記錄 API](/api/records) - 處理具有參考的記錄
- [自訂欄位概述](/api/custom-fields/list-custom-fields) - 一般概念