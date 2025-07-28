---
title: 電子郵件自訂欄位
description: 創建電子郵件欄位以儲存和驗證電子郵件地址
---

電子郵件自訂欄位允許您在記錄中儲存電子郵件地址並進行內建驗證。它們非常適合追蹤聯絡資訊、受讓人電子郵件或任何與電子郵件相關的數據。

## 基本範例

創建一個簡單的電子郵件欄位：

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

## 進階範例

創建一個帶有描述的電子郵件欄位：

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

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 電子郵件欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `EMAIL` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

## 設定電子郵件值

要在記錄上設定或更新電子郵件值：

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

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 電子郵件自訂欄位的 ID |
| `text` | String | 否 | 要儲存的電子郵件地址 |

## 使用電子郵件值創建記錄

當使用電子郵件值創建新記錄時：

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

## 回應欄位

### CustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | ID! | 自訂欄位的唯一識別碼 |
| `name` | String! | 電子郵件欄位的顯示名稱 |
| `type` | CustomFieldType! | 欄位類型 (EMAIL) |
| `description` | String | 欄位的幫助文本 |
| `value` | JSON | 包含電子郵件值 (見下文) |
| `createdAt` | DateTime! | 欄位創建的時間 |
| `updatedAt` | DateTime! | 欄位最後修改的時間 |

**重要**：電子郵件值是通過 `customField.value.text` 欄位訪問，而不是直接在回應上。

## 查詢電子郵件值

當查詢具有電子郵件自訂欄位的記錄時，通過 `customField.value.text` 路徑訪問電子郵件：

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

回應將包含嵌套結構中的電子郵件：

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

## 電子郵件驗證

### 表單驗證
當電子郵件欄位在表單中使用時，它們會自動驗證電子郵件格式：
- 使用標準電子郵件驗證規則
- 剪裁輸入中的空白
- 拒絕無效的電子郵件格式

### 驗證規則
- 必須包含 `@` 符號
- 必須具有有效的域格式
- 自動刪除前導/尾隨空白
- 接受常見的電子郵件格式

### 有效的電子郵件範例
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### 無效的電子郵件範例
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## 重要注意事項

### 直接 API 與表單
- **表單**：自動應用電子郵件驗證
- **直接 API**：無驗證 - 可以儲存任何文本
- **建議**：使用表單進行用戶輸入以確保驗證

### 儲存格式
- 電子郵件地址以純文本儲存
- 無特殊格式或解析
- 大小寫敏感性：EMAIL 自訂欄位以大小寫敏感方式儲存（與用戶身份驗證電子郵件不同，後者會標準化為小寫）
- 除了資料庫限制外，沒有最大長度限制（16MB 限制）

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## 錯誤回應

### 無效的電子郵件格式（僅限表單）
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

### 找不到欄位
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

## 最佳實踐

### 數據輸入
- 始終在應用程序中驗證電子郵件地址
- 僅將電子郵件欄位用於實際的電子郵件地址
- 考慮使用表單進行用戶輸入以獲得自動驗證

### 數據質量
- 在儲存之前剪裁空白
- 考慮大小寫標準化（通常為小寫）
- 在重要操作之前驗證電子郵件格式

### 隱私考量
- 電子郵件地址以純文本儲存
- 考慮數據隱私法規（GDPR、CCPA）
- 實施適當的訪問控制

## 常見用例

1. **聯絡管理**
   - 客戶電子郵件地址
   - 供應商聯絡資訊
   - 團隊成員電子郵件
   - 支援聯絡詳細資訊

2. **專案管理**
   - 利害關係人電子郵件
   - 批准聯絡電子郵件
   - 通知接收者
   - 外部合作者電子郵件

3. **客戶支援**
   - 客戶電子郵件地址
   - 支援票證聯絡人
   - 升級聯絡人
   - 反饋電子郵件地址

4. **銷售與行銷**
   - 潛在客戶電子郵件地址
   - 活動聯絡名單
   - 夥伴聯絡資訊
   - 轉介來源電子郵件

## 整合功能

### 與自動化
- 當電子郵件欄位更新時觸發操作
- 向儲存的電子郵件地址發送通知
- 根據電子郵件變更創建後續任務

### 與查詢
- 參考其他記錄中的電子郵件數據
- 從多個來源聚合電子郵件列表
- 通過電子郵件地址查找記錄

### 與表單
- 自動電子郵件驗證
- 電子郵件格式檢查
- 剪裁空白

## 限制

- 除了格式檢查外，沒有內建的電子郵件驗證或驗證
- 沒有電子郵件特定的 UI 功能（如可點擊的電子郵件鏈接）
- 以純文本儲存，未加密
- 沒有電子郵件撰寫或發送功能
- 沒有電子郵件元數據儲存（顯示名稱等）
- 直接 API 調用繞過驗證（僅表單進行驗證）

## 相關資源

- [文本欄位](/api/custom-fields/text-single) - 用於非電子郵件文本數據
- [URL 欄位](/api/custom-fields/url) - 用於網站地址
- [電話欄位](/api/custom-fields/phone) - 用於電話號碼
- [自訂欄位概述](/api/custom-fields/list-custom-fields) - 一般概念