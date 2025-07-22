package main

import (
	"net/http"
	"github.com/llopes05/RESTeSOAP/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func listarLivros(c *gin.Context, db *gorm.DB) {
	var livros []models.Livro
	if err := db.Preload("Emprestimos").Find(&livros).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar livros"})
		return
	}
	c.JSON(http.StatusOK, livros)
}

func criarLivro(c *gin.Context, db *gorm.DB) {
	var livro models.Livro
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := db.Create(&livro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar livro"})
		return
	}
	c.JSON(http.StatusCreated, livro)
}

func listarEmprestimos(c *gin.Context, db *gorm.DB) {
	var emprestimos []models.Emprestimo
	if err := db.Preload("Livro").Preload("Usuario").Find(&emprestimos).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar empréstimos"})
		return
	}
	c.JSON(http.StatusOK, emprestimos)
}

func criarEmprestimo(c *gin.Context, db *gorm.DB) {
	var emprestimo models.Emprestimo
	if err := c.ShouldBindJSON(&emprestimo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := db.Create(&emprestimo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar empréstimo"})
		return
	}
	c.JSON(http.StatusCreated, emprestimo)
}