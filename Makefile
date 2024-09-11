.PHONY: all lint test test-coverage coverage setup setup-name setup-githooks setup-tools


GO:=$(shell which go)
GOBIN:=$(shell go env GOPATH)/bin
GOLANGCILINT:=$(GOBIN)/golangci-lint
GOTESTCOVERAGE:=$(GOBIN)/go-test-coverage
DIST?=./dist
GOOS?=$(shell go env GOOS)
GOARCH?=$(shell go env GOARCH)

NAME?=name-placeholder
VERSION?=v0.0.0

default: help

setup: setup-name setup-githooks setup-tools
	@if [ -f ./tools/setup-complete ]; then \
		echo "Setup already completed"; \
		exit 0; \
	fi; \

setup-name:
	@echo "What is the name of your module?"
	@read -p "Name: " name; \
	if [ -z "$$name" ]; then \
		echo "Name is required"; \
		exit 1; \
	fi; \
	echo "What is the version of your module?"
	@read -p "Version: (default: v0.0.0) " version; \
	if [ -z "$$version" ]; then \
		VERSION=v0.0.0; \
	fi; \
	sed -i "s/template-module-placeholder/$$name/" template-README.md; \
	sed -i "s/template-module-placeholder/$$name/" .github/workflows/main.yml; \
	sed -i "s/template-module-placeholder/$$name/" .golangci.yml; \
	sed -i "s/template-module-placeholder/$$name/" .testcoverage.yml; \
	sed -i "s/template-module-placeholder/$$name/" CONTRIBUTING.md; \
	sed -i "s/template-module-placeholder/$$name/" LICENSE; \
	mv template-README.md README.md;
	git tag $$version

setup-tools:
	@echo "Setting up development tools..."
	@cat tools/tools.go | grep _ | awk -F'"' '{print $$2}' | xargs -tI % go install %@latest

setup-githooks:
	@echo "Setting up git hooks..."
	@mkdir -p .git/hooks
	@mv .githooks/pre-push .git/hooks/pre-push
	@chmod +x .git/hooks/pre-push
	@rm -rf .githooks

quality-control: lint test test-coverage coverage

lint: 
	$(GOLANGCILINT) --config .golangci.yml run
	
test:
	$(GO) test -v -race ./...

test-coverage:
	$(GO) test -v -coverprofile coverage.out --covermode=atomic ./...
	$(GO) tool cover -html=coverage.out -o coverage.html
	
coverage: test-coverage
	@echo "Running test coverage..."
	@echo ""
	@echo "--------------------------------------------------"
	@echo "----------[template-module-placeholder]----------"
	$(GOTESTCOVERAGE) --config .testcoverage.yml
	@echo "--------------------------------------------------"
	@echo ""

help:
	@echo "template-module-placeholder"
	@echo ""
	@echo "Usage:"
	@echo "  make help : Show this help message"
	@echo "  make setup : Setup the development environment"
	@echo "  make lint : Lint the code"
	@echo "  make test : Run the tests"
	@echo "  make test-coverage : Run the tests and generate a coverage report"
	@echo "  make coverage : Run the tests and generate a coverage report"
	@echo ""