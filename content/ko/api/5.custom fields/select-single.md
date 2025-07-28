---
title: 단일 선택 사용자 정의 필드
description: 미리 정의된 목록에서 사용자가 하나의 옵션을 선택할 수 있도록 단일 선택 필드를 생성합니다.
---

단일 선택 사용자 정의 필드는 사용자가 미리 정의된 목록에서 정확히 하나의 옵션을 선택할 수 있도록 합니다. 이는 상태 필드, 카테고리, 우선 순위 또는 제어된 옵션 집합에서 하나의 선택만 이루어져야 하는 모든 시나리오에 이상적입니다.

## 기본 예제

간단한 단일 선택 필드를 생성합니다:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

미리 정의된 옵션으로 단일 선택 필드를 생성합니다:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 단일 선택 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `SELECT_SINGLE` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | 아니요 | 필드의 초기 옵션 |

### CreateCustomFieldOptionInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `title` | String! | ✅ 예 | 옵션의 표시 텍스트 |
| `color` | String | 아니요 | 옵션의 헥스 색상 코드 |

## 기존 필드에 옵션 추가

기존 단일 선택 필드에 새 옵션을 추가합니다:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## 단일 선택 값 설정

레코드에서 선택된 옵션을 설정합니다:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 단일 선택 사용자 정의 필드의 ID |
| `customFieldOptionId` | String | 아니요 | 선택할 옵션의 ID (단일 선택에 선호됨) |
| `customFieldOptionIds` | [String!] | 아니요 | 옵션 ID의 배열 (단일 선택에 대해 첫 번째 요소 사용) |

## 단일 선택 값 쿼리

레코드의 단일 선택 값을 쿼리합니다:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

`value` 필드는 선택된 옵션의 세부 정보를 포함하는 JSON 객체를 반환합니다.

## 단일 선택 값으로 레코드 생성

단일 선택 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
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
| `value` | JSON | ID, 제목, 색상, 위치가 포함된 선택된 옵션 객체 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

### CustomFieldOption 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 옵션의 고유 식별자 |
| `title` | String! | 옵션의 표시 텍스트 |
| `color` | String | 시각적 표현을 위한 헥스 색상 코드 |
| `position` | Float | 옵션의 정렬 순서 |
| `customField` | CustomField! | 이 옵션이 속한 사용자 정의 필드 |

### CustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드의 고유 식별자 |
| `name` | String! | 단일 선택 필드의 표시 이름 |
| `type` | CustomFieldType! | 항상 `SELECT_SINGLE` |
| `description` | String | 필드에 대한 도움말 텍스트 |
| `customFieldOptions` | [CustomFieldOption!] | 사용 가능한 모든 옵션 |

## 값 형식

### 입력 형식
- **API 매개변수**: 단일 옵션 ID에 `customFieldOptionId` 사용
- **대안**: `customFieldOptionIds` 배열 사용 (첫 번째 요소 사용)
- **선택 지우기**: 두 필드를 생략하거나 빈 값을 전달

### 출력 형식
- **GraphQL 응답**: {id, title, color, position}을 포함하는 `value` 필드의 JSON 객체
- **활동 로그**: 문자열로 된 옵션 제목
- **자동화 데이터**: 문자열로 된 옵션 제목

## 선택 동작

### 독점 선택
- 새 옵션을 설정하면 이전 선택이 자동으로 제거됩니다.
- 한 번에 하나의 옵션만 선택할 수 있습니다.
- `null` 또는 빈 값을 설정하면 선택이 지워집니다.

### 폴백 논리
- `customFieldOptionIds` 배열이 제공되면 첫 번째 옵션만 사용됩니다.
- 이는 다중 선택 입력 형식과의 호환성을 보장합니다.
- 빈 배열이나 null 값은 선택을 지웁니다.

## 옵션 관리

### 옵션 속성 업데이트
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
  }) {
    id
    title
    color
  }
}
```

### 옵션 삭제
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

**참고**: 옵션을 삭제하면 선택된 모든 레코드에서 지워집니다.

### 옵션 재정렬
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## 검증 규칙

### 옵션 검증
- 제공된 옵션 ID는 존재해야 합니다.
- 옵션은 지정된 사용자 정의 필드에 속해야 합니다.
- 한 번에 하나의 옵션만 선택할 수 있습니다 (자동으로 시행됨).
- Null/빈 값은 유효합니다 (선택 없음).

### 필드 검증
- 사용 가능하려면 정의된 옵션이 최소 하나 이상 있어야 합니다.
- 옵션 제목은 필드 내에서 고유해야 합니다.
- 색상 코드는 유효한 헥스 형식이어야 합니다 (제공된 경우).

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## 오류 응답

### 잘못된 옵션 ID
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### 옵션이 필드에 속하지 않음
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

### 값을 구문 분석할 수 없음
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 모범 사례

### 옵션 디자인
- 명확하고 설명적인 옵션 제목 사용
- 의미 있는 색상 코딩 적용
- 옵션 목록을 집중적이고 관련성 있게 유지
- 옵션을 논리적으로 정렬 (우선 순위, 빈도 등)

### 상태 필드 패턴
- 프로젝트 전반에 걸쳐 일관된 상태 워크플로우 사용
- 옵션의 자연스러운 진행 고려
- 명확한 최종 상태 포함 (완료, 취소 등)
- 옵션 의미를 반영하는 색상 사용

### 데이터 관리
- 주기적으로 사용되지 않는 옵션 검토 및 정리
- 일관된 명명 규칙 사용
- 옵션 삭제가 기존 레코드에 미치는 영향 고려
- 옵션 업데이트 및 마이그레이션 계획

## 일반 사용 사례

1. **상태 및 워크플로우**
   - 작업 상태 (할 일, 진행 중, 완료)
   - 승인 상태 (대기 중, 승인됨, 거부됨)
   - 프로젝트 단계 (계획, 개발, 테스트, 출시)
   - 문제 해결 상태

2. **분류 및 카테고리화**
   - 우선 순위 수준 (낮음, 중간, 높음, 중요)
   - 작업 유형 (버그, 기능, 개선, 문서화)
   - 프로젝트 카테고리 (내부, 클라이언트, 연구)
   - 부서 할당

3. **품질 및 평가**
   - 검토 상태 (시작되지 않음, 검토 중, 승인됨)
   - 품질 등급 (불량, 보통, 좋음, 우수)
   - 위험 수준 (낮음, 중간, 높음)
   - 신뢰 수준

4. **할당 및 소유권**
   - 팀 할당
   - 부서 소유권
   - 역할 기반 할당
   - 지역 할당

## 통합 기능

### 자동화와 함께
- 특정 옵션이 선택될 때 작업 트리거
- 선택된 카테고리에 따라 작업 라우팅
- 상태 변경에 대한 알림 전송
- 선택에 따라 조건부 워크플로우 생성

### 조회와 함께
- 선택된 옵션으로 레코드 필터링
- 다른 레코드의 옵션 데이터 참조
- 옵션 선택에 따라 보고서 생성
- 선택된 값으로 레코드 그룹화

### 양식과 함께
- 드롭다운 입력 제어
- 라디오 버튼 인터페이스
- 옵션 검증 및 필터링
- 선택에 따라 조건부 필드 표시

## 활동 추적

단일 선택 필드 변경 사항은 자동으로 추적됩니다:
- 이전 및 새로운 옵션 선택 표시
- 활동 로그에 옵션 제목 표시
- 모든 선택 변경에 대한 타임스탬프
- 수정에 대한 사용자 귀속

## 다중 선택과의 차이점

| 기능 | 단일 선택 | 다중 선택 |
|---------|---------------|--------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## 제한 사항

- 한 번에 하나의 옵션만 선택할 수 있습니다.
- 계층적 또는 중첩된 옵션 구조가 없습니다.
- 옵션은 필드를 사용하는 모든 레코드에서 공유됩니다.
- 내장된 옵션 분석 또는 사용 추적이 없습니다.
- 색상 코드는 표시용만 있으며 기능적 영향을 미치지 않습니다.
- 옵션별로 다른 권한을 설정할 수 없습니다.

## 관련 리소스

- [다중 선택 필드](/api/custom-fields/select-multi) - 여러 선택을 위한
- [체크박스 필드](/api/custom-fields/checkbox) - 간단한 불리언 선택을 위한
- [텍스트 필드](/api/custom-fields/text-single) - 자유 형식 텍스트 입력을 위한
- [사용자 정의 필드 개요](/api/custom-fields/1.index) - 일반 개념