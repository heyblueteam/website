---
title: 수식 사용자 정의 필드
description: 다른 데이터를 기반으로 값을 자동으로 계산하는 계산된 필드를 생성합니다.
---

수식 사용자 정의 필드는 Blue 내에서 차트 및 대시보드 계산에 사용됩니다. 이들은 사용자 정의 필드 데이터에서 작동하는 집계 함수(SUM, AVERAGE, COUNT 등)를 정의하여 차트에서 계산된 메트릭을 표시합니다. 수식은 개별 할 일 수준에서 계산되지 않고 시각화 목적으로 여러 레코드에 걸쳐 데이터를 집계합니다.

## 기본 예제

차트 계산을 위한 수식 필드를 생성합니다:

```graphql
mutation CreateFormulaField {
  createCustomField(input: {
    name: "Budget Total"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Budget)"
        html: "<span>SUM(Budget)</span>"
      }
      display: {
        type: NUMBER
        precision: 2
        function: SUM
      }
    }
  }) {
    id
    name
    type
    formula
  }
}
```

## 고급 예제

복잡한 계산을 포함하는 통화 수식을 생성합니다:

```graphql
mutation CreateCurrencyFormula {
  createCustomField(input: {
    name: "Profit Margin"
    type: FORMULA
    projectId: "proj_123"
    formula: {
      logic: {
        text: "SUM(Revenue) - SUM(Costs)"
        html: "<span>SUM(Revenue) - SUM(Costs)</span>"
      }
      display: {
        type: CURRENCY
        currency: {
          code: "USD"
          name: "US Dollar"
        }
        precision: 2
      }
    }
    description: "Automatically calculates profit by subtracting costs from revenue"
  }) {
    id
    name
    type
    formula
  }
}
```

## 입력 매개변수

### CreateCustomFieldInput

| 매개변수 | 유형 | 필수 | 설명 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 예 | 수식 필드의 표시 이름 |
| `type` | CustomFieldType! | ✅ 예 | `FORMULA` 여야 합니다. |
| `projectId` | String! | ✅ 예 | 이 필드가 생성될 프로젝트 ID |
| `formula` | JSON | 아니오 | 차트 계산을 위한 수식 정의 |
| `description` | String | 아니오 | 사용자에게 표시되는 도움말 텍스트 |

### 수식 구조

```json
{
  "logic": {
    "text": "Display text for the formula",
    "html": "HTML formatted display text"
  },
  "display": {
    "type": "NUMBER|CURRENCY|PERCENTAGE",
    "currency": {
      "code": "USD",
      "name": "US Dollar"  
    },
    "precision": 2,
    "function": "SUM|AVERAGE|AVERAGEA|COUNT|COUNTA|MAX|MIN"
  }
}
```

## 지원되는 함수

### 차트 집계 함수

수식 필드는 차트 계산을 위한 다음 집계 함수를 지원합니다:

| 함수 | 설명 | ChartFunction Enum |
|----------|-------------|-------------------|
| `SUM` | 모든 값의 합계 | `SUM` |
| `AVERAGE` | 숫자 값의 평균 | `AVERAGE` |
| `AVERAGEA` | 0과 null을 제외한 평균 | `AVERAGEA` |
| `COUNT` | 값의 개수 | `COUNT` |
| `COUNTA` | 0과 null을 제외한 개수 | `COUNTA` |
| `MAX` | 최대 값 | `MAX` |
| `MIN` | 최소 값 | `MIN` |

**참고**: 이러한 함수는 `display.function` 필드에서 사용되며 차트 시각화를 위한 집계 데이터에서 작동합니다. 복잡한 수학적 표현식이나 필드 수준 계산은 지원되지 않습니다.

## 표시 유형

### 숫자 표시

```json
{
  "display": {
    "type": "NUMBER",
    "precision": 2
  }
}
```

결과: `1250.75`

### 통화 표시

```json
{
  "display": {
    "type": "CURRENCY",
    "currency": {
      "code": "USD",
      "name": "US Dollar"
    },
    "precision": 2
  }
}
```

결과: `$1,250.75`

### 백분율 표시

```json
{
  "display": {
    "type": "PERCENTAGE",
    "precision": 1
  }
}
```

결과: `87.5%`

## 수식 필드 편집

기존 수식 필드를 업데이트합니다:

```graphql
mutation EditFormulaField {
  editCustomField(input: {
    customFieldId: "field_456"
    formula: {
      logic: {
        text: "AVERAGE(Score)"
        html: "<span>AVERAGE(Score)</span>"
      }
      display: {
        type: PERCENTAGE
        precision: 1
      }
    }
  }) {
    id
    formula
  }
}
```

## 수식 처리

### 차트 계산 컨텍스트

수식 필드는 차트 세그먼트 및 대시보드의 컨텍스트에서 처리됩니다:
- 차트가 렌더링되거나 업데이트될 때 계산이 발생합니다
- 결과는 `ChartSegment.formulaResult`에 소수 값으로 저장됩니다
- 처리는 'formula'라는 전용 BullMQ 큐를 통해 처리됩니다
- 업데이트는 대시보드 구독자에게 실시간 업데이트를 위해 게시됩니다

### 표시 형식

`getFormulaDisplayValue` 함수는 표시 유형에 따라 계산된 결과를 형식화합니다:
- **NUMBER**: 선택적 정밀도로 일반 숫자로 표시
- **PERCENTAGE**: 선택적 정밀도와 함께 % 접미사 추가  
- **CURRENCY**: 지정된 통화 코드를 사용하여 형식화

## 수식 결과 저장

결과는 `formulaResult` 필드에 저장됩니다:

```json
{
  "number": 1250.75,
  "formulaResult": {
    "number": 1250.75,
    "display": {
      "type": "CURRENCY",
      "currency": {
        "code": "USD",
        "name": "US Dollar"
      },
      "precision": 2
    }
  }
}
```

## 응답 필드

### TodoCustomField 응답

| 필드 | 유형 | 설명 |
|-------|------|-------------|
| `id` | String! | 필드 값에 대한 고유 식별자 |
| `customField` | CustomField! | 수식 필드 정의 |
| `number` | Float | 계산된 숫자 결과 |
| `formulaResult` | JSON | 표시 형식이 적용된 전체 결과 |
| `todo` | Todo! | 이 값이 속한 레코드 |
| `createdAt` | DateTime! | 값이 생성된 시점 |
| `updatedAt` | DateTime! | 값이 마지막으로 계산된 시점 |

## 데이터 컨텍스트

### 차트 데이터 소스

수식 필드는 차트 데이터 소스 컨텍스트 내에서 작동합니다:
- 수식은 프로젝트 내의 할 일에 걸쳐 사용자 정의 필드 값을 집계합니다
- `display.function`에 지정된 집계 함수가 계산을 결정합니다
- 결과는 SQL 집계 함수(avg, sum, count 등)를 사용하여 계산됩니다
- 효율성을 위해 데이터베이스 수준에서 계산이 수행됩니다

## 일반 수식 예제

### 총 예산 (차트 표시)

```json
{
  "logic": {
    "text": "Total Budget",
    "html": "<span>Total Budget</span>"
  },
  "display": {
    "type": "CURRENCY",
    "currency": { "code": "USD", "name": "US Dollar" },
    "precision": 2,
    "function": "SUM"
  }
}
```

### 평균 점수 (차트 표시)

```json
{
  "logic": {
    "text": "Average Quality Score",
    "html": "<span>Average Quality Score</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 1,
    "function": "AVERAGE"
  }
}
```

### 작업 수 (차트 표시)

```json
{
  "logic": {
    "text": "Total Tasks",
    "html": "<span>Total Tasks</span>"
  },
  "display": {
    "type": "NUMBER",
    "precision": 0,
    "function": "COUNT"
  }
}
```

## 필수 권한

사용자 정의 필드 작업은 표준 역할 기반 권한을 따릅니다:

| 작업 | 필요한 역할 |
|--------|---------------|
| Create formula field | Project member with appropriate role |
| Update formula field | Project member with appropriate role |
| View formula results | Project member with view permissions |
| Delete formula field | Project member with appropriate role |

**참고**: 필요한 특정 역할은 프로젝트의 사용자 정의 역할 구성에 따라 다릅니다. CUSTOM_FIELDS_CREATE와 같은 특별한 권한 상수는 없습니다.

## 오류 처리

### 유효성 검사 오류
```json
{
  "errors": [{
    "message": "Validation error message",
    "extensions": {
      "code": "VALIDATION_ERROR"
    }
  }]
}
```

### 사용자 정의 필드 찾을 수 없음
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

### 수식 설계
- 수식 필드에 대해 명확하고 설명적인 이름을 사용합니다
- 계산 논리를 설명하는 설명을 추가합니다
- 배포 전에 샘플 데이터로 수식을 테스트합니다
- 수식을 간단하고 읽기 쉽게 유지합니다

### 성능 최적화
- 깊게 중첩된 수식 종속성을 피합니다
- 와일드카드 대신 특정 필드 참조를 사용합니다
- 복잡한 계산을 위한 캐싱 전략을 고려합니다
- 대규모 프로젝트에서 수식 성능을 모니터링합니다

### 데이터 품질
- 수식을 사용하기 전에 원본 데이터를 검증합니다
- 빈 값이나 null 값을 적절하게 처리합니다
- 표시 유형에 적합한 정밀도를 사용합니다
- 계산에서 엣지 케이스를 고려합니다

## 일반 사용 사례

1. **재무 추적**
   - 예산 계산
   - 손익 계산서
   - 비용 분석
   - 수익 예측

2. **프로젝트 관리**
   - 완료 비율
   - 자원 활용
   - 일정 계산
   - 성과 메트릭

3. **품질 관리**
   - 평균 점수
   - 합격/불합격 비율
   - 품질 메트릭
   - 준수 추적

4. **비즈니스 인텔리전스**
   - KPI 계산
   - 추세 분석
   - 비교 메트릭
   - 대시보드 값

## 제한 사항

- 수식은 차트/대시보드 집계 전용이며, 할 일 수준 계산에는 사용되지 않습니다
- 지원되는 7개의 집계 함수(SUM, AVERAGE 등)로 제한됩니다
- 복잡한 수학적 표현식이나 필드 간 계산은 지원되지 않습니다
- 단일 수식에서 여러 필드를 참조할 수 없습니다
- 결과는 차트 및 대시보드에서만 볼 수 있습니다
- `logic` 필드는 표시 텍스트 전용이며 실제 계산 논리는 아닙니다

## 관련 리소스

- [숫자 필드](/api/5.custom%20fields/number) - 정적 숫자 값용
- [통화 필드](/api/5.custom%20fields/currency) - 금전적 값용
- [참조 필드](/api/5.custom%20fields/reference) - 프로젝트 간 데이터용
- [조회 필드](/api/5.custom%20fields/lookup) - 집계 데이터용
- [사용자 정의 필드 개요](/api/5.custom%20fields/2.list-custom-fields) - 일반 개념