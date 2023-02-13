.PHONY: all
all: prepare

.PHONY: prepare
prepare:
	@go mod download
	@go mod tidy

.PHONY: build
build:  ## Build executable binary file
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/linux-amd64/music-get main.go
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o bin/linux-arm64/music-get main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o bin/darwin-amd64/music-get main.go
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/darwin-arm64/music-get main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o bin/windows-amd64/music-get.exe main.go
	CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o bin/windows-arm64/music-get.exe main.go

.PHONY: help
help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)