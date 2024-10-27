package main

import (
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
		pontoSuperiorEsquerdo := e.entradaPontoSuperiorEsquerdo.ponto
		if ponto.X < pontoSuperiorEsquerdo.X {
			return color.Black, true
		}
		if ponto.Y > pontoSuperiorEsquerdo.Y {
			return color.Black, true
		}
		if _, evaluated := e.entradaPontoInferiorDireito.Evaluated(); evaluated {
			pontoInferiorDireito := e.entradaPontoInferiorDireito.ponto
			if pontoInferiorDireito.X < ponto.X {
				return color.Black, true
			}
			if pontoInferiorDireito.Y > ponto.Y {
				return color.Black, true
			}
		} else {
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
	return "Selecione o ponto inferior direito:"
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
	for i, inp := range []*entradaPonto{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return stamps, false
		} else {
			stamps[inp.ponto] = color.RGBA{R: 0x63, G: 0x63, B: 0x63, A: 0xFF}
		}

		for j_ := -11; j_ <= 11; j_++ {
			for i_ := -11; i_ <= 11; i_++ {
				ponto := ufpa_cg.Ponto{X: i_, Y: -j_}
				if i == 0 && (ponto.X < inp.ponto.X || ponto.Y > inp.ponto.Y) {
					stamps[ponto] = color.Black
				} else if i == 1 && (ponto.X > inp.ponto.X || ponto.Y < inp.ponto.Y) {
					stamps[ponto] = color.Black
				}
			}
		}
	}
	return stamps, true
}

func (e *entradaJanela) Reset() {
	for _, inp := range []EntradaModulo{e.entradaPontoSuperiorEsquerdo, e.entradaPontoInferiorDireito} {
		inp.Reset()
	}
}
