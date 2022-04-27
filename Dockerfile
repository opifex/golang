FROM golang:1.18.1-alpine

WORKDIR /app

COPY . .

RUN go mod download

EXPOSE 80

ENTRYPOINT [ "go", "run", "cmd/app/main.go" ]
