build:
	@go build -o bin/goecom cmd/main.go

test:
	@go test -v ./...

run: build
	@./bin/gocom
