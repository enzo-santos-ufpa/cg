package main

import (
	"fmt"
	"image/color"
	"strings"
	"ufpa_cg"
)

type entradaPonto3D struct {
	Label string

	entradaX *entradaInteiro
	entradaY *entradaInteiro
	entradaZ *entradaInteiro

	ponto ufpa_cg.Ponto3D
	ok    bool
}

func (e *entradaPonto3D) entradas() []*entradaInteiro {
	return []*entradaInteiro{e.entradaX, e.entradaY, e.entradaZ}
}

func (e *entradaPonto3D) Selected(ufpa_cg.Ponto) bool {
	return false
}

func (e *entradaPonto3D) OnUpdate() {
	for _, inp := range e.entradas() {
		if _, evaluated := inp.Evaluated(); !evaluated {
			inp.OnUpdate()
			return
		}
	}
	e.ponto = ufpa_cg.Ponto3D{
		X: e.entradaX.valor,
		Y: e.entradaY.valor,
		Z: e.entradaZ.valor,
	}
	e.ok = true
}

func (e *entradaPonto3D) OnDraw(ufpa_cg.Ponto, int, int, int) (color.Color, bool) {
	return nil, false
}

func (e *entradaPonto3D) DescribeState() (string, bool) {
	buffer := strings.Builder{}
	buffer.WriteString("(")
	for i, inp := range e.entradas() {
		if _, evaluated := inp.Evaluated(); !evaluated {
			if label, ok := inp.DescribeState(); ok {
				return fmt.Sprintf("%s %s", buffer.String(), label), true
			} else {
				return buffer.String(), true
			}
		}
		buffer.WriteString(inp.DescribeValue())
		if i < len(e.entradas())-1 {
			buffer.WriteString(",")
		}
	}
	return "", false
}

func (e *entradaPonto3D) DescribePrompt() string {
	for i, inp := range e.entradas() {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return fmt.Sprintf("Selecione o eixo %c:", rune(i+'X'))
		}
	}
	return ""
}

func (e *entradaPonto3D) DescribeAction() (string, bool) {
	return "", false
}

func (e *entradaPonto3D) DescribeValue() string {
	return fmt.Sprintf("(%d, %d, %d)", e.ponto.X, e.ponto.Y, e.ponto.Z)
}

func (e *entradaPonto3D) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	return nil, e.ok
}

func (e *entradaPonto3D) Reset() {
	for _, inp := range e.entradas() {
		inp.Reset()
	}
	e.ponto = ufpa_cg.Ponto3D{}
	e.ok = false
}
