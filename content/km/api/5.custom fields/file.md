---
title: សមាសភាគកំណត់ឯកសារ
description: បង្កើតកំណត់ឯកសារដើម្បីភ្ជាប់ឯកសារ រូបភាព និងឯកសារផ្សេងទៀតទៅកាន់កំណត់ត្រា
---

សមាសភាគកំណត់ឯកសារអនុញ្ញាតឱ្យអ្នកភ្ជាប់ឯកសារច្រើនទៅកាន់កំណត់ត្រា។ ឯកសារត្រូវបានរក្សាទុកយ៉ាងសុវត្ថិភាពនៅក្នុង AWS S3 ជាមួយនឹងការតាមដានមេតាដាតាប្រភេទពេញលេញ ការផ្ទៀងផ្ទាត់ប្រភេទឯកសារ និងការគ្រប់គ្រងការចូលប្រើយ៉ាងត្រឹមត្រូវ។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតសមាសភាគឯកសារងាយៗ៖

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

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតសមាសភាគឯកសារដែលមានការពិពណ៌នា៖

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

## ព័ត៌មានចូល

### CreateCustomFieldInput

| ព័ត៌មាន | ប្រភេទ | ត្រូវការ | ពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃសមាសភាគឯកសារ |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `FILE` |
| `description` | String | អត់ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |

**ចំណាំ**: សមាសភាគកំណត់ត្រូវបានភ្ជាប់ដោយស្វ័យប្រវត្តិជាមួយគម្រោងដោយផ្អែកលើបរិបទគម្រោងបច្ចុប្បន្នរបស់អ្នកប្រើ។ មិនត្រូវការព័ត៌មាន `projectId` ទេ។

## ដំណើរការបញ្ចូលឯកសារ

### ជំហានទី 1: បញ្ចូលឯកសារ

ដំបូង បញ្ចូលឯកសារដើម្បីទទួលបាន UID ឯកសារ៖

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

### ជំហានទី 2: ភ្ជាប់ឯកសារទៅកាន់កំណត់ត្រា

បន្ទាប់មកភ្ជាប់ឯកសារដែលបានបញ្ចូលទៅកាន់កំណត់ត្រា៖

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

## ការគ្រប់គ្រងការភ្ជាប់ឯកសារ

### ការបន្ថែមឯកសារតែមួយ

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

### ការលុបឯកសារ

```graphql
mutation RemoveFileFromField {
  deleteTodoCustomFieldFile(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    fileUid: "file_uid_789"
  })
}
```

### ប្រតិបត្តិការឯកសារប្រកបដោយមាស

ធ្វើអាប់ដេតឯកសារច្រើនក្នុងមួយពេលដោយប្រើ customFieldOptionIds:

```graphql
mutation SetMultipleFiles {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["file_uid_1", "file_uid_2", "file_uid_3"]
  })
}
```

## ព័ត៌មានចូលសម្រាប់បញ្ចូលឯកសារ

### UploadFileInput

| ព័ត៌មាន | ប្រភេទ | ត្រូវការ | ពិពណ៌នា |
|-----------|------|----------|-------------|
| `file` | Upload! | ✅ បាទ | ឯកសារដើម្បីបញ្ចូល |
| `companyId` | String! | ✅ បាទ | ID ក្រុមហ៊ុនសម្រាប់ការផ្ទុកឯកសារ |
| `projectId` | String | អត់ | ID គម្រោងសម្រាប់ឯកសារដែលមានលក្ខណៈពិសេស |

### ព័ត៌មានចូលសម្រាប់ការគ្រប់គ្រងឯកសារ

| ព័ត៌មាន | ប្រភេទ | ត្រូវការ | ពិពណ៌នា |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ បាទ | ID នៃកំណត់ត្រា |
| `customFieldId` | String! | ✅ បាទ | ID នៃសមាសភាគឯកសារ |
| `fileUid` | String! | ✅ បាទ | អត្តសញ្ញាណឯកតាដែលមានភាពឯកទេសនៃឯកសារដែលបានបញ្ចូល |

## ការផ្ទុកឯកសារ និងកំណត់

### កំណត់ទំហំឯកសារ

| ប្រភេទកំណត់ | ទំហំ |
|------------|------|
| Maximum file size | 256MB per file |
| Batch upload limit | 10 files max, 1GB total |
| GraphQL upload limit | 256MB |

### ប្រភេទឯកសារដែលគាំទ្រ

#### រូបភាព
- `jpg`, `jpeg`, `png`, `gif`, `bmp`, `webp`, `svg`, `ico`, `tiff`, `tif`

#### វីដេអូ
- `mp4`, `avi`, `mov`, `wmv`, `flv`, `webm`, `mkv`, `3gp`

#### សូរ
- `mp3`, `wav`, `flac`, `aac`, `ogg`, `wma`

#### ឯកសារ
- `pdf`, `doc`, `docx`, `xls`, `xlsx`, `ppt`, `pptx`, `txt`, `rtf`

#### សារពើភ័ណ្ឌ
- `zip`, `rar`, `7z`, `tar`, `gz`

#### កូដ/អត្ថបទ
- `json`, `xml`, `csv`, `md`, `yaml`, `yml`

### ស្ថាបត្យកម្មការផ្ទុក

- **ការផ្ទុក**: AWS S3 ជាមួយរចនាសម្ព័ន្ធថតដែលបានរៀបចំ
- **ទម្រង់ផ្លូវ**: `companies/{companySlug}/projects/{projectSlug}/uploads/{year}/{month}/{username}/{fileUid}_{filename}`
- **សុវត្ថិភាព**: URLs ដែលបានចុះហត្ថលេខាសម្រាប់ការចូលប្រើយ៉ាងសុវត្ថិភាព
- **ការបម្រុងទុក**: ការបម្រុងទុក S3 ដោយស្វ័យប្រវត្តិ

## វាលឆ្លើយតប

### File Response

| វាល | ប្រភេទ | ពិពណ៌នា |
|-------|------|-------------|
| `id` | ID! | ID ទិន្នន័យ |
| `uid` | String! | អត្តសញ្ញាណឯកសារដែលមានភាពឯកទេស |
| `name` | String! | ឈ្មោះឯកសារដើម |
| `size` | Float! | ទំហំឯកសារក្នុងបាយត្រា |
| `type` | String! | ប្រភេទ MIME |
| `extension` | String! | ប្រភេទឯកសារ |
| `status` | FileStatus | កំពុងរង់ចាំ ឬ បានបញ្ជាក់ (អាចទទេ) |
| `shared` | Boolean! | ថាតើឯកសារនេះត្រូវបានចែករំលែកឬអត់ |
| `createdAt` | DateTime! | ម៉ោងបញ្ចូល |

### TodoCustomFieldFile Response

| វាល | ប្រភេទ | ពិពណ៌នា |
|-------|------|-------------|
| `id` | ID! | ID កំណត់ត្រាសម្ព័ន្ធ |
| `uid` | String! | អត្តសញ្ញាណឯកទេស |
| `position` | Float! | លំដាប់បង្ហាញ |
| `file` | File! | វត្ថុឯកសារដែលភ្ជាប់ |
| `todoCustomField` | TodoCustomField! | សមាសភាគកំណត់ឪពុក |
| `createdAt` | DateTime! | ពេលណាដែលឯកសារត្រូវបានភ្ជាប់ |

## ការបង្កើតកំណត់ត្រាជាមួយឯកសារ

នៅពេលបង្កើតកំណត់ត្រា អ្នកអាចភ្ជាប់ឯកសារដោយប្រើ UID របស់ពួកវា៖

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

## ការផ្ទៀងផ្ទាត់ឯកសារ និងសុវត្ថិភាព

### ការផ្ទៀងផ្ទាត់ការបញ្ចូល

- **ការត្រួតពិនិត្យប្រភេទ MIME**: ផ្ទៀងផ្ទាត់ប្រឆាំងនឹងប្រភេទដែលអនុញ្ញាត
- **ការផ្ទៀងផ្ទាត់ប្រភេទឯកសារ**: ការបង្វិលសម្រាប់ `application/octet-stream`
- **កំណត់ទំហំ**: ត្រូវបានអនុវត្តនៅពេលបញ្ចូល
- **ការសម្អាតឈ្មោះឯកសារ**: លុបអក្សរពិសេស

### ការគ្រប់គ្រងការចូល

- **អនុសាសន៍បញ្ចូល**: ត្រូវការការជាសមាជិកគម្រោង/ក្រុមហ៊ុន
- **ការភ្ជាប់ឯកសារ**: តួនាទី ADMIN, OWNER, MEMBER, CLIENT
- **ការចូលប្រើឯកសារ**: ទទួលបានពីសិទ្ធិគម្រោង/ក្រុមហ៊ុន
- **URLs សុវត្ថិភាព**: URLs ដែលបានចុះហត្ថលេខាដែលមានកំណត់ពេលសម្រាប់ការចូលប្រើឯកសារ

## សិទ្ធិដែលត្រូវការ

| សកម្មភាព | សិទ្ធិដែលត្រូវការ |
|--------|-------------------|
| Create file field | `OWNER` or `ADMIN` project-level role |
| Update file field | `OWNER` or `ADMIN` project-level role |
| Upload files | Project or company membership |
| Attach files | ADMIN, OWNER, MEMBER, or CLIENT role |
| View files | Standard record view permissions |
| Delete files | Same as attach permissions |

## ការឆ្លើយតបកំហុស

### ឯកសារធំជាងគេ
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

### ឯកសារមិនឃើញ
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

### វាលមិនឃើញ
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

## អនុសាសន៍ល្អបំផុត

### ការគ្រប់គ្រងឯកសារ
- បញ្ចូលឯកសារមុនពេលភ្ជាប់ទៅកាន់កំណត់ត្រា
- ប្រើឈ្មោះឯកសារដែលពិពណ៌នាដោយច្បាស់
- រៀបចំឯកសារតាមគម្រោង/គោលបំណង
- សម្អាតឯកសារដែលមិនបានប្រើប្រាស់ជាប្រចាំ

### សមត្ថភាព
- បញ្ចូលឯកសារនៅក្នុងក្រុមពេលដែលអាចធ្វើទៅបាន
- ប្រើទ្រង់ទ្រាយឯកសារដែលសមស្របសម្រាប់ប្រភេទមាតិកា
- បង្ហាប់ឯកសារធំមុនពេលបញ្ចូល
- ពិចារណាអំពីការទាមទារបង្ហាញឯកសារ

### សុវត្ថិភាព
- ផ្ទៀងផ្ទាត់មាតិកាឯកសារ មិនត្រឹមតែប្រភេទទេ
- ប្រើការត្រួតពិនិត្យវីរុសសម្រាប់ឯកសារដែលបានបញ្ចូល
- អនុវត្តការគ្រប់គ្រងការចូលប្រើយ៉ាងត្រឹមត្រូវ
- តាមដានលំនាំការបញ្ចូលឯកសារ

## ករណីប្រើប្រាស់ទូទៅ

1. **ការគ្រប់គ្រងឯកសារ**
   - ការបញ្ជាក់គម្រោង
   - កិច្ចសន្យា និងកិច្ចព្រមព្រៀង
   - កំណត់ត្រាសន្និសីទ និងការបង្ហាញ
   - ឯកសារបច្ចេកទេស

2. **ការគ្រប់គ្រងទ្រព្យសម្បត្តិ**
   - ឯកសាររចនា និងការបង្ហាញ
   - ទ្រព្យសម្បត្តិម៉ាក និងឡូហ្គូ
   - សម្ភារៈទីផ្សារ
   - រូបភាពផលិតផល

3. **ការអនុវត្តន៍ និងកំណត់ត្រា**
   - ឯកសារប្រកាសច្បាប់
   - ផ្លូវពិនិត្យ
   - វិញ្ញាបនប័ត្រ និងអាជ្ញាប័ណ្ណ
   - កំណត់ត្រាហិរញ្ញវត្ថុ

4. **ការសហការណ៍**
   - ធនធានដែលបានចែករំលែក
   - ឯកសារដែលមានការគ្រប់គ្រងកំណែ
   - មតិយោបល់ និងការអនុវត្ត
   - សម្ភារៈយោង

## លក្ខណៈពិសេសនៃការបញ្ចូល

### ជាមួយការបង្កើតស្វ័យប្រវត្តិ
- បង្កើតសកម្មភាពនៅពេលដែលឯកសារត្រូវបានបន្ថែម/លុប
- ដំណើរការឯកសារតាមប្រភេទឬមេតាដាតា
- ផ្ញើការជូនដំណឹងសម្រាប់ការផ្លាស់ប្តូរឯកសារ
- បម្រុងឯកសារតាមលក្ខខណ្ឌ

### ជាមួយរូបភាពគម្រប
- ប្រើសមាសភាគឯកសារជាដើមរូបភាពគម្រប
- ការបង្ហាញរូបភាព និងរូបភាពតូចដោយស្វ័យប្រវត្តិ
- ការអាប់ដេតគម្របដោយស្វ័យប្រវត្តិពេលដែលឯកសារប្រែប្រួល

### ជាមួយការស្វែងរក
- យោងឯកសារពីកំណត់ត្រាផ្សេងទៀត
- បូកចំនួន និងទំហំឯកសារ
- ស្វែងរកកំណត់ត្រាដោយមេតាដាតារបស់ឯកសារ
- យោងឯកសារភ្ជាប់

## ការកំណត់

- ទំហំអតិបរមា 256MB សម្រាប់ឯកសារមួយ
- អាស្រ័យលើការមានស្រាប់នៃ S3
- មិនមានការគ្រប់គ្រងកំណែឯកសារដែលមានស្រាប់
- មិនមានការបម្លែងឯកសារដោយស្វ័យប្រវត្តិ
- សមត្ថភាពបង្ហាញឯកសារដែលមានកំណត់
- មិនមានការកែប្រែសហការយ៉ាងពេលវេលាពិត

## ធនធានដែលពាក់ព័ន្ធ

- [API បញ្ចូលឯកសារ](/api/upload-files) - ចំណុចបញ្ចូលឯកសារ
- [ទិដ្ឋភាពទូទៅសមាសភាគកំណត់](/api/custom-fields/list-custom-fields) - គំនិតទូទៅ
- [API ស្វ័យប្រវត្តិ](/api/automations) - ការបង្កើតស្វ័យប្រវត្តិដែលមានឯកសារ
- [ឯកសាររបស់ AWS S3](https://docs.aws.amazon.com/s3/) - ការផ្ទុកក្រោយ