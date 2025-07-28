---
title: 電話カスタムフィールド
description: 国際フォーマットで電話番号を保存および検証するための電話フィールドを作成します
---

電話カスタムフィールドを使用すると、レコードに電話番号を保存し、組み込みの検証および国際フォーマットを使用できます。連絡先情報、緊急連絡先、またはプロジェクト内の電話関連データを追跡するのに最適です。

## 基本例

シンプルな電話フィールドを作成します：

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きの電話フィールドを作成します：

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ はい | 電話フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `PHONE` である必要があります |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: カスタムフィールドは、ユーザーの現在のプロジェクトコンテキストに基づいてプロジェクトに自動的に関連付けられます。`projectId` パラメータは必要ありません。

## 電話値の設定

レコードに電話値を設定または更新するには：

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | 電話カスタムフィールドのID |
| `text` | String | いいえ | 国コード付きの電話番号 |
| `regionCode` | String | いいえ | 国コード（自動検出） |

**注意**: `text` はスキーマでオプションですが、フィールドが意味を持つためには電話番号が必要です。`setTodoCustomField` を使用する場合、検証は行われません - 任意のテキスト値とregionCodeを保存できます。自動検出はレコード作成時にのみ行われます。

## 電話値を持つレコードの作成

電話値を持つ新しいレコードを作成する場合：

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      regionCode
    }
  }
}
```

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールド定義 |
| `text` | String | フォーマットされた電話番号（国際フォーマット） |
| `regionCode` | String | 国コード（例： "US", "GB", "CA"） |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に変更された日時 |

## 電話番号の検証

**重要**: 電話番号の検証とフォーマットは、`createTodo` を介して新しいレコードを作成する際にのみ行われます。`setTodoCustomField` を使用して既存の電話値を更新する場合、検証は行われず、値は提供された通りに保存されます。

### 受け入れられるフォーマット（レコード作成時）
電話番号には、次のいずれかのフォーマットで国コードを含める必要があります：

- **E.164フォーマット（推奨）**: `+12345678900`
- **国際フォーマット**: `+1 234 567 8900`
- **句読点付き国際フォーマット**: `+1 (234) 567-8900`
- **ダッシュ付き国コード**: `+1-234-567-8900`

**注意**: 国コードなしの国内フォーマット（例： `(234) 567-8900`）は、レコード作成時に拒否されます。

### 検証ルール（レコード作成時）
- libphonenumber-jsを使用して解析と検証を行います
- 様々な国際電話番号フォーマットを受け入れます
- 番号から国を自動的に検出します
- 国際表示フォーマットで番号をフォーマットします（例： `+1 234 567 8900`）
- 国コードを別々に抽出して保存します（例： `US`）

### 有効な電話の例
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### 無効な電話の例
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## ストレージフォーマット

電話番号を持つレコードを作成する際：
- **text**: 検証後に国際フォーマットで保存されます（例： `+1 234 567 8900`）
- **regionCode**: ISO国コードとして保存されます（例： `US`, `GB`, `CA`）自動的に検出されます

`setTodoCustomField` を介して更新する場合：
- **text**: 提供された通りに保存されます（フォーマットなし）
- **regionCode**: 提供された通りに保存されます（検証なし）

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## エラーレスポンス

### 無効な電話フォーマット
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### 国コードが欠落しています
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## ベストプラクティス

### データ入力
- 電話番号には常に国コードを含める
- 一貫性のためにE.164フォーマットを使用する
- 重要な操作のために保存する前に番号を検証する
- 表示フォーマットの地域的な好みを考慮する

### データ品質
- グローバル互換性のために国際フォーマットで番号を保存する
- 国特有の機能のためにregionCodeを使用する
- 重要な操作（SMS、通話）の前に電話番号を検証する
- 連絡タイミングのためにタイムゾーンの影響を考慮する

### 国際的考慮事項
- 国コードは自動的に検出され、保存される
- 番号は国際標準でフォーマットされる
- 表示の地域的な好みはregionCodeを使用できる
- 表示時にローカルダイヤル規則を考慮する

## 一般的なユースケース

1. **連絡先管理**
   - クライアントの電話番号
   - ベンダーの連絡先情報
   - チームメンバーの電話番号
   - サポートの連絡先詳細

2. **緊急連絡先**
   - 緊急連絡先番号
   - 当番の連絡先情報
   - 危機対応の連絡先
   - エスカレーション電話番号

3. **カスタマーサポート**
   - 顧客の電話番号
   - サポートのコールバック番号
   - 検証用電話番号
   - フォローアップの連絡先番号

4. **営業とマーケティング**
   - リードの電話番号
   - キャンペーンの連絡先リスト
   - パートナーの連絡先情報
   - 参照元の電話番号

## 統合機能

### 自動化との統合
- 電話フィールドが更新されたときにアクションをトリガーする
- 保存された電話番号にSMS通知を送信する
- 電話の変更に基づいてフォローアップタスクを作成する
- 電話番号データに基づいて通話をルーティングする

### ルックアップとの統合
- 他のレコードから電話データを参照する
- 複数のソースから電話リストを集約する
- 電話番号でレコードを見つける
- 連絡先情報をクロスリファレンスする

### フォームとの統合
- 自動電話検証
- 国際フォーマットのチェック
- 国コードの検出
- リアルタイムのフォーマットフィードバック

## 制限事項

- すべての番号に国コードが必要
- SMSや通話の機能は組み込まれていない
- フォーマットチェックを超えた電話番号の検証は行われない
- 電話メタデータ（キャリア、タイプなど）は保存されない
- 国コードなしの国内フォーマットの番号は拒否される
- UIでの国際標準を超えた自動電話番号フォーマットは行われない

## 関連リソース

- [テキストフィールド](/api/custom-fields/text-single) - 電話以外のテキストデータ用
- [メールフィールド](/api/custom-fields/email) - メールアドレス用
- [URLフィールド](/api/custom-fields/url) - ウェブサイトアドレス用
- [カスタムフィールドの概要](/custom-fields/list-custom-fields) - 一般的な概念