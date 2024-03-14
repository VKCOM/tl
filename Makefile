BUILD_VERSION    := $(if $(BUILD_VERSION),$(BUILD_VERSION),$(shell git describe --tags --always --dirty))
BUILD_COMMIT     := $(if $(BUILD_COMMIT),$(BUILD_COMMIT),$(shell git log --format="%H" -n 1))
BUILD_COMMIT_TS  := $(if $(BUILD_COMMIT_TS),$(BUILD_COMMIT_TS),$(shell git log --format="%ct" -n 1))
BUILD_BRANCH     := $(if $(BUILD_BRANCH),$(BUILD_BRANCH),$(shell git rev-parse --abbrev-ref HEAD))
BUILD_TIME       := $(if $(BUILD_TIME),$(BUILD_TIME),$(shell date +%FT%T%z))
BUILD_MACHINE    := $(if $(BUILD_MACHINE),$(BUILD_MACHINE),$(shell uname -n -m -r -s))
BUILD_GO_VERSION := $(if $(BUILD_GO_VERSION),$(BUILD_GO_VERSION),$(shell go version | cut -d' ' -f3))

COMMON_BUILD_VARS := \
  -X 'github.com/vkcom/tl/pkg/build.buildTimestamp=$(BUILD_TIME)' \
  -X 'github.com/vkcom/tl/pkg/build.machine=$(BUILD_MACHINE)' \
  -X 'github.com/vkcom/tl/pkg/build.commit=$(BUILD_COMMIT)' \
  -X 'github.com/vkcom/tl/pkg/build.version=$(BUILD_VERSION)' \
  -X 'github.com/vkcom/tl/pkg/build.commitTimestamp=$(BUILD_COMMIT_TS)' \
  -X 'github.com/vkcom/tl/pkg/build.branchName=$(BUILD_BRANCH)' \

COMMON_LDFLAGS = $(COMMON_BUILD_VARS) -extldflags '-O2'

GO = go

SHA256_CHECKSUM := $(shell go run ./cmd/sha256sum ./internal)
ifndef SHA256_CHECKSUM
$(error SHA256_CHECKSUM failed to set, problem with go run cmd/sha256sum internal)
endif

.PHONY: build

all: build

build:
	@echo "Building tlgen with sha256 checksum $(SHA256_CHECKSUM)"
	@$(GO) build -ldflags "$(COMMON_LDFLAGS)  -X 'github.com/vkcom/tl/internal/tlcodegen.buildSHA256Checksum=$(SHA256_CHECKSUM)'" -buildvcs=false -o target/bin/tlgen ./cmd/tlgen

tlo-bootstrap: build
	@./target/bin/tlgen \
		-pkgPath=github.com/vkcom/tl/internal/tlast/gentlo/tl \
		-basicPkgPath=github.com/vkcom/tl/pkg/basictl \
		-generateRPCCode=false \
		-outdir=./internal/tlast/gentlo \
		-schema=./internal/tlast/tls.tl

qtpl:
	@if ! [ -x "$(command -v qtc)"]; then \
		echo "qtc could not be found"; \
		echo "install it using"; \
		echo "go get -u github.com/valyala/quicktemplate"; \
		echo "go get -u github.com/valyala/quicktemplate/qtc"; \
	else \
		qtc -dir=./internal -skipLineComments; \
	fi

