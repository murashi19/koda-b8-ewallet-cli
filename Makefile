APP_NAME=ewallet
BINARY=bin/$(APP_NAME)

.PHONY: help run build clean tidy \
		docker-build docker-run \
		migrate-up migrate-down migrate-force migrate-version migrate-create

help:
	@echo "========== E-Wallet Makefile =========="
	@echo "  make run               Run application"
	@echo "  make build             Build binary"
	@echo "  make clean             Remove build artifacts"
	@echo "  make tidy              Clean Go dependencies"
	@echo ""
	@echo "Migration:"
	@echo "  make migrate-up"
	@echo "  make migrate-down"
	@echo "  make migrate-force version=<version>"
	@echo "  make migrate-version"
	@echo "  make migrate-create name=<migration_name>"

run:
	go run ./cmd/$(APP_NAME)

build:
	mkdir -p bin
	go build -o $(BINARY) ./cmd/$(APP_NAME)

clean:
	rm -rf bin

tidy:
	go mod tidy

docker-build:
	docker build -t $(APP_NAME):latest .

docker-run:
	docker run -it --rm $(APP_NAME):latest

migrate-create:
	migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down

migrate-version:
	migrate -path migrations -database "$(DATABASE_URL)" version
