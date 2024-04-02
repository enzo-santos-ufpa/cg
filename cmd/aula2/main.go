package main

import (
	"fmt"
	"ufpa_cg"
)

func main() {
	p1 := ufpa_cg.Ponto{X: 0, Y: 0}
	p2 := ufpa_cg.Ponto{X: 5, Y: 2}
	algoritmo := ufpa_cg.NewAlgoritmoBresenham(p1, p2)
	for algoritmo.Move() {
		ponto := algoritmo.PontoAtual()
		fmt.Printf("(%d, %d)\n", ponto.X, ponto.Y)
	}
}
