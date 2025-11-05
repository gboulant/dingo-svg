all: test demos

test:
	@go test

demos: demos.testall

demos.%:
	@make -C demos/d01.convexhull $*

clean: demos.clean
	@go clean
	@rm -f output.*