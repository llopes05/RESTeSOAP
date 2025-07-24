package main

import (
	"fmt"
	"github.com/llopes05/RESTeSOAP/internal/database"
	"github.com/llopes05/RESTeSOAP/internal/models"
)

func main() {
	// Conecta ao banco de dados, como já fazia antes
	db, err := database.InitDB()
	if err != nil {
		panic("Falha ao conectar ao banco: " + err.Error())
	}
	fmt.Println("Banco conectado com sucesso.")

	// Apaga os dados antigos para evitar duplicatas
	fmt.Println("Limpando dados antigos...")
	db.Exec("DELETE FROM emprestimos")
	db.Exec("DELETE FROM livros")
	db.Exec("DELETE FROM usuarios")

	// Roda as migrações para garantir que as tabelas existam
	db.AutoMigrate(&models.Livro{}, &models.Usuario{}, &models.Emprestimo{})
	fmt.Println("Tabelas migradas.")

	// --- INÍCIO: POPULANDO O BANCO ---

	// Criando Usuários
	usuarios := []models.Usuario{
		{Nome: "Alice", Email: "alice@example.com"},
		{Nome: "Beto", Email: "beto@example.com"},
	}
	db.Create(&usuarios)
	fmt.Printf("%d usuários criados.\n", len(usuarios))

	// Criando Livros
	livros := []models.Livro{
		{Titulo: "1984", Autor: "George Orwell", Ano: 1949, Disponivel: true},
		{Titulo: "Admirável Mundo Novo", Autor: "Aldous Huxley", Ano: 1932, Disponivel: true},
		{Titulo: "O Senhor dos Anéis", Autor: "J.R.R. Tolkien", Ano: 1954, Disponivel: true},
		{Titulo: "A Metamorfose", Autor: "Franz Kafka", Ano: 1915, Disponivel: false}, // Um livro já emprestado
	}
	db.Create(&livros)
	fmt.Printf("%d livros criados.\n", len(livros))

	// Criando um Empréstimo de exemplo para o livro indisponível
	emprestimoExemplo := models.Emprestimo{
		LivroID:   4, // ID do livro "A Metamorfose"
		UsuarioID: 1, // ID da "Alice"
		Devolvido: false,
	}
	db.Create(&emprestimoExemplo)
	fmt.Println("1 empréstimo de exemplo criado.")

	fmt.Println("\nBanco populado com sucesso!")
}