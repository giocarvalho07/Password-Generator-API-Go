.PHONY: build run test lint clean docker-build docker-run swagger

APP_NAME=password-generator
CMD_PATH=./cmd/api
BUILD_DIR=./bin

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) $(CMD_PATH)/main.go

run:
	go run $(CMD_PATH)/main.go

test:
	go test -v ./...

lint:
	golangci-lint run

clean:
	rm -rf $(BUILD_DIR)

docker-build:
	docker build -t $(APP_NAME) .

docker-run:
	docker run -p 8080:8080 $(APP_NAME)

swagger:
	swag init -g $(CMD_PATH)/main.go

tidy:
	go mod tidy

fmt:
	gofmt -s -w .

vet:
	go vet ./...