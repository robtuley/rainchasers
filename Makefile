COMMIT_HASH=$(shell git rev-parse HEAD)
GO_BUILD=go build -ldflags "-s -w -X main.build=$(COMMIT_HASH)" -a -installsuffix netgo -tags netgo -o ./app

eaday: test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./cmd/eaday

ea: test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./cmd/ea

sepa: test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./cmd/sepa 

nrw: test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./cmd/nrw 

store: test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./cmd/store 

check: test vet lint errcheck

generate:
	GO111MODULE=on go generate ./...

vet: generate 
	GO111MODULE=on go vet ./...

test: generate vet
	GO111MODULE=on go test -race ./...

lint: generate
	GO111MODULE=on golint ./...

errcheck:
	go get -u github.com/kisielk/errcheck
	CGO_ENABLED=0 errcheck ./...

.PHONY: check generate vet test lint errcheck eaday ea sepa store