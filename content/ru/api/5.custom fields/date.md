---
title: Пользовательское поле даты
description: Создайте поля даты для отслеживания одиночных дат или диапазонов дат с поддержкой часовых поясов
---

Пользовательские поля даты позволяют вам хранить одиночные даты или диапазоны дат для записей. Они поддерживают обработку часовых поясов, интеллектуальное форматирование и могут использоваться для отслеживания сроков, дат событий или любой информации, связанной со временем.

## Простой пример

Создайте простое поле даты:

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

## Расширенный пример

Создайте поле срока с описанием:

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

## Входные параметры

### CreateCustomFieldInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Да | Отображаемое имя поля даты |
| `type` | CustomFieldType! | ✅ Да | Должно быть `DATE` |
| `isDueDate` | Boolean | Нет | Является ли это поле сроком |
| `description` | String | Нет | Текст помощи, отображаемый пользователям |

**Примечание**: Пользовательские поля автоматически ассоциируются с проектом на основе текущего контекста проекта пользователя. Параметр `projectId` не требуется.

## Установка значений даты

Поля даты могут хранить либо одиночную дату, либо диапазон дат:

### Одиночная дата

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

### Диапазон дат

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

### Целый день

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

### Параметры SetTodoCustomFieldInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Да | ID записи для обновления |
| `customFieldId` | String! | ✅ Да | ID пользовательского поля даты |
| `startDate` | DateTime | Нет | Дата/время начала в формате ISO 8601 |
| `endDate` | DateTime | Нет | Дата/время окончания в формате ISO 8601 |
| `timezone` | String | Нет | Идентификатор часового пояса (например, "America/New_York") |

**Примечание**: Если предоставлено только `startDate`, то `endDate` автоматически устанавливается на то же значение.

## Форматы даты

### Формат ISO 8601
Все даты должны быть предоставлены в формате ISO 8601:
- `2025-01-15T14:30:00Z` - время UTC
- `2025-01-15T14:30:00+05:00` - с учетом смещения часового пояса
- `2025-01-15T14:30:00.123Z` - с миллисекундами

### Идентификаторы часовых поясов
Используйте стандартные идентификаторы часовых поясов:
- `America/New_York`
- `Europe/London`
- `Asia/Tokyo`
- `Australia/Sydney`

Если часовой пояс не предоставлен, система по умолчанию использует обнаруженный часовой пояс пользователя.

## Создание записей с значениями даты

При создании новой записи со значениями даты:

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

### Поддерживаемые форматы ввода

При создании записей даты могут быть предоставлены в различных форматах:

| Формат | Пример | Результат |
|--------|---------|---------|
| ISO Date | `"2025-01-15"` | Single date (start and end same) |
| ISO DateTime | `"2025-01-15T10:00:00Z"` | Single date/time |
| Date Range | `"2025-01-01,2025-01-31"` | Start and end dates |

## Поля ответа

### Ответ TodoCustomField

| Поле | Тип | Описание |
|-------|------|-------------|
| `id` | ID! | Уникальный идентификатор для значения поля |
| `uid` | String! | Уникальная строка идентификатора |
| `customField` | CustomField! | Определение пользовательского поля (содержит значения даты) |
| `todo` | Todo! | Запись, к которой принадлежит это значение |
| `createdAt` | DateTime! | Когда значение было создано |
| `updatedAt` | DateTime! | Когда значение было в последний раз изменено |

**Важно**: Значения даты (`startDate`, `endDate`, `timezone`) доступны через поле `customField.value`, а не напрямую в TodoCustomField.

### Структура объекта значения

Значения даты возвращаются через поле `customField.value` как объект JSON:

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

**Примечание**: Поле `value` имеет тип `CustomField`, а не `TodoCustomField`.

## Запрос значений даты

При запросе записей с пользовательскими полями даты доступ к значениям даты осуществляется через поле `customField.value`:

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

Ответ будет включать значения даты в поле `value`:

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

## Интеллектуальное отображение даты

Система автоматически форматирует даты на основе диапазона:

| Сценарий | Формат отображения |
|----------|----------------|
| Single date | `Jan 15, 2025` |
| All-day event | `Jan 15, 2025` (время не отображается) |
| Same day with times | `Jan 15, 2025 10:00 AM - 5:00 PM` |
| Multi-day range | `Jan 1 → Jan 31, 2025` |

**Обнаружение целого дня**: События с 00:00 до 23:59 автоматически определяются как события на весь день.

## Обработка часовых поясов

### Хранение
- Все даты хранятся в UTC в базе данных
- Информация о часовом поясе сохраняется отдельно
- Конвертация происходит при отображении

### Рекомендации
- Всегда указывайте часовой пояс для точности
- Используйте согласованные часовые пояса в рамках проекта
- Учитывайте местоположение пользователей для глобальных команд

### Общие часовые пояса

| Регион | Идентификатор часового пояса | Смещение UTC |
|--------|-------------|------------|
| US Eastern | `America/New_York` | UTC-5/-4 |
| US Pacific | `America/Los_Angeles` | UTC-8/-7 |
| UK | `Europe/London` | UTC+0/+1 |
| EU Central | `Europe/Berlin` | UTC+1/+2 |
| Japan | `Asia/Tokyo` | UTC+9 |
| Australia Eastern | `Australia/Sydney` | UTC+10/+11 |

## Фильтрация и запросы

Поля даты поддерживают сложную фильтрацию:

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

### Проверка на пустые поля даты

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

### Поддерживаемые операторы

| Оператор | Использование | Описание |
|----------|-------|-------------|
| `EQ` | С диапазоном дат | Дата пересекается с указанным диапазоном (любое пересечение) |
| `NE` | С диапазоном дат | Дата не пересекается с диапазоном |
| `IS` | С `values: null` | Поле даты пустое (startDate или endDate равно null) |
| `NOT` | С `values: null` | Поле даты имеет значение (обе даты не равны null) |

## Необходимые разрешения

| Действие | Необходимое разрешение |
|--------|-------------------|
| Create date field | `OWNER` or `ADMIN` role at company or project level |
| Update date field | `OWNER` or `ADMIN` role at company or project level |
| Set date value | Standard record edit permissions |
| View date value | Standard record view permissions |

## Ответы об ошибках

### Неверный формат даты
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

### Поле не найдено
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


## Ограничения

- Нет поддержки повторяющихся дат (используйте автоматизацию для повторяющихся событий)
- Нельзя установить время без даты
- Нет встроенного расчета рабочих дней
- Диапазоны дат не проверяют автоматически, что конец > начало
- Максимальная точность - до секунды (без хранения миллисекунд)

## Связанные ресурсы

- [Обзор пользовательских полей](/api/custom-fields/list-custom-fields) - Общие концепции пользовательских полей
- [API автоматизации](/api/automations/index) - Создание автоматизаций на основе дат