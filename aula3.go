package ufpa_cg

type AlgoritmoBresenham struct {
	Algoritmo *AlgoritmoBresenham1Octante
	TrocaX    bool
	TrocaY    bool
	TrocaXY   bool
}

func NewAlgoritmoBresenham2(p1, p2 Ponto) *AlgoritmoBresenham {
	var trocaX, trocaY, trocaXY bool
	m := float64(p2.Y-p1.Y) / float64(p2.X-p1.X)
	pA := p1
	pB := p2
	if m > 1 || m < -1 {
		pA = Ponto{X: pA.Y, Y: pA.X}
		pB = Ponto{X: pB.Y, Y: pB.X}
		trocaXY = true
	}
	if p1.X > p2.X {
		pA = Ponto{X: -pA.X, Y: pA.Y}
		pB = Ponto{X: -pB.X, Y: pB.Y}
		trocaX = true
	}
	if p1.Y > p2.Y {
		pA = Ponto{X: pA.X, Y: -pA.Y}
		pB = Ponto{X: pB.X, Y: -pB.Y}
		trocaY = true
	}
	return &AlgoritmoBresenham{
		Algoritmo: NewAlgoritmoBresenham1Octante(pA, pB),
		TrocaX:    trocaX,
		TrocaY:    trocaY,
		TrocaXY:   trocaXY,
	}
}

func (a *AlgoritmoBresenham) Move() bool {
	return a.Algoritmo.Move()
}

func (a *AlgoritmoBresenham) PontoAtual() Ponto {
	ponto := a.Algoritmo.PontoAtual()
	if a.TrocaX {
		ponto = Ponto{X: -ponto.X, Y: ponto.Y}
	}
	if a.TrocaY {
		ponto = Ponto{X: ponto.X, Y: -ponto.Y}
	}
	if a.TrocaXY {
		ponto = Ponto{X: ponto.Y, Y: ponto.X}
	}
	return ponto
}
