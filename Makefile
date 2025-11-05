all: testall

test:
	@go test

demos.%:
	@make -C demos/d01.convexhull $*

testall: test demos.testall

clean: demos.clean
	@go clean
	@rm -f output.*