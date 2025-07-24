package soap

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/llopes05/RESTeSOAP/internal/models"
	"gorm.io/gorm"
)


func SoapHandler(c *gin.Context, db *gorm.DB) {
	rawXML, err := io.ReadAll(c.Request.Body)
	if err != nil {
		respondWithError(c, "Corpo da requisição inválido")
		return
	}
	var requestEnvelope EmprestarLivroEnvelope
	if err := xml.Unmarshal(rawXML, &requestEnvelope); err != nil {
		respondWithError(c, "XML mal formatado ou corpo da requisição inesperado")
		return
	}
	requestData := requestEnvelope.Body.Request
	if requestData.XMLName.Local != "EmprestarLivroRequest" {
		respondWithError(c, "Ação SOAP 'EmprestarLivroRequest' não encontrada no Body")
		return
	}

	var livro models.Livro
	if err := db.First(&livro, requestData.LivroID).Error; err != nil {
		respondWithError(c, "Livro não encontrado")
		return
	}

	if !livro.Disponivel {
		respondWithError(c, "Livro não está disponível para empréstimo")
		return
	}

	emprestimo := models.Emprestimo{
		LivroID:    requestData.LivroID,
		UsuarioID:  requestData.UsuarioID,
		DataInicio: time.Now(),
		DataFim:    time.Now().AddDate(0, 0, 14),
		Devolvido:  false,
	}

	if err := db.Create(&emprestimo).Error; err != nil {
		respondWithError(c, "Erro ao salvar o empréstimo no banco de dados")
		return
	}

	livro.Disponivel = false
	db.Save(&livro)

	responseContent := EmprestarLivroResponse{
		XMLName:      xml.Name{Space: "http://biblioteca.com/soap/definitions", Local: "EmprestarLivroResponse"},
		Sucesso:      true,
		Mensagem:     fmt.Sprintf("SUCESSO! Empréstimo do livro '%s' realizado!", livro.Titulo),
		EmprestimoID: emprestimo.ID,
	}
	responseEnvelope := SOAPResponseEnvelope{Body: SOAPResponseBody{Response: responseContent}}

	c.Header("Content-Type", "text/xml; charset=utf-8") // 'text/xml' é mais comum para SOAP 1.1
	c.XML(http.StatusOK, responseEnvelope)
}

func respondWithError(c *gin.Context, message string) {
	fault := SOAPFault{
		FaultCode:   "Client",
		FaultString: message,
	}
	responseEnvelope := SOAPResponseEnvelope{Body: SOAPResponseBody{Response: fault}}
	c.Header("Content-Type", "text/xml; charset=utf-8")
	c.XML(http.StatusInternalServerError, responseEnvelope)
}