---
title: 다중 행 텍스트 사용자 정의 필드
description: 설명, 메모 및 댓글과 같은 긴 콘텐츠를 위한 다중 행 텍스트 필드를 생성합니다.
---

다중 행 텍스트 사용자 정의 필드는 줄 바꿈 및 서식을 포함하여 긴 텍스트 콘텐츠를 저장할 수 있게 해줍니다. 설명, 메모, 댓글 또는 여러 줄이 필요한 모든 텍스트 데이터에 적합합니다.

## 기본 예제

간단한 다중 행 텍스트 필드를 생성합니다:

```graphql
mutation CreateTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Description"
      type: TEXT_MULTI
    }
  ) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 다중 행 텍스트 필드를 생성합니다:

```graphql
mutation CreateDetailedTextMultiField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Notes"
      type: TEXT_MULTI
      description: "Detailed notes and observations about the project"
    }
  ) {
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
| `type` | CustomFieldType! | ✅ 예 | 반드시 `TEXT_MULTI` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

**참고:** `projectId`는 입력 객체의 일부가 아니라 변형에 대한 별도의 인수로 전달됩니다. 또는 프로젝트 컨텍스트는 GraphQL 요청의 `X-Bloo-Project-ID` 헤더에서 결정할 수 있습니다.

## 텍스트 값 설정

레코드에서 다중 행 텍스트 값을 설정하거나 업데이트하려면:

```graphql
mutation SetTextMultiValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "This is a multi-line text value.\n\nIt can contain line breaks and longer content."
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 텍스트 사용자 정의 필드의 ID |
| `text` | String | 아니요 | 저장할 다중 행 텍스트 콘텐츠 |

## 텍스트 값으로 레코드 생성

다중 행 텍스트 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithTextMulti {
  createTodo(input: {
    title: "Project Planning"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "text_multi_field_id"
      value: "Project Overview:\n\n1. Research phase\n2. Design phase\n3. Implementation phase\n\nKey considerations:\n- Budget constraints\n- Timeline requirements\n- Resource allocation"
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
      text
    }
  }
}
```

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값에 대한 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `text` | String | 저장된 다중 행 텍스트 콘텐츠 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

## 텍스트 유효성 검사

### 양식 유효성 검사
다중 행 텍스트 필드가 양식에서 사용될 때:
- 앞뒤 공백이 자동으로 잘립니다.
- 필드가 필수로 표시된 경우 필수 유효성 검사가 적용됩니다.
- 특정 형식 유효성 검사는 적용되지 않습니다.

### 유효성 검사 규칙
- 줄 바꿈을 포함한 모든 문자열 콘텐츠를 허용합니다.
- 문자 길이 제한 없음 (데이터베이스 제한까지)
- 유니코드 문자 및 특수 기호 지원
- 줄 바꿈은 저장 시 보존됩니다.

### 유효한 텍스트 예제
```
Single line text

Multi-line text with
line breaks

Text with special characters:
- Bullets
- Numbers: 123
- Symbols: @#$%
- Unicode: 🚀 ✅ ⭐

Code snippets:
function example() {
  return "hello world";
}
```

## 중요한 참고 사항

### 저장 용량
- MySQL `MediumText` 유형으로 저장됩니다.
- 최대 16MB의 텍스트 콘텐츠를 지원합니다.
- 줄 바꿈 및 서식이 보존됩니다.
- 국제 문자를 위한 UTF-8 인코딩

### 직접 API 대 양식
- **양식**: 자동 공백 잘림 및 필수 유효성 검사
- **직접 API**: 텍스트는 제공된 대로 정확히 저장됩니다.
- **권장 사항**: 사용자 입력을 위해 양식을 사용하여 일관된 서식을 보장합니다.

### TEXT_MULTI 대 TEXT_SINGLE
- **TEXT_MULTI**: 다중 행 텍스트 영역 입력, 긴 콘텐츠에 적합
- **TEXT_SINGLE**: 단일 행 텍스트 입력, 짧은 값에 적합
- **백엔드**: 두 유형은 동일합니다 - 동일한 저장 필드, 유효성 검사 및 처리
- **프론트엔드**: 데이터 입력을 위한 서로 다른 UI 구성 요소 (텍스트 영역 대 입력 필드)
- **중요**: TEXT_MULTI와 TEXT_SINGLE의 구분은 순전히 UI 목적을 위해 존재합니다.

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create text field | `OWNER` or `ADMIN` project-level role |
| Update text field | `OWNER` or `ADMIN` project-level role |
| Set text value | Any role except `VIEW_ONLY` or `COMMENT_ONLY` |
| View text value | Any project-level role |

## 오류 응답

### 필수 필드 유효성 검사 (양식 전용)
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

### 필드 없음
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

### 콘텐츠 구성
- 구조화된 콘텐츠에 대해 일관된 서식을 사용합니다.
- 가독성을 위해 마크다운과 유사한 구문 사용을 고려합니다.
- 긴 콘텐츠를 논리적인 섹션으로 나눕니다.
- 가독성을 개선하기 위해 줄 바꿈을 사용합니다.

### 데이터 입력
- 사용자 안내를 위한 명확한 필드 설명을 제공합니다.
- 유효성을 보장하기 위해 사용자 입력에 양식을 사용합니다.
- 사용 사례에 따라 문자 제한을 고려합니다.
- 필요 시 애플리케이션에서 콘텐츠 형식을 검증합니다.

### 성능 고려 사항
- 매우 긴 텍스트 콘텐츠는 쿼리 성능에 영향을 미칠 수 있습니다.
- 큰 텍스트 필드를 표시하기 위해 페이지 매김을 고려합니다.
- 검색 기능을 위한 인덱스 고려 사항
- 큰 콘텐츠가 있는 필드의 저장 사용량을 모니터링합니다.

## 필터링 및 검색

### 포함 검색
다중 행 텍스트 필드는 사용자 정의 필드 필터를 통해 부분 문자열 검색을 지원합니다:

```graphql
query SearchTextMulti {
  todos(
    customFieldFilters: [{
      customFieldId: "text_multi_field_id"
      operation: CONTAINS
      value: "project"
    }]
  ) {
    id
    title
    customFields {
      customField {
        name
        type
      }
      text
    }
  }
}
```

### 검색 기능
- `CONTAINS` 연산자를 사용하여 텍스트 필드 내에서 부분 문자열 일치
- `NCONTAINS` 연산자를 사용한 대소문자 구분 없는 검색
- `IS` 연산자를 사용한 정확한 일치
- `NOT` 연산자를 사용한 부정 일치
- 모든 텍스트 줄을 가로질러 검색
- 부분 단어 일치 지원

## 일반 사용 사례

1. **프로젝트 관리**
   - 작업 설명
   - 프로젝트 요구 사항
   - 회의 메모
   - 상태 업데이트

2. **고객 지원**
   - 문제 설명
   - 해결 메모
   - 고객 피드백
   - 커뮤니케이션 로그

3. **콘텐츠 관리**
   - 기사 콘텐츠
   - 제품 설명
   - 사용자 댓글
   - 리뷰 세부 정보

4. **문서화**
   - 프로세스 설명
   - 지침
   - 가이드라인
   - 참조 자료

## 통합 기능

### 자동화와 함께
- 텍스트 콘텐츠가 변경될 때 작업 트리거
- 텍스트 콘텐츠에서 키워드 추출
- 요약 또는 알림 생성
- 외부 서비스로 텍스트 콘텐츠 처리

### 조회와 함께
- 다른 레코드의 텍스트 데이터 참조
- 여러 소스에서 텍스트 콘텐츠 집계
- 텍스트 콘텐츠로 레코드 찾기
- 관련 텍스트 정보 표시

### 양식과 함께
- 자동 공백 잘림
- 필수 필드 유효성 검사
- 다중 행 텍스트 영역 UI
- 문자 수 표시 (구성된 경우)

## 제한 사항

- 내장 텍스트 서식 또는 리치 텍스트 편집 없음
- 자동 링크 감지 또는 변환 없음
- 맞춤법 검사 또는 문법 유효성 검사 없음
- 내장 텍스트 분석 또는 처리 없음
- 버전 관리 또는 변경 추적 없음
- 제한된 검색 기능 (전체 텍스트 검색 없음)
- 매우 큰 텍스트에 대한 콘텐츠 압축 없음

## 관련 리소스

- [단일 행 텍스트 필드](/api/custom-fields/text-single) - 짧은 텍스트 값용
- [이메일 필드](/api/custom-fields/email) - 이메일 주소용
- [URL 필드](/api/custom-fields/url) - 웹사이트 주소용
- [사용자 정의 필드 개요](/api/custom-fields/2.list-custom-fields) - 일반 개념