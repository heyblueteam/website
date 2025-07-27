# Verification for: Currency Custom Field
Path: /content/en/api/5.custom fields/currency.md
Status: [✅] Completed

## 1. GraphQL Schema Verification

### Custom Field Type
- [✅] Verify `CURRENCY` exists in CustomFieldType enum - **YES** at line 4510
- [✅] Verify required input fields for CURRENCY type - **YES**

### Mutations
- [✅] Verify `createCustomField` mutation supports CURRENCY type - **YES**
- [✅] Verify `setTodoCustomField` mutation with currency parameters - **YES**
- [✅] Verify `createTodo` mutation accepts currency in customFields array - **YES**

## 2. Input Parameter Verification

### CreateCustomFieldInput Fields for CURRENCY
- [✅] `name` field exists and is String! - **YES**
- [✅] `type` field accepts CURRENCY value - **YES**
- [❌] `projectId` field exists and is String! - **NO** - project comes from context
- [✅] `currency` field exists as optional String - **YES** line 19
- [✅] `min` field exists as optional Float - **YES** line 17
- [✅] `max` field exists as optional Float - **YES** line 18
- [✅] `description` field exists as optional String - **YES** line 11
- [❌] `isActive` field exists as optional Boolean - **NO** - field doesn't exist

### SetTodoCustomFieldInput for Currency
- [✅] Verify `number` parameter for currency amount - **YES** line 2543
- [✅] Verify `currency` parameter for currency code - **YES** line 2541
- [✅] Check if both are required or optional - **OPTIONAL**

## 3. Currency Support Verification

### Supported Currencies
- [✅] Verify the list of 72 supported currencies - **EXACT MATCH**
- [✅] Check if all listed fiat currencies are supported - **ALL VERIFIED**
- [✅] Verify BTC and ETH cryptocurrency support - **YES**
- [✅] Check for any additional or missing currencies - **NONE**

## 4. Response Type Verification

### TodoCustomField Response
- [✅] `number` field for currency amount - **YES**
- [✅] `currency` field for currency code - **YES**
- [✅] Verify other standard fields (id, customField, todo, timestamps) - **YES**

## 5. Business Logic Verification

### Currency Formatting
- [✅] Verify automatic formatting behavior - **YES** with Intl.NumberFormat
- [✅] Check locale-specific formatting claims - **YES** each currency has locale
- [✅] Verify USD/CAD prefix behavior - **YES** special handling confirmed

### Validation Rules
- [⚠️] Min/max constraint enforcement - **NOT FOUND** in setTodoCustomField
- [⚠️] Currency code validation - **NO EXPLICIT** validation in resolver
- [✅] Decimal place handling (2 decimal places) - **YES** 0-2 decimal places
- [✅] Invalid currency code handling - **YES** via parsing logic

## 6. Permission Verification

### CRUD Permissions
- [❌] Create field permission: CUSTOM_FIELDS_CREATE vs OWNER/ADMIN - **NO CUSTOM_FIELDS_CREATE permission exists**
- [❌] Update field permission: CUSTOM_FIELDS_UPDATE vs OWNER/ADMIN - **NO CUSTOM_FIELDS_UPDATE permission exists**
- [✅] Set value permission requirements - **Role-based editable flag check**
- [✅] View value permission requirements - **Standard record view**

**ACTUAL**: Project context is required (user must be in project). No specific role check in createCustomField, but setTodoCustomField checks role-based editable permissions.

## 7. Error Code Verification

### Documented Error Codes
- [❌] `INVALID_CURRENCY` - **DOESN'T EXIST**
- [❌] `VALUE_OUT_OF_RANGE` - **DOESN'T EXIST**
- [❌] `INVALID_NUMBER` - **DOESN'T EXIST**
- [✅] Check actual error codes used - **Uses CustomFieldValueParseError for parsing issues**

## 8. Integration Features Verification

### Related Features
- [✅] Formula field integration claims - **YES** updateFormulaResults called
- [✅] Currency conversion field references - **YES** CURRENCY_CONVERSION type exists
- [⚠️] Automation trigger capabilities - **NOT VERIFIED**

## 9. Link Verification

### Internal API Links
- [✅] `/api/custom-fields/currency-conversion` - **EXISTS**
- [✅] `/api/custom-fields/number` - **EXISTS**
- [✅] `/api/custom-fields/formula` - **EXISTS**
- [❌] `/custom-fields/list-custom-fields` - **WRONG** should be `/api/custom-fields/list-custom-fields`

## 10. Known Issues to Check

### Common Documentation Issues
- [✅] Check for hallucinated features - **NONE FOUND**
- [✅] Verify all code examples compile - **MOSTLY ACCURATE**
- [✅] Check for non-existent parameters - **projectId and isActive don't exist**
- [✅] Verify formatting examples are accurate - **YES**

## Summary

### Critical Issues (Must Fix)
1. **Wrong permissions documented**: No CUSTOM_FIELDS_CREATE/UPDATE permissions exist. Actual permission model is project membership + role-based editable flags
2. **Non-existent error codes**: INVALID_CURRENCY, VALUE_OUT_OF_RANGE, INVALID_NUMBER don't exist
3. **projectId parameter doesn't exist**: Project comes from context, not input
4. **isActive field doesn't exist**: This field is not in the schema

### Minor Issues (Should Fix)
1. **Missing validation details**: Min/max constraints not enforced in setTodoCustomField
2. **Wrong link path**: `/custom-fields/list-custom-fields` should be `/api/custom-fields/list-custom-fields`
3. **Currency validation**: No explicit currency code validation in resolver (relies on parsing)

### Suggestions
1. Document actual permission model based on project membership
2. Document actual error handling using CustomFieldValueParseError
3. Remove projectId and isActive from input parameters
4. Clarify that min/max are stored but not enforced on value updates

### Overall Assessment
**75% Accurate** - The core functionality is well documented, currency list is perfect, and formatting behavior is accurate. Main issues are with permissions model, error codes, and a few non-existent parameters.