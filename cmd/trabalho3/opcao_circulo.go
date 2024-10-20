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

func (c *configuracoesDesenharCirculo) Evaluate() []ufpa_cg.Ponto {
	pontoA := c.pontoA.ponto

	pontos := make([]ufpa_cg.Ponto, 0)
	algoritmo := ufpa_cg.NewAlgoritmoBresenham(pontoA, ufpa_cg.Ponto{X: pontoA.X + 3, Y: pontoA.Y - 3})
	for algoritmo.Move() {
		pontos = append(pontos, algoritmo.PontoAtual())
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
			pontoA: &entradaPonto{},
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
