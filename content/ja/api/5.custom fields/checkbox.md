---
title: チェックボックスカスタムフィールド
description: はい/いいえまたは真/偽データのためのブールチェックボックスフィールドを作成する
---

チェックボックスカスタムフィールドは、タスクに対してシンプルなブール（真/偽）入力を提供します。これは、バイナリ選択、ステータスインジケーター、または何かが完了したかどうかを追跡するのに最適です。

## 基本例

シンプルなチェックボックスフィールドを作成します：

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明と検証を伴うチェックボックスフィールドを作成します：

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ はい | チェックボックスの表示名 |
| `type` | CustomFieldType! | ✅ はい | `CHECKBOX` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

## チェックボックス値の設定

タスクのチェックボックス値を設定または更新するには：

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

チェックボックスのチェックを外すには：

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するタスクのID |
| `customFieldId` | String! | ✅ はい | チェックボックスカスタムフィールドのID |
| `checked` | Boolean | いいえ | チェックするにはtrue、チェックを外すにはfalse |

## チェックボックス値を持つタスクの作成

チェックボックス値を持つ新しいタスクを作成する場合：

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### 受け入れられる文字列値

タスクを作成する際、チェックボックス値は文字列として渡す必要があります：

| 文字列値 | 結果 |
|--------------|---------|
| `"true"` | ✅ チェック済み（大文字と小文字を区別） |
| `"1"` | ✅ チェック済み |
| `"checked"` | ✅ チェック済み（大文字と小文字を区別） |
| Any other value | ❌ チェックなし |

**注意**: タスク作成時の文字列比較は大文字と小文字を区別します。値は`"true"`、`"1"`、または`"checked"`と正確に一致する必要があります。

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | ID! | フィールド値の一意の識別子 |
| `uid` | String! | 代替の一意の識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `checked` | Boolean | チェックボックスの状態（真/偽/ヌル） |
| `todo` | Todo! | この値が属するタスク |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

## 自動化統合

チェックボックスフィールドは、状態の変化に基づいて異なる自動化イベントをトリガーします：

| アクション | トリガーされるイベント | 説明 |
|--------|----------------|-------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | チェックボックスがチェックされたときにトリガーされます |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | チェックボックスのチェックが外されたときにトリガーされます |

これにより、チェックボックスの状態変化に応じて反応する自動化を作成できます。例えば：
- アイテムが承認されたときに通知を送信
- レビューのチェックボックスがチェックされたときにタスクを移動
- チェックボックスの状態に基づいて関連フィールドを更新

## データのインポート/エクスポート

### チェックボックス値のインポート

CSVや他の形式でデータをインポートする場合：
- `"true"`、`"yes"` → チェック済み（大文字と小文字を区別しない）
- その他の値（`"false"`、`"no"`、`"0"`、空） → チェックなし

### チェックボックス値のエクスポート

データをエクスポートする場合：
- チェック済みのボックスは`"X"`としてエクスポートされます
- チェックなしのボックスは空文字列`""`としてエクスポートされます

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## エラー応答

### 無効な値タイプ
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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

## ベストプラクティス

### 命名規則
- 明確でアクション指向の名前を使用する： "承認済み"、"レビュー済み"、"完了"
- ユーザーを混乱させる否定的な名前は避ける： "非アクティブ"よりも"アクティブ"を好む
- チェックボックスが何を表すかを具体的にする

### チェックボックスを使用するタイミング
- **バイナリ選択**：はい/いいえ、真/偽、完了/未完了
- **ステータスインジケーター**：承認済み、レビュー済み、公開済み
- **機能フラグ**：優先サポートあり、署名が必要
- **シンプルな追跡**：メール送信済み、請求書支払い済み、アイテム発送済み

### チェックボックスを使用しないタイミング
- 2つ以上のオプションが必要な場合（代わりにSELECT_SINGLEを使用）
- 数値またはテキストデータの場合（NUMBERまたはTEXTフィールドを使用）
- 誰がいつチェックしたかを追跡する必要がある場合（監査ログを使用）

## 一般的なユースケース

1. **承認ワークフロー**
   - "マネージャー承認"
   - "クライアントサインオフ"
   - "法務レビュー完了"

2. **タスク管理**
   - "ブロック中"
   - "レビューの準備完了"
   - "高優先度"

3. **品質管理**
   - "QA合格"
   - "ドキュメント完了"
   - "テスト作成済み"

4. **管理フラグ**
   - "請求書送信済み"
   - "契約署名済み"
   - "フォローアップが必要"

## 制限事項

- チェックボックスフィールドは真/偽の値のみを保存できます（初期設定後の三状態またはヌルは不可）
- デフォルト値の設定はできません（常にヌルから始まります）
- 誰がいつチェックしたかなどの追加メタデータを保存できません
- 他のフィールド値に基づく条件付き表示はありません

## 関連リソース

- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的なカスタムフィールドの概念
- [自動化API](/api/automations) - チェックボックスの変更によってトリガーされる自動化を作成