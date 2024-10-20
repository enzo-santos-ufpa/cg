package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"image/color"
	"slices"
	"ufpa_cg"
)

type moduloDesenhaLinha struct {
	pontoA *ufpa_cg.Ponto
	pontoB *ufpa_cg.Ponto

	pontoAtual  *ufpa_cg.Ponto
	pontosLinha []ufpa_cg.Ponto
	cursorX     int
	cursorY     int
}

type opcaoDesenharLinha struct{}

func NewOpcaoDesenharLinha() OpcaoMenu {
	return &opcaoDesenharLinha{}
}

func (o *opcaoDesenharLinha) Title() string {
	return "Desenhar linha"
}

func (o *opcaoDesenharLinha) Create() ModuloJogo {
	return &moduloDesenhaLinha{}
}

func (m *moduloDesenhaLinha) Update() error {
	x, y := ebiten.CursorPosition()
	m.cursorX = x
	m.cursorY = y

	// Seleciona o ponto atual da tela
	if ponto := m.pontoAtual; inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) && ponto != nil {
		if pontoA := m.pontoA; pontoA == nil {
			m.pontoA = &ufpa_cg.Ponto{X: ponto.X, Y: ponto.Y}
		} else if m.pontoB == nil {
			pontoB := ufpa_cg.Ponto{X: ponto.X, Y: ponto.Y}
			m.pontoB = &pontoB

			// Calcula a reta para os dois pontos selecionados
			pontos := make([]ufpa_cg.Ponto, 0)
			algoritmo := ufpa_cg.NewAlgoritmoBresenham(*pontoA, pontoB)
			for algoritmo.Move() {
				pontos = append(pontos, algoritmo.PontoAtual())
			}
			m.pontosLinha = pontos
		}
	}
	// Reseta o estado da tela após os dois pontos tiverem sidos selecionados
	if m.pontoA != nil && m.pontoB != nil && repeatingKeyPressed(ebiten.KeySpace) {
		m.pontoA = nil
		m.pontoB = nil
		m.pontosLinha = nil
	}
	return nil
}

func (m *moduloDesenhaLinha) Draw(screen *ebiten.Image, textFace *ebitentext.GoTextFace, dy int) {
	heightOffset := dy

	const pointSize = 10       // Tamanho de um ponto, em pixels
	const pointSpacing = 3     // Espaçamento entre um ponto e outro, em pixels
	const minX, maxX = -11, 11 // Intervalo de X para o grid de pontos
	const minY, maxY = -11, 11 // Intervalo de Y para o grid de pontos

	var hoveredColor = color.RGBA{R: 0x00, G: 0x00, B: 0xFF, A: 0xFF}       // Cor para um ponto com o cursor apontado
	var selectedColor = color.RGBA{R: 0xFF, G: 0x00, B: 0x00, A: 0xFF}      // Cor para um ponto selecionado
	var filledColor = color.RGBA{R: 0xFF, G: 0x62, B: 0x00, A: 0xFF}        // Cor para um ponto na linha formada
	var defaultColor = color.RGBA{R: 0xFF, G: 0xFF, B: 0xFF, A: 0xFF}       // Cor para um ponto x != 0 && y != 0
	var defaultOriginColor = color.RGBA{R: 0x91, G: 0x91, B: 0x91, A: 0xFF} // Cor para um ponto x == 0 || y == 0

	pontoA := m.pontoA
	pontoB := m.pontoB

	var hoveringPoint *ufpa_cg.Ponto
	// Percorre o grid de pontos
	for j := minY; j <= maxY; j++ {
		for i := minX; i <= maxX; i++ {
			x := (i-minX)*(pointSize+pointSpacing) + 20
			y := heightOffset + j - minY

			ponto := ufpa_cg.Ponto{X: i, Y: -j}
			pixel := ebiten.NewImage(pointSize, pointSize)
			if (pontoA == nil || pontoB == nil) && // Se algum ponto não tiver sido selecionado
				(x <= m.cursorX && m.cursorX <= x+pointSize) && // Se a posição X do cursor estiver dentro deste ponto
				(y <= m.cursorY && m.cursorY <= y+pointSize) { // Se a posição Y do cursor estiver dentro deste ponto
				pixel.Fill(hoveredColor) // Marca este ponto como "EM SELEÇÃO"
				hoveringPoint = &ponto
			} else if (pontoA != nil && *pontoA == ponto) || // Se este ponto for o ponto A selecionado
				(pontoB != nil && *pontoB == ponto) { // Se este ponto for o ponto B selecionado
				pixel.Fill(selectedColor) // Marca este ponto como "SELECIONADO"
			} else if len(m.pontosLinha) > 0 && slices.Contains(m.pontosLinha, ponto) { // Se este ponto estiver na linha formada pelos pontos A e B
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
	for i, ponto := range []*ufpa_cg.Ponto{m.pontoA, m.pontoB} {
		const dx = 20

		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(dx, float64(heightOffset))
		text := fmt.Sprintf("Selecione o ponto %c:", rune(i+'A'))
		ebitentext.Draw(screen, text, textFace, op)
		w, h := ebitentext.Measure(text, textFace, 0)
		if ponto == nil {
			hasHoveringEntry = true
			hoveringEntryDx = dx + int(w) + 10
			hoveringEntryDy = heightOffset
			break
		} else {
			op := &ebitentext.DrawOptions{}
			op.ColorScale.ScaleWithColor(color.White)
			op.GeoM.Translate(dx+w+10, float64(heightOffset))
			ebitentext.Draw(screen, fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), textFace, op)
		}
		heightOffset += int(h) + 5
	}
	if ponto := hoveringPoint; ponto != nil && hasHoveringEntry {
		m.pontoAtual = ponto
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("(%d, %d)", ponto.X, ponto.Y), hoveringEntryDx, hoveringEntryDy)
	} else {
		m.pontoAtual = nil
	}
	if pontoA != nil && pontoB != nil {
		op := &ebitentext.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.White)
		op.GeoM.Translate(20, float64(heightOffset))
		ebitentext.Draw(screen, "Pressione ESPAÇO para refazer", textFace, op)
	}
}
