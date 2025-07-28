---
title: 대시보드 생성
description: Blue에서 데이터 시각화 및 보고를 위한 새로운 대시보드를 생성합니다.
---

## 대시보드 생성

`createDashboard` 변형을 사용하면 회사나 프로젝트 내에서 새로운 대시보드를 생성할 수 있습니다. 대시보드는 팀이 지표를 추적하고, 진행 상황을 모니터링하며, 데이터 기반 결정을 내리는 데 도움을 주는 강력한 시각화 도구입니다.

### 기본 예제

```graphql
mutation CreateDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      title: "Sales Performance Dashboard"
    }
  ) {
    id
    title
    createdBy {
      id
      email
      firstName
      lastName
    }
    createdAt
  }
}
```

### 프로젝트 전용 대시보드

특정 프로젝트와 연관된 대시보드를 생성합니다:

```graphql
mutation CreateProjectDashboard {
  createDashboard(
    input: {
      companyId: "comp_abc123"
      projectId: "proj_xyz789"
      title: "Q4 Project Metrics"
    }
  ) {
    id
    title
    project {
      id
      name
    }
    createdBy {
      id
      email
    }
    dashboardUsers {
      id
      user {
        id
        email
      }
      role
    }
    createdAt
  }
}
```

## 입력 매개변수

### CreateDashboardInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `companyId` | String! | ✅ 예 | 대시보드가 생성될 회사의 ID |
| `title` | String! | ✅ 예 | 대시보드의 이름. 비어 있지 않은 문자열이어야 합니다. |
| `projectId` | String | 아니오 | 이 대시보드와 연관될 프로젝트의 선택적 ID |

## 응답 필드

변형은 완전한 `Dashboard` 객체를 반환합니다:

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 생성된 대시보드의 고유 식별자 |
| `title` | String! | 제공된 대시보드 제목 |
| `companyId` | String! | 이 대시보드가 속한 회사 |
| `projectId` | String | 연관된 프로젝트 ID (제공된 경우) |
| `project` | Project | 연관된 프로젝트 객체 (projectId가 제공된 경우) |
| `createdBy` | User! | 대시보드를 생성한 사용자 (당신) |
| `dashboardUsers` | [DashboardUser!]! | 접근 권한이 있는 사용자 목록 (초기에는 생성자만) |
| `createdAt` | DateTime! | 대시보드가 생성된 시간의 타임스탬프 |
| `updatedAt` | DateTime! | 마지막 수정 시간의 타임스탬프 (새 대시보드의 경우 createdAt과 동일) |

### DashboardUser 필드

대시보드가 생성되면 생성자가 자동으로 대시보드 사용자로 추가됩니다:

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 대시보드 사용자 관계의 고유 식별자 |
| `user` | User! | 대시보드에 접근할 수 있는 사용자 객체 |
| `role` | DashboardRole! | 사용자의 역할 (생성자는 전체 접근 권한을 가짐) |
| `dashboard` | Dashboard! | 대시보드에 대한 참조 |

## 필수 권한

지정된 회사에 속한 모든 인증된 사용자는 대시보드를 생성할 수 있습니다. 특별한 역할 요구 사항은 없습니다.

| 사용자 상태 | 대시보드 생성 가능 |
|-------------|-------------------|
| Company Member | ✅ 예 |
| 비회사의 구성원 | ❌ 아니오 |
| Unauthenticated | ❌ 아니오 |

## 오류 응답

### 유효하지 않은 회사
```json
{
  "errors": [{
    "message": "Company not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### 회사에 없는 사용자
```json
{
  "errors": [{
    "message": "You don't have access to this company",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

### 유효하지 않은 프로젝트
```json
{
  "errors": [{
    "message": "Project not found or doesn't belong to the specified company",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```

### 비어 있는 제목
```json
{
  "errors": [{
    "message": "Dashboard title cannot be empty",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## 중요 사항

- **자동 소유권**: 대시보드를 생성하는 사용자는 자동으로 전체 권한을 가진 소유자가 됩니다.
- **프로젝트 연관**: `projectId`를 제공하는 경우, 동일한 회사에 속해야 합니다.
- **초기 권한**: 초기에는 생성자만 접근할 수 있습니다. `editDashboard`를 사용하여 추가 사용자들을 초대할 수 있습니다.
- **제목 요구 사항**: 대시보드 제목은 비어 있지 않은 문자열이어야 합니다. 고유성 요구 사항은 없습니다.
- **회사 구성원 자격**: 대시보드를 생성하려면 해당 회사의 구성원이어야 합니다.

## 대시보드 생성 워크플로우

1. **이 변형을 사용하여 대시보드 생성**
2. **대시보드 빌더 UI를 사용하여 차트 및 위젯 구성**
3. **`editDashboard` 변형을 사용하여 팀원 추가** ( `dashboardUsers` )
4. **대시보드 인터페이스를 통해 필터 및 날짜 범위 설정**
5. **대시보드를 공유하거나 포함**하여 고유 ID 사용

## 사용 사례

1. **임원 대시보드**: 회사 지표에 대한 고수준 개요 생성
2. **프로젝트 추적**: 진행 상황을 모니터링하기 위한 프로젝트 전용 대시보드 구축
3. **팀 성과**: 팀 생산성과 성과 지표 추적
4. **클라이언트 보고**: 클라이언트 대상 보고서를 위한 대시보드 생성
5. **실시간 모니터링**: 실시간 운영 데이터를 위한 대시보드 설정

## 모범 사례

1. **명명 규칙**: 대시보드의 목적을 나타내는 명확하고 설명적인 제목 사용
2. **프로젝트 연관**: 프로젝트 전용 대시보드일 경우 프로젝트에 연결
3. **접근 관리**: 생성 직후 팀원을 추가하여 협업 촉진
4. **조직화**: 일관된 명명 패턴을 사용하여 대시보드 계층 구조 생성

## 관련 작업

- [대시보드 목록](/api/dashboards/) - 회사 또는 프로젝트의 모든 대시보드 검색
- [대시보드 편집](/api/dashboards/rename-dashboard) - 대시보드 이름 변경 또는 사용자 관리
- [대시보드 복사](/api/dashboards/copy-dashboard) - 기존 대시보드 복제
- [대시보드 삭제](/api/dashboards/delete-dashboard) - 대시보드 제거