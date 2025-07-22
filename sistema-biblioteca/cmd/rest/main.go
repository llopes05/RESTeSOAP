package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llopes05/RESTeSOAP/internal/database"
)

func main() {
	if _, err := database.InitDB(); err != nil {
		panic("Falha ao conectar ao banco: " + err.Error())
	}

	r := gin.Default()
	configurarRotas(r)
	r.Run(":8080") 
}

func configurarRotas(r *gin.Engine) {
	r.GET("/livros", listarLivros)
	r.POST("/livros", criarLivro)
	r.GET("/livros/:id", buscarLivro)
	r.PUT("/livros/:id", atualizarLivro)
	r.DELETE("/livros/:id", deletarLivro)
}