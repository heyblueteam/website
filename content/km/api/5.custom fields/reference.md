---
title: ការយោងវាលផ្ទាល់ខ្លួន
description: បង្កើតវាលយោងដែលភ្ជាប់ទៅកាន់កំណត់ត្រានៅក្នុងគម្រោងផ្សេងៗសម្រាប់ទំនាក់ទំនងអន្ដរាគមន៍
---

វាលផ្ទាល់ខ្លួនដែលមានការយោងអនុញ្ញាតឱ្យអ្នកបង្កើតតំណភ្ជាប់រវាងកំណត់ត្រានៅក្នុងគម្រោងផ្សេងៗ ដែលអាចធ្វើឱ្យមានទំនាក់ទំនងអន្ដរាគមន៍ និងការចែករំលែកទិន្នន័យ។ វាផ្តល់នូវវិធីដ៏មានអំណាចក្នុងការតភ្ជាប់ការងារដែលពាក់ព័ន្ធនៅក្នុងរចនាសម្ព័ន្ធគម្រោងរបស់អង្គភាពរបស់អ្នក។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលយោងសាមញ្ញ៖

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

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលយោងដែលមានការតម្រង និងការជ្រើសរើសច្រើន៖

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

## ប៉ារ៉ាម៉ែត្រនាំចូល

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលយោង |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `REFERENCE` |
| `referenceProjectId` | String | មិន | អត្តសញ្ញាណនៃគម្រោងដែលត្រូវបានយោង |
| `referenceMultiple` | Boolean | មិន | អនុញ្ញាតឱ្យមានការជ្រើសរើសកំណត់ត្រាច្រើន (លំនាំដើម: មិន) |
| `referenceFilter` | TodoFilterInput | មិន | លក្ខខណ្ឌតម្រងសម្រាប់កំណត់ត្រាដែលបានយោង |
| `description` | String | មិន | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |

**ចំណាំ**: វាលផ្ទាល់ខ្លួនត្រូវបានភ្ជាប់ដោយស្វ័យប្រវត្តិជាមួយគម្រោងដោយផ្អែកលើបរិបទគម្រោងបច្ចុប្បន្នរបស់អ្នកប្រើ។

## ការកំណត់យោង

### ការយោងតែមួយប៉ុណ្ណោះ និងការយោងច្រើន

**ការយោងតែមួយ (លំនាំដើម):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- អ្នកប្រើអាចជ្រើសរើសកំណត់ត្រាមួយពីគម្រោងដែលបានយោង
- ត្រឡប់មកវិញជាវត្ថុ Todo មួយ

**ការយោងច្រើន:**
```graphql
{
  referenceMultiple: true
}
```
- អ្នកប្រើអាចជ្រើសរើសកំណត់ត្រាច្រើនពីគម្រោងដែលបានយោង
- ត្រឡប់មកវិញជាអារ៉ាយនៃវត្ថុ Todo

### ការតម្រងយោង

ប្រើ `referenceFilter` ដើម្បីកំណត់កំណត់ត្រាដែលអាចជ្រើសរើសបាន៖

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

## ការកំណត់តម្លៃយោង

### ការយោងតែមួយ

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### ការយោងច្រើន

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

### SetTodoCustomFieldInput ប៉ារ៉ាម៉ែត្រ

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ បាទ | អត្តសញ្ញាណនៃកំណត់ត្រាដែលត្រូវបានធ្វើបច្ចុប្បន្នភាព |
| `customFieldId` | String! | ✅ បាទ | អត្តសញ្ញាណនៃវាលផ្ទាល់ខ្លួនដែលបានយោង |
| `customFieldReferenceTodoIds` | [String!] | ✅ បាទ | អារ៉ាយនៃអត្តសញ្ញាណនៃកំណត់ត្រាដែលបានយោង |

## ការបង្កើតកំណត់ត្រាជាមួយការយោង

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

## វាលឆ្លើយតប

### TodoCustomField ឆ្លើយតប

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | ID! | អត្តសញ្ញាណឯកតាសម្រាប់តម្លៃវាល |
| `customField` | CustomField! | ការកំណត់វាលយោង |
| `todo` | Todo! | កំណត់ត្រាដែលតម្លៃនេះពាក់ព័ន្ធ |
| `createdAt` | DateTime! | ពេលដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលដែលតម្លៃត្រូវបានកែប្រែចុងក្រោយ |

**ចំណាំ**: ការងារដែលបានយោងត្រូវបានចូលដំណើរការតាមរយៈ `customField.selectedTodos` មិនមែនតាមរយៈ TodoCustomField ដោយផ្ទាល់ទេ។

### វាល Todo ដែលបានយោង

Todo ដែលបានយោងមួយៗមាន៖

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | ID! | អត្តសញ្ញាណឯកតារបស់កំណត់ត្រាដែលបានយោង |
| `title` | String! | ចំណងជើងនៃកំណត់ត្រាដែលបានយោង |
| `status` | TodoStatus! | ស្ថានភាពបច្ចុប្បន្ន (ACTIVE, COMPLETED, ល.) |
| `description` | String | ការពិពណ៌នានៃកំណត់ត្រាដែលបានយោង |
| `dueDate` | DateTime | ថ្ងៃកំណត់ប្រសិនបើបានកំណត់ |
| `assignees` | [User!] | អ្នកប្រើដែលបានចាត់តាំង |
| `tags` | [Tag!] | ស្លាកដែលពាក់ព័ន្ធ |
| `project` | Project! | គម្រោងដែលមានកំណត់ត្រាដែលបានយោង |

## ការស្វែងរកទិន្នន័យយោង

### សំណើមូលដ្ឋាន

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

### សំណើកម្រិតខ្ពស់ជាមួយទិន្នន័យដែលមានស្រទាប់

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

## អាជ្ញាប័ណ្ណដែលត្រូវការ

| សកម្មភាព | អាជ្ញាប័ណ្ណដែលត្រូវការ |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**សំខាន់**: អ្នកប្រើត្រូវមានអាជ្ញាប័ណ្ណមើលនៅលើគម្រោងដែលបានយោងដើម្បីមើលកំណត់ត្រាដែលភ្ជាប់។

## ការចូលដំណើរការអន្ដរាគមន៍

### ការមើលឃើញគម្រោង

- អ្នកប្រើអាចយោងតែទៅកាន់កំណត់ត្រានៅក្នុងគម្រោងដែលពួកគេមានការចូលដំណើរការ
- កំណត់ត្រាដែលបានយោងគោរពអាជ្ញាប័ណ្ណនៃគម្រោងដើម
- ការផ្លាស់ប្តូរទៅកាន់កំណត់ត្រាដែលបានយោងបង្ហាញនៅក្នុងពេលពិត
- ការលុបកំណត់ត្រាដែលបានយោងនឹងលុបចេញពីវាលយោង

### ការធ្វើឱ្យអាជ្ញាប័ណ្ណមកពីមាត្រា

- វាលយោងទទួលអាជ្ញាប័ណ្ណពីគម្រោងទាំងពីរ
- អ្នកប្រើត្រូវការចូលមើលទៅគម្រោងដែលបានយោង
- អាជ្ញាប័ណ្ណកែប្រែផ្អែកលើច្បាប់នៃគម្រោងបច្ចុប្បន្ន
- ទិន្នន័យដែលបានយោងគឺអានតែប៉ុណ្ណោះនៅក្នុងបរិបទនៃវាលយោង

## ចម្លើយកំហុស

### គម្រោងយោងមិនត្រឹមត្រូវ

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

### កំណត់ត្រាដែលបានយោងមិនឃើញ

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

### អាជ្ញាប័ណ្ណត្រូវបានបដិសេធ

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

## អនុវត្តន៍ល្អ

### ការរចនាវាល

1. **ឈ្មោះច្បាស់** - ប្រើឈ្មោះដែលពិពណ៌នាដែលបង្ហាញពីទំនាក់ទំនង
2. **ការតម្រងដែលសមស្រប** - កំណត់តម្រងដើម្បីបង្ហាញតែកំណត់ត្រាដែលពាក់ព័ន្ធ
3. **ពិចារណាអាជ្ញាប័ណ្ណ** - ធានាថាអ្នកប្រើមានការចូលដំណើរការទៅគម្រោងដែលបានយោង
4. **ឯកសារទំនាក់ទំនង** - ផ្តល់នូវការពិពណ៌នាដែលច្បាស់លាស់អំពីការតភ្ជាប់

### ការពិចារណាអំពីសមត្ថភាព

1. **កំណត់វិសាលភាពយោង** - ប្រើតម្រងដើម្បីកាត់បន្ថយចំនួនកំណត់ត្រាដែលអាចជ្រើសរើសបាន
2. **ជៀសវាងការស្រទាប់ជ្រៅ** - កុំបង្កើតខ្សែភ្ជាប់ដែលស្មុគស្មាញ
3. **ពិចារណាការផ្ទុក** - ទិន្នន័យដែលបានយោងត្រូវបានផ្ទុកសម្រាប់សមត្ថភាព
4. **តាមដានការប្រើប្រាស់** - តាមដានរបៀបដែលការយោងត្រូវបានប្រើនៅក្នុងគម្រោង

### សុវត្ថិភាពទិន្នន័យ

1. **ដោះស្រាយការលុប** - គ្រោងសម្រាប់ពេលដែលកំណត់ត្រាដែលបានយោងត្រូវបានលុប
2. **ផ្ទៀងផ្ទាត់អាជ្ញាប័ណ្ណ** - ធានាថាអ្នកប្រើអាចចូលដំណើរការទៅគម្រោងដែលបានយោង
3. **ធ្វើបច្ចុប្បន្នភាពអាស្រ័យភាព** - ពិចារណាភាពឥទ្ធិពលពេលប្តូរកំណត់ត្រាដែលបានយោង
4. **តាមដានការត្រួតពិនិត្យ** - តាមដានទំនាក់ទំនងយោងសម្រាប់ការអនុវត្តន៍

## ករណីប្រើប្រាស់ទូទៅ

### ការពឹងផ្អែកលើគម្រោង

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

### តម្រូវការរបស់អតិថិជន

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

### ការចែកចាយធនធាន

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

### ការធានាគុណភាព

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

## ការរួមបញ្ចូលជាមួយការស្វែងរក

វាលយោងធ្វើការជាមួយ [វាលស្វែងរក](/api/custom-fields/lookup) ដើម្បីទាញយកទិន្នន័យពីកំណត់ត្រាដែលបានយោង។ វាលស្វែងរកអាចទាញយកតម្លៃពីកំណត់ត្រាដែលបានជ្រើសរើសនៅក្នុងវាលយោង ប៉ុន្តែវាជាឧបករណ៍ទាញយកទិន្នន័យតែប៉ុណ្ណោះ (មិនមានមុខងារបូកដូចជា SUM ត្រូវបានគាំទ្រ)។

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

## កំណត់

- គម្រោងដែលបានយោងត្រូវតែអាចចូលដំណើរការបានសម្រាប់អ្នកប្រើ
- ការផ្លាស់ប្តូរទៅអាជ្ញាប័ណ្ណគម្រោងដែលបានយោងមានអត្ថពលន៍លើការចូលដំណើរការវាលយោង
- ការស្រទាប់ជ្រៅនៃការយោងអាចប៉ះពាល់ដល់សមត្ថភាព
- មិនមានការផ្ទៀងផ្ទាត់ក្នុងសម្រាប់ការយោងជុំវិញ
- មិនមានការកំណត់ស្វ័យប្រវត្តិដែលរារាំងការយោងពីគម្រោងដូចគ្នា
- ការផ្ទៀងផ្ទាត់តម្រងមិនត្រូវបានអនុវត្តនៅពេលកំណត់តម្លៃយោង

## ឯកសារដែលពាក់ព័ន្ធ

- [វាលស្វែងរក](/api/custom-fields/lookup) - ទាញយកទិន្នន័យពីកំណត់ត្រាដែលបានយោង
- [API គម្រោង](/api/projects) - គ្រប់គ្រងគម្រោងដែលមានការយោង
- [API កំណត់ត្រា](/api/records) - ការងារជាមួយកំណត់ត្រាដែលមានការយោង
- [ទិដ្ឋភាពទូទៅនៃវាលផ្ទាល់ខ្លួន](/api/custom-fields/list-custom-fields) - គំនិតទូទៅ