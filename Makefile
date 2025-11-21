################################################################################
### Infra-Lab cli ##############################################################
################################################################################
MAKEFLAGS = --no-print-directory # __hidethis__
####

.PHONY: help
help: ## 				Show this help
	@sed -e '/__hidethis__/d; /^\.PHONY.*/d; /[A-Z0-9#]?*/!d; /^\t/d; s/:.##/\t/g; s/^####.*//; s/#/-/g; s/^\([A-Z0-9_]*=.*\)/| \1/g; s/^\([a-zA-Z0-9]\)/* \1/g; s/^| \(.*\)/\1/' $(MAKEFILE_LIST)

################################################################################
### Go #########################################################################
################################################################################

OUTPUT_DIR="outputs"
####

.PHONY: build
build: ##				Build binary of main package
	@go build -o ${OUTPUT_DIR}/infra-lab-cli main.go

.PHONY: test
test: ##				Run tests
	@go test -v ./...

.PHONY: test-json
test-json: ##			Run tests and print output in JSON
	@go test -json ./...

.PHONY: tidy
tidy: ##				Update go.mod and go.sum files
	@go mod tidy

.PHONY: lint
lint: ##				Lint the code
	@golangci-lint run -c .golangci.yaml --timeout 5m
