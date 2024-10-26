package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"slices"
	"ufpa_cg"
)

type AlgoritmoPreenchimento interface {
	Evaluate(vertices []ufpa_cg.Ponto, ponto ufpa_cg.Ponto) []ufpa_cg.Ponto
}

type algoritmoPreenchimentoRecursao struct{}

func NewAlgoritmoPreenchimentoRecursao() AlgoritmoPreenchimento {
	return &algoritmoPreenchimentoRecursao{}
}

func (a *algoritmoPreenchimentoRecursao) preenche(miolo *[]ufpa_cg.Ponto, borda []ufpa_cg.Ponto, ponto ufpa_cg.Ponto) {
	if slices.Contains(*miolo, ponto) || slices.Contains(borda, ponto) {
		return
	}
	*miolo = append(*miolo, ponto)
	a.preenche(miolo, borda, ufpa_cg.Ponto{X: ponto.X + 1, Y: ponto.Y})
	a.preenche(miolo, borda, ufpa_cg.Ponto{X: ponto.X, Y: ponto.Y + 1})
	a.preenche(miolo, borda, ufpa_cg.Ponto{X: ponto.X - 1, Y: ponto.Y})
	a.preenche(miolo, borda, ufpa_cg.Ponto{X: ponto.X, Y: ponto.Y - 1})
}

func (a *algoritmoPreenchimentoRecursao) Evaluate(vertices []ufpa_cg.Ponto, ponto ufpa_cg.Ponto) []ufpa_cg.Ponto {
	borda := make([]ufpa_cg.Ponto, 0)
	for i := 0; i < len(vertices); i++ {
		ponto1 := vertices[i]
		var ponto2 ufpa_cg.Ponto
		if i < len(vertices)-1 {
			ponto2 = vertices[i+1]
		} else {
			ponto2 = vertices[0]
		}
		algoritmo := ufpa_cg.NewAlgoritmoBresenham(ponto1, ponto2)
		for algoritmo.Move() {
			borda = append(borda, algoritmo.PontoAtual())
		}
	}
	miolo := make([]ufpa_cg.Ponto, 0)
	a.preenche(&miolo, borda, ponto)
	return miolo
}

type configuracoesPreenchimento struct {
	Algoritmo AlgoritmoPreenchimento

	pontos *entradaPontos
	centro *entradaPonto
}

func (c *configuracoesPreenchimento) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontos, c.centro}
}

func (c *configuracoesPreenchimento) Evaluate() []ufpa_cg.Ponto {
	inps := c.pontos.entradas[:len(c.pontos.entradas)-1]
	ponto := c.centro.ponto

	vertices := make([]ufpa_cg.Ponto, len(inps))
	for i := 0; i < len(inps); i++ {
		vertices[i] = inps[i].ponto
	}
	return c.Algoritmo.Evaluate(vertices, ponto)
}

type moduloPreenchimento struct {
	settings *configuracoesPreenchimento
}

type opcaoPreenchimento struct {
	Algoritmo AlgoritmoPreenchimento
	Label     string
}

func NewOpcaoPreenchimento(label string, algoritmo AlgoritmoPreenchimento) OpcaoMenu {
	return &opcaoPreenchimento{
		Algoritmo: algoritmo,
		Label:     label,
	}
}

func (o *opcaoPreenchimento) Title() string {
	return o.Label
}

func (o *opcaoPreenchimento) Create() ModuloJogo {
	return &moduloPreenchimento{
		settings: &configuracoesPreenchimento{
			Algoritmo: o.Algoritmo,
			pontos:    &entradaPontos{Minimo: 3, entradas: []*entradaPonto{new(entradaPonto)}},
			centro:    &entradaPonto{Label: "ponto interno"},
		},
	}
}

func (m *moduloPreenchimento) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloPreenchimento) Update() error {
	return nil
}

func (m *moduloPreenchimento) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
