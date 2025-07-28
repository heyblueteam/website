---
title: シングルラインテキストカスタムフィールド
description: 名前、タイトル、ラベルなどの短いテキスト値のためのシングルラインテキストフィールドを作成します
---

シングルラインテキストカスタムフィールドは、シングルライン入力用に設計された短いテキスト値を保存することを可能にします。名前、タイトル、ラベル、または1行で表示されるべき任意のテキストデータに最適です。

## 基本的な例

シンプルなシングルラインテキストフィールドを作成します：

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きのシングルラインテキストフィールドを作成します：

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
| `name` | String! | ✅ はい | テキストフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `TEXT_SINGLE` である必要があります |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: プロジェクトコンテキストは、認証ヘッダーから自動的に決定されます。`projectId` パラメータは必要ありません。

## テキスト値の設定

レコードにシングルラインテキスト値を設定または更新するには：

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | テキストカスタムフィールドのID |
| `text` | String | いいえ | 保存するシングルラインテキスト内容 |

## テキスト値を持つレコードの作成

シングルラインテキスト値を持つ新しいレコードを作成する際：

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | ID! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールド定義（テキスト値を含む） |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

**重要**: テキスト値は、`customField.value.text` フィールドを通じてアクセスされ、TodoCustomField では直接アクセスされません。

## テキスト値のクエリ

テキストカスタムフィールドを持つレコードをクエリする際は、`customField.value.text` パスを通じてテキストにアクセスします：

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

レスポンスには、ネストされた構造内のテキストが含まれます：

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## テキスト検証

### フォーム検証
シングルラインテキストフィールドがフォームで使用される場合：
- 前後の空白は自動的にトリムされます
- フィールドが必須としてマークされている場合、必須検証が適用されます
- 特定のフォーマット検証は適用されません

### 検証ルール
- 改行を含む任意の文字列コンテンツを受け入れます（ただし推奨されません）
- 文字数制限はありません（データベースの制限まで）
- Unicode 文字と特殊記号をサポートします
- 改行は保持されますが、このフィールドタイプには意図されていません

### 一般的なテキストの例
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## 重要な注意事項

### ストレージ容量
- MySQL `MediumText` タイプを使用して保存されます
- 最大16MBのテキストコンテンツをサポートします
- マルチラインテキストフィールドと同じストレージ
- 国際文字用のUTF-8エンコーディング

### 直接APIとフォーム
- **フォーム**: 自動的な空白トリミングと必須検証
- **直接API**: テキストは提供された通りに保存されます
- **推奨**: 一貫したフォーマットを確保するためにユーザー入力にはフォームを使用してください

### TEXT_SINGLE と TEXT_MULTI
- **TEXT_SINGLE**: シングルラインテキスト入力、短い値に最適
- **TEXT_MULTI**: マルチラインテキストエリア入力、長いコンテンツに最適
- **バックエンド**: 両方は同じストレージと検証を使用します
- **フロントエンド**: データ入力用の異なるUIコンポーネント
- **意図**: TEXT_SINGLEは意味的にシングルライン値を意図しています

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## エラーレスポンス

### 必須フィールド検証（フォームのみ）
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

### コンテンツガイドライン
- テキストは簡潔でシングルラインに適したものにしてください
- 意図されたシングルライン表示のために改行を避けてください
- 類似のデータタイプに対して一貫したフォーマットを使用してください
- UI要件に基づいて文字数制限を考慮してください

### データ入力
- ユーザーを導くために明確なフィールド説明を提供してください
- 検証を確保するためにユーザー入力にはフォームを使用してください
- 必要に応じてアプリケーション内でコンテンツフォーマットを検証してください
- 標準化された値のためにドロップダウンを使用することを検討してください

### パフォーマンスの考慮事項
- シングルラインテキストフィールドは軽量でパフォーマンスが良好です
- 頻繁に検索されるフィールドにはインデックスを考慮してください
- UI内で適切な表示幅を使用してください
- 表示目的のためにコンテンツの長さを監視してください

## フィルタリングと検索

### 含む検索
シングルラインテキストフィールドは部分文字列検索をサポートします：

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### 検索機能
- 大文字と小文字を区別しない部分一致
- 部分的な単語一致をサポート
- 正確な値の一致
- フルテキスト検索やランキングはなし

## 一般的なユースケース

1. **識別子とコード**
   - 商品SKU
   - 注文番号
   - 参照コード
   - バージョン番号

2. **名前とタイトル**
   - クライアント名
   - プロジェクトタイトル
   - 商品名
   - カテゴリラベル

3. **短い説明**
   - 簡潔な要約
   - ステータスラベル
   - 優先度インジケーター
   - 分類タグ

4. **外部参照**
   - チケット番号
   - 請求書参照
   - 外部システムID
   - ドキュメント番号

## 統合機能

### ルックアップとの統合
- 他のレコードからテキストデータを参照
- テキストコンテンツでレコードを検索
- 関連するテキスト情報を表示
- 複数のソースからテキスト値を集約

### フォームとの統合
- 自動的な空白トリミング
- 必須フィールド検証
- シングルラインテキスト入力UI
- 文字数制限の表示（設定されている場合）

### インポート/エクスポートとの統合
- 直接CSV列マッピング
- 自動的なテキスト値の割り当て
- バルクデータインポートのサポート
- スプレッドシート形式へのエクスポート

## 制限事項

### 自動化制限
- 自動化トリガーフィールドとして直接利用できません
- 自動化フィールドの更新には使用できません
- 自動化条件で参照できます
- メールテンプレートやウェブフックで利用可能です

### 一般的な制限
- 組み込みのテキストフォーマットやスタイリングはありません
- 必須フィールドを超えた自動検証はありません
- 組み込みの一意性の強制はありません
- 非常に大きなテキストのコンテンツ圧縮はありません
- バージョン管理や変更追跡はありません
- 検索機能は制限されています（フルテキスト検索なし）

## 関連リソース

- [マルチラインテキストフィールド](/api/custom-fields/text-multi) - より長いテキストコンテンツ用
- [メールフィールド](/api/custom-fields/email) - メールアドレス用
- [URLフィールド](/api/custom-fields/url) - ウェブサイトアドレス用
- [ユニークIDフィールド](/api/custom-fields/unique-id) - 自動生成された識別子用
- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的な概念