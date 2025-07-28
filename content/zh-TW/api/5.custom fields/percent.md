---
title: 百分比自訂欄位
description: 創建百分比欄位以儲存數值，自動處理 % 符號和顯示格式
---

百分比自訂欄位允許您為記錄儲存百分比值。它們自動處理輸入和顯示的 % 符號，同時在內部儲存原始數值。非常適合完成率、成功率或任何基於百分比的指標。

## 基本範例

創建一個簡單的百分比欄位：

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## 進階範例

創建一個帶有描述的百分比欄位：

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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
| `name` | String! | ✅ 是 | 百分比欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `PERCENT` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：項目上下文是從您的身份驗證標頭自動確定的。無需 `projectId` 參數。

**注意**：PERCENT 欄位不支持最小/最大約束或像 NUMBER 欄位那樣的前綴格式。

## 設定百分比值

百分比欄位儲存數值，自動處理 % 符號：

### 帶有百分比符號

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### 直接數值

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 百分比自訂欄位的 ID |
| `number` | Float | 否 | 數值百分比 (例如，75.5 代表 75.5%) |

## 值儲存和顯示

### 儲存格式
- **內部儲存**：原始數值 (例如，75.5)
- **數據庫**：儲存為 `Decimal` 在 `number` 欄位中
- **GraphQL**：返回為 `Float` 類型

### 顯示格式
- **用戶界面**：客戶端應用程序必須附加 % 符號 (例如，"75.5%")
- **圖表**：當輸出類型為 PERCENTAGE 時顯示 % 符號
- **API 響應**：原始數值不帶 % 符號 (例如，75.5)

## 創建帶有百分比值的記錄

當創建帶有百分比值的新記錄時：

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### 支援的輸入格式

| 格式 | 範例 | 結果 |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**注意**：% 符號會自動從輸入中剝離，並在顯示時重新添加。

## 查詢百分比值

當查詢帶有百分比自訂欄位的記錄時，通過 `customField.value.number` 路徑訪問值：

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

響應將包括百分比作為原始數字：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## 響應欄位

### TodoCustomField 響應

| 欄位 | 類型 | 描述 |
|-------|------|-------------|
| `id` | ID! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 (包含百分比值) |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

**重要**：百分比值通過 `customField.value.number` 欄位訪問。儲存的值不包含 % 符號，必須由客戶端應用程序添加以進行顯示。

## 過濾和查詢

百分比欄位支持與 NUMBER 欄位相同的過濾：

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### 支援的運算符

| 運算符 | 描述 | 範例 |
|----------|-------------|---------|
| `EQ` | 等於 | `percentage = 75` |
| `NE` | 不等於 | `percentage ≠ 75` |
| `GT` | 大於 | `percentage > 75` |
| `GTE` | 大於或等於 | `percentage ≥ 75` |
| `LT` | 小於 | `percentage < 75` |
| `LTE` | 小於或等於 | `percentage ≤ 75` |
| `IN` | 列表中的值 | `percentage in [50, 75, 100]` |
| `NIN` | 列表中不包含的值 | `percentage not in [0, 25]` |
| `IS` | 檢查是否為空值，使用 `values: null` | `percentage is null` |
| `NOT` | 檢查是否不為空值，使用 `values: null` | `percentage is not null` |

### 範圍過濾

對於範圍過濾，使用多個運算符：

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## 百分比值範圍

### 常見範圍

| 範圍 | 描述 | 使用案例 |
|-------|-------------|----------|
| `0-100` | 標準百分比 | Completion rates, success rates |
| `0-∞` | 無限制百分比 | Growth rates, performance metrics |
| `-∞-∞` | 任何值 | Change rates, variance |

### 範例值

| 輸入 | 儲存 | 顯示 |
|-------|--------|---------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## 圖表聚合

百分比欄位支持在儀表板圖表和報告中進行聚合。可用的函數包括：

- `AVERAGE` - 平均百分比值
- `COUNT` - 擁有值的記錄數量
- `MIN` - 最低百分比值
- `MAX` - 最高百分比值 
- `SUM` - 所有百分比值的總和

這些聚合在創建圖表和儀表板時可用，而不是在直接的 GraphQL 查詢中。

## 所需權限

| 操作 | 所需權限 |
|--------|-------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## 錯誤響應

### 無效的百分比格式
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 不是數字
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 最佳實踐

### 值輸入
- 允許用戶輸入帶或不帶 % 符號
- 驗證合理的範圍以符合您的使用案例
- 提供清晰的上下文，說明 100% 代表什麼

### 顯示
- 在用戶界面中始終顯示 % 符號
- 使用適當的小數精度
- 考慮對範圍進行顏色編碼 (紅色/黃色/綠色)

### 數據解釋
- 記錄 100% 在您的上下文中意味著什麼
- 適當處理超過 100% 的值
- 考慮負值是否有效

## 常見使用案例

1. **項目管理**
   - 任務完成率
   - 項目進度
   - 資源利用率
   - 衝刺速度

2. **性能追蹤**
   - 成功率
   - 錯誤率
   - 效率指標
   - 質量分數

3. **財務指標**
   - 增長率
   - 利潤率
   - 折扣金額
   - 變化百分比

4. **分析**
   - 轉換率
   - 點擊率
   - 參與指標
   - 性能指標

## 整合功能

### 使用公式
- 在計算中引用 PERCENT 欄位
- 公式輸出中的自動 % 符號格式
- 與其他數值欄位結合

### 使用自動化
- 根據百分比閾值觸發操作
- 發送里程碑百分比的通知
- 根據完成率更新狀態

### 使用查詢
- 從相關記錄聚合百分比
- 計算平均成功率
- 找出表現最佳/最差的項目

### 使用圖表
- 創建基於百分比的可視化
- 隨時間跟踪進度
- 比較性能指標

## 與 NUMBER 欄位的區別

### 有什麼不同
- **輸入處理**：自動剝離 % 符號
- **顯示**：自動添加 % 符號
- **約束**：不進行最小/最大驗證
- **格式**：不支持前綴

### 有什麼相同
- **儲存**：相同的數據庫欄位和類型
- **過濾**：相同的查詢運算符
- **聚合**：相同的聚合函數
- **權限**：相同的權限模型

## 限制

- 無最小/最大值約束
- 無前綴格式選項
- 無自動驗證 0-100% 範圍
- 不支持百分比格式之間的轉換 (例如，0.75 ↔ 75%)
- 允許超過 100% 的值

## 相關資源

- [自訂欄位概述](/api/custom-fields/list-custom-fields) - 一般自訂欄位概念
- [數字自訂欄位](/api/custom-fields/number) - 用於原始數值
- [自動化 API](/api/automations/index) - 創建基於百分比的自動化