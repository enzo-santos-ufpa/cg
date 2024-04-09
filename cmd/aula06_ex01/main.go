package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"ufpa_cg"
)

func main() {
	if err := run(os.Stdout); err != nil {
		log.Fatal(err)
	}
}

func run(stdout io.Writer) error {
	pA := ufpa_cg.Ponto{X: -3, Y: 1}
	pB := ufpa_cg.Ponto{X: 4, Y: 3}
	pC := ufpa_cg.Ponto{X: 3, Y: -3}

	if err := ufpa_cg.ExibePoligono([]ufpa_cg.AlgoritmoLinha{
		ufpa_cg.NewAlgoritmoBresenham(pA, pB),
		ufpa_cg.NewAlgoritmoBresenham(pB, pC),
		ufpa_cg.NewAlgoritmoBresenham(pA, pC),
	}, stdout); err != nil {
		return err
	}

	// Como 'pA' é o pivô, move 'pA' para origem e todos os outros pontos
	origemDx := 0 - pA.X
	origemDy := 0 - pA.Y
	pA = pA.MoveMatricial(origemDx, origemDy)
	pB = pB.MoveMatricial(origemDx, origemDy)
	pC = pC.MoveMatricial(origemDx, origemDy)

	// Rotaciona todos os pontos 45º
	const anguloOperacao1 float64 = 45
	anguloRad := anguloOperacao1 * math.Pi / 180
	pA = pA.RotacionaMatricial(anguloRad)
	pB = pB.RotacionaMatricial(anguloRad)
	pC = pC.RotacionaMatricial(anguloRad)

	// Escala todos os pontos
	const exOperacao2 float64 = 1.5
	const eyOperacao2 float64 = 0.5
	pA = pA.RedimensionaMatricial(exOperacao2, eyOperacao2)
	pB = pB.RedimensionaMatricial(exOperacao2, eyOperacao2)
	pC = pC.RedimensionaMatricial(exOperacao2, eyOperacao2)

	// Move todos os pontos
	const dxOperacao3 int = 3
	const dyOperacao3 int = 2
	pA = pA.MoveMatricial(dxOperacao3, dyOperacao3)
	pB = pB.MoveMatricial(dxOperacao3, dyOperacao3)
	pC = pC.MoveMatricial(dxOperacao3, dyOperacao3)

	// Reverte o movimento feito para 'pA' como pivô
	pA = pA.MoveMatricial(-origemDx, -origemDy)
	pB = pB.MoveMatricial(-origemDx, -origemDy)
	pC = pC.MoveMatricial(-origemDx, -origemDy)

	for i := 0; i < 2; i++ {
		if _, err := fmt.Fprintln(stdout); err != nil {
			return err
		}
	}

	if err := ufpa_cg.ExibePoligono([]ufpa_cg.AlgoritmoLinha{
		ufpa_cg.NewAlgoritmoBresenham(pA, pB),
		ufpa_cg.NewAlgoritmoBresenham(pB, pC),
		ufpa_cg.NewAlgoritmoBresenham(pA, pC),
	}, stdout); err != nil {
		return err
	}
	return nil
}
