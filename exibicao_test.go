package ufpa_cg

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestExibicao_1Quadrante_Decrescente(t *testing.T) {
	var writer strings.Builder
	require.Nil(t, Exibe(
		NewAlgoritmoBresenham(Ponto{X: -4, Y: 5}, Ponto{X: 0, Y: 2}),
		&writer,
	))
	fmt.Println("\n" + writer.String())
	lines := strings.Split(writer.String(), "\n")
	require.Equal(t, "     │ ", lines[0])
	require.Equal(t, " █   │ ", lines[1])
	require.Equal(t, "  ░░ │ ", lines[2])
	require.Equal(t, "    ░│ ", lines[3])
	require.Equal(t, "     █ ", lines[4])
	require.Equal(t, "     │ ", lines[5])
}

func TestExibicao_1Octante_Crescente(t *testing.T) {
	var writer strings.Builder
	require.Nil(t, Exibe(
		NewAlgoritmoBresenham(Ponto{X: 0, Y: 0}, Ponto{X: 5, Y: 3}),
		&writer,
	))
	lines := strings.Split(writer.String(), "\n")
	require.Equal(t, " │      ", lines[0])
	require.Equal(t, " │    █ ", lines[1])
	require.Equal(t, " │  ░░  ", lines[2])
	require.Equal(t, " │░░    ", lines[3])
	require.Equal(t, "─█──────", lines[4])
	require.Equal(t, " │      ", lines[5])
	require.Len(t, lines, 7)
}

func TestExibicao_2Octante_Crescente(t *testing.T) {
	var writer strings.Builder
	require.Nil(t, Exibe(
		NewAlgoritmoBresenham(Ponto{X: 0, Y: 3}, Ponto{X: 3, Y: 9}),
		&writer,
	))
	lines := strings.Split(writer.String(), "\n")
	require.Equal(t, " │    ", lines[0])
	require.Equal(t, " │  █ ", lines[1])
	require.Equal(t, " │ ░  ", lines[2])
	require.Equal(t, " │ ░  ", lines[3])
	require.Equal(t, " │░   ", lines[4])
	require.Equal(t, " │░   ", lines[5])
	require.Equal(t, " ░    ", lines[6])
	require.Equal(t, " █    ", lines[7])
	require.Equal(t, " │    ", lines[8])
	require.Len(t, lines, 10)
}
