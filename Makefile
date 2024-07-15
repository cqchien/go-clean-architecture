.PHONY: run build test


ifeq (create-migrate,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  MIGRATION_NAME := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(MIGRATION_NAME):;@:)
endif

TIMESTAMP = $(shell date +%Y%m%d%H%M%S)
MIGRATION_FILE = migrations/$(TIMESTAMP)_$(MIGRATION_NAME).go

# ==============================================================================
# Dev

run:
	go run ./cmd/api/main.go

build:
	go build ./cmd/api/main.go

create-migrate:
	@cp migrations/migration_template.go $(MIGRATION_FILE)
	@sed -i '' "s/TIMESTAMP_DESCRIPTION/$(TIMESTAMP)_$(MIGRATION_NAME)/g" $(MIGRATION_FILE)
	@echo "Created migration $(MIGRATION_FILE)"

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
