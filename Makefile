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

TEST_PATH := internal/tlcodegen/test
TLS_PATH := $(TEST_PATH)/tls
GEN_PATH := $(TEST_PATH)/gen
TLOS_PATH := $(GEN_PATH)
BASIC_TL_PATH := github.com/vkcom/tl/pkg/basictl

TL_BYTE_VERSIONS := ch_proxy.,ab.
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
		--copyrightPath=./COPYRIGHT \
		--pkgPath=github.com/vkcom/tl/internal/tlast/gentlo/tl \
		--basicPkgPath=github.com/vkcom/tl/pkg/basictl \
		--generateRPCCode=false \
		--outdir=./internal/tlast/gentlo \
		./internal/tlast/tls.tl

.PHONY: gen_check
gen_check: build
	@./target/bin/tlgen --split-internal -v --language=go \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/schema \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/schema/tl \
		--tloPath=./$(TLOS_PATH)/test.tlo \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		--generateLegacyJsonRead=false \
		--generateRPCCode=false \
		./$(TLS_PATH)/schema.tl

.PHONY: gen
gen: gen_check
	@echo "Checking that generated code actually compiles..."
	@$(GO) build ./$(GEN_PATH)/schema/...

.PHONY: gen_dev
gen_dev: qtpl gen

.PHONY: goldmaster_nocompile
goldmaster_nocompile: build
	@./target/bin/tlgen --language=go --split-internal -v \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/cases \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/cases/tl \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=cases_bytes. \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		--generateRPCCode=false \
		./$(TLS_PATH)/cases.tl
	@./target/bin/tlgen --language=go --split-internal -v \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/goldmaster \
		--generateSchemaDocumentation \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/goldmaster/tl \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		--generateRPCCode=false \
		--canonicalFormPath=./$(TLS_PATH)/goldmaster_canonical.tl \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
	@./target/bin/tlgen --language=go -v \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/goldmaster_nosplit \
		--generateSchemaDocumentation \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/goldmaster_nosplit/tl \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		--tloPath=./$(TLOS_PATH)/goldmaster.tlo \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		--generateRPCCode=false \
		--canonicalFormPath=./$(TLS_PATH)/goldmaster_canonical.tl \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl

.PHONY: goldmaster
goldmaster: goldmaster_nocompile
	@echo "Checking that generated code actually compiles..."
	$(GO) build ./$(GEN_PATH)/cases/...
	$(GO) build ./$(GEN_PATH)/goldmaster/...
	$(GO) build ./$(GEN_PATH)/goldmaster_nosplit/...

.PHONY: gen_tlo
gen_tlo: build # do not set --basicPkgPath, or you'll have hard time updating basictl
	@./target/bin/tlgen --language=go \
		--copyrightPath=./COPYRIGHT \
		--pkgPath=github.com/vkcom/tl/internal/tlast/gentlo/tl \
		--generateRPCCode=false \
		--outdir=./internal/tlast/gentlo \
		./internal/tlast/tls.tl


.PHONY: gen_all
gen_all: tlo-bootstrap gen goldmaster

qtpl:
	@if ! [ -x "$(command -v qtc)"]; then \
		echo "qtc could not be found"; \
		echo "install it using"; \
		echo "go get -u github.com/valyala/quicktemplate"; \
		echo "go get -u github.com/valyala/quicktemplate/qtc"; \
	else \
		qtc -dir=./internal -skipLineComments; \
	fi

# TODO: in progress...
#.PHONY: cpp
#cpp: build
#	@./target/bin/tlgen -language=cpp -v \
#		--outdir=./test/cpp \
#		--basicPkgPath=gitlab.mvk.com/go/vkgo/pkg/basictl \
#		./test/cpp.tl
#	g++ -o test/test_cpp test/test_cpp.cpp test/cpp/all.cpp -std=c++17 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter
