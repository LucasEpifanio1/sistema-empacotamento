package models

// Dimensoes representa as medidas físicas de um produto.
// Essas informações são essenciais para calcular se o produto
// cabe ou não dentro de uma determinada caixa.

type Dimensoes struct {
	Altura      float64 `json:"altura"`
	Largura     float64 `json:"largura"`
	Comprimento float64 `json:"comprimento"`
}

// Produto representa um item do pedido.
// O campo ProdutoID é usado como identificador e também como nome,
// conforme definido no contrato do teste técnico.

type Produto struct {
	ProdutoID string    `json:"produto_id"`
	Dimensoes Dimensoes `json:"dimensoes"`
}
