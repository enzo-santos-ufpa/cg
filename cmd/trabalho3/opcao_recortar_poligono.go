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

func (j *JanelaRecorte) intersectionPoint(p1, p2 ufpa_cg.Ponto, aresta TipoAresta) ufpa_cg.Ponto {
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
			y = j.PontoSuperiorEsquerdo.Y
		} else {
			y = j.PontoInferiorDireito.Y
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

// TODO Make it work

func (j *JanelaRecorte) CalculaIntersecaoPoligono(vertices []ufpa_cg.Ponto) []ufpa_cg.Ponto {

	pontos := make([]ufpa_cg.Ponto, len(vertices))
	for i, vertice := range vertices {
		pontos[i] = vertice
	}

	for _, aresta := range []TipoAresta{ArestaEsquerda, ArestaDireita, ArestaInferior, ArestaSuperior} {
		pontosAresta := make([]ufpa_cg.Ponto, 0)
		p1 := pontos[len(pontos)-1]
		for _, p2 := range pontos {
			if j.contains(p2, aresta) {
				if !j.contains(p1, aresta) {
					pontosAresta = append(pontosAresta, j.intersectionPoint(p1, p2, aresta))
				}
				pontosAresta = append(pontosAresta, p2)
			} else if j.contains(p1, aresta) {
				pontosAresta = append(pontosAresta, j.intersectionPoint(p1, p2, aresta))
			}
			p1 = p2
		}

		pontos = make([]ufpa_cg.Ponto, len(pontosAresta))
		for i, ponto := range pontosAresta {
			pontos[i] = ponto
		}
	}
	return pontos
}

func (c *configuracoesRecortarPoligono) Evaluate() []ufpa_cg.Ponto {
	vertices := make([]ufpa_cg.Ponto, len(c.pontos.entradas)-1)
	for i, inp := range c.pontos.entradas[:len(c.pontos.entradas)-1] {
		vertices[i] = inp.ponto
	}
	janela := JanelaRecorte(ufpa_cg.JanelaRecorte{
		PontoSuperiorEsquerdo: c.janela.entradaPonto1.ponto,
		PontoInferiorDireito:  c.janela.entradaPonto2.ponto,
	})
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
