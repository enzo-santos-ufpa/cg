package main

import (
	"image/color"
	"ufpa_cg"
)

type entradaPontoPolilinha struct {
	pontos *entradaPontos
	ponto  *entradaPonto

	ok bool
}

func (e *entradaPontoPolilinha) Selected(ponto ufpa_cg.Ponto) bool {
	if e.pontos.Selected(ponto) {
		return true
	}
	if e.ponto.Selected(ponto) {
		return true
	}
	return false
}

// pnpoly verifica se um dado ponto está contido num polígono formado por vertices.
//
// Fonte: https://wrfranklin.org/Research/Short_Notes/pnpoly.html
func pnpoly(vertices []ufpa_cg.Ponto, ponto ufpa_cg.Ponto) bool {
	var inside bool
	for i, j := 0, len(vertices)-1; i < len(vertices); j, i = i, i+1 {
		if ((vertices[i].Y > ponto.Y) != (vertices[j].Y > ponto.Y)) &&
			(ponto.X < (vertices[j].X-vertices[i].X)*(ponto.Y-vertices[i].Y)/(vertices[j].Y-vertices[i].Y)+vertices[i].X) {
			inside = !inside
		}
	}
	return inside
}

func (e *entradaPontoPolilinha) OnUpdate() bool {
	if _, evaluated := e.pontos.Evaluated(); !evaluated {
		return e.pontos.OnUpdate()
	}
	if _, evaluated := e.ponto.Evaluated(); !evaluated {
		return e.ponto.OnUpdate()
	}
	vertices := make([]ufpa_cg.Ponto, len(e.pontos.entradas)-1)
	for i, inp := range e.pontos.entradas[:len(e.pontos.entradas)-1] {
		vertices[i] = inp.ponto
	}
	if pnpoly(vertices, e.ponto.ponto) {
		e.ok = true
	} else {
		e.ponto.Reset()
	}
	return true
}

func (e *entradaPontoPolilinha) OnDraw(ponto ufpa_cg.Ponto, x, y int, size int) (color.Color, bool) {
	if stamp, evaluated := e.pontos.Evaluated(); evaluated {
		stampColor, ok := stamp[ponto]
		if !ok {
			return e.ponto.OnDraw(ponto, x, y, size)
		}
		return stampColor, ok
	}
	return e.pontos.OnDraw(ponto, x, y, size)
}

func (e *entradaPontoPolilinha) DescribeState() (string, bool) {
	for _, inp := range []EntradaModulo{e.pontos, e.ponto} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			if label, ok := inp.DescribeState(); ok {
				return label, true
			}
			break
		}

	}
	return "", false
}

func (e *entradaPontoPolilinha) DescribePrompt() string {
	if _, evaluated := e.pontos.Evaluated(); !evaluated {
		return e.pontos.DescribePrompt()
	}
	return e.ponto.DescribePrompt()
}

func (e *entradaPontoPolilinha) DescribeActions() []AcaoEntrada {
	if _, evaluated := e.pontos.Evaluated(); !evaluated {
		return e.pontos.DescribeActions()
	}
	return e.ponto.DescribeActions()
}

func (e *entradaPontoPolilinha) DescribeValue() string {
	for _, inp := range []EntradaModulo{e.ponto, e.pontos} {
		if _, evaluated := inp.Evaluated(); evaluated {
			return inp.DescribeValue()
		}
	}
	return ""
}

func (e *entradaPontoPolilinha) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	if !e.ok {
		return nil, false
	}
	stamps := make(map[ufpa_cg.Ponto]color.Color)
	for _, inp := range []EntradaModulo{e.pontos, e.ponto} {
		stamp, evaluated := inp.Evaluated()
		if !evaluated {
			continue
		}
		for point, pointColor := range stamp {
			stamps[point] = pointColor
		}
	}
	return stamps, true
}

func (e *entradaPontoPolilinha) Reset() {
	e.pontos.Reset()
	e.ponto.Reset()
	e.ok = false
}
