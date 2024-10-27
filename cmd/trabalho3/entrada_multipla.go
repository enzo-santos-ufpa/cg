package main

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"strings"
	"ufpa_cg"
)

type entradaMultipla[E EntradaModulo] struct {
	Minimo      int
	Prompt      string
	Create      func() E
	OnEvaluated func() (map[ufpa_cg.Ponto]color.Color, bool)

	entradas []E
	ok       bool
	borda    map[ufpa_cg.Ponto]color.Color
}

func (e *entradaMultipla[E]) Selected(ponto ufpa_cg.Ponto) bool {
	for _, inp := range e.entradas {
		if inp.Selected(ponto) {
			return true
		}
	}
	return false
}

func (e *entradaMultipla[E]) OnUpdate() {
	switch {
	case repeatingKeyPressed(ebiten.KeyNumpadEnter):
		count := 0
		for _, inp := range e.entradas {
			if _, evaluated := inp.Evaluated(); evaluated {
				count++
			}
		}
		if count >= e.Minimo {
			e.ok = true
		}
	default:
		for i, inp := range e.entradas {
			if _, evaluated := inp.Evaluated(); !evaluated {
				inp.OnUpdate()
				break
			} else if i == len(e.entradas)-1 {
				e.entradas = append(e.entradas, e.Create())
			}
		}
	}
}

func (e *entradaMultipla[E]) OnDraw(ponto ufpa_cg.Ponto, x, y int, size int) (color.Color, bool) {
	for _, inp := range e.entradas {
		if _, evaluated := inp.Evaluated(); !evaluated {
			if customColor, ok := inp.OnDraw(ponto, x, y, size); ok {
				return customColor, true
			}
			break
		}
	}
	return nil, false
}

func (e *entradaMultipla[E]) DescribeState() (string, bool) {
	for _, inp := range e.entradas {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return inp.DescribeState()
		}
	}
	return "", false
}

func (e *entradaMultipla[E]) DescribePrompt() string {
	if _, evaluated := e.Evaluated(); evaluated {
		return e.Prompt
	}
	buffer := strings.Builder{}
	for i, inp := range e.entradas {
		if _, evaluated := inp.Evaluated(); !evaluated {
			return fmt.Sprintf("%s %s", inp.DescribePrompt(), buffer.String())
		}
		buffer.WriteString(inp.DescribeValue())
		if i < len(e.entradas)-1 {
			buffer.WriteString(", ")
		}
	}
	return fmt.Sprintf("%s: %s", e.Prompt, buffer.String())
}

func (e *entradaMultipla[E]) DescribeActions() []AcaoEntrada {
	actions := make([]AcaoEntrada, 0)
	for _, inp := range e.entradas {
		if _, evaluated := inp.Evaluated(); !evaluated {
			actions = append(actions, inp.DescribeActions()...)
			break
		}
	}
	if len(e.entradas)-1 >= e.Minimo {
		actions = append(actions, AcaoEntrada{Titulo: "ENTERpad", Descricao: "concluir"})
	}
	return actions
}

func (e *entradaMultipla[E]) DescribeValue() string {
	buffer := strings.Builder{}
	for i, inp := range e.entradas {
		if _, evaluated := inp.Evaluated(); !evaluated {
			break
		}
		buffer.WriteString(inp.DescribeValue())
		if i < len(e.entradas)-2 {
			buffer.WriteString(", ")
		}
	}
	return buffer.String()
}

func (e *entradaMultipla[E]) Evaluated() (map[ufpa_cg.Ponto]color.Color, bool) {
	if !e.ok {
		return nil, false
	}
	return nil, e.ok
	return e.OnEvaluated()
}

func (e *entradaMultipla[E]) Reset() {
	e.entradas = []E{e.Create()}
	e.borda = nil
	e.ok = false
}
