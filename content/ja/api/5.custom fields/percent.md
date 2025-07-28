---
title: パーセントカスタムフィールド
description: 数値を保存するためのパーセントフィールドを作成し、自動的な%記号の処理と表示フォーマットを提供します
---

パーセントカスタムフィールドは、レコードのためにパーセント値を保存することを可能にします。入力と表示のために%記号を自動的に処理し、内部では生の数値を保存します。完了率、成功率、または任意のパーセントベースのメトリクスに最適です。

## 基本的な例

シンプルなパーセントフィールドを作成します：

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きのパーセントフィールドを作成します：

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
  }) {
    id
    name
    type
    description
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | パーセントフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `PERCENT` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: プロジェクトコンテキストは、認証ヘッダーから自動的に決定されます。`projectId` パラメータは必要ありません。

**注意**: PERCENTフィールドは、NUMBERフィールドのような最小/最大制約やプレフィックスフォーマットをサポートしていません。

## パーセント値の設定

パーセントフィールドは、自動的な%記号の処理を伴う数値を保存します：

### パーセント記号付き

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### 直接数値

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | パーセントカスタムフィールドのID |
| `number` | Float | いいえ | 数値パーセント値（例：75.5は75.5%を表します） |

## 値の保存と表示

### 保存形式
- **内部保存**: 生の数値（例：75.5）
- **データベース**: `Decimal` 列の中に `number` として保存されます
- **GraphQL**: `Float` 型として返されます

### 表示形式
- **ユーザーインターフェース**: クライアントアプリケーションは%記号を追加する必要があります（例："75.5%"）
- **チャート**: 出力タイプがPERCENTAGEのときに%記号付きで表示されます
- **APIレスポンス**: %記号なしの生の数値（例：75.5）

## パーセント値を持つレコードの作成

パーセント値を持つ新しいレコードを作成する場合：

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### サポートされている入力形式

| 形式 | 例 | 結果 |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**注意**: %記号は自動的に入力から削除され、表示時に再追加されます。

## パーセント値のクエリ

パーセントカスタムフィールドを持つレコードをクエリする際は、`customField.value.number` パスを通じて値にアクセスします：

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

レスポンスには、生の数値としてパーセントが含まれます：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## レスポンスフィールド

### TodoCustomFieldレスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | ID! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールド定義（パーセント値を含む） |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

**重要**: パーセント値は、`customField.value.number` フィールドを通じてアクセスされます。%記号は保存された値には含まれておらず、表示のためにクライアントアプリケーションによって追加される必要があります。

## フィルタリングとクエリ

パーセントフィールドは、NUMBERフィールドと同じフィルタリングをサポートします：

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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

### サポートされている演算子

| 演算子 | 説明 | 例 |
|----------|-------------|---------|
| `EQ` | 等しい | `percentage = 75` |
| `NE` | 等しくない | `percentage ≠ 75` |
| `GT` | より大きい | `percentage > 75` |
| `GTE` | より大きいか等しい | `percentage ≥ 75` |
| `LT` | より小さい | `percentage < 75` |
| `LTE` | より小さいか等しい | `percentage ≤ 75` |
| `IN` | リスト内の値 | `percentage in [50, 75, 100]` |
| `NIN` | リスト内にない値 | `percentage not in [0, 25]` |
| `IS` | `values: null` でnullをチェック | `percentage is null` |
| `NOT` | `values: null` でnullでないことをチェック | `percentage is not null` |

### 範囲フィルタリング

範囲フィルタリングには、複数の演算子を使用します：

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## パーセント値の範囲

### 一般的な範囲

| 範囲 | 説明 | 使用例 |
|-------|-------------|----------|
| `0-100` | 標準パーセント | Completion rates, success rates |
| `0-∞` | 無制限パーセント | Growth rates, performance metrics |
| `-∞-∞` | 任意の値 | Change rates, variance |

### 例の値

| 入力 | 保存された | 表示 |
|-------|--------|---------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## チャート集計

パーセントフィールドは、ダッシュボードチャートやレポートでの集計をサポートします。利用可能な関数には以下が含まれます：

- `AVERAGE` - 平均パーセント値
- `COUNT` - 値を持つレコードの数
- `MIN` - 最低パーセント値
- `MAX` - 最高パーセント値 
- `SUM` - すべてのパーセント値の合計

これらの集計は、チャートやダッシュボードを作成する際に利用可能であり、直接のGraphQLクエリでは利用できません。

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## エラーレスポンス

### 無効なパーセント形式
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 数値ではない
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

### 値の入力
- ユーザーが%記号ありまたはなしで入力できるようにします
- 使用例に対して合理的な範囲を検証します
- 100%が何を表すかについて明確なコンテキストを提供します

### 表示
- ユーザーインターフェースでは常に%記号を表示します
- 適切な小数点精度を使用します
- 範囲に対して色分けを検討します（赤/黄/緑）

### データの解釈
- 100%があなたのコンテキストで何を意味するかを文書化します
- 100%を超える値を適切に処理します
- 負の値が有効かどうかを検討します

## 一般的な使用例

1. **プロジェクト管理**
   - タスク完了率
   - プロジェクトの進捗
   - リソースの利用率
   - スプリントの速度

2. **パフォーマンス追跡**
   - 成功率
   - エラー率
   - 効率メトリクス
   - 品質スコア

3. **財務メトリクス**
   - 成長率
   - 利益率
   - 割引額
   - 変化率

4. **分析**
   - コンバージョン率
   - クリック率
   - エンゲージメントメトリクス
   - パフォーマンス指標

## 統合機能

### 数式との併用
- 計算にPERCENTフィールドを参照します
- 数式出力で自動的に%記号のフォーマットを行います
- 他の数値フィールドと組み合わせます

### 自動化との併用
- パーセントの閾値に基づいてアクションをトリガーします
- マイルストーンパーセントの通知を送信します
- 完了率に基づいてステータスを更新します

### ルックアップとの併用
- 関連レコードからのパーセントを集計します
- 平均成功率を計算します
- 最高/最低のパフォーマンスアイテムを見つけます

### チャートとの併用
- パーセントベースの視覚化を作成します
- 時間の経過に伴う進捗を追跡します
- パフォーマンスメトリクスを比較します

## NUMBERフィールドとの違い

### 何が異なるか
- **入力処理**: 自動的に%記号を削除します
- **表示**: 自動的に%記号を追加します
- **制約**: 最小/最大の検証はありません
- **フォーマット**: プレフィックスサポートはありません

### 何が同じか
- **保存**: 同じデータベースの列と型
- **フィルタリング**: 同じクエリ演算子
- **集計**: 同じ集計関数
- **権限**: 同じ権限モデル

## 制限事項

- 最小/最大値の制約はありません
- プレフィックスフォーマットオプションはありません
- 0-100%範囲の自動検証はありません
- パーセント形式間の変換はありません（例：0.75 ↔ 75%）
- 100%を超える値が許可されています

## 関連リソース

- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的なカスタムフィールドの概念
- [数値カスタムフィールド](/api/custom-fields/number) - 生の数値のために
- [自動化API](/api/automations/index) - パーセントベースの自動化を作成します