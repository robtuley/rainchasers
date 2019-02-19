COMMIT_HASH=$(shell git rev-parse HEAD)
GO_BUILD=go build -ldflags "-s -w -X main.build=$(COMMIT_HASH)" -a -installsuffix netgo -tags netgo -o ./app

ea: vet test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./ea/latest 

sepa: vet test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./sepa 

firestore: vet test
	CGO_ENABLED=0 GO111MODULE=on $(GO_BUILD) ./firestore 

check: test vet lint errcheck

vet:
	GO111MODULE=on go vet ./...

test:
	GO111MODULE=on go test -race ./...

lint:
	GO111MODULE=on golint ./...

errcheck:
	go get -u github.com/kisielk/errcheck
	CGO_ENABLED=0 errcheck ./...

.PHONY: check vet test lint errcheck ea sepa firestore