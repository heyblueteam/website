---
title: ユニークIDカスタムフィールド
description: 連続番号とカスタムフォーマットを持つ自動生成されたユニーク識別子フィールドを作成します
---

ユニークIDカスタムフィールドは、レコードのために連続的でユニークな識別子を自動的に生成します。チケット番号、注文ID、請求書番号、またはワークフロー内の任意の連続識別子システムを作成するのに最適です。

## 基本的な例

自動シーケンシングを使用してシンプルなユニークIDフィールドを作成します：

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## 高度な例

プレフィックスとゼロパディングを持つフォーマットされたユニークIDフィールドを作成します：

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## 入力パラメータ

### CreateCustomFieldInput (UNIQUE_ID)

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | ユニークIDフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `UNIQUE_ID` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |
| `useSequenceUniqueId` | Boolean | いいえ | 自動シーケンシングを有効にする（デフォルト：false） |
| `prefix` | String | いいえ | 生成されたIDのテキストプレフィックス（例："TASK-"） |
| `sequenceDigits` | Int | いいえ | ゼロパディングのための桁数 |
| `sequenceStartingNumber` | Int | いいえ | シーケンスの開始番号 |

## 設定オプション

### 自動シーケンシング (`useSequenceUniqueId`)
- **true**: レコードが作成されると自動的に連続IDが生成されます
- **false** または **undefined**: 手動入力が必要です（テキストフィールドのように機能します）

### プレフィックス (`prefix`)
- 生成されたすべてのIDに追加されるオプションのテキストプレフィックス
- 例："TASK-", "ORD-", "BUG-", "REQ-"
- 長さ制限はありませんが、表示に適したものにしてください

### シーケンス桁数 (`sequenceDigits`)
- シーケンス番号のゼロパディングのための桁数
- 例: `sequenceDigits: 3` は `001`、 `002`、 `003` を生成します
- 指定しない場合、パディングは適用されません

### 開始番号 (`sequenceStartingNumber`)
- シーケンスの最初の番号
- 例: `sequenceStartingNumber: 1000` は1000、1001、1002...から始まります
- 指定しない場合、1から始まります（デフォルトの動作）

## 生成されたIDフォーマット

最終的なIDフォーマットはすべての設定オプションを組み合わせます：

```
{prefix}{paddedSequenceNumber}
```

### フォーマット例

| 設定 | 生成されたID |
|---------------|---------------|
| オプションなし | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## ユニークID値の読み取り

### ユニークIDでレコードをクエリ
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### レスポンスフォーマット
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## 自動ID生成

### IDが生成されるタイミング
- **レコード作成**: 新しいレコードが作成されると自動的にIDが割り当てられます
- **フィールド追加**: 既存のレコードにUNIQUE_IDフィールドを追加する際、バックグラウンドジョブがキューに追加されます（ワーカーの実装は保留中）
- **バックグラウンド処理**: 新しいレコードのID生成はデータベーストリガーを介して同期的に行われます

### 生成プロセス
1. **トリガー**: 新しいレコードが作成されるか、UNIQUE_IDフィールドが追加されます
2. **シーケンスルックアップ**: システムは次に使用可能なシーケンス番号を見つけます
3. **ID割り当て**: シーケンス番号がレコードに割り当てられます
4. **カウンター更新**: 将来のレコードのためにシーケンスカウンターがインクリメントされます
5. **フォーマット**: IDは表示時にプレフィックスとパディングでフォーマットされます

### ユニーク性の保証
- **データベース制約**: 各フィールド内のシーケンスIDに対するユニーク制約
- **アトミック操作**: シーケンス生成はデータベースロックを使用して重複を防ぎます
- **プロジェクトスコープ**: シーケンスはプロジェクトごとに独立しています
- **レースコンディション保護**: 同時リクエストは安全に処理されます

## 手動モードと自動モード

### 自動モード (`useSequenceUniqueId: true`)
- IDはデータベーストリガーを介して自動的に生成されます
- 連続番号が保証されます
- アトミックなシーケンス生成により重複が防止されます
- フォーマットされたIDはプレフィックス + パディングされたシーケンス番号を組み合わせます

### 手動モード (`useSequenceUniqueId: false` または `undefined`)
- 通常のテキストフィールドのように機能します
- ユーザーは `setTodoCustomField` と `text` パラメータを使用してカスタム値を入力できます
- 自動生成はありません
- データベース制約を超えたユニーク性の強制はありません

## 手動値の設定（手動モードのみ）

`useSequenceUniqueId` がfalseの場合、値を手動で設定できます：

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## レスポンスフィールド

### TodoCustomFieldレスポンス (UNIQUE_ID)

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値のユニーク識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `sequenceId` | Int | 生成されたシーケンス番号（UNIQUE_IDフィールドに対して入力されます） |
| `text` | String | フォーマットされたテキスト値（プレフィックス + パディングされたシーケンスを組み合わせます） |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に更新された日時 |

### CustomFieldレスポンス (UNIQUE_ID)

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | 自動シーケンシングが有効かどうか |
| `prefix` | String | 生成されたIDのテキストプレフィックス |
| `sequenceDigits` | Int | ゼロパディングのための桁数 |
| `sequenceStartingNumber` | Int | シーケンスの開始番号 |

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## エラーレスポンス

### フィールド設定エラー
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### 権限エラー
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

## 重要な注意事項

### 自動生成されたID
- **読み取り専用**: 自動生成されたIDは手動で編集できません
- **永久**: 一度割り当てられると、シーケンスIDは変更されません
- **時系列**: IDは作成順序を反映します
- **スコープ**: シーケンスはプロジェクトごとに独立しています

### パフォーマンスの考慮事項
- 新しいレコードのID生成はデータベーストリガーを介して同期的に行われます
- シーケンス生成は `FOR UPDATE` ロックを使用してアトミック操作を行います
- バックグラウンドジョブシステムは存在しますが、ワーカーの実装は保留中です
- 高ボリュームプロジェクトのためにシーケンスの開始番号を考慮してください

### 移行と更新
- 既存のレコードに自動シーケンシングを追加すると、バックグラウンドジョブがキューに追加されます（ワーカー保留中）
- シーケンス設定の変更は将来のレコードにのみ影響します
- 設定の更新時に既存のIDは変更されません
- シーケンスカウンターは現在の最大値から続行されます

## ベストプラクティス

### 設定設計
- 他のシステムと衝突しない説明的なプレフィックスを選択してください
- 予想されるボリュームに適した桁パディングを使用してください
- 衝突を避けるために合理的な開始番号を設定してください
- デプロイ前にサンプルデータで設定をテストしてください

### プレフィックスガイドライン
- プレフィックスは短く、記憶に残るものにしてください（2-5文字）
- 一貫性のために大文字を使用してください
- 読みやすさのために区切り（ハイフン、アンダースコア）を含めてください
- URLやシステムで問題を引き起こす可能性のある特殊文字は避けてください

### シーケンス計画
- レコードのボリュームを見積もり、適切な桁パディングを選択してください
- 開始番号を設定する際に将来の成長を考慮してください
- 異なるレコードタイプのために異なるシーケンス範囲を計画してください
- チームの参照のためにIDスキームを文書化してください

## 一般的な使用例

1. **サポートシステム**
   - チケット番号: `TICK-001`, `TICK-002`
   - ケースID: `CASE-2024-001`
   - サポートリクエスト: `SUP-001`

2. **プロジェクト管理**
   - タスクID: `TASK-001`, `TASK-002`
   - スプリントアイテム: `SPRINT-001`
   - 納品物番号: `DEL-001`

3. **ビジネスオペレーション**
   - 注文番号: `ORD-2024-001`
   - 請求書ID: `INV-001`
   - 購入注文: `PO-001`

4. **品質管理**
   - バグレポート: `BUG-001`
   - テストケースID: `TEST-001`
   - レビュー番号: `REV-001`

## 統合機能

### 自動化との統合
- ユニークIDが割り当てられたときにアクションをトリガーします
- 自動化ルールでIDパターンを使用します
- メールテンプレートや通知でIDを参照します

### ルックアップとの統合
- 他のレコードからユニークIDを参照します
- ユニークIDでレコードを検索します
- 関連レコードの識別子を表示します

### レポーティングとの統合
- IDパターンでグループ化およびフィルタリングします
- ID割り当ての傾向を追跡します
- シーケンスの使用状況とギャップを監視します

## 制限事項

- **連続のみ**: IDは時系列で割り当てられます
- **ギャップなし**: 削除されたレコードはシーケンスにギャップを残します
- **再利用なし**: シーケンス番号は再利用されません
- **プロジェクトスコープ**: プロジェクト間でシーケンスを共有できません
- **フォーマット制約**: 限られたフォーマットオプション
- **バルク更新なし**: 既存のシーケンスIDを一括更新できません
- **カスタムロジックなし**: カスタムID生成ルールを実装できません

## 関連リソース

- [テキストフィールド](/api/custom-fields/text-single) - 手動テキスト識別子用
- [数値フィールド](/api/custom-fields/number) - 数値シーケンス用
- [カスタムフィールドの概要](/api/custom-fields/2.list-custom-fields) - 一般的な概念
- [自動化](/api/automations) - IDベースの自動化ルール用