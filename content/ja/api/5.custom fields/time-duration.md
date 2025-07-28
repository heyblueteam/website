---
title: 時間の長さカスタムフィールド
description: ワークフロー内のイベント間の時間を追跡する計算された時間の長さフィールドを作成します
---

時間の長さカスタムフィールドは、ワークフロー内の2つのイベント間の期間を自動的に計算して表示します。これらは、処理時間、応答時間、サイクル時間、またはプロジェクト内の時間に基づくメトリックを追跡するのに最適です。

## 基本的な例

タスクの完了にかかる時間を追跡するシンプルな時間の長さフィールドを作成します：

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

## 高度な例

SLAターゲットを持つカスタムフィールドの変更間の時間を追跡する複雑な時間の長さフィールドを作成します：

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

## 入力パラメータ

### CreateCustomFieldInput (TIME_DURATION)

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 時間の長さフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `TIME_DURATION` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ はい | 時間の長さを表示する方法 |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ はい | 開始イベントの設定 |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ はい | 終了イベントの設定 |
| `timeDurationTargetTime` | Float | いいえ | SLA監視のための目標時間（秒） |

### CustomFieldTimeDurationInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ はい | 追跡するイベントのタイプ |
| `condition` | CustomFieldTimeDurationCondition! | ✅ はい | `FIRST` または `LAST` の発生 |
| `customFieldId` | String | Conditional | `TODO_CUSTOM_FIELD` タイプに必要 |
| `customFieldOptionIds` | [String!] | Conditional | 選択フィールドの変更に必要 |
| `todoListId` | String | Conditional | `TODO_MOVED` タイプに必要 |
| `tagId` | String | Conditional | `TODO_TAG_ADDED` タイプに必要 |
| `assigneeId` | String | Conditional | `TODO_ASSIGNEE_ADDED` タイプに必要 |

### CustomFieldTimeDurationType 値

| 値 | 説明 |
|-------|-------------|
| `TODO_CREATED_AT` | レコードが作成されたとき |
| `TODO_CUSTOM_FIELD` | カスタムフィールドの値が変更されたとき |
| `TODO_DUE_DATE` | 期限が設定されたとき |
| `TODO_MARKED_AS_COMPLETE` | レコードが完了としてマークされたとき |
| `TODO_MOVED` | レコードが別のリストに移動されたとき |
| `TODO_TAG_ADDED` | レコードにタグが追加されたとき |
| `TODO_ASSIGNEE_ADDED` | レコードに担当者が追加されたとき |

### CustomFieldTimeDurationCondition 値

| 値 | 説明 |
|-------|-------------|
| `FIRST` | イベントの最初の発生を使用 |
| `LAST` | イベントの最後の発生を使用 |

### CustomFieldTimeDurationDisplayType 値

| 値 | 説明 | 例 |
|-------|-------------|---------|
| `FULL_DATE` | 日:時間:分:秒形式 | `"01:02:03:04"` |
| `FULL_DATE_STRING` | 完全な単語で書かれた形式 | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | 単位付きの数値 | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | 日数のみの期間 | `"2.5"` (2.5 days) |
| `HOURS` | 時間のみの期間 | `"60"` (60 hours) |
| `MINUTES` | 分のみの期間 | `"3600"` (3600 minutes) |
| `SECONDS` | 秒のみの期間 | `"216000"` (216000 seconds) |

## 応答フィールド

### TodoCustomField 応答

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `number` | Float | 秒単位の期間 |
| `value` | Float | 数値のエイリアス（秒単位の期間） |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成されたとき |
| `updatedAt` | DateTime! | 値が最後に更新されたとき |

### CustomField 応答 (TIME_DURATION)

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | 期間の表示形式 |
| `timeDurationStart` | CustomFieldTimeDuration | 開始イベントの設定 |
| `timeDurationEnd` | CustomFieldTimeDuration | 終了イベントの設定 |
| `timeDurationTargetTime` | Float | SLA監視のための目標時間（秒） |

## 期間計算

### 仕組み
1. **開始イベント**: システムは指定された開始イベントを監視します
2. **終了イベント**: システムは指定された終了イベントを監視します
3. **計算**: 期間 = 終了時間 - 開始時間
4. **保存**: 期間は数値として秒単位で保存されます
5. **表示**: `timeDurationDisplay` 設定に従ってフォーマットされます

### 更新トリガー
期間値は以下のときに自動的に再計算されます：
- レコードが作成または更新されたとき
- カスタムフィールドの値が変更されたとき
- タグが追加または削除されたとき
- 担当者が追加または削除されたとき
- レコードがリスト間で移動されたとき
- レコードが完了/未完了としてマークされたとき

## 期間値の読み取り

### 期間フィールドのクエリ
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

### フォーマットされた表示値
期間値は `timeDurationDisplay` 設定に基づいて自動的にフォーマットされます：

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

## 一般的な設定例

### タスク完了時間
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

### ステータス変更の期間
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

### 特定のリスト内の時間
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

### 割り当て応答時間
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

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## エラー応答

### 無効な設定
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

### 参照されたフィールドが見つかりません
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

### 必要なオプションが不足しています
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

## 重要な注意事項

### 自動計算
- 期間フィールドは **読み取り専用** です - 値は自動的に計算されます
- APIを介して手動で期間値を設定することはできません
- 計算はバックグラウンドジョブを介して非同期に行われます
- トリガーイベントが発生すると、値は自動的に更新されます

### パフォーマンスの考慮事項
- 期間計算はキューに入れられ、非同期に処理されます
- 大量の期間フィールドはパフォーマンスに影響を与える可能性があります
- 期間フィールドを設計する際には、トリガーイベントの頻度を考慮してください
- 不要な再計算を避けるために特定の条件を使用してください

### Null値
期間フィールドは以下のときに `null` を表示します：
- 開始イベントがまだ発生していない
- 終了イベントがまだ発生していない
- 設定が存在しないエンティティを参照している
- 計算中にエラーが発生した

## ベストプラクティス

### 設定設計
- 可能な限り一般的なイベントタイプではなく、特定のイベントタイプを使用してください
- ワークフローに基づいて適切な `FIRST` と `LAST` 条件を選択してください
- デプロイ前にサンプルデータで期間計算をテストしてください
- チームメンバーのために期間フィールドのロジックを文書化してください

### 表示フォーマット
- 最も読みやすい形式には `FULL_DATE_SUBSTRING` を使用してください
- 一貫した幅のコンパクトな表示には `FULL_DATE` を使用してください
- 公式なレポートや文書には `FULL_DATE_STRING` を使用してください
- シンプルな数値表示には `DAYS`、`HOURS`、`MINUTES`、または `SECONDS` を使用してください
- フォーマットを選択する際にはUIスペースの制約を考慮してください

### SLA監視と目標時間
`timeDurationTargetTime` を使用する場合：
- 目標期間を秒単位で設定してください
- SLA遵守のために実際の期間を目標と比較してください
- ダッシュボードで期限切れの項目を強調表示するために使用してください
- 例：24時間応答SLA = 86400秒

### ワークフロー統合
- 実際のビジネスプロセスに合わせて期間フィールドを設計してください
- プロセス改善と最適化のために期間データを使用してください
- ワークフローのボトルネックを特定するために期間の傾向を監視してください
- 必要に応じて期間のしきい値に対してアラートを設定してください

## 一般的な使用例

1. **プロセスパフォーマンス**
   - タスク完了時間
   - レビューサイクル時間
   - 承認処理時間
   - 応答時間

2. **SLA監視**
   - 最初の応答までの時間
   - 解決時間
   - エスカレーションの時間枠
   - サービスレベルの遵守

3. **ワークフロー分析**
   - ボトルネックの特定
   - プロセスの最適化
   - チームパフォーマンスメトリック
   - 品質保証のタイミング

4. **プロジェクト管理**
   - フェーズの期間
   - マイルストーンのタイミング
   - リソース配分時間
   - 配送時間枠

## 制限事項

- 期間フィールドは **読み取り専用** であり、手動で設定することはできません
- 値は非同期に計算され、すぐに利用できない場合があります
- ワークフロー内で適切なイベントトリガーが設定されている必要があります
- 発生していないイベントの期間を計算することはできません
- 離散イベント間の時間を追跡することに限定されています（連続時間の追跡ではありません）
- 組み込みのSLAアラートや通知はありません
- 複数の期間計算を単一のフィールドに集約することはできません

## 関連リソース

- [数値フィールド](/api/custom-fields/number) - 手動の数値値用
- [日付フィールド](/api/custom-fields/date) - 特定の日付の追跡用
- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的な概念
- [自動化](/api/automations) - 期間のしきい値に基づいてアクションをトリガーするためのもの