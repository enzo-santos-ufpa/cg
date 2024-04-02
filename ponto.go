package ufpa_cg

type Ponto struct {
	X int
	Y int
}

func CoeficienteLinear(p1, p2 Ponto) float64 {
	return float64(p2.Y-p1.Y) / float64(p2.X-p1.X)
}
