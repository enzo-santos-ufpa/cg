package ufpa_cg

import "math"

func (p Ponto) MoveMatricial(dx, dy int) Ponto {
	m0, _ := NewMatriz([][]float64{
		{float64(dx)},
		{float64(dy)},
	})
	m1, _ := NewMatriz([][]float64{
		{float64(p.X)},
		{float64(p.Y)},
	})
	m2, _ := m0.Soma(m1)
	return Ponto{
		X: int(math.Round(m2.Get(0, 0))),
		Y: int(math.Round(m2.Get(1, 0))),
	}
}

func (p Ponto) RedimensionaMatricial(ex, ey float64) Ponto {
	m0, _ := NewMatriz([][]float64{
		{ex, 0},
		{0, ey},
	})
	m1, _ := NewMatriz([][]float64{
		{float64(p.X)},
		{float64(p.Y)},
	})
	m2, _ := m0.Multiplica(m1)
	return Ponto{
		X: int(math.Round(m2.Get(0, 0))),
		Y: int(math.Round(m2.Get(1, 0))),
	}
}

func (p Ponto) RotacionaMatricial(angulo float64) Ponto {
	m0, _ := NewMatriz([][]float64{
		{math.Cos(angulo), -math.Sin(angulo)},
		{math.Sin(angulo), math.Cos(angulo)},
	})
	m1, _ := NewMatriz([][]float64{
		{float64(p.X)},
		{float64(p.Y)},
	})
	m2, _ := m0.Multiplica(m1)
	return Ponto{
		X: int(math.Round(m2.Get(0, 0))),
		Y: int(math.Round(m2.Get(1, 0))),
	}
}
