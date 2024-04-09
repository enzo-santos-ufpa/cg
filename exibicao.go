package ufpa_cg

import (
	"fmt"
	"io"
	"slices"
)

func Exibe(algoritmo AlgoritmoLinha, w io.Writer) error {
	return ExibePoligono([]AlgoritmoLinha{algoritmo}, w)
}

func ExibePoligono(algoritmos []AlgoritmoLinha, w io.Writer) error {
	pontos := make([]Ponto, 0)
	vertices := make([]Ponto, 0)
	for _, algoritmo := range algoritmos {
		for algoritmo.Move() {
			pontos = append(pontos, algoritmo.PontoAtual())
		}
		p1 := pontos[0]
		p2 := pontos[len(pontos)-1]
		vertices = append(vertices, p1, p2)
	}

	chars := []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}

	var minX, maxX, minY, maxY int
	var minXOk, maxXOk, minYOk, maxYOk bool
	for _, vertice := range vertices {
		if !minXOk || vertice.X < minX {
			minX = vertice.X
			minXOk = true
		}
		if !maxXOk || vertice.X > maxX {
			maxX = vertice.X
			maxXOk = true
		}
		if !minYOk || vertice.Y < minY {
			minY = vertice.Y
			minYOk = true
		}
		if !maxYOk || vertice.Y > maxY {
			maxY = vertice.Y
			maxYOk = true
		}
	}

	for y := maxY + 1; y >= minY-1; y-- {
		n := y % 10
		if n < 0 {
			n = -n
		}
		prefix := chars[n]
		if _, err := fmt.Fprint(w, prefix); err != nil {
			return err
		}

		for x := minX - 1; x <= maxX+1; x++ {
			ponto := Ponto{X: x, Y: y}
			var text string
			switch {
			case slices.Contains(vertices, ponto):
				text = "█"
			case slices.Contains(pontos, ponto):
				text = "░"
			case ponto.X == 0 && ponto.Y == 0:
				text = "╋"
			case ponto.X == 0:
				text = "┃"
			case ponto.Y == 0:
				text = "━"
			case slices.ContainsFunc(vertices, func(p Ponto) bool {
				return ponto.X == p.X
			}):
				text = "┊"
			case slices.ContainsFunc(vertices, func(p Ponto) bool {
				return ponto.Y == p.Y
			}):
				text = "╌"
			default:
				text = " "
			}
			if _, err := fmt.Fprint(w, text); err != nil {
				return err
			}
		}
		if _, err := fmt.Fprint(w, "\n"); err != nil {
			return err
		}
	}

	if _, err := fmt.Fprint(w, " "); err != nil {
		return err
	}
	for x := minX - 1; x <= maxX+1; x++ {
		n := x % 10
		if n < 0 {
			n = -n
		}
		text := chars[n]

		if _, err := fmt.Fprint(w, text); err != nil {
			return err
		}
	}
	return nil
}
