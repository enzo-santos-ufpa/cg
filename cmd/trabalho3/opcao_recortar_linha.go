package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesRecortarLinha struct {
	pontoA        *entradaPonto
	pontoB        *entradaPonto
	entradaJanela *entradaJanela
}

func (c *configuracoesRecortarLinha) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontoA, c.pontoB, c.entradaJanela}
}

func (c *configuracoesRecortarLinha) Evaluate() []ufpa_cg.Ponto {
	janela := c.entradaJanela.JanelaRecorte()
	pontoA := c.pontoA.ponto
	pontoB := c.pontoB.ponto
	pontos := make([]ufpa_cg.Ponto, 0)
	ponto1, ponto2, ok := janela.CalculaIntersecao(pontoA, pontoB)
	if !ok {
		return pontos
	}
	algoritmo := ufpa_cg.NewAlgoritmoBresenham(ponto1, ponto2)
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
			entradaJanela: &entradaJanela{
				entradaPonto1: &entradaPonto{},
				entradaPonto2: &entradaPonto{},
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
