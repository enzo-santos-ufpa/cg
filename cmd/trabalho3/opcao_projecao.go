package main

import (
	"cmp"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"math"
	"slices"
	"ufpa_cg"
)

type AlgoritmoProjecao interface {
	EvaluateInputs() []EntradaModulo
	Transform(vertices []ufpa_cg.Ponto3D) []ufpa_cg.Ponto
}

type algoritmoProjecaoOrtogonal struct {
	entradaEixo *entradaItens[ufpa_cg.Eixo3D]
}

func NewAlgoritmoProjecaoOrtogonal() AlgoritmoProjecao {
	return &algoritmoProjecaoOrtogonal{
		entradaEixo: &entradaItens[ufpa_cg.Eixo3D]{
			Label: "plano de projeção",
			Itens: []ufpa_cg.Eixo3D{ufpa_cg.EixoX, ufpa_cg.EixoY, ufpa_cg.EixoZ},
			Labeler: func(value ufpa_cg.Eixo3D) string {
				switch value {
				case ufpa_cg.EixoX:
					return "Plano YZ"
				case ufpa_cg.EixoY:
					return "Plano XZ"
				case ufpa_cg.EixoZ:
					return "Plano XY"
				}
				panic(fmt.Sprintf("invalid ufpa_cg.Eixo3D value: %v", value))
			},
		},
	}
}

func (a *algoritmoProjecaoOrtogonal) EvaluateInputs() []EntradaModulo {
	return []EntradaModulo{a.entradaEixo}
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

type algoritmoProjecaoObliqua struct {
	entradaAnguloInclinacao *entradaInteiro
	entradaFatorCompressao  *entradaInteiro
}

func NewAlgoritmoProjecaoObliqua() AlgoritmoProjecao {
	return &algoritmoProjecaoObliqua{
		entradaAnguloInclinacao: &entradaInteiro{
			Label:        "angulo de inclinação",
			Sufixo:       "°",
			PossuiMinimo: true,
			Minimo:       -360,
			PossuiMaximo: true,
			Maximo:       360,
		},
		entradaFatorCompressao: &entradaInteiro{
			Label:        "fator de compressão",
			Sufixo:       "%",
			PossuiMinimo: true,
			Minimo:       0,
			PossuiMaximo: true,
			Maximo:       100,
		},
	}
}

func (a *algoritmoProjecaoObliqua) EvaluateInputs() []EntradaModulo {
	return []EntradaModulo{a.entradaAnguloInclinacao, a.entradaFatorCompressao}
}

func (a *algoritmoProjecaoObliqua) Transform(vertices []ufpa_cg.Ponto3D) []ufpa_cg.Ponto {
	anguloInclinacao := float64(a.entradaAnguloInclinacao.valor) * math.Pi / 180.0
	fatorCompressao := float64(a.entradaFatorCompressao.valor) / 100.0

	vertices2D := make([]ufpa_cg.Ponto, len(vertices))
	for i, vertice := range vertices {
		vertices2D[i] = ufpa_cg.Ponto{
			X: int(math.Round(float64(vertice.X) + fatorCompressao*float64(vertice.Z)*math.Cos(anguloInclinacao))),
			Y: int(math.Round(float64(vertice.Y) + fatorCompressao*float64(vertice.Z)*math.Sin(anguloInclinacao))),
		}
	}
	return vertices2D
}

type algoritmoProjecaoPerspectiva struct {
	entradaCentro *entradaPonto3D
	entradaPlanoZ *entradaInteiro
}

func NewAlgoritmoProjecaoPerspectiva() AlgoritmoProjecao {
	return &algoritmoProjecaoPerspectiva{
		entradaCentro: &entradaPonto3D{
			Label:    "centro de projeção",
			entradaX: &entradaInteiro{},
			entradaY: &entradaInteiro{},
			entradaZ: &entradaInteiro{},
		},
		entradaPlanoZ: &entradaInteiro{Label: "plano de projeção"},
	}
}

func (a *algoritmoProjecaoPerspectiva) EvaluateInputs() []EntradaModulo {
	return []EntradaModulo{a.entradaCentro, a.entradaPlanoZ}
}

func (a *algoritmoProjecaoPerspectiva) Transform(vertices []ufpa_cg.Ponto3D) []ufpa_cg.Ponto {
	centro := a.entradaCentro.ponto
	distanciaZ := a.entradaPlanoZ.valor

	vertices2D := make([]ufpa_cg.Ponto, len(vertices))
	for i, vertice := range vertices {
		dx := vertice.X - centro.X
		dy := vertice.Y - centro.Y
		dz := vertice.Z - centro.Z
		if dz == 0 {
			return []ufpa_cg.Ponto{}
		}
		fator := float64(distanciaZ) / float64(dz)
		vertices2D[i] = ufpa_cg.Ponto{
			X: int(math.Round(float64(centro.X) + float64(dx)*fator)),
			Y: int(math.Round(float64(centro.Y) + float64(dy)*fator)),
		}
	}
	return vertices2D
}

type configuracoesProjetarPoligono3D struct {
	Algoritmo AlgoritmoProjecao

	entrada *entradaMultipla[*entradaPonto3D]
}

func (c *configuracoesProjetarPoligono3D) Inputs() []EntradaModulo {
	return append([]EntradaModulo{c.entrada}, c.Algoritmo.EvaluateInputs()...)
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
