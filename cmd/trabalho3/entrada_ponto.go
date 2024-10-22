package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image/color"
	"ufpa_cg"
)

type estadoEntradaPonto struct {
	cursorX    int
	cursorY    int
	pontoAtual *ufpa_cg.Ponto
}

type entradaPonto struct {
	Label string

	estado estadoEntradaPonto
	ponto  ufpa_cg.Ponto
	ok     bool
}

func (e *entradaPonto) Selected(ponto ufpa_cg.Ponto) bool {
	return e.ok && e.ponto == ponto
}

func (e *entradaPonto) OnUpdate() {
	x, y := ebiten.CursorPosition()
	e.estado.cursorX = x
	e.estado.cursorY = y

	if ponto := e.estado.pontoAtual; inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && ponto != nil {
		e.ponto = *ponto
		e.ok = true
	}
}

func (e *entradaPonto) OnDraw(ponto ufpa_cg.Ponto, x, y int, size int) (color.Color, bool) {
	var hoveredColor = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF} // Cor para um ponto com o cursor apontado
	if (x <= e.estado.cursorX && e.estado.cursorX <= x+size) &&       // Se a posição X do cursor estiver dentro deste ponto
		(y <= e.estado.cursorY && e.estado.cursorY <= y+size) { // Se a posição Y do cursor estiver dentro deste ponto
		e.estado.pontoAtual = &ufpa_cg.Ponto{X: ponto.X, Y: ponto.Y}
		return hoveredColor, true
	}
	return nil, false
}

func (e *entradaPonto) OnDisplay() (string, bool) {
	if ponto := e.estado.pontoAtual; ponto != nil {
		return fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), true
	} else {
		return "", false
	}
}

func (e *entradaPonto) DescribeLabel() string {
	return e.Label
}

func (e *entradaPonto) DescribeValue() string {
	return fmt.Sprintf("(%d, %d)", e.ponto.X, e.ponto.Y)
}

func (e *entradaPonto) Evaluated() bool {
	return e.ok
}

func (e *entradaPonto) Reset() {
	e.ponto = ufpa_cg.Ponto{}
	e.ok = false
}