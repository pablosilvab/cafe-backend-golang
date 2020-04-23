
# build stage
FROM golang:latest as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o run cmd/main.go

# final stage
FROM alpine:latest 
COPY --from=builder /app/run /app/
COPY config/ "/app/config"
WORKDIR /app
EXPOSE 8080
ENTRYPOINT ["./run"]