package ufpa_cg

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRotaciona_Origem(t *testing.T) {
	require.Equal(t, Ponto{X: 1, Y: 5}, Ponto{X: 4, Y: 4}.Rotaciona(30.0*math.Pi/180.0, Ponto{X: 0, Y: 0}))
}

func TestRotaciona_QualquerPonto(t *testing.T) {
	require.Equal(t, Ponto{X: 3, Y: 6}, Ponto{X: 6, Y: 5}.Rotaciona(30.0*math.Pi/180.0, Ponto{X: 2, Y: 1}))
}
