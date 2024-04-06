package ufpa_cg

import (
	"bufio"
	"io"
	"slices"
)

type Sequencia struct {
	Inicio int
	Fim    int

	valor int
	ok    bool
}

func NewSequencia(inicio, fim int) *Sequencia {
	return &Sequencia{
		Inicio: inicio,
		Fim:    fim,
		valor:  0,
		ok:     false,
	}
}

func (s *Sequencia) Move() bool {
	var valorFim, valorInicio int
	if s.Inicio <= s.Fim {
		valorInicio = s.Inicio - 1
		valorFim = s.Fim + 1
	} else {
		valorInicio = s.Fim - 1
		valorFim = s.Inicio + 1
	}

	valor := s.valor
	if !s.ok {
		s.valor = valorFim
		s.ok = true
		return true
	}
	if valor == valorInicio {
		return false
	}
	s.valor--
	return true
}

func (s *Sequencia) Value() int {
	return s.valor
}

func Exibe(algoritmo AlgoritmoLinha, w io.Writer) error {
	pontos := make([]Ponto, 0)
	for algoritmo.Move() {
		pontos = append(pontos, algoritmo.PontoAtual())
	}
	p1 := pontos[0]
	p2 := pontos[len(pontos)-1]

	minX := p1.X
	maxX := p2.X

	writer := bufio.NewWriter(w)
	defer writer.Flush()

	chars := []string{"⁰", "¹", "²", "³", "⁴", "⁵", "⁶", "⁷", "⁸", "⁹"}

	sequenciaY := NewSequencia(p2.Y, p1.Y)
	for sequenciaY.Move() {
		y := sequenciaY.Value()

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
			if ponto == p1 || ponto == p2 {
				text = "█"
			} else if slices.Contains(pontos, ponto) {
				text = "░"
			} else if ponto.X == 0 && ponto.Y == 0 {
				text = "╋1"
			} else if ponto.X == 0 {
				text = "┃"
			} else if ponto.Y == 0 {
				text = "━"
			} else if ponto.X == p1.X || ponto.X == p2.X {
				text = "┊"
			} else if ponto.Y == p1.Y || ponto.Y == p2.Y {
				text = "╌"
			} else {
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
