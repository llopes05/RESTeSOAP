package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llopes05/RESTeSOAP/internal/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	r := gin.Default()

	setupRoutes(r, db)

	r.Run(":8080")
}