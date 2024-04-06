package ufpa_cg

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAlgoritmoBresenham_1(t *testing.T) {
	algoritmo := NewAlgoritmoBresenham(Ponto{X: 0, Y: 3}, Ponto{X: 3, Y: 9})
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 0, Y: 3}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 0, Y: 4}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 1, Y: 5}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 1, Y: 6}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 2, Y: 7}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 2, Y: 8}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 3, Y: 9}, algoritmo.PontoAtual())
	require.False(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 3, Y: 9}, algoritmo.PontoAtual())
}

func TestAlgoritmoBresenham_2(t *testing.T) {
	algoritmo := NewAlgoritmoBresenham(Ponto{X: 0, Y: 0}, Ponto{X: 2, Y: 5})
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 0, Y: 0}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 0, Y: 1}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 1, Y: 2}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 1, Y: 3}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 2, Y: 4}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 2, Y: 5}, algoritmo.PontoAtual())
	require.False(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 2, Y: 5}, algoritmo.PontoAtual())
}

func TestAlgoritmoBresenham_3(t *testing.T) {
	algoritmo := NewAlgoritmoBresenham(Ponto{X: -4, Y: 5}, Ponto{X: 0, Y: 2})
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -4, Y: 5}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -3, Y: 4}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -2, Y: 4}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -1, Y: 3}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 0, Y: 2}, algoritmo.PontoAtual())
	require.False(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 0, Y: 2}, algoritmo.PontoAtual())
}

func TestAlgoritmoBresenham_4(t *testing.T) {
	algoritmo := NewAlgoritmoBresenham(Ponto{X: 0, Y: 2}, Ponto{X: -4, Y: 5})
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: 0, Y: 2}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -1, Y: 3}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -2, Y: 3}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -3, Y: 4}, algoritmo.PontoAtual())
	require.True(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -4, Y: 5}, algoritmo.PontoAtual())
	require.False(t, algoritmo.Move())
	require.Equal(t, Ponto{X: -4, Y: 5}, algoritmo.PontoAtual())
}
