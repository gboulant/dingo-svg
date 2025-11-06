all: testall

test:
	@go test

build:
	@go build

exec: build
	@./d02.surface3d

testall: test exec

clean:
	@go clean
	@rm -f output.*
