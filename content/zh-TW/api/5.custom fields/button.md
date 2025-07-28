---
title: 按鈕自訂欄位
description: 創建互動式按鈕欄位，當被點擊時觸發自動化
---

按鈕自訂欄位提供互動式 UI 元素，當被點擊時觸發自動化。與其他存儲數據的自訂欄位類型不同，按鈕欄位作為操作觸發器來執行配置的工作流程。

## 基本範例

創建一個簡單的按鈕欄位以觸發自動化：

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

## 進階範例

創建一個需要確認的按鈕：

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

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 按鈕的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `BUTTON` |
| `projectId` | String! | ✅ 是 | 將創建欄位的專案 ID |
| `buttonType` | String | 否 | 確認行為（見下方按鈕類型） |
| `buttonConfirmText` | String | 否 | 用戶必須輸入的硬確認文本 |
| `description` | String | 否 | 顯示給用戶的幫助文本 |
| `required` | Boolean | 否 | 欄位是否為必填（默認為 false） |
| `isActive` | Boolean | 否 | 欄位是否為活動（默認為 true） |

### 按鈕類型欄位

`buttonType` 欄位是一個自由格式的字符串，可以被 UI 客戶端用來確定確認行為。常見值包括：

- `""`（空） - 無確認
- `"soft"` - 簡單的確認對話框
- `"hard"` - 需要輸入確認文本

**注意**：這些僅僅是 UI 提示。API 不會驗證或強制特定值。

## 觸發按鈕點擊

要觸發按鈕點擊並執行相關的自動化：

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### 點擊輸入參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 包含按鈕的任務 ID |
| `customFieldId` | String! | ✅ 是 | 按鈕自訂欄位的 ID |

### 重要：API 行為

**所有通過 API 的按鈕點擊會立即執行**，無論任何 `buttonType` 或 `buttonConfirmText` 設置。這些欄位是為了讓 UI 客戶端實現確認對話框，但 API 本身：

- 不會驗證確認文本
- 不會強制任何確認要求
- 在調用時立即執行按鈕操作

確認純粹是客戶端 UI 的安全功能。

### 範例：點擊不同的按鈕類型

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

上述三個變異在通過 API 調用時會立即執行按鈕操作，繞過任何確認要求。

## 回應欄位

### 自訂欄位回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 自訂欄位的唯一標識符 |
| `name` | String! | 按鈕的顯示名稱 |
| `type` | CustomFieldType! | 對於按鈕欄位始終是 `BUTTON` |
| `buttonType` | String | 確認行為設置 |
| `buttonConfirmText` | String | 所需的確認文本（如果使用硬確認） |
| `description` | String | 用戶的幫助文本 |
| `required` | Boolean! | 欄位是否為必填 |
| `isActive` | Boolean! | 欄位是否當前有效 |
| `projectId` | String! | 此欄位所屬專案的 ID |
| `createdAt` | DateTime! | 欄位創建的時間 |
| `updatedAt` | DateTime! | 欄位最後修改的時間 |

## 按鈕欄位的工作原理

### 自動化整合

按鈕欄位旨在與 Blue 的自動化系統協同工作：

1. **使用上述變異創建按鈕欄位**
2. **配置自動化**，以監聽 `CUSTOM_FIELD_BUTTON_CLICKED` 事件
3. **用戶在 UI 中點擊按鈕**
4. **自動化執行**配置的操作

### 事件流程

當按鈕被點擊時：

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### 無數據存儲

重要：按鈕欄位不存儲任何值數據。它們純粹作為操作觸發器。每次點擊：
- 生成一個事件
- 觸發相關的自動化
- 在任務歷史中記錄一個操作
- 不修改任何欄位值

## 所需權限

用戶需要適當的專案角色來創建和使用按鈕欄位：

| 操作 | 所需角色 |
|------|----------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## 錯誤回應

### 權限被拒絕
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

### 自訂欄位未找到
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

**注意**：API 不會對缺少的自動化或確認不匹配返回特定錯誤。

## 最佳實踐

### 命名慣例
- 使用以行動為導向的名稱：“發送發票”、“創建報告”、“通知團隊”
- 清楚說明按鈕的功能
- 避免使用“按鈕 1”或“點擊這裡”等通用名稱

### 確認設置
- 對於安全、可逆的操作，將 `buttonType` 留空
- 設置 `buttonType` 以建議 UI 客戶端的確認行為
- 使用 `buttonConfirmText` 指定用戶在 UI 確認中應輸入的內容
- 記住：這些僅僅是 UI 提示 - API 調用始終立即執行

### 自動化設計
- 將按鈕操作集中於單一工作流程
- 提供清晰的反饋，告知點擊後發生了什麼
- 考慮添加描述文本以解釋按鈕的目的

## 常見用例

1. **工作流程轉換**
   - “標記為完成”
   - “發送批准”
   - “存檔任務”

2. **外部整合**
   - “同步到 CRM”
   - “生成發票”
   - “發送電子郵件更新”

3. **批量操作**
   - “更新所有子任務”
   - “複製到專案”
   - “應用模板”

4. **報告操作**
   - “生成報告”
   - “導出數據”
   - “創建摘要”

## 限制

- 按鈕不能存儲或顯示數據值
- 每個按鈕只能觸發自動化，而不是直接的 API 調用（不過，自動化可以包含 HTTP 請求操作來調用外部 API 或 Blue 的 API）
- 按鈕的可見性不能有條件地控制
- 每次點擊最多執行一次自動化（儘管該自動化可以觸發多個操作）

## 相關資源

- [自動化 API](/api/automations/index) - 配置按鈕觸發的操作
- [自訂欄位概述](/custom-fields/list-custom-fields) - 一般自訂欄位概念