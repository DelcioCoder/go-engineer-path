package main

import "fmt"

func main() {
	// idade := 10
	// if idade >= 18 {
	// 	fmt.Println("Maior de idade")
	// } else {
	// 	fmt.Println("Menor de idade")
	// }
	nota := 18
	if nota >= 16 {
		fmt.Println("Excelente!")
	} else if nota >= 10 {
		fmt.Println("Aprovado!")
	} else {
		fmt.Println("Reprovado")
	}
}
