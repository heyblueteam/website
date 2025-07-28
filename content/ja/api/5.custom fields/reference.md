---
title: 参照カスタムフィールド
description: クロスプロジェクトの関係のために他のプロジェクトのレコードにリンクする参照フィールドを作成します
---

参照カスタムフィールドを使用すると、異なるプロジェクトのレコード間にリンクを作成でき、クロスプロジェクトの関係やデータ共有が可能になります。これは、組織のプロジェクト構造全体で関連する作業を接続するための強力な方法を提供します。

## 基本的な例

シンプルな参照フィールドを作成します：

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## 高度な例

フィルタリングと複数選択を持つ参照フィールドを作成します：

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 参照フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `REFERENCE` である必要があります |
| `referenceProjectId` | String | いいえ | 参照するプロジェクトのID |
| `referenceMultiple` | Boolean | いいえ | 複数のレコード選択を許可する（デフォルト：false） |
| `referenceFilter` | TodoFilterInput | いいえ | 参照されるレコードのフィルタ基準 |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注**: カスタムフィールドは、ユーザーの現在のプロジェクトコンテキストに基づいてプロジェクトに自動的に関連付けられます。

## 参照設定

### 単一参照と複数参照

**単一参照（デフォルト）:**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- ユーザーは参照プロジェクトから1つのレコードを選択できます
- 単一のTodoオブジェクトを返します

**複数参照:**
```graphql
{
  referenceMultiple: true
}
```
- ユーザーは参照プロジェクトから複数のレコードを選択できます
- Todoオブジェクトの配列を返します

### 参照フィルタリング

`referenceFilter` を使用して、選択できるレコードを制限します：

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## 参照値の設定

### 単一参照

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### 複数参照

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | 参照カスタムフィールドのID |
| `customFieldReferenceTodoIds` | [String!] | ✅ はい | 参照されるレコードIDの配列 |

## 参照を持つレコードの作成

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
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
| `customField` | CustomField! | 参照フィールドの定義 |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に変更された日時 |

**注**: 参照されたTodoは、`customField.selectedTodos`を介してアクセスされ、TodoCustomField上では直接アクセスされません。

### 参照されたTodoフィールド

各参照されたTodoには以下が含まれます：

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | ID! | 参照されたレコードの一意の識別子 |
| `title` | String! | 参照されたレコードのタイトル |
| `status` | TodoStatus! | 現在のステータス（ACTIVE、COMPLETEDなど） |
| `description` | String | 参照されたレコードの説明 |
| `dueDate` | DateTime | 設定されている場合の期限 |
| `assignees` | [User!] | 割り当てられたユーザー |
| `tags` | [Tag!] | 関連するタグ |
| `project` | Project! | 参照されたレコードを含むプロジェクト |

## 参照データのクエリ

### 基本クエリ

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### ネストされたデータを持つ高度なクエリ

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**重要**: ユーザーは、リンクされたレコードを見るために参照プロジェクトの表示権限を持っている必要があります。

## クロスプロジェクトアクセス

### プロジェクトの可視性

- ユーザーは、自分がアクセスできるプロジェクトからのみレコードを参照できます
- 参照されたレコードは、元のプロジェクトの権限を尊重します
- 参照されたレコードへの変更はリアルタイムで表示されます
- 参照されたレコードを削除すると、参照フィールドからも削除されます

### 権限の継承

- 参照フィールドは両方のプロジェクトから権限を継承します
- ユーザーは、参照プロジェクトへの表示アクセスが必要です
- 編集権限は、現在のプロジェクトのルールに基づいています
- 参照データは、参照フィールドのコンテキストでは読み取り専用です

## エラー応答

### 無効な参照プロジェクト

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### 参照されたレコードが見つかりません

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

### 権限が拒否されました

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## ベストプラクティス

### フィールド設計

1. **明確な命名** - 関係を示す説明的な名前を使用します
2. **適切なフィルタリング** - 関連するレコードのみを表示するようにフィルタを設定します
3. **権限を考慮する** - ユーザーが参照プロジェクトにアクセスできることを確認します
4. **関係を文書化する** - 接続の明確な説明を提供します

### パフォーマンスの考慮事項

1. **参照スコープを制限する** - フィルタを使用して選択可能なレコードの数を減らします
2. **深いネストを避ける** - 複雑な参照の連鎖を作成しないでください
3. **キャッシングを考慮する** - 参照データはパフォーマンスのためにキャッシュされます
4. **使用状況を監視する** - プロジェクト間で参照がどのように使用されているかを追跡します

### データの整合性

1. **削除を処理する** - 参照されたレコードが削除される場合の計画を立てます
2. **権限を検証する** - ユーザーが参照プロジェクトにアクセスできることを確認します
3. **依存関係を更新する** - 参照されたレコードを変更する際の影響を考慮します
4. **監査証跡** - 準拠のために参照関係を追跡します

## 一般的なユースケース

### プロジェクト依存関係

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### クライアント要件

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### リソース割り当て

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### 品質保証

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## ルックアップとの統合

参照フィールドは、[ルックアップフィールド](/api/custom-fields/lookup)と連携して、参照されたレコードからデータを取得します。ルックアップフィールドは、参照フィールドで選択されたレコードから値を抽出できますが、データ抽出専用であり（SUMなどの集計関数はサポートされていません）。

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## 制限事項

- 参照されたプロジェクトはユーザーがアクセスできる必要があります
- 参照されたプロジェクトの権限の変更は、参照フィールドのアクセスに影響します
- 参照の深いネストはパフォーマンスに影響を与える可能性があります
- 循環参照に対する組み込みの検証はありません
- 同じプロジェクトの参照を防ぐ自動制限はありません
- 参照値を設定する際のフィルタ検証は強制されません

## 関連リソース

- [ルックアップフィールド](/api/custom-fields/lookup) - 参照されたレコードからデータを抽出します
- [プロジェクトAPI](/api/projects) - 参照を含むプロジェクトの管理
- [レコードAPI](/api/records) - 参照を持つレコードの操作
- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的な概念