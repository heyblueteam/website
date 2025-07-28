---
title: វាលបន្ថែមប្រទេស
description: បង្កើតវាលជ្រើសរើសប្រទេសជាមួយនឹងការត្រួតពិនិត្យកូដប្រទេស ISO
---

វាលបន្ថែមប្រទេសអនុញ្ញាតឱ្យអ្នករក្សាទុកនិងគ្រប់គ្រងព័ត៌មានប្រទេសសម្រាប់កំណត់ត្រា។ វាលនេះគាំទ្រដល់ឈ្មោះប្រទេសនិងកូដប្រទេស ISO Alpha-2។

**សំខាន់**: ការត្រួតពិនិត្យនិងអាកប្បកិរិយាការបម្លែងប្រទេសមានភាពខុសគ្នាខ្លាំងនៅក្នុងការប្រែប្រួល:
- **createTodo**: បញ្ជាក់និងបម្លែងឈ្មោះប្រទេសទៅកូដ ISO យ៉ាងស្វ័យប្រវត្តិ
- **setTodoCustomField**: ទទួលយកតម្លៃណាមួយដោយគ្មានការត្រួតពិនិត្យ

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលប្រទេសសាមញ្ញ:

```graphql
mutation CreateCountryField {
  createCustomField(input: {
    name: "Country of Origin"
    type: COUNTRY
    projectId: "proj_123"
  }) {
    id
    name
    type
  }
}
```

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលប្រទេសជាមួយនឹងការពិពណ៌នា:

```graphql
mutation CreateDetailedCountryField {
  createCustomField(input: {
    name: "Customer Location"
    type: COUNTRY
    projectId: "proj_123"
    description: "Primary country where the customer is located"
    isActive: true
  }) {
    id
    name
    type
    description
    isActive
  }
}
```

## ប៉ារ៉ាម៉ែត្រ​ចូល

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលប្រទេស |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវមាន `COUNTRY` |
| `description` | String | មិនត្រូវការ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |

**កំណត់ចំណាំ**: `projectId` មិនត្រូវបានផ្ញើក្នុងការចូលទេ ប៉ុន្តែត្រូវបានកំណត់ដោយបរិបទ GraphQL (ធម្មតាពីក្បាលសំណើឬការផ្ទៀងផ្ទាត់)។

## ការកំណត់តម្លៃប្រទេស

វាលប្រទេសរក្សាទុកទិន្នន័យក្នុងវាលមូលដ្ឋានពីរនៅក្នុងមូលដ្ឋានទិន្នន័យ:
- **`countryCodes`**: រក្សាទុកកូដប្រទេស ISO Alpha-2 ជាស្រទាប់ខ្សែអក្សរដែលបំបែកដោយសញ្ញាកម្មា (ត្រូវបានត្រឡប់ជាអារ៉ាយតាម API)
- **`text`**: រក្សាទុកអត្ថបទបង្ហាញឬឈ្មោះប្រទេសជាខ្សែអក្សរ

### ការយល់ដឹងអំពីប៉ារ៉ាម៉ែត្រ

ការប្រែប្រួល `setTodoCustomField` ទទួលយកប៉ារ៉ាម៉ែត្រជាជម្រើសពីរសម្រាប់វាលប្រទេស:

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា | អ្វីដែលវាធ្វើ |
|-----------|------|----------|-------------|--------------|
| `todoId` | String! | ✅ បាទ | ID នៃកំណត់ត្រាដែលត្រូវកែប្រែ | - |
| `customFieldId` | String! | ✅ បាទ | ID នៃវាលបន្ថែមប្រទេស | - |
| `countryCodes` | [String!] | មិនត្រូវការ | អារ៉ាយនៃកូដប្រទេស ISO Alpha-2 | Stored in the `countryCodes` field |
| `text` | String | មិនត្រូវការ | អត្ថបទបង្ហាញឬឈ្មោះប្រទេស | Stored in the `text` field |

**សំខាន់**: 
- ក្នុង `setTodoCustomField`: ប៉ារ៉ាម៉ែត្រទាំងពីរជាជម្រើសនិងរក្សាទុកដោយឯករាជ្យ
- ក្នុង `createTodo`: ប្រព័ន្ធកំណត់វាលទាំងពីរដោយស្វ័យប្រវត្តិផ្អែកលើការចូលរបស់អ្នក (អ្នកមិនអាចគ្រប់គ្រងពួកវាដោយឯករាជ្យទេ)

### ជម្រើស 1: ប្រើតែកូដប្រទេស

រក្សាទុកកូដ ISO ដែលបានបញ្ជាក់ដោយគ្មានអត្ថបទបង្ហាញ:

```graphql
mutation SetCountryByCode {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
  })
}
```

លទ្ធផល: `countryCodes` = `["US"]`, `text` = `null`

### ជម្រើស 2: ប្រើតែអត្ថបទ

រក្សាទុកអត្ថបទបង្ហាញដោយគ្មានកូដដែលបានបញ្ជាក់:

```graphql
mutation SetCountryByText {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "United States"
  })
}
```

លទ្ធផល: `countryCodes` = `null`, `text` = `"United States"`

**កំណត់ចំណាំ**: នៅពេលប្រើ `setTodoCustomField`, មិនមានការត្រួតពិនិត្យកើតឡើងទេ មិនថាអ្នកប្រើប៉ារ៉ាម៉ែត្រ​អ្វី។ តម្លៃត្រូវបានរក្សាទុកយ៉ាងត្រឹមត្រូវដូចដែលបានផ្តល់។

### ជម្រើស 3: ប្រើទាំងពីរ (បានណែនាំ)

រក្សាទុកទាំងកូដដែលបានបញ្ជាក់និងអត្ថបទបង្ហាញ:

```graphql
mutation SetCountryComplete {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US"]
    text: "United States"
  })
}
```

លទ្ធផល: `countryCodes` = `["US"]`, `text` = `"United States"`

### ប្រទេសច្រើន

រក្សាទុកប្រទេសច្រើនដោយប្រើអារ៉ាយ:

```graphql
mutation SetMultipleCountries {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    countryCodes: ["US", "CA", "MX"]
    text: "North American Markets"  # Can be any descriptive text
  })
}
```

## ការបង្កើតកំណត់ត្រាជាមួយតម្លៃប្រទេស

នៅពេលបង្កើតកំណត់ត្រា, ការប្រែប្រួល `createTodo` **ធ្វើការត្រួតពិនិត្យនិងបម្លែង** តម្លៃប្រទេសដោយស្វ័យប្រវត្តិ។ នេះគឺជាការប្រែប្រួលតែមួយដែលអនុវត្តការត្រួតពិនិត្យប្រទេស:

```graphql
mutation CreateRecordWithCountry {
  createTodo(input: {
    title: "International Client"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "country_field_id"
      value: "France"  # Can use country name or code
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
      text
      countryCodes
    }
  }
}
```

### ទ្រង់ទ្រាយចូលដែលទទួលយក

| ប្រភេទចូល | ឧទាហរណ៍ | លទ្ធផល |
|------------|---------|---------|
| Country Name | `"United States"` | Stored as `US` |
| ISO Alpha-2 Code | `"GB"` | Stored as `GB` |
| Multiple (comma-separated) | `"US, CA"` | **មិនគាំទ្រ** - ត្រូវបានចាត់ទុកថាជាតម្លៃមិនត្រឹមត្រូវតែមួយ |
| Mixed format | `"United States, CA"` | **មិនគាំទ្រ** - ត្រូវបានចាត់ទុកថាជាតម្លៃមិនត្រឹមត្រូវតែមួយ |

## វាលចម្លើយ

### TodoCustomField ចម្លើយ

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកត្តសម្រាប់តម្លៃវាល |
| `customField` | CustomField! | ការបកស្រាយវាលបន្ថែម |
| `text` | String | អត្ថបទបង្ហាញ (ឈ្មោះប្រទេស) |
| `countryCodes` | [String!] | អារ៉ាយនៃកូដប្រទេស ISO Alpha-2 |
| `todo` | Todo! | កំណត់ត្រានេះមានតម្លៃ |
| `createdAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលវេលាដែលតម្លៃបានកែប្រែចុងក្រោយ |

## ស្តង់ដារប្រទេស

Blue ប្រើស្តង់ដារ **ISO 3166-1 Alpha-2** សម្រាប់កូដប្រទេស:

- កូដប្រទេសពីរអក្សរ (ឧ. US, GB, FR, DE)
- ការត្រួតពិនិត្យដោយប្រើបណ្ណាល័យ `i18n-iso-countries` **កើតឡើងតែក្នុង createTodo**
- គាំទ្រទាំងអស់ប្រទេសដែលត្រូវបានទទួលស្គាល់យ៉ាងផ្លូវការទាំងអស់

### ឧទាហរណ៍កូដប្រទេស

| ប្រទេស | កូដ ISO |
|---------|----------|
| United States | `US` |
| United Kingdom | `GB` |
| Canada | `CA` |
| Germany | `DE` |
| France | `FR` |
| Japan | `JP` |
| Australia | `AU` |
| Brazil | `BR` |

សម្រាប់បញ្ជីពេញលេញនៃកូដប្រទេស ISO 3166-1 alpha-2 សូមចូលទៅកាន់ [ISO Online Browsing Platform](https://www.iso.org/obp/ui/#search/code/)។

## ការត្រួតពិនិត្យ

**ការត្រួតពិនិត្យកើតឡើងតែក្នុងការប្រែប្រួល `createTodo`**:

1. **កូដ ISO ដែលមានសុពលភាព**: ទទួលយកកូដ ISO Alpha-2 ដែលមានសុពលភាពណាមួយ
2. **ឈ្មោះប្រទេស**: បម្លែងឈ្មោះប្រទេសដែលត្រូវបានទទួលស្គាល់ទៅកូដដោយស្វ័យប្រវត្តិ
3. **ការចូលមិនត្រឹមត្រូវ**: បោះបង់ `CustomFieldValueParseError` សម្រាប់តម្លៃដែលមិនត្រូវបានទទួលស្គាល់

**កំណត់ចំណាំ**: ការប្រែប្រួល `setTodoCustomField` មិនអនុវត្តការត្រួតពិនិត្យនិងទទួលយកតម្លៃខ្សែអក្សរណាមួយ។

### ឧទាហរណ៍កំហុស

```json
{
  "errors": [{
    "message": "Invalid country value.",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

## លក្ខណៈពិសេសនៃការបញ្ចូល

### វាលស្វែងរក
វាលប្រទេសអាចត្រូវបានយោងដោយវាលបន្ថែម LOOKUP ដែលអនុញ្ញាតឱ្យអ្នកទាញយកទិន្នន័យប្រទេសពីកំណត់ត្រាដែលពាក់ព័ន្ធ។

### ការប្រតិបត្តិ
ប្រើតម្លៃប្រទេសក្នុងលក្ខខណ្ឌស្វ័យប្រវត្តិ:
- បំបែកសកម្មភាពដោយប្រទេសជាក់លាក់
- ផ្ញើការជូនដំណឹងផ្អែកលើប្រទេស
- ផ្លូវការងារផ្អែកលើតំបន់ភូមិសាស្ត្រ

### ទម្រង់
វាលប្រទេសនៅក្នុងទម្រង់ធ្វើការត្រួតពិនិត្យការចូលរបស់អ្នកប្រើដោយស្វ័យប្រវត្តិ និងបម្លែងឈ្មោះប្រទេសទៅកូដ។

## អាជ្ញាប័ណ្ណដែលត្រូវការ

| សកម្មភាព | អាជ្ញាប័ណ្ណដែលត្រូវការ |
|--------|-------------------|
| Create country field | Project `OWNER` or `ADMIN` role |
| Update country field | Project `OWNER` or `ADMIN` role |
| Set country value | Standard record edit permissions |
| View country value | Standard record view permissions |

## ចម្លើយកំហុស

### តម្លៃប្រទេសមិនត្រឹមត្រូវ
```json
{
  "errors": [{
    "message": "Invalid country value provided",
    "extensions": {
      "code": "CUSTOM_FIELD_VALUE_PARSE_ERROR"
    }
  }]
}
```

### ការមិនត្រូវគ្នានៃប្រភេទវាល
```json
{
  "errors": [{
    "message": "Field type mismatch: expected COUNTRY",
    "extensions": {
      "code": "INVALID_FIELD_TYPE"
    }
  }]
}
```

## អនុសាសន៍ល្អ

### ការដោះស្រាយចូល
- ប្រើ `createTodo` សម្រាប់ការត្រួតពិនិត្យនិងបម្លែងដោយស្វ័យប្រវត្តិ
- ប្រើ `setTodoCustomField` ដោយប្រុងប្រយ័ត្នព្រោះវាប៉ះពាល់ការត្រួតពិនិត្យ
- ពិចារណាការត្រួតពិនិត្យការចូលនៅក្នុងកម្មវិធីរបស់អ្នកមុននឹងប្រើ `setTodoCustomField`
- បង្ហាញឈ្មោះប្រទេសពេញលេញនៅក្នុង UI សម្រាប់ភាពច្បាស់លាស់

### គុណភាពទិន្នន័យ
- ត្រួតពិនិត្យការចូលប្រទេសនៅចំណុចចូល
- ប្រើទ្រង់ទ្រាយដែលស្របគ្នានៅក្នុងប្រព័ន្ធរបស់អ្នក
- ពិចារណាការបែងចែកតំបន់សម្រាប់ការរាយការណ៍

### ប្រទេសច្រើន
- ប្រើការគាំទ្រអារ៉ាយនៅក្នុង `setTodoCustomField` សម្រាប់ប្រទេសច្រើន
- ប្រទេសច្រើនក្នុង `createTodo` មិន **គាំទ្រ** តាមវាលតម្លៃ
- រក្សាទុកកូដប្រទេសជាអារ៉ាយនៅក្នុង `setTodoCustomField` សម្រាប់ការដោះស្រាយត្រឹមត្រូវ

## ករណីប្រើប្រាស់ទូទៅ

1. **ការគ្រប់គ្រងអតិថិជន**
   - ទីតាំងកណ្តាលរបស់អតិថិជន
   - ទីកន្លែងដឹកជញ្ជូន
   - តំបន់ពន្ធ

2. **ការតាមដានគម្រោង**
   - ទីតាំងគម្រោង
   - ទីកន្លែងសមាជិកក្រុម
   - គោលដៅទីផ្សារ

3. **ការអនុវត្តន៍ និងច្បាប់**
   - តំបន់ច្បាប់
   - ការទាមទារទិន្នន័យ
   - ការត្រួតពិនិត្យការនាំចេញ

4. **ការលក់ និងទីផ្សារ**
   - ការបែងចែកតំបន់
   - ការបែងចែកទីផ្សារ
   - ការទាក់ទាញយុទ្ធនាការ

## កំណត់

- គាំទ្រតែគ្រាប់កូដ ISO 3166-1 Alpha-2 (កូដពីរអក្សរ)
- គ្មានការគាំទ្រដែលបានកំណត់សម្រាប់ការបែងចែកប្រទេស (រដ្ឋ/ខេត្ត)
- គ្មានរូបភាពតំណាងប្រទេសដោយស្វ័យប្រវត្តិ (គ្រាន់តែជាអត្ថបទ)
- មិនអាចត្រួតពិនិត្យកូដប្រទេសប្រវត្តិសាស្ត្រ
- គ្មានការបែងចែកតំបន់ឬទ្វីបដែលបានកំណត់
- **ការត្រួតពិនិត្យត្រឹមតែធ្វើការនៅក្នុង `createTodo`, មិននៅក្នុង `setTodoCustomField`**
- **ប្រទេសច្រើនមិនគាំទ្រនៅក្នុងវាលតម្លៃ `createTodo`**
- **កូដប្រទេសត្រូវបានរក្សាទុកជាខ្សែអក្សរដែលបំបែកដោយសញ្ញាកម្មា មិនមែនជាអារ៉ាយពិតទេ**

## ឯកសារដែលពាក់ព័ន្ធ

- [ទិដ្ឋភាពទូទៅនៃវាលបន្ថែម](/custom-fields/list-custom-fields) - គំនិតទូទៅអំពីវាលបន្ថែម
- [វាលស្វែងរក](/api/custom-fields/lookup) - យោងទិន្នន័យប្រទេសពីកំណត់ត្រាផ្សេងទៀត
- [Forms API](/api/forms) - រួមបញ្ចូលវាលប្រទេសនៅក្នុងទម្រង់បន្ថែម