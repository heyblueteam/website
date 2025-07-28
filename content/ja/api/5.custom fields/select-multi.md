---
title: マルチセレクトカスタムフィールド
description: ユーザーが事前定義されたリストから複数のオプションを選択できるようにするマルチセレクトフィールドを作成します
---

マルチセレクトカスタムフィールドは、ユーザーが事前定義されたリストから複数のオプションを選択できるようにします。カテゴリ、タグ、スキル、機能、または制御されたオプションセットから複数の選択が必要なシナリオに最適です。

## 基本例

シンプルなマルチセレクトフィールドを作成します：

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高度な例

マルチセレクトフィールドを作成し、オプションを別々に追加します：

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | マルチセレクトフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `SELECT_MULTI` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |
| `projectId` | String! | ✅ はい | このフィールドのプロジェクトのID |

### CreateCustomFieldOptionInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ はい | カスタムフィールドのID |
| `title` | String! | ✅ はい | オプションの表示テキスト |
| `color` | String | いいえ | オプションの色（任意の文字列） |
| `position` | Float | いいえ | オプションのソート順 |

## 既存フィールドへのオプション追加

既存のマルチセレクトフィールドに新しいオプションを追加します：

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## マルチセレクト値の設定

レコードに複数の選択されたオプションを設定するには：

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | マルチセレクトカスタムフィールドのID |
| `customFieldOptionIds` | [String!] | ✅ はい | 選択するオプションIDの配列 |

## マルチセレクト値を持つレコードの作成

マルチセレクト値を持つ新しいレコードを作成する場合：

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
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
| `selectedOptions` | [CustomFieldOption!] | 選択されたオプションの配列 |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

### CustomFieldOption レスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | オプションの一意の識別子 |
| `title` | String! | オプションの表示テキスト |
| `color` | String | 視覚的表現のための16進数カラーコード |
| `position` | Float | オプションのソート順 |
| `customField` | CustomField! | このオプションが属するカスタムフィールド |

### CustomField レスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | フィールドの一意の識別子 |
| `name` | String! | マルチセレクトフィールドの表示名 |
| `type` | CustomFieldType! | 常に `SELECT_MULTI` |
| `description` | String | フィールドのヘルプテキスト |
| `customFieldOptions` | [CustomFieldOption!] | 利用可能なすべてのオプション |

## 値のフォーマット

### 入力フォーマット
- **APIパラメータ**: オプションIDの配列 (`["option1", "option2", "option3"]`)
- **文字列フォーマット**: カンマ区切りのオプションID (`"option1,option2,option3"`)

### 出力フォーマット
- **GraphQLレスポンス**: CustomFieldOptionオブジェクトの配列
- **アクティビティログ**: カンマ区切りのオプションタイトル
- **自動化データ**: オプションタイトルの配列

## オプションの管理

### オプションプロパティの更新
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
  }) {
    id
    title
    color
  }
}
```

### オプションの削除
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### オプションの再順序付け
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## バリデーションルール

### オプションのバリデーション
- 提供されたすべてのオプションIDは存在しなければなりません
- オプションは指定されたカスタムフィールドに属する必要があります
- SELECT_MULTIフィールドのみが複数のオプションを選択できます
- 空の配列は有効です（選択なし）

### フィールドのバリデーション
- 使用可能にするためには、少なくとも1つのオプションが定義されている必要があります
- オプションタイトルはフィールド内で一意でなければなりません
- カラーフィールドは任意の文字列値を受け入れます（16進数のバリデーションなし）

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## エラーレスポンス

### 無効なオプションID
```json
{
  "errors": [{
    "message": "Custom field option not found",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### オプションがフィールドに属していない
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### フィールドが見つからない
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### 非マルチフィールドでの複数オプション
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## ベストプラクティス

### オプションデザイン
- 説明的で簡潔なオプションタイトルを使用する
- 一貫したカラコーディングスキームを適用する
- オプションリストを管理可能に保つ（通常は3〜20オプション）
- オプションを論理的に順序付ける（アルファベット順、頻度順など）

### データ管理
- 定期的に未使用のオプションをレビューしてクリーンアップする
- プロジェクト間で一貫した命名規則を使用する
- フィールドを作成する際にオプションの再利用性を考慮する
- オプションの更新と移行を計画する

### ユーザーエクスペリエンス
- 明確なフィールド説明を提供する
- 視覚的な区別を改善するために色を使用する
- 関連するオプションをグループ化する
- 一般的なケースのためのデフォルト選択を考慮する

## 一般的なユースケース

1. **プロジェクト管理**
   - タスクのカテゴリやタグ
   - 優先度レベルやタイプ
   - チームメンバーの割り当て
   - ステータスインジケーター

2. **コンテンツ管理**
   - 記事のカテゴリやトピック
   - コンテンツタイプやフォーマット
   - 出版チャネル
   - 承認ワークフロー

3. **カスタマーサポート**
   - 問題のカテゴリやタイプ
   - 影響を受ける製品やサービス
   - 解決方法
   - 顧客セグメント

4. **製品開発**
   - 機能のカテゴリ
   - 技術要件
   - テスト環境
   - リリースチャネル

## 統合機能

### 自動化との統合
- 特定のオプションが選択されたときにアクションをトリガーする
- 選択されたカテゴリに基づいて作業をルーティングする
- 高優先度の選択に対して通知を送信する
- オプションの組み合わせに基づいてフォローアップタスクを作成する

### ルックアップとの統合
- 選択されたオプションでレコードをフィルタリングする
- オプション選択に基づいてデータを集約する
- 他のレコードからオプションデータを参照する
- オプションの組み合わせに基づいてレポートを作成する

### フォームとの統合
- マルチセレクト入力コントロール
- オプションのバリデーションとフィルタリング
- 動的オプションの読み込み
- 条件付きフィールド表示

## アクティビティトラッキング

マルチセレクトフィールドの変更は自動的に追跡されます：
- 追加されたオプションと削除されたオプションを表示
- アクティビティログにオプションタイトルを表示
- すべての選択変更のタイムスタンプ
- 修正のユーザー帰属

## 制限事項

- オプションの最大実用限界はUIパフォーマンスに依存します
- 階層的またはネストされたオプション構造はありません
- オプションはフィールドを使用するすべてのレコードで共有されます
- 組み込みのオプション分析や使用状況追跡はありません
- カラーフィールドは任意の文字列を受け入れます（16進数のバリデーションなし）
- オプションごとに異なる権限を設定することはできません
- オプションはフィールド作成と同時に作成することはできません
- 専用の再順序付けミューテーションはありません（positionを使用してeditCustomFieldOptionを使用）

## 関連リソース

- [シングルセレクトフィールド](/api/custom-fields/select-single) - 単一選択のため
- [チェックボックスフィールド](/api/custom-fields/checkbox) - シンプルなブール選択のため
- [テキストフィールド](/api/custom-fields/text-single) - 自由形式のテキスト入力のため
- [カスタムフィールドの概要](/api/custom-fields/2.list-custom-fields) - 一般的な概念