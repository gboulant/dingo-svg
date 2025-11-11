all: testall

test:
	@go test

demos.%:
	@make -C demos/d01.convexhull $*
	@make -C demos/d02.surface3d $*
	@make -C demos/d03.remarkable $*
	@make -C demos/d04.guitarneck $*

testall: test demos.testall

clean: demos.clean
	@go clean
	@rm -f output.*
