all: run

build: build_encode

build_encode:
	@go build -o tmp/encode ./cmd/encode/encode.go

run:
	@go run ./cmd/encode/encode.go
