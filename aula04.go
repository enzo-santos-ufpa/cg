package ufpa_cg

import (
	"math"
)

func (p Ponto) Rotaciona(angulo float64, pivo Ponto) Ponto {
	ponto := p
	ponto = Ponto{X: ponto.X - pivo.X, Y: ponto.Y - pivo.Y}
	ponto = Ponto{
		X: int(math.Round(float64(ponto.X)*math.Cos(angulo) - float64(ponto.Y)*math.Sin(angulo))),
		Y: int(math.Round(float64(ponto.X)*math.Sin(angulo) + float64(ponto.Y)*math.Cos(angulo))),
	}
	ponto = Ponto{X: ponto.X + pivo.X, Y: ponto.Y + pivo.Y}
	return ponto
}
