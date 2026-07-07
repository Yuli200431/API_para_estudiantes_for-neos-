.PHONY: run tidy test

# Corre la API localmente con SQLite (sin Docker)
run:
	go run ./cmd/api

# Descarga y ordena dependencias
tidy:
	go mod tidy

# Corre todos los tests con cobertura
test:
	go test ./... -cover