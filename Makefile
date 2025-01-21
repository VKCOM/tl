GO = go

BUILD_VERSION    := $(if $(BUILD_VERSION),$(BUILD_VERSION),$(shell git describe --tags --always --dirty))
BUILD_COMMIT     := $(if $(BUILD_COMMIT),$(BUILD_COMMIT),$(shell git log --format="%H" -n 1))
BUILD_COMMIT_TS  := $(if $(BUILD_COMMIT_TS),$(BUILD_COMMIT_TS),$(shell git log --format="%ct" -n 1))
BUILD_BRANCH     := $(if $(BUILD_BRANCH),$(BUILD_BRANCH),$(shell git rev-parse --abbrev-ref HEAD))
BUILD_TIME       := $(if $(BUILD_TIME),$(BUILD_TIME),$(shell date +%FT%T%z))
BUILD_MACHINE    := $(if $(BUILD_MACHINE),$(BUILD_MACHINE),$(shell uname -n -m -r -s))
BUILD_GO_VERSION := $(if $(BUILD_GO_VERSION),$(BUILD_GO_VERSION),$(shell $(GO) version | cut -d' ' -f3))

COMMON_BUILD_VARS := \
  -X 'github.com/vkcom/tl/pkg/build.buildTimestamp=$(BUILD_TIME)' \
  -X 'github.com/vkcom/tl/pkg/build.machine=$(BUILD_MACHINE)' \
  -X 'github.com/vkcom/tl/pkg/build.commit=$(BUILD_COMMIT)' \
  -X 'github.com/vkcom/tl/pkg/build.version=$(BUILD_VERSION)' \
  -X 'github.com/vkcom/tl/pkg/build.commitTimestamp=$(BUILD_COMMIT_TS)' \
  -X 'github.com/vkcom/tl/pkg/build.branchName=$(BUILD_BRANCH)' \

COMMON_LDFLAGS = $(COMMON_BUILD_VARS) -extldflags '-O2'

TEST_PATH := internal/tlcodegen/test
TLS_PATH := $(TEST_PATH)/tls
GEN_PATH := $(TEST_PATH)/gen
TLOS_PATH := $(GEN_PATH)
BASIC_TL_PATH := github.com/vkcom/tl/pkg/basictl

TL_BYTE_VERSIONS := ch_proxy.,ab.

all: build

.PHONY: build
build:
	@$(GO) build -ldflags "$(COMMON_LDFLAGS)" -buildvcs=false -o target/bin/tlgen ./cmd/tlgen

tlo-bootstrap: build
	@./target/bin/tlgen -v --language=go \
		--copyrightPath=./COPYRIGHT \
		--pkgPath=github.com/vkcom/tl/internal/tlast/gentlo/tl \
		--basicPkgPath=github.com/vkcom/tl/pkg/basictl \
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
		./$(TLS_PATH)/cases.tl
	@./target/bin/tlgen --language=go --split-internal -v \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/goldmaster \
		--generateSchemaDocumentation \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/goldmaster/tl \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		--canonicalFormPath=./$(TLS_PATH)/goldmaster_canonical.tl \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
	@./target/bin/tlgen --language=go -v \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/goldmaster_nosplit \
		--generateSchemaDocumentation \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/goldmaster_nosplit/tl \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		--tloPath=./$(TLOS_PATH)/goldmaster.tlo \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
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

.PHONY: cpp_build
cpp_build:
	g++ -o $(GEN_PATH)/test_cpp $(GEN_PATH)/test_cpp.cpp $(GEN_PATH)/cpp/all.cpp -std=c++17 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter

cpp_template_gen: build
	@./target/bin/tlgen -language=cpp -v \
		--outdir=./$(GEN_PATH)/$(OUTPUT_PATH) \
		--basicPkgPath=$(BASIC_TL_PATH) \
		./$(TLS_PATH)/$(TL_FILE); \

.PHONY: cpp_gen
cpp_gen:
	$(MAKE) cpp_template_gen OUTPUT_PATH=cpp TL_FILE=cpp.tl

.PHONY: cpp_gen
cpp_gen_test_data:
	$(MAKE) cpp_template_gen OUTPUT_PATH=cases_cpp TL_FILE=cases.tl
	$(MAKE) cpp_template_gen OUTPUT_PATH=schema_cpp TL_FILE=schema.tl

.PHONY: cpp
cpp:
	$(MAKE) cpp_gen
	$(MAKE) cpp_build

.PHONY: test_multi_lang_cases
test_multi_lang_cases:
	@cd internal/tlcodegen/test/codegen_test/; \
	$(MAKE) run-all-languages-tests;

.PHONY: test
test:
	@$(GO) test $(shell $(GO) list ./... | grep -v internal/tlcodegen/test/gen/)

# target should be as close as possible to github actions used to enable merge
.PHONY: check
check: build test
	@$(GO) run honnef.co/go/tools/cmd/staticcheck@v0.5.1 ./... # update version together with github actions