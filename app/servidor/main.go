package main

import (
	"log" // Para logs no terminal

	"sistema-empacotamento/internal/controllers" // Controllers da aplicação

	"github.com/gin-gonic/gin" // Framework Gin
)

func main() {
	// Cria o servidor Gin com middlewares padrão (logger e recovery)
	r := gin.Default()

	// Desabilita proxies confiáveis (boa prática em ambiente local)
	r.SetTrustedProxies(nil)

	// Rota de health check
	// Usada para verificar se a API está no ar
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	// Rota principal do sistema
	// Recebe pedidos para empacotamento
	r.POST("/pack", controllers.PackPedidos)

	// Log informando que o servidor iniciou
	log.Println("Servidor iniciado na porta 8081")

	// Inicia o servidor HTTP
	r.Run(":8081")
}
