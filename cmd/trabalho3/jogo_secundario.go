package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
	"image/color"
	"slices"
	"ufpa_cg"
)

type JogoSecundario struct {
	Source *ebitentext.GoTextFaceSource
	Modulo ModuloJogo

	output []ufpa_cg.Ponto
}

func (j *JogoSecundario) Update() error {
	settings := j.Modulo.Settings()
	if settings == nil {
		return nil
	}

	// Seleciona o ponto atual da tela
	allEvaluated := true
	for _, inp := range settings.Inputs() {
		if _, evaluated := inp.Evaluated(); evaluated {
			continue
		}
		inp.OnUpdate()
		allEvaluated = false
		break
	}
	if allEvaluated {
		// Calcula a reta para as entradas selecionadas
		if output := j.output; output == nil {
			j.output = settings.Evaluate()
		}

		// Reseta o estado da tela após todas as entradas tiverem sidos selecionadas
		if repeatingKeyPressed(ebiten.KeySpace) {
			for _, inp := range settings.Inputs() {
				inp.Reset()
			}
			j.output = nil
		}
	}
	return nil
}

func (j *JogoSecundario) Draw(screen *ebiten.Image) {
	textFace := &ebitentext.GoTextFace{
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
	ebitentext.Draw(screen, text, textFace, op)
	_, h := ebitentext.Measure(text, textFace, 0)

	heightOffset := int(h) + 20

	const pointSize = 10       // Tamanho de um ponto, em pixels
	const pointSpacing = 3     // Espaçamento entre um ponto e outro, em pixels
	const minX, maxX = -11, 11 // Intervalo de X para o grid de pontos
	const minY, maxY = -11, 11 // Intervalo de Y para o grid de pontos

	var selectedColor = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}      // Cor para um ponto selecionado
	var filledColor = color.RGBA{R: 0xFF, G: 0x62, B: 0x00, A: 0xFF}        // Cor para um ponto na linha formada
	var defaultColor = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}       // Cor para um ponto x != 0 && y != 0
	var defaultOriginColor = color.RGBA{R: 0x91, G: 0x91, B: 0x91, A: 0xFF} // Cor para um ponto x == 0 || y == 0

	settings := j.Modulo.Settings()
	if settings == nil {
		return
	}

	stamps := make(map[ufpa_cg.Ponto]color.Color)
	var currentInput EntradaModulo
	for _, inp := range settings.Inputs() {
		stamp, evaluated := inp.Evaluated()
		if !evaluated {
			currentInput = inp
			break
		}
		for ponto, cor := range stamp {
			stamps[ponto] = cor
		}
	}

	// Percorre o grid de pontos
	for j_ := minY; j_ <= maxY; j_++ {
		for i := minX; i <= maxX; i++ {
			x := (i-minX)*(pointSize+pointSpacing) + 20
			y := heightOffset + j_ - minY

			ponto := ufpa_cg.Ponto{X: i, Y: -j_}

			selected := false
			for _, inp := range settings.Inputs() {
				if inp.Selected(ponto) {
					selected = true
					break
				}
			}

			var customColor color.Color
			var ok bool
			if currentInput == nil {
				customColor, ok = nil, false
			} else {
				customColor, ok = currentInput.OnDraw(ponto, x, y, pointSize)
			}

			pixel := ebiten.NewImage(pointSize, pointSize)
			if stampColor, isStamp := stamps[ponto]; isStamp {
				pixel.Fill(stampColor)
			} else if ok {
				pixel.Fill(customColor)
			} else if selected { // Se este ponto for algum selecionado
				pixel.Fill(selectedColor) // Marca este ponto como "SELECIONADO"
			} else if len(j.output) > 0 && slices.Contains(j.output, ponto) { // Se este ponto estiver na linha formada pelos pontos A e B
				pixel.Fill(filledColor) // Marca este ponto como "NA RETA"
			} else if i == 0 || j_ == 0 { // Se este ponto estiver em algum dos eixos X ou Y
				pixel.Fill(defaultOriginColor) // Marca este ponto como "NÃO SELECIONADO, NA ORIGEM"
			} else {
				pixel.Fill(defaultColor) // Marca este ponto como "NÃO SELECIONADO"
			}

			// Desenha o ponto
			op := ebiten.GeoM{}
			op.Translate(float64(x), float64(y))
			screen.DrawImage(pixel, &ebiten.DrawImageOptions{GeoM: op})
		}
		heightOffset += pointSize + pointSpacing
	}
	heightOffset += 40

	for _, inp := range settings.Inputs() {
		const dx = 20

		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(dx, float64(heightOffset))
		promptLabel := inp.DescribePrompt()
		ebitentext.Draw(screen, promptLabel, textFace, op)
		w, h := ebitentext.Measure(promptLabel, textFace, 0)
		if _, evaluated := inp.Evaluated(); !evaluated {
			if text, ok := inp.OnDisplay(); ok {
				ebitenutil.DebugPrintAt(screen, text, dx+int(w)+10, heightOffset)
			}
			heightOffset += int(h) + 5
			break
		} else {
			op := &ebitentext.DrawOptions{}
			op.ColorScale.ScaleWithColor(color.White)
			op.GeoM.Translate(dx+w+10, float64(heightOffset))
			ebitentext.Draw(screen, inp.DescribeValue(), textFace, op)
		}
		heightOffset += int(h) + 5
	}
	if currentInput != nil {
		if hintLabel, ok := currentInput.DescribeAction(); ok {
			op := &ebitentext.DrawOptions{}
			op.ColorScale.ScaleWithColor(color.White)
			op.GeoM.Translate(20, float64(heightOffset))
			ebitentext.Draw(screen, hintLabel, textFace, op)
			_, h := ebitentext.Measure(hintLabel, textFace, 0)
			heightOffset += int(h) + 5
		}
	}
	if len(settings.Inputs()) > 0 && currentInput == nil {
		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(20, float64(heightOffset))
		ebitentext.Draw(screen, "Pressione ESPAÇO para refazer", textFace, op)
	}
}

func (j *JogoSecundario) Layout(_, _ int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
