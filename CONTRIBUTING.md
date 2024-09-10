# Contributing to template-module-placeholder

## Prerequisites

- [Go](https://golang.org/doc/install)
- [golangci-lint](https://github.com/golangci/golangci-lint)
- [go-test-coverage](https://github.com/vladopajic/go-test-coverage)

These can be installed with the following command:

```bash
make setup
```

This uses the tools defined in the `tools/tools.go` file. If you need to add more tools, add them to this file.

## Testing

Testing should use the standard Go testing framework and not rely on any external dependencies including assertion libraries.

To run tests, use the following command:

```bash
make test
```

This is just a wrapper around the following command:

```bash
go test -v --race ./...
```

## Linting

Linting should use the standard golangci-lint tool. To run linting, use the following command:

```bash
make lint
```

This is just a wrapper around the following command:

```bash
golangci-lint --config .golangci.yml run ./...
```

The configuration file is located in the root of the repository under `.golangci.yml`.

## Code Coverage

Code coverage should be used to measure the percentage of code that is covered by tests. To run code coverage, use the following command:

```bash
make test-coverage
```

This is just a wrapper around the following command:

```bash
go test -v --race -coverprofile=coverage.out -covermode=atomic ./...
```

The coverage report will be generated in the `coverage.out` file. We then use the `go-test-coverage` tool to generate a summary report in the terminal. We also use the standard Go tool to generate a HTML report. The HTML report is generated in the `coverage.html` file.

