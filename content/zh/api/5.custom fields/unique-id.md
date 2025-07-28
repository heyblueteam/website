---
title: 唯一 ID 自定义字段
description: 创建具有顺序编号和自定义格式的自动生成唯一标识符字段
---

唯一 ID 自定义字段自动为您的记录生成顺序的唯一标识符。它们非常适合创建票证号码、订单 ID、发票号码或工作流程中的任何顺序标识符系统。

## 基本示例

创建一个简单的唯一 ID 字段，具有自动顺序：

```graphql
mutation CreateUniqueIdField {
  createCustomField(input: {
    name: "Ticket Number"
    type: UNIQUE_ID
    useSequenceUniqueId: true
  }) {
    id
    name
    type
    useSequenceUniqueId
  }
}
```

## 高级示例

创建一个带有前缀和零填充的格式化唯一 ID 字段：

```graphql
mutation CreateFormattedUniqueIdField {
  createCustomField(input: {
    name: "Order ID"
    type: UNIQUE_ID
    description: "Auto-generated order identifier"
    useSequenceUniqueId: true
    prefix: "ORD-"
    sequenceDigits: 4
    sequenceStartingNumber: 1000
  }) {
    id
    name
    type
    description
    useSequenceUniqueId
    prefix
    sequenceDigits
    sequenceStartingNumber
  }
}
```

## 输入参数

### CreateCustomFieldInput (UNIQUE_ID)

| 参数 | 类型 | 必需 | 描述 |
|-----------|------|----------|-------------|
| `name` | String! | ✅ 是 | 唯一 ID 字段的显示名称 |
| `type` | CustomFieldType! | ✅ 是 | 必须是 `UNIQUE_ID` |
| `description` | String | 否 | 显示给用户的帮助文本 |
| `useSequenceUniqueId` | Boolean | 否 | 启用自动顺序（默认：false） |
| `prefix` | String | 否 | 生成 ID 的文本前缀（例如，“TASK-”） |
| `sequenceDigits` | Int | 否 | 零填充的位数 |
| `sequenceStartingNumber` | Int | 否 | 序列的起始数字 |

## 配置选项

### 自动顺序 (`useSequenceUniqueId`)
- **true**: 在创建记录时自动生成顺序 ID
- **false** 或 **undefined**: 需要手动输入（像文本字段一样工作）

### 前缀 (`prefix`)
- 添加到所有生成 ID 的可选文本前缀
- 示例：“TASK-”，“ORD-”，“BUG-”，“REQ-”
- 没有长度限制，但请保持合理以便于显示

### 序列数字 (`sequenceDigits`)
- 零填充序列号的位数
- 示例：`sequenceDigits: 3` 生成 `001`，`002`，`003`
- 如果未指定，则不应用填充

### 起始数字 (`sequenceStartingNumber`)
- 序列中的第一个数字
- 示例：`sequenceStartingNumber: 1000` 从 1000 开始，1001，1002...
- 如果未指定，则从 1 开始（默认行为）

## 生成的 ID 格式

最终的 ID 格式结合了所有配置选项：

```
{prefix}{paddedSequenceNumber}
```

### 格式示例

| 配置 | 生成的 ID |
|---------------|---------------|
| 无选项 | `1`, `2`, `3` |
| `prefix: "TASK-"` | `TASK-1`, `TASK-2`, `TASK-3` |
| `sequenceDigits: 3` | `001`, `002`, `003` |
| `prefix: "ORD-", sequenceDigits: 4` | `ORD-0001`, `ORD-0002`, `ORD-0003` |
| `prefix: "BUG-", sequenceStartingNumber: 500` | `BUG-500`, `BUG-501`, `BUG-502` |
| All options combined | `TASK-1001`, `TASK-1002`, `TASK-1003` |

## 读取唯一 ID 值

### 查询具有唯一 ID 的记录
```graphql
query GetRecordsWithUniqueIds {
  todos(filter: { projectIds: ["proj_123"] }) {
    id
    title
    customFields {
      id
      customField {
        name
        type
        prefix
        sequenceDigits
      }
      sequenceId    # The generated sequence number
      text         # The text value for UNIQUE_ID fields
    }
  }
}
```

### 响应格式
```json
{
  "data": {
    "todos": [
      {
        "id": "todo_123",
        "title": "Fix login issue",
        "customFields": [
          {
            "id": "field_value_456",
            "customField": {
              "name": "Ticket Number",
              "type": "UNIQUE_ID",
              "prefix": "TASK-",
              "sequenceDigits": 3
            },
            "sequenceId": 42,
            "text": "TASK-042"
          }
        ]
      }
    ]
  }
}
```

## 自动 ID 生成

### 何时生成 ID
- **记录创建**：在创建新记录时自动分配 ID
- **字段添加**：当将 UNIQUE_ID 字段添加到现有记录时，排队一个后台作业（工作者实现待定）
- **后台处理**：新记录的 ID 生成通过数据库触发器同步进行

### 生成过程
1. **触发**：创建新记录或添加 UNIQUE_ID 字段
2. **序列查找**：系统查找下一个可用的序列号
3. **ID 分配**：序列号分配给记录
4. **计数器更新**：为未来的记录递增序列计数器
5. **格式化**：ID 在显示时使用前缀和填充进行格式化

### 唯一性保证
- **数据库约束**：每个字段内序列 ID 的唯一约束
- **原子操作**：序列生成使用数据库锁以防止重复
- **项目范围**：序列在每个项目中是独立的
- **竞争条件保护**：并发请求安全处理

## 手动模式与自动模式

### 自动模式 (`useSequenceUniqueId: true`)
- ID 通过数据库触发器自动生成
- 保证顺序编号
- 原子序列生成防止重复
- 格式化 ID 结合前缀 + 填充序列号

### 手动模式 (`useSequenceUniqueId: false` 或 `undefined`)
- 像常规文本字段一样工作
- 用户可以通过 `setTodoCustomField` 和 `text` 参数输入自定义值
- 没有自动生成
- 除数据库约束外不强制唯一性

## 设置手动值（仅限手动模式）

当 `useSequenceUniqueId` 为 false 时，您可以手动设置值：

```graphql
mutation SetUniqueIdValue {
  setTodoCustomField(input: {
    todoId: "todo_123"
    customFieldId: "field_456"
    text: "CUSTOM-ID-001"
  })
}
```

## 响应字段

### TodoCustomField 响应 (UNIQUE_ID)

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `id` | String! | 字段值的唯一标识符 |
| `customField` | CustomField! | 自定义字段定义 |
| `sequenceId` | Int | 生成的序列号（为 UNIQUE_ID 字段填充） |
| `text` | String | 格式化的文本值（结合前缀 + 填充序列） |
| `todo` | Todo! | 此值所属的记录 |
| `createdAt` | DateTime! | 值创建的时间 |
| `updatedAt` | DateTime! | 值最后更新的时间 |

### CustomField 响应 (UNIQUE_ID)

| 字段 | 类型 | 描述 |
|-------|------|-------------|
| `useSequenceUniqueId` | Boolean | 是否启用自动顺序 |
| `prefix` | String | 生成 ID 的文本前缀 |
| `sequenceDigits` | Int | 零填充的位数 |
| `sequenceStartingNumber` | Int | 序列的起始数字 |

## 所需权限

| 操作 | 所需权限 |
|--------|-------------------|
| Create unique ID field | `OWNER` or `ADMIN` role at project level |
| Update unique ID field | `OWNER` or `ADMIN` role at project level |
| Set manual value | Standard record edit permissions |
| View unique ID value | Standard record view permissions |

## 错误响应

### 字段配置错误
```json
{
  "errors": [{
    "message": "Invalid sequence configuration",
    "extensions": {
      "code": "BAD_USER_INPUT"
    }
  }]
}
```

### 权限错误
```json
{
  "errors": [{
    "message": "CustomField not found",
    "extensions": {
      "code": "CUSTOM_FIELD_NOT_FOUND"
    }
  }]
}
```

## 重要说明

### 自动生成的 ID
- **只读**：自动生成的 ID 不能手动编辑
- **永久**：一旦分配，序列 ID 不会更改
- **按时间顺序**：ID 反映创建顺序
- **范围**：序列在每个项目中是独立的

### 性能考虑
- 新记录的 ID 生成通过数据库触发器同步进行
- 序列生成使用 `FOR UPDATE` 锁进行原子操作
- 存在后台作业系统，但工作者实现待定
- 考虑高容量项目的序列起始数字

### 迁移和更新
- 将自动顺序添加到现有记录会排队后台作业（工作者待定）
- 更改序列设置仅影响未来记录
- 配置更新时现有 ID 保持不变
- 序列计数器从当前最大值继续

## 最佳实践

### 配置设计
- 选择不会与其他系统冲突的描述性前缀
- 根据预期的数量使用适当的数字填充
- 设置合理的起始数字以避免冲突
- 在部署前使用示例数据测试配置

### 前缀指南
- 保持前缀简短且易于记忆（2-5 个字符）
- 使用大写字母以保持一致性
- 包含分隔符（连字符、下划线）以提高可读性
- 避免可能在 URL 或系统中引起问题的特殊字符

### 序列规划
- 估计记录数量以选择适当的数字填充
- 在设置起始数字时考虑未来增长
- 为不同记录类型规划不同的序列范围
- 记录您的 ID 方案以供团队参考

## 常见用例

1. **支持系统**
   - 票证号码： `TICK-001`， `TICK-002`
   - 案例 ID： `CASE-2024-001`
   - 支持请求： `SUP-001`

2. **项目管理**
   - 任务 ID： `TASK-001`， `TASK-002`
   - 冲刺项目： `SPRINT-001`
   - 可交付物编号： `DEL-001`

3. **业务运营**
   - 订单编号： `ORD-2024-001`
   - 发票 ID： `INV-001`
   - 采购订单： `PO-001`

4. **质量管理**
   - 错误报告： `BUG-001`
   - 测试用例 ID： `TEST-001`
   - 审查编号： `REV-001`

## 集成功能

### 与自动化
- 在分配唯一 ID 时触发操作
- 在自动化规则中使用 ID 模式
- 在电子邮件模板和通知中引用 ID

### 与查找
- 从其他记录引用唯一 ID
- 按唯一 ID 查找记录
- 显示相关记录标识符

### 与报告
- 按 ID 模式分组和过滤
- 跟踪 ID 分配趋势
- 监控序列使用和间隙

## 限制

- **仅顺序**：ID 按时间顺序分配
- **无间隙**：删除的记录在序列中留下间隙
- **无重用**：序列号永远不会重用
- **项目范围**：不能跨项目共享序列
- **格式约束**：格式选项有限
- **无批量更新**：无法批量更新现有序列 ID
- **无自定义逻辑**：无法实现自定义 ID 生成规则

## 相关资源

- [文本字段](/api/custom-fields/text-single) - 用于手动文本标识符
- [数字字段](/api/custom-fields/number) - 用于数字序列
- [自定义字段概述](/api/custom-fields/2.list-custom-fields) - 一般概念
- [自动化](/api/automations) - 用于基于 ID 的自动化规则