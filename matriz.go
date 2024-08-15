package ufpa_cg

import (
	"fmt"
	"math"
)

func GrausParaRadianos(a float64) float64 {
	return a * math.Pi / 180.0
}

type DimensaoMatriz struct {
	NumLinhas  int
	NumColunas int
}

type Matriz struct {
	Dimensao DimensaoMatriz

	valor [][]float64
}

func NewMatrizVazia(dimensao DimensaoMatriz) *Matriz {
	valor := make([][]float64, dimensao.NumLinhas)
	for l := 0; l < dimensao.NumLinhas; l++ {
		valor[l] = make([]float64, dimensao.NumColunas)
	}
	return &Matriz{
		Dimensao: dimensao,
		valor:    valor,
	}
}

func NewMatriz(valor [][]float64) (*Matriz, error) {
	numsCols := make([]int, len(valor))
	numsColsOk := true
	for i, linha := range valor {
		numCols := len(linha)
		if numCols == 0 {
			return nil, fmt.Errorf("given array cannot be converted to a matrix: expected row at index %d to be non-empty", i)
		}
		numsCols[i] = numCols
		if numsColsOk && i > 0 && numCols != numsCols[0] {
			numsColsOk = false
		}
	}
	if !numsColsOk {
		return nil, fmt.Errorf("given array cannot be converted to a matrix: expected %v to be all equal", numsCols)
	}

	return &Matriz{
		Dimensao: DimensaoMatriz{
			NumLinhas:  len(valor),
			NumColunas: numsCols[0],
		},
		valor: valor,
	}, nil
}

func (m *Matriz) Get(l, c int) float64 {
	return m.valor[l][c]
}

func (m *Matriz) Set(l, c int, valor float64) {
	m.valor[l][c] = valor
}

func (m *Matriz) Soma(m2 *Matriz) (*Matriz, error) {
	if m.Dimensao != m2.Dimensao {
		return nil, fmt.Errorf(
			"matrices are incompatible for matrix addition: expected (%d, %d) + (%d, %d) to be equal",
			m.Dimensao.NumLinhas,
			m.Dimensao.NumColunas,
			m2.Dimensao.NumLinhas,
			m2.Dimensao.NumColunas,
		)
	}
	resultado := NewMatrizVazia(m.Dimensao)
	for l := 0; l < m.Dimensao.NumLinhas; l++ {
		for c := 0; c < m.Dimensao.NumColunas; c++ {
			resultado.Set(l, c, m.Get(l, c)+m2.Get(l, c))
		}
	}
	return resultado, nil
}

func (m *Matriz) MultiplicaEscalar(valor float64) *Matriz {
	resultado := NewMatrizVazia(m.Dimensao)
	for l := 0; l < m.Dimensao.NumLinhas; l++ {
		for c := 0; c < m.Dimensao.NumColunas; c++ {
			resultado.Set(l, c, m.Get(l, c)*valor)
		}
	}
	return resultado
}

func (m *Matriz) Transposta() *Matriz {
	resultado := NewMatrizVazia(DimensaoMatriz{
		NumLinhas:  m.Dimensao.NumColunas,
		NumColunas: m.Dimensao.NumLinhas,
	})
	for l := 0; l < m.Dimensao.NumLinhas; l++ {
		for c := 0; c < m.Dimensao.NumColunas; c++ {
			resultado.Set(c, l, m.Get(l, c))
		}
	}
	return resultado
}

func (m *Matriz) Multiplica(m2 *Matriz) (*Matriz, error) {
	if m.Dimensao.NumColunas != m2.Dimensao.NumLinhas {
		return nil, fmt.Errorf(
			"matrices are incompatible for matrix multiplication: expected (_, %d) x (%d, _) to be equal",
			m.Dimensao.NumColunas,
			m2.Dimensao.NumLinhas,
		)
	}
	resultado := NewMatrizVazia(DimensaoMatriz{
		NumLinhas:  m.Dimensao.NumLinhas,
		NumColunas: m2.Dimensao.NumColunas,
	})
	for l := 0; l < resultado.Dimensao.NumLinhas; l++ {
		for c := 0; c < resultado.Dimensao.NumColunas; c++ {
			var soma float64
			for i := 0; i < m.Dimensao.NumColunas; i++ {
				soma += m.Get(l, i) * m2.Get(i, c)
			}
			resultado.Set(l, c, soma)
		}
	}
	return resultado, nil
}
