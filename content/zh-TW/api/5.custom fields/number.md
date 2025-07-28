---
title: 數字自定義欄位
description: 創建數字欄位以存儲數值，並可選擇性設置最小/最大約束和前綴格式
---

數字自定義欄位允許您為記錄存儲數值。它們支持驗證約束、小數精度，並可用於數量、分數、測量或任何不需要特殊格式的數據。

## 基本範例

創建一個簡單的數字欄位：

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 進階範例

創建一個帶有約束和前綴的數字欄位：

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 數字欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `NUMBER` |
| `projectId` | String! | ✅ 是 | 創建欄位的項目 ID |
| `min` | Float | 否 | 最小值約束（僅供 UI 指導） |
| `max` | Float | 否 | 最大值約束（僅供 UI 指導） |
| `prefix` | String | 否 | 顯示前綴（例如，“#”，“~”，“$”） |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

## 設置數字值

數字欄位存儲小數值並可選擇性進行驗證：

### 簡單數字值

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### 整數值

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 數字自定義欄位的 ID |
| `number` | Float | 否 | 要存儲的數值 |

## 值約束

### 最小/最大約束（UI 指導）

**重要**：最小/最大約束被存儲但不會在伺服器端強制執行。它們作為前端應用程序的 UI 指導。

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**需要客戶端驗證**：前端應用程序必須實現驗證邏輯以強制執行最小/最大約束。

### 支持的值類型

| 類型 | 示例 | 描述 |
|------|------|------|
| Integer | `42` | 整數 |
| Decimal | `42.5` | 帶小數的數字 |
| Negative | `-10` | 負值（如果沒有最小約束） |
| Zero | `0` | 零值 |

**注意**：最小/最大約束不會在伺服器端進行驗證。超出指定範圍的值將被接受並存儲。

## 使用數字值創建記錄

當使用數字值創建新記錄時：

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### 支持的輸入格式

在創建記錄時，請在自定義欄位數組中使用 `number` 參數（而不是 `value`）：

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自定義欄位定義 |
| `number` | Float | 數值 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

### CustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位定義的唯一標識符 |
| `name` | String! | 欄位的顯示名稱 |
| `type` | CustomFieldType! | 始終是 `NUMBER` |
| `min` | Float | 允許的最小值 |
| `max` | Float | 允許的最大值 |
| `prefix` | String | 顯示前綴 |
| `description` | String | 幫助文本 |

**注意**：如果數字值未設置，則 `number` 欄位將是 `null`。

## 過濾和查詢

數字欄位支持全面的數字過濾：

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
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

### 支持的運算符

| 運算符 | 描述 | 示例 |
|--------|------|------|
| `EQ` | 等於 | `number = 42` |
| `NE` | 不等於 | `number ≠ 42` |
| `GT` | 大於 | `number > 42` |
| `GTE` | 大於或等於 | `number ≥ 42` |
| `LT` | 小於 | `number < 42` |
| `LTE` | 小於或等於 | `number ≤ 42` |
| `IN` | 在數組中 | `number in [1, 2, 3]` |
| `NIN` | 不在數組中 | `number not in [1, 2, 3]` |
| `IS` | 為空/不為空 | `number is null` |

### 範圍過濾

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## 顯示格式

### 帶前綴

如果設置了前綴，則會顯示：

| 值 | 前綴 | 顯示 |
|----|------|------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### 小數精度

數字保持其小數精度：

| 輸入 | 存儲 | 顯示 |
|------|------|------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## 錯誤回應

### 無效的數字格式
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**注意**：最小/最大驗證錯誤不會在伺服器端發生。約束驗證必須在您的前端應用程序中實現。

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

### 約束設計
- 設置現實的最小/最大值以供 UI 指導
- 實施客戶端驗證以強制執行約束
- 使用約束在表單中提供用戶反饋
- 考慮負值是否對您的用例有效

### 值精度
- 根據需要使用適當的小數精度
- 考慮顯示目的的四捨五入
- 在相關欄位之間保持精度一致

### 顯示增強
- 使用有意義的前綴以提供上下文
- 考慮在欄位名稱中使用單位（例如，“重量（公斤）”）
- 提供清晰的驗證規則描述

## 常見用例

1. **評分系統**
   - 性能評級
   - 質量分數
   - 優先級別
   - 客戶滿意度評級

2. **測量**
   - 數量和金額
   - 尺寸和大小
   - 持續時間（以數字格式）
   - 容量和限制

3. **業務指標**
   - 收入數字
   - 轉換率
   - 預算分配
   - 目標數字

4. **技術數據**
   - 版本號
   - 配置值
   - 性能指標
   - 閾值設置

## 集成功能

### 與圖表和儀表板
- 在圖表計算中使用數字欄位
- 創建數字可視化
- 隨時間跟踪趨勢

### 與自動化
- 根據數字閾值觸發操作
- 根據數字變更更新相關欄位
- 對特定值發送通知

### 與查詢
- 從相關記錄聚合數字
- 計算總數和平均值
- 找到關係中的最小/最大值

### 與圖表
- 創建數字可視化
- 隨時間跟踪趨勢
- 比較記錄之間的值

## 限制

- **不進行伺服器端的最小/最大約束驗證**
- **需要客戶端驗證**以強制執行約束
- 無內置的貨幣格式（請使用貨幣類型）
- 無自動百分比符號（請使用百分比類型）
- 無單位轉換功能
- 小數精度受數據庫小數類型的限制
- 欄位本身不支持數學公式評估

## 相關資源

- [自定義欄位概述](/api/custom-fields/1.index) - 一般自定義欄位概念
- [貨幣自定義欄位](/api/custom-fields/currency) - 用於貨幣值
- [百分比自定義欄位](/api/custom-fields/percent) - 用於百分比值
- [自動化 API](/api/automations/1.index) - 創建基於數字的自動化