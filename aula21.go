package ufpa_cg

func sign(x float64) int {
	if x >= 0 {
		return 0
	}
	return 1
}

type JanelaRecorte struct {
	PontoSuperiorEsquerdo Ponto
	PontoInferiorDireito  Ponto
}

func (j JanelaRecorte) calculaBits(ponto Ponto) int {
	b := 0
	b |= sign(float64(j.PontoSuperiorEsquerdo.Y)-float64(ponto.Y)) << 3
	b |= sign(float64(ponto.Y)-float64(j.PontoInferiorDireito.Y)) << 2
	b |= sign(float64(j.PontoInferiorDireito.X)-float64(ponto.X)) << 1
	b |= sign(float64(ponto.X)-float64(j.PontoSuperiorEsquerdo.X)) << 0
	return b
}

func (j JanelaRecorte) CalculaIntersecao(p1, p2 Ponto) (Ponto, Ponto, bool) {
	b1 := j.calculaBits(p1)
	b2 := j.calculaBits(p2)

	for {
		if b1|b2 == 0 {
			// Totalmente dentro
			return p1, p2, true
		}
		if b1&b2 != 0 {
			// Totalmente fora
			return Ponto{}, Ponto{}, false
		}

		x, y := 0.0, 0.0
		var b int
		if b1 != 0 {
			b = b1
		} else {
			b = b2
		}
		switch {
		case b>>3&1 == 1:
			x = float64(p1.X) + float64(p2.X-p1.X)*float64(j.PontoSuperiorEsquerdo.Y-p1.Y)/float64(p2.Y-p1.Y)
			y = float64(j.PontoSuperiorEsquerdo.Y)
		case b>>2&1 == 1:
			x = float64(p1.X) + float64(p2.X-p1.X)*float64(j.PontoInferiorDireito.Y-p1.Y)/float64(p2.Y-p1.Y)
			y = float64(j.PontoInferiorDireito.Y)
		case b>>1&1 == 1:
			x = float64(j.PontoInferiorDireito.X)
			y = float64(p1.Y) + float64(p2.Y-p1.Y)*float64(j.PontoInferiorDireito.X-p1.X)/float64(p2.X-p1.X)
		case b>>0&1 == 1:
			x = float64(j.PontoSuperiorEsquerdo.X)
			y = float64(p1.Y) + float64(p2.Y-p1.Y)*float64(j.PontoSuperiorEsquerdo.X-p1.X)/float64(p2.X-p1.X)
		}
		if b == b1 {
			p1 = Ponto{X: int(x), Y: int(y)}
			b1 = j.calculaBits(p1)
		} else {
			p2 = Ponto{X: int(x), Y: int(y)}
			b2 = j.calculaBits(p2)
		}
	}
}
