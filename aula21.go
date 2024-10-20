package ufpa_cg

import (
	"log"
	"math"
)

func sign(x float64) int {
	if x >= 0 {
		return 0
	}
	return 1
}

type ResultadoIntersecao struct {
	EhRetaCompleta   bool
	EhRetaVisivel    bool
	PontoIntersecaoX Ponto
	PontoIntersecaoY Ponto
}

type JanelaRecorte struct {
	PontoInferiorEsquerdo Ponto
	PontoSuperiorDireito  Ponto
}

func (j JanelaRecorte) calculaBits(ponto Ponto) int {
	b := 0
	b |= sign(float64(j.PontoSuperiorDireito.Y)-float64(ponto.Y)) << 3
	b |= sign(float64(ponto.Y)-float64(j.PontoInferiorEsquerdo.Y)) << 2
	b |= sign(float64(j.PontoSuperiorDireito.X)-float64(ponto.X)) << 1
	b |= sign(float64(ponto.X)-float64(j.PontoInferiorEsquerdo.X)) << 0
	return b
}

func (j JanelaRecorte) CalculaIntersecao(p1, p2 Ponto) ResultadoIntersecao {
	b1 := j.calculaBits(p1)
	b2 := j.calculaBits(p2)
	log.Printf("P1(%d, %d) = %04b", p1.X, p1.Y, b1)
	log.Printf("P2(%d, %d) = %04b", p2.X, p2.Y, b2)
	if b1|b2 == 0 {
		// Totalmente dentro
		return ResultadoIntersecao{
			EhRetaCompleta: true,
			EhRetaVisivel:  true,
		}
	}
	if b1&b2 != 0 {
		// Totalmente fora
		return ResultadoIntersecao{
			EhRetaCompleta: true,
			EhRetaVisivel:  false,
		}
	}

	var x int
	var y int
	cr := b1 | b2
	if cr>>0&1 == 1 {
		x = j.PontoInferiorEsquerdo.X
	} else /*if cr>>2&1 == 1*/ {
		x = j.PontoSuperiorDireito.X
	}
	if cr>>2&1 == 1 {
		y = j.PontoInferiorEsquerdo.Y
	} else /*if cr>>4&1 == 1*/ {
		y = j.PontoSuperiorDireito.Y
	}

	xi := float64(y-p1.Y)*float64(p2.X-p1.X)/float64(p2.Y-p1.Y) + float64(p1.X)
	yi := float64(x-p1.X)*float64(p2.Y-p1.Y)/float64(p2.X-p1.X) + float64(p1.Y)
	return ResultadoIntersecao{
		EhRetaCompleta: false,
		EhRetaVisivel:  true,
		PontoIntersecaoX: Ponto{
			X: int(math.Round(xi)),
			Y: y,
		},
		PontoIntersecaoY: Ponto{
			X: x,
			Y: int(math.Round(yi)),
		},
	}

}
