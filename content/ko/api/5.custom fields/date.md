---
title: 날짜 사용자 정의 필드
description: 단일 날짜 또는 날짜 범위를 추적하기 위한 날짜 필드를 생성하며, 시간대 지원을 포함합니다.
---

날짜 사용자 정의 필드는 레코드에 단일 날짜 또는 날짜 범위를 저장할 수 있게 해줍니다. 이들은 시간대 처리, 지능형 형식을 지원하며, 마감일, 이벤트 날짜 또는 시간 기반 정보를 추적하는 데 사용할 수 있습니다.

## 기본 예제

간단한 날짜 필드를 생성합니다:

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 마감일 필드를 생성합니다:

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 날짜 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `DATE` 여야 합니다. |
| `isDueDate` | Boolean | 아니오 | 이 필드가 마감일을 나타내는지 여부 |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 사용자 정의 필드는 사용자의 현재 프로젝트 컨텍스트에 따라 자동으로 프로젝트와 연결됩니다. `projectId` 매개변수가 필요하지 않습니다.

## 날짜 값 설정

날짜 필드는 단일 날짜 또는 날짜 범위를 저장할 수 있습니다:

### 단일 날짜

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### 날짜 범위

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### 종일 이벤트

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 날짜 사용자 정의 필드의 ID |
| `startDate` | DateTime | 아니오 | ISO 8601 형식의 시작 날짜/시간 |
| `endDate` | DateTime | 아니오 | ISO 8601 형식의 종료 날짜/시간 |
| `timezone` | String | 아니오 | 시간대 식별자 (예: "America/New_York") |

**참고**: 만약 `startDate`만 제공되면, `endDate`는 자동으로 동일한 값으로 기본 설정됩니다.

## 날짜 형식

### ISO 8601 형식
모든 날짜는 ISO 8601 형식으로 제공되어야 합니다:
- `2025-01-15T14:30:00Z` - UTC 시간
- `2025-01-15T14:30:00+05:00` - 시간대 오프셋 포함
- `2025-01-15T14:30:00.123Z` - 밀리초 포함

### 시간대 식별자
표준 시간대 식별자를 사용하십시오:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

시간대가 제공되지 않으면 시스템은 사용자의 감지된 시간대로 기본 설정됩니다.

## 날짜 값을 가진 레코드 생성

날짜 값을 가진 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### 지원되는 입력 형식

레코드를 생성할 때 날짜는 다양한 형식으로 제공될 수 있습니다:

| 형식 | 예제 | 결과 |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 필드 값의 고유 식별자 |
| `uid` | String! | 고유 식별자 문자열 |
| `customField` | CustomField! | 사용자 정의 필드 정의 (날짜 값 포함) |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

**중요**: 날짜 값 (`startDate`, `endDate`, `timezone`)은 `customField.value` 필드를 통해 접근하며, TodoCustomField에서 직접 접근하지 않습니다.

### 값 객체 구조

날짜 값은 `customField.value` 필드를 통해 JSON 객체로 반환됩니다:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**참고**: `value` 필드는 `CustomField` 유형에 있으며, `TodoCustomField`에는 없습니다.

## 날짜 값 쿼리

날짜 사용자 정의 필드가 있는 레코드를 쿼리할 때, 날짜 값은 `customField.value` 필드를 통해 접근합니다:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

응답에는 `value` 필드에 날짜 값이 포함됩니다:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## 날짜 표시 지능

시스템은 범위에 따라 날짜를 자동으로 형식화합니다:

| 시나리오 | 표시 형식 |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (시간 표시 안 함) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**종일 감지**: 00:00부터 23:59까지의 이벤트는 자동으로 종일 이벤트로 감지됩니다.

## 시간대 처리

### 저장
- 모든 날짜는 데이터베이스에 UTC로 저장됩니다.
- 시간대 정보는 별도로 보존됩니다.
- 변환은 표시 시 발생합니다.

### 모범 사례
- 정확성을 위해 항상 시간대를 제공하십시오.
- 프로젝트 내에서 일관된 시간대를 사용하십시오.
- 글로벌 팀의 사용자 위치를 고려하십시오.

### 일반적인 시간대

| 지역 | 시간대 ID | UTC 오프셋 |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## 필터링 및 쿼리

날짜 필드는 복잡한 필터링을 지원합니다:

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### 빈 날짜 필드 확인

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### 지원되는 연산자

| 연산자 | 사용법 | 설명 |
|----------|-------|-------------|
| `EQ` | dateRange와 함께 | 날짜가 지정된 범위와 겹침 (모든 교차) |
| `NE` | dateRange와 함께 | 날짜가 범위와 겹치지 않음 |
| `IS` | `values: null`와 함께 | 날짜 필드가 비어 있음 (startDate 또는 endDate가 null) |
| `NOT` | `values: null`와 함께 | 날짜 필드에 값이 있음 (두 날짜가 null이 아님) |

## 필요한 권한

| 작업 | 필요한 권한 |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## 오류 응답

### 잘못된 날짜 형식
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
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
      "code": "NOT_FOUND"
    }
  }]
}
```


## 제한 사항

- 반복 날짜 지원 없음 (반복 이벤트에 대한 자동화 사용)
- 날짜 없이 시간 설정 불가
- 내장된 근무일 계산 없음
- 날짜 범위가 종료 > 시작을 자동으로 검증하지 않음
- 최대 정밀도는 초 단위 (밀리초 저장 없음)

## 관련 리소스

- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 사용자 정의 필드 개념
- [자동화 API](/api/automations/index) - 날짜 기반 자동화 생성