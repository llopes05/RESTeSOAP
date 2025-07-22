package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/livros", func(c *gin.Context) { listarLivros(c, db) })
	r.POST("/livros", func(c *gin.Context) { criarLivro(c, db) })
	r.GET("/livros/:id", func(c *gin.Context) { obterLivro(c, db) })
	r.PUT("/livros/:id", func(c *gin.Context) { atualizarLivro(c, db) })
	r.DELETE("/livros/:id", func(c *gin.Context) { deletarLivro(c, db) })

	r.GET("/emprestimos", func(c *gin.Context) { listarEmprestimos(c, db) })
	r.POST("/emprestimos", func(c *gin.Context) { criarEmprestimo(c, db) })
	r.GET("/emprestimos/:id", func(c *gin.Context) { obterEmprestimo(c, db) })
	r.PUT("/emprestimos/:id", func(c *gin.Context) { atualizarEmprestimo(c, db) })
	r.DELETE("/emprestimos/:id", func(c *gin.Context) { deletarEmprestimo(c, db) })
}