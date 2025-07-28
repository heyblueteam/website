---
title: 時間持續時間自定義字段
description: 創建計算的時間持續時間字段，以跟踪工作流程中事件之間的時間
---

時間持續時間自定義字段自動計算並顯示工作流程中兩個事件之間的持續時間。它們非常適合跟踪處理時間、響應時間、週期時間或項目中的任何基於時間的指標。

## 基本示例

創建一個簡單的時間持續時間字段，以跟踪任務完成所需的時間：

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## 高級示例

創建一個複雜的時間持續時間字段，以跟踪自定義字段變更之間的時間，並設置 SLA 目標：

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## 輸入參數

### CreateCustomFieldInput (TIME_DURATION)

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 持續時間字段的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `TIME_DURATION` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ 是 | 如何顯示持續時間 |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ 是 | 開始事件配置 |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ 是 | 結束事件配置 |
| `timeDurationTargetTime` | Float | 否 | 用於 SLA 監控的目標持續時間（以秒為單位） |

### CustomFieldTimeDurationInput

| 參數 | 類型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ 是 | 要跟踪的事件類型 |
| `condition` | CustomFieldTimeDurationCondition! | ✅ 是 | `FIRST` 或 `LAST` 發生 |
| `customFieldId` | String | Conditional | 對於 `TODO_CUSTOM_FIELD` 類型是必需的 |
| `customFieldOptionIds` | [String!] | Conditional | 對於選擇字段變更是必需的 |
| `todoListId` | String | Conditional | 對於 `TODO_MOVED` 類型是必需的 |
| `tagId` | String | Conditional | 對於 `TODO_TAG_ADDED` 類型是必需的 |
| `assigneeId` | String | Conditional | 對於 `TODO_ASSIGNEE_ADDED` 類型是必需的 |

### CustomFieldTimeDurationType 值

| 值 | 描述 |
|-------|-------------|
| `TODO_CREATED_AT` | 記錄創建的時間 |
| `TODO_CUSTOM_FIELD` | 自定義字段值變更的時間 |
| `TODO_DUE_DATE` | 設置截止日期的時間 |
| `TODO_MARKED_AS_COMPLETE` | 記錄標記為完成的時間 |
| `TODO_MOVED` | 記錄移動到不同列表的時間 |
| `TODO_TAG_ADDED` | 向記錄添加標籤的時間 |
| `TODO_ASSIGNEE_ADDED` | 向記錄添加指派人的時間 |

### CustomFieldTimeDurationCondition 值

| 值 | 描述 |
|-------|-------------|
| `FIRST` | 使用事件的第一次發生 |
| `LAST` | 使用事件的最後一次發生 |

### CustomFieldTimeDurationDisplayType 值

| 值 | 描述 | 示例 |
|-------|-------------|---------|
| `FULL_DATE` | 天:小時:分鐘:秒格式 | `"01:02:03:04"` |
| `FULL_DATE_STRING` | 完整單詞書寫 | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | 帶單位的數字 | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | 僅以天為單位的持續時間 | `"2.5"` (2.5 days) |
| `HOURS` | 僅以小時為單位的持續時間 | `"60"` (60 hours) |
| `MINUTES` | 僅以分鐘為單位的持續時間 | `"3600"` (3600 minutes) |
| `SECONDS` | 僅以秒為單位的持續時間 | `"216000"` (216000 seconds) |

## 響應字段

### TodoCustomField 響應

| 字段 | 類型 | 描述 |
|-------|------|-------------|
| `id` | String! | 字段值的唯一標識符 |
| `customField` | CustomField! | 自定義字段定義 |
| `number` | Float | 持續時間（以秒為單位） |
| `value` | Float | 數字的別名（持續時間以秒為單位） |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後更新的時間 |

### CustomField 響應 (TIME_DURATION)

| 字段 | 類型 | 描述 |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | 持續時間的顯示格式 |
| `timeDurationStart` | CustomFieldTimeDuration | 開始事件配置 |
| `timeDurationEnd` | CustomFieldTimeDuration | 結束事件配置 |
| `timeDurationTargetTime` | Float | 目標持續時間（以秒為單位，用於 SLA 監控） |

## 持續時間計算

### 工作原理
1. **開始事件**：系統監控指定的開始事件
2. **結束事件**：系統監控指定的結束事件
3. **計算**：持續時間 = 結束時間 - 開始時間
4. **存儲**：持續時間以數字形式存儲（以秒為單位）
5. **顯示**：根據 `timeDurationDisplay` 設置進行格式化

### 更新觸發器
當以下情況發生時，持續時間值會自動重新計算：
- 記錄被創建或更新
- 自定義字段值發生變更
- 標籤被添加或移除
- 指派人被添加或移除
- 記錄在列表之間移動
- 記錄被標記為完成/未完成

## 讀取持續時間值

### 查詢持續時間字段
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### 格式化顯示值
持續時間值會根據 `timeDurationDisplay` 設置自動格式化：

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## 常見配置示例

### 任務完成時間
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### 狀態變更持續時間
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### 特定列表中的時間
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### 指派響應時間
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## 所需權限

| 操作 | 所需權限 |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## 錯誤響應

### 配置無效
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### 找不到引用的字段
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

### 缺少所需選項
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 重要說明

### 自動計算
- 持續時間字段是 **只讀的** - 值會自動計算
- 您不能通過 API 手動設置持續時間值
- 計算是通過後台作業異步進行的
- 當觸發事件發生時，值會自動更新

### 性能考量
- 持續時間計算會排隊並異步處理
- 大量的持續時間字段可能會影響性能
- 在設計持續時間字段時，考慮觸發事件的頻率
- 使用特定條件以避免不必要的重新計算

### 空值
當以下情況發生時，持續時間字段將顯示 `null`：
- 開始事件尚未發生
- 結束事件尚未發生
- 配置引用不存在的實體
- 計算遇到錯誤

## 最佳實踐

### 配置設計
- 儘可能使用特定事件類型，而不是通用類型
- 根據您的工作流程選擇適當的 `FIRST` 與 `LAST` 條件
- 在部署之前使用示例數據測試持續時間計算
- 為團隊成員記錄您的持續時間字段邏輯

### 顯示格式
- 使用 `FULL_DATE_SUBSTRING` 以獲得最易讀的格式
- 使用 `FULL_DATE` 以獲得緊湊且一致寬度的顯示
- 使用 `FULL_DATE_STRING` 用於正式報告和文件
- 使用 `DAYS`、`HOURS`、`MINUTES` 或 `SECONDS` 進行簡單的數字顯示
- 在選擇格式時考慮您的 UI 空間限制

### SLA 監控與目標時間
使用 `timeDurationTargetTime` 時：
- 設置目標持續時間（以秒為單位）
- 將實際持續時間與目標進行比較以滿足 SLA 要求
- 在儀表板中使用以突出顯示逾期項目
- 示例：24 小時響應 SLA = 86400 秒

### 工作流程集成
- 設計持續時間字段以匹配您的實際業務流程
- 使用持續時間數據進行流程改進和優化
- 監控持續時間趨勢以識別工作流程瓶頸
- 如有需要，設置持續時間閾值的警報

## 常見用例

1. **流程性能**
   - 任務完成時間
   - 審核週期時間
   - 批准處理時間
   - 響應時間

2. **SLA 監控**
   - 首次響應時間
   - 解決時間
   - 升級時間框架
   - 服務水平合規性

3. **工作流程分析**
   - 瓶頸識別
   - 流程優化
   - 團隊性能指標
   - 質量保證時間

4. **項目管理**
   - 階段持續時間
   - 里程碑時間
   - 資源分配時間
   - 交付時間框架

## 限制

- 持續時間字段是 **只讀的**，無法手動設置
- 值是異步計算的，可能不會立即可用
- 需要在您的工作流程中設置適當的事件觸發器
- 無法計算尚未發生的事件的持續時間
- 限於跟踪離散事件之間的時間（不進行連續時間跟踪）
- 沒有內置的 SLA 警報或通知
- 無法將多個持續時間計算聚合到單個字段中

## 相關資源

- [數字字段](/api/custom-fields/number) - 用於手動數值
- [日期字段](/api/custom-fields/date) - 用於特定日期跟踪
- [自定義字段概述](/api/custom-fields/list-custom-fields) - 一般概念
- [自動化](/api/automations) - 用於根據持續時間閾值觸發操作