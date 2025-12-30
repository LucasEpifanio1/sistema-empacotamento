package models

type Produto struct {
	Nome        string  `json:"nome"`
	Altura      float64 `json:"altura"`
	Largura     float64 `json:"largura"`
	Comprimento float64 `json:"comprimento"`
}
