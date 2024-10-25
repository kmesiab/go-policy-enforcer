# Creating Policy JSON Files and Understanding Operators in go-policy-enforcer

## Table of Contents

- [Policy JSON File Structure](#policy-json-file-structure)
- [Policy Operators](#policy-operators)
- [Handling Nested Values](#handling-nested-values)

## Policy JSON File Structure

Policy JSON files should follow the following structure:

```json
{
  "name": "Policy Name",
  "rules": [
    {
      "field": "Field Name",
      "operator": "Operator",
      "value": "Value"
    },
    {
      "field": "Field Name",
      "operator": "Operator",
      "value": "Value"
    },
    ...
  ]
}
```

- `name`: The name of the policy.
- `rules`: An array of rules that define the conditions for enforcing the
  policy.
- - `field`: The name of the field to compare. Nested fields can be accessed
using dot notation (e.g., `address.city`).
- - `operator`: The operator to use for comparison.
- - `value`: The value to compare against.

## Policy Operators

️ℹ️See the [POLICIES.md](POLICIES.md) file for a detailed explanation of
the operators used in the go-policy-enforcer library.

The library supports the following policy operators:

- `==`: Equal to.
- `!=`: Not equal to.
- `>`: Greater than.
- `<`: Less than.
- `>=`: Greater than or equal to.
- `<=`: Less than or equal to.
- `in`: Check if a value is present in a slice.
- `not in`: Check if a value is not present in a slice.

## Handling Nested Values

To access nested values in the policy rules, use dot notation in the `field`
property. For example, if you have a JSON structure like this:

```json
{
  "name": "John Doe",
  "address": {
    "city": "New York",
    "state": "NY"
  }
}
```

You can create a policy rule to check if the `city` field is equal to
"New York" like this:

```json
{
  "field": "address.city",
  "operator": "==",
  "value": "New York"
}
```

In this case, the policy will be enforced if the `city` field of the nested
`address` object is equal to "New York".

Remember to update the policy JSON file accordingly when using nested values
in your rules.
