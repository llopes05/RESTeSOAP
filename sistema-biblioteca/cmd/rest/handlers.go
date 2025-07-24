package main

import (
	"net/http"
	"strconv"

	"github.com/llopes05/RESTeSOAP/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ErrorResponse representa uma resposta de erro para Swagger
// swagger:model
type ErrorResponse struct {
	Error string `json:"error"`
}

// LivroSwagger representa um livro para documentação Swagger
// swagger:model
type LivroSwagger struct {
	ID        uint   `json:"id"`
	Titulo    string `json:"titulo"`
	Autor     string `json:"autor"`
	Ano       int    `json:"ano"`
	Emprestimos []models.Emprestimo `json:"emprestimos,omitempty"` // Campo opcional, apenas para documentação
}

// EmprestimoSwagger representa um empréstimo para documentação Swagger
// swagger:model
type EmprestimoSwagger struct {
	ID        uint   `json:"id"`
	LivroID   uint   `json:"livro_id"`
	UsuarioID uint   `json:"usuario_id"`
	DataInicio string `json:"data_inicio"`
	DataFim    string `json:"data_fim"`
}

// UsuarioSwagger representa um usuário para documentação Swagger
// swagger:model
// Utilize este tipo nos comentários Swagger para evitar erro com gorm.Model
// Adapte os campos conforme o modelo real
//
type UsuarioSwagger struct {
	ID    uint   `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

// @title Biblioteca API
// @version 1.0
// @description API para gerenciar livros e empréstimos de uma biblioteca
// @host localhost:8080
// @BasePath /

func listarLivros(c *gin.Context, db *gorm.DB) {
	// @Summary Listar todos os livros
	// @Description Retorna uma lista de todos os livros disponíveis
	// @Tags livros
	// @Produce json
	// @Success 200 {array} LivroSwagger
	// @Router /livros [get]
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
// @Param livro body LivroSwagger true "Dados do livro"
// @Success 201 {object} LivroSwagger
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

// @Summary Obter um livro por ID
// @Description Retorna os detalhes de um livro específico
// @Tags livros
// @Produce json
// @Param id path int true "ID do livro"
// @Success 200 {object} LivroSwagger
// @Failure 404 {object} ErrorResponse "Livro não encontrado"
// @Router /livros/{id} [get]
func obterLivro(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var livro models.Livro
	if err := db.Preload("Emprestimos").First(&livro, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Livro não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livro"})
		return
	}
	c.JSON(http.StatusOK, livro)
}

// @Summary Atualizar um livro
// @Description Atualiza os dados de um livro existente
// @Tags livros
// @Accept json
// @Produce json
// @Param id path int true "ID do livro"
// @Param livro body LivroSwagger true "Dados atualizados do livro"
// @Success 200 {object} LivroSwagger
// @Failure 404 {object} ErrorResponse "Livro não encontrado"
// @Router /livros/{id} [put]
func atualizarLivro(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var livro models.Livro
	if err := db.First(&livro, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Livro não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livro"})
		return
	}
	if err := c.ShouldBindJSON(&livro); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := db.Save(&livro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar livro"})
		return
	}
	c.JSON(http.StatusOK, livro)
}

// @Summary Deletar um livro
// @Description Remove um livro da biblioteca
// @Tags livros
// @Param id path int true "ID do livro"
// @Success 204 "Nenhum conteúdo"
// @Failure 404 {object} ErrorResponse "Livro não encontrado"
// @Router /livros/{id} [delete]
func deletarLivro(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var livro models.Livro
	if err := db.First(&livro, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Livro não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livro"})
		return
	}
	if err := db.Delete(&livro).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar livro"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// @Summary Listar todos os empréstimos
// @Description Retorna uma lista de todos os empréstimos registrados
// @Tags emprestimos
// @Produce json
// @Success 200 {array} EmprestimoSwagger
// @Router /emprestimos [get]
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
// @Failure 400 {object} gin.H "Dados inválidos ou livro indisponível"
// @Router /emprestimos [post]
func criarEmprestimo(c *gin.Context, db *gorm.DB) {
	var emprestimo models.Emprestimo
	if err := c.ShouldBindJSON(&emprestimo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// livro precisa existir e estar disponível
	var livro models.Livro
	if err := db.First(&livro, emprestimo.LivroID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Livro não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar livro"})
		return
	}
	if !livro.Disponivel {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Livro não está disponível"})
		return
	}

	// se empréstimo criado, livro: indisponível
	emprestimo.Devolvido = false 
	if err := db.Create(&emprestimo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar empréstimo"})
		return
	}
	livro.Disponivel = false
	db.Save(&livro)
	c.JSON(http.StatusCreated, emprestimo)
}

// @Summary Obter um empréstimo por ID
// @Description Retorna os detalhes de um empréstimo específico
// @Tags emprestimos
// @Produce json
// @Param id path int true "ID do empréstimo"
// @Success 200 {object} EmprestimoSwagger
// @Failure 404 {object} ErrorResponse "Empréstimo não encontrado"
// @Router /emprestimos/{id} [get]
func obterEmprestimo(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var emprestimo models.Emprestimo
	if err := db.Preload("Livro").Preload("Usuario").First(&emprestimo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Empréstimo não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar empréstimo"})
		return
	}
	c.JSON(http.StatusOK, emprestimo)
}

// @Summary Atualizar um empréstimo
// @Description Atualiza os dados de um empréstimo existente
// @Tags emprestimos
// @Accept json
// @Produce json
// @Param id path int true "ID do empréstimo"
// @Param emprestimo body EmprestimoSwagger true "Dados atualizados do empréstimo"
// @Success 200 {object} EmprestimoSwagger
// @Failure 404 {object} ErrorResponse "Empréstimo não encontrado"
// @Router /emprestimos/{id} [put]
func atualizarEmprestimo(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var emprestimo models.Emprestimo
	if err := db.First(&emprestimo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Empréstimo não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar empréstimo"})
		return
	}
	if err := c.ShouldBindJSON(&emprestimo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := db.Save(&emprestimo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar empréstimo"})
		return
	}
	c.JSON(http.StatusOK, emprestimo)
}

// @Summary Deletar um empréstimo
// @Description Remove um empréstimo da biblioteca
// @Tags emprestimos
// @Param id path int true "ID do empréstimo"
// @Success 204 "Nenhum conteúdo"
// @Failure 404 {object} ErrorResponse "Empréstimo não encontrado"
// @Router /emprestimos/{id} [delete]
func deletarEmprestimo(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var emprestimo models.Emprestimo
	if err := db.First(&emprestimo, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Empréstimo não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar empréstimo"})
		return
	}
	if err := db.Delete(&emprestimo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar empréstimo"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}

// @Summary Listar todos os usuários
// @Description Retorna uma lista de todos os usuários cadastrados
// @Tags usuarios
// @Produce json
// @Success 200 {array} UsuarioSwagger
// @Router /usuarios [get]
func listarUsuarios(c *gin.Context, db *gorm.DB) {
	var usuarios []models.Usuario
	if err := db.Find(&usuarios).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar usuários"})
		return
	}
	c.JSON(http.StatusOK, usuarios)
}

// @Summary Criar um novo usuário
// @Description Adiciona um novo usuário ao sistema
// @Tags usuarios
// @Accept json
// @Produce json
// @Param usuario body UsuarioSwagger true "Dados do usuário"
// @Success 201 {object} UsuarioSwagger
// @Router /usuarios [post]
func criarUsuario(c *gin.Context, db *gorm.DB) {
	var usuario models.Usuario
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := db.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar usuário"})
		return
	}
	c.JSON(http.StatusCreated, usuario)
}

// @Summary Obter um usuário por ID
// @Description Retorna os detalhes de um usuário específico
// @Tags usuarios
// @Produce json
// @Param id path int true "ID do usuário"
// @Success 200 {object} UsuarioSwagger
// @Failure 404 {object} ErrorResponse "Usuário não encontrado"
// @Router /usuarios/{id} [get]
func obterUsuario(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var usuario models.Usuario
	if err := db.First(&usuario, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

// @Summary Atualizar um usuário
// @Description Atualiza os dados de um usuário existente
// @Tags usuarios
// @Accept json
// @Produce json
// @Param id path int true "ID do usuário"
// @Param usuario body UsuarioSwagger true "Dados atualizados do usuário"
// @Success 200 {object} UsuarioSwagger
// @Failure 404 {object} ErrorResponse "Usuário não encontrado"
// @Router /usuarios/{id} [put]
func atualizarUsuario(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var usuario models.Usuario
	if err := db.First(&usuario, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
		return
	}
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}
	if err := db.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}
	c.JSON(http.StatusOK, usuario)
}

// @Summary Deletar um usuário
// @Description Remove um usuário do sistema
// @Tags usuarios
// @Param id path int true "ID do usuário"
// @Success 204 "Nenhum conteúdo"
// @Failure 404 {object} ErrorResponse "Usuário não encontrado"
// @Router /usuarios/{id} [delete]
func deletarUsuario(c *gin.Context, db *gorm.DB) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
	var usuario models.Usuario
	if err := db.First(&usuario, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, ErrorResponse{"Usuário não encontrado"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar usuário"})
		return
	}
	if err := db.Delete(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar usuário"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}