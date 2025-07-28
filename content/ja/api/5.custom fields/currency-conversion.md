---
title: 通貨変換カスタムフィールド
description: リアルタイム為替レートを使用して通貨値を自動的に変換するフィールドを作成します
---

通貨変換カスタムフィールドは、ソースCURRENCYフィールドから異なるターゲット通貨に値を自動的に変換します。これらのフィールドは、ソース通貨の値が変更されるたびに自動的に更新されます。

変換レートは、[Frankfurter API](https://github.com/hakanensari/frankfurter)によって提供され、これは[欧州中央銀行](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html)が発表した基準為替レートを追跡するオープンソースサービスです。これにより、国際ビジネスのニーズに対して正確で信頼性が高く、最新の通貨変換が保証されます。

## 基本例

シンプルな通貨変換フィールドを作成します：

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## 高度な例

特定の日付の履歴レートを使用した変換フィールドを作成します：

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## 完全なセットアッププロセス

通貨変換フィールドを設定するには、3つのステップが必要です：

### ステップ1: ソースCURRENCYフィールドを作成

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### ステップ2: CURRENCY_CONVERSIONフィールドを作成

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### ステップ3: 変換オプションを作成

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 変換フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `CURRENCY_CONVERSION` である必要があります |
| `currencyFieldId` | String | いいえ | 変換元のソースCURRENCYフィールドのID |
| `conversionDateType` | String | いいえ | 為替レートのための日付戦略（以下を参照） |
| `conversionDate` | String | いいえ | 変換のための日付文字列（conversionDateTypeに基づく） |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: カスタムフィールドは、ユーザーの現在のプロジェクトコンテキストに基づいて自動的にプロジェクトに関連付けられます。`projectId` パラメータは必要ありません。

### 変換日タイプ

| タイプ | 説明 | conversionDate パラメータ |
|------|-------------|-------------------------|
| `currentDate` | リアルタイム為替レートを使用 | 必要なし |
| `specificDate` | 固定日付のレートを使用 | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | 別のフィールドの日付を使用 | "todoDueDate" or DATE field ID |

## 変換オプションの作成

変換オプションは、どの通貨ペアが変換できるかを定義します：

### CreateCustomFieldOptionInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ はい | CURRENCY_CONVERSIONフィールドのID |
| `title` | String! | ✅ はい | この変換オプションの表示名 |
| `currencyConversionFrom` | String! | ✅ はい | ソース通貨コードまたは「Any」 |
| `currencyConversionTo` | String! | ✅ はい | ターゲット通貨コード |

### ソースとして「Any」を使用

特別な値「Any」を`currencyConversionFrom`として使用すると、フォールバックオプションが作成されます：

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

このオプションは、特定の通貨ペアの一致が見つからない場合に使用されます。

## 自動変換の仕組み

1. **値の更新**: ソースCURRENCYフィールドに値が設定されると
2. **オプションの一致**: システムはソース通貨に基づいて一致する変換オプションを見つけます
3. **レートの取得**: Frankfurter APIから為替レートを取得します
4. **計算**: ソース金額に為替レートを掛けます
5. **保存**: ターゲット通貨コードとともに変換された値を保存します

### 例のフロー

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## 日付ベースの変換

### 現在の日付を使用

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

ソース値が変更されるたびに、変換は現在の為替レートで更新されます。

### 特定の日付を使用

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

常に指定された日付の為替レートを使用します。

### フィールドからの日付を使用

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

別のフィールドからの日付を使用します（todoの期限日またはDATEカスタムフィールドのいずれか）。

## レスポンスフィールド

### TodoCustomFieldレスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | 変換フィールドの定義 |
| `number` | Float | 変換された金額 |
| `currency` | String | ターゲット通貨コード |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に更新された日時 |

## 為替レートのソース

Blueは為替レートに**Frankfurter API**を使用します：
- 欧州中央銀行がホストするオープンソースAPI
- 公式為替レートで毎日更新
- 1999年までの履歴レートをサポート
- ビジネス利用に無料で信頼性があります

## エラーハンドリング

### 変換の失敗

変換が失敗した場合（APIエラー、無効な通貨など）：
- 変換された値は`0`に設定されます
- ターゲット通貨はまだ保存されます
- ユーザーにエラーは表示されません

### 一般的なシナリオ

| シナリオ | 結果 |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| 一致するオプションなし | Uses "Any" option if available |
| Missing source value | 変換は実行されませんでした |

## 必要な権限

カスタムフィールド管理にはプロジェクトレベルのアクセスが必要です：

| 役割 | フィールドの作成/更新が可能 |
|------|-------------------------|
| `OWNER` | ✅ はい |
| `ADMIN` | ✅ はい |
| `MEMBER` | ❌ いいえ |
| `CLIENT` | ❌ いいえ |

変換された値の表示権限は標準のレコードアクセスルールに従います。

## ベストプラクティス

### オプション設定
- 一般的な変換のために特定の通貨ペアを作成します
- 柔軟性のために「Any」フォールバックオプションを追加します
- オプションに説明的なタイトルを使用します

### 日付戦略の選択
- リアルタイムの財務追跡には`currentDate`を使用します
- 履歴報告には`specificDate`を使用します
- 取引特有のレートには`fromDateField`を使用します

### パフォーマンスの考慮事項
- 複数の変換フィールドが並行して更新されます
- ソース値が変更されるときのみAPI呼び出しが行われます
- 同じ通貨の変換はAPI呼び出しをスキップします

## 一般的なユースケース

1. **多通貨プロジェクト**
   - プロジェクトコストを現地通貨で追跡
   - 会社通貨で総予算を報告
   - 地域間での値を比較

2. **国際販売**
   - 取引値を報告通貨に変換
   - 複数の通貨で収益を追跡
   - クローズした取引の履歴変換

3. **財務報告**
   - 期末の通貨変換
   - 統合財務諸表
   - 現地通貨での予算対実績

4. **契約管理**
   - 契約値を署名日で変換
   - 複数の通貨での支払いスケジュールを追跡
   - 通貨リスク評価

## 制限事項

- 暗号通貨の変換はサポートされていません
- 変換された値を手動で設定することはできません（常に計算されます）
- すべての変換金額に対して固定の2桁の精度
- カスタム為替レートはサポートされていません
- 為替レートのキャッシュはありません（各変換ごとに新しいAPI呼び出し）
- Frankfurter APIの可用性に依存します

## 関連リソース

- [通貨フィールド](/api/custom-fields/currency) - 変換のためのソースフィールド
- [日付フィールド](/api/custom-fields/date) - 日付ベースの変換のため
- [数式フィールド](/api/custom-fields/formula) - 代替計算
- [カスタムフィールドの概要](/custom-fields/list-custom-fields) - 一般的な概念