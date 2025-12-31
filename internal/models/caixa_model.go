package models

// Caixa representa uma caixa física disponível para empacotamento.
// Esse model não possui tags JSON porque ele não vem do request;
// é uma estrutura interna usada apenas na lógica do sistema.

type Caixa struct {
	ID          string
	Altura      float64
	Largura     float64
	Comprimento float64
}
