.PHONY: test

test:
	go test -cover ./...

test-cover:
	mkdir bin || true
	go test -cover ./...
	go test -coverprofile=./bin/cover.out -json  $(shell go list ./... | grep -v /cmd/) > ./bin/testreport.json

vet:
	go vet ./...

lint:
	golangci-lint run

checks: vet lint