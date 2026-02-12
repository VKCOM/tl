GO = go

TEST_PATH := internal/tlcodegen/test
TLS_PATH := $(TEST_PATH)/tls
GEN_PATH := $(TEST_PATH)/gen
TLOS_PATH := $(GEN_PATH)
BASIC_TL_PATH := github.com/vkcom/tl/pkg/basictl

TL_BYTE_VERSIONS := ch_proxy.,ab.

all: build

.PHONY: build
build: # build static binary to run on many linux variants
	CGO_ENABLED=0 $(GO) build -buildvcs=true -o target/bin/tlgen ./cmd/tlgen
	CGO_ENABLED=0 $(GO) build -buildvcs=false -o target/bin/tl2gen ./cmd/tl2gen
	CGO_ENABLED=0 $(GO) build -buildvcs=false -o target/bin/tl2client ./cmd/tl2client

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
		--outdir=./$(GEN_PATH)/casesTL1 \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/casesTL1/tl \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=cases_bytes. \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		--checkLengthSanity=false \
		./$(TLS_PATH)/cases.tl
	@./target/bin/tlgen --language=go --split-internal -v \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/cases \
		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/cases/tl \
		--basicPkgPath=$(BASIC_TL_PATH) \
		--generateByteVersions=cases_bytes. \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		--checkLengthSanity=false \
		./$(TLS_PATH)/cases.tl
	@./target/bin/tlgen --language=go --split-internal -v \
		--tl2WhiteList=* \
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
		--checkLengthSanity=false \
		--canonicalFormPath=./$(TLS_PATH)/goldmaster_canonical.tl \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
	@./target/bin/tlgen --language=go -v \
		--tl2WhiteList=* \
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

.PHONY: migrate_to_tl2
migrate_to_tl2: build
	@./target/bin/tlgen -v \
		--tl2-migration-file=./$(TLS_PATH)/cases.tl2 \
		./$(TLS_PATH)/cases.tl

.PHONY: goldmaster_tl2_nocompile
goldmaster_tl2_nocompile: build migrate_to_tl2
	#	@./target/bin/tlgen --language=go --split-internal -v \
	#		--tl2WhiteList=* \
	#		--copyrightPath=./COPYRIGHT \
	#		--outdir=./$(GEN_PATH)/casesTL2 \
	#		--pkgPath=github.com/vkcom/tl/$(GEN_PATH)/casesTL2/tl \
	#		--basicPkgPath=$(BASIC_TL_PATH) \
	#		--generateByteVersions=cases_bytes. \
	#		--generateRandomCode \
	#		--generateLegacyJsonRead=false \
	#		--checkLengthSanity=false \
	#		./$(TLS_PATH)/cases.tl2

.PHONY: goldmaster
goldmaster: goldmaster_nocompile goldmaster_tl2_nocompile
	@echo "Checking that generated code actually compiles..."
	$(GO) build ./$(GEN_PATH)/cases/...
	$(GO) build ./$(GEN_PATH)/casesTL1/...
	# $(GO) build ./$(GEN_PATH)/casesTL2/...
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
	go run github.com/valyala/quicktemplate/qtc -dir=./internal -skipLineComments;

.PHONY: cpp_move_basictl
cpp_move_basictl:
	@cd scripts && \
	bash move-basictl-cpp.sh

.PHONY: cpp_build
cpp_build:
	g++ -o $(GEN_PATH)/test_cpp $(GEN_PATH)/test_cpp.cpp $(GEN_PATH)/cpp/all.cpp -std=c++17 -O3 -Wno-noexcept-type -g -Wall -Wextra -Werror=return-type -Wno-unused-parameter

cpp_template_gen: build
	@./target/bin/tlgen -language=cpp -v \
		--cpp-generate-meta=true \
		--cpp-generate-factory=true \
		--outdir=./$(GEN_PATH)/$(OUTPUT_PATH) \
		./$(TLS_PATH)/$(TL_FILE); \

.PHONY: cpp_gen
cpp_gen:
	$(MAKE) cpp_template_gen OUTPUT_PATH=cpp TL_FILE=cpp.tl

.PHONY: cpp_gen
cpp_gen_test_data:
	$(MAKE) cpp_template_gen OUTPUT_PATH=cases_cpp TL_FILE=cases.tl
	$(MAKE) cpp_template_gen OUTPUT_PATH=schema_cpp TL_FILE=schema.tl
	$(MAKE) cpp_template_gen OUTPUT_PATH=goldmaster_cpp TL_FILE=goldmaster.tl

.PHONY: cpp
cpp:
	$(MAKE) cpp_gen
	#$(MAKE) cpp_build

.PHONY: test_multi_lang_cases
test_multi_lang_cases:
	@cd internal/tlcodegen/test/codegen_test/; \
	$(MAKE) run-all-languages-tests;

.PHONY: test
test:
	@$(GO) test $(shell $(GO) list ./... | grep -v internal/tlcodegen/test/gen/)

# target should be as close as possible to github actions used to enable merge
.PHONY: check
check: build test lint
	@$(GO) run honnef.co/go/tools/cmd/staticcheck@v0.5.1 --version
	@$(GO) run honnef.co/go/tools/cmd/staticcheck@v0.5.1 ./... # update version together with github actions

.PHONY: lint
lint:
	@$(GO) run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6 run internal/tlcodegen/test/gen/...

.PHONY: testpure
testpure: build
	@./target/bin/tlgen --language=go -v --split-internal \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./cmd/tl2gen/genold \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/vkcom/tl/cmd/tl2gen/genold/tl \
		--basicPkgPath=github.com/vkcom/tl/pkg/basictl \
		--generateByteVersions=ch_proxy.,ab. \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		./cmd/tl2client/test.tl
	@./target/bin/tl2gen --language=go -v --split-internal \
		--tl2WhiteList=* \
		--newDicts \
		--copyrightPath=./COPYRIGHT \
		--outdir=./cmd/tl2gen/gennew \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/vkcom/tl/cmd/tl2gen/gennew/tl \
		--basicPkgPath=github.com/vkcom/tl/pkg/basictl \
		--generateByteVersions=ch_proxy.,ab. \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		./cmd/tl2client/test.tl
	@echo "Checking that generated code actually compiles..."
	time $(GO) build ./cmd/tl2gen/genold/...
	time $(GO) build ./cmd/tl2gen/gennew/...

# /home/user/go/src/gitlab.mvk.com/go/vkgo/pkg/vktl/combined.tl
# ./cmd/tl2client/test.tl
