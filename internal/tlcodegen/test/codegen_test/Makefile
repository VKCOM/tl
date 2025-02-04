
run-all-languages-tests:
	make run-php-tests; \
 	make run-cpp-tests; \
 	make run-golang-tests;

run-php-tests:
	cd php && make run-all-tests

run-cpp-tests:
	cd cpp && make run-all-tests

run-golang-tests:
	cd ../../../.. && make goldmaster_nocompile;
	cd golang;
	go test ./...

INCORRECT_TLS := \
	backward_compatibilty_samples/schema-incorrect-remove-union-variants.tl \
	backward_compatibilty_samples/schema-incorrect-add-constructors-to-nonunion-with-bare-usages.tl \

run-linter-check-diff-test:
	cd ../../../.. && make build;
	for sample in $(INCORRECT_TLS); do \
	    if ../../../../target/bin/tlgen \
			--linter-mode=check-diff \
			--schema-to-compare=../tls/backward_compatibilty_samples/schema-correct.tl \
			../tls/$${sample}; \
		then \
		  	echo DIFF CHECK MUST FAIL ON $${sample}; \
			exit 1; \
		fi \
	done