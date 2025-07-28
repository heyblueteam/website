---
title: 국가 사용자 정의 필드
description: ISO 국가 코드 유효성 검사를 통한 국가 선택 필드 생성
---

국가 사용자 정의 필드는 레코드에 대한 국가 정보를 저장하고 관리할 수 있게 해줍니다. 이 필드는 국가 이름과 ISO Alpha-2 국가 코드를 모두 지원합니다.

**중요**: 국가 유효성 검사 및 변환 동작은 변형에 따라 크게 다릅니다:
- **createTodo**: 국가 이름을 자동으로 검증하고 ISO 코드로 변환합니다.
- **setTodoCustomField**: 검증 없이 모든 값을 허용합니다.

## 기본 예제

간단한 국가 필드를 생성합니다:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 국가 필드를 생성합니다:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 국가 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `COUNTRY` 여야 합니다. |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: `projectId`는 입력으로 전달되지 않지만 GraphQL 컨텍스트(일반적으로 요청 헤더 또는 인증에서)에서 결정됩니다.

## 국가 값 설정

국가 필드는 두 개의 데이터베이스 필드에 데이터를 저장합니다:
- **`countryCodes`**: ISO Alpha-2 국가 코드를 데이터베이스에 쉼표로 구분된 문자열로 저장합니다(API를 통해 배열로 반환됨).
- **`text`**: 표시 텍스트 또는 국가 이름을 문자열로 저장합니다.

### 매개변수 이해하기

`setTodoCustomField` 변형은 국가 필드에 대해 두 개의 선택적 매개변수를 허용합니다:

| 매개변수 | 유형 | 필수 | 설명 | 기능 |
|-----------|------|----------|-------------|--------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID | - |
| `customFieldId` | String! | ✅ 예 | 국가 사용자 정의 필드의 ID | - |
| `countryCodes` | [String!] | 아니오 | ISO Alpha-2 국가 코드의 배열 | Stored in the `countryCodes` field |
| `text` | String | 아니오 | 표시 텍스트 또는 국가 이름 | Stored in the `text` field |

**중요**: 
- `setTodoCustomField`에서는 두 매개변수가 선택적이며 독립적으로 저장됩니다.
- `createTodo`에서는 시스템이 입력에 따라 두 필드를 자동으로 설정합니다(독립적으로 제어할 수 없습니다).

### 옵션 1: 국가 코드만 사용

표시 텍스트 없이 검증된 ISO 코드를 저장합니다:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

결과: `countryCodes` = `["US"]`, `text` = `null`

### 옵션 2: 텍스트만 사용

검증된 코드 없이 표시 텍스트를 저장합니다:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

결과: `countryCodes` = `null`, `text` = `"United States"`

**참고**: `setTodoCustomField`를 사용할 때, 어떤 매개변수를 사용하든지 유효성 검사가 발생하지 않습니다. 값은 제공된 대로 정확하게 저장됩니다.

### 옵션 3: 둘 다 사용 (권장)

검증된 코드와 표시 텍스트를 모두 저장합니다:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

결과: `countryCodes` = `["US"]`, `text` = `"United States"`

### 여러 국가

배열을 사용하여 여러 국가를 저장합니다:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## 국가 값으로 레코드 생성

레코드를 생성할 때 `createTodo` 변형은 **자동으로 국가 값을 검증하고 변환합니다**. 이는 국가 유효성 검사를 수행하는 유일한 변형입니다:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      countryCodes
    }
  }
}
```

### 허용되는 입력 형식

| 입력 유형 | 예제 | 결과 |
|------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **지원되지 않음** - 단일 잘못된 값으로 처리됨 |
| Mixed format | `"United States, CA"` | **지원되지 않음** - 단일 잘못된 값으로 처리됨 |

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `text` | String | 표시 텍스트 (국가 이름) |
| `countryCodes` | [String!] | ISO Alpha-2 국가 코드의 배열 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

## 국가 표준

Blue는 국가 코드에 대해 **ISO 3166-1 Alpha-2** 표준을 사용합니다:

- 두 글자 국가 코드(예: US, GB, FR, DE)
- `i18n-iso-countries` 라이브러리를 사용한 유효성 검사는 **createTodo에서만 발생합니다.**
- 공식적으로 인정된 모든 국가를 지원합니다.

### 예제 국가 코드

| 국가 | ISO 코드 |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

ISO 3166-1 alpha-2 국가 코드의 공식 전체 목록은 [ISO 온라인 브라우징 플랫폼](https://www.iso.org/obp/ui/#search/code/)을 방문하세요.

## 유효성 검사

**유효성 검사는 `createTodo` 변형에서만 발생합니다**:

1. **유효한 ISO 코드**: 유효한 ISO Alpha-2 코드를 수용합니다.
2. **국가 이름**: 인식된 국가 이름을 자동으로 코드로 변환합니다.
3. **잘못된 입력**: 인식되지 않은 값에 대해 `CustomFieldValueParseError`를 발생시킵니다.

**참고**: `setTodoCustomField` 변형은 유효성 검사를 수행하지 않으며 모든 문자열 값을 수용합니다.

### 오류 예제

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## 통합 기능

### 조회 필드
국가 필드는 LOOKUP 사용자 정의 필드에 의해 참조될 수 있어 관련 레코드에서 국가 데이터를 가져올 수 있습니다.

### 자동화
자동화 조건에서 국가 값을 사용하세요:
- 특정 국가에 따라 작업 필터링
- 국가에 따라 알림 전송
- 지리적 지역에 따라 작업 라우팅

### 양식
양식의 국가 필드는 사용자 입력을 자동으로 검증하고 국가 이름을 코드로 변환합니다.

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## 오류 응답

### 잘못된 국가 값
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### 필드 유형 불일치
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## 모범 사례

### 입력 처리
- 자동 검증 및 변환을 위해 `createTodo` 사용
- 유효성 검사를 우회하므로 `setTodoCustomField`를 신중하게 사용
- `setTodoCustomField` 사용 전에 응용 프로그램에서 입력을 검증하는 것을 고려
- 명확성을 위해 UI에서 전체 국가 이름을 표시

### 데이터 품질
- 입력 지점에서 국가 입력을 검증
- 시스템 전반에 걸쳐 일관된 형식 사용
- 보고를 위해 지역 그룹화를 고려

### 여러 국가
- 여러 국가를 위해 `setTodoCustomField`에서 배열 지원 사용
- `createTodo`의 값 필드에서 여러 국가는 **지원되지 않음**
- 적절한 처리를 위해 `setTodoCustomField`에 국가 코드를 배열로 저장

## 일반 사용 사례

1. **고객 관리**
   - 고객 본사 위치
   - 배송지
   - 세금 관할권

2. **프로젝트 추적**
   - 프로젝트 위치
   - 팀원 위치
   - 시장 목표

3. **규정 준수 및 법률**
   - 규제 관할권
   - 데이터 거주 요건
   - 수출 통제

4. **판매 및 마케팅**
   - 영토 할당
   - 시장 세분화
   - 캠페인 타겟팅

## 제한 사항

- ISO 3166-1 Alpha-2 코드(2글자 코드)만 지원
- 국가 하위 구역(주/도)에 대한 기본 지원 없음
- 자동 국가 플래그 아이콘 없음(텍스트 기반만 지원)
- 역사적 국가 코드를 검증할 수 없음
- 지역 또는 대륙 그룹화에 대한 기본 지원 없음
- **유효성 검사는 `createTodo`에서만 작동하며, `setTodoCustomField`에서는 작동하지 않음**
- **`createTodo` 값 필드에서 여러 국가 지원되지 않음**
- **국가 코드는 쉼표로 구분된 문자열로 저장되며, 실제 배열이 아님**

## 관련 리소스

- [사용자 정의 필드 개요](/custom-fields/list-custom-fields) - 일반 사용자 정의 필드 개념
- [조회 필드](/api/custom-fields/lookup) - 다른 레코드에서 국가 데이터 참조
- [양식 API](/api/forms) - 사용자 정의 양식에 국가 필드 포함