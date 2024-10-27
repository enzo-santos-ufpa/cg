package main

import (
	"cmp"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"math"
	"slices"
	"ufpa_cg"
)

type AlgoritmoProjecao interface {
	Transform(vertices []ufpa_cg.Ponto3D) []ufpa_cg.Ponto
}

type algoritmoProjecaoOrtogonal struct{}

func NewAlgoritmoProjecaoOrtogonal() AlgoritmoProjecao {
	return &algoritmoProjecaoOrtogonal{}
}

func (a *algoritmoProjecaoOrtogonal) Transform(vertices []ufpa_cg.Ponto3D) []ufpa_cg.Ponto {
	const planoProjecao ufpa_cg.Eixo3D = ufpa_cg.EixoX

	vertices2D := make([]ufpa_cg.Ponto, len(vertices))
	for _, vertice := range vertices {
		var vertice2D ufpa_cg.Ponto
		switch planoProjecao {
		case ufpa_cg.EixoX:
			vertice2D = ufpa_cg.Ponto{X: vertice.Y, Y: vertice.Z}
		case ufpa_cg.EixoY:
			vertice2D = ufpa_cg.Ponto{X: vertice.X, Y: vertice.Z}
		case ufpa_cg.EixoZ:
			vertice2D = ufpa_cg.Ponto{X: vertice.X, Y: vertice.Y}
		}
		vertices2D = append(vertices2D, vertice2D)
	}
	return vertices2D
}

type algoritmoProjecaoObliqua struct{}

func NewAlgoritmoProjecaoObliqua() AlgoritmoProjecao {
	return &algoritmoProjecaoObliqua{}
}

func (a *algoritmoProjecaoObliqua) Transform(vertices []ufpa_cg.Ponto3D) []ufpa_cg.Ponto {
	return []ufpa_cg.Ponto{}
}

type algoritmoProjecaoPerspectiva struct {
	entradaEx *entradaInteiro
	entradaEy *entradaInteiro
}

func NewAlgoritmoProjecaoPerspectiva() AlgoritmoProjecao {
	return &algoritmoProjecaoPerspectiva{}
}

func (a *algoritmoProjecaoPerspectiva) Transform(vertices []ufpa_cg.Ponto3D) []ufpa_cg.Ponto {
	return []ufpa_cg.Ponto{}
}

type configuracoesProjetarPoligono3D struct {
	Algoritmo AlgoritmoProjecao

	entrada *entradaMultipla[*entradaPonto3D]
}

func (c *configuracoesProjetarPoligono3D) Inputs() []EntradaModulo {
	return []EntradaModulo{c.entrada}
}

func (c *configuracoesProjetarPoligono3D) Evaluate() []ufpa_cg.Ponto {
	vertices3D := make([]ufpa_cg.Ponto3D, len(c.entrada.entradas)-1)
	for i, inp := range c.entrada.entradas[:len(c.entrada.entradas)-1] {
		vertices3D[i] = inp.ponto
	}
	vertices2D := c.Algoritmo.Transform(vertices3D)

	// Ordena vértices em sentido horário
	somaX, somaY := 0, 0
	for _, vertice2D := range vertices2D {
		somaX += vertice2D.X
		somaY += vertice2D.Y
	}
	centroX, centroY := float64(somaX)/float64(len(vertices2D)), float64(somaY)/float64(len(vertices2D))
	slices.SortFunc(vertices2D, func(p0, p1 ufpa_cg.Ponto) int {
		return cmp.Compare(
			math.Atan2(float64(p0.Y)-centroY, float64(p0.X)-centroX),
			math.Atan2(float64(p1.Y)-centroY, float64(p1.X)-centroX),
		)
	})
	borda := make([]ufpa_cg.Ponto, 0)
	for i, ponto1 := range vertices2D {
		ponto2 := vertices2D[(i+1)%len(vertices2D)]
		algoritmo := ufpa_cg.NewAlgoritmoBresenham(ponto1, ponto2)
		for algoritmo.Move() {
			borda = append(borda, algoritmo.PontoAtual())
		}
	}
	return borda
}

type moduloProjetaPoligono3D struct {
	settings *configuracoesProjetarPoligono3D
}

type opcaoProjetarPoligono3D struct {
	Label     string
	Algoritmo AlgoritmoProjecao
}

func NewOpcaoProjetarPoligono3D(label string, algoritmo AlgoritmoProjecao) OpcaoMenu {
	return &opcaoProjetarPoligono3D{
		Label:     label,
		Algoritmo: algoritmo,
	}
}

func (o *opcaoProjetarPoligono3D) Title() string {
	return o.Label
}

func (o *opcaoProjetarPoligono3D) Create() ModuloJogo {
	return &moduloProjetaPoligono3D{
		settings: &configuracoesProjetarPoligono3D{
			Algoritmo: o.Algoritmo,
			entrada: &entradaMultipla[*entradaPonto3D]{
				Minimo: 4,
				Prompt: "Selecione o ponto:",
				OnEvaluated: func() (map[ufpa_cg.Ponto]color.Color, bool) {
					return nil, false
				},
				Create: func() *entradaPonto3D {
					return &entradaPonto3D{
						entradaX: &entradaInteiro{},
						entradaY: &entradaInteiro{},
						entradaZ: &entradaInteiro{},
					}
				},
				entradas: []*entradaPonto3D{{
					entradaX: &entradaInteiro{},
					entradaY: &entradaInteiro{},
					entradaZ: &entradaInteiro{},
				}},
			},
		},
	}
}

func (m *moduloProjetaPoligono3D) Settings() ConfiguracaoModulo {
	return m.settings
}

func (m *moduloProjetaPoligono3D) Update() error {
	return nil
}

func (m *moduloProjetaPoligono3D) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}
