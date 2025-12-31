// controllers/pack_controller.go
package controllers

import (
	"net/http"

	"sistema-empacotamento/internal/models"
	"sistema-empacotamento/internal/services"

	"github.com/gin-gonic/gin"
)

// PackRequest representa o body do request para empacotamento
type PackRequest struct {
	Pedidos []models.Pedido `json:"pedidos"`
}

// CaixaResponse representa uma caixa retornada no response
type CaixaResponse struct {
	CaixaID  string   `json:"caixa_id"`
	Produtos []string `json:"produtos"`
}

// PedidoResponse representa cada pedido no response
type PedidoResponse struct {
	PedidoID int             `json:"pedido_id"`
	Caixas   []CaixaResponse `json:"caixas"`
}

// PackResponse representa o response completo do endpoint /pack
type PackResponse struct {
	Pedidos []PedidoResponse `json:"pedidos"`
}

// PackPedidos godoc
// @Summary Empacota produtos em caixas
// @Description Recebe pedidos com produtos e retorna a alocação em caixas disponíveis
// @Tags Empacotamento
// @Accept json
// @Produce json
// @Param pedidos body PackRequest true "Lista de pedidos"
// @Success 200 {object} PackResponse
// @Failure 400 {object} map[string]string
// @Router /pack [post]
func PackPedidos(c *gin.Context) {
	var request PackRequest

	// Converte o JSON recebido em struct
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	var response []PedidoResponse

	// Processa cada pedido individualmente
	for _, pedido := range request.Pedidos {
		caixas := services.EmpacotarPedido(pedido)

		var caixasResponse []CaixaResponse

		// Converte o retorno do service para o formato do JSON
		for _, caixa := range caixas {
			var produtos []string
			for _, produto := range caixa.Produtos {
				produtos = append(produtos, produto.ProdutoID)
			}

			caixasResponse = append(caixasResponse, CaixaResponse{
				CaixaID:  caixa.Caixa.ID,
				Produtos: produtos,
			})
		}

		response = append(response, PedidoResponse{
			PedidoID: pedido.PedidoID,
			Caixas:   caixasResponse,
		})
	}

	c.JSON(http.StatusOK, PackResponse{Pedidos: response})
}
