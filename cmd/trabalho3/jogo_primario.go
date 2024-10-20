package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
	"image/color"
)

type JogoPrimario struct {
	TextFont *TextFont

	menu           ebiten.Game
	options        []OpcaoMenu
	selectingIndex int
}

func (j *JogoPrimario) Update() error {
	menu := j.menu
	if menu == nil {
		switch {
		case repeatingKeyPressed(ebiten.KeyDown) && j.selectingIndex < len(j.options)-1:
			j.selectingIndex++
		case repeatingKeyPressed(ebiten.KeyUp) && j.selectingIndex > 0:
			j.selectingIndex--
		case repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter):
			j.menu = &JogoSecundario{
				Source: j.TextFont.Source,
				Modulo: j.options[j.selectingIndex].Create(),
			}
		}

	} else {
		if repeatingKeyPressed(ebiten.KeyEscape) {
			j.menu = nil
		} else {
			return menu.Update()
		}
	}
	return nil
}

func (j *JogoPrimario) Draw(screen *ebiten.Image) {
	menu := j.menu
	if menu != nil {
		menu.Draw(screen)
		return
	}

	const headerHeightOffset = 10.0 // Quanto o cabeçalho "Menu de opções" deve ficar deslocado para baixo
	const headerWidthOffset = 15.0  // Quanto o cabeçalho "Menu de opções" deve ficar deslocado à direita
	const headerLineSpacing = 10.0  // Quanto o cabeçalho "Menu de opções" deve ficar separado do corpo de opções
	const widthOffset = 20.0        // Quanto cada opão deve ficar deslocada à direita
	const lineSpacing = 5.0         // Quanto cada opção deve ficar sepadada uma da outra
	const fontSize = 16             // Tamanho da fonte de cada texto nessa tela

	heightOffset := headerHeightOffset

	// Constrói lista de opções do menu principal
	labels := make([]string, len(j.options)+1)
	labels[0] = "Menu de opções"
	for i, module := range j.options {
		labels[i+1] = module.Title()
	}
	for n, label := range labels {
		i := n - 1
		op := &ebitentext.DrawOptions{}
		if n == 0 {
			op.GeoM.Translate(headerWidthOffset, heightOffset)
		} else {
			op.GeoM.Translate(widthOffset, heightOffset)
		}
		if j.selectingIndex == i {
			op.ColorScale.ScaleWithColor(color.RGBA{R: 0x42, G: 0x87, B: 0xf5})
		} else {
			op.ColorScale.ScaleWithColor(color.White)
		}
		f := &ebitentext.GoTextFace{
			Source:    j.TextFont.Source,
			Direction: ebitentext.DirectionLeftToRight,
			Size:      fontSize,
			Language:  language.BrazilianPortuguese,
		}
		var text string
		if n == 0 {
			text = fmt.Sprintf("%s\n", label)
		} else {
			text = fmt.Sprintf("%d. %s\n", i+1, label)
		}
		ebitentext.Draw(screen, text, f, op)
		_, height := ebitentext.Measure(text, f, 0)
		if n == 0 {
			heightOffset += height + headerLineSpacing
		} else {
			heightOffset += height + lineSpacing
		}
	}
}

func (j *JogoPrimario) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}
