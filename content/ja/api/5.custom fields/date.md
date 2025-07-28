---
title: 日付カスタムフィールド
description: タイムゾーンサポートを備えた単一の日付または日付範囲を追跡するための日付フィールドを作成します
---

日付カスタムフィールドを使用すると、レコードの単一の日付または日付範囲を保存できます。これらはタイムゾーン処理、インテリジェントフォーマットをサポートしており、締切、イベント日付、または任意の時間ベースの情報を追跡するために使用できます。

## 基本例

シンプルな日付フィールドを作成します：

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

## 高度な例

説明付きの期限日フィールドを作成します：

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

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 日付フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `DATE` でなければなりません |
| `isDueDate` | Boolean | いいえ | このフィールドが期限日を表すかどうか |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: カスタムフィールドは、ユーザーの現在のプロジェクトコンテキストに基づいてプロジェクトに自動的に関連付けられます。`projectId` パラメータは必要ありません。

## 日付値の設定

日付フィールドは、単一の日付または日付範囲のいずれかを保存できます：

### 単一の日付

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

### 日付範囲

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

### 終日イベント

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

### SetTodoCustomFieldInput パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | 日付カスタムフィールドのID |
| `startDate` | DateTime | いいえ | ISO 8601形式の開始日/時刻 |
| `endDate` | DateTime | いいえ | ISO 8601形式の終了日/時刻 |
| `timezone` | String | いいえ | タイムゾーン識別子（例："America/New_York"） |

**注意**: もし`startDate`のみが提供された場合、`endDate`は自動的に同じ値にデフォルト設定されます。

## 日付フォーマット

### ISO 8601フォーマット
すべての日付はISO 8601形式で提供する必要があります：
- `2025-01-15T14:30:00Z` - UTC時間
- `2025-01-15T14:30:00+05:00` - タイムゾーンオフセット付き
- `2025-01-15T14:30:00.123Z` - ミリ秒付き

### タイムゾーン識別子
標準のタイムゾーン識別子を使用します：
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

タイムゾーンが提供されない場合、システムはユーザーの検出されたタイムゾーンにデフォルト設定されます。

## 日付値を持つレコードの作成

日付値を持つ新しいレコードを作成する際：

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

### サポートされる入力フォーマット

レコードを作成する際、日付はさまざまなフォーマットで提供できます：

| フォーマット | 例 | 結果 |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## レスポンスフィールド

### TodoCustomFieldレスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | ID! | フィールド値の一意の識別子 |
| `uid` | String! | 一意の識別子文字列 |
| `customField` | CustomField! | カスタムフィールド定義（日付値を含む） |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に変更された日時 |

**重要**: 日付値（`startDate`、`endDate`、`timezone`）は、TodoCustomField上ではなく、`customField.value`フィールドを通じてアクセスされます。

### 値オブジェクト構造

日付値は、`customField.value`フィールドを通じてJSONオブジェクトとして返されます：

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

**注意**: `value`フィールドは、`CustomField`タイプにあり、`TodoCustomField`にはありません。

## 日付値のクエリ

日付カスタムフィールドを持つレコードをクエリする際は、`customField.value`フィールドを通じて日付値にアクセスします：

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

レスポンスには、`value`フィールドに日付値が含まれます：

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

## 日付表示インテリジェンス

システムは、範囲に基づいて日付を自動的にフォーマットします：

| シナリオ | 表示フォーマット |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025`（時間は表示されません） |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**終日検出**: 00:00から23:59までのイベントは、自動的に終日イベントとして検出されます。

## タイムゾーン処理

### ストレージ
- すべての日付はデータベースにUTCで保存されます
- タイムゾーン情報は別々に保持されます
- 表示時に変換が行われます

### ベストプラクティス
- 正確性のために常にタイムゾーンを提供する
- プロジェクト内で一貫したタイムゾーンを使用する
- グローバルチームのユーザーの位置を考慮する

### 一般的なタイムゾーン

| 地域 | タイムゾーンID | UTCオフセット |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## フィルタリングとクエリ

日付フィールドは複雑なフィルタリングをサポートします：

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

### 空の日付フィールドの確認

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

### サポートされている演算子

| 演算子 | 使用法 | 説明 |
|----------|-------|-------------|
| `EQ` | dateRangeを使用 | 日付が指定された範囲と重複する（任意の交差） |
| `NE` | dateRangeを使用 | 日付が範囲と重複しない |
| `IS` | `values: null`を使用 | 日付フィールドが空である（startDateまたはendDateがnull） |
| `NOT` | `values: null`を使用 | 日付フィールドに値がある（両方の日付がnullでない） |

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## エラーレスポンス

### 無効な日付フォーマット
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

### フィールドが見つかりません
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


## 制限事項

- 繰り返し日付のサポートはありません（繰り返しイベントには自動化を使用）
- 日付なしで時間を設定することはできません
- 組み込みの営業日計算はありません
- 日付範囲は自動的に終了 > 開始を検証しません
- 最大精度は秒まで（ミリ秒の保存はありません）

## 関連リソース

- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的なカスタムフィールドの概念
- [自動化API](/api/automations/index) - 日付ベースの自動化を作成