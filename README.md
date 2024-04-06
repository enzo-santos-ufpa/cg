# Computação Gráfica

Conteúdo da disciplina de Computação Gráfica.

A linguagem escolhida para as implementações é o [Go](https://go.dev).

## Uso

Para utilizar este projeto, é necessário ter o Go >= 1.22 instalado ([link](https://go.dev/doc/install)):

```shell
$ go version
```

```none
go version go1.22.0 windows/amd64
```

Clone este repositório:

```shell
$ git clone https://github.com/enzo-santos-ufpa/cg
$ cd cg
```

Execute os testes:

```shell
$ go test -cover ./...
```

```none
        ufpa_cg/cmd/exibicao            coverage: 0.0% of statements
ok      ufpa_cg 0.490s  coverage: 95.7% of statements
```

## Implementações

- [**Aula 2**](aula02.go) (21/03/2024): algoritmos bruto e de Bresenham para gerar uma linha entre dois pontos
- [**Aula 3**](aula03.go) (26/03/2024): algoritmo de Bresenham considerando pontos fora do primeiro octante
- [**Aula 4**](aula04.go) (02/04/2024): algoritmo de rotação em torno de um ponto pivô
- [**Aula 5**](aula05.go) (04/04/2024): operações de translação, transformação de escala e rotação utilizando matrizes
