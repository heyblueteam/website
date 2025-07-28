---
title: 通貨カスタムフィールド
description: 適切なフォーマットとバリデーションで金銭的価値を追跡するための通貨フィールドを作成します
---

通貨カスタムフィールドを使用すると、関連する通貨コードとともに金銭的価値を保存および管理できます。このフィールドは、主要な法定通貨および暗号通貨を含む72種類の異なる通貨をサポートし、自動フォーマットおよびオプションの最小/最大制約を提供します。

## 基本例

シンプルな通貨フィールドを作成します：

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## 高度な例

バリデーション制約を持つ通貨フィールドを作成します：

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 通貨フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `CURRENCY` である必要があります |
| `currency` | String | いいえ | デフォルトの通貨コード（3文字のISOコード） |
| `min` | Float | いいえ | 最小許可値（保存されるが、更新時には強制されない） |
| `max` | Float | いいえ | 最大許可値（保存されるが、更新時には強制されない） |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: プロジェクトコンテキストは、認証から自動的に決定されます。フィールドを作成するプロジェクトにアクセスできる必要があります。

## 通貨値の設定

レコードに通貨値を設定または更新するには：

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | 通貨カスタムフィールドのID |
| `number` | Float! | ✅ はい | 金銭的な金額 |
| `currency` | String! | ✅ はい | 3文字の通貨コード |

## 通貨値を持つレコードの作成

通貨値を持つ新しいレコードを作成する場合：

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### 作成のための入力フォーマット

レコードを作成する際、通貨値は異なる方法で渡されます：

| パラメータ | 型 | 説明 |
|-----------|------|-------------|
| `customFieldId` | String! | 通貨フィールドのID |
| `value` | String! | 文字列としての金額（例："1500.50"） |
| `currency` | String! | 3文字の通貨コード |

## サポートされている通貨

Blueは、70の法定通貨と2の暗号通貨を含む72の通貨をサポートしています：

### 法定通貨

#### アメリカ大陸
| 通貨 | コード | 名前 |
|----------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | ベネズエラ・ボリバル・ソベラノ |
| ボリビア・ボリビアーノ | `BOB` | ボリビア・ボリビアーノ |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### ヨーロッパ
| 通貨 | コード | 名前 |
|----------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| ノルウェー・クローネ | `NOK` | ノルウェー・クローネ |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### アジア太平洋
| 通貨 | コード | 名前 |
|----------|------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### 中東およびアフリカ
| 通貨 | コード | 名前 |
|----------|------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### 暗号通貨
| 通貨 | コード |
|----------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `number` | Float | 金銭的な金額 |
| `currency` | String | 3文字の通貨コード |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

## 通貨フォーマット

システムはロケールに基づいて通貨値を自動的にフォーマットします：

- **シンボルの配置**: 通貨シンボルを正しく配置（前/後）
- **小数点区切り**: ロケール固有の区切りを使用（. または ,）
- **千の区切り**: 適切なグルーピングを適用
- **小数点以下の桁数**: 金額に基づいて0-2桁を表示
- **特別な処理**: USD/CADは明確さのために通貨コードの接頭辞を表示

### フォーマット例

| 値 | 通貨 | 表示 |
|-------|----------|---------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## バリデーション

### 金額のバリデーション
- 有効な数値である必要があります
- 最小/最大制約はフィールド定義と共に保存されますが、値の更新時には強制されません
- 表示のために最大2桁の小数をサポート（内部的には完全な精度が保存されます）

### 通貨コードのバリデーション
- 72のサポートされている通貨コードのいずれかである必要があります
- 大文字と小文字を区別します（大文字を使用）
- 無効なコードはエラーを返します

## 統合機能

### 数式
通貨フィールドは、計算のためにFORMULAカスタムフィールドで使用できます：
- 複数の通貨フィールドの合計
- パーセンテージの計算
- 算術演算の実行

### 通貨換算
CURRENCY_CONVERSIONフィールドを使用して、通貨間の自動換算を行います（[通貨換算フィールド](/api/custom-fields/currency-conversion)を参照）

### 自動化
通貨値は以下に基づいて自動化をトリガーできます：
- 金額の閾値
- 通貨の種類
- 値の変更

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**注意**: すべてのプロジェクトメンバーがカスタムフィールドを作成できますが、値を設定する能力は各フィールドに対して設定された役割ベースの権限に依存します。

## エラー応答

### 無効な通貨値
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

このエラーは、以下の場合に発生します：
- 通貨コードが72のサポートされているコードのいずれでもない
- 数字のフォーマットが無効
- 値を正しく解析できない

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

### 通貨の選択
- 主な市場に合ったデフォルトの通貨を設定します
- ISO 4217通貨コードを一貫して使用します
- デフォルトを選択する際にユーザーの所在地を考慮します

### 値の制約
- データ入力エラーを防ぐために合理的な最小/最大値を設定します
- 負の値を持たないフィールドには最小値として0を使用します
- 最大値を設定する際には使用ケースを考慮します

### 多通貨プロジェクト
- 報告用に一貫した基準通貨を使用します
- 自動換算のためにCURRENCY_CONVERSIONフィールドを実装します
- 各フィールドに使用すべき通貨を文書化します

## 一般的なユースケース

1. **プロジェクト予算**
   - プロジェクト予算の追跡
   - コスト見積もり
   - 費用の追跡

2. **販売と取引**
   - 取引の金額
   - 契約金額
   - 収益の追跡

3. **財務計画**
   - 投資金額
   - 資金調達ラウンド
   - 財務目標

4. **国際ビジネス**
   - 多通貨価格設定
   - 外国為替の追跡
   - 国境を越えた取引

## 制限事項

- 表示のための最大2桁の小数（ただし、より多くの精度が保存されます）
- 標準のCURRENCYフィールドには組み込みの通貨換算はありません
- 単一のフィールド値に通貨を混在させることはできません
- 自動為替レートの更新はありません（これにはCURRENCY_CONVERSIONを使用してください）
- 通貨シンボルはカスタマイズできません

## 関連リソース

- [通貨換算フィールド](/api/custom-fields/currency-conversion) - 自動通貨換算
- [数値フィールド](/api/custom-fields/number) - 非金銭的な数値のために
- [数式フィールド](/api/custom-fields/formula) - 通貨値で計算
- [リストカスタムフィールド](/api/custom-fields/list-custom-fields) - プロジェクト内のすべてのカスタムフィールドをクエリ