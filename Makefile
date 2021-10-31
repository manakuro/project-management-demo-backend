# Set up tools
.PHONY: install
install:
	brew install pre-commit
	pre-commit --version
	pre-commit install
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1
	go install golang.org/x/tools/cmd/goimports@latest

.PHONY: setup_db
setup_db:
	./bin/init_db.sh
