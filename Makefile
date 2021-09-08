.PHONY: build
build:
	go build -o ./bin/gotp ./cmd

.PHONY: test
test:
	go test -v ./cmd