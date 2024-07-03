.PHONY: run build test

# ==============================================================================
# Dev

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

test:
	go test -cover ./...

# ==============================================================================
# Ops
up:
	docker compose -f docker-compose.yml up -d --force-recreate --remove-orphans

down:
	docker compose down

clean:
	docker system prune -f
