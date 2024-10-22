package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesDesenharBezier2 struct {
	pontoA *entradaPonto
	pontoB *entradaPonto
	pontoC *entradaPonto
}

func (c *configuracoesDesenharBezier2) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontoA, c.pontoB, c.pontoC}
}

func bezier(pontos []ufpa_cg.Ponto, t float64) ufpa_cg.Ponto {
	if len(pontos) == 2 {
		pontoA := pontos[0]
		pontoB := pontos[1]
		return ufpa_cg.Ponto{
			X: int((1-t)*float64(pontoA.X) + t*float64(pontoB.X)),
			Y: int((1-t)*float64(pontoA.Y) + t*float64(pontoB.Y)),
		}
	}
	p := make([]ufpa_cg.Ponto, len(pontos)-1)
	for i := 0; i < len(pontos)-1; i++ {
		p[i] = bezier([]ufpa_cg.Ponto{pontos[i], pontos[i+1]}, t)
	}
	return bezier(p, t)
}

func (c *configuracoesDesenharBezier2) Evaluate() []ufpa_cg.Ponto {
	pontoA := c.pontoA.ponto
	pontoB := c.pontoB.ponto
	pontoC := c.pontoC.ponto

	arestas := make([]ufpa_cg.Ponto, 5)
	for i := 0; i <= 4; i++ {
		t := float64(i) / 4
		arestas[i] = bezier([]ufpa_cg.Ponto{pontoA, pontoB, pontoC}, t)
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

type moduloDesenhaBezier2 struct {
	settings ConfiguracaoModulo
}

type opcaoDesenharBezier2 struct{}

func NewOpcaoDesenharBezier2() OpcaoMenu {
	return &opcaoDesenharBezier2{}
}

func (o *opcaoDesenharBezier2) Title() string {
	return "Desenhar curva de Bezier (grau 2)"
}

func (o *opcaoDesenharBezier2) Create() ModuloJogo {
	return &moduloDesenhaBezier2{
		settings: &configuracoesDesenharBezier2{
			pontoA: &entradaPonto{Label: "ponto A"},
			pontoB: &entradaPonto{Label: "ponto B"},
			pontoC: &entradaPonto{Label: "ponto C"},
		},
	}
}

func (m *moduloDesenhaBezier2) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloDesenhaBezier2) Update() error {
	return nil
}

func (m *moduloDesenhaBezier2) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
