---
title: ការបម្លែងរូបិយប័ណ្ណកំណត់ផ្ទាល់
description: បង្កើតកន្លែងដែលបម្លែងតម្លៃរូបិយប័ណ្ណដោយស្វ័យប្រវត្តិដោយប្រើអត្រាប្តូរពេលវេលាពិត
---

កន្លែងបម្លែងរូបិយប័ណ្ណកំណត់ផ្ទាល់បម្លែងតម្លៃពីកន្លែងរូបិយប័ណ្ណមួយទៅរូបិយប័ណ្ណគោលដៅផ្សេងៗដោយប្រើអត្រាប្តូរពេលវេលាពិត។ កន្លែងទាំងនេះធ្វើការអាប់ដេតដោយស្វ័យប្រវត្តិពេលណាដែលតម្លៃរូបិយប័ណ្ណប្រភពផ្លាស់ប្តូរ។

អត្រាបម្លែងត្រូវបានផ្តល់ដោយ [Frankfurter API](https://github.com/hakanensari/frankfurter) ដែលជាសេវាកម្មប្រភពបើកដែលតាមដានអត្រាប្តូរពីការបោះពុម្ពដែលបានផ្សាយដោយ [ធនាគារមជ្ឍមណ្ឌលអឺរ៉ុប](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html)។ នេះធានាថាការបម្លែងរូបិយប័ណ្ណមានភាពត្រឹមត្រូវ, អាចទុកចិត្តបាន និងមានព័ត៌មានថ្មីសម្រាប់តម្រូវការអាជីវកម្មអន្តរជាតិរបស់អ្នក។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតកន្លែងបម្លែងរូបិយប័ណ្ណសាមញ្ញ៖

```graphql
mutation CreateCurrencyConversionField {
  createCustomField(input: {
    name: "Price in EUR"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_currency_field_id"
    conversionDateType: "currentDate"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
  }
}
```

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតកន្លែងបម្លែងជាមួយកាលបរិច្ឆេទជាក់លាក់សម្រាប់អត្រាប្រវត្តិសាស្ត្រ៖

```graphql
mutation CreateHistoricalConversionField {
  createCustomField(input: {
    name: "Q1 Budget in Local Currency"
    type: CURRENCY_CONVERSION
    currencyFieldId: "budget_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2024-01-01T00:00:00Z"
    description: "Budget converted at Q1 exchange rates"
  }) {
    id
    name
    type
    currencyFieldId
    conversionDateType
    conversionDate
  }
}
```

## ដំណើរការកំណត់រចនាសម្ព័ន្ធពេញលេញ

ការកំណត់កន្លែងបម្លែងរូបិយប័ណ្ណត្រូវការជំហានបី៖

### ជំហាន ១: បង្កើតកន្លែងរូបិយប័ណ្ណប្រភព

```graphql
mutation CreateSourceCurrencyField {
  createCustomField(input: {
    name: "Contract Value"
    type: CURRENCY
    currency: "USD"
  }) {
    id  # Save this ID for Step 2
    name
    type
  }
}
```

### ជំហាន ២: បង្កើតកន្លែង CURRENCY_CONVERSION

```graphql
mutation CreateConversionField {
  createCustomField(input: {
    name: "Contract Value (Local Currency)"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id_from_step_1"
    conversionDateType: "currentDate"
  }) {
    id  # Save this ID for Step 3
    name
    type
  }
}
```

### ជំហាន ៣: បង្កើតជម្រើសបម្លែង

```graphql
mutation CreateConversionOptions {
  createCustomFieldOptions(input: {
    customFieldId: "conversion_field_id_from_step_2"
    customFieldOptions: [
      {
        title: "USD to EUR"
        currencyConversionFrom: "USD"
        currencyConversionTo: "EUR"
      },
      {
        title: "USD to GBP"
        currencyConversionFrom: "USD"
        currencyConversionTo: "GBP"
      },
      {
        title: "Any to JPY"
        currencyConversionFrom: "Any"
        currencyConversionTo: "JPY"
      }
    ]
  }) {
    id
    title
    currencyConversionFrom
    currencyConversionTo
  }
}
```

## ប៉ារ៉ាម៉ែត្រ Input

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃកន្លែងបម្លែង |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | មិនទាន់ | ID នៃកន្លែងរូបិយប័ណ្ណប្រភពដើម្បីបម្លែងពី |
| `conversionDateType` | String | មិនទាន់ | យុទ្ធសាស្ត្រកាលបរិច្ឆេទសម្រាប់អត្រាប្តូរ (មើលខាងក្រោម) |
| `conversionDate` | String | មិនទាន់ | ខ្សែអក្សរកាលបរិច្ឆេទសម្រាប់ការបម្លែង (ផ្អែកលើ conversionDateType) |
| `description` | String | មិនទាន់ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |

**កំណត់ចំណាំ**: កន្លែងកំណត់ផ្ទាល់ត្រូវបានភ្ជាប់ដោយស្វ័យប្រវត្តិជាមួយគម្រោងផ្អែកលើបរិបទគម្រោងបច្ចុប្បន្នរបស់អ្នកប្រើ។ មិនត្រូវការប៉ារ៉ាម៉ែត្រ `projectId` ទេ។

### ប្រភេទកាលបរិច្ឆេទបម្លែង

| ប្រភេទ | ការពិពណ៌នា | ប៉ារ៉ាម៉ែត្រ conversionDate |
|------|-------------|-------------------------|
| `currentDate` | ប្រើអត្រាប្តូរពេលវេលាពិត | មិនត្រូវការ |
| `specificDate` | ប្រើអត្រាពីកាលបរិច្ឆេទថេរ | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | ប្រើកាលបរិច្ឆេទពីកន្លែងផ្សេងទៀត | "todoDueDate" or DATE field ID |

## ការបង្កើតជម្រើសបម្លែង

ជម្រើសបម្លែងកំណត់ថាតើគូរូបិយប័ណ្ណណាដែលអាចបម្លែងបាន៖

### CreateCustomFieldOptionInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ បាទ | ID នៃកន្លែង CURRENCY_CONVERSION |
| `title` | String! | ✅ បាទ | ឈ្មោះបង្ហាញសម្រាប់ជម្រើសបម្លែងនេះ |
| `currencyConversionFrom` | String! | ✅ បាទ | កូដរូបិយប័ណ្ណប្រភពឬ "Any" |
| `currencyConversionTo` | String! | ✅ បាទ | កូដរូបិយប័ណ្ណគោលដៅ |

### ការប្រើប្រាស់ "Any" ជារូបិយប័ណ្ណប្រភព

តម្លៃពិសេស "Any" ជា `currencyConversionFrom` បង្កើតជម្រើសជំនួយ៖

```graphql
mutation CreateUniversalConversion {
  createCustomFieldOption(input: {
    customFieldId: "conversion_field_id"
    title: "Any currency to EUR"
    currencyConversionFrom: "Any"
    currencyConversionTo: "EUR"
  }) {
    id
  }
}
```

ជម្រើសនេះនឹងត្រូវបានប្រើនៅពេលដែលមិនមានការប៉ាន់ស្មានគូរូបិយប័ណ្ណជាក់លាក់ណាមួយ។

## របៀបដែលការបម្លែងដោយស្វ័យប្រវត្តិធ្វើការងារ

1. **ការអាប់ដេតតម្លៃ**: នៅពេលដែលតម្លៃត្រូវបានកំណត់នៅក្នុងកន្លែងរូបិយប័ណ្ណប្រភព
2. **ការប៉ាន់ស្មានជម្រើស**: ប្រព័ន្ធស្វែងរកជម្រើសបម្លែងដែលផ្គូផ្គងដោយផ្អែកលើរូបិយប័ណ្ណប្រភព
3. **ការទាញយកអត្រា**: ទាញយកអត្រាប្តូរពី Frankfurter API
4. **ការគណនា**: គុណចំនួនប្រភពដោយអត្រាប្តូរ
5. **ការផ្ទុក**: រក្សាទុកតម្លៃដែលបានបម្លែងជាមួយកូដរូបិយប័ណ្ណគោលដៅ

### ឧទាហរណ៍ចលនាដែលធ្វើការងារ

```graphql
# 1. Set value in source CURRENCY field
mutation SetSourceValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "source_currency_field_id"
    number: 1000
    currency: "USD"
  })
}

# 2. CURRENCY_CONVERSION fields automatically update
# If you have USD→EUR and USD→GBP options configured,
# both conversion fields will calculate and store their values
```

## ការបម្លែងផ្អែកលើកាលបរិច្ឆេទ

### ការប្រើប្រាស់កាលបរិច្ឆេទបច្ចុប្បន្ន

```graphql
mutation CreateRealtimeConversion {
  createCustomField(input: {
    name: "Current EUR Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "currentDate"
  })
}
```

ការបម្លែងអាប់ដេតជាមួយអត្រាប្តូរបច្ចុប្បន្នរាល់ពេលដែលតម្លៃប្រភពផ្លាស់ប្តូរ។

### ការប្រើប្រាស់កាលបរិច្ឆេទជាក់លាក់

```graphql
mutation CreateFixedDateConversion {
  createCustomField(input: {
    name: "Year-End 2023 Value"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "specificDate"
    conversionDate: "2023-12-31T00:00:00Z"
  })
}
```

ប្រើអត្រាប្តូរពីកាលបរិច្ឆេទដែលបានកំណត់ជានិច្ច។

### ការប្រើប្រាស់កាលបរិច្ឆេទពីកន្លែង

```graphql
mutation CreateDateFieldConversion {
  createCustomField(input: {
    name: "Value at Contract Date"
    type: CURRENCY_CONVERSION
    currencyFieldId: "source_field_id"
    conversionDateType: "fromDateField"
    conversionDate: "contract_date_field_id"  # ID of a DATE custom field
  })
}
```

ប្រើកាលបរិច្ឆេទពីកន្លែងផ្សេងទៀត (មិនថាជាកាលបរិច្ឆេទកំណត់ឬកន្លែងកាលបរិច្ឆេទ)។

## កន្លែងឆ្លើយតប

### TodoCustomField Response

| កន្លែង | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណឯកតាដែលមានសម្រាប់តម្លៃកន្លែង |
| `customField` | CustomField! | ការកំណត់កន្លែងបម្លែង |
| `number` | Float | ចំនួនដែលបានបម្លែង |
| `currency` | String | កូដរូបិយប័ណ្ណគោលដៅ |
| `todo` | Todo! | កំណត់ត្រាដែលតម្លៃនេះស្ថិតនៅក្នុង |
| `createdAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលវេលាដែលតម្លៃបានអាប់ដេតចុងក្រោយ |

## ប្រភពអត្រាប្តូរ

Blue ប្រើ **Frankfurter API** សម្រាប់អត្រាប្តូរ៖
- API បើកដែលមានសេវាដោយធនាគារមជ្ឍមណ្ឌលអឺរ៉ុប
- អាប់ដេតរៀងរាល់ថ្ងៃជាមួយអត្រាប្តូរផ្លូវការនៅ
- គាំទ្រអត្រាប្រវត្តិសាស្ត្រចាប់តាំងពីឆ្នាំ 1999
- ឥតគិតថ្លៃ និងអាចទុកចិត្តបានសម្រាប់ការប្រើប្រាស់អាជីវកម្ម

## ការគ្រប់គ្រងកំហុស

### ការបរាជ័យក្នុងការបម្លែង

នៅពេលដែលការបម្លែងបរាជ័យ (កំហុស API, រូបិយប័ណ្ណមិនត្រឹមត្រូវ, ល។):
- តម្លៃដែលបានបម្លែងត្រូវបានកំណត់ទៅ `0`
- កូដរូបិយប័ណ្ណគោលដៅនៅតែត្រូវបានរក្សាទុក
- មិនមានកំហុសណាមួយត្រូវបានបោះបង់ទៅអ្នកប្រើ

### ស្ថានភាពទូទៅ

| ស្ថានភាព | លទ្ធផល |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| មិនមានជម្រើសផ្គូផ្គង | Uses "Any" option if available |
| Missing source value | មិនមានការបម្លែងធ្វើឡើង |

## អាជ្ញាប័ណ្ណដែលត្រូវការ

ការគ្រប់គ្រងកន្លែងកំណត់ត្រូវការការចូលដំណើរការនៅកម្រិតគម្រោង៖

| តួនាទី | អាចបង្កើត/អាប់ដេតកន្លែង |
|------|-------------------------|
| `OWNER` | ✅ បាទ |
| `ADMIN` | ✅ បាទ |
| `MEMBER` | ❌ ទេ |
| `CLIENT` | ❌ ទេ |

ការអនុញ្ញាតមើលសម្រាប់តម្លៃដែលបានបម្លែងតាមដានច្បាប់ចូលដំណើរការទូទៅ។

## អនុសាសន៍ល្អ

### ការកំណត់ជម្រើស
- បង្កើតគូរូបិយប័ណ្ណជាក់លាក់សម្រាប់ការបម្លែងទូទៅ
- បន្ថែមជម្រើសជំនួយ "Any" សម្រាប់ភាពបត់បែន
- ប្រើឈ្មោះពិពណ៌នាសម្រាប់ជម្រើស

### ការជ្រើសរើសយុទ្ធសាស្ត្រកាលបរិច្ឆេទ
- ប្រើ `currentDate` សម្រាប់ការតាមដានហិរញ្ញវត្ថុបច្ចុប្បន្ន
- ប្រើ `specificDate` សម្រាប់របាយការណ៍ប្រវត្តិសាស្ត្រ
- ប្រើ `fromDateField` សម្រាប់អត្រាដែលជាក់លាក់ក្នុងប្រតិបត្តិការ

### ការពិចារណានៅលើប្រសិទ្ធភាព
- កន្លែងបម្លែងច្រើនធ្វើការអាប់ដេតនៅក្នុងពេលតែមួយ
- ការហៅ API ត្រូវបានធ្វើឡើងតែពេលដែលតម្លៃប្រភពផ្លាស់ប្តូរ
- ការបម្លែងរូបិយប័ណ្ណដូចគ្នាអាចឆ្លងកាត់ការហៅ API

## ករណីប្រើប្រាស់ទូទៅ

1. **គម្រោងរូបិយប័ណ្ណច្រើន**
   - តាមដានចំណាយគម្រោងនៅក្នុងរូបិយប័ណ្ណក្នុងស្រុក
   - រាយការណ៍ថវិកាសរុបនៅក្នុងរូបិយប័ណ្ណក្រុមហ៊ុន
   - ប្រៀបធៀបតម្លៃនៅតាមតំបន់

2. **ការលក់អន្តរជាតិ**
   - បម្លែងតម្លៃកិច្ចសន្យាទៅកាន់រូបិយប័ណ្ណរាយការណ៍
   - តាមដានប្រាក់ចំណូលនៅក្នុងរូបិយប័ណ្ណច្រើន
   - ការបម្លែងប្រវត្តិសាស្ត្រសម្រាប់កិច្ចសន្យាដែលបានបិទ

3. **របាយការណ៍ហិរញ្ញវត្ថុ**
   - ការបម្លែងរូបិយប័ណ្ណនៅចុងរដូវ
   - របាយការណ៍ហិរញ្ញវត្ថុដែលបានបញ្ចូល
   - ថវិកាដោយប្រាក់ចំណូលក្នុងរូបិយប័ណ្ណក្នុងស្រុក

4. **ការគ្រប់គ្រងកិច្ចសន្យា**
   - បម្លែងតម្លៃកិច្ចសន្យានៅពេលចុះហត្ថលេខា
   - តាមដានកាលវិភាគការទូទាត់នៅក្នុងរូបិយប័ណ្ណច្រើន
   - ការវាយតម្លៃហានិភ័យរូបិយប័ណ្ណ

## ការកំណត់

- មិនគាំទ្រការបម្លែងរូបិយប័ណ្ណឌីជីថល
- មិនអាចកំណត់តម្លៃដែលបានបម្លែងដោយដៃ (តែងតែគណនា)
- កំណត់ភាពត្រឹមត្រូវ ២ ចំនួនទសភាគសម្រាប់តម្លៃដែលបានបម្លែងទាំងអស់
- មិនគាំទ្រអត្រាប្តូរដោយផ្ទាល់
- មិនមានការផ្ទុកអត្រាប្តូរ (ការហៅ API ថ្មីសម្រាប់ការបម្លែងនីមួយៗ)
- អាស្រ័យលើភាពអាចប្រើប្រាស់នៃ Frankfurter API

## ឯកសារទាក់ទង

- [កន្លែងរូបិយប័ណ្ណ](/api/custom-fields/currency) - កន្លែងប្រភពសម្រាប់ការបម្លែង
- [កន្លែងកាលបរិច្ឆេទ](/api/custom-fields/date) - សម្រាប់ការបម្លែងផ្អែកលើកាលបរិច្ឆេទ
- [កន្លែងគណនាផលិតផល](/api/custom-fields/formula) - ការគណនាផ្សេងទៀត
- [ទិដ្ឋភាពទូទៅនៃកន្លែងកំណត់ផ្ទាល់](/custom-fields/list-custom-fields) - គំនិតទូទៅ