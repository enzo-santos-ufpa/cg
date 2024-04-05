package ufpa_cg

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPonto_MoveMatricial(t *testing.T) {
	p1 := Ponto{X: 1, Y: 2}
	p2 := p1.MoveMatricial(3, 5)
	require.Equal(t, Ponto{X: 4, Y: 7}, p2)
}

func TestPonto_RedimensionaMatricial(t *testing.T) {
	p1 := Ponto{X: 2, Y: 3}
	p2 := p1.RedimensionaMatricial(4, 5)
	require.Equal(t, Ponto{X: 8, Y: 15}, p2)
}

func TestPonto_RotacionaMatricial(t *testing.T) {
	p1 := Ponto{X: 4, Y: 4}
	p2 := p1.RotacionaMatricial(30.0 * math.Pi / 180.0)
	require.Equal(t, Ponto{X: 1, Y: 5}, p2)
}
