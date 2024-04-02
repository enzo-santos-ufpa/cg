package ufpa_cg

type AlgoritmoLinha interface {
	Move() bool
	PontoAtual() Ponto
}
