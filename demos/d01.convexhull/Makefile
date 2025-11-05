all: test

test:
	@go test

build:
	@go build

exec: build
	@./d01.convexhull

testall: test exec

clean:
	@go clean
	@rm -f output.*