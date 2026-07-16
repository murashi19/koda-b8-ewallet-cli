APP_NAME=ewallet

run:
	go run ./cmd/$(APP_NAME)

build:
	go build -o bin/$(APP_NAME) ./cmd/$(APP_NAME)

clean:
	rm -rf bin