package ufpa_cg

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPonto3D_MoveMatricial(t *testing.T) {
	p1 := Ponto3D{X: 10, Y: 20, Z: 30}
	p2 := p1.MoveMatricial(1, 2, 3)
	require.Equal(t, Ponto3D{X: 11, Y: 22, Z: 33}, p2)
}

func TestPonto3D_RedimensionaMatricial(t *testing.T) {
	p1 := Ponto3D{X: 10, Y: 20, Z: 30}
	p2 := p1.RedimensionaMatricial(1, 2, 3)
	require.Equal(t, Ponto3D{X: 10, Y: 40, Z: 90}, p2)
}

// https://www.mathforengineers.com/math-calculators/3D-point-rotation-calculator.html
func TestPonto3D_RotacionaMatricial_X(t *testing.T) {
	p1 := Ponto3D{X: 10, Y: 20, Z: 30}
	p2 := p1.RotacionaMatricial(math.Pi/4, EixoX)
	require.Equal(t, Ponto3D{X: 10, Y: -7, Z: 35}, p2)
}

// https://www.mathforengineers.com/math-calculators/3D-point-rotation-calculator.html
func TestPonto3D_RotacionaMatricial_Y(t *testing.T) {
	p1 := Ponto3D{X: 10, Y: 20, Z: 30}
	p2 := p1.RotacionaMatricial(math.Pi/4, EixoY)
	require.Equal(t, Ponto3D{X: 28, Y: 20, Z: 14}, p2)
}

// https://www.mathforengineers.com/math-calculators/3D-point-rotation-calculator.html
func TestPonto3D_RotacionaMatricial_Z(t *testing.T) {
	p1 := Ponto3D{X: 10, Y: 20, Z: 30}
	p2 := p1.RotacionaMatricial(math.Pi/4, EixoZ)
	require.Equal(t, Ponto3D{X: -7, Y: 21, Z: 30}, p2)
}
