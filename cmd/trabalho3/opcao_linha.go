package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesDesenharLinha struct {
	pontoA *entradaPonto
	pontoB *entradaPonto
}

func (c *configuracoesDesenharLinha) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontoA, c.pontoB}
}

func (c *configuracoesDesenharLinha) Evaluate() []ufpa_cg.Ponto {
	pontoA := c.pontoA.ponto
	pontoB := c.pontoB.ponto

	pontos := make([]ufpa_cg.Ponto, 0)
	algoritmo := ufpa_cg.NewAlgoritmoBresenham(pontoA, pontoB)
	for algoritmo.Move() {
		pontos = append(pontos, algoritmo.PontoAtual())
	}
	return pontos
}

type moduloDesenhaLinha struct {
	settings *configuracoesDesenharLinha
}

type opcaoDesenharLinha struct{}

func NewOpcaoDesenharLinha() OpcaoMenu {
	return &opcaoDesenharLinha{}
}

func (o *opcaoDesenharLinha) Title() string {
	return "Desenhar linha"
}

func (o *opcaoDesenharLinha) Create() ModuloJogo {
	return &moduloDesenhaLinha{
		settings: &configuracoesDesenharLinha{
			pontoA: &entradaPonto{Label: "ponto A"},
			pontoB: &entradaPonto{Label: "ponto B"},
		},
	}
}

func (m *moduloDesenhaLinha) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloDesenhaLinha) Update() error {
	return nil
}

func (m *moduloDesenhaLinha) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
