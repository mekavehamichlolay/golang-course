FROM golang:1.20-alpine3.18

WORKDIR /app
RUN apk update && \
    apk add --no-cache git && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

COPY go.mod .
COPY go.sum .

RUN go mod tidy && go mod vendor

COPY . .

RUN go build -o goapp

CMD ["./goapp"]
