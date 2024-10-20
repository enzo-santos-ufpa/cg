package main

import (
	"bytes"
	"context"
	_ "embed"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"image/color"
	"log"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

func repeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

type Game struct {
	TextFont *TextFont

	menu           ebiten.Game
	options        []OpcaoMenu
	selectingIndex int
}

func (g *Game) Update() error {
	menu := g.menu
	if menu == nil {
		switch {
		case repeatingKeyPressed(ebiten.KeyDown) && g.selectingIndex < len(g.options)-1:
			g.selectingIndex++
		case repeatingKeyPressed(ebiten.KeyUp) && g.selectingIndex > 0:
			g.selectingIndex--
		case repeatingKeyPressed(ebiten.KeyEnter) || repeatingKeyPressed(ebiten.KeyNumpadEnter):
			g.menu = &SubGame{
				Source: g.TextFont.Source,
				Modulo: g.options[g.selectingIndex].Create(),
			}
		}

	} else {
		if repeatingKeyPressed(ebiten.KeyEscape) {
			g.menu = nil
		} else {
			return menu.Update()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	menu := g.menu
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
	labels := make([]string, len(g.options)+1)
	labels[0] = "Menu de opções"
	for i, module := range g.options {
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
		if g.selectingIndex == i {
			op.ColorScale.ScaleWithColor(color.RGBA{R: 0x42, G: 0x87, B: 0xf5})
		} else {
			op.ColorScale.ScaleWithColor(color.White)
		}
		f := &ebitentext.GoTextFace{
			Source:    g.TextFont.Source,
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

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

type OpcaoMenu interface {
	Title() string
	Create() ModuloGame
}

type ModuloGame interface {
	Update() error
	Draw(screen *ebiten.Image, face *ebitentext.GoTextFace, dy int)
}

type opcaoVazia struct {
	Label string
}

func NewOpcaoVazia(label string) OpcaoMenu {
	return &opcaoVazia{Label: label}
}

func (o *opcaoVazia) Title() string {
	return o.Label
}

func (o *opcaoVazia) Create() ModuloGame {
	return &moduloVazio{}
}

type moduloVazio struct{}

func (m moduloVazio) Update() error {
	return nil
}

func (m moduloVazio) Draw(_ *ebiten.Image, _ *ebitentext.GoTextFace, _ int) {}

type SubGame struct {
	Source *ebitentext.GoTextFaceSource
	Modulo ModuloGame
}

func (m *SubGame) Update() error {
	return m.Modulo.Update()
}

func (m *SubGame) Draw(screen *ebiten.Image) {
	f := &ebitentext.GoTextFace{
		Source:    m.Source,
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

	m.Modulo.Draw(screen, f, int(h)+20)
}

func (m *SubGame) Layout(_, _ int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}

type TextFont struct {
	Source *ebitentext.GoTextFaceSource
}

//go:embed Arial.ttf
var arialTtf []byte

type AppData struct {
	Logger    *zap.Logger
	ErrorChan <-chan error
}

func main() {
	var data *AppData
	app := fx.New(
		fx.Provide(func(lc fx.Lifecycle) (*TextFont, error) {
			source, err := ebitentext.NewGoTextFaceSource(bytes.NewReader(arialTtf))
			if err != nil {
				return nil, err
			}
			return &TextFont{Source: source}, nil
		}),
		fx.Provide(zap.NewProduction),
		fx.Provide(func(lc fx.Lifecycle, textFont *TextFont, logger *zap.Logger) *AppData {
			game := &Game{
				TextFont: textFont,
				options: []OpcaoMenu{
					NewOpcaoDesenharLinha(),
					NewOpcaoVazia("Desenhar círculo"),
					NewOpcaoVazia("Desenhar elipse"),
					NewOpcaoVazia("Desenhar curva de Bezier (grau 2)"),
					NewOpcaoVazia("Desenhar curva de Bezier (grau 3)"),
					NewOpcaoVazia("Desenhar polilinha"),
					NewOpcaoVazia("Preencher por recursão"),
					NewOpcaoVazia("Preencher por varredura"),
					NewOpcaoVazia("Recortar linha"),
					NewOpcaoVazia("Recortar polígono"),
					NewOpcaoVazia("Transformar por rotação"),
					NewOpcaoVazia("Transformar por translação"),
					NewOpcaoVazia("Transformar por escola"),
					NewOpcaoVazia("Realizar projeção ortogonal"),
					NewOpcaoVazia("Realizar projeção oblíqua"),
					NewOpcaoVazia("Realizar projeção perspectiva"),
				},
			}

			errc := make(chan error, 1)
			lc.Append(fx.Hook{
				OnStart: func(ctx context.Context) error {
					go func() {
						defer close(errc)

						ebiten.SetWindowSize(screenWidth, screenHeight)
						ebiten.SetWindowTitle("Computação Gráfica - Trabalho 3")
						if err := ebiten.RunGame(game); err != nil {
							errc <- err
						}
					}()

					select {
					case err := <-errc:
						return err
					default:
						return nil
					}
				},
			})
			return &AppData{
				Logger:    logger,
				ErrorChan: errc,
			}
		}),
		fx.Populate(&data),
	)

	ctx := context.Background()
	if err := app.Start(ctx); err != nil {
		log.Fatal(err)
	}
	select {
	case signal := <-app.Done():
		data.Logger.Info("Received signal", zap.String("signal", signal.String()))
	case err := <-data.ErrorChan:
		data.Logger.Error("ok", zap.Error(err))
	}
	if err := app.Stop(ctx); err != nil {
		log.Fatal(err)
	}
}
