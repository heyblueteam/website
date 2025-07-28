---
title: 單行文本自定義字段
description: 創建單行文本字段以存儲短文本值，例如名稱、標題和標籤
---

單行文本自定義字段允許您存儲用於單行輸入的短文本值。它們非常適合名稱、標題、標籤或任何應該顯示在單行上的文本數據。

## 基本示例

創建一個簡單的單行文本字段：

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## 高級示例

創建一個帶描述的單行文本字段：

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 文本字段的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `TEXT_SINGLE` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：項目上下文會自動從您的身份驗證標頭中確定。無需 `projectId` 參數。

## 設置文本值

要在記錄上設置或更新單行文本值：

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的記錄的 ID |
| `customFieldId` | String! | ✅ 是 | 文本自定義字段的 ID |
| `text` | String | 否 | 要存儲的單行文本內容 |

## 使用文本值創建記錄

在創建帶有單行文本值的新記錄時：

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## 響應字段

### TodoCustomField 響應

| 字段 | 類型 | 描述 |
|-------|------|-------------|
| `id` | ID! | 字段值的唯一標識符 |
| `customField` | CustomField! | 自定義字段定義（包含文本值） |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

**重要**：文本值通過 `customField.value.text` 字段訪問，而不是直接在 TodoCustomField 上。

## 查詢文本值

在查詢帶有文本自定義字段的記錄時，通過 `customField.value.text` 路徑訪問文本：

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

響應將包含嵌套結構中的文本：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## 文本驗證

### 表單驗證
當單行文本字段用於表單時：
- 自動修剪前導和尾隨空格
- 如果字段標記為必需，則應用必需驗證
- 不應用特定格式驗證

### 驗證規則
- 接受任何字符串內容，包括換行符（雖然不建議）
- 沒有字符長度限制（最多到數據庫限制）
- 支持 Unicode 字符和特殊符號
- 換行符被保留，但不適用於此字段類型

### 典型文本示例
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## 重要注意事項

### 存儲容量
- 使用 MySQL `MediumText` 類型存儲
- 支持最多 16MB 的文本內容
- 與多行文本字段的存儲相同
- 用於國際字符的 UTF-8 編碼

### 直接 API 與表單
- **表單**：自動修剪空格和必需驗證
- **直接 API**：文本按提供的方式存儲
- **建議**：使用表單進行用戶輸入，以確保一致的格式

### TEXT_SINGLE 與 TEXT_MULTI
- **TEXT_SINGLE**：單行文本輸入，適合短值
- **TEXT_MULTI**：多行文本區域輸入，適合較長內容
- **後端**：兩者使用相同的存儲和驗證
- **前端**：不同的 UI 組件用於數據輸入
- **意圖**：TEXT_SINGLE 在語義上適合單行值

## 所需權限

| 操作 | 所需權限 |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## 錯誤響應

### 必需字段驗證（僅限表單）
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

## 最佳實踐

### 內容指南
- 保持文本簡潔且適合單行
- 避免換行以便於顯示在單行上
- 對於相似數據類型使用一致的格式
- 根據您的 UI 要求考慮字符限制

### 數據輸入
- 提供清晰的字段描述以指導用戶
- 使用表單進行用戶輸入以確保驗證
- 如果需要，驗證應用程序中的內容格式
- 考慮使用下拉選單來標準化值

### 性能考量
- 單行文本字段輕量且性能良好
- 考慮對經常搜索的字段進行索引
- 在您的 UI 中使用適當的顯示寬度
- 監控顯示目的的內容長度

## 過濾和搜索

### 包含搜索
單行文本字段支持子字符串搜索：

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### 搜索能力
- 不區分大小寫的子字符串匹配
- 支持部分單詞匹配
- 精確值匹配
- 不支持全文搜索或排名

## 常見用例

1. **標識符和代碼**
   - 產品 SKU
   - 訂單號
   - 參考代碼
   - 版本號

2. **名稱和標題**
   - 客戶名稱
   - 項目標題
   - 產品名稱
   - 類別標籤

3. **簡短描述**
   - 簡要摘要
   - 狀態標籤
   - 優先級指示
   - 分類標籤

4. **外部參考**
   - 票據號
   - 發票參考
   - 外部系統 ID
   - 文檔號

## 集成功能

### 與查找
- 從其他記錄引用文本數據
- 根據文本內容查找記錄
- 顯示相關的文本信息
- 從多個來源聚合文本值

### 與表單
- 自動修剪空格
- 必需字段驗證
- 單行文本輸入 UI
- 字符限制顯示（如果配置）

### 與導入/導出
- 直接 CSV 列映射
- 自動文本值分配
- 批量數據導入支持
- 導出到電子表格格式

## 限制

### 自動化限制
- 不能直接作為自動化觸發字段
- 不能用於自動化字段更新
- 可以在自動化條件中引用
- 可用於電子郵件模板和 Webhook

### 一般限制
- 沒有內置文本格式或樣式
- 除了必需字段外，沒有自動驗證
- 沒有內置的唯一性強制
- 對於非常大的文本沒有內容壓縮
- 沒有版本控制或變更跟踪
- 限制的搜索能力（不支持全文搜索）

## 相關資源

- [多行文本字段](/api/custom-fields/text-multi) - 用於較長的文本內容
- [電子郵件字段](/api/custom-fields/email) - 用於電子郵件地址
- [URL 字段](/api/custom-fields/url) - 用於網站地址
- [唯一 ID 字段](/api/custom-fields/unique-id) - 用於自動生成的標識符
- [自定義字段概述](/api/custom-fields/list-custom-fields) - 一般概念