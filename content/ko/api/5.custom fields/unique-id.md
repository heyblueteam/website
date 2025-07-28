---
title: 고유 ID 사용자 정의 필드
description: 순차 번호 및 사용자 정의 형식을 가진 자동 생성된 고유 식별자 필드를 생성합니다.
---

고유 ID 사용자 정의 필드는 기록을 위해 순차적이고 고유한 식별자를 자동으로 생성합니다. 티켓 번호, 주문 ID, 송장 번호 또는 워크플로우 내의 모든 순차 식별자 시스템을 만드는 데 적합합니다.

## 기본 예제

자동 순차 생성으로 간단한 고유 ID 필드를 생성합니다:

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## 고급 예제

접두사와 제로 패딩이 있는 형식화된 고유 ID 필드를 생성합니다:

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput (UNIQUE_ID)

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 고유 ID 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `UNIQUE_ID` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |
| `useSequenceUniqueId` | Boolean | 아니요 | 자동 순차 생성 활성화 (기본값: false) |
| `prefix` | String | 아니요 | 생성된 ID의 텍스트 접두사 (예: "TASK-") |
| `sequenceDigits` | Int | 아니요 | 제로 패딩을 위한 숫자 자리수 |
| `sequenceStartingNumber` | Int | 아니요 | 시퀀스의 시작 번호 |

## 구성 옵션

### 자동 순차 생성 (`useSequenceUniqueId`)
- **true**: 기록이 생성될 때 자동으로 순차 ID가 생성됩니다.
- **false** 또는 **undefined**: 수동 입력이 필요합니다 (텍스트 필드처럼 작동).

### 접두사 (`prefix`)
- 생성된 모든 ID에 추가되는 선택적 텍스트 접두사
- 예시: "TASK-", "ORD-", "BUG-", "REQ-"
- 길이 제한은 없지만 표시를 위해 합리적으로 유지하십시오.

### 시퀀스 자리수 (`sequenceDigits`)
- 시퀀스 번호의 제로 패딩을 위한 자리수
- 예시: `sequenceDigits: 3`는 `001`, `002`, `003`를 생성합니다.
- 지정하지 않으면 패딩이 적용되지 않습니다.

### 시작 번호 (`sequenceStartingNumber`)
- 시퀀스의 첫 번째 번호
- 예시: `sequenceStartingNumber: 1000`는 1000, 1001, 1002...에서 시작합니다.
- 지정하지 않으면 1에서 시작합니다 (기본 동작).

## 생성된 ID 형식

최종 ID 형식은 모든 구성 옵션을 결합합니다:

```
{prefix}{paddedSequenceNumber}
```

### 형식 예제

| 구성 | 생성된 ID |
|---------------|---------------|
| 옵션 없음 | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## 고유 ID 값 읽기

### 고유 ID로 기록 쿼리
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### 응답 형식
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## 자동 ID 생성

### ID가 생성되는 경우
- **기록 생성**: 새 기록이 생성될 때 ID가 자동으로 할당됩니다.
- **필드 추가**: 기존 기록에 UNIQUE_ID 필드를 추가할 때 백그라운드 작업이 대기열에 추가됩니다 (작업자 구현 대기 중).
- **백그라운드 처리**: 새 기록에 대한 ID 생성은 데이터베이스 트리거를 통해 동기적으로 발생합니다.

### 생성 프로세스
1. **트리거**: 새 기록이 생성되거나 UNIQUE_ID 필드가 추가됩니다.
2. **시퀀스 조회**: 시스템이 다음 사용 가능한 시퀀스 번호를 찾습니다.
3. **ID 할당**: 시퀀스 번호가 기록에 할당됩니다.
4. **카운터 업데이트**: 향후 기록을 위해 시퀀스 카운터가 증가합니다.
5. **형식화**: ID가 표시될 때 접두사와 패딩으로 형식화됩니다.

### 고유성 보장
- **데이터베이스 제약 조건**: 각 필드 내 시퀀스 ID에 대한 고유 제약 조건.
- **원자적 작업**: 시퀀스 생성은 중복을 방지하기 위해 데이터베이스 잠금을 사용합니다.
- **프로젝트 범위**: 시퀀스는 프로젝트별로 독립적입니다.
- **경쟁 조건 보호**: 동시 요청이 안전하게 처리됩니다.

## 수동 모드 대 자동 모드

### 자동 모드 (`useSequenceUniqueId: true`)
- ID가 데이터베이스 트리거를 통해 자동으로 생성됩니다.
- 순차 번호가 보장됩니다.
- 원자적 시퀀스 생성이 중복을 방지합니다.
- 형식화된 ID는 접두사 + 패딩된 시퀀스 번호를 결합합니다.

### 수동 모드 (`useSequenceUniqueId: false` 또는 `undefined`)
- 일반 텍스트 필드처럼 작동합니다.
- 사용자는 `setTodoCustomField`와 `text` 매개변수를 통해 사용자 정의 값을 입력할 수 있습니다.
- 자동 생성이 없습니다.
- 데이터베이스 제약 조건을 초과하는 고유성 강제가 없습니다.

## 수동 값 설정 (수동 모드 전용)

`useSequenceUniqueId`가 false일 때, 값을 수동으로 설정할 수 있습니다:

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## 응답 필드

### TodoCustomField 응답 (UNIQUE_ID)

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `sequenceId` | Int | 생성된 시퀀스 번호 (UNIQUE_ID 필드에 대해 채워짐) |
| `text` | String | 형식화된 텍스트 값 (접두사 + 패딩된 시퀀스를 결합) |
| `todo` | Todo! | 이 값이 속한 기록 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 업데이트된 시간 |

### CustomField 응답 (UNIQUE_ID)

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | 자동 순차 생성이 활성화되어 있는지 여부 |
| `prefix` | String | 생성된 ID의 텍스트 접두사 |
| `sequenceDigits` | Int | 제로 패딩을 위한 자리수 |
| `sequenceStartingNumber` | Int | 시퀀스의 시작 번호 |

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## 오류 응답

### 필드 구성 오류
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### 권한 오류
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## 중요 참고 사항

### 자동 생성된 ID
- **읽기 전용**: 자동 생성된 ID는 수동으로 편집할 수 없습니다.
- **영구적**: 할당된 후 시퀀스 ID는 변경되지 않습니다.
- **시간 순서**: ID는 생성 순서를 반영합니다.
- **범위 지정**: 시퀀스는 프로젝트별로 독립적입니다.

### 성능 고려 사항
- 새 기록에 대한 ID 생성은 데이터베이스 트리거를 통해 동기적으로 발생합니다.
- 시퀀스 생성은 `FOR UPDATE` 잠금을 사용하여 원자적 작업을 수행합니다.
- 백그라운드 작업 시스템이 존재하지만 작업자 구현은 대기 중입니다.
- 대량 프로젝트를 위해 시퀀스 시작 번호를 고려하십시오.

### 마이그레이션 및 업데이트
- 기존 기록에 자동 순차 생성을 추가하면 백그라운드 작업이 대기열에 추가됩니다 (작업자 대기 중).
- 시퀀스 설정 변경은 향후 기록에만 영향을 미칩니다.
- 구성 업데이트 시 기존 ID는 변경되지 않습니다.
- 시퀀스 카운터는 현재 최대값에서 계속 진행됩니다.

## 모범 사례

### 구성 설계
- 다른 시스템과 충돌하지 않을 설명적인 접두사를 선택하십시오.
- 예상되는 볼륨에 적절한 자리수 패딩을 사용하십시오.
- 충돌을 피하기 위해 합리적인 시작 번호를 설정하십시오.
- 배포 전에 샘플 데이터로 구성을 테스트하십시오.

### 접두사 지침
- 접두사는 짧고 기억하기 쉽게 유지하십시오 (2-5자).
- 일관성을 위해 대문자를 사용하십시오.
- 가독성을 위해 구분 기호(하이픈, 밑줄)를 포함하십시오.
- URL이나 시스템에서 문제를 일으킬 수 있는 특수 문자는 피하십시오.

### 시퀀스 계획
- 기록 볼륨을 추정하여 적절한 자리수 패딩을 선택하십시오.
- 시작 번호를 설정할 때 향후 성장을 고려하십시오.
- 서로 다른 기록 유형에 대해 서로 다른 시퀀스 범위를 계획하십시오.
- 팀 참조를 위해 ID 스킴을 문서화하십시오.

## 일반 사용 사례

1. **지원 시스템**
   - 티켓 번호: `TICK-001`, `TICK-002`
   - 케이스 ID: `CASE-2024-001`
   - 지원 요청: `SUP-001`

2. **프로젝트 관리**
   - 작업 ID: `TASK-001`, `TASK-002`
   - 스프린트 항목: `SPRINT-001`
   - 납품 번호: `DEL-001`

3. **비즈니스 운영**
   - 주문 번호: `ORD-2024-001`
   - 송장 ID: `INV-001`
   - 구매 주문: `PO-001`

4. **품질 관리**
   - 버그 보고서: `BUG-001`
   - 테스트 케이스 ID: `TEST-001`
   - 리뷰 번호: `REV-001`

## 통합 기능

### 자동화와 함께
- 고유 ID가 할당될 때 작업 트리거
- 자동화 규칙에서 ID 패턴 사용
- 이메일 템플릿 및 알림에서 ID 참조

### 조회와 함께
- 다른 기록의 고유 ID 참조
- 고유 ID로 기록 찾기
- 관련 기록 식별자 표시

### 보고와 함께
- ID 패턴으로 그룹화 및 필터링
- ID 할당 추세 추적
- 시퀀스 사용 및 간격 모니터링

## 제한 사항

- **순차적만 가능**: ID는 시간 순서로 할당됩니다.
- **간격 없음**: 삭제된 기록은 시퀀스에 간격을 남깁니다.
- **재사용 없음**: 시퀀스 번호는 절대 재사용되지 않습니다.
- **프로젝트 범위**: 프로젝트 간 시퀀스를 공유할 수 없습니다.
- **형식 제약**: 제한된 형식 옵션.
- **대량 업데이트 없음**: 기존 시퀀스 ID를 대량 업데이트할 수 없습니다.
- **사용자 정의 논리 없음**: 사용자 정의 ID 생성 규칙을 구현할 수 없습니다.

## 관련 리소스

- [텍스트 필드](/api/custom-fields/text-single) - 수동 텍스트 식별자를 위한
- [숫자 필드](/api/custom-fields/number) - 숫자 시퀀스를 위한
- [사용자 정의 필드 개요](/api/custom-fields/2.list-custom-fields) - 일반 개념
- [자동화](/api/automations) - ID 기반 자동화 규칙을 위한