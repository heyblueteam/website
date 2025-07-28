---
title: 평점 사용자 정의 필드
description: 구성 가능한 척도와 유효성 검사를 통해 숫자 평점을 저장하는 평점 필드를 생성합니다.
---

평점 사용자 정의 필드는 구성 가능한 최소 및 최대 값을 가진 레코드에 숫자 평점을 저장할 수 있게 해줍니다. 성과 평점, 만족도 점수, 우선 순위 수준 또는 프로젝트의 숫자 기반 데이터에 이상적입니다.

## 기본 예제

기본 0-5 척도로 간단한 평점 필드를 생성합니다:

```graphql
mutation CreateRatingField {
  createCustomField(input: {
    name: "Performance Rating"
    type: RATING
    projectId: "proj_123"
    max: 5
  }) {
    id
    name
    type
    min
    max
  }
}
```

## 고급 예제

사용자 정의 척도와 설명이 있는 평점 필드를 생성합니다:

```graphql
mutation CreateDetailedRatingField {
  createCustomField(input: {
    name: "Customer Satisfaction"
    type: RATING
    projectId: "proj_123"
    description: "Rate customer satisfaction from 1-10"
    min: 1
    max: 10
  }) {
    id
    name
    type
    description
    min
    max
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 평점 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `RATING` 이어야 합니다. |
| `projectId` | String! | ✅ 예 | 이 필드가 생성될 프로젝트 ID |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |
| `min` | Float | 아니오 | 최소 평점 값 (기본값 없음) |
| `max` | Float | 아니오 | 최대 평점 값 |

## 평점 값 설정

레코드에서 평점 값을 설정하거나 업데이트하려면:

```graphql
mutation SetRatingValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    value: "4.5"
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 평점 사용자 정의 필드의 ID |
| `value` | String! | ✅ 예 | 문자열로 된 평점 값 (구성된 범위 내) |

## 평점 값으로 레코드 생성

평점 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithRating {
  createTodo(input: {
    title: "Review customer feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "rating_field_id"
      value: "4.5"
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
      }
      value
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
| `value` | Float | 저장된 평점 값 (customField.value를 통해 접근) |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

**참고**: 평점 값은 실제로 `customField.value.number`를 통해 쿼리에서 접근됩니다.

### CustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드의 고유 식별자 |
| `name` | String! | 평점 필드의 표시 이름 |
| `type` | CustomFieldType! | 항상 `RATING` |
| `min` | Float | 허용되는 최소 평점 값 |
| `max` | Float | 허용되는 최대 평점 값 |
| `description` | String | 필드에 대한 도움말 텍스트 |

## 평점 유효성 검사

### 값 제약 조건
- 평점 값은 숫자여야 합니다 (Float 유형)
- 값은 구성된 최소/최대 범위 내에 있어야 합니다
- 최소값이 지정되지 않은 경우 기본값이 없습니다
- 최대값은 선택 사항이지만 권장됩니다

### 유효성 검사 규칙
**중요**: 유효성 검사는 양식을 제출할 때만 발생하며, `setTodoCustomField`를 직접 사용할 때는 발생하지 않습니다.

- 입력은 부동 소수점 숫자로 구문 분석됩니다 (양식 사용 시)
- 최소값보다 크거나 같아야 합니다 (양식 사용 시)
- 최대값보다 작거나 같아야 합니다 (양식 사용 시)
- `setTodoCustomField`는 유효성 검사 없이 모든 문자열 값을 허용합니다

### 유효한 평점 예제
최소=1, 최대=5인 필드의 경우:
```
1       # Minimum value
5       # Maximum value
3.5     # Decimal values allowed
2.75    # Precise decimal ratings
```

### 유효하지 않은 평점 예제
최소=1, 최대=5인 필드의 경우:
```
0       # Below minimum
6       # Above maximum
-1      # Negative value (below min)
abc     # Non-numeric value
```

## 구성 옵션

### 평점 척도 설정
```graphql
# 1-5 star rating
mutation CreateStarRating {
  createCustomField(input: {
    name: "Star Rating"
    type: RATING
    projectId: "proj_123"
    min: 1
    max: 5
  }) {
    id
    min
    max
  }
}

# 0-100 percentage rating
mutation CreatePercentageRating {
  createCustomField(input: {
    name: "Completion Percentage"
    type: RATING
    projectId: "proj_123"
    min: 0
    max: 100
  }) {
    id
    min
    max
  }
}
```

### 일반 평점 척도
- **1-5 별**: `min: 1, max: 5`
- **0-10 NPS**: `min: 0, max: 10`
- **1-10 성과**: `min: 1, max: 10`
- **0-100 백분율**: `min: 0, max: 100`
- **사용자 정의 척도**: 임의의 숫자 범위

## 필수 권한

사용자 정의 필드 작업은 표준 역할 기반 권한을 따릅니다:

| 작업 | 필수 역할 |
|--------|---------------|
| Create rating field | Project member with appropriate role |
| Update rating field | Project member with appropriate role |
| Set rating value | Project member with field edit permissions |
| View rating value | Project member with view permissions |

**참고**: 특정 역할 요구 사항은 프로젝트의 사용자 정의 역할 구성 및 필드 수준 권한에 따라 다릅니다.

## 오류 응답

### 유효성 검사 오류 (양식 전용)
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

**중요**: 평점 값 유효성 검사 (최소/최대 제약 조건)는 양식을 제출할 때만 발생하며, `setTodoCustomField`를 직접 사용할 때는 발생하지 않습니다.

### 사용자 정의 필드 찾을 수 없음
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

## 모범 사례

### 척도 설계
- 유사한 필드에서 일관된 평점 척도를 사용하십시오
- 사용자 친숙성을 고려하십시오 (1-5 별, 0-10 NPS)
- 적절한 최소 값을 설정하십시오 (0 대 1)
- 각 평점 수준에 대한 명확한 의미를 정의하십시오

### 데이터 품질
- 저장하기 전에 평점 값을 검증하십시오
- 소수점 정밀도를 적절하게 사용하십시오
- 표시 목적으로 반올림을 고려하십시오
- 평점 의미에 대한 명확한 지침을 제공하십시오

### 사용자 경험
- 평점 척도를 시각적으로 표시하십시오 (별, 진행률 표시줄)
- 현재 값과 척도 한계를 표시하십시오
- 평점 의미에 대한 맥락을 제공하십시오
- 새 레코드에 대한 기본값을 고려하십시오

## 일반 사용 사례

1. **성과 관리**
   - 직원 성과 평점
   - 프로젝트 품질 점수
   - 작업 완료 평점
   - 기술 수준 평가

2. **고객 피드백**
   - 만족도 평점
   - 제품 품질 점수
   - 서비스 경험 평점
   - 넷 프로모터 점수 (NPS)

3. **우선 순위 및 중요성**
   - 작업 우선 순위 수준
   - 긴급성 평점
   - 위험 평가 점수
   - 영향 평점

4. **품질 보증**
   - 코드 검토 평점
   - 테스트 품질 점수
   - 문서 품질
   - 프로세스 준수 평점

## 통합 기능

### 자동화와 함께
- 평점 임계값에 따라 작업 트리거
- 낮은 평점에 대한 알림 전송
- 높은 평점에 대한 후속 작업 생성
- 평점 값에 따라 작업 라우팅

### 조회와 함께
- 레코드 간 평균 평점 계산
- 평점 범위로 레코드 찾기
- 다른 레코드의 평점 데이터 참조
- 평점 통계 집계

### Blue 프론트엔드와 함께
- 양식 컨텍스트에서 자동 범위 유효성 검사
- 시각적 평점 입력 컨트롤
- 실시간 유효성 검사 피드백
- 별 또는 슬라이더 입력 옵션

## 활동 추적

평점 필드 변경 사항은 자동으로 추적됩니다:
- 이전 및 새로운 평점 값이 기록됩니다
- 활동은 숫자 변경 사항을 보여줍니다
- 모든 평점 업데이트에 대한 타임스탬프
- 변경 사항에 대한 사용자 귀속

## 제한 사항

- 숫자 값만 지원됩니다
- 내장된 시각적 평점 표시 (별 등)가 없습니다
- 소수점 정밀도는 데이터베이스 구성에 따라 다릅니다
- 평점 메타데이터 저장 (댓글, 맥락)이 없습니다
- 자동 평점 집계 또는 통계가 없습니다
- 척도 간 평점 변환이 내장되어 있지 않습니다
- **중요**: 최소/최대 유효성 검사는 양식에서만 작동하며, `setTodoCustomField`를 통해서는 작동하지 않습니다.

## 관련 리소스

- [숫자 필드](/api/5.custom%20fields/number) - 일반 숫자 데이터용
- [백분율 필드](/api/5.custom%20fields/percent) - 백분율 값용
- [선택 필드](/api/5.custom%20fields/select-single) - 이산 선택 평점용
- [사용자 정의 필드 개요](/api/5.custom%20fields/2.list-custom-fields) - 일반 개념