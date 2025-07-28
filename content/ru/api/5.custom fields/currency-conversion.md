---
title: Пользовательское поле конверсии валюты
description: Создавайте поля, которые автоматически конвертируют значения валюты с использованием актуальных курсов обмена
---

Пользовательские поля конверсии валюты автоматически конвертируют значения из исходного поля ВАЛЮТА в различные целевые валюты, используя актуальные курсы обмена. Эти поля обновляются автоматически каждый раз, когда изменяется значение исходной валюты.

Курсы конверсии предоставляются [Frankfurter API](https://github.com/hakanensari/frankfurter), открытым сервисом, который отслеживает справочные курсы обмена, публикуемые [Европейским центральным банком](https://www.ecb.europa.eu/stats/policy_and_exchange_rates/euro_reference_exchange_rates/html/index.en.html). Это обеспечивает точные, надежные и актуальные конверсии валют для ваших международных бизнес-потребностей.

## Простой пример

Создайте простое поле конверсии валюты:

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

## Расширенный пример

Создайте поле конверсии с конкретной датой для исторических курсов:

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

## Полный процесс настройки

Настройка поля конверсии валюты требует три шага:

### Шаг 1: Создайте исходное поле ВАЛЮТА

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

### Шаг 2: Создайте поле CURRENCY_CONVERSION

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

### Шаг 3: Создайте параметры конверсии

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

## Входные параметры

### CreateCustomFieldInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Да | Отображаемое имя поля конверсии |
| `type` | CustomFieldType! | ✅ Да | Должен быть `CURRENCY_CONVERSION` |
| `currencyFieldId` | String | Нет | ID исходного поля ВАЛЮТА, из которого нужно конвертировать |
| `conversionDateType` | String | Нет | Стратегия даты для курсов обмена (см. ниже) |
| `conversionDate` | String | Нет | Строка даты для конверсии (на основе conversionDateType) |
| `description` | String | Нет | Текст помощи, отображаемый пользователям |

**Примечание**: Пользовательские поля автоматически ассоциируются с проектом на основе текущего контекста проекта пользователя. Параметр `projectId` не требуется.

### Типы дат конверсии

| Тип | Описание | Параметр conversionDate |
|------|-------------|-------------------------|
| `currentDate` | Использует актуальные курсы обмена | Не требуется |
| `specificDate` | Использует курсы с фиксированной даты | ISO date string (e.g., "2024-01-01T00:00:00Z") |
| `fromDateField` | Использует дату из другого поля | "todoDueDate" or DATE field ID |

## Создание параметров конверсии

Параметры конверсии определяют, какие валютные пары могут быть конвертированы:

### CreateCustomFieldOptionInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `customFieldId` | String! | ✅ Да | ID поля CURRENCY_CONVERSION |
| `title` | String! | ✅ Да | Отображаемое имя для этого варианта конверсии |
| `currencyConversionFrom` | String! | ✅ Да | Код исходной валюты или "Любая" |
| `currencyConversionTo` | String! | ✅ Да | Код целевой валюты |

### Использование "Любая" в качестве источника

Специальное значение "Любая" как `currencyConversionFrom` создает резервный вариант:

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

Этот вариант будет использоваться, когда не будет найдено конкретного соответствия валютной пары.

## Как работает автоматическая конверсия

1. **Обновление значения**: Когда значение устанавливается в исходном поле ВАЛЮТА
2. **Сопоставление варианта**: Система находит соответствующий вариант конверсии на основе исходной валюты
3. **Получение курса**: Извлекает курс обмена из Frankfurter API
4. **Расчет**: Умножает исходную сумму на курс обмена
5. **Хранение**: Сохраняет конвертированное значение с кодом целевой валюты

### Пример потока

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

## Конверсии на основе даты

### Использование текущей даты

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

Конверсии обновляются с актуальными курсами обмена каждый раз, когда изменяется исходное значение.

### Использование конкретной даты

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

Всегда использует курсы обмена с указанной даты.

### Использование даты из поля

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

Использует дату из другого поля (либо срок выполнения задачи, либо пользовательское поле ДАТА).

## Поля ответа

### Ответ TodoCustomField

| Поле | Тип | Описание |
|-------|------|-------------|
| `id` | String! | Уникальный идентификатор для значения поля |
| `customField` | CustomField! | Определение поля конверсии |
| `number` | Float | Конвертированная сумма |
| `currency` | String | Код целевой валюты |
| `todo` | Todo! | Запись, к которой принадлежит это значение |
| `createdAt` | DateTime! | Когда было создано значение |
| `updatedAt` | DateTime! | Когда значение было в последний раз обновлено |

## Источник курсов обмена

Blue использует **Frankfurter API** для курсов обмена:
- Открытый API, размещенный Европейским центральным банком
- Обновляется ежедневно с официальными курсами обмена
- Поддерживает исторические курсы с 1999 года
- Бесплатен и надежен для бизнес-использования

## Обработка ошибок

### Ошибки конверсии

Когда конверсия не удалась (ошибка API, недопустимая валюта и т. д.):
- Конвертированное значение устанавливается в `0`
- Целевая валюта все еще сохраняется
- Пользователю не выдается ошибка

### Общие сценарии

| Сценарий | Результат |
|----------|---------|
| Same currency (USD→USD) | Value copied without API call |
| Invalid currency code | Conversion returns 0 |
| API unavailable | Conversion returns 0 |
| Нет подходящего варианта | Uses "Any" option if available |
| Missing source value | Конверсия не выполнена |

## Необходимые разрешения

Управление пользовательскими полями требует доступа на уровне проекта:

| Роль | Может создавать/обновлять поля |
|------|-------------------------|
| `OWNER` | ✅ Да |
| `ADMIN` | ✅ Да |
| `MEMBER` | ❌ Нет |
| `CLIENT` | ❌ Нет |

Разрешения на просмотр конвертированных значений следуют стандартным правилам доступа к записям.

## Рекомендации по лучшим практикам

### Конфигурация параметров
- Создавайте конкретные валютные пары для распространенных конверсий
- Добавьте резервный вариант "Любая" для гибкости
- Используйте описательные названия для параметров

### Выбор стратегии даты
- Используйте `currentDate` для отслеживания финансов в реальном времени
- Используйте `specificDate` для исторической отчетности
- Используйте `fromDateField` для ставок, специфичных для транзакций

### Учет производительности
- Несколько полей конверсии обновляются параллельно
- Вызовы API выполняются только при изменении исходного значения
- Конверсии в одной валюте пропускают вызовы API

## Общие случаи использования

1. **Мультивалютные проекты**
   - Отслеживание затрат проекта в местных валютах
   - Отчет о общем бюджете в валюте компании
   - Сравнение значений по регионам

2. **Международные продажи**
   - Конвертация значений сделок в валюту отчетности
   - Отслеживание доходов в нескольких валютах
   - Историческая конверсия для закрытых сделок

3. **Финансовая отчетность**
   - Конверсии валют в конце периода
   - Консолидированные финансовые отчеты
   - Бюджет против фактических данных в местной валюте

4. **Управление контрактами**
   - Конвертация значений контрактов на дату подписания
   - Отслеживание графиков платежей в нескольких валютах
   - Оценка валютных рисков

## Ограничения

- Нет поддержки конверсий криптовалют
- Невозможно вручную устанавливать конвертированные значения (всегда рассчитываются)
- Фиксированная точность 2 десятичных знака для всех конвертированных сумм
- Нет поддержки пользовательских курсов обмена
- Нет кэширования курсов обмена (свежий вызов API для каждой конверсии)
- Зависит от доступности Frankfurter API

## Связанные ресурсы

- [Поля валюты](/api/custom-fields/currency) - Исходные поля для конверсий
- [Поля даты](/api/custom-fields/date) - Для конверсий на основе даты
- [Поля формул](/api/custom-fields/formula) - Альтернативные расчеты
- [Обзор пользовательских полей](/custom-fields/list-custom-fields) - Общие концепции