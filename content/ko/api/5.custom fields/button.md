---
title: 버튼 사용자 정의 필드
description: 클릭 시 자동화를 트리거하는 인터랙티브 버튼 필드를 생성합니다.
---

버튼 사용자 정의 필드는 클릭 시 자동화를 트리거하는 인터랙티브 UI 요소를 제공합니다. 데이터를 저장하는 다른 사용자 정의 필드 유형과 달리, 버튼 필드는 구성된 워크플로를 실행하기 위한 액션 트리거 역할을 합니다.

## 기본 예제

자동화를 트리거하는 간단한 버튼 필드를 생성합니다:

```graphql
mutation CreateButtonField {
  createCustomField(input: {
    name: "Send Invoice"
    type: BUTTON
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

확인 요구 사항이 있는 버튼을 생성합니다:

```graphql
mutation CreateButtonWithConfirmation {
  createCustomField(input: {
    name: "Delete All Attachments"
    type: BUTTON
    projectId: "proj_123"
    buttonType: "hardConfirmation"
    buttonConfirmText: "DELETE"
    description: "Permanently removes all attachments from this task"
  }) {
    id
    name
    type
    buttonType
    buttonConfirmText
    description
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 버튼의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `BUTTON` 여야 합니다. |
| `projectId` | String! | ✅ 예 | 필드가 생성될 프로젝트 ID |
| `buttonType` | String | 아니요 | 확인 동작 (아래 버튼 유형 참조) |
| `buttonConfirmText` | String | 아니요 | 하드 확인을 위한 사용자가 입력해야 하는 텍스트 |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |
| `required` | Boolean | 아니요 | 필드가 필수인지 여부 (기본값은 false) |
| `isActive` | Boolean | 아니요 | 필드가 활성인지 여부 (기본값은 true) |

### 버튼 유형 필드

`buttonType` 필드는 UI 클라이언트가 확인 동작을 결정하는 데 사용할 수 있는 자유 형식 문자열입니다. 일반적인 값은 다음과 같습니다:

- `""` (비어 있음) - 확인 없음
- `"soft"` - 간단한 확인 대화 상자
- `"hard"` - 확인 텍스트 입력 요구

**참고**: 이는 UI 힌트일 뿐입니다. API는 특정 값을 검증하거나 강제하지 않습니다.

## 버튼 클릭 트리거

버튼 클릭을 트리거하고 관련된 자동화를 실행하려면:

```graphql
mutation ClickButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
  })
}
```

### 클릭 입력 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 버튼이 포함된 작업의 ID |
| `customFieldId` | String! | ✅ 예 | 버튼 사용자 정의 필드의 ID |

### 중요: API 동작

**API를 통한 모든 버튼 클릭은 즉시 실행됩니다** `buttonType` 또는 `buttonConfirmText` 설정과 관계없이. 이러한 필드는 UI 클라이언트가 확인 대화 상자를 구현하기 위해 저장되지만, API 자체는:

- 확인 텍스트를 검증하지 않습니다.
- 확인 요구 사항을 강제하지 않습니다.
- 호출 시 버튼 동작을 즉시 실행합니다.

확인은 순전히 클라이언트 측 UI 안전 기능입니다.

### 예제: 다양한 버튼 유형 클릭

```graphql
# Button with no confirmation
mutation ClickSimpleButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "simple_button_id"
  })
}

# Button with soft confirmation (API call is the same!)
mutation ClickSoftConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "soft_confirm_button_id"
  })
}

# Button with hard confirmation (API call is still the same!)
mutation ClickHardConfirmButton {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "hard_confirm_button_id"
  })
}
```

위의 세 가지 변형은 API를 통해 호출될 때 버튼 동작을 즉시 실행하며, 확인 요구 사항을 우회합니다.

## 응답 필드

### 사용자 정의 필드 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 사용자 정의 필드의 고유 식별자 |
| `name` | String! | 버튼의 표시 이름 |
| `type` | CustomFieldType! | 버튼 필드의 경우 항상 `BUTTON` |
| `buttonType` | String | 확인 동작 설정 |
| `buttonConfirmText` | String | 필수 확인 텍스트 (하드 확인을 사용하는 경우) |
| `description` | String | 사용자에 대한 도움말 텍스트 |
| `required` | Boolean! | 필수가 있는지 여부 |
| `isActive` | Boolean! | 현재 활성인지 여부 |
| `projectId` | String! | 이 필드가 속한 프로젝트의 ID |
| `createdAt` | DateTime! | 필드가 생성된 시간 |
| `updatedAt` | DateTime! | 필드가 마지막으로 수정된 시간 |

## 버튼 필드 작동 방식

### 자동화 통합

버튼 필드는 Blue의 자동화 시스템과 함께 작동하도록 설계되었습니다:

1. **버튼 필드 생성** 위의 변형을 사용하여
2. **자동화 구성** `CUSTOM_FIELD_BUTTON_CLICKED` 이벤트를 수신합니다.
3. **사용자가 UI에서 버튼 클릭**
4. **자동화가 구성된 작업 실행**

### 이벤트 흐름

버튼이 클릭될 때:

```
User Click → setTodoCustomField mutation → CUSTOM_FIELD_BUTTON_CLICKED event → Automation execution
```

### 데이터 저장 없음

중요: 버튼 필드는 어떤 값 데이터도 저장하지 않습니다. 순전히 액션 트리거 역할을 합니다. 각 클릭은:
- 이벤트를 생성합니다.
- 관련된 자동화를 트리거합니다.
- 작업 기록에 액션을 기록합니다.
- 어떤 필드 값도 수정하지 않습니다.

## 필수 권한

사용자는 버튼 필드를 생성하고 사용하기 위해 적절한 프로젝트 역할이 필요합니다:

| 작업 | 필수 역할 |
|--------|-------------------|
| Create button field | `OWNER` or `ADMIN` at project level |
| Update button field | `OWNER` or `ADMIN` at project level |
| Click button | `OWNER`, `ADMIN`, `MEMBER`, or `CLIENT` (based on field permissions) |
| Configure automations | `OWNER` or `ADMIN` at project level |

## 오류 응답

### 권한 거부
```json
{
  "errors": [{
    "message": "You don't have permission to edit this custom field",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### 사용자 정의 필드 찾을 수 없음
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

**참고**: API는 누락된 자동화나 확인 불일치에 대해 특정 오류를 반환하지 않습니다.

## 모범 사례

### 명명 규칙
- 행동 지향적인 이름 사용: "송장 보내기", "보고서 생성", "팀 알림"
- 버튼이 수행하는 작업에 대해 구체적으로 설명
- "버튼 1" 또는 "여기를 클릭"과 같은 일반적인 이름 피하기

### 확인 설정
- 안전하고 되돌릴 수 있는 작업을 위해 `buttonType` 비워 두기
- UI 클라이언트에 확인 동작을 제안하기 위해 `buttonType` 설정
- 사용자가 UI 확인에서 입력해야 하는 내용을 지정하기 위해 `buttonConfirmText` 사용
- 기억하세요: 이는 UI 힌트일 뿐입니다 - API 호출은 항상 즉시 실행됩니다.

### 자동화 설계
- 버튼 작업을 단일 워크플로에 집중
- 클릭 후 어떤 일이 발생했는지에 대한 명확한 피드백 제공
- 버튼의 목적을 설명하는 설명 텍스트 추가 고려

## 일반적인 사용 사례

1. **워크플로 전환**
   - "완료로 표시"
   - "승인 요청"
   - "작업 보관"

2. **외부 통합**
   - "CRM에 동기화"
   - "송장 생성"
   - "이메일 업데이트 전송"

3. **배치 작업**
   - "모든 하위 작업 업데이트"
   - "프로젝트에 복사"
   - "템플릿 적용"

4. **보고 작업**
   - "보고서 생성"
   - "데이터 내보내기"
   - "요약 생성"

## 제한 사항

- 버튼은 데이터 값을 저장하거나 표시할 수 없습니다.
- 각 버튼은 자동화만 트리거할 수 있으며, 직접 API 호출은 할 수 없습니다 (그러나 자동화는 외부 API 또는 Blue의 API를 호출하는 HTTP 요청 작업을 포함할 수 있습니다).
- 버튼 가시성은 조건부로 제어할 수 없습니다.
- 클릭당 최대 하나의 자동화 실행 (그러나 해당 자동화는 여러 작업을 트리거할 수 있습니다).

## 관련 리소스

- [자동화 API](/api/automations/index) - 버튼에 의해 트리거되는 작업 구성
- [사용자 정의 필드 개요](/custom-fields/list-custom-fields) - 일반 사용자 정의 필드 개념