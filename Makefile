APP_NAME=cafe-backend-golang

go-run-build:
	./build/${APP_NAME}

go-build:
	go build -o ./build/${APP_NAME} cmd/main.go

go-run:
	go run cmd/main.go