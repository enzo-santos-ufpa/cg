package main

import (
	"fmt"
	"image/color"
	"ufpa_cg"
)

type entradaJanela struct {
	entradaPontoSuperiorEsquerdo *entradaPonto
	entradaPontoInferiorDireito  *entradaPonto
}

func (e *entradaJanela) Selected(ponto ufpa_cg.Ponto) bool {
	for _, inp := range []*entradaPonto{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		if _, evaluated := inp.Evaluated(); evaluated && inp.ponto == ponto {
			return true
		}
	}
	return false
}

func (e *entradaJanela) OnUpdate() {
	for _, inp := range []EntradaModulo{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			inp.OnUpdate()
			return
		}
	}
}

func (e *entradaJanela) OnDraw(ponto ufpa_cg.Ponto, x, y int, size int) (color.Color, bool) {
	if _, evaluated := e.entradaPontoSuperiorEsquerdo.Evaluated(); evaluated {
		if _, evaluated := e.entradaPontoInferiorDireito.Evaluated(); !evaluated {
			return e.entradaPontoInferiorDireito.OnDraw(ponto, x, y, size)
		}
	} else {
		return e.entradaPontoSuperiorEsquerdo.OnDraw(ponto, x, y, size)
	}
	return nil, false
}

func (e *entradaJanela) DescribeState() (string, bool) {
	for _, inp := range []*entradaPonto{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return inp.DescribeState()
		}
	}
	return "", false
}

func (e *entradaJanela) DescribePrompt() string {
	if _, evaluated := e.entradaPontoSuperiorEsquerdo.Evaluated(); !evaluated {
		return "Selecione o ponto superior esquerdo:"
	}
	pontoSuperiorEsquerdo := e.entradaPontoSuperiorEsquerdo.ponto
	if _, evaluated := e.entradaPontoInferiorDireito.Evaluated(); !evaluated {
		return fmt.Sprintf(
			"Selecione o ponto inferior direito: (%d, %d), ",
			pontoSuperiorEsquerdo.X,
			pontoSuperiorEsquerdo.Y,
		)
	}
	pontoInferiorDireito := e.entradaPontoInferiorDireito.ponto
	return fmt.Sprintf(
		"Selecione o ponto inferior direito: (%d, %d), (%d, %d)",
		pontoSuperiorEsquerdo.X,
		pontoSuperiorEsquerdo.Y,
		pontoInferiorDireito.X,
		pontoInferiorDireito.Y,
	)
}

func (e *entradaJanela) DescribeAction() (string, bool) {
	return "", false
}

func (e *entradaJanela) DescribeValue() string {
	for _, inp := range []*entradaPonto{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return inp.DescribeValue()
		}
	}
	return ""
}

func (e *entradaJanela) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	stamps := make(map[ufpa_cg.Ponto]color.Color)
	for _, inp := range []*entradaPonto{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return stamps, false
		} else {
			stamps[inp.ponto] = color.RGBA{R: 0x63, G: 0x63, B: 0x63, A: 0xFF}
		}
	}
	pontoSuperiorEsquerdo := e.entradaPontoSuperiorEsquerdo.ponto
	pontoInferiorDireito := e.entradaPontoInferiorDireito.ponto
	for x := pontoSuperiorEsquerdo.X - 1; x <= pontoInferiorDireito.X+1; x++ {
		stamps[ufpa_cg.Ponto{X: x, Y: pontoSuperiorEsquerdo.Y + 1}] = color.Black
		stamps[ufpa_cg.Ponto{X: x, Y: pontoInferiorDireito.Y - 1}] = color.Black
	}
	for y := pontoInferiorDireito.Y - 1; y <= pontoSuperiorEsquerdo.Y+1; y++ {
		stamps[ufpa_cg.Ponto{X: pontoSuperiorEsquerdo.X - 1, Y: y}] = color.Black
		stamps[ufpa_cg.Ponto{X: pontoInferiorDireito.X + 1, Y: y}] = color.Black
	}
	return stamps, true
}

func (e *entradaJanela) Reset() {
	for _, inp := range []EntradaModulo{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		inp.Reset()
	}
}
