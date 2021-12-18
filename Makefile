build:
	@go build -v ./...

clean:
	@rm -rf coverage.out

fmt:
	@gofmt -w .

lint:
	@$(shell go env GOPATH)/bin/golangci-lint run --color always --concurrency 4 --verbose
test:
	@go test -v -failfast -parallel 1 -p 1 -coverprofile=coverage.out $(packages)
ifeq ($(coverage), yes)
	@go tool cover -func=coverage.out
endif

view_test_coverage:
	@go tool cover -html=coverage.out
