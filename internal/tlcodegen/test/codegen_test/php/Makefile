build-tlgen:
	@cd ../../../../../; \
	$(MAKE) build;

build-php-files: build-tlgen
	@cd ../../../../../ && \
	./target/bin/tlgen \
	   --language=php \
	   --php-rpc-support=true \
	   --php-serialization-bodies=true \
	   --php-generate-fetchers=true \
	   --php-generate-switcher=true \
	   --php-use-builtin-data-providers=true \
	   --php-add-type-comments=true \
	   --php-generate-fetchers-echo-comment=false \
	   --php-serialization-bodies-whitelist= \
	   \
	   --outdir=./internal/tlcodegen/test/gen/cases_php \
	   ./internal/tlcodegen/test/tls/cases.tl

run-all-tests: build-php-files
	php test_tl_switcher.php;