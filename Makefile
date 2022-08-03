all: buildAndRun

run:
	./bin/cmd

build:
	go build -o bin/ ./cmd/

buildAndRun:
	go build -o bin/ ./cmd/ && ./bin/cmd

up:
	migrate -database postgres://admin:fuad1234@localhost:5432/go-postgres?sslmode=disable -path internal/database/migrations up
