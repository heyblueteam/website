---
title: 시간 지속 시간 사용자 정의 필드
description: 워크플로우 내 이벤트 간의 시간을 추적하는 계산된 시간 지속 필드를 생성합니다.
---

시간 지속 시간 사용자 정의 필드는 워크플로우 내 두 이벤트 간의 지속 시간을 자동으로 계산하고 표시합니다. 이 필드는 처리 시간, 응답 시간, 사이클 시간 또는 프로젝트 내의 시간 기반 메트릭을 추적하는 데 이상적입니다.

## 기본 예제

작업 완료에 걸리는 시간을 추적하는 간단한 시간 지속 필드를 생성합니다:

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## 고급 예제

SLA 목표와 함께 사용자 정의 필드 변경 간의 시간을 추적하는 복잡한 시간 지속 필드를 생성합니다:

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput (TIME_DURATION)

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 지속 시간 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `TIME_DURATION` 여야 함 |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ 예 | 지속 시간을 표시하는 방법 |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ 예 | 시작 이벤트 구성 |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ 예 | 종료 이벤트 구성 |
| `timeDurationTargetTime` | Float | 아니오 | SLA 모니터링을 위한 목표 지속 시간(초) |

### CustomFieldTimeDurationInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ 예 | 추적할 이벤트 유형 |
| `condition` | CustomFieldTimeDurationCondition! | ✅ 예 | `FIRST` 또는 `LAST` 발생 |
| `customFieldId` | String | Conditional | `TODO_CUSTOM_FIELD` 유형에 필요 |
| `customFieldOptionIds` | [String!] | Conditional | 선택 필드 변경에 필요 |
| `todoListId` | String | Conditional | `TODO_MOVED` 유형에 필요 |
| `tagId` | String | Conditional | `TODO_TAG_ADDED` 유형에 필요 |
| `assigneeId` | String | Conditional | `TODO_ASSIGNEE_ADDED` 유형에 필요 |

### CustomFieldTimeDurationType 값

| 값 | 설명 |
|-------|-------------|
| `TODO_CREATED_AT` | 레코드가 생성된 시간 |
| `TODO_CUSTOM_FIELD` | 사용자 정의 필드 값이 변경된 시간 |
| `TODO_DUE_DATE` | 마감일이 설정된 시간 |
| `TODO_MARKED_AS_COMPLETE` | 레코드가 완료로 표시된 시간 |
| `TODO_MOVED` | 레코드가 다른 목록으로 이동된 시간 |
| `TODO_TAG_ADDED` | 레코드에 태그가 추가된 시간 |
| `TODO_ASSIGNEE_ADDED` | 레코드에 담당자가 추가된 시간 |

### CustomFieldTimeDurationCondition 값

| 값 | 설명 |
|-------|-------------|
| `FIRST` | 이벤트의 첫 번째 발생을 사용 |
| `LAST` | 이벤트의 마지막 발생을 사용 |

### CustomFieldTimeDurationDisplayType 값

| 값 | 설명 | 예제 |
|-------|-------------|---------|
| `FULL_DATE` | 일:시간:분:초 형식 | `"01:02:03:04"` |
| `FULL_DATE_STRING` | 전체 단어로 작성됨 | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | 단위가 있는 숫자 | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | 일 단위의 지속 시간 | `"2.5"` (2.5 days) |
| `HOURS` | 시간 단위의 지속 시간 | `"60"` (60 hours) |
| `MINUTES` | 분 단위의 지속 시간 | `"3600"` (3600 minutes) |
| `SECONDS` | 초 단위의 지속 시간 | `"216000"` (216000 seconds) |

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `number` | Float | 지속 시간(초) |
| `value` | Float | 숫자에 대한 별칭(지속 시간(초)) |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 업데이트된 시간 |

### CustomField 응답 (TIME_DURATION)

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | 지속 시간의 표시 형식 |
| `timeDurationStart` | CustomFieldTimeDuration | 시작 이벤트 구성 |
| `timeDurationEnd` | CustomFieldTimeDuration | 종료 이벤트 구성 |
| `timeDurationTargetTime` | Float | SLA 모니터링을 위한 목표 지속 시간(초) |

## 지속 시간 계산

### 작동 원리
1. **시작 이벤트**: 시스템이 지정된 시작 이벤트를 모니터링합니다.
2. **종료 이벤트**: 시스템이 지정된 종료 이벤트를 모니터링합니다.
3. **계산**: 지속 시간 = 종료 시간 - 시작 시간
4. **저장**: 지속 시간이 숫자로 초 단위로 저장됩니다.
5. **표시**: `timeDurationDisplay` 설정에 따라 형식화됩니다.

### 업데이트 트리거
지속 시간 값은 다음과 같은 경우 자동으로 재계산됩니다:
- 레코드가 생성되거나 업데이트될 때
- 사용자 정의 필드 값이 변경될 때
- 태그가 추가되거나 제거될 때
- 담당자가 추가되거나 제거될 때
- 레코드가 목록 간에 이동될 때
- 레코드가 완료/미완료로 표시될 때

## 지속 시간 값 읽기

### 지속 시간 필드 쿼리
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### 형식화된 표시 값
지속 시간 값은 `timeDurationDisplay` 설정에 따라 자동으로 형식화됩니다:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## 일반 구성 예제

### 작업 완료 시간
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### 상태 변경 지속 시간
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### 특정 목록의 시간
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### 할당 응답 시간
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## 필수 권한

| 작업 | 필요한 권한 |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## 오류 응답

### 잘못된 구성
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### 참조된 필드를 찾을 수 없음
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### 필수 옵션 누락
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 중요 사항

### 자동 계산
- 지속 시간 필드는 **읽기 전용**입니다 - 값은 자동으로 계산됩니다.
- API를 통해 지속 시간 값을 수동으로 설정할 수 없습니다.
- 계산은 백그라운드 작업을 통해 비동기적으로 발생합니다.
- 트리거 이벤트가 발생할 때 값이 자동으로 업데이트됩니다.

### 성능 고려사항
- 지속 시간 계산은 대기열에 추가되고 비동기적으로 처리됩니다.
- 많은 수의 지속 시간 필드는 성능에 영향을 미칠 수 있습니다.
- 지속 시간 필드를 설계할 때 트리거 이벤트의 빈도를 고려하십시오.
- 불필요한 재계산을 피하기 위해 특정 조건을 사용하십시오.

### 널 값
지속 시간 필드는 다음과 같은 경우 `null`를 표시합니다:
- 시작 이벤트가 아직 발생하지 않았습니다.
- 종료 이벤트가 아직 발생하지 않았습니다.
- 구성에서 존재하지 않는 엔터티를 참조합니다.
- 계산 중 오류가 발생했습니다.

## 모범 사례

### 구성 설계
- 가능하면 일반적인 것보다 특정 이벤트 유형을 사용하십시오.
- 워크플로우에 따라 적절한 `FIRST`와 `LAST` 조건을 선택하십시오.
- 배포 전에 샘플 데이터를 사용하여 지속 시간 계산을 테스트하십시오.
- 팀원들을 위해 지속 시간 필드 로직을 문서화하십시오.

### 표시 형식
- 가장 읽기 쉬운 형식을 위해 `FULL_DATE_SUBSTRING`를 사용하십시오.
- 일관된 너비의 간결한 표시를 위해 `FULL_DATE`를 사용하십시오.
- 공식 보고서 및 문서에는 `FULL_DATE_STRING`를 사용하십시오.
- 간단한 숫자 표시를 위해 `DAYS`, `HOURS`, `MINUTES` 또는 `SECONDS`를 고려하십시오.
- 형식을 선택할 때 UI 공간 제약을 고려하십시오.

### SLA 모니터링과 목표 시간
`timeDurationTargetTime`를 사용할 때:
- 목표 지속 시간을 초 단위로 설정하십시오.
- SLA 준수를 위해 실제 지속 시간을 목표와 비교하십시오.
- 대시보드에서 연체 항목을 강조 표시하는 데 사용하십시오.
- 예: 24시간 응답 SLA = 86400초

### 워크플로우 통합
- 실제 비즈니스 프로세스에 맞게 지속 시간 필드를 설계하십시오.
- 프로세스 개선 및 최적화를 위해 지속 시간 데이터를 사용하십시오.
- 워크플로우 병목 현상을 식별하기 위해 지속 시간 추세를 모니터링하십시오.
- 필요할 경우 지속 시간 임계값에 대한 알림을 설정하십시오.

## 일반 사용 사례

1. **프로세스 성과**
   - 작업 완료 시간
   - 검토 사이클 시간
   - 승인 처리 시간
   - 응답 시간

2. **SLA 모니터링**
   - 첫 번째 응답까지의 시간
   - 해결 시간
   - 에스컬레이션 시간
   - 서비스 수준 준수

3. **워크플로우 분석**
   - 병목 현상 식별
   - 프로세스 최적화
   - 팀 성과 메트릭
   - 품질 보증 타이밍

4. **프로젝트 관리**
   - 단계 지속 시간
   - 이정표 타이밍
   - 자원 할당 시간
   - 납기 시간

## 제한 사항

- 지속 시간 필드는 **읽기 전용**이며 수동으로 설정할 수 없습니다.
- 값은 비동기적으로 계산되며 즉시 사용할 수 없을 수 있습니다.
- 워크플로우에 적절한 이벤트 트리거가 설정되어야 합니다.
- 발생하지 않은 이벤트에 대한 지속 시간을 계산할 수 없습니다.
- 이산 이벤트 간의 시간 추적에만 제한됩니다(연속 시간 추적 아님).
- 내장 SLA 알림이나 알림이 없습니다.
- 여러 지속 시간 계산을 단일 필드로 집계할 수 없습니다.

## 관련 리소스

- [숫자 필드](/api/custom-fields/number) - 수동 숫자 값에 대해
- [날짜 필드](/api/custom-fields/date) - 특정 날짜 추적에 대해
- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 개념
- [자동화](/api/automations) - 지속 시간 임계값에 따라 작업 트리거하기 위해