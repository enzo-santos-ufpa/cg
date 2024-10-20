package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type opcaoVazia struct {
	Label string
}

func NewOpcaoVazia(label string) OpcaoMenu {
	return &opcaoVazia{Label: label}
}

func (o *opcaoVazia) Title() string {
	return o.Label
}

func (o *opcaoVazia) Create() ModuloJogo {
	return &moduloVazio{}
}

type configuracaoModuloVazio struct {
}

func (c configuracaoModuloVazio) Inputs() []EntradaModulo {
	return []EntradaModulo{}
}

func (c configuracaoModuloVazio) Evaluate() []ufpa_cg.Ponto {
	return []ufpa_cg.Ponto{}
}

type moduloVazio struct{}

func (m *moduloVazio) Settings() ConfiguracaoModulo {
	return &configuracaoModuloVazio{}
}

func (m *moduloVazio) Update() error {
	return nil
}

func (m *moduloVazio) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
