package main

import (
	"cmp"
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"math"
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

type algoritmoPreenchimentoVarredura struct{}

func NewAlgoritmoPreenchimentoVarredura() AlgoritmoPreenchimento {
	return &algoritmoPreenchimentoVarredura{}
}

func (a *algoritmoPreenchimentoVarredura) Evaluate(vertices []ufpa_cg.Ponto, _ ufpa_cg.Ponto) []ufpa_cg.Ponto {
	type PontoCritico struct {
		Indice   int
		Direcao  int
		X        float64
		MInverso float64
	}

	miolo := make([]ufpa_cg.Ponto, 0)

	yMin := 999
	yMax := -999
	pontosCriticos := make([]PontoCritico, 0)
	for i, vertice := range vertices {
		if vertice.Y < yMin {
			yMin = vertice.Y
		}
		if yMax < vertice.Y {
			yMax = vertice.Y
		}
		vertice2 := vertices[(i+1)%len(vertices)]
		if vertice.Y < vertice2.Y {
			pontosCriticos = append(pontosCriticos, PontoCritico{
				Indice:   i,
				Direcao:  1,
				X:        float64(vertice.X),
				MInverso: float64(vertice2.X-vertice.X) / float64(vertice2.Y-vertice.Y),
			})
		}
		vertice0 := vertices[(i-1+len(vertices))%len(vertices)]
		if vertice.Y < vertice0.Y {
			pontosCriticos = append(pontosCriticos, PontoCritico{
				Indice:   i,
				Direcao:  -1,
				X:        float64(vertice.X),
				MInverso: float64(vertice0.X-vertice.X) / float64(vertice0.Y-vertice.Y),
			})
		}
	}

	pontosCriticosAtuais := make([]*PontoCritico, 0)
	for y := yMin; y <= yMax; y++ {
		for _, pontoCritico := range pontosCriticosAtuais {
			pontoCritico.X += pontoCritico.MInverso
		}
		for _, pontoCritico := range pontosCriticos {
			if vertices[pontoCritico.Indice].Y == y {
				pontosCriticosAtuais = append(pontosCriticosAtuais, &pontoCritico)
			}
		}

		novosPontosCriticosAtuais := make([]*PontoCritico, 0)
		for _, pontoCritico := range pontosCriticosAtuais {
			pontoMax := vertices[(pontoCritico.Indice+pontoCritico.Direcao+len(vertices))%len(vertices)]
			if pontoMax.Y == y {
				continue
			}
			novosPontosCriticosAtuais = append(novosPontosCriticosAtuais, pontoCritico)
		}
		pontosCriticosAtuais = novosPontosCriticosAtuais[:]
		slices.SortStableFunc(pontosCriticosAtuais, func(p0, p1 *PontoCritico) int {
			return cmp.Compare(p0.X, p1.X)
		})
		for i := 0; i < len(pontosCriticosAtuais); i += 2 {
			p0 := pontosCriticosAtuais[i]
			p1 := pontosCriticosAtuais[i+1]

			for x := int(math.Round(p0.X)); x <= int(math.Round(p1.X)); x++ {
				miolo = append(miolo, ufpa_cg.Ponto{X: x, Y: y})
			}
		}
	}
	return miolo
}

type configuracoesPreenchimento struct {
	Algoritmo AlgoritmoPreenchimento

	// TODO Valida se ponto selecionado está dentro da polilinha
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
