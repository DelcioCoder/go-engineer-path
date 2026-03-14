// Semana 1 — Calculadora CLI
// Execute: go run main.go <numero1> <operador> <numero2>
// Exemplo: go run main.go 10 + 5
package main

// TODO: Implementar a calculadora aqui
import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func soma(a, b float64) float64 {
	return a + b
}

func multiplicacao(a, b float64) float64 {
	return a * b
}

func subtracao(a, b float64) float64 {
	return a - b
}

func divisao(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("divisão por zero não permitida!")
	}

	return a / b, nil
}

func main() {
	fmt.Println("Calculadora CLI")
	num1, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Erro:", err)
	}
	operacao := os.Args[2]
	num2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Erro:", err)
	}

	switch operacao {
	case "+":
		fmt.Println(soma(num1, num2))

	case "*":
		fmt.Println(multiplicacao(num1, num2))

	case "-":
		fmt.Println(subtracao(num1, num2))

	case "/":
		resultado, err := divisao(num1, num2)
		if err != nil {
			fmt.Println("Erro:", err)
		}
		fmt.Println(resultado)

	default:
		fmt.Println("Operação Inválida!!")
	}
}
