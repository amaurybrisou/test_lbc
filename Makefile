build:
	@docker	build . --target bin -t test_lbc

test:
	@docker build --no-cache . --target unit-test -t test_lbc_tests

run: clean
	@docker run -p 8000:8000 --name test_lbc_container test_lbc

clean:
	docker rm test_lbc_container