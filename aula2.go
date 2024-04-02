package ufpa_cg

import (
	"math"
)

type AlgoritmoBruto struct {
	PontoInicial Ponto
	PontoFinal   Ponto

	m float64

	x int
}

func NewAlgoritmoBruto(p1, p2 Ponto) *AlgoritmoBruto {
	return &AlgoritmoBruto{
		PontoInicial: p1,
		PontoFinal:   p2,

		x: p1.X,
		m: float64(p2.Y-p1.Y) / float64(p2.X-p1.X),
	}
}

func (a *AlgoritmoBruto) Move() bool {
	if a.x > a.PontoFinal.X {
		return false
	}
	a.x++
	return true
}

func (a *AlgoritmoBruto) PontoAtual() Ponto {
	// Atua sobre o `x` anterior, visto que `Move` sempre aponta o `a.x` para a próxima coordenada X a ser gerada
	x := a.x - 1

	// Considerando que m = (yb-ya)/(xb-xa), então
	//     (yb-ya)/(xb-xa) = m
	//     yb - ya = m*(xb-xa)
	//     yb = ya + m*(xb-xa)
	//
	// Definindo (xa, ya) como a.PontoInicial, portanto
	//     yb = y0 + m*(xb-x0)
	return Ponto{
		X: x,
		Y: a.PontoInicial.Y + int(math.Round(a.m*float64(x-a.PontoInicial.X))),
	}
}

type AlgoritmoBresenham struct {
	PontoInicial Ponto
	PontoFinal   Ponto

	m float64

	x int
	y int
	e float64

	pontoAtual Ponto
}

func NewAlgoritmoBresenham(p1, p2 Ponto) *AlgoritmoBresenham {
	m := float64(p2.Y-p1.Y) / float64(p2.X-p1.X)
	return &AlgoritmoBresenham{
		PontoInicial: p1,
		PontoFinal:   p2,

		m: m,

		pontoAtual: p1,
		x:          p1.X,
		y:          p1.Y,
		e:          m - .5,
	}
}

func (a *AlgoritmoBresenham) Move() bool {
	if a.x > a.PontoFinal.X {
		return false
	}
	a.pontoAtual = Ponto{X: a.x, Y: a.y}

	if a.e >= 0 {
		a.y++
		a.e--
	}
	a.x++
	a.e += a.m
	return true
}

func (a *AlgoritmoBresenham) PontoAtual() Ponto {
	return a.pontoAtual
}
