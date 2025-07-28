---
title: វាលកំណត់រូបិយវត្ថុ
description: បង្កើតវាលរូបិយវត្ថុដើម្បីតាមដានតម្លៃប្រាក់ដោយមានទ្រង់ទ្រាយនិងការផ្ទៀងផ្ទាត់ត្រឹមត្រូវ
---

វាលកំណត់រូបិយវត្ថុអនុញ្ញាតឱ្យអ្នករក្សាទុកនិងគ្រប់គ្រងតម្លៃប្រាក់ជាមួយកូដរូបិយវត្ថុដែលទាក់ទង។ វាលនេះគាំទ្រទៅនឹងរូបិយវត្ថុ 72 ប្រភេទដែលរួមមានរូបិយវត្ថុ fiat សំខាន់ៗ និងរូបិយវត្ថុឌីជីថល ជាមួយនឹងការទ្រង់ទ្រាយស្វ័យប្រវត្តិ និងកំណត់ min/max ជាជម្រើស។

## ឧទាហរណ៍មូលដ្ឋាន

បង្កើតវាលរូបិយវត្ថុសាមញ្ញ៖

```graphql
mutation CreateCurrencyField {
  createCustomField(input: {
    name: "Budget"
    type: CURRENCY
    projectId: "proj_123"
    currency: "USD"
  }) {
    id
    name
    type
    currency
  }
}
```

## ឧទាហរណ៍កម្រិតខ្ពស់

បង្កើតវាលរូបិយវត្ថុជាមួយនឹងការកំណត់ការផ្ទៀងផ្ទាត់៖

```graphql
mutation CreateConstrainedCurrencyField {
  createCustomField(input: {
    name: "Deal Value"
    type: CURRENCY
    projectId: "proj_123"
    currency: "EUR"
    min: 0
    max: 1000000
    description: "Estimated deal value in euros"
    isActive: true
  }) {
    id
    name
    type
    currency
    min
    max
    description
  }
}
```

## ប៉ារ៉ាម៉ែត្របញ្ចូល

### CreateCustomFieldInput

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `name` | String! | ✅ បាទ | ឈ្មោះបង្ហាញនៃវាលរូបិយវត្ថុ |
| `type` | CustomFieldType! | ✅ បាទ | ត្រូវតែជា `CURRENCY` |
| `currency` | String | ទេ | កូដរូបិយវត្ថុលំនាំដើម (កូដ ISO 3 អក្សរ) |
| `min` | Float | ទេ | តម្លៃអប្បបរមាដែលអនុញ្ញាត (រក្សាទុកប៉ុន្តែមិនអនុវត្តនៅពេលធ្វើបច្ចុប្បន្នភាព) |
| `max` | Float | ទេ | តម្លៃអតិបរិមាដែលអនុញ្ញាត (រក្សាទុកប៉ុន្តែមិនអនុវត្តនៅពេលធ្វើបច្ចុប្បន្នភាព) |
| `description` | String | ទេ | អត្ថបទជំនួយដែលបង្ហាញទៅអ្នកប្រើ |

**ចំណាំ**: បរិបទគម្រោងត្រូវបានកំណត់ដោយស្វ័យប្រវត្តិពីការផ្ទៀងផ្ទាត់របស់អ្នក។ អ្នកត្រូវមានសិទ្ធិចូលទៅកាន់គម្រោងដែលអ្នកកំពុងបង្កើតវាលនេះ។

## កំណត់តម្លៃរូបិយវត្ថុ

ដើម្បីកំណត់ឬធ្វើបច្ចុប្បន្នភាពតម្លៃរូបិយវត្ថុនៅលើកំណត់ត្រា៖

```graphql
mutation SetCurrencyValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    number: 1500.50
    currency: "USD"
  })
}
```

### SetTodoCustomFieldInput Parameters

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ត្រូវការ | ការពិពណ៌នា |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ បាទ | ID នៃកំណត់ត្រាដើម្បីធ្វើបច្ចុប្បន្នភាព |
| `customFieldId` | String! | ✅ បាទ | ID នៃវាលកំណត់រូបិយវត្ថុ |
| `number` | Float! | ✅ បាទ | តម្លៃប្រាក់ |
| `currency` | String! | ✅ បាទ | កូដរូបិយវត្ថុ 3 អក្សរ |

## បង្កើតកំណត់ត្រាជាមួយតម្លៃរូបិយវត្ថុ

នៅពេលបង្កើតកំណត់ត្រាថ្មីជាមួយតម្លៃរូបិយវត្ថុ៖

```graphql
mutation CreateRecordWithCurrency {
  createTodo(input: {
    title: "Q4 Marketing Campaign"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "currency_field_id"
      value: "25000.00"
      currency: "GBP"
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
      number
      currency
    }
  }
}
```

### របៀបបញ្ចូលសម្រាប់បង្កើត

នៅពេលបង្កើតកំណត់ត្រា តម្លៃរូបិយវត្ថុត្រូវបានបញ្ជូនខុសគ្នា៖

| ប៉ារ៉ាម៉ែត្រ | ប្រភេទ | ការពិពណ៌នា |
|-----------|------|-------------|
| `customFieldId` | String! | ID នៃវាលរូបិយវត្ថុ |
| `value` | String! | តម្លៃជាអក្សរ (ឧ. "1500.50") |
| `currency` | String! | កូដរូបិយវត្ថុ 3 អក្សរ |

## រូបិយវត្ថុដែលគាំទ្រ

Blue គាំទ្ររូបិយវត្ថុ 72 ប្រភេទដែលរួមមានរូបិយវត្ថុ fiat 70 ប្រភេទ និងរូបិយវត្ថុឌីជីថល 2 ប្រភេទ៖

### រូបិយវត្ថុ Fiat

#### អាមេរិក
| រូបិយវត្ថុ | កូដ | ឈ្មោះ |
|----------|------|------|
| US Dollar | `USD` | US Dollar |
| Canadian Dollar | `CAD` | Canadian Dollar |
| Mexican Peso | `MXN` | Mexican Peso |
| Brazilian Real | `BRL` | Brazilian Real |
| Argentine Peso | `ARS` | Argentine Peso |
| Chilean Peso | `CLP` | Chilean Peso |
| Colombian Peso | `COP` | Colombian Peso |
| Peruvian Sol | `PEN` | Peruvian Sol |
| Uruguayan Peso | `UYU` | Uruguayan Peso |
| Venezuelan Bolívar | `VES` | Venezuelan Bolívar Soberano |
| Bolivian Boliviano | `BOB` | Bolivian Boliviano |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### អឺរ៉ុប
| រូបិយវត្ថុ | កូដ | ឈ្មោះ |
|----------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| Norwegian Krone | `NOK` | Norwegian Krone |
| Danish Krone | `DKK` | Danish Krone |
| Polish Złoty | `PLN` | Polish Złoty |
| Czech Koruna | `CZK` | Czech Koruna |
| Hungarian Forint | `HUF` | Hungarian Forint |
| Romanian Leu | `RON` | Romanian Leu |
| Bulgarian Lev | `BGN` | Bulgarian Lev |
| Turkish Lira | `TRY` | Turkish Lira |
| Ukrainian Hryvnia | `UAH` | Ukrainian Hryvnia |
| Russian Ruble | `RUB` | Russian Ruble |
| Georgian Lari | `GEL` | Georgian Lari |
| Icelandic króna | `ISK` | Icelandic króna |
| Bosnia-Herzegovina Mark | `BAM` | Bosnia-Herzegovina Convertible Mark |

#### តំបន់អាស៊ី-ប៉ាស៊ីហ្វិក
| រូបិយវត្ថុ | កូដ | ឈ្មោះ |
|----------|------|------|
| Japanese Yen | `JPY` | Yen |
| Chinese Yuan | `CNY` | Yuan |
| Hong Kong Dollar | `HKD` | Hong Kong Dollar |
| Singapore Dollar | `SGD` | Singapore Dollar |
| Australian Dollar | `AUD` | Australian Dollar |
| New Zealand Dollar | `NZD` | New Zealand Dollar |
| South Korean Won | `KRW` | South Korean Won |
| Indian Rupee | `INR` | Indian Rupee |
| Indonesian Rupiah | `IDR` | Indonesian Rupiah |
| Thai Baht | `THB` | Thai Baht |
| Malaysian Ringgit | `MYR` | Malaysian Ringgit |
| Philippine Peso | `PHP` | Philippine Peso |
| Vietnamese Dong | `VND` | Vietnamese Dong |
| Taiwanese Dollar | `TWD` | New Taiwan Dollar |
| Pakistani Rupee | `PKR` | Pakistani Rupee |
| Sri Lankan Rupee | `LKR` | Sri Lankan Rupee |
| Cambodian Riel | `KHR` | Cambodian Riel |
| Kazakhstani Tenge | `KZT` | Kazakhstani Tenge |

#### តំបន់មជ្ឈមណ្ឌល និងអាហ្វ្រិក
| រូបិយវត្ថុ | កូដ | ឈ្មោះ |
|----------|------|------|
| UAE Dirham | `AED` | UAE Dirham |
| Saudi Riyal | `SAR` | Saudi Riyal |
| Kuwaiti Dinar | `KWD` | Kuwaiti Dinar |
| Bahraini Dinar | `BHD` | Bahraini Dinar |
| Qatari Riyal | `QAR` | Qatari Riyal |
| Israeli Shekel | `ILS` | Israeli New Shekel |
| Egyptian Pound | `EGP` | Egyptian Pound |
| Moroccan Dirham | `MAD` | Moroccan Dirham |
| Tunisian Dinar | `TND` | Tunisian Dinar |
| South African Rand | `ZAR` | South African Rand |
| Kenyan Shilling | `KES` | Kenyan Shilling |
| Nigerian Naira | `NGN` | Nigerian Naira |
| Ghanaian Cedi | `GHS` | Ghanaian Cedi |
| Zambian Kwacha | `ZMW` | Zambian Kwacha |
| Malagasy Ariary | `MGA` | Malagasy Ariary |

### រូបិយវត្ថុឌីជីថល
| រូបិយវត្ថុ | កូដ |
|----------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## វាលឆ្លើយតប

### TodoCustomField Response

| វាល | ប្រភេទ | ការពិពណ៌នា |
|-------|------|-------------|
| `id` | String! | អត្តសញ្ញាណដ៏ឯកតាសម្រាប់តម្លៃវាល |
| `customField` | CustomField! | ការកំណត់វាលកំណត់ |
| `number` | Float | តម្លៃប្រាក់ |
| `currency` | String | កូដរូបិយវត្ថុ 3 អក្សរ |
| `todo` | Todo! | កំណត់ត្រាដែលតម្លៃនេះជាប់ពាក់ព័ន្ធ |
| `createdAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានបង្កើត |
| `updatedAt` | DateTime! | ពេលវេលាដែលតម្លៃត្រូវបានកែប្រែចុងក្រោយ |

## ការទ្រង់ទ្រាយរូបិយវត្ថុ

ប្រព័ន្ធនឹងទ្រង់ទ្រាយតម្លៃរូបិយវត្ថុដោយស្វ័យប្រវត្តិផ្អែកលើភូមិសាស្ត្រ៖

- **ទីតាំងសញ្ញា**: ដាក់សញ្ញារូបិយវត្ថុឲ្យត្រឹមត្រូវ (មុន/ក្រោយ)
- **អ្នកបំបែកទសភាគ**: ប្រើអ្នកបំបែកដែលមានលក្ខណៈជាក់លាក់ (. ឬ ,)
- **អ្នកបំបែកពាន់**: អនុវត្តការបែងចែកដែលសមស្រប
- **ទីតាំងទសភាគ**: បង្ហាញទីតាំងទសភាគ 0-2 ផ្អែកលើចំនួន
- **ការប្រតិបត្តិពិសេស**: USD/CAD បង្ហាញកូដរូបិយវត្ថុជាមុនសម្រាប់ភាពច្បាស់លាស់

### ឧទាហរណ៍ការទ្រង់ទ្រាយ

| តម្លៃ | រូបិយវត្ថុ | បង្ហាញ |
|-------|----------|---------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## ការផ្ទៀងផ្ទាត់

### ការផ្ទៀងផ្ទាត់តម្លៃ
- ត្រូវតែជាចំនួនត្រឹមត្រូវ
- កំណត់ min/max ត្រូវបានរក្សាទុកជាមួយនឹងការកំណត់វាល ប៉ុន្តែមិនអនុវត្តនៅពេលធ្វើបច្ចុប្បន្នភាពតម្លៃ
- គាំទ្រទៅនឹងទីតាំងទសភាគ 2 សម្រាប់ការបង្ហាញ (ការពិតត្រូវបានរក្សាទុកនៅក្នុង)

### ការផ្ទៀងផ្ទាត់កូដរូបិយវត្ថុ
- ត្រូវតែជាកូដរូបិយវត្ថុ 72 ដែលគាំទ្រ
- មានអក្សរធំ (ប្រើអក្សរធំ)
- កូដមិនត្រឹមត្រូវនឹងត្រឡប់មកវិញកំហុស

## លក្ខណៈពិសេសនៃការបញ្ចូល

### ធរណីមាត្រ
វាលរូបិយវត្ថុអាចប្រើបានក្នុងវាលកំណត់ FORMULA សម្រាប់ការគណនា៖
- បូកវាលរូបិយវត្ថុច្រើន
- គណនាភាគរយ
- ប្រតិបត្តិការគណនា

### ការបម្លែងរូបិយវត្ថុ
ប្រើវាល CURRENCY_CONVERSION ដើម្បីបម្លែងដោយស្វ័យប្រវត្តិរវាងរូបិយវត្ថុ (មើល [វាលការបម្លែងរូបិយវត្ថុ](/api/custom-fields/currency-conversion))

### ការប្រតិបត្តិ
តម្លៃរូបិយវត្ថុអាចជំរុញការប្រតិបត្តិផ្អែកលើ៖
- កម្រិតតម្លៃ
- ប្រភេទរូបិយវត្ថុ
- ការផ្លាស់ប្តូរតម្លៃ

## សិទ្ធិដែលត្រូវការ

| សកម្មភាព | សិទ្ធិដែលត្រូវការ |
|--------|-------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**ចំណាំ**: ខណៈពេលដែលសមាជិកគម្រោងណាមួយអាចបង្កើតវាលកំណត់បាន វាសិទ្ធិក្នុងការកំណត់តម្លៃអាស្រ័យលើសិទ្ធិដែលបានកំណត់ដោយមូលដ្ឋានសម្រាប់វាលនីមួយៗ។

## ការឆ្លើយតបកំហុស

### តម្លៃរូបិយវត្ថុមិនត្រឹមត្រូវ
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

កំហុសនេះកើតឡើងនៅពេល៖
- កូដរូបិយវត្ថុមិនមែនជាកូដ 72 ដែលគាំទ្រទេ
- ទ្រង់ទ្រាយលេខមិនត្រឹមត្រូវ
- តម្លៃមិនអាចត្រូវបានបញ្ចូលបានត្រឹមត្រូវ

### វាលកំណត់មិនឃើញ
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

## អនុសាសន៍ល្អ

### ការជ្រើសរូបិយវត្ថុ
- កំណត់រូបិយវត្ថុលំនាំដើមដែលសមស្របជាមួយទីផ្សារដែលអ្នកប្រើប្រាស់
- ប្រើកូដរូបិយវត្ថុ ISO 4217 ជាប្រពៃណី
- ពិចារណាទីតាំងអ្នកប្រើនៅពេលជ្រើសរូបិយវត្ថុលំនាំដើម

### ការកំណត់តម្លៃ
- កំណត់តម្លៃ min/max ដែលសមស្របដើម្បីចៀសវាងកំហុសក្នុងការបញ្ចូលទិន្នន័យ
- ប្រើ 0 ជា min សម្រាប់វាលដែលមិនគួរត្រូវមានតម្លៃអវិជ្ជមាន
- ពិចារណាករណីប្រើប្រាស់របស់អ្នកនៅពេលកំណត់តម្លៃអតិបរិមា

### គម្រោងរូបិយវត្ថុច្រើន
- ប្រើរូបិយវត្ថុមូលដ្ឋានដែលសមស្របសម្រាប់ការរាយការណ៍
- អនុវត្តវាល CURRENCY_CONVERSION សម្រាប់ការបម្លែងដោយស្វ័យប្រវត្តិ
- ឯកសារថាតើរូបិយវត្ថុណាដែលគួរប្រើសម្រាប់វាលនីមួយៗ

## ករណីប្រើប្រាស់ទូទៅ

1. **ការប្រាក់គម្រោង**
   - ការតាមដានថវិកាគម្រោង
   - ការប៉ាន់ស្មានថ្លៃ
   - ការតាមដានចំណាយ

2. **ការលក់ និងកិច្ចព្រមព្រៀង**
   - តម្លៃកិច្ចព្រមព្រៀង
   - ចំនួនកិច្ចសន្យា
   - ការតាមដានប្រាក់ចំណូល

3. **ការធ្វើផែនការហិរញ្ញវត្ថុ**
   - តម្លៃវិនិយោគ
   - ជ Runde Funding
   - គោលដៅហិរញ្ញវត្ថុ

4. **អាជីវកម្មអន្តរជាតិ**
   - តម្លៃច្រើនរូបិយវត្ថុ
   - ការតាមដានការប្តូរប្រាក់
   - ប្រតិបត្តិការប្រទេសក្រៅ

## កំណត់

- ទីតាំងទសភាគ 2 សម្រាប់ការបង្ហាញ (ប៉ុន្តែមានភាពច្បាស់លាស់បន្ថែមត្រូវបានរក្សាទុក)
- មិនមានការបម្លែងរូបិយវត្ថុក្នុងវាល CURRENCY ស្តង់ដា
- មិនអាចបញ្ចូលរូបិយវត្ថុច្រើននៅក្នុងតម្លៃវាលមួយ
- មិនមានការអាប់ដេតអត្រាប្តូរដោយស្វ័យប្រវត្តិ (ប្រើ CURRENCY_CONVERSION សម្រាប់នេះ)
- សញ្ញារូបិយវត្ថុមិនអាចកំណត់បាន

## ឯកសារដែលទាក់ទង

- [វាលការបម្លែងរូបិយវត្ថុ](/api/custom-fields/currency-conversion) - ការបម្លែងរូបិយវត្ថុដោយស្វ័យប្រវត្តិ
- [វាលលេខ](/api/custom-fields/number) - សម្រាប់តម្លៃលេខដែលមិនមែនជាប្រាក់
- [វាលធរណីមាត្រ](/api/custom-fields/formula) - គណនាជាមួយតម្លៃរូបិយវត្ថុ
- [វាលកំណត់បញ្ជី](/api/custom-fields/list-custom-fields) - ស្វែងរកវាលកំណត់ទាំងអស់នៅក្នុងគម្រោង