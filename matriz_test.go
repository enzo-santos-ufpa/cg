package ufpa_cg

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGrausParaRadianos(t *testing.T) {
	const delta = 0.00001
	require.InDelta(t, GrausParaRadianos(0), 0, delta)
	require.InDelta(t, GrausParaRadianos(45), math.Pi/4, delta)
	require.InDelta(t, GrausParaRadianos(90), math.Pi/2, delta)
	require.InDelta(t, GrausParaRadianos(135), 3*math.Pi/4, delta)
	require.InDelta(t, GrausParaRadianos(180), math.Pi, delta)
	require.InDelta(t, GrausParaRadianos(360), 2*math.Pi, delta)
}

func TestMatriz_NewMatriz_Linha(t *testing.T) {
	m, err := NewMatriz([][]float64{{1}, {2}, {3}})
	require.Nil(t, err)
	require.NotNil(t, m)
	require.Equal(t, DimensaoMatriz{NumLinhas: 3, NumColunas: 1}, m.Dimensao)
}

func TestMatriz_NewMatriz_Coluna(t *testing.T) {
	m, err := NewMatriz([][]float64{{1, 2, 3}})
	require.Nil(t, err)
	require.NotNil(t, m)
	require.Equal(t, DimensaoMatriz{NumLinhas: 1, NumColunas: 3}, m.Dimensao)
}

func TestMatriz_NewMatriz_Quadrada(t *testing.T) {
	m, err := NewMatriz([][]float64{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})
	require.Nil(t, err)
	require.NotNil(t, m)
	require.Equal(t, DimensaoMatriz{NumLinhas: 3, NumColunas: 3}, m.Dimensao)
}

func TestMatriz_NewMatriz_LinhaSublinhasVazia(t *testing.T) {
	m, err := NewMatriz([][]float64{{}})
	require.Nil(t, m)
	require.NotNil(t, err)
	require.Equal(t, "given array cannot be converted to a matrix: expected row at index 0 to be non-empty", err.Error())
}

func TestMatriz_NewMatriz_ColunaSublinhasVazia(t *testing.T) {
	m, err := NewMatriz([][]float64{{}, {}, {}})
	require.Nil(t, m)
	require.NotNil(t, err)
	require.Equal(t, "given array cannot be converted to a matrix: expected row at index 0 to be non-empty", err.Error())
}

func TestMatriz_NewMatriz_SublinhasIncompativel(t *testing.T) {
	m, err := NewMatriz([][]float64{
		{1, 11, 3},
		{13, 5},
		{7, 17, 9},
	})
	require.Nil(t, m)
	require.NotNil(t, err)
	require.Equal(t, "given array cannot be converted to a matrix: expected [3 2 3] to be all equal", err.Error())
}

func TestMatriz_Soma(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 11, 3},
		{13, 5, 15},
		{7, 17, 9},
	})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{10, 2, 12},
		{4, 14, 6},
		{16, 8, 18},
	})
	require.Nil(t, err)

	m3Atual, err := m1.Soma(m2)
	require.Nil(t, err)

	m3Esperada, err := NewMatriz([][]float64{
		{11, 13, 15},
		{17, 19, 21},
		{23, 25, 27},
	})
	require.Nil(t, err)

	require.Equal(t, m3Esperada, m3Atual)
}

func TestMatriz_Soma_Incompativel(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 11, 3},
		{13, 5, 15},
		{7, 17, 9},
	})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{10, 2},
		{4, 14},
	})
	require.Nil(t, err)

	m3, err := m1.Soma(m2)
	require.Nil(t, m3)
	require.NotNil(t, err)
	require.Equal(t, "matrices are incompatible for matrix addition: expected (3, 3) + (2, 2) to be equal", err.Error())
}

func TestMatriz_MultiplicaEscalar(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{10, 12, 14},
		{16, 18, 20},
		{22, 24, 26},
	})
	require.Nil(t, err)

	m2Atual := m1.MultiplicaEscalar(0.5)
	m2Esperada, err := NewMatriz([][]float64{
		{5, 6, 7},
		{8, 9, 10},
		{11, 12, 13},
	})
	require.Nil(t, err)
	require.Equal(t, m2Esperada, m2Atual)
}

func TestMatriz_Transposta_Linha(t *testing.T) {
	m1, err := NewMatriz([][]float64{{1}, {2}, {3}})
	require.Nil(t, err)

	m2Atual := m1.Transposta()
	m2Esperada, err := NewMatriz([][]float64{{1, 2, 3}})
	require.Nil(t, err)
	require.Equal(t, m2Esperada, m2Atual)
}

func TestMatriz_Transposta_Coluna(t *testing.T) {
	m1, err := NewMatriz([][]float64{{1, 2, 3}})
	require.Nil(t, err)

	m2Atual := m1.Transposta()
	m2Esperada, err := NewMatriz([][]float64{{1}, {2}, {3}})
	require.Nil(t, err)
	require.Equal(t, m2Esperada, m2Atual)
}

func TestMatriz_Transposta_Quadrada(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	})
	require.Nil(t, err)

	m2Atual := m1.Transposta()
	m2Esperada, err := NewMatriz([][]float64{
		{1, 4, 7},
		{2, 5, 8},
		{3, 6, 9},
	})
	require.Nil(t, err)
	require.Equal(t, m2Esperada, m2Atual)
}

func TestMatriz_Transposta(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 2},
		{3, 4},
		{5, 6},
	})
	require.Nil(t, err)

	m2Atual := m1.Transposta()
	m2Esperada, err := NewMatriz([][]float64{
		{1, 3, 5},
		{2, 4, 6},
	})
	require.Nil(t, err)
	require.Equal(t, m2Esperada, m2Atual)
}

func TestMatriz_Multiplica_Quadrada(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 10, 3},
		{12, 5, 14},
		{7, 16, 9},
	})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{11, 2, 13},
		{4, 15, 6},
		{17, 8, 19},
	})
	require.Nil(t, err)

	m3Atual, err := m1.Multiplica(m2)
	require.Nil(t, err)

	m3Esperada, err := NewMatriz([][]float64{
		{102, 176, 130},
		{390, 211, 452},
		{294, 326, 358},
	})
	require.Nil(t, err)

	require.Equal(t, m3Esperada, m3Atual)
}

func TestMatriz_Multiplica_ResultadoQuadrada(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 10},
		{12, 3},
		{5, 14},
	})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{11, 2, 13},
		{4, 15, 6},
	})
	require.Nil(t, err)

	m3Atual, err := m1.Multiplica(m2)
	require.Nil(t, err)

	m3Esperada, err := NewMatriz([][]float64{
		{51, 152, 73},
		{144, 69, 174},
		{111, 220, 149},
	})
	require.Nil(t, err)

	require.Equal(t, m3Esperada, m3Atual)
}

func TestMatriz_Multiplica_ResultadoLinha(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 2},
		{4, 3},
		{5, 6},
	})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{8},
		{7},
	})
	require.Nil(t, err)

	m3Atual, err := m1.Multiplica(m2)
	require.Nil(t, err)

	m3Esperada, err := NewMatriz([][]float64{
		{22},
		{53},
		{82},
	})
	require.Nil(t, err)

	require.Equal(t, m3Esperada, m3Atual)
}

func TestMatriz_Multiplica_ResultadoColuna(t *testing.T) {
	m1, err := NewMatriz([][]float64{{1, 3}})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{2, 5, 6},
		{7, 4, 9},
	})
	require.Nil(t, err)

	m3Atual, err := m1.Multiplica(m2)
	require.Nil(t, err)

	m3Esperada, err := NewMatriz([][]float64{{23, 17, 33}})
	require.Nil(t, err)

	require.Equal(t, m3Esperada, m3Atual)
}

func TestMatriz_Multiplica(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{5, 15, 6, 16},
		{17, 7, 18, 8},
		{9, 19, 10, 20},
	})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{1, 11},
		{12, 2},
		{3, 13},
		{14, 4},
	})
	require.Nil(t, err)

	m3Atual, err := m1.Multiplica(m2)
	require.Nil(t, err)

	m3Esperada, err := NewMatriz([][]float64{
		{427, 227},
		{267, 467},
		{547, 347},
	})
	require.Nil(t, err)

	require.Equal(t, m3Esperada, m3Atual)
}

func TestMatriz_Multiplica_Incompativel(t *testing.T) {
	m1, err := NewMatriz([][]float64{
		{1, 2, 3},
		{4, 5, 6},
	})
	require.Nil(t, err)

	m2, err := NewMatriz([][]float64{
		{4},
		{6},
	})
	require.Nil(t, err)

	m3, err := m1.Multiplica(m2)
	require.Nil(t, m3)
	require.NotNil(t, err)
	require.Equal(t, "matrices are incompatible for matrix multiplication: expected (_, 3) x (2, _) to be equal", err.Error())
}
