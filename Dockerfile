FROM golang:1.24-alpine

WORKDIR /app

COPY src/go.mod src/go.sum ./

RUN go mod download

COPY src/ ./

RUN go build -o server cmd/main.go

EXPOSE 5000

CMD ["./server"]