package runner

import (
	"fmt"

	"github.com/fabrigod939/Calculadora.go/operacao"
)

type Operacao interface {
	Calcular(a, b float64) float64
}

type Runner struct {
	Operacoes map[string]Operacao
}

func (r *Runner) Executar() {
	var a, b float64
	var operacao string
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Println("Digite outro número: ")
	fmt.Scanln(&b)
	fmt.Println("Escolha entre esses operadores: (+ | - | * | /): ")
	fmt.Scanln(&operacao)
	op, existe := r.Operacoes[operacao]
	if !existe {
		fmt.Println("Opção invalida")
		return
	}
	resultado := op.Calcular(a, b)
	fmt.Printf("Resultado: %.2f\n", resultado)

}

func NewRunner() *Runner {
	return &Runner{
		Operacoes: map[string]Operacao{
			"+": operacao.Soma{},
			"-": operacao.Subtração{},
			"*": operacao.Multiplicação{},
			"/": operacao.Divisão{},
		},
	}
}
