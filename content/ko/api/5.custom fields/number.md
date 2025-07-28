---
title: 숫자 사용자 정의 필드
description: 선택적 최소/최대 제약 조건 및 접두사 형식을 사용하여 숫자 값을 저장할 숫자 필드를 생성합니다.
---

숫자 사용자 정의 필드는 레코드에 숫자 값을 저장할 수 있게 해줍니다. 이들은 유효성 검사 제약 조건, 소수점 정밀도를 지원하며, 수량, 점수, 측정값 또는 특별한 형식이 필요하지 않은 모든 숫자 데이터에 사용할 수 있습니다.

## 기본 예제

간단한 숫자 필드를 생성합니다:

```graphql
mutation CreateNumberField {
  createCustomField(input: {
    name: "Priority Score"
    type: NUMBER
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

제약 조건 및 접두사가 있는 숫자 필드를 생성합니다:

```graphql
mutation CreateConstrainedNumberField {
  createCustomField(input: {
    name: "Team Size"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 100
    prefix: "#"
    description: "Number of team members assigned to this project"
  }) {
    id
    name
    type
    min
    max
    prefix
    description
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 숫자 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `NUMBER` 여야 합니다. |
| `projectId` | String! | ✅ 예 | 필드를 생성할 프로젝트의 ID |
| `min` | Float | 아니오 | 최소값 제약 조건 (UI 안내용) |
| `max` | Float | 아니오 | 최대값 제약 조건 (UI 안내용) |
| `prefix` | String | 아니오 | 표시 접두사 (예: "#", "~", "$") |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

## 숫자 값 설정

숫자 필드는 선택적 유효성 검사와 함께 소수값을 저장합니다:

### 간단한 숫자 값

```graphql
mutation SetNumberValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 42.5
  })
}
```

### 정수 값

```graphql
mutation SetIntegerValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 숫자 사용자 정의 필드의 ID |
| `number` | Float | 아니오 | 저장할 숫자 값 |

## 값 제약 조건

### 최소/최대 제약 조건 (UI 안내)

**중요**: 최소/최대 제약 조건은 저장되지만 서버 측에서 강제 적용되지 않습니다. 이들은 프론트엔드 애플리케이션을 위한 UI 안내 역할을 합니다.

```graphql
mutation CreateConstrainedField {
  createCustomField(input: {
    name: "Rating"
    type: NUMBER
    projectId: "proj_123"
    min: 1
    max: 10
    description: "Rating from 1 to 10"
  }) {
    id
    name
    min
    max
  }
}
```

**클라이언트 측 유효성 검사 필요**: 프론트엔드 애플리케이션은 최소/최대 제약 조건을 강제하기 위해 유효성 검사 로직을 구현해야 합니다.

### 지원되는 값 유형

| 유형 | 예제 | 설명 |
|------|---------|-------------|
| Integer | `42` | 정수 |
| Decimal | `42.5` | 소수점이 있는 숫자 |
| Negative | `-10` | 음수 (최소 제약 조건이 없는 경우) |
| Zero | `0` | 제로 값 |

**참고**: 최소/최대 제약 조건은 서버 측에서 유효성 검증되지 않습니다. 지정된 범위를 벗어난 값은 수락되고 저장됩니다.

## 숫자 값으로 레코드 생성

숫자 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithNumber {
  createTodo(input: {
    title: "Performance Review"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "score_field_id"
      number: 85.5
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        min
        max
        prefix
      }
      number
      value
    }
  }
}
```

### 지원되는 입력 형식

레코드를 생성할 때 사용자 정의 필드 배열에서 `number` 매개변수를 사용합니다 (`value` 아님):

```graphql
customFields: [{
  customFieldId: "field_id"
  number: 42.5  # Use number parameter, not value
}]
```

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `number` | Float | 숫자 값 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

### CustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 정의의 고유 식별자 |
| `name` | String! | 필드의 표시 이름 |
| `type` | CustomFieldType! | 항상 `NUMBER` |
| `min` | Float | 허용되는 최소값 |
| `max` | Float | 허용되는 최대값 |
| `prefix` | String | 표시 접두사 |
| `description` | String | 도움말 텍스트 |

**참고**: 숫자 값이 설정되지 않은 경우 `number` 필드는 `null`입니다.

## 필터링 및 쿼리

숫자 필드는 포괄적인 숫자 필터링을 지원합니다:

```graphql
query FilterByNumberRange {
  todos(filter: {
    customFields: [{
      customFieldId: "score_field_id"
      operator: GTE
      number: 80
    }]
  }) {
    id
    title
    customFields {
      number
    }
  }
}
```

### 지원되는 연산자

| 연산자 | 설명 | 예제 |
|----------|-------------|---------|
| `EQ` | 같음 | `number = 42` |
| `NE` | 같지 않음 | `number ≠ 42` |
| `GT` | 큼 | `number > 42` |
| `GTE` | 크거나 같음 | `number ≥ 42` |
| `LT` | 작음 | `number < 42` |
| `LTE` | 작거나 같음 | `number ≤ 42` |
| `IN` | 배열에 포함 | `number in [1, 2, 3]` |
| `NIN` | 배열에 포함되지 않음 | `number not in [1, 2, 3]` |
| `IS` | null/비 null | `number is null` |

### 범위 필터링

```graphql
query FilterByRange {
  todos(filter: {
    customFields: [{
      customFieldId: "priority_field_id"
      operator: GTE
      number: 5
    }]
  }) {
    id
    title
  }
}
```

## 표시 형식

### 접두사와 함께

접두사가 설정된 경우 표시됩니다:

| 값 | 접두사 | 표시 |
|-------|--------|---------|
| `42` | `"#"` | `#42` |
| `100` | `"~"` | `~100` |
| `3.14` | `"π"` | `π3.14` |

### 소수점 정밀도

숫자는 소수점 정밀도를 유지합니다:

| 입력 | 저장됨 | 표시됨 |
|-------|--------|-----------|
| `42` | `42.0` | `42` |
| `42.5` | `42.5` | `42.5` |
| `42.123` | `42.123` | `42.123` |

## 필수 권한

| 작업 | 필요한 권한 |
|--------|--------------------|
| Create number field | Company role: `OWNER` or `ADMIN` |
| Update number field | Company role: `OWNER` or `ADMIN` |
| Set number value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View number value | Standard record view permissions |
| Use in filtering | Standard record view permissions |

## 오류 응답

### 잘못된 숫자 형식
```json
{
  "errors": [{
    "message": "Invalid number format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 필드를 찾을 수 없음
```json
{
  "errors": [{
    "message": "Custom field was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

**참고**: 최소/최대 유효성 검사 오류는 서버 측에서 발생하지 않습니다. 제약 조건 유효성 검사는 프론트엔드 애플리케이션에서 구현해야 합니다.

### 숫자가 아님
```json
{
  "errors": [{
    "message": "Value is not a valid number",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 모범 사례

### 제약 조건 설계
- UI 안내를 위해 현실적인 최소/최대 값을 설정합니다.
- 제약 조건을 강제하기 위해 클라이언트 측 유효성 검사를 구현합니다.
- 양식에서 사용자 피드백을 제공하기 위해 제약 조건을 사용합니다.
- 음수 값이 사용 사례에 적합한지 고려합니다.

### 값 정밀도
- 필요에 맞는 적절한 소수점 정밀도를 사용합니다.
- 표시 목적으로 반올림을 고려합니다.
- 관련 필드 간에 정밀도를 일관되게 유지합니다.

### 표시 향상
- 맥락에 맞는 의미 있는 접두사를 사용합니다.
- 필드 이름에 단위를 고려합니다 (예: "무게 (kg)").
- 유효성 검사 규칙에 대한 명확한 설명을 제공합니다.

## 일반적인 사용 사례

1. **점수 시스템**
   - 성과 평가
   - 품질 점수
   - 우선 순위 수준
   - 고객 만족도 평가

2. **측정**
   - 수량 및 금액
   - 치수 및 크기
   - 기간 (숫자 형식으로)
   - 용량 및 한계

3. **비즈니스 메트릭**
   - 수익 수치
   - 전환율
   - 예산 배정
   - 목표 수치

4. **기술 데이터**
   - 버전 번호
   - 구성 값
   - 성능 메트릭
   - 임계값 설정

## 통합 기능

### 차트 및 대시보드와 함께
- 차트 계산에서 숫자 필드를 사용합니다.
- 숫자 시각화를 생성합니다.
- 시간에 따른 추세를 추적합니다.

### 자동화와 함께
- 숫자 임계값에 따라 작업을 트리거합니다.
- 숫자 변경에 따라 관련 필드를 업데이트합니다.
- 특정 값에 대한 알림을 보냅니다.

### 조회와 함께
- 관련 레코드에서 숫자를 집계합니다.
- 총계 및 평균을 계산합니다.
- 관계 전반에 걸쳐 최소/최대 값을 찾습니다.

### 차트와 함께
- 숫자 시각화를 생성합니다.
- 시간에 따른 추세를 추적합니다.
- 레코드 간의 값을 비교합니다.

## 제한 사항

- **서버 측에서 최소/최대 제약 조건의 유효성 검사 없음**
- **제약 조건 강제를 위한 클라이언트 측 유효성 검사 필요**
- 내장된 통화 형식 없음 (대신 CURRENCY 유형 사용)
- 자동 백분율 기호 없음 (대신 PERCENT 유형 사용)
- 단위 변환 기능 없음
- 소수점 정밀도는 데이터베이스 Decimal 유형에 의해 제한됨
- 필드 자체에서 수학적 공식 평가 없음

## 관련 리소스

- [사용자 정의 필드 개요](/api/custom-fields/1.index) - 일반 사용자 정의 필드 개념
- [통화 사용자 정의 필드](/api/custom-fields/currency) - 금전적 값에 대한 필드
- [백분율 사용자 정의 필드](/api/custom-fields/percent) - 백분율 값에 대한 필드
- [자동화 API](/api/automations/1.index) - 숫자 기반 자동화 생성