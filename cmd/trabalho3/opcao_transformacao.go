package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"math"
	"ufpa_cg"
)

type AlgoritmoTransformacao interface {
	EvaluateInputs() []EntradaModulo
	Transform(ponto ufpa_cg.Ponto) ufpa_cg.Ponto
}

type algoritmoRotacao struct {
	entradaAngulo *entradaInteiro
}

func NewAlgoritmoTransformacaoRotacao() AlgoritmoTransformacao {
	return &algoritmoRotacao{
		entradaAngulo: &entradaInteiro{
			Label:        "ângulo (º)",
			PossuiMinimo: true,
			Minimo:       -360,
			PossuiMaximo: true,
			Maximo:       360,
		},
	}
}

func (a *algoritmoRotacao) EvaluateInputs() []EntradaModulo {
	return []EntradaModulo{a.entradaAngulo}
}

func (a *algoritmoRotacao) Transform(ponto ufpa_cg.Ponto) ufpa_cg.Ponto {
	angulo := a.entradaAngulo.valor
	return ponto.RotacionaMatricial(float64(angulo) * (math.Pi / 180))
}

type algoritmoTranslacao struct {
	entradaDx *entradaInteiro
	entradaDy *entradaInteiro
}

func NewAlgoritmoTransformacaoTranslacao() AlgoritmoTransformacao {
	return &algoritmoTranslacao{
		entradaDx: &entradaInteiro{
			Label: "deslocamento X",
		},
		entradaDy: &entradaInteiro{
			Label: "deslocamento Y",
		},
	}
}

func (a *algoritmoTranslacao) EvaluateInputs() []EntradaModulo {
	return []EntradaModulo{a.entradaDx, a.entradaDy}
}
func (a *algoritmoTranslacao) Transform(ponto ufpa_cg.Ponto) ufpa_cg.Ponto {
	dx := a.entradaDx.valor
	dy := a.entradaDy.valor
	return ponto.MoveMatricial(dx, dy)
}

type algoritmoEscala struct {
	entradaEx *entradaInteiro
	entradaEy *entradaInteiro
}

func NewAlgoritmoTransformacaoEscala() AlgoritmoTransformacao {
	return &algoritmoEscala{
		entradaEx: &entradaInteiro{
			Label: "fator X",
		},
		entradaEy: &entradaInteiro{
			Label: "fator Y",
		},
	}
}

func (a *algoritmoEscala) EvaluateInputs() []EntradaModulo {
	return []EntradaModulo{a.entradaEx, a.entradaEy}
}

func (a *algoritmoEscala) Transform(ponto ufpa_cg.Ponto) ufpa_cg.Ponto {
	ex := a.entradaEx.valor
	ey := a.entradaEy.valor
	return ponto.RedimensionaMatricial(float64(ex), float64(ey))
}

type configuracoesTransformarPoligono struct {
	Algoritmo AlgoritmoTransformacao

	pontos *entradaPontos
}

func (c *configuracoesTransformarPoligono) Inputs() []EntradaModulo {
	return append([]EntradaModulo{c.pontos}, c.Algoritmo.EvaluateInputs()...)
}

func (c *configuracoesTransformarPoligono) Evaluate() []ufpa_cg.Ponto {
	pontos := make([]ufpa_cg.Ponto, 0)
	vertices := make([]ufpa_cg.Ponto, len(c.pontos.entradas)-1)
	inps := c.pontos.entradas[:len(c.pontos.entradas)-1]
	for i := 0; i < len(inps); i++ {
		vertices[i] = c.Algoritmo.Transform(inps[i].ponto)
	}
	for i, ponto1 := range vertices {
		var ponto2 ufpa_cg.Ponto
		if i < len(vertices)-1 {
			ponto2 = vertices[i+1]
		} else {
			ponto2 = vertices[0]
		}
		algoritmo := ufpa_cg.NewAlgoritmoBresenham(ponto1, ponto2)
		for algoritmo.Move() {
			pontos = append(pontos, algoritmo.PontoAtual())
		}
	}
	return pontos
}

type moduloTransformaPoligono struct {
	settings *configuracoesTransformarPoligono
}

type opcaoTransformarPoligono struct {
	Label     string
	Algoritmo AlgoritmoTransformacao
}

func NewOpcaoTransformarPoligono(label string, algoritmo AlgoritmoTransformacao) OpcaoMenu {
	return &opcaoTransformarPoligono{
		Label:     label,
		Algoritmo: algoritmo,
	}
}

func (o *opcaoTransformarPoligono) Title() string {
	return o.Label
}

func (o *opcaoTransformarPoligono) Create() ModuloJogo {
	return &moduloTransformaPoligono{
		settings: &configuracoesTransformarPoligono{
			Algoritmo: o.Algoritmo,
			pontos:    &entradaPontos{Minimo: 3, entradas: []*entradaPonto{new(entradaPonto)}},
		},
	}
}

func (m *moduloTransformaPoligono) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloTransformaPoligono) Update() error {
	return nil
}

func (m *moduloTransformaPoligono) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
