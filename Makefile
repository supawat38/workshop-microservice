 go get  github.com/googollee/go-socket.io@v1.0.1
 
 .PHONY: clean test security build run

APP_NAME = apiserver
BUILD_DIR = $(PWD)/build

clean:
	rm -rf ./build

security:
	gosec -quiet ./...

test: security
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o $(BUILD_DIR)/$(APP_NAME) main.go

docker.run: docker.network docker.postgres swag docker.fiber migrate.up

docker.network:
	docker network inspect dev-network >/dev/null 2>&1 || \
	docker network create -d bridge dev-network

docker.fiber.build:
	docker build -t fiber .

docker.fiber: docker.fiber.build
	docker run --rm -d \
		--name dev-fiber \
		--network dev-network \
		-p 5000:5000 \
		fiber

docker.stop: docker.stop.fiber

docker.stop.fiber:
	docker stop dev-fiber