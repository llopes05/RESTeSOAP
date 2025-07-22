package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func setupRoutes(r *gin.Engine, db *gorm.DB) {

	r.GET("/livros", func(c *gin.Context) { listarLivros(c, db) })
	r.POST("/livros", func(c *gin.Context) { criarLivro(c, db) })


	r.GET("/emprestimos", func(c *gin.Context) { listarEmprestimos(c, db) })
	r.POST("/emprestimos", func(c *gin.Context) { criarEmprestimo(c, db) })
}