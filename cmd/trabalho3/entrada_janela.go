package main

import (
	"fmt"
	"image/color"
	"ufpa_cg"
)

type entradaJanela struct {
	entradaPonto1 *entradaPonto
	entradaPonto2 *entradaPonto
}

func (e *entradaJanela) JanelaRecorte() ufpa_cg.JanelaRecorte {
	janelaPonto1 := e.entradaPonto1.ponto
	janelaPonto2 := e.entradaPonto2.ponto
	pontoSuperiorEsquerdo := ufpa_cg.Ponto{
		X: min(janelaPonto1.X, janelaPonto2.X),
		Y: max(janelaPonto1.Y, janelaPonto2.Y),
	}
	pontoInferiorDireito := ufpa_cg.Ponto{
		X: max(janelaPonto1.X, janelaPonto2.X),
		Y: min(janelaPonto1.Y, janelaPonto2.Y),
	}
	return ufpa_cg.JanelaRecorte{
		PontoSuperiorEsquerdo: pontoSuperiorEsquerdo,
		PontoInferiorDireito:  pontoInferiorDireito,
	}
}

func (e *entradaJanela) Selected(ponto ufpa_cg.Ponto) bool {
	for _, inp := range []*entradaPonto{e.entradaPonto1, e.entradaPonto2} {
		if _, evaluated := inp.Evaluated(); evaluated && inp.ponto == ponto {
			return true
		}
	}
	return false
}

func (e *entradaJanela) OnUpdate() bool {
	for _, inp := range []EntradaModulo{e.entradaPonto1, e.entradaPonto2} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return inp.OnUpdate()
		}
	}
	return false
}

func (e *entradaJanela) OnDraw(ponto ufpa_cg.Ponto, x, y int, size int) (color.Color, bool) {
	if _, evaluated := e.entradaPonto1.Evaluated(); evaluated {
		if _, evaluated := e.entradaPonto2.Evaluated(); !evaluated {
			return e.entradaPonto2.OnDraw(ponto, x, y, size)
		}
	} else {
		return e.entradaPonto1.OnDraw(ponto, x, y, size)
	}
	return nil, false
}

func (e *entradaJanela) DescribeState() (string, bool) {
	for _, inp := range []*entradaPonto{e.entradaPonto1, e.entradaPonto2} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return inp.DescribeState()
		}
	}
	return "", false
}

func (e *entradaJanela) DescribePrompt() string {
	if _, evaluated := e.entradaPonto1.Evaluated(); !evaluated {
		return "Selecione o ponto 1 de recorte:"
	}
	ponto1 := e.entradaPonto1.ponto
	if _, evaluated := e.entradaPonto2.Evaluated(); !evaluated {
		return fmt.Sprintf(
			"Selecione o ponto 2 de recorte: (%d, %d), ",
			ponto1.X,
			ponto1.Y,
		)
	}
	ponto2 := e.entradaPonto2.ponto
	return fmt.Sprintf(
		"Selecione o ponto 2 de recorte: (%d, %d), (%d, %d)",
		ponto1.X,
		ponto1.Y,
		ponto2.X,
		ponto2.Y,
	)
}

func (e *entradaJanela) DescribeActions() []AcaoEntrada {
	return nil
}

func (e *entradaJanela) DescribeValue() string {
	for _, inp := range []*entradaPonto{e.entradaPonto1, e.entradaPonto2} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return inp.DescribeValue()
		}
	}
	return ""
}

func (e *entradaJanela) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	stamps := make(map[ufpa_cg.Ponto]color.Color)
	for _, inp := range []*entradaPonto{e.entradaPonto1, e.entradaPonto2} {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return stamps, false
		} else {
			stamps[inp.ponto] = color.RGBA{R: 0x63, G: 0x63, B: 0x63, A: 0xFF}
		}
	}
	janela := e.JanelaRecorte()
	for x := janela.PontoSuperiorEsquerdo.X - 1; x <= janela.PontoInferiorDireito.X+1; x++ {
		stamps[ufpa_cg.Ponto{X: x, Y: janela.PontoSuperiorEsquerdo.Y + 1}] = color.Black
		stamps[ufpa_cg.Ponto{X: x, Y: janela.PontoInferiorDireito.Y - 1}] = color.Black
	}
	for y := janela.PontoInferiorDireito.Y - 1; y <= janela.PontoSuperiorEsquerdo.Y+1; y++ {
		stamps[ufpa_cg.Ponto{X: janela.PontoSuperiorEsquerdo.X - 1, Y: y}] = color.Black
		stamps[ufpa_cg.Ponto{X: janela.PontoInferiorDireito.X + 1, Y: y}] = color.Black
	}
	return stamps, true
}

func (e *entradaJanela) Reset() {
	for _, inp := range []EntradaModulo{e.entradaPonto1, e.entradaPonto2} {
		inp.Reset()
	}
}
