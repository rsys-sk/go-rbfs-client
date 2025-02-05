.PHONY: test

test:
	go test -cover ./...

fumpt:
	gofumpt -l -w .
	gci write -s Standard -s Default -s "Prefix(gitlab.rtbrick.net)" .

lint:
	golangci-lint run

checks: fumpt lint test