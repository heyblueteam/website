---
title: シングルセレクトカスタムフィールド
description: ユーザーが事前定義されたリストから1つのオプションを選択できるようにするシングルセレクトフィールドを作成します
---

シングルセレクトカスタムフィールドは、ユーザーが事前定義されたリストから正確に1つのオプションを選択できるようにします。これは、ステータスフィールド、カテゴリ、優先順位、または制御されたオプションセットから1つの選択のみを行う必要があるシナリオに最適です。

## 基本的な例

シンプルなシングルセレクトフィールドを作成します：

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高度な例

事前定義されたオプションを持つシングルセレクトフィールドを作成します：

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | シングルセレクトフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `SELECT_SINGLE` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | いいえ | フィールドの初期オプション |

### CreateCustomFieldOptionInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `title` | String! | ✅ はい | オプションの表示テキスト |
| `color` | String | いいえ | オプションの16進カラーコード |

## 既存フィールドへのオプションの追加

既存のシングルセレクトフィールドに新しいオプションを追加します：

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## シングルセレクト値の設定

レコードに選択されたオプションを設定するには：

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | シングルセレクトカスタムフィールドのID |
| `customFieldOptionId` | String | いいえ | 選択するオプションのID（シングルセレクトに推奨） |
| `customFieldOptionIds` | [String!] | いいえ | オプションIDの配列（シングルセレクトには最初の要素を使用） |

## シングルセレクト値のクエリ

レコードのシングルセレクト値をクエリします：

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

`value` フィールドは、選択されたオプションの詳細を含むJSONオブジェクトを返します。

## シングルセレクト値を持つレコードの作成

シングルセレクト値を持つ新しいレコードを作成する場合：

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
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
| `value` | JSON | ID、タイトル、色、位置を持つ選択されたオプションオブジェクトを含みます |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

### CustomFieldOption レスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | オプションの一意の識別子 |
| `title` | String! | オプションの表示テキスト |
| `color` | String | 視覚的表現のための16進カラーコード |
| `position` | Float | オプションの並び順 |
| `customField` | CustomField! | このオプションが属するカスタムフィールド |

### CustomField レスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | フィールドの一意の識別子 |
| `name` | String! | シングルセレクトフィールドの表示名 |
| `type` | CustomFieldType! | 常に `SELECT_SINGLE` |
| `description` | String | フィールドのヘルプテキスト |
| `customFieldOptions` | [CustomFieldOption!] | 利用可能なすべてのオプション |

## 値のフォーマット

### 入力フォーマット
- **APIパラメータ**: シングルオプションIDには `customFieldOptionId` を使用します
- **代替**: `customFieldOptionIds` 配列を使用します（最初の要素を取ります）
- **選択のクリア**: 両方のフィールドを省略するか、空の値を渡します

### 出力フォーマット
- **GraphQLレスポンス**: {id, title, color, position} を含む `value` フィールド内のJSONオブジェクト
- **アクティビティログ**: オプションタイトルを文字列として
- **自動化データ**: オプションタイトルを文字列として

## 選択の動作

### 排他的選択
- 新しいオプションを設定すると、以前の選択が自動的に削除されます
- 一度に選択できるオプションは1つだけです
- `null` または空の値を設定すると選択がクリアされます

### フォールバックロジック
- `customFieldOptionIds` 配列が提供されると、最初のオプションのみが使用されます
- これにより、マルチセレクト入力フォーマットとの互換性が確保されます
- 空の配列やnull値は選択をクリアします

## オプションの管理

### オプションプロパティの更新
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
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

**注意**: オプションを削除すると、それが選択されたすべてのレコードからクリアされます。

### オプションの並び替え
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## バリデーションルール

### オプションのバリデーション
- 提供されたオプションIDは存在しなければなりません
- オプションは指定されたカスタムフィールドに属していなければなりません
- 選択できるオプションは1つだけです（自動的に強制されます）
- null/空の値は有効です（選択なし）

### フィールドのバリデーション
- 使用可能にするためには、少なくとも1つのオプションが定義されている必要があります
- オプションタイトルはフィールド内で一意でなければなりません
- カラーコードは提供された場合、有効な16進形式でなければなりません

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## エラーレスポンス

### 無効なオプションID
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
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
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### 値を解析できません
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

## ベストプラクティス

### オプションデザイン
- 明確で説明的なオプションタイトルを使用します
- 意味のあるカラーコーディングを適用します
- オプションリストを焦点を絞って関連性を保ちます
- オプションを論理的に並べます（優先順位、頻度など）

### ステータスフィールドパターン
- プロジェクト全体で一貫したステータスワークフローを使用します
- オプションの自然な進行を考慮します
- 明確な最終状態を含めます（完了、キャンセルなど）
- オプションの意味を反映する色を使用します

### データ管理
- 定期的に未使用のオプションをレビューして整理します
- 一貫した命名規則を使用します
- オプション削除が既存のレコードに与える影響を考慮します
- オプションの更新と移行を計画します

## 一般的な使用例

1. **ステータスとワークフロー**
   - タスクステータス（未着手、進行中、完了）
   - 承認ステータス（保留、承認、拒否）
   - プロジェクトフェーズ（計画、開発、テスト、リリース）
   - 問題解決ステータス

2. **分類とカテゴライズ**
   - 優先度レベル（低、中、高、クリティカル）
   - タスクタイプ（バグ、機能、改善、文書）
   - プロジェクトカテゴリ（内部、クライアント、研究）
   - 部門割り当て

3. **品質と評価**
   - レビューステータス（未開始、レビュー中、承認）
   - 品質評価（不良、普通、良好、優秀）
   - リスクレベル（低、中、高）
   - 信頼レベル

4. **割り当てと所有権**
   - チーム割り当て
   - 部門所有権
   - 役割ベースの割り当て
   - 地域割り当て

## 統合機能

### 自動化との統合
- 特定のオプションが選択されたときにアクションをトリガーします
- 選択されたカテゴリに基づいて作業をルーティングします
- ステータス変更の通知を送信します
- 選択に基づいて条件付きワークフローを作成します

### ルックアップとの統合
- 選択されたオプションでレコードをフィルタリングします
- 他のレコードからオプションデータを参照します
- オプション選択に基づいてレポートを作成します
- 選択された値でレコードをグループ化します

### フォームとの統合
- ドロップダウン入力コントロール
- ラジオボタンインターフェース
- オプションのバリデーションとフィルタリング
- 選択に基づく条件付きフィールド表示

## アクティビティトラッキング

シングルセレクトフィールドの変更は自動的に追跡されます：
- 古いオプション選択と新しいオプション選択を表示します
- アクティビティログにオプションタイトルを表示します
- すべての選択変更のタイムスタンプ
- 修正のためのユーザー帰属

## マルチセレクトとの違い

| 機能 | シングルセレクト | マルチセレクト |
|---------|---------------|--------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## 制限事項

- 一度に選択できるオプションは1つだけです
- 階層的またはネストされたオプション構造はありません
- オプションはフィールドを使用するすべてのレコードで共有されます
- 組み込みのオプション分析や使用追跡はありません
- カラーコードは表示専用で、機能的な影響はありません
- オプションごとに異なる権限を設定することはできません

## 関連リソース

- [マルチセレクトフィールド](/api/custom-fields/select-multi) - 複数選択用
- [チェックボックスフィールド](/api/custom-fields/checkbox) - シンプルなブール選択用
- [テキストフィールド](/api/custom-fields/text-single) - 自由形式のテキスト入力用
- [カスタムフィールドの概要](/api/custom-fields/1.index) - 一般的な概念