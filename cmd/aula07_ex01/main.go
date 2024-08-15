package main

import (
	"fmt"
	"log"
	"ufpa_cg"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	const arestaCubo = 2

	vertices := make([]ufpa_cg.Ponto3D, 0)
	verticePivo := ufpa_cg.Ponto3D{X: 0, Y: 0, Z: 0}

	for kAltura := 0; kAltura <= 1; kAltura++ {
		for kLargura := 0; kLargura <= 1; kLargura++ {
			for kProfundidade := 0; kProfundidade <= 1; kProfundidade++ {
				vertices = append(vertices, ufpa_cg.Ponto3D{
					X: verticePivo.X + kAltura*arestaCubo,
					Y: verticePivo.Y + kLargura*arestaCubo,
					Z: verticePivo.Z + kProfundidade*arestaCubo,
				})
			}
		}
	}
	fmt.Println(vertices)

	var verticesModificados []ufpa_cg.Ponto3D

	// Escala
	verticesModificados = make([]ufpa_cg.Ponto3D, len(vertices))
	for i, vertice := range vertices {
		verticesModificados[i] = vertice.RedimensionaMatricial(1, 2, 0.5)
	}
	fmt.Printf("A) Escala {1x 2y 0.5z}: %v\n", verticesModificados)

	// Rotação
	verticesModificados = make([]ufpa_cg.Ponto3D, len(vertices))
	for i, vertice := range vertices {
		verticesModificados[i] = vertice.RotacionaMatricial(ufpa_cg.GrausParaRadianos(30), ufpa_cg.EixoZ)
	}
	fmt.Printf("B) Rotação {30ºz}: %v\n", verticesModificados)

	// Rotação dupla
	verticesModificados = make([]ufpa_cg.Ponto3D, len(vertices))
	for i, vertice := range vertices {
		vertice = vertice.RotacionaMatricial(ufpa_cg.GrausParaRadianos(45), ufpa_cg.EixoX)
		vertice = vertice.RotacionaMatricial(ufpa_cg.GrausParaRadianos(45), ufpa_cg.EixoY)
		verticesModificados[i] = vertice
	}
	fmt.Printf("C) Rotação {45ºx}, Rotação {45ºz}: %v\n", verticesModificados)

	return nil
}
