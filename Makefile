all: build

build:
	go build -o bin/ ./cmd/

up:
	migrate -database postgres://admin:fuad1234@localhost:5432/go-postgres?sslmode=disable -path internal/database/migrations up
