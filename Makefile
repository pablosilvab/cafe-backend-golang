APP_NAME=cafe-backend-golang
APP_VERSION=0.0.1.SNAPSHOT
PROJECT_FOLDER=.
GIT_DIR=$(shell pwd)

docker-shell:
	docker run --rm -it -v $(GIT_DIR):/app -w /app/$(PROJECT_FOLDER) --entrypoint=/bin/sh ${APP_NAME}:${APP_VERSION}

docker-build:
	docker build -t ${APP_NAME}:${APP_VERSION} ${GIT_DIR}

go-run-build:
	./build/${APP_NAME}

go-build:
	go build -o ./build/${APP_NAME} cmd/main.go

go-run:
	go run cmd/main.go