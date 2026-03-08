package main

import "fmt"

func main() {
	//Tipos primitivos
	var itens int = 10
	var preco float64 = 5.50

	total := float64(itens) * preco
	fmt.Println(total)
}
