---
title: URLカスタムフィールド
description: ウェブサイトのアドレスやリンクを保存するためのURLフィールドを作成します
---

URLカスタムフィールドを使用すると、レコードにウェブサイトのアドレスやリンクを保存できます。プロジェクトのウェブサイト、リファレンスリンク、ドキュメントのURL、または作業に関連するウェブベースのリソースを追跡するのに最適です。

## 基本的な例

シンプルなURLフィールドを作成します：

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きのURLフィールドを作成します：

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | URLフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `URL` である必要があります |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意:** `projectId` は、入力オブジェクトの一部ではなく、ミューテーションに別の引数として渡されます。

## URL値の設定

レコードにURL値を設定または更新するには：

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | URLカスタムフィールドのID |
| `text` | String! | ✅ はい | 保存するURLアドレス |

## URL値を持つレコードの作成

URL値を持つ新しいレコードを作成する際：

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `text` | String | 保存されたURLアドレス |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

## URLの検証

### 現在の実装
- **直接API**: 現在、URL形式の検証は強制されていません
- **フォーム**: URL検証は計画されていますが、現在はアクティブではありません
- **ストレージ**: 任意の文字列値をURLフィールドに保存できます

### 計画された検証
将来のバージョンには以下が含まれます：
- HTTP/HTTPSプロトコルの検証
- 有効なURL形式のチェック
- ドメイン名の検証
- 自動プロトコルプレフィックスの追加

### 推奨されるURL形式
現在強制されていませんが、以下の標準形式を使用してください：

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## 重要な注意事項

### ストレージ形式
- URLは変更なしのプレーンテキストとして保存されます
- 自動プロトコル追加なし (http://, https://)
- 入力された通りに大文字小文字が保持されます
- URLエンコーディング/デコーディングは行われません

### 直接APIとフォーム
- **フォーム**: 計画されたURL検証（現在はアクティブではありません）
- **直接API**: 検証なし - 任意のテキストを保存できます
- **推奨**: 保存する前にアプリケーション内でURLを検証してください

### URLとテキストフィールド
- **URL**: ウェブアドレス用に意味的に意図されています
- **TEXT_SINGLE**: 一般的な単一行テキスト
- **バックエンド**: 現在、ストレージと検証は同一です
- **フロントエンド**: データ入力用の異なるUIコンポーネント

## 必要な権限

カスタムフィールド操作はロールベースの権限を使用します：

| アクション | 必要なロール |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**注意:** 権限はプロジェクト内のユーザーロールに基づいてチェックされ、特定の権限定数ではありません。

## エラー応答

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

## ベストプラクティス

### URL形式の標準
- プロトコルを常に含める（http://またはhttps://）
- セキュリティのために可能な限りHTTPSを使用する
- 保存する前にURLをテストしてアクセス可能であることを確認する
- 表示目的で短縮URLの使用を検討する

### データ品質
- 保存する前にアプリケーション内でURLを検証する
- 一般的なタイプミス（プロトコルの欠落、不正確なドメイン）をチェックする
- 組織全体でURL形式を標準化する
- URLのアクセシビリティと可用性を考慮する

### セキュリティに関する考慮事項
- ユーザー提供のURLには注意する
- 特定のサイトに制限する場合はドメインを検証する
- 悪意のあるコンテンツのためにURLスキャンを考慮する
- 機密データを扱う際はHTTPS URLを使用する

## フィルタリングと検索

### 含む検索
URLフィールドは部分文字列検索をサポートします：

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### 検索機能
- 大文字小文字を区別しない部分文字列一致
- 部分ドメイン一致
- パスとパラメータの検索
- プロトコル特有のフィルタリングなし

## 一般的な使用例

1. **プロジェクト管理**
   - プロジェクトのウェブサイト
   - ドキュメントリンク
   - リポジトリのURL
   - デモサイト

2. **コンテンツ管理**
   - リファレンス資料
   - ソースリンク
   - メディアリソース
   - 外部記事

3. **カスタマーサポート**
   - 顧客のウェブサイト
   - サポートドキュメント
   - ナレッジベースの記事
   - ビデオチュートリアル

4. **営業およびマーケティング**
   - 会社のウェブサイト
   - 製品ページ
   - マーケティング資料
   - ソーシャルメディアプロフィール

## 統合機能

### ルックアップとの統合
- 他のレコードからのリファレンスURL
- ドメインまたはURLパターンによるレコードの検索
- 関連するウェブリソースの表示
- 複数のソースからのリンクの集約

### フォームとの統合
- URL特有の入力コンポーネント
- 適切なURL形式のための計画された検証
- リンクプレビュー機能（フロントエンド）
- クリック可能なURL表示

### レポートとの統合
- URLの使用状況とパターンを追跡
- 壊れたリンクやアクセスできないリンクを監視
- ドメインまたはプロトコルによる分類
- 分析のためのURLリストのエクスポート

## 制限事項

### 現在の制限
- アクティブなURL形式の検証なし
- 自動プロトコル追加なし
- リンクの検証やアクセシビリティチェックなし
- URLの短縮や展開なし
- faviconやプレビューの生成なし

### 自動化制限
- 自動化トリガーフィールドとしては利用できません
- 自動化フィールドの更新には使用できません
- 自動化条件で参照できます
- メールテンプレートやWebhookで利用可能です

### 一般的な制約
- ビルトインのリンクプレビューフィーチャーなし
- 自動URL短縮なし
- クリックトラッキングや分析なし
- URLの有効期限チェックなし
- 悪意のあるURLスキャンなし

## 将来の拡張

### 計画された機能
- HTTP/HTTPSプロトコルの検証
- カスタム正規表現検証パターン
- 自動プロトコルプレフィックスの追加
- URLのアクセシビリティチェック

### 潜在的な改善
- リンクプレビュー生成
- favicon表示
- URL短縮統合
- クリックトラッキング機能
- 壊れたリンク検出

## 関連リソース

- [テキストフィールド](/api/custom-fields/text-single) - 非URLテキストデータ用
- [メールフィールド](/api/custom-fields/email) - メールアドレス用
- [カスタムフィールドの概要](/api/custom-fields/2.list-custom-fields) - 一般的な概念

## テキストフィールドからの移行

テキストフィールドからURLフィールドに移行する場合：

1. **同じ名前と構成のURLフィールドを作成**
2. **既存のテキスト値をエクスポートして有効なURLであることを確認**
3. **レコードを更新して新しいURLフィールドを使用**
4. **成功した移行後に古いテキストフィールドを削除**
5. **アプリケーションを更新してURL特有のUIコンポーネントを使用**

### 移行例
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```