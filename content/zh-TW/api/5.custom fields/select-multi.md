---
title: 多選自定義欄位
description: 創建多選欄位以允許用戶從預定義列表中選擇多個選項
---

多選自定義欄位允許用戶從預定義列表中選擇多個選項。它們非常適合用於類別、標籤、技能、功能或任何需要從受控選項集中進行多重選擇的情況。

## 基本範例

創建一個簡單的多選欄位：

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 進階範例

創建一個多選欄位，然後單獨添加選項：

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 多選欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `SELECT_MULTI` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |
| `projectId` | String! | ✅ 是 | 此欄位的項目 ID |

### CreateCustomFieldOptionInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `customFieldId` | String! | ✅ 是 | 自定義欄位的 ID |
| `title` | String! | ✅ 是 | 選項的顯示文本 |
| `color` | String | 否 | 選項的顏色（任何字符串） |
| `position` | Float | 否 | 選項的排序順序 |

## 向現有欄位添加選項

向現有的多選欄位添加新選項：

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## 設定多選值

要在記錄上設置多個選定選項：

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄 ID |
| `customFieldId` | String! | ✅ 是 | 多選自定義欄位的 ID |
| `customFieldOptionIds` | [String!] | ✅ 是 | 要選擇的選項 ID 陣列 |

## 使用多選值創建記錄

在創建帶有多選值的新記錄時：

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
    }
  }
}
```

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自定義欄位定義 |
| `selectedOptions` | [CustomFieldOption!] | 選定選項的陣列 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

### CustomFieldOption 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 選項的唯一標識符 |
| `title` | String! | 選項的顯示文本 |
| `color` | String | 用於視覺表示的十六進制顏色代碼 |
| `position` | Float | 選項的排序順序 |
| `customField` | CustomField! | 此選項所屬的自定義欄位 |

### CustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位的唯一標識符 |
| `name` | String! | 多選欄位的顯示名稱 |
| `type` | CustomFieldType! | 始終是 `SELECT_MULTI` |
| `description` | String | 此欄位的幫助文本 |
| `customFieldOptions` | [CustomFieldOption!] | 所有可用選項 |

## 值格式

### 輸入格式
- **API 參數**: 選項 ID 陣列 (`["option1", "option2", "option3"]`)
- **字符串格式**: 逗號分隔的選項 ID (`"option1,option2,option3"`)

### 輸出格式
- **GraphQL 回應**: CustomFieldOption 物件的陣列
- **活動日誌**: 逗號分隔的選項標題
- **自動化數據**: 選項標題的陣列

## 管理選項

### 更新選項屬性
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
  }) {
    id
    title
    color
  }
}
```

### 刪除選項
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### 重新排序選項
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## 驗證規則

### 選項驗證
- 所有提供的選項 ID 必須存在
- 選項必須屬於指定的自定義欄位
- 只有 SELECT_MULTI 欄位可以選擇多個選項
- 空陣列是有效的（無選擇）

### 欄位驗證
- 必須定義至少一個選項才能使用
- 選項標題在欄位內必須唯一
- 顏色欄位接受任何字符串值（無十六進制驗證）

## 所需權限

| 行動 | 所需權限 |
|------|----------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## 錯誤回應

### 無效的選項 ID
```json
{
  "errors": [{
    "message": "Custom field option not found",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### 選項不屬於欄位
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### 找不到欄位
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

### 非多選欄位上的多個選項
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 最佳實踐

### 選項設計
- 使用描述性且簡潔的選項標題
- 應用一致的顏色編碼方案
- 保持選項列表可管理（通常 3-20 個選項）
- 合理排序選項（按字母順序、按頻率等）

### 數據管理
- 定期審查並清理未使用的選項
- 在項目中使用一致的命名約定
- 在創建欄位時考慮選項的可重用性
- 計劃選項的更新和遷移

### 用戶體驗
- 提供清晰的欄位描述
- 使用顏色改善視覺區分
- 將相關選項分組
- 考慮常見情況的默認選擇

## 常見用例

1. **項目管理**
   - 任務類別和標籤
   - 優先級別和類型
   - 團隊成員分配
   - 狀態指示器

2. **內容管理**
   - 文章類別和主題
   - 內容類型和格式
   - 發布渠道
   - 審批工作流程

3. **客戶支持**
   - 問題類別和類型
   - 受影響的產品或服務
   - 解決方法
   - 客戶細分

4. **產品開發**
   - 功能類別
   - 技術要求
   - 測試環境
   - 發布渠道

## 集成功能

### 與自動化
- 在選擇特定選項時觸發行動
- 根據選定類別路由工作
- 對高優先級選擇發送通知
- 根據選項組合創建後續任務

### 與查詢
- 根據選定選項過濾記錄
- 聚合選項選擇的數據
- 從其他記錄引用選項數據
- 根據選項組合創建報告

### 與表單
- 多選輸入控件
- 選項驗證和過濾
- 動態選項加載
- 條件欄位顯示

## 活動追蹤

多選欄位的變更會自動追蹤：
- 顯示添加和刪除的選項
- 在活動日誌中顯示選項標題
- 所有選擇變更的時間戳
- 修改的用戶歸屬

## 限制

- 選項的最大實際限制取決於 UI 性能
- 無層次或嵌套的選項結構
- 選項在使用該欄位的所有記錄中共享
- 無內置的選項分析或使用追蹤
- 顏色欄位接受任何字符串（無十六進制驗證）
- 不能為每個選項設置不同的權限
- 選項必須單獨創建，而不能與欄位創建同時進行
- 無專用的重新排序變更（使用 editCustomFieldOption 與位置）

## 相關資源

- [單選欄位](/api/custom-fields/select-single) - 用於單選選擇
- [複選框欄位](/api/custom-fields/checkbox) - 用於簡單布林選擇
- [文本欄位](/api/custom-fields/text-single) - 用於自由格式文本輸入
- [自定義欄位概述](/api/custom-fields/2.list-custom-fields) - 一般概念