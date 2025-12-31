package models

// Pedido representa uma compra feita pelo cliente.
// Cada pedido possui um identificador único e uma lista de produtos.

type Pedido struct {
	PedidoID int       `json:"pedido_id"`
	Produtos []Produto `json:"produtos"`
}
