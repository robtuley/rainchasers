COMMIT_HASH=$(shell git rev-parse HEAD)
GO_BUILD=go build -ldflags "-s -w -X main.version=$(COMMIT_HASH)" -a -installsuffix netgo -tags netgo -o ./app

eaday: test
	CGO_ENABLED=0 $(GO_BUILD) ./cmd/eaday

ea: test
	CGO_ENABLED=0 $(GO_BUILD) ./cmd/ea

sepa: test
	CGO_ENABLED=0 $(GO_BUILD) ./cmd/sepa 

nrw: test
	CGO_ENABLED=0 $(GO_BUILD) ./cmd/nrw 

store: test
	CGO_ENABLED=0 $(GO_BUILD) ./cmd/store 

web: generate
	CGO_ENABLED=0 $(GO_BUILD) ./cmd/web 

check: test vet lint errcheck

generate:
	go generate ./...

vet: generate 
	go vet ./...

test: generate vet
	go test -race ./...

.PHONY: check generate vet test eaday ea sepa store web