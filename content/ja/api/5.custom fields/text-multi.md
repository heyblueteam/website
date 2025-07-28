---
title: マルチラインテキストカスタムフィールド
description: 説明、ノート、コメントなどの長いコンテンツ用のマルチラインテキストフィールドを作成します
---

マルチラインテキストカスタムフィールドを使用すると、改行やフォーマットを含む長いテキストコンテンツを保存できます。これは、説明、ノート、コメント、または複数行が必要な任意のテキストデータに最適です。

## 基本的な例

シンプルなマルチラインテキストフィールドを作成します：

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きのマルチラインテキストフィールドを作成します：

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `name` | String! | ✅ はい | テキストフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `TEXT_MULTI` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意:** `projectId` は、入力オブジェクトの一部ではなく、ミューテーションに別の引数として渡されます。あるいは、プロジェクトコンテキストは、GraphQLリクエスト内の `X-Bloo-Project-ID` ヘッダーから決定できます。

## テキスト値の設定

レコードにマルチラインテキスト値を設定または更新するには：

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | テキストカスタムフィールドのID |
| `text` | String | いいえ | 保存するマルチラインテキストコンテンツ |

## テキスト値を持つレコードの作成

マルチラインテキスト値を持つ新しいレコードを作成する場合：

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
    }
  }
}
```

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `text` | String | 保存されたマルチラインテキストコンテンツ |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

## テキスト検証

### フォーム検証
マルチラインテキストフィールドがフォームで使用される場合：
- 前後の空白は自動的にトリムされます
- フィールドが必須としてマークされている場合、必須検証が適用されます
- 特定のフォーマット検証は適用されません

### 検証ルール
- 改行を含む任意の文字列コンテンツを受け入れます
- 文字数制限はありません（データベースの制限まで）
- Unicode文字と特殊記号をサポートします
- 改行はストレージに保存されます

### 有効なテキストの例
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## 重要な注意事項

### ストレージ容量
- MySQL `MediumText` タイプを使用して保存
- 最大16MBのテキストコンテンツをサポート
- 改行とフォーマットが保持されます
- 国際文字用のUTF-8エンコーディング

### 直接APIとフォーム
- **フォーム**: 自動的な空白トリミングと必須検証
- **直接API**: テキストは提供された通りに保存されます
- **推奨**: 一貫したフォーマットを確保するためにユーザー入力にはフォームを使用してください

### TEXT_MULTIとTEXT_SINGLE
- **TEXT_MULTI**: マルチラインテキストエリア入力、長いコンテンツに最適
- **TEXT_SINGLE**: シングルラインテキスト入力、短い値に最適
- **バックエンド**: 両方のタイプは同一 - 同じストレージフィールド、検証、処理
- **フロントエンド**: データ入力のための異なるUIコンポーネント（テキストエリアと入力フィールド）
- **重要**: TEXT_MULTIとTEXT_SINGLEの区別はUI目的のためだけに存在します

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

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
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## ベストプラクティス

### コンテンツの整理
- 構造化されたコンテンツのために一貫したフォーマットを使用
- 読みやすさのためにマークダウンのような構文を使用することを検討
- 長いコンテンツを論理的なセクションに分ける
- 読みやすさを向上させるために改行を使用

### データ入力
- ユーザーをガイドするために明確なフィールド説明を提供
- 検証を確保するためにユーザー入力にはフォームを使用
- 使用ケースに基づいて文字数制限を考慮
- 必要に応じてアプリケーション内でコンテンツフォーマットを検証

### パフォーマンスの考慮事項
- 非常に長いテキストコンテンツはクエリパフォーマンスに影響を与える可能性があります
- 大きなテキストフィールドを表示するためにページネーションを検討
- 検索機能のためのインデックスの考慮
- 大きなコンテンツを持つフィールドのストレージ使用量を監視

## フィルタリングと検索

### 含む検索
マルチラインテキストフィールドはカスタムフィールドフィルターを介して部分文字列検索をサポートします：

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### 検索機能
- `CONTAINS` 演算子を使用したテキストフィールド内の部分文字列一致
- `NCONTAINS` 演算子を使用した大文字小文字を区別しない検索
- `IS` 演算子を使用した完全一致
- `NOT` 演算子を使用した否定一致
- テキストのすべての行を横断して検索
- 部分単語一致をサポート

## 一般的な使用例

1. **プロジェクト管理**
   - タスクの説明
   - プロジェクトの要件
   - 会議のノート
   - ステータス更新

2. **カスタマーサポート**
   - 問題の説明
   - 解決ノート
   - 顧客のフィードバック
   - コミュニケーションログ

3. **コンテンツ管理**
   - 記事コンテンツ
   - 製品説明
   - ユーザーコメント
   - レビューの詳細

4. **ドキュメンテーション**
   - プロセスの説明
   - 指示
   - ガイドライン
   - 参考資料

## 統合機能

### 自動化との統合
- テキストコンテンツが変更されたときにアクションをトリガー
- テキストコンテンツからキーワードを抽出
- 要約や通知を作成
- 外部サービスでテキストコンテンツを処理

### ルックアップとの統合
- 他のレコードからテキストデータを参照
- 複数のソースからテキストコンテンツを集約
- テキストコンテンツでレコードを検索
- 関連するテキスト情報を表示

### フォームとの統合
- 自動的な空白トリミング
- 必須フィールドの検証
- マルチラインテキストエリアUI
- 文字数表示（設定されている場合）

## 制限事項

- 組み込みのテキストフォーマットやリッチテキスト編集はありません
- 自動リンク検出や変換はありません
- スペルチェックや文法検証はありません
- 組み込みのテキスト分析や処理はありません
- バージョン管理や変更追跡はありません
- 検索機能は制限されています（全文検索はありません）
- 非常に大きなテキストのコンテンツ圧縮はありません

## 関連リソース

- [シングルラインテキストフィールド](/api/custom-fields/text-single) - 短いテキスト値用
- [メールフィールド](/api/custom-fields/email) - メールアドレス用
- [URLフィールド](/api/custom-fields/url) - ウェブサイトアドレス用
- [カスタムフィールドの概要](/api/custom-fields/2.list-custom-fields) - 一般的な概念