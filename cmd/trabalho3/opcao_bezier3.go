package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesDesenharBezier3 struct {
	pontoA *entradaPonto
	pontoB *entradaPonto
	pontoC *entradaPonto
	pontoD *entradaPonto
}

func (c *configuracoesDesenharBezier3) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontoA, c.pontoB, c.pontoC, c.pontoD}
}

func (c *configuracoesDesenharBezier3) Evaluate() []ufpa_cg.Ponto {
	pontoA := c.pontoA.ponto
	pontoB := c.pontoB.ponto
	pontoC := c.pontoC.ponto
	pontoD := c.pontoD.ponto

	arestas := make([]ufpa_cg.Ponto, 5)
	for i := 0; i <= 4; i++ {
		t := float64(i) / 4
		arestas[i] = bezier([]ufpa_cg.Ponto{pontoA, pontoB, pontoC, pontoD}, t)
	}
	pontos := make([]ufpa_cg.Ponto, 0)
	for i := 0; i < len(arestas)-1; i++ {
		aresta1 := arestas[i]
		aresta2 := arestas[i+1]
		algoritmo := ufpa_cg.NewAlgoritmoBresenham(aresta1, aresta2)
		for algoritmo.Move() {
			pontos = append(pontos, algoritmo.PontoAtual())
		}
	}
	return pontos
}

type moduloDesenhaBezier3 struct {
	settings ConfiguracaoModulo
}

type opcaoDesenharBezier3 struct{}

func NewOpcaoDesenharBezier3() OpcaoMenu {
	return &opcaoDesenharBezier3{}
}

func (o *opcaoDesenharBezier3) Title() string {
	return "Desenhar curva de Bezier (grau 3)"
}

func (o *opcaoDesenharBezier3) Create() ModuloJogo {
	return &moduloDesenhaBezier3{
		settings: &configuracoesDesenharBezier3{
			pontoA: &entradaPonto{Label: "ponto A"},
			pontoB: &entradaPonto{Label: "ponto B"},
			pontoC: &entradaPonto{Label: "ponto C"},
			pontoD: &entradaPonto{Label: "ponto D"},
		},
	}
}

func (m *moduloDesenhaBezier3) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloDesenhaBezier3) Update() error {
	return nil
}

func (m *moduloDesenhaBezier3) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
