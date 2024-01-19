generate:
	go generate ./...

lint:
	golangci-lint run ./...

test:
	go test ./... -count=1


.PHONY: generate lint test
