---
title: 백분율 사용자 정의 필드
description: 자동 % 기호 처리 및 표시 형식을 갖춘 숫자 값을 저장하는 백분율 필드를 생성합니다.
---

백분율 사용자 정의 필드는 레코드에 대한 백분율 값을 저장할 수 있게 해줍니다. 이 필드는 입력 및 표시를 위한 % 기호를 자동으로 처리하며, 내부적으로는 원시 숫자 값을 저장합니다. 완료율, 성공률 또는 기타 백분율 기반 메트릭에 적합합니다.

## 기본 예제

간단한 백분율 필드를 생성합니다:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Completion Rate"
    type: PERCENT
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 백분율 필드를 생성합니다:

```graphql
mutation CreatePercentField {
  createCustomField(input: {
    name: "Success Rate"
    type: PERCENT
    description: "Percentage of successful outcomes for this process"
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
| `name` | String! | ✅ 예 | 백분율 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `PERCENT` 여야 합니다. |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 프로젝트 컨텍스트는 인증 헤더에서 자동으로 결정됩니다. `projectId` 매개변수가 필요하지 않습니다.

**참고**: PERCENT 필드는 NUMBER 필드와 같은 최소/최대 제약 조건이나 접두사 형식을 지원하지 않습니다.

## 백분율 값 설정

백분율 필드는 자동 % 기호 처리를 통해 숫자 값을 저장합니다:

### 백분율 기호와 함께

```graphql
mutation SetPercentWithSymbol {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 75.5
  }) {
    id
    customField {
      value  # Returns { number: 75.5 }
    }
  }
}
```

### 직접 숫자 값

```graphql
mutation SetPercentNumeric {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 100
  }) {
    id
    customField {
      value  # Returns { number: 100.0 }
    }
  }
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 백분율 사용자 정의 필드의 ID |
| `number` | Float | 아니오 | 숫자 백분율 값 (예: 75.5는 75.5%에 해당) |

## 값 저장 및 표시

### 저장 형식
- **내부 저장**: 원시 숫자 값 (예: 75.5)
- **데이터베이스**: `Decimal` 열에 `number`로 저장됨
- **GraphQL**: `Float` 유형으로 반환됨

### 표시 형식
- **사용자 인터페이스**: 클라이언트 애플리케이션은 % 기호를 추가해야 합니다 (예: "75.5%")
- **차트**: 출력 유형이 PERCENTAGE일 때 % 기호와 함께 표시됨
- **API 응답**: % 기호 없이 원시 숫자 값 (예: 75.5)

## 백분율 값으로 레코드 생성

백분율 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithPercent {
  createTodo(input: {
    title: "Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "success_rate_field_id"
      value: "85.5%"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Percent is accessed here as { number: 85.5 }
      }
    }
  }
}
```

### 지원되는 입력 형식

| 형식 | 예제 | 결과 |
|--------|---------|---------|
| With % symbol | `"75.5%"` | Stored as 75.5 |
| Without % symbol | `"75.5"` | Stored as 75.5 |
| Integer percentage | `"100"` | Stored as 100.0 |
| Decimal percentage | `"33.333"` | Stored as 33.333 |

**참고**: % 기호는 입력에서 자동으로 제거되고 표시할 때 다시 추가됩니다.

## 백분율 값 쿼리

백분율 사용자 정의 필드가 있는 레코드를 쿼리할 때, `customField.value.number` 경로를 통해 값을 접근합니다:

```graphql
query GetRecordWithPercent {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For PERCENT type, contains { number: 75.5 }
      }
    }
  }
}
```

응답에는 원시 숫자로서 백분율이 포함됩니다:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Completion Rate",
          "type": "PERCENT",
          "value": {
            "number": 75.5
          }
        }
      }]
    }
  }
}
```

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 필드 값에 대한 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 (백분율 값 포함) |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시점 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시점 |

**중요**: 백분율 값은 `customField.value.number` 필드를 통해 접근됩니다. % 기호는 저장된 값에 포함되지 않으며, 표시를 위해 클라이언트 애플리케이션에서 추가해야 합니다.

## 필터링 및 쿼리

백분율 필드는 NUMBER 필드와 동일한 필터링을 지원합니다:

```graphql
query FilterByPercentRange {
  todos(filter: {
    customFields: [{
      customFieldId: "completion_rate_field_id"
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
| `EQ` | 같음 | `percentage = 75` |
| `NE` | 같지 않음 | `percentage ≠ 75` |
| `GT` | 초과 | `percentage > 75` |
| `GTE` | 이상 | `percentage ≥ 75` |
| `LT` | 미만 | `percentage < 75` |
| `LTE` | 이하 | `percentage ≤ 75` |
| `IN` | 목록에 있는 값 | `percentage in [50, 75, 100]` |
| `NIN` | 목록에 없는 값 | `percentage not in [0, 25]` |
| `IS` | `values: null`로 null 확인 | `percentage is null` |
| `NOT` | `values: null`로 not null 확인 | `percentage is not null` |

### 범위 필터링

범위 필터링을 위해 여러 연산자를 사용합니다:

```graphql
query FilterHighPerformers {
  todos(filter: {
    customFields: [{
      customFieldId: "success_rate_field_id"
      operator: GTE
      number: 90
    }]
  }) {
    id
    title
    customFields {
      customField {
        value  # Returns { number: 95.5 } for example
      }
    }
  }
}
```

## 백분율 값 범위

### 일반 범위

| 범위 | 설명 | 사용 사례 |
|-------|-------------|----------|
| `0-100` | 표준 백분율 | Completion rates, success rates |
| `0-∞` | 무제한 백분율 | Growth rates, performance metrics |
| `-∞-∞` | 모든 값 | Change rates, variance |

### 예제 값

| 입력 | 저장 | 표시 |
|-------|--------|---------|
| `"50%"` | `50.0` | `50%` |
| `"100"` | `100.0` | `100%` |
| `"150.5"` | `150.5` | `150.5%` |
| `"-25"` | `-25.0` | `-25%` |

## 차트 집계

백분율 필드는 대시보드 차트 및 보고서에서 집계를 지원합니다. 사용 가능한 함수는 다음과 같습니다:

- `AVERAGE` - 평균 백분율 값
- `COUNT` - 값이 있는 레코드 수
- `MIN` - 가장 낮은 백분율 값
- `MAX` - 가장 높은 백분율 값 
- `SUM` - 모든 백분율 값의 합계

이러한 집계는 차트 및 대시보드를 생성할 때 사용할 수 있으며, 직접 GraphQL 쿼리에서는 사용할 수 없습니다.

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create percent field | `OWNER` or `ADMIN` role at project level |
| Update percent field | `OWNER` or `ADMIN` role at project level |
| Set percent value | Standard record edit permissions |
| View percent value | Standard record view permissions |
| Use chart aggregation | Standard chart viewing permissions |

## 오류 응답

### 잘못된 백분율 형식
```json
{
  "errors": [{
    "message": "Invalid percentage value",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

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

### 값 입력
- 사용자가 % 기호가 있거나 없는 상태로 입력할 수 있도록 허용
- 사용 사례에 대한 합리적인 범위를 검증
- 100%가 나타내는 것에 대한 명확한 맥락 제공

### 표시
- 사용자 인터페이스에서 항상 % 기호를 표시
- 적절한 소수점 정밀도 사용
- 범위에 대한 색상 코딩 고려 (빨강/노랑/초록)

### 데이터 해석
- 100%가 귀하의 맥락에서 의미하는 바를 문서화
- 100%를 초과하는 값을 적절하게 처리
- 음수 값이 유효한지 고려

## 일반 사용 사례

1. **프로젝트 관리**
   - 작업 완료율
   - 프로젝트 진행 상황
   - 자원 활용도
   - 스프린트 속도

2. **성능 추적**
   - 성공률
   - 오류율
   - 효율성 메트릭
   - 품질 점수

3. **재무 메트릭**
   - 성장률
   - 이익률
   - 할인 금액
   - 변화 비율

4. **분석**
   - 전환율
   - 클릭률
   - 참여 메트릭
   - 성과 지표

## 통합 기능

### 수식과 함께
- 계산에서 PERCENT 필드 참조
- 수식 출력에서 자동 % 기호 형식
- 다른 숫자 필드와 결합

### 자동화와 함께
- 백분율 임계값에 따라 작업 트리거
- 이정표 백분율에 대한 알림 전송
- 완료율에 따라 상태 업데이트

### 조회와 함께
- 관련 레코드에서 백분율 집계
- 평균 성공률 계산
- 가장 높은/낮은 성과 항목 찾기

### 차트와 함께
- 백분율 기반 시각화 생성
- 시간 경과에 따른 진행 상황 추적
- 성과 메트릭 비교

## NUMBER 필드와의 차이점

### 다른 점
- **입력 처리**: % 기호를 자동으로 제거
- **표시**: % 기호를 자동으로 추가
- **제약 조건**: 최소/최대 검증 없음
- **형식**: 접두사 지원 없음

### 동일한 점
- **저장**: 동일한 데이터베이스 열 및 유형
- **필터링**: 동일한 쿼리 연산자
- **집계**: 동일한 집계 함수
- **권한**: 동일한 권한 모델

## 제한 사항

- 최소/최대 값 제약 없음
- 접두사 형식 옵션 없음
- 0-100% 범위에 대한 자동 검증 없음
- 백분율 형식 간 변환 없음 (예: 0.75 ↔ 75%)
- 100%를 초과하는 값 허용

## 관련 리소스

- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 사용자 정의 필드 개념
- [숫자 사용자 정의 필드](/api/custom-fields/number) - 원시 숫자 값에 대한 필드
- [자동화 API](/api/automations/index) - 백분율 기반 자동화 생성