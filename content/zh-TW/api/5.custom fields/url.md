---
title: URL 自訂欄位
description: 建立 URL 欄位以儲存網站地址和連結
---

URL 自訂欄位允許您在記錄中儲存網站地址和連結。它們非常適合追蹤專案網站、參考連結、文件 URL 或與您工作相關的任何基於網路的資源。

## 基本範例

建立一個簡單的 URL 欄位：

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## 進階範例

建立一個帶有描述的 URL 欄位：

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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

| 參數 | 類型 | 必填 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | URL 欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `URL` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意：** `projectId` 作為單獨的參數傳遞給變更，而不是作為輸入對象的一部分。

## 設定 URL 值

要在記錄上設定或更新 URL 值：

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必填 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | URL 自訂欄位的 ID |
| `text` | String! | ✅ 是 | 要儲存的 URL 地址 |

## 使用 URL 值創建記錄

當使用 URL 值創建新記錄時：

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
|-------|------|-------------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 |
| `text` | String | 儲存的 URL 地址 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

## URL 驗證

### 當前實施
- **直接 API**：目前不強制執行 URL 格式驗證
- **表單**：計劃進行 URL 驗證，但目前未啟用
- **儲存**：任何字串值都可以儲存在 URL 欄位中

### 計劃中的驗證
未來版本將包括：
- HTTP/HTTPS 協議驗證
- 有效 URL 格式檢查
- 網域名稱驗證
- 自動協議前綴添加

### 推薦的 URL 格式
雖然目前未強制執行，但請使用這些標準格式：

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## 重要注意事項

### 儲存格式
- URL 以純文本形式儲存，無修改
- 不自動添加協議 (http://, https://)
- 保持輸入的大小寫敏感
- 不進行 URL 編碼/解碼

### 直接 API 與表單
- **表單**：計劃中的 URL 驗證（目前未啟用）
- **直接 API**：無驗證 - 任何文本均可儲存
- **建議**：在儲存之前在您的應用程式中驗證 URL

### URL 與文本欄位
- **URL**：語義上用於網頁地址
- **TEXT_SINGLE**：一般單行文本
- **後端**：目前儲存和驗證相同
- **前端**：不同的 UI 組件用於數據輸入

## 所需權限

自訂欄位操作使用基於角色的權限：

| 操作 | 所需角色 |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**注意：** 權限是根據項目中的用戶角色進行檢查，而不是特定的權限常數。

## 錯誤回應

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

### 必填欄位驗證（僅限表單）
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

## 最佳實踐

### URL 格式標準
- 始終包括協議（http:// 或 https://）
- 儘可能使用 HTTPS 以提高安全性
- 在儲存之前測試 URL 以確保其可訪問
- 考慮使用縮短的 URL 以便於顯示

### 數據質量
- 在儲存之前在您的應用程式中驗證 URL
- 檢查常見的拼寫錯誤（缺少協議、不正確的域名）
- 在您的組織中標準化 URL 格式
- 考慮 URL 的可訪問性和可用性

### 安全考量
- 對用戶提供的 URL 要謹慎
- 如果限制於特定網站，則驗證域名
- 考慮對 URL 進行惡意內容掃描
- 在處理敏感數據時使用 HTTPS URL

## 過濾和搜索

### 包含搜索
URL 欄位支持子字串搜索：

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### 搜索能力
- 不區分大小寫的子字串匹配
- 部分域名匹配
- 路徑和參數搜索
- 無協議特定的過濾

## 常見用例

1. **專案管理**
   - 專案網站
   - 文件連結
   - 儲存庫 URL
   - 演示網站

2. **內容管理**
   - 參考材料
   - 來源連結
   - 媒體資源
   - 外部文章

3. **客戶支持**
   - 客戶網站
   - 支持文檔
   - 知識庫文章
   - 視頻教程

4. **銷售與市場營銷**
   - 公司網站
   - 產品頁面
   - 市場材料
   - 社交媒體資料

## 整合功能

### 與查詢
- 參考其他記錄中的 URL
- 通過域名或 URL 模式查找記錄
- 顯示相關的網頁資源
- 從多個來源聚合連結

### 與表單
- URL 特定的輸入組件
- 計劃中的驗證以確保正確的 URL 格式
- 連結預覽功能（前端）
- 可點擊的 URL 顯示

### 與報告
- 追蹤 URL 使用情況和模式
- 監控損壞或無法訪問的連結
- 按域名或協議分類
- 將 URL 列表導出以進行分析

## 限制

### 當前限制
- 無主動的 URL 格式驗證
- 無自動協議添加
- 無連結驗證或可訪問性檢查
- 無 URL 縮短或擴展
- 無 favicon 或預覽生成

### 自動化限制
- 不可用作自動化觸發欄位
- 不能用於自動化欄位更新
- 可以在自動化條件中引用
- 可用於電子郵件模板和網路鉤子

### 一般約束
- 無內建的連結預覽功能
- 無自動 URL 縮短
- 無點擊追蹤或分析
- 無 URL 到期檢查
- 無惡意 URL 掃描

## 未來增強

### 計劃中的功能
- HTTP/HTTPS 協議驗證
- 自訂正則表達式驗證模式
- 自動協議前綴添加
- URL 可訪問性檢查

### 潛在改進
- 連結預覽生成
- favicon 顯示
- URL 縮短整合
- 點擊追蹤能力
- 損壞連結檢測

## 相關資源

- [文本欄位](/api/custom-fields/text-single) - 用於非 URL 文本數據
- [電子郵件欄位](/api/custom-fields/email) - 用於電子郵件地址
- [自訂欄位概述](/api/custom-fields/2.list-custom-fields) - 一般概念

## 從文本欄位遷移

如果您正在從文本欄位遷移到 URL 欄位：

1. **建立 URL 欄位**，使用相同的名稱和配置
2. **導出現有文本值**以驗證它們是否為有效的 URL
3. **更新記錄**以使用新的 URL 欄位
4. **在成功遷移後刪除舊的文本欄位**
5. **更新應用程式**以使用 URL 特定的 UI 組件

### 遷移範例
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```