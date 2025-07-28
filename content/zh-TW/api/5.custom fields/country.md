---
title: 國家自訂欄位
description: 創建具有 ISO 國家代碼驗證的國家選擇欄位
---

國家自訂欄位允許您存儲和管理記錄的國家資訊。該欄位支持國家名稱和 ISO Alpha-2 國家代碼。

**重要**：國家驗證和轉換行為在不同的變更中有顯著差異：
- **createTodo**：自動驗證並將國家名稱轉換為 ISO 代碼
- **setTodoCustomField**：接受任何值而不進行驗證

## 基本範例

創建一個簡單的國家欄位：

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 進階範例

創建一個帶有描述的國家欄位：

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 國家欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `COUNTRY` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：`projectId` 不會在輸入中傳遞，而是由 GraphQL 上下文決定（通常來自請求標頭或身份驗證）。

## 設定國家值

國家欄位在兩個數據庫欄位中存儲數據：
- **`countryCodes`**：將 ISO Alpha-2 國家代碼作為以逗號分隔的字符串存儲在數據庫中（通過 API 返回為數組）
- **`text`**：將顯示文本或國家名稱作為字符串存儲

### 理解參數

`setTodoCustomField` 變更接受兩個可選參數用於國家欄位：

| 參數 | 類型 | 必需 | 描述 | 功能 |
|------|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID | - |
| `customFieldId` | String! | ✅ 是 | 國家自訂欄位的 ID | - |
| `countryCodes` | [String!] | 否 | ISO Alpha-2 國家代碼的數組 | Stored in the `countryCodes` field |
| `text` | String | 否 | 顯示文本或國家名稱 | Stored in the `text` field |

**重要**：
- 在 `setTodoCustomField`：兩個參數都是可選的，並且獨立存儲
- 在 `createTodo`：系統根據您的輸入自動設置這兩個欄位（您無法獨立控制它們）

### 選項 1：僅使用國家代碼

存儲經過驗證的 ISO 代碼而不顯示文本：

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

結果：`countryCodes` = `["US"]`, `text` = `null`

### 選項 2：僅使用文本

存儲顯示文本而不經過驗證的代碼：

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

結果：`countryCodes` = `null`, `text` = `"United States"`

**注意**：使用 `setTodoCustomField` 時，無論您使用哪個參數，都不會進行驗證。值將按提供的方式存儲。

### 選項 3：同時使用兩者（推薦）

同時存儲經過驗證的代碼和顯示文本：

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

結果：`countryCodes` = `["US"]`, `text` = `"United States"`

### 多個國家

使用數組存儲多個國家：

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## 使用國家值創建記錄

在創建記錄時，`createTodo` 變更 **自動驗證並轉換** 國家值。這是唯一執行國家驗證的變更：

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      countryCodes
    }
  }
}
```

### 接受的輸入格式

| 輸入類型 | 範例 | 結果 |
|-----------|------|------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **不支持** - 被視為單一無效值 |
| Mixed format | `"United States, CA"` | **不支持** - 被視為單一無效值 |

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 |
| `text` | String | 顯示文本（國家名稱） |
| `countryCodes` | [String!] | ISO Alpha-2 國家代碼的數組 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

## 國家標準

Blue 使用 **ISO 3166-1 Alpha-2** 標準來定義國家代碼：

- 兩個字母的國家代碼（例如，美國、英國、法國、德國）
- 驗證使用 `i18n-iso-countries` 庫 **僅在 createTodo 中發生**
- 支持所有官方認可的國家

### 範例國家代碼

| 國家 | ISO 代碼 |
|------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

要查看完整的官方 ISO 3166-1 alpha-2 國家代碼列表，請訪問 [ISO 在線瀏覽平台](https://www.iso.org/obp/ui/#search/code/)。

## 驗證

**驗證僅在 `createTodo` 變更中發生**：

1. **有效的 ISO 代碼**：接受任何有效的 ISO Alpha-2 代碼
2. **國家名稱**：自動將已識別的國家名稱轉換為代碼
3. **無效輸入**：對於未識別的值拋出 `CustomFieldValueParseError`

**注意**：`setTodoCustomField` 變更不執行任何驗證，並接受任何字符串值。

### 錯誤範例

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 整合功能

### 查詢欄位
國家欄位可以被查詢自訂欄位引用，允許您從相關記錄中提取國家數據。

### 自動化
在自動化條件中使用國家值：
- 按特定國家過濾操作
- 根據國家發送通知
- 根據地理區域路由任務

### 表單
表單中的國家欄位自動驗證用戶輸入並將國家名稱轉換為代碼。

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## 錯誤回應

### 無效的國家值
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 欄位類型不匹配
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## 最佳實踐

### 輸入處理
- 使用 `createTodo` 進行自動驗證和轉換
- 小心使用 `setTodoCustomField`，因為它會繞過驗證
- 考慮在您的應用程序中驗證輸入，然後再使用 `setTodoCustomField`
- 在 UI 中顯示完整的國家名稱以提高清晰度

### 數據質量
- 在輸入點驗證國家輸入
- 在您的系統中使用一致的格式
- 考慮報告的區域分組

### 多個國家
- 在 `setTodoCustomField` 中使用數組支持以處理多個國家
- 在 `createTodo` 中的多個國家 **不支持** 通過值欄位
- 在 `setTodoCustomField` 中將國家代碼存儲為數組以便正確處理

## 常見用例

1. **客戶管理**
   - 客戶總部位置
   - 運送目的地
   - 稅務管轄區

2. **項目追蹤**
   - 項目位置
   - 團隊成員位置
   - 市場目標

3. **合規與法律**
   - 監管管轄區
   - 數據居住要求
   - 出口管制

4. **銷售與市場營銷**
   - 領土分配
   - 市場細分
   - 活動目標

## 限制

- 只支持 ISO 3166-1 Alpha-2 代碼（2 字母代碼）
- 不支持國家細分（州/省）
- 不支持自動國旗圖標（僅基於文本）
- 無法驗證歷史國家代碼
- 不支持內建的區域或大陸分組
- **驗證僅在 `createTodo` 中有效，而在 `setTodoCustomField` 中無效**
- **在 `createTodo` 值欄位中不支持多個國家**
- **國家代碼存儲為以逗號分隔的字符串，而不是實際的數組**

## 相關資源

- [自訂欄位概述](/custom-fields/list-custom-fields) - 一般自訂欄位概念
- [查詢欄位](/api/custom-fields/lookup) - 從其他記錄引用國家數據
- [表單 API](/api/forms) - 在自訂表單中包含國家欄位