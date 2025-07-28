---
title: 電話自訂欄位
description: 創建電話欄位以儲存和驗證具有國際格式的電話號碼
---

電話自訂欄位允許您在記錄中儲存電話號碼，並內建驗證和國際格式。它們非常適合追蹤聯絡資訊、緊急聯絡人或任何與電話相關的數據。

## 基本範例

創建一個簡單的電話欄位：

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

## 進階範例

創建一個帶有描述的電話欄位：

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

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 電話欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `PHONE` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：自訂欄位會根據用戶當前的項目上下文自動與項目關聯。無需 `projectId` 參數。

## 設定電話值

要在記錄上設定或更新電話值：

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄的 ID |
| `customFieldId` | String! | ✅ 是 | 電話自訂欄位的 ID |
| `text` | String | 否 | 帶有國家代碼的電話號碼 |
| `regionCode` | String | 否 | 國家代碼（自動檢測） |

**注意**：雖然 `text` 在架構中是可選的，但該欄位需要電話號碼才能有意義。使用 `setTodoCustomField` 時，不會執行驗證 - 您可以儲存任何文本值和 regionCode。自動檢測僅在創建記錄時發生。

## 創建帶有電話值的記錄

當創建帶有電話值的新記錄時：

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

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 |
| `text` | String | 格式化的電話號碼（國際格式） |
| `regionCode` | String | 國家代碼（例如，“US”，“GB”，“CA”） |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

## 電話號碼驗證

**重要**：電話號碼的驗證和格式化僅在通過 `createTodo` 創建新記錄時發生。使用 `setTodoCustomField` 更新現有電話值時，不會執行驗證，並且值將按提供的方式儲存。

### 接受的格式（在記錄創建期間）
電話號碼必須包含國家代碼，格式如下：

- **E.164 格式（首選）**： `+12345678900`
- **國際格式**： `+1 234 567 8900`
- **帶標點的國際格式**： `+1 (234) 567-8900`
- **帶破折號的國家代碼**： `+1-234-567-8900`

**注意**：不帶國家代碼的國內格式（如 `(234) 567-8900`）在創建記錄時將被拒絕。

### 驗證規則（在記錄創建期間）
- 使用 libphonenumber-js 進行解析和驗證
- 接受各種國際電話號碼格式
- 自動從號碼中檢測國家
- 以國際顯示格式格式化號碼（例如， `+1 234 567 8900`）
- 單獨提取並儲存國家代碼（例如， `US`）

### 有效的電話範例
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### 無效的電話範例
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## 儲存格式

當創建帶有電話號碼的記錄時：
- **text**：在驗證後以國際格式儲存（例如， `+1 234 567 8900`）
- **regionCode**：以 ISO 國家代碼儲存（例如， `US`， `GB`， `CA`）自動檢測

當通過 `setTodoCustomField` 更新時：
- **text**：按提供的方式儲存（無格式）
- **regionCode**：按提供的方式儲存（無驗證）

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## 錯誤回應

### 無效的電話格式
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

### 缺少國家代碼
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

## 最佳實踐

### 數據輸入
- 始終在電話號碼中包含國家代碼
- 使用 E.164 格式以保持一致性
- 在儲存重要操作之前驗證號碼
- 考慮顯示格式的區域偏好

### 數據質量
- 以國際格式儲存號碼以便於全球兼容性
- 使用 regionCode 以獲得國家特定功能
- 在關鍵操作（短信、通話）之前驗證電話號碼
- 考慮聯絡時間的時區影響

### 國際考量
- 國家代碼會自動檢測並儲存
- 號碼以國際標準格式化
- 區域顯示偏好可以使用 regionCode
- 考慮顯示時的本地撥號慣例

## 常見用例

1. **聯絡管理**
   - 客戶電話號碼
   - 供應商聯絡資訊
   - 團隊成員電話號碼
   - 支援聯絡詳情

2. **緊急聯絡人**
   - 緊急聯絡電話
   - 隨叫隨到的聯絡資訊
   - 危機應對聯絡人
   - 升級電話號碼

3. **客戶支援**
   - 客戶電話號碼
   - 支援回撥電話號碼
   - 驗證電話號碼
   - 跟進聯絡電話號碼

4. **銷售與行銷**
   - 潛在客戶電話號碼
   - 活動聯絡名單
   - 合作夥伴聯絡資訊
   - 轉介來源電話

## 整合功能

### 與自動化
- 當電話欄位被更新時觸發操作
- 向儲存的電話號碼發送短信通知
- 根據電話變更創建跟進任務
- 根據電話號碼數據路由通話

### 與查詢
- 參考其他記錄中的電話數據
- 從多個來源匯總電話列表
- 通過電話號碼查找記錄
- 交叉參考聯絡資訊

### 與表單
- 自動電話驗證
- 國際格式檢查
- 國家代碼檢測
- 實時格式反饋

## 限制

- 所有號碼都需要國家代碼
- 沒有內建的短信或通話功能
- 除格式檢查外，沒有電話號碼驗證
- 不儲存電話元數據（運營商、類型等）
- 沒有國家代碼的國內格式號碼將被拒絕
- 在 UI 中不會自動格式化電話號碼，僅限國際標準

## 相關資源

- [文本欄位](/api/custom-fields/text-single) - 用於非電話文本數據
- [電子郵件欄位](/api/custom-fields/email) - 用於電子郵件地址
- [網址欄位](/api/custom-fields/url) - 用於網站地址
- [自訂欄位概述](/custom-fields/list-custom-fields) - 一般概念