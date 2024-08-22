build:
	@go build -o bin/gocom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/gocom
