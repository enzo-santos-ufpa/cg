package ufpa_cg

import (
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
	lines := strings.Split(writer.String(), "\n")
	require.Equal(t, "⁶ ┊   ┃ ", lines[0])
	require.Equal(t, "⁵╌█╌╌╌┃╌", lines[1])
	require.Equal(t, "⁴ ┊░░ ┃ ", lines[2])
	require.Equal(t, "³ ┊  ░┃ ", lines[3])
	require.Equal(t, "²╌┊╌╌╌█╌", lines[4])
	require.Equal(t, "¹ ┊   ┃ ", lines[5])
	require.Equal(t, " ⁵⁴³²¹⁰¹", lines[6])
}

func TestExibicao_1Octante_Crescente(t *testing.T) {
	var writer strings.Builder
	require.Nil(t, Exibe(
		NewAlgoritmoBresenham(Ponto{X: 0, Y: 0}, Ponto{X: 5, Y: 3}),
		&writer,
	))
	lines := strings.Split(writer.String(), "\n")
	require.Equal(t, "⁴ ┃    ┊ ", lines[0])
	require.Equal(t, "³╌┃╌╌╌╌█╌", lines[1])
	require.Equal(t, "² ┃  ░░┊ ", lines[2])
	require.Equal(t, "¹ ┃░░  ┊ ", lines[3])
	require.Equal(t, "⁰━█━━━━━━", lines[4])
	require.Equal(t, "¹ ┃    ┊ ", lines[5])
	require.Equal(t, " ¹⁰¹²³⁴⁵⁶", lines[6])
}

func TestExibicao_123Quadrantes_Crescente(t *testing.T) {
	var writer strings.Builder
	require.Nil(t, Exibe(
		NewAlgoritmoBresenham(Ponto{X: -4, Y: -3}, Ponto{X: 2, Y: 5}),
		&writer,
	))
	lines := strings.Split(writer.String(), "\n")
	require.Equal(t, "⁶ ┊   ┃ ┊ ", lines[0])
	require.Equal(t, "⁵╌┊╌╌╌┃╌█╌", lines[1])
	require.Equal(t, "⁴ ┊   ┃░┊ ", lines[2])
	require.Equal(t, "³ ┊   ░ ┊ ", lines[3])
	require.Equal(t, "² ┊   ░ ┊ ", lines[4])
	require.Equal(t, "¹ ┊  ░┃ ┊ ", lines[5])
	require.Equal(t, "⁰━━━░━╋━━━", lines[6])
	require.Equal(t, "¹ ┊░  ┃ ┊ ", lines[7])
	require.Equal(t, "² ┊░  ┃ ┊ ", lines[8])
	require.Equal(t, "³╌█╌╌╌┃╌┊╌", lines[9])
	require.Equal(t, "⁴ ┊   ┃ ┊ ", lines[10])
	require.Equal(t, " ⁵⁴³²¹⁰¹²³", lines[11])
}

func TestExibicao_2Octante_Crescente(t *testing.T) {
	var writer strings.Builder
	require.Nil(t, Exibe(
		NewAlgoritmoBresenham(Ponto{X: 0, Y: 3}, Ponto{X: 3, Y: 9}),
		&writer,
	))
	lines := strings.Split(writer.String(), "\n")
	require.Equal(t, "⁰ ┃  ┊ ", lines[0])
	require.Equal(t, "⁹╌┃╌╌█╌", lines[1])
	require.Equal(t, "⁸ ┃ ░┊ ", lines[2])
	require.Equal(t, "⁷ ┃ ░┊ ", lines[3])
	require.Equal(t, "⁶ ┃░ ┊ ", lines[4])
	require.Equal(t, "⁵ ┃░ ┊ ", lines[5])
	require.Equal(t, "⁴ ░  ┊ ", lines[6])
	require.Equal(t, "³╌█╌╌┊╌", lines[7])
	require.Equal(t, "² ┃  ┊ ", lines[8])
	require.Equal(t, " ¹⁰¹²³⁴", lines[9])
}
