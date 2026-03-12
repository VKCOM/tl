include ../../Common.mk

.PHONY: gen

gen:
	@$(TL_GEN) --language=go -v \
		--outdir=../rpc/internal/gen \
		--pkgPath=gitlab.corp.mail.ru/vkgo/vkgo/pkg/rpc/internal/gen/tl \
		--basicPkgPath=gitlab.corp.mail.ru/vkgo/vkgo/pkg/basictl \
		--generateByteVersions= \
		--generateRandomCode \
		--generateLegacyJsonRead=false \
		../rpc/rpc.tl
	@echo "Checking that generated code actually compiles..."
	@go build ../rpc/internal/gen/...

fuzz_udp:
	cd udp && go test -fuzz=FuzzTransport -fuzztime=240h
