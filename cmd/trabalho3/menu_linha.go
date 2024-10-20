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

type desenharLinhaGame struct {
	TextFont *TextFont

	pontoA *ufpa_cg.Ponto
	pontoB *ufpa_cg.Ponto

	pontoAtual  *ufpa_cg.Ponto
	pontosLinha []ufpa_cg.Ponto
	cursorX     int
	cursorY     int
}

func (g *desenharLinhaGame) Update() error {
	x, y := ebiten.CursorPosition()
	g.cursorX = x
	g.cursorY = y

	// Seleciona o ponto atual da tela
	if ponto := g.pontoAtual; inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && ponto != nil {
		if pontoA := g.pontoA; pontoA == nil {
			g.pontoA = &ufpa_cg.Ponto{X: ponto.X, Y: ponto.Y}
		} else if g.pontoB == nil {
			pontoB := ufpa_cg.Ponto{X: ponto.X, Y: ponto.Y}
			g.pontoB = &pontoB

			// Calcula a reta para os dois pontos selecionados
			pontos := make([]ufpa_cg.Ponto, 0)
			algoritmo := ufpa_cg.NewAlgoritmoBresenham(*pontoA, pontoB)
			for algoritmo.Move() {
				pontos = append(pontos, algoritmo.PontoAtual())
			}
			g.pontosLinha = pontos
		}
	}
	// Reseta o estado da tela após os dois pontos tiverem sidos selecionados
	if g.pontoA != nil && g.pontoB != nil && repeatingKeyPressed(ebiten.KeySpace) {
		g.pontoA = nil
		g.pontoB = nil
		g.pontosLinha = nil
	}
	return nil
}

func (g *desenharLinhaGame) Draw(screen *ebiten.Image) {
	f := &ebitentext.GoTextFace{
		Source:    g.TextFont.Source,
		Direction: ebitentext.DirectionLeftToRight,
		Size:      16,
		Language:  language.BrazilianPortuguese,
	}

	const pointSize = 10       // Tamanho de um ponto, em pixels
	const pointSpacing = 3     // Espaçamento entre um ponto e outro, em pixels
	const minX, maxX = -11, 11 // Intervalo de X para o grid de pontos
	const minY, maxY = -11, 11 // Intervalo de Y para o grid de pontos

	var hoveredColor = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}       // Cor para um ponto com o cursor apontado
	var selectedColor = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}      // Cor para um ponto selecionado
	var filledColor = color.RGBA{R: 0xFF, G: 0x62, B: 0x00, A: 0xFF}        // Cor para um ponto na linha formada
	var defaultColor = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}       // Cor para um ponto x != 0 && y != 0
	var defaultOriginColor = color.RGBA{R: 0x91, G: 0x91, B: 0x91, A: 0xFF} // Cor para um ponto x == 0 || y == 0

	// Desenha texto "Pressione ESC para voltar" no topo
	op := &ebitentext.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)
	op.GeoM.Translate(10, 10)
	text := "Pressione ESC para voltar"
	ebitentext.Draw(screen, text, f, op)
	_, h := ebitentext.Measure(text, f, 0)
	heightOffset := int(h) + 20

	pontoA := g.pontoA
	pontoB := g.pontoB

	// Percorre o grid de pontos
	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			x := (i-minX)*(pointSize+pointSpacing) + 20
			y := heightOffset + j - minY

			ponto := ufpa_cg.Ponto{X: i, Y: -j}
			pixel := ebiten.NewImage(pointSize, pointSize)
			if (pontoA == nil || pontoB == nil) && // Se algum ponto não tiver sido selecionado
				(x <= g.cursorX && g.cursorX <= x+pointSize) && // Se a posição X do cursor estiver dentro deste ponto
				(y <= g.cursorY && g.cursorY <= y+pointSize) { // Se a posição Y do cursor estiver dentro deste ponto
				pixel.Fill(hoveredColor) // Marca este ponto como "EM SELEÇÃO"
				g.pontoAtual = &ponto
			} else if (pontoA != nil && *pontoA == ponto) || // Se este ponto for o ponto A selecionado
				(pontoB != nil && *pontoB == ponto) { // Se este ponto for o ponto B selecionado
				pixel.Fill(selectedColor) // Marca este ponto como "SELECIONADO"
			} else if len(g.pontosLinha) > 0 && slices.Contains(g.pontosLinha, ponto) { // Se este ponto estiver na linha formada pelos pontos A e B
				pixel.Fill(filledColor) // Marca este ponto como "NA RETA"
			} else if i == 0 || j == 0 { // Se este ponto estiver em algum dos eixos X ou Y
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
	for i, ponto := range []*ufpa_cg.Ponto{g.pontoA, g.pontoB} {
		const dx = 20

		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(dx, float64(heightOffset))
		text := fmt.Sprintf("Selecione o ponto %c:", rune(i+'A'))
		ebitentext.Draw(screen, text, f, op)
		w, h := ebitentext.Measure(text, f, 0)
		if ponto == nil {
			hasHoveringEntry = true
			hoveringEntryDx = dx + int(w) + 10
			hoveringEntryDy = heightOffset
			break
		} else {
			op := &ebitentext.DrawOptions{}
			op.ColorScale.ScaleWithColor(color.White)
			op.GeoM.Translate(dx+w+10, float64(heightOffset))
			ebitentext.Draw(screen, fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), f, op)
		}
		heightOffset += int(h) + 5
	}
	if ponto := g.pontoAtual; ponto != nil && hasHoveringEntry {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), hoveringEntryDx, hoveringEntryDy)
	}
	if pontoA != nil && pontoB != nil {
		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(20, float64(heightOffset))
		ebitentext.Draw(screen, "Pressione ESPAÇO para refazer", f, op)
	}
}

func (g *desenharLinhaGame) Layout(_, _ int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}

type telaDesenharLinha struct {
	TextFont *TextFont
}

func NewTelaDesenharLinha(textFont *TextFont) TelaOpcao {
	return &telaDesenharLinha{
		TextFont: textFont,
	}
}

func (l telaDesenharLinha) Game() ebiten.Game {
	return &desenharLinhaGame{
		TextFont: l.TextFont,
	}
}

func (l telaDesenharLinha) Titulo() string {
	return "Desenhar linha"
}
