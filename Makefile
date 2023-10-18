test:
	@go test -v

run:
	@go run cmd/main.go

lint: 
	@golangci-lint run --fix --print-resources-usage --allow-parallel-runners --color always