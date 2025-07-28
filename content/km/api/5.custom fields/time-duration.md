---
title: វាលកំណត់អាយុកាល
description: បង្កើតវាលអាយុកាលដែលគណនាដែលតាមដានពេលវេលារវាងព្រឹត្តិការណ៍ក្នុងការងាររបស់អ្នក
---

វាលកំណត់អាយុកាលអាចគណនានិងបង្ហាញអាយុកាលរវាងព្រឹត្តិការណ៍ពីរនៅក្នុងការងាររបស់អ្នក។ វាអាចប្រើបានសម្រាប់តាមដានពេលវេលាការប្រតិបត្តិ, ពេលវេលាចម្លើយ, ពេលវេលាច្រកឬមាត្រពេលវេលាផ្សេងៗនៅក្នុងគម្រោងរបស់អ្នក។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលអាយុកាលសាមញ្ញដែលតាមដានថាតើការងារបានចំណាយពេលប៉ុន្មានដើម្បីបញ្ចប់៖

```graphql
mutation CreateTimeDurationField {
  createCustomField(input: {
    name: "Processing Time"
    type: TIME_DURATION
    projectId: "proj_123"
    timeDurationDisplay: FULL_DATE_SUBSTRING
    timeDurationStartInput: {
      type: TODO_CREATED_AT
      condition: FIRST
    }
    timeDurationEndInput: {
      type: TODO_MARKED_AS_COMPLETE
      condition: FIRST
    }
  }) {
    id
    name
    type
    timeDurationDisplay
    timeDurationStart {
      type
      condition
    }
    timeDurationEnd {
      type
      condition
    }
  }
}
```

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលអាយុកាលស្មុគស្មាញដែលតាមដានពេលវេលារវាងការផ្លាស់ប្តូរវាលកំណត់ជាមួយគោលដៅ SLA៖

```graphql
mutation CreateAdvancedTimeDurationField {
  createCustomField(input: {
    name: "Review Cycle Time"
    type: TIME_DURATION
    projectId: "proj_123"
    description: "Time from review request to approval"
    timeDurationDisplay: FULL_DATE_STRING
    timeDurationTargetTime: 86400  # 24 hour SLA target
    timeDurationStartInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["review_requested_option_id"]
    }
    timeDurationEndInput: {
      type: TODO_CUSTOM_FIELD
      condition: FIRST
      customFieldId: "status_field_id"
      customFieldOptionIds: ["approved_option_id"]
    }
  }) {
    id
    name
    type
    description
    timeDurationDisplay
    timeDurationStart {
      type
      condition
      customField {
        name
      }
    }
    timeDurationEnd {
      type
      condition
      customField {
        name
      }
    }
  }
}
```

## ប៉ារ៉ាម៉ែត្រ Input

### CreateCustomFieldInput (TIME_DURATION)

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលអាយុកាល |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `TIME_DURATION` |
| `description` | String | ទេ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ បាទ | របៀបបង្ហាញអាយុកាល |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ បាទ | ការកំណត់ព្រឹត្តិការណ៍ចាប់ផ្តើម |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ បាទ | ការកំណត់ព្រឹត្តិការណ៍បញ្ចប់ |
| `timeDurationTargetTime` | Float | ទេ | គោលដៅអាយុកាលក្នុងវិនាទីសម្រាប់ការតាមដាន SLA |

### CustomFieldTimeDurationInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ បាទ | ប្រភេទព្រឹត្តិការណ៍ដើម្បីតាមដាន |
| `condition` | CustomFieldTimeDurationCondition! | ✅ បាទ | `FIRST` ឬ `LAST` ការកើតឡើង |
| `customFieldId` | String | Conditional | ត្រូវការសម្រាប់ `TODO_CUSTOM_FIELD` ប្រភេទ |
| `customFieldOptionIds` | [String!] | Conditional | ត្រូវការសម្រាប់ការផ្លាស់ប្តូរវាលជ្រើសរើស |
| `todoListId` | String | Conditional | ត្រូវការសម្រាប់ `TODO_MOVED` ប្រភេទ |
| `tagId` | String | Conditional | ត្រូវការសម្រាប់ `TODO_TAG_ADDED` ប្រភេទ |
| `assigneeId` | String | Conditional | ត្រូវការសម្រាប់ `TODO_ASSIGNEE_ADDED` ប្រភេទ |

### CustomFieldTimeDurationType Values

| តម្លៃ | ការពិពណ៌នា |
|-------|-------------|
| `TODO_CREATED_AT` | ពេលដែលកំណត់ត្រូវបានបង្កើត |
| `TODO_CUSTOM_FIELD` | ពេលដែលតម្លៃវាលកំណត់ផ្លាស់ប្តូរ |
| `TODO_DUE_DATE` | ពេលដែលកាលបរិច្ឆេទកំណត់ត្រូវបានកំណត់ |
| `TODO_MARKED_AS_COMPLETE` | ពេលដែលកំណត់ត្រូវបានសម្គាល់ថាបញ្ចប់ |
| `TODO_MOVED` | ពេលដែលកំណត់ត្រូវបានផ្លាស់ទីទៅបញ្ជីផ្សេងទៀត |
| `TODO_TAG_ADDED` | ពេលដែលស្លាកត្រូវបានបន្ថែមទៅកំណត់ |
| `TODO_ASSIGNEE_ADDED` | ពេលដែលអ្នកតែងតាំងត្រូវបានបន្ថែមទៅកំណត់ |

### CustomFieldTimeDurationCondition Values

| តម្លៃ | ការពិពណ៌នា |
|-------|-------------|
| `FIRST` | ប្រើការកើតឡើងដំបូងនៃព្រឹត្តិការណ៍ |
| `LAST` | ប្រើការកើតឡើងចុងក្រោយនៃព្រឹត្តិការណ៍ |

### CustomFieldTimeDurationDisplayType Values

| តម្លៃ | ការពិពណ៌នា | ឧទាហរណ៍ |
|-------|-------------|---------|
| `FULL_DATE` | រូបមន្តថ្ងៃ:ម៉ោង:នាទី:វិនាទី | `"01:02:03:04"` |
| `FULL_DATE_STRING` | សរសេរដោយពេញលេញ | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | លេខជាមួយឯកតា | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | អាយុកាលនៅក្នុងថ្ងៃតែប៉ុណ្ណោះ | `"2.5"` (2.5 days) |
| `HOURS` | អាយុកាលនៅក្នុងម៉ោងតែប៉ុណ្ណោះ | `"60"` (60 hours) |
| `MINUTES` | អាយុកាលនៅក្នុងនាទីតែប៉ុណ្ណោះ | `"3600"` (3600 minutes) |
| `SECONDS` | អាយុកាលនៅក្នុងវិនាទីតែប៉ុណ្ណោះ | `"216000"` (216000 seconds) |

## វាលចម្លើយ

### TodoCustomField Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកត្តសាស្ត្រសម្រាប់តម្លៃវាល |
| `customField` | CustomField! | ការកំណត់វាលកំណត់ |
| `number` | Float | អាយុកាលក្នុងវិនាទី |
| `value` | Float | ឈ្មោះសម្រាប់លេខ (អាយុកាលក្នុងវិនាទី) |
| `todo` | Todo! | កំណត់ដែលតម្លៃនេះជាប់ពាក់ |
| `createdAt` | DateTime! | ពេលដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលដែលតម្លៃត្រូវបានអាប់ដេតចុងក្រោយ |

### CustomField Response (TIME_DURATION)

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | របៀបបង្ហាញសម្រាប់អាយុកាល |
| `timeDurationStart` | CustomFieldTimeDuration | ការកំណត់ព្រឹត្តិការណ៍ចាប់ផ្តើម |
| `timeDurationEnd` | CustomFieldTimeDuration | ការកំណត់ព្រឹត្តិការណ៍បញ្ចប់ |
| `timeDurationTargetTime` | Float | គោលដៅអាយុកាលក្នុងវិនាទី (សម្រាប់ការតាមដាន SLA) |

## ការគណនាអាយុកាល

### របៀបដែលវាធ្វើការ
1. **ព្រឹត្តិការណ៍ចាប់ផ្តើម**: ប្រព័ន្ធតាមដានសម្រាប់ព្រឹត្តិការណ៍ចាប់ផ្តើមដែលបានកំណត់
2. **ព្រឹត្តិការណ៍បញ្ចប់**: ប្រព័ន្ធតាមដានសម្រាប់ព្រឹត្តិការណ៍បញ្ចប់ដែលបានកំណត់
3. **ការគណនា**: អាយុកាល = ពេលវេលាបញ្ចប់ - ពេលវេលាចាប់ផ្តើម
4. **ការផ្ទុក**: អាយុកាលត្រូវបានផ្ទុកនៅក្នុងវិនាទីជាលេខ
5. **ការបង្ហាញ**: ត្រូវបានរៀបចំតាមការកំណត់ `timeDurationDisplay`

### ការជំរុញអាប់ដេត
តម្លៃអាយុកាលត្រូវបានគណនាឡើងវិញដោយស្វ័យប្រវត្តិពេលដែល:
- កំណត់ត្រូវបានបង្កើតឬអាប់ដេត
- តម្លៃវាលកំណត់ផ្លាស់ប្តូរ
- ស្លាកត្រូវបានបន្ថែមឬយកចេញ
- អ្នកតែងតាំងត្រូវបានបន្ថែមឬយកចេញ
- កំណត់ត្រូវបានផ្លាស់ទីរវាងបញ្ជី
- កំណត់ត្រូវបានសម្គាល់ថាបញ្ចប់/មិនបានបញ្ចប់

## ការអានតម្លៃអាយុកាល

### សំណើរវាលអាយុកាល
```graphql
query GetTaskWithDuration {
  todo(id: "todo_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        timeDurationDisplay
      }
      number    # Duration in seconds
      value     # Same as number
    }
  }
}
```

### តម្លៃបង្ហាញដែលបានរៀបចំ
តម្លៃអាយុកាលត្រូវបានរៀបចំដោយស្វ័យប្រវត្តិដោយផ្អែកលើការកំណត់ `timeDurationDisplay`:

```javascript
// FULL_DATE format
93784 seconds → "01:02:03:04" (1 day, 2 hours, 3 minutes, 4 seconds)

// FULL_DATE_STRING format
7323 seconds → "Two hours, two minutes, three seconds"

// FULL_DATE_SUBSTRING format
3723 seconds → "1 hour, 2 minutes, 3 seconds"

// DAYS format
216000 seconds → "2.5" (2.5 days)

// HOURS format
7200 seconds → "2" (2 hours)

// MINUTES format
180 seconds → "3" (3 minutes)

// SECONDS format
3661 seconds → "3661" (raw seconds)
```

## ឧទាហរណ៍កំណត់រួមទូទៅ

### ពេលវេលាបញ្ចប់ការងារ
```graphql
timeDurationStartInput: {
  type: TODO_CREATED_AT
  condition: FIRST
}
timeDurationEndInput: {
  type: TODO_MARKED_AS_COMPLETE
  condition: FIRST
}
```

### ពេលវេលាប្រែប្រួលស្ថានភាព
```graphql
timeDurationStartInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["in_progress_option_id"]
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["completed_option_id"]
}
```

### ពេលវេលានៅក្នុងបញ្ជីជាក់លាក់
```graphql
timeDurationStartInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "review_list_id"
}
timeDurationEndInput: {
  type: TODO_MOVED
  condition: FIRST
  todoListId: "approved_list_id"
}
```

### ពេលវេលាចម្លើយនៃការបែងចែក
```graphql
timeDurationStartInput: {
  type: TODO_ASSIGNEE_ADDED
  condition: FIRST
  assigneeId: "user_123"
}
timeDurationEndInput: {
  type: TODO_CUSTOM_FIELD
  condition: FIRST
  customFieldId: "status_field_id"
  customFieldOptionIds: ["started_option_id"]
}
```

## អាជ្ញាប័ណ្ណដែលត្រូវការ

| សកម្មភាព | អាជ្ញាប័ណ្ណដែលត្រូវការ |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## ការឆ្លើយតបកំហុស

### ការកំណត់មិនត្រឹមត្រូវ
```json
{
  "errors": [{
    "message": "Custom field is required for TODO_CUSTOM_FIELD type",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### វាលដែលបានយោងមិនបានរកឃើញ
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

### ជម្រើសដែលត្រូវការមិនមាន
```json
{
  "errors": [{
    "message": "Custom field options are required for select field changes",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

## កំណត់ចំណាំសំខាន់

### ការគណនាដោយស្វ័យប្រវត្តិ
- វាលអាយុកាលគឺ **អានតែ** - តម្លៃត្រូវបានគណនាដោយស្វ័យប្រវត្តិ
- អ្នកមិនអាចកំណត់តម្លៃអាយុកាលដោយដៃតាម API
- ការគណនាដំណើរការនៅក្នុងការងារប្រព័ន្ធ
- តម្លៃត្រូវបានអាប់ដេតដោយស្វ័យប្រវត្តិពេលដែលមានព្រឹត្តិការណ៍ជំរុញកើតឡើង

### ការពិចារណាអំពីសមត្ថភាព
- ការគណនាអាយុកាលត្រូវបានចុះបញ្ជីនិងដំណើរការដោយស្វ័យប្រវត្តិ
- ចំនួនធំនៃវាលអាយុកាលអាចប៉ះពាល់ដល់សមត្ថភាព
- ពិចារណាពេលវេលានៃព្រឹត្តិការណ៍ជំរុញពេលដែលរចនាវាលអាយុកាល
- ប្រើលក្ខខណ្ឌជាក់លាក់ដើម្បីជៀសវាងការគណនាដោយស្វ័យប្រវត្តិដែលមិនចាំបាច់

### តម្លៃ Null
វាលអាយុកាលនឹងបង្ហាញ `null` នៅពេល:
- ព្រឹត្តិការណ៍ចាប់ផ្តើមមិនទាន់កើតឡើងទេ
- ព្រឹត្តិការណ៍បញ្ចប់មិនទាន់កើតឡើងទេ
- ការកំណត់យោងទៅកាន់អង្គភាពដែលមិនមាន
- ការគណនាប្រឈមមុខនឹងកំហុស

## អនុវត្តន៍ល្អបំផុត

### ការរចនាការកំណត់
- ប្រើប្រភេទព្រឹត្តិការណ៍ជាក់លាក់ជាងប្រភេទទូទៅនៅពេលដែលអាចធ្វើទៅបាន
- ជ្រើសរើស `FIRST` បើប្រៀបធៀបជាមួយ `LAST` លក្ខខណ្ឌផ្អែកលើការងាររបស់អ្នក
- សាកល្បងការគណនាអាយុកាលជាមួយទិន្នន័យគំរូមុនពេលបញ្ចូលប្រព័ន្ធ
- ឯកសារពីយុទ្ធសាស្ត្រវាលអាយុកាលរបស់អ្នកសម្រាប់សមាជិកក្រុម

### ការបង្ហាញការរៀបចំ
- ប្រើ `FULL_DATE_SUBSTRING` សម្រាប់ទ្រង់ទ្រាយដែលអាចអានបានយ៉ាងល្អបំផុត
- ប្រើ `FULL_DATE` សម្រាប់ការបង្ហាញទំហំស្រួលនិងស្រប
- ប្រើ `FULL_DATE_STRING` សម្រាប់របាយការណ៍និងឯកសារផ្លូវការនានា
- ប្រើ `DAYS`, `HOURS`, `MINUTES`, ឬ `SECONDS` សម្រាប់ការបង្ហាញលេខសាមញ្ញ
- ពិចារណាពីកំណត់កន្លែង UI របស់អ្នកពេលជ្រើសរើសទ្រង់ទ្រាយ

### ការតាមដាន SLA ជាមួយពេលវេលាគោលដៅ
ពេលប្រើ `timeDurationTargetTime`:
- កំណត់អាយុកាលគោលដៅនៅក្នុងវិនាទី
- ប្រៀបធៀបអាយុកាលពិតប្រាកដនឹងគោលដៅសម្រាប់ការអនុវត្ត SLA
- ប្រើនៅក្នុងផ្ទាំងព័ត៌មានដើម្បីបង្ហាញអត្ថបទដែលលើសកាលកំណត់
- ឧទាហរណ៍: SLA ចម្លើយ 24 ម៉ោង = 86400 វិនាទី

### ការរួមបញ្ចូលក្នុងការងារ
- រចនាវាលអាយុកាលឱ្យសមស្របនឹងដំណើរការអាជីវកម្មរបស់អ្នក
- ប្រើទិន្នន័យអាយុកាលសម្រាប់ការកែលម្អនិងបង្កើនប្រសិទ្ធភាពដំណើរការ
- តាមដាននិន្នាការអាយុកាលដើម្បីកំណត់ចំណុចបញ្ឈប់ក្នុងការងារ
- កំណត់ការជូនដំណឹងសម្រាប់កម្រិតអាយុកាលប្រសិនបើត្រូវការ

## ករណីប្រើទូទៅ

1. **ការសមត្ថភាពដំណើរការ**
   - ពេលវេលាបញ្ចប់ការងារ
   - ពេលវេលាច្រកពិនិត្យ
   - ពេលវេលាការអនុម័ត
   - ពេលវេលាចម្លើយ

2. **ការតាមដាន SLA**
   - ពេលវេលាចម្លើយដំបូង
   - ពេលវេលាដោះស្រាយ
   - ពេលវេលាឡើងកម្រិត
   - ការអនុវត្តកម្រិតសេវាកម្ម

3. **វិភាគការងារ**
   - ការកំណត់ចំណុចបញ្ឈប់
   - ការកែលម្អដំណើរការ
   - ម៉ាត្រនៃសមត្ថភាពក្រុម
   - ពេលវេលាធានាគុណភាព

4. **ការគ្រប់គ្រងគម្រោង**
   - អាយុកាលដំណាក់កាល
   - ពេលវេលាសម្គាល់
   - ពេលវេលាការបែងចែកធនធាន
   - ពេលវេលាដឹកជញ្ជូន

## ការបង្ហាញ

- វាលអាយុកាលគឺ **អានតែ** ហើយមិនអាចកំណត់ដោយដៃបានទេ
- តម្លៃត្រូវបានគណនាដោយស្វ័យប្រវត្តិហើយអាចមិនមាននៅក្នុងពេលវេលាបច្ចុប្បន្ន
- ត្រូវការការជំរុញព្រឹត្តិការណ៍ដែលត្រឹមត្រូវដើម្បីកំណត់នៅក្នុងការងាររបស់អ្នក
- មិនអាចគណនាអាយុកាលសម្រាប់ព្រឹត្តិការណ៍ដែលមិនទាន់កើតឡើងទេ
- មានកំណត់ក្នុងការតាមដានពេលវេលារវាងព្រឹត្តិការណ៍ដែលច្បាស់ (មិនមែនការតាមដានពេលវេលាបន្ត)
- មិនមានការជូនដំណឹងឬការជូនដំណឹង SLA ដែលបានបង្កើតឡើង
- មិនអាចប្រមូលការគណនាអាយុកាលច្រើនទៅក្នុងវាលតែមួយបាន

## ឯកសារដែលពាក់ព័ន្ធ

- [វាលលេខ](/api/custom-fields/number) - សម្រាប់តម្លៃលេខដោយដៃ
- [វាលកាលបរិច្ឆេទ](/api/custom-fields/date) - សម្រាប់ការតាមដានកាលបរិច្ឆេទជាក់លាក់
- [ទិដ្ឋភាពទូទៅនៃវាលកំណត់](/api/custom-fields/list-custom-fields) - គំនិតទូទៅ
- [Automations](/api/automations) - សម្រាប់ការជំរុញសកម្មភាពដោយផ្អែកលើកម្រិតអាយុកាល