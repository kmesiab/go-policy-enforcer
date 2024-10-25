# OPERATORS.md

## Introduction

Welcome to the `go-policy-enforcer` library! This document provides a
comprehensive guide to the built-in operators, design principles, and
the steps to extend the library with custom operators. By understanding
and contributing to the operators, you can help improve the policy engine’s
flexibility and applicability.

## Getting Started

Before diving into operators, ensure you have the project set up locally:

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-policy-enforcer.git  
   ```

2. Install dependencies and build the project:

   ```bash  
   go mod tidy  
   go build  
   ```

## Built-in Policy Operators

The `go-policy-enforcer` library includes several predefined operators for
evaluating policy rule conditions. Below is a breakdown of each:

**Equality Operators**:
    - `==`: Checks if two values are equal (`reflect.DeepEqual`).
    - `!=`: Checks if two values are not equal.

**Comparison Operators**:
    - `>`: Verifies if the left value is greater than the right.
    - `<`: Checks if the left value is less than the right.
    - `>=`: Determines if the left value is greater than or equal to the right.
    - `<=`: Ensures the left value is less than or equal to the right.

**Membership Operators**:
    - `in`: Validates if a value is present within a slice.
    - `not in`: Confirms a value is absent from a slice.

These operators offer flexibility in policy enforcement, supporting a wide
range of comparisons across data types.

## Design Philosophy

The operator framework in `go-policy-enforcer` is built around key principles:

- **Extensibility**: Operators are function-based, allowing easy extension.
- **Simplicity**: Using function pointers for operator definitions simplifies
the extension process.
- **Reusability**: Operators use Go’s type assertion and reflection, handling
various data types and making them versatile.

## File Structure

Here is an overview of the core files associated with operators to guide new
contributors:

- `operators.go`: Contains predefined operators.
- `operators_map.go`: Maintains the mapping between operator keys and their
functions.
- `custom_operators/`: Directory where custom operators should be created.
- `operators_test.go`: Houses test cases for validating operators.

## Creating Custom Operators

To extend the operator functionality, follow these steps to create and register
a custom operator.

### Step 1: Define the Operator Function

Create a new file in the `custom_operators` folder and implement your operator
function to match the `PolicyCheckOperator` signature.

```go
package custom_operators

// Example of a custom operator function
var CustomOperator = func(leftVal, rightVal any) bool {
// Your custom logic here
return false // Placeholder logic
}
```

### Step 2: Register the Operator

Add the custom operator to `PolicyCheckOperatorMap` in `operators_map.go`. This
links the operator’s string representation to its function:

```go
"custom": custom_operators.CustomOperator,
```

### Step 3: Use the Custom Operator

Once registered, you can use the custom operator within policy JSON files with
the specified key:

```json
{
  "field": "exampleField",
  "operator": "custom",
  "value": "exampleValue"
}
```

## Best Practices for Custom Operators

To maintain quality and consistency in custom operators:

1. **Follow Established Patterns**: Use the structure of existing operators for
consistency.
2. **Thorough Testing**: Test custom operators under various scenarios in a
   custom est file.
3. **Document Behavior**: Include comments detailing the purpose, input types,
and logic of the operator for clarity and usability.
