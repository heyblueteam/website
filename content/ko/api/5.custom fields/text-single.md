---
title: 단일 행 텍스트 사용자 정의 필드
description: 이름, 제목 및 레이블과 같은 짧은 텍스트 값을 위한 단일 행 텍스트 필드를 생성합니다.
---

단일 행 텍스트 사용자 정의 필드는 단일 행 입력을 위해 설계된 짧은 텍스트 값을 저장할 수 있게 해줍니다. 이름, 제목, 레이블 또는 단일 행에 표시되어야 하는 모든 텍스트 데이터에 적합합니다.

## 기본 예제

간단한 단일 행 텍스트 필드를 생성합니다:

```graphql
mutation CreateTextSingleField {
  createCustomField(input: {
    name: "Client Name"
    type: TEXT_SINGLE
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 단일 행 텍스트 필드를 생성합니다:

```graphql
mutation CreateDetailedTextSingleField {
  createCustomField(input: {
    name: "Product SKU"
    type: TEXT_SINGLE
    description: "Unique product identifier code"
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
| `name` | String! | ✅ 예 | 텍스트 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `TEXT_SINGLE` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 프로젝트 컨텍스트는 인증 헤더에서 자동으로 결정됩니다. `projectId` 매개변수가 필요하지 않습니다.

## 텍스트 값 설정

레코드에서 단일 행 텍스트 값을 설정하거나 업데이트하려면:

```graphql
mutation SetTextSingleValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "ABC-123-XYZ"
  }) {
    id
    customField {
      value  # Returns { text: "ABC-123-XYZ" }
    }
  }
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 텍스트 사용자 정의 필드의 ID |
| `text` | String | 아니요 | 저장할 단일 행 텍스트 내용 |

## 텍스트 값으로 레코드 생성

단일 행 텍스트 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithTextSingle {
  createTodo(input: {
    title: "Process Order"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_single_field_id"
      value: "ORD-2024-001"
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Text is accessed here as { text: "ORD-2024-001" }
      }
    }
  }
}
```

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의(텍스트 값 포함) |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

**중요**: 텍스트 값은 TodoCustomField에서 직접 접근하는 것이 아니라 `customField.value.text` 필드를 통해 접근합니다.

## 텍스트 값 쿼리

텍스트 사용자 정의 필드가 있는 레코드를 쿼리할 때, 텍스트는 `customField.value.text` 경로를 통해 접근합니다:

```graphql
query GetRecordWithText {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For TEXT_SINGLE type, contains { text: "your text value" }
      }
    }
  }
}
```

응답에는 중첩된 구조로 텍스트가 포함됩니다:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Product SKU",
          "type": "TEXT_SINGLE",
          "value": {
            "text": "ABC-123-XYZ"
          }
        }
      }]
    }
  }
}
```

## 텍스트 검증

### 양식 검증
단일 행 텍스트 필드가 양식에서 사용될 때:
- 선행 및 후행 공백이 자동으로 제거됩니다.
- 필드가 필수로 표시된 경우 필수 검증이 적용됩니다.
- 특정 형식 검증은 적용되지 않습니다.

### 검증 규칙
- 줄 바꿈을 포함한 모든 문자열 내용을 허용합니다(권장하지 않음).
- 문자 길이 제한 없음(데이터베이스 제한까지).
- 유니코드 문자 및 특수 기호 지원.
- 줄 바꿈은 유지되지만 이 필드 유형에는 적합하지 않습니다.

### 일반 텍스트 예제
```
Product Name
SKU-123-ABC
Client Reference #2024-001
Version 1.2.3
Project Alpha
Status: Active
```

## 중요 참고 사항

### 저장 용량
- MySQL `MediumText` 유형으로 저장됩니다.
- 최대 16MB의 텍스트 내용을 지원합니다.
- 다중 행 텍스트 필드와 동일한 저장소.
- 국제 문자를 위한 UTF-8 인코딩.

### 직접 API vs 양식
- **양식**: 자동 공백 제거 및 필수 검증.
- **직접 API**: 텍스트는 제공된 대로 정확히 저장됩니다.
- **권장 사항**: 사용자 입력을 위해 양식을 사용하여 일관된 형식을 보장합니다.

### TEXT_SINGLE vs TEXT_MULTI
- **TEXT_SINGLE**: 단일 행 텍스트 입력, 짧은 값에 적합.
- **TEXT_MULTI**: 다중 행 텍스트 영역 입력, 긴 내용에 적합.
- **백엔드**: 두 가지 모두 동일한 저장소 및 검증 사용.
- **프론트엔드**: 데이터 입력을 위한 서로 다른 UI 구성 요소.
- **의도**: TEXT_SINGLE은 의미상 단일 행 값을 위해 설계되었습니다.

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` role at project level |
| Update text field | `OWNER` or `ADMIN` role at project level |
| Set text value | Standard record edit permissions |
| View text value | Standard record view permissions |

## 오류 응답

### 필수 필드 검증 (양식 전용)
```json
{
  "errors": [{
    "message": "This field is required",
    "extensions": {
      "code": "VALIDATION_ERROR"
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

## 모범 사례

### 콘텐츠 가이드라인
- 텍스트를 간결하고 단일 행에 적합하게 유지합니다.
- 의도된 단일 행 표시를 위해 줄 바꿈을 피합니다.
- 유사한 데이터 유형에 대해 일관된 형식을 사용합니다.
- UI 요구 사항에 따라 문자 제한을 고려합니다.

### 데이터 입력
- 사용자 안내를 위한 명확한 필드 설명을 제공합니다.
- 사용자 입력을 위해 양식을 사용하여 검증을 보장합니다.
- 필요한 경우 애플리케이션에서 콘텐츠 형식 검증을 수행합니다.
- 표준화된 값을 위해 드롭다운 사용을 고려합니다.

### 성능 고려 사항
- 단일 행 텍스트 필드는 가볍고 성능이 뛰어납니다.
- 자주 검색되는 필드에 대한 인덱싱을 고려합니다.
- UI에서 적절한 표시 너비를 사용합니다.
- 표시 목적을 위해 콘텐츠 길이를 모니터링합니다.

## 필터링 및 검색

### 포함 검색
단일 행 텍스트 필드는 부분 문자열 검색을 지원합니다:

```graphql
query SearchTextSingle {
  todos(
    customFieldFilters: [{
      customFieldId: "text_single_field_id"
      operation: CONTAINS
      value: "SKU"
    }]
  ) {
    id
    title
    customFields {
      customField {
        value  # Access text via value.text
      }
    }
  }
}
```

### 검색 기능
- 대소문자를 구분하지 않는 부분 문자열 일치.
- 부분 단어 일치 지원.
- 정확한 값 일치.
- 전체 텍스트 검색 또는 순위 없음.

## 일반 사용 사례

1. **식별자 및 코드**
   - 제품 SKU
   - 주문 번호
   - 참조 코드
   - 버전 번호

2. **이름 및 제목**
   - 클라이언트 이름
   - 프로젝트 제목
   - 제품 이름
   - 카테고리 레이블

3. **짧은 설명**
   - 간단한 요약
   - 상태 레이블
   - 우선 순위 표시기
   - 분류 태그

4. **외부 참조**
   - 티켓 번호
   - 송장 참조
   - 외부 시스템 ID
   - 문서 번호

## 통합 기능

### 조회와 함께
- 다른 레코드의 텍스트 데이터 참조.
- 텍스트 내용을 통해 레코드 찾기.
- 관련 텍스트 정보 표시.
- 여러 출처에서 텍스트 값 집계.

### 양식과 함께
- 자동 공백 제거.
- 필수 필드 검증.
- 단일 행 텍스트 입력 UI.
- 문자 제한 표시(구성된 경우).

### 가져오기/내보내기와 함께
- 직접 CSV 열 매핑.
- 자동 텍스트 값 할당.
- 대량 데이터 가져오기 지원.
- 스프레드시트 형식으로 내보내기.

## 제한 사항

### 자동화 제한
- 자동화 트리거 필드로 직접 사용 불가.
- 자동화 필드 업데이트에 사용 불가.
- 자동화 조건에서 참조 가능.
- 이메일 템플릿 및 웹훅에서 사용 가능.

### 일반 제한 사항
- 내장된 텍스트 형식 지정 또는 스타일링 없음.
- 필수 필드 외에 자동 검증 없음.
- 내장된 고유성 강제 없음.
- 매우 큰 텍스트에 대한 콘텐츠 압축 없음.
- 버전 관리 또는 변경 추적 없음.
- 제한된 검색 기능(전체 텍스트 검색 없음).

## 관련 리소스

- [다중 행 텍스트 필드](/api/custom-fields/text-multi) - 긴 텍스트 콘텐츠용
- [이메일 필드](/api/custom-fields/email) - 이메일 주소용
- [URL 필드](/api/custom-fields/url) - 웹사이트 주소용
- [고유 ID 필드](/api/custom-fields/unique-id) - 자동 생성 식별자용
- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 개념