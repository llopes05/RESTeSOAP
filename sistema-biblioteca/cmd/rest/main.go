package main

import (
   "github.com/gin-gonic/gin"
   "github.com/llopes05/RESTeSOAP/internal/database"
   "github.com/swaggo/gin-swagger"
   "github.com/swaggo/files"
   _ "github.com/llopes05/RESTeSOAP/docs"
   "github.com/llopes05/RESTeSOAP/internal/soap"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		panic("Falha ao conectar ao banco de dados: " + err.Error())
	}

   r := gin.Default()

   r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

   setupRoutes(r, db)

   // Rota SOAP (POST)
   r.POST("/soap", func(c *gin.Context) {
	   soap.SoapHandler(c, db)
   })

   // Rota para servir o WSDL (GET)
   r.GET("/soap/biblioteca.wsdl", func(c *gin.Context) {
	   c.File("docs/biblioteca.wsdl")
   })

   r.Run(":8080")
}