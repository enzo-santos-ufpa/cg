package ufpa_cg

type Ponto struct {
	X int
	Y int
}

func CoeficienteLinear(p1, p2 Ponto) float64 {
	return float64(p2.Y-p1.Y) / float64(p2.X-p1.X)
}

type Eixo3D int

const (
	EixoX Eixo3D = iota
	EixoY
	EixoZ
)

type Ponto3D struct {
	X int
	Y int
	Z int
}
