---
title: 公式自訂欄位
description: 創建根據其他數據自動計算值的計算欄位
---

公式自訂欄位用於 Blue 中的圖表和儀表板計算。它們定義在自訂欄位數據上運行的聚合函數（SUM、AVERAGE、COUNT 等），以在圖表中顯示計算的指標。公式不是在單個待辦事項層級計算的，而是聚合多條記錄的數據以便於視覺化。

## 基本範例

為圖表計算創建一個公式欄位：

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## 進階範例

創建一個具有複雜計算的貨幣公式：

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 公式欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `FORMULA` |
| `projectId` | String! | ✅ 是 | 此欄位將被創建的專案 ID |
| `formula` | JSON | 否 | 用於圖表計算的公式定義 |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

### 公式結構

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## 支援的函數

### 圖表聚合函數

公式欄位支援以下聚合函數以進行圖表計算：

| 函數 | 描述 | ChartFunction 枚舉 |
|------|------|---------------------|
| `SUM` | 所有值的總和 | `SUM` |
| `AVERAGE` | 數值的平均值 | `AVERAGE` |
| `AVERAGEA` | 排除零和空值的平均值 | `AVERAGEA` |
| `COUNT` | 值的計數 | `COUNT` |
| `COUNTA` | 排除零和空值的計數 | `COUNTA` |
| `MAX` | 最大值 | `MAX` |
| `MIN` | 最小值 | `MIN` |

**注意**：這些函數用於 `display.function` 欄位，並在圖表視覺化中對聚合數據進行操作。不支援複雜的數學表達式或欄位級計算。

## 顯示類型

### 數字顯示

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

結果： `1250.75`

### 貨幣顯示

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

結果： `$1,250.75`

### 百分比顯示

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

結果： `87.5%`

## 編輯公式欄位

更新現有的公式欄位：

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## 公式處理

### 圖表計算上下文

公式欄位在圖表區段和儀表板的上下文中進行處理：
- 當圖表被渲染或更新時進行計算
- 結果以十進制值存儲在 `ChartSegment.formulaResult` 中
- 處理通過名為 'formula' 的專用 BullMQ 隊列進行
- 更新會發布給儀表板訂閱者以實現實時更新

### 顯示格式

`getFormulaDisplayValue` 函數根據顯示類型格式化計算結果：
- **NUMBER**：顯示為普通數字，並可選擇精度
- **PERCENTAGE**：添加 % 後綴，並可選擇精度  
- **CURRENCY**：使用指定的貨幣代碼格式化

## 公式結果存儲

結果存儲在 `formulaResult` 欄位中：

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 公式欄位定義 |
| `number` | Float | 計算的數值結果 |
| `formulaResult` | JSON | 帶有顯示格式的完整結果 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後計算的時間 |

## 數據上下文

### 圖表數據來源

公式欄位在圖表數據來源上下文中運作：
- 公式聚合專案中待辦事項的自訂欄位值
- `display.function` 中指定的聚合函數決定計算
- 結果使用 SQL 聚合函數（avg、sum、count 等）計算
- 計算在數據庫層級進行以提高效率

## 常見公式範例

### 總預算（圖表顯示）

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### 平均分數（圖表顯示）

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### 任務計數（圖表顯示）

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## 所需權限

自訂欄位操作遵循標準的基於角色的權限：

| 操作 | 所需角色 |
|------|----------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**注意**：所需的具體角色取決於您專案的自訂角色配置。沒有像 CUSTOM_FIELDS_CREATE 這樣的特殊權限常量。

## 錯誤處理

### 驗證錯誤
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

### 公式設計
- 使用清晰、描述性的名稱來命名公式欄位
- 添加描述以解釋計算邏輯
- 在部署之前使用示例數據測試公式
- 保持公式簡單且易於閱讀

### 性能優化
- 避免深度嵌套的公式依賴
- 使用具體的欄位引用，而不是通配符
- 考慮對複雜計算的緩存策略
- 在大型專案中監控公式性能

### 數據質量
- 在公式中使用之前驗證源數據
- 適當處理空值或 null 值
- 為顯示類型使用適當的精度
- 考慮計算中的邊緣情況

## 常見用例

1. **財務追蹤**
   - 預算計算
   - 利潤/損失報表
   - 成本分析
   - 收入預測

2. **專案管理**
   - 完成百分比
   - 資源利用率
   - 時間線計算
   - 性能指標

3. **質量控制**
   - 平均分數
   - 通過/失敗率
   - 質量指標
   - 合規追蹤

4. **商業智慧**
   - KPI 計算
   - 趨勢分析
   - 比較指標
   - 儀表板值

## 限制

- 公式僅用於圖表/儀表板聚合，不進行待辦事項層級計算
- 限制於七個支援的聚合函數（SUM、AVERAGE 等）
- 不支援複雜的數學表達式或欄位到欄位的計算
- 不能在單個公式中引用多個欄位
- 結果僅在圖表和儀表板中可見
- `logic` 欄位僅用於顯示文本，而非實際計算邏輯

## 相關資源

- [數字欄位](/api/5.custom%20fields/number) - 用於靜態數值
- [貨幣欄位](/api/5.custom%20fields/currency) - 用於貨幣值
- [參考欄位](/api/5.custom%20fields/reference) - 用於跨專案數據
- [查找欄位](/api/5.custom%20fields/lookup) - 用於聚合數據
- [自訂欄位概述](/api/5.custom%20fields/2.list-custom-fields) - 一般概念