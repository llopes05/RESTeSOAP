package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/protobuf/types/known/emptypb"
	"github.com/llopes05/RESTeSOAP/internal/database"
	pb "github.com/llopes05/RESTeSOAP/proto" 
	"google.golang.org/grpc"
	"gorm.io/gorm"
)

type server struct {
	pb.UnimplementedBibliotecaServiceServer
	db *gorm.DB
}

func (s *server) ListarLivros(ctx context.Context, req *emptypb.Empty) (*pb.ListaLivrosResponse, error) {
	var livros []struct {
		ID        uint   `gorm:"column:id"`
		Titulo    string `gorm:"column:titulo"`
		Autor     string `gorm:"column:autor"`
		Ano       int    `gorm:"column:ano"`
		Disponivel bool   `gorm:"column:disponivel"`
	}
	if err := s.db.Table("livros").Select("id, titulo, autor, ano, disponivel").Find(&livros).Error; err != nil {
		return nil, err
	}

	response := &pb.ListaLivrosResponse{}
	for _, l := range livros {
		response.Livros = append(response.Livros, &pb.Livro{
			Id:        int32(l.ID),
			Titulo:    l.Titulo,
			Autor:     l.Autor,
			Ano:       int32(l.Ano),
			Disponivel: l.Disponivel,
		})
	}
	return response, nil
}

func (s *server) CriarLivro(ctx context.Context, req *pb.CriarLivroRequest) (*pb.CriarLivroResponse, error) {
	livro := &pb.Livro{
		Titulo:    req.Titulo,
		Autor:     req.Autor,
		Ano:       req.Ano,
		Disponivel: req.Disponivel,
	}
	if err := s.db.Exec("INSERT INTO livros (titulo, autor, ano, disponivel) VALUES (?, ?, ?, ?)",
		livro.Titulo, livro.Autor, livro.Ano, livro.Disponivel).Error; err != nil {
		return nil, err
	}
	var newID int
	s.db.Raw("SELECT last_insert_rowid()").Scan(&newID)
	livro.Id = int32(newID)

	return &pb.CriarLivroResponse{Livro: livro}, nil
}

func (s *server) CriarEmprestimo(ctx context.Context, req *pb.CriarEmprestimoRequest) (*pb.CriarEmprestimoResponse, error) {
    emprestimo := &pb.Emprestimo{
        LivroId:    req.LivroId,
        UsuarioId:  req.UsuarioId,
        DataInicio: req.DataInicio,
        DataFim:    req.DataFim,
        Devolvido:  false,
    }

    if err := s.db.Create(emprestimo).Error; err != nil {
        return nil, err
    }
    var newID int
    s.db.Raw("SELECT last_insert_rowid()").Scan(&newID)
    emprestimo.Id = int32(newID)
    return &pb.CriarEmprestimoResponse{Emprestimo: emprestimo}, nil
}

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao ouvir: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBibliotecaServiceServer(s, &server{db: db})
	log.Printf("Servidor gRPC rodando na porta :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir: %v", err)
	}
}