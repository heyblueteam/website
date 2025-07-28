---
title: 檔案自訂欄位
description: 建立檔案欄位以將文件、圖片和其他檔案附加到記錄上
---

檔案自訂欄位允許您將多個檔案附加到記錄上。檔案安全地儲存在 AWS S3 中，並具有全面的元資料追蹤、檔案類型驗證和適當的存取控制。

## 基本範例

建立一個簡單的檔案欄位：

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

## 進階範例

建立一個帶有描述的檔案欄位：

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

## 輸入參數

### CreateCustomFieldInput

| 參數 | 類型 | 必填 | 描述 |
|------|------|------|------|
| `name` | String! | ✅ 是 | 檔案欄位的顯示名稱 |
| `type` | CustomFieldType! | ✅ 是 | 必須是 `FILE` |
| `description` | String | 否 | 顯示給使用者的幫助文字 |

**注意**：自訂欄位會根據使用者當前的專案上下文自動與專案關聯。無需 `projectId` 參數。

## 檔案上傳過程

### 步驟 1：上傳檔案

首先，上傳檔案以獲取檔案 UID：

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

### 步驟 2：將檔案附加到記錄

然後將上傳的檔案附加到記錄上：

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

## 管理檔案附件

### 添加單個檔案

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

### 移除檔案

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### 批量檔案操作

使用 customFieldOptionIds 同時更新多個檔案：

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## 檔案上傳輸入參數

### UploadFileInput

| 參數 | 類型 | 必填 | 描述 |
|------|------|------|------|
| `file` | Upload! | ✅ 是 | 要上傳的檔案 |
| `companyId` | String! | ✅ 是 | 用於檔案儲存的公司 ID |
| `projectId` | String | 否 | 用於專案特定檔案的專案 ID |

### 檔案管理輸入參數

| 參數 | 類型 | 必填 | 描述 |
|------|------|------|------|
| `todoId` | String! | ✅ 是 | 記錄的 ID |
| `customFieldId` | String! | ✅ 是 | 檔案自訂欄位的 ID |
| `fileUid` | String! | ✅ 是 | 上傳檔案的唯一識別碼 |

## 檔案儲存和限制

### 檔案大小限制

| 限制類型 | 大小 |
|----------|------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### 支援的檔案類型

#### 圖片
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### 影片
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### 音訊
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### 文件
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### 壓縮檔
- `zip`, `rar`, `7z`, `tar`, `gz`

#### 代碼/文本
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### 儲存架構

- **儲存**：AWS S3，具有組織的資料夾結構
- **路徑格式**： `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **安全性**：簽名 URL 以安全訪問
- **備份**：自動 S3 冗餘

## 回應欄位

### 檔案回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | ID! | 資料庫 ID |
| `uid` | String! | 唯一檔案識別碼 |
| `name` | String! | 原始檔名 |
| `size` | Float! | 檔案大小（以位元組為單位） |
| `type` | String! | MIME 類型 |
| `extension` | String! | 檔案擴展名 |
| `status` | FileStatus | PENDING 或 CONFIRMED（可為空） |
| `shared` | Boolean! | 檔案是否共享 |
| `createdAt` | DateTime! | 上傳時間戳記 |

### TodoCustomFieldFile 回應

| 欄位 | 類型 | 描述 |
|------|------|------|
| `id` | ID! | 交接記錄 ID |
| `uid` | String! | 唯一識別碼 |
| `position` | Float! | 顯示順序 |
| `file` | File! | 關聯檔案物件 |
| `todoCustomField` | TodoCustomField! | 父自訂欄位 |
| `createdAt` | DateTime! | 附加檔案的時間 |

## 使用檔案創建記錄

在創建記錄時，您可以使用其 UID 附加檔案：

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

## 檔案驗證和安全性

### 上傳驗證

- **MIME 類型檢查**：驗證是否符合允許的類型
- **檔案擴展名驗證**：用於 `application/octet-stream` 的備用
- **大小限制**：在上傳時強制執行
- **檔名清理**：移除特殊字符

### 存取控制

- **上傳權限**：需要專案/公司成員資格
- **檔案關聯**：ADMIN、OWNER、MEMBER、CLIENT 角色
- **檔案存取**：繼承自專案/公司權限
- **安全 URL**：時間限制的簽名 URL 用於檔案存取

## 所需權限

| 行動 | 所需權限 |
|------|----------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## 錯誤回應

### 檔案過大
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

### 檔案未找到
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

### 欄位未找到
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

## 最佳實踐

### 檔案管理
- 在附加到記錄之前上傳檔案
- 使用描述性檔名
- 按專案/目的組織檔案
- 定期清理未使用的檔案

### 性能
- 盡可能批量上傳檔案
- 為內容類型使用適當的檔案格式
- 在上傳之前壓縮大型檔案
- 考慮檔案預覽需求

### 安全性
- 驗證檔案內容，而不僅僅是擴展名
- 對上傳的檔案進行病毒掃描
- 實施適當的存取控制
- 監控檔案上傳模式

## 常見用例

1. **文件管理**
   - 專案規範
   - 合同和協議
   - 會議記錄和簡報
   - 技術文檔

2. **資產管理**
   - 設計檔案和模型
   - 品牌資產和標誌
   - 行銷材料
   - 產品圖片

3. **合規和記錄**
   - 法律文件
   - 審計記錄
   - 證書和許可證
   - 財務記錄

4. **協作**
   - 共享資源
   - 版本控制的文件
   - 反饋和註解
   - 參考材料

## 整合功能

### 與自動化
- 當檔案被添加/移除時觸發動作
- 根據類型或元資料處理檔案
- 發送檔案變更的通知
- 根據條件歸檔檔案

### 與封面圖片
- 使用檔案欄位作為封面圖片來源
- 自動圖像處理和縮略圖
- 當檔案變更時動態更新封面

### 與查詢
- 從其他記錄引用檔案
- 聚合檔案數量和大小
- 根據檔案元資料查找記錄
- 交叉引用檔案附件

## 限制

- 每個檔案最大 256MB
- 依賴於 S3 可用性
- 無內建檔案版本控制
- 無自動檔案轉換
- 有限的檔案預覽能力
- 無實時協作編輯

## 相關資源

- [上傳檔案 API](/api/upload-files) - 檔案上傳端點
- [自訂欄位概述](/api/custom-fields/list-custom-fields) - 一般概念
- [自動化 API](/api/automations) - 基於檔案的自動化
- [AWS S3 文檔](https://docs.aws.amazon.com/s3/) - 儲存後端