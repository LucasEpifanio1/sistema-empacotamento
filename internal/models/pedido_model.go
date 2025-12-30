package models

type Pedido struct {
	ID       string    `json:"id"`
	Produtos []Produto `json:"produtos"`
}
