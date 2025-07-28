---
title: 체크박스 사용자 정의 필드
description: 예/아니오 또는 참/거짓 데이터에 대한 부울 체크박스 필드 생성
---

체크박스 사용자 정의 필드는 작업에 대한 간단한 부울(참/거짓) 입력을 제공합니다. 이들은 이진 선택, 상태 표시기 또는 무언가가 완료되었는지 추적하는 데 적합합니다.

## 기본 예제

간단한 체크박스 필드를 생성합니다:

```graphql
mutation CreateCheckboxField {
  createCustomField(input: {
    name: "Reviewed"
    type: CHECKBOX
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명 및 유효성 검사가 포함된 체크박스 필드를 생성합니다:

```graphql
mutation CreateDetailedCheckbox {
  createCustomField(input: {
    name: "Customer Approved"
    type: CHECKBOX
    description: "Check this box when the customer has approved the work"
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
| `name` | String! | ✅ 예 | 체크박스의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `CHECKBOX` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

## 체크박스 값 설정

작업에서 체크박스 값을 설정하거나 업데이트하려면:

```graphql
mutation CheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: true
  })
}
```

체크박스를 선택 해제하려면:

```graphql
mutation UncheckTheBox {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    checked: false
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 작업의 ID |
| `customFieldId` | String! | ✅ 예 | 체크박스 사용자 정의 필드의 ID |
| `checked` | Boolean | 아니요 | 선택하려면 true, 선택 해제하려면 false |

## 체크박스 값으로 작업 생성

체크박스 값으로 새 작업을 생성할 때:

```graphql
mutation CreateTaskWithCheckbox {
  createTodo(input: {
    title: "Review contract"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "checkbox_field_id"
      value: "true"  # Pass as string
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
      checked
    }
  }
}
```

### 허용된 문자열 값

작업을 생성할 때 체크박스 값은 문자열로 전달되어야 합니다:

| 문자열 값 | 결과 |
|--------------|---------|
| `"true"` | ✅ 선택됨 (대소문자 구분) |
| `"1"` | ✅ 선택됨 |
| `"checked"` | ✅ 선택됨 (대소문자 구분) |
| Any other value | ❌ 선택 해제됨 |

**참고**: 작업 생성 중 문자열 비교는 대소문자를 구분합니다. 값은 반드시 `"true"`, `"1"` 또는 `"checked"`와 정확히 일치해야 선택된 상태가 됩니다.

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 필드 값의 고유 식별자 |
| `uid` | String! | 대체 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `checked` | Boolean | 체크박스 상태 (참/거짓/널) |
| `todo` | Todo! | 이 값이 속한 작업 |
| `createdAt` | DateTime! | 값이 생성된 시점 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시점 |

## 자동화 통합

체크박스 필드는 상태 변경에 따라 다양한 자동화 이벤트를 트리거합니다:

| 작업 | 트리거된 이벤트 | 설명 |
|--------|----------------|-------------|
| Check (false → true) | `CUSTOM_FIELD_ADDED` | 체크박스가 선택되었을 때 트리거됨 |
| Uncheck (true → false) | `CUSTOM_FIELD_REMOVED` | 체크박스가 선택 해제되었을 때 트리거됨 |

이렇게 하면 체크박스 상태 변경에 응답하는 자동화를 생성할 수 있습니다. 예를 들어:
- 항목이 승인될 때 알림 전송
- 검토 체크박스가 선택되었을 때 작업 이동
- 체크박스 상태에 따라 관련 필드 업데이트

## 데이터 가져오기/내보내기

### 체크박스 값 가져오기

CSV 또는 기타 형식을 통해 데이터를 가져올 때:
- `"true"`, `"yes"` → 선택됨 (대소문자 구분 안 함)
- 기타 모든 값 (예: `"false"`, `"no"`, `"0"`, 빈 값) → 선택 해제됨

### 체크박스 값 내보내기

데이터를 내보낼 때:
- 선택된 체크박스는 `"X"`로 내보냄
- 선택 해제된 체크박스는 빈 문자열 `""`로 내보냄

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create checkbox field | `OWNER` or `ADMIN` role at project level |
| Update checkbox field | `OWNER` or `ADMIN` role at project level |
| Set checkbox value | Standard task edit permissions (excluding VIEW_ONLY and COMMENT_ONLY roles) |
| View checkbox value | Standard task view permissions (authenticated users in company/project) |

## 오류 응답

### 잘못된 값 유형
```json
{
  "errors": [{
    "message": "Invalid value type for checkbox field",
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

## 모범 사례

### 명명 규칙
- 명확하고 행동 지향적인 이름 사용: "승인됨", "검토됨", "완료됨"
- 사용자에게 혼란을 주는 부정적인 이름 피하기: "비활성"보다는 "활성" 선호
- 체크박스가 나타내는 내용을 구체적으로 설명

### 체크박스를 사용할 때
- **이진 선택**: 예/아니오, 참/거짓, 완료/미완료
- **상태 표시기**: 승인됨, 검토됨, 게시됨
- **기능 플래그**: 우선 지원 있음, 서명 필요
- **간단한 추적**: 이메일 발송, 청구서 지불, 항목 배송

### 체크박스를 사용하지 말아야 할 때
- 두 가지 이상의 옵션이 필요할 때 (대신 SELECT_SINGLE 사용)
- 숫자 또는 텍스트 데이터의 경우 (NUMBER 또는 TEXT 필드 사용)
- 누가 체크했는지 또는 언제 체크했는지 추적해야 할 때 (감사 로그 사용)

## 일반 사용 사례

1. **승인 워크플로**
   - "관리자 승인"
   - "클라이언트 서명"
   - "법률 검토 완료"

2. **작업 관리**
   - "차단됨"
   - "검토 준비 완료"
   - "높은 우선 순위"

3. **품질 관리**
   - "QA 통과"
   - "문서 완료"
   - "테스트 작성됨"

4. **관리 플래그**
   - "청구서 발송"
   - "계약 서명"
   - "후속 조치 필요"

## 제한 사항

- 체크박스 필드는 참/거짓 값만 저장할 수 있습니다 (초기 설정 후 삼중 상태 또는 널 불가)
- 기본값 구성 불가 (설정될 때까지 항상 널로 시작)
- 누가 체크했는지 또는 언제 체크했는지와 같은 추가 메타데이터 저장 불가
- 다른 필드 값에 따라 조건부 가시성 불가

## 관련 리소스

- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 사용자 정의 필드 개념
- [자동화 API](/api/automations) - 체크박스 변경으로 트리거되는 자동화 생성