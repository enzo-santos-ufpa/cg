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
	Label        string
	Minimo       int
	PossuiMinimo bool
	Maximo       int
	PossuiMaximo bool
	Sufixo       string

	estado estadoEntradaInteiro
	valor  int
	ok     bool
}

func (e *entradaInteiro) Selected(_ ufpa_cg.Ponto) bool {
	return false
}

func (e *entradaInteiro) OnUpdate() bool {
	switch {
	case repeatingKeyPressed(ebiten.KeyDown) &&
		(!e.PossuiMinimo || e.estado.valorAtual >= e.Minimo):
		e.estado.valorAtual--
		return true
	case repeatingKeyPressed(ebiten.KeyUp) &&
		(!e.PossuiMaximo || e.estado.valorAtual <= e.Maximo):
		e.estado.valorAtual++
		return true
	case repeatingKeyPressed(ebiten.KeyEnter) &&
		(!e.PossuiMinimo || e.estado.valorAtual >= e.Minimo) &&
		(!e.PossuiMaximo || e.estado.valorAtual <= e.Maximo):
		e.valor = e.estado.valorAtual
		e.ok = true
		return true
	}
	return false
}

func (e *entradaInteiro) OnDraw(_ ufpa_cg.Ponto, _, _ int, _ int) (color.Color, bool) {
	return nil, false
}

func (e *entradaInteiro) DescribeState() (string, bool) {
	return fmt.Sprintf("%d%s", e.estado.valorAtual, e.Sufixo), true
}

func (e *entradaInteiro) DescribePrompt() string {
	return fmt.Sprintf("Selecione o %s:", e.Label)
}

func (e *entradaInteiro) DescribeActions() []AcaoEntrada {
	return []AcaoEntrada{
		{Titulo: "↑ ou ↓", Descricao: "selecionar"},
		{Titulo: "ENTER", Descricao: "prosseguir"},
	}
}

func (e *entradaInteiro) DescribeValue() string {
	return fmt.Sprintf("%d%s", e.valor, e.Sufixo)
}

func (e *entradaInteiro) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	return nil, e.ok
}

func (e *entradaInteiro) Reset() {
	e.estado = estadoEntradaInteiro{}
	e.valor = 0
	e.ok = false
}
