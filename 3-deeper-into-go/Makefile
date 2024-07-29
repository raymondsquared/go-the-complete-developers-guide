.PHONY: fmt
fmt:
	gofmt -w .

.PHONY: test
test:
	go test

.PHONY: run
run:
	go run main.go deck.go

.PHONY: build
build:
	go build main.go deck.go
	chmod +x main

.PHONY: execute
execute:
	./main
