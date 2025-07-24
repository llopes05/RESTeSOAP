package main

import (
	"context"
	"log"
	"time"

	pb "github.com/llopes05/RESTeSOAP/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Falha ao conectar: %v", err)
	}
	defer conn.Close()

	c := pb.NewBibliotecaServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ListarLivros(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Erro ao listar livros: %v", err)
	}
	log.Println("Livros:", r.Livros)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r2, err := c.CriarLivro(ctx, &pb.CriarLivroRequest{
		Titulo:    "Naruto",
		Autor:     "Masashi Kishimoto",
		Ano:       1999,
		Disponivel: true,
	})
	if err != nil {
		log.Fatalf("Erro ao criar livro: %v", err)
	}
	log.Println("Livro criado:", r2.Livro)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	emprestimoResp, err := c.CriarEmprestimo(ctx, &pb.CriarEmprestimoRequest{
		LivroId:    2,
		UsuarioId:  1,
		DataInicio: "2025-07-22",
		DataFim:    "2025-07-29",
	})
	if err != nil {
		log.Fatalf("Erro ao criar empréstimo: %v", err)
	}
	log.Println("Empréstimo criado:", emprestimoResp.Emprestimo)

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = c.DeletarLivro(ctx, &pb.DeletarLivroRequest{Id: 1})
	if err != nil {
		log.Fatalf("Erro ao deletar livro: %v", err)
	}
	log.Println("Livro deletado com sucesso!")
}
