---
title: 数値カスタムフィールド
description: 数値を格納するための数値フィールドを作成し、オプションの最小/最大制約およびプレフィックスフォーマットを設定します
---

数値カスタムフィールドを使用すると、レコードの数値を格納できます。これらは検証制約、少数精度をサポートし、数量、スコア、測定値、または特別なフォーマットを必要としない任意の数値データに使用できます。

## 基本的な例

シンプルな数値フィールドを作成します：

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高度な例

制約とプレフィックスを持つ数値フィールドを作成します：

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 数値フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `NUMBER` でなければなりません |
| `projectId` | String! | ✅ はい | フィールドを作成するプロジェクトのID |
| `min` | Float | いいえ | 最小値制約（UIガイダンスのみ） |
| `max` | Float | いいえ | 最大値制約（UIガイダンスのみ） |
| `prefix` | String | いいえ | 表示プレフィックス（例："#", "~", "$"） |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

## 数値の設定

数値フィールドは、オプションの検証を伴う少数値を格納します：

### シンプルな数値

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### 整数値

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | 数値カスタムフィールドのID |
| `number` | Float | いいえ | 格納する数値 |

## 値の制約

### 最小/最大制約（UIガイダンス）

**重要**: 最小/最大制約は保存されますが、サーバー側で強制されません。これらはフロントエンドアプリケーションのUIガイダンスとして機能します。

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**クライアント側の検証が必要**: フロントエンドアプリケーションは、最小/最大制約を強制するための検証ロジックを実装する必要があります。

### サポートされる値の型

| 型 | 例 | 説明 |
|------|---------|-------------|
| Integer | `42` | 整数 |
| Decimal | `42.5` | 小数点を含む数値 |
| Negative | `-10` | 負の値（最小制約がない場合） |
| Zero | `0` | ゼロ値 |

**注意**: 最小/最大制約はサーバー側で検証されません。指定された範囲外の値は受け入れられ、保存されます。

## 数値値を持つレコードの作成

数値値を持つ新しいレコードを作成する際：

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
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
        prefix
      }
      number
      value
    }
  }
}
```

### サポートされる入力形式

レコードを作成する際は、カスタムフィールド配列内で `number` パラメータを使用します（`value` ではありません）：

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `number` | Float | 数値 |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

### CustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド定義の一意の識別子 |
| `name` | String! | フィールドの表示名 |
| `type` | CustomFieldType! | 常に `NUMBER` |
| `min` | Float | 許可される最小値 |
| `max` | Float | 許可される最大値 |
| `prefix` | String | 表示プレフィックス |
| `description` | String | ヘルプテキスト |

**注意**: 数値が設定されていない場合、`number` フィールドは `null` になります。

## フィルタリングとクエリ

数値フィールドは包括的な数値フィルタリングをサポートします：

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### サポートされる演算子

| 演算子 | 説明 | 例 |
|----------|-------------|---------|
| `EQ` | 等しい | `number = 42` |
| `NE` | 等しくない | `number ≠ 42` |
| `GT` | より大きい | `number > 42` |
| `GTE` | より大きいまたは等しい | `number ≥ 42` |
| `LT` | より小さい | `number < 42` |
| `LTE` | より小さいまたは等しい | `number ≤ 42` |
| `IN` | 配列内 | `number in [1, 2, 3]` |
| `NIN` | 配列外 | `number not in [1, 2, 3]` |
| `IS` | NULLである/NULLでない | `number is null` |

### 範囲フィルタリング

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## 表示フォーマット

### プレフィックス付き

プレフィックスが設定されている場合、それが表示されます：

| 値 | プレフィックス | 表示 |
|-------|--------|---------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### 少数精度

数値はその少数精度を維持します：

| 入力 | 保存 | 表示 |
|-------|--------|-----------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## 必要な権限

| アクション | 必要な権限 |
|--------|--------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## エラーレスポンス

### 無効な数値形式
```json
{
  "errors": [{
    "message": "Invalid number format",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**注意**: 最小/最大検証エラーはサーバー側では発生しません。制約の検証はフロントエンドアプリケーションで実装する必要があります。

### 数値ではありません
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## ベストプラクティス

### 制約設計
- UIガイダンスのために現実的な最小/最大値を設定する
- 制約を強制するためにクライアント側の検証を実装する
- フォームでユーザーにフィードバックを提供するために制約を使用する
- 負の値がユースケースに有効かどうかを考慮する

### 値の精度
- 必要に応じた適切な少数精度を使用する
- 表示目的での丸めを考慮する
- 関連フィールド間で精度を一貫させる

### 表示の強化
- コンテキストに意味のあるプレフィックスを使用する
- フィールド名に単位を考慮する（例："重量 (kg)"）
- 検証ルールの明確な説明を提供する

## 一般的なユースケース

1. **スコアリングシステム**
   - パフォーマンス評価
   - 品質スコア
   - 優先度レベル
   - 顧客満足度評価

2. **測定**
   - 数量と金額
   - 寸法とサイズ
   - 期間（数値形式）
   - 容量と制限

3. **ビジネスメトリクス**
   - 収益数値
   - コンバージョン率
   - 予算配分
   - 目標数値

4. **技術データ**
   - バージョン番号
   - 設定値
   - パフォーマンスメトリクス
   - 閾値設定

## 統合機能

### チャートとダッシュボードとの連携
- チャート計算に数値フィールドを使用する
- 数値の視覚化を作成する
- 時間の経過に伴うトレンドを追跡する

### 自動化との連携
- 数値の閾値に基づいてアクションをトリガーする
- 数値の変更に基づいて関連フィールドを更新する
- 特定の値に対して通知を送信する

### ルックアップとの連携
- 関連レコードから数値を集計する
- 合計と平均を計算する
- 関係全体で最小/最大値を見つける

### チャートとの連携
- 数値の視覚化を作成する
- 時間の経過に伴うトレンドを追跡する
- レコード間で値を比較する

## 制限事項

- **最小/最大制約のサーバー側検証なし**
- **制約の強制にはクライアント側の検証が必要**
- 組み込みの通貨フォーマットなし（代わりにCURRENCY型を使用）
- 自動的なパーセント記号なし（代わりにPERCENT型を使用）
- 単位変換機能なし
- データベースのDecimal型によって制限された少数精度
- フィールド自体での数学的な式評価なし

## 関連リソース

- [カスタムフィールドの概要](/api/custom-fields/1.index) - 一般的なカスタムフィールドの概念
- [通貨カスタムフィールド](/api/custom-fields/currency) - 金銭的価値用
- [パーセントカスタムフィールド](/api/custom-fields/percent) - パーセント値用
- [自動化API](/api/automations/1.index) - 数値に基づく自動化を作成