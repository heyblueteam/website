---
title: フォーミュラカスタムフィールド
description: 他のデータに基づいて自動的に値を計算する計算フィールドを作成します
---

フォーミュラカスタムフィールドは、Blue内のチャートおよびダッシュボードの計算に使用されます。これらは、カスタムフィールドデータに対して動作する集計関数（SUM、AVERAGE、COUNTなど）を定義し、チャート内に計算されたメトリックを表示します。フォーミュラは、個々のtodoレベルで計算されるのではなく、視覚化の目的で複数のレコードにわたってデータを集約します。

## 基本例

チャート計算用のフォーミュラフィールドを作成します：

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

## 高度な例

複雑な計算を伴う通貨フォーミュラを作成します：

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

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | フォーミュラフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `FORMULA` である必要があります |
| `projectId` | String! | ✅ はい | このフィールドが作成されるプロジェクトのID |
| `formula` | JSON | いいえ | チャート計算用のフォーミュラ定義 |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

### フォーミュラ構造

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

## サポートされている関数

### チャート集計関数

フォーミュラフィールドは、チャート計算のために以下の集計関数をサポートしています：

| 関数 | 説明 | ChartFunction Enum |
|----------|-------------|-------------------|
| `SUM` | すべての値の合計 | `SUM` |
| `AVERAGE` | 数値の平均 | `AVERAGE` |
| `AVERAGEA` | ゼロとNULLを除外した平均 | `AVERAGEA` |
| `COUNT` | 値のカウント | `COUNT` |
| `COUNTA` | ゼロとNULLを除外したカウント | `COUNTA` |
| `MAX` | 最大値 | `MAX` |
| `MIN` | 最小値 | `MIN` |

**注意**: これらの関数は、`display.function` フィールドで使用され、チャートの視覚化のために集約データに対して動作します。複雑な数学的表現やフィールドレベルの計算はサポートされていません。

## 表示タイプ

### 数字表示

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

結果: `1250.75`

### 通貨表示

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

結果: `$1,250.75`

### パーセンテージ表示

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

結果: `87.5%`

## フォーミュラフィールドの編集

既存のフォーミュラフィールドを更新します：

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

## フォーミュラ処理

### チャート計算コンテキスト

フォーミュラフィールドは、チャートセグメントおよびダッシュボードのコンテキストで処理されます：
- チャートがレンダリングまたは更新されるときに計算が行われます
- 結果は、`ChartSegment.formulaResult` に小数値として保存されます
- 処理は、「formula」という名前の専用BullMQキューを通じて行われます
- 更新は、リアルタイム更新のためにダッシュボードの購読者に公開されます

### 表示フォーマット

`getFormulaDisplayValue` 関数は、表示タイプに基づいて計算結果をフォーマットします：
- **NUMBER**: オプションの精度を持つ通常の数字として表示
- **PERCENTAGE**: オプションの精度を持つ % サフィックスを追加  
- **CURRENCY**: 指定された通貨コードを使用してフォーマット

## フォーミュラ結果の保存

結果は、`formulaResult` フィールドに保存されます：

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

## レスポンスフィールド

### TodoCustomFieldレスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | フォーミュラフィールドの定義 |
| `number` | Float | 計算された数値結果 |
| `formulaResult` | JSON | 表示フォーマット付きの完全な結果 |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に計算された日時 |

## データコンテキスト

### チャートデータソース

フォーミュラフィールドは、チャートデータソースのコンテキスト内で動作します：
- フォーミュラは、プロジェクト内のtodosにわたってカスタムフィールド値を集約します
- `display.function` に指定された集計関数が計算を決定します
- 結果は、SQL集計関数（avg、sum、countなど）を使用して計算されます
- 効率のためにデータベースレベルで計算が行われます

## 一般的なフォーミュラの例

### 総予算（チャート表示）

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

### 平均スコア（チャート表示）

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

### タスク数（チャート表示）

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

## 必要な権限

カスタムフィールド操作は、標準の役割ベースの権限に従います：

| アクション | 必要な役割 |
|--------|---------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**注意**: 必要な特定の役割は、プロジェクトのカスタム役割設定に依存します。CUSTOM_FIELDS_CREATEのような特別な権限定数はありません。

## エラーハンドリング

### バリデーションエラー
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

### フォーミュラ設計
- フォーミュラフィールドには明確で説明的な名前を使用します
- 計算ロジックを説明する説明を追加します
- デプロイ前にサンプルデータでフォーミュラをテストします
- フォーミュラはシンプルで読みやすく保ちます

### パフォーマンス最適化
- 深くネストされたフォーミュラ依存関係を避けます
- ワイルドカードではなく特定のフィールド参照を使用します
- 複雑な計算のためのキャッシング戦略を検討します
- 大規模プロジェクトでのフォーミュラのパフォーマンスを監視します

### データ品質
- フォーミュラで使用する前にソースデータを検証します
- 空またはNULLの値を適切に処理します
- 表示タイプに適切な精度を使用します
- 計算におけるエッジケースを考慮します

## 一般的なユースケース

1. **財務追跡**
   - 予算計算
   - 利益/損失計算
   - コスト分析
   - 収益予測

2. **プロジェクト管理**
   - 完了率
   - リソース利用率
   - タイムライン計算
   - パフォーマンスメトリック

3. **品質管理**
   - 平均スコア
   - 合格/不合格率
   - 品質メトリック
   - コンプライアンス追跡

4. **ビジネスインテリジェンス**
   - KPI計算
   - トレンド分析
   - 比較メトリック
   - ダッシュボード値

## 制限事項

- フォーミュラはチャート/ダッシュボードの集計のみに使用され、todoレベルの計算には使用されません
- サポートされている7つの集計関数（SUM、AVERAGEなど）に制限されています
- 複雑な数学的表現やフィールド間の計算はありません
- 単一のフォーミュラ内で複数のフィールドを参照することはできません
- 結果はチャートおよびダッシュボードでのみ表示されます
- `logic` フィールドは表示テキスト専用で、実際の計算ロジックではありません

## 関連リソース

- [数値フィールド](/api/5.custom%20fields/number) - 静的数値値用
- [通貨フィールド](/api/5.custom%20fields/currency) - 金銭的値用
- [参照フィールド](/api/5.custom%20fields/reference) - プロジェクト間データ用
- [ルックアップフィールド](/api/5.custom%20fields/lookup) - 集約データ用
- [カスタムフィールドの概要](/api/5.custom%20fields/2.list-custom-fields) - 一般的な概念