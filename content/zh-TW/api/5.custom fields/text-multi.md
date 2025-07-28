---
title: 多行文本自定義欄位
description: 創建多行文本欄位以容納較長的內容，如描述、備註和評論
---

多行文本自定義欄位允許您儲存帶有換行和格式的較長文本內容。它們非常適合用於描述、備註、評論或任何需要多行的文本數據。

## 基本範例

創建一個簡單的多行文本欄位：

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## 進階範例

創建一個帶有描述的多行文本欄位：

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `name` | String! | ✅ 是 | 文本欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `TEXT_MULTI` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意：** `projectId` 作為單獨的參數傳遞給變更，而不是作為輸入對象的一部分。或者，可以從您的 GraphQL 請求中的 `X-Bloo-Project-ID` 標頭中確定項目上下文。

## 設定文本值

要在記錄上設置或更新多行文本值：

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄的 ID |
| `customFieldId` | String! | ✅ 是 | 文本自定義欄位的 ID |
| `text` | String | 否 | 要儲存的多行文本內容 |

## 使用文本值創建記錄

當使用多行文本值創建新記錄時：

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
    }
  }
}
```

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一識別碼 |
| `customField` | CustomField! | 自定義欄位定義 |
| `text` | String | 儲存的多行文本內容 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

## 文本驗證

### 表單驗證
當多行文本欄位用於表單時：
- 自動修剪前後空白
- 如果欄位標記為必需，則應用必需驗證
- 不應用特定格式驗證

### 驗證規則
- 接受任何字符串內容，包括換行
- 沒有字符長度限制（最多到數據庫限制）
- 支援 Unicode 字符和特殊符號
- 換行在儲存中被保留

### 有效文本範例
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## 重要注意事項

### 儲存容量
- 使用 MySQL `MediumText` 類型儲存
- 支援最多 16MB 的文本內容
- 換行和格式被保留
- 針對國際字符使用 UTF-8 編碼

### 直接 API 與表單
- **表單**：自動修剪空白和必需驗證
- **直接 API**：文本按提供的方式儲存
- **建議**：使用表單進行用戶輸入以確保一致的格式

### TEXT_MULTI 與 TEXT_SINGLE
- **TEXT_MULTI**：多行文本區輸入，適合較長內容
- **TEXT_SINGLE**：單行文本輸入，適合較短值
- **後端**：兩種類型是相同的 - 相同的儲存欄位、驗證和處理
- **前端**：不同的 UI 組件用於數據輸入（文本區 vs 輸入欄位）
- **重要**：TEXT_MULTI 和 TEXT_SINGLE 之間的區別純粹是為了 UI 目的

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

## 錯誤回應

### 必需欄位驗證（僅限表單）
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

### 內容組織
- 對於結構化內容使用一致的格式
- 考慮使用類似 markdown 的語法以提高可讀性
- 將長內容拆分為邏輯部分
- 使用換行以改善可讀性

### 數據輸入
- 提供清晰的欄位描述以指導用戶
- 使用表單進行用戶輸入以確保驗證
- 根據您的使用案例考慮字符限制
- 如有需要，在您的應用程序中驗證內容格式

### 性能考量
- 非常長的文本內容可能會影響查詢性能
- 考慮為顯示大型文本欄位進行分頁
- 搜索功能的索引考量
- 監控大型內容欄位的儲存使用情況

## 過濾和搜索

### 包含搜索
多行文本欄位支援通過自定義欄位過濾器進行子字符串搜索：

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### 搜索能力
- 使用 `CONTAINS` 操作符在文本欄位中進行子字符串匹配
- 使用 `NCONTAINS` 操作符進行不區分大小寫的搜索
- 使用 `IS` 操作符進行精確匹配
- 使用 `NOT` 操作符進行負匹配
- 在所有文本行中進行搜索
- 支援部分單詞匹配

## 常見用例

1. **項目管理**
   - 任務描述
   - 項目需求
   - 會議記錄
   - 狀態更新

2. **客戶支持**
   - 問題描述
   - 解決方案備註
   - 客戶反饋
   - 通信記錄

3. **內容管理**
   - 文章內容
   - 產品描述
   - 用戶評論
   - 評論詳情

4. **文檔**
   - 流程描述
   - 指導說明
   - 指南
   - 參考材料

## 整合功能

### 與自動化
- 當文本內容變更時觸發操作
- 從文本內容中提取關鍵字
- 創建摘要或通知
- 使用外部服務處理文本內容

### 與查詢
- 參考來自其他記錄的文本數據
- 從多個來源聚合文本內容
- 通過文本內容查找記錄
- 顯示相關文本信息

### 與表單
- 自動修剪空白
- 必需欄位驗證
- 多行文本區 UI
- 字符計數顯示（如果配置）

## 限制

- 無內建文本格式或豐富文本編輯
- 無自動鏈接檢測或轉換
- 無拼寫檢查或語法驗證
- 無內建文本分析或處理
- 無版本控制或變更追蹤
- 限制的搜索能力（無全文搜索）
- 對於非常大的文本無內容壓縮

## 相關資源

- [單行文本欄位](/api/custom-fields/text-single) - 用於短文本值
- [電子郵件欄位](/api/custom-fields/email) - 用於電子郵件地址
- [URL 欄位](/api/custom-fields/url) - 用於網站地址
- [自定義欄位概述](/api/custom-fields/2.list-custom-fields) - 一般概念