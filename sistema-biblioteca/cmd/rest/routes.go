package main

import "github.com/gin-gonic/gin"

func setupRoutes(r *gin.Engine) {
	r.GET("/livros", listarLivros)
	r.POST("/livros", criarLivro)
	r.GET("/livros/:id", buscarLivro)
	r.PUT("/livros/:id", atualizarLivro)
	r.DELETE("/livros/:id", deletarLivro)
}

