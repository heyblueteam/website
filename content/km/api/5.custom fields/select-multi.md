---
title: វាលកំណត់ជម្រើសច្រើន
description: បង្កើតវាលជម្រើសច្រើនដើម្បីអនុញ្ញាតឱ្យអ្នកប្រើជ្រើសរើសជម្រើសច្រើនពីបញ្ជីដែលបានកំណត់
---

វាលកំណត់ជម្រើសច្រើនអនុញ្ញាតឱ្យអ្នកប្រើជ្រើសរើសជម្រើសច្រើនពីបញ្ជីដែលបានកំណត់។ វាគឺជាដំណោះស្រាយល្អសម្រាប់ប្រភេទ, ស្លាក, ជំនាញ, លក្ខណៈពិសេស, ឬស្ថានការណ៍ណាមួយដែលត្រូវការជម្រើសច្រើនពីសំណុំជម្រើសដែលត្រូវគ្រប់គ្រង។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលជម្រើសច្រើនសាមញ្ញ៖

```graphql
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Project Categories"
    type: SELECT_MULTI
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលជម្រើសច្រើនហើយបន្ថែមជម្រើសដោយឡែក៖

```graphql
# Step 1: Create the multi-select field
mutation CreateMultiSelectField {
  createCustomField(input: {
    name: "Required Skills"
    type: SELECT_MULTI
    projectId: "proj_123"
    description: "Select all skills required for this task"
  }) {
    id
    name
    type
    description
  }
}

# Step 2: Add options to the field
mutation AddOptions {
  createCustomFieldOptions(input: [
    { customFieldId: "field_123", title: "JavaScript", color: "#f7df1e" }
    { customFieldId: "field_123", title: "React", color: "#61dafb" }
    { customFieldId: "field_123", title: "Node.js", color: "#339933" }
    { customFieldId: "field_123", title: "GraphQL", color: "#e10098" }
  ]) {
    id
    title
    color
    position
  }
}
```

## ប៉ារ៉ាម៉ែត្រ Input

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលជម្រើសច្រើន |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `SELECT_MULTI` |
| `description` | String | ទេ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |
| `projectId` | String! | ✅ បាទ | អត្តសញ្ញាណនៃគម្រោងសម្រាប់វាលនេះ |

### CreateCustomFieldOptionInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ បាទ | អត្តសញ្ញាណនៃវាលកំណត់ |
| `title` | String! | ✅ បាទ | អត្ថបទបង្ហាញសម្រាប់ជម្រើស |
| `color` | String | ទេ | ពណ៌សម្រាប់ជម្រើស (អត្ថបទណាមួយ) |
| `position` | Float | ទេ | លំដាប់តម្រៀបសម្រាប់ជម្រើស |

## បន្ថែមជម្រើសទៅវាលដែលមានស្រាប់

បន្ថែមជម្រើសថ្មីទៅវាលជម្រើសច្រើនដែលមានស្រាប់៖

```graphql
mutation AddMultiSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Python"
    color: "#3776ab"
  }) {
    id
    title
    color
    position
  }
}
```

## កំណត់តម្លៃជម្រើសច្រើន

ដើម្បីកំណត់ជម្រើសដែលបានជ្រើសច្រើននៅលើកំណត់ត្រា៖

```graphql
mutation SetMultiSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionIds: ["option_1", "option_2", "option_3"]
  })
}
```

### SetTodoCustomFieldInput Parameters

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ បាទ | អត្តសញ្ញាណនៃកំណត់ត្រាដែលត្រូវអាប់ដេត |
| `customFieldId` | String! | ✅ បាទ | អត្តសញ្ញាណនៃវាលកំណត់ជម្រើសច្រើន |
| `customFieldOptionIds` | [String!] | ✅ បាទ | អារ៉ាយនៃអត្តសញ្ញាណជម្រើសដើម្បីជ្រើស |

## បង្កើតកំណត់ត្រាជាមួយតម្លៃជម្រើសច្រើន

នៅពេលបង្កើតកំណត់ត្រាថ្មីជាមួយតម្លៃជម្រើសច្រើន៖

```graphql
mutation CreateRecordWithMultiSelect {
  createTodo(input: {
    title: "Develop new feature"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "skills_field_id"
      value: "option1,option2,option3"
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
      selectedOptions {
        id
        title
        color
      }
    }
  }
}
```

## វាលចម្លើយ

### TodoCustomField Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកតាសម្រាប់តម្លៃវាល |
| `customField` | CustomField! | ការកំណត់វាលកំណត់ |
| `selectedOptions` | [CustomFieldOption!] | អារ៉ាយនៃជម្រើសដែលបានជ្រើស |
| `todo` | Todo! | កំណត់ត្រាដែលតម្លៃនេះជាប់នឹង |
| `createdAt` | DateTime! | ពេលដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលដែលតម្លៃត្រូវបានកែប្រែចុងក្រោយ |

### CustomFieldOption Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកតាសម្រាប់ជម្រើស |
| `title` | String! | អត្ថបទបង្ហាញសម្រាប់ជម្រើស |
| `color` | String | កូដពណ៌ Hex សម្រាប់ការតំណាងវិស្វកម្ម |
| `position` | Float | លំដាប់តម្រៀបសម្រាប់ជម្រើស |
| `customField` | CustomField! | វាលកំណត់ដែលជាជម្រើសនេះ |

### CustomField Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកតាសម្រាប់វាល |
| `name` | String! | ឈ្មោះបង្ហាញនៃវាលជម្រើសច្រើន |
| `type` | CustomFieldType! | តែងតែ `SELECT_MULTI` |
| `description` | String | អត្ថបទជំនួយសម្រាប់វាល |
| `customFieldOptions` | [CustomFieldOption!] | ជម្រើសទាំងអស់ដែលមានស្រាប់ |

## ទ្រង់ទ្រាយតម្លៃ

### ទ្រង់ទ្រាយ Input
- **API Parameter**: អារ៉ាយនៃអត្តសញ្ញាណជម្រើស (`["option1", "option2", "option3"]`)
- **ទ្រង់ទ្រាយអត្ថបទ**: អត្តសញ្ញាណជម្រើសដែលបំបែកដោយក្បាល (`"option1,option2,option3"`)

### ទ្រង់ទ្រាយ Output
- **GraphQL Response**: អារ៉ាយនៃវត្ថុ CustomFieldOption
- **Activity Log**: អត្ថបទជម្រើសដែលបំបែកដោយក្បាល
- **ទិន្នន័យស្វ័យប្រវត្តិ**: អារ៉ាយនៃអត្ថបទជម្រើស

## ការគ្រប់គ្រងជម្រើស

### កែប្រែលក្ខណៈជម្រើស
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Title"
    color: "#ff0000"
  }) {
    id
    title
    color
  }
}
```

### លុបជម្រើស
```graphql
mutation DeleteOption {
  deleteCustomFieldOption(id: "option_123")
}
```

### កែប្រែលំដាប់ជម្រើស
```graphql
# Update position values to reorder options
mutation UpdateOptionPosition {
  editCustomFieldOption(input: {
    id: "option_123"
    position: 1.5  # Position between 1.0 and 2.0
  }) {
    id
    position
  }
}
```

## ច្បាប់ផ្ទៀងផ្ទាត់

### ការផ្ទៀងផ្ទាត់ជម្រើស
- អត្តសញ្ញាណជម្រើសទាំងអស់ដែលផ្តល់ត្រូវតែមានស្រាប់
- ជម្រើសត្រូវតែជាប់នឹងវាលកំណត់ដែលបានកំណត់
- គ្រាន់តែ SELECT_MULTI វាលអាចមានជម្រើសច្រើនដែលបានជ្រើស
- អារ៉ាយទទេគឺមានសុពលភាព (គ្មានការជ្រើស)

### ការផ្ទៀងផ្ទាត់វាល
- ត្រូវមានជម្រើសយ៉ាងហោចណាស់មួយដែលបានកំណត់ដើម្បីអាចប្រើបាន
- អត្ថបទជម្រើសត្រូវតែមានឯកត្តភាពនៅក្នុងវាល
- វាលពណ៌ទទួលយកតម្លៃអត្ថបទណាមួយ (គ្មានការផ្ទៀងផ្ទាត់ hex)

## អាជ្ញាប័ណ្ណដែលត្រូវការ

| សកម្មភាព | អាជ្ញាប័ណ្ណដែលត្រូវការ |
|--------|-------------------|
| Create multi-select field | `OWNER` or `ADMIN` role at project level |
| Update multi-select field | `OWNER` or `ADMIN` role at project level |
| Add/edit options | `OWNER` or `ADMIN` role at project level |
| Set selected values | Standard record edit permissions |
| View selected values | Standard record view permissions |

## ចម្លើយកំហុស

### អត្តសញ្ញាណជម្រើសមិនត្រឹមត្រូវ
```json
{
  "errors": [{
    "message": "Custom field option not found",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### ជម្រើសមិន属于វាល
```json
{
  "errors": [{
    "message": "Option does not belong to this custom field",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### វាលមិនមាន
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### ជម្រើសច្រើននៅលើវាលដែលមិនមែនជាច្រើន
```json
{
  "errors": [{
    "message": "custom fields can only have one option",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## អនុសាសន៍ល្អ

### ការរចនាជម្រើស
- ប្រើឈ្មោះជម្រើសដែលពិពណ៌នាដោយច្បាស់និងខ្លី
- អនុវត្តន៍ស្កេមពណ៌ដែលមានភាពស្រប
- រក្សាទុកបញ្ជីជម្រើសឱ្យគ្រប់គ្រាន់ (ធម្មតា 3-20 ជម្រើស)
- តម្រៀបជម្រើសដោយមានលំនាំ (តាមអក្សរ, តាមឧត្តមភាព, ល។)

### ការគ្រប់គ្រងទិន្នន័យ
- ពិនិត្យមើលនិងសម្អាតជម្រើសដែលមិនបានប្រើជាប្រចាំ
- ប្រើសេចក្តីថ្លែងការណ៍ឈ្មោះដែលមានភាពស្របនៅក្នុងគម្រោង
- ពិចារណាអំពីការប្រើប្រាស់ជម្រើសឡើងវិញនៅពេលបង្កើតវាល
- គ្រោងសម្រាប់ការអាប់ដេតនិងការផ្លាស់ប្តូរជម្រើស

### បទពិសោធន៍អ្នកប្រើ
- ផ្តល់ការពិពណ៌នាវាលដែលច្បាស់
- ប្រើពណ៌ដើម្បីធ្វើឱ្យមានភាពខុសគ្នាដោយច្បាស់
- ក្រុមជម្រើសដែលពាក់ព័ន្ធគ្នា
- ពិចារណាជម្រើសលំនាំដើមសម្រាប់ករណីទូទៅ

## ករណីប្រើទូទៅ

1. **ការគ្រប់គ្រងគម្រោង**
   - ប្រភេទកិច្ចការនិងស្លាក
   - កម្រិតនិងប្រភេទអាទិភាព
   - ការបែងចែកសមាជិកក្រុម
   - សញ្ញាស្ថានភាព

2. **ការគ្រប់គ្រងមាតិកា**
   - ប្រភេទនិងប្រធានបទអត្ថបទ
   - ប្រភេទនិងទ្រង់ទ្រាយមាតិកា
   - ช่องทางផ្សព្វផ្សាយ
   - ការបញ្ជូនការអនុម័ត

3. **ការគាំទ្រអតិថិជន**
   - ប្រភេទនិងប្រភេទបញ្ហា
   - ផលិតផលឬសេវាកម្មដែលទាក់ទង
   - វិធីសាស្ត្រដោះស្រាយ
   - ក្រុមអតិថិជន

4. **ការអភិវឌ្ឍផលិតផល**
   - ប្រភេទលក្ខណៈពិសេស
   - តម្រូវការបច្ចេកទេស
   - បរិយាកាសសាកល្បង
   - ช่องทางចេញផ្សាយ

## លក្ខណៈពិសេសនៃការបញ្ចូល

### ជាមួយស្វ័យប្រវត្តិ
- បង្កើតសកម្មភាពនៅពេលជម្រើសជាក់លាក់ត្រូវបានជ្រើស
- ផ្លូវការងារយោងទៅលើប្រភេទដែលបានជ្រើស
- ផ្ញើការជូនដំណឹងសម្រាប់ការជ្រើសដែលមានអាទិភាពខ្ពស់
- បង្កើតកិច្ចការតាមដាននៅលើការបញ្ចូលជម្រើស

### ជាមួយការស្វែងរក
- បង្ហាញកំណត់ត្រាដោយជម្រើសដែលបានជ្រើស
- ប្រមូលទិន្នន័យនៅជុំវិញការជ្រើសជម្រើស
- យោងទៅកាន់ទិន្នន័យជម្រើសពីកំណត់ត្រាផ្សេងទៀត
- បង្កើតរបាយការណ៍នៅលើការបញ្ចូលជម្រើស

### ជាមួយទម្រង់
- គ្រប់គ្រងការបញ្ចូលជម្រើសច្រើន
- ការផ្ទៀងផ្ទាត់និងការប្រមូលជម្រើស
- ការផ្ទុកជម្រើសដោយឆាប់រហ័ស
- ការបង្ហាញវាលលក្ខខណ្ឌ

## ការតាមដានសកម្មភាព

ការផ្លាស់ប្តូរវាលជម្រើសច្រើនត្រូវបានតាមដានដោយស្វ័យប្រវត្តិ៖
- បង្ហាញជម្រើសដែលបានបន្ថែមនិងលុប
- បង្ហាញអត្ថបទជម្រើសនៅក្នុងកំណត់ហេតុសកម្មភាព
- ម៉ោងសម្រាប់ការផ្លាស់ប្តូរទាំងអស់
- ការបញ្ជាក់អ្នកប្រើសម្រាប់ការកែប្រែ

## ការកំណត់

- លំហឱ្យមានប្រយោជន៍អតិបរិមានៃជម្រើសអាស្រ័យលើសមត្ថភាព UI
- គ្មានរចនាសម្ព័ន្ធជម្រើសដែលមានលំដាប់ឬជាន់
- ជម្រើសត្រូវបានចែករំលែកនៅលើកំណត់ត្រាទាំងអស់ដែលប្រើវាល
- គ្មានវិភាគជម្រើសឬការតាមដានការប្រើប្រាស់ដែលមានស្រាប់
- វាលពណ៌ទទួលយកតម្លៃអត្ថបទណាមួយ (គ្មានការផ្ទៀងផ្ទាត់ hex)
- មិនអាចកំណត់អាជ្ញាប័ណ្ណផ្សេងៗសម្រាប់ជម្រើសមួយៗ
- ជម្រើសត្រូវបានបង្កើតដោយឡែក, មិនមែននៅក្នុងការបង្កើតវាល
- គ្មានការប្រែប្រួលលំដាប់ដែលមានស្រាប់ (ប្រើ editCustomFieldOption ជាមួយទីតាំង)

## ឯកសារដែលទាក់ទង

- [វាលជម្រើសតែមួយ](/api/custom-fields/select-single) - សម្រាប់ការជ្រើសរើសជម្រើសតែមួយ
- [វាលប្រអប់ស្លាក](/api/custom-fields/checkbox) - សម្រាប់ជម្រើស boolean ងាយស្រួល
- [វាលអត្ថបទ](/api/custom-fields/text-single) - សម្រាប់ការបញ្ចូលអត្ថបទដោយសេរី
- [ទិដ្ឋភាពទូទៅនៃវាលកំណត់](/api/custom-fields/2.list-custom-fields) - គំនិតទូទៅ