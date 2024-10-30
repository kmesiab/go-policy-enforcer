# Contributing to Go Policy Enforcer

Thank you for considering contributing to Go Policy Enforcer! 🎉 This
guide helps you navigate the contribution process and ensures code consistency
across the project.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [How Can I Contribute?](#how-can-i-contribute)
- [Reporting Bugs](#reporting-bugs)
- [Suggesting Enhancements](#suggesting-enhancements)
- [Contributing Code](#contributing-code)
- [Development Setup](#development-setup)
- [Style Guidelines](#style-guidelines)
- [Coding Standards](#coding-standards)
  - [Commit Messages](#commit-messages)
- [Testing](#testing)
- [Pull Request Process](#pull-request-process)
- [Acknowledgments](#acknowledgments)

## Code of Conduct

This project follows a [Contributor Code of Conduct](CODE_OF_CONDUCT.md).
By participating, you agree to abide by its terms.

## How Can I Contribute?

### Reporting Bugs

When reporting a bug, please include:

- **Description**: A concise summary of the bug.
- **Steps to Reproduce**: Describe how to trigger the bug.
- **Expected Behavior**: What should happen?
- **Screenshots or Logs**: Attach relevant details.
- **Environment**: OS, Go version, and any other relevant information.

Please file bug reports in the
[GitHub Issues](https://github.com/kmesiab/go-policy-enforcer/issues).

### Suggesting Enhancements

Enhancement requests are welcome! When suggesting improvements:

- **Title**: A clear and descriptive title.
- **Description**: A detailed explanation.
- **Benefits**: Explain the benefit to users or developers.

### Contributing Code

We appreciate code contributions! Follow these guidelines to help your
code align with the library's standards.

## Development Setup

1. **Fork the repository** by clicking the "Fork" button on the GitHub page.
2. **Clone your forked repository**:

   ```sh
   git clone https://github.com/kmesiab/go-policy-enforcer.git
   cd go-policy-enforcer
   ```

3. **Create a new branch**:

   ```sh
   git checkout -b feature/your-feature-name
   ```

4. **Install dependencies**: Ensure you have installed Go 1.16+ and that
   dependencies are managed by Go modules. Verify your Go version with:

```sh
   go version
```

## Style Guidelines

### Coding Standards

We follow idiomatic Go patterns. Key practices include:

- **Formatting**: Run `go fmt` on all code.
- +**Linting**: Use `golangci-lint` for comprehensive static analysis:
  
```sh
# Install golangci-lint
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

# Run linting
golangci-lint run
```

- **Naming**: Use descriptive names for functions and variables; avoid
  abbreviations unless widely understood (e.g., `len`, `fmt`).
- **Error Handling**: Return explicit error values, and log errors with
  meaningful messages.
- **Code Structure**:
  - Functions that operate on common data types (e.g., slices) should be
    designed to be reusable.
  - Group related functions in a single file, and separate logical
    units into distinct files (e.g., operators, helpers).
- **Documentation**: Document all exported functions and types with Go-style
  comments.

### Commit Messages

We use **Conventional Commits** to keep history clean and searchable:

- **Format**: `<type>(scope): <description>`
- **Types**: Use types like `feat`, `fix`, `refactor`, `docs`, `test`, and
  `chore`.
- **Examples**:
  - `feat(policy): add deep equality operator`
  - `fix(operators): handle nil cases in deep equality`
- **Link Issues**: Mention related issues (e.g., `Fixes #15`).

## Testing

Testing is crucial for this library, especially given its logical complexity.
Adhere to the following:

- **Write Tests**: All new code must be covered by tests, especially custom
- operators and utility functions.
- **Use Table-Driven Tests**: For testing functions with multiple cases
- (e.g., operators), use table-driven tests for clarity.
- **Run Tests**: Verify functionality with `go test ./...`:

  ```sh
  go test ./...
  ```

- **Test Coverage**: Aim for high coverage, especially for critical logic
- in policy evaluation and operator functions.

## Pull Request Process

1. **Run Tests**:

   ```sh
   go test ./...
   ```

2. **Commit Changes**:

   ```sh
   git add .
   git commit -m "feat(policy): add deep equality operator"
   ```

3. **Push to Your Fork**:

   ```sh
   git push origin feature/your-feature-name
   ```

4. **Submit a Pull Request**:

    - Go to the original repository.
    - Click "New Pull Request".
    - Select your branch and submit.

5. **Review and Update**: Be responsive to feedback and make adjustments as
   necessary.

## Acknowledgments

Thank you for considering contributing to the Go Policy Enforcer project.
We value your time and effort.
