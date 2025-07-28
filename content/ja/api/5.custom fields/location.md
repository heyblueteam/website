---
title: ロケーションカスタムフィールド
description: レコードの地理座標を保存するためのロケーションフィールドを作成します
---

ロケーションカスタムフィールドは、レコードの地理座標（緯度と経度）を保存します。これにより、正確な座標の保存、地理空間クエリ、および効率的な位置ベースのフィルタリングがサポートされます。

## 基本的な例

シンプルなロケーションフィールドを作成します：

```graphql
mutation CreateLocationField {
  createCustomField(input: {
    name: "Meeting Location"
    type: LOCATION
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 高度な例

説明付きのロケーションフィールドを作成します：

```graphql
mutation CreateDetailedLocationField {
  createCustomField(input: {
    name: "Office Location"
    type: LOCATION
    projectId: "proj_123"
    description: "Primary office location coordinates"
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

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ はい | ロケーションフィールドの表示名 |
| `type` | CustomFieldType! | ✅ はい | `LOCATION` でなければなりません |
| `description` | String | いいえ | ユーザーに表示されるヘルプテキスト |

## ロケーション値の設定

ロケーションフィールドは緯度と経度の座標を保存します：

```graphql
mutation SetLocationValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    latitude: 40.7128
    longitude: -74.0060
  })
}
```

### SetTodoCustomFieldInput パラメータ

| パラメータ | 型 | 必須 | 説明 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ はい | 更新するレコードのID |
| `customFieldId` | String! | ✅ はい | ロケーションカスタムフィールドのID |
| `latitude` | Float | いいえ | 緯度座標（-90から90） |
| `longitude` | Float | いいえ | 経度座標（-180から180） |

**注意**: スキーマでは両方のパラメータがオプションですが、有効なロケーションには両方の座標が必要です。どちらか一方のみが提供された場合、ロケーションは無効になります。

## 座標の検証

### 有効な範囲

| 座標 | 範囲 | 説明 |
|------------|-------|-------------|
| Latitude | -90 to 90 | 北/南の位置 |
| Longitude | -180 to 180 | 東/西の位置 |

### 例の座標

| ロケーション | 緯度 | 経度 |
|----------|----------|-----------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## ロケーション値を持つレコードの作成

ロケーションデータを持つ新しいレコードを作成する場合：

```graphql
mutation CreateRecordWithLocation {
  createTodo(input: {
    title: "Site Visit"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "location_field_id"
      value: "40.7128,-74.0060"  # Format: "latitude,longitude"
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
      latitude
      longitude
    }
  }
}
```

### 作成のための入力形式

レコードを作成する際、ロケーション値はカンマ区切り形式を使用します：

| 形式 | 例 | 説明 |
|--------|---------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | 標準座標形式 |
| `"51.5074,-0.1278"` | London coordinates | カンマの周りにスペースなし |
| `"-33.8688,151.2093"` | Sydney coordinates | 負の値が許可されます |

## レスポンスフィールド

### TodoCustomField レスポンス

| フィールド | 型 | 説明 |
|-------|------|-------------|
| `id` | String! | フィールド値のユニークな識別子 |
| `customField` | CustomField! | カスタムフィールドの定義 |
| `latitude` | Float | 緯度座標 |
| `longitude` | Float | 経度座標 |
| `todo` | Todo! | この値が属するレコード |
| `createdAt` | DateTime! | 値が作成された日時 |
| `updatedAt` | DateTime! | 値が最後に修正された日時 |

## 重要な制限事項

### 組み込みのジオコーディングなし

ロケーションフィールドは座標のみを保存します - それらは **含まれません**：
- 住所から座標への変換
- 逆ジオコーディング（座標から住所）
- 住所の検証または検索
- マッピングサービスとの統合
- 地名のルックアップ

### 外部サービスが必要

住所機能には、外部サービスとの統合が必要です：
- **Google Maps API** ジオコーディング用
- **OpenStreetMap Nominatim** 無料ジオコーディング用
- **MapBox** マッピングおよびジオコーディング用
- **Here API** ロケーションサービス用

### 統合の例

```javascript
// Client-side geocoding example (not part of Blue API)
async function geocodeAddress(address) {
  const response = await fetch(
    `https://maps.googleapis.com/maps/api/geocode/json?address=${encodeURIComponent(address)}&key=${API_KEY}`
  );
  const data = await response.json();
  
  if (data.results.length > 0) {
    const { lat, lng } = data.results[0].geometry.location;
    
    // Now set the location field in Blue
    await setTodoCustomField({
      todoId: "todo_123",
      customFieldId: "location_field_456",
      latitude: lat,
      longitude: lng
    });
  }
}
```

## 必要な権限

| アクション | 必要な役割 |
|--------|---------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## エラーレスポンス

### 無効な座標
```json
{
  "errors": [{
    "message": "Invalid coordinates: latitude must be between -90 and 90",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 無効な経度
```json
{
  "errors": [{
    "message": "Invalid coordinates: longitude must be between -180 and 180",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## ベストプラクティス

### データ収集
- 正確な位置のためにGPS座標を使用
- 保存前に座標を検証
- 座標の精度ニーズを考慮（6桁の小数点 ≈ 10cmの精度）
- 座標を小数度で保存（度/分/秒ではなく）

### ユーザーエクスペリエンス
- 座標選択のためのマップインターフェースを提供
- 座標を表示する際にロケーションプレビューを表示
- API呼び出しの前にクライアント側で座標を検証
- ロケーションデータのタイムゾーンの影響を考慮

### パフォーマンス
- 効率的なクエリのために空間インデックスを使用
- 必要な精度に座標の精度を制限
- 頻繁にアクセスされるロケーションのキャッシュを考慮
- 可能な場合はロケーションのバッチ更新を行う

## 一般的なユースケース

1. **フィールドオペレーション**
   - 設備の位置
   - サービスコールの住所
   - 検査サイト
   - 配送場所

2. **イベント管理**
   - イベント会場
   - 会議の場所
   - 会議サイト
   - ワークショップの場所

3. **資産追跡**
   - 設備の位置
   - 施設の場所
   - 車両追跡
   - 在庫の場所

4. **地理的分析**
   - サービスカバレッジエリア
   - 顧客分布
   - 市場分析
   - テリトリー管理

## 統合機能

### ルックアップとの統合
- 他のレコードからロケーションデータを参照
- 地理的近接によるレコードの検索
- 位置ベースのデータの集約
- 座標のクロスリファレンス

### 自動化との統合
- ロケーションの変更に基づいてアクションをトリガー
- ジオフェンス通知を作成
- ロケーションが変更されたときに関連レコードを更新
- 位置ベースのレポートを生成

### フォーミュラとの統合
- ロケーション間の距離を計算
- 地理的中心を決定
- ロケーションパターンを分析
- 位置ベースのメトリックを作成

## 制限事項

- 組み込みのジオコーディングや住所変換なし
- マッピングインターフェースは提供されていない
- 住所機能には外部サービスが必要
- 座標の保存のみが制限されている
- 範囲チェックを超えた自動ロケーション検証なし

## 関連リソース

- [カスタムフィールドの概要](/api/custom-fields/list-custom-fields) - 一般的な概念
- [Google Maps API](https://developers.google.com/maps) - ジオコーディングサービス
- [OpenStreetMap Nominatim](https://nominatim.org/) - 無料ジオコーディング
- [MapBox API](https://docs.mapbox.com/) - マッピングおよびジオコーディングサービス