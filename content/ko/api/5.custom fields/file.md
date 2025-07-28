---
title: 파일 사용자 정의 필드
description: 문서, 이미지 및 기타 파일을 레코드에 첨부하기 위한 파일 필드를 생성합니다.
---

파일 사용자 정의 필드를 사용하면 레코드에 여러 파일을 첨부할 수 있습니다. 파일은 AWS S3에 안전하게 저장되며, 포괄적인 메타데이터 추적, 파일 유형 검증 및 적절한 접근 제어가 이루어집니다.

## 기본 예제

간단한 파일 필드를 생성합니다:

```graphql
mutation CreateFileField {
  createCustomField(input: {
    name: "Attachments"
    type: FILE
  }) {
    id
    name
    type
  }
}
```

## 고급 예제

설명이 포함된 파일 필드를 생성합니다:

```graphql
mutation CreateDetailedFileField {
  createCustomField(input: {
    name: "Project Documents"
    type: FILE
    description: "Upload project-related documents, images, and files"
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
| `name` | String! | ✅ 예 | 파일 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `FILE` 여야 합니다. |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 사용자 정의 필드는 사용자의 현재 프로젝트 컨텍스트에 따라 프로젝트와 자동으로 연결됩니다. `projectId` 매개변수가 필요하지 않습니다.

## 파일 업로드 프로세스

### 1단계: 파일 업로드

먼저, 파일을 업로드하여 파일 UID를 가져옵니다:

```graphql
mutation UploadFile {
  uploadFile(input: {
    file: $file  # File upload variable
    companyId: "company_123"
    projectId: "proj_123"
  }) {
    id
    uid
    name
    size
    type
    extension
    status
  }
}
```

### 2단계: 레코드에 파일 첨부

그런 다음 업로드된 파일을 레코드에 첨부합니다:

```graphql
mutation AttachFileToRecord {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "file_field_456"
    fileUid: "file_uid_from_upload"
  }) {
    id
    file {
      uid
      name
      size
      type
    }
  }
}
```

## 파일 첨부 관리

### 단일 파일 추가

```graphql
mutation AddFileToField {
  createTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  }) {
    id
    position
    file {
      uid
      name
      size
      type
      extension
    }
  }
}
```

### 파일 제거

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### 대량 파일 작업

customFieldOptionIds를 사용하여 여러 파일을 한 번에 업데이트합니다:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## 파일 업로드 입력 매개변수

### UploadFileInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `file` | Upload! | ✅ 예 | 업로드할 파일 |
| `companyId` | String! | ✅ 예 | 파일 저장을 위한 회사 ID |
| `projectId` | String | 아니요 | 프로젝트별 파일을 위한 프로젝트 ID |

### 파일 관리 입력 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 파일 사용자 정의 필드의 ID |
| `fileUid` | String! | ✅ 예 | 업로드된 파일의 고유 식별자 |

## 파일 저장 및 한계

### 파일 크기 한계

| 한계 유형 | 크기 |
|------------|------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### 지원되는 파일 유형

#### 이미지
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### 비디오
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### 오디오
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### 문서
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### 아카이브
- `zip`, `rar`, `7z`, `tar`, `gz`

#### 코드/텍스트
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### 저장 아키텍처

- **저장소**: 정리된 폴더 구조를 가진 AWS S3
- **경로 형식**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **보안**: 안전한 접근을 위한 서명된 URL
- **백업**: 자동 S3 중복성

## 응답 필드

### 파일 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 데이터베이스 ID |
| `uid` | String! | 고유 파일 식별자 |
| `name` | String! | 원래 파일 이름 |
| `size` | Float! | 바이트 단위의 파일 크기 |
| `type` | String! | MIME 유형 |
| `extension` | String! | 파일 확장자 |
| `status` | FileStatus | PENDING 또는 CONFIRMED (nullable) |
| `shared` | Boolean! | 파일이 공유되는지 여부 |
| `createdAt` | DateTime! | 업로드 타임스탬프 |

### TodoCustomFieldFile 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | ID! | 연결 레코드 ID |
| `uid` | String! | 고유 식별자 |
| `position` | Float! | 표시 순서 |
| `file` | File! | 관련 파일 객체 |
| `todoCustomField` | TodoCustomField! | 부모 사용자 정의 필드 |
| `createdAt` | DateTime! | 파일이 첨부된 시점 |

## 파일이 포함된 레코드 생성

레코드를 생성할 때 파일의 UID를 사용하여 파일을 첨부할 수 있습니다:

```graphql
mutation CreateRecordWithFiles {
  createTodo(input: {
    title: "Project deliverables"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "file_field_id"
      customFieldOptionIds: ["file_uid_1", "file_uid_2"]
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
      todoCustomFieldFiles {
        id
        position
        file {
          uid
          name
          size
          type
        }
      }
    }
  }
}
```

## 파일 검증 및 보안

### 업로드 검증

- **MIME 유형 검사**: 허용된 유형에 대해 검증합니다.
- **파일 확장자 검증**: `application/octet-stream`에 대한 대체 수단입니다.
- **크기 한계**: 업로드 시 enforced됩니다.
- **파일 이름 정리**: 특수 문자를 제거합니다.

### 접근 제어

- **업로드 권한**: 프로젝트/회사 멤버십이 필요합니다.
- **파일 연결**: ADMIN, OWNER, MEMBER, CLIENT 역할
- **파일 접근**: 프로젝트/회사 권한에서 상속됩니다.
- **안전한 URL**: 파일 접근을 위한 시간 제한 서명된 URL

## 필요한 권한

| 작업 | 필요한 권한 |
|--------|-------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## 오류 응답

### 파일이 너무 큽니다
```json
{
  "errors": [{
    "message": "File \"filename.pdf\": Size exceeds maximum limit of 256MB",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### 파일을 찾을 수 없습니다
```json
{
  "errors": [{
    "message": "File not found",
    "extensions": {
      "code": "FILE_NOT_FOUND"
    }
  }]
}
```

### 필드를 찾을 수 없습니다
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

### 파일 관리
- 레코드에 첨부하기 전에 파일을 업로드합니다.
- 설명적인 파일 이름을 사용합니다.
- 프로젝트/목적별로 파일을 정리합니다.
- 사용하지 않는 파일을 주기적으로 정리합니다.

### 성능
- 가능할 경우 파일을 배치로 업로드합니다.
- 콘텐츠 유형에 적합한 파일 형식을 사용합니다.
- 업로드 전에 큰 파일을 압축합니다.
- 파일 미리보기 요구 사항을 고려합니다.

### 보안
- 파일 내용을 검증하고 확장자만 검증하지 않습니다.
- 업로드된 파일에 대해 바이러스 검사를 사용합니다.
- 적절한 접근 제어를 구현합니다.
- 파일 업로드 패턴을 모니터링합니다.

## 일반적인 사용 사례

1. **문서 관리**
   - 프로젝트 사양
   - 계약 및 합의
   - 회의 노트 및 프레젠테이션
   - 기술 문서

2. **자산 관리**
   - 디자인 파일 및 목업
   - 브랜드 자산 및 로고
   - 마케팅 자료
   - 제품 이미지

3. **규정 준수 및 기록**
   - 법적 문서
   - 감사 기록
   - 인증서 및 라이센스
   - 재무 기록

4. **협업**
   - 공유 리소스
   - 버전 관리 문서
   - 피드백 및 주석
   - 참조 자료

## 통합 기능

### 자동화와 함께
- 파일이 추가/제거될 때 작업 트리거
- 유형 또는 메타데이터에 따라 파일 처리
- 파일 변경에 대한 알림 전송
- 조건에 따라 파일 아카이브

### 커버 이미지와 함께
- 파일 필드를 커버 이미지 소스로 사용
- 자동 이미지 처리 및 썸네일
- 파일 변경 시 동적 커버 업데이트

### 조회와 함께
- 다른 레코드에서 파일 참조
- 파일 수 및 크기 집계
- 파일 메타데이터로 레코드 찾기
- 파일 첨부에 대한 교차 참조

## 한계

- 파일당 최대 256MB
- S3 가용성에 의존
- 내장 파일 버전 관리 없음
- 자동 파일 변환 없음
- 제한된 파일 미리보기 기능
- 실시간 협업 편집 없음

## 관련 리소스

- [파일 업로드 API](/api/upload-files) - 파일 업로드 엔드포인트
- [사용자 정의 필드 개요](/api/custom-fields/list-custom-fields) - 일반 개념
- [자동화 API](/api/automations) - 파일 기반 자동화
- [AWS S3 문서](https://docs.aws.amazon.com/s3/) - 저장소 백엔드