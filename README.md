# 🚀 Go Policy Enforcer

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)

![Build](https://github.com/kmesiab/go-policy-enforcer/actions/workflows/go-build.yml/badge.svg)
![Build](https://github.com/kmesiab/go-policy-enforcer/actions/workflows/go-lint.yml/badge.svg)
![Build](https://github.com/kmesiab/go-policy-enforcer/actions/workflows/go-test.yml/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/kmesiab/go-policy-enforcer)](https://goreportcard.com/report/github.com/kmesiab/go-policy-enforcer)

Go Policy Enforcer is a flexible, lightweight library that allows
developers to dynamically enforce policies on Go structs using
customizable rules and reflection. This library is useful for
scenarios such as access control, validation, and policy-based
filtering of data.

## ✨ Features

- ⚡ **Dynamic Policy Enforcement**: Apply rules dynamically without
hardcoding logic.
- 🔍 **Reflection-Based Field Access**: Leverages Go's reflection
to access and evaluate struct fields.
- 🛠️ **Custom Operators**: Supports custom operators like
`==`, `!=`, `<`, `>`, and more.
- 📦 **Extensible**: Easily add new rules and policies to
adapt to your use case.

## 📚 Installation

Install the package via `go get`:

```bash
go get github.com/kmesiab/go-policy-enforcer
```

## 🛠️ Usage

Here's a quick example of how to use Go Policy Enforcer to enforce rules
on your Go structs:

### Define Your Struct

```go
type Asset struct {
    ID        string
    Type      string
    Finalized bool
}
```

### Define Your Policies

Policies can be written to define what rules need to be enforced on
a given struct.

```go
policies := []Policy{
    {
        Name: "FinalizedPolicy",
        Rules: []Rule{
        {Field: "Finalized", Operator: "==", Value: true},
    },
    {
        Name: "TypePolicy",
        Rules: []Rule{
        {Field: "Type", Operator: "==", Value: "asset"},
    },
}
```

### Enforce Policies

Create a policy enforcer and apply the rules to your struct:

```go
resource := Asset{
    ID:        "1",
    Type:      "asset",
    Finalized: true,
}

enforcer := NewPolicyEnforcer(&policies)

if enforcer.Enforce(resource) {
    fmt.Println("Asset passes all policies")
} else {
    fmt.Println("Asset failed one or more policies")
}
```

### Learn More About Policy Files

For more information about creating policy JSON files and understanding
operators in go-policy-enforcer, refer to the [POLICIES.md](POLICIES.md) file.

### Learn More About Policy Operators

For more information about the operators used in the go-policy-enforcer library,
refer to the [OPERATORS.md](OPERATORS.md) file.

---

## PolicyEnforcer `Match` Function

The `Match` function in the `PolicyEnforcer` class is designed to evaluate
resources against a set of policies. It checks if the provided resource
satisfies the conditions defined in any of the policies and returns a
list of the matching policies.

### How it Works

The `Match` function takes a resource as input and evaluates each policy's
rules against that resource. Each policy contains a set of rules, and the
resource is checked to see if it satisfies all the rules in a policy. If
a resource satisfies all the rules in a policy, that policy is considered
a match, and it is added to the list of matched policies.

The `Match` function does not stop at the first match. It continues evaluating
the resource against all available policies and returns all the matching
policies.

#### Usage Example

Here is an example of how to use the `Match` function:

```go
// Define some policies
policies := []Policy{
    {
        Name: "FinalizedPolicy",
        Rules: []Rule{
        {Field: "Finalized", Operator: "==", Value: true},
    },
    {
        Name: "TypePolicy",
        Rules: []Rule{
        {Field: "Type", Operator: "==", Value: "asset"},
    },
}

// Define a resource that will be evaluated against the policies
resource := struct {
    Age    int
    Status string
}{
    Age:    30,
    Status: "active",
}

// Create a PolicyEnforcer instance with the policies
enforcer := NewPolicyEnforcer(&policies)

// Use the Match function to find matching policies
matchedPolicies := enforcer.Match(resource)

// Output the names of the matched policies
for _, policy := range matchedPolicies {
    fmt.Println("Matched policy:", policy.Name)
}
```

## Handling Nested and Complex Structs

```go
type Address struct {
    City    string
    Country string
}

type Person struct {
    Name    string
    Address Address
}

policies := []Policy{
    {
        Name: "CityPolicy",
        Rules: []Rule{
            {Field: "Address.City", Operator: "==", Value: "Seattle"},
        },
    },
}
```

The `go-policy-enforcer` library is designed to handle nested and complex
struct types seamlessly through the use of reflection and dot notation for
field access.

This means you can enforce policies not only on simple flat
structs but also on deeply nested structs. By specifying fields using dot
notation, such as `address.city`, you can perform evaluations on nested
fields within your data structures.

This powerful feature allows the library to accommodate complex data
hierarchies and intricate JSON structures, ensuring that your policies
can be as detailed and comprehensive as needed.

---

## ✅ Running Tests

Run the following command to execute tests:

```bash
go test ./...
```

---

## 🤝 Contributors

We appreciate the contributions from the community that help to make
`go-policy-enforcer` better. If you would like to contribute, please
feel free to fork the repository and submit a pull request.

Visit the [CONTRIBUTING.md](CONTRIBUTING.md) guid.

### How to Contribute

1. **Fork the Repository**: Click the "Fork" button in the top-right corner
of the GitHub page.
2. **Clone Your Fork**: Clone your forked repository to your local machine.

   ```bash
   git clone https://github.com/kmesiab/go-policy-enforcer.git

---

## 📝 License

This project is licensed under the MIT License. See the
[LICENSE](./LICENSE) file for more details.
