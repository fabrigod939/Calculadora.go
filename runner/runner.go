package runner

import (
	"fmt"

	"github.com/fabrigod939/Calculadora.go/operação"
)

type Operação interface {
	Calcular(a, b float64) float64
}

type Runner struct {
	Operação map[string]Operação
}

func (r *Runner) Executar() {
	var a, b float64
	var operação string
	fmt.Println("Digite o primeiro número: ")
	fmt.Scanln(&a)
	fmt.Println("Digite outro número: ")
	fmt.Scanln(&b)
	fmt.Println("Escolha entre esses operadores: (+ | - | * | /): ")
	fmt.Scanln(&operação)
	op, existe := r.Operação[operação]
	if !existe {
		fmt.Println("Opção invalida")
		return
	}
	resultado := op.Calcular(a, b)
	fmt.Printf("Resultado: %.2f\n", resultado)

}

func NewRunner() *Runner {
	return &Runner{
		Operações: map[string]operação.Operação{
			"+": operação.Soma{},
			"-": operação.Subtração{},
			"*": operação.Multiplicação{},
			"/": operação.Divisão{},
		},
	}
}
