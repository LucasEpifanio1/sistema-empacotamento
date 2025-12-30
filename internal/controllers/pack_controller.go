package controllers

// Importações necessárias
import (
	"net/http" // Códigos HTTP (200, 400, etc)

	"sistema-empacotamento/internal/models" // Models do domínio

	"github.com/gin-gonic/gin" // Framework Gin
)

// PackRequest representa o formato do JSON esperado na requisição
// Exemplo de JSON:
//
//	{
//	  "pedidos": [ { ... } ]
//	}
type PackRequest struct {
	Pedidos []models.Pedido `json:"pedidos"`
}

// PackPedidos é o handler da rota POST /pack
func PackPedidos(c *gin.Context) {
	// Variável que irá armazenar os dados vindos do JSON
	var req PackRequest

	// Tenta converter (bindar) o JSON da requisição para a struct PackRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		// Se o JSON estiver inválido ou fora do padrão esperado
		c.JSON(http.StatusBadRequest, gin.H{
			"erro": "JSON inválido",
		})
		return
	}

	// Por enquanto, não existe lógica de empacotamento
	// Apenas devolvemos os pedidos recebidos
	c.JSON(http.StatusOK, gin.H{
		"mensagem": "Pedidos recebidos com sucesso",
		"pedidos":  req.Pedidos,
	})
}
