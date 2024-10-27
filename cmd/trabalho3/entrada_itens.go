package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"ufpa_cg"
)

type entradaItens[T any] struct {
	Label   string
	Itens   []T
	Labeler func(value T) string

	index int
	ok    bool
}

func (e *entradaItens[T]) Selected(_ ufpa_cg.Ponto) bool {
	return false
}

func (e *entradaItens[T]) OnUpdate() {
	switch {
	case repeatingKeyPressed(ebiten.KeyDown):
		e.index = (e.index - 1 + len(e.Itens)) % len(e.Itens)
	case repeatingKeyPressed(ebiten.KeyUp):
		e.index = (e.index + 1) % len(e.Itens)
	case repeatingKeyPressed(ebiten.KeyEnter):
		e.ok = true
	}
}

func (e *entradaItens[T]) OnDraw(_ ufpa_cg.Ponto, _, _ int, _ int) (color.Color, bool) {
	return nil, false
}

func (e *entradaItens[T]) DescribeState() (string, bool) {
	return fmt.Sprintf("%s", e.Labeler(e.Itens[e.index])), true
}

func (e *entradaItens[T]) DescribePrompt() string {
	return fmt.Sprintf("Selecione o %s:", e.Label)
}

func (e *entradaItens[T]) DescribeActions() []AcaoEntrada {
	return []AcaoEntrada{
		{Titulo: "↑ ou ↓", Descricao: "selecionar"},
		{Titulo: "ENTER", Descricao: "prosseguir"},
	}
}

func (e *entradaItens[T]) DescribeValue() string {
	return fmt.Sprintf("%s", e.Labeler(e.Itens[e.index]))
}

func (e *entradaItens[T]) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	return nil, e.ok
}

func (e *entradaItens[T]) Reset() {
	e.index = 0
	e.ok = false
}
