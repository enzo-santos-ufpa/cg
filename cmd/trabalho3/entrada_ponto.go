package main

import (
	"fmt"
	"ufpa_cg"
)

type entradaPonto struct {
	ponto ufpa_cg.Ponto
	ok    bool
}

func (e *entradaPonto) Selected(ponto ufpa_cg.Ponto) bool {
	return e.ok && e.ponto == ponto
}

func (e *entradaPonto) Describe() string {
	return fmt.Sprintf("(%d, %d)", e.ponto.X, e.ponto.Y)
}

func (e *entradaPonto) Consume(value ufpa_cg.Ponto) {
	e.ponto = value
	e.ok = true
}

func (e *entradaPonto) Evaluated() bool {
	return e.ok
}

func (e *entradaPonto) Reset() {
	e.ponto = ufpa_cg.Ponto{}
	e.ok = false
}
