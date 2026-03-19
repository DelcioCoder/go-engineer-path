package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estudo de Array")
	// // Forma 1: Declaração + Zero value padrão
	// fmt.Println("Declaração sem valores")
	// var numeros [5]int
	// fmt.Println(numeros)
	//
	// fmt.Println("Declaração com valores")
	// nomes := [4]string{"Ana", "Bruno", "Dercio", "Armando"}
	// fmt.Println(nomes)
	//
	// // Forma 3: Go infere o tamanho automaticamente (usando ...)
	// fmt.Println("Inferência de tamanho automático")
	// cores := [...]string{"Vermelho", "Azul", "Amarelo"}
	// fmt.Println(cores)
	// fmt.Println(len(cores))

	fmt.Println("Acessando ou modificando elementos")
	var notas [3]float64

	notas[0] = 10.5
	notas[1] = 11.5
	notas[0] = 12.5
	notas[0] = 13.5

	for _, v := range notas {
		fmt.Println(v)

	}
}
