package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"ufpa_cg"

	"github.com/erikgeiser/promptkit/textinput"
)

func inputPonto(prompt string) (ufpa_cg.Ponto, error) {
	pattern := regexp.MustCompile(`^(-?\d+), (-?\d+)$`)

	input := textinput.New(prompt)
	input.Placeholder = "X, Y"
	input.Validate = func(text string) error {
		if pattern.MatchString(text) {
			return nil
		}
		return fmt.Errorf("ponto inválido; formato: `X, Y`")
	}
	text, err := input.RunPrompt()
	if err != nil {
		return ufpa_cg.Ponto{}, err
	}
	groups := pattern.FindStringSubmatch(text)
	x, err := strconv.Atoi(groups[1])
	if err != nil {
		return ufpa_cg.Ponto{}, err
	}
	y, err := strconv.Atoi(groups[2])
	if err != nil {
		return ufpa_cg.Ponto{}, err
	}
	return ufpa_cg.Ponto{X: x, Y: y}, nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	p1, err := inputPonto("Ponto inicial: ")
	if err != nil {
		return err
	}
	p2, err := inputPonto("  Ponto final: ")
	if err != nil {
		return err
	}
	if err := ufpa_cg.Exibe(ufpa_cg.NewAlgoritmoBresenham(p1, p2), os.Stdout); err != nil {
		return err
	}
	return nil
}
