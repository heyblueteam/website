---
title: 통화 사용자 정의 필드
description: 적절한 형식과 유효성 검사를 통해 금전적 가치를 추적하기 위한 통화 필드를 생성합니다.
---

통화 사용자 정의 필드는 관련 통화 코드와 함께 금전적 가치를 저장하고 관리할 수 있게 해줍니다. 이 필드는 주요 법정 통화와 암호화폐를 포함하여 72개의 다양한 통화를 지원하며, 자동 형식 지정 및 선택적 최소/최대 제약 조건을 제공합니다.

## 기본 예제

간단한 통화 필드를 생성합니다:

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

## 고급 예제

유효성 검사 제약 조건이 있는 통화 필드를 생성합니다:

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

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 통화 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | 반드시 `CURRENCY` 여야 합니다. |
| `currency` | String | 아니요 | 기본 통화 코드 (3자리 ISO 코드) |
| `min` | Float | 아니요 | 허용되는 최소값 (저장되지만 업데이트 시 적용되지 않음) |
| `max` | Float | 아니요 | 허용되는 최대값 (저장되지만 업데이트 시 적용되지 않음) |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 프로젝트 컨텍스트는 인증에서 자동으로 결정됩니다. 필드를 생성하는 프로젝트에 대한 액세스 권한이 있어야 합니다.

## 통화 값 설정

레코드에서 통화 값을 설정하거나 업데이트하려면:

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

### SetTodoCustomFieldInput 매개변수

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ 예 | 업데이트할 레코드의 ID |
| `customFieldId` | String! | ✅ 예 | 통화 사용자 정의 필드의 ID |
| `number` | Float! | ✅ 예 | 금전적 금액 |
| `currency` | String! | ✅ 예 | 3자리 통화 코드 |

## 통화 값이 포함된 레코드 생성

통화 값이 포함된 새 레코드를 생성할 때:

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

### 생성 입력 형식

레코드를 생성할 때 통화 값은 다르게 전달됩니다:

| 매개변수 | 유형 | 설명 |
|-----------|------|-------------|
| `customFieldId` | String! | 통화 필드의 ID |
| `value` | String! | 문자열로서의 금액 (예: "1500.50") |
| `currency` | String! | 3자리 통화 코드 |

## 지원되는 통화

Blue는 70개의 법정 통화와 2개의 암호화폐를 포함하여 72개의 통화를 지원합니다:

### 법정 통화

#### 아메리카
| 통화 | 코드 | 이름 |
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
| Venezuelan Bolívar | `VES` | 베네수엘라 볼리바르 수베라노 |
| 볼리비아 볼리비아노 | `BOB` | 볼리비아 볼리비아노 |
| Costa Rican Colón | `CRC` | Costa Rican Colón |
| Dominican Peso | `DOP` | Dominican Peso |
| Guatemalan Quetzal | `GTQ` | Guatemalan Quetzal |
| Jamaican Dollar | `JMD` | Jamaican Dollar |

#### 유럽
| 통화 | 코드 | 이름 |
|----------|------|------|
| Euro | `EUR` | Euro |
| British Pound | `GBP` | Pound Sterling |
| Swiss Franc | `CHF` | Swiss Franc |
| Swedish Krona | `SEK` | Swedish Krona |
| 노르웨이 크로네 | `NOK` | 노르웨이 크로네 |
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

#### 아시아-태평양
| 통화 | 코드 | 이름 |
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

#### 중동 및 아프리카
| 통화 | 코드 | 이름 |
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

### 암호화폐
| 통화 | 코드 |
|----------|------|
| Bitcoin | `BTC` |
| Ethereum | `ETH` |

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 사용자 정의 필드 정의 |
| `number` | Float | 금전적 금액 |
| `currency` | String | 3자리 통화 코드 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 수정된 시간 |

## 통화 형식 지정

시스템은 로케일에 따라 통화 값을 자동으로 형식화합니다:

- **기호 위치**: 통화 기호를 올바르게 배치합니다 (앞/뒤)
- **소수 구분 기호**: 로케일에 따라 구분 기호를 사용합니다 (. 또는 ,)
- **천 단위 구분 기호**: 적절한 그룹화를 적용합니다
- **소수 자리수**: 금액에 따라 0-2 소수 자리수를 표시합니다
- **특별 처리**: USD/CAD는 명확성을 위해 통화 코드 접두사를 표시합니다

### 형식 지정 예제

| 값 | 통화 | 표시 |
|-------|----------|---------|
| 1500.50 | USD | USD $1,500.50 |
| 1500.50 | EUR | €1.500,50 |
| 1500 | JPY | ¥1,500 |
| 1500.99 | GBP | £1,500.99 |

## 유효성 검사

### 금액 유효성 검사
- 유효한 숫자여야 합니다
- 최소/최대 제약 조건은 필드 정의와 함께 저장되지만 값 업데이트 시 적용되지 않습니다
- 표시를 위해 최대 2자리 소수를 지원합니다 (내부적으로는 전체 정밀도가 저장됨)

### 통화 코드 유효성 검사
- 72개의 지원되는 통화 코드 중 하나여야 합니다
- 대소문자를 구분합니다 (대문자 사용)
- 잘못된 코드는 오류를 반환합니다

## 통합 기능

### 수식
통화 필드는 계산을 위해 FORMULA 사용자 정의 필드에서 사용할 수 있습니다:
- 여러 통화 필드의 합계
- 백분율 계산
- 산술 연산 수행

### 통화 변환
CURRENCY_CONVERSION 필드를 사용하여 통화 간 자동 변환을 수행합니다 (자세한 내용은 [통화 변환 필드](/api/custom-fields/currency-conversion) 참조)

### 자동화
통화 값은 다음을 기반으로 자동화를 트리거할 수 있습니다:
- 금액 임계값
- 통화 유형
- 값 변경

## 필요한 권한

| 작업 | 필요한 권한 |
|--------|-------------------|
| Create currency field | Must be a member of the project (any role) |
| Update currency field | Must be a member of the project (any role) |
| Set currency value | Must have edit permissions based on project role |
| View currency value | Standard record view permissions |

**참고**: 모든 프로젝트 구성원이 사용자 정의 필드를 생성할 수 있지만, 값을 설정하는 기능은 각 필드에 대해 구성된 역할 기반 권한에 따라 달라집니다.

## 오류 응답

### 잘못된 통화 값
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

이 오류는 다음과 같은 경우에 발생합니다:
- 통화 코드가 지원되는 72개 코드 중 하나가 아닙니다
- 숫자 형식이 잘못되었습니다
- 값을 올바르게 구문 분석할 수 없습니다

### 사용자 정의 필드를 찾을 수 없음
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

## 모범 사례

### 통화 선택
- 기본 시장에 맞는 기본 통화를 설정합니다
- ISO 4217 통화 코드를 일관되게 사용합니다
- 기본값을 선택할 때 사용자 위치를 고려합니다

### 값 제약 조건
- 데이터 입력 오류를 방지하기 위해 합리적인 최소/최대 값을 설정합니다
- 부정적인 값을 허용하지 않는 필드에는 최소값으로 0을 사용합니다
- 최대값 설정 시 사용 사례를 고려합니다

### 다중 통화 프로젝트
- 보고를 위해 일관된 기본 통화를 사용합니다
- 자동 변환을 위해 CURRENCY_CONVERSION 필드를 구현합니다
- 각 필드에 대해 사용해야 할 통화를 문서화합니다

## 일반적인 사용 사례

1. **프로젝트 예산 책정**
   - 프로젝트 예산 추적
   - 비용 추정
   - 지출 추적

2. **판매 및 거래**
   - 거래 값
   - 계약 금액
   - 수익 추적

3. **재무 계획**
   - 투자 금액
   - 자금 조달 라운드
   - 재무 목표

4. **국제 비즈니스**
   - 다중 통화 가격 책정
   - 외환 추적
   - 국경 간 거래

## 제한 사항

- 표시를 위해 최대 2자리 소수 (더 많은 정밀도는 저장됨)
- 표준 CURRENCY 필드에서 내장된 통화 변환 없음
- 단일 필드 값에서 통화를 혼합할 수 없음
- 자동 환율 업데이트 없음 (이를 위해 CURRENCY_CONVERSION 사용)
- 통화 기호는 사용자 정의할 수 없음

## 관련 리소스

- [통화 변환 필드](/api/custom-fields/currency-conversion) - 자동 통화 변환
- [숫자 필드](/api/custom-fields/number) - 비금전적 숫자 값용
- [수식 필드](/api/custom-fields/formula) - 통화 값으로 계산
- [목록 사용자 정의 필드](/api/custom-fields/list-custom-fields) - 프로젝트의 모든 사용자 정의 필드 쿼리