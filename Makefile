all: testall

test:
	@go test

demos.%:
	@make -C demos/d01.convexhull $*
	@make -C demos/d02.isometry $*
	@make -C demos/d03.remarkable $*
	@make -C demos/d04.guitarneck $*
	@make -C demos/d05.appartmap $*

testall: test demos.testall

clean: demos.clean
	@go clean
	@rm -f output.*
