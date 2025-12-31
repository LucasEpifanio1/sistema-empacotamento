package services

import "sistema-empacotamento/internal/models"

// Estrutura simples para representar uma caixa em uso
type CaixaEmUso struct {
	Caixa    models.Caixa
	Produtos []models.Produto
}

// Caixas disponíveis no sistema
var caixasDisponiveis = []models.Caixa{
	{ID: "Caixa 1", Altura: 30, Largura: 40, Comprimento: 80},
	{ID: "Caixa 2", Altura: 50, Largura: 50, Comprimento: 40},
	{ID: "Caixa 3", Altura: 50, Largura: 80, Comprimento: 60},
}

// Verifica se o produto cabe fisicamente dentro da caixa
func produtoCabeNaCaixa(produto models.Produto, caixa models.Caixa) bool {
	return produto.Dimensoes.Altura <= caixa.Altura &&
		produto.Dimensoes.Largura <= caixa.Largura &&
		produto.Dimensoes.Comprimento <= caixa.Comprimento
}

// Função principal de empacotamento
func EmpacotarPedido(pedido models.Pedido) []CaixaEmUso {
	var caixasEmUso []CaixaEmUso

	// Percorre todos os produtos do pedido
	for _, produto := range pedido.Produtos {

		colocado := false

		// Tenta colocar o produto em uma caixa já aberta
		for i := range caixasEmUso {
			if produtoCabeNaCaixa(produto, caixasEmUso[i].Caixa) {
				caixasEmUso[i].Produtos = append(caixasEmUso[i].Produtos, produto)
				colocado = true
				break
			}
		}

		// Se não couber em nenhuma caixa aberta, abre uma nova
		if !colocado {
			for _, caixa := range caixasDisponiveis {
				if produtoCabeNaCaixa(produto, caixa) {
					caixasEmUso = append(caixasEmUso, CaixaEmUso{
						Caixa:    caixa,
						Produtos: []models.Produto{produto},
					})
					break
				}
			}
		}
	}

	return caixasEmUso
}
