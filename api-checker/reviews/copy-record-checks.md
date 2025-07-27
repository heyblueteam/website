# Verification for: Copy Record
Path: /content/en/api/3.records/8.copy-record.md
Status: [✅] Completed

## 1. GraphQL Schema Verification

### Mutation Operation
- [✅] Verify `copyTodo` mutation exists
  - Operation: `copyTodo`
  - Location in schema: `/bloo-api/src/schema.graphql`
  - Actual vs Documented: **MOSTLY MATCHES** - `copyTodo(input: CopyTodoInput!): Boolean!`

## 2. Input Parameter Verification

### CopyTodo Input Fields
- [❌] `title` field
  - Documented type: `String!` (Required: Yes)
  - Actual type: **String** (OPTIONAL)
  - Required status matches: **NO** - Documentation wrong

- [✅] `todoId` field
  - Documented type: `String!`
  - Actual type: **String!**
  - Required status matches: **YES**

- [✅] `todoListId` field
  - Documented type: `String!`
  - Actual type: **String!**
  - Required status matches: **YES**

- [✅] `options` field
  - Documented type: `Array!`
  - Actual type: **[CopyTodoOption!]!**
  - Required status matches: **YES**

## 3. Copy Options Enum Verification

### CopyTodoOptions Enum
- [✅] `DESCRIPTION` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `DUE_DATE` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `CHECKLISTS` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `ASSIGNEES` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `TAGS` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [✅] `CUSTOM_FIELDS` option
  - Exists in schema: **YES**
  - Description matches: **YES**

- [⚠️] **Missing from docs**: `COMMENTS` option exists in schema but not documented

## 4. Response Field Verification

### CopyTodo Response
- [❌] Response structure mismatch
  - Documented: `{ success: Boolean }`
  - Actual type: **Boolean!** (direct boolean, not object)
  - Documentation shows wrong response format

## 5. Implementation Verification

### Resolver Check
- [✅] Resolver exists for `copyTodo`
  - Location: `/bloo-api/src/resolvers/Mutation/copyTodo.ts`
  - Handler function: `copyTodo`

### Business Logic Verification
- [✅] Copy functionality implemented correctly
  - Options handling: **EXCELLENT** - all options work as documented
  - Positioning logic: **CORRECT** - places at bottom of list
  - Implementation: `/bloo-api/src/datasources/TodoDatasource.ts`

## 6. Permission Verification

### Required Permissions
- [✅] Edit permissions on source list
  - Permission check exists: **YES** - OWNER/ADMIN/MEMBER required
  - Location: `/bloo-api/src/permissions/permissions.ts`

- [✅] Edit permissions on target list
  - Permission check exists: **YES** - cross-project restrictions for MEMBER
  - Location: `copyTodo.ts` resolver

- [✅] **Additional restrictions found**:
  - MEMBER role can only copy within same project
  - Assignees filtered by target project membership

## 7. Error Code Verification

### Documented Error Codes
- [❌] `BAD_USER_INPUT`
  - Exists in codebase: **YES** but not used for invalid IDs
  - Used for invalid IDs: **NO** - uses specific errors instead
  - Actual errors: `TODO_NOT_FOUND`, `TODO_LIST_NOT_FOUND`

- [✅] `FORBIDDEN`
  - Exists in codebase: **YES**
  - Used for permission errors: **YES** (UnauthorizedError)
  - Location: `/bloo-api/src/lib/errors.ts:139`

- [❌] `GRAPHQL_VALIDATION_FAILED`
  - Exists in codebase: **NOT FOUND**
  - Used for missing fields: **NO** - GraphQL handles validation
  - This error code doesn't exist in the codebase

## 8. Headers Verification

### Required Headers
- [✅] `x-bloo-token-id` header requirement
  - Actually required: **YES**
  - Part of authentication: **YES**

- [✅] `x-bloo-token-secret` header requirement
  - Actually required: **YES**
  - Part of authentication: **YES**

- [✅] `x-bloo-project-id` header requirement
  - Actually required: **YES**
  - Used for project context: **YES**

- [✅] `x-bloo-company-id` header requirement
  - Actually required: **YES**
  - Used for company context: **YES**

## 9. Link Verification

### Internal API Links
- [✅] `/api/records/move-record-list` - **EXISTS** (5.move-record-list.md)
- [✅] `/api/error-codes` - **EXISTS** (12.error-codes.md)

## 10. Code Example Verification

### Basic Example
- [✅] GraphQL syntax is valid
- [✅] All fields in mutation exist
- [❌] Required fields issue - title marked as required but actually optional
- [❌] Response structure doesn't match - shows object but returns boolean

## 11. Special Considerations

### Business Rules
- [✅] "Copied record placed at bottom of target list" - **VERIFIED** in implementation
- [✅] Copy options actually copy the specified data elements - **ALL WORK CORRECTLY**
- [✅] Cross-list copying is supported - **YES** with permission restrictions

### Additional Features Found
- [✅] **Cross-project copying** supported (with restrictions)
- [✅] **Assignee filtering** by target project membership
- [✅] **Custom field file handling** with proper storage copying
- [✅] **Automation triggers** for cross-project copies

## Summary

### Critical Issues (Must Fix)
1. **Title field requirement**: Documentation says required, but schema shows optional
2. **Response format**: Documentation shows `{ success: Boolean }` but actual return is `Boolean!`
3. **Missing copy option**: `COMMENTS` option exists in schema but not documented
4. **Wrong error codes**: Documentation claims `BAD_USER_INPUT` for invalid IDs, but code uses specific errors like `TODO_NOT_FOUND`

### Minor Issues (Should Fix)
1. **Non-existent error code**: `GRAPHQL_VALIDATION_FAILED` doesn't exist in codebase
2. **Incomplete permission description**: Missing details about cross-project restrictions for MEMBER role

### Suggestions
1. **Add COMMENTS option**: Document the missing copy option
2. **Enhance permission documentation**: Add details about cross-project copying restrictions
3. **Update error handling**: Use correct error codes (`TODO_NOT_FOUND`, `TODO_LIST_NOT_FOUND`)
4. **Fix response format**: Show correct boolean return type

### Overall Assessment
The implementation is **excellent and comprehensive** with advanced features like:
- ✅ Cross-project copying with proper restrictions
- ✅ Assignee filtering by project membership  
- ✅ Custom field file handling with storage copying
- ✅ Automation triggers for cross-project operations
- ✅ All copy options work exactly as intended

However, the documentation has several **accuracy issues** that need fixing to match the actual robust implementation. The core functionality is solid but documentation needs updates.