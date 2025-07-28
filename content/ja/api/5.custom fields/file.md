---
title: ファイルカスタムフィールド
description: レコードにドキュメント、画像、その他のファイルを添付するためのファイルフィールドを作成します
---

ファイルカスタムフィールドを使用すると、レコードに複数のファイルを添付できます。ファイルはAWS S3に安全に保存され、包括的なメタデータ追跡、ファイルタイプの検証、および適切なアクセス制御が行われます。

## 基本的な例

シンプルなファイルフィールドを作成します：

```graphql
mutation CreateFileField {
  createCustomField(input: {
    name: "Attachments"
    type: FILE
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きのファイルフィールドを作成します：

```graphql
mutation CreateDetailedFileField {
  createCustomField(input: {
    name: "Project Documents"
    type: FILE
    description: "Upload project-related documents, images, and files"
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
| `name` | String! | ✅ はい | ファイルフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `FILE` である必要があります |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

**注意**: カスタムフィールドは、ユーザーの現在のプロジェクトコンテキストに基づいてプロジェクトに自動的に関連付けられます。`projectId` パラメータは必要ありません。

## ファイルアップロードプロセス

### ステップ1: ファイルをアップロード

最初に、ファイルをアップロードしてファイルUIDを取得します：

```graphql
mutation UploadFile {
  uploadFile(input: {
    file: $file  # File upload variable
    companyId: "company_123"
    projectId: "proj_123"
  }) {
    id
    uid
    name
    size
    type
    extension
    status
  }
}
```

### ステップ2: ファイルをレコードに添付

次に、アップロードしたファイルをレコードに添付します：

```graphql
mutation AttachFileToRecord {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "file_field_456"
    fileUid: "file_uid_from_upload"
  }) {
    id
    file {
      uid
      name
      size
      type
    }
  }
}
```

## ファイル添付の管理

### 単一ファイルの追加

```graphql
mutation AddFileToField {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  }) {
    id
    position
    file {
      uid
      name
      size
      type
      extension
    }
  }
}
```

### ファイルの削除

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### バルクファイル操作

customFieldOptionIdsを使用して複数のファイルを一度に更新します：

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## ファイルアップロード入力パラメータ

### UploadFileInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `file` | Upload! | ✅ はい | アップロードするファイル |
| `companyId` | String! | ✅ はい | ファイルストレージの会社ID |
| `projectId` | String | いいえ | プロジェクト固有のファイルのプロジェクトID |

### ファイル管理入力パラメータ

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | レコードのID |
| `customFieldId` | String! | ✅ はい | ファイルカスタムフィールドのID |
| `fileUid` | String! | ✅ はい | アップロードされたファイルの一意の識別子 |

## ファイルストレージと制限

### ファイルサイズ制限

| 制限タイプ | サイズ |
|------------|------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### サポートされているファイルタイプ

#### 画像
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### 動画
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### 音声
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### ドキュメント
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### アーカイブ
- `zip`, `rar`, `7z`, `tar`, `gz`

#### コード/テキスト
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### ストレージアーキテクチャ

- **ストレージ**: AWS S3、整理されたフォルダ構造
- **パス形式**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **セキュリティ**: 安全なアクセスのための署名付きURL
- **バックアップ**: 自動S3冗長性

## レスポンスフィールド

### ファイルレスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | ID! | データベースID |
| `uid` | String! | 一意のファイル識別子 |
| `name` | String! | 元のファイル名 |
| `size` | Float! | バイト単位のファイルサイズ |
| `type` | String! | MIMEタイプ |
| `extension` | String! | ファイル拡張子 |
| `status` | FileStatus | 保留中または確認済み（nullable） |
| `shared` | Boolean! | ファイルが共有されているかどうか |
| `createdAt` | DateTime! | アップロードのタイムスタンプ |

### TodoCustomFieldFileレスポンス

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | ID! | ジャンクションレコードID |
| `uid` | String! | 一意の識別子 |
| `position` | Float! | 表示順序 |
| `file` | File! | 関連付けられたファイルオブジェクト |
| `todoCustomField` | TodoCustomField! | 親カスタムフィールド |
| `createdAt` | DateTime! | ファイルが添付された日時 |

## ファイル付きレコードの作成

レコードを作成する際に、ファイルのUIDを使用して添付できます：

```graphql
mutation CreateRecordWithFiles {
  createTodo(input: {
    title: "Project deliverables"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "file_field_id"
      customFieldOptionIds: ["file_uid_1", "file_uid_2"]
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
      todoCustomFieldFiles {
        id
        position
        file {
          uid
          name
          size
          type
        }
      }
    }
  }
}
```

## ファイルの検証とセキュリティ

### アップロード検証

- **MIMEタイプチェック**: 許可されたタイプに対して検証します
- **ファイル拡張子検証**: `application/octet-stream` のフォールバック
- **サイズ制限**: アップロード時に強制されます
- **ファイル名のサニタイズ**: 特殊文字を削除します

### アクセス制御

- **アップロード権限**: プロジェクト/会社のメンバーシップが必要です
- **ファイルの関連付け**: ADMIN、OWNER、MEMBER、CLIENTの役割
- **ファイルアクセス**: プロジェクト/会社の権限から継承されます
- **安全なURL**: ファイルアクセスのための時間制限付き署名URL

## 必要な権限

| アクション | 必要な権限 |
|--------|-------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## エラーレスポンス

### ファイルが大きすぎます
```json
{
  "errors": [{
    "message": "File \"filename.pdf\": Size exceeds maximum limit of 256MB",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### ファイルが見つかりません
```json
{
  "errors": [{
    "message": "File not found",
    "extensions": {
      "code": "FILE_NOT_FOUND"
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

### ファイル管理
- レコードに添付する前にファイルをアップロードします
- 説明的なファイル名を使用します
- プロジェクト/目的ごとにファイルを整理します
- 定期的に未使用のファイルをクリーンアップします

### パフォーマンス
- 可能な場合はファイルをバッチでアップロードします
- コンテンツタイプに適したファイル形式を使用します
- アップロード前に大きなファイルを圧縮します
- ファイルのプレビュー要件を考慮します

### セキュリティ
- 拡張子だけでなくファイルの内容を検証します
- アップロードされたファイルにウイルススキャンを実施します
- 適切なアクセス制御を実装します
- ファイルアップロードのパターンを監視します

## 一般的なユースケース

1. **ドキュメント管理**
   - プロジェクト仕様
   - 契約書および合意書
   - 会議のメモおよびプレゼンテーション
   - 技術文書

2. **資産管理**
   - デザインファイルおよびモックアップ
   - ブランド資産およびロゴ
   - マーケティング資料
   - 製品画像

3. **コンプライアンスおよび記録**
   - 法的文書
   - 監査証跡
   - 証明書およびライセンス
   - 財務記録

4. **コラボレーション**
   - 共有リソース
   - バージョン管理された文書
   - フィードバックおよび注釈
   - 参考資料

## 統合機能

### 自動化との統合
- ファイルが追加/削除されたときにアクションをトリガーします
- タイプまたはメタデータに基づいてファイルを処理します
- ファイルの変更について通知を送信します
- 条件に基づいてファイルをアーカイブします

### カバー画像との統合
- ファイルフィールドをカバー画像のソースとして使用します
- 自動画像処理およびサムネイル
- ファイルが変更されたときの動的カバー更新

### ルックアップとの統合
- 他のレコードからファイルを参照します
- ファイルのカウントとサイズを集計します
- ファイルメタデータによってレコードを検索します
- ファイル添付を相互参照します

## 制限事項

- ファイルごとの最大256MB
- S3の可用性に依存
- 組み込みのファイルバージョン管理なし
- 自動ファイル変換なし
- 限定的なファイルプレビュー機能
- リアルタイムの共同編集なし

## 関連リソース

- [ファイルアップロードAPI](/api/upload-files) - ファイルアップロードエンドポイント
- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的な概念
- [自動化API](/api/automations) - ファイルベースの自動化
- [AWS S3ドキュメント](https://docs.aws.amazon.com/s3/) - ストレージバックエンド