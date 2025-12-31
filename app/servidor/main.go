// main.go
package main

import (
	"log"
	"sistema-empacotamento/internal/controllers"

	_ "sistema-empacotamento/docs" // <- necessário para registrar os docs

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title API de Empacotamento
// @version 1.0
// @description API para empacotamento de produtos em caixas
// @host localhost:8081
// @BasePath /
func main() {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// Endpoint principal
	r.POST("/pack", controllers.PackPedidos)

	// Rota Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Println("Servidor iniciado na porta 8081")
	r.Run(":8081")
}
