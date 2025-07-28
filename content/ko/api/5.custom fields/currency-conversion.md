---
title: 통화 변환 사용자 정의 필드
description: 실시간 환율을 사용하여 통화 값을 자동으로 변환하는 필드를 생성합니다.
---

통화 변환 사용자 정의 필드는 소스 CURRENCY 필드의 값을 실시간 환율을 사용하여 다른 대상 통화로 자동 변환합니다. 이러한 필드는 소스 통화 값이 변경될 때마다 자동으로 업데이트됩니다.

변환 비율은 [Frankfurter API](https://github.com/hakanensari/frankfurter)에서 제공되며, 이는 [유럽 중앙은행](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html)에서 발표한 기준 환율을 추적하는 오픈 소스 서비스입니다. 이를 통해 귀하의 국제 비즈니스 요구에 대한 정확하고 신뢰할 수 있으며 최신의 통화 변환을 보장합니다.

## 기본 예제

간단한 통화 변환 필드를 생성합니다:

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

## 고급 예제

특정 날짜에 대한 역사적 비율로 변환 필드를 생성합니다:

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

## 전체 설정 프로세스

통화 변환 필드를 설정하려면 세 가지 단계가 필요합니다:

### 단계 1: 소스 CURRENCY 필드 생성

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

### 단계 2: CURRENCY_CONVERSION 필드 생성

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

### 단계 3: 변환 옵션 생성

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

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 변환 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `CURRENCY_CONVERSION` 여야 합니다. |
| `currencyFieldId` | String | 아니요 | 변환할 소스 CURRENCY 필드의 ID |
| `conversionDateType` | String | 아니요 | 환율에 대한 날짜 전략 (아래 참조) |
| `conversionDate` | String | 아니요 | 변환을 위한 날짜 문자열 (conversionDateType에 따라 다름) |
| `description` | String | 아니요 | 사용자에게 표시되는 도움말 텍스트 |

**참고**: 사용자 정의 필드는 사용자의 현재 프로젝트 컨텍스트에 따라 프로젝트와 자동으로 연결됩니다. `projectId` 매개변수는 필요하지 않습니다.

### 변환 날짜 유형

| 유형 | 설명 | conversionDate 매개변수 |
|------|-------------|-------------------------|
| `currentDate` | 실시간 환율 사용 | 필요하지 않음 |
| `specificDate` | 고정 날짜의 비율 사용 | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | 다른 필드의 날짜 사용 | "todoDueDate" or DATE field ID |

## 변환 옵션 생성

변환 옵션은 어떤 통화 쌍이 변환될 수 있는지를 정의합니다:

### CreateCustomFieldOptionInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ 예 | CURRENCY_CONVERSION 필드의 ID |
| `title` | String! | ✅ 예 | 이 변환 옵션의 표시 이름 |
| `currencyConversionFrom` | String! | ✅ 예 | 소스 통화 코드 또는 "모두" |
| `currencyConversionTo` | String! | ✅ 예 | 대상 통화 코드 |

### "모두"를 소스로 사용

특별 값 "모두"는 `currencyConversionFrom`로 대체 옵션을 생성합니다:

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

이 옵션은 특정 통화 쌍 일치가 발견되지 않을 때 사용됩니다.

## 자동 변환 작동 방식

1. **값 업데이트**: 소스 CURRENCY 필드에 값이 설정될 때
2. **옵션 일치**: 시스템이 소스 통화에 따라 일치하는 변환 옵션을 찾습니다.
3. **환율 가져오기**: Frankfurter API에서 환율을 검색합니다.
4. **계산**: 소스 금액에 환율을 곱합니다.
5. **저장**: 변환된 값을 대상 통화 코드와 함께 저장합니다.

### 예제 흐름

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

## 날짜 기반 변환

### 현재 날짜 사용

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

변환은 소스 값이 변경될 때마다 현재 환율로 업데이트됩니다.

### 특정 날짜 사용

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

항상 지정된 날짜의 환율을 사용합니다.

### 필드에서 날짜 사용

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

다른 필드(할 일 기한 또는 DATE 사용자 정의 필드)에서 날짜를 사용합니다.

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값의 고유 식별자 |
| `customField` | CustomField! | 변환 필드 정의 |
| `number` | Float | 변환된 금액 |
| `currency` | String | 대상 통화 코드 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시간 |
| `updatedAt` | DateTime! | 값이 마지막으로 업데이트된 시간 |

## 환율 출처

Blue는 환율에 대해 **Frankfurter API**를 사용합니다:
- 유럽 중앙은행이 호스팅하는 오픈 소스 API
- 공식 환율로 매일 업데이트
- 1999년부터의 역사적 비율 지원
- 비즈니스 사용에 무료 및 신뢰할 수 있음

## 오류 처리

### 변환 실패

변환이 실패할 경우 (API 오류, 유효하지 않은 통화 등):
- 변환된 값은 `0`로 설정됩니다.
- 대상 통화는 여전히 저장됩니다.
- 사용자에게 오류가 발생하지 않습니다.

### 일반 시나리오

| 시나리오 | 결과 |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| 일치하는 옵션 없음 | Uses "Any" option if available |
| Missing source value | 변환이 수행되지 않음 |

## 필요한 권한

사용자 정의 필드 관리는 프로젝트 수준의 접근이 필요합니다:

| 역할 | 필드 생성/업데이트 가능 |
|------|-------------------------|
| `OWNER` | ✅ 예 |
| `ADMIN` | ✅ 예 |
| `MEMBER` | ❌ 아니요 |
| `CLIENT` | ❌ 아니요 |

변환된 값에 대한 보기 권한은 표준 레코드 접근 규칙을 따릅니다.

## 모범 사례

### 옵션 구성
- 일반 변환을 위한 특정 통화 쌍을 생성합니다.
- 유연성을 위한 "모두" 대체 옵션을 추가합니다.
- 옵션에 대한 설명적인 제목을 사용합니다.

### 날짜 전략 선택
- 실시간 재무 추적을 위해 `currentDate`를 사용합니다.
- 역사적 보고를 위해 `specificDate`를 사용합니다.
- 거래 특정 비율을 위해 `fromDateField`를 사용합니다.

### 성능 고려 사항
- 여러 변환 필드가 병렬로 업데이트됩니다.
- 소스 값이 변경될 때만 API 호출이 이루어집니다.
- 동일 통화 변환은 API 호출을 건너뜁니다.

## 일반 사용 사례

1. **다중 통화 프로젝트**
   - 프로젝트 비용을 현지 통화로 추적합니다.
   - 회사 통화로 총 예산을 보고합니다.
   - 지역 간 값을 비교합니다.

2. **국제 판매**
   - 거래 값을 보고 통화로 변환합니다.
   - 여러 통화로 수익을 추적합니다.
   - 종료된 거래에 대한 역사적 변환.

3. **재무 보고**
   - 기간 종료 통화 변환.
   - 통합 재무 제표.
   - 현지 통화로 예산 대비 실제.

4. **계약 관리**
   - 계약 체결 날짜에 계약 값을 변환합니다.
   - 여러 통화로 지불 일정을 추적합니다.
   - 통화 위험 평가.

## 제한 사항

- 암호화폐 변환을 지원하지 않습니다.
- 변환된 값을 수동으로 설정할 수 없습니다 (항상 계산됨).
- 모든 변환 금액에 대해 고정 2자리 소수점 정확도.
- 사용자 정의 환율을 지원하지 않습니다.
- 환율 캐싱을 지원하지 않습니다 (각 변환에 대해 새 API 호출).
- Frankfurter API의 가용성에 따라 다릅니다.

## 관련 리소스

- [통화 필드](/api/custom-fields/currency) - 변환을 위한 소스 필드
- [날짜 필드](/api/custom-fields/date) - 날짜 기반 변환을 위한 필드
- [수식 필드](/api/custom-fields/formula) - 대체 계산
- [사용자 정의 필드 개요](/custom-fields/list-custom-fields) - 일반 개념