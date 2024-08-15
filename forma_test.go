package ufpa_cg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewForma2DFromMatriz(t *testing.T) {
	matriz, _ := NewMatriz([][]float64{
		{0, 2, 2, 0},
		{0, 0, 2, 2},
		{1, 1, 1, 1},
	})
	forma, err := NewForma2DFromMatriz(matriz)
	require.NoError(t, err)
	require.Len(t, forma.Pontos, 4)
	require.Equal(t, Ponto{X: 0, Y: 0}, forma.Pontos[0])
	require.Equal(t, Ponto{X: 2, Y: 0}, forma.Pontos[1])
	require.Equal(t, Ponto{X: 2, Y: 2}, forma.Pontos[2])
	require.Equal(t, Ponto{X: 0, Y: 2}, forma.Pontos[3])
}

func TestNewForma2DFromMatriz_NumLinhasInvalido(t *testing.T) {
	matriz, _ := NewMatriz([][]float64{
		{0, 2, 2, 0},
		{0, 0, 2, 2},
	})
	_, err := NewForma2DFromMatriz(matriz)
	require.EqualError(t, err, "incorrect number of lines: expected 3, got 2")
}

func TestNewForma2DFromMatriz_Linha3Invalida(t *testing.T) {
	matriz, _ := NewMatriz([][]float64{
		{0, 2, 2, 0},
		{0, 0, 2, 2},
		{1, 0, 1, 1},
	})
	_, err := NewForma2DFromMatriz(matriz)
	require.EqualError(t, err, "incorrect value at (3, 2): expected 1, got 0.00")
}

func TestForma2D_ToMatriz(t *testing.T) {
	forma := Forma2D{
		Pontos: []Ponto{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		},
	}
	matriz := forma.ToMatriz()
	require.Equal(t, matriz.Dimensao, DimensaoMatriz{
		NumLinhas:  3,
		NumColunas: 4,
	})
	require.Equal(t, 0.0, matriz.Get(0, 0))
	require.Equal(t, 0.0, matriz.Get(1, 0))
	require.Equal(t, 1.0, matriz.Get(2, 0))

	require.Equal(t, 2.0, matriz.Get(0, 1))
	require.Equal(t, 0.0, matriz.Get(1, 1))
	require.Equal(t, 1.0, matriz.Get(2, 1))

	require.Equal(t, 2.0, matriz.Get(0, 2))
	require.Equal(t, 2.0, matriz.Get(1, 2))
	require.Equal(t, 1.0, matriz.Get(2, 2))

	require.Equal(t, 0.0, matriz.Get(0, 3))
	require.Equal(t, 2.0, matriz.Get(1, 3))
	require.Equal(t, 1.0, matriz.Get(2, 3))
}

func TestForma2D_IniciaOperacao_Vazia(t *testing.T) {
	f0 := Forma2D{}
	f1 := f0.IniciaOperacao().Calcula()
	require.Equal(t, &f0, &f1)
}

// Rotação de 30°
func TestForma2D_Exemplo01SlidesAula10(t *testing.T) {
	f1 := Forma2D{
		Pontos: []Ponto{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		},
	}
	f2 := f1.IniciaOperacao().RotacionaMatricial(
		GrausParaRadianos(30),
	).Calcula()
	require.Len(t, f2.Pontos, 4)
	require.Equal(t, Ponto{X: 0, Y: 0}, f2.Pontos[0])
	require.Equal(t, Ponto{X: 2, Y: 1}, f2.Pontos[1])
	require.Equal(t, Ponto{X: 1, Y: 3}, f2.Pontos[2])
	require.Equal(t, Ponto{X: -1, Y: 2}, f2.Pontos[3])
}

// Transladar (-1, 1) depois rotacionar 30°
func TestForma2D_Exemplo02SlidesAula10(t *testing.T) {
	f1 := Forma2D{
		Pontos: []Ponto{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		},
	}
	f2 := f1.IniciaOperacao().MoveMatricial(
		-1, 1,
	).RotacionaMatricial(
		GrausParaRadianos(30),
	).Calcula()
	require.Len(t, f2.Pontos, 4)
	require.Equal(t, Ponto{X: -1, Y: 1}, f2.Pontos[0])
	require.Equal(t, Ponto{X: 1, Y: 2}, f2.Pontos[1])
	require.Equal(t, Ponto{X: 0, Y: 4}, f2.Pontos[2])
	require.Equal(t, Ponto{X: -2, Y: 3}, f2.Pontos[3])
}

// Rotacionar 30º depois Transladar (-1, 1)
func TestForma2D_Exemplo03SlidesAula10(t *testing.T) {
	f1 := Forma2D{
		Pontos: []Ponto{
			{X: 0, Y: 0},
			{X: 2, Y: 0},
			{X: 2, Y: 2},
			{X: 0, Y: 2},
		},
	}
	f2 := f1.IniciaOperacao().RotacionaMatricial(
		GrausParaRadianos(30),
	).MoveMatricial(
		-1, 1,
	).Calcula()
	require.Len(t, f2.Pontos, 4)
	require.Equal(t, Ponto{X: -1, Y: 0}, f2.Pontos[0])
	require.Equal(t, Ponto{X: 0, Y: 1}, f2.Pontos[1])
	require.Equal(t, Ponto{X: -1, Y: 3}, f2.Pontos[2])
	require.Equal(t, Ponto{X: -2, Y: 2}, f2.Pontos[3])
}
