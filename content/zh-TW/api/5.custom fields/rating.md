---
title: 評分自訂欄位
description: 創建評分欄位以儲存具有可配置刻度和驗證的數值評分
---

評分自訂欄位允許您在記錄中儲存數值評分，並可配置最小值和最大值。它們非常適合用於績效評分、滿意度分數、優先級別或任何基於數值刻度的數據在您的項目中。

## 基本範例

創建一個具有默認 0-5 刻度的簡單評分欄位：

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## 進階範例

創建一個具有自訂刻度和描述的評分欄位：

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 評分欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `RATING` |
| `projectId` | String! | ✅ 是 | 此欄位將創建的項目 ID |
| `description` | String | 否 | 顯示給用戶的幫助文本 |
| `min` | Float | 否 | 最小評分值（無默認值） |
| `max` | Float | 否 | 最大評分值 |

## 設定評分值

要在記錄上設置或更新評分值：

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 評分自訂欄位的 ID |
| `value` | String! | ✅ 是 | 評分值作為字符串（在配置範圍內） |

## 使用評分值創建記錄

當使用評分值創建新記錄時：

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
    }
  }
}
```

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一識別碼 |
| `customField` | CustomField! | 自訂欄位定義 |
| `value` | Float | 儲存的評分值（通過 customField.value 訪問） |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

**注意**: 評分值實際上是通過 `customField.value.number` 在查詢中訪問的。

### CustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位的唯一識別碼 |
| `name` | String! | 評分欄位的顯示名稱 |
| `type` | CustomFieldType! | 始終是 `RATING` |
| `min` | Float | 允許的最小評分值 |
| `max` | Float | 允許的最大評分值 |
| `description` | String | 欄位的幫助文本 |

## 評分驗證

### 值約束
- 評分值必須是數字（浮點類型）
- 值必須在配置的最小/最大範圍內
- 如果未指定最小值，則沒有默認值
- 最大值是可選的，但建議使用

### 驗證規則
**重要**: 驗證僅在提交表單時發生，而不是在直接使用 `setTodoCustomField` 時。

- 輸入被解析為浮點數（使用表單時）
- 必須大於或等於最小值（使用表單時）
- 必須小於或等於最大值（使用表單時）
- `setTodoCustomField` 接受任何字符串值而不進行驗證

### 有效評分範例
對於最小值=1，最大值=5 的欄位：
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### 無效評分範例
對於最小值=1，最大值=5 的欄位：
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## 配置選項

### 評分刻度設置
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### 常見評分刻度
- **1-5 星**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 績效**: `min: 1, max: 10`
- **0-100 百分比**: `min: 0, max: 100`
- **自訂刻度**: 任何數值範圍

## 所需權限

自訂欄位操作遵循標準基於角色的權限：

| 行動 | 所需角色 |
|------|----------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**注意**: 所需的具體角色取決於您的項目的自訂角色配置和欄位級別權限。

## 錯誤回應

### 驗證錯誤（僅限表單）
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

**重要**: 評分值驗證（最小/最大約束）僅在提交表單時發生，而不是在直接使用 `setTodoCustomField` 時。

### 找不到自訂欄位
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

## 最佳實踐

### 刻度設計
- 在相似欄位中使用一致的評分刻度
- 考慮用戶的熟悉度（1-5 星，0-10 NPS）
- 設定適當的最小值（0 與 1）
- 為每個評分級別定義明確的意義

### 數據質量
- 在儲存之前驗證評分值
- 適當使用小數精度
- 考慮顯示目的的四捨五入
- 提供清晰的評分意義指導

### 用戶體驗
- 以視覺方式顯示評分刻度（星星、進度條）
- 顯示當前值和刻度限制
- 提供評分意義的上下文
- 考慮新記錄的默認值

## 常見用例

1. **績效管理**
   - 員工績效評分
   - 項目質量分數
   - 任務完成評分
   - 技能水平評估

2. **客戶反饋**
   - 滿意度評分
   - 產品質量分數
   - 服務體驗評分
   - 淨推薦值（NPS）

3. **優先級和重要性**
   - 任務優先級別
   - 緊急程度評分
   - 風險評估分數
   - 影響評分

4. **質量保證**
   - 代碼審查評分
   - 測試質量分數
   - 文檔質量
   - 流程遵循評分

## 整合功能

### 與自動化
- 根據評分閾值觸發行動
- 對低評分發送通知
- 為高評分創建後續任務
- 根據評分值路由工作

### 與查詢
- 計算記錄的平均評分
- 按評分範圍查找記錄
- 參考其他記錄的評分數據
- 聚合評分統計

### 與 Blue 前端
- 在表單上下文中自動範圍驗證
- 視覺評分輸入控件
- 實時驗證反饋
- 星星或滑塊輸入選項

## 活動追蹤

評分欄位的變更會自動追蹤：
- 舊的和新的評分值被記錄
- 活動顯示數值變更
- 所有評分更新的時間戳
- 用戶歸屬於變更

## 限制

- 僅支持數值
- 沒有內建的視覺評分顯示（星星等）
- 小數精度取決於數據庫配置
- 沒有評分元數據儲存（評論、上下文）
- 沒有自動評分聚合或統計
- 沒有內建的評分轉換功能
- **關鍵**: 最小/最大驗證僅在表單中有效，無法通過 `setTodoCustomField` 使用

## 相關資源

- [數字欄位](/api/5.custom%20fields/number) - 用於一般數值數據
- [百分比欄位](/api/5.custom%20fields/percent) - 用於百分比值
- [選擇欄位](/api/5.custom%20fields/select-single) - 用於離散選擇評分
- [自訂欄位概述](/api/5.custom%20fields/2.list-custom-fields) - 一般概念