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
ok      ufpa_cg X.XXXs                  coverage: 97.6% of statements
        ufpa_cg/cmd/exibicao            coverage: 0.0% of statements
ok      ufpa_cg/cmd/aula06_ex01 X.XXXs  coverage: 86.1% of statements
        ufpa_cg/cmd/aula07_ex01         coverage: 0.0% of statements
```

## Implementações

- [**Aula 2**](aula02.go) (21/03/2024): algoritmos bruto e de Bresenham para gerar uma linha entre dois pontos
- [**Aula 3**](aula03.go) (26/03/2024): algoritmo de Bresenham considerando pontos fora do primeiro octante
- [**Aula 4**](aula04.go) (02/04/2024): algoritmo de rotação em torno de um ponto pivô
- [**Aula 5**](aula05.go) (04/04/2024): operações de translação, transformação de escala e rotação utilizando matrizes
- [**Aula 6**](cmd/aula06_ex01/main.go) (09/04/2024): exercícios com conteúdo da aula anterior
- [**Aula 7**](aula07.go) + [exercícios](cmd/aula07_ex01/main.go) (11/04/2024): operações de translação, transformação
  de escala e usando matrizes para pontos 3D
- **Aula 8** (06/08/2024)
- **Aula 9** (08/08/2024)
- **Aula 10** (13/08/2024)
- **Aula 11** (15/08/2024)
- **Aula 12** (20/08/2024)
- **Aula 13** (22/08/2024)
- **Aula 14** (27/08/2024)
- **Aula 15** (29/08/2024)
- **Aula 16** (27/08/2024)
- **Aula 17** (29/08/2024)
- **Aula 18** (10/09/2024)
- **Aula 19** (12/09/2024)
- **Aula 20** (17/09/2024)
- [**Aula 21**](aula21.go) (19/09/2024): recorte utilizando algoritmo de Cohen-Sutherland
