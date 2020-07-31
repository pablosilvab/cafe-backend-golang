APP_NAME=cafe-backend-golang
APP_VERSION=0.0.1.SNAPSHOT
APP_DOCKER_HOST = pablon27
PROJECT_FOLDER=.
GIT_DIR=$(shell pwd)
HOST_PORT=8080

helm-install:
	helm upgrade -i ${APP_NAME} ./charts/

helm-replace:
	helm delete ${APP_NAME} && helm install ${APP_NAME} ./charts/

docker-shell:
	docker run --rm -it -v $(GIT_DIR):/app -w /app/$(PROJECT_FOLDER) --entrypoint=/bin/sh ${APP_NAME}:${APP_VERSION}

docker-run:
	docker run --rm -p ${HOST_PORT}:8080 ${APP_NAME}:${APP_VERSION}

docker-build:
	docker build -t ${APP_NAME}:${APP_VERSION} ${GIT_DIR}

docker-push:
	docker build -t ${APP_DOCKER_HOST}/${APP_NAME}:${APP_VERSION} ${GIT_DIR}
	docker push ${APP_DOCKER_HOST}/${APP_NAME}:${APP_VERSION} 

go-run-build:
	./build/${APP_NAME}

go-build:
	go build -o ./build/${APP_NAME} cmd/main.go

go-run:
	go run cmd/main.go

go-download:
	export GO111MODULE=on
	go mod download
