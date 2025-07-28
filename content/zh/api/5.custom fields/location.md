---
title: 位置自定义字段
description: 创建位置字段以存储记录的地理坐标
---

位置自定义字段存储记录的地理坐标（纬度和经度）。它们支持精确的坐标存储、地理空间查询和高效的基于位置的过滤。

## 基本示例

创建一个简单的位置字段：

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

## 高级示例

创建一个带描述的位置字段：

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

## 输入参数

### CreateCustomFieldInput

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 位置字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `LOCATION` |
| `description` | String | 否 | 显示给用户的帮助文本 |

## 设置位置值

位置字段存储纬度和经度坐标：

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

### SetTodoCustomFieldInput 参数

| 参数 | 类型 | 必需 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 要更新的记录的 ID |
| `customFieldId` | String! | ✅ 是 | 位置自定义字段的 ID |
| `latitude` | Float | 否 | 纬度坐标（-90 到 90） |
| `longitude` | Float | 否 | 经度坐标（-180 到 180） |

**注意**：虽然这两个参数在模式中都是可选的，但有效位置需要两个坐标。如果只提供一个，位置将无效。

## 坐标验证

### 有效范围

| 坐标 | 范围 | 描述 |
|------|------|------|
| Latitude | -90 to 90 | 南北位置 |
| Longitude | -180 to 180 | 东西位置 |

### 示例坐标

| 位置 | 纬度 | 经度 |
|------|------|------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## 使用位置值创建记录

创建带有位置数据的新记录时：

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

### 创建的输入格式

创建记录时，位置值使用逗号分隔格式：

| 格式 | 示例 | 描述 |
|------|------|------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | 标准坐标格式 |
| `"51.5074,-0.1278"` | London coordinates | 逗号周围不允许有空格 |
| `"-33.8688,151.2093"` | Sydney coordinates | 允许负值 |

## 响应字段

### TodoCustomField 响应

| 字段 | 类型 | 描述 |
|------|------|------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `latitude` | Float | 纬度坐标 |
| `longitude` | Float | 经度坐标 |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后修改的时间 |

## 重要限制

### 无内置地理编码

位置字段仅存储坐标 - 它们不包括：
- 地址到坐标的转换
- 反向地理编码（坐标到地址）
- 地址验证或搜索
- 与地图服务的集成
- 地名查找

### 需要外部服务

要实现地址功能，您需要集成外部服务：
- **Google Maps API** 用于地理编码
- **OpenStreetMap Nominatim** 用于免费地理编码
- **MapBox** 用于地图和地理编码
- **Here API** 用于位置服务

### 示例集成

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

## 所需权限

| 操作 | 所需角色 |
|------|----------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## 错误响应

### 无效坐标
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

### 无效经度
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

## 最佳实践

### 数据收集
- 使用 GPS 坐标以获取精确位置
- 在存储之前验证坐标
- 考虑坐标精度需求（6 位小数 ≈ 10 厘米精度）
- 以十进制度存储坐标（而不是度/分钟/秒）

### 用户体验
- 提供地图界面以选择坐标
- 显示坐标时提供位置预览
- 在 API 调用之前进行客户端坐标验证
- 考虑位置数据的时区影响

### 性能
- 使用空间索引以提高查询效率
- 限制坐标精度到所需的准确性
- 考虑对频繁访问的位置进行缓存
- 尽可能批量更新位置

## 常见用例

1. **现场操作**
   - 设备位置
   - 服务呼叫地址
   - 检查地点
   - 交付地点

2. **事件管理**
   - 事件场地
   - 会议地点
   - 会议场所
   - 研讨会地点

3. **资产跟踪**
   - 设备位置
   - 设施位置
   - 车辆跟踪
   - 库存位置

4. **地理分析**
   - 服务覆盖区域
   - 客户分布
   - 市场分析
   - 领土管理

## 集成功能

### 与查找
- 从其他记录引用位置数据
- 按地理接近性查找记录
- 聚合基于位置的数据
- 交叉引用坐标

### 与自动化
- 根据位置变化触发操作
- 创建地理围栏通知
- 当位置变化时更新相关记录
- 生成基于位置的报告

### 与公式
- 计算位置之间的距离
- 确定地理中心
- 分析位置模式
- 创建基于位置的指标

## 限制

- 无内置地理编码或地址转换
- 不提供地图界面
- 地址功能需要外部服务
- 仅限于坐标存储
- 除范围检查外没有自动位置验证

## 相关资源

- [自定义字段概述](/api/custom-fields/list-custom-fields) - 一般概念
- [Google Maps API](https://developers.google.com/maps) - 地理编码服务
- [OpenStreetMap Nominatim](https://nominatim.org/) - 免费地理编码
- [MapBox API](https://docs.mapbox.com/) - 地图和地理编码服务