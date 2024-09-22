FROM golang:1.23-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . /app

RUN apk add --no-cache make
RUN go install github.com/a-h/templ/cmd/templ@latest
RUN go install github.com/air-verse/air@latest

RUN templ generate
RUN go build -o ./tmp/main ./src/main.go

EXPOSE 8080
CMD [ "air", "-c", "/app/.air.toml" ]

