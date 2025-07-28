---
title: 貨幣轉換自訂欄位
description: 創建自動使用實時匯率轉換貨幣值的欄位
---

貨幣轉換自訂欄位自動將來源貨幣欄位的值轉換為不同的目標貨幣，使用實時匯率。當來源貨幣值變更時，這些欄位會自動更新。

轉換匯率由 [Frankfurter API](https://github.com/hakanensari/frankfurter) 提供，這是一個開源服務，跟踪由 [歐洲中央銀行](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html) 發布的參考匯率。這確保了您國際業務需求的準確、可靠和最新的貨幣轉換。

## 基本範例

創建一個簡單的貨幣轉換欄位：

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## 進階範例

創建一個具有特定日期的轉換欄位，以獲取歷史匯率：

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## 完整設置過程

設置貨幣轉換欄位需要三個步驟：

### 步驟 1：創建來源貨幣欄位

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### 步驟 2：創建貨幣轉換欄位

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### 步驟 3：創建轉換選項

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 轉換欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | 否 | 要轉換的來源貨幣欄位的 ID |
| `conversionDateType` | String | 否 | 匯率的日期策略（見下文） |
| `conversionDate` | String | 否 | 用於轉換的日期字串（基於 conversionDateType） |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：自訂欄位會根據用戶當前的專案上下文自動與專案關聯。無需 `projectId` 參數。

### 轉換日期類型

| 類型 | 描述 | conversionDate 參數 |
|------|------|---------------------|
| `currentDate` | 使用實時匯率 | 不需要 |
| `specificDate` | 使用固定日期的匯率 | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | 使用來自另一個欄位的日期 | "todoDueDate" or DATE field ID |

## 創建轉換選項

轉換選項定義可以轉換的貨幣對：

### CreateCustomFieldOptionInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `customFieldId` | String! | ✅ 是 | 貨幣轉換欄位的 ID |
| `title` | String! | ✅ 是 | 此轉換選項的顯示名稱 |
| `currencyConversionFrom` | String! | ✅ 是 | 來源貨幣代碼或「任何」 |
| `currencyConversionTo` | String! | ✅ 是 | 目標貨幣代碼 |

### 使用「任何」作為來源

特殊值「任何」作為 `currencyConversionFrom` 創建一個備用選項：

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

當找不到特定貨幣對匹配時，將使用此選項。

## 自動轉換的工作原理

1. **值更新**：當在來源貨幣欄位中設置值時
2. **選項匹配**：系統根據來源貨幣找到匹配的轉換選項
3. **匯率獲取**：從 Frankfurter API 獲取匯率
4. **計算**：將來源金額乘以匯率
5. **儲存**：將轉換後的值與目標貨幣代碼一起保存

### 範例流程

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## 基於日期的轉換

### 使用當前日期

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

每次來源值變更時，轉換會使用當前匯率更新。

### 使用特定日期

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

始終使用指定日期的匯率。

### 使用來自欄位的日期

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

使用來自另一個欄位的日期（無論是待辦事項到期日還是日期自訂欄位）。

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 轉換欄位的定義 |
| `number` | Float | 轉換後的金額 |
| `currency` | String | 目標貨幣代碼 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後更新的時間 |

## 匯率來源

Blue 使用 **Frankfurter API** 來獲取匯率：
- 由歐洲中央銀行主辦的開源 API
- 每日更新官方匯率
- 支持自 1999 年以來的歷史匯率
- 免費且可靠，適用於商業用途

## 錯誤處理

### 轉換失敗

當轉換失敗（API 錯誤、無效貨幣等）時：
- 轉換後的值設置為 `0`
- 目標貨幣仍然被存儲
- 不會向用戶拋出錯誤

### 常見場景

| 場景 | 結果 |
|------|------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| 沒有匹配的選項 | Uses "Any" option if available |
| Missing source value | 沒有執行轉換 |

## 所需權限

自訂欄位管理需要專案級別的訪問：

| 角色 | 可以創建/更新欄位 |
|------|--------------------|
| `OWNER` | ✅ 是 |
| `ADMIN` | ✅ 是 |
| `MEMBER` | ❌ 否 |
| `CLIENT` | ❌ 否 |

轉換後值的查看權限遵循標準記錄訪問規則。

## 最佳實踐

### 選項配置
- 為常見轉換創建特定的貨幣對
- 添加「任何」備用選項以提高靈活性
- 為選項使用描述性標題

### 日期策略選擇
- 使用 `currentDate` 進行實時財務跟踪
- 使用 `specificDate` 進行歷史報告
- 使用 `fromDateField` 進行交易特定的匯率

### 性能考量
- 多個轉換欄位並行更新
- 只有當來源值變更時才會進行 API 調用
- 相同貨幣的轉換跳過 API 調用

## 常見用例

1. **多貨幣專案**
   - 以當地貨幣跟踪專案成本
   - 以公司貨幣報告總預算
   - 比較不同地區的值

2. **國際銷售**
   - 將交易值轉換為報告貨幣
   - 以多種貨幣跟踪收入
   - 對已結束交易進行歷史轉換

3. **財務報告**
   - 期末貨幣轉換
   - 綜合財務報表
   - 當地貨幣的預算與實際對比

4. **合同管理**
   - 在簽署日期轉換合同值
   - 以多種貨幣跟踪付款計劃
   - 貨幣風險評估

## 限制

- 不支持加密貨幣轉換
- 無法手動設置轉換值（始終計算）
- 所有轉換金額固定為 2 位小數精度
- 不支持自訂匯率
- 不會緩存匯率（每次轉換都會進行新的 API 調用）
- 依賴於 Frankfurter API 的可用性

## 相關資源

- [貨幣欄位](/api/custom-fields/currency) - 轉換的來源欄位
- [日期欄位](/api/custom-fields/date) - 用於基於日期的轉換
- [公式欄位](/api/custom-fields/formula) - 替代計算
- [自訂欄位概述](/custom-fields/list-custom-fields) - 一般概念