---
title: 單選自訂欄位
description: 創建單選欄位以允許用戶從預定義列表中選擇一個選項
---

單選自訂欄位允許用戶從預定義列表中選擇一個選項。它們非常適合用於狀態欄位、類別、優先級或任何只能從受控選項集中做出一個選擇的情況。

## 基本範例

創建一個簡單的單選欄位：

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 進階範例

創建一個具有預定義選項的單選欄位：

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 單選欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `SELECT_SINGLE` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | 否 | 欄位的初始選項 |

### CreateCustomFieldOptionInput

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `title` | String! | ✅ 是 | 選項的顯示文本 |
| `color` | String | 否 | 選項的十六進制顏色代碼 |

## 向現有欄位添加選項

向現有的單選欄位添加新選項：

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## 設置單選值

要在記錄上設置所選選項：

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 單選自訂欄位的 ID |
| `customFieldOptionId` | String | 否 | 要選擇的選項 ID（單選時首選） |
| `customFieldOptionIds` | [String!] | 否 | 選項 ID 的數組（單選時使用第一個元素） |

## 查詢單選值

查詢記錄的單選值：

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

`value` 欄位返回一個 JSON 對象，包含所選選項的詳細信息。

## 使用單選值創建記錄

當使用單選值創建新記錄時：

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
    }
  }
}
```

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|-------|------|-------------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 |
| `value` | JSON | 包含所選選項對象的 ID、標題、顏色、位置 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

### CustomFieldOption 回應

| 欄位 | 類型 | 描述 |
|-------|------|-------------|
| `id` | String! | 選項的唯一標識符 |
| `title` | String! | 選項的顯示文本 |
| `color` | String | 用於視覺表示的十六進制顏色代碼 |
| `position` | Float | 選項的排序順序 |
| `customField` | CustomField! | 此選項所屬的自訂欄位 |

### CustomField 回應

| 欄位 | 類型 | 描述 |
|-------|------|-------------|
| `id` | String! | 欄位的唯一標識符 |
| `name` | String! | 單選欄位的顯示名稱 |
| `type` | CustomFieldType! | 始終是 `SELECT_SINGLE` |
| `description` | String | 欄位的幫助文本 |
| `customFieldOptions` | [CustomFieldOption!] | 所有可用選項 |

## 值格式

### 輸入格式
- **API 參數**: 使用 `customFieldOptionId` 作為單個選項 ID
- **替代**: 使用 `customFieldOptionIds` 數組（取第一個元素）
- **清除選擇**: 省略兩個欄位或傳遞空值

### 輸出格式
- **GraphQL 回應**: JSON 對象在 `value` 欄位中包含 {id, title, color, position}
- **活動日誌**: 選項標題作為字符串
- **自動化數據**: 選項標題作為字符串

## 選擇行為

### 獨佔選擇
- 設置新選項會自動移除先前的選擇
- 一次只能選擇一個選項
- 設置 `null` 或空值會清除選擇

### 回退邏輯
- 如果提供 `customFieldOptionIds` 數組，則僅使用第一個選項
- 這確保與多選輸入格式的兼容性
- 空數組或 null 值會清除選擇

## 管理選項

### 更新選項屬性
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### 刪除選項
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**注意**: 刪除選項將從所有選擇了該選項的記錄中清除它。

### 重新排序選項
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## 驗證規則

### 選項驗證
- 提供的選項 ID 必須存在
- 選項必須屬於指定的自訂欄位
- 只能選擇一個選項（自動強制執行）
- Null/空值是有效的（無選擇）

### 欄位驗證
- 必須定義至少一個選項才能可用
- 選項標題在欄位內必須唯一
- 顏色代碼必須是有效的十六進制格式（如果提供）

## 所需權限

| 操作 | 所需權限 |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## 錯誤回應

### 無效的選項 ID
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### 選項不屬於欄位
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### 找不到欄位
```json
{
  "errors": [{
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### 無法解析值
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 最佳實踐

### 選項設計
- 使用清晰、描述性的選項標題
- 應用有意義的顏色編碼
- 保持選項列表集中且相關
- 按邏輯排序選項（按優先級、頻率等）

### 狀態欄位模式
- 在項目中使用一致的狀態工作流程
- 考慮選項的自然進展
- 包括清晰的最終狀態（完成、已取消等）
- 使用反映選項意義的顏色

### 數據管理
- 定期檢查和清理未使用的選項
- 使用一致的命名慣例
- 考慮刪除選項對現有記錄的影響
- 計劃選項的更新和遷移

## 常見用例

1. **狀態和工作流程**
   - 任務狀態（待辦、進行中、完成）
   - 審批狀態（待處理、已批准、已拒絕）
   - 項目階段（計劃、開發、測試、已發布）
   - 問題解決狀態

2. **分類和類別**
   - 優先級別（低、中、高、關鍵）
   - 任務類型（錯誤、功能、增強、文檔）
   - 項目類別（內部、客戶、研究）
   - 部門分配

3. **質量和評估**
   - 審查狀態（未開始、審查中、已批准）
   - 質量評級（差、一般、好、優秀）
   - 風險級別（低、中、高）
   - 信心級別

4. **分配和擁有權**
   - 團隊分配
   - 部門擁有權
   - 基於角色的分配
   - 區域分配

## 整合功能

### 與自動化
- 當選擇特定選項時觸發操作
- 根據選擇的類別路由工作
- 發送狀態變更的通知
- 根據選擇創建條件工作流程

### 與查詢
- 根據選擇的選項過濾記錄
- 從其他記錄引用選項數據
- 根據選項選擇創建報告
- 按選擇的值分組記錄

### 與表單
- 下拉輸入控件
- 單選按鈕界面
- 選項驗證和過濾
- 根據選擇顯示條件欄位

## 活動追蹤

單選欄位的變更會自動追蹤：
- 顯示舊的和新的選項選擇
- 在活動日誌中顯示選項標題
- 所有選擇變更的時間戳
- 修改的用戶歸屬

## 與多選的區別

| 特徵 | 單選 | 多選 |
|---------|---------------|--------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## 限制

- 一次只能選擇一個選項
- 無層級或嵌套的選項結構
- 選項在使用該欄位的所有記錄中共享
- 無內建的選項分析或使用追蹤
- 顏色代碼僅用於顯示，無功能影響
- 不能為每個選項設置不同的權限

## 相關資源

- [多選欄位](/api/custom-fields/select-multi) - 用於多選擇
- [複選框欄位](/api/custom-fields/checkbox) - 用於簡單的布林選擇
- [文本欄位](/api/custom-fields/text-single) - 用於自由格式文本輸入
- [自訂欄位概述](/api/custom-fields/1.index) - 一般概念