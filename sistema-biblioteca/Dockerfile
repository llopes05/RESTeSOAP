FROM golang:1.23.1 as builder

WORKDIR /app
COPY sistema-biblioteca/ ./sistema-biblioteca/
WORKDIR /app/sistema-biblioteca
RUN go mod tidy
RUN go build -o /app/biblioteca-server ./cmd/rest
RUN go build -o /app/biblioteca-grpc ./cmd/grpc/server.go

FROM golang:1.23.1
WORKDIR /app
COPY --from=builder /app/biblioteca-server /app/biblioteca-server
COPY --from=builder /app/biblioteca-grpc /app/biblioteca-grpc
COPY --from=builder /app/sistema-biblioteca/biblioteca.db /app/biblioteca.db
COPY --from=builder /app/sistema-biblioteca/docs /app/docs

RUN apt-get update && apt-get install -y sqlite3 && rm -rf /var/lib/apt/lists/*

EXPOSE 8080 50051

CMD ["/app/biblioteca-server"]