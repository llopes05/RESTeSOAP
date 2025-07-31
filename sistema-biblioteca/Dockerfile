FROM golang:1.23.1 AS builder

WORKDIR /app
COPY . ./
RUN go mod tidy
RUN go build -o /app/biblioteca-server ./cmd/rest
RUN go build -o /app/biblioteca-grpc ./cmd/grpc/server.go

FROM golang:1.23.1
WORKDIR /app
COPY --from=builder /app/biblioteca-server /app/biblioteca-server
COPY --from=builder /app/biblioteca-grpc /app/biblioteca-grpc
COPY --from=builder /app/biblioteca.db /app/biblioteca.db
COPY --from=builder /app/docs /app/docs

RUN apt-get update && apt-get install -y sqlite3 && rm -rf /var/lib/apt/lists/*

EXPOSE 8080 50051

CMD ["/app/biblioteca-server"]