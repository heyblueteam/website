---
title: 이메일 사용자 정의 필드
description: 이메일 주소를 저장하고 검증하기 위한 이메일 필드 생성
---

이메일 사용자 정의 필드는 내장 검증 기능을 통해 레코드에 이메일 주소를 저장할 수 있게 해줍니다. 연락처 정보, 담당자 이메일 또는 프로젝트의 이메일 관련 데이터를 추적하는 데 이상적입니다.

## 기본 예제

간단한 이메일 필드를 생성합니다:

```graphql
mutation CreateEmailField {
  createCustomField(input: {
    name: "Contact Email"
    type: EMAIL
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 이메일 필드를 생성합니다:

```graphql
mutation CreateDetailedEmailField {
  createCustomField(input: {
    name: "Client Email"
    type: EMAIL
    description: "Primary email address for client communications"
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
| `name` | String! | ✅ 예 | 이메일 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `EMAIL` 여야 합니다. |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

## 이메일 값 설정

레코드에서 이메일 값을 설정하거나 업데이트하려면:

```graphql
mutation SetEmailValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "john.doe@example.com"
  }) {
    id
    customField {
      value  # Returns { text: "john.doe@example.com" }
    }
  }
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 이메일 사용자 정의 필드의 ID |
| `text` | String | 아니오 | 저장할 이메일 주소 |

## 이메일 값으로 레코드 생성

이메일 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithEmail {
  createTodo(input: {
    title: "Follow up with client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "email_field_id"
      value: "client@company.com"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Email is accessed here as { text: "client@company.com" }
      }
    }
  }
}
```

## 응답 필드

### CustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 사용자 정의 필드의 고유 식별자 |
| `name` | String! | 이메일 필드의 표시 이름 |
| `type` | CustomFieldType! | 필드 유형 (EMAIL) |
| `description` | String | 필드에 대한 도움말 텍스트 |
| `value` | JSON | 이메일 값을 포함합니다 (아래 참조) |
| `createdAt` | DateTime! | 필드가 생성된 시간 |
| `updatedAt` | DateTime! | 필드가 마지막으로 수정된 시간 |

**중요**: 이메일 값은 응답에서 직접 접근하는 것이 아니라 `customField.value.text` 필드를 통해 접근합니다.

## 이메일 값 쿼리

이메일 사용자 정의 필드가 있는 레코드를 쿼리할 때, 이메일은 `customField.value.text` 경로를 통해 접근합니다:

```graphql
query GetRecordWithEmail {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For EMAIL type, contains { text: "email@example.com" }
      }
    }
  }
}
```

응답에는 중첩 구조로 이메일이 포함됩니다:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Contact Email",
          "type": "EMAIL",
          "value": {
            "text": "john.doe@example.com"
          }
        }
      }]
    }
  }
}
```

## 이메일 검증

### 양식 검증
이메일 필드가 양식에서 사용될 때, 자동으로 이메일 형식을 검증합니다:
- 표준 이메일 검증 규칙 사용
- 입력에서 공백을 제거
- 잘못된 이메일 형식 거부

### 검증 규칙
- `@` 기호를 포함해야 함
- 유효한 도메인 형식이어야 함
- 앞뒤 공백은 자동으로 제거됨
- 일반적인 이메일 형식이 허용됨

### 유효한 이메일 예시
```
user@example.com
john.doe@company.co.uk
test+tag@domain.org
first.last@sub.domain.com
```

### 잘못된 이메일 예시
```
plainaddress          # Missing @ symbol
@domain.com          # Missing local part
user@                # Missing domain
user@domain          # Missing TLD
user name@domain.com # Spaces not allowed
```

## 중요한 노트

### 직접 API 대 양식
- **양식**: 자동 이메일 검증이 적용됨
- **직접 API**: 검증 없음 - 어떤 텍스트도 저장 가능
- **권장 사항**: 사용자 입력을 위해 양식을 사용하여 검증을 보장

### 저장 형식
- 이메일 주소는 일반 텍스트로 저장됨
- 특별한 형식이나 파싱 없음
- 대소문자 구분: EMAIL 사용자 정의 필드는 대소문자를 구분하여 저장됨 (사용자 인증 이메일은 소문자로 정규화됨)
- 데이터베이스 제약을 초과하는 최대 길이 제한 없음 (16MB 제한)

## 필요한 권한

| 작업 | 필요한 권한 |
|--------|-------------------|
| Create email field | `OWNER` or `ADMIN` project-level role |
| Update email field | `OWNER` or `ADMIN` project-level role |
| Delete email field | `OWNER` or `ADMIN` project-level role |
| Set email value | Any role except `VIEW_ONLY` and `COMMENT_ONLY` |
| View email value | Any project role with field access |

## 오류 응답

### 잘못된 이메일 형식 (양식 전용)
```json
{
  "errors": [{
    "message": "ValidationError",
    "extensions": {
      "code": "BAD_USER_INPUT",
      "data": {
        "errors": [{
          "field": "email",
          "message": "Email format is invalid"
        }]
      }
    }
  }]
}
```

### 필드를 찾을 수 없음
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

## 모범 사례

### 데이터 입력
- 항상 애플리케이션에서 이메일 주소를 검증하십시오.
- 실제 이메일 주소에만 이메일 필드를 사용하십시오.
- 자동 검증을 위해 사용자 입력에 양식을 사용하는 것을 고려하십시오.

### 데이터 품질
- 저장하기 전에 공백을 제거하십시오.
- 대소문자 정규화를 고려하십시오 (일반적으로 소문자).
- 중요한 작업 전에 이메일 형식을 검증하십시오.

### 개인 정보 고려 사항
- 이메일 주소는 일반 텍스트로 저장됩니다.
- 데이터 개인 정보 규정을 고려하십시오 (GDPR, CCPA).
- 적절한 접근 제어를 구현하십시오.

## 일반 사용 사례

1. **연락처 관리**
   - 클라이언트 이메일 주소
   - 공급업체 연락처 정보
   - 팀원 이메일
   - 지원 연락처 세부정보

2. **프로젝트 관리**
   - 이해관계자 이메일
   - 승인 연락처 이메일
   - 알림 수신자
   - 외부 협력자 이메일

3. **고객 지원**
   - 고객 이메일 주소
   - 지원 티켓 연락처
   - 에스컬레이션 연락처
   - 피드백 이메일 주소

4. **영업 및 마케팅**
   - 리드 이메일 주소
   - 캠페인 연락처 목록
   - 파트너 연락처 정보
   - 추천 출처 이메일

## 통합 기능

### 자동화와 함께
- 이메일 필드가 업데이트될 때 작업 트리거
- 저장된 이메일 주소로 알림 전송
- 이메일 변경에 따라 후속 작업 생성

### 조회와 함께
- 다른 레코드의 이메일 데이터 참조
- 여러 출처에서 이메일 목록 집계
- 이메일 주소로 레코드 찾기

### 양식과 함께
- 자동 이메일 검증
- 이메일 형식 확인
- 공백 제거

## 제한 사항

- 형식 확인 외에 내장 이메일 검증 또는 검증 없음
- 클릭 가능한 이메일 링크와 같은 이메일 전용 UI 기능 없음
- 암호화 없이 일반 텍스트로 저장됨
- 이메일 작성 또는 전송 기능 없음
- 이메일 메타데이터 저장 없음 (표시 이름 등)
- 직접 API 호출은 검증을 우회함 (양식만 검증)

## 관련 리소스

- [텍스트 필드](/api/custom-fields/text-single) - 비이메일 텍스트 데이터용
- [URL 필드](/api/custom-fields/url) - 웹사이트 주소용
- [전화 필드](/api/custom-fields/phone) - 전화번호용
- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 개념