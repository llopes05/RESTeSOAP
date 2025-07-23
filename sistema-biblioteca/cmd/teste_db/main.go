package main

import (
	"fmt"
	"github.com/llopes05/RESTeSOAP/internal/database"
	"github.com/llopes05/RESTeSOAP/internal/models"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("Falha ao conectar ao banco: " + err.Error())
	}
	fmt.Println("Banco conectado")

	db.AutoMigrate(&models.Livro{}, &models.Usuario{}, &models.Emprestimo{})
	fmt.Println("Tabelas 'livros', 'usuarios' e 'emprestimos' criadas.")
}