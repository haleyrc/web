.PHONY: fmt test tidy verify vet

all: fmt tidy verify vet test

fmt:
	gofmt -s -w .

test:
	go test -v -count=1 -shuffle=on ./...

tidy:
	go mod tidy

verify:
	go mod verify

vet:
	go vet ./...
