package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesRecortarLinha struct {
	pontoA *entradaPonto
	pontoB *entradaPonto
	janela *entradaJanela
}

func (c *configuracoesRecortarLinha) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontoA, c.pontoB, c.janela}
}

func (c *configuracoesRecortarLinha) Evaluate() []ufpa_cg.Ponto {
	pontoA := c.pontoA.ponto
	pontoB := c.pontoB.ponto

	pontos := make([]ufpa_cg.Ponto, 0)
	algoritmo := ufpa_cg.NewAlgoritmoBresenham(pontoA, pontoB)
	for algoritmo.Move() {
		pontos = append(pontos, algoritmo.PontoAtual())
	}
	return pontos
}

type moduloRecortaLinha struct {
	settings *configuracoesRecortarLinha
}

type opcaoRecortarLinha struct{}

func NewOpcaoRecortarLinha() OpcaoMenu {
	return &opcaoRecortarLinha{}
}

func (o *opcaoRecortarLinha) Title() string {
	return "Recortar linha"
}

func (o *opcaoRecortarLinha) Create() ModuloJogo {
	return &moduloRecortaLinha{
		settings: &configuracoesRecortarLinha{
			pontoA: &entradaPonto{Label: "ponto A"},
			pontoB: &entradaPonto{Label: "ponto B"},
			janela: &entradaJanela{
				entradaPontoSuperiorEsquerdo: &entradaPonto{},
				entradaPontoInferiorDireito:  &entradaPonto{},
			},
		},
	}
}

func (m *moduloRecortaLinha) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloRecortaLinha) Update() error {
	return nil
}

func (m *moduloRecortaLinha) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
