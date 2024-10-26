package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"ufpa_cg"
)

type estadoEntradaInteiro struct {
	valorAtual int
}

type entradaInteiro struct {
	Label string

	estado estadoEntradaInteiro
	valor  int
	ok     bool
}

func (e *entradaInteiro) Selected(_ ufpa_cg.Ponto) bool {
	return false
}

func (e *entradaInteiro) OnUpdate() {
	switch {
	case repeatingKeyPressed(ebiten.KeyDown) && e.estado.valorAtual > 0:
		e.estado.valorAtual--
	case repeatingKeyPressed(ebiten.KeyUp):
		e.estado.valorAtual++
	case repeatingKeyPressed(ebiten.KeyEnter) && e.estado.valorAtual > 0:
		e.valor = e.estado.valorAtual
		e.ok = true
	}
}

func (e *entradaInteiro) OnDraw(_ ufpa_cg.Ponto, _, _ int, _ int) (color.Color, bool) {
	return nil, false
}

func (e *entradaInteiro) OnDisplay() (string, bool) {
	return fmt.Sprintf("%d", e.estado.valorAtual), true
}

func (e *entradaInteiro) DescribePrompt() string {
	return fmt.Sprintf("Selecione o %s:", e.Label)
}

func (e *entradaInteiro) DescribeValue() string {
	return fmt.Sprintf("%d", e.valor)
}

func (e *entradaInteiro) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	return nil, e.ok
}

func (e *entradaInteiro) Reset() {
	e.estado = estadoEntradaInteiro{}
	e.valor = 0
	e.ok = false
}
