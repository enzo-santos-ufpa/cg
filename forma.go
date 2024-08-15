package ufpa_cg

import (
	"fmt"
	"math"
)

type OperacaoForma2D struct {
	Forma    Forma2D
	Matrizes []*Matriz
}

func (o OperacaoForma2D) Calcula() Forma2D {
	matrizes := o.Matrizes
	if len(matrizes) == 0 {
		return o.Forma
	}

	m := o.Matrizes[0]
	for _, m1 := range append(matrizes[1:], o.Forma.ToMatriz()) {
		m, _ = m.Multiplica(m1)
	}
	f, _ := NewForma2DFromMatriz(m)
	return f
}

type Forma2D struct {
	Pontos []Ponto
}

func NewForma2DFromMatriz(matriz *Matriz) (Forma2D, error) {
	if matriz.Dimensao.NumLinhas != 3 {
		return Forma2D{}, fmt.Errorf("incorrect number of lines: expected 3, got %d", matriz.Dimensao.NumLinhas)
	}
	pontos := make([]Ponto, matriz.Dimensao.NumColunas)
	for c := 0; c < matriz.Dimensao.NumColunas; c++ {
		z := matriz.Get(2, c)
		if z != 1 {
			return Forma2D{}, fmt.Errorf("incorrect value at (3, %d): expected 1, got %.2f", c+1, z)
		}
		pontos[c] = Ponto{
			X: int(math.Round(matriz.Get(0, c))),
			Y: int(math.Round(matriz.Get(1, c))),
		}
	}
	return Forma2D{Pontos: pontos}, nil
}

func (f Forma2D) ToMatriz() *Matriz {
	matriz := NewMatrizVazia(DimensaoMatriz{
		NumLinhas:  3,
		NumColunas: len(f.Pontos),
	})
	for i, ponto := range f.Pontos {
		matriz.Set(0, i, float64(ponto.X))
		matriz.Set(1, i, float64(ponto.Y))
		matriz.Set(2, i, 1)
	}
	return matriz
}

func (f Forma2D) IniciaOperacao() OperacaoForma2D {
	return OperacaoForma2D{
		Forma:    f,
		Matrizes: nil,
	}
}

func (o OperacaoForma2D) RotacionaMatricial(angulo float64) OperacaoForma2D {
	senA := math.Sin(angulo)
	cosA := math.Cos(angulo)
	m, _ := NewMatriz([][]float64{
		{cosA, -senA, 0},
		{senA, cosA, 0},
		{0, 0, 1},
	})
	return OperacaoForma2D{
		Forma:    o.Forma,
		Matrizes: append(o.Matrizes, m),
	}
}

func (o OperacaoForma2D) MoveMatricial(dx, dy int) OperacaoForma2D {
	m, _ := NewMatriz([][]float64{
		{1, 0, float64(dx)},
		{0, 1, float64(dy)},
		{0, 0, 1},
	})
	return OperacaoForma2D{
		Forma:    o.Forma,
		Matrizes: append(o.Matrizes, m),
	}
}

// TODO func RedimensionaMatricial(ex, ey float64)
