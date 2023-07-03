FROM golang:1.20-alpine3.18

WORKDIR /app
RUN apk update && \
    apk add --no-cache git && \
    go install github.com/kyleconroy/sqlc/cmd/sqlc@latest && \
    go install github.com/pressly/goose/v3/cmd/goose@latest

COPY ./vendor .

COPY go.mod .
COPY go.sum .

RUN go mod tidy && go mod vendor

COPY . .

RUN go build -o goapp
ENV CGO_ENABLED=1

CMD ["./goapp"]
