---
title: 다중 선택 사용자 정의 필드
description: 미리 정의된 목록에서 사용자가 여러 옵션을 선택할 수 있도록 다중 선택 필드를 생성합니다.
---

다중 선택 사용자 정의 필드는 사용자가 미리 정의된 목록에서 여러 옵션을 선택할 수 있도록 합니다. 이는 카테고리, 태그, 기술, 기능 또는 제어된 옵션 집합에서 여러 선택이 필요한 모든 시나리오에 적합합니다.

## 기본 예제

간단한 다중 선택 필드를 생성합니다:

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

다중 선택 필드를 생성한 후 옵션을 별도로 추가합니다:

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 다중 선택 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `SELECT_MULTI` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |
| `projectId` | String! | ✅ 예 | 이 필드의 프로젝트 ID |

### CreateCustomFieldOptionInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ 예 | 사용자 정의 필드의 ID |
| `title` | String! | ✅ 예 | 옵션의 표시 텍스트 |
| `color` | String | 아니요 | 옵션의 색상 (임의의 문자열) |
| `position` | Float | 아니요 | 옵션의 정렬 순서 |

## 기존 필드에 옵션 추가

기존 다중 선택 필드에 새 옵션을 추가합니다:

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## 다중 선택 값 설정

레코드에서 여러 선택된 옵션을 설정하려면:

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 다중 선택 사용자 정의 필드의 ID |
| `customFieldOptionIds` | [String!] | ✅ 예 | 선택할 옵션 ID의 배열 |

## 다중 선택 값으로 레코드 생성

다중 선택 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
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
| `selectedOptions` | [CustomFieldOption!] | 선택된 옵션의 배열 |
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
| `name` | String! | 다중 선택 필드의 표시 이름 |
| `type` | CustomFieldType! | 항상 `SELECT_MULTI` |
| `description` | String | 필드에 대한 도움말 텍스트 |
| `customFieldOptions` | [CustomFieldOption!] | 사용 가능한 모든 옵션 |

## 값 형식

### 입력 형식
- **API 매개변수**: 옵션 ID의 배열 (`["option1", "option2", "option3"]`)
- **문자열 형식**: 쉼표로 구분된 옵션 ID (`"option1,option2,option3"`)

### 출력 형식
- **GraphQL 응답**: CustomFieldOption 객체의 배열
- **활동 로그**: 쉼표로 구분된 옵션 제목
- **자동화 데이터**: 옵션 제목의 배열

## 옵션 관리

### 옵션 속성 업데이트
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
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

### 옵션 재정렬
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## 유효성 검사 규칙

### 옵션 유효성 검사
- 제공된 모든 옵션 ID는 존재해야 합니다.
- 옵션은 지정된 사용자 정의 필드에 속해야 합니다.
- SELECT_MULTI 필드만 여러 옵션을 선택할 수 있습니다.
- 빈 배열은 유효합니다 (선택 없음).

### 필드 유효성 검사
- 사용 가능하려면 정의된 옵션이 최소 하나 이상 있어야 합니다.
- 옵션 제목은 필드 내에서 고유해야 합니다.
- 색상 필드는 임의의 문자열 값을 허용합니다 (헥스 유효성 검사 없음).

## 필요한 권한

| 작업 | 필요한 권한 |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## 오류 응답

### 잘못된 옵션 ID
```json
{
  "errors": [{
    "message": "Custom field option not found",
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
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### 비다중 필드에서 여러 옵션
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 모범 사례

### 옵션 디자인
- 설명적이고 간결한 옵션 제목 사용
- 일관된 색상 코딩 체계 적용
- 옵션 목록을 관리 가능하게 유지 (일반적으로 3-20 옵션)
- 옵션을 논리적으로 정렬 (알파벳순, 빈도순 등)

### 데이터 관리
- 사용하지 않는 옵션을 주기적으로 검토하고 정리
- 프로젝트 전반에 걸쳐 일관된 명명 규칙 사용
- 필드를 생성할 때 옵션 재사용성을 고려
- 옵션 업데이트 및 마이그레이션 계획

### 사용자 경험
- 명확한 필드 설명 제공
- 시각적 구분을 개선하기 위해 색상 사용
- 관련 옵션을 함께 그룹화
- 일반적인 경우에 대한 기본 선택 고려

## 일반 사용 사례

1. **프로젝트 관리**
   - 작업 카테고리 및 태그
   - 우선순위 수준 및 유형
   - 팀원 할당
   - 상태 표시기

2. **콘텐츠 관리**
   - 기사 카테고리 및 주제
   - 콘텐츠 유형 및 형식
   - 게시 채널
   - 승인 워크플로우

3. **고객 지원**
   - 문제 카테고리 및 유형
   - 영향을 받는 제품 또는 서비스
   - 해결 방법
   - 고객 세그먼트

4. **제품 개발**
   - 기능 카테고리
   - 기술 요구 사항
   - 테스트 환경
   - 릴리스 채널

## 통합 기능

### 자동화와 함께
- 특정 옵션이 선택될 때 작업 트리거
- 선택된 카테고리에 따라 작업 라우팅
- 높은 우선순위 선택에 대한 알림 전송
- 옵션 조합에 따라 후속 작업 생성

### 조회와 함께
- 선택된 옵션으로 레코드 필터링
- 옵션 선택에 따른 데이터 집계
- 다른 레코드에서 옵션 데이터 참조
- 옵션 조합에 따라 보고서 생성

### 양식과 함께
- 다중 선택 입력 제어
- 옵션 유효성 검사 및 필터링
- 동적 옵션 로딩
- 조건부 필드 표시

## 활동 추적

다중 선택 필드 변경 사항은 자동으로 추적됩니다:
- 추가 및 제거된 옵션 표시
- 활동 로그에 옵션 제목 표시
- 모든 선택 변경에 대한 타임스탬프
- 수정에 대한 사용자 귀속

## 제한 사항

- 옵션의 최대 실용 한계는 UI 성능에 따라 다름
- 계층적 또는 중첩된 옵션 구조 없음
- 옵션은 필드를 사용하는 모든 레코드에서 공유됨
- 내장된 옵션 분석 또는 사용 추적 없음
- 색상 필드는 임의의 문자열을 허용 (헥스 유효성 검사 없음)
- 옵션별로 다른 권한을 설정할 수 없음
- 옵션은 필드 생성과 함께 인라인으로 생성할 수 없음
- 전용 재정렬 변형 없음 (position과 함께 editCustomFieldOption 사용)

## 관련 리소스

- [단일 선택 필드](/api/custom-fields/select-single) - 단일 선택을 위한
- [체크박스 필드](/api/custom-fields/checkbox) - 간단한 불리언 선택을 위한
- [텍스트 필드](/api/custom-fields/text-single) - 자유 형식 텍스트 입력을 위한
- [사용자 정의 필드 개요](/api/custom-fields/2.list-custom-fields) - 일반 개념