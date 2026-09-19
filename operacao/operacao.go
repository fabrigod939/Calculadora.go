package operacao

import "fmt"

type Soma struct{}

func (s Soma) Calcular(a, b float64) float64 {
	return a + b
}

type Subtração struct{}

func (s Subtração) Calcular(a, b float64) float64 {
	return a - b
}

type Multiplicação struct{}

func (m Multiplicação) Calcular(a, b float64) float64 {
	return a * b
}

type Divisão struct{}

func (d Divisão) Calcular(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Operação Invalida")
		return 0
	}
	return a / b
}
