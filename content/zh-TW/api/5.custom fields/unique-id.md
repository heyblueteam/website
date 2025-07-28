---
title: 唯一識別碼自訂欄位
description: 創建自動生成的唯一識別碼欄位，具有順序編號和自訂格式
---

唯一識別碼自訂欄位自動生成順序的唯一識別碼，適用於您的記錄。它們非常適合創建票證號碼、訂單 ID、發票號碼或您工作流程中的任何順序識別系統。

## 基本範例

創建一個簡單的唯一識別碼欄位，具有自動排序功能：

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## 進階範例

創建一個具有前綴和零填充的格式化唯一識別碼欄位：

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## 輸入參數

### CreateCustomFieldInput (UNIQUE_ID)

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 唯一識別碼欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `UNIQUE_ID` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |
| `useSequenceUniqueId` | Boolean | 否 | 啟用自動排序（默認：假） |
| `prefix` | String | 否 | 生成 ID 的文本前綴（例如，“TASK-”） |
| `sequenceDigits` | Int | 否 | 零填充的位數 |
| `sequenceStartingNumber` | Int | 否 | 順序的起始數字 |

## 配置選項

### 自動排序 (`useSequenceUniqueId`)
- **true**: 當記錄被創建時，自動生成順序 ID
- **false** 或 **undefined**: 需要手動輸入（類似於文本欄位）

### 前綴 (`prefix`)
- 添加到所有生成 ID 的可選文本前綴
- 例子：“TASK-”、“ORD-”、“BUG-”、“REQ-”
- 沒有長度限制，但請保持合理以便顯示

### 順序位數 (`sequenceDigits`)
- 用於零填充順序號碼的位數
- 例子：`sequenceDigits: 3` 生成 `001`、`002`、`003`
- 如果未指定，則不應用填充

### 起始數字 (`sequenceStartingNumber`)
- 順序中的第一個數字
- 例子：`sequenceStartingNumber: 1000` 從 1000 開始，1001，1002...
- 如果未指定，則從 1 開始（默認行為）

## 生成的 ID 格式

最終的 ID 格式結合了所有配置選項：

```
{prefix}{paddedSequenceNumber}
```

### 格式範例

| 配置 | 生成的 ID |
|------|-----------|
| 無選項 | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## 讀取唯一識別碼值

### 查詢具有唯一識別碼的記錄
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### 回應格式
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## 自動 ID 生成

### 何時生成 ID
- **記錄創建**：當新記錄被創建時，自動分配 ID
- **欄位添加**：當將 UNIQUE_ID 欄位添加到現有記錄時，會排隊一個後台作業（工作者實現待定）
- **後台處理**：新記錄的 ID 生成通過數據庫觸發器同步進行

### 生成過程
1. **觸發**：創建新記錄或添加 UNIQUE_ID 欄位
2. **順序查找**：系統查找下一個可用的順序號碼
3. **ID 分配**：順序號碼分配給記錄
4. **計數器更新**：順序計數器為未來的記錄遞增
5. **格式化**：ID 在顯示時用前綴和填充格式化

### 唯一性保證
- **數據庫約束**：每個欄位內的順序 ID 具有唯一約束
- **原子操作**：順序生成使用數據庫鎖以防止重複
- **項目範圍**：每個項目的序列是獨立的
- **競爭條件保護**：安全處理並發請求

## 手動模式與自動模式

### 自動模式 (`useSequenceUniqueId: true`)
- ID 通過數據庫觸發器自動生成
- 保證順序編號
- 原子順序生成防止重複
- 格式化 ID 結合前綴 + 填充的順序號碼

### 手動模式 (`useSequenceUniqueId: false` 或 `undefined`)
- 像常規文本欄位一樣運作
- 用戶可以通過 `setTodoCustomField` 與 `text` 參數輸入自訂值
- 無自動生成
- 除了數據庫約束外，無唯一性強制

## 設置手動值（僅限手動模式）

當 `useSequenceUniqueId` 為假時，您可以手動設置值：

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## 回應欄位

### TodoCustomField 回應 (UNIQUE_ID)

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一識別碼 |
| `customField` | CustomField! | 自訂欄位定義 |
| `sequenceId` | Int | 生成的順序號碼（為 UNIQUE_ID 欄位填充） |
| `text` | String | 格式化的文本值（結合前綴 + 填充的順序） |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後更新的時間 |

### CustomField 回應 (UNIQUE_ID)

| 欄位 | 類型 | 描述 |
|------|------|------|
| `useSequenceUniqueId` | Boolean | 是否啟用自動排序 |
| `prefix` | String | 生成 ID 的文本前綴 |
| `sequenceDigits` | Int | 用於零填充的位數 |
| `sequenceStartingNumber` | Int | 順序的起始數字 |

## 所需權限

| 操作 | 所需權限 |
|------|----------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## 錯誤回應

### 欄位配置錯誤
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### 權限錯誤
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## 重要說明

### 自動生成的 ID
- **唯讀**：自動生成的 ID 不能手動編輯
- **永久**：一旦分配，順序 ID 不會改變
- **按時間順序**：ID 反映創建順序
- **範圍**：序列在每個項目中是獨立的

### 性能考量
- 新記錄的 ID 生成是通過數據庫觸發器同步進行的
- 順序生成使用 `FOR UPDATE` 鎖進行原子操作
- 存在後台作業系統，但工作者實現待定
- 考慮高容量項目的序列起始數字

### 遷移和更新
- 將自動排序添加到現有記錄會排隊後台作業（工作者待定）
- 更改順序設置僅影響未來的記錄
- 當配置更新時，現有 ID 保持不變
- 順序計數器從當前最大值繼續

## 最佳實踐

### 配置設計
- 選擇不會與其他系統衝突的描述性前綴
- 根據預期的數量使用適當的位數填充
- 設置合理的起始數字以避免衝突
- 在部署之前使用示例數據測試配置

### 前綴指南
- 保持前綴簡短且易記（2-5 個字符）
- 使用大寫字母以保持一致性
- 包含分隔符（連字符、下劃線）以提高可讀性
- 避免可能在 URL 或系統中引起問題的特殊字符

### 順序規劃
- 估算您的記錄量以選擇適當的位數填充
- 在設置起始數字時考慮未來增長
- 為不同的記錄類型規劃不同的序列範圍
- 為團隊參考記錄您的 ID 計劃

## 常見用例

1. **支持系統**
   - 票證號碼：`TICK-001`、`TICK-002`
   - 案例 ID：`CASE-2024-001`
   - 支持請求：`SUP-001`

2. **項目管理**
   - 任務 ID：`TASK-001`、`TASK-002`
   - 衝刺項目：`SPRINT-001`
   - 可交付成果編號：`DEL-001`

3. **業務運營**
   - 訂單號碼：`ORD-2024-001`
   - 發票 ID：`INV-001`
   - 購買訂單：`PO-001`

4. **質量管理**
   - 錯誤報告：`BUG-001`
   - 測試案例 ID：`TEST-001`
   - 審查編號：`REV-001`

## 整合功能

### 與自動化
- 當分配唯一 ID 時觸發操作
- 在自動化規則中使用 ID 模式
- 在電子郵件模板和通知中引用 ID

### 與查詢
- 參考來自其他記錄的唯一 ID
- 通過唯一 ID 查找記錄
- 顯示相關記錄識別碼

### 與報告
- 按 ID 模式分組和過濾
- 跟踪 ID 分配趨勢
- 監控序列使用情況和缺口

## 限制

- **僅順序**：ID 按時間順序分配
- **無間隙**：刪除的記錄在序列中留下間隙
- **無重用**：順序號碼永遠不會重用
- **項目範圍**：無法跨項目共享序列
- **格式約束**：格式選項有限
- **無批量更新**：無法批量更新現有的順序 ID
- **無自訂邏輯**：無法實施自訂 ID 生成規則

## 相關資源

- [文本欄位](/api/custom-fields/text-single) - 用於手動文本識別碼
- [數字欄位](/api/custom-fields/number) - 用於數字序列
- [自訂欄位概述](/api/custom-fields/2.list-custom-fields) - 一般概念
- [自動化](/api/automations) - 用於基於 ID 的自動化規則