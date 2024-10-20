package ufpa_cg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_JanelaRecorte_Intersecao_InferiorCentralParaCentralDireito(t *testing.T) {
	j := JanelaRecorte{
		PontoInferiorEsquerdo: Ponto{X: 0, Y: 0},
		PontoSuperiorDireito:  Ponto{X: 25, Y: 25},
	}
	p1 := Ponto{X: 10, Y: -5}
	p2 := Ponto{X: 30, Y: 10}
	require.Equal(t, j.CalculaIntersecao(p1, p2), ResultadoIntersecao{
		EhRetaCompleta:   false,
		EhRetaVisivel:    true,
		PontoIntersecaoX: Ponto{X: 17, Y: 0},
		PontoIntersecaoY: Ponto{X: 25, Y: 6},
	})
}

func Test_JanelaRecorte_Intersecao_SuperiorCentralParaCentralDireito(t *testing.T) {
	j := JanelaRecorte{
		PontoInferiorEsquerdo: Ponto{X: 0, Y: 0},
		PontoSuperiorDireito:  Ponto{X: 25, Y: 25},
	}
	p1 := Ponto{X: 10, Y: 30}
	p2 := Ponto{X: 30, Y: 15}
	require.Equal(t, j.CalculaIntersecao(p1, p2), ResultadoIntersecao{
		EhRetaCompleta:   false,
		EhRetaVisivel:    true,
		PontoIntersecaoX: Ponto{X: 17, Y: 25},
		PontoIntersecaoY: Ponto{X: 25, Y: 19},
	})
}

func Test_JanelaRecorte_Intersecao_CentralEsquerdoParaInferiorCentral(t *testing.T) {
	j := JanelaRecorte{
		PontoInferiorEsquerdo: Ponto{X: 0, Y: 0},
		PontoSuperiorDireito:  Ponto{X: 25, Y: 25},
	}
	p1 := Ponto{X: -5, Y: 10}
	p2 := Ponto{X: 15, Y: -5}
	require.Equal(t, ResultadoIntersecao{
		EhRetaCompleta:   false,
		EhRetaVisivel:    true,
		PontoIntersecaoX: Ponto{X: 8, Y: 0},
		PontoIntersecaoY: Ponto{X: 0, Y: 6},
	}, j.CalculaIntersecao(p1, p2))
}

func Test_JanelaRecorte_Intersecao_InferiorCentralParaSuperiorCentral(t *testing.T) {
	j := JanelaRecorte{
		PontoInferiorEsquerdo: Ponto{X: 0, Y: 0},
		PontoSuperiorDireito:  Ponto{X: 25, Y: 25},
	}
	p1 := Ponto{X: 10, Y: -5}
	p2 := Ponto{X: 15, Y: 30}
	require.Equal(t, ResultadoIntersecao{
		EhRetaCompleta:   false,
		EhRetaVisivel:    true,
		PontoIntersecaoX: Ponto{X: 11, Y: 0},
		PontoIntersecaoY: Ponto{X: 14, Y: 25},
	}, j.CalculaIntersecao(p1, p2))
}
