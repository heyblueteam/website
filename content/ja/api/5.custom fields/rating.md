---
title: 評価カスタムフィールド
description: 設定可能なスケールと検証を持つ数値評価を保存するための評価フィールドを作成します
---

評価カスタムフィールドを使用すると、設定可能な最小値と最大値を持つレコードに数値評価を保存できます。これは、パフォーマンス評価、満足度スコア、優先度レベル、またはプロジェクト内の任意の数値スケールベースのデータに最適です。

## 基本例

デフォルトの0-5スケールを持つシンプルな評価フィールドを作成します：

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

## 高度な例

カスタムスケールと説明を持つ評価フィールドを作成します：

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

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 評価フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `RATING`である必要があります |
| `projectId` | String! | ✅ はい | このフィールドが作成されるプロジェクトID |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |
| `min` | Float | いいえ | 最小評価値（デフォルトなし） |
| `max` | Float | いいえ | 最大評価値 |

## 評価値の設定

レコードに評価値を設定または更新するには：

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | 評価カスタムフィールドのID |
| `value` | String! | ✅ はい | 文字列としての評価値（設定された範囲内） |

## 評価値を持つレコードの作成

評価値を持つ新しいレコードを作成する場合：

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

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `value` | Float | 保存された評価値（customField.valueを介してアクセス） |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

**注意**: 評価値は実際には`customField.value.number`を介してクエリでアクセスされます。

### CustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールドの一意の識別子 |
| `name` | String! | 評価フィールドの表示名 |
| `type` | CustomFieldType! | 常に`RATING` |
| `min` | Float | 許可される最小評価値 |
| `max` | Float | 許可される最大評価値 |
| `description` | String | フィールドのヘルプテキスト |

## 評価の検証

### 値の制約
- 評価値は数値でなければなりません（Float型）
- 値は設定された最小/最大範囲内でなければなりません
- 最小値が指定されていない場合、デフォルト値はありません
- 最大値はオプションですが推奨されます

### 検証ルール
**重要**: 検証はフォームを送信する際にのみ発生し、`setTodoCustomField`を直接使用する際には発生しません。

- 入力は浮動小数点数として解析されます（フォームを使用する場合）
- 最小値以上でなければなりません（フォームを使用する場合）
- 最大値以下でなければなりません（フォームを使用する場合）
- `setTodoCustomField`は検証なしで任意の文字列値を受け入れます

### 有効な評価の例
最小=1、最大=5のフィールドの場合：
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### 無効な評価の例
最小=1、最大=5のフィールドの場合：
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## 設定オプション

### 評価スケールの設定
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

### 一般的な評価スケール
- **1-5 スター**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 パフォーマンス**: `min: 1, max: 10`
- **0-100 パーセンテージ**: `min: 0, max: 100`
- **カスタムスケール**: 任意の数値範囲

## 必要な権限

カスタムフィールド操作は標準の役割ベースの権限に従います：

| アクション | 必要な役割 |
|--------|---------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**注意**: 必要な特定の役割は、プロジェクトのカスタム役割設定およびフィールドレベルの権限に依存します。

## エラーレスポンス

### 検証エラー（フォームのみ）
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

**重要**: 評価値の検証（最小/最大制約）は、フォームを送信する際にのみ発生し、`setTodoCustomField`を直接使用する際には発生しません。

### カスタムフィールドが見つかりません
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

## ベストプラクティス

### スケールデザイン
- 類似のフィールド間で一貫した評価スケールを使用する
- ユーザーの慣れを考慮する（1-5スター、0-10 NPS）
- 適切な最小値を設定する（0対1）
- 各評価レベルの明確な意味を定義する

### データ品質
- 保存する前に評価値を検証する
- 小数点精度を適切に使用する
- 表示目的での丸めを考慮する
- 評価の意味について明確なガイダンスを提供する

### ユーザーエクスペリエンス
- 評価スケールを視覚的に表示する（スター、プログレスバー）
- 現在の値とスケールの制限を表示する
- 評価の意味に関するコンテキストを提供する
- 新しいレコードのデフォルト値を考慮する

## 一般的なユースケース

1. **パフォーマンス管理**
   - 従業員のパフォーマンス評価
   - プロジェクトの品質スコア
   - タスク完了評価
   - スキルレベルの評価

2. **顧客フィードバック**
   - 満足度評価
   - 製品品質スコア
   - サービス体験評価
   - ネットプロモータースコア（NPS）

3. **優先度と重要性**
   - タスクの優先度レベル
   - 緊急度評価
   - リスク評価スコア
   - 影響評価

4. **品質保証**
   - コードレビュー評価
   - テスト品質スコア
   - ドキュメントの品質
   - プロセス遵守評価

## 統合機能

### 自動化との統合
- 評価の閾値に基づいてアクションをトリガーする
- 低評価の通知を送信する
- 高評価のフォローアップタスクを作成する
- 評価値に基づいて作業をルーティングする

### ルックアップとの統合
- レコード間での平均評価を計算する
- 評価範囲でレコードを検索する
- 他のレコードから評価データを参照する
- 評価統計を集計する

### Blue フロントエンドとの統合
- フォームコンテキストでの自動範囲検証
- 視覚的な評価入力コントロール
- リアルタイムの検証フィードバック
- スターまたはスライダー入力オプション

## アクティビティトラッキング

評価フィールドの変更は自動的に追跡されます：
- 古いおよび新しい評価値が記録されます
- アクティビティは数値の変更を示します
- すべての評価更新のタイムスタンプ
- 変更のユーザー帰属

## 制限事項

- 数値値のみがサポートされています
- ビジュアル評価表示（スターなど）は組み込まれていません
- 小数点精度はデータベース設定に依存します
- 評価メタデータの保存（コメント、コンテキスト）はありません
- 自動評価集計や統計はありません
- スケール間の評価変換は組み込まれていません
- **重要**: 最小/最大検証はフォームでのみ機能し、`setTodoCustomField`を介しては機能しません

## 関連リソース

- [数値フィールド](/api/5.custom%20fields/number) - 一般的な数値データ用
- [パーセントフィールド](/api/5.custom%20fields/percent) - パーセンテージ値用
- [選択フィールド](/api/5.custom%20fields/select-single) - 離散選択評価用
- [カスタムフィールドの概要](/api/5.custom%20fields/2.list-custom-fields) - 一般的な概念