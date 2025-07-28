---
title: 참조 사용자 정의 필드
description: 교차 프로젝트 관계를 위한 다른 프로젝트의 레코드에 연결되는 참조 필드를 생성합니다.
---

참조 사용자 정의 필드를 사용하면 서로 다른 프로젝트의 레코드 간에 링크를 생성할 수 있어 교차 프로젝트 관계 및 데이터 공유가 가능합니다. 이는 조직의 프로젝트 구조 전반에 걸쳐 관련 작업을 연결하는 강력한 방법을 제공합니다.

## 기본 예제

간단한 참조 필드를 생성합니다:

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## 고급 예제

필터링 및 다중 선택이 가능한 참조 필드를 생성합니다:

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 참조 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `REFERENCE` 여야 합니다. |
| `referenceProjectId` | String | 아니요 | 참조할 프로젝트의 ID |
| `referenceMultiple` | Boolean | 아니요 | 여러 레코드 선택 허용 (기본값: false) |
| `referenceFilter` | TodoFilterInput | 아니요 | 참조된 레코드에 대한 필터 기준 |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 사용자 정의 필드는 사용자의 현재 프로젝트 컨텍스트에 따라 프로젝트와 자동으로 연결됩니다.

## 참조 구성

### 단일 참조 대 다중 참조

**단일 참조 (기본값):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- 사용자는 참조된 프로젝트에서 하나의 레코드를 선택할 수 있습니다.
- 단일 Todo 객체를 반환합니다.

**다중 참조:**
```graphql
{
  referenceMultiple: true
}
```
- 사용자는 참조된 프로젝트에서 여러 레코드를 선택할 수 있습니다.
- Todo 객체의 배열을 반환합니다.

### 참조 필터링

`referenceFilter`를 사용하여 선택할 수 있는 레코드를 제한합니다:

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## 참조 값 설정

### 단일 참조

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### 다중 참조

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 참조 사용자 정의 필드의 ID |
| `customFieldReferenceTodoIds` | [String!] | ✅ 예 | 참조된 레코드 ID의 배열 |

## 참조가 있는 레코드 생성

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 필드 값에 대한 고유 식별자 |
| `customField` | CustomField! | 참조 필드 정의 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

**참고**: 참조된 todos는 `customField.selectedTodos`를 통해 접근하며, TodoCustomField에서 직접 접근할 수 없습니다.

### 참조된 Todo 필드

각 참조된 Todo에는 다음이 포함됩니다:

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 참조된 레코드의 고유 식별자 |
| `title` | String! | 참조된 레코드의 제목 |
| `status` | TodoStatus! | 현재 상태 (ACTIVE, COMPLETED 등) |
| `description` | String | 참조된 레코드의 설명 |
| `dueDate` | DateTime | 설정된 경우 마감일 |
| `assignees` | [User!] | 할당된 사용자 |
| `tags` | [Tag!] | 관련 태그 |
| `project` | Project! | 참조된 레코드를 포함하는 프로젝트 |

## 참조 데이터 쿼리

### 기본 쿼리

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### 중첩 데이터가 있는 고급 쿼리

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## 필수 권한

| 작업 | 필수 권한 |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**중요**: 사용자는 연결된 레코드를 보기 위해 참조된 프로젝트에 대한 보기 권한이 있어야 합니다.

## 교차 프로젝트 접근

### 프로젝트 가시성

- 사용자는 접근할 수 있는 프로젝트의 레코드만 참조할 수 있습니다.
- 참조된 레코드는 원래 프로젝트의 권한을 준수합니다.
- 참조된 레코드에 대한 변경 사항은 실시간으로 나타납니다.
- 참조된 레코드를 삭제하면 참조 필드에서 제거됩니다.

### 권한 상속

- 참조 필드는 두 프로젝트의 권한을 상속받습니다.
- 사용자는 참조된 프로젝트에 대한 보기 접근이 필요합니다.
- 편집 권한은 현재 프로젝트의 규칙에 기반합니다.
- 참조된 데이터는 참조 필드의 맥락에서 읽기 전용입니다.

## 오류 응답

### 잘못된 참조 프로젝트

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### 참조된 레코드를 찾을 수 없음

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

### 권한 거부

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## 모범 사례

### 필드 설계

1. **명확한 명명** - 관계를 나타내는 설명적인 이름을 사용합니다.
2. **적절한 필터링** - 관련 레코드만 표시하도록 필터를 설정합니다.
3. **권한 고려** - 사용자가 참조된 프로젝트에 접근할 수 있는지 확인합니다.
4. **관계 문서화** - 연결에 대한 명확한 설명을 제공합니다.

### 성능 고려 사항

1. **참조 범위 제한** - 필터를 사용하여 선택할 수 있는 레코드 수를 줄입니다.
2. **깊은 중첩 피하기** - 복잡한 참조 체인을 만들지 않습니다.
3. **캐싱 고려** - 성능을 위해 참조된 데이터는 캐시됩니다.
4. **사용량 모니터링** - 프로젝트 간에 참조가 어떻게 사용되는지 추적합니다.

### 데이터 무결성

1. **삭제 처리** - 참조된 레코드가 삭제될 때를 계획합니다.
2. **권한 검증** - 사용자가 참조된 프로젝트에 접근할 수 있는지 확인합니다.
3. **종속성 업데이트** - 참조된 레코드를 변경할 때 영향을 고려합니다.
4. **감사 추적** - 준수를 위해 참조 관계를 추적합니다.

## 일반 사용 사례

### 프로젝트 의존성

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### 클라이언트 요구 사항

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### 리소스 할당

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### 품질 보증

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## 조회와의 통합

참조 필드는 [조회 필드](/api/custom-fields/lookup)와 함께 작동하여 참조된 레코드에서 데이터를 가져옵니다. 조회 필드는 참조 필드에서 선택된 레코드의 값을 추출할 수 있지만, 데이터 추출기일 뿐입니다 (SUM과 같은 집계 함수는 지원되지 않습니다).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## 제한 사항

- 참조된 프로젝트는 사용자에게 접근 가능해야 합니다.
- 참조된 프로젝트 권한에 대한 변경 사항은 참조 필드 접근에 영향을 미칩니다.
- 참조의 깊은 중첩은 성능에 영향을 미칠 수 있습니다.
- 순환 참조에 대한 내장 검증이 없습니다.
- 동일 프로젝트 참조를 방지하는 자동 제한이 없습니다.
- 참조 값을 설정할 때 필터 검증이 시행되지 않습니다.

## 관련 리소스

- [조회 필드](/api/custom-fields/lookup) - 참조된 레코드에서 데이터 추출
- [프로젝트 API](/api/projects) - 참조를 포함하는 프로젝트 관리
- [레코드 API](/api/records) - 참조가 있는 레코드 작업
- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 개념