
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
	schema-incorrect-remove-union-variants.tl \
	schema-incorrect-add-constructors-to-nonunion-with-bare-usages.tl \
	schema-incorrect-remove-fields.tl \
	schema-incorrect-new-fields-without-fieldmask.tl \
	schema-incorrect-remove-template-arguments.tl \
	schema-incorrect-remove-fieldmask.tl \
	schema-incorrect-add-fieldmask-to-old-field.tl \
	schema-incorrect-change-bit-in-fieldmask.tl \
    schema-incorrect-change-reference-in-fieldmask.tl \
    schema-incorrect-all-fieldmask-bits-are-used.tl \
    schema-incorrect-bit-that-already-used-1.tl \
    schema-incorrect-bit-that-already-used-2-from-outer-scope.tl \
    schema-incorrect-bit-that-already-used-3-not-used-but-forbiden.tl \
    schema-incorrect-bit-that-already-used-4-new-fieldmask-with-some-bits-used-in-schema-already.tl \
    schema-incorrect-add-only-fieldmask-but-without-usage.tl \
    schema-incorrect-new-arguments-added-by-fieldmask-forgotten.tl \
    schema-incorrect-new-arguments-added-with-new-fieldmask-but-old-one-used.tl \
    schema-incorrect-new-arguments-added-without-new-fieldmask.tl \
    schema-incorrect-new-function-first-arg-not-nat.tl \
    schema-incorrect-new-function-non-first-args-without-fieldmask.tl \

run-linter-check-diff-test:
	cd ../../../.. && make build;
	for sample in $(INCORRECT_TLS); do \
	    if ../../../../target/bin/tlgen \
			--schema-to-compare=../tls/backward_compatibilty_samples/schema-correct.tl \
			../tls/backward_compatibilty_samples/$${sample}; \
		then \
		  	echo DIFF CHECK MUST FAIL ON $${sample}; \
			exit 1; \
		fi \
	done