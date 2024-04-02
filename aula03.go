package ufpa_cg

type AlgoritmoBresenham struct {
	algoritmo AlgoritmoLinha
	trocaX    bool
	trocaY    bool
	trocaXY   bool
}

func NewAlgoritmoBresenham(p1, p2 Ponto) AlgoritmoLinha {
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
		algoritmo: NewAlgoritmoBresenham1Octante(pA, pB),
		trocaX:    trocaX,
		trocaY:    trocaY,
		trocaXY:   trocaXY,
	}
}

func (a *AlgoritmoBresenham) Move() bool {
	return a.algoritmo.Move()
}

func (a *AlgoritmoBresenham) PontoAtual() Ponto {
	ponto := a.algoritmo.PontoAtual()
	if a.trocaX {
		ponto = Ponto{X: -ponto.X, Y: ponto.Y}
	}
	if a.trocaY {
		ponto = Ponto{X: ponto.X, Y: -ponto.Y}
	}
	if a.trocaXY {
		ponto = Ponto{X: ponto.Y, Y: ponto.X}
	}
	return ponto
}
