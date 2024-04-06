package ufpa_cg

import (
	"bufio"
	"io"
	"slices"
)

func Exibe(algoritmo AlgoritmoLinha, w io.Writer) error {
	pontos := make([]Ponto, 0)
	for algoritmo.Move() {
		pontos = append(pontos, algoritmo.PontoAtual())
	}
	p1 := pontos[0]
	p2 := pontos[len(pontos)-1]

	writer := bufio.NewWriter(w)
	defer writer.Flush()

	chars := []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}

	var minX, maxX int
	if p2.X > p1.X {
		minX = p1.X
		maxX = p2.X
	} else {
		minX = p2.X
		maxX = p1.X
	}

	var minY, maxY int
	if p2.Y > p1.Y {
		minY = p1.Y
		maxY = p2.Y
	} else {
		minY = p2.Y
		maxY = p1.Y
	}

	for y := maxY + 1; y >= minY-1; y-- {
		n := y % 10
		if n < 0 {
			n = -n
		}
		prefix := chars[n]
		if _, err := writer.WriteString(prefix); err != nil {
			return err
		}

		for x := minX - 1; x <= maxX+1; x++ {
			ponto := Ponto{X: x, Y: y}
			var text string
			switch {
			case ponto == p1 || ponto == p2:
				text = "█"
			case slices.Contains(pontos, ponto):
				text = "░"
			case ponto.X == 0 && ponto.Y == 0:
				text = "╋"
			case ponto.X == 0:
				text = "┃"
			case ponto.Y == 0:
				text = "━"
			case ponto.X == p1.X || ponto.X == p2.X:
				text = "┊"
			case ponto.Y == p1.Y || ponto.Y == p2.Y:
				text = "╌"
			default:
				text = " "
			}
			if _, err := writer.WriteString(text); err != nil {
				return err
			}
		}
		if _, err := writer.WriteString("\n"); err != nil {
			return err
		}
	}

	if _, err := writer.WriteString(" "); err != nil {
		return err
	}
	for x := minX - 1; x <= maxX+1; x++ {
		n := x % 10
		if n < 0 {
			n = -n
		}
		text := chars[n]

		if _, err := writer.WriteString(text); err != nil {
			return err
		}
	}
	return nil
}
