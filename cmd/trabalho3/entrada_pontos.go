package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"strings"
	"ufpa_cg"
)

type entradaPontos struct {
	Minimo int

	entradas []*entradaPonto
	ok       bool
}

func (e *entradaPontos) Selected(ponto ufpa_cg.Ponto) bool {
	for _, inp := range e.entradas {
		if inp.Selected(ponto) {
			return true
		}
	}
	return false
}

func (e *entradaPontos) OnUpdate() {
	switch {
	case repeatingKeyPressed(ebiten.KeyEnter):
		count := 0
		for _, inp := range e.entradas {
			if inp.Evaluated() {
				count++
			}
		}
		if count >= e.Minimo {
			e.ok = true
		}
	default:
		for i, inp := range e.entradas {
			if !inp.Evaluated() {
				inp.OnUpdate()
				break
			} else if i == len(e.entradas)-1 {
				e.entradas = append(e.entradas, new(entradaPonto))
			}
		}
	}
}

func (e *entradaPontos) OnDraw(ponto ufpa_cg.Ponto, x, y int, size int) (color.Color, bool) {
	for _, inp := range e.entradas {
		if !inp.Evaluated() {
			if customColor, ok := inp.OnDraw(ponto, x, y, size); ok {
				return customColor, true
			}
			break
		}
	}
	return nil, false
}

func (e *entradaPontos) OnDisplay() (string, bool) {
	for _, inp := range e.entradas {
		if !inp.Evaluated() {
			if ponto := inp.estado.pontoAtual; ponto != nil {
				return fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), true
			}
			break
		}

	}
	return "", false
}

func (e *entradaPontos) DescribeLabel() string {
	return "conjunto de pontos"
}

func (e *entradaPontos) DescribeValue() string {
	buffer := strings.Builder{}
	for i, inp := range e.entradas {
		if !inp.Evaluated() {
			break
		}
		ponto := inp.ponto
		buffer.WriteString(fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y))
		if i < len(e.entradas)-2 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}

func (e *entradaPontos) Evaluated() bool {
	return e.ok
}

func (e *entradaPontos) Reset() {
	e.entradas = []*entradaPonto{{}}
	e.ok = false
}
