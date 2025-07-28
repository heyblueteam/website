---
title: URL 사용자 정의 필드
description: 웹사이트 주소 및 링크를 저장하기 위한 URL 필드를 생성합니다.
---

URL 사용자 정의 필드는 기록에 웹사이트 주소 및 링크를 저장할 수 있게 해줍니다. 이는 프로젝트 웹사이트, 참조 링크, 문서 URL 또는 작업과 관련된 웹 기반 리소스를 추적하는 데 이상적입니다.

## 기본 예제

간단한 URL 필드를 생성합니다:

```graphql
mutation CreateUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Project Website"
      type: URL
    }
  ) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 URL 필드를 생성합니다:

```graphql
mutation CreateDetailedUrlField($projectId: String!) {
  createCustomField(
    projectId: $projectId
    input: {
      name: "Reference Link"
      type: URL
      description: "Link to external documentation or resources"
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
| `name` | String! | ✅ 예 | URL 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | 반드시 `URL` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

**참고:** `projectId`는 입력 객체의 일부가 아니라 변형에 대한 별도의 인수로 전달됩니다.

## URL 값 설정

기록에 URL 값을 설정하거나 업데이트하려면:

```graphql
mutation SetUrlValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "https://example.com/documentation"
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 기록의 ID |
| `customFieldId` | String! | ✅ 예 | URL 사용자 정의 필드의 ID |
| `text` | String! | ✅ 예 | 저장할 URL 주소 |

## URL 값으로 레코드 생성

URL 값으로 새 레코드를 생성할 때:

```graphql
mutation CreateRecordWithUrl {
  createTodo(input: {
    title: "Review documentation"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "url_field_id"
      value: "https://docs.example.com/api"
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
| `text` | String | 저장된 URL 주소 |
| `todo` | Todo! | 이 값이 속한 기록 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

## URL 검증

### 현재 구현
- **직접 API**: 현재 URL 형식 검증이 시행되지 않음
- **양식**: URL 검증이 계획되었으나 현재 활성화되지 않음
- **저장소**: URL 필드에 임의의 문자열 값을 저장할 수 있음

### 계획된 검증
향후 버전에는 다음이 포함될 예정입니다:
- HTTP/HTTPS 프로토콜 검증
- 유효한 URL 형식 확인
- 도메인 이름 검증
- 자동 프로토콜 접두사 추가

### 권장 URL 형식
현재 시행되지 않지만, 다음 표준 형식을 사용하십시오:

```
https://example.com
https://www.example.com
https://subdomain.example.com
https://example.com/path
https://example.com/path?param=value
http://localhost:3000
https://docs.example.com/api/v1
```

## 중요 참고 사항

### 저장 형식
- URL은 수정 없이 일반 텍스트로 저장됨
- 자동 프로토콜 추가 없음 (http://, https://)
- 입력한 대로 대소문자 구분 유지
- URL 인코딩/디코딩 수행되지 않음

### 직접 API 대 양식
- **양식**: 계획된 URL 검증 (현재 활성화되지 않음)
- **직접 API**: 검증 없음 - 모든 텍스트를 저장할 수 있음
- **권장 사항**: 저장하기 전에 애플리케이션에서 URL을 검증하십시오.

### URL 대 텍스트 필드
- **URL**: 웹 주소를 위한 의미론적으로 의도된 필드
- **TEXT_SINGLE**: 일반 단일 행 텍스트
- **백엔드**: 현재 동일한 저장 및 검증
- **프론트엔드**: 데이터 입력을 위한 다른 UI 구성 요소

## 필요한 권한

사용자 정의 필드 작업은 역할 기반 권한을 사용합니다:

| 작업 | 필요한 역할 |
|--------|-------------------|
| Create URL field | `OWNER` or `ADMIN` role in the project |
| Update URL field | `OWNER` or `ADMIN` role in the project |
| Set URL value | User must have edit permissions for the record |
| View URL value | User must have view permissions for the record |

**참고:** 권한은 특정 권한 상수가 아닌 프로젝트 내 사용자 역할에 따라 확인됩니다.

## 오류 응답

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

## 모범 사례

### URL 형식 표준
- 항상 프로토콜을 포함하십시오 (http:// 또는 https://)
- 보안을 위해 가능한 경우 HTTPS를 사용하십시오.
- 저장하기 전에 URL을 테스트하여 접근 가능성을 확인하십시오.
- 표시 목적으로 단축 URL 사용을 고려하십시오.

### 데이터 품질
- 저장하기 전에 애플리케이션에서 URL을 검증하십시오.
- 일반적인 오타(프로토콜 누락, 잘못된 도메인)를 확인하십시오.
- 조직 전체에서 URL 형식을 표준화하십시오.
- URL 접근성과 가용성을 고려하십시오.

### 보안 고려 사항
- 사용자 제공 URL에 주의하십시오.
- 특정 사이트로 제한할 경우 도메인을 검증하십시오.
- 악성 콘텐츠에 대한 URL 스캔을 고려하십시오.
- 민감한 데이터를 처리할 때 HTTPS URL을 사용하십시오.

## 필터링 및 검색

### 포함 검색
URL 필드는 부분 문자열 검색을 지원합니다:

```graphql
query SearchUrls {
  todos(
    customFieldFilters: [{
      customFieldId: "url_field_id"
      operation: CONTAINS
      value: "docs.example.com"
    }]
  ) {
    id
    title
    customFields {
      text
    }
  }
}
```

### 검색 기능
- 대소문자 구분 없는 부분 문자열 일치
- 부분 도메인 일치
- 경로 및 매개변수 검색
- 프로토콜별 필터링 없음

## 일반 사용 사례

1. **프로젝트 관리**
   - 프로젝트 웹사이트
   - 문서 링크
   - 저장소 URL
   - 데모 사이트

2. **콘텐츠 관리**
   - 참조 자료
   - 소스 링크
   - 미디어 리소스
   - 외부 기사

3. **고객 지원**
   - 고객 웹사이트
   - 지원 문서
   - 지식 기반 기사
   - 비디오 튜토리얼

4. **판매 및 마케팅**
   - 회사 웹사이트
   - 제품 페이지
   - 마케팅 자료
   - 소셜 미디어 프로필

## 통합 기능

### 조회와 함께
- 다른 기록의 참조 URL
- 도메인 또는 URL 패턴으로 기록 찾기
- 관련 웹 리소스 표시
- 여러 출처의 링크 집계

### 양식과 함께
- URL 전용 입력 구성 요소
- 적절한 URL 형식에 대한 계획된 검증
- 링크 미리보기 기능 (프론트엔드)
- 클릭 가능한 URL 표시

### 보고와 함께
- URL 사용 및 패턴 추적
- 끊어진 링크 또는 접근 불가능한 링크 모니터링
- 도메인 또는 프로토콜별로 분류
- 분석을 위한 URL 목록 내보내기

## 제한 사항

### 현재 제한 사항
- 활성 URL 형식 검증 없음
- 자동 프로토콜 추가 없음
- 링크 검증 또는 접근성 확인 없음
- URL 단축 또는 확장 없음
- 파비콘 또는 미리보기 생성 없음

### 자동화 제한
- 자동화 트리거 필드로 사용할 수 없음
- 자동화 필드 업데이트에 사용할 수 없음
- 자동화 조건에서 참조 가능
- 이메일 템플릿 및 웹훅에서 사용 가능

### 일반 제약
- 내장 링크 미리보기 기능 없음
- 자동 URL 단축 없음
- 클릭 추적 또는 분석 없음
- URL 만료 확인 없음
- 악성 URL 스캔 없음

## 향후 개선 사항

### 계획된 기능
- HTTP/HTTPS 프로토콜 검증
- 사용자 정의 정규 표현식 검증 패턴
- 자동 프로토콜 접두사 추가
- URL 접근성 확인

### 잠재적 개선 사항
- 링크 미리보기 생성
- 파비콘 표시
- URL 단축 통합
- 클릭 추적 기능
- 끊어진 링크 감지

## 관련 리소스

- [텍스트 필드](/api/custom-fields/text-single) - 비 URL 텍스트 데이터용
- [이메일 필드](/api/custom-fields/email) - 이메일 주소용
- [사용자 정의 필드 개요](/api/custom-fields/2.list-custom-fields) - 일반 개념

## 텍스트 필드에서 URL 필드로 마이그레이션

텍스트 필드에서 URL 필드로 마이그레이션하는 경우:

1. **동일한 이름과 구성으로 URL 필드 생성**
2. **기존 텍스트 값을 내보내어 유효한 URL인지 확인**
3. **레코드를 업데이트하여 새 URL 필드 사용**
4. **성공적인 마이그레이션 후 이전 텍스트 필드 삭제**
5. **애플리케이션을 업데이트하여 URL 전용 UI 구성 요소 사용**

### 마이그레이션 예제
```graphql
# Step 1: Create URL field
mutation CreateUrlField {
  createCustomField(input: {
    name: "Website Link"
    type: URL
    projectId: "proj_123"
  }) {
    id
  }
}

# Step 2: Update records (repeat for each record)
mutation MigrateToUrlField {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "new_url_field_id"
    text: "https://example.com"  # Value from old text field
  })
}
```