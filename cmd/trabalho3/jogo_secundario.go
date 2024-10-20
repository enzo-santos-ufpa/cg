package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/text/language"
	"image/color"
	"slices"
	"ufpa_cg"
)

type JogoSecundario struct {
	Source *ebitentext.GoTextFaceSource
	Modulo ModuloJogo

	cursorX    int
	cursorY    int
	pontoAtual *ufpa_cg.Ponto
	output     []ufpa_cg.Ponto
}

func (j *JogoSecundario) Update() error {
	x, y := ebiten.CursorPosition()
	j.cursorX = x
	j.cursorY = y

	settings := j.Modulo.Settings()
	if settings == nil {
		return nil
	}

	// Seleciona o ponto atual da tela
	allEvaluated := true
	for _, inp := range settings.Inputs() {
		if inp.Evaluated() {
			continue
		}
		allEvaluated = false
		if ponto := j.pontoAtual; inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && ponto != nil {
			inp.Consume(*ponto)
			break
		}
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

	var hoveredColor = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}       // Cor para um ponto com o cursor apontado
	var selectedColor = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}      // Cor para um ponto selecionado
	var filledColor = color.RGBA{R: 0xFF, G: 0x62, B: 0x00, A: 0xFF}        // Cor para um ponto na linha formada
	var defaultColor = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}       // Cor para um ponto x != 0 && y != 0
	var defaultOriginColor = color.RGBA{R: 0x91, G: 0x91, B: 0x91, A: 0xFF} // Cor para um ponto x == 0 || y == 0

	settings := j.Modulo.Settings()
	if settings == nil {
		return
	}

	hasPendingInput := false
	for _, inp := range settings.Inputs() {
		if !inp.Evaluated() {
			hasPendingInput = true
			break
		}
	}

	var hoveringPoint *ufpa_cg.Ponto
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

			pixel := ebiten.NewImage(pointSize, pointSize)
			if hasPendingInput && // Se algum ponto não tiver sido selecionado
				(x <= j.cursorX && j.cursorX <= x+pointSize) && // Se a posição X do cursor estiver dentro deste ponto
				(y <= j.cursorY && j.cursorY <= y+pointSize) { // Se a posição Y do cursor estiver dentro deste ponto
				pixel.Fill(hoveredColor) // Marca este ponto como "EM SELEÇÃO"
				hoveringPoint = &ponto
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

	hasHoveringEntry := false
	var hoveringEntryDx, hoveringEntryDy int
	for i, inp := range settings.Inputs() {
		const dx = 20

		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(dx, float64(heightOffset))
		text := fmt.Sprintf("Selecione o ponto %c:", rune(i+'A'))
		ebitentext.Draw(screen, text, textFace, op)
		w, h := ebitentext.Measure(text, textFace, 0)
		if !inp.Evaluated() {
			hasHoveringEntry = true
			hoveringEntryDx = dx + int(w) + 10
			hoveringEntryDy = heightOffset
			break
		} else {
			op := &ebitentext.DrawOptions{}
			op.ColorScale.ScaleWithColor(color.White)
			op.GeoM.Translate(dx+w+10, float64(heightOffset))
			ebitentext.Draw(screen, inp.Describe(), textFace, op)
		}
		heightOffset += int(h) + 5
	}
	if ponto := hoveringPoint; ponto != nil && hasHoveringEntry {
		j.pontoAtual = ponto
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), hoveringEntryDx, hoveringEntryDy)
	} else {
		j.pontoAtual = nil
	}

	if len(settings.Inputs()) > 0 && !hasPendingInput {
		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(20, float64(heightOffset))
		ebitentext.Draw(screen, "Pressione ESPAÇO para refazer", textFace, op)
	}
}

func (j *JogoSecundario) Layout(_, _ int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}
