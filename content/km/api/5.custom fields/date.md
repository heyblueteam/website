---
title: វាលកាលបរិច្ឆេទផ្ទាល់ខ្លួន
description: បង្កើតវាលកាលបរិច្ឆេទដើម្បីតាមដានកាលបរិច្ឆេទតែមួយឬចន្លោះកាលបរិច្ឆេទជាមួយការគាំទ្រពេលវេលា
---

វាលកាលបរិច្ឆេទផ្ទាល់ខ្លួនអនុញ្ញាតឱ្យអ្នករក្សាទុកកាលបរិច្ឆេទតែមួយឬចន្លោះកាលបរិច្ឆេទសម្រាប់កំណត់ត្រា។ វាគាំទ្រការដោះសោពេលវេលា ការបង្ហាញយ៉ាងឆ្លាតវៃ ហើយអាចប្រើដើម្បីតាមដានកាលកំណត់ កាលបរិច្ឆេទព្រឹត្តិការណ៍ ឬព័ត៌មានដែលមានមូលដ្ឋានលើពេលវេលាណាមួយ។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលកាលបរិច្ឆេទសាមញ្ញ៖

```graphql
mutation CreateDateField {
  createCustomField(input: {
    name: "Deadline"
    type: DATE
  }) {
    id
    name
    type
  }
}
```

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលកាលបរិច្ឆេទកំណត់ជាមួយការពិពណ៌នា៖

```graphql
mutation CreateDueDateField {
  createCustomField(input: {
    name: "Contract Expiration"
    type: DATE
    isDueDate: true
    description: "When the contract expires and needs renewal"
  }) {
    id
    name
    type
    isDueDate
    description
  }
}
```

## ប៉ារ៉ាម៉ែត្រនាំចូល

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលកាលបរិច្ឆេទ |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `DATE` |
| `isDueDate` | Boolean | ទេ | ថាតើវាលនេះតំណាងឱ្យកាលបរិច្ឆេទកំណត់ |
| `description` | String | ទេ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |

**កំណត់ចំណាំ**: វាលផ្ទាល់ខ្លួនត្រូវបានភ្ជាប់ដោយស្វ័យប្រវត្តិជាមួយគម្រោងមួយដោយផ្អែកលើបរិបទគម្រោងបច្ចុប្បន្នរបស់អ្នកប្រើ។ មិនត្រូវការប៉ារ៉ាម៉ែត្រ `projectId` ទេ។

## ការកំណត់តម្លៃកាលបរិច្ឆេទ

វាលកាលបរិច្ឆេទអាចរក្សាទុកទាំងកាលបរិច្ឆេទតែមួយឬចន្លោះកាលបរិច្ឆេទ៖

### កាលបរិច្ឆេទតែមួយ

```graphql
mutation SetSingleDate {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T10:00:00Z"
    endDate: "2025-01-15T10:00:00Z"
    timezone: "America/New_York"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### ចន្លោះកាលបរិច្ឆេទ

```graphql
mutation SetDateRange {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-01T09:00:00Z"
    endDate: "2025-01-31T17:00:00Z"
    timezone: "Europe/London"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### ព្រឹត្តិការណ៍ទាំងថ្ងៃ

```graphql
mutation SetAllDayEvent {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    startDate: "2025-01-15T00:00:00Z"
    endDate: "2025-01-15T23:59:59Z"
    timezone: "Asia/Tokyo"
  }) {
    id
    customField {
      value  # Contains { startDate, endDate, timezone }
    }
  }
}
```

### ប៉ារ៉ាម៉ែត្រ SetTodoCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ បាទ | ID នៃកំណត់ត្រាដែលត្រូវកែប្រែ |
| `customFieldId` | String! | ✅ បាទ | ID នៃវាលកាលបរិច្ឆេទផ្ទាល់ខ្លួន |
| `startDate` | DateTime | ទេ | ថ្ងៃ/ម៉ោងចាប់ផ្តើមក្នុងទ្រង់ទ្រាយ ISO 8601 |
| `endDate` | DateTime | ទេ | ថ្ងៃ/ម៉ោងបញ្ចប់ក្នុងទ្រង់ទ្រាយ ISO 8601 |
| `timezone` | String | ទេ | អត្តសញ្ញាណពេលវេលា (ឧ. "America/New_York") |

**កំណត់ចំណាំ**: ប្រសិនបើមានតែ `startDate` ត្រូវបានផ្តល់ជូន, `endDate` នឹងត្រូវបានកំណត់ជាដំណាក់កាលដូចគ្នា។

## ទ្រង់ទ្រាយកាលបរិច្ឆេទ

### ទ្រង់ទ្រាយ ISO 8601
កាលបរិច្ឆេទទាំងអស់ត្រូវតែផ្តល់ជូនក្នុងទ្រង់ទ្រាយ ISO 8601:
- `2025-01-15T14:30:00Z` - ពេលវេលា UTC
- `2025-01-15T14:30:00+05:00` - ជាមួយការបន្ថែមពេលវេលា
- `2025-01-15T14:30:00.123Z` - ជាមួយមីលីសេកុន

### អត្តសញ្ញាណពេលវេលា
ប្រើអត្តសញ្ញាណពេលវេលាមានស្តង់ដា:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

ប្រសិនបើមិនមានពេលវេលាដែលបានផ្តល់, ប្រព័ន្ធនឹងត្រឡប់ទៅកាន់ពេលវេលាដែលបានរកឃើញរបស់អ្នកប្រើ។

## ការបង្កើតកំណត់ត្រាជាមួយតម្លៃកាលបរិច្ឆេទ

នៅពេលបង្កើតកំណត់ត្រាថ្មីជាមួយតម្លៃកាលបរិច្ឆេទ៖

```graphql
mutation CreateRecordWithDate {
  createTodo(input: {
    title: "Project Milestone"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "date_field_id"
      value: "2025-02-15"  # Simple date format
    }]
  }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # Date values are accessed here
      }
    }
  }
}
```

### ទ្រង់ទ្រាយនាំចូលដែលគាំទ្រ

នៅពេលបង្កើតកំណត់ត្រា, កាលបរិច្ឆេទអាចផ្តល់ជូនក្នុងទ្រង់ទ្រាយផ្សេងៗ៖

| ទ្រង់ទ្រាយ | ឧទាហរណ៍ | លទ្ធផល |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## វាលឆ្លើយតប

### TodoCustomField ឆ្លើយតប

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | ID! | អត្តសញ្ញាណឯកតាសម្រាប់តម្លៃវាល |
| `uid` | String! | ស្រទាប់អត្តសញ្ញាណឯកតា |
| `customField` | CustomField! | ការកំណត់វាលផ្ទាល់ខ្លួន (មានតម្លៃកាលបរិច្ឆេទ) |
| `todo` | Todo! | កំណត់ត្រាដែលតម្លៃនេះស្ថិតនៅក្នុង |
| `createdAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានកែប្រែចុងក្រោយ |

**សំខាន់**: តម្លៃកាលបរិច្ឆេទ (`startDate`, `endDate`, `timezone`) ត្រូវបានចូលដំណើរការតាមរយៈវាល `customField.value`, មិនមែនដោយផ្ទាល់លើ TodoCustomField ទេ។

### រចនាសម្ព័ន្ធវាលតម្លៃ

តម្លៃកាលបរិច្ឆេទត្រូវបានត្រឡប់តាមរយៈវាល `customField.value` ជារូបវន្ត JSON:

```json
{
  "customField": {
    "value": {
      "startDate": "2025-01-15T10:00:00.000Z",
      "endDate": "2025-01-15T17:00:00.000Z",
      "timezone": "America/New_York"
    }
  }
}
```

**កំណត់ចំណាំ**: វាល `value` ស្ថិតនៅលើប្រភេទ `CustomField`, មិនមែននៅលើ `TodoCustomField`។

## ការស្វែងរកតម្លៃកាលបរិច្ឆេទ

នៅពេលស្វែងរកកំណត់ត្រាជាមួយវាលកាលបរិច្ឆេទផ្ទាល់ខ្លួន, ចូលដំណើរការតម្លៃកាលបរិច្ឆេទតាមរយៈវាល `customField.value`:

```graphql
query GetRecordWithDateField {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        value  # For DATE type, contains { startDate, endDate, timezone }
      }
    }
  }
}
```

ការឆ្លើយតបនឹងរួមបញ្ចូលតម្លៃកាលបរិច្ឆេទនៅក្នុងវាល `value`:

```json
{
  "data": {
    "todo": {
      "customFields": [{
        "customField": {
          "name": "Deadline",
          "type": "DATE",
          "value": {
            "startDate": "2025-01-15T10:00:00.000Z",
            "endDate": "2025-01-15T10:00:00.000Z",
            "timezone": "America/New_York"
          }
        }
      }]
    }
  }
}
```

## បញ្ញាណបង្ហាញកាលបរិច្ឆេទ

ប្រព័ន្ធធ្វើការបង្ហាញកាលបរិច្ឆេទដោយស្វ័យប្រវត្តិផ្អែកលើចន្លោះ៖

| ស្ថានភាព | ទ្រង់ទ្រាយបង្ហាញ |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (មិនបង្ហាញពេលវេលា) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**ការស្គាល់ទាំងថ្ងៃ**: ព្រឹត្តិការណ៍ពីម៉ោង 00:00 ដល់ 23:59 ត្រូវបានស្គាល់ដោយស្វ័យប្រវត្តិថាជាព្រឹត្តិការណ៍ទាំងថ្ងៃ។

## ការដោះសោពេលវេលា

### ការផ្ទុក
- កាលបរិច្ឆេទទាំងអស់ត្រូវបានផ្ទុកនៅក្នុង UTC ក្នុងមូលដ្ឋានទិន្នន័យ
- ព័ត៌មានពេលវេលាត្រូវបានរក្សាទុកដោយឡែក
- ការបម្លែងកើតឡើងនៅពេលបង្ហាញ

### អនុសាសន៍ល្អបំផុត
- តែងតែផ្តល់ពេលវេលាសម្រាប់ភាពត្រឹមត្រូវ
- ប្រើពេលវេលាដូចគ្នានៅក្នុងគម្រោង
- ពិចារណាទីតាំងអ្នកប្រើសម្រាប់ក្រុមអន្តរជាតិ

### ពេលវេលាធម្មតា

| តំបន់ | អត្តសញ្ញាណពេលវេលា | ការបន្ថែម UTC |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## ការតម្រងនិងការស្វែងរក

វាលកាលបរិច្ឆេទគាំទ្រការតម្រងស្មុគស្មាញ៖

```graphql
query FilterByDateRange {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      dateRange: {
        startDate: "2025-01-01T00:00:00Z"
        endDate: "2025-12-31T23:59:59Z"
      }
      operator: EQ  # Returns todos whose dates overlap with this range
    }]
  }) {
    id
    title
  }
}
```

### ការត្រួតពិនិត្យវាលកាលបរិច្ឆេទទទេ

```graphql
query FilterEmptyDates {
  todos(filter: {
    customFields: [{
      customFieldId: "date_field_id"
      values: null
      operator: IS  # Returns todos with no date set
    }]
  }) {
    id
    title
  }
}
```

### ប្រើប្រាស់ប្រភេទដែលគាំទ្រ

| ប្រភេទ | ការប្រើប្រាស់ | ការពិពណ៌នា |
|----------|-------|-------------|
| `EQ` | ជាមួយ dateRange | កាលបរិច្ឆេទមានការប្រហែលជាមួយចន្លោះដែលបានកំណត់ (ការប្រហែលណាមួយ) |
| `NE` | ជាមួយ dateRange | កាលបរិច្ឆេទមិនមានការប្រហែលជាមួយចន្លោះ |
| `IS` | ជាមួយ `values: null` | វាលកាលបរិច្ឆេទទទេ (startDate ឬ endDate គឺ null) |
| `NOT` | ជាមួយ `values: null` | វាលកាលបរិច្ឆេទមានតម្លៃ (ទាំងពីរទីកាលមិនមែនជារូបភាព null) |

## អាជ្ញាប័ណ្ណដែលត្រូវការ

| សកម្មភាព | អាជ្ញាប័ណ្ណដែលត្រូវការ |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## ការឆ្លើយតបកំហុស

### ទ្រង់ទ្រាយកាលបរិច្ឆេទមិនត្រឹមត្រូវ
```json
{
  "errors": [{
    "message": "Invalid date format. Use ISO 8601 format",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### វាលមិនបានរកឃើញ
```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "NOT_FOUND"
    }
  }]
}
```


## ការកំណត់

- មិនគាំទ្រកាលបរិច្ឆេទដែលកើតឡើងឡើយ (ប្រើការបង្កើតស្វ័យប្រវត្តិសម្រាប់ព្រឹត្តិការណ៍ដែលកើតឡើងឡើយ)
- មិនអាចកំណត់ពេលវេលាបានដោយគ្មានកាលបរិច្ឆេទ
- មិនមានការគណនាថ្ងៃធ្វើការដែលបានបង្កើតឡើង
- ចន្លោះកាលបរិច្ឆេទមិនត្រូវបានផ្ទៀងផ្ទាត់ end > start ដោយស្វ័យប្រវត្តិ
- កម្រិតអតិបរមាគឺដល់វិនាទី (មិនមានការផ្ទុកមីលីសេកុន)

## ធនធានដែលពាក់ព័ន្ធ

- [ទិដ្ឋភាពទូទៅអំពីវាលផ្ទាល់ខ្លួន](/api/custom-fields/list-custom-fields) - គំនិតទូទៅអំពីវាលផ្ទាល់ខ្លួន
- [API ស្វ័យប្រវត្តិ](/api/automations/index) - បង្កើតស្វ័យប្រវត្តិដែលមានមូលដ្ឋានលើកាលបរិច្ឆេទ