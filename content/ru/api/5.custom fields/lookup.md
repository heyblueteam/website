---
title: Поиск пользовательского поля
description: Создайте поля поиска, которые автоматически извлекают данные из ссылочных записей
---

Пользовательские поля поиска автоматически извлекают данные из записей, на которые ссылаются [Ссылочные поля](/api/custom-fields/reference), отображая информацию из связанных записей без ручного копирования. Они обновляются автоматически при изменении ссылочных данных.

## Простой пример

Создайте поле поиска для отображения тегов из ссылочных записей:

```graphql
mutation CreateLookupField {
  createCustomField(input: {
    name: "Related Todo Tags"
    type: LOOKUP
    lookupOption: {
      referenceId: "reference_field_id"
      lookupType: TODO_TAG
    }
    description: "Tags from related todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Расширенный пример

Создайте поле поиска для извлечения значений пользовательских полей из ссылочных записей:

```graphql
mutation CreateCustomFieldLookup {
  createCustomField(input: {
    name: "Referenced Budget Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "project_reference_field_id"
      lookupId: "budget_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
    description: "Budget values from referenced todos"
  }) {
    id
    name
    type
    lookupOption
  }
}
```

## Входные параметры

### CreateCustomFieldInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|--------------|-------------|
| `name` | String! | ✅ Да | Отображаемое имя поля поиска |
| `type` | CustomFieldType! | ✅ Да | Должно быть `LOOKUP` |
| `lookupOption` | CustomFieldLookupOptionInput! | ✅ Да | Конфигурация поиска |
| `description` | String | Нет | Текст помощи, отображаемый пользователям |

## Конфигурация поиска

### CustomFieldLookupOptionInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|--------------|-------------|
| `referenceId` | String! | ✅ Да | ID ссылочного поля для извлечения данных |
| `lookupId` | String | Нет | ID конкретного пользовательского поля для поиска (обязательно для типа TODO_CUSTOM_FIELD) |
| `lookupType` | CustomFieldLookupType! | ✅ Да | Тип данных для извлечения из ссылочных записей |

## Типы поиска

### Значения CustomFieldLookupType

| Тип | Описание | Возвращает |
|------|-------------|---------|
| `TODO_DUE_DATE` | Сроки выполнения из ссылочных задач | Array of date objects with start/end dates and timezone |
| `TODO_CREATED_AT` | Даты создания из ссылочных задач | Array of creation timestamps |
| `TODO_UPDATED_AT` | Даты последнего обновления из ссылочных задач | Array of update timestamps |
| `TODO_TAG` | Теги из ссылочных задач | Array of tag objects with id, name, and color |
| `TODO_ASSIGNEE` | Исполнители из ссылочных задач | Array of user objects |
| `TODO_DESCRIPTION` | Описания из ссылочных задач | Array of text descriptions (empty values filtered out) |
| `TODO_LIST` | Имена списков задач из ссылочных задач | Array of list titles |
| `TODO_CUSTOM_FIELD` | Значения пользовательских полей из ссылочных задач | Array of values based on the field type |

## Поля ответа

### Ответ CustomField (для полей поиска)

| Поле | Тип | Описание |
|-------|------|-------------|
| `id` | String! | Уникальный идентификатор поля |
| `name` | String! | Отображаемое имя поля поиска |
| `type` | CustomFieldType! | Будет `LOOKUP` |
| `customFieldLookupOption` | CustomFieldLookupOption | Конфигурация поиска и результаты |
| `createdAt` | DateTime! | Когда поле было создано |
| `updatedAt` | DateTime! | Когда поле было в последний раз обновлено |

### Структура CustomFieldLookupOption

| Поле | Тип | Описание |
|-------|------|-------------|
| `lookupType` | CustomFieldLookupType! | Тип выполняемого поиска |
| `lookupResult` | JSON | Извлеченные данные из ссылочных записей |
| `reference` | CustomField | Ссылочное поле, используемое в качестве источника |
| `lookup` | CustomField | Конкретное поле, которое ищется (для TODO_CUSTOM_FIELD) |
| `parentCustomField` | CustomField | Родительское поле поиска |
| `parentLookup` | CustomField | Родительский поиск в цепочке (для вложенных поисков) |

## Как работают поиски

1. **Извлечение данных**: Поиски извлекают конкретные данные из всех записей, связанных через ссылочное поле
2. **Автоматические обновления**: Когда ссылочные записи изменяются, значения поиска обновляются автоматически
3. **Только для чтения**: Поля поиска нельзя редактировать напрямую - они всегда отражают текущие ссылочные данные
4. **Без вычислений**: Поиски извлекают и отображают данные как есть без агрегаций или вычислений

## Поиски TODO_CUSTOM_FIELD

При использовании типа `TODO_CUSTOM_FIELD` вы должны указать, какое пользовательское поле извлекать, используя параметр `lookupId`:

```graphql
mutation CreateCustomFieldValueLookup {
  createCustomField(input: {
    name: "Project Status Values"
    type: LOOKUP
    lookupOption: {
      referenceId: "linked_projects_reference_field"
      lookupId: "status_custom_field_id"
      lookupType: TODO_CUSTOM_FIELD
    }
  }) {
    id
  }
}
```

Это извлекает значения указанного пользовательского поля из всех ссылочных записей.

## Запрос данных поиска

```graphql
query GetLookupValues {
  todo(id: "todo_123") {
    customFields {
      id
      customField {
        name
        type
        customFieldLookupOption {
          lookupType
          lookupResult
          reference {
            id
            name
          }
          lookup {
            id
            name
            type
          }
        }
      }
    }
  }
}
```

## Примеры результатов поиска

### Результат поиска тегов
```json
{
  "lookupResult": [
    {
      "id": "tag_123",
      "title": "urgent",
      "color": "#ff0000"
    },
    {
      "id": "tag_456",
      "title": "development",
      "color": "#00ff00"
    }
  ]
}
```

### Результат поиска исполнителей
```json
{
  "lookupResult": [
    {
      "id": "user_123",
      "name": "John Doe",
      "email": "john@example.com"
    }
  ]
}
```

### Результат поиска пользовательского поля
Результаты варьируются в зависимости от типа пользовательского поля, которое ищется. Например, поиск валютного поля может вернуть:
```json
{
  "lookupResult": [
    {
      "value": 1000,
      "currency": "USD"
    },
    {
      "value": 2500,
      "currency": "EUR"
    }
  ]
}
```

## Необходимые разрешения

| Действие | Необходимое разрешение |
|--------|-------------------|
| Create lookup field | `OWNER` or `ADMIN` role at project level |
| Update lookup field | `OWNER` or `ADMIN` role at project level |
| View lookup results | Standard record view permissions |
| Access source data | View permissions on referenced project required |

**Важно**: Пользователи должны иметь разрешения на просмотр как текущего проекта, так и ссылочного проекта, чтобы видеть результаты поиска.

## Ответы на ошибки

### Неверное ссылочное поле
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

### Обнаружен круговой поиск
```json
{
  "errors": [{
    "message": "Circular lookup detected",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### Отсутствует ID поиска для TODO_CUSTOM_FIELD
```json
{
  "errors": [{
    "message": "lookupId is required when lookupType is TODO_CUSTOM_FIELD",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

## Рекомендации по лучшим практикам

1. **Четкие названия**: Используйте описательные названия, которые указывают, какие данные ищутся
2. **Соответствующие типы**: Выбирайте тип поиска, который соответствует вашим потребностям в данных
3. **Производительность**: Поиски обрабатывают все ссылочные записи, поэтому будьте внимательны к ссылочным полям с большим количеством ссылок
4. **Разрешения**: Убедитесь, что пользователи имеют доступ к ссылочным проектам для корректной работы поисков

## Общие случаи использования

### Видимость между проектами
Отображайте теги, исполнителей или статусы из связанных проектов без ручной синхронизации.

### Отслеживание зависимостей
Показывайте сроки выполнения или статус завершения задач, от которых зависит текущая работа.

### Обзор ресурсов
Отображайте всех членов команды, назначенных на ссылочные задачи, для планирования ресурсов.

### Агрегация статусов
Соберите все уникальные статусы из связанных задач, чтобы увидеть состояние проекта на первый взгляд.

## Ограничения

- Поля поиска являются только для чтения и не могут быть отредактированы напрямую
- Нет функций агрегации (SUM, COUNT, AVG) - поиски только извлекают данные
- Нет опций фильтрации - все ссылочные записи включены
- Цепочки круговых поисков предотвращаются, чтобы избежать бесконечных циклов
- Результаты отражают текущие данные и обновляются автоматически

## Связанные ресурсы

- [Ссылочные поля](/api/custom-fields/reference) - создайте ссылки на записи для источников поиска
- [Значения пользовательских полей](/api/custom-fields/custom-field-values) - установите значения для редактируемых пользовательских полей
- [Список пользовательских полей](/api/custom-fields/list-custom-fields) - запросите все пользовательские поля в проекте