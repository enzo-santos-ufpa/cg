package ufpa_cg

import "math"

func (p Ponto3D) MoveMatricial(tx, ty, tz int) Ponto3D {
	m0, _ := NewMatriz([][]float64{
		{1, 0, 0, float64(tx)},
		{0, 1, 0, float64(ty)},
		{0, 0, 1, float64(tz)},
		{0, 0, 0, 1},
	})
	m1, _ := NewMatriz([][]float64{
		{float64(p.X)},
		{float64(p.Y)},
		{float64(p.Z)},
		{1},
	})
	m2, _ := m0.Multiplica(m1)
	return Ponto3D{
		X: int(math.Round(m2.Get(0, 0))),
		Y: int(math.Round(m2.Get(1, 0))),
		Z: int(math.Round(m2.Get(2, 0))),
	}
}

func (p Ponto3D) RedimensionaMatricial(ex, ey, ez float64) Ponto3D {
	m0, _ := NewMatriz([][]float64{
		{ex, 0, 0, 0},
		{0, ey, 0, 0},
		{0, 0, ez, 0},
		{0, 0, 0, 1},
	})
	m1, _ := NewMatriz([][]float64{
		{float64(p.X)},
		{float64(p.Y)},
		{float64(p.Z)},
		{1},
	})
	m2, _ := m0.Multiplica(m1)
	return Ponto3D{
		X: int(math.Round(m2.Get(0, 0))),
		Y: int(math.Round(m2.Get(1, 0))),
		Z: int(math.Round(m2.Get(2, 0))),
	}
}

func (p Ponto3D) RotacionaMatricial(angulo float64, eixo Eixo3D) Ponto3D {
	senA := math.Sin(angulo)
	cosA := math.Cos(angulo)

	var m0 *Matriz
	switch eixo {
	case EixoX:
		m0, _ = NewMatriz([][]float64{
			{1, 0, 0, 0},
			{0, cosA, -senA, 0},
			{0, senA, cosA, 0},
			{0, 0, 0, 1},
		})
	case EixoY:
		m0, _ = NewMatriz([][]float64{
			{cosA, 0, senA, 0},
			{0, 1, 0, 0},
			{-senA, 0, cosA, 0},
			{0, 0, 0, 1},
		})
	case EixoZ:
		m0, _ = NewMatriz([][]float64{
			{cosA, -senA, 0, 0},
			{senA, cosA, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		})
	}
	m1, _ := NewMatriz([][]float64{
		{float64(p.X)},
		{float64(p.Y)},
		{float64(p.Z)},
		{1},
	})
	m2, _ := m0.Multiplica(m1)
	return Ponto3D{
		X: int(math.Round(m2.Get(0, 0))),
		Y: int(math.Round(m2.Get(1, 0))),
		Z: int(math.Round(m2.Get(2, 0))),
	}
}
