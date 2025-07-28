---
title: メールカスタムフィールド
description: メールアドレスを保存および検証するためのメールフィールドを作成します
---

メールカスタムフィールドを使用すると、組み込みの検証機能を持つレコードにメールアドレスを保存できます。これは、連絡先情報、担当者のメール、またはプロジェクト内のメール関連データを追跡するのに最適です。

## 基本的な例

シンプルなメールフィールドを作成します：

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きのメールフィールドを作成します：

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | メールフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `EMAIL` である必要があります |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

## メール値の設定

レコードにメール値を設定または更新するには：

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | メールカスタムフィールドのID |
| `text` | String | いいえ | 保存するメールアドレス |

## メール値を持つレコードの作成

メール値を持つ新しいレコードを作成する際：

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## レスポンスフィールド

### CustomField レスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | ID! | カスタムフィールドの一意の識別子 |
| `name` | String! | メールフィールドの表示名 |
| `type` | CustomFieldType! | フィールドタイプ (EMAIL) |
| `description` | String | フィールドのヘルプテキスト |
| `value` | JSON | メール値を含む (以下を参照) |
| `createdAt` | DateTime! | フィールドが作成された日時 |
| `updatedAt` | DateTime! | フィールドが最後に変更された日時 |

**重要**: メール値は、`customField.value.text` フィールドを通じてアクセスされ、レスポンスに直接表示されません。

## メール値のクエリ

メールカスタムフィールドを持つレコードをクエリする際は、`customField.value.text` パスを通じてメールにアクセスします：

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

レスポンスには、ネストされた構造内のメールが含まれます：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## メール検証

### フォーム検証
メールフィールドがフォームで使用されると、自動的にメール形式を検証します：
- 標準のメール検証ルールを使用
- 入力から空白をトリム
- 無効なメール形式を拒否

### 検証ルール
- `@` シンボルを含む必要があります
- 有効なドメイン形式である必要があります
- 前後の空白は自動的に削除されます
- 一般的なメール形式が受け入れられます

### 有効なメールの例
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### 無効なメールの例
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## 重要な注意事項

### 直接APIとフォーム
- **フォーム**: 自動メール検証が適用されます
- **直接API**: 検証なし - 任意のテキストを保存できます
- **推奨**: ユーザー入力にはフォームを使用して検証を確実に行います

### ストレージ形式
- メールアドレスはプレーンテキストとして保存されます
- 特別なフォーマットやパースはありません
- 大文字小文字の区別: EMAILカスタムフィールドは大文字小文字を区別して保存されます (ユーザー認証メールは小文字に正規化されます)
- データベースの制約を超える最大長の制限はありません (16MB制限)

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## エラーレスポンス

### 無効なメール形式 (フォームのみ)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
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

## ベストプラクティス

### データ入力
- アプリケーション内で常にメールアドレスを検証します
- 実際のメールアドレスのみにメールフィールドを使用します
- 自動検証を得るためにユーザー入力にはフォームを使用することを検討してください

### データ品質
- 保存する前に空白をトリムします
- 大文字小文字の正規化を検討します (通常は小文字)
- 重要な操作の前にメール形式を検証します

### プライバシーに関する考慮事項
- メールアドレスはプレーンテキストとして保存されます
- データプライバシー規制 (GDPR、CCPA) を考慮します
- 適切なアクセス制御を実装します

## 一般的なユースケース

1. **連絡先管理**
   - クライアントのメールアドレス
   - ベンダーの連絡先情報
   - チームメンバーのメール
   - サポートの連絡先詳細

2. **プロジェクト管理**
   - ステークホルダーのメール
   - 承認連絡先のメール
   - 通知受信者
   - 外部コラボレーターのメール

3. **カスタマーサポート**
   - 顧客のメールアドレス
   - サポートチケットの連絡先
   - エスカレーションの連絡先
   - フィードバックのメールアドレス

4. **営業およびマーケティング**
   - リードのメールアドレス
   - キャンペーンの連絡先リスト
   - パートナーの連絡先情報
   - 紹介元のメール

## 統合機能

### 自動化との統合
- メールフィールドが更新されたときにアクションをトリガー
- 保存されたメールアドレスに通知を送信
- メールの変更に基づいてフォローアップタスクを作成

### ルックアップとの統合
- 他のレコードからメールデータを参照
- 複数のソースからメールリストを集約
- メールアドレスでレコードを検索

### フォームとの統合
- 自動メール検証
- メール形式のチェック
- 空白のトリミング

## 制限事項

- フォーマットチェックを超える組み込みのメール検証や検証はありません
- クリック可能なメールリンクのようなメール特有のUI機能はありません
- プレーンテキストとして保存され、暗号化はされていません
- メールの作成や送信機能はありません
- メールメタデータの保存 (表示名など) はありません
- 直接API呼び出しは検証をバイパスします (フォームのみが検証します)

## 関連リソース

- [テキストフィールド](/api/custom-fields/text-single) - 非メールのテキストデータ用
- [URLフィールド](/api/custom-fields/url) - ウェブサイトのアドレス用
- [電話フィールド](/api/custom-fields/phone) - 電話番号用
- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的な概念