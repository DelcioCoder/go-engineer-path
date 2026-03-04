package main

import "fmt"

func main() {
	fmt.Println("Aprender a declarar variaveis em Go")

	// Forma explicita de declarar variáveis em Go
	var name string = "Dercio"
	fmt.Println("Nome: ", name)


	// Declaração com inferência de tipo 
	var idade = 22
	fmt.Println("Idade: ", idade)
}
