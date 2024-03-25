package ufpa_cg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlgoritmoBruto(t *testing.T) {
	algoritmo := NewAlgoritmoBruto(Ponto{X: 0, Y: 0}, Ponto{X: 5, Y: 3})
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 0, Y: 0}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 1, Y: 1}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 2, Y: 1}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 3, Y: 2}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 4, Y: 2}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 5, Y: 3}, algoritmo.PontoAtual())
	assert.False(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 5, Y: 3}, algoritmo.PontoAtual())
}

func TestAlgoritmoBresenham(t *testing.T) {
	algoritmo := NewAlgoritmoBresenham(Ponto{X: 0, Y: 0}, Ponto{X: 5, Y: 3})
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 0, Y: 0}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 1, Y: 1}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 2, Y: 1}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 3, Y: 2}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 4, Y: 2}, algoritmo.PontoAtual())
	assert.True(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 5, Y: 3}, algoritmo.PontoAtual())
	assert.False(t, algoritmo.Move())
	assert.Equal(t, Ponto{X: 5, Y: 3}, algoritmo.PontoAtual())
}
