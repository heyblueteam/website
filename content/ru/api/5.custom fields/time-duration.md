---
title: Пользовательское поле продолжительности времени
description: Создайте вычисляемые поля продолжительности времени, которые отслеживают время между событиями в вашем рабочем процессе
---

Пользовательские поля продолжительности времени автоматически вычисляют и отображают продолжительность между двумя событиями в вашем рабочем процессе. Они идеально подходят для отслеживания времени обработки, времени ответа, времени цикла или любых метрик, основанных на времени, в ваших проектах.

## Простой пример

Создайте простое поле продолжительности времени, которое отслеживает, сколько времени требуется для выполнения задач:

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

## Продвинутый пример

Создайте сложное поле продолжительности времени, которое отслеживает время между изменениями пользовательских полей с целевым SLA:

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

## Входные параметры

### CreateCustomFieldInput (TIME_DURATION)

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Да | Отображаемое имя поля продолжительности |
| `type` | CustomFieldType! | ✅ Да | Должно быть `TIME_DURATION` |
| `description` | String | Нет | Текст помощи, отображаемый пользователям |
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType! | ✅ Да | Как отображать продолжительность |
| `timeDurationStartInput` | CustomFieldTimeDurationInput! | ✅ Да | Конфигурация начального события |
| `timeDurationEndInput` | CustomFieldTimeDurationInput! | ✅ Да | Конфигурация конечного события |
| `timeDurationTargetTime` | Float | Нет | Целевая продолжительность в секундах для мониторинга SLA |

### CustomFieldTimeDurationInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `type` | CustomFieldTimeDurationType! | ✅ Да | Тип события для отслеживания |
| `condition` | CustomFieldTimeDurationCondition! | ✅ Да | `FIRST` или `LAST` событие |
| `customFieldId` | String | Conditional | Обязательный для типа `TODO_CUSTOM_FIELD` |
| `customFieldOptionIds` | [String!] | Conditional | Обязательный для изменений полей выбора |
| `todoListId` | String | Conditional | Обязательный для типа `TODO_MOVED` |
| `tagId` | String | Conditional | Обязательный для типа `TODO_TAG_ADDED` |
| `assigneeId` | String | Conditional | Обязательный для типа `TODO_ASSIGNEE_ADDED` |

### Значения CustomFieldTimeDurationType

| Значение | Описание |
|-------|-------------|
| `TODO_CREATED_AT` | Когда запись была создана |
| `TODO_CUSTOM_FIELD` | Когда изменилось значение пользовательского поля |
| `TODO_DUE_DATE` | Когда была установлена дата выполнения |
| `TODO_MARKED_AS_COMPLETE` | Когда запись была отмечена как завершенная |
| `TODO_MOVED` | Когда запись была перемещена в другой список |
| `TODO_TAG_ADDED` | Когда к записи был добавлен тег |
| `TODO_ASSIGNEE_ADDED` | Когда к записи был добавлен исполнитель |

### Значения CustomFieldTimeDurationCondition

| Значение | Описание |
|-------|-------------|
| `FIRST` | Использовать первое вхождение события |
| `LAST` | Использовать последнее вхождение события |

### Значения CustomFieldTimeDurationDisplayType

| Значение | Описание | Пример |
|-------|-------------|---------|
| `FULL_DATE` | Формат Дни:Часы:Минуты:Секунды | `"01:02:03:04"` |
| `FULL_DATE_STRING` | Записано полностью словами | `"Two hours, two minutes, three seconds"` |
| `FULL_DATE_SUBSTRING` | Числовой с единицами | `"1 hour, 2 minutes, 3 seconds"` |
| `DAYS` | Продолжительность только в днях | `"2.5"` (2.5 days) |
| `HOURS` | Продолжительность только в часах | `"60"` (60 hours) |
| `MINUTES` | Продолжительность только в минутах | `"3600"` (3600 minutes) |
| `SECONDS` | Продолжительность только в секундах | `"216000"` (216000 seconds) |

## Поля ответа

### Ответ TodoCustomField

| Поле | Тип | Описание |
|-------|------|-------------|
| `id` | String! | Уникальный идентификатор для значения поля |
| `customField` | CustomField! | Определение пользовательского поля |
| `number` | Float | Продолжительность в секундах |
| `value` | Float | Псевдоним для числа (продолжительность в секундах) |
| `todo` | Todo! | Запись, к которой принадлежит это значение |
| `createdAt` | DateTime! | Когда значение было создано |
| `updatedAt` | DateTime! | Когда значение было в последний раз обновлено |

### Ответ CustomField (TIME_DURATION)

| Поле | Тип | Описание |
|-------|------|-------------|
| `timeDurationDisplay` | CustomFieldTimeDurationDisplayType | Формат отображения для продолжительности |
| `timeDurationStart` | CustomFieldTimeDuration | Конфигурация начального события |
| `timeDurationEnd` | CustomFieldTimeDuration | Конфигурация конечного события |
| `timeDurationTargetTime` | Float | Целевая продолжительность в секундах (для мониторинга SLA) |

## Вычисление продолжительности

### Как это работает
1. **Начальное событие**: Система отслеживает указанное начальное событие
2. **Конечное событие**: Система отслеживает указанное конечное событие
3. **Вычисление**: Продолжительность = Время окончания - Время начала
4. **Хранение**: Продолжительность хранится в секундах как число
5. **Отображение**: Отформатировано в соответствии с настройкой `timeDurationDisplay`

### Триггеры обновления
Значения продолжительности автоматически пересчитываются, когда:
- Записи создаются или обновляются
- Изменяются значения пользовательских полей
- Теги добавляются или удаляются
- Исполнители добавляются или удаляются
- Записи перемещаются между списками
- Записи отмечаются как завершенные/незавершенные

## Чтение значений продолжительности

### Запрос полей продолжительности
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

### Отформатированные значения отображения
Значения продолжительности автоматически форматируются в зависимости от настройки `timeDurationDisplay`:

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

## Общие примеры конфигурации

### Время завершения задачи
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

### Продолжительность изменения статуса
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

### Время в конкретном списке
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

### Время ответа на назначение
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

## Необходимые разрешения

| Действие | Необходимое разрешение |
|--------|-------------------|
| Create duration field | Project-level `OWNER` or `ADMIN` role |
| Update duration field | Project-level `OWNER` or `ADMIN` role |
| View duration value | Any project member role |

## Ответы на ошибки

### Неверная конфигурация
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

### Ссылающееся поле не найдено
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

### Отсутствуют обязательные параметры
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

## Важные заметки

### Автоматическое вычисление
- Поля продолжительности являются **только для чтения** - значения вычисляются автоматически
- Вы не можете вручную устанавливать значения продолжительности через API
- Вычисления происходят асинхронно через фоновые задания
- Значения обновляются автоматически, когда происходят триггерные события

### Соображения по производительности
- Вычисления продолжительности ставятся в очередь и обрабатываются асинхронно
- Большое количество полей продолжительности может повлиять на производительность
- Учитывайте частоту триггерных событий при проектировании полей продолжительности
- Используйте конкретные условия, чтобы избежать ненужных пересчетов

### Нулевые значения
Поля продолжительности будут показывать `null`, когда:
- Начальное событие еще не произошло
- Конечное событие еще не произошло
- Конфигурация ссылается на несуществующие сущности
- Вычисление сталкивается с ошибкой

## Лучшие практики

### Проектирование конфигурации
- Используйте конкретные типы событий, а не общие, когда это возможно
- Выбирайте подходящие условия `FIRST` против `LAST` в зависимости от вашего рабочего процесса
- Тестируйте вычисления продолжительности с образцовыми данными перед развертыванием
- Документируйте логику вашего поля продолжительности для членов команды

### Форматирование отображения
- Используйте `FULL_DATE_SUBSTRING` для наиболее читаемого формата
- Используйте `FULL_DATE` для компактного, постоянного отображения ширины
- Используйте `FULL_DATE_STRING` для официальных отчетов и документов
- Используйте `DAYS`, `HOURS`, `MINUTES` или `SECONDS` для простых числовых отображений
- Учитывайте ограничения пространства вашего пользовательского интерфейса при выборе формата

### Мониторинг SLA с целевым временем
При использовании `timeDurationTargetTime`:
- Установите целевую продолжительность в секундах
- Сравните фактическую продолжительность с целевой для соблюдения SLA
- Используйте в панелях мониторинга для выделения просроченных элементов
- Пример: SLA на ответ 24 часа = 86400 секунд

### Интеграция в рабочий процесс
- Проектируйте поля продолжительности, чтобы они соответствовали вашим реальным бизнес-процессам
- Используйте данные о продолжительности для улучшения и оптимизации процессов
- Мониторьте тенденции продолжительности, чтобы выявить узкие места в рабочем процессе
- Настройте оповещения о порогах продолжительности, если это необходимо

## Общие случаи использования

1. **Производительность процессов**
   - Время завершения задач
   - Время цикла проверки
   - Время обработки одобрения
   - Время ответа

2. **Мониторинг SLA**
   - Время до первого ответа
   - Время разрешения
   - Сроки эскалации
   - Соблюдение уровня обслуживания

3. **Аналитика рабочего процесса**
   - Выявление узких мест
   - Оптимизация процессов
   - Метрики производительности команды
   - Тайминг контроля качества

4. **Управление проектами**
   - Продолжительности фаз
   - Тайминг вех
   - Время распределения ресурсов
   - Сроки доставки

## Ограничения

- Поля продолжительности являются **только для чтения** и не могут быть установлены вручную
- Значения вычисляются асинхронно и могут быть недоступны немедленно
- Требуется правильная настройка триггеров событий в вашем рабочем процессе
- Невозможно вычислить продолжительности для событий, которые еще не произошли
- Ограничено отслеживанием времени между дискретными событиями (не для непрерывного отслеживания времени)
- Нет встроенных оповещений или уведомлений SLA
- Невозможно агрегировать несколько вычислений продолжительности в одно поле

## Связанные ресурсы

- [Числовые поля](/api/custom-fields/number) - Для ручных числовых значений
- [Поля даты](/api/custom-fields/date) - Для отслеживания конкретных дат
- [Обзор пользовательских полей](/api/custom-fields/list-custom-fields) - Общие концепции
- [Автоматизация](/api/automations) - Для триггеров действий на основе порогов продолжительности