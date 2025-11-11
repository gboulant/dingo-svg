modname=$(shell go list -f '{{.Module}}')
pkgname=$(shell go list -f '{{.Name}}')
pathname=$(shell go list)
basename=$(shell basename ${pathname})
PROGNAME=${basename}

all: testall

info:
	@echo "Module name  : ${modname}"
	@echo "Package name : ${pkgname}"
	@echo "Path name    : ${pathname}"
	@echo "Base base    : ${basename}"
	@echo "PROGNAME     : ${PROGNAME}"

test::
	@go test

build:
	@go build

exec: build
	@./${PROGNAME}

testall: test exec

clean:
	@go clean
	@rm -f output.*
