package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
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

type moduloVazio struct{}

func (m *moduloVazio) Update() error {
	return nil
}

func (m *moduloVazio) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
