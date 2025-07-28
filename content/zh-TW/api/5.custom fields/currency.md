---
title: 貨幣自訂欄位
description: 創建貨幣欄位以追蹤金額值，並提供適當的格式和驗證
---

貨幣自訂欄位允許您儲存和管理與貨幣代碼相關的金額值。該欄位支持72種不同的貨幣，包括主要法定貨幣和加密貨幣，並具有自動格式化和可選的最小/最大限制。

## 基本範例

創建一個簡單的貨幣欄位：

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## 進階範例

創建一個具有驗證限制的貨幣欄位：

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 貨幣欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `CURRENCY` |
| `currency` | String | 否 | 預設貨幣代碼（3 字母 ISO 代碼） |
| `min` | Float | 否 | 允許的最小值（儲存但不在更新時強制執行） |
| `max` | Float | 否 | 允許的最大值（儲存但不在更新時強制執行） |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

**注意**：專案上下文是根據您的身份驗證自動確定的。您必須有權訪問您正在創建欄位的專案。

## 設定貨幣值

要在記錄上設置或更新貨幣值：

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 貨幣自訂欄位的 ID |
| `number` | Float! | ✅ 是 | 金額 |
| `currency` | String! | ✅ 是 | 3 字母貨幣代碼 |

## 使用貨幣值創建記錄

當使用貨幣值創建新記錄時：

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### 創建的輸入格式

在創建記錄時，貨幣值的傳遞方式不同：

| 參數 | 類型 | 描述 |
|------|------|------|
| `customFieldId` | String! | 貨幣欄位的 ID |
| `value` | String! | 以字串形式表示的金額（例如："1500.50"） |
| `currency` | String! | 3 字母貨幣代碼 |

## 支持的貨幣

Blue 支持 72 種貨幣，包括 70 種法定貨幣和 2 種加密貨幣：

### 法定貨幣

#### 美洲
| 貨幣 | 代碼 | 名稱 |
|------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | 委內瑞拉主權玻利瓦爾 |
| Bolivian Boliviano | `BOB` | Bolivian Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### 歐洲
| 貨幣 | 代碼 | 名稱 |
|------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| 挪威克朗 | `NOK` | 挪威克朗 |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### 亞太地區
| 貨幣 | 代碼 | 名稱 |
|------|------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### 中東與非洲
| 貨幣 | 代碼 | 名稱 |
|------|------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### 加密貨幣
| 貨幣 | 代碼 |
|------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 |
| `number` | Float | 金額 |
| `currency` | String | 3 字母貨幣代碼 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

## 貨幣格式化

系統根據區域自動格式化貨幣值：

- **符號位置**：正確放置貨幣符號（在前/在後）
- **小數分隔符**：使用區域特定的分隔符（. 或 ,）
- **千位分隔符**：應用適當的分組
- **小數位數**：根據金額顯示 0-2 位小數
- **特殊處理**：美元/加元顯示貨幣代碼前綴以便於理解

### 格式化範例

| 值 | 貨幣 | 顯示 |
|----|------|------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## 驗證

### 金額驗證
- 必須是有效的數字
- 最小/最大限制與欄位定義一起儲存，但在值更新時不強制執行
- 支持最多 2 位小數的顯示（完整精度在內部儲存）

### 貨幣代碼驗證
- 必須是 72 種支持的貨幣代碼之一
- 大小寫敏感（使用大寫）
- 無效代碼會返回錯誤

## 整合功能

### 公式
貨幣欄位可以用於公式自訂欄位進行計算：
- 總和多個貨幣欄位
- 計算百分比
- 執行算術運算

### 貨幣轉換
使用貨幣轉換欄位自動在貨幣之間轉換（請參見 [貨幣轉換欄位](/api/custom-fields/currency-conversion)）

### 自動化
貨幣值可以根據以下條件觸發自動化：
- 金額閾值
- 貨幣類型
- 值變更

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**注意**：雖然任何專案成員都可以創建自訂欄位，但設置值的能力取決於為每個欄位配置的基於角色的權限。

## 錯誤回應

### 無效的貨幣值
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

當以下情況發生時會出現此錯誤：
- 貨幣代碼不是 72 種支持的代碼之一
- 數字格式無效
- 值無法正確解析

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

### 貨幣選擇
- 設置與您的主要市場匹配的預設貨幣
- 一致使用 ISO 4217 貨幣代碼
- 考慮用戶位置以選擇預設值

### 值限制
- 設置合理的最小/最大值以防止數據輸入錯誤
- 對於不應為負數的欄位，使用 0 作為最小值
- 設置最大值時考慮您的使用案例

### 多貨幣專案
- 使用一致的基礎貨幣進行報告
- 實施貨幣轉換欄位以自動轉換
- 記錄每個欄位應使用的貨幣

## 常見使用案例

1. **專案預算**
   - 專案預算追蹤
   - 成本估算
   - 開支追蹤

2. **銷售與交易**
   - 交易值
   - 合約金額
   - 收入追蹤

3. **財務規劃**
   - 投資金額
   - 融資輪次
   - 財務目標

4. **國際業務**
   - 多貨幣定價
   - 外匯追蹤
   - 跨境交易

## 限制

- 顯示最多 2 位小數（儘管儲存了更多精度）
- 標準貨幣欄位中沒有內建的貨幣轉換
- 不能在單個欄位值中混合貨幣
- 沒有自動匯率更新（使用貨幣轉換進行此操作）
- 貨幣符號不可自定義

## 相關資源

- [貨幣轉換欄位](/api/custom-fields/currency-conversion) - 自動貨幣轉換
- [數字欄位](/api/custom-fields/number) - 用於非貨幣數值
- [公式欄位](/api/custom-fields/formula) - 使用貨幣值進行計算
- [列表自訂欄位](/api/custom-fields/list-custom-fields) - 查詢專案中的所有自訂欄位