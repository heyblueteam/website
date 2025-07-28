---
title: វាលកំណត់តែមួយដែលអាចប្ដូរបាន
description: បង្កើតវាលកំណត់តែមួយដើម្បីអនុញ្ញាតឱ្យអ្នកប្រើជ្រើសរើសជម្រើសមួយពីបញ្ជីដែលបានកំណត់ជាមុន
---

វាលកំណត់តែមួយអនុញ្ញាតឱ្យអ្នកប្រើជ្រើសរើសជម្រើសតែមួយពីបញ្ជីដែលបានកំណត់ជាមុន។ វាជាការបញ្ជាក់សម្រាប់វាលស្ថានភាព, ប្រភេទ, អាទិភាព, ឬស្ថានភាពណាមួយដែលត្រូវបានជ្រើសរើសតែមួយពីកំណត់ជម្រើសដែលត្រូវគ្រប់គ្រង។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលកំណត់តែមួយដែលសាមញ្ញ:

```graphql
mutation CreateSingleSelectField {
  createCustomField(input: {
    name: "Project Status"
    type: SELECT_SINGLE
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលកំណត់តែមួយជាមួយជម្រើសដែលបានកំណត់ជាមុន:

```graphql
mutation CreateDetailedSingleSelectField {
  createCustomField(input: {
    name: "Priority Level"
    type: SELECT_SINGLE
    projectId: "proj_123"
    description: "Set the priority level for this task"
    customFieldOptions: [
      { title: "Low", color: "#28a745" }
      { title: "Medium", color: "#ffc107" }
      { title: "High", color: "#fd7e14" }
      { title: "Critical", color: "#dc3545" }
    ]
  }) {
    id
    name
    type
    description
    customFieldOptions {
      id
      title
      color
      position
    }
  }
}
```

## ព័ត៌មានបញ្ចូល

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការទេ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលកំណត់តែមួយ |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `SELECT_SINGLE` |
| `description` | String | ទេ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |
| `customFieldOptions` | [CreateCustomFieldOptionInput!] | ទេ | ជម្រើសដំបូងសម្រាប់វាល |

### CreateCustomFieldOptionInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការទេ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `title` | String! | ✅ បាទ | អត្ថបទបង្ហាញសម្រាប់ជម្រើស |
| `color` | String | ទេ | កូដពណ៌ Hex សម្រាប់ជម្រើស |

## បន្ថែមជម្រើសទៅវាលដែលមានស្រាប់

បន្ថែមជម្រើសថ្មីទៅវាលកំណត់តែមួយដែលមានស្រាប់:

```graphql
mutation AddSingleSelectOption {
  createCustomFieldOption(input: {
    customFieldId: "field_123"
    title: "Urgent"
    color: "#6f42c1"
  }) {
    id
    title
    color
    position
  }
}
```

## កំណត់តម្លៃកំណត់តែមួយ

ដើម្បីកំណត់ជម្រើសដែលបានជ្រើសនៅលើកំណត់ត្រា:

```graphql
mutation SetSingleSelectValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldOptionId: "option_789"
  })
}
```

### SetTodoCustomFieldInput Parameters

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការទេ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ បាទ | ID នៃកំណត់ត្រាដែលត្រូវធ្វើបច្ចុប្បន្នភាព |
| `customFieldId` | String! | ✅ បាទ | ID នៃវាលកំណត់តែមួយ |
| `customFieldOptionId` | String | ទេ | ID នៃជម្រើសដែលត្រូវជ្រើស (មិនគួរឱ្យចាប់អារម្មណ៍សម្រាប់កំណត់តែមួយ) |
| `customFieldOptionIds` | [String!] | ទេ | អារ៉ៃនៃ ID ជម្រើស (ប្រើធាតុដំបូងសម្រាប់កំណត់តែមួយ) |

## ការស្វែងរកតម្លៃកំណត់តែមួយ

ស្វែងរកតម្លៃកំណត់តែមួយរបស់កំណត់ត្រា:

```graphql
query GetRecordWithSingleSelect {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
      }
      value  # For SELECT_SINGLE, contains: {"id": "opt_123", "title": "High", "color": "#dc3545", "position": 3}
    }
  }
}
```

វាល `value` ត្រឡប់មកវិញជាអ objeck JSON ជាមួយព័ត៌មានលម្អិតនៃជម្រើសដែលបានជ្រើស។

## ការបង្កើតកំណត់ត្រាជាមួយតម្លៃកំណត់តែមួយ

នៅពេលបង្កើតកំណត់ត្រាថ្មីជាមួយតម្លៃកំណត់តែមួយ:

```graphql
mutation CreateRecordWithSingleSelect {
  createTodo(input: {
    title: "Review user feedback"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "priority_field_id"
      customFieldOptionId: "option_high_priority"
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
      value  # Contains the selected option object
    }
  }
}
```

## វាលឆ្លើយតប

### TodoCustomField Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកត្តសម្រាប់តម្លៃវាល |
| `customField` | CustomField! | ការបញ្ជាក់វាលកំណត់ |
| `value` | JSON | រក្សាទុកវត្ថុជម្រើសដែលបានជ្រើសដែលមាន id, title, color, position |
| `todo` | Todo! | កំណត់ត្រាដែលតម្លៃនេះស្ថិតនៅក្នុង |
| `createdAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានកែប្រែចុងក្រោយ |

### CustomFieldOption Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកត្តសម្រាប់ជម្រើស |
| `title` | String! | អត្ថបទបង្ហាញសម្រាប់ជម្រើស |
| `color` | String | កូដពណ៌ Hex សម្រាប់ការតំណាងវិស្វកម្ម |
| `position` | Float | លំដាប់ការរៀបចំសម្រាប់ជម្រើស |
| `customField` | CustomField! | វាលកំណត់ដែលជាជម្រើសនេះ |

### CustomField Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកត្តសម្រាប់វាល |
| `name` | String! | ឈ្មោះបង្ហាញនៃវាលកំណត់តែមួយ |
| `type` | CustomFieldType! | តែងតែ `SELECT_SINGLE` |
| `description` | String | អត្ថបទជំនួយសម្រាប់វាល |
| `customFieldOptions` | [CustomFieldOption!] | ជម្រើសទាំងអស់ដែលមានស្រាប់ |

## ទម្រង់តម្លៃ

### ទម្រង់បញ្ចូល
- **API Parameter**: ប្រើ `customFieldOptionId` សម្រាប់ ID ជម្រើសតែមួយ
- **ជំនួស**: ប្រើ `customFieldOptionIds` អារ៉ៃ (យកធាតុដំបូង)
- **ការលុបចេញជម្រើស**: បោះបង់ទាំងពីរវាលឬបញ្ជូនតម្លៃទទេ

### ទម្រង់ចេញ
- **GraphQL Response**: វត្ថុ JSON នៅក្នុង `value` វាលមាន {id, title, color, position}
- **កំណត់ហេតុសកម្មភាព**: ឈ្មោះជម្រើសជាស្ត្រី
- **ទិន្នន័យស្វ័យប្រវត្តិ**: ឈ្មោះជម្រើសជាស្ត្រី

## អាកប្បកិរិយាជម្រើស

### ការជ្រើសរើសតែមួយ
- ការកំណត់ជម្រើសថ្មីនឹងលុបចេញជម្រើសមុន
- តែជម្រើសតែមួយអាចត្រូវបានជ្រើសនៅពេលណាមួយ
- ការកំណត់ `null` ឬតម្លៃទទេនឹងលុបចេញជម្រើស

### តុល្យភាព
- ប្រសិនបើ `customFieldOptionIds` អារ៉ៃត្រូវបានផ្តល់, តែជម្រើសដំបូងត្រូវបានប្រើ
- នេះធានាថាមានភាពសមស្របជាមួយទម្រង់បញ្ចូលជម្រើសច្រើន
- អារ៉ៃទទេឬតម្លៃ null លុបចេញជម្រើស

## ការគ្រប់គ្រងជម្រើស

### ការអាប់ដេតលក្ខណៈជម្រើស
```graphql
mutation UpdateOption {
  editCustomFieldOption(input: {
    id: "option_123"
    title: "Updated Priority"
    color: "#ff6b6b"
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

**កំណត់ចំណាំ**: ការលុបជម្រើសនឹងលុបវាពីកំណត់ត្រាទាំងអស់ដែលវាត្រូវបានជ្រើស។

### ការរៀបចំជម្រើស
```graphql
mutation ReorderOptions {
  reorderCustomFieldOptions(input: {
    customFieldId: "field_123"
    optionIds: ["option_1", "option_3", "option_2"]
  }) {
    id
    position
  }
}
```

## ច្បាប់ផ្ទៀងផ្ទាត់

### ការផ្ទៀងផ្ទាត់ជម្រើស
- ID ជម្រើសដែលផ្តល់ត្រូវតែមានស្រាប់
- ជម្រើសត្រូវតែជាការពាក់ព័ន្ធនឹងវាលកំណត់ដែលបានកំណត់
- តែជម្រើសតែមួយអាចត្រូវបានជ្រើស (បានអនុវត្តដោយស្វ័យប្រវត្តិ)
- តម្លៃ null/ទទេគឺមានសុពលភាព (គ្មានការជ្រើស)

### ការផ្ទៀងផ្ទាត់វាល
- ត្រូវតែមានជម្រើសយ៉ាងហោចណាស់មួយដែលបានកំណត់ដើម្បីអាចប្រើបាន
- ឈ្មោះជម្រើសត្រូវតែមានឯកត្តភាពក្នុងវាល
- កូដពណ៌ត្រូវតែមានទម្រង់ hex ដែលមានសុពលភាព (ប្រសិនបើផ្តល់)

## អាជ្ញាប័ណ្ណដែលត្រូវការ

| សកម្មភាព | អាជ្ញាប័ណ្ណដែលត្រូវការ |
|--------|-------------------|
| Create single-select field | Company role: `OWNER` or `ADMIN` |
| Update single-select field | Company role: `OWNER` or `ADMIN` |
| Add/edit options | Company role: `OWNER` or `ADMIN` |
| Set selected value | Any company role (`OWNER`, `ADMIN`, `MEMBER`, `CLIENT`) or custom project role with edit permission |
| View selected value | Standard record view permissions |

## ចម្លើយកំហុស

### ID ជម្រើសមិនត្រឹមត្រូវ
```json
{
  "errors": [{
    "message": "Custom field option was not found.",
    "extensions": {
      "code": "CUSTOM_FIELD_OPTION_NOT_FOUND"
    }
  }]
}
```

### ជម្រើសមិនពាក់ព័ន្ធនឹងវាល
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

### វាលមិនឃើញ
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

### មិនអាចបកប្រែតម្លៃ
```json
{
  "errors": [{
    "message": "Unable to parse custom field value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## អនុសាសន៍ល្អ

### ការរចនាជម្រើស
- ប្រើឈ្មោះជម្រើសដែលច្បាស់និងពិពណ៌នាដោយច្បាស់
- អនុវត្តកូដពណ៌ដែលមានន័យ
- រក្សាបញ្ជីជម្រើសឱ្យមានការយកចិត្តទុកដាក់និងសមស្រប
- រៀបចំជម្រើសដោយមានលំដាប់ (តាមអាទិភាព, អត្រា, ល។)

### លំនាំវាលស្ថានភាព
- ប្រើការងារស្ថានភាពដែលមានភាពស្របគ្នានៅក្នុងគម្រោង
- ពិចារណាពីការរីកចម្រើនធម្មជាតិរបស់ជម្រើស
- រួមបញ្ចូលស្ថានភាពចុងក្រោយដែលច្បាស់ (បានធ្វើ, បានបោះបង់, ល។)
- ប្រើពណ៌ដែលបង្ហាញពីន័យនៃជម្រើស

### ការគ្រប់គ្រងទិន្នន័យ
- ពិនិត្យនិងសម្អាតជម្រើសដែលមិនបានប្រើជាប្រចាំ
- ប្រើការកំណត់ឈ្មោះដែលមានភាពស្របគ្នា
- ពិចារណាអំពីផលប៉ះពាល់នៃការលុបជម្រើសលើកំណត់ត្រាដែលមានស្រាប់
- គ្រោងសម្រាប់ការអាប់ដេតនិងការផ្លាស់ប្តូរជម្រើស

## ករណីប្រើប្រាស់ទូទៅ

1. **ស្ថានភាពនិងការងារ**
   - ស្ថានភាពភារកិច្ច (ត្រូវធ្វើ, កំពុងដំណើរការ, បានធ្វើ)
   - ស្ថានភាពអនុម័ត (កំពុងរង់ចាំ, បានអនុម័ត, បានបដិសេធ)
   - ជំហានគម្រោង (កំពុងរៀបចំ, កំពុងអភិវឌ្ឍ, កំពុងសាកល្បង, បានចេញផ្សាយ)
   - ស្ថានភាពដោះស្រាយបញ្ហា

2. **ការបែងចែកនិងការបែងចែកប្រភេទ**
   - កម្រិតអាទិភាព (ទាប, មធ្យម, ខ្ពស់, សំខាន់)
   - ប្រភេទភារកិច្ច (កំហុស, លក្ខណៈពិសេស, ការកែលម្អ, ឯកសារ)
   - ប្រភេទគម្រោង (ក្នុងស្ថាប័ន, អតិថិជន, ស្រាវជ្រាវ)
   - ការបែងចែកនាយកដ្ឋាន

3. **គុណភាពនិងការវាយតម្លៃ**
   - ស្ថានភាពពិនិត្យ (មិនបានចាប់ផ្តើម, កំពុងពិនិត្យ, បានអនុម័ត)
   - កម្រិតគុណភាព (អាក្រក់, ធម្មតា, ល្អ, ល្អឥតខ្ចោះ)
   - កម្រិតហានិភ័យ (ទាប, មធ្យម, ខ្ពស់)
   - កម្រិតជំនឿ

4. **ការបែងចែកនិងការកាន់កាប់**
   - ការបែងចែកក្រុម
   - ការកាន់កាប់នាយកដ្ឋាន
   - ការបែងចែកដែលមានមូលដ្ឋានលើតួនាទី
   - ការបែងចែកតំបន់

## លក្ខណៈពិសេសនៃការបញ្ចូល

### ជាមួយស្វ័យប្រវត្តិ
- បង្កើតសកម្មភាពនៅពេលជម្រើសជាក់លាក់ត្រូវបានជ្រើស
- ផ្ទេរងារយោងទៅលើប្រភេទដែលបានជ្រើស
- ផ្ញើការជូនដំណឹងសម្រាប់ការផ្លាស់ប្តូរស្ថានភាព
- បង្កើតការងារដែលមានលក្ខខណ្ឌដោយផ្អែកលើការជ្រើស

### ជាមួយការស្វែងរក
- ការប្រមូលកំណត់ត្រាដោយប្រើជម្រើសដែលបានជ្រើស
- យោងទិន្នន័យជម្រើសពីកំណត់ត្រាផ្សេងទៀត
- បង្កើតរបាយការណ៍ដោយផ្អែកលើការជ្រើសជម្រើស
- ក្រុមកំណត់ត្រាដោយប្រើតម្លៃដែលបានជ្រើស

### ជាមួយបែបបទ
- ការត្រួតពិនិត្យបញ្ចូលជម្រើស
- អ៊ុតប៊ូតុងរ៉ាឌីយ៉ូ
- ការផ្ទៀងផ្ទាត់និងការប្រមូលជម្រើស
- ការបង្ហាញវាលដែលមានលក្ខខណ្ឌដោយផ្អែកលើការជ្រើស

## ការតាមដានសកម្មភាព

ការផ្លាស់ប្តូរវាលកំណត់តែមួយត្រូវបានតាមដានដោយស្វ័យប្រវត្តិ:
- បង្ហាញជម្រើសចាស់និងថ្មី
- បង្ហាញឈ្មោះជម្រើសនៅក្នុងកំណត់ហេតុសកម្មភាព
- ម៉ោងសម្រាប់ការផ្លាស់ប្តូរជម្រើសទាំងអស់
- ការបញ្ជាក់អ្នកប្រើសម្រាប់ការកែប្រែ

## ភាពខុសគ្នាពីការជ្រើសរើសច្រើន

| លក្ខណៈ | ការជ្រើសរើសតែមួយ | ការជ្រើសរើសច្រើន |
|---------|---------------|--------------|
| **Selection Limit** | Exactly 1 option | Multiple options |
| **Input Parameter** | `customFieldOptionId` | `customFieldOptionIds` |
| **Response Field** | `value` (single option object) | `value` (array of option objects) |
| **Storage Behavior** | Replaces existing selection | Adds to existing selections |
| **Common Use Cases** | Status, category, priority | Tags, skills, categories |

## ការដាក់កំណត់

- តែជម្រើសតែមួយអាចត្រូវបានជ្រើសនៅពេលណាមួយ
- គ្មានរចនាសម្ព័ន្ធជម្រើសជាន់ឬជាន់
- ជម្រើសត្រូវបានចែករំលែកនៅក្នុងកំណត់ត្រាទាំងអស់ដែលប្រើវាល
- គ្មានវិភាគជម្រើសឬការតាមដានការប្រើប្រាស់ដែលមានស្រាប់
- កូដពណ៌គឺសម្រាប់ការបង្ហាញតែប៉ុណ្ណោះ, មិនមានផលប៉ះពាល់ដំណើរការ
- មិនអាចកំណត់អាជ្ញាប័ណ្ណផ្សេងៗសម្រាប់ជម្រើសនីមួយៗ

## ឯកសារដែលពាក់ព័ន្ធ

- [វាលជ្រើសច្រើន](/api/custom-fields/select-multi) - សម្រាប់ការជ្រើសរើសជាច្រើន
- [វាលប្រអប់បញ្ជាក់](/api/custom-fields/checkbox) - សម្រាប់ជម្រើស boolean សាមញ្ញ
- [វាលអត្ថបទ](/api/custom-fields/text-single) - សម្រាប់ការបញ្ចូលអត្ថបទដោយឥតគិតថ្លៃ
- [ទិដ្ឋភាពទូទៅនៃវាលកំណត់](/api/custom-fields/1.index) - គំនិតទូទៅ