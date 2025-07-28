---
title: カスタムフィールドの参照
description: 参照されたレコードから自動的にデータを取得する参照フィールドを作成します
---

カスタムフィールドの参照は、[参照フィールド](/api/custom-fields/reference)によって参照されたレコードから自動的にデータを取得し、手動でのコピーなしにリンクされたレコードの情報を表示します。参照データが変更されると、自動的に更新されます。

## 基本的な例

参照されたレコードからタグを表示する参照フィールドを作成します：

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 高度な例

参照されたレコードからカスタムフィールドの値を抽出する参照フィールドを作成します：

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | 参照フィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `LOOKUP` でなければなりません |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ はい | 参照設定 |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

## 参照設定

### CustomFieldLookupOptionInput

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `referenceId` | String! | ✅ はい | データを取得するための参照フィールドのID |
| `lookupId` | String | いいえ | 参照する特定のカスタムフィールドのID（TODO_CUSTOM_FIELDタイプの場合は必須） |
| `lookupType` | CustomFieldLookupType! | ✅ はい | 参照されたレコードから抽出するデータの型 |

## 参照タイプ

### CustomFieldLookupType 値

| 型 | 説明 | 戻り値 |
|------|-------------|---------|
| `TODO_DUE_DATE` | 参照されたタスクの期限 | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | 参照されたタスクの作成日 | Array of creation timestamps |
| `TODO_UPDATED_AT` | 参照されたタスクの最終更新日 | Array of update timestamps |
| `TODO_TAG` | 参照されたタスクのタグ | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | 参照されたタスクの担当者 | Array of user objects |
| `TODO_DESCRIPTION` | 参照されたタスクの説明 | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | 参照されたタスクのタスクリスト名 | Array of list titles |
| `TODO_CUSTOM_FIELD` | 参照されたタスクのカスタムフィールドの値 | Array of values based on the field type |

## 応答フィールド

### CustomField 応答（参照フィールド用）

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールドの一意の識別子 |
| `name` | String! | 参照フィールドの表示名 |
| `type` | CustomFieldType! | `LOOKUP` になります |
| `customFieldLookupOption` | CustomFieldLookupOption | 参照設定と結果 |
| `createdAt` | DateTime! | フィールドが作成された日時 |
| `updatedAt` | DateTime! | フィールドが最後に更新された日時 |

### CustomFieldLookupOption 構造

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | 実行されている参照の型 |
| `lookupResult` | JSON | 参照されたレコードから抽出されたデータ |
| `reference` | CustomField | ソースとして使用される参照フィールド |
| `lookup` | CustomField | 参照される特定のフィールド（TODO_CUSTOM_FIELD用） |
| `parentCustomField` | CustomField | 親の参照フィールド |
| `parentLookup` | CustomField | チェーン内の親参照（ネストされた参照用） |

## 参照の動作

1. **データ抽出**: 参照は、参照フィールドを通じてリンクされたすべてのレコードから特定のデータを抽出します
2. **自動更新**: 参照されたレコードが変更されると、参照値が自動的に更新されます
3. **読み取り専用**: 参照フィールドは直接編集できず、常に現在の参照データを反映します
4. **計算なし**: 参照はデータをそのまま抽出して表示し、集計や計算は行いません

## TODO_CUSTOM_FIELD 参照

`TODO_CUSTOM_FIELD` タイプを使用する場合、`lookupId` パラメータを使用して抽出するカスタムフィールドを指定する必要があります：

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

これにより、指定されたカスタムフィールドの値がすべての参照されたレコードから抽出されます。

## 参照データのクエリ

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## 例の参照結果

### タグ参照結果
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### 担当者参照結果
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### カスタムフィールド参照結果
結果は、参照されるカスタムフィールドの型に基づいて異なります。たとえば、通貨フィールドの参照は次のようになります：
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**重要**: ユーザーは、参照結果を表示するために、現在のプロジェクトと参照されたプロジェクトの両方に対する表示権限を持っている必要があります。

## エラー応答

### 無効な参照フィールド
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

### 循環参照が検出されました
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### TODO_CUSTOM_FIELDの参照IDが不足しています
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## ベストプラクティス

1. **明確な命名**: どのデータが参照されているかを示す説明的な名前を使用します
2. **適切な型**: データニーズに合った参照タイプを選択します
3. **パフォーマンス**: 参照はすべての参照されたレコードを処理するため、多くのリンクを持つ参照フィールドには注意が必要です
4. **権限**: 参照が機能するために、ユーザーが参照されたプロジェクトにアクセスできることを確認します

## 一般的なユースケース

### プロジェクト間の可視性
手動での同期なしに、関連プロジェクトからタグ、担当者、またはステータスを表示します。

### 依存関係の追跡
現在の作業が依存しているタスクの期限や完了状況を表示します。

### リソースの概要
リソース計画のために、参照されたタスクに割り当てられたすべてのチームメンバーを表示します。

### ステータスの集約
関連タスクからすべてのユニークなステータスを収集し、プロジェクトの健康状態を一目で確認します。

## 制限事項

- 参照フィールドは読み取り専用で、直接編集することはできません
- 集計関数（SUM、COUNT、AVG）はありません - 参照はデータを抽出するだけです
- フィルタリングオプションはありません - すべての参照されたレコードが含まれます
- 無限ループを避けるために、循環参照チェーンは防止されます
- 結果は現在のデータを反映し、自動的に更新されます

## 関連リソース

- [参照フィールド](/api/custom-fields/reference) - 参照ソースのためのレコードへのリンクを作成します
- [カスタムフィールドの値](/api/custom-fields/custom-field-values) - 編集可能なカスタムフィールドに値を設定します
- [カスタムフィールドの一覧](/api/custom-fields/list-custom-fields) - プロジェクト内のすべてのカスタムフィールドをクエリします