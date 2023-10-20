build:
	@go build

test:
	@go test -v

run:
	@go run main.go

lint: 
	@golangci-lint run --fix --print-resources-usage --allow-parallel-runners --color always

dbrun:
	@surreal start memory -A --auth --user root --pass root
