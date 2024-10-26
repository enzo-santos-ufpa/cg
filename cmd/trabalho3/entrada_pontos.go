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
	borda    map[ufpa_cg.Ponto]color.Color
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
			if _, evaluated := inp.Evaluated(); evaluated {
				count++
			}
		}
		if count >= e.Minimo {
			e.ok = true
		}
	default:
		for i, inp := range e.entradas {
			if _, evaluated := inp.Evaluated(); !evaluated {
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
		if _, evaluated := inp.Evaluated(); !evaluated {
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
		if _, evaluated := inp.Evaluated(); !evaluated {
			if ponto := inp.estado.pontoAtual; ponto != nil {
				return fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), true
			}
			break
		}

	}
	return "", false
}

func (e *entradaPontos) DescribePrompt() string {
	buffer := strings.Builder{}
	for i, inp := range e.entradas {
		if _, evaluated := inp.Evaluated(); !evaluated {
			break
		}
		ponto := inp.ponto
		buffer.WriteString(fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y))
		if i < len(e.entradas)-1 {
			buffer.WriteString(", ")
		}
	}
	return fmt.Sprintf("Selecione o conjunto de pontos: %s", buffer.String())
}

func (e *entradaPontos) DescribeValue() string {
	buffer := strings.Builder{}
	for i, inp := range e.entradas {
		if _, evaluated := inp.Evaluated(); !evaluated {
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

func (e *entradaPontos) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	if !e.ok {
		return nil, false
	}
	var hoveredColor = color.RGBA{R: 0xFF, G: 0xFF, B: 0x00, A: 0xFF} // Cor para um ponto com o cursor apontado
	if borda := e.borda; borda == nil {
		b := make(map[ufpa_cg.Ponto]color.Color)
		vertices := make([]ufpa_cg.Ponto, len(e.entradas)-1)
		for i, inp := range e.entradas[:len(e.entradas)-1] {
			vertices[i] = inp.ponto
		}
		for i := 0; i < len(vertices); i++ {
			ponto1 := vertices[i]
			var ponto2 ufpa_cg.Ponto
			if i < len(vertices)-1 {
				ponto2 = vertices[i+1]
			} else {
				ponto2 = vertices[0]
			}
			algoritmo := ufpa_cg.NewAlgoritmoBresenham(ponto1, ponto2)
			for algoritmo.Move() {
				b[algoritmo.PontoAtual()] = hoveredColor
			}
		}
		e.borda = b
		return b, true
	} else {
		return borda, true
	}
}

func (e *entradaPontos) Reset() {
	e.entradas = []*entradaPonto{{}}
	e.borda = nil
	e.ok = false
}
