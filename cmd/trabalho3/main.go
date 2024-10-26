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
	"image/color"
	"log"
	"ufpa_cg"
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

type EntradaModulo interface {
	Evaluated() (map[ufpa_cg.Ponto]color.Color, bool)
	Selected(ponto ufpa_cg.Ponto) bool

	// DescribePrompt informa ao usuário qual tipo de entrada ele seleciona no passo atual.
	//
	// Por exemplo, uma entrada de polígono pode retornar algo como "Selecione o conjunto de pontos:".
	//
	// Note que este método pode ser implementado para se basear no menu atual do usuário. Por exemplo, uma entrada
	// de ponto pode retornar "Selecione o ponto 1:" para o primeiro ponto no menu de desenhar uma reta, mas também
	// retornar "Selecione o centro:" no menu de desenhar um círculo.
	DescribePrompt() string

	// DescribeAction informa ao usuário quais ações ele deve executar para ir para o próximo passo.
	//
	// Por exemplo, uma entrade de polígono pode precisar que o usuário clique na grade de pontos para selecionar os
	// vértices e pressione Enter para prosseguir para a próxima entrada. Neste caso, este método pode retornar algo
	// como ("Pressione ENTER para prosseguir", true).
	//
	// Este método também pode descrever ações adicionais que o usuário pode executar para modificar a entrada atual.
	// Por exemplo, uma entrada de valor inteiro pode precisar que o usuário pressione a tecla de seta para cima no
	// teclado para incrementar o valor inteiro atual. Neste caso, este método pode retornar algo como ("Pressione a
	// seta para cima para prosseguir", true).
	//
	// Caso o segundo valor de retorno seja falso, supõe-se que a ação atual é simplesmente clicar em algum ponto na
	// grade de pontos principais. Neste caso, nenhum texto de ação será exibido.
	DescribeAction() (string, bool)

	DescribeValue() string
	Reset()
	OnDisplay() (string, bool)
	OnUpdate()
	OnDraw(ponto ufpa_cg.Ponto, x, y int, size int) (color.Color, bool)
}

type ConfiguracaoModulo interface {
	Inputs() []EntradaModulo

	Evaluate() []ufpa_cg.Ponto
}
type OpcaoMenu interface {
	Title() string
	Create() ModuloJogo
}

type ModuloJogo interface {
	Update() error
	Settings() ConfiguracaoModulo
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
			jogo := &JogoPrimario{
				TextFont: textFont,
				options: []OpcaoMenu{
					NewOpcaoDesenharLinha(),
					NewOpcaoDesenharCirculo(),
					NewOpcaoDesenharElipse(),
					NewOpcaoDesenharBezier2(),
					NewOpcaoDesenharBezier3(),
					NewOpcaoDesenharPolilinha(),
					NewOpcaoPreenchimento("Preencher por recursão", NewAlgoritmoPreenchimentoRecursao()),
					NewOpcaoPreenchimento("Preencher por varredura", NewAlgoritmoPreenchimentoVarredura()),
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
