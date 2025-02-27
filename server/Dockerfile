FROM golang:alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

CMD ["go", "run", "./cmd/server/main.go"]