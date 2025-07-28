---
title: ボタンカスタムフィールド
description: クリック時に自動化をトリガーするインタラクティブなボタンフィールドを作成する
---

ボタンカスタムフィールドは、クリック時に自動化をトリガーするインタラクティブなUI要素を提供します。他のカスタムフィールドタイプがデータを保存するのに対し、ボタンフィールドは構成されたワークフローを実行するためのアクショントリガーとして機能します。

## 基本的な例

自動化をトリガーするシンプルなボタンフィールドを作成します：

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高度な例

確認要件を持つボタンを作成します：

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | ボタンの表示名 |
| `type` | CustomFieldType! | ✅ はい | `BUTTON` である必要があります |
| `projectId` | String! | ✅ はい | フィールドが作成されるプロジェクトID |
| `buttonType` | String | いいえ | 確認動作（下記のボタンタイプを参照） |
| `buttonConfirmText` | String | いいえ | 確認のためにユーザーが入力する必要があるテキスト |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |
| `required` | Boolean | いいえ | フィールドが必須かどうか（デフォルトはfalse） |
| `isActive` | Boolean | いいえ | フィールドがアクティブかどうか（デフォルトはtrue） |

### ボタンタイプフィールド

`buttonType` フィールドは、UIクライアントが確認動作を決定するために使用できる自由形式の文字列です。一般的な値には以下が含まれます：

- `""`（空） - 確認なし
- `"soft"` - シンプルな確認ダイアログ
- `"hard"` - 確認テキストの入力を要求

**注意**: これらはUIのヒントに過ぎません。APIは特定の値を検証または強制しません。

## ボタンクリックのトリガー

ボタンクリックをトリガーし、関連する自動化を実行するには：

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### クリック入力パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | ボタンを含むタスクのID |
| `customFieldId` | String! | ✅ はい | ボタンカスタムフィールドのID |

### 重要: APIの動作

**APIを介したすべてのボタンクリックは即座に実行されます** `buttonType` または `buttonConfirmText` 設定に関係なく。これらのフィールドはUIクライアントが確認ダイアログを実装するために保存されますが、API自体は：

- 確認テキストを検証しません
- 確認要件を強制しません
- 呼び出されたときにボタンアクションを即座に実行します

確認は純粋にクライアント側のUI安全機能です。

### 例: 異なるボタンタイプのクリック

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

上記の3つのミューテーションはすべて、APIを介して呼び出されたときにボタンアクションを即座に実行し、確認要件をバイパスします。

## レスポンスフィールド

### カスタムフィールドレスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | カスタムフィールドの一意の識別子 |
| `name` | String! | ボタンの表示名 |
| `type` | CustomFieldType! | ボタンフィールドの場合は常に `BUTTON` |
| `buttonType` | String | 確認動作設定 |
| `buttonConfirmText` | String | 必要な確認テキスト（ハード確認を使用する場合） |
| `description` | String | ユーザー向けのヘルプテキスト |
| `required` | Boolean! | フィールドが必須かどうか |
| `isActive` | Boolean! | フィールドが現在アクティブかどうか |
| `projectId` | String! | このフィールドが属するプロジェクトのID |
| `createdAt` | DateTime! | フィールドが作成された日時 |
| `updatedAt` | DateTime! | フィールドが最後に変更された日時 |

## ボタンフィールドの動作

### 自動化統合

ボタンフィールドはBlueの自動化システムと連携するように設計されています：

1. **ボタンフィールドを作成** 上記のミューテーションを使用
2. **自動化を構成** `CUSTOM_FIELD_BUTTON_CLICKED` イベントをリッスン
3. **ユーザーがUIでボタンをクリック**
4. **自動化が構成されたアクションを実行**

### イベントフロー

ボタンがクリックされると：

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### データストレージなし

重要: ボタンフィールドは値データを保存しません。純粋にアクショントリガーとして機能します。各クリックは：
- イベントを生成
- 関連する自動化をトリガー
- タスク履歴にアクションを記録
- いかなるフィールド値も変更しません

## 必要な権限

ユーザーはボタンフィールドを作成および使用するために適切なプロジェクトロールが必要です：

| アクション | 必要なロール |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## エラーレスポンス

### 権限が拒否されました
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### カスタムフィールドが見つかりません
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

**注意**: APIは自動化が欠落している場合や確認の不一致に対して特定のエラーを返しません。

## ベストプラクティス

### 命名規則
- アクション指向の名前を使用: "請求書を送信", "レポートを作成", "チームに通知"
- ボタンが何をするのか具体的にする
- "ボタン1"や"ここをクリック"のような一般的な名前は避ける

### 確認設定
- 安全で可逆的なアクションのために `buttonType` を空のままにする
- UIクライアントに確認動作を提案するために `buttonType` を設定する
- UI確認でユーザーが入力すべき内容を指定するために `buttonConfirmText` を使用する
- 覚えておいてください: これらはUIのヒントに過ぎません - API呼び出しは常に即座に実行されます

### 自動化設計
- ボタンアクションは単一のワークフローに集中させる
- クリック後に何が起こったかについて明確なフィードバックを提供する
- ボタンの目的を説明するために説明テキストを追加することを検討する

## 一般的な使用例

1. **ワークフローの遷移**
   - "完了としてマーク"
   - "承認のために送信"
   - "タスクをアーカイブ"

2. **外部統合**
   - "CRMに同期"
   - "請求書を生成"
   - "メール更新を送信"

3. **バッチ操作**
   - "すべてのサブタスクを更新"
   - "プロジェクトにコピー"
   - "テンプレートを適用"

4. **報告アクション**
   - "レポートを生成"
   - "データをエクスポート"
   - "要約を作成"

## 制限事項

- ボタンはデータ値を保存または表示できません
- 各ボタンは自動化をトリガーすることしかできず、直接API呼び出しはできません（ただし、自動化には外部APIやBlueのAPIを呼び出すHTTPリクエストアクションを含めることができます）
- ボタンの可視性は条件に基づいて制御できません
- クリックごとに最大1つの自動化実行（ただし、その自動化は複数のアクションをトリガーできます）

## 関連リソース

- [自動化API](/api/automations/index) - ボタンによってトリガーされるアクションを構成する
- [カスタムフィールドの概要](/custom-fields/list-custom-fields) - 一般的なカスタムフィールドの概念