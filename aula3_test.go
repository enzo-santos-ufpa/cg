package ufpa_cg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgoritmoBresenham2_1(t *testing.T) {
	algoritmo := NewAlgoritmoBresenham(Ponto{X: 0, Y: 3}, Ponto{X: 3, Y: 9})
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 0, Y: 3}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 0, Y: 4}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 1, Y: 5}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 1, Y: 6}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 2, Y: 7}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 2, Y: 8}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 3, Y: 9}, algoritmo.PontoAtual())
	assert.False(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 3, Y: 9}, algoritmo.PontoAtual())
}

func TestAlgoritmoBresenham2_2(t *testing.T) {
	algoritmo := NewAlgoritmoBresenham(Ponto{X: 0, Y: 0}, Ponto{X: 2, Y: 5})
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 0, Y: 0}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 0, Y: 1}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 1, Y: 2}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 1, Y: 3}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 2, Y: 4}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 2, Y: 5}, algoritmo.PontoAtual())
	assert.False(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 2, Y: 5}, algoritmo.PontoAtual())
}
