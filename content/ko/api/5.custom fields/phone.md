---
title: 전화 사용자 정의 필드
description: 국제 형식으로 전화번호를 저장하고 검증하기 위한 전화 필드 생성
---

전화 사용자 정의 필드는 레코드에 전화번호를 저장하고 내장된 검증 및 국제 형식을 제공합니다. 이는 연락처 정보, 비상 연락처 또는 프로젝트의 전화 관련 데이터를 추적하는 데 이상적입니다.

## 기본 예제

간단한 전화 필드를 생성합니다:

```graphql
mutation CreatePhoneField {
  createCustomField(input: {
    name: "Contact Phone"
    type: PHONE
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 전화 필드를 생성합니다:

```graphql
mutation CreateDetailedPhoneField {
  createCustomField(input: {
    name: "Emergency Contact"
    type: PHONE
    description: "Emergency contact number with country code"
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
| `name` | String! | ✅ 예 | 전화 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `PHONE` 여야 합니다. |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 사용자 정의 필드는 사용자의 현재 프로젝트 컨텍스트에 따라 프로젝트와 자동으로 연결됩니다. `projectId` 매개변수는 필요하지 않습니다.

## 전화 값 설정

레코드에서 전화 값을 설정하거나 업데이트하려면:

```graphql
mutation SetPhoneValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "+1 234 567 8900"
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 전화 사용자 정의 필드의 ID |
| `text` | String | 아니오 | 국가 코드가 포함된 전화번호 |
| `regionCode` | String | 아니오 | 국가 코드 (자동으로 감지됨) |

**참고**: `text`는 스키마에서 선택 사항이지만, 필드가 의미를 가지려면 전화번호가 필요합니다. `setTodoCustomField`를 사용할 때는 검증이 수행되지 않으며, 아무 텍스트 값과 regionCode를 저장할 수 있습니다. 자동 감지는 레코드 생성 중에만 발생합니다.

## 전화 값으로 레코드 생성

전화 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithPhone {
  createTodo(input: {
    title: "Call client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "phone_field_id"
      value: "+1-555-123-4567"
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
      text
      regionCode
    }
  }
}
```

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `text` | String | 형식이 지정된 전화번호 (국제 형식) |
| `regionCode` | String | 국가 코드 (예: "US", "GB", "CA") |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

## 전화번호 검증

**중요**: 전화번호 검증 및 형식 지정은 `createTodo`를 통해 새 레코드를 생성할 때만 발생합니다. `setTodoCustomField`를 사용하여 기존 전화 값을 업데이트할 때는 검증이 수행되지 않으며, 값은 제공된 대로 저장됩니다.

### 허용되는 형식 (레코드 생성 중)
전화번호는 다음 형식 중 하나로 국가 코드를 포함해야 합니다:

- **E.164 형식 (선호)**: `+12345678900`
- **국제 형식**: `+1 234 567 8900`
- **구두점이 있는 국제 형식**: `+1 (234) 567-8900`
- **대시가 있는 국가 코드**: `+1-234-567-8900`

**참고**: 국가 코드가 없는 국가 형식 (예: `(234) 567-8900`)은 레코드 생성 중에 거부됩니다.

### 검증 규칙 (레코드 생성 중)
- libphonenumber-js를 사용하여 구문 분석 및 검증
- 다양한 국제 전화번호 형식 수용
- 번호에서 국가 자동 감지
- 국제 표시 형식으로 번호 형식 지정 (예: `+1 234 567 8900`)
- 국가 코드를 별도로 추출하여 저장 (예: `US`)

### 유효한 전화 예제
```
+12345678900           # E.164 format
+1 234 567 8900        # International format
+1 (234) 567-8900      # With parentheses
+1-234-567-8900        # With dashes
+44 20 7946 0958       # UK number
+33 1 42 86 83 26      # French number
```

### 유효하지 않은 전화 예제
```
(234) 567-8900         # Missing country code
234-567-8900           # Missing country code
123                    # Too short
invalid-phone          # Not a number
+1 234                 # Incomplete number
```

## 저장 형식

전화번호로 레코드를 생성할 때:
- **text**: 검증 후 국제 형식으로 저장 (예: `+1 234 567 8900`)
- **regionCode**: ISO 국가 코드로 저장 (예: `US`, `GB`, `CA`) 자동 감지됨

`setTodoCustomField`를 통해 업데이트할 때:
- **text**: 제공된 대로 정확히 저장 (형식 없음)
- **regionCode**: 제공된 대로 정확히 저장 (검증 없음)

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create phone field | `OWNER` or `ADMIN` role at project level |
| Update phone field | `OWNER` or `ADMIN` role at project level |
| Set phone value | Standard record edit permissions |
| View phone value | Standard record view permissions |

## 오류 응답

### 잘못된 전화 형식
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 필드 찾을 수 없음
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

### 국가 코드 누락
```json
{
  "errors": [{
    "message": "Invalid phone number format.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 모범 사례

### 데이터 입력
- 전화번호에 항상 국가 코드를 포함하세요.
- 일관성을 위해 E.164 형식을 사용하세요.
- 중요한 작업을 위해 저장하기 전에 번호를 검증하세요.
- 표시 형식에 대한 지역 선호도를 고려하세요.

### 데이터 품질
- 글로벌 호환성을 위해 국제 형식으로 번호를 저장하세요.
- 국가별 기능을 위해 regionCode를 사용하세요.
- 중요한 작업 (SMS, 통화) 전에 전화번호를 검증하세요.
- 연락처 타이밍에 대한 시간대 영향을 고려하세요.

### 국제 고려 사항
- 국가 코드는 자동으로 감지되어 저장됩니다.
- 번호는 국제 표준으로 형식이 지정됩니다.
- 지역 표시 선호도는 regionCode를 사용할 수 있습니다.
- 표시할 때 지역 전화 규칙을 고려하세요.

## 일반 사용 사례

1. **연락처 관리**
   - 클라이언트 전화번호
   - 공급업체 연락처 정보
   - 팀원 전화번호
   - 지원 연락처 세부정보

2. **비상 연락처**
   - 비상 연락처 번호
   - 대기 연락처 정보
   - 위기 대응 연락처
   - 에스컬레이션 전화번호

3. **고객 지원**
   - 고객 전화번호
   - 지원 콜백 번호
   - 검증 전화번호
   - 후속 연락처 번호

4. **판매 및 마케팅**
   - 리드 전화번호
   - 캠페인 연락처 목록
   - 파트너 연락처 정보
   - 추천 소스 전화

## 통합 기능

### 자동화와 함께
- 전화 필드가 업데이트될 때 작업 트리거
- 저장된 전화번호로 SMS 알림 전송
- 전화 변경에 따라 후속 작업 생성
- 전화번호 데이터에 따라 통화 라우팅

### 조회와 함께
- 다른 레코드의 전화 데이터 참조
- 여러 출처에서 전화 목록 집계
- 전화번호로 레코드 찾기
- 연락처 정보 교차 참조

### 양식과 함께
- 자동 전화 검증
- 국제 형식 확인
- 국가 코드 감지
- 실시간 형식 피드백

## 제한 사항

- 모든 번호에 국가 코드가 필요합니다.
- 내장된 SMS 또는 통화 기능이 없습니다.
- 형식 검사를 제외한 전화번호 검증이 없습니다.
- 전화 메타데이터 (통신사, 유형 등)를 저장하지 않습니다.
- 국가 코드가 없는 국가 형식 번호는 거부됩니다.
- UI에서 국제 표준을 넘어 자동 전화번호 형식 지정이 없습니다.

## 관련 리소스

- [텍스트 필드](/api/custom-fields/text-single) - 비전화 텍스트 데이터용
- [이메일 필드](/api/custom-fields/email) - 이메일 주소용
- [URL 필드](/api/custom-fields/url) - 웹사이트 주소용
- [사용자 정의 필드 개요](/custom-fields/list-custom-fields) - 일반 개념