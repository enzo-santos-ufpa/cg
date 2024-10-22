package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"ufpa_cg"
)

type configuracoesDesenharPolilinha struct {
	pontos *entradaPontos
}

func (c *configuracoesDesenharPolilinha) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontos}
}

func (c *configuracoesDesenharPolilinha) Evaluate() []ufpa_cg.Ponto {
	pontos := make([]ufpa_cg.Ponto, 0)
	inps := c.pontos.entradas[:len(c.pontos.entradas)-1]
	for i := 0; i < len(inps); i++ {
		ponto1 := inps[i].ponto
		var ponto2 ufpa_cg.Ponto
		if i < len(inps)-1 {
			ponto2 = inps[i+1].ponto
		} else {
			ponto2 = inps[0].ponto
		}
		algoritmo := ufpa_cg.NewAlgoritmoBresenham(ponto1, ponto2)
		for algoritmo.Move() {
			pontos = append(pontos, algoritmo.PontoAtual())
		}
	}
	return pontos
}

type moduloDesenhaPolilinha struct {
	settings *configuracoesDesenharPolilinha
}

type opcaoDesenharPolilinha struct{}

func NewOpcaoDesenharPolilinha() OpcaoMenu {
	return &opcaoDesenharPolilinha{}
}

func (o *opcaoDesenharPolilinha) Title() string {
	return "Desenhar polilinha"
}

func (o *opcaoDesenharPolilinha) Create() ModuloJogo {
	return &moduloDesenhaPolilinha{
		settings: &configuracoesDesenharPolilinha{
			pontos: &entradaPontos{Minimo: 3, entradas: []*entradaPonto{new(entradaPonto)}},
		},
	}
}

func (m *moduloDesenhaPolilinha) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloDesenhaPolilinha) Update() error {
	return nil
}

func (m *moduloDesenhaPolilinha) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
