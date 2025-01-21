
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