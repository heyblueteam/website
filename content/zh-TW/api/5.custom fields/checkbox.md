---
title: 复選框自訂欄位
description: 創建布林值復選框欄位以用於是/否或真/假數據
---

復選框自訂欄位為任務提供了一個簡單的布林（真/假）輸入。它們非常適合二元選擇、狀態指示器或跟蹤某項任務是否已完成。

## 基本範例

創建一個簡單的復選框欄位：

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

## 進階範例

創建一個帶有描述和驗證的復選框欄位：

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

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 復選框的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `CHECKBOX` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

## 設定復選框值

要在任務上設置或更新復選框值：

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

要取消選中復選框：

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的任務 ID |
| `customFieldId` | String! | ✅ 是 | 復選框自訂欄位的 ID |
| `checked` | Boolean | 否 | 設為 true 以選中，false 以取消選中 |

## 使用復選框值創建任務

在創建帶有復選框值的新任務時：

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

### 接受的字串值

在創建任務時，復選框值必須作為字串傳遞：

| 字串值 | 結果 |
|--------|------|
| `"true"` | ✅ 已選中（區分大小寫） |
| `"1"` | ✅ 已選中 |
| `"checked"` | ✅ 已選中（區分大小寫） |
| Any other value | ❌ 未選中 |

**注意**：在創建任務時，字串比較是區分大小寫的。值必須完全匹配 `"true"`、`"1"` 或 `"checked"` 才能導致選中狀態。

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | ID! | 欄位值的唯一標識符 |
| `uid` | String! | 替代唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 |
| `checked` | Boolean | 復選框狀態（真/假/空） |
| `todo` | Todo! | 此值所屬的任務 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

## 自動化整合

復選框欄位根據狀態變更觸發不同的自動化事件：

| 行動 | 觸發事件 | 描述 |
|------|----------|------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | 當復選框被選中時觸發 |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | 當復選框被取消選中時觸發 |

這使您能夠創建對復選框狀態變更作出響應的自動化，例如：
- 當項目獲得批准時發送通知
- 當審核復選框被選中時移動任務
- 根據復選框狀態更新相關欄位

## 數據導入/導出

### 導入復選框值

通過 CSV 或其他格式導入數據時：
- `"true"`、`"yes"` → 已選中（不區分大小寫）
- 任何其他值（包括 `"false"`、`"no"`、`"0"`、空） → 未選中

### 導出復選框值

導出數據時：
- 已選中的框導出為 `"X"`
- 未選中的框導出為空字串 `""`

## 所需權限

| 行動 | 所需權限 |
|------|----------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## 錯誤回應

### 無效的值類型
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

### 找不到欄位
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

## 最佳實踐

### 命名慣例
- 使用清晰、以行動為導向的名稱： "已批准"、"已審核"、"已完成"
- 避免混淆用戶的負面名稱：更喜歡使用 "是否啟用" 而不是 "是否禁用"
- 明確說明復選框所代表的內容

### 何時使用復選框
- **二元選擇**：是/否、真/假、已完成/未完成
- **狀態指示器**：已批准、已審核、已發布
- **功能標誌**：具有優先支持、需要簽名
- **簡單跟蹤**：電子郵件已發送、發票已支付、項目已發貨

### 何時不使用復選框
- 當您需要超過兩個選項時（請使用 SELECT_SINGLE）
- 對於數字或文本數據（請使用 NUMBER 或 TEXT 欄位）
- 當您需要跟蹤誰選中或何時選中時（請使用審計日誌）

## 常見用例

1. **批准工作流程**
   - "經理批准"
   - "客戶簽字"
   - "法律審查完成"

2. **任務管理**
   - "被阻塞"
   - "準備審核"
   - "高優先級"

3. **質量控制**
   - "QA 通過"
   - "文件完成"
   - "測試已編寫"

4. **行政標誌**
   - "發票已發送"
   - "合同已簽署"
   - "需要跟進"

## 限制

- 復選框欄位只能存儲真/假值（不支持三狀態或初始設置後的空值）
- 無法配置默認值（始終在設置之前為空）
- 無法存儲額外的元數據，如誰選中或何時選中
- 無法根據其他欄位值進行條件可見性設置

## 相關資源

- [自訂欄位概述](/api/custom-fields/list-custom-fields) - 一般自訂欄位概念
- [自動化 API](/api/automations) - 創建由復選框變更觸發的自動化