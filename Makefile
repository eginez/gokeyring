.PHONY: build test clean install all

BINARY_NAME = gokeyring

all: build

build: clean
	go build -o $(BINARY_NAME) main.go

test:
	go test -v ./...

clean:
	go clean -cache
	go clean -i
	rm -f $(BINARY_NAME) gokeyring

install: build
	GOBIN=$$HOME/bin go install ./...
