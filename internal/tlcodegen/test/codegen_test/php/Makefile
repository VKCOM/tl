build-tlgen:
	@cd ../../../../../; \
	$(MAKE) build;

build-php-files: build-tlgen
	@cd ../../../../../; \
	./target/bin/tlgen \
	   --language=php \
	   --php-serialization-bodies \
	   --php-generate-meta \
       --php-generate-factory \
       --php-ignore-unused-types=false \
       --php-rpc-support=false \
       --php-inplace-simple-structs=false \
	   \
	   --outdir=./internal/tlcodegen/test/gen/cases_php \
	   \
	   --schemaURL="https://github.com/VKCOM/tl/blob/master/internal/tlcodegen/test/tls/cases.tl" \
       --schemaCommit="php gen cases test" \
       --schemaTimestamp=1737019773 \
	   \
	   ./internal/tlcodegen/test/tls/cases.tl

run-all-tests: build-php-files
	php test_cases.php;
	php test_basictl.php;