TARGET = life
VERSION = $(shell git describe --tags --always)
GO_BUILD_ARGS = -ldflags "-X main.version=$(VERSION)"

all: clean build test
env:
	@GOPATH=$(shell pwd)
build: env
	@echo building...
	go build $(GO_BUILD_ARGS) -o $(TARGET)
run: env
	@echo running...
	go run $(GO_BUILD_ARGS) main.go test/small
test: env
	@echo testing...
	go test $(GO_BUILD_ARGS) -v ./...
install: env
	@echo installing...
	go install $(GO_BUILD_ARGS) -o $(TARGET)
clean:
	@echo cleaning...
ifneq ("$(wildcard $(TARGET))","") 
	rm -r $(TARGET)
endif

