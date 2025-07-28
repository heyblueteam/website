---
title: 위치 사용자 정의 필드
description: 레코드의 지리적 좌표를 저장하기 위한 위치 필드 생성
---

위치 사용자 정의 필드는 레코드의 지리적 좌표(위도 및 경도)를 저장합니다. 이들은 정밀한 좌표 저장, 지리 공간 쿼리 및 효율적인 위치 기반 필터링을 지원합니다.

## 기본 예제

간단한 위치 필드를 생성합니다:

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

## 고급 예제

설명이 포함된 위치 필드를 생성합니다:

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

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 위치 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `LOCATION` 여야 합니다. |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

## 위치 값 설정

위치 필드는 위도 및 경도 좌표를 저장합니다:

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

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 위치 사용자 정의 필드의 ID |
| `latitude` | Float | 아니오 | 위도 좌표 (-90에서 90) |
| `longitude` | Float | 아니오 | 경도 좌표 (-180에서 180) |

**참고**: 두 매개변수 모두 스키마에서 선택 사항이지만, 유효한 위치를 위해서는 두 좌표가 모두 필요합니다. 하나만 제공되면 위치가 유효하지 않습니다.

## 좌표 검증

### 유효 범위

| 좌표 | 범위 | 설명 |
|------------|-------|-------------|
| Latitude | -90 to 90 | 북/남 위치 |
| Longitude | -180 to 180 | 동/서 위치 |

### 예제 좌표

| 위치 | 위도 | 경도 |
|----------|----------|-----------|
| New York City | 40.7128 | -74.0060 |
| London | 51.5074 | -0.1278 |
| Sydney | -33.8688 | 151.2093 |
| Tokyo | 35.6762 | 139.6503 |
| São Paulo | -23.5505 | -46.6333 |

## 위치 값으로 레코드 생성

위치 데이터로 새 레코드를 생성할 때:

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

### 생성 입력 형식

레코드를 생성할 때 위치 값은 쉼표로 구분된 형식을 사용합니다:

| 형식 | 예제 | 설명 |
|--------|---------|-------------|
| `"latitude,longitude"` | `"40.7128,-74.0060"` | 표준 좌표 형식 |
| `"51.5074,-0.1278"` | London coordinates | 쉼표 주위에 공백 없음 |
| `"-33.8688,151.2093"` | Sydney coordinates | 음수 값 허용 |

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `latitude` | Float | 위도 좌표 |
| `longitude` | Float | 경도 좌표 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

## 중요한 제한 사항

### 내장 지오코딩 없음

위치 필드는 좌표만 저장합니다 - 다음은 포함되지 않습니다:
- 주소-좌표 변환
- 역 지오코딩 (좌표-주소)
- 주소 검증 또는 검색
- 매핑 서비스와의 통합
- 장소 이름 조회

### 외부 서비스 필요

주소 기능을 위해서는 외부 서비스를 통합해야 합니다:
- **Google Maps API** 지오코딩용
- **OpenStreetMap Nominatim** 무료 지오코딩용
- **MapBox** 매핑 및 지오코딩용
- **Here API** 위치 서비스용

### 예제 통합

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

## 필요한 권한

| 작업 | 필요한 역할 |
|--------|---------------|
| Create location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Update location field | `OWNER`, `ADMIN`, or `MEMBER` in the project |
| Set location value | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` with edit permissions on the record |
| View location value | Any project member with read access to the record |

## 오류 응답

### 잘못된 좌표
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

### 잘못된 경도
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

## 모범 사례

### 데이터 수집
- 정확한 위치를 위해 GPS 좌표 사용
- 저장하기 전에 좌표 검증
- 좌표 정밀도 요구 사항 고려 (소수점 6자리 ≈ 10cm 정확도)
- 좌표를 십진수로 저장 (도/분/초가 아님)

### 사용자 경험
- 좌표 선택을 위한 지도 인터페이스 제공
- 좌표를 표시할 때 위치 미리보기 표시
- API 호출 전에 클라이언트 측에서 좌표 검증
- 위치 데이터에 대한 시간대 영향 고려

### 성능
- 효율적인 쿼리를 위해 공간 인덱스 사용
- 필요한 정확도로 좌표 정밀도 제한
- 자주 접근하는 위치에 대한 캐싱 고려
- 가능할 경우 위치 업데이트를 일괄 처리

## 일반 사용 사례

1. **현장 작업**
   - 장비 위치
   - 서비스 호출 주소
   - 검사 장소
   - 배송 위치

2. **이벤트 관리**
   - 이벤트 장소
   - 회의 위치
   - 컨퍼런스 장소
   - 워크숍 위치

3. **자산 추적**
   - 장비 위치
   - 시설 위치
   - 차량 추적
   - 재고 위치

4. **지리적 분석**
   - 서비스 범위 지역
   - 고객 분포
   - 시장 분석
   - 지역 관리

## 통합 기능

### 조회와 함께
- 다른 레코드의 위치 데이터 참조
- 지리적 근접성으로 레코드 찾기
- 위치 기반 데이터 집계
- 좌표 교차 참조

### 자동화와 함께
- 위치 변경에 따라 작업 트리거
- 지오펜스 알림 생성
- 위치가 변경될 때 관련 레코드 업데이트
- 위치 기반 보고서 생성

### 수식과 함께
- 위치 간 거리 계산
- 지리적 중심 결정
- 위치 패턴 분석
- 위치 기반 메트릭 생성

## 제한 사항

- 내장 지오코딩 또는 주소 변환 없음
- 제공된 매핑 인터페이스 없음
- 주소 기능을 위해 외부 서비스 필요
- 좌표 저장만 가능
- 범위 검사를 넘어서는 자동 위치 검증 없음

## 관련 리소스

- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 개념
- [Google Maps API](https://developers.google.com/maps) - 지오코딩 서비스
- [OpenStreetMap Nominatim](https://nominatim.org/) - 무료 지오코딩
- [MapBox API](https://docs.mapbox.com/) - 매핑 및 지오코딩 서비스