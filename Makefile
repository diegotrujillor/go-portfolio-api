APP_NAME=api
PORT=8080
IMAGE=go-portfolio-api:local

.PHONY: run test tidy build clean docker-build docker-run

tidy:
	go mod tidy

run:
	go run ./cmd/api

test:
	go test ./... -count=1

build:
	go build -o bin/$(APP_NAME) ./cmd/api

clean:
	rm -rf bin

docker-build:
	docker build -t $(IMAGE) .

docker-run:
	docker run --rm -p $(PORT):8080 $(IMAGE)