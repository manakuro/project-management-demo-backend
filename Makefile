# Set up tools
install:
	brew install yq
	brew install pre-commit
	pre-commit --version
	pre-commit install
	go1.16.9 install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go1.16.9 install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1
	go1.16.9 install golang.org/x/tools/cmd/goimports@latest
	go1.16.9 install github.com/cosmtrek/air@v1.27.3

# Start dev server
start:
	air

# Set up database
setup_db:
	./bin/init_db.sh

# Migrate scheme in ent to database
migrate_schema:
	go1.16.9 run ./cmd/migration/main.go
migrate_schema_staging:
	APP_ENV=staging go1.16.9 run ./cmd/migration/main.go

migrate_up:
	migrate -path $$(yq e '.development.path' db/config.yaml) -database $$(yq e '.development.database' db/config.yaml) up

migrate_down:
	migrate -path $$(yq e '.development.path' db/config.yaml) -database $$(yq e '.development.database' db/config.yaml) down

# Connect database through proxy.
connect_db_staging:
	pscale connect project_management_demo staging

# Seed data
seed:
	go1.16.9 run ./cmd/seed/main.go
seed_staging:
	AppEnv=staging go1.16.9 run ./cmd/seed/main.go

ent_generate:
	go1.16.9 generate ./ent --feature sql/upsert --idtype string

generate:
	go1.16.9 generate ./...

schema_description:
	ent describe ./ent/schema

# Testing
setup_test_db:
	./bin/init_db_test.sh

test_repository:
	go1.16.9 test ./pkg/adapter/repository/...

# E2E
setup_e2e_db:
	./bin/init_db_e2e.sh

e2e:
	go1.16.9 test ./test/e2e/...

.PHONY: install setup_db migrate_up migrate_down start migrate_schema schema_description ent_generate setup_test_db setup_e2e_db e2e test_repository seed migrate_schema_staging seed_staging
