.PHONY: fmt
fmt:
	gofmt -w .

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
