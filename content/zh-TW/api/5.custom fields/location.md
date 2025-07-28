---
title: 位置自訂欄位
description: 創建位置欄位以儲存記錄的地理座標
---

位置自訂欄位儲存記錄的地理座標（緯度和經度）。它們支持精確的座標儲存、地理空間查詢和高效的基於位置的過濾。

## 基本範例

創建一個簡單的位置信息欄位：

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

## 進階範例

創建一個帶有描述的位置信息欄位：

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

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 位置欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `LOCATION` |
| `description` | String | 否 | 顯示給用戶的幫助文本 |

## 設定位置值

位置欄位儲存緯度和經度座標：

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

### SetTodoCustomFieldInput 參數

| 參數 | 類型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的記錄的 ID |
| `customFieldId` | String! | ✅ 是 | 位置自訂欄位的 ID |
| `latitude` | Float | 否 | 緯度座標（-90 到 90） |
| `longitude` | Float | 否 | 經度座標（-180 到 180） |

**注意**：雖然這兩個參數在架構中都是可選的，但有效的位置需要兩個座標。如果只提供一個，則位置將無效。

## 座標驗證

### 有效範圍

| 座標 | 範圍 | 描述 |
|------|------|------|
| Latitude | -90 to 90 | 南北位置 |
| Longitude | -180 to 180 | 東西位置 |

### 範例座標

| 位置 | 緯度 | 經度 |
|------|------|------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## 使用位置值創建記錄

當創建一個帶有位置數據的新記錄時：

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

### 創建的輸入格式

在創建記錄時，位置值使用逗號分隔格式：

| 格式 | 範例 | 描述 |
|------|------|------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | 標準座標格式 |
| `"51.5074,-0.1278"` | London coordinates | 逗號周圍無空格 |
| `"-33.8688,151.2093"` | Sydney coordinates | 允許負值 |

## 回應欄位

### TodoCustomField 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | String! | 欄位值的唯一標識符 |
| `customField` | CustomField! | 自訂欄位定義 |
| `latitude` | Float | 緯度座標 |
| `longitude` | Float | 經度座標 |
| `todo` | Todo! | 此值所屬的記錄 |
| `createdAt` | DateTime! | 值創建的時間 |
| `updatedAt` | DateTime! | 值最後修改的時間 |

## 重要限制

### 無內建地理編碼

位置欄位僅儲存座標 - 它們不包括：
- 地址到座標的轉換
- 反向地理編碼（座標到地址）
- 地址驗證或搜索
- 與地圖服務的整合
- 地名查詢

### 需要外部服務

對於地址功能，您需要整合外部服務：
- **Google Maps API** 用於地理編碼
- **OpenStreetMap Nominatim** 用於免費地理編碼
- **MapBox** 用於地圖和地理編碼
- **Here API** 用於位置服務

### 整合範例

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

## 所需權限

| 行動 | 所需角色 |
|------|----------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## 錯誤回應

### 無效座標
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

### 無效經度
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

## 最佳實踐

### 數據收集
- 使用 GPS 座標以獲得精確位置
- 在儲存之前驗證座標
- 考慮座標精度需求（6 位小數 ≈ 10 公分精度）
- 以十進制度數儲存座標（而不是度/分鐘/秒）

### 用戶體驗
- 提供地圖介面以選擇座標
- 顯示座標時顯示位置預覽
- 在 API 調用之前進行客戶端座標驗證
- 考慮位置數據的時區影響

### 性能
- 使用空間索引以提高查詢效率
- 限制座標精度至所需的準確度
- 考慮對經常訪問的位置進行緩存
- 在可能的情況下批量更新位置

## 常見用例

1. **現場操作**
   - 設備位置
   - 服務呼叫地址
   - 檢查地點
   - 交付位置

2. **事件管理**
   - 活動場地
   - 會議地點
   - 會議場所
   - 工作坊地點

3. **資產追蹤**
   - 設備位置
   - 設施位置
   - 車輛追蹤
   - 庫存位置

4. **地理分析**
   - 服務覆蓋區域
   - 客戶分佈
   - 市場分析
   - 領土管理

## 整合功能

### 與查詢
- 參考其他記錄的位置信息
- 按地理接近度查找記錄
- 聚合基於位置的數據
- 交叉參考座標

### 與自動化
- 根據位置變更觸發行動
- 創建地理圍欄通知
- 當位置變更時更新相關記錄
- 生成基於位置的報告

### 與公式
- 計算位置之間的距離
- 確定地理中心
- 分析位置模式
- 創建基於位置的指標

## 限制

- 無內建地理編碼或地址轉換
- 不提供地圖介面
- 需要外部服務以實現地址功能
- 僅限於座標儲存
- 除範圍檢查外，無自動位置驗證

## 相關資源

- [自訂欄位概述](/api/custom-fields/list-custom-fields) - 一般概念
- [Google Maps API](https://developers.google.com/maps) - 地理編碼服務
- [OpenStreetMap Nominatim](https://nominatim.org/) - 免費地理編碼
- [MapBox API](https://docs.mapbox.com/) - 地圖和地理編碼服務