---
title: 日期自訂欄位
description: 創建日期欄位以追蹤單一日期或日期範圍，並支援時區
---

日期自訂欄位允許您為記錄儲存單一日期或日期範圍。它們支援時區處理、智能格式化，並可用於追蹤截止日期、事件日期或任何基於時間的信息。

## 基本範例

創建一個簡單的日期欄位：

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## 進階範例

創建一個帶描述的到期日期欄位：

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 日期欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `DATE` |
| `isDueDate` | Boolean | 否 | 此欄位是否表示到期日期 |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：自訂欄位會根據用戶當前的項目上下文自動與項目關聯。無需 `projectId` 參數。

## 設定日期值

日期欄位可以儲存單一日期或日期範圍：

### 單一日期

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### 日期範圍

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### 全天事件

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 日期自訂欄位的 ID |
| `startDate` | DateTime | 否 | ISO 8601 格式的開始日期/時間 |
| `endDate` | DateTime | 否 | ISO 8601 格式的結束日期/時間 |
| `timezone` | String | 否 | 時區標識符（例如，“America/New_York”） |

**注意**：如果僅提供 `startDate`，則 `endDate` 會自動默認為相同的值。

## 日期格式

### ISO 8601 格式
所有日期必須以 ISO 8601 格式提供：
- `2025-01-15T14:30:00Z` - UTC 時間
- `2025-01-15T14:30:00+05:00` - 帶有時區偏移
- `2025-01-15T14:30:00.123Z` - 帶有毫秒

### 時區標識符
使用標準時區標識符：
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

如果未提供時區，系統將默認為用戶檢測到的時區。

## 使用日期值創建記錄

當使用日期值創建新記錄時：

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### 支援的輸入格式

在創建記錄時，可以以各種格式提供日期：

| 格式 | 範例 | 結果 |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|-------|------|-------------|
| `id` | ID! | 欄位值的唯一標識符 |
| `uid` | String! | 唯一標識符字符串 |
| `customField` | CustomField! | 自訂欄位定義（包含日期值） |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

**重要**：日期值（`startDate`、`endDate`、`timezone`）是通過 `customField.value` 欄位訪問的，而不是直接在 TodoCustomField 上。

### 值對象結構

日期值通過 `customField.value` 欄位作為 JSON 對象返回：

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**注意**：`value` 欄位屬於 `CustomField` 類型，而不是 `TodoCustomField`。

## 查詢日期值

在查詢具有日期自訂欄位的記錄時，通過 `customField.value` 欄位訪問日期值：

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

回應將包含 `value` 欄位中的日期值：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## 日期顯示智能

系統根據範圍自動格式化日期：

| 情境 | 顯示格式 |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (不顯示時間) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**全天檢測**：從 00:00 到 23:59 的事件會自動檢測為全天事件。

## 時區處理

### 儲存
- 所有日期在數據庫中以 UTC 儲存
- 時區信息單獨保留
- 顯示時進行轉換

### 最佳實踐
- 始終提供時區以確保準確性
- 在項目內使用一致的時區
- 考慮全球團隊的用戶位置

### 常見時區

| 區域 | 時區 ID | UTC 偏移 |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## 過濾和查詢

日期欄位支援複雜過濾：

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### 檢查空日期欄位

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### 支援的運算符

| 運算符 | 用法 | 描述 |
|----------|-------|-------------|
| `EQ` | 與 dateRange | 日期與指定範圍重疊（任何交集） |
| `NE` | 與 dateRange | 日期不與範圍重疊 |
| `IS` | 與 `values: null` | 日期欄位為空（startDate 或 endDate 為 null） |
| `NOT` | 與 `values: null` | 日期欄位有值（兩個日期都不為 null） |

## 所需權限

| 操作 | 所需權限 |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## 錯誤回應

### 無效的日期格式
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
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
      "code": "NOT_FOUND"
    }
  }]
}
```


## 限制

- 不支援重複日期（使用自動化處理重複事件）
- 無法在沒有日期的情況下設定時間
- 沒有內建的工作日計算
- 日期範圍不會自動驗證結束 > 開始
- 最大精度為秒（不儲存毫秒）

## 相關資源

- [自訂欄位概述](/api/custom-fields/list-custom-fields) - 一般自訂欄位概念
- [自動化 API](/api/automations/index) - 創建基於日期的自動化