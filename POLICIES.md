# Creating Policy JSON Files and Understanding Operators in go-policy-enforcer

## Table of Contents

- [Policy JSON File Structure](#policy-json-file-structure)
- [Policy Operators](#policy-operators)

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
- - `field`: The name of the field to compare.
- - `operator`: The operator to use for comparison.
- - `value`: The value to compare against.

## Policy Operators

The library supports the following policy operators:

- `==`: Equal to.
- `!=`: Not equal to.
- `>`: Greater than.
- `<`: Less than.
- `>=`: Greater than or equal to.
- `<=`: Less than or equal to.
- `in`: Check if a value is present in a slice.
- `not in`: Check if a value is not present in a slice.

When evaluating a policy, the library compares the value of the specified
field with the provided value using the specified operator. If the comparison
is true, the policy is enforced; otherwise, it is not enforced.

For example, consider the following policy rule:

```json
{
  "field": "age",
  "operator": ">",
  "value": 18
}
```

In this case, the policy will be enforced if the `age` field of the asset
is greater than 18.

To enforce a policy on an asset, you can use the `Enforce` method of
the `PolicyEnforcer` struct. The method takes an asset as input and
returns a boolean value indicating whether the policy is enforced or not.

Here's an example of how to enforce a policy on an asset:

```go
func main() {
   // Load a policy from a JSON file
   policy, err := gopolicyenforcer.LoadPolicy("path/to/policy.json")
   if err != nil {
      log.Fatalf("Error loading policy: %v", err)
   }

   // Create an asset to test enforcement
   asset := &Asset{ID: 1, Name: "John Doe", Age: 25}

   // Create a PolicyEnforcer instance with the policy
   e := gopolicyenforcer.NewPolicyEnforcer(policy)

   // Enforce the policy on the asset and print the result
   if e.Enforce(asset) {
      fmt.Println("Policy is enforced")
   } else {
      fmt.Println("Policy is not enforced")
   }
}
```

In this example, the policy will be enforced because the `age` field
of the asset is greater than 18.
