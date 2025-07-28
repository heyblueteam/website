---
title: 사용자 정의 필드 조회
description: 참조된 레코드에서 데이터를 자동으로 가져오는 조회 필드 생성
---

사용자 정의 조회 필드는 [참조 필드](/api/custom-fields/reference)에서 참조된 레코드의 데이터를 자동으로 가져와, 수동 복사 없이 연결된 레코드의 정보를 표시합니다. 참조된 데이터가 변경될 때 자동으로 업데이트됩니다.

## 기본 예제

참조된 레코드에서 태그를 표시하는 조회 필드를 생성합니다:

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 고급 예제

참조된 레코드에서 사용자 정의 필드 값을 추출하는 조회 필드를 생성합니다:

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 조회 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `LOOKUP` 여야 함 |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ 예 | 조회 구성 |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

## 조회 구성

### CustomFieldLookupOptionInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `referenceId` | String! | ✅ 예 | 데이터를 가져올 참조 필드의 ID |
| `lookupId` | String | 아니오 | 조회할 특정 사용자 정의 필드의 ID (TODO_CUSTOM_FIELD 유형에 필요) |
| `lookupType` | CustomFieldLookupType! | ✅ 예 | 참조된 레코드에서 추출할 데이터 유형 |

## 조회 유형

### CustomFieldLookupType 값

| 유형 | 설명 | 반환 |
|------|-------------|---------|
| `TODO_DUE_DATE` | 참조된 할 일의 기한 | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | 참조된 할 일의 생성 날짜 | Array of creation timestamps |
| `TODO_UPDATED_AT` | 참조된 할 일의 마지막 업데이트 날짜 | Array of update timestamps |
| `TODO_TAG` | 참조된 할 일의 태그 | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | 참조된 할 일의 담당자 | Array of user objects |
| `TODO_DESCRIPTION` | 참조된 할 일의 설명 | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | 참조된 할 일의 할 일 목록 이름 | Array of list titles |
| `TODO_CUSTOM_FIELD` | 참조된 할 일의 사용자 정의 필드 값 | Array of values based on the field type |

## 응답 필드

### CustomField 응답 (조회 필드용)

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드의 고유 식별자 |
| `name` | String! | 조회 필드의 표시 이름 |
| `type` | CustomFieldType! | `LOOKUP`가 됩니다 |
| `customFieldLookupOption` | CustomFieldLookupOption | 조회 구성 및 결과 |
| `createdAt` | DateTime! | 필드가 생성된 시간 |
| `updatedAt` | DateTime! | 필드가 마지막으로 업데이트된 시간 |

### CustomFieldLookupOption 구조

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | 수행 중인 조회 유형 |
| `lookupResult` | JSON | 참조된 레코드에서 추출된 데이터 |
| `reference` | CustomField | 소스로 사용되는 참조 필드 |
| `lookup` | CustomField | 조회되는 특정 필드 (TODO_CUSTOM_FIELD용) |
| `parentCustomField` | CustomField | 상위 조회 필드 |
| `parentLookup` | CustomField | 체인의 상위 조회 (중첩 조회용) |

## 조회 작동 방식

1. **데이터 추출**: 조회는 참조 필드를 통해 연결된 모든 레코드에서 특정 데이터를 추출합니다.
2. **자동 업데이트**: 참조된 레코드가 변경되면 조회 값이 자동으로 업데이트됩니다.
3. **읽기 전용**: 조회 필드는 직접 편집할 수 없으며 항상 현재 참조된 데이터를 반영합니다.
4. **계산 없음**: 조회는 집계나 계산 없이 데이터를 있는 그대로 추출하고 표시합니다.

## TODO_CUSTOM_FIELD 조회

`TODO_CUSTOM_FIELD` 유형을 사용할 때, `lookupId` 매개변수를 사용하여 추출할 사용자 정의 필드를 지정해야 합니다:

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

이것은 모든 참조된 레코드에서 지정된 사용자 정의 필드의 값을 추출합니다.

## 조회 데이터 쿼리

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## 예제 조회 결과

### 태그 조회 결과
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### 담당자 조회 결과
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### 사용자 정의 필드 조회 결과
조회 결과는 조회되는 사용자 정의 필드 유형에 따라 다릅니다. 예를 들어, 통화 필드 조회는 다음과 같은 결과를 반환할 수 있습니다:
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**중요**: 사용자는 조회 결과를 보기 위해 현재 프로젝트와 참조된 프로젝트 모두에 대한 보기 권한이 있어야 합니다.

## 오류 응답

### 잘못된 참조 필드
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

### 순환 조회 감지됨
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### TODO_CUSTOM_FIELD에 대한 조회 ID 누락
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## 모범 사례

1. **명확한 명명**: 어떤 데이터가 조회되는지를 나타내는 설명적인 이름을 사용하세요.
2. **적절한 유형**: 데이터 요구 사항에 맞는 조회 유형을 선택하세요.
3. **성능**: 조회는 모든 참조된 레코드를 처리하므로 많은 링크가 있는 참조 필드에 유의하세요.
4. **권한**: 조회가 작동하려면 사용자가 참조된 프로젝트에 대한 액세스 권한이 있어야 합니다.

## 일반 사용 사례

### 프로젝트 간 가시성
수동 동기화 없이 관련 프로젝트의 태그, 담당자 또는 상태를 표시합니다.

### 의존성 추적
현재 작업이 의존하는 작업의 기한 또는 완료 상태를 표시합니다.

### 리소스 개요
리소스 계획을 위해 참조된 작업에 할당된 모든 팀원을 표시합니다.

### 상태 집계
관련 작업의 모든 고유 상태를 수집하여 프로젝트 건강 상태를 한눈에 확인합니다.

## 제한 사항

- 조회 필드는 읽기 전용이며 직접 편집할 수 없습니다.
- 집계 함수(SUM, COUNT, AVG)가 없습니다 - 조회는 데이터만 추출합니다.
- 필터링 옵션이 없습니다 - 모든 참조된 레코드가 포함됩니다.
- 무한 루프를 방지하기 위해 순환 조회 체인이 방지됩니다.
- 결과는 현재 데이터를 반영하며 자동으로 업데이트됩니다.

## 관련 리소스

- [참조 필드](/api/custom-fields/reference) - 조회 소스를 위한 레코드 링크 생성
- [사용자 정의 필드 값](/api/custom-fields/custom-field-values) - 편집 가능한 사용자 정의 필드에 값 설정
- [사용자 정의 필드 목록](/api/custom-fields/list-custom-fields) - 프로젝트의 모든 사용자 정의 필드 쿼리