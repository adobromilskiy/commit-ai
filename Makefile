.DEFAULT_GOAL := build

export GOFLAGS = -mod=vendor

.PHONY: fmt
fmt:
	@gofumpt -l -w .

.PHONY: vet
vet: fmt
	@go vet ./...
	@staticcheck ./...
	@shadow ./...

.PHONY: lint
lint: vet
	@deadcode ./...
	@golangci-lint run ./...

.PHONY: build
build: vet
	@go install && echo "installation successful"
