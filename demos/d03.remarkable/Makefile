all: testall

test:
build:
	@go build

exec: build
	@./d03.remarkable

testall: test exec

clean:
	@go clean
	@rm -f output.*
