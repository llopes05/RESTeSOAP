# Sistema de Biblioteca

Um sistema simples para gerenciar livros e empréstimos, com API REST e suporte a gRPC.

## Descrição
Este projeto oferece uma API REST para CRUD de livros e empréstimos, além de um serviço gRPC para listar e criar livros. Utiliza Go, GORM (SQLite) e RabbitMQ (opcional para mensageria).

## Dependências
- **Go**: Versão 1.23.1
- **SQLite**: Para o banco de dados.
- **Protobuf Compiler (protoc)**: Para gerar código gRPC.
- **RabbitMQ**: Opcional, para mensageria assíncrona.
- **Pacotes Go:**
  - `go get github.com/gin-gonic/gin`
  - `go get gorm.io/driver/sqlite`
  - `go get gorm.io/gorm`
  - `go get github.com/swaggo/swag`
  - `go get github.com/swaggo/gin-swagger`
  - `go get github.com/swaggo/files`
  - `go get google.golang.org/grpc`
  - `go get google.golang.org/protobuf/cmd/protoc-gen-go`
  - `go get google.golang.org/grpc/cmd/protoc-gen-go-grpc`
  - `go get github.com/streadway/amqp`

## Instalação das dependências
```bash
go mod tidy
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

## Instale o protoc
```bash
sudo apt install protobuf-compiler
```

## Gere os arquivos gRPC
```bash
protoc --go_out=./proto --go-grpc_out=./proto --proto_path=. proto/biblioteca.proto
```

# Como Rodar

## API REST
Inicie o servidor REST na porta 8080:
```bash
go run cmd/rest/*.go
```
Acesse a documentação Swagger em: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

## Servidor gRPC
Inicie o servidor gRPC na porta 50051:
```bash
go run cmd/grpc/server.go
```

## Cliente gRPC
Teste o cliente gRPC:
```bash
go run cmd/grpc/client.go
```

## Docker

Para rodar o backend (REST/gRPC) em container Docker:

### Build da imagem
```bash
docker build -t biblioteca-app .
```

### Rodar o servidor REST (porta 8080)
```bash
docker run -p 8080:8080 biblioteca-app
```

### Rodar o servidor gRPC (porta 50051)
```bash
docker run -p 50051:50051 biblioteca-app /app/biblioteca-grpc
```

> Obs: Certifique-se de que as portas desejadas estejam livres no seu sistema antes de rodar o container.