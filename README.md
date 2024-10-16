# 🚀 Go Policy Enforcer

![Golang](https://img.shields.io/badge/Go-00add8.svg?labelColor=171e21&style=for-the-badge&logo=go)

![Build](https://github.com/kmesiab/go-policy-enforcer/actions/workflows/go
-build.yml/badge.svg)
![Build](https://github.com/kmesiab/go-policy-enforcer/actions/workflows/go-lint.yml/badge.svg)
![Build](https://github.com/kmesiab/go-policy-enforcer/actions/workflows/go-test.yml/badge.svg)
[![Go Report Card](htt

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

## ✅ Running Tests

Run the following command to execute tests:

```bash
go test ./...
```

## 📝 License

This project is licensed under the MIT License. See the
[LICENSE](./LICENSE) file for more details.
