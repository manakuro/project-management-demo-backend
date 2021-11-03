# Set up tools
install:
	brew install yq
	brew install pre-commit
	pre-commit --version
	pre-commit install
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.15.1
	go install golang.org/x/tools/cmd/goimports@latest

setup_db:
	./bin/init_db.sh

migrate_up:
	migrate -path $$(yq e '.development.path' db/config.yaml) -database $$(yq e '.development.database' db/config.yaml) up

migrate_down:
	migrate -path $$(yq e '.development.path' db/config.yaml) -database $$(yq e '.development.database' db/config.yaml) down

.PHONY: install setup_db migrate_up migrate_down
