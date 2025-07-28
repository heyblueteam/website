---
title: 国別カスタムフィールド
description: ISO国コードの検証を伴う国の選択フィールドを作成する
---

国別カスタムフィールドは、レコードの国情報を保存および管理するためのものです。このフィールドは、国名とISO Alpha-2国コードの両方をサポートしています。

**重要**: 国の検証および変換の動作は、ミューテーションによって大きく異なります：
- **createTodo**: 国名を自動的に検証およびISOコードに変換します
- **setTodoCustomField**: 検証なしで任意の値を受け入れます

## 基本的な例

シンプルな国フィールドを作成します：

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きの国フィールドを作成します：

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 国フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `COUNTRY` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: `projectId` は入力として渡されず、GraphQLコンテキスト（通常はリクエストヘッダーや認証から）によって決定されます。

## 国の値を設定する

国フィールドは、2つのデータベースフィールドにデータを保存します：
- **`countryCodes`**: ISO Alpha-2国コードをカンマ区切りの文字列としてデータベースに保存します（APIを介して配列として返されます）
- **`text`**: 表示テキストまたは国名を文字列として保存します

### パラメータの理解

`setTodoCustomField` ミューテーションは、国フィールド用に2つのオプションのパラメータを受け入れます：

| パラメータ | 型 | 必須 | 説明 | 何をするか |
|-----------|------|----------|-------------|--------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID | - |
| `customFieldId` | String! | ✅ はい | 国カスタムフィールドのID | - |
| `countryCodes` | [String!] | いいえ | ISO Alpha-2国コードの配列 | Stored in the `countryCodes` field |
| `text` | String | いいえ | 表示テキストまたは国名 | Stored in the `text` field |

**重要**: 
- `setTodoCustomField` では、両方のパラメータはオプションであり、独立して保存されます
- `createTodo` では、システムが自動的に入力に基づいて両方のフィールドを設定します（独立して制御することはできません）

### オプション1: 国コードのみを使用

表示テキストなしで検証されたISOコードを保存します：

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

結果: `countryCodes` = `["US"]`, `text` = `null`

### オプション2: テキストのみを使用

検証されたコードなしで表示テキストを保存します：

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

結果: `countryCodes` = `null`, `text` = `"United States"`

**注意**: `setTodoCustomField` を使用する場合、どのパラメータを使用しても検証は行われません。値は提供された通りに保存されます。

### オプション3: 両方を使用（推奨）

検証されたコードと表示テキストの両方を保存します：

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

結果: `countryCodes` = `["US"]`, `text` = `"United States"`

### 複数の国

配列を使用して複数の国を保存します：

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## 国の値を持つレコードの作成

レコードを作成する際、`createTodo` ミューテーションは**自動的に国の値を検証および変換**します。これは国の検証を行う唯一のミューテーションです：

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      text
      countryCodes
    }
  }
}
```

### 受け入れられる入力形式

| 入力タイプ | 例 | 結果 |
|------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **サポートされていません** - 単一の無効な値として扱われます |
| Mixed format | `"United States, CA"` | **サポートされていません** - 単一の無効な値として扱われます |

## レスポンスフィールド

### TodoCustomFieldレスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `text` | String | 表示テキスト（国名） |
| `countryCodes` | [String!] | ISO Alpha-2国コードの配列 |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

## 国の基準

Blueは、**ISO 3166-1 Alpha-2**標準を国コードに使用しています：

- 2文字の国コード（例：US、GB、FR、DE）
- `i18n-iso-countries`ライブラリを使用した検証は**createTodo**でのみ行われます
- すべての公式に認められた国をサポートします

### 例国コード

| 国 | ISOコード |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

ISO 3166-1 alpha-2国コードの公式リストの完全なリストについては、[ISOオンラインブラウジングプラットフォーム](https://www.iso.org/obp/ui/#search/code/)をご覧ください。

## 検証

**検証は`createTodo`ミューテーションでのみ行われます**：

1. **有効なISOコード**: 有効なISO Alpha-2コードを受け入れます
2. **国名**: 認識された国名を自動的にコードに変換します
3. **無効な入力**: 認識されていない値に対して`CustomFieldValueParseError`をスローします

**注意**: `setTodoCustomField`ミューテーションは検証を行わず、任意の文字列値を受け入れます。

### エラー例

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 統合機能

### ルックアップフィールド
国フィールドはLOOKUPカスタムフィールドによって参照され、関連レコードから国データを取得できます。

### 自動化
自動化条件で国の値を使用します：
- 特定の国によるアクションのフィルタリング
- 国に基づいた通知の送信
- 地理的地域に基づくタスクのルーティング

### フォーム
フォーム内の国フィールドは、自動的にユーザー入力を検証し、国名をコードに変換します。

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## エラー応答

### 無効な国の値
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### フィールドタイプの不一致
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## ベストプラクティス

### 入力処理
- 自動検証と変換のために`createTodo`を使用します
- 検証をバイパスするため、`setTodoCustomField`を注意深く使用します
- `setTodoCustomField`を使用する前に、アプリケーション内で入力を検証することを検討してください
- 明確さのためにUIで完全な国名を表示します

### データ品質
- 入力ポイントで国の入力を検証します
- システム全体で一貫した形式を使用します
- レポートのために地域グループを考慮します

### 複数の国
- 複数の国のために`setTodoCustomField`で配列サポートを使用します
- `createTodo`の値フィールドでは複数の国は**サポートされていません**
- 適切な処理のために`setTodoCustomField`に国コードを配列として保存します

## 一般的なユースケース

1. **顧客管理**
   - 顧客本社の所在地
   - 配送先
   - 税管轄区域

2. **プロジェクト追跡**
   - プロジェクトの所在地
   - チームメンバーの所在地
   - 市場ターゲット

3. **コンプライアンスと法務**
   - 規制管轄区域
   - データ居住要件
   - 輸出管理

4. **営業とマーケティング**
   - テリトリーの割り当て
   - 市場セグメンテーション
   - キャンペーンターゲティング

## 制限事項

- ISO 3166-1 Alpha-2コード（2文字コード）のみをサポート
- 国の下位区分（州/県）に対する組み込みサポートなし
- 自動国旗アイコンなし（テキストベースのみ）
- 歴史的国コードを検証できない
- 組み込みの地域または大陸グループなし
- **検証は`createTodo`でのみ機能し、`setTodoCustomField`では機能しません**
- **`createTodo`の値フィールドでは複数の国はサポートされていません**
- **国コードはカンマ区切りの文字列として保存され、真の配列ではありません**

## 関連リソース

- [カスタムフィールドの概要](/custom-fields/list-custom-fields) - 一般的なカスタムフィールドの概念
- [ルックアップフィールド](/api/custom-fields/lookup) - 他のレコードから国データを参照
- [フォームAPI](/api/forms) - カスタムフォームに国フィールドを含める