package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesDesenharElipse struct {
	centro *entradaPonto
	eixoA  *entradaInteiro
	eixoB  *entradaInteiro
}

func (c *configuracoesDesenharElipse) Inputs() []EntradaModulo {
	return []EntradaModulo{c.centro, c.eixoA, c.eixoB}
}

func gen4(ponto ufpa_cg.Ponto, centro ufpa_cg.Ponto) []ufpa_cg.Ponto {
	return []ufpa_cg.Ponto{
		{X: ponto.X + centro.X, Y: ponto.Y + centro.Y},
		{X: ponto.X + centro.X, Y: -ponto.Y + centro.Y},
		{X: -ponto.X + centro.X, Y: -ponto.Y + centro.Y},
		{X: -ponto.X + centro.X, Y: ponto.Y + centro.Y},
	}
}

func (c *configuracoesDesenharElipse) Evaluate() []ufpa_cg.Ponto {
	centro := c.centro.ponto
	a := c.eixoA.valor
	b := c.eixoB.valor

	pontos := make([]ufpa_cg.Ponto, 0)

	x := 0
	y := b
	aSq := a * a
	bSq := b * b
	xDv := 2 * bSq * x
	yDv := 2 * aSq * y
	var e float64
	e = -float64(b*aSq) + float64(aSq)*0.25
	for xDv < yDv {
		pontos = append(pontos, gen4(ufpa_cg.Ponto{X: x, Y: y}, centro)...)

		x++
		e += float64(xDv + bSq)
		xDv += 2 * bSq
		if e > 0 {
			y--
			e += float64(aSq - yDv)
			yDv -= 2 * aSq
		}
	}
	e = float64(bSq)*(float64(x)+0.5)*(float64(x)+0.5) + float64(aSq*y*y) - float64(aSq*bSq)
	for y >= 0 {
		pontos = append(pontos, gen4(ufpa_cg.Ponto{X: x, Y: y}, centro)...)

		y--
		e += float64(aSq - yDv)
		yDv -= 2 * aSq
		if e < 0 {
			x++
			e += float64(xDv + bSq)
			xDv += 2 * bSq
		}
	}
	return pontos
}

type moduloDesenhaElipse struct {
	settings ConfiguracaoModulo
}

type opcaoDesenharElipse struct{}

func NewOpcaoDesenharElipse() OpcaoMenu {
	return &opcaoDesenharElipse{}
}

func (o *opcaoDesenharElipse) Title() string {
	return "Desenhar elipse"
}

func (o *opcaoDesenharElipse) Create() ModuloJogo {
	return &moduloDesenhaElipse{
		settings: &configuracoesDesenharElipse{
			centro: &entradaPonto{Label: "centro"},
			eixoA: &entradaInteiro{
				Label:        "eixo A",
				PossuiMinimo: true,
				Minimo:       1,
			},
			eixoB: &entradaInteiro{
				Label:        "eixo B",
				PossuiMinimo: true,
				Minimo:       1,
			},
		},
	}
}

func (m *moduloDesenhaElipse) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloDesenhaElipse) Update() error {
	return nil
}

func (m *moduloDesenhaElipse) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
