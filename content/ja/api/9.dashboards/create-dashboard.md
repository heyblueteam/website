---
title: ダッシュボードの作成
description: Blueでのデータ可視化と報告のための新しいダッシュボードを作成します
---

## ダッシュボードの作成

`createDashboard` ミューテーションを使用すると、会社またはプロジェクト内に新しいダッシュボードを作成できます。ダッシュボードは、チームが指標を追跡し、進捗を監視し、データに基づいた意思決定を行うのに役立つ強力な可視化ツールです。

### 基本的な例

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### プロジェクト特有のダッシュボード

特定のプロジェクトに関連付けられたダッシュボードを作成します：

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## 入力パラメータ

### CreateDashboardInput

| パラメータ | タイプ | 必須 | 説明 |
|-----------|------|----------|-------------|
| `companyId` | String! | ✅ はい | ダッシュボードが作成される会社のID |
| `title` | String! | ✅ はい | ダッシュボードの名前。空でない文字列である必要があります |
| `projectId` | String | いいえ | このダッシュボードに関連付けるオプションのプロジェクトID |

## レスポンスフィールド

ミューテーションは完全な `Dashboard` オブジェクトを返します：

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | 作成されたダッシュボードの一意の識別子 |
| `title` | String! | 提供されたダッシュボードのタイトル |
| `companyId` | String! | このダッシュボードが属する会社 |
| `projectId` | String | 関連するプロジェクトID（提供された場合） |
| `project` | Project | 関連するプロジェクトオブジェクト（projectIdが提供された場合） |
| `createdBy` | User! | ダッシュボードを作成したユーザー（あなた） |
| `dashboardUsers` | [DashboardUser!]! | アクセス権を持つユーザーのリスト（最初は作成者のみ） |
| `createdAt` | DateTime! | ダッシュボードが作成された日時 |
| `updatedAt` | DateTime! | 最後の修正のタイムスタンプ（新しいダッシュボードの場合はcreatedAtと同じ） |

### DashboardUser フィールド

ダッシュボードが作成されると、作成者は自動的にダッシュボードユーザーとして追加されます：

| フィールド | タイプ | 説明 |
|-------|------|-------------|
| `id` | String! | ダッシュボードユーザー関係の一意の識別子 |
| `user` | User! | ダッシュボードにアクセスするユーザーオブジェクト |
| `role` | DashboardRole! | ユーザーの役割（作成者は完全なアクセス権を持つ） |
| `dashboard` | Dashboard! | ダッシュボードへの参照 |

## 必要な権限

指定された会社に属する認証されたユーザーは、ダッシュボードを作成できます。特別な役割の要件はありません。

| ユーザーステータス | ダッシュボードを作成できる |
|-------------|-------------------|
| Company Member | ✅ はい |
| 非会社メンバー | ❌ いいえ |
| Unauthenticated | ❌ いいえ |

## エラー応答

### 無効な会社
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### ユーザーが会社にいない
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### 無効なプロジェクト
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### 空のタイトル
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 重要な注意事項

- **自動所有権**: ダッシュボードを作成するユーザーは、自動的にその所有者となり、完全な権限を持ちます
- **プロジェクトの関連付け**: `projectId` を提供する場合、それは同じ会社に属する必要があります
- **初期権限**: 初めは作成者のみがアクセスできます。`editDashboard` を使用して他のユーザーを追加します
- **タイトル要件**: ダッシュボードのタイトルは空でない文字列でなければなりません。ユニーク性の要件はありません
- **会社のメンバーシップ**: ダッシュボードを作成するには、その会社のメンバーである必要があります

## ダッシュボード作成ワークフロー

1. **このミューテーションを使用してダッシュボードを作成します**
2. **ダッシュボードビルダーUIを使用してチャートやウィジェットを構成します**
3. **`editDashboard` ミューテーションを使用してチームメンバーを追加します（`dashboardUsers` とともに）**
4. **ダッシュボードインターフェースを通じてフィルターや日付範囲を設定します**
5. **ダッシュボードを共有または埋め込みます、そのユニークIDを使用して**

## ユースケース

1. **エグゼクティブダッシュボード**: 会社の指標の高レベルの概要を作成します
2. **プロジェクト追跡**: 進捗を監視するためのプロジェクト特有のダッシュボードを構築します
3. **チームパフォーマンス**: チームの生産性と達成指標を追跡します
4. **クライアント報告**: クライアント向けの報告のためのダッシュボードを作成します
5. **リアルタイム監視**: ライブ運用データのためのダッシュボードを設定します

## ベストプラクティス

1. **命名規則**: ダッシュボードの目的を示す明確で説明的なタイトルを使用します
2. **プロジェクトの関連付け**: プロジェクト特有のダッシュボードにはプロジェクトにリンクします
3. **アクセス管理**: 作成後すぐにチームメンバーを追加してコラボレーションを行います
4. **組織化**: 一貫した命名パターンを使用してダッシュボードの階層を作成します

## 関連操作

- [ダッシュボードの一覧](/api/dashboards/) - 会社またはプロジェクトのすべてのダッシュボードを取得します
- [ダッシュボードの編集](/api/dashboards/rename-dashboard) - ダッシュボードの名前を変更するか、ユーザーを管理します
- [ダッシュボードのコピー](/api/dashboards/copy-dashboard) - 既存のダッシュボードを複製します
- [ダッシュボードの削除](/api/dashboards/delete-dashboard) - ダッシュボードを削除します