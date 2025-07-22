package main

import (
	"net/http"
	"github.com/llopes05/RESTeSOAP/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @title Biblioteca API
// @version 1.0
// @description API para gerenciar livros e empréstimos de uma biblioteca
// @host localhost:8080
// @BasePath /

func listarLivros(c *gin.Context, db *gorm.DB) {
	var livros []models.Livro
	if err := db.Preload("Emprestimos").Find(&livros).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar livros"})
		return
	}
	c.JSON(http.StatusOK, livros)
}

// @Summary Criar um novo livro
// @Description Adiciona um novo livro à biblioteca
// @Tags livros
// @Accept json
// @Produce json
// @Param livro body models.Livro true "Dados do livro"
// @Success 201 {object} models.Livro
// @Router /livros [post]
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

// @Summary Criar um novo empréstimo
// @Description Registra um novo empréstimo de livro
// @Tags emprestimos
// @Accept json
// @Produce json
// @Param emprestimo body models.Emprestimo true "Dados do empréstimo"
// @Success 201 {object} models.Emprestimo
// @Router /emprestimos [post]
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