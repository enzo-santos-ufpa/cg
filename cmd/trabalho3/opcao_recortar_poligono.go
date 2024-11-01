package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"math"
	"ufpa_cg"
)

type configuracoesRecortarPoligono struct {
	pontos *entradaPontos
	janela *entradaJanela
}

func (c *configuracoesRecortarPoligono) Inputs() []EntradaModulo {
	return []EntradaModulo{c.pontos, c.janela}
}

type TipoAresta int

const (
	ArestaEsquerda TipoAresta = iota
	ArestaDireita  TipoAresta = iota
	ArestaInferior TipoAresta = iota
	ArestaSuperior TipoAresta = iota
)

type JanelaRecorte ufpa_cg.JanelaRecorte

func (j *JanelaRecorte) intersect(p1, p2 ufpa_cg.Ponto, aresta TipoAresta) ufpa_cg.Ponto {
	switch aresta {
	case ArestaEsquerda, ArestaDireita:
		var x int
		if aresta == ArestaEsquerda {
			x = j.PontoSuperiorEsquerdo.X
		} else {
			x = j.PontoInferiorDireito.X
		}
		return ufpa_cg.Ponto{
			X: x,
			Y: int(math.Round(float64(p1.Y) + float64(p2.Y-p1.Y)*float64(x-p1.X)/float64(p2.X-p1.X))),
		}
	case ArestaInferior, ArestaSuperior:
		var y int
		if aresta == ArestaInferior {
			y = j.PontoInferiorDireito.Y
		} else {
			y = j.PontoSuperiorEsquerdo.Y
		}
		return ufpa_cg.Ponto{
			X: int(math.Round(float64(p1.X) + float64(p2.X-p1.X)*float64(y-p1.Y)/float64(p2.Y-p1.Y))),
			Y: y,
		}
	}
	panic(fmt.Sprintf("unexpected value for `aresta`: %v", aresta))
}

func (j *JanelaRecorte) contains(ponto ufpa_cg.Ponto, aresta TipoAresta) bool {
	switch aresta {
	case ArestaEsquerda:
		return ponto.X >= j.PontoSuperiorEsquerdo.X
	case ArestaDireita:
		return ponto.X <= j.PontoInferiorDireito.X
	case ArestaInferior:
		return ponto.Y >= j.PontoInferiorDireito.Y
	case ArestaSuperior:
		return ponto.Y <= j.PontoSuperiorEsquerdo.Y
	}
	panic(fmt.Sprintf("unexpected value for `aresta`: %v", aresta))
}

func (j *JanelaRecorte) CalculaIntersecaoPoligono(vertices []ufpa_cg.Ponto) []ufpa_cg.Ponto {
	poligono := make([]ufpa_cg.Ponto, len(vertices))
	for i, vertice := range vertices {
		poligono[i] = vertice
	}

	for _, aresta := range []TipoAresta{ArestaEsquerda, ArestaDireita, ArestaInferior, ArestaSuperior} {
		subPoligono := make([]ufpa_cg.Ponto, 0)
		for i, ponto1 := range poligono {
			ponto2 := poligono[(i+1)%len(poligono)]
			if j.contains(ponto1, aresta) {
				if j.contains(ponto2, aresta) {
					subPoligono = append(subPoligono, ponto2)
				} else {
					subPoligono = append(subPoligono, j.intersect(ponto1, ponto2, aresta))
				}
			} else if j.contains(ponto2, aresta) {
				subPoligono = append(subPoligono, j.intersect(ponto1, ponto2, aresta), ponto2)
			}
		}
		poligono = make([]ufpa_cg.Ponto, len(subPoligono))
		for i, ponto := range subPoligono {
			poligono[i] = ponto
		}
	}
	return poligono
}

func (c *configuracoesRecortarPoligono) Evaluate() []ufpa_cg.Ponto {
	vertices := make([]ufpa_cg.Ponto, len(c.pontos.entradas)-1)
	for i, inp := range c.pontos.entradas[:len(c.pontos.entradas)-1] {
		vertices[i] = inp.ponto
	}
	janela := JanelaRecorte(c.janela.JanelaRecorte())
	vertices2 := janela.CalculaIntersecaoPoligono(vertices)

	borda := make([]ufpa_cg.Ponto, 0)
	for i := 0; i < len(vertices2); i++ {
		ponto1 := vertices2[i]
		var ponto2 ufpa_cg.Ponto
		if i < len(vertices2)-1 {
			ponto2 = vertices2[i+1]
		} else {
			ponto2 = vertices2[0]
		}
		algoritmo := ufpa_cg.NewAlgoritmoBresenham(ponto1, ponto2)
		for algoritmo.Move() {
			borda = append(borda, algoritmo.PontoAtual())
		}
	}
	return borda
}

type moduloRecortaPoligono struct {
	settings *configuracoesRecortarPoligono
}

type opcaoRecortarPoligono struct{}

func NewOpcaoRecortarPoligono() OpcaoMenu {
	return &opcaoRecortarPoligono{}
}

func (o *opcaoRecortarPoligono) Title() string {
	return "Recortar polígono"
}

func (o *opcaoRecortarPoligono) Create() ModuloJogo {
	return &moduloRecortaPoligono{
		settings: &configuracoesRecortarPoligono{
			pontos: &entradaPontos{Minimo: 3, entradas: []*entradaPonto{new(entradaPonto)}},
			janela: &entradaJanela{
				entradaPonto1: &entradaPonto{},
				entradaPonto2: &entradaPonto{},
			},
		},
	}
}

func (m *moduloRecortaPoligono) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloRecortaPoligono) Update() error {
	return nil
}

func (m *moduloRecortaPoligono) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
