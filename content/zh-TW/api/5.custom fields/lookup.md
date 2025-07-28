---
title: 查詢自定義字段
description: 創建自動從引用記錄中提取數據的查詢字段
---

查詢自定義字段自動從由 [引用字段](/api/custom-fields/reference) 引用的記錄中提取數據，顯示來自鏈接記錄的信息，而無需手動複製。當引用數據發生變化時，它們會自動更新。

## 基本示例

創建一個查詢字段以顯示來自引用記錄的標籤：

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 高級示例

創建一個查詢字段以從引用記錄中提取自定義字段值：

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 查詢字段的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ 是 | 查詢配置 |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

## 查詢配置

### CustomFieldLookupOptionInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `referenceId` | String! | ✅ 是 | 用於提取數據的引用字段的 ID |
| `lookupId` | String | 否 | 要查詢的特定自定義字段的 ID（對於 TODO_CUSTOM_FIELD 類型是必需的） |
| `lookupType` | CustomFieldLookupType! | ✅ 是 | 要從引用記錄中提取的數據類型 |

## 查詢類型

### CustomFieldLookupType 值

| 類型 | 描述 | 返回 |
|------|------|------|
| `TODO_DUE_DATE` | 來自引用待辦事項的截止日期 | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | 來自引用待辦事項的創建日期 | Array of creation timestamps |
| `TODO_UPDATED_AT` | 來自引用待辦事項的最後更新日期 | Array of update timestamps |
| `TODO_TAG` | 來自引用待辦事項的標籤 | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | 來自引用待辦事項的指派人 | Array of user objects |
| `TODO_DESCRIPTION` | 來自引用待辦事項的描述 | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | 來自引用待辦事項的待辦清單名稱 | Array of list titles |
| `TODO_CUSTOM_FIELD` | 來自引用待辦事項的自定義字段值 | Array of values based on the field type |

## 響應字段

### CustomField 響應（對於查詢字段）

| 字段 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 字段的唯一標識符 |
| `name` | String! | 查詢字段的顯示名稱 |
| `type` | CustomFieldType! | 將是 `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | 查詢配置和結果 |
| `createdAt` | DateTime! | 字段創建的時間 |
| `updatedAt` | DateTime! | 字段最後更新的時間 |

### CustomFieldLookupOption 結構

| 字段 | 類型 | 描述 |
|------|------|------|
| `lookupType` | CustomFieldLookupType! | 正在執行的查詢類型 |
| `lookupResult` | JSON | 從引用記錄中提取的數據 |
| `reference` | CustomField | 作為來源使用的引用字段 |
| `lookup` | CustomField | 正在查詢的特定字段（對於 TODO_CUSTOM_FIELD） |
| `parentCustomField` | CustomField | 父查詢字段 |
| `parentLookup` | CustomField | 鏈中的父查詢（對於嵌套查詢） |

## 查詢如何運作

1. **數據提取**：查詢從所有通過引用字段鏈接的記錄中提取特定數據
2. **自動更新**：當引用記錄發生變化時，查詢值會自動更新
3. **只讀**：查詢字段不能直接編輯 - 它們始終反映當前的引用數據
4. **無計算**：查詢提取並顯示數據，無需聚合或計算

## TODO_CUSTOM_FIELD 查詢

使用 `TODO_CUSTOM_FIELD` 類型時，您必須使用 `lookupId` 參數指定要提取的自定義字段：

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

這將從所有引用記錄中提取指定自定義字段的值。

## 查詢查詢數據

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## 示例查詢結果

### 標籤查詢結果
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### 指派人查詢結果
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### 自定義字段查詢結果
結果根據正在查詢的自定義字段類型而異。例如，貨幣字段查詢可能返回：
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**重要**：用戶必須對當前項目和引用項目擁有查看權限才能查看查詢結果。

## 錯誤響應

### 無效的引用字段
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

### 檢測到循環查詢
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### TODO_CUSTOM_FIELD 缺少查詢 ID
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## 最佳實踐

1. **清晰命名**：使用描述性名稱以指示正在查詢的數據
2. **適當類型**：選擇符合您數據需求的查詢類型
3. **性能**：查詢處理所有引用記錄，因此要注意具有多個鏈接的引用字段
4. **權限**：確保用戶對引用項目有訪問權限，以便查詢正常運行

## 常見用例

### 跨項目可見性
顯示來自相關項目的標籤、指派人或狀態，而無需手動同步。

### 依賴跟蹤
顯示當前工作依賴的任務的截止日期或完成狀態。

### 資源概覽
顯示分配給引用任務的所有團隊成員以進行資源規劃。

### 狀態聚合
收集來自相關任務的所有唯一狀態，以便一目了然地查看項目健康狀況。

## 限制

- 查詢字段是只讀的，不能直接編輯
- 無聚合函數（SUM、COUNT、AVG） - 查詢僅提取數據
- 無過濾選項 - 所有引用記錄均包含在內
- 防止循環查詢鏈以避免無限循環
- 結果反映當前數據並自動更新

## 相關資源

- [引用字段](/api/custom-fields/reference) - 創建指向記錄的鏈接以用作查詢來源
- [自定義字段值](/api/custom-fields/custom-field-values) - 在可編輯的自定義字段上設置值
- [列出自定義字段](/api/custom-fields/list-custom-fields) - 查詢項目中的所有自定義字段