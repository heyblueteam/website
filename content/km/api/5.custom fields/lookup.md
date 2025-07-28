---
title: ការស្វែងរកវាលកំណត់ផ្ទាល់ខ្លួន
description: បង្កើតវាលស្វែងរកដែលទាញយកទិន្នន័យដោយស្វ័យប្រវត្តិពីកំណត់ត្រាដែលបានយោង
---

វាលកំណត់ផ្ទាល់ខ្លួនស្វែងរកដោយស្វ័យប្រវត្តិទាញយកទិន្នន័យពីកំណត់ត្រាដែលបានយោងដោយ [វាលយោង](/api/custom-fields/reference) បង្ហាញព័ត៌មានពីកំណត់ត្រាដែលភ្ជាប់គ្នាโดยគ្មានការចម្លងដោយដៃ។ ពួកវានឹងធ្វើការអាប់ដេតដោយស្វ័យប្រវត្តិពេលដែលទិន្នន័យដែលបានយោងផ្លាស់ប្តូរ។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលស្វែងរកដើម្បីបង្ហាញតាក់សពីកំណត់ត្រាដែលបានយោង:

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

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលស្វែងរកដើម្បីទាញយកតម្លៃវាលកំណត់ផ្ទាល់ខ្លួនពីកំណត់ត្រាដែលបានយោង:

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

## ប៉ារ៉ាម៉ែត្រ​ចូល

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលស្វែងរក |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ បាទ | ការកំណត់ស្វែងរក |
| `description` | String | មិន | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |

## ការកំណត់ស្វែងរក

### CustomFieldLookupOptionInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `referenceId` | String! | ✅ បាទ | ID នៃវាលយោងដើម្បីទាញយកទិន្នន័យ |
| `lookupId` | String | មិន | ID នៃវាលកំណត់ផ្ទាល់ខ្លួនជាក់លាក់ដែលត្រូវស្វែងរក (ត្រូវការសម្រាប់ប្រភេទ TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ បាទ | ប្រភេទទិន្នន័យដើម្បីទាញយកពីកំណត់ត្រាដែលបានយោង |

## ប្រភេទស្វែងរក

### CustomFieldLookupType Values

| ប្រភេទ | ការពិពណ៌នា | ត្រឡប់ |
|------|-------------|---------|
| `TODO_DUE_DATE` | ថ្ងៃកំណត់ដែលបានយោងពី TODO | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | ថ្ងៃបង្កើតពី TODO ដែលបានយោង | Array of creation timestamps |
| `TODO_UPDATED_AT` | ថ្ងៃដែលបានអាប់ដេតចុងក្រោយពី TODO ដែលបានយោង | Array of update timestamps |
| `TODO_TAG` | តាក់សពី TODO ដែលបានយោង | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | អ្នកចាត់ការពី TODO ដែលបានយោង | Array of user objects |
| `TODO_DESCRIPTION` | ការពិពណ៌នាពី TODO ដែលបានយោង | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | ឈ្មោះបញ្ជី TODO ពី TODO ដែលបានយោង | Array of list titles |
| `TODO_CUSTOM_FIELD` | តម្លៃវាលកំណត់ផ្ទាល់ខ្លួនពី TODO ដែលបានយោង | Array of values based on the field type |

## វាលឆ្លើយតប

### CustomField Response (សម្រាប់វាលស្វែងរក)

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកតាសម្រាប់វាល |
| `name` | String! | ឈ្មោះបង្ហាញនៃវាលស្វែងរក |
| `type` | CustomFieldType! | នឹងជា `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | ការកំណត់ស្វែងរក និងលទ្ធផល |
| `createdAt` | DateTime! | ពេលវេលាដែលបានបង្កើតវាល |
| `updatedAt` | DateTime! | ពេលវេលាដែលបានអាប់ដេតចុងក្រោយវាល |

### CustomFieldLookupOption Structure

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | ប្រភេទស្វែងរកដែលកំពុងអនុវត្ត |
| `lookupResult` | JSON | ទិន្នន័យដែលបានទាញយកពីកំណត់ត្រាដែលបានយោង |
| `reference` | CustomField | វាលយោងដែលកំពុងប្រើជាភាគី |
| `lookup` | CustomField | វាលជាក់លាក់ដែលត្រូវបានស្វែងរក (សម្រាប់ TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | វាលស្វែងរកមាត្រ |
| `parentLookup` | CustomField | វាលស្វែងរកមាត្រ ក្នុងខ្សែ (សម្រាប់ស្វែងរកដែលមានរបៀបបន្ថែម) |

## របៀបដែលស្វែងរកធ្វើការ

1. **ការទាញយកទិន្នន័យ**: ស្វែងរកទាញយកទិន្នន័យជាក់លាក់ពីកំណត់ត្រាទាំងអស់ដែលភ្ជាប់តាមរយៈវាលយោង
2. **ការអាប់ដេតដោយស្វ័យប្រវត្តិ**: ពេលដែលកំណត់ត្រាដែលបានយោងផ្លាស់ប្តូរ តម្លៃស្វែងរកនឹងអាប់ដេតដោយស្វ័យប្រវត្តិ
3. **អានតែប៉ុណ្ណោះ**: វាលស្វែងរកមិនអាចកែប្រែបានដោយផ្ទាល់ - ពួកវានឹងតែងតែបង្ហាញទិន្នន័យដែលបានយោងបច្ចុប្បន្ន
4. **គ្មានការគណនា**: ស្វែងរកទាញយក និងបង្ហាញទិន្នន័យដូចដែលវា មានដោយគ្មានការបូកបញ្ចូល ឬការគណនា

## TODO_CUSTOM_FIELD ស្វែងរក

ពេលដែលប្រើ `TODO_CUSTOM_FIELD` ប្រភេទ អ្នកត្រូវតែបញ្ជាក់ថាតើវាលកំណត់ផ្ទាល់ខ្លួនណាដែលត្រូវទាញយកដោយប្រើ `lookupId` ប៉ារ៉ាម៉ែត្រ:

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

នេះនឹងទាញយកតម្លៃនៃវាលកំណត់ផ្ទាល់ខ្លួនដែលបានបញ្ជាក់ពីកំណត់ត្រាដែលបានយោងទាំងអស់។

## ការស្វែងរកទិន្នន័យស្វែងរក

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

## លទ្ធផលស្វែងរកឧទាហរណ៍

### លទ្ធផលស្វែងរកតាក់ស
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

### លទ្ធផលស្វែងរកអ្នកចាត់ការ
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

### លទ្ធផលស្វែងរកវាលកំណត់ផ្ទាល់ខ្លួន
លទ្ធផលអាចខុសគ្នាដោយផ្អែកលើប្រភេទវាលកំណត់ផ្ទាល់ខ្លួនដែលកំពុងស្វែងរក។ ឧទាហរណ៍ វាលស្វែងរករូបិយវត្ថុអាចត្រឡប់មកវិញ៖
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

## អាជ្ញាប័ណ្ណដែលត្រូវការ

| សកម្មភាព | អាជ្ញាប័ណ្ណដែលត្រូវការ |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**សំខាន់**: អ្នកប្រើត្រូវតែមានអាជ្ញាប័ណ្ណមើលលើគម្រោងបច្ចុប្បន្ន និងគម្រោងដែលបានយោងដើម្បីឃើញលទ្ធផលស្វែងរក។

## ការឆ្លើយតបកំហុស

### វាលយោងមិនត្រឹមត្រូវ
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

### ការស្វែងរកច្រកចូលត្រូវបានរកឃើញ
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

### គ្មាន ID ស្វែងរកសម្រាប់ TODO_CUSTOM_FIELD
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

## អនុសាសន៍ល្អ

1. **ឈ្មោះច្បាស់**: ប្រើឈ្មោះដែលពិពណ៌នាដែលបង្ហាញថាទិន្នន័យអ្វីដែលកំពុងស្វែងរក
2. **ប្រភេទសមស្រប**: ជ្រើសប្រភេទស្វែងរកដែលសមស្របនឹងតម្រូវការទិន្នន័យរបស់អ្នក
3. **ការអនុវត្តន៍**: ស្វែងរកដំណើរការទាំងអស់នៃកំណត់ត្រាដែលបានយោង ដូច្នេះត្រូវប្រុងប្រយ័ត្នចំពោះវាលយោងដែលមានតំណភ្ជាប់ច្រើន
4. **អាជ្ញាប័ណ្ណ**: ធានាថាអ្នកប្រើមានការចូលដំណើរការទៅកាន់គម្រោងដែលបានយោងសម្រាប់ស្វែងរកដើម្បីធ្វើការងារ

## ករណីប្រើប្រាស់ទូទៅ

### ការមើលឃើញឆ្លងគម្រោង
បង្ហាញតាក់ស អ្នកចាត់ការ ឬស្ថានភាពពីគម្រោងដែលពាក់ព័ន្ធដោយគ្មានការសម្របសម្រួលដោយដៃ។

### ការតាមដានការពឹងផ្អែក
បង្ហាញថ្ងៃកំណត់ឬស្ថានភាពការបញ្ចប់នៃភារកិច្ចដែលការងារបច្ចុប្បន្នពឹងផ្អែក។

### ទិដ្ឋភាពធនធាន
បង្ហាញសមាជិកក្រុមទាំងអស់ដែលត្រូវបានចាត់តាំងទៅកាន់ភារកិច្ចដែលបានយោងសម្រាប់ការធ្វើផែនការធនធាន។

### ការប្រមូលស្ថានភាព
ប្រមូលស្ថានភាពដ៏ឯកតាពីភារកិច្ចដែលពាក់ព័ន្ធដើម្បីឃើញសុខភាពគម្រោងនៅក្នុងមួយភ្លែត។

## ការដាក់កំណត់

- វាលស្វែងរកគឺអានតែប៉ុណ្ណោះ ហើយមិនអាចកែប្រែបានដោយផ្ទាល់
- គ្មានមុខងារបូកបញ្ចូល (SUM, COUNT, AVG) - ស្វែងរកតែទាញយកទិន្នន័យ
- គ្មានជម្រើសតម្រង - កំណត់ត្រាទាំងអស់ដែលបានយោងត្រូវបានរួមបញ្ចូល
- ខ្សែស្វែងរកច្រកចូលត្រូវបានរារាំងដើម្បីជៀសវាងការបង្វិលអស់កម្លាំង
- លទ្ធផលបង្ហាញទិន្នន័យបច្ចុប្បន្ន និងអាប់ដេតដោយស្វ័យប្រវត្តិ

## ធនធានដែលពាក់ព័ន្ធ

- [វាលយោង](/api/custom-fields/reference) - បង្កើតតំណទៅកាន់កំណត់ត្រាសម្រាប់ប្រភពស្វែងរក
- [តម្លៃវាលកំណត់ផ្ទាល់ខ្លួន](/api/custom-fields/custom-field-values) - កំណត់តម្លៃលើវាលកំណត់ផ្ទាល់ខ្លួនដែលអាចកែប្រែបាន
- [បញ្ជីវាលកំណត់ផ្ទាល់ខ្លួន](/api/custom-fields/list-custom-fields) - ស្វែងរកវាលកំណត់ផ្ទាល់ខ្លួនទាំងអស់ក្នុងគម្រោង