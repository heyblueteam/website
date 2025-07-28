---
title: Справочный пользовательский поле
description: Создайте справочные поля, которые ссылаются на записи в других проектах для межпроектных отношений
---

Справочные пользовательские поля позволяют вам создавать ссылки между записями в разных проектах, что обеспечивает межпроектные отношения и обмен данными. Они предоставляют мощный способ соединения связанных работ в структуре проектов вашей организации.

## Простой пример

Создайте простое справочное поле:

```graphql
mutation CreateReferenceField {
  createCustomField(input: {
    name: "Related Project"
    type: REFERENCE
    referenceProjectId: "proj_456"
    description: "Link to related project records"
  }) {
    id
    name
    type
    referenceProjectId
  }
}
```

## Расширенный пример

Создайте справочное поле с фильтрацией и множественным выбором:

```graphql
mutation CreateFilteredReferenceField {
  createCustomField(input: {
    name: "Dependencies"
    type: REFERENCE
    referenceProjectId: "proj_456"
    referenceMultiple: true
    referenceFilter: {
      status: ACTIVE
      tags: ["dependency"]
    }
    description: "Select multiple dependency records from the project"
  }) {
    id
    name
    type
    referenceProjectId
    referenceMultiple
    referenceFilter
  }
}
```

## Входные параметры

### CreateCustomFieldInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `name` | String! | ✅ Да | Отображаемое имя справочного поля |
| `type` | CustomFieldType! | ✅ Да | Должен быть `REFERENCE` |
| `referenceProjectId` | String | Нет | ID проекта для ссылки |
| `referenceMultiple` | Boolean | Нет | Разрешить множественный выбор записей (по умолчанию: ложь) |
| `referenceFilter` | TodoFilterInput | Нет | Критерии фильтрации для ссылочных записей |
| `description` | String | Нет | Текст помощи, отображаемый пользователям |

**Примечание**: Пользовательские поля автоматически ассоциируются с проектом на основе текущего контекста проекта пользователя.

## Конфигурация ссылки

### Одиночные и множественные ссылки

**Одиночная ссылка (по умолчанию):**
```graphql
{
  referenceMultiple: false  # or omit this field
}
```
- Пользователи могут выбрать одну запись из ссылочного проекта
- Возвращает один объект Todo

**Множественные ссылки:**
```graphql
{
  referenceMultiple: true
}
```
- Пользователи могут выбрать несколько записей из ссылочного проекта
- Возвращает массив объектов Todo

### Фильтрация ссылок

Используйте `referenceFilter`, чтобы ограничить, какие записи могут быть выбраны:

```graphql
{
  referenceFilter: {
    assigneeIds: ["user_123"]
    tagIds: ["tag_123"]
    dueStart: "2024-01-01"
    dueEnd: "2024-12-31"
    showCompleted: false
  }
}
```

## Установка значений ссылок

### Одиночная ссылка

```graphql
mutation SetSingleReference {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: ["referenced_todo_789"]
  })
}
```

### Множественные ссылки

```graphql
mutation SetMultipleReferences {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    customFieldReferenceTodoIds: [
      "referenced_todo_789",
      "referenced_todo_012",
      "referenced_todo_345"
    ]
  })
}
```

### Параметры SetTodoCustomFieldInput

| Параметр | Тип | Обязательный | Описание |
|-----------|------|----------|-------------|
| `todoId` | String! | ✅ Да | ID записи для обновления |
| `customFieldId` | String! | ✅ Да | ID справочного пользовательского поля |
| `customFieldReferenceTodoIds` | [String!] | ✅ Да | Массив ID ссылочных записей |

## Создание записей со ссылками

```graphql
mutation CreateRecordWithReference {
  createTodo(input: {
    title: "Implementation Task"
    todoListId: "list_123"
    customFields: [{
      customFieldId: "reference_field_id"
      value: "referenced_todo_789"
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
      selectedTodos {
        id
        title
        status
      }
    }
  }
}
```

## Поля ответа

### Ответ TodoCustomField

| Поле | Тип | Описание |
|-------|------|-------------|
| `id` | ID! | Уникальный идентификатор для значения поля |
| `customField` | CustomField! | Определение справочного поля |
| `todo` | Todo! | Запись, к которой принадлежит это значение |
| `createdAt` | DateTime! | Когда значение было создано |
| `updatedAt` | DateTime! | Когда значение было в последний раз изменено |

**Примечание**: Ссылочные задачи доступны через `customField.selectedTodos`, а не напрямую в TodoCustomField.

### Ссылочные поля Todo

Каждая ссылочная задача включает:

| Поле | Тип | Описание |
|-------|------|-------------|
| `id` | ID! | Уникальный идентификатор ссылочной записи |
| `title` | String! | Заголовок ссылочной записи |
| `status` | TodoStatus! | Текущий статус (АКТИВЕН, ЗАВЕРШЕН и т.д.) |
| `description` | String | Описание ссылочной записи |
| `dueDate` | DateTime | Срок выполнения, если установлен |
| `assignees` | [User!] | Назначенные пользователи |
| `tags` | [Tag!] | Связанные теги |
| `project` | Project! | Проект, содержащий ссылочную запись |

## Запрос данных ссылок

### Простой запрос

```graphql
query GetRecordsWithReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        selectedTodos {
          id
          title
          status
          project {
            id
            name
          }
        }
      }
    }
  }
}
```

### Расширенный запрос с вложенными данными

```graphql
query GetDetailedReferences {
  todos(projectId: "project_123") {
    id
    title
    customFields {
      id
      customField {
        name
        type
        referenceProjectId
        referenceMultiple
      }
      selectedTodos {
        id
        title
        description
        status
        dueDate
        assignees {
          id
          name
          email
        }
        tags {
          id
          name
          color
        }
        project {
          id
          name
        }
      }
    }
  }
}
```

## Требуемые разрешения

| Действие | Требуемое разрешение |
|--------|-------------------|
| Create reference field | `OWNER` or `ADMIN` role at project level |
| Update reference field | `OWNER` or `ADMIN` role at project level |
| Set reference value | Standard record edit permissions |
| View reference value | Standard record view permissions |
| Access referenced records | View permissions on referenced project |

**Важно**: Пользователи должны иметь разрешения на просмотр в ссылочном проекте, чтобы видеть связанные записи.

## Доступ между проектами

### Видимость проекта

- Пользователи могут ссылаться только на записи из проектов, к которым у них есть доступ
- Ссылочные записи подчиняются разрешениям оригинального проекта
- Изменения в ссылочных записях отображаются в реальном времени
- Удаление ссылочных записей удаляет их из справочных полей

### Наследование разрешений

- Справочные поля наследуют разрешения от обоих проектов
- Пользователи должны иметь доступ к просмотру ссылочного проекта
- Разрешения на редактирование основываются на правилах текущего проекта
- Ссылочные данные доступны только для чтения в контексте справочного поля

## Ответы на ошибки

### Неверный проект ссылки

```json
{
  "errors": [{
    "message": "Project not found",
    "extensions": {
      "code": "PROJECT_NOT_FOUND"
    }
  }]
}
```

### Ссылочная запись не найдена

```json
{
  "errors": [{
    "message": "Custom field not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

### Доступ запрещен

```json
{
  "errors": [{
    "message": "Forbidden",
    "extensions": {
      "code": "FORBIDDEN"
    }
  }]
}
```

## Лучшие практики

### Дизайн полей

1. **Четкие названия** - Используйте описательные названия, которые указывают на взаимосвязь
2. **Соответствующая фильтрация** - Установите фильтры, чтобы показывать только релевантные записи
3. **Учитывайте разрешения** - Убедитесь, что пользователи имеют доступ к ссылочным проектам
4. **Документируйте отношения** - Предоставьте четкие описания связи

### Соображения по производительности

1. **Ограничьте область ссылки** - Используйте фильтры, чтобы уменьшить количество выбираемых записей
2. **Избегайте глубокого вложения** - Не создавайте сложные цепочки ссылок
3. **Учитывайте кэширование** - Ссылочные данные кэшируются для повышения производительности
4. **Мониторинг использования** - Отслеживайте, как ссылки используются в проектах

### Целостность данных

1. **Обрабатывайте удаления** - Планируйте, что делать, когда ссылочные записи удаляются
2. **Проверяйте разрешения** - Убедитесь, что пользователи могут получить доступ к ссылочным проектам
3. **Обновляйте зависимости** - Учитывайте влияние при изменении ссылочных записей
4. **Аудиторские следы** - Отслеживайте ссылки для соблюдения требований

## Общие случаи использования

### Зависимости проекта

```graphql
# Link to prerequisite tasks in other projects
{
  name: "Prerequisites"
  type: REFERENCE
  referenceProjectId: "infrastructure_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: true
    tagIds: ["prerequisite_tag_id"]
  }
}
```

### Требования клиента

```graphql
# Reference client requirements from a requirements project
{
  name: "Client Requirements"
  type: REFERENCE
  referenceProjectId: "requirements_project"
  referenceFilter: {
    assigneeIds: ["client_user_id"]
    showCompleted: false
  }
}
```

### Распределение ресурсов

```graphql
# Link to resource records in a resource management project
{
  name: "Assigned Resources"
  type: REFERENCE
  referenceProjectId: "resources_project"
  referenceMultiple: true
  referenceFilter: {
    tagIds: ["available_tag_id"]
  }
}
```

### Обеспечение качества

```graphql
# Reference QA test cases from a testing project
{
  name: "Test Cases"
  type: REFERENCE
  referenceProjectId: "qa_project"
  referenceMultiple: true
  referenceFilter: {
    showCompleted: false
    tagIds: ["test_case_tag_id"]
  }
}
```

## Интеграция с поисковыми полями

Справочные поля работают с [Поисковыми полями](/api/custom-fields/lookup) для извлечения данных из ссылочных записей. Поисковые поля могут извлекать значения из записей, выбранных в справочных полях, но они являются только извлекателями данных (агрегирующие функции, такие как SUM, не поддерживаются).

```graphql
# Reference field links to records
{
  name: "Related Tasks"
  type: REFERENCE
  referenceProjectId: "other_project"
}

# Lookup field extracts data from referenced records
{
  name: "Task Status"
  type: LOOKUP
  lookupOption: {
    customFieldId: "related_tasks_field_id"
    targetField: "status"
  }
}
```

## Ограничения

- Ссылочные проекты должны быть доступны пользователю
- Изменения в разрешениях ссылочного проекта влияют на доступ к справочным полям
- Глубокое вложение ссылок может повлиять на производительность
- Нет встроенной проверки на циклические ссылки
- Нет автоматического ограничения, предотвращающего ссылки на один и тот же проект
- Проверка фильтров не применяется при установке значений ссылок

## Связанные ресурсы

- [Поисковые поля](/api/custom-fields/lookup) - Извлечение данных из ссылочных записей
- [API проектов](/api/projects) - Управление проектами, содержащими ссылки
- [API записей](/api/records) - Работа с записями, имеющими ссылки
- [Обзор пользовательских полей](/api/custom-fields/list-custom-fields) - Общие концепции