package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
	"image/color"
)

type JogoSecundario struct {
	Source *ebitentext.GoTextFaceSource
	Modulo ModuloGame
}

func (j *JogoSecundario) Update() error {
	return j.Modulo.Update()
}

func (j *JogoSecundario) Draw(screen *ebiten.Image) {
	f := &ebitentext.GoTextFace{
		Source:    j.Source,
		Direction: ebitentext.DirectionLeftToRight,
		Size:      16,
		Language:  language.BrazilianPortuguese,
	}

	// Desenha texto "Pressione ESC para voltar" no topo
	op := &ebitentext.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(10, 10)
	text := "Pressione ESC para voltar"
	ebitentext.Draw(screen, text, f, op)
	_, h := ebitentext.Measure(text, f, 0)

	j.Modulo.Draw(screen, f, int(h)+20)
}

func (j *JogoSecundario) Layout(_, _ int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
