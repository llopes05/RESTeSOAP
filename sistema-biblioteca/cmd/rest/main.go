package main

import (
	"github.com/gin-gonic/gin"
	"github.com/llopes05/RESTeSOAP/internal/database"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/files"
	_ "github.com/llopes05/RESTeSOAP/docs" 
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	setupRoutes(r, db)

	r.Run(":8080")
}