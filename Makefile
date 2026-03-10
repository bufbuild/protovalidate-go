# See https://tech.davis-hansson.com/p/make/
SHELL := bash
.DELETE_ON_ERROR:
.SHELLFLAGS := -eu -o pipefail -c
.DEFAULT_GOAL := all
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules
MAKEFLAGS += --no-print-directory
TMP := .tmp
BIN := $(TMP)/bin
BENCH_TMP := $(TMP)/bench
COPYRIGHT_YEARS := 2023-2025
LICENSE_IGNORE := -e internal/testdata/
# Set to use a different compiler. For example, `GO=go1.18rc1 make test`.
GO ?= go
ARGS ?= --strict_message --strict_error
GOLANGCI_LINT_VERSION ?= v2.9.0
# Set the protovalidate version, can be a commit, branch, or tag
PROTOVALIDATE_VERSION ?= 8b6fa0a9cbff5b0e600f1d2a3024e135b5594db7
PROTOVALIDATE_PROTO_VERSION := $(shell PATH="$(abspath $(BIN)):$(PATH)" ./make/scripts/resolve_bsr_commit.sh buf.build/bufbuild/protovalidate $(PROTOVALIDATE_VERSION))
PROTOVALIDATE_TESTING_PROTO_VERSION := $(shell PATH="$(abspath $(BIN)):$(PATH)" ./make/scripts/resolve_bsr_commit.sh buf.build/bufbuild/protovalidate-testing $(PROTOVALIDATE_VERSION))
PROTOBUF_GO_VERSION := $(shell go list -json -m google.golang.org/protobuf  | jq -r '.Version')
PROTOVALIDATE_GEN_SDK_VERSION := $(shell buf registry sdk version --module=buf.build/bufbuild/protovalidate:$(PROTOVALIDATE_PROTO_VERSION) --plugin=buf.build/protocolbuffers/go:$(PROTOBUF_GO_VERSION)) 

.PHONY: help
help: ## Describe useful make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "%-15s %s\n", $$1, $$2}'

.PHONY: all
all: sync-upstream generate test conformance lint ## Generate and run all tests and lint (default)

.PHONY: clean
clean: ## Delete intermediate build artifacts
	@# -X only removes untracked files, -d recurses into directories, -f actually removes files/dirs
	git clean -Xdf

.PHONY: test
test: ## Run all unit tests
	$(GO) test -race -cover ./...

.PHONY: test-opaque
test-opaque: ## Test proto opaque API support
	$(GO) test --tags=protoopaque ./...

.PHONY: lint
lint: lint-proto lint-go  ## Lint code and protos

.PHONY: lint-go
lint-go: $(BIN)/golangci-lint
	$(BIN)/golangci-lint run --modules-download-mode=readonly --timeout=3m0s ./...
	$(BIN)/golangci-lint fmt --diff

.PHONY: lint-proto
lint-proto: $(BIN)/buf
	$(BIN)/buf lint

.PHONY: lint-fix
lint-fix:
	$(BIN)/golangci-lint run --fix --modules-download-mode=readonly --timeout=3m0s ./...
	$(BIN)/golangci-lint fmt

.PHONY: conformance
conformance: $(BIN)/protovalidate-conformance protovalidate-conformance-go ## Run conformance tests
	$(BIN)/protovalidate-conformance $(ARGS) $(BIN)/protovalidate-conformance-go --expected_failures=conformance/expected_failures.yaml

.PHONY: conformance-hyperpb
conformance-hyperpb: ## Run conformance tests against hyperpb
	HYPERPB=true $(MAKE) conformance

.PHONY: generate
generate: generate-proto generate-license ## Regenerate code and license headers
	$(GO) mod tidy

.PHONY: generate-proto
generate-proto: $(BIN)/buf
	rm -rf internal/gen/*/
	$(BIN)/buf generate buf.build/bufbuild/protovalidate-testing:$(PROTOVALIDATE_TESTING_PROTO_VERSION)
	$(BIN)/buf generate

.PHONY: sync-upstream
sync-upstream:
	yq -i '(.deps[] | select(. == "buf.build/bufbuild/protovalidate:*")) = "buf.build/bufbuild/protovalidate:$(PROTOVALIDATE_PROTO_VERSION)"' buf.yaml
	buf dep update
	go get buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go@$(PROTOVALIDATE_GEN_SDK_VERSION)

.PHONY: generate-license
generate-license: $(BIN)/license-header
	@# We want to operate on a list of modified and new files, excluding
	@# deleted and ignored files. git-ls-files can't do this alone. comm -23 takes
	@# two files and prints the union, dropping lines common to both (-3) and
	@# those only in the second file (-2). We make one git-ls-files call for
	@# the modified, cached, and new (--others) files, and a second for the
	@# deleted files.
	comm -23 \
		<(git ls-files --cached --modified --others --no-empty-directory --exclude-standard | sort -u | grep -v $(LICENSE_IGNORE) ) \
		<(git ls-files --deleted | sort -u) | \
		xargs $(BIN)/license-header \
			--license-type apache \
			--copyright-holder "Buf Technologies, Inc." \
			--year-range "$(COPYRIGHT_YEARS)"

.PHONY: checkgenerate
checkgenerate: generate
	@# Used in CI to verify that `make generate` doesn't produce a diff.
	test -z "$$(git status --porcelain | tee /dev/stderr)"


BENCH ?= .
BENCH_COUNT ?= 10
BENCH_NAME ?= $(shell date +%F:%T)
.PHONY: bench
bench: $(BENCH_TMP)
	go test -run ^$$ -bench="$(BENCH)" -benchmem \
		-memprofile "$(BENCH_TMP)/$(BENCH_NAME).mem.profile" \
		-cpuprofile "$(BENCH_TMP)/$(BENCH_NAME).cpu.profile" \
		-count $(BENCH_COUNT) \
		| tee "$(BENCH_TMP)/$(BENCH_NAME).bench.txt"


.PHONY: upgrade-go
upgrade-go:
	$(GO) get -u -t ./... && $(GO) mod tidy -v

$(BENCH_TMP):
	@mkdir -p $(BENCH_TMP)

$(BIN):
	@mkdir -p $(BIN)

$(BIN)/buf: $(BIN) Makefile
	GOBIN=$(abspath $(@D)) $(GO) install github.com/bufbuild/buf/cmd/buf@latest

$(BIN)/license-header: $(BIN) Makefile
	GOBIN=$(abspath $(@D)) $(GO) install \
		  github.com/bufbuild/buf/private/pkg/licenseheader/cmd/license-header@latest

$(BIN)/golangci-lint: $(BIN) Makefile
	GOBIN=$(abspath $(@D)) $(GO) install \
		github.com/golangci/golangci-lint/v2/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

$(BIN)/protovalidate-conformance: $(BIN) Makefile
	GOBIN=$(abspath $(BIN)) $(GO) install \
    	github.com/bufbuild/protovalidate/tools/protovalidate-conformance@$(PROTOVALIDATE_VERSION)

.PHONY: protovalidate-conformance-go
protovalidate-conformance-go: $(BIN)
	GOBIN=$(abspath $(BIN)) $(GO) install ./internal/cmd/protovalidate-conformance-go
