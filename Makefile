GO = go

BUILD_VCS ?= true

TEST_PATH := internal/tlcodegen/test
TLS_PATH := $(TEST_PATH)/tls
GEN_PATH := $(TEST_PATH)/gen
TLOS_PATH := $(GEN_PATH)
BASIC_TL_PATH := github.com/VKCOM/tl/pkg/basictl

TL_BYTE_VERSIONS := ch_proxy.,ab.

all: build

.PHONY: build
build: # build static binary to run on many linux variants
	CGO_ENABLED=0 $(GO) build -buildvcs=$(BUILD_VCS) -o target/bin/tlgen ./cmd/tlgen
	CGO_ENABLED=0 $(GO) build -buildvcs=$(BUILD_VCS) -o target/bin/tl2gen ./cmd/tl2gen
	CGO_ENABLED=0 $(GO) build -buildvcs=$(BUILD_VCS) -o target/bin/tl2client ./cmd/tl2client

.PHONY: gen_bootstrap
gen_bootstrap: build # do not set --basicPkgPath, or you'll have hard time updating basictl
	@./target/bin/tl2gen --language=go \
		--copyrightPath=./COPYRIGHT \
		--pkgPath=github.com/VKCOM/tl/internal/tlast/gentlo/tl \
		--outdir=./internal/tlast/gentlo \
		./internal/tlast/tls.tl

.PHONY: gen_check
gen_check: build
	@./target/bin/tl2gen --split-internal -v --language=go \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/schema \
		--pkgPath=github.com/VKCOM/tl/$(GEN_PATH)/schema/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		./$(TLS_PATH)/schema.tl
	@./target/bin/tl2gen --language=tlo \
		--outfile=./$(TLOS_PATH)/test.tlo \
		--schemaTimestamp=1761907954 \
		./$(TLS_PATH)/schema.tl

.PHONY: gen
gen: gen_check
	@echo "Checking that generated code actually compiles..."
	@$(GO) build ./$(GEN_PATH)/schema/...

.PHONY: gen_dev
gen_dev: qtpl gen

.PHONY: goldmaster_nocompile
goldmaster_nocompile: build
	@./target/bin/tl2gen --language=go -v \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/casesTL1 \
		--pkgPath=github.com/VKCOM/tl/$(GEN_PATH)/casesTL1/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--basicRPCPath=github.com/VKCOM/tl/pkg/rpc \
		--generateRPCCode \
		--generateByteVersions=cases_bytes. \
		--generateRandomCode \
		--checkLengthSanity=false \
		./$(TLS_PATH)/cases.tl
	@./target/bin/tl2gen --language=go -v \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/cases \
		--pkgPath=github.com/VKCOM/tl/$(GEN_PATH)/cases/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--basicRPCPath=github.com/VKCOM/tl/pkg/rpc \
		--generateRPCCode \
		--generateByteVersions=cases_bytes. \
		--generateRandomCode \
		--checkLengthSanity=false \
		./$(TLS_PATH)/cases.tl
	@./target/bin/tl2gen --language=go --split-internal -v \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/goldmaster \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/VKCOM/tl/$(GEN_PATH)/goldmaster/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--basicRPCPath=github.com/VKCOM/tl/pkg/rpc \
		--generateRPCCode \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		--generateRandomCode \
		--checkLengthSanity=false \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
	@./target/bin/tl2gen --language=go -v \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./$(GEN_PATH)/goldmaster_nosplit \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/VKCOM/tl/$(GEN_PATH)/goldmaster_nosplit/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--basicRPCPath=github.com/VKCOM/tl/pkg/rpc \
		--generateRPCCode \
		--generateByteVersions=$(TL_BYTE_VERSIONS) \
		--generateRandomCode \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
	@./target/bin/tl2gen --language=tlo -v \
		--outfile=./$(TLOS_PATH)/goldmaster.tlo \
		--schemaTimestamp=301822800 \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
	@./target/bin/tl2gen --language=canonical -v \
		--outfile=./$(TLS_PATH)/goldmaster_canonical.tl \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
	@./target/bin/tl2gen --language=tljson.html -v \
		--outfile=./$(GEN_PATH)/goldmaster_nosplit/tljson.html \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl

.PHONY: migrate_to_tl2
migrate_to_tl2: build
	rm ./$(TLS_PATH)/cases.tl2
	@./target/bin/tl2gen --language=tl2migration -v \
  		--tl2WhiteList=* \
		./$(TLS_PATH)/cases.tl
	git checkout ./$(TLS_PATH)/cases.tl # restore cases, which would contain only TL1 Bool/vector/etc.

.PHONY: goldmaster_tl2_nocompile
goldmaster_tl2_nocompile: build migrate_to_tl2
	@./target/bin/tl2gen --language=go -v \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./internal/tlcodegen/test/gen/casesTL2 \
		--pkgPath=github.com/VKCOM/tl/internal/tlcodegen/test/gen/casesTL2/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--basicRPCPath=github.com/VKCOM/tl/pkg/rpc \
		--generateRPCCode \
		--generateByteVersions=cases_bytes. \
		--generateRandomCode \
		--checkLengthSanity=false \
		./$(TLS_PATH)/cases.tl2

.PHONY: goldmaster
goldmaster: goldmaster_nocompile goldmaster_tl2_nocompile
	@echo "Checking that generated code actually compiles..."
	$(GO) build ./$(GEN_PATH)/cases/...
	$(GO) build ./$(GEN_PATH)/casesTL1/...
	$(GO) build ./$(GEN_PATH)/casesTL2/...
	$(GO) build ./$(GEN_PATH)/goldmaster/...
	$(GO) build ./$(GEN_PATH)/goldmaster_nosplit/...

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

.PHONY: cpp_test
cpp_test:
	@echo "TODO"

.PHONY: test_multi_lang_cases
test_multi_lang_cases:
	@cd internal/tlcodegen/test/codegen_test/; \
	$(MAKE) run-all-languages-tests;

.PHONY: go_test
go_test:
	@$(GO) test $(shell $(GO) list ./... | grep -v internal/tlcodegen/test/gen/)

.PHONY: test
test: go_test php_test

# target should be as close as possible to github actions used to enable merge
.PHONY: check
check: build test lint
	@$(GO) run honnef.co/go/tools/cmd/staticcheck@v0.7.0 --version
	@$(GO) run honnef.co/go/tools/cmd/staticcheck@v0.7.0 ./cmd/... ./internal/... ./pkg/... # update version together with github actions

.PHONY: lint
lint:
	@$(GO) run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.6 run internal/tlcodegen/test/gen/...

.PHONY: testpure
testpure: build
	@./target/bin/tl2gen --language=tl2migration -v \
		--tl2migrationDevMode \
		--tl2WhiteList=curl.,a.,statshouseApi.,videoAdsService. \
		./cmd/tl2client/test.tl ./cmd/tl2client/test.tl2
	@./target/bin/tl2gen --language=go -v --split-internal \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./target/genold \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/VKCOM/tl/target/genold/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--generateByteVersions=ch_proxy.,ab. \
		--generateRandomCode \
		./cmd/tl2client/test.tl ./cmd/tl2client/test.tl2
		# genold will not compile due to import statements with gennew
		# @echo "Checking that generated code actually compiles..."
		# time $(GO) build ./target/genold/...

.PHONY: testpure
testpuremigr: build
	@./target/bin/tl2gen \
		--language=go -v --split-internal \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./target/gennew \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		--pkgPath=github.com/VKCOM/tl/target/gennew/tl \
		--basicPkgPath=github.com/VKCOM/tl/pkg/basictl \
		--generateByteVersions=ch_proxy.,ab. \
		--generateRandomCode \
		./cmd/tl2client/test_migr.tl ./cmd/tl2client/test_migr.tl2
	@echo "Checking that generated code actually compiles..."
	time $(GO) build ./target/gennew/...


.PHONY: testrust
testrust: build
	@./target/bin/tl2gen --language=rust -v \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--outdir=./target/gentl \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		./cmd/tl2client/test.tl ./cmd/tl2client/test.tl2
	@./target/bin/tl2gen --language=rust -v \
		--tl2WhiteList=* \
		--copyrightPath=./COPYRIGHT \
		--crate-name=gengoldmaster \
		--outdir=./target/gengoldmaster \
		--schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/goldmaster.tl" \
		--schemaCommit=abcdefgh \
		--schemaTimestamp=301822800 \
		./$(TLS_PATH)/goldmaster.tl ./$(TLS_PATH)/goldmaster2.tl ./$(TLS_PATH)/goldmaster3.tl
		# genold will not compile due to import statements with gennew
		# @echo "Checking that generated code actually compiles..."
		# time $(GO) build ./target/genold/...

# /home/user/go/src/gitlab.mvk.com/go/vkgo/pkg/vktl/combined.tl
# ./cmd/tl2client/test.tl

# если тега нет, то будет текущий + 1, иначе надо написать TAG={тег}
.PHONY: update-tag
update-tag:
	@if [ -n "$(TAG)" ]; then \
		bash scripts/update-tag.sh "$(TAG)"; \
	else \
		bash scripts/update-tag.sh; \
	fi

.PHONY: php_test
php_test: build
	$(MAKE) -C internal/tlcodegen/test/codegen_test run-php-tests