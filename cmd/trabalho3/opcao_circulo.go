package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesDesenharCirculo struct {
	pontoA *entradaPonto
}

func (c *configuracoesDesenharCirculo) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontoA}
}

func gen8(ponto ufpa_cg.Ponto, centro ufpa_cg.Ponto) []ufpa_cg.Ponto {
	return []ufpa_cg.Ponto{
		{X: ponto.X + centro.X, Y: ponto.Y + centro.Y},
		{X: ponto.Y + centro.X, Y: ponto.X + centro.Y},
		{X: ponto.Y + centro.X, Y: -ponto.X + centro.Y},
		{X: ponto.X + centro.X, Y: -ponto.Y + centro.Y},
		{X: -ponto.X + centro.X, Y: -ponto.Y + centro.Y},
		{X: -ponto.Y + centro.X, Y: -ponto.X + centro.Y},
		{X: -ponto.Y + centro.X, Y: ponto.X + centro.Y},
		{X: -ponto.X + centro.X, Y: ponto.Y + centro.Y},
	}
}

func (c *configuracoesDesenharCirculo) Evaluate() []ufpa_cg.Ponto {
	centro := c.pontoA.ponto

	r := 5
	pontos := make([]ufpa_cg.Ponto, 0)
	x := 0
	y := r
	e := -r
	pontos = append(pontos, gen8(ufpa_cg.Ponto{X: 0, Y: r}, centro)...)
	for x <= y {
		e += 2*x + 1
		x++
		if e >= 0 {
			e += 2 - 2*y
			y--
		}
		pontos = append(pontos, gen8(ufpa_cg.Ponto{X: x, Y: y}, centro)...)
	}
	return pontos
}

type moduloDesenhaCirculo struct {
	settings ConfiguracaoModulo
}

type opcaoDesenharCirculo struct{}

func NewOpcaoDesenharCirculo() OpcaoMenu {
	return &opcaoDesenharCirculo{}
}

func (o *opcaoDesenharCirculo) Title() string {
	return "Desenhar círculo"
}

func (o *opcaoDesenharCirculo) Create() ModuloJogo {
	return &moduloDesenhaCirculo{
		settings: &configuracoesDesenharCirculo{
			pontoA: &entradaPonto{Label: "centro"},
		},
	}
}

func (m *moduloDesenhaCirculo) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloDesenhaCirculo) Update() error {
	return nil
}

func (m *moduloDesenhaCirculo) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
