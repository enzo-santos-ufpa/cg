package main

import (
	"bytes"
	"context"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	ebitentext "github.com/hajimehoshi/ebiten/v2/text/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
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

type OpcaoMenu interface {
	Title() string
	Create() ModuloGame
}

type ModuloGame interface {
	Update() error
	Draw(screen *ebiten.Image, face *ebitentext.GoTextFace, dy int)
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
			jogo := &Jogo{
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
						if err := ebiten.RunGame(jogo); err != nil {
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
